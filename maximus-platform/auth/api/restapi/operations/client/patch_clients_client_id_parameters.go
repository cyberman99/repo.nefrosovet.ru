// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPatchClientsClientIDParams creates a new PatchClientsClientIDParams object
// no default values defined in spec.
func NewPatchClientsClientIDParams() PatchClientsClientIDParams {

	return PatchClientsClientIDParams{}
}

// PatchClientsClientIDParams contains all the bound params for the patch clients client ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters PatchClientsClientID
type PatchClientsClientIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Body PatchClientsClientIDBody
	/*Идентификатор клиента
	  Required: true
	  In: path
	*/
	ClientID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPatchClientsClientIDParams() beforehand.
func (o *PatchClientsClientIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PatchClientsClientIDBody
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
	rClientID, rhkClientID, _ := route.Params.GetOK("clientID")
	if err := o.bindClientID(rClientID, rhkClientID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClientID binds and validates parameter ClientID from path.
func (o *PatchClientsClientIDParams) bindClientID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ClientID = raw

	return nil
}