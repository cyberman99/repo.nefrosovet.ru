// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetBackendsBackendIDTestOKCode is the HTTP code returned for type GetBackendsBackendIDTestOK
const GetBackendsBackendIDTestOKCode int = 200

/*GetBackendsBackendIDTestOK Объект теста бэкенда

swagger:response getBackendsBackendIdTestOK
*/
type GetBackendsBackendIDTestOK struct {

	/*
	  In: Body
	*/
	Payload *GetBackendsBackendIDTestOKBody `json:"body,omitempty"`
}

// NewGetBackendsBackendIDTestOK creates GetBackendsBackendIDTestOK with default headers values
func NewGetBackendsBackendIDTestOK() *GetBackendsBackendIDTestOK {

	return &GetBackendsBackendIDTestOK{}
}

// WithPayload adds the payload to the get backends backend Id test o k response
func (o *GetBackendsBackendIDTestOK) WithPayload(payload *GetBackendsBackendIDTestOKBody) *GetBackendsBackendIDTestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backends backend Id test o k response
func (o *GetBackendsBackendIDTestOK) SetPayload(payload *GetBackendsBackendIDTestOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackendsBackendIDTestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBackendsBackendIDTestBadRequestCode is the HTTP code returned for type GetBackendsBackendIDTestBadRequest
const GetBackendsBackendIDTestBadRequestCode int = 400

/*GetBackendsBackendIDTestBadRequest Объект теста бэкенда

swagger:response getBackendsBackendIdTestBadRequest
*/
type GetBackendsBackendIDTestBadRequest struct {

	/*
	  In: Body
	*/
	Payload *GetBackendsBackendIDTestBadRequestBody `json:"body,omitempty"`
}

// NewGetBackendsBackendIDTestBadRequest creates GetBackendsBackendIDTestBadRequest with default headers values
func NewGetBackendsBackendIDTestBadRequest() *GetBackendsBackendIDTestBadRequest {

	return &GetBackendsBackendIDTestBadRequest{}
}

// WithPayload adds the payload to the get backends backend Id test bad request response
func (o *GetBackendsBackendIDTestBadRequest) WithPayload(payload *GetBackendsBackendIDTestBadRequestBody) *GetBackendsBackendIDTestBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backends backend Id test bad request response
func (o *GetBackendsBackendIDTestBadRequest) SetPayload(payload *GetBackendsBackendIDTestBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackendsBackendIDTestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBackendsBackendIDTestNotFoundCode is the HTTP code returned for type GetBackendsBackendIDTestNotFound
const GetBackendsBackendIDTestNotFoundCode int = 404

/*GetBackendsBackendIDTestNotFound Not found

swagger:response getBackendsBackendIdTestNotFound
*/
type GetBackendsBackendIDTestNotFound struct {

	/*
	  In: Body
	*/
	Payload *GetBackendsBackendIDTestNotFoundBody `json:"body,omitempty"`
}

// NewGetBackendsBackendIDTestNotFound creates GetBackendsBackendIDTestNotFound with default headers values
func NewGetBackendsBackendIDTestNotFound() *GetBackendsBackendIDTestNotFound {

	return &GetBackendsBackendIDTestNotFound{}
}

// WithPayload adds the payload to the get backends backend Id test not found response
func (o *GetBackendsBackendIDTestNotFound) WithPayload(payload *GetBackendsBackendIDTestNotFoundBody) *GetBackendsBackendIDTestNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backends backend Id test not found response
func (o *GetBackendsBackendIDTestNotFound) SetPayload(payload *GetBackendsBackendIDTestNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackendsBackendIDTestNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBackendsBackendIDTestMethodNotAllowedCode is the HTTP code returned for type GetBackendsBackendIDTestMethodNotAllowed
const GetBackendsBackendIDTestMethodNotAllowedCode int = 405

/*GetBackendsBackendIDTestMethodNotAllowed Invalid Method

swagger:response getBackendsBackendIdTestMethodNotAllowed
*/
type GetBackendsBackendIDTestMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *GetBackendsBackendIDTestMethodNotAllowedBody `json:"body,omitempty"`
}

// NewGetBackendsBackendIDTestMethodNotAllowed creates GetBackendsBackendIDTestMethodNotAllowed with default headers values
func NewGetBackendsBackendIDTestMethodNotAllowed() *GetBackendsBackendIDTestMethodNotAllowed {

	return &GetBackendsBackendIDTestMethodNotAllowed{}
}

// WithPayload adds the payload to the get backends backend Id test method not allowed response
func (o *GetBackendsBackendIDTestMethodNotAllowed) WithPayload(payload *GetBackendsBackendIDTestMethodNotAllowedBody) *GetBackendsBackendIDTestMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backends backend Id test method not allowed response
func (o *GetBackendsBackendIDTestMethodNotAllowed) SetPayload(payload *GetBackendsBackendIDTestMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackendsBackendIDTestMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBackendsBackendIDTestInternalServerErrorCode is the HTTP code returned for type GetBackendsBackendIDTestInternalServerError
const GetBackendsBackendIDTestInternalServerErrorCode int = 500

/*GetBackendsBackendIDTestInternalServerError Internal server error

swagger:response getBackendsBackendIdTestInternalServerError
*/
type GetBackendsBackendIDTestInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetBackendsBackendIDTestInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetBackendsBackendIDTestInternalServerError creates GetBackendsBackendIDTestInternalServerError with default headers values
func NewGetBackendsBackendIDTestInternalServerError() *GetBackendsBackendIDTestInternalServerError {

	return &GetBackendsBackendIDTestInternalServerError{}
}

// WithPayload adds the payload to the get backends backend Id test internal server error response
func (o *GetBackendsBackendIDTestInternalServerError) WithPayload(payload *GetBackendsBackendIDTestInternalServerErrorBody) *GetBackendsBackendIDTestInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get backends backend Id test internal server error response
func (o *GetBackendsBackendIDTestInternalServerError) SetPayload(payload *GetBackendsBackendIDTestInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBackendsBackendIDTestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
