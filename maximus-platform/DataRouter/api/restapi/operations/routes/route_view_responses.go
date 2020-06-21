// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// RouteViewOKCode is the HTTP code returned for type RouteViewOK
const RouteViewOKCode int = 200

/*RouteViewOK Коллекция маршрутов

swagger:response routeViewOK
*/
type RouteViewOK struct {

	/*
	  In: Body
	*/
	Payload *RouteViewOKBody `json:"body,omitempty"`
}

// NewRouteViewOK creates RouteViewOK with default headers values
func NewRouteViewOK() *RouteViewOK {

	return &RouteViewOK{}
}

// WithPayload adds the payload to the route view o k response
func (o *RouteViewOK) WithPayload(payload *RouteViewOKBody) *RouteViewOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the route view o k response
func (o *RouteViewOK) SetPayload(payload *RouteViewOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RouteViewOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RouteViewNotFoundCode is the HTTP code returned for type RouteViewNotFound
const RouteViewNotFoundCode int = 404

/*RouteViewNotFound Not found

swagger:response routeViewNotFound
*/
type RouteViewNotFound struct {

	/*
	  In: Body
	*/
	Payload *RouteViewNotFoundBody `json:"body,omitempty"`
}

// NewRouteViewNotFound creates RouteViewNotFound with default headers values
func NewRouteViewNotFound() *RouteViewNotFound {

	return &RouteViewNotFound{}
}

// WithPayload adds the payload to the route view not found response
func (o *RouteViewNotFound) WithPayload(payload *RouteViewNotFoundBody) *RouteViewNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the route view not found response
func (o *RouteViewNotFound) SetPayload(payload *RouteViewNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RouteViewNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RouteViewMethodNotAllowedCode is the HTTP code returned for type RouteViewMethodNotAllowed
const RouteViewMethodNotAllowedCode int = 405

/*RouteViewMethodNotAllowed Invalid Method

swagger:response routeViewMethodNotAllowed
*/
type RouteViewMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *RouteViewMethodNotAllowedBody `json:"body,omitempty"`
}

// NewRouteViewMethodNotAllowed creates RouteViewMethodNotAllowed with default headers values
func NewRouteViewMethodNotAllowed() *RouteViewMethodNotAllowed {

	return &RouteViewMethodNotAllowed{}
}

// WithPayload adds the payload to the route view method not allowed response
func (o *RouteViewMethodNotAllowed) WithPayload(payload *RouteViewMethodNotAllowedBody) *RouteViewMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the route view method not allowed response
func (o *RouteViewMethodNotAllowed) SetPayload(payload *RouteViewMethodNotAllowedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RouteViewMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RouteViewInternalServerErrorCode is the HTTP code returned for type RouteViewInternalServerError
const RouteViewInternalServerErrorCode int = 500

/*RouteViewInternalServerError Internal sersver error

swagger:response routeViewInternalServerError
*/
type RouteViewInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *RouteViewInternalServerErrorBody `json:"body,omitempty"`
}

// NewRouteViewInternalServerError creates RouteViewInternalServerError with default headers values
func NewRouteViewInternalServerError() *RouteViewInternalServerError {

	return &RouteViewInternalServerError{}
}

// WithPayload adds the payload to the route view internal server error response
func (o *RouteViewInternalServerError) WithPayload(payload *RouteViewInternalServerErrorBody) *RouteViewInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the route view internal server error response
func (o *RouteViewInternalServerError) SetPayload(payload *RouteViewInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RouteViewInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
