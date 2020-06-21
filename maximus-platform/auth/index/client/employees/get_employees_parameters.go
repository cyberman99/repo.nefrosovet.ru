// Code generated by go-swagger; DO NOT EDIT.

package employees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetEmployeesParams creates a new GetEmployeesParams object
// with the default values initialized.
func NewGetEmployeesParams() *GetEmployeesParams {
	var ()
	return &GetEmployeesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmployeesParamsWithTimeout creates a new GetEmployeesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmployeesParamsWithTimeout(timeout time.Duration) *GetEmployeesParams {
	var ()
	return &GetEmployeesParams{

		timeout: timeout,
	}
}

// NewGetEmployeesParamsWithContext creates a new GetEmployeesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmployeesParamsWithContext(ctx context.Context) *GetEmployeesParams {
	var ()
	return &GetEmployeesParams{

		Context: ctx,
	}
}

// NewGetEmployeesParamsWithHTTPClient creates a new GetEmployeesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmployeesParamsWithHTTPClient(client *http.Client) *GetEmployeesParams {
	var ()
	return &GetEmployeesParams{
		HTTPClient: client,
	}
}

/*GetEmployeesParams contains all the parameters to send to the API endpoint
for the get employees operation typically these are written to a http.Request
*/
type GetEmployeesParams struct {

	/*Limit
	  Лимит выдачи

	*/
	Limit *int64
	/*Offset
	  Шаг выдачи

	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get employees params
func (o *GetEmployeesParams) WithTimeout(timeout time.Duration) *GetEmployeesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get employees params
func (o *GetEmployeesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get employees params
func (o *GetEmployeesParams) WithContext(ctx context.Context) *GetEmployeesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get employees params
func (o *GetEmployeesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get employees params
func (o *GetEmployeesParams) WithHTTPClient(client *http.Client) *GetEmployeesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get employees params
func (o *GetEmployeesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get employees params
func (o *GetEmployeesParams) WithLimit(limit *int64) *GetEmployeesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get employees params
func (o *GetEmployeesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get employees params
func (o *GetEmployeesParams) WithOffset(offset *int64) *GetEmployeesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get employees params
func (o *GetEmployeesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmployeesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
