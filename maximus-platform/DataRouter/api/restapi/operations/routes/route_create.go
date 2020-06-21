// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "repo.nefrosovet.ru/maximus-platform/DataRouter/api/models"
)

// RouteCreateHandlerFunc turns a function with the right signature into a route create handler
type RouteCreateHandlerFunc func(RouteCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RouteCreateHandlerFunc) Handle(params RouteCreateParams) middleware.Responder {
	return fn(params)
}

// RouteCreateHandler interface for that can handle valid route create params
type RouteCreateHandler interface {
	Handle(RouteCreateParams) middleware.Responder
}

// NewRouteCreate creates a new http.Handler for the route create operation
func NewRouteCreate(ctx *middleware.Context, handler RouteCreateHandler) *RouteCreate {
	return &RouteCreate{Context: ctx, Handler: handler}
}

/*RouteCreate swagger:route POST /routes Routes routeCreate

Создание маршрута

*/
type RouteCreate struct {
	Context *middleware.Context
	Handler RouteCreateHandler
}

func (o *RouteCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRouteCreateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// RouteCreateBadRequestBody route create bad request body
// swagger:model RouteCreateBadRequestBody
type RouteCreateBadRequestBody struct {
	models.Error400Data

	RouteCreateBadRequestBodyAllOf1

	// errors
	Errors *RouteCreateBadRequestBodyAO2Errors `json:"errors,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RouteCreateBadRequestBody) UnmarshalJSON(raw []byte) error {
	// RouteCreateBadRequestBodyAO0
	var routeCreateBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &routeCreateBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = routeCreateBadRequestBodyAO0

	// RouteCreateBadRequestBodyAO1
	var routeCreateBadRequestBodyAO1 RouteCreateBadRequestBodyAllOf1
	if err := swag.ReadJSON(raw, &routeCreateBadRequestBodyAO1); err != nil {
		return err
	}
	o.RouteCreateBadRequestBodyAllOf1 = routeCreateBadRequestBodyAO1

	// RouteCreateBadRequestBodyAO2
	var dataRouteCreateBadRequestBodyAO2 struct {
		Errors *RouteCreateBadRequestBodyAO2Errors `json:"errors,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataRouteCreateBadRequestBodyAO2); err != nil {
		return err
	}

	o.Errors = dataRouteCreateBadRequestBodyAO2.Errors

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RouteCreateBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	routeCreateBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeCreateBadRequestBodyAO0)

	routeCreateBadRequestBodyAO1, err := swag.WriteJSON(o.RouteCreateBadRequestBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeCreateBadRequestBodyAO1)

	var dataRouteCreateBadRequestBodyAO2 struct {
		Errors *RouteCreateBadRequestBodyAO2Errors `json:"errors,omitempty"`
	}

	dataRouteCreateBadRequestBodyAO2.Errors = o.Errors

	jsonDataRouteCreateBadRequestBodyAO2, errRouteCreateBadRequestBodyAO2 := swag.WriteJSON(dataRouteCreateBadRequestBodyAO2)
	if errRouteCreateBadRequestBodyAO2 != nil {
		return nil, errRouteCreateBadRequestBodyAO2
	}
	_parts = append(_parts, jsonDataRouteCreateBadRequestBodyAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route create bad request body
func (o *RouteCreateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error400Data
	if err := o.Error400Data.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with RouteCreateBadRequestBodyAllOf1

	if err := o.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RouteCreateBadRequestBody) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.Errors) { // not required
		return nil
	}

	if o.Errors != nil {
		if err := o.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routeCreateBadRequest" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RouteCreateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteCreateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res RouteCreateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RouteCreateBadRequestBodyAO2Errors route create bad request body a o2 errors
// swagger:model RouteCreateBadRequestBodyAO2Errors
type RouteCreateBadRequestBodyAO2Errors struct {

	// validation
	Validation *RouteCreateBadRequestBodyAO2ErrorsValidation `json:"validation,omitempty"`
}

// Validate validates this route create bad request body a o2 errors
func (o *RouteCreateBadRequestBodyAO2Errors) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValidation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RouteCreateBadRequestBodyAO2Errors) validateValidation(formats strfmt.Registry) error {

	if swag.IsZero(o.Validation) { // not required
		return nil
	}

	if o.Validation != nil {
		if err := o.Validation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routeCreateBadRequest" + "." + "errors" + "." + "validation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RouteCreateBadRequestBodyAO2Errors) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteCreateBadRequestBodyAO2Errors) UnmarshalBinary(b []byte) error {
	var res RouteCreateBadRequestBodyAO2Errors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RouteCreateBadRequestBodyAO2ErrorsValidation route create bad request body a o2 errors validation
// swagger:model RouteCreateBadRequestBodyAO2ErrorsValidation
type RouteCreateBadRequestBodyAO2ErrorsValidation struct {

	// dst
	// Enum: [object required]
	Dst string `json:"dst,omitempty"`

	// dst qos
	// Enum: [int required]
	DstQos string `json:"dst.qos,omitempty"`

	// dst topic
	// Enum: [object required]
	DstTopic string `json:"dst.topic,omitempty"`

	// reply ID
	// Enum: [string not_found]
	ReplyID string `json:"replyID,omitempty"`

	// src
	// Enum: [object required]
	Src string `json:"src,omitempty"`

	// src payload
	// Enum: [object]
	SrcPayload string `json:"src.payload,omitempty"`

	// src topic
	// Enum: [object required]
	SrcTopic string `json:"src.topic,omitempty"`
}

// Validate validates this route create bad request body a o2 errors validation
func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDst(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDstQos(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDstTopic(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReplyID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSrc(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSrcPayload(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSrcTopic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var routeCreateBadRequestBodyAO2ErrorsValidationTypeDstPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["object","required"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeCreateBadRequestBodyAO2ErrorsValidationTypeDstPropEnum = append(routeCreateBadRequestBodyAO2ErrorsValidationTypeDstPropEnum, v)
	}
}

const (

	// RouteCreateBadRequestBodyAO2ErrorsValidationDstObject captures enum value "object"
	RouteCreateBadRequestBodyAO2ErrorsValidationDstObject string = "object"

	// RouteCreateBadRequestBodyAO2ErrorsValidationDstRequired captures enum value "required"
	RouteCreateBadRequestBodyAO2ErrorsValidationDstRequired string = "required"
)

// prop value enum
func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateDstEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, routeCreateBadRequestBodyAO2ErrorsValidationTypeDstPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateDst(formats strfmt.Registry) error {

	if swag.IsZero(o.Dst) { // not required
		return nil
	}

	// value enum
	if err := o.validateDstEnum("routeCreateBadRequest"+"."+"errors"+"."+"validation"+"."+"dst", "body", o.Dst); err != nil {
		return err
	}

	return nil
}

var routeCreateBadRequestBodyAO2ErrorsValidationTypeDstQosPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["int","required"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeCreateBadRequestBodyAO2ErrorsValidationTypeDstQosPropEnum = append(routeCreateBadRequestBodyAO2ErrorsValidationTypeDstQosPropEnum, v)
	}
}

const (

	// RouteCreateBadRequestBodyAO2ErrorsValidationDstQosInt captures enum value "int"
	RouteCreateBadRequestBodyAO2ErrorsValidationDstQosInt string = "int"

	// RouteCreateBadRequestBodyAO2ErrorsValidationDstQosRequired captures enum value "required"
	RouteCreateBadRequestBodyAO2ErrorsValidationDstQosRequired string = "required"
)

// prop value enum
func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateDstQosEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, routeCreateBadRequestBodyAO2ErrorsValidationTypeDstQosPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateDstQos(formats strfmt.Registry) error {

	if swag.IsZero(o.DstQos) { // not required
		return nil
	}

	// value enum
	if err := o.validateDstQosEnum("routeCreateBadRequest"+"."+"errors"+"."+"validation"+"."+"dst.qos", "body", o.DstQos); err != nil {
		return err
	}

	return nil
}

var routeCreateBadRequestBodyAO2ErrorsValidationTypeDstTopicPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["object","required"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeCreateBadRequestBodyAO2ErrorsValidationTypeDstTopicPropEnum = append(routeCreateBadRequestBodyAO2ErrorsValidationTypeDstTopicPropEnum, v)
	}
}

const (

	// RouteCreateBadRequestBodyAO2ErrorsValidationDstTopicObject captures enum value "object"
	RouteCreateBadRequestBodyAO2ErrorsValidationDstTopicObject string = "object"

	// RouteCreateBadRequestBodyAO2ErrorsValidationDstTopicRequired captures enum value "required"
	RouteCreateBadRequestBodyAO2ErrorsValidationDstTopicRequired string = "required"
)

// prop value enum
func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateDstTopicEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, routeCreateBadRequestBodyAO2ErrorsValidationTypeDstTopicPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateDstTopic(formats strfmt.Registry) error {

	if swag.IsZero(o.DstTopic) { // not required
		return nil
	}

	// value enum
	if err := o.validateDstTopicEnum("routeCreateBadRequest"+"."+"errors"+"."+"validation"+"."+"dst.topic", "body", o.DstTopic); err != nil {
		return err
	}

	return nil
}

var routeCreateBadRequestBodyAO2ErrorsValidationTypeReplyIDPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string","not_found"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeCreateBadRequestBodyAO2ErrorsValidationTypeReplyIDPropEnum = append(routeCreateBadRequestBodyAO2ErrorsValidationTypeReplyIDPropEnum, v)
	}
}

const (

	// RouteCreateBadRequestBodyAO2ErrorsValidationReplyIDString captures enum value "string"
	RouteCreateBadRequestBodyAO2ErrorsValidationReplyIDString string = "string"

	// RouteCreateBadRequestBodyAO2ErrorsValidationReplyIDNotFound captures enum value "not_found"
	RouteCreateBadRequestBodyAO2ErrorsValidationReplyIDNotFound string = "not_found"
)

// prop value enum
func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateReplyIDEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, routeCreateBadRequestBodyAO2ErrorsValidationTypeReplyIDPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateReplyID(formats strfmt.Registry) error {

	if swag.IsZero(o.ReplyID) { // not required
		return nil
	}

	// value enum
	if err := o.validateReplyIDEnum("routeCreateBadRequest"+"."+"errors"+"."+"validation"+"."+"replyID", "body", o.ReplyID); err != nil {
		return err
	}

	return nil
}

var routeCreateBadRequestBodyAO2ErrorsValidationTypeSrcPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["object","required"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeCreateBadRequestBodyAO2ErrorsValidationTypeSrcPropEnum = append(routeCreateBadRequestBodyAO2ErrorsValidationTypeSrcPropEnum, v)
	}
}

