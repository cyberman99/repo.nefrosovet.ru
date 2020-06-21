// Code generated by go-swagger; DO NOT EDIT.

package utils

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

// NewPostLinkParams creates a new PostLinkParams object
// with the default values initialized.
func NewPostLinkParams() *PostLinkParams {
	var ()
	return &PostLinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLinkParamsWithTimeout creates a new PostLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLinkParamsWithTimeout(timeout time.Duration) *PostLinkParams {
	var ()
	return &PostLinkParams{

		timeout: timeout,
	}
}

// NewPostLinkParamsWithContext creates a new PostLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLinkParamsWithContext(ctx context.Context) *PostLinkParams {
	var ()
	return &PostLinkParams{

		Context: ctx,
	}
}

// NewPostLinkParamsWithHTTPClient creates a new PostLinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLinkParamsWithHTTPClient(client *http.Client) *PostLinkParams {
	var ()
	return &PostLinkParams{
		HTTPClient: client,
	}
}

/*PostLinkParams contains all the parameters to send to the API endpoint
for the post link operation typically these are written to a http.Request
*/
type PostLinkParams struct {

	/*Link
	  <patientGUID>;rel=firstName

	*/
	Link string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post link params
func (o *PostLinkParams) WithTimeout(timeout time.Duration) *PostLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post link params
func (o *PostLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post link params
func (o *PostLinkParams) WithContext(ctx context.Context) *PostLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post link params
func (o *PostLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post link params
func (o *PostLinkParams) WithHTTPClient(client *http.Client) *PostLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post link params
func (o *PostLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLink adds the link to the post link params
func (o *PostLinkParams) WithLink(link string) *PostLinkParams {
	o.SetLink(link)
	return o
}

// SetLink adds the link to the post link params
func (o *PostLinkParams) SetLink(link string) {
	o.Link = link
}

// WriteToRequest writes these params to a swagger request
func (o *PostLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Link
	if err := r.SetHeaderParam("Link", o.Link); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
