// Code generated by go-swagger; DO NOT EDIT.

package patients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "repo.nefrosovet.ru/maximus-platform/auth/index/models"
)

// GetPatientsPatientGUIDReader is a Reader for the GetPatientsPatientGUID structure.
type GetPatientsPatientGUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPatientsPatientGUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPatientsPatientGUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetPatientsPatientGUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPatientsPatientGUIDOK creates a GetPatientsPatientGUIDOK with default headers values
func NewGetPatientsPatientGUIDOK() *GetPatientsPatientGUIDOK {
	return &GetPatientsPatientGUIDOK{}
}

/*GetPatientsPatientGUIDOK handles this case with default header values.

Коллекция пациентов
*/
type GetPatientsPatientGUIDOK struct {
	Payload *GetPatientsPatientGUIDOKBody
}

func (o *GetPatientsPatientGUIDOK) Error() string {
	return fmt.Sprintf("[GET /patients/{patientGUID}][%d] getPatientsPatientGuidOK  %+v", 200, o.Payload)
}

func (o *GetPatientsPatientGUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPatientsPatientGUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPatientsPatientGUIDNotFound creates a GetPatientsPatientGUIDNotFound with default headers values
func NewGetPatientsPatientGUIDNotFound() *GetPatientsPatientGUIDNotFound {
	return &GetPatientsPatientGUIDNotFound{}
}

/*GetPatientsPatientGUIDNotFound handles this case with default header values.

Not found
*/
type GetPatientsPatientGUIDNotFound struct {
	Payload *GetPatientsPatientGUIDNotFoundBody
}

func (o *GetPatientsPatientGUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /patients/{patientGUID}][%d] getPatientsPatientGuidNotFound  %+v", 404, o.Payload)
}

func (o *GetPatientsPatientGUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPatientsPatientGUIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetPatientsPatientGUIDNotFoundBody get patients patient GUID not found body
swagger:model GetPatientsPatientGUIDNotFoundBody
*/
type GetPatientsPatientGUIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetPatientsPatientGUIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// GetPatientsPatientGUIDNotFoundBodyAO0
	var getPatientsPatientGUIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &getPatientsPatientGUIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = getPatientsPatientGUIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetPatientsPatientGUIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getPatientsPatientGUIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getPatientsPatientGUIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get patients patient GUID not found body
func (o *GetPatientsPatientGUIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *GetPatientsPatientGUIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPatientsPatientGUIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetPatientsPatientGUIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetPatientsPatientGUIDOKBody get patients patient GUID o k body
swagger:model GetPatientsPatientGUIDOKBody
*/
type GetPatientsPatientGUIDOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetPatientsPatientGUIDOKBody) UnmarshalJSON(raw []byte) error {
	// GetPatientsPatientGUIDOKBodyAO0
	var getPatientsPatientGUIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &getPatientsPatientGUIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = getPatientsPatientGUIDOKBodyAO0

	// GetPatientsPatientGUIDOKBodyAO1
	var dataGetPatientsPatientGUIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetPatientsPatientGUIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetPatientsPatientGUIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetPatientsPatientGUIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getPatientsPatientGUIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getPatientsPatientGUIDOKBodyAO0)

	var dataGetPatientsPatientGUIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}

	dataGetPatientsPatientGUIDOKBodyAO1.Data = o.Data

	jsonDataGetPatientsPatientGUIDOKBodyAO1, errGetPatientsPatientGUIDOKBodyAO1 := swag.WriteJSON(dataGetPatientsPatientGUIDOKBodyAO1)
	if errGetPatientsPatientGUIDOKBodyAO1 != nil {
		return nil, errGetPatientsPatientGUIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetPatientsPatientGUIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get patients patient GUID o k body
func (o *GetPatientsPatientGUIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.SuccessData
	if err := o.SuccessData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPatientsPatientGUIDOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPatientsPatientGuidOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPatientsPatientGUIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPatientsPatientGUIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetPatientsPatientGUIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
