// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "repo.nefrosovet.ru/maximus-platform/auth/api/models"
)

// GetBackendsBackendIDHandlerFunc turns a function with the right signature into a get backends backend ID handler
type GetBackendsBackendIDHandlerFunc func(GetBackendsBackendIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBackendsBackendIDHandlerFunc) Handle(params GetBackendsBackendIDParams) middleware.Responder {
	return fn(params)
}

// GetBackendsBackendIDHandler interface for that can handle valid get backends backend ID params
type GetBackendsBackendIDHandler interface {
	Handle(GetBackendsBackendIDParams) middleware.Responder
}

// NewGetBackendsBackendID creates a new http.Handler for the get backends backend ID operation
func NewGetBackendsBackendID(ctx *middleware.Context, handler GetBackendsBackendIDHandler) *GetBackendsBackendID {
	return &GetBackendsBackendID{Context: ctx, Handler: handler}
}

/*GetBackendsBackendID swagger:route GET /backends/{backendID} Backend getBackendsBackendId

Информация о бэкенде

*/
type GetBackendsBackendID struct {
	Context *middleware.Context
	Handler GetBackendsBackendIDHandler
}

func (o *GetBackendsBackendID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBackendsBackendIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetBackendsBackendIDInternalServerErrorBody get backends backend ID internal server error body
// swagger:model GetBackendsBackendIDInternalServerErrorBody
type GetBackendsBackendIDInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetBackendsBackendIDInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// GetBackendsBackendIDInternalServerErrorBodyAO0
	var getBackendsBackendIDInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &getBackendsBackendIDInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = getBackendsBackendIDInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetBackendsBackendIDInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getBackendsBackendIDInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getBackendsBackendIDInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get backends backend ID internal server error body
func (o *GetBackendsBackendIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *GetBackendsBackendIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackendsBackendIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetBackendsBackendIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetBackendsBackendIDMethodNotAllowedBody get backends backend ID method not allowed body
// swagger:model GetBackendsBackendIDMethodNotAllowedBody
type GetBackendsBackendIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetBackendsBackendIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// GetBackendsBackendIDMethodNotAllowedBodyAO0
	var getBackendsBackendIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &getBackendsBackendIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = getBackendsBackendIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetBackendsBackendIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getBackendsBackendIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getBackendsBackendIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get backends backend ID method not allowed body
func (o *GetBackendsBackendIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *GetBackendsBackendIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackendsBackendIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res GetBackendsBackendIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetBackendsBackendIDNotFoundBody get backends backend ID not found body
// swagger:model GetBackendsBackendIDNotFoundBody
type GetBackendsBackendIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetBackendsBackendIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// GetBackendsBackendIDNotFoundBodyAO0
	var getBackendsBackendIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &getBackendsBackendIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = getBackendsBackendIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetBackendsBackendIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getBackendsBackendIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getBackendsBackendIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get backends backend ID not found body
func (o *GetBackendsBackendIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *GetBackendsBackendIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackendsBackendIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetBackendsBackendIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetBackendsBackendIDOKBody get backends backend ID o k body
// swagger:model GetBackendsBackendIDOKBody
type GetBackendsBackendIDOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []interface{} `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetBackendsBackendIDOKBody) UnmarshalJSON(raw []byte) error {
	// GetBackendsBackendIDOKBodyAO0
	var getBackendsBackendIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &getBackendsBackendIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = getBackendsBackendIDOKBodyAO0

	// GetBackendsBackendIDOKBodyAO1
	var dataGetBackendsBackendIDOKBodyAO1 struct {
		Data []interface{} `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataGetBackendsBackendIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetBackendsBackendIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetBackendsBackendIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getBackendsBackendIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getBackendsBackendIDOKBodyAO0)

	var dataGetBackendsBackendIDOKBodyAO1 struct {
		Data []interface{} `json:"data"`
	}

	dataGetBackendsBackendIDOKBodyAO1.Data = o.Data

	jsonDataGetBackendsBackendIDOKBodyAO1, errGetBackendsBackendIDOKBodyAO1 := swag.WriteJSON(dataGetBackendsBackendIDOKBodyAO1)
	if errGetBackendsBackendIDOKBodyAO1 != nil {
		return nil, errGetBackendsBackendIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetBackendsBackendIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get backends backend ID o k body
func (o *GetBackendsBackendIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetBackendsBackendIDOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getBackendsBackendIdOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBackendsBackendIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackendsBackendIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetBackendsBackendIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
