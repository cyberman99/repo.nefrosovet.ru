// Code generated by go-swagger; DO NOT EDIT.

package manage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostTokensParams creates a new PostTokensParams object
// no default values defined in spec.
func NewPostTokensParams() PostTokensParams {

	return PostTokensParams{}
}

// PostTokensParams contains all the bound params for the post tokens operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostTokens
type PostTokensParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Body PostTokensBody
	/*ROOT Токен доступа
	  Required: true
	  In: query
	*/
	MasterToken string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostTokensParams() beforehand.
func (o *PostTokensParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PostTokensBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	qMasterToken, qhkMasterToken, _ := qs.GetOK("masterToken")
	if err := o.bindMasterToken(qMasterToken, qhkMasterToken, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMasterToken binds and validates parameter MasterToken from query.
func (o *PostTokensParams) bindMasterToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("masterToken", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("masterToken", "query", raw); err != nil {
		return err
	}

	o.MasterToken = raw

	return nil
}
