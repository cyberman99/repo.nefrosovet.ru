// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ClientViewOKCode is the HTTP code returned for type ClientViewOK
const ClientViewOKCode int = 200

/*ClientViewOK Коллекция клиентов

swagger:response clientViewOK
*/
type ClientViewOK struct {

	/*
	  In: Body
	*/
	Payload *ClientViewOKBody `json:"body,omitempty"`
}

// NewClientViewOK creates ClientViewOK with default headers values
func NewClientViewOK() *ClientViewOK {

	return &ClientViewOK{}
}

// WithPayload adds the payload to the client view o k response
func (o *ClientViewOK) WithPayload(payload *ClientViewOKBody) *ClientViewOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the client view o k response
func (o *ClientViewOK) SetPayload(payload *ClientViewOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClientViewOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ClientViewNotFoundCode is the HTTP code returned for type ClientViewNotFound
const ClientViewNotFoundCode int = 404

/*ClientViewNotFound Not found

swagger:response clientViewNotFound
*/
type ClientViewNotFound struct {

	/*
	  In: Body
	*/
	Payload *ClientViewNotFoundBody `json:"body,omitempty"`
}

// NewClientViewNotFound creates ClientViewNotFound with default headers values
func NewClientViewNotFound() *ClientViewNotFound {

	return &ClientViewNotFound{}
}

// WithPayload adds the payload to the client view not found response
func (o *ClientViewNotFound) WithPayload(payload *ClientViewNotFoundBody) *ClientViewNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the client view not found response
func (o *ClientViewNotFound) SetPayload(payload *ClientViewNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClientViewNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ClientViewMethodNotAllowedCode is the HTTP code returned for type ClientViewMethodNotAllowed
const ClientViewMethodNotAllowedCode int = 405

/*ClientViewMethodNotAllowed Invalid Method

swagger:response clientViewMethodNotAllowed
*/
type ClientViewMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *ClientViewMethodNotAllowedBody `json:"body,omitempty"`
}

// NewClientViewMethodNotAllowed creates ClientViewMethodNotAllowed with default headers values
func NewClientViewMethodNotAllowed() *ClientViewMethodNotAllowed {

	return &ClientViewMethodNotAllowed{}
}

// WithPayload adds the payload to the client view method not allowed response
func (o *ClientViewMethodNotAllowed) WithPayload(payload *ClientViewMethodNotAllowedBody) *ClientViewMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the client view method not allowed response
func (o *ClientViewMethodNotAllowed) SetPayload(payload *ClientViewMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClientViewMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ClientViewInternalServerErrorCode is the HTTP code returned for type ClientViewInternalServerError
const ClientViewInternalServerErrorCode int = 500

/*ClientViewInternalServerError Internal sersver error

swagger:response clientViewInternalServerError
*/
type ClientViewInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *ClientViewInternalServerErrorBody `json:"body,omitempty"`
}

// NewClientViewInternalServerError creates ClientViewInternalServerError with default headers values
func NewClientViewInternalServerError() *ClientViewInternalServerError {

	return &ClientViewInternalServerError{}
}

// WithPayload adds the payload to the client view internal server error response
func (o *ClientViewInternalServerError) WithPayload(payload *ClientViewInternalServerErrorBody) *ClientViewInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the client view internal server error response
func (o *ClientViewInternalServerError) SetPayload(payload *ClientViewInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClientViewInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
