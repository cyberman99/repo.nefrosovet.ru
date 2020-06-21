// Code generated by go-swagger; DO NOT EDIT.

package photo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "repo.nefrosovet.ru/maximus-platform/recognition/api/models"
)

// DeleteHandlerFunc turns a function with the right signature into a delete handler
type DeleteHandlerFunc func(DeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteHandlerFunc) Handle(params DeleteParams) middleware.Responder {
	return fn(params)
}

// DeleteHandler interface for that can handle valid delete params
type DeleteHandler interface {
	Handle(DeleteParams) middleware.Responder
}

// NewDelete creates a new http.Handler for the delete operation
func NewDelete(ctx *middleware.Context, handler DeleteHandler) *Delete {
	return &Delete{Context: ctx, Handler: handler}
}

/*Delete swagger:route DELETE /photos/{photoID} Photo delete

Удаление фотографии

*/
type Delete struct {
	Context *middleware.Context
	Handler DeleteHandler
}

func (o *Delete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteBadRequestBody delete bad request body
// swagger:model DeleteBadRequestBody
type DeleteBadRequestBody struct {
	models.Error400Data

	// errors
	// Required: true
	Errors *DeleteBadRequestBodyAO1Errors `json:"errors"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteBadRequestBody) UnmarshalJSON(raw []byte) error {
	// DeleteBadRequestBodyAO0
	var deleteBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &deleteBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = deleteBadRequestBodyAO0

	// DeleteBadRequestBodyAO1
	var dataDeleteBadRequestBodyAO1 struct {
		Errors *DeleteBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}
	if err := swag.ReadJSON(raw, &dataDeleteBadRequestBodyAO1); err != nil {
		return err
	}

	o.Errors = dataDeleteBadRequestBodyAO1.Errors

	o.Message = dataDeleteBadRequestBodyAO1.Message

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	deleteBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteBadRequestBodyAO0)

	var dataDeleteBadRequestBodyAO1 struct {
		Errors *DeleteBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}

	dataDeleteBadRequestBodyAO1.Errors = o.Errors

	dataDeleteBadRequestBodyAO1.Message = o.Message

	jsonDataDeleteBadRequestBodyAO1, errDeleteBadRequestBodyAO1 := swag.WriteJSON(dataDeleteBadRequestBodyAO1)
	if errDeleteBadRequestBodyAO1 != nil {
		return nil, errDeleteBadRequestBodyAO1
	}
	_parts = append(_parts, jsonDataDeleteBadRequestBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete bad request body
func (o *DeleteBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error400Data
	if err := o.Error400Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteBadRequestBody) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("deleteBadRequest"+"."+"errors", "body", o.Errors); err != nil {
		return err
	}

	if o.Errors != nil {
		if err := o.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteBadRequest" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

func (o *DeleteBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("deleteBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteBadRequestBodyAO1Errors delete bad request body a o1 errors
// swagger:model DeleteBadRequestBodyAO1Errors
type DeleteBadRequestBodyAO1Errors struct {

	// core
	Core string `json:"core,omitempty"`

	// json
	JSON string `json:"json,omitempty"`

	// validation
	Validation interface{} `json:"validation,omitempty"`
}

// Validate validates this delete bad request body a o1 errors
func (o *DeleteBadRequestBodyAO1Errors) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteBadRequestBodyAO1Errors) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteBadRequestBodyAO1Errors) UnmarshalBinary(b []byte) error {
	var res DeleteBadRequestBodyAO1Errors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteInternalServerErrorBody delete internal server error body
// swagger:model DeleteInternalServerErrorBody
type DeleteInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// DeleteInternalServerErrorBodyAO0
	var deleteInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &deleteInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = deleteInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	deleteInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete internal server error body
func (o *DeleteInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *DeleteInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DeleteInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteMethodNotAllowedBody delete method not allowed body
// swagger:model DeleteMethodNotAllowedBody
type DeleteMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// DeleteMethodNotAllowedBodyAO0
	var deleteMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &deleteMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = deleteMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	deleteMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete method not allowed body
func (o *DeleteMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *DeleteMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res DeleteMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteNotFoundBody delete not found body
// swagger:model DeleteNotFoundBody
type DeleteNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteNotFoundBody) UnmarshalJSON(raw []byte) error {
	// DeleteNotFoundBodyAO0
	var deleteNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &deleteNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = deleteNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	deleteNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete not found body
func (o *DeleteNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *DeleteNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteOKBody delete o k body
// swagger:model DeleteOKBody
type DeleteOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []interface{} `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DeleteOKBody) UnmarshalJSON(raw []byte) error {
	// DeleteOKBodyAO0
	var deleteOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &deleteOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = deleteOKBodyAO0

	// DeleteOKBodyAO1
	var dataDeleteOKBodyAO1 struct {
		Data []interface{} `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataDeleteOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataDeleteOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DeleteOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	deleteOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, deleteOKBodyAO0)

	var dataDeleteOKBodyAO1 struct {
		Data []interface{} `json:"data"`
	}

	dataDeleteOKBodyAO1.Data = o.Data

	jsonDataDeleteOKBodyAO1, errDeleteOKBodyAO1 := swag.WriteJSON(dataDeleteOKBodyAO1)
	if errDeleteOKBodyAO1 != nil {
		return nil, errDeleteOKBodyAO1
	}
	_parts = append(_parts, jsonDataDeleteOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete o k body
func (o *DeleteOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DeleteOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("deleteOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
