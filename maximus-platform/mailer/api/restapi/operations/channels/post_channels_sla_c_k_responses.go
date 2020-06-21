// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostChannelsSLACKOKCode is the HTTP code returned for type PostChannelsSLACKOK
const PostChannelsSLACKOKCode int = 200

/*PostChannelsSLACKOK Коллекция каналов

swagger:response postChannelsSlaCKOK
*/
type PostChannelsSLACKOK struct {

	/*
	  In: Body
	*/
	Payload *PostChannelsSLACKOKBody `json:"body,omitempty"`
}

// NewPostChannelsSLACKOK creates PostChannelsSLACKOK with default headers values
func NewPostChannelsSLACKOK() *PostChannelsSLACKOK {

	return &PostChannelsSLACKOK{}
}

// WithPayload adds the payload to the post channels Sla c k o k response
func (o *PostChannelsSLACKOK) WithPayload(payload *PostChannelsSLACKOKBody) *PostChannelsSLACKOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post channels Sla c k o k response
func (o *PostChannelsSLACKOK) SetPayload(payload *PostChannelsSLACKOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostChannelsSLACKOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostChannelsSLACKBadRequestCode is the HTTP code returned for type PostChannelsSLACKBadRequest
const PostChannelsSLACKBadRequestCode int = 400

/*PostChannelsSLACKBadRequest Validation error

swagger:response postChannelsSlaCKBadRequest
*/
type PostChannelsSLACKBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostChannelsSLACKBadRequestBody `json:"body,omitempty"`
}

// NewPostChannelsSLACKBadRequest creates PostChannelsSLACKBadRequest with default headers values
func NewPostChannelsSLACKBadRequest() *PostChannelsSLACKBadRequest {

	return &PostChannelsSLACKBadRequest{}
}

// WithPayload adds the payload to the post channels Sla c k bad request response
func (o *PostChannelsSLACKBadRequest) WithPayload(payload *PostChannelsSLACKBadRequestBody) *PostChannelsSLACKBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post channels Sla c k bad request response
func (o *PostChannelsSLACKBadRequest) SetPayload(payload *PostChannelsSLACKBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostChannelsSLACKBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostChannelsSLACKForbiddenCode is the HTTP code returned for type PostChannelsSLACKForbidden
const PostChannelsSLACKForbiddenCode int = 403

/*PostChannelsSLACKForbidden Forbidden

swagger:response postChannelsSlaCKForbidden
*/
type PostChannelsSLACKForbidden struct {

	/*
	  In: Body
	*/
	Payload *PostChannelsSLACKForbiddenBody `json:"body,omitempty"`
}

// NewPostChannelsSLACKForbidden creates PostChannelsSLACKForbidden with default headers values
func NewPostChannelsSLACKForbidden() *PostChannelsSLACKForbidden {

	return &PostChannelsSLACKForbidden{}
}

// WithPayload adds the payload to the post channels Sla c k forbidden response
func (o *PostChannelsSLACKForbidden) WithPayload(payload *PostChannelsSLACKForbiddenBody) *PostChannelsSLACKForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post channels Sla c k forbidden response
func (o *PostChannelsSLACKForbidden) SetPayload(payload *PostChannelsSLACKForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostChannelsSLACKForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostChannelsSLACKMethodNotAllowedCode is the HTTP code returned for type PostChannelsSLACKMethodNotAllowed
const PostChannelsSLACKMethodNotAllowedCode int = 405

/*PostChannelsSLACKMethodNotAllowed Invalid Method

swagger:response postChannelsSlaCKMethodNotAllowed
*/
type PostChannelsSLACKMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *PostChannelsSLACKMethodNotAllowedBody `json:"body,omitempty"`
}

// NewPostChannelsSLACKMethodNotAllowed creates PostChannelsSLACKMethodNotAllowed with default headers values
func NewPostChannelsSLACKMethodNotAllowed() *PostChannelsSLACKMethodNotAllowed {

	return &PostChannelsSLACKMethodNotAllowed{}
}

// WithPayload adds the payload to the post channels Sla c k method not allowed response
func (o *PostChannelsSLACKMethodNotAllowed) WithPayload(payload *PostChannelsSLACKMethodNotAllowedBody) *PostChannelsSLACKMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post channels Sla c k method not allowed response
func (o *PostChannelsSLACKMethodNotAllowed) SetPayload(payload *PostChannelsSLACKMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostChannelsSLACKMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}