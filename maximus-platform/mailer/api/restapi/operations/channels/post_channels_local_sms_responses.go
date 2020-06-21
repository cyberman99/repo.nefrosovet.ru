// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostChannelsLocalSmsOKCode is the HTTP code returned for type PostChannelsLocalSmsOK
const PostChannelsLocalSmsOKCode int = 200

/*PostChannelsLocalSmsOK Коллекция каналов

swagger:response postChannelsLocalSmsOK
*/
type PostChannelsLocalSmsOK struct {

	/*
	  In: Body
	*/
	Payload *PostChannelsLocalSmsOKBody `json:"body,omitempty"`
}

// NewPostChannelsLocalSmsOK creates PostChannelsLocalSmsOK with default headers values
func NewPostChannelsLocalSmsOK() *PostChannelsLocalSmsOK {

	return &PostChannelsLocalSmsOK{}
}

// WithPayload adds the payload to the post channels local sms o k response
func (o *PostChannelsLocalSmsOK) WithPayload(payload *PostChannelsLocalSmsOKBody) *PostChannelsLocalSmsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post channels local sms o k response
func (o *PostChannelsLocalSmsOK) SetPayload(payload *PostChannelsLocalSmsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostChannelsLocalSmsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostChannelsLocalSmsBadRequestCode is the HTTP code returned for type PostChannelsLocalSmsBadRequest
const PostChannelsLocalSmsBadRequestCode int = 400

/*PostChannelsLocalSmsBadRequest Validation error

swagger:response postChannelsLocalSmsBadRequest
*/
type PostChannelsLocalSmsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostChannelsLocalSmsBadRequestBody `json:"body,omitempty"`
}

// NewPostChannelsLocalSmsBadRequest creates PostChannelsLocalSmsBadRequest with default headers values
func NewPostChannelsLocalSmsBadRequest() *PostChannelsLocalSmsBadRequest {

	return &PostChannelsLocalSmsBadRequest{}
}

// WithPayload adds the payload to the post channels local sms bad request response
func (o *PostChannelsLocalSmsBadRequest) WithPayload(payload *PostChannelsLocalSmsBadRequestBody) *PostChannelsLocalSmsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post channels local sms bad request response
func (o *PostChannelsLocalSmsBadRequest) SetPayload(payload *PostChannelsLocalSmsBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostChannelsLocalSmsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostChannelsLocalSmsForbiddenCode is the HTTP code returned for type PostChannelsLocalSmsForbidden
const PostChannelsLocalSmsForbiddenCode int = 403

/*PostChannelsLocalSmsForbidden Forbidden

swagger:response postChannelsLocalSmsForbidden
*/
type PostChannelsLocalSmsForbidden struct {

	/*
	  In: Body
	*/
	Payload *PostChannelsLocalSmsForbiddenBody `json:"body,omitempty"`
}

// NewPostChannelsLocalSmsForbidden creates PostChannelsLocalSmsForbidden with default headers values
func NewPostChannelsLocalSmsForbidden() *PostChannelsLocalSmsForbidden {

	return &PostChannelsLocalSmsForbidden{}
}

// WithPayload adds the payload to the post channels local sms forbidden response
func (o *PostChannelsLocalSmsForbidden) WithPayload(payload *PostChannelsLocalSmsForbiddenBody) *PostChannelsLocalSmsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post channels local sms forbidden response
func (o *PostChannelsLocalSmsForbidden) SetPayload(payload *PostChannelsLocalSmsForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostChannelsLocalSmsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostChannelsLocalSmsMethodNotAllowedCode is the HTTP code returned for type PostChannelsLocalSmsMethodNotAllowed
const PostChannelsLocalSmsMethodNotAllowedCode int = 405

/*PostChannelsLocalSmsMethodNotAllowed Invalid Method

swagger:response postChannelsLocalSmsMethodNotAllowed
*/
type PostChannelsLocalSmsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *PostChannelsLocalSmsMethodNotAllowedBody `json:"body,omitempty"`
}

// NewPostChannelsLocalSmsMethodNotAllowed creates PostChannelsLocalSmsMethodNotAllowed with default headers values
func NewPostChannelsLocalSmsMethodNotAllowed() *PostChannelsLocalSmsMethodNotAllowed {

	return &PostChannelsLocalSmsMethodNotAllowed{}
}

// WithPayload adds the payload to the post channels local sms method not allowed response
func (o *PostChannelsLocalSmsMethodNotAllowed) WithPayload(payload *PostChannelsLocalSmsMethodNotAllowedBody) *PostChannelsLocalSmsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post channels local sms method not allowed response
func (o *PostChannelsLocalSmsMethodNotAllowed) SetPayload(payload *PostChannelsLocalSmsMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostChannelsLocalSmsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
