// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostChannelsEmailOKCode is the HTTP code returned for type PostChannelsEmailOK
const PostChannelsEmailOKCode int = 200

/*PostChannelsEmailOK Коллекция каналов

swagger:response postChannelsEmailOK
*/
type PostChannelsEmailOK struct {

	/*
	  In: Body
	*/
	Payload *PostChannelsEmailOKBody `json:"body,omitempty"`
}

// NewPostChannelsEmailOK creates PostChannelsEmailOK with default headers values
func NewPostChannelsEmailOK() *PostChannelsEmailOK {

	return &PostChannelsEmailOK{}
}

// WithPayload adds the payload to the post channels email o k response
func (o *PostChannelsEmailOK) WithPayload(payload *PostChannelsEmailOKBody) *PostChannelsEmailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post channels email o k response
func (o *PostChannelsEmailOK) SetPayload(payload *PostChannelsEmailOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostChannelsEmailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostChannelsEmailBadRequestCode is the HTTP code returned for type PostChannelsEmailBadRequest
const PostChannelsEmailBadRequestCode int = 400

/*PostChannelsEmailBadRequest Validation error

swagger:response postChannelsEmailBadRequest
*/
type PostChannelsEmailBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostChannelsEmailBadRequestBody `json:"body,omitempty"`
}

// NewPostChannelsEmailBadRequest creates PostChannelsEmailBadRequest with default headers values
func NewPostChannelsEmailBadRequest() *PostChannelsEmailBadRequest {

	return &PostChannelsEmailBadRequest{}
}

// WithPayload adds the payload to the post channels email bad request response
func (o *PostChannelsEmailBadRequest) WithPayload(payload *PostChannelsEmailBadRequestBody) *PostChannelsEmailBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post channels email bad request response
func (o *PostChannelsEmailBadRequest) SetPayload(payload *PostChannelsEmailBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostChannelsEmailBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostChannelsEmailForbiddenCode is the HTTP code returned for type PostChannelsEmailForbidden
const PostChannelsEmailForbiddenCode int = 403

/*PostChannelsEmailForbidden Forbidden

swagger:response postChannelsEmailForbidden
*/
type PostChannelsEmailForbidden struct {

	/*
	  In: Body
	*/
	Payload *PostChannelsEmailForbiddenBody `json:"body,omitempty"`
}

// NewPostChannelsEmailForbidden creates PostChannelsEmailForbidden with default headers values
func NewPostChannelsEmailForbidden() *PostChannelsEmailForbidden {

	return &PostChannelsEmailForbidden{}
}

// WithPayload adds the payload to the post channels email forbidden response
func (o *PostChannelsEmailForbidden) WithPayload(payload *PostChannelsEmailForbiddenBody) *PostChannelsEmailForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post channels email forbidden response
func (o *PostChannelsEmailForbidden) SetPayload(payload *PostChannelsEmailForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostChannelsEmailForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostChannelsEmailMethodNotAllowedCode is the HTTP code returned for type PostChannelsEmailMethodNotAllowed
const PostChannelsEmailMethodNotAllowedCode int = 405

/*PostChannelsEmailMethodNotAllowed Invalid Method

swagger:response postChannelsEmailMethodNotAllowed
*/
type PostChannelsEmailMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *PostChannelsEmailMethodNotAllowedBody `json:"body,omitempty"`
}

// NewPostChannelsEmailMethodNotAllowed creates PostChannelsEmailMethodNotAllowed with default headers values
func NewPostChannelsEmailMethodNotAllowed() *PostChannelsEmailMethodNotAllowed {

	return &PostChannelsEmailMethodNotAllowed{}
}

// WithPayload adds the payload to the post channels email method not allowed response
func (o *PostChannelsEmailMethodNotAllowed) WithPayload(payload *PostChannelsEmailMethodNotAllowedBody) *PostChannelsEmailMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post channels email method not allowed response
func (o *PostChannelsEmailMethodNotAllowed) SetPayload(payload *PostChannelsEmailMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostChannelsEmailMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
