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

// GetEmployeesReader is a Reader for the GetEmployees structure.
type GetEmployeesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmployeesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEmployeesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmployeesOK creates a GetEmployeesOK with default headers values
func NewGetEmployeesOK() *GetEmployeesOK {
	return &GetEmployeesOK{}
}

/*GetEmployeesOK handles this case with default header values.

Коллекция сотрудников
*/
type GetEmployeesOK struct {
	Payload *GetEmployeesOKBody
}

func (o *GetEmployeesOK) Error() string {
	return fmt.Sprintf("[GET /employees][%d] getEmployeesOK  %+v", 200, o.Payload)
}

func (o *GetEmployeesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEmployeesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DataItems0 data items0
swagger:model DataItems0
*/
type DataItems0 struct {
	models.MainData

	models.ExtendedData

	models.EmployeeObject

	// class
	Class interface{} `json:"class,omitempty"`

	// Коллекция клиник сотрудника
	ClinicGUID []string `json:"clinicGUID"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DataItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 models.MainData
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	o.MainData = aO0

	// AO1
	var aO1 models.ExtendedData
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	o.ExtendedData = aO1

	// AO2
	var aO2 models.EmployeeObject
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	o.EmployeeObject = aO2

	// AO3
	var dataAO3 struct {
		Class interface{} `json:"class,omitempty"`

		ClinicGUID []string `json:"clinicGUID,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO3); err != nil {
		return err
	}

	o.Class = dataAO3.Class

	o.ClinicGUID = dataAO3.ClinicGUID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DataItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 4)

	aO0, err := swag.WriteJSON(o.MainData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(o.ExtendedData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(o.EmployeeObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	var dataAO3 struct {
		Class interface{} `json:"class,omitempty"`

		ClinicGUID []string `json:"clinicGUID,omitempty"`
	}

	dataAO3.Class = o.Class

	dataAO3.ClinicGUID = o.ClinicGUID

	jsonDataAO3, errAO3 := swag.WriteJSON(dataAO3)
	if errAO3 != nil {
		return nil, errAO3
	}
	_parts = append(_parts, jsonDataAO3)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this data items0
func (o *DataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.MainData
	if err := o.MainData.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.ExtendedData
	if err := o.ExtendedData.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.EmployeeObject
	if err := o.EmployeeObject.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *DataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DataItems0) UnmarshalBinary(b []byte) error {
	var res DataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetEmployeesOKBody get employees o k body
swagger:model GetEmployeesOKBody
*/
type GetEmployeesOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetEmployeesOKBody) UnmarshalJSON(raw []byte) error {
	// GetEmployeesOKBodyAO0
	var getEmployeesOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &getEmployeesOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = getEmployeesOKBodyAO0

	// GetEmployeesOKBodyAO1
	var dataGetEmployeesOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetEmployeesOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetEmployeesOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetEmployeesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getEmployeesOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getEmployeesOKBodyAO0)

	var dataGetEmployeesOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}

	dataGetEmployeesOKBodyAO1.Data = o.Data

	jsonDataGetEmployeesOKBodyAO1, errGetEmployeesOKBodyAO1 := swag.WriteJSON(dataGetEmployeesOKBodyAO1)
	if errGetEmployeesOKBodyAO1 != nil {
		return nil, errGetEmployeesOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetEmployeesOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get employees o k body
func (o *GetEmployeesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetEmployeesOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("getEmployeesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEmployeesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEmployeesOKBody) UnmarshalBinary(b []byte) error {
	var res GetEmployeesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
