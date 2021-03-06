// Code generated by go-swagger; DO NOT EDIT.

package manage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "repo.nefrosovet.ru/maximus-platform/mailer/api/models"
)

// DeleteTokensTokenIDHandlerFunc turns a function with the right signature into a delete tokens token ID handler
type DeleteTokensTokenIDHandlerFunc func(DeleteTokensTokenIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTokensTokenIDHandlerFunc) Handle(params DeleteTokensTokenIDParams) middleware.Responder {
	return fn(params)
}

// DeleteTokensTokenIDHandler interface for that can handle valid delete tokens token ID params
type DeleteTokensTokenIDHandler interface {
	Handle(DeleteTokensTokenIDParams) middleware.Responder
}

// NewDeleteTokensTokenID creates a new http.Handler for the delete tokens token ID operation
func NewDeleteTokensTokenID(ctx *middleware.Context, handler DeleteTokensTokenIDHandler) *DeleteTokensTokenID {
	return &DeleteTokensTokenID{Context: ctx, Handler: handler}
}

/*DeleteTokensTokenID swagger:route DELETE /tokens/{tokenID} Manage deleteTokensTokenId

Удаление токена

*/
type DeleteTokensTokenID struct {
	Context *middleware.Context
	Handler DeleteTokensTokenIDHandler
}

func (o *DeleteTokensTokenID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteTokensTokenIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DataItems0 data items0
// swagger:model DataItems0
type DataItems0 struct {

	// Идентификатор токена
	ID string `json:"ID,omitempty"`

	models.TokenParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DataItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ID string `json:"ID,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	o.ID = dataAO0.ID

	// AO1
	var aO1 models.TokenParams
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	o.TokenParams = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DataItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		ID string `json:"ID,omitempty"`
	}

	dataAO0.ID = o.ID

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(o.TokenParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this data items0
func (o *DataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.TokenParams
	if err := o.TokenParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *DataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DataItems0) UnmarshalBinary(b []byte) error {
	var res DataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteTokensTokenIDForbiddenBody delete tokens token ID forbidden body
// swagger:model DeleteTokensTokenIDForbiddenBody
type DeleteTokensTokenIDForbiddenBody struct {
	models.Error403Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteTokensTokenIDForbiddenBody) UnmarshalJSON(raw []byte) error {
	// DeleteTokensTokenIDForbiddenBodyAO0
	var deleteTokensTokenIDForbiddenBodyAO0 models.Error403Data
	if err := swag.ReadJSON(raw, &deleteTokensTokenIDForbiddenBodyAO0); err != nil {
		return err
	}
	o.Error403Data = deleteTokensTokenIDForbiddenBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteTokensTokenIDForbiddenBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	deleteTokensTokenIDForbiddenBodyAO0, err := swag.WriteJSON(o.Error403Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteTokensTokenIDForbiddenBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete tokens token ID forbidden body
func (o *DeleteTokensTokenIDForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error403Data
	if err := o.Error403Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteTokensTokenIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteTokensTokenIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res DeleteTokensTokenIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteTokensTokenIDMethodNotAllowedBody delete tokens token ID method not allowed body
// swagger:model DeleteTokensTokenIDMethodNotAllowedBody
type DeleteTokensTokenIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteTokensTokenIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// DeleteTokensTokenIDMethodNotAllowedBodyAO0
	var deleteTokensTokenIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &deleteTokensTokenIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = deleteTokensTokenIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteTokensTokenIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	deleteTokensTokenIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteTokensTokenIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete tokens token ID method not allowed body
func (o *DeleteTokensTokenIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *DeleteTokensTokenIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteTokensTokenIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res DeleteTokensTokenIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteTokensTokenIDNotFoundBody delete tokens token ID not found body
// swagger:model DeleteTokensTokenIDNotFoundBody
type DeleteTokensTokenIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteTokensTokenIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// DeleteTokensTokenIDNotFoundBodyAO0
	var deleteTokensTokenIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &deleteTokensTokenIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = deleteTokensTokenIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteTokensTokenIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	deleteTokensTokenIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteTokensTokenIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete tokens token ID not found body
func (o *DeleteTokensTokenIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *DeleteTokensTokenIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteTokensTokenIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteTokensTokenIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteTokensTokenIDOKBody delete tokens token ID o k body
// swagger:model DeleteTokensTokenIDOKBody
type DeleteTokensTokenIDOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteTokensTokenIDOKBody) UnmarshalJSON(raw []byte) error {
	// DeleteTokensTokenIDOKBodyAO0
	var deleteTokensTokenIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &deleteTokensTokenIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = deleteTokensTokenIDOKBodyAO0

	// DeleteTokensTokenIDOKBodyAO1
	var dataDeleteTokensTokenIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataDeleteTokensTokenIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataDeleteTokensTokenIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteTokensTokenIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	deleteTokensTokenIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteTokensTokenIDOKBodyAO0)

	var dataDeleteTokensTokenIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataDeleteTokensTokenIDOKBodyAO1.Data = o.Data

	jsonDataDeleteTokensTokenIDOKBodyAO1, errDeleteTokensTokenIDOKBodyAO1 := swag.WriteJSON(dataDeleteTokensTokenIDOKBodyAO1)
	if errDeleteTokensTokenIDOKBodyAO1 != nil {
		return nil, errDeleteTokensTokenIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataDeleteTokensTokenIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete tokens token ID o k body
func (o *DeleteTokensTokenIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DeleteTokensTokenIDOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("deleteTokensTokenIdOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteTokensTokenIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteTokensTokenIDOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteTokensTokenIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
