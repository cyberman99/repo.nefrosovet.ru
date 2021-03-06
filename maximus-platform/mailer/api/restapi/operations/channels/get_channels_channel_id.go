// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "repo.nefrosovet.ru/maximus-platform/mailer/api/models"
)

// GetChannelsChannelIDHandlerFunc turns a function with the right signature into a get channels channel ID handler
type GetChannelsChannelIDHandlerFunc func(GetChannelsChannelIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetChannelsChannelIDHandlerFunc) Handle(params GetChannelsChannelIDParams) middleware.Responder {
	return fn(params)
}

// GetChannelsChannelIDHandler interface for that can handle valid get channels channel ID params
type GetChannelsChannelIDHandler interface {
	Handle(GetChannelsChannelIDParams) middleware.Responder
}

// NewGetChannelsChannelID creates a new http.Handler for the get channels channel ID operation
func NewGetChannelsChannelID(ctx *middleware.Context, handler GetChannelsChannelIDHandler) *GetChannelsChannelID {
	return &GetChannelsChannelID{Context: ctx, Handler: handler}
}

/*GetChannelsChannelID swagger:route GET /channels/{channelID} Channels getChannelsChannelId

Информация о канале

*/
type GetChannelsChannelID struct {
	Context *middleware.Context
	Handler GetChannelsChannelIDHandler
}

func (o *GetChannelsChannelID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetChannelsChannelIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetChannelsChannelIDForbiddenBody get channels channel ID forbidden body
// swagger:model GetChannelsChannelIDForbiddenBody
type GetChannelsChannelIDForbiddenBody struct {
	models.Error403Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetChannelsChannelIDForbiddenBody) UnmarshalJSON(raw []byte) error {
	// GetChannelsChannelIDForbiddenBodyAO0
	var getChannelsChannelIDForbiddenBodyAO0 models.Error403Data
	if err := swag.ReadJSON(raw, &getChannelsChannelIDForbiddenBodyAO0); err != nil {
		return err
	}
	o.Error403Data = getChannelsChannelIDForbiddenBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetChannelsChannelIDForbiddenBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getChannelsChannelIDForbiddenBodyAO0, err := swag.WriteJSON(o.Error403Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getChannelsChannelIDForbiddenBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get channels channel ID forbidden body
func (o *GetChannelsChannelIDForbiddenBody) Validate(formats strfmt.Registry) error {
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
func (o *GetChannelsChannelIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetChannelsChannelIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetChannelsChannelIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetChannelsChannelIDMethodNotAllowedBody get channels channel ID method not allowed body
// swagger:model GetChannelsChannelIDMethodNotAllowedBody
type GetChannelsChannelIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetChannelsChannelIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// GetChannelsChannelIDMethodNotAllowedBodyAO0
	var getChannelsChannelIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &getChannelsChannelIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = getChannelsChannelIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetChannelsChannelIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getChannelsChannelIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getChannelsChannelIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get channels channel ID method not allowed body
func (o *GetChannelsChannelIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *GetChannelsChannelIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetChannelsChannelIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res GetChannelsChannelIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetChannelsChannelIDNotFoundBody get channels channel ID not found body
// swagger:model GetChannelsChannelIDNotFoundBody
type GetChannelsChannelIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetChannelsChannelIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// GetChannelsChannelIDNotFoundBodyAO0
	var getChannelsChannelIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &getChannelsChannelIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = getChannelsChannelIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetChannelsChannelIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getChannelsChannelIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getChannelsChannelIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get channels channel ID not found body
func (o *GetChannelsChannelIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *GetChannelsChannelIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetChannelsChannelIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetChannelsChannelIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetChannelsChannelIDOKBody get channels channel ID o k body
// swagger:model GetChannelsChannelIDOKBody
type GetChannelsChannelIDOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []interface{} `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetChannelsChannelIDOKBody) UnmarshalJSON(raw []byte) error {
	// GetChannelsChannelIDOKBodyAO0
	var getChannelsChannelIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &getChannelsChannelIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = getChannelsChannelIDOKBodyAO0

	// GetChannelsChannelIDOKBodyAO1
	var dataGetChannelsChannelIDOKBodyAO1 struct {
		Data []interface{} `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataGetChannelsChannelIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetChannelsChannelIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetChannelsChannelIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getChannelsChannelIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getChannelsChannelIDOKBodyAO0)

	var dataGetChannelsChannelIDOKBodyAO1 struct {
		Data []interface{} `json:"data"`
	}

	dataGetChannelsChannelIDOKBodyAO1.Data = o.Data

	jsonDataGetChannelsChannelIDOKBodyAO1, errGetChannelsChannelIDOKBodyAO1 := swag.WriteJSON(dataGetChannelsChannelIDOKBodyAO1)
	if errGetChannelsChannelIDOKBodyAO1 != nil {
		return nil, errGetChannelsChannelIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetChannelsChannelIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get channels channel ID o k body
func (o *GetChannelsChannelIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetChannelsChannelIDOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getChannelsChannelIdOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetChannelsChannelIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetChannelsChannelIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetChannelsChannelIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
