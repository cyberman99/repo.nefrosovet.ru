// Code generated by go-swagger; DO NOT EDIT.

package clinic_employees

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

// NewPostClinicsClinicGUIDEmployeesEmployeeGUIDParams creates a new PostClinicsClinicGUIDEmployeesEmployeeGUIDParams object
// with the default values initialized.
func NewPostClinicsClinicGUIDEmployeesEmployeeGUIDParams() *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	var ()
	return &PostClinicsClinicGUIDEmployeesEmployeeGUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostClinicsClinicGUIDEmployeesEmployeeGUIDParamsWithTimeout creates a new PostClinicsClinicGUIDEmployeesEmployeeGUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostClinicsClinicGUIDEmployeesEmployeeGUIDParamsWithTimeout(timeout time.Duration) *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	var ()
	return &PostClinicsClinicGUIDEmployeesEmployeeGUIDParams{

		timeout: timeout,
	}
}

// NewPostClinicsClinicGUIDEmployeesEmployeeGUIDParamsWithContext creates a new PostClinicsClinicGUIDEmployeesEmployeeGUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostClinicsClinicGUIDEmployeesEmployeeGUIDParamsWithContext(ctx context.Context) *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	var ()
	return &PostClinicsClinicGUIDEmployeesEmployeeGUIDParams{

		Context: ctx,
	}
}

// NewPostClinicsClinicGUIDEmployeesEmployeeGUIDParamsWithHTTPClient creates a new PostClinicsClinicGUIDEmployeesEmployeeGUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostClinicsClinicGUIDEmployeesEmployeeGUIDParamsWithHTTPClient(client *http.Client) *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	var ()
	return &PostClinicsClinicGUIDEmployeesEmployeeGUIDParams{
		HTTPClient: client,
	}
}

/*PostClinicsClinicGUIDEmployeesEmployeeGUIDParams contains all the parameters to send to the API endpoint
for the post clinics clinic GUID employees employee GUID operation typically these are written to a http.Request
*/
type PostClinicsClinicGUIDEmployeesEmployeeGUIDParams struct {

	/*ClinicGUID
	  GUID клиники

	*/
	ClinicGUID string
	/*EmployeeGUID
	  GUID сотрудника

	*/
	EmployeeGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post clinics clinic GUID employees employee GUID params
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams) WithTimeout(timeout time.Duration) *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post clinics clinic GUID employees employee GUID params
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post clinics clinic GUID employees employee GUID params
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams) WithContext(ctx context.Context) *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post clinics clinic GUID employees employee GUID params
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post clinics clinic GUID employees employee GUID params
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams) WithHTTPClient(client *http.Client) *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post clinics clinic GUID employees employee GUID params
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClinicGUID adds the clinicGUID to the post clinics clinic GUID employees employee GUID params
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams) WithClinicGUID(clinicGUID string) *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	o.SetClinicGUID(clinicGUID)
	return o
}

// SetClinicGUID adds the clinicGuid to the post clinics clinic GUID employees employee GUID params
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams) SetClinicGUID(clinicGUID string) {
	o.ClinicGUID = clinicGUID
}

// WithEmployeeGUID adds the employeeGUID to the post clinics clinic GUID employees employee GUID params
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams) WithEmployeeGUID(employeeGUID string) *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams {
	o.SetEmployeeGUID(employeeGUID)
	return o
}

// SetEmployeeGUID adds the employeeGuid to the post clinics clinic GUID employees employee GUID params
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams) SetEmployeeGUID(employeeGUID string) {
	o.EmployeeGUID = employeeGUID
}

// WriteToRequest writes these params to a swagger request
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clinicGUID
	if err := r.SetPathParam("clinicGUID", o.ClinicGUID); err != nil {
		return err
	}

	// path param employeeGUID
	if err := r.SetPathParam("employeeGUID", o.EmployeeGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}