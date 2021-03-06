// Code generated by go-swagger; DO NOT EDIT.

package webhook

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

// NewPostWebhooksWebhookIDParams creates a new PostWebhooksWebhookIDParams object
// with the default values initialized.
func NewPostWebhooksWebhookIDParams() *PostWebhooksWebhookIDParams {
	var ()
	return &PostWebhooksWebhookIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWebhooksWebhookIDParamsWithTimeout creates a new PostWebhooksWebhookIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWebhooksWebhookIDParamsWithTimeout(timeout time.Duration) *PostWebhooksWebhookIDParams {
	var ()
	return &PostWebhooksWebhookIDParams{

		timeout: timeout,
	}
}

// NewPostWebhooksWebhookIDParamsWithContext creates a new PostWebhooksWebhookIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWebhooksWebhookIDParamsWithContext(ctx context.Context) *PostWebhooksWebhookIDParams {
	var ()
	return &PostWebhooksWebhookIDParams{

		Context: ctx,
	}
}

// NewPostWebhooksWebhookIDParamsWithHTTPClient creates a new PostWebhooksWebhookIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWebhooksWebhookIDParamsWithHTTPClient(client *http.Client) *PostWebhooksWebhookIDParams {
	var ()
	return &PostWebhooksWebhookIDParams{
		HTTPClient: client,
	}
}

/*PostWebhooksWebhookIDParams contains all the parameters to send to the API endpoint
for the post webhooks webhook ID operation typically these are written to a http.Request
*/
type PostWebhooksWebhookIDParams struct {

	/*Body*/
	Body *models.PostWebhooksWebhookIDParamsBody
	/*WebhookID
	  Идентификатор хука

	*/
	WebhookID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post webhooks webhook ID params
func (o *PostWebhooksWebhookIDParams) WithTimeout(timeout time.Duration) *PostWebhooksWebhookIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post webhooks webhook ID params
func (o *PostWebhooksWebhookIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post webhooks webhook ID params
func (o *PostWebhooksWebhookIDParams) WithContext(ctx context.Context) *PostWebhooksWebhookIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post webhooks webhook ID params
func (o *PostWebhooksWebhookIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post webhooks webhook ID params
func (o *PostWebhooksWebhookIDParams) WithHTTPClient(client *http.Client) *PostWebhooksWebhookIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post webhooks webhook ID params
func (o *PostWebhooksWebhookIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post webhooks webhook ID params
func (o *PostWebhooksWebhookIDParams) WithBody(body *models.PostWebhooksWebhookIDParamsBody) *PostWebhooksWebhookIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post webhooks webhook ID params
func (o *PostWebhooksWebhookIDParams) SetBody(body *models.PostWebhooksWebhookIDParamsBody) {
	o.Body = body
}

// WithWebhookID adds the webhookID to the post webhooks webhook ID params
func (o *PostWebhooksWebhookIDParams) WithWebhookID(webhookID string) *PostWebhooksWebhookIDParams {
	o.SetWebhookID(webhookID)
	return o
}

// SetWebhookID adds the webhookId to the post webhooks webhook ID params
func (o *PostWebhooksWebhookIDParams) SetWebhookID(webhookID string) {
	o.WebhookID = webhookID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWebhooksWebhookIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param webhookID
	if err := r.SetPathParam("webhookID", o.WebhookID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