const (

	// RouteCreateBadRequestBodyAO2ErrorsValidationSrcObject captures enum value "object"
	RouteCreateBadRequestBodyAO2ErrorsValidationSrcObject string = "object"

	// RouteCreateBadRequestBodyAO2ErrorsValidationSrcRequired captures enum value "required"
	RouteCreateBadRequestBodyAO2ErrorsValidationSrcRequired string = "required"
)

// prop value enum
func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateSrcEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, routeCreateBadRequestBodyAO2ErrorsValidationTypeSrcPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateSrc(formats strfmt.Registry) error {

	if swag.IsZero(o.Src) { // not required
		return nil
	}

	// value enum
	if err := o.validateSrcEnum("routeCreateBadRequest"+"."+"errors"+"."+"validation"+"."+"src", "body", o.Src); err != nil {
		return err
	}

	return nil
}

var routeCreateBadRequestBodyAO2ErrorsValidationTypeSrcPayloadPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["object"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeCreateBadRequestBodyAO2ErrorsValidationTypeSrcPayloadPropEnum = append(routeCreateBadRequestBodyAO2ErrorsValidationTypeSrcPayloadPropEnum, v)
	}
}

const (

	// RouteCreateBadRequestBodyAO2ErrorsValidationSrcPayloadObject captures enum value "object"
	RouteCreateBadRequestBodyAO2ErrorsValidationSrcPayloadObject string = "object"
)

