// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetChannelsOKCode is the HTTP code returned for type GetChannelsOK
const GetChannelsOKCode int = 200

/*GetChannelsOK Коллекция каналов

swagger:response getChannelsOK
*/
type GetChannelsOK struct {

	/*
	  In: Body
	*/
	Payload *GetChannelsOKBody `json:"body,omitempty"`
}

// NewGetChannelsOK creates GetChannelsOK with default headers values
func NewGetChannelsOK() *GetChannelsOK {

	return &GetChannelsOK{}
}

// WithPayload adds the payload to the get channels o k response
func (o *GetChannelsOK) WithPayload(payload *GetChannelsOKBody) *GetChannelsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get channels o k response
func (o *GetChannelsOK) SetPayload(payload *GetChannelsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChannelsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChannelsForbiddenCode is the HTTP code returned for type GetChannelsForbidden
const GetChannelsForbiddenCode int = 403

/*GetChannelsForbidden Forbidden

swagger:response getChannelsForbidden
*/
type GetChannelsForbidden struct {

	/*
	  In: Body
	*/
	Payload *GetChannelsForbiddenBody `json:"body,omitempty"`
}

// NewGetChannelsForbidden creates GetChannelsForbidden with default headers values
func NewGetChannelsForbidden() *GetChannelsForbidden {

	return &GetChannelsForbidden{}
}

// WithPayload adds the payload to the get channels forbidden response
func (o *GetChannelsForbidden) WithPayload(payload *GetChannelsForbiddenBody) *GetChannelsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get channels forbidden response
func (o *GetChannelsForbidden) SetPayload(payload *GetChannelsForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChannelsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
