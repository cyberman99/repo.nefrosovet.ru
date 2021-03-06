// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetChannelsChannelIDOKCode is the HTTP code returned for type GetChannelsChannelIDOK
const GetChannelsChannelIDOKCode int = 200

/*GetChannelsChannelIDOK Коллекция каналов

swagger:response getChannelsChannelIdOK
*/
type GetChannelsChannelIDOK struct {

	/*
	  In: Body
	*/
	Payload *GetChannelsChannelIDOKBody `json:"body,omitempty"`
}

// NewGetChannelsChannelIDOK creates GetChannelsChannelIDOK with default headers values
func NewGetChannelsChannelIDOK() *GetChannelsChannelIDOK {

	return &GetChannelsChannelIDOK{}
}

// WithPayload adds the payload to the get channels channel Id o k response
func (o *GetChannelsChannelIDOK) WithPayload(payload *GetChannelsChannelIDOKBody) *GetChannelsChannelIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get channels channel Id o k response
func (o *GetChannelsChannelIDOK) SetPayload(payload *GetChannelsChannelIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChannelsChannelIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChannelsChannelIDForbiddenCode is the HTTP code returned for type GetChannelsChannelIDForbidden
const GetChannelsChannelIDForbiddenCode int = 403

/*GetChannelsChannelIDForbidden Forbidden

swagger:response getChannelsChannelIdForbidden
*/
type GetChannelsChannelIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *GetChannelsChannelIDForbiddenBody `json:"body,omitempty"`
}

// NewGetChannelsChannelIDForbidden creates GetChannelsChannelIDForbidden with default headers values
func NewGetChannelsChannelIDForbidden() *GetChannelsChannelIDForbidden {

	return &GetChannelsChannelIDForbidden{}
}

// WithPayload adds the payload to the get channels channel Id forbidden response
func (o *GetChannelsChannelIDForbidden) WithPayload(payload *GetChannelsChannelIDForbiddenBody) *GetChannelsChannelIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get channels channel Id forbidden response
func (o *GetChannelsChannelIDForbidden) SetPayload(payload *GetChannelsChannelIDForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChannelsChannelIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChannelsChannelIDNotFoundCode is the HTTP code returned for type GetChannelsChannelIDNotFound
const GetChannelsChannelIDNotFoundCode int = 404

/*GetChannelsChannelIDNotFound Not found

swagger:response getChannelsChannelIdNotFound
*/
type GetChannelsChannelIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *GetChannelsChannelIDNotFoundBody `json:"body,omitempty"`
}

// NewGetChannelsChannelIDNotFound creates GetChannelsChannelIDNotFound with default headers values
func NewGetChannelsChannelIDNotFound() *GetChannelsChannelIDNotFound {

	return &GetChannelsChannelIDNotFound{}
}

// WithPayload adds the payload to the get channels channel Id not found response
func (o *GetChannelsChannelIDNotFound) WithPayload(payload *GetChannelsChannelIDNotFoundBody) *GetChannelsChannelIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get channels channel Id not found response
func (o *GetChannelsChannelIDNotFound) SetPayload(payload *GetChannelsChannelIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChannelsChannelIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChannelsChannelIDMethodNotAllowedCode is the HTTP code returned for type GetChannelsChannelIDMethodNotAllowed
const GetChannelsChannelIDMethodNotAllowedCode int = 405

/*GetChannelsChannelIDMethodNotAllowed Invalid Method

swagger:response getChannelsChannelIdMethodNotAllowed
*/
type GetChannelsChannelIDMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *GetChannelsChannelIDMethodNotAllowedBody `json:"body,omitempty"`
}

// NewGetChannelsChannelIDMethodNotAllowed creates GetChannelsChannelIDMethodNotAllowed with default headers values
func NewGetChannelsChannelIDMethodNotAllowed() *GetChannelsChannelIDMethodNotAllowed {

	return &GetChannelsChannelIDMethodNotAllowed{}
}

// WithPayload adds the payload to the get channels channel Id method not allowed response
func (o *GetChannelsChannelIDMethodNotAllowed) WithPayload(payload *GetChannelsChannelIDMethodNotAllowedBody) *GetChannelsChannelIDMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get channels channel Id method not allowed response
func (o *GetChannelsChannelIDMethodNotAllowed) SetPayload(payload *GetChannelsChannelIDMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChannelsChannelIDMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
