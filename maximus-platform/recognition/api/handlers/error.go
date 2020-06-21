package handlers

import (
	"fmt"
	errs "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"io"
	"net/http"
	"regexp"
	"repo.nefrosovet.ru/maximus-platform/recognition/api/models"
	"repo.nefrosovet.ru/maximus-platform/recognition/monitor"
	"strings"
)

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
		`cannot unmarshal \w+ into Go struct field \w+\.(\w+) of type (\w+)`,
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

type ErrorHandler struct {
	models.Error400Data

	version string
}

func NewErrorHandler(version string) *ErrorHandler {
	return &ErrorHandler{
		Error400Data: models.Error400Data{},
		version:      version,
	}
}

func (eh *ErrorHandler) ServeError(rw http.ResponseWriter, req *http.Request, err error) {
	var (
		msg  string
		errH = new(ErrorHandler)
		data = make([]interface{}, 0)
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

		errH.Version = &eh.version
		errH.Data = data
		errH.Message = &msg
		errH.Errors = New400Errors(nil, req)
		errH.WriteResponse(rw, runtime.JSONProducer())

	case *errs.CompositeError:
		rw.WriteHeader(400)

		errH.Version = &eh.version
		errH.Message = &PayloadValidationErrorMessage
		errH.Data = data
		errH.Errors = New400Errors(e, req)
		errH.WriteResponse(rw, runtime.JSONProducer())

	case errs.Error:
		rw.WriteHeader(asHTTPCode(int(e.Code())))

		errH.Version = &eh.version
		errH.Data = data
		if e.Code() == http.StatusNotFound {
			errH.Message = &NotFoundMessage
		} else {
			errH.Message = &InternalServerErrorMessage
		}
		compErr := errs.CompositeValidationError(e)
		errH.Errors = New400Errors(compErr, req)
		errH.WriteResponse(rw, runtime.JSONProducer())

	case nil:
		rw.WriteHeader(http.StatusInternalServerError)
		errH.Version = &eh.version
		errH.Data = data
		errH.Message = &InternalServerErrorMessage
		errH.Errors = New400Errors(nil, req)

		errH.WriteResponse(rw, runtime.JSONProducer())

	default:
		defer func() {
			if r := recover(); r != nil {
				monitor.Capture(r, req)
				err := errs.New(500, fmt.Sprint(r))
				compErr := errs.CompositeValidationError(err)
				errH.Errors = New400Errors(compErr, req)
				errH.WriteResponse(rw, runtime.JSONProducer())
			}
		}()

		rw.WriteHeader(http.StatusInternalServerError)
		errH.Version = &eh.version
		errH.Data = data
		errH.Message = &InternalServerErrorMessage
		errH.Errors = New400Errors(nil, req)

		errH.WriteResponse(rw, runtime.JSONProducer())
	}
}

// interface protocol method
func (eh *ErrorHandler) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	if err := producer.Produce(rw, eh); err != nil {
		panic(err)
	}
}

type validatorResp map[string]string

func (v validatorResp) SetFields(field string, rule string) {
	v[field] = rule
}

func New400Errors(err *errs.CompositeError, req *http.Request) *models.Error400DataAO1Errors {
	var result = new(models.Error400DataAO1Errors)
	if err == nil {
		return result
	}
	validation := make(validatorResp)

	for _, subErr := range flattenComposite(err).Errors {
		isProcessed := true
		switch e := subErr.(type) {
		case *errs.ParseError:
			if m := errorBadComplexTypePattern.FindStringSubmatch(e.Error()); m != nil {
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
				isProcessed = false
			}
		case *errs.Validation:
			if m := errorRequiredPattern.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], required)
			} else if strings.Contains(e.Error(), errorDuplicationString) {
				validation.SetFields("PhotoID", unique)
			} else if m := errorEnumPattern.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], enum)
			} else if m := errorShortLength.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], min)
			} else if m := errorUUIDWithBodyString.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], format)
			} else if m := errorUUIDString.FindStringSubmatch(e.Error()); m != nil {
				validation.SetFields(m[1], format)
			} else {
				isProcessed = false
			}
		default:
			validation.SetFields("ERROR: ", e.Error())
		}

		if !isProcessed {
			validation.SetFields("UNPROCESSED ERROR: ", subErr.Error())
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