// prop value enum
func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateSrcPayloadEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, routeCreateBadRequestBodyAO2ErrorsValidationTypeSrcPayloadPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateSrcPayload(formats strfmt.Registry) error {

	if swag.IsZero(o.SrcPayload) { // not required
		return nil
	}

	// value enum
	if err := o.validateSrcPayloadEnum("routeCreateBadRequest"+"."+"errors"+"."+"validation"+"."+"src.payload", "body", o.SrcPayload); err != nil {
		return err
	}

	return nil
}

var routeCreateBadRequestBodyAO2ErrorsValidationTypeSrcTopicPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["object","required"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeCreateBadRequestBodyAO2ErrorsValidationTypeSrcTopicPropEnum = append(routeCreateBadRequestBodyAO2ErrorsValidationTypeSrcTopicPropEnum, v)
	}
}

const (

	// RouteCreateBadRequestBodyAO2ErrorsValidationSrcTopicObject captures enum value "object"
	RouteCreateBadRequestBodyAO2ErrorsValidationSrcTopicObject string = "object"

	// RouteCreateBadRequestBodyAO2ErrorsValidationSrcTopicRequired captures enum value "required"
	RouteCreateBadRequestBodyAO2ErrorsValidationSrcTopicRequired string = "required"
)

// prop value enum
func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateSrcTopicEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, routeCreateBadRequestBodyAO2ErrorsValidationTypeSrcTopicPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) validateSrcTopic(formats strfmt.Registry) error {

	if swag.IsZero(o.SrcTopic) { // not required
		return nil
	}

	// value enum
	if err := o.validateSrcTopicEnum("routeCreateBadRequest"+"."+"errors"+"."+"validation"+"."+"src.topic", "body", o.SrcTopic); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteCreateBadRequestBodyAO2ErrorsValidation) UnmarshalBinary(b []byte) error {
	var res RouteCreateBadRequestBodyAO2ErrorsValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RouteCreateBadRequestBodyAllOf1 route create bad request body all of1
// swagger:model RouteCreateBadRequestBodyAllOf1
type RouteCreateBadRequestBodyAllOf1 interface{}

// RouteCreateBody route create body
// swagger:model RouteCreateBody
type RouteCreateBody struct {
	models.RouteObject

	RouteCreateParamsBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RouteCreateBody) UnmarshalJSON(raw []byte) error {
	// RouteCreateParamsBodyAO0
	var routeCreateParamsBodyAO0 models.RouteObject
	if err := swag.ReadJSON(raw, &routeCreateParamsBodyAO0); err != nil {
		return err
	}
	o.RouteObject = routeCreateParamsBodyAO0

	// RouteCreateParamsBodyAO1
	var routeCreateParamsBodyAO1 RouteCreateParamsBodyAllOf1
	if err := swag.ReadJSON(raw, &routeCreateParamsBodyAO1); err != nil {
		return err
	}
	o.RouteCreateParamsBodyAllOf1 = routeCreateParamsBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RouteCreateBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	routeCreateParamsBodyAO0, err := swag.WriteJSON(o.RouteObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeCreateParamsBodyAO0)

	routeCreateParamsBodyAO1, err := swag.WriteJSON(o.RouteCreateParamsBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeCreateParamsBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route create body
func (o *RouteCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.RouteObject
	if err := o.RouteObject.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with RouteCreateParamsBodyAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *RouteCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteCreateBody) UnmarshalBinary(b []byte) error {
	var res RouteCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RouteCreateInternalServerErrorBody route create internal server error body
// swagger:model RouteCreateInternalServerErrorBody
type RouteCreateInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RouteCreateInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// RouteCreateInternalServerErrorBodyAO0
	var routeCreateInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &routeCreateInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = routeCreateInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RouteCreateInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	routeCreateInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeCreateInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route create internal server error body
func (o *RouteCreateInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error500Data
	if err := o.Error500Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *RouteCreateInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteCreateInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res RouteCreateInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RouteCreateMethodNotAllowedBody route create method not allowed body
// swagger:model RouteCreateMethodNotAllowedBody
type RouteCreateMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RouteCreateMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// RouteCreateMethodNotAllowedBodyAO0
	var routeCreateMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &routeCreateMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = routeCreateMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RouteCreateMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	routeCreateMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeCreateMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route create method not allowed body
func (o *RouteCreateMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error405Data
	if err := o.Error405Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *RouteCreateMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteCreateMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res RouteCreateMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RouteCreateOKBody route create o k body
// swagger:model RouteCreateOKBody
type RouteCreateOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RouteCreateOKBody) UnmarshalJSON(raw []byte) error {
	// RouteCreateOKBodyAO0
	var routeCreateOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &routeCreateOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = routeCreateOKBodyAO0

	// RouteCreateOKBodyAO1
	var dataRouteCreateOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataRouteCreateOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataRouteCreateOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RouteCreateOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	routeCreateOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeCreateOKBodyAO0)

	var dataRouteCreateOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataRouteCreateOKBodyAO1.Data = o.Data

	jsonDataRouteCreateOKBodyAO1, errRouteCreateOKBodyAO1 := swag.WriteJSON(dataRouteCreateOKBodyAO1)
	if errRouteCreateOKBodyAO1 != nil {
		return nil, errRouteCreateOKBodyAO1
	}
	_parts = append(_parts, jsonDataRouteCreateOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route create o k body
func (o *RouteCreateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.SuccessData
	if err := o.SuccessData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RouteCreateOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routeCreateOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *RouteCreateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteCreateOKBody) UnmarshalBinary(b []byte) error {
	var res RouteCreateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RouteCreateParamsBodyAllOf1 route create params body all of1
// swagger:model RouteCreateParamsBodyAllOf1
type RouteCreateParamsBodyAllOf1 interface{}
