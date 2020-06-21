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

// PostBackendsLdapHandlerFunc turns a function with the right signature into a post backends ldap handler
type PostBackendsLdapHandlerFunc func(PostBackendsLdapParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostBackendsLdapHandlerFunc) Handle(params PostBackendsLdapParams) middleware.Responder {
	return fn(params)
}

// PostBackendsLdapHandler interface for that can handle valid post backends ldap params
type PostBackendsLdapHandler interface {
	Handle(PostBackendsLdapParams) middleware.Responder
}

// NewPostBackendsLdap creates a new http.Handler for the post backends ldap operation
func NewPostBackendsLdap(ctx *middleware.Context, handler PostBackendsLdapHandler) *PostBackendsLdap {
	return &PostBackendsLdap{Context: ctx, Handler: handler}
}

/*PostBackendsLdap swagger:route POST /backends/ldap Backend postBackendsLdap

Создание LDAP бэкенда

*/
type PostBackendsLdap struct {
	Context *middleware.Context
	Handler PostBackendsLdapHandler
}

func (o *PostBackendsLdap) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostBackendsLdapParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostBackendsLdapBody post backends ldap body
// swagger:model PostBackendsLdapBody
type PostBackendsLdapBody struct {
	models.BackendLdapParams

	models.PasswordObject

	models.BackendParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostBackendsLdapBody) UnmarshalJSON(raw []byte) error {
	// PostBackendsLdapParamsBodyAO0
	var postBackendsLdapParamsBodyAO0 models.BackendLdapParams
	if err := swag.ReadJSON(raw, &postBackendsLdapParamsBodyAO0); err != nil {
		return err
	}
	o.BackendLdapParams = postBackendsLdapParamsBodyAO0

	// PostBackendsLdapParamsBodyAO1
	var postBackendsLdapParamsBodyAO1 models.PasswordObject
	if err := swag.ReadJSON(raw, &postBackendsLdapParamsBodyAO1); err != nil {
		return err
	}
	o.PasswordObject = postBackendsLdapParamsBodyAO1

	// PostBackendsLdapParamsBodyAO2
	var postBackendsLdapParamsBodyAO2 models.BackendParams
	if err := swag.ReadJSON(raw, &postBackendsLdapParamsBodyAO2); err != nil {
		return err
	}
	o.BackendParams = postBackendsLdapParamsBodyAO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostBackendsLdapBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	postBackendsLdapParamsBodyAO0, err := swag.WriteJSON(o.BackendLdapParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postBackendsLdapParamsBodyAO0)

	postBackendsLdapParamsBodyAO1, err := swag.WriteJSON(o.PasswordObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postBackendsLdapParamsBodyAO1)

	postBackendsLdapParamsBodyAO2, err := swag.WriteJSON(o.BackendParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postBackendsLdapParamsBodyAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post backends ldap body
func (o *PostBackendsLdapBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.BackendLdapParams
	if err := o.BackendLdapParams.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.PasswordObject
	if err := o.PasswordObject.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.BackendParams
	if err := o.BackendParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostBackendsLdapBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBackendsLdapBody) UnmarshalBinary(b []byte) error {
	var res PostBackendsLdapBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostBackendsLdapInternalServerErrorBody post backends ldap internal server error body
// swagger:model PostBackendsLdapInternalServerErrorBody
type PostBackendsLdapInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostBackendsLdapInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// PostBackendsLdapInternalServerErrorBodyAO0
	var postBackendsLdapInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &postBackendsLdapInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = postBackendsLdapInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostBackendsLdapInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postBackendsLdapInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postBackendsLdapInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post backends ldap internal server error body
func (o *PostBackendsLdapInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *PostBackendsLdapInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBackendsLdapInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostBackendsLdapInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostBackendsLdapMethodNotAllowedBody post backends ldap method not allowed body
// swagger:model PostBackendsLdapMethodNotAllowedBody
type PostBackendsLdapMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostBackendsLdapMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PostBackendsLdapMethodNotAllowedBodyAO0
	var postBackendsLdapMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &postBackendsLdapMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = postBackendsLdapMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostBackendsLdapMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postBackendsLdapMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postBackendsLdapMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post backends ldap method not allowed body
func (o *PostBackendsLdapMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PostBackendsLdapMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBackendsLdapMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PostBackendsLdapMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostBackendsLdapOKBody post backends ldap o k body
// swagger:model PostBackendsLdapOKBody
type PostBackendsLdapOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []*LdapDataItem `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostBackendsLdapOKBody) UnmarshalJSON(raw []byte) error {
	// PostBackendsLdapOKBodyAO0
	var postBackendsLdapOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postBackendsLdapOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postBackendsLdapOKBodyAO0

	// PostBackendsLdapOKBodyAO1
	var dataPostBackendsLdapOKBodyAO1 struct {
		Data []*LdapDataItem `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataPostBackendsLdapOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostBackendsLdapOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostBackendsLdapOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postBackendsLdapOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postBackendsLdapOKBodyAO0)

	var dataPostBackendsLdapOKBodyAO1 struct {
		Data []*LdapDataItem `json:"data"`
	}

	dataPostBackendsLdapOKBodyAO1.Data = o.Data

	jsonDataPostBackendsLdapOKBodyAO1, errPostBackendsLdapOKBodyAO1 := swag.WriteJSON(dataPostBackendsLdapOKBodyAO1)
	if errPostBackendsLdapOKBodyAO1 != nil {
		return nil, errPostBackendsLdapOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostBackendsLdapOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post backends ldap o k body
func (o *PostBackendsLdapOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostBackendsLdapOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("postBackendsLdapOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postBackendsLdapOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostBackendsLdapOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBackendsLdapOKBody) UnmarshalBinary(b []byte) error {
	var res PostBackendsLdapOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}