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

// GetPatientsReader is a Reader for the GetPatients structure.
type GetPatientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPatientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPatientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPatientsOK creates a GetPatientsOK with default headers values
func NewGetPatientsOK() *GetPatientsOK {
	return &GetPatientsOK{}
}

/*GetPatientsOK handles this case with default header values.

Коллекция пациентов
*/
type GetPatientsOK struct {
	Payload *GetPatientsOKBody
}

func (o *GetPatientsOK) Error() string {
	return fmt.Sprintf("[GET /patients][%d] getPatientsOK  %+v", 200, o.Payload)
}

func (o *GetPatientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPatientsOKBody)

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

	models.PatientObject

	// class
	Class interface{} `json:"class,omitempty"`

	// source GUID
	SourceGUID []string `json:"sourceGUID"`

	// Ссылка на GUID смерженной записи
	TargetGUID string `json:"targetGUID,omitempty"`
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
	var aO2 models.PatientObject
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	o.PatientObject = aO2

	// AO3
	var dataAO3 struct {
		Class interface{} `json:"class,omitempty"`

		SourceGUID []string `json:"sourceGUID,omitempty"`

		TargetGUID string `json:"targetGUID,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO3); err != nil {
		return err
	}

	o.Class = dataAO3.Class

	o.SourceGUID = dataAO3.SourceGUID

	o.TargetGUID = dataAO3.TargetGUID

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

	aO2, err := swag.WriteJSON(o.PatientObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	var dataAO3 struct {
		Class interface{} `json:"class,omitempty"`

		SourceGUID []string `json:"sourceGUID,omitempty"`

		TargetGUID string `json:"targetGUID,omitempty"`
	}

	dataAO3.Class = o.Class

	dataAO3.SourceGUID = o.SourceGUID

	dataAO3.TargetGUID = o.TargetGUID

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
	// validation for a type composition with models.PatientObject
	if err := o.PatientObject.Validate(formats); err != nil {
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

/*GetPatientsOKBody get patients o k body
swagger:model GetPatientsOKBody
*/
type GetPatientsOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetPatientsOKBody) UnmarshalJSON(raw []byte) error {
	// GetPatientsOKBodyAO0
	var getPatientsOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &getPatientsOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = getPatientsOKBodyAO0

	// GetPatientsOKBodyAO1
	var dataGetPatientsOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetPatientsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetPatientsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetPatientsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getPatientsOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getPatientsOKBodyAO0)

	var dataGetPatientsOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}

	dataGetPatientsOKBodyAO1.Data = o.Data

	jsonDataGetPatientsOKBodyAO1, errGetPatientsOKBodyAO1 := swag.WriteJSON(dataGetPatientsOKBodyAO1)
	if errGetPatientsOKBodyAO1 != nil {
		return nil, errGetPatientsOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetPatientsOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get patients o k body
func (o *GetPatientsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetPatientsOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("getPatientsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPatientsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPatientsOKBody) UnmarshalBinary(b []byte) error {
	var res GetPatientsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
