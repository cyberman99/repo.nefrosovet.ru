// Code generated by go-swagger; DO NOT EDIT.

package replies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "repo.nefrosovet.ru/maximus-platform/DataRouter/api/models"
)

// ReplyDeleteHandlerFunc turns a function with the right signature into a reply delete handler
type ReplyDeleteHandlerFunc func(ReplyDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplyDeleteHandlerFunc) Handle(params ReplyDeleteParams) middleware.Responder {
	return fn(params)
}

// ReplyDeleteHandler interface for that can handle valid reply delete params
type ReplyDeleteHandler interface {
	Handle(ReplyDeleteParams) middleware.Responder
}

// NewReplyDelete creates a new http.Handler for the reply delete operation
func NewReplyDelete(ctx *middleware.Context, handler ReplyDeleteHandler) *ReplyDelete {
	return &ReplyDelete{Context: ctx, Handler: handler}
}

/*ReplyDelete swagger:route DELETE /replies/{replyID} Replies replyDelete

Удаление шаблона ответа

*/
type ReplyDelete struct {
	Context *middleware.Context
	Handler ReplyDeleteHandler
}

func (o *ReplyDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplyDeleteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ReplyDeleteInternalServerErrorBody reply delete internal server error body
// swagger:model ReplyDeleteInternalServerErrorBody
type ReplyDeleteInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ReplyDeleteInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// ReplyDeleteInternalServerErrorBodyAO0
	var replyDeleteInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &replyDeleteInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = replyDeleteInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ReplyDeleteInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	replyDeleteInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, replyDeleteInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this reply delete internal server error body
func (o *ReplyDeleteInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *ReplyDeleteInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplyDeleteInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ReplyDeleteInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ReplyDeleteMethodNotAllowedBody reply delete method not allowed body
// swagger:model ReplyDeleteMethodNotAllowedBody
type ReplyDeleteMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ReplyDeleteMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// ReplyDeleteMethodNotAllowedBodyAO0
	var replyDeleteMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &replyDeleteMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = replyDeleteMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ReplyDeleteMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	replyDeleteMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, replyDeleteMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this reply delete method not allowed body
func (o *ReplyDeleteMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *ReplyDeleteMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplyDeleteMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res ReplyDeleteMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ReplyDeleteNotFoundBody reply delete not found body
// swagger:model ReplyDeleteNotFoundBody
type ReplyDeleteNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ReplyDeleteNotFoundBody) UnmarshalJSON(raw []byte) error {
	// ReplyDeleteNotFoundBodyAO0
	var replyDeleteNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &replyDeleteNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = replyDeleteNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ReplyDeleteNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	replyDeleteNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, replyDeleteNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this reply delete not found body
func (o *ReplyDeleteNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *ReplyDeleteNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplyDeleteNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ReplyDeleteNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ReplyDeleteOKBody reply delete o k body
// swagger:model ReplyDeleteOKBody
type ReplyDeleteOKBody struct {
	models.SuccessData

	// data
	Data interface{} `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ReplyDeleteOKBody) UnmarshalJSON(raw []byte) error {
	// ReplyDeleteOKBodyAO0
	var replyDeleteOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &replyDeleteOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = replyDeleteOKBodyAO0

	// ReplyDeleteOKBodyAO1
	var dataReplyDeleteOKBodyAO1 struct {
		Data interface{} `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataReplyDeleteOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataReplyDeleteOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ReplyDeleteOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	replyDeleteOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, replyDeleteOKBodyAO0)

	var dataReplyDeleteOKBodyAO1 struct {
		Data interface{} `json:"data,omitempty"`
	}

	dataReplyDeleteOKBodyAO1.Data = o.Data

	jsonDataReplyDeleteOKBodyAO1, errReplyDeleteOKBodyAO1 := swag.WriteJSON(dataReplyDeleteOKBodyAO1)
	if errReplyDeleteOKBodyAO1 != nil {
		return nil, errReplyDeleteOKBodyAO1
	}
	_parts = append(_parts, jsonDataReplyDeleteOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this reply delete o k body
func (o *ReplyDeleteOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.SuccessData
	if err := o.SuccessData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *ReplyDeleteOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReplyDeleteOKBody) UnmarshalBinary(b []byte) error {
	var res ReplyDeleteOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
