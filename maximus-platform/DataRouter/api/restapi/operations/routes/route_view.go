// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "repo.nefrosovet.ru/maximus-platform/DataRouter/api/models"
)

// RouteViewHandlerFunc turns a function with the right signature into a route view handler
type RouteViewHandlerFunc func(RouteViewParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RouteViewHandlerFunc) Handle(params RouteViewParams) middleware.Responder {
	return fn(params)
}

// RouteViewHandler interface for that can handle valid route view params
type RouteViewHandler interface {
	Handle(RouteViewParams) middleware.Responder
}

// NewRouteView creates a new http.Handler for the route view operation
func NewRouteView(ctx *middleware.Context, handler RouteViewHandler) *RouteView {
	return &RouteView{Context: ctx, Handler: handler}
}

/*RouteView swagger:route GET /routes/{routeID} Routes routeView

Информация о маршруте

*/
type RouteView struct {
	Context *middleware.Context
	Handler RouteViewHandler
}

func (o *RouteView) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRouteViewParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// RouteViewInternalServerErrorBody route view internal server error body
// swagger:model RouteViewInternalServerErrorBody
type RouteViewInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RouteViewInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// RouteViewInternalServerErrorBodyAO0
	var routeViewInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &routeViewInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = routeViewInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RouteViewInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	routeViewInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeViewInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route view internal server error body
func (o *RouteViewInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *RouteViewInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteViewInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res RouteViewInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RouteViewMethodNotAllowedBody route view method not allowed body
// swagger:model RouteViewMethodNotAllowedBody
type RouteViewMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RouteViewMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// RouteViewMethodNotAllowedBodyAO0
	var routeViewMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &routeViewMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = routeViewMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RouteViewMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	routeViewMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeViewMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route view method not allowed body
func (o *RouteViewMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *RouteViewMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteViewMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res RouteViewMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RouteViewNotFoundBody route view not found body
// swagger:model RouteViewNotFoundBody
type RouteViewNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RouteViewNotFoundBody) UnmarshalJSON(raw []byte) error {
	// RouteViewNotFoundBodyAO0
	var routeViewNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &routeViewNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = routeViewNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RouteViewNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	routeViewNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeViewNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route view not found body
func (o *RouteViewNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *RouteViewNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteViewNotFoundBody) UnmarshalBinary(b []byte) error {
	var res RouteViewNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RouteViewOKBody route view o k body
// swagger:model RouteViewOKBody
type RouteViewOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RouteViewOKBody) UnmarshalJSON(raw []byte) error {
	// RouteViewOKBodyAO0
	var routeViewOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &routeViewOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = routeViewOKBodyAO0

	// RouteViewOKBodyAO1
	var dataRouteViewOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataRouteViewOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataRouteViewOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RouteViewOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	routeViewOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeViewOKBodyAO0)

	var dataRouteViewOKBodyAO1 struct {
		Data []*DataItems0 `json:"data"`
	}

	dataRouteViewOKBodyAO1.Data = o.Data

	jsonDataRouteViewOKBodyAO1, errRouteViewOKBodyAO1 := swag.WriteJSON(dataRouteViewOKBodyAO1)
	if errRouteViewOKBodyAO1 != nil {
		return nil, errRouteViewOKBodyAO1
	}
	_parts = append(_parts, jsonDataRouteViewOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route view o k body
func (o *RouteViewOKBody) Validate(formats strfmt.Registry) error {
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

func (o *RouteViewOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("routeViewOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *RouteViewOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteViewOKBody) UnmarshalBinary(b []byte) error {
	var res RouteViewOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
