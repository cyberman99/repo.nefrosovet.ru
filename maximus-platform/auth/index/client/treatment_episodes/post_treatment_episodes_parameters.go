// Code generated by go-swagger; DO NOT EDIT.

package treatment_episodes

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

// NewPostTreatmentEpisodesParams creates a new PostTreatmentEpisodesParams object
// with the default values initialized.
func NewPostTreatmentEpisodesParams() *PostTreatmentEpisodesParams {
	var ()
	return &PostTreatmentEpisodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTreatmentEpisodesParamsWithTimeout creates a new PostTreatmentEpisodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTreatmentEpisodesParamsWithTimeout(timeout time.Duration) *PostTreatmentEpisodesParams {
	var ()
	return &PostTreatmentEpisodesParams{

		timeout: timeout,
	}
}

// NewPostTreatmentEpisodesParamsWithContext creates a new PostTreatmentEpisodesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTreatmentEpisodesParamsWithContext(ctx context.Context) *PostTreatmentEpisodesParams {
	var ()
	return &PostTreatmentEpisodesParams{

		Context: ctx,
	}
}

// NewPostTreatmentEpisodesParamsWithHTTPClient creates a new PostTreatmentEpisodesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTreatmentEpisodesParamsWithHTTPClient(client *http.Client) *PostTreatmentEpisodesParams {
	var ()
	return &PostTreatmentEpisodesParams{
		HTTPClient: client,
	}
}

/*PostTreatmentEpisodesParams contains all the parameters to send to the API endpoint
for the post treatment episodes operation typically these are written to a http.Request
*/
type PostTreatmentEpisodesParams struct {

	/*Body*/
	Body PostTreatmentEpisodesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post treatment episodes params
func (o *PostTreatmentEpisodesParams) WithTimeout(timeout time.Duration) *PostTreatmentEpisodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post treatment episodes params
func (o *PostTreatmentEpisodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post treatment episodes params
func (o *PostTreatmentEpisodesParams) WithContext(ctx context.Context) *PostTreatmentEpisodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post treatment episodes params
func (o *PostTreatmentEpisodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post treatment episodes params
func (o *PostTreatmentEpisodesParams) WithHTTPClient(client *http.Client) *PostTreatmentEpisodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post treatment episodes params
func (o *PostTreatmentEpisodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post treatment episodes params
func (o *PostTreatmentEpisodesParams) WithBody(body PostTreatmentEpisodesBody) *PostTreatmentEpisodesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post treatment episodes params
func (o *PostTreatmentEpisodesParams) SetBody(body PostTreatmentEpisodesBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostTreatmentEpisodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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