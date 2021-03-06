// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "repo.nefrosovet.ru/maximus-platform/auth/api/models"
)

// PostBackendsBackendIDGroupsHandlerFunc turns a function with the right signature into a post backends backend ID groups handler
type PostBackendsBackendIDGroupsHandlerFunc func(PostBackendsBackendIDGroupsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostBackendsBackendIDGroupsHandlerFunc) Handle(params PostBackendsBackendIDGroupsParams) middleware.Responder {
	return fn(params)
}

// PostBackendsBackendIDGroupsHandler interface for that can handle valid post backends backend ID groups params
type PostBackendsBackendIDGroupsHandler interface {
	Handle(PostBackendsBackendIDGroupsParams) middleware.Responder
}

// NewPostBackendsBackendIDGroups creates a new http.Handler for the post backends backend ID groups operation
func NewPostBackendsBackendIDGroups(ctx *middleware.Context, handler PostBackendsBackendIDGroupsHandler) *PostBackendsBackendIDGroups {
	return &PostBackendsBackendIDGroups{Context: ctx, Handler: handler}
}

/*PostBackendsBackendIDGroups swagger:route POST /backends/{backendID}/groups Backend postBackendsBackendIdGroups

Редактироание соответствия групп бэкенда

*/
type PostBackendsBackendIDGroups struct {
	Context *middleware.Context
	Handler PostBackendsBackendIDGroupsHandler
}

func (o *PostBackendsBackendIDGroups) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostBackendsBackendIDGroupsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostBackendsBackendIDGroupsInternalServerErrorBody post backends backend ID groups internal server error body
// swagger:model PostBackendsBackendIDGroupsInternalServerErrorBody
type PostBackendsBackendIDGroupsInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostBackendsBackendIDGroupsInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// PostBackendsBackendIDGroupsInternalServerErrorBodyAO0
	var postBackendsBackendIDGroupsInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &postBackendsBackendIDGroupsInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = postBackendsBackendIDGroupsInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostBackendsBackendIDGroupsInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postBackendsBackendIDGroupsInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postBackendsBackendIDGroupsInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post backends backend ID groups internal server error body
func (o *PostBackendsBackendIDGroupsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *PostBackendsBackendIDGroupsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBackendsBackendIDGroupsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostBackendsBackendIDGroupsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostBackendsBackendIDGroupsMethodNotAllowedBody post backends backend ID groups method not allowed body
// swagger:model PostBackendsBackendIDGroupsMethodNotAllowedBody
type PostBackendsBackendIDGroupsMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostBackendsBackendIDGroupsMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PostBackendsBackendIDGroupsMethodNotAllowedBodyAO0
	var postBackendsBackendIDGroupsMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &postBackendsBackendIDGroupsMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = postBackendsBackendIDGroupsMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostBackendsBackendIDGroupsMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postBackendsBackendIDGroupsMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postBackendsBackendIDGroupsMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post backends backend ID groups method not allowed body
func (o *PostBackendsBackendIDGroupsMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PostBackendsBackendIDGroupsMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBackendsBackendIDGroupsMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PostBackendsBackendIDGroupsMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostBackendsBackendIDGroupsNotFoundBody post backends backend ID groups not found body
// swagger:model PostBackendsBackendIDGroupsNotFoundBody
type PostBackendsBackendIDGroupsNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostBackendsBackendIDGroupsNotFoundBody) UnmarshalJSON(raw []byte) error {
	// PostBackendsBackendIDGroupsNotFoundBodyAO0
	var postBackendsBackendIDGroupsNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &postBackendsBackendIDGroupsNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = postBackendsBackendIDGroupsNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostBackendsBackendIDGroupsNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postBackendsBackendIDGroupsNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postBackendsBackendIDGroupsNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post backends backend ID groups not found body
func (o *PostBackendsBackendIDGroupsNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error404Data
	if err := o.Error404Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostBackendsBackendIDGroupsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBackendsBackendIDGroupsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PostBackendsBackendIDGroupsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostBackendsBackendIDGroupsOKBody post backends backend ID groups o k body
// swagger:model PostBackendsBackendIDGroupsOKBody
type PostBackendsBackendIDGroupsOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostBackendsBackendIDGroupsOKBody) UnmarshalJSON(raw []byte) error {
	// PostBackendsBackendIDGroupsOKBodyAO0
	var postBackendsBackendIDGroupsOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postBackendsBackendIDGroupsOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postBackendsBackendIDGroupsOKBodyAO0

	// PostBackendsBackendIDGroupsOKBodyAO1
	var dataPostBackendsBackendIDGroupsOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataPostBackendsBackendIDGroupsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostBackendsBackendIDGroupsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostBackendsBackendIDGroupsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postBackendsBackendIDGroupsOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postBackendsBackendIDGroupsOKBodyAO0)

	var dataPostBackendsBackendIDGroupsOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataPostBackendsBackendIDGroupsOKBodyAO1.Data = o.Data

	jsonDataPostBackendsBackendIDGroupsOKBodyAO1, errPostBackendsBackendIDGroupsOKBodyAO1 := swag.WriteJSON(dataPostBackendsBackendIDGroupsOKBodyAO1)
	if errPostBackendsBackendIDGroupsOKBodyAO1 != nil {
		return nil, errPostBackendsBackendIDGroupsOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostBackendsBackendIDGroupsOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post backends backend ID groups o k body
func (o *PostBackendsBackendIDGroupsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostBackendsBackendIDGroupsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("postBackendsBackendIdGroupsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postBackendsBackendIdGroupsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostBackendsBackendIDGroupsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBackendsBackendIDGroupsOKBody) UnmarshalBinary(b []byte) error {
	var res PostBackendsBackendIDGroupsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostBackendsBackendIDGroupsParamsBodyItems0 post backends backend ID groups params body items0
// swagger:model PostBackendsBackendIDGroupsParamsBodyItems0
type PostBackendsBackendIDGroupsParamsBodyItems0 struct {
	models.BackendGroupParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostBackendsBackendIDGroupsParamsBodyItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 models.BackendGroupParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	o.BackendGroupParams = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostBackendsBackendIDGroupsParamsBodyItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(o.BackendGroupParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post backends backend ID groups params body items0
func (o *PostBackendsBackendIDGroupsParamsBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.BackendGroupParams
	if err := o.BackendGroupParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostBackendsBackendIDGroupsParamsBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBackendsBackendIDGroupsParamsBodyItems0) UnmarshalBinary(b []byte) error {
	var res PostBackendsBackendIDGroupsParamsBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
