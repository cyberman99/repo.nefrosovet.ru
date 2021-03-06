// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetClientsClientIDOKCode is the HTTP code returned for type GetClientsClientIDOK
const GetClientsClientIDOKCode int = 200

/*GetClientsClientIDOK Объект клиента

swagger:response getClientsClientIdOK
*/
type GetClientsClientIDOK struct {

	/*
	  In: Body
	*/
	Payload *GetClientsClientIDOKBody `json:"body,omitempty"`
}

// NewGetClientsClientIDOK creates GetClientsClientIDOK with default headers values
func NewGetClientsClientIDOK() *GetClientsClientIDOK {

	return &GetClientsClientIDOK{}
}

// WithPayload adds the payload to the get clients client Id o k response
func (o *GetClientsClientIDOK) WithPayload(payload *GetClientsClientIDOKBody) *GetClientsClientIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get clients client Id o k response
func (o *GetClientsClientIDOK) SetPayload(payload *GetClientsClientIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClientsClientIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetClientsClientIDNotFoundCode is the HTTP code returned for type GetClientsClientIDNotFound
const GetClientsClientIDNotFoundCode int = 404

/*GetClientsClientIDNotFound Not found

swagger:response getClientsClientIdNotFound
*/
type GetClientsClientIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *GetClientsClientIDNotFoundBody `json:"body,omitempty"`
}

// NewGetClientsClientIDNotFound creates GetClientsClientIDNotFound with default headers values
func NewGetClientsClientIDNotFound() *GetClientsClientIDNotFound {

	return &GetClientsClientIDNotFound{}
}

// WithPayload adds the payload to the get clients client Id not found response
func (o *GetClientsClientIDNotFound) WithPayload(payload *GetClientsClientIDNotFoundBody) *GetClientsClientIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get clients client Id not found response
func (o *GetClientsClientIDNotFound) SetPayload(payload *GetClientsClientIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClientsClientIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetClientsClientIDMethodNotAllowedCode is the HTTP code returned for type GetClientsClientIDMethodNotAllowed
const GetClientsClientIDMethodNotAllowedCode int = 405

/*GetClientsClientIDMethodNotAllowed Invalid Method

swagger:response getClientsClientIdMethodNotAllowed
*/
type GetClientsClientIDMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *GetClientsClientIDMethodNotAllowedBody `json:"body,omitempty"`
}

// NewGetClientsClientIDMethodNotAllowed creates GetClientsClientIDMethodNotAllowed with default headers values
func NewGetClientsClientIDMethodNotAllowed() *GetClientsClientIDMethodNotAllowed {

	return &GetClientsClientIDMethodNotAllowed{}
}

// WithPayload adds the payload to the get clients client Id method not allowed response
func (o *GetClientsClientIDMethodNotAllowed) WithPayload(payload *GetClientsClientIDMethodNotAllowedBody) *GetClientsClientIDMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get clients client Id method not allowed response
func (o *GetClientsClientIDMethodNotAllowed) SetPayload(payload *GetClientsClientIDMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClientsClientIDMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetClientsClientIDInternalServerErrorCode is the HTTP code returned for type GetClientsClientIDInternalServerError
const GetClientsClientIDInternalServerErrorCode int = 500

/*GetClientsClientIDInternalServerError Internal server error

swagger:response getClientsClientIdInternalServerError
*/
type GetClientsClientIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetClientsClientIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetClientsClientIDInternalServerError creates GetClientsClientIDInternalServerError with default headers values
func NewGetClientsClientIDInternalServerError() *GetClientsClientIDInternalServerError {

	return &GetClientsClientIDInternalServerError{}
}

// WithPayload adds the payload to the get clients client Id internal server error response
func (o *GetClientsClientIDInternalServerError) WithPayload(payload *GetClientsClientIDInternalServerErrorBody) *GetClientsClientIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get clients client Id internal server error response
func (o *GetClientsClientIDInternalServerError) SetPayload(payload *GetClientsClientIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClientsClientIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
