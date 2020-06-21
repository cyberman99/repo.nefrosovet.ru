// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "repo.nefrosovet.ru/maximus-platform/auth/api/models"
)

// PostIdentifyHandlerFunc turns a function with the right signature into a post identify handler
type PostIdentifyHandlerFunc func(PostIdentifyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostIdentifyHandlerFunc) Handle(params PostIdentifyParams) middleware.Responder {
	return fn(params)
}

// PostIdentifyHandler interface for that can handle valid post identify params
type PostIdentifyHandler interface {
	Handle(PostIdentifyParams) middleware.Responder
}

// NewPostIdentify creates a new http.Handler for the post identify operation
func NewPostIdentify(ctx *middleware.Context, handler PostIdentifyHandler) *PostIdentify {
	return &PostIdentify{Context: ctx, Handler: handler}
}

/*PostIdentify swagger:route POST /identify Auth postIdentify

Идентификация пользователя

*/
type PostIdentify struct {
	Context *middleware.Context
	Handler PostIdentifyHandler
}

func (o *PostIdentify) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostIdentifyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostIdentifyBody post identify body
// swagger:model PostIdentifyBody
type PostIdentifyBody struct {
	models.IdentifyParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostIdentifyBody) UnmarshalJSON(raw []byte) error {
	// PostIdentifyParamsBodyAO0
	var postIdentifyParamsBodyAO0 models.IdentifyParams
	if err := swag.ReadJSON(raw, &postIdentifyParamsBodyAO0); err != nil {
		return err
	}
	o.IdentifyParams = postIdentifyParamsBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostIdentifyBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postIdentifyParamsBodyAO0, err := swag.WriteJSON(o.IdentifyParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postIdentifyParamsBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post identify body
func (o *PostIdentifyBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.IdentifyParams
	if err := o.IdentifyParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostIdentifyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostIdentifyBody) UnmarshalBinary(b []byte) error {
	var res PostIdentifyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostIdentifyInternalServerErrorBody post identify internal server error body
// swagger:model PostIdentifyInternalServerErrorBody
type PostIdentifyInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostIdentifyInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// PostIdentifyInternalServerErrorBodyAO0
	var postIdentifyInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &postIdentifyInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = postIdentifyInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostIdentifyInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postIdentifyInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postIdentifyInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post identify internal server error body
func (o *PostIdentifyInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *PostIdentifyInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostIdentifyInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostIdentifyInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostIdentifyMethodNotAllowedBody post identify method not allowed body
// swagger:model PostIdentifyMethodNotAllowedBody
type PostIdentifyMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostIdentifyMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PostIdentifyMethodNotAllowedBodyAO0
	var postIdentifyMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &postIdentifyMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = postIdentifyMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostIdentifyMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postIdentifyMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postIdentifyMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post identify method not allowed body
func (o *PostIdentifyMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PostIdentifyMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostIdentifyMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PostIdentifyMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostIdentifyOKBody post identify o k body
// swagger:model PostIdentifyOKBody
type PostIdentifyOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostIdentifyOKBody) UnmarshalJSON(raw []byte) error {
	// PostIdentifyOKBodyAO0
	var postIdentifyOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postIdentifyOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postIdentifyOKBodyAO0

	// PostIdentifyOKBodyAO1
	var dataPostIdentifyOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataPostIdentifyOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostIdentifyOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostIdentifyOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postIdentifyOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postIdentifyOKBodyAO0)

	var dataPostIdentifyOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataPostIdentifyOKBodyAO1.Data = o.Data

	jsonDataPostIdentifyOKBodyAO1, errPostIdentifyOKBodyAO1 := swag.WriteJSON(dataPostIdentifyOKBodyAO1)
	if errPostIdentifyOKBodyAO1 != nil {
		return nil, errPostIdentifyOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostIdentifyOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post identify o k body
func (o *PostIdentifyOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostIdentifyOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("postIdentifyOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostIdentifyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostIdentifyOKBody) UnmarshalBinary(b []byte) error {
	var res PostIdentifyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostIdentifyUnauthorizedBody post identify unauthorized body
// swagger:model PostIdentifyUnauthorizedBody
type PostIdentifyUnauthorizedBody struct {
	models.Error401Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostIdentifyUnauthorizedBody) UnmarshalJSON(raw []byte) error {
	// PostIdentifyUnauthorizedBodyAO0
	var postIdentifyUnauthorizedBodyAO0 models.Error401Data
	if err := swag.ReadJSON(raw, &postIdentifyUnauthorizedBodyAO0); err != nil {
		return err
	}
	o.Error401Data = postIdentifyUnauthorizedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostIdentifyUnauthorizedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postIdentifyUnauthorizedBodyAO0, err := swag.WriteJSON(o.Error401Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postIdentifyUnauthorizedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post identify unauthorized body
func (o *PostIdentifyUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error401Data
	if err := o.Error401Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostIdentifyUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostIdentifyUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res PostIdentifyUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
