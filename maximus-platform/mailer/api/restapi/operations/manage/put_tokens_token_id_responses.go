// Code generated by go-swagger; DO NOT EDIT.

package manage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutTokensTokenIDOKCode is the HTTP code returned for type PutTokensTokenIDOK
const PutTokensTokenIDOKCode int = 200

/*PutTokensTokenIDOK Коллекция токенов

swagger:response putTokensTokenIdOK
*/
type PutTokensTokenIDOK struct {

	/*
	  In: Body
	*/
	Payload *PutTokensTokenIDOKBody `json:"body,omitempty"`
}

// NewPutTokensTokenIDOK creates PutTokensTokenIDOK with default headers values
func NewPutTokensTokenIDOK() *PutTokensTokenIDOK {

	return &PutTokensTokenIDOK{}
}

// WithPayload adds the payload to the put tokens token Id o k response
func (o *PutTokensTokenIDOK) WithPayload(payload *PutTokensTokenIDOKBody) *PutTokensTokenIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put tokens token Id o k response
func (o *PutTokensTokenIDOK) SetPayload(payload *PutTokensTokenIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTokensTokenIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutTokensTokenIDBadRequestCode is the HTTP code returned for type PutTokensTokenIDBadRequest
const PutTokensTokenIDBadRequestCode int = 400

/*PutTokensTokenIDBadRequest Validation error

swagger:response putTokensTokenIdBadRequest
*/
type PutTokensTokenIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PutTokensTokenIDBadRequestBody `json:"body,omitempty"`
}

// NewPutTokensTokenIDBadRequest creates PutTokensTokenIDBadRequest with default headers values
func NewPutTokensTokenIDBadRequest() *PutTokensTokenIDBadRequest {

	return &PutTokensTokenIDBadRequest{}
}

// WithPayload adds the payload to the put tokens token Id bad request response
func (o *PutTokensTokenIDBadRequest) WithPayload(payload *PutTokensTokenIDBadRequestBody) *PutTokensTokenIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put tokens token Id bad request response
func (o *PutTokensTokenIDBadRequest) SetPayload(payload *PutTokensTokenIDBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTokensTokenIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutTokensTokenIDForbiddenCode is the HTTP code returned for type PutTokensTokenIDForbidden
const PutTokensTokenIDForbiddenCode int = 403

/*PutTokensTokenIDForbidden Forbidden

swagger:response putTokensTokenIdForbidden
*/
type PutTokensTokenIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *PutTokensTokenIDForbiddenBody `json:"body,omitempty"`
}

// NewPutTokensTokenIDForbidden creates PutTokensTokenIDForbidden with default headers values
func NewPutTokensTokenIDForbidden() *PutTokensTokenIDForbidden {

	return &PutTokensTokenIDForbidden{}
}

// WithPayload adds the payload to the put tokens token Id forbidden response
func (o *PutTokensTokenIDForbidden) WithPayload(payload *PutTokensTokenIDForbiddenBody) *PutTokensTokenIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put tokens token Id forbidden response
func (o *PutTokensTokenIDForbidden) SetPayload(payload *PutTokensTokenIDForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTokensTokenIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutTokensTokenIDMethodNotAllowedCode is the HTTP code returned for type PutTokensTokenIDMethodNotAllowed
const PutTokensTokenIDMethodNotAllowedCode int = 405

/*PutTokensTokenIDMethodNotAllowed Invalid Method

swagger:response putTokensTokenIdMethodNotAllowed
*/
type PutTokensTokenIDMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *PutTokensTokenIDMethodNotAllowedBody `json:"body,omitempty"`
}

// NewPutTokensTokenIDMethodNotAllowed creates PutTokensTokenIDMethodNotAllowed with default headers values
func NewPutTokensTokenIDMethodNotAllowed() *PutTokensTokenIDMethodNotAllowed {

	return &PutTokensTokenIDMethodNotAllowed{}
}

// WithPayload adds the payload to the put tokens token Id method not allowed response
func (o *PutTokensTokenIDMethodNotAllowed) WithPayload(payload *PutTokensTokenIDMethodNotAllowedBody) *PutTokensTokenIDMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put tokens token Id method not allowed response
func (o *PutTokensTokenIDMethodNotAllowed) SetPayload(payload *PutTokensTokenIDMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutTokensTokenIDMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
