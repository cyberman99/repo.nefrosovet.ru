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

// PatchBackendsOauth2BackendIDHandlerFunc turns a function with the right signature into a patch backends oauth2 backend ID handler
type PatchBackendsOauth2BackendIDHandlerFunc func(PatchBackendsOauth2BackendIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchBackendsOauth2BackendIDHandlerFunc) Handle(params PatchBackendsOauth2BackendIDParams) middleware.Responder {
	return fn(params)
}

// PatchBackendsOauth2BackendIDHandler interface for that can handle valid patch backends oauth2 backend ID params
type PatchBackendsOauth2BackendIDHandler interface {
	Handle(PatchBackendsOauth2BackendIDParams) middleware.Responder
}

// NewPatchBackendsOauth2BackendID creates a new http.Handler for the patch backends oauth2 backend ID operation
func NewPatchBackendsOauth2BackendID(ctx *middleware.Context, handler PatchBackendsOauth2BackendIDHandler) *PatchBackendsOauth2BackendID {
	return &PatchBackendsOauth2BackendID{Context: ctx, Handler: handler}
}

/*PatchBackendsOauth2BackendID swagger:route PATCH /backends/oauth2/{backendID} Backend patchBackendsOauth2BackendId

Изменение oauth2 бэкенда

*/
type PatchBackendsOauth2BackendID struct {
	Context *middleware.Context
	Handler PatchBackendsOauth2BackendIDHandler
}

