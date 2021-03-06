// Code generated by go-swagger; DO NOT EDIT.

package photo

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

	models "repo.nefrosovet.ru/maximus-platform/recognition/api/models"
)

// ViewHandlerFunc turns a function with the right signature into a view handler
type ViewHandlerFunc func(ViewParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ViewHandlerFunc) Handle(params ViewParams) middleware.Responder {
	return fn(params)
}

// ViewHandler interface for that can handle valid view params
type ViewHandler interface {
	Handle(ViewParams) middleware.Responder
}

// NewView creates a new http.Handler for the view operation
func NewView(ctx *middleware.Context, handler ViewHandler) *View {
	return &View{Context: ctx, Handler: handler}
}

/*View swagger:route GET /photos/{photoID} Photo view

Информация о фотографии

*/
type View struct {
	Context *middleware.Context
	Handler ViewHandler
}

func (o *View) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewViewParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ViewBadRequestBody view bad request body
// swagger:model ViewBadRequestBody
type ViewBadRequestBody struct {
	models.Error400Data

	// errors
	// Required: true
	Errors *ViewBadRequestBodyAO1Errors `json:"errors"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ViewBadRequestBody) UnmarshalJSON(raw []byte) error {
	// ViewBadRequestBodyAO0
	var viewBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &viewBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = viewBadRequestBodyAO0

	// ViewBadRequestBodyAO1
	var dataViewBadRequestBodyAO1 struct {
		Errors *ViewBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}
	if err := swag.ReadJSON(raw, &dataViewBadRequestBodyAO1); err != nil {
		return err
	}

	o.Errors = dataViewBadRequestBodyAO1.Errors

	o.Message = dataViewBadRequestBodyAO1.Message

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ViewBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	viewBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, viewBadRequestBodyAO0)

	var dataViewBadRequestBodyAO1 struct {
		Errors *ViewBadRequestBodyAO1Errors `json:"errors"`

		Message *string `json:"message"`
	}

	dataViewBadRequestBodyAO1.Errors = o.Errors

	dataViewBadRequestBodyAO1.Message = o.Message

	jsonDataViewBadRequestBodyAO1, errViewBadRequestBodyAO1 := swag.WriteJSON(dataViewBadRequestBodyAO1)
	if errViewBadRequestBodyAO1 != nil {
		return nil, errViewBadRequestBodyAO1
	}
	_parts = append(_parts, jsonDataViewBadRequestBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this view bad request body
func (o *ViewBadRequestBody) Validate(formats strfmt.Registry) error {
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

func (o *ViewBadRequestBody) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("viewBadRequest"+"."+"errors", "body", o.Errors); err != nil {
		return err
	}

	if o.Errors != nil {
		if err := o.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("viewBadRequest" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

func (o *ViewBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("viewBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ViewBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ViewBadRequestBody) UnmarshalBinary(b []byte) error {
	var res ViewBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ViewBadRequestBodyAO1Errors view bad request body a o1 errors
// swagger:model ViewBadRequestBodyAO1Errors
type ViewBadRequestBodyAO1Errors struct {

	// core
	Core string `json:"core,omitempty"`

	// json
	JSON string `json:"json,omitempty"`

	// validation
	Validation interface{} `json:"validation,omitempty"`
}

// Validate validates this view bad request body a o1 errors
func (o *ViewBadRequestBodyAO1Errors) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ViewBadRequestBodyAO1Errors) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ViewBadRequestBodyAO1Errors) UnmarshalBinary(b []byte) error {
	var res ViewBadRequestBodyAO1Errors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ViewInternalServerErrorBody view internal server error body
// swagger:model ViewInternalServerErrorBody
type ViewInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ViewInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// ViewInternalServerErrorBodyAO0
	var viewInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &viewInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = viewInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ViewInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	viewInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, viewInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this view internal server error body
func (o *ViewInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *ViewInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ViewInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ViewInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ViewMethodNotAllowedBody view method not allowed body
// swagger:model ViewMethodNotAllowedBody
type ViewMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ViewMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// ViewMethodNotAllowedBodyAO0
	var viewMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &viewMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = viewMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ViewMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	viewMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, viewMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this view method not allowed body
func (o *ViewMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *ViewMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ViewMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res ViewMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ViewNotFoundBody view not found body
// swagger:model ViewNotFoundBody
type ViewNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ViewNotFoundBody) UnmarshalJSON(raw []byte) error {
	// ViewNotFoundBodyAO0
	var viewNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &viewNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = viewNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ViewNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	viewNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, viewNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this view not found body
func (o *ViewNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *ViewNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ViewNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ViewNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ViewOKBody view o k body
// swagger:model ViewOKBody
type ViewOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ViewOKBody) UnmarshalJSON(raw []byte) error {
	// ViewOKBodyAO0
	var viewOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &viewOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = viewOKBodyAO0

	// ViewOKBodyAO1
	var dataViewOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataViewOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataViewOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ViewOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	viewOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, viewOKBodyAO0)

	var dataViewOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataViewOKBodyAO1.Data = o.Data

	jsonDataViewOKBodyAO1, errViewOKBodyAO1 := swag.WriteJSON(dataViewOKBodyAO1)
	if errViewOKBodyAO1 != nil {
		return nil, errViewOKBodyAO1
	}
	_parts = append(_parts, jsonDataViewOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this view o k body
func (o *ViewOKBody) Validate(formats strfmt.Registry) error {
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

func (o *ViewOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("viewOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("viewOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ViewOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ViewOKBody) UnmarshalBinary(b []byte) error {
	var res ViewOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
