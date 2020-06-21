package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/getsentry/raven-go"
	apiErrors "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"repo.nefrosovet.ru/maximus-platform/mailer/api/restapi/operations/manage"
	"repo.nefrosovet.ru/maximus-platform/mailer/sender"
)

// ServeError the error handler interface implementation
func ServeError(rw http.ResponseWriter, r *http.Request, err error) {
	rw.Header().Set("Content-Type", "application/json")
	switch e := err.(type) {
	case *apiErrors.CompositeError:
		errorsStruct := new(manage.PostTokensBadRequestBody)
		errorsStruct.Version = sender.Version

		errorsStruct.Message = PayloadValidationErrorMessage
		errorsStruct.Errors = parseComposite(e)

		res := manage.NewPostTokensBadRequest().WithPayload(errorsStruct)
		res.WriteResponse(rw, runtime.JSONProducer())

	case *apiErrors.MethodNotAllowedError:
		rw.Header().Add("Allow", strings.Join(err.(*apiErrors.MethodNotAllowedError).Allowed, ","))
		// Cause a warning about superfluous WriteHeader call
		// rw.WriteHeader(asHTTPCode(int(e.Code())))
		if r == nil || r.Method != "HEAD" {
			errorsStruct := new(manage.PostTokensMethodNotAllowedBody)
			errorsStruct.Version = sender.Version

			message := "Method " + r.Method + " not allowed"
			errorsStruct.Message = message

			res := manage.NewPostTokensMethodNotAllowed().WithPayload(errorsStruct)
			res.WriteResponse(rw, runtime.JSONProducer())
		}
	default:
		defer func() {
			if recoverValue := recover(); recoverValue != nil {
				str := fmt.Sprint(recoverValue)
				packet := raven.NewPacket(str, raven.NewException(errors.New(str), raven.GetOrNewStacktrace(recoverValue.(error), 2, 3, nil)), raven.NewHttp(r))
				raven.Capture(packet, nil)
				//w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		errorsStruct := new(manage.PostTokensForbiddenBody)
		errorsStruct.Version = sender.Version
		errorsStruct.Message = "Access denied"

		errMap := map[string]interface{}{
			"core": err.Error(),
		}
		errorsStruct.Errors = errMap

		res := manage.NewPostTokensForbidden().WithPayload(errorsStruct)
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

func parseComposite(err *apiErrors.CompositeError) interface{} {
	var res interface{}
	validation := make(map[string]interface{})
	iface := make(map[string]interface{})

	for _, subErr := range flattenComposite(err).Errors {
		switch e := subErr.(type) {
		case *apiErrors.ParseError:
			if m := errorBadTypePattern.FindStringSubmatch(e.Error()); m != nil {
				iface[m[1]] = m[2]
			} else if errorEOFPattern.MatchString(e.Error()) {
				validation["JSON"] = "EOF"
				res = validation
			} else {
				validation["ERROR"] = e.Error() + "\n"
				res = validation
			}
		case *apiErrors.Validation:
			if m := errorRequiredPattern.FindStringSubmatch(e.Error()); m != nil {
				iface[m[1]] = "required"
			} else if m := errorEnumPattern.FindStringSubmatch(e.Error()); m != nil {
				iface[m[1]] = "enum"
			} else {
				validation["ERROR"] = e.Error() + "\n"
				res = validation
			}
		default:
			validation["ERROR"] = e.Error() + "\n"
			res = validation
		}
	}

	if len(iface) != 0 {
		validation["validation"] = iface
		res = validation
	}

	return res
}
