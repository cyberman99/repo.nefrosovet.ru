package handlers

import (
	"fmt"
	errs "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/prometheus/common/log"
	"io"
	"net/http"
	"regexp"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/api/models"
	"repo.nefrosovet.ru/maximus-platform/DataRouter/report"
	"strings"
)

type DataRouter400Err struct {
	models.Error400Data
}

func (dr *DataRouter400Err) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	if err := producer.Produce(rw, dr); err != nil {
		panic(err)
	}
}

func ServeDataRouterError(rw http.ResponseWriter, req *http.Request, err error) {
	var (
		msg   string
		drErr = new(DataRouter400Err)
		data  = make([]interface{}, 0)
	)

	switch e := err.(type) {
	case *errs.MethodNotAllowedError:
		if req == nil {
			return
		}
		if req.Method == http.MethodHead {
			return
		}

		msg = fmt.Sprintf(
			"Method %v not allowed", req.Method,
		)
		rw.Header().Add("Allow", strings.Join(e.Allowed, ","))
		rw.WriteHeader(asHTTPCode(int(e.Code())))

		drErr.Version = &Version
		drErr.Data = data
		drErr.Message = &msg
		drErr.Errors = New400Errors(nil, req)
		drErr.WriteResponse(rw, runtime.JSONProducer())

	case *errs.CompositeError:
		rw.WriteHeader(400)

		drErr.Version = &Version
		drErr.Message = PayloadValidationErrorMessage
		drErr.Data = data
		drErr.Errors = New400Errors(e, req)
		drErr.WriteResponse(rw, runtime.JSONProducer())

	case errs.Error:
		rw.WriteHeader(asHTTPCode(int(e.Code())))

		drErr.Version = &Version
		drErr.Data = data
		if e.Code() == http.StatusNotFound {
			drErr.Message = NotFoundMessage
		} else {
			drErr.Message = InternalServerErrorMessage
		}
		compErr := errs.CompositeValidationError(e)
		drErr.Errors = New400Errors(compErr, req)
		drErr.WriteResponse(rw, runtime.JSONProducer())

	case nil:
		rw.WriteHeader(http.StatusInternalServerError)
		drErr.Version = &Version
		drErr.Data = data
		drErr.Message = InternalServerErrorMessage
		drErr.Errors = New400Errors(nil, req)

		drErr.WriteResponse(rw, runtime.JSONProducer())

	default:
		defer func() {
			if r := recover(); r != nil {
				report.Capture(r, req)
				err := errs.New(500, fmt.Sprint(r))
				compErr := errs.CompositeValidationError(err)
				drErr.Errors = New400Errors(compErr, req)
				drErr.WriteResponse(rw, runtime.JSONProducer())
			}
		}()

		rw.WriteHeader(http.StatusInternalServerError)
		drErr.Version = &Version
		drErr.Data = data
		drErr.Message = InternalServerErrorMessage
		drErr.Errors = New400Errors(nil, req)

		drErr.WriteResponse(rw, runtime.JSONProducer())
	}
}

const (
	format   = "format"
	required = "required"
	enum     = "oneof"
	unique   = "unique"
	min      = "min"
	array    = "array"
	object   = "object"
)

var (
	errorBadComplexTypePattern = regexp.MustCompile(
		`cannot unmarshal \w+ into Go struct field .*\.(\w+) of type \w+\.\w+`,
	)
	errorBadSequenceTypePattern = regexp.MustCompile(
		`cannot unmarshal \w+ into Go struct field \w+\.(\w+) of type \[].*`,
	)
	errorBadTypePattern = regexp.MustCompile(
		`cannot unmarshal \w+ into Go struct field \w+(?:\.(\w+))+ of type (\w+)`,
	)
	errorNoStructBadTypePattern = regexp.MustCompile(
		`cannot unmarshal \w+ into Go struct field \.(\w+) of type (\w+)`,
	)
	errorRequiredPattern        = regexp.MustCompile(`(.*) in \w+ is required`)
	errorEnumPattern            = regexp.MustCompile(`(\w+) in \w+ should be one of .*`)
	errorBodyTypePattern        = regexp.MustCompile(`parsing (\w+) .* failed, because parse error: expected (\w+) .*`)
	errorInvalidBodyTypePattern = regexp.MustCompile(`parsing body body from "" failed, because parse error: expected .*`)
	errorUUIDWithBodyString     = regexp.MustCompile(`\w+\.(\w+) in \w+ must be of type uuid: .*`)
	errorUUIDString             = regexp.MustCompile(`(\w+) in \w+ must be of type uuid: .*`)

	errorEOFString         = `failed, because unexpected EOF`
	errorDuplicationString = `shouldn't contain duplicates`
	errorShortLength       = regexp.MustCompile(`\w+\.(\w+) in \w+ should be at least .*`)
)

func New400Errors(err *errs.CompositeError, req *http.Request) *models.Error400DataAO1Errors {
	var result = new(models.Error400DataAO1Errors)
	if err == nil {
		return result
	}
	validation := make(validatorResp)

	for _, subErr := range flattenComposite(err).Errors {
		switch e := subErr.(type) {
		case *errs.ParseError:
			if errorInvalidBodyTypePattern.MatchString(e.Error()) {
				// hardcoded. Go-Swagger can't parse uuid
				if strings.Contains(req.RequestURI, "repl") {
					validation.SetFields("replyID", "string")
				}
				if strings.Contains(req.RequestURI, "route") {
					if req.Method == http.MethodPost {
						validation.SetFields("replyID", "string")
					} else if req.Method == http.MethodPut {
						validation.SetFields("replyID", "string")
						validation.SetFields("routeID", "string")
					} else {
						validation.SetFields("routeID", "string")
					}
				}
				if strings.Contains(req.RequestURI, "client") {
					validation.SetFields("ID", "string")
				}

			} else if m := errorBadComplexTypePattern.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], object)
			} else if m := errorBadSequenceTypePattern.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], array)
			} else if m := errorBadTypePattern.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], m[2])
			} else if m := errorNoStructBadTypePattern.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], m[2])
			} else if m := errorBodyTypePattern.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], m[2])
			} else if strings.Contains(e.Error(), errorEOFString) {
				result.Validation = io.EOF
			} else {
				log.Warn(" ODD ERROR:", e)
			}
		case *errs.Validation:
			if m := errorRequiredPattern.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], required)
			} else if strings.Contains(e.Error(), errorDuplicationString) {
				validation.SetFields("ID", unique)
			} else if m := errorEnumPattern.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], enum)
			} else if m := errorShortLength.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], min)
			} else if m := errorUUIDWithBodyString.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], format)
			} else if m := errorUUIDString.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], format)
			} else {
				log.Warn(" ODD ERROR:", e)
			}
		default:

			validation.SetFields("error", e.Error())
		}
	}

	result.Validation = validation
	return result
}

func flattenComposite(errors *errs.CompositeError) *errs.CompositeError {
	var res []error
	for _, er := range errors.Errors {
		switch e := er.(type) {
		case *errs.CompositeError:
			if len(e.Errors) > 0 {
				flat := flattenComposite(e)
				if len(flat.Errors) > 0 {
					res = append(res, flat.Errors...)
				}
			}
		default:
			if e != nil {
				res = append(res, e)
			}
		}
	}
	return errs.CompositeValidationError(res...)
}

func asHTTPCode(input int) int {
	if input >= 600 {
		return 422
	}
	return input
}

type validatorResp map[string]string

func (v validatorResp) SetFields(field string, rule string) {

	if field == "topic" {
		field = "dst.topic"
	}
	if field == "qos" {
		field = "dst.qos"
	}
	v[field] = rule
}
