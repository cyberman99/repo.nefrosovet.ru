// Code generated by go-swagger; DO NOT EDIT.

package patients

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

// NewPatchPatientsPatientGUIDParams creates a new PatchPatientsPatientGUIDParams object
// with the default values initialized.
func NewPatchPatientsPatientGUIDParams() *PatchPatientsPatientGUIDParams {
	var ()
	return &PatchPatientsPatientGUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPatientsPatientGUIDParamsWithTimeout creates a new PatchPatientsPatientGUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchPatientsPatientGUIDParamsWithTimeout(timeout time.Duration) *PatchPatientsPatientGUIDParams {
	var ()
	return &PatchPatientsPatientGUIDParams{

		timeout: timeout,
	}
}

// NewPatchPatientsPatientGUIDParamsWithContext creates a new PatchPatientsPatientGUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchPatientsPatientGUIDParamsWithContext(ctx context.Context) *PatchPatientsPatientGUIDParams {
	var ()
	return &PatchPatientsPatientGUIDParams{

		Context: ctx,
	}
}

// NewPatchPatientsPatientGUIDParamsWithHTTPClient creates a new PatchPatientsPatientGUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchPatientsPatientGUIDParamsWithHTTPClient(client *http.Client) *PatchPatientsPatientGUIDParams {
	var ()
	return &PatchPatientsPatientGUIDParams{
		HTTPClient: client,
	}
}

/*PatchPatientsPatientGUIDParams contains all the parameters to send to the API endpoint
for the patch patients patient GUID operation typically these are written to a http.Request
*/
type PatchPatientsPatientGUIDParams struct {

	/*Body*/
	Body PatchPatientsPatientGUIDBody
	/*PatientGUID
	  GUID пациента

	*/
	PatientGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch patients patient GUID params
func (o *PatchPatientsPatientGUIDParams) WithTimeout(timeout time.Duration) *PatchPatientsPatientGUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch patients patient GUID params
func (o *PatchPatientsPatientGUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch patients patient GUID params
func (o *PatchPatientsPatientGUIDParams) WithContext(ctx context.Context) *PatchPatientsPatientGUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch patients patient GUID params
func (o *PatchPatientsPatientGUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch patients patient GUID params
func (o *PatchPatientsPatientGUIDParams) WithHTTPClient(client *http.Client) *PatchPatientsPatientGUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch patients patient GUID params
func (o *PatchPatientsPatientGUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch patients patient GUID params
func (o *PatchPatientsPatientGUIDParams) WithBody(body PatchPatientsPatientGUIDBody) *PatchPatientsPatientGUIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch patients patient GUID params
func (o *PatchPatientsPatientGUIDParams) SetBody(body PatchPatientsPatientGUIDBody) {
	o.Body = body
}

// WithPatientGUID adds the patientGUID to the patch patients patient GUID params
func (o *PatchPatientsPatientGUIDParams) WithPatientGUID(patientGUID string) *PatchPatientsPatientGUIDParams {
	o.SetPatientGUID(patientGUID)
	return o
}

// SetPatientGUID adds the patientGuid to the patch patients patient GUID params
func (o *PatchPatientsPatientGUIDParams) SetPatientGUID(patientGUID string) {
	o.PatientGUID = patientGUID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPatientsPatientGUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param patientGUID
	if err := r.SetPathParam("patientGUID", o.PatientGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