func (o *PatchBackendsOauth2BackendID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchBackendsOauth2BackendIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PatchBackendsOauth2BackendIDBody patch backends oauth2 backend ID body
// swagger:model PatchBackendsOauth2BackendIDBody
type PatchBackendsOauth2BackendIDBody struct {
	models.BackendPatchOauth2Params

	models.BackendOauth2IDParam
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchBackendsOauth2BackendIDBody) UnmarshalJSON(raw []byte) error {
	// PatchBackendsOauth2BackendIDParamsBodyAO0
	var patchBackendsOauth2BackendIDParamsBodyAO0 models.BackendPatchOauth2Params
	if err := swag.ReadJSON(raw, &patchBackendsOauth2BackendIDParamsBodyAO0); err != nil {
		return err
	}
	o.BackendPatchOauth2Params = patchBackendsOauth2BackendIDParamsBodyAO0

	// PatchBackendsOauth2BackendIDParamsBodyAO1
	var patchBackendsOauth2BackendIDParamsBodyAO1 models.BackendOauth2IDParam
	if err := swag.ReadJSON(raw, &patchBackendsOauth2BackendIDParamsBodyAO1); err != nil {
		return err
	}
	o.BackendOauth2IDParam = patchBackendsOauth2BackendIDParamsBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchBackendsOauth2BackendIDBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	patchBackendsOauth2BackendIDParamsBodyAO0, err := swag.WriteJSON(o.BackendPatchOauth2Params)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchBackendsOauth2BackendIDParamsBodyAO0)

	patchBackendsOauth2BackendIDParamsBodyAO1, err := swag.WriteJSON(o.BackendOauth2IDParam)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchBackendsOauth2BackendIDParamsBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch backends oauth2 backend ID body
func (o *PatchBackendsOauth2BackendIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.BackendPatchOauth2Params
	if err := o.BackendPatchOauth2Params.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.BackendOauth2IDParam
	if err := o.BackendOauth2IDParam.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PatchBackendsOauth2BackendIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchBackendsOauth2BackendIDBody) UnmarshalBinary(b []byte) error {
	var res PatchBackendsOauth2BackendIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatchBackendsOauth2BackendIDInternalServerErrorBody patch backends oauth2 backend ID internal server error body
// swagger:model PatchBackendsOauth2BackendIDInternalServerErrorBody
type PatchBackendsOauth2BackendIDInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchBackendsOauth2BackendIDInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// PatchBackendsOauth2BackendIDInternalServerErrorBodyAO0
	var patchBackendsOauth2BackendIDInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &patchBackendsOauth2BackendIDInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = patchBackendsOauth2BackendIDInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchBackendsOauth2BackendIDInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	patchBackendsOauth2BackendIDInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchBackendsOauth2BackendIDInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch backends oauth2 backend ID internal server error body
func (o *PatchBackendsOauth2BackendIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *PatchBackendsOauth2BackendIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchBackendsOauth2BackendIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PatchBackendsOauth2BackendIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatchBackendsOauth2BackendIDMethodNotAllowedBody patch backends oauth2 backend ID method not allowed body
// swagger:model PatchBackendsOauth2BackendIDMethodNotAllowedBody
type PatchBackendsOauth2BackendIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchBackendsOauth2BackendIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PatchBackendsOauth2BackendIDMethodNotAllowedBodyAO0
	var patchBackendsOauth2BackendIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &patchBackendsOauth2BackendIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = patchBackendsOauth2BackendIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchBackendsOauth2BackendIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	patchBackendsOauth2BackendIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchBackendsOauth2BackendIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch backends oauth2 backend ID method not allowed body
func (o *PatchBackendsOauth2BackendIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PatchBackendsOauth2BackendIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchBackendsOauth2BackendIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PatchBackendsOauth2BackendIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatchBackendsOauth2BackendIDNotFoundBody patch backends oauth2 backend ID not found body
// swagger:model PatchBackendsOauth2BackendIDNotFoundBody
type PatchBackendsOauth2BackendIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchBackendsOauth2BackendIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// PatchBackendsOauth2BackendIDNotFoundBodyAO0
	var patchBackendsOauth2BackendIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &patchBackendsOauth2BackendIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = patchBackendsOauth2BackendIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchBackendsOauth2BackendIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	patchBackendsOauth2BackendIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchBackendsOauth2BackendIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch backends oauth2 backend ID not found body
func (o *PatchBackendsOauth2BackendIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *PatchBackendsOauth2BackendIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchBackendsOauth2BackendIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PatchBackendsOauth2BackendIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatchBackendsOauth2BackendIDOKBody patch backends oauth2 backend ID o k body
// swagger:model PatchBackendsOauth2BackendIDOKBody
type PatchBackendsOauth2BackendIDOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []*Oauth2DataItem `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchBackendsOauth2BackendIDOKBody) UnmarshalJSON(raw []byte) error {
	// PatchBackendsOauth2BackendIDOKBodyAO0
	var patchBackendsOauth2BackendIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &patchBackendsOauth2BackendIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = patchBackendsOauth2BackendIDOKBodyAO0

	// PatchBackendsOauth2BackendIDOKBodyAO1
	var dataPatchBackendsOauth2BackendIDOKBodyAO1 struct {
		Data []*Oauth2DataItem `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataPatchBackendsOauth2BackendIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPatchBackendsOauth2BackendIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchBackendsOauth2BackendIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	patchBackendsOauth2BackendIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchBackendsOauth2BackendIDOKBodyAO0)

	var dataPatchBackendsOauth2BackendIDOKBodyAO1 struct {
		Data []*Oauth2DataItem `json:"data"`
	}

	dataPatchBackendsOauth2BackendIDOKBodyAO1.Data = o.Data

	jsonDataPatchBackendsOauth2BackendIDOKBodyAO1, errPatchBackendsOauth2BackendIDOKBodyAO1 := swag.WriteJSON(dataPatchBackendsOauth2BackendIDOKBodyAO1)
	if errPatchBackendsOauth2BackendIDOKBodyAO1 != nil {
		return nil, errPatchBackendsOauth2BackendIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataPatchBackendsOauth2BackendIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch backends oauth2 backend ID o k body
func (o *PatchBackendsOauth2BackendIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PatchBackendsOauth2BackendIDOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("patchBackendsOauth2BackendIdOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patchBackendsOauth2BackendIdOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchBackendsOauth2BackendIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchBackendsOauth2BackendIDOKBody) UnmarshalBinary(b []byte) error {
	var res PatchBackendsOauth2BackendIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// oauth2-data-item oauth2 data item
// swagger:model oauth2-data-item
type Oauth2DataItem struct {
	models.BackendIDObject

	models.BackendOauth2Params

	models.BackendParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *Oauth2DataItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 models.BackendIDObject
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	o.BackendIDObject = aO0

	// AO1
	var aO1 models.BackendOauth2Params
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	o.BackendOauth2Params = aO1

	// AO2
	var aO2 models.BackendParams
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	o.BackendParams = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o Oauth2DataItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(o.BackendIDObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(o.BackendOauth2Params)
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

// Validate validates this oauth2 data item
func (o *Oauth2DataItem) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.BackendIDObject
	if err := o.BackendIDObject.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.BackendOauth2Params
	if err := o.BackendOauth2Params.Validate(formats); err != nil {
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
func (o *Oauth2DataItem) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *Oauth2DataItem) UnmarshalBinary(b []byte) error {
	var res Oauth2DataItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
