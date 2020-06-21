// Code generated by go-swagger; DO NOT EDIT.

package messages

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

// NewGetMessagesMessageIDParams creates a new GetMessagesMessageIDParams object
// no default values defined in spec.
func NewGetMessagesMessageIDParams() GetMessagesMessageIDParams {

	return GetMessagesMessageIDParams{}
}

// GetMessagesMessageIDParams contains all the bound params for the get messages message ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetMessagesMessageID
type GetMessagesMessageIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Токен доступа
	  Required: true
	  In: query
	*/
	AccessToken string
	/*Идентификатор сообщения
	  Required: true
	  In: path
	*/
	MessageID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetMessagesMessageIDParams() beforehand.
func (o *GetMessagesMessageIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAccessToken, qhkAccessToken, _ := qs.GetOK("accessToken")
	if err := o.bindAccessToken(qAccessToken, qhkAccessToken, route.Formats); err != nil {
		res = append(res, err)
	}

	rMessageID, rhkMessageID, _ := route.Params.GetOK("messageID")
	if err := o.bindMessageID(rMessageID, rhkMessageID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAccessToken binds and validates parameter AccessToken from query.
func (o *GetMessagesMessageIDParams) bindAccessToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("accessToken", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("accessToken", "query", raw); err != nil {
		return err
	}

	o.AccessToken = raw

	return nil
}

// bindMessageID binds and validates parameter MessageID from path.
func (o *GetMessagesMessageIDParams) bindMessageID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.MessageID = raw

	return nil
}
