// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteBackendsBackendIDGroupsParams creates a new DeleteBackendsBackendIDGroupsParams object
// no default values defined in spec.
func NewDeleteBackendsBackendIDGroupsParams() DeleteBackendsBackendIDGroupsParams {

	return DeleteBackendsBackendIDGroupsParams{}
}

// DeleteBackendsBackendIDGroupsParams contains all the bound params for the delete backends backend ID groups operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteBackendsBackendIDGroups
type DeleteBackendsBackendIDGroupsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Идентификатор бэкенда
	  Required: true
	  In: path
	*/
	BackendID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteBackendsBackendIDGroupsParams() beforehand.
func (o *DeleteBackendsBackendIDGroupsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBackendID, rhkBackendID, _ := route.Params.GetOK("backendID")
	if err := o.bindBackendID(rBackendID, rhkBackendID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBackendID binds and validates parameter BackendID from path.
func (o *DeleteBackendsBackendIDGroupsParams) bindBackendID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.BackendID = raw

	return nil
}
