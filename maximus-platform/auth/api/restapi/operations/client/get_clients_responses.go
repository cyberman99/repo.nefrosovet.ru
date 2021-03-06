// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetClientsOKCode is the HTTP code returned for type GetClientsOK
const GetClientsOKCode int = 200

/*GetClientsOK Объект клиента

swagger:response getClientsOK
*/
type GetClientsOK struct {

	/*
	  In: Body
	*/
	Payload *GetClientsOKBody `json:"body,omitempty"`
}

// NewGetClientsOK creates GetClientsOK with default headers values
func NewGetClientsOK() *GetClientsOK {

	return &GetClientsOK{}
}

// WithPayload adds the payload to the get clients o k response
func (o *GetClientsOK) WithPayload(payload *GetClientsOKBody) *GetClientsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get clients o k response
func (o *GetClientsOK) SetPayload(payload *GetClientsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClientsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetClientsMethodNotAllowedCode is the HTTP code returned for type GetClientsMethodNotAllowed
const GetClientsMethodNotAllowedCode int = 405

/*GetClientsMethodNotAllowed Invalid Method

swagger:response getClientsMethodNotAllowed
*/
type GetClientsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *GetClientsMethodNotAllowedBody `json:"body,omitempty"`
}

// NewGetClientsMethodNotAllowed creates GetClientsMethodNotAllowed with default headers values
func NewGetClientsMethodNotAllowed() *GetClientsMethodNotAllowed {

	return &GetClientsMethodNotAllowed{}
}

// WithPayload adds the payload to the get clients method not allowed response
func (o *GetClientsMethodNotAllowed) WithPayload(payload *GetClientsMethodNotAllowedBody) *GetClientsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get clients method not allowed response
func (o *GetClientsMethodNotAllowed) SetPayload(payload *GetClientsMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClientsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetClientsInternalServerErrorCode is the HTTP code returned for type GetClientsInternalServerError
const GetClientsInternalServerErrorCode int = 500

/*GetClientsInternalServerError Internal server error

swagger:response getClientsInternalServerError
*/
type GetClientsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetClientsInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetClientsInternalServerError creates GetClientsInternalServerError with default headers values
func NewGetClientsInternalServerError() *GetClientsInternalServerError {

	return &GetClientsInternalServerError{}
}

// WithPayload adds the payload to the get clients internal server error response
func (o *GetClientsInternalServerError) WithPayload(payload *GetClientsInternalServerErrorBody) *GetClientsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get clients internal server error response
func (o *GetClientsInternalServerError) SetPayload(payload *GetClientsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClientsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
