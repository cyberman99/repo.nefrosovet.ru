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

// NewPostUnlinkPatientGUIDParams creates a new PostUnlinkPatientGUIDParams object
// with the default values initialized.
func NewPostUnlinkPatientGUIDParams() *PostUnlinkPatientGUIDParams {
	var ()
	return &PostUnlinkPatientGUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUnlinkPatientGUIDParamsWithTimeout creates a new PostUnlinkPatientGUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUnlinkPatientGUIDParamsWithTimeout(timeout time.Duration) *PostUnlinkPatientGUIDParams {
	var ()
	return &PostUnlinkPatientGUIDParams{

		timeout: timeout,
	}
}

// NewPostUnlinkPatientGUIDParamsWithContext creates a new PostUnlinkPatientGUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUnlinkPatientGUIDParamsWithContext(ctx context.Context) *PostUnlinkPatientGUIDParams {
	var ()
	return &PostUnlinkPatientGUIDParams{

		Context: ctx,
	}
}

// NewPostUnlinkPatientGUIDParamsWithHTTPClient creates a new PostUnlinkPatientGUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUnlinkPatientGUIDParamsWithHTTPClient(client *http.Client) *PostUnlinkPatientGUIDParams {
	var ()
	return &PostUnlinkPatientGUIDParams{
		HTTPClient: client,
	}
}

/*PostUnlinkPatientGUIDParams contains all the parameters to send to the API endpoint
for the post unlink patient GUID operation typically these are written to a http.Request
*/
type PostUnlinkPatientGUIDParams struct {

	/*PatientGUID
	  GUID пациента

	*/
	PatientGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post unlink patient GUID params
func (o *PostUnlinkPatientGUIDParams) WithTimeout(timeout time.Duration) *PostUnlinkPatientGUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post unlink patient GUID params
func (o *PostUnlinkPatientGUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post unlink patient GUID params
func (o *PostUnlinkPatientGUIDParams) WithContext(ctx context.Context) *PostUnlinkPatientGUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post unlink patient GUID params
func (o *PostUnlinkPatientGUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post unlink patient GUID params
func (o *PostUnlinkPatientGUIDParams) WithHTTPClient(client *http.Client) *PostUnlinkPatientGUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post unlink patient GUID params
func (o *PostUnlinkPatientGUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatientGUID adds the patientGUID to the post unlink patient GUID params
func (o *PostUnlinkPatientGUIDParams) WithPatientGUID(patientGUID string) *PostUnlinkPatientGUIDParams {
	o.SetPatientGUID(patientGUID)
	return o
}

// SetPatientGUID adds the patientGuid to the post unlink patient GUID params
func (o *PostUnlinkPatientGUIDParams) SetPatientGUID(patientGUID string) {
	o.PatientGUID = patientGUID
}

// WriteToRequest writes these params to a swagger request
func (o *PostUnlinkPatientGUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param patientGUID
	if err := r.SetPathParam("patientGUID", o.PatientGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
