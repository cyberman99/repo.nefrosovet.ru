// Code generated by go-swagger; DO NOT EDIT.

package channels

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

// PostChannelsMtsSmsHandlerFunc turns a function with the right signature into a post channels mts sms handler
type PostChannelsMtsSmsHandlerFunc func(PostChannelsMtsSmsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostChannelsMtsSmsHandlerFunc) Handle(params PostChannelsMtsSmsParams) middleware.Responder {
	return fn(params)
}

// PostChannelsMtsSmsHandler interface for that can handle valid post channels mts sms params
type PostChannelsMtsSmsHandler interface {
	Handle(PostChannelsMtsSmsParams) middleware.Responder
}

// NewPostChannelsMtsSms creates a new http.Handler for the post channels mts sms operation
func NewPostChannelsMtsSms(ctx *middleware.Context, handler PostChannelsMtsSmsHandler) *PostChannelsMtsSms {
	return &PostChannelsMtsSms{Context: ctx, Handler: handler}
}

/*PostChannelsMtsSms swagger:route POST /channels/mts_sms Channels postChannelsMtsSms

Создание MTS SMS канала

*/
type PostChannelsMtsSms struct {
	Context *middleware.Context
	Handler PostChannelsMtsSmsHandler
}

func (o *PostChannelsMtsSms) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostChannelsMtsSmsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostChannelsMtsSmsBadRequestBody post channels mts sms bad request body
// swagger:model PostChannelsMtsSmsBadRequestBody
type PostChannelsMtsSmsBadRequestBody struct {
	models.Error400Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostChannelsMtsSmsBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PostChannelsMtsSmsBadRequestBodyAO0
	var postChannelsMtsSmsBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &postChannelsMtsSmsBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = postChannelsMtsSmsBadRequestBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostChannelsMtsSmsBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postChannelsMtsSmsBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postChannelsMtsSmsBadRequestBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post channels mts sms bad request body
func (o *PostChannelsMtsSmsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error400Data
	if err := o.Error400Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostChannelsMtsSmsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostChannelsMtsSmsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostChannelsMtsSmsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostChannelsMtsSmsBody post channels mts sms body
// swagger:model PostChannelsMtsSmsBody
type PostChannelsMtsSmsBody struct {
	models.ChannelParamsMtsSms

	PostChannelsMtsSmsParamsBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostChannelsMtsSmsBody) UnmarshalJSON(raw []byte) error {
	// PostChannelsMtsSmsParamsBodyAO0
	var postChannelsMtsSmsParamsBodyAO0 models.ChannelParamsMtsSms
	if err := swag.ReadJSON(raw, &postChannelsMtsSmsParamsBodyAO0); err != nil {
		return err
	}
	o.ChannelParamsMtsSms = postChannelsMtsSmsParamsBodyAO0

	// PostChannelsMtsSmsParamsBodyAO1
	var postChannelsMtsSmsParamsBodyAO1 PostChannelsMtsSmsParamsBodyAllOf1
	if err := swag.ReadJSON(raw, &postChannelsMtsSmsParamsBodyAO1); err != nil {
		return err
	}
	o.PostChannelsMtsSmsParamsBodyAllOf1 = postChannelsMtsSmsParamsBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostChannelsMtsSmsBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postChannelsMtsSmsParamsBodyAO0, err := swag.WriteJSON(o.ChannelParamsMtsSms)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postChannelsMtsSmsParamsBodyAO0)

	postChannelsMtsSmsParamsBodyAO1, err := swag.WriteJSON(o.PostChannelsMtsSmsParamsBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postChannelsMtsSmsParamsBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post channels mts sms body
func (o *PostChannelsMtsSmsBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ChannelParamsMtsSms
	if err := o.ChannelParamsMtsSms.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PostChannelsMtsSmsParamsBodyAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostChannelsMtsSmsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostChannelsMtsSmsBody) UnmarshalBinary(b []byte) error {
	var res PostChannelsMtsSmsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostChannelsMtsSmsForbiddenBody post channels mts sms forbidden body
// swagger:model PostChannelsMtsSmsForbiddenBody
type PostChannelsMtsSmsForbiddenBody struct {
	models.Error403Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostChannelsMtsSmsForbiddenBody) UnmarshalJSON(raw []byte) error {
	// PostChannelsMtsSmsForbiddenBodyAO0
	var postChannelsMtsSmsForbiddenBodyAO0 models.Error403Data
	if err := swag.ReadJSON(raw, &postChannelsMtsSmsForbiddenBodyAO0); err != nil {
		return err
	}
	o.Error403Data = postChannelsMtsSmsForbiddenBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostChannelsMtsSmsForbiddenBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postChannelsMtsSmsForbiddenBodyAO0, err := swag.WriteJSON(o.Error403Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postChannelsMtsSmsForbiddenBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post channels mts sms forbidden body
func (o *PostChannelsMtsSmsForbiddenBody) Validate(formats strfmt.Registry) error {
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
func (o *PostChannelsMtsSmsForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostChannelsMtsSmsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res PostChannelsMtsSmsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostChannelsMtsSmsMethodNotAllowedBody post channels mts sms method not allowed body
// swagger:model PostChannelsMtsSmsMethodNotAllowedBody
type PostChannelsMtsSmsMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostChannelsMtsSmsMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PostChannelsMtsSmsMethodNotAllowedBodyAO0
	var postChannelsMtsSmsMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &postChannelsMtsSmsMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = postChannelsMtsSmsMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostChannelsMtsSmsMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postChannelsMtsSmsMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postChannelsMtsSmsMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post channels mts sms method not allowed body
func (o *PostChannelsMtsSmsMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PostChannelsMtsSmsMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostChannelsMtsSmsMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PostChannelsMtsSmsMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostChannelsMtsSmsOKBody post channels mts sms o k body
// swagger:model PostChannelsMtsSmsOKBody
type PostChannelsMtsSmsOKBody struct {
	models.SuccessData

	// data
	Data []*DataItemMtsSms `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostChannelsMtsSmsOKBody) UnmarshalJSON(raw []byte) error {
	// PostChannelsMtsSmsOKBodyAO0
	var postChannelsMtsSmsOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postChannelsMtsSmsOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postChannelsMtsSmsOKBodyAO0

	// PostChannelsMtsSmsOKBodyAO1
	var dataPostChannelsMtsSmsOKBodyAO1 struct {
		Data []*DataItemMtsSms `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataPostChannelsMtsSmsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostChannelsMtsSmsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostChannelsMtsSmsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postChannelsMtsSmsOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postChannelsMtsSmsOKBodyAO0)

	var dataPostChannelsMtsSmsOKBodyAO1 struct {
		Data []*DataItemMtsSms `json:"data"`
	}

	dataPostChannelsMtsSmsOKBodyAO1.Data = o.Data

	jsonDataPostChannelsMtsSmsOKBodyAO1, errPostChannelsMtsSmsOKBodyAO1 := swag.WriteJSON(dataPostChannelsMtsSmsOKBodyAO1)
	if errPostChannelsMtsSmsOKBodyAO1 != nil {
		return nil, errPostChannelsMtsSmsOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostChannelsMtsSmsOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post channels mts sms o k body
func (o *PostChannelsMtsSmsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostChannelsMtsSmsOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("postChannelsMtsSmsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostChannelsMtsSmsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostChannelsMtsSmsOKBody) UnmarshalBinary(b []byte) error {
	var res PostChannelsMtsSmsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostChannelsMtsSmsParamsBodyAllOf1 post channels mts sms params body all of1
// swagger:model PostChannelsMtsSmsParamsBodyAllOf1
type PostChannelsMtsSmsParamsBodyAllOf1 interface{}

// data-item-mts-sms data item mts sms
// swagger:model data-item-mts-sms
type DataItemMtsSms struct {

	// Идентификатор канала
	ID string `json:"ID,omitempty"`

	// Тип канала
	Type string `json:"type,omitempty"`

	models.ChannelParamsMtsSms
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DataItemMtsSms) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ID string `json:"ID,omitempty"`

		Type string `json:"type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	o.ID = dataAO0.ID

	o.Type = dataAO0.Type

	// AO1
	var aO1 models.ChannelParamsMtsSms
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	o.ChannelParamsMtsSms = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DataItemMtsSms) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		ID string `json:"ID,omitempty"`

		Type string `json:"type,omitempty"`
	}

	dataAO0.ID = o.ID

	dataAO0.Type = o.Type

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(o.ChannelParamsMtsSms)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this data item mts sms
func (o *DataItemMtsSms) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ChannelParamsMtsSms
	if err := o.ChannelParamsMtsSms.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *DataItemMtsSms) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DataItemMtsSms) UnmarshalBinary(b []byte) error {
	var res DataItemMtsSms
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
