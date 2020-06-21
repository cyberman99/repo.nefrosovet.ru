// Code generated by go-swagger; DO NOT EDIT.

package messages

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

// GetMessagesMessageIDHandlerFunc turns a function with the right signature into a get messages message ID handler
type GetMessagesMessageIDHandlerFunc func(GetMessagesMessageIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMessagesMessageIDHandlerFunc) Handle(params GetMessagesMessageIDParams) middleware.Responder {
	return fn(params)
}

// GetMessagesMessageIDHandler interface for that can handle valid get messages message ID params
type GetMessagesMessageIDHandler interface {
	Handle(GetMessagesMessageIDParams) middleware.Responder
}

// NewGetMessagesMessageID creates a new http.Handler for the get messages message ID operation
func NewGetMessagesMessageID(ctx *middleware.Context, handler GetMessagesMessageIDHandler) *GetMessagesMessageID {
	return &GetMessagesMessageID{Context: ctx, Handler: handler}
}

/*GetMessagesMessageID swagger:route GET /messages/{messageID} Messages getMessagesMessageId

Информация о сообщении

*/
type GetMessagesMessageID struct {
	Context *middleware.Context
	Handler GetMessagesMessageIDHandler
}

func (o *GetMessagesMessageID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMessagesMessageIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetMessagesMessageIDForbiddenBody get messages message ID forbidden body
// swagger:model GetMessagesMessageIDForbiddenBody
type GetMessagesMessageIDForbiddenBody struct {
	models.Error403Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetMessagesMessageIDForbiddenBody) UnmarshalJSON(raw []byte) error {
	// GetMessagesMessageIDForbiddenBodyAO0
	var getMessagesMessageIDForbiddenBodyAO0 models.Error403Data
	if err := swag.ReadJSON(raw, &getMessagesMessageIDForbiddenBodyAO0); err != nil {
		return err
	}
	o.Error403Data = getMessagesMessageIDForbiddenBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetMessagesMessageIDForbiddenBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getMessagesMessageIDForbiddenBodyAO0, err := swag.WriteJSON(o.Error403Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getMessagesMessageIDForbiddenBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get messages message ID forbidden body
func (o *GetMessagesMessageIDForbiddenBody) Validate(formats strfmt.Registry) error {
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
func (o *GetMessagesMessageIDForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMessagesMessageIDForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetMessagesMessageIDForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetMessagesMessageIDMethodNotAllowedBody get messages message ID method not allowed body
// swagger:model GetMessagesMessageIDMethodNotAllowedBody
type GetMessagesMessageIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetMessagesMessageIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// GetMessagesMessageIDMethodNotAllowedBodyAO0
	var getMessagesMessageIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &getMessagesMessageIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = getMessagesMessageIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetMessagesMessageIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getMessagesMessageIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getMessagesMessageIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get messages message ID method not allowed body
func (o *GetMessagesMessageIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *GetMessagesMessageIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMessagesMessageIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res GetMessagesMessageIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetMessagesMessageIDNotFoundBody get messages message ID not found body
// swagger:model GetMessagesMessageIDNotFoundBody
type GetMessagesMessageIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetMessagesMessageIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// GetMessagesMessageIDNotFoundBodyAO0
	var getMessagesMessageIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &getMessagesMessageIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = getMessagesMessageIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetMessagesMessageIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getMessagesMessageIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getMessagesMessageIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get messages message ID not found body
func (o *GetMessagesMessageIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *GetMessagesMessageIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMessagesMessageIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetMessagesMessageIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetMessagesMessageIDOKBody get messages message ID o k body
// swagger:model GetMessagesMessageIDOKBody
type GetMessagesMessageIDOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetMessagesMessageIDOKBody) UnmarshalJSON(raw []byte) error {
	// GetMessagesMessageIDOKBodyAO0
	var getMessagesMessageIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &getMessagesMessageIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = getMessagesMessageIDOKBodyAO0

	// GetMessagesMessageIDOKBodyAO1
	var dataGetMessagesMessageIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataGetMessagesMessageIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetMessagesMessageIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetMessagesMessageIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getMessagesMessageIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getMessagesMessageIDOKBodyAO0)

	var dataGetMessagesMessageIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataGetMessagesMessageIDOKBodyAO1.Data = o.Data

	jsonDataGetMessagesMessageIDOKBodyAO1, errGetMessagesMessageIDOKBodyAO1 := swag.WriteJSON(dataGetMessagesMessageIDOKBodyAO1)
	if errGetMessagesMessageIDOKBodyAO1 != nil {
		return nil, errGetMessagesMessageIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetMessagesMessageIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get messages message ID o k body
func (o *GetMessagesMessageIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetMessagesMessageIDOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("getMessagesMessageIdOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMessagesMessageIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMessagesMessageIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetMessagesMessageIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
