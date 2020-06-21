// Code generated by go-swagger; DO NOT EDIT.

package departments

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

// NewPutClinicsClinicGUIDDepartmentsDepartmentGUIDParams creates a new PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams object
// with the default values initialized.
func NewPutClinicsClinicGUIDDepartmentsDepartmentGUIDParams() *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams {
	var ()
	return &PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutClinicsClinicGUIDDepartmentsDepartmentGUIDParamsWithTimeout creates a new PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutClinicsClinicGUIDDepartmentsDepartmentGUIDParamsWithTimeout(timeout time.Duration) *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams {
	var ()
	return &PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams{

		timeout: timeout,
	}
}

// NewPutClinicsClinicGUIDDepartmentsDepartmentGUIDParamsWithContext creates a new PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutClinicsClinicGUIDDepartmentsDepartmentGUIDParamsWithContext(ctx context.Context) *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams {
	var ()
	return &PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams{

		Context: ctx,
	}
}

// NewPutClinicsClinicGUIDDepartmentsDepartmentGUIDParamsWithHTTPClient creates a new PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutClinicsClinicGUIDDepartmentsDepartmentGUIDParamsWithHTTPClient(client *http.Client) *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams {
	var ()
	return &PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams{
		HTTPClient: client,
	}
}

/*PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams contains all the parameters to send to the API endpoint
for the put clinics clinic GUID departments department GUID operation typically these are written to a http.Request
*/
type PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams struct {

	/*Body*/
	Body PutClinicsClinicGUIDDepartmentsDepartmentGUIDBody
	/*ClinicGUID
	  GUID клиники

	*/
	ClinicGUID string
	/*DepartmentGUID
	  GUID подразделения клиники

	*/
	DepartmentGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put clinics clinic GUID departments department GUID params
func (o *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams) WithTimeout(timeout time.Duration) *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put clinics clinic GUID departments department GUID params
func (o *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put clinics clinic GUID departments department GUID params
func (o *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams) WithContext(ctx context.Context) *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put clinics clinic GUID departments department GUID params
func (o *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put clinics clinic GUID departments department GUID params
func (o *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams) WithHTTPClient(client *http.Client) *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put clinics clinic GUID departments department GUID params
func (o *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put clinics clinic GUID departments department GUID params
func (o *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams) WithBody(body PutClinicsClinicGUIDDepartmentsDepartmentGUIDBody) *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put clinics clinic GUID departments department GUID params
func (o *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams) SetBody(body PutClinicsClinicGUIDDepartmentsDepartmentGUIDBody) {
	o.Body = body
}

// WithClinicGUID adds the clinicGUID to the put clinics clinic GUID departments department GUID params
func (o *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams) WithClinicGUID(clinicGUID string) *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams {
	o.SetClinicGUID(clinicGUID)
	return o
}

// SetClinicGUID adds the clinicGuid to the put clinics clinic GUID departments department GUID params
func (o *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams) SetClinicGUID(clinicGUID string) {
	o.ClinicGUID = clinicGUID
}

// WithDepartmentGUID adds the departmentGUID to the put clinics clinic GUID departments department GUID params
func (o *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams) WithDepartmentGUID(departmentGUID string) *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams {
	o.SetDepartmentGUID(departmentGUID)
	return o
}

// SetDepartmentGUID adds the departmentGuid to the put clinics clinic GUID departments department GUID params
func (o *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams) SetDepartmentGUID(departmentGUID string) {
	o.DepartmentGUID = departmentGUID
}

// WriteToRequest writes these params to a swagger request
func (o *PutClinicsClinicGUIDDepartmentsDepartmentGUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param clinicGUID
	if err := r.SetPathParam("clinicGUID", o.ClinicGUID); err != nil {
		return err
	}

	// path param departmentGUID
	if err := r.SetPathParam("departmentGUID", o.DepartmentGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}