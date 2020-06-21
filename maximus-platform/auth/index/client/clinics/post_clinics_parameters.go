// Code generated by go-swagger; DO NOT EDIT.

package clinics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostClinicsParams creates a new PostClinicsParams object
// with the default values initialized.
func NewPostClinicsParams() *PostClinicsParams {
	var ()
	return &PostClinicsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostClinicsParamsWithTimeout creates a new PostClinicsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostClinicsParamsWithTimeout(timeout time.Duration) *PostClinicsParams {
	var ()
	return &PostClinicsParams{

		timeout: timeout,
	}
}

// NewPostClinicsParamsWithContext creates a new PostClinicsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostClinicsParamsWithContext(ctx context.Context) *PostClinicsParams {
	var ()
	return &PostClinicsParams{

		Context: ctx,
	}
}

// NewPostClinicsParamsWithHTTPClient creates a new PostClinicsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostClinicsParamsWithHTTPClient(client *http.Client) *PostClinicsParams {
	var ()
	return &PostClinicsParams{
		HTTPClient: client,
	}
}

/*PostClinicsParams contains all the parameters to send to the API endpoint
for the post clinics operation typically these are written to a http.Request
*/
type PostClinicsParams struct {

	/*Body*/
	Body PostClinicsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post clinics params
func (o *PostClinicsParams) WithTimeout(timeout time.Duration) *PostClinicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post clinics params
func (o *PostClinicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post clinics params
func (o *PostClinicsParams) WithContext(ctx context.Context) *PostClinicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post clinics params
func (o *PostClinicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post clinics params
func (o *PostClinicsParams) WithHTTPClient(client *http.Client) *PostClinicsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post clinics params
func (o *PostClinicsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post clinics params
func (o *PostClinicsParams) WithBody(body PostClinicsBody) *PostClinicsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post clinics params
func (o *PostClinicsParams) SetBody(body PostClinicsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostClinicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}