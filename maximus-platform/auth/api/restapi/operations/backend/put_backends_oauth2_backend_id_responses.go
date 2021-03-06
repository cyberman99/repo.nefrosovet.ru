// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutBackendsOauth2BackendIDOKCode is the HTTP code returned for type PutBackendsOauth2BackendIDOK
const PutBackendsOauth2BackendIDOKCode int = 200

/*PutBackendsOauth2BackendIDOK Объект oauth2 бэкенда

swagger:response putBackendsOauth2BackendIdOK
*/
type PutBackendsOauth2BackendIDOK struct {

	/*
	  In: Body
	*/
	Payload *PutBackendsOauth2BackendIDOKBody `json:"body,omitempty"`
}

// NewPutBackendsOauth2BackendIDOK creates PutBackendsOauth2BackendIDOK with default headers values
func NewPutBackendsOauth2BackendIDOK() *PutBackendsOauth2BackendIDOK {

	return &PutBackendsOauth2BackendIDOK{}
}

// WithPayload adds the payload to the put backends oauth2 backend Id o k response
func (o *PutBackendsOauth2BackendIDOK) WithPayload(payload *PutBackendsOauth2BackendIDOKBody) *PutBackendsOauth2BackendIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put backends oauth2 backend Id o k response
func (o *PutBackendsOauth2BackendIDOK) SetPayload(payload *PutBackendsOauth2BackendIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutBackendsOauth2BackendIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutBackendsOauth2BackendIDNotFoundCode is the HTTP code returned for type PutBackendsOauth2BackendIDNotFound
const PutBackendsOauth2BackendIDNotFoundCode int = 404

/*PutBackendsOauth2BackendIDNotFound Not found

swagger:response putBackendsOauth2BackendIdNotFound
*/
type PutBackendsOauth2BackendIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *PutBackendsOauth2BackendIDNotFoundBody `json:"body,omitempty"`
}

// NewPutBackendsOauth2BackendIDNotFound creates PutBackendsOauth2BackendIDNotFound with default headers values
func NewPutBackendsOauth2BackendIDNotFound() *PutBackendsOauth2BackendIDNotFound {

	return &PutBackendsOauth2BackendIDNotFound{}
}

// WithPayload adds the payload to the put backends oauth2 backend Id not found response
func (o *PutBackendsOauth2BackendIDNotFound) WithPayload(payload *PutBackendsOauth2BackendIDNotFoundBody) *PutBackendsOauth2BackendIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put backends oauth2 backend Id not found response
func (o *PutBackendsOauth2BackendIDNotFound) SetPayload(payload *PutBackendsOauth2BackendIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutBackendsOauth2BackendIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutBackendsOauth2BackendIDMethodNotAllowedCode is the HTTP code returned for type PutBackendsOauth2BackendIDMethodNotAllowed
const PutBackendsOauth2BackendIDMethodNotAllowedCode int = 405

/*PutBackendsOauth2BackendIDMethodNotAllowed Invalid Method

swagger:response putBackendsOauth2BackendIdMethodNotAllowed
*/
type PutBackendsOauth2BackendIDMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *PutBackendsOauth2BackendIDMethodNotAllowedBody `json:"body,omitempty"`
}

// NewPutBackendsOauth2BackendIDMethodNotAllowed creates PutBackendsOauth2BackendIDMethodNotAllowed with default headers values
func NewPutBackendsOauth2BackendIDMethodNotAllowed() *PutBackendsOauth2BackendIDMethodNotAllowed {

	return &PutBackendsOauth2BackendIDMethodNotAllowed{}
}

// WithPayload adds the payload to the put backends oauth2 backend Id method not allowed response
func (o *PutBackendsOauth2BackendIDMethodNotAllowed) WithPayload(payload *PutBackendsOauth2BackendIDMethodNotAllowedBody) *PutBackendsOauth2BackendIDMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put backends oauth2 backend Id method not allowed response
func (o *PutBackendsOauth2BackendIDMethodNotAllowed) SetPayload(payload *PutBackendsOauth2BackendIDMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutBackendsOauth2BackendIDMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutBackendsOauth2BackendIDInternalServerErrorCode is the HTTP code returned for type PutBackendsOauth2BackendIDInternalServerError
const PutBackendsOauth2BackendIDInternalServerErrorCode int = 500

/*PutBackendsOauth2BackendIDInternalServerError Internal server error

swagger:response putBackendsOauth2BackendIdInternalServerError
*/
type PutBackendsOauth2BackendIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PutBackendsOauth2BackendIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewPutBackendsOauth2BackendIDInternalServerError creates PutBackendsOauth2BackendIDInternalServerError with default headers values
func NewPutBackendsOauth2BackendIDInternalServerError() *PutBackendsOauth2BackendIDInternalServerError {

	return &PutBackendsOauth2BackendIDInternalServerError{}
}

// WithPayload adds the payload to the put backends oauth2 backend Id internal server error response
func (o *PutBackendsOauth2BackendIDInternalServerError) WithPayload(payload *PutBackendsOauth2BackendIDInternalServerErrorBody) *PutBackendsOauth2BackendIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put backends oauth2 backend Id internal server error response
func (o *PutBackendsOauth2BackendIDInternalServerError) SetPayload(payload *PutBackendsOauth2BackendIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutBackendsOauth2BackendIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
