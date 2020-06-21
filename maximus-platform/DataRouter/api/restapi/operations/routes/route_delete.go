// Code generated by go-swagger; DO NOT EDIT.

package routes

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

// RouteDeleteHandlerFunc turns a function with the right signature into a route delete handler
type RouteDeleteHandlerFunc func(RouteDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RouteDeleteHandlerFunc) Handle(params RouteDeleteParams) middleware.Responder {
	return fn(params)
}

// RouteDeleteHandler interface for that can handle valid route delete params
type RouteDeleteHandler interface {
	Handle(RouteDeleteParams) middleware.Responder
}

// NewRouteDelete creates a new http.Handler for the route delete operation
func NewRouteDelete(ctx *middleware.Context, handler RouteDeleteHandler) *RouteDelete {
	return &RouteDelete{Context: ctx, Handler: handler}
}

/*RouteDelete swagger:route DELETE /routes/{routeID} Routes routeDelete

Удаление маршрута

*/
type RouteDelete struct {
	Context *middleware.Context
	Handler RouteDeleteHandler
}

func (o *RouteDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRouteDeleteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// RouteDeleteInternalServerErrorBody route delete internal server error body
// swagger:model RouteDeleteInternalServerErrorBody
type RouteDeleteInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RouteDeleteInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// RouteDeleteInternalServerErrorBodyAO0
	var routeDeleteInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &routeDeleteInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = routeDeleteInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RouteDeleteInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	routeDeleteInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeDeleteInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route delete internal server error body
func (o *RouteDeleteInternalServerErrorBody) Validate(formats strfmt.Registry) error {
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
func (o *RouteDeleteInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteDeleteInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res RouteDeleteInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RouteDeleteMethodNotAllowedBody route delete method not allowed body
// swagger:model RouteDeleteMethodNotAllowedBody
type RouteDeleteMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RouteDeleteMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// RouteDeleteMethodNotAllowedBodyAO0
	var routeDeleteMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &routeDeleteMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = routeDeleteMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RouteDeleteMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	routeDeleteMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeDeleteMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route delete method not allowed body
func (o *RouteDeleteMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *RouteDeleteMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteDeleteMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res RouteDeleteMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RouteDeleteNotFoundBody route delete not found body
// swagger:model RouteDeleteNotFoundBody
type RouteDeleteNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RouteDeleteNotFoundBody) UnmarshalJSON(raw []byte) error {
	// RouteDeleteNotFoundBodyAO0
	var routeDeleteNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &routeDeleteNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = routeDeleteNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RouteDeleteNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	routeDeleteNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeDeleteNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route delete not found body
func (o *RouteDeleteNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *RouteDeleteNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteDeleteNotFoundBody) UnmarshalBinary(b []byte) error {
	var res RouteDeleteNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RouteDeleteOKBody route delete o k body
// swagger:model RouteDeleteOKBody
type RouteDeleteOKBody struct {
	models.SuccessData

	// data
	Data interface{} `json:"data,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RouteDeleteOKBody) UnmarshalJSON(raw []byte) error {
	// RouteDeleteOKBodyAO0
	var routeDeleteOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &routeDeleteOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = routeDeleteOKBodyAO0

	// RouteDeleteOKBodyAO1
	var dataRouteDeleteOKBodyAO1 struct {
		Data interface{} `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataRouteDeleteOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataRouteDeleteOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RouteDeleteOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	routeDeleteOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, routeDeleteOKBodyAO0)

	var dataRouteDeleteOKBodyAO1 struct {
		Data interface{} `json:"data,omitempty"`
	}

	dataRouteDeleteOKBodyAO1.Data = o.Data

	jsonDataRouteDeleteOKBodyAO1, errRouteDeleteOKBodyAO1 := swag.WriteJSON(dataRouteDeleteOKBodyAO1)
	if errRouteDeleteOKBodyAO1 != nil {
		return nil, errRouteDeleteOKBodyAO1
	}
	_parts = append(_parts, jsonDataRouteDeleteOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this route delete o k body
func (o *RouteDeleteOKBody) Validate(formats strfmt.Registry) error {
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
func (o *RouteDeleteOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RouteDeleteOKBody) UnmarshalBinary(b []byte) error {
	var res RouteDeleteOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
