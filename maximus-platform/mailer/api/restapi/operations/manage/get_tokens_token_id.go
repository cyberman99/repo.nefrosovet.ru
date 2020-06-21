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

// GetTokensTokenIDHandlerFunc turns a function with the right signature into a get tokens token ID handler
type GetTokensTokenIDHandlerFunc func(GetTokensTokenIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTokensTokenIDHandlerFunc) Handle(params GetTokensTokenIDParams) middleware.Responder {
	return fn(params)
}

// GetTokensTokenIDHandler interface for that can handle valid get tokens token ID params
type GetTokensTokenIDHandler interface {
	Handle(GetTokensTokenIDParams) middleware.Responder
}

// NewGetTokensTokenID creates a new http.Handler for the get tokens token ID operation
func NewGetTokensTokenID(ctx *middleware.Context, handler GetTokensTokenIDHandler) *GetTokensTokenID {
	return &GetTokensTokenID{Context: ctx, Handler: handler}
}

/*GetTokensTokenID swagger:route GET /tokens/{tokenID} Manage getTokensTokenId

Информация о токене

*/
type GetTokensTokenID struct {
	Context *middleware.Context
	Handler GetTokensTokenIDHandler
}

func (o *GetTokensTokenID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTokensTokenIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetTokensTokenIDForbiddenBody get tokens token ID forbidden body
// swagger:model GetTokensTokenIDForbiddenBody
type GetTokensTokenIDForbiddenBody struct {
	models.Error403Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetTokensTokenIDForbiddenBody) UnmarshalJSON(raw []byte) error {
	// GetTokensTokenIDForbiddenBodyAO0
	var getTokensTokenIDForbiddenBodyAO0 models.Error403Data
	if err := swag.ReadJSON(raw, &getTokensTokenIDForbiddenBodyAO0); err != nil {
		return err
	}
	o.Error403Data = getTokensTokenIDForbiddenBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetTokensTokenIDForbiddenBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getTokensTokenIDForbiddenBodyAO0, err := swag.WriteJSON(o.Error403Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getTokensTokenIDForbiddenBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get tokens token ID forbidden body
func (o *GetTokensTokenIDForbiddenBody) Validate(formats strfmt.Registry) error {
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
func (o *GetTokensTokenIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTokensTokenIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetTokensTokenIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetTokensTokenIDMethodNotAllowedBody get tokens token ID method not allowed body
// swagger:model GetTokensTokenIDMethodNotAllowedBody
type GetTokensTokenIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetTokensTokenIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// GetTokensTokenIDMethodNotAllowedBodyAO0
	var getTokensTokenIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &getTokensTokenIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = getTokensTokenIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetTokensTokenIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getTokensTokenIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getTokensTokenIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get tokens token ID method not allowed body
func (o *GetTokensTokenIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *GetTokensTokenIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTokensTokenIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res GetTokensTokenIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetTokensTokenIDNotFoundBody get tokens token ID not found body
// swagger:model GetTokensTokenIDNotFoundBody
type GetTokensTokenIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetTokensTokenIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// GetTokensTokenIDNotFoundBodyAO0
	var getTokensTokenIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &getTokensTokenIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = getTokensTokenIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetTokensTokenIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getTokensTokenIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getTokensTokenIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get tokens token ID not found body
func (o *GetTokensTokenIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *GetTokensTokenIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTokensTokenIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetTokensTokenIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetTokensTokenIDOKBody get tokens token ID o k body
// swagger:model GetTokensTokenIDOKBody
type GetTokensTokenIDOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetTokensTokenIDOKBody) UnmarshalJSON(raw []byte) error {
	// GetTokensTokenIDOKBodyAO0
	var getTokensTokenIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &getTokensTokenIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = getTokensTokenIDOKBodyAO0

	// GetTokensTokenIDOKBodyAO1
	var dataGetTokensTokenIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataGetTokensTokenIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetTokensTokenIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetTokensTokenIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getTokensTokenIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getTokensTokenIDOKBodyAO0)

	var dataGetTokensTokenIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataGetTokensTokenIDOKBodyAO1.Data = o.Data

	jsonDataGetTokensTokenIDOKBodyAO1, errGetTokensTokenIDOKBodyAO1 := swag.WriteJSON(dataGetTokensTokenIDOKBodyAO1)
	if errGetTokensTokenIDOKBodyAO1 != nil {
		return nil, errGetTokensTokenIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetTokensTokenIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get tokens token ID o k body
func (o *GetTokensTokenIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetTokensTokenIDOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("getTokensTokenIdOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTokensTokenIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTokensTokenIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetTokensTokenIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
