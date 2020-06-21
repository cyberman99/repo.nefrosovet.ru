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

// PatchBackendsLdapBackendIDHandlerFunc turns a function with the right signature into a patch backends ldap backend ID handler
type PatchBackendsLdapBackendIDHandlerFunc func(PatchBackendsLdapBackendIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchBackendsLdapBackendIDHandlerFunc) Handle(params PatchBackendsLdapBackendIDParams) middleware.Responder {
	return fn(params)
}

// PatchBackendsLdapBackendIDHandler interface for that can handle valid patch backends ldap backend ID params
type PatchBackendsLdapBackendIDHandler interface {
	Handle(PatchBackendsLdapBackendIDParams) middleware.Responder
}

// NewPatchBackendsLdapBackendID creates a new http.Handler for the patch backends ldap backend ID operation
func NewPatchBackendsLdapBackendID(ctx *middleware.Context, handler PatchBackendsLdapBackendIDHandler) *PatchBackendsLdapBackendID {
	return &PatchBackendsLdapBackendID{Context: ctx, Handler: handler}
}

/*PatchBackendsLdapBackendID swagger:route PATCH /backends/ldap/{backendID} Backend patchBackendsLdapBackendId

Изменение LDAP бэкенда

*/
type PatchBackendsLdapBackendID struct {
	Context *middleware.Context
	Handler PatchBackendsLdapBackendIDHandler
}

