// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewEventViewParams creates a new EventViewParams object
// no default values defined in spec.
func NewEventViewParams() EventViewParams {

	return EventViewParams{}
}

// EventViewParams contains all the bound params for the event view operation
// typically these are obtained from a http.Request
//
// swagger:parameters eventView
type EventViewParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентификатор события
	  Required: true
	  In: path
	*/
	EventID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewEventViewParams() beforehand.
func (o *EventViewParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rEventID, rhkEventID, _ := route.Params.GetOK("eventID")
	if err := o.bindEventID(rEventID, rhkEventID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEventID binds and validates parameter EventID from path.
func (o *EventViewParams) bindEventID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("eventID", "path", "strfmt.UUID", raw)
	}
	o.EventID = *(value.(*strfmt.UUID))

	if err := o.validateEventID(formats); err != nil {
		return err
	}

	return nil
}

// validateEventID carries on validations for parameter EventID
func (o *EventViewParams) validateEventID(formats strfmt.Registry) error {

	if err := validate.FormatOf("eventID", "path", "uuid", o.EventID.String(), formats); err != nil {
		return err
	}
	return nil
}