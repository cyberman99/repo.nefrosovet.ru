// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostBackendsBackendIDGroupsOKCode is the HTTP code returned for type PostBackendsBackendIDGroupsOK
const PostBackendsBackendIDGroupsOKCode int = 200

/*PostBackendsBackendIDGroupsOK Объект групп бэкенда

swagger:response postBackendsBackendIdGroupsOK
*/
type PostBackendsBackendIDGroupsOK struct {

	/*
	  In: Body
	*/
	Payload *PostBackendsBackendIDGroupsOKBody `json:"body,omitempty"`
}

// NewPostBackendsBackendIDGroupsOK creates PostBackendsBackendIDGroupsOK with default headers values
func NewPostBackendsBackendIDGroupsOK() *PostBackendsBackendIDGroupsOK {

	return &PostBackendsBackendIDGroupsOK{}
}

// WithPayload adds the payload to the post backends backend Id groups o k response
func (o *PostBackendsBackendIDGroupsOK) WithPayload(payload *PostBackendsBackendIDGroupsOKBody) *PostBackendsBackendIDGroupsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post backends backend Id groups o k response
func (o *PostBackendsBackendIDGroupsOK) SetPayload(payload *PostBackendsBackendIDGroupsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBackendsBackendIDGroupsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostBackendsBackendIDGroupsNotFoundCode is the HTTP code returned for type PostBackendsBackendIDGroupsNotFound
const PostBackendsBackendIDGroupsNotFoundCode int = 404

/*PostBackendsBackendIDGroupsNotFound Not found

swagger:response postBackendsBackendIdGroupsNotFound
*/
type PostBackendsBackendIDGroupsNotFound struct {

	/*
	  In: Body
	*/
	Payload *PostBackendsBackendIDGroupsNotFoundBody `json:"body,omitempty"`
}

// NewPostBackendsBackendIDGroupsNotFound creates PostBackendsBackendIDGroupsNotFound with default headers values
func NewPostBackendsBackendIDGroupsNotFound() *PostBackendsBackendIDGroupsNotFound {

	return &PostBackendsBackendIDGroupsNotFound{}
}

// WithPayload adds the payload to the post backends backend Id groups not found response
func (o *PostBackendsBackendIDGroupsNotFound) WithPayload(payload *PostBackendsBackendIDGroupsNotFoundBody) *PostBackendsBackendIDGroupsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post backends backend Id groups not found response
func (o *PostBackendsBackendIDGroupsNotFound) SetPayload(payload *PostBackendsBackendIDGroupsNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBackendsBackendIDGroupsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostBackendsBackendIDGroupsMethodNotAllowedCode is the HTTP code returned for type PostBackendsBackendIDGroupsMethodNotAllowed
const PostBackendsBackendIDGroupsMethodNotAllowedCode int = 405

/*PostBackendsBackendIDGroupsMethodNotAllowed Invalid Method

swagger:response postBackendsBackendIdGroupsMethodNotAllowed
*/
type PostBackendsBackendIDGroupsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *PostBackendsBackendIDGroupsMethodNotAllowedBody `json:"body,omitempty"`
}

// NewPostBackendsBackendIDGroupsMethodNotAllowed creates PostBackendsBackendIDGroupsMethodNotAllowed with default headers values
func NewPostBackendsBackendIDGroupsMethodNotAllowed() *PostBackendsBackendIDGroupsMethodNotAllowed {

	return &PostBackendsBackendIDGroupsMethodNotAllowed{}
}

// WithPayload adds the payload to the post backends backend Id groups method not allowed response
func (o *PostBackendsBackendIDGroupsMethodNotAllowed) WithPayload(payload *PostBackendsBackendIDGroupsMethodNotAllowedBody) *PostBackendsBackendIDGroupsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post backends backend Id groups method not allowed response
func (o *PostBackendsBackendIDGroupsMethodNotAllowed) SetPayload(payload *PostBackendsBackendIDGroupsMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBackendsBackendIDGroupsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostBackendsBackendIDGroupsInternalServerErrorCode is the HTTP code returned for type PostBackendsBackendIDGroupsInternalServerError
const PostBackendsBackendIDGroupsInternalServerErrorCode int = 500

/*PostBackendsBackendIDGroupsInternalServerError Internal server error

swagger:response postBackendsBackendIdGroupsInternalServerError
*/
type PostBackendsBackendIDGroupsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PostBackendsBackendIDGroupsInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostBackendsBackendIDGroupsInternalServerError creates PostBackendsBackendIDGroupsInternalServerError with default headers values
func NewPostBackendsBackendIDGroupsInternalServerError() *PostBackendsBackendIDGroupsInternalServerError {

	return &PostBackendsBackendIDGroupsInternalServerError{}
}

// WithPayload adds the payload to the post backends backend Id groups internal server error response
func (o *PostBackendsBackendIDGroupsInternalServerError) WithPayload(payload *PostBackendsBackendIDGroupsInternalServerErrorBody) *PostBackendsBackendIDGroupsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post backends backend Id groups internal server error response
func (o *PostBackendsBackendIDGroupsInternalServerError) SetPayload(payload *PostBackendsBackendIDGroupsInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostBackendsBackendIDGroupsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