func (o *PatchBackendsLdapBackendID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchBackendsLdapBackendIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PatchBackendsLdapBackendIDBody patch backends ldap backend ID body
// swagger:model PatchBackendsLdapBackendIDBody
type PatchBackendsLdapBackendIDBody struct {
	models.BackendPatchLdapParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchBackendsLdapBackendIDBody) UnmarshalJSON(raw []byte) error {
	// PatchBackendsLdapBackendIDParamsBodyAO0
	var patchBackendsLdapBackendIDParamsBodyAO0 models.BackendPatchLdapParams
	if err := swag.ReadJSON(raw, &patchBackendsLdapBackendIDParamsBodyAO0); err != nil {
		return err
	}
	o.BackendPatchLdapParams = patchBackendsLdapBackendIDParamsBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchBackendsLdapBackendIDBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	patchBackendsLdapBackendIDParamsBodyAO0, err := swag.WriteJSON(o.BackendPatchLdapParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchBackendsLdapBackendIDParamsBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch backends ldap backend ID body
func (o *PatchBackendsLdapBackendIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.BackendPatchLdapParams
	if err := o.BackendPatchLdapParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PatchBackendsLdapBackendIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchBackendsLdapBackendIDBody) UnmarshalBinary(b []byte) error {
	var res PatchBackendsLdapBackendIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatchBackendsLdapBackendIDInternalServerErrorBody patch backends ldap backend ID internal server error body
// swagger:model PatchBackendsLdapBackendIDInternalServerErrorBody
type PatchBackendsLdapBackendIDInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchBackendsLdapBackendIDInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// PatchBackendsLdapBackendIDInternalServerErrorBodyAO0
	var patchBackendsLdapBackendIDInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &patchBackendsLdapBackendIDInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = patchBackendsLdapBackendIDInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchBackendsLdapBackendIDInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	patchBackendsLdapBackendIDInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchBackendsLdapBackendIDInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch backends ldap backend ID internal server error body
func (o *PatchBackendsLdapBackendIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *PatchBackendsLdapBackendIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchBackendsLdapBackendIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PatchBackendsLdapBackendIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatchBackendsLdapBackendIDMethodNotAllowedBody patch backends ldap backend ID method not allowed body
// swagger:model PatchBackendsLdapBackendIDMethodNotAllowedBody
type PatchBackendsLdapBackendIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchBackendsLdapBackendIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PatchBackendsLdapBackendIDMethodNotAllowedBodyAO0
	var patchBackendsLdapBackendIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &patchBackendsLdapBackendIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = patchBackendsLdapBackendIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchBackendsLdapBackendIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	patchBackendsLdapBackendIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchBackendsLdapBackendIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch backends ldap backend ID method not allowed body
func (o *PatchBackendsLdapBackendIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PatchBackendsLdapBackendIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchBackendsLdapBackendIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PatchBackendsLdapBackendIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatchBackendsLdapBackendIDNotFoundBody patch backends ldap backend ID not found body
// swagger:model PatchBackendsLdapBackendIDNotFoundBody
type PatchBackendsLdapBackendIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchBackendsLdapBackendIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// PatchBackendsLdapBackendIDNotFoundBodyAO0
	var patchBackendsLdapBackendIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &patchBackendsLdapBackendIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = patchBackendsLdapBackendIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchBackendsLdapBackendIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	patchBackendsLdapBackendIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchBackendsLdapBackendIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch backends ldap backend ID not found body
func (o *PatchBackendsLdapBackendIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *PatchBackendsLdapBackendIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchBackendsLdapBackendIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PatchBackendsLdapBackendIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatchBackendsLdapBackendIDOKBody patch backends ldap backend ID o k body
// swagger:model PatchBackendsLdapBackendIDOKBody
type PatchBackendsLdapBackendIDOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []*LdapDataItem `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchBackendsLdapBackendIDOKBody) UnmarshalJSON(raw []byte) error {
	// PatchBackendsLdapBackendIDOKBodyAO0
	var patchBackendsLdapBackendIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &patchBackendsLdapBackendIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = patchBackendsLdapBackendIDOKBodyAO0

	// PatchBackendsLdapBackendIDOKBodyAO1
	var dataPatchBackendsLdapBackendIDOKBodyAO1 struct {
		Data []*LdapDataItem `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataPatchBackendsLdapBackendIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPatchBackendsLdapBackendIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchBackendsLdapBackendIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	patchBackendsLdapBackendIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchBackendsLdapBackendIDOKBodyAO0)

	var dataPatchBackendsLdapBackendIDOKBodyAO1 struct {
		Data []*LdapDataItem `json:"data"`
	}

	dataPatchBackendsLdapBackendIDOKBodyAO1.Data = o.Data

	jsonDataPatchBackendsLdapBackendIDOKBodyAO1, errPatchBackendsLdapBackendIDOKBodyAO1 := swag.WriteJSON(dataPatchBackendsLdapBackendIDOKBodyAO1)
	if errPatchBackendsLdapBackendIDOKBodyAO1 != nil {
		return nil, errPatchBackendsLdapBackendIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataPatchBackendsLdapBackendIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch backends ldap backend ID o k body
func (o *PatchBackendsLdapBackendIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PatchBackendsLdapBackendIDOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("patchBackendsLdapBackendIdOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patchBackendsLdapBackendIdOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchBackendsLdapBackendIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchBackendsLdapBackendIDOKBody) UnmarshalBinary(b []byte) error {
	var res PatchBackendsLdapBackendIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ldap-data-item ldap data item
// swagger:model ldap-data-item
type LdapDataItem struct {
	models.BackendIDObject

	models.BackendLdapParams

	models.BackendParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *LdapDataItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 models.BackendIDObject
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	o.BackendIDObject = aO0

	// AO1
	var aO1 models.BackendLdapParams
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	o.BackendLdapParams = aO1

	// AO2
	var aO2 models.BackendParams
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	o.BackendParams = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o LdapDataItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(o.BackendIDObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(o.BackendLdapParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(o.BackendParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this ldap data item
func (o *LdapDataItem) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.BackendIDObject
	if err := o.BackendIDObject.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.BackendLdapParams
	if err := o.BackendLdapParams.Validate(formats); err != nil {
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
func (o *LdapDataItem) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *LdapDataItem) UnmarshalBinary(b []byte) error {
	var res LdapDataItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
