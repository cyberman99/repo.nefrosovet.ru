package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/getsentry/raven-go"
	apiErrors "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"repo.nefrosovet.ru/maximus-platform/auth/api/models"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/auth"
	"repo.nefrosovet.ru/maximus-platform/auth/api/restapi/operations/token"
)

// ServeError the error handler interface implementation
func ServeError(rw http.ResponseWriter, r *http.Request, err error) {
	var data []string
	var dataErrors []string

	rw.Header().Set("Content-Type", "application/json")
	switch e := err.(type) {
	case *apiErrors.CompositeError:
		errorsStruct := new(auth.PostUserBadRequestBody)
		errorsStruct.Version = &Version
		errorsStruct.Data = nil
		errorsStruct.Message = &PayloadValidationErrorMessage
		errorsStruct.Data = data
		errorsStruct.Errors = parseComposite(e)

		res := auth.NewPostUserBadRequest().WithPayload(errorsStruct)
		res.WriteResponse(rw, runtime.JSONProducer())
	case *apiErrors.MethodNotAllowedError:
		rw.Header().Add("Allow", strings.Join(err.(*apiErrors.MethodNotAllowedError).Allowed, ","))
		//rw.WriteHeader(asHTTPCode(int(e.Code())))
		if r == nil || r.Method != "HEAD" {
			errorsStruct := new(auth.PostUserMethodNotAllowedBody)
			errorsStruct.Version = &Version
			errorsStruct.Data = nil
			errorsStruct.Message = "Method " + r.Method + " not allowed"

			errorsStruct.Errors = dataErrors
			errorsStruct.Data = data

			res := auth.NewPostUserMethodNotAllowed().WithPayload(errorsStruct)
			res.WriteResponse(rw, runtime.JSONProducer())
		}
	case *jwt.ValidationError:
		errorsStruct := new(token.GetWhoamiUnauthorizedBody)
		errorsStruct.Version = &Version
		errorsStruct.Data = nil
		errorsStruct.Message = &AccessDeniedMessage
		errorsStruct.Data = data
		errorsStruct.Errors = append(errorsStruct.Errors, e.Error())

		res := token.NewGetWhoamiUnauthorized().WithPayload(errorsStruct)
		res.WriteResponse(rw, runtime.JSONProducer())
	default:
		defer func() {
			if recoverValue := recover(); recoverValue != nil {
				str := fmt.Sprint(recoverValue)
				packet := raven.NewPacket(str, raven.NewException(errors.New(str), raven.GetOrNewStacktrace(recoverValue.(error), 2, 3, nil)), raven.NewHttp(r))
				raven.Capture(packet, nil)
				// w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		errorsStruct := new(auth.PostUserInternalServerErrorBody)
		errorsStruct.Version = &Version
		errorsStruct.Data = data
		errorsStruct.Errors = dataErrors
		errorsStruct.Message = &InternalServerErrorMessage

		errMap := map[string]interface{}{
			"core": err.Error(),
		}
		errorsStruct.Errors = errMap

		res := auth.NewPostUserInternalServerError().WithPayload(errorsStruct)
		res.WriteResponse(rw, runtime.JSONProducer())
	}
}

func asHTTPCode(input int) int {
	if input >= 600 {
		return 422
	}
	return input
}

func flattenComposite(errs *apiErrors.CompositeError) *apiErrors.CompositeError {
	var res []error
	for _, er := range errs.Errors {
		switch e := er.(type) {
		case *apiErrors.CompositeError:
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
	return apiErrors.CompositeValidationError(res...)
}

var errorBadTypePattern = regexp.MustCompile(`cannot unmarshal \w+ into Go struct field \w+\.(\w+) of type (\w+)`)
var errorEOFPattern = regexp.MustCompile(`parsing .* failed, because unexpected EOF`)
var errorRequiredPattern = regexp.MustCompile(`(\w+) in \w+ is required`)
var errorEnumPattern = regexp.MustCompile(`(\w+) in \w+ should be one of .*`)
var errorMinLengthPattern = regexp.MustCompile(`(\w+) in \w+ should be at least \d+ chars long`)

func parseComposite(err *apiErrors.CompositeError) *auth.PostUserBadRequestBodyAO1Errors {
	res := new(auth.PostUserBadRequestBodyAO1Errors)
	validation := make(map[string]interface{})
	var core string

	for _, subErr := range flattenComposite(err).Errors {

		switch e := subErr.(type) {
		case *apiErrors.ParseError:
			if m := errorBadTypePattern.FindStringSubmatch(e.Error()); m != nil {
				validation[m[1]] = m[2]
			} else if errorEOFPattern.MatchString(e.Error()) {
				res.JSON = "EOF"
			} else {
				core = core + e.Error() + "\n"
			}
		case *apiErrors.Validation:
			if m := errorRequiredPattern.FindStringSubmatch(e.Error()); m != nil {
				validation[m[1]] = "required"
			} else if m := errorEnumPattern.FindStringSubmatch(e.Error()); m != nil {
				validation[m[1]] = "enum"
			} else if m := errorMinLengthPattern.FindStringSubmatch(e.Error()); m != nil {
				validation[m[1]] = "min"
			} else {
				core = core + e.Error() + "\n"
			}
		default:
			core = core + e.Error() + "\n"
		}
	}

	if len(validation) != 0 {
		res.Validation = validation
	}
	res.Core = core

	return res
}

// RavenHandler is http.Handler with raven control
type RavenHandler struct {
	o http.Handler
}

// NewRavenHandler returns RavenHandler
func NewRavenHandler(h http.Handler) *RavenHandler {
	return &RavenHandler{
		o: h,
	}
}

// ServeHTTP realizes http.Handler interface
func (rh *RavenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			err, ok := rec.(error)
			if !ok {
				err = fmt.Errorf("%s", rec)
			}

			packet := raven.NewPacket(err.Error(),
				raven.NewException(err,
					raven.GetOrNewStacktrace(err, 2, 3, nil),
				),
				raven.NewHttp(r),
			)

			raven.Capture(packet, nil)

			// Return response
			response := new(models.Error500Data)
			response.Version = &Version
			response.Data = []string{}
			response.Errors = map[string]interface{}{
				"core": err.Error(),
			}
			response.Message = &InternalServerErrorMessage

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)

			_ = runtime.JSONProducer().Produce(w, &response)
		}
	}()

	rh.o.ServeHTTP(w, r)
}
