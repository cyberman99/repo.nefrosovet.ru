// Code generated by go-swagger; DO NOT EDIT.

package clinic_employees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "repo.nefrosovet.ru/maximus-platform/auth/index/models"
)

// PostClinicsClinicGUIDEmployeesEmployeeGUIDReader is a Reader for the PostClinicsClinicGUIDEmployeesEmployeeGUID structure.
type PostClinicsClinicGUIDEmployeesEmployeeGUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostClinicsClinicGUIDEmployeesEmployeeGUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostClinicsClinicGUIDEmployeesEmployeeGUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostClinicsClinicGUIDEmployeesEmployeeGUIDOK creates a PostClinicsClinicGUIDEmployeesEmployeeGUIDOK with default headers values
func NewPostClinicsClinicGUIDEmployeesEmployeeGUIDOK() *PostClinicsClinicGUIDEmployeesEmployeeGUIDOK {
	return &PostClinicsClinicGUIDEmployeesEmployeeGUIDOK{}
}

/*PostClinicsClinicGUIDEmployeesEmployeeGUIDOK handles this case with default header values.

Коллекция сотрудников в клинике
*/
type PostClinicsClinicGUIDEmployeesEmployeeGUIDOK struct {
	Payload *PostClinicsClinicGUIDEmployeesEmployeeGUIDOKBody
}

func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDOK) Error() string {
	return fmt.Sprintf("[POST /clinics/{clinicGUID}/employees/{employeeGUID}][%d] postClinicsClinicGuidEmployeesEmployeeGuidOK  %+v", 200, o.Payload)
}

func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostClinicsClinicGUIDEmployeesEmployeeGUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostClinicsClinicGUIDEmployeesEmployeeGUIDNotFound creates a PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFound with default headers values
func NewPostClinicsClinicGUIDEmployeesEmployeeGUIDNotFound() *PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFound {
	return &PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFound{}
}

/*PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFound handles this case with default header values.

Not found
*/
type PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFound struct {
	Payload *PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBody
}

func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFound) Error() string {
	return fmt.Sprintf("[POST /clinics/{clinicGUID}/employees/{employeeGUID}][%d] postClinicsClinicGuidEmployeesEmployeeGuidNotFound  %+v", 404, o.Payload)
}

func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBody post clinics clinic GUID employees employee GUID not found body
swagger:model PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBody
*/
type PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBodyAO0
	var postClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &postClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = postClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post clinics clinic GUID employees employee GUID not found body
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error404Data
	if err := o.Error404Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PostClinicsClinicGUIDEmployeesEmployeeGUIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostClinicsClinicGUIDEmployeesEmployeeGUIDOKBody post clinics clinic GUID employees employee GUID o k body
swagger:model PostClinicsClinicGUIDEmployeesEmployeeGUIDOKBody
*/
type PostClinicsClinicGUIDEmployeesEmployeeGUIDOKBody struct {
	models.SuccessData

	// data
	Data []string `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDOKBody) UnmarshalJSON(raw []byte) error {
	// PostClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO0
	var postClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO0

	// PostClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO1
	var dataPostClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO1 struct {
		Data []string `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostClinicsClinicGUIDEmployeesEmployeeGUIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO0)

	var dataPostClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO1 struct {
		Data []string `json:"data,omitempty"`
	}

	dataPostClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO1.Data = o.Data

	jsonDataPostClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO1, errPostClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO1 := swag.WriteJSON(dataPostClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO1)
	if errPostClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO1 != nil {
		return nil, errPostClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostClinicsClinicGUIDEmployeesEmployeeGUIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post clinics clinic GUID employees employee GUID o k body
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.SuccessData
	if err := o.SuccessData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostClinicsClinicGUIDEmployeesEmployeeGUIDOKBody) UnmarshalBinary(b []byte) error {
	var res PostClinicsClinicGUIDEmployeesEmployeeGUIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
