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

// PostRefreshHandlerFunc turns a function with the right signature into a post refresh handler
type PostRefreshHandlerFunc func(PostRefreshParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostRefreshHandlerFunc) Handle(params PostRefreshParams) middleware.Responder {
	return fn(params)
}

// PostRefreshHandler interface for that can handle valid post refresh params
type PostRefreshHandler interface {
	Handle(PostRefreshParams) middleware.Responder
}

// NewPostRefresh creates a new http.Handler for the post refresh operation
func NewPostRefresh(ctx *middleware.Context, handler PostRefreshHandler) *PostRefresh {
	return &PostRefresh{Context: ctx, Handler: handler}
}

/*PostRefresh swagger:route POST /refresh Auth postRefresh

Регенерация токенов

*/
type PostRefresh struct {
	Context *middleware.Context
	Handler PostRefreshHandler
}

func (o *PostRefresh) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostRefreshParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostRefreshBody post refresh body
// swagger:model PostRefreshBody
type PostRefreshBody struct {
	models.AuthRefreshParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRefreshBody) UnmarshalJSON(raw []byte) error {
	// PostRefreshParamsBodyAO0
	var postRefreshParamsBodyAO0 models.AuthRefreshParams
	if err := swag.ReadJSON(raw, &postRefreshParamsBodyAO0); err != nil {
		return err
	}
	o.AuthRefreshParams = postRefreshParamsBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRefreshBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postRefreshParamsBodyAO0, err := swag.WriteJSON(o.AuthRefreshParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRefreshParamsBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post refresh body
func (o *PostRefreshBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.AuthRefreshParams
	if err := o.AuthRefreshParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostRefreshBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRefreshBody) UnmarshalBinary(b []byte) error {
	var res PostRefreshBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostRefreshInternalServerErrorBody post refresh internal server error body
// swagger:model PostRefreshInternalServerErrorBody
type PostRefreshInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRefreshInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// PostRefreshInternalServerErrorBodyAO0
	var postRefreshInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &postRefreshInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = postRefreshInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRefreshInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postRefreshInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRefreshInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post refresh internal server error body
func (o *PostRefreshInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *PostRefreshInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRefreshInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostRefreshInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostRefreshMethodNotAllowedBody post refresh method not allowed body
// swagger:model PostRefreshMethodNotAllowedBody
type PostRefreshMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRefreshMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PostRefreshMethodNotAllowedBodyAO0
	var postRefreshMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &postRefreshMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = postRefreshMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRefreshMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postRefreshMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRefreshMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post refresh method not allowed body
func (o *PostRefreshMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PostRefreshMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRefreshMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PostRefreshMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostRefreshOKBody post refresh o k body
// swagger:model PostRefreshOKBody
type PostRefreshOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRefreshOKBody) UnmarshalJSON(raw []byte) error {
	// PostRefreshOKBodyAO0
	var postRefreshOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postRefreshOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postRefreshOKBodyAO0

	// PostRefreshOKBodyAO1
	var dataPostRefreshOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataPostRefreshOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostRefreshOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRefreshOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postRefreshOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRefreshOKBodyAO0)

	var dataPostRefreshOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataPostRefreshOKBodyAO1.Data = o.Data

	jsonDataPostRefreshOKBodyAO1, errPostRefreshOKBodyAO1 := swag.WriteJSON(dataPostRefreshOKBodyAO1)
	if errPostRefreshOKBodyAO1 != nil {
		return nil, errPostRefreshOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostRefreshOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post refresh o k body
func (o *PostRefreshOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostRefreshOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("postRefreshOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostRefreshOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRefreshOKBody) UnmarshalBinary(b []byte) error {
	var res PostRefreshOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostRefreshUnauthorizedBody post refresh unauthorized body
// swagger:model PostRefreshUnauthorizedBody
type PostRefreshUnauthorizedBody struct {
	models.Error401Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRefreshUnauthorizedBody) UnmarshalJSON(raw []byte) error {
	// PostRefreshUnauthorizedBodyAO0
	var postRefreshUnauthorizedBodyAO0 models.Error401Data
	if err := swag.ReadJSON(raw, &postRefreshUnauthorizedBodyAO0); err != nil {
		return err
	}
	o.Error401Data = postRefreshUnauthorizedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRefreshUnauthorizedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postRefreshUnauthorizedBodyAO0, err := swag.WriteJSON(o.Error401Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRefreshUnauthorizedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post refresh unauthorized body
func (o *PostRefreshUnauthorizedBody) Validate(formats strfmt.Registry) error {
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
func (o *PostRefreshUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRefreshUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res PostRefreshUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}