// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutChannelsMtsSmsChannelIDOKCode is the HTTP code returned for type PutChannelsMtsSmsChannelIDOK
const PutChannelsMtsSmsChannelIDOKCode int = 200

/*PutChannelsMtsSmsChannelIDOK Коллекция каналов

swagger:response putChannelsMtsSmsChannelIdOK
*/
type PutChannelsMtsSmsChannelIDOK struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsMtsSmsChannelIDOKBody `json:"body,omitempty"`
}

// NewPutChannelsMtsSmsChannelIDOK creates PutChannelsMtsSmsChannelIDOK with default headers values
func NewPutChannelsMtsSmsChannelIDOK() *PutChannelsMtsSmsChannelIDOK {

	return &PutChannelsMtsSmsChannelIDOK{}
}

// WithPayload adds the payload to the put channels mts sms channel Id o k response
func (o *PutChannelsMtsSmsChannelIDOK) WithPayload(payload *PutChannelsMtsSmsChannelIDOKBody) *PutChannelsMtsSmsChannelIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels mts sms channel Id o k response
func (o *PutChannelsMtsSmsChannelIDOK) SetPayload(payload *PutChannelsMtsSmsChannelIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsMtsSmsChannelIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutChannelsMtsSmsChannelIDBadRequestCode is the HTTP code returned for type PutChannelsMtsSmsChannelIDBadRequest
const PutChannelsMtsSmsChannelIDBadRequestCode int = 400

/*PutChannelsMtsSmsChannelIDBadRequest Validation error

swagger:response putChannelsMtsSmsChannelIdBadRequest
*/
type PutChannelsMtsSmsChannelIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsMtsSmsChannelIDBadRequestBody `json:"body,omitempty"`
}

// NewPutChannelsMtsSmsChannelIDBadRequest creates PutChannelsMtsSmsChannelIDBadRequest with default headers values
func NewPutChannelsMtsSmsChannelIDBadRequest() *PutChannelsMtsSmsChannelIDBadRequest {

	return &PutChannelsMtsSmsChannelIDBadRequest{}
}

// WithPayload adds the payload to the put channels mts sms channel Id bad request response
func (o *PutChannelsMtsSmsChannelIDBadRequest) WithPayload(payload *PutChannelsMtsSmsChannelIDBadRequestBody) *PutChannelsMtsSmsChannelIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels mts sms channel Id bad request response
func (o *PutChannelsMtsSmsChannelIDBadRequest) SetPayload(payload *PutChannelsMtsSmsChannelIDBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsMtsSmsChannelIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutChannelsMtsSmsChannelIDForbiddenCode is the HTTP code returned for type PutChannelsMtsSmsChannelIDForbidden
const PutChannelsMtsSmsChannelIDForbiddenCode int = 403

/*PutChannelsMtsSmsChannelIDForbidden Forbidden

swagger:response putChannelsMtsSmsChannelIdForbidden
*/
type PutChannelsMtsSmsChannelIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsMtsSmsChannelIDForbiddenBody `json:"body,omitempty"`
}

// NewPutChannelsMtsSmsChannelIDForbidden creates PutChannelsMtsSmsChannelIDForbidden with default headers values
func NewPutChannelsMtsSmsChannelIDForbidden() *PutChannelsMtsSmsChannelIDForbidden {

	return &PutChannelsMtsSmsChannelIDForbidden{}
}

// WithPayload adds the payload to the put channels mts sms channel Id forbidden response
func (o *PutChannelsMtsSmsChannelIDForbidden) WithPayload(payload *PutChannelsMtsSmsChannelIDForbiddenBody) *PutChannelsMtsSmsChannelIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels mts sms channel Id forbidden response
func (o *PutChannelsMtsSmsChannelIDForbidden) SetPayload(payload *PutChannelsMtsSmsChannelIDForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsMtsSmsChannelIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutChannelsMtsSmsChannelIDNotFoundCode is the HTTP code returned for type PutChannelsMtsSmsChannelIDNotFound
const PutChannelsMtsSmsChannelIDNotFoundCode int = 404

/*PutChannelsMtsSmsChannelIDNotFound Not found

swagger:response putChannelsMtsSmsChannelIdNotFound
*/
type PutChannelsMtsSmsChannelIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsMtsSmsChannelIDNotFoundBody `json:"body,omitempty"`
}

// NewPutChannelsMtsSmsChannelIDNotFound creates PutChannelsMtsSmsChannelIDNotFound with default headers values
func NewPutChannelsMtsSmsChannelIDNotFound() *PutChannelsMtsSmsChannelIDNotFound {

	return &PutChannelsMtsSmsChannelIDNotFound{}
}

// WithPayload adds the payload to the put channels mts sms channel Id not found response
func (o *PutChannelsMtsSmsChannelIDNotFound) WithPayload(payload *PutChannelsMtsSmsChannelIDNotFoundBody) *PutChannelsMtsSmsChannelIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels mts sms channel Id not found response
func (o *PutChannelsMtsSmsChannelIDNotFound) SetPayload(payload *PutChannelsMtsSmsChannelIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsMtsSmsChannelIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutChannelsMtsSmsChannelIDMethodNotAllowedCode is the HTTP code returned for type PutChannelsMtsSmsChannelIDMethodNotAllowed
const PutChannelsMtsSmsChannelIDMethodNotAllowedCode int = 405

/*PutChannelsMtsSmsChannelIDMethodNotAllowed Invalid Method

swagger:response putChannelsMtsSmsChannelIdMethodNotAllowed
*/
type PutChannelsMtsSmsChannelIDMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *PutChannelsMtsSmsChannelIDMethodNotAllowedBody `json:"body,omitempty"`
}

// NewPutChannelsMtsSmsChannelIDMethodNotAllowed creates PutChannelsMtsSmsChannelIDMethodNotAllowed with default headers values
func NewPutChannelsMtsSmsChannelIDMethodNotAllowed() *PutChannelsMtsSmsChannelIDMethodNotAllowed {

	return &PutChannelsMtsSmsChannelIDMethodNotAllowed{}
}

// WithPayload adds the payload to the put channels mts sms channel Id method not allowed response
func (o *PutChannelsMtsSmsChannelIDMethodNotAllowed) WithPayload(payload *PutChannelsMtsSmsChannelIDMethodNotAllowedBody) *PutChannelsMtsSmsChannelIDMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put channels mts sms channel Id method not allowed response
func (o *PutChannelsMtsSmsChannelIDMethodNotAllowed) SetPayload(payload *PutChannelsMtsSmsChannelIDMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutChannelsMtsSmsChannelIDMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}