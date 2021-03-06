// Code generated by go-swagger; DO NOT EDIT.

package services

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

// GetClinicsClinicGUIDServicesServiceGUIDReader is a Reader for the GetClinicsClinicGUIDServicesServiceGUID structure.
type GetClinicsClinicGUIDServicesServiceGUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClinicsClinicGUIDServicesServiceGUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClinicsClinicGUIDServicesServiceGUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetClinicsClinicGUIDServicesServiceGUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClinicsClinicGUIDServicesServiceGUIDOK creates a GetClinicsClinicGUIDServicesServiceGUIDOK with default headers values
func NewGetClinicsClinicGUIDServicesServiceGUIDOK() *GetClinicsClinicGUIDServicesServiceGUIDOK {
	return &GetClinicsClinicGUIDServicesServiceGUIDOK{}
}

/*GetClinicsClinicGUIDServicesServiceGUIDOK handles this case with default header values.

Коллекция сервисов
*/
type GetClinicsClinicGUIDServicesServiceGUIDOK struct {
	Payload *GetClinicsClinicGUIDServicesServiceGUIDOKBody
}

func (o *GetClinicsClinicGUIDServicesServiceGUIDOK) Error() string {
	return fmt.Sprintf("[GET /clinics/{clinicGUID}/services/{serviceGUID}][%d] getClinicsClinicGuidServicesServiceGuidOK  %+v", 200, o.Payload)
}

func (o *GetClinicsClinicGUIDServicesServiceGUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetClinicsClinicGUIDServicesServiceGUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClinicsClinicGUIDServicesServiceGUIDNotFound creates a GetClinicsClinicGUIDServicesServiceGUIDNotFound with default headers values
func NewGetClinicsClinicGUIDServicesServiceGUIDNotFound() *GetClinicsClinicGUIDServicesServiceGUIDNotFound {
	return &GetClinicsClinicGUIDServicesServiceGUIDNotFound{}
}

/*GetClinicsClinicGUIDServicesServiceGUIDNotFound handles this case with default header values.

Not found
*/
type GetClinicsClinicGUIDServicesServiceGUIDNotFound struct {
	Payload *GetClinicsClinicGUIDServicesServiceGUIDNotFoundBody
}

func (o *GetClinicsClinicGUIDServicesServiceGUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /clinics/{clinicGUID}/services/{serviceGUID}][%d] getClinicsClinicGuidServicesServiceGuidNotFound  %+v", 404, o.Payload)
}

func (o *GetClinicsClinicGUIDServicesServiceGUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetClinicsClinicGUIDServicesServiceGUIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetClinicsClinicGUIDServicesServiceGUIDNotFoundBody get clinics clinic GUID services service GUID not found body
swagger:model GetClinicsClinicGUIDServicesServiceGUIDNotFoundBody
*/
type GetClinicsClinicGUIDServicesServiceGUIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetClinicsClinicGUIDServicesServiceGUIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// GetClinicsClinicGUIDServicesServiceGUIDNotFoundBodyAO0
	var getClinicsClinicGUIDServicesServiceGUIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &getClinicsClinicGUIDServicesServiceGUIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = getClinicsClinicGUIDServicesServiceGUIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetClinicsClinicGUIDServicesServiceGUIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getClinicsClinicGUIDServicesServiceGUIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getClinicsClinicGUIDServicesServiceGUIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get clinics clinic GUID services service GUID not found body
func (o *GetClinicsClinicGUIDServicesServiceGUIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *GetClinicsClinicGUIDServicesServiceGUIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetClinicsClinicGUIDServicesServiceGUIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetClinicsClinicGUIDServicesServiceGUIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetClinicsClinicGUIDServicesServiceGUIDOKBody get clinics clinic GUID services service GUID o k body
swagger:model GetClinicsClinicGUIDServicesServiceGUIDOKBody
*/
type GetClinicsClinicGUIDServicesServiceGUIDOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetClinicsClinicGUIDServicesServiceGUIDOKBody) UnmarshalJSON(raw []byte) error {
	// GetClinicsClinicGUIDServicesServiceGUIDOKBodyAO0
	var getClinicsClinicGUIDServicesServiceGUIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &getClinicsClinicGUIDServicesServiceGUIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = getClinicsClinicGUIDServicesServiceGUIDOKBodyAO0

	// GetClinicsClinicGUIDServicesServiceGUIDOKBodyAO1
	var dataGetClinicsClinicGUIDServicesServiceGUIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetClinicsClinicGUIDServicesServiceGUIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetClinicsClinicGUIDServicesServiceGUIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetClinicsClinicGUIDServicesServiceGUIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getClinicsClinicGUIDServicesServiceGUIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getClinicsClinicGUIDServicesServiceGUIDOKBodyAO0)

	var dataGetClinicsClinicGUIDServicesServiceGUIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}

	dataGetClinicsClinicGUIDServicesServiceGUIDOKBodyAO1.Data = o.Data

	jsonDataGetClinicsClinicGUIDServicesServiceGUIDOKBodyAO1, errGetClinicsClinicGUIDServicesServiceGUIDOKBodyAO1 := swag.WriteJSON(dataGetClinicsClinicGUIDServicesServiceGUIDOKBodyAO1)
	if errGetClinicsClinicGUIDServicesServiceGUIDOKBodyAO1 != nil {
		return nil, errGetClinicsClinicGUIDServicesServiceGUIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetClinicsClinicGUIDServicesServiceGUIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get clinics clinic GUID services service GUID o k body
func (o *GetClinicsClinicGUIDServicesServiceGUIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetClinicsClinicGUIDServicesServiceGUIDOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("getClinicsClinicGuidServicesServiceGuidOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetClinicsClinicGUIDServicesServiceGUIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetClinicsClinicGUIDServicesServiceGUIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetClinicsClinicGUIDServicesServiceGUIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
