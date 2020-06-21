// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOauth2BackendIDParams creates a new GetOauth2BackendIDParams object
// no default values defined in spec.
func NewGetOauth2BackendIDParams() GetOauth2BackendIDParams {

	return GetOauth2BackendIDParams{}
}

// GetOauth2BackendIDParams contains all the bound params for the get oauth2 backend ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetOauth2BackendID
type GetOauth2BackendIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентификатор бэкенда
	  Required: true
	  In: path
	*/
	BackendID string
	/*URI, на который будет перенаправлен клиент после аутентификации
	  Required: true
	  In: query
	*/
	RedirectURI string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetOauth2BackendIDParams() beforehand.
func (o *GetOauth2BackendIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rBackendID, rhkBackendID, _ := route.Params.GetOK("backendID")
	if err := o.bindBackendID(rBackendID, rhkBackendID, route.Formats); err != nil {
		res = append(res, err)
	}

	qRedirectURI, qhkRedirectURI, _ := qs.GetOK("redirectURI")
	if err := o.bindRedirectURI(qRedirectURI, qhkRedirectURI, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBackendID binds and validates parameter BackendID from path.
func (o *GetOauth2BackendIDParams) bindBackendID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.BackendID = raw

	return nil
}

// bindRedirectURI binds and validates parameter RedirectURI from query.
func (o *GetOauth2BackendIDParams) bindRedirectURI(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("redirectURI", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("redirectURI", "query", raw); err != nil {
		return err
	}

	o.RedirectURI = raw

	return nil
}
