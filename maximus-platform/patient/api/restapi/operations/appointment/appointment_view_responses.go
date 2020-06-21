// Code generated by go-swagger; DO NOT EDIT.

package appointment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AppointmentViewOKCode is the HTTP code returned for type AppointmentViewOK
const AppointmentViewOKCode int = 200

/*AppointmentViewOK Коллекция медицинских назначений

swagger:response appointmentViewOK
*/
type AppointmentViewOK struct {

	/*
	  In: Body
	*/
	Payload *AppointmentViewOKBody `json:"body,omitempty"`
}

// NewAppointmentViewOK creates AppointmentViewOK with default headers values
func NewAppointmentViewOK() *AppointmentViewOK {

	return &AppointmentViewOK{}
}

// WithPayload adds the payload to the appointment view o k response
func (o *AppointmentViewOK) WithPayload(payload *AppointmentViewOKBody) *AppointmentViewOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the appointment view o k response
func (o *AppointmentViewOK) SetPayload(payload *AppointmentViewOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppointmentViewOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppointmentViewNotFoundCode is the HTTP code returned for type AppointmentViewNotFound
const AppointmentViewNotFoundCode int = 404

/*AppointmentViewNotFound Not found

swagger:response appointmentViewNotFound
*/
type AppointmentViewNotFound struct {

	/*
	  In: Body
	*/
	Payload *AppointmentViewNotFoundBody `json:"body,omitempty"`
}

// NewAppointmentViewNotFound creates AppointmentViewNotFound with default headers values
func NewAppointmentViewNotFound() *AppointmentViewNotFound {

	return &AppointmentViewNotFound{}
}

// WithPayload adds the payload to the appointment view not found response
func (o *AppointmentViewNotFound) WithPayload(payload *AppointmentViewNotFoundBody) *AppointmentViewNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the appointment view not found response
func (o *AppointmentViewNotFound) SetPayload(payload *AppointmentViewNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppointmentViewNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppointmentViewMethodNotAllowedCode is the HTTP code returned for type AppointmentViewMethodNotAllowed
const AppointmentViewMethodNotAllowedCode int = 405

/*AppointmentViewMethodNotAllowed Invalid Method

swagger:response appointmentViewMethodNotAllowed
*/
type AppointmentViewMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *AppointmentViewMethodNotAllowedBody `json:"body,omitempty"`
}

// NewAppointmentViewMethodNotAllowed creates AppointmentViewMethodNotAllowed with default headers values
func NewAppointmentViewMethodNotAllowed() *AppointmentViewMethodNotAllowed {

	return &AppointmentViewMethodNotAllowed{}
}

// WithPayload adds the payload to the appointment view method not allowed response
func (o *AppointmentViewMethodNotAllowed) WithPayload(payload *AppointmentViewMethodNotAllowedBody) *AppointmentViewMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the appointment view method not allowed response
func (o *AppointmentViewMethodNotAllowed) SetPayload(payload *AppointmentViewMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppointmentViewMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AppointmentViewInternalServerErrorCode is the HTTP code returned for type AppointmentViewInternalServerError
const AppointmentViewInternalServerErrorCode int = 500

/*AppointmentViewInternalServerError Internal server error

swagger:response appointmentViewInternalServerError
*/
type AppointmentViewInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *AppointmentViewInternalServerErrorBody `json:"body,omitempty"`
}

// NewAppointmentViewInternalServerError creates AppointmentViewInternalServerError with default headers values
func NewAppointmentViewInternalServerError() *AppointmentViewInternalServerError {

	return &AppointmentViewInternalServerError{}
}

// WithPayload adds the payload to the appointment view internal server error response
func (o *AppointmentViewInternalServerError) WithPayload(payload *AppointmentViewInternalServerErrorBody) *AppointmentViewInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the appointment view internal server error response
func (o *AppointmentViewInternalServerError) SetPayload(payload *AppointmentViewInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AppointmentViewInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
