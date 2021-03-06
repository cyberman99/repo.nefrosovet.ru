// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostRefreshOKCode is the HTTP code returned for type PostRefreshOK
const PostRefreshOKCode int = 200

/*PostRefreshOK Результат аутентификации

swagger:response postRefreshOK
*/
type PostRefreshOK struct {

	/*
	  In: Body
	*/
	Payload *PostRefreshOKBody `json:"body,omitempty"`
}

// NewPostRefreshOK creates PostRefreshOK with default headers values
func NewPostRefreshOK() *PostRefreshOK {

	return &PostRefreshOK{}
}

// WithPayload adds the payload to the post refresh o k response
func (o *PostRefreshOK) WithPayload(payload *PostRefreshOKBody) *PostRefreshOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post refresh o k response
func (o *PostRefreshOK) SetPayload(payload *PostRefreshOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRefreshOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostRefreshUnauthorizedCode is the HTTP code returned for type PostRefreshUnauthorized
const PostRefreshUnauthorizedCode int = 401

/*PostRefreshUnauthorized Access denied

swagger:response postRefreshUnauthorized
*/
type PostRefreshUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *PostRefreshUnauthorizedBody `json:"body,omitempty"`
}

// NewPostRefreshUnauthorized creates PostRefreshUnauthorized with default headers values
func NewPostRefreshUnauthorized() *PostRefreshUnauthorized {

	return &PostRefreshUnauthorized{}
}

// WithPayload adds the payload to the post refresh unauthorized response
func (o *PostRefreshUnauthorized) WithPayload(payload *PostRefreshUnauthorizedBody) *PostRefreshUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post refresh unauthorized response
func (o *PostRefreshUnauthorized) SetPayload(payload *PostRefreshUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRefreshUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostRefreshMethodNotAllowedCode is the HTTP code returned for type PostRefreshMethodNotAllowed
const PostRefreshMethodNotAllowedCode int = 405

/*PostRefreshMethodNotAllowed Invalid Method

swagger:response postRefreshMethodNotAllowed
*/
type PostRefreshMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *PostRefreshMethodNotAllowedBody `json:"body,omitempty"`
}

// NewPostRefreshMethodNotAllowed creates PostRefreshMethodNotAllowed with default headers values
func NewPostRefreshMethodNotAllowed() *PostRefreshMethodNotAllowed {

	return &PostRefreshMethodNotAllowed{}
}

// WithPayload adds the payload to the post refresh method not allowed response
func (o *PostRefreshMethodNotAllowed) WithPayload(payload *PostRefreshMethodNotAllowedBody) *PostRefreshMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post refresh method not allowed response
func (o *PostRefreshMethodNotAllowed) SetPayload(payload *PostRefreshMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRefreshMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostRefreshInternalServerErrorCode is the HTTP code returned for type PostRefreshInternalServerError
const PostRefreshInternalServerErrorCode int = 500

/*PostRefreshInternalServerError Internal server error

swagger:response postRefreshInternalServerError
*/
type PostRefreshInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PostRefreshInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostRefreshInternalServerError creates PostRefreshInternalServerError with default headers values
func NewPostRefreshInternalServerError() *PostRefreshInternalServerError {

	return &PostRefreshInternalServerError{}
}

// WithPayload adds the payload to the post refresh internal server error response
func (o *PostRefreshInternalServerError) WithPayload(payload *PostRefreshInternalServerErrorBody) *PostRefreshInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post refresh internal server error response
func (o *PostRefreshInternalServerError) SetPayload(payload *PostRefreshInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRefreshInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
