// Code generated by go-swagger; DO NOT EDIT.

package employees

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

// GetEmployeesEmployeeGUIDReader is a Reader for the GetEmployeesEmployeeGUID structure.
type GetEmployeesEmployeeGUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmployeesEmployeeGUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEmployeesEmployeeGUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetEmployeesEmployeeGUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmployeesEmployeeGUIDOK creates a GetEmployeesEmployeeGUIDOK with default headers values
func NewGetEmployeesEmployeeGUIDOK() *GetEmployeesEmployeeGUIDOK {
	return &GetEmployeesEmployeeGUIDOK{}
}

/*GetEmployeesEmployeeGUIDOK handles this case with default header values.

Коллекция сотрудников
*/
type GetEmployeesEmployeeGUIDOK struct {
	Payload *GetEmployeesEmployeeGUIDOKBody
}

func (o *GetEmployeesEmployeeGUIDOK) Error() string {
	return fmt.Sprintf("[GET /employees/{employeeGUID}][%d] getEmployeesEmployeeGuidOK  %+v", 200, o.Payload)
}

func (o *GetEmployeesEmployeeGUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEmployeesEmployeeGUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmployeesEmployeeGUIDNotFound creates a GetEmployeesEmployeeGUIDNotFound with default headers values
func NewGetEmployeesEmployeeGUIDNotFound() *GetEmployeesEmployeeGUIDNotFound {
	return &GetEmployeesEmployeeGUIDNotFound{}
}

/*GetEmployeesEmployeeGUIDNotFound handles this case with default header values.

Not found
*/
type GetEmployeesEmployeeGUIDNotFound struct {
	Payload *GetEmployeesEmployeeGUIDNotFoundBody
}

func (o *GetEmployeesEmployeeGUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /employees/{employeeGUID}][%d] getEmployeesEmployeeGuidNotFound  %+v", 404, o.Payload)
}

func (o *GetEmployeesEmployeeGUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEmployeesEmployeeGUIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetEmployeesEmployeeGUIDNotFoundBody get employees employee GUID not found body
swagger:model GetEmployeesEmployeeGUIDNotFoundBody
*/
type GetEmployeesEmployeeGUIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetEmployeesEmployeeGUIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// GetEmployeesEmployeeGUIDNotFoundBodyAO0
	var getEmployeesEmployeeGUIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &getEmployeesEmployeeGUIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = getEmployeesEmployeeGUIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetEmployeesEmployeeGUIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getEmployeesEmployeeGUIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getEmployeesEmployeeGUIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get employees employee GUID not found body
func (o *GetEmployeesEmployeeGUIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *GetEmployeesEmployeeGUIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEmployeesEmployeeGUIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetEmployeesEmployeeGUIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetEmployeesEmployeeGUIDOKBody get employees employee GUID o k body
swagger:model GetEmployeesEmployeeGUIDOKBody
*/
type GetEmployeesEmployeeGUIDOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetEmployeesEmployeeGUIDOKBody) UnmarshalJSON(raw []byte) error {
	// GetEmployeesEmployeeGUIDOKBodyAO0
	var getEmployeesEmployeeGUIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &getEmployeesEmployeeGUIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = getEmployeesEmployeeGUIDOKBodyAO0

	// GetEmployeesEmployeeGUIDOKBodyAO1
	var dataGetEmployeesEmployeeGUIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetEmployeesEmployeeGUIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetEmployeesEmployeeGUIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetEmployeesEmployeeGUIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getEmployeesEmployeeGUIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getEmployeesEmployeeGUIDOKBodyAO0)

	var dataGetEmployeesEmployeeGUIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}

	dataGetEmployeesEmployeeGUIDOKBodyAO1.Data = o.Data

	jsonDataGetEmployeesEmployeeGUIDOKBodyAO1, errGetEmployeesEmployeeGUIDOKBodyAO1 := swag.WriteJSON(dataGetEmployeesEmployeeGUIDOKBodyAO1)
	if errGetEmployeesEmployeeGUIDOKBodyAO1 != nil {
		return nil, errGetEmployeesEmployeeGUIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetEmployeesEmployeeGUIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get employees employee GUID o k body
func (o *GetEmployeesEmployeeGUIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetEmployeesEmployeeGUIDOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("getEmployeesEmployeeGuidOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEmployeesEmployeeGUIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEmployeesEmployeeGUIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetEmployeesEmployeeGUIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
