// Code generated by go-swagger; DO NOT EDIT.

package viber

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

	models "repo.nefrosovet.ru/maximus-platform/mailer/bot_proxy/models"
)

// NewPostWebhooksViberParams creates a new PostWebhooksViberParams object
// with the default values initialized.
func NewPostWebhooksViberParams() *PostWebhooksViberParams {
	var ()
	return &PostWebhooksViberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWebhooksViberParamsWithTimeout creates a new PostWebhooksViberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWebhooksViberParamsWithTimeout(timeout time.Duration) *PostWebhooksViberParams {
	var ()
	return &PostWebhooksViberParams{

		timeout: timeout,
	}
}

// NewPostWebhooksViberParamsWithContext creates a new PostWebhooksViberParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWebhooksViberParamsWithContext(ctx context.Context) *PostWebhooksViberParams {
	var ()
	return &PostWebhooksViberParams{

		Context: ctx,
	}
}

// NewPostWebhooksViberParamsWithHTTPClient creates a new PostWebhooksViberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWebhooksViberParamsWithHTTPClient(client *http.Client) *PostWebhooksViberParams {
	var ()
	return &PostWebhooksViberParams{
		HTTPClient: client,
	}
}

/*PostWebhooksViberParams contains all the parameters to send to the API endpoint
for the post webhooks viber operation typically these are written to a http.Request
*/
type PostWebhooksViberParams struct {

	/*Body*/
	Body *models.PostWebhooksViberParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post webhooks viber params
func (o *PostWebhooksViberParams) WithTimeout(timeout time.Duration) *PostWebhooksViberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post webhooks viber params
func (o *PostWebhooksViberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post webhooks viber params
func (o *PostWebhooksViberParams) WithContext(ctx context.Context) *PostWebhooksViberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post webhooks viber params
func (o *PostWebhooksViberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post webhooks viber params
func (o *PostWebhooksViberParams) WithHTTPClient(client *http.Client) *PostWebhooksViberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post webhooks viber params
func (o *PostWebhooksViberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post webhooks viber params
func (o *PostWebhooksViberParams) WithBody(body *models.PostWebhooksViberParamsBody) *PostWebhooksViberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post webhooks viber params
func (o *PostWebhooksViberParams) SetBody(body *models.PostWebhooksViberParamsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostWebhooksViberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
