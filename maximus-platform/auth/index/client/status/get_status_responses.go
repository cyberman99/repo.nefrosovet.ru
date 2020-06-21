// Code generated by go-swagger; DO NOT EDIT.

package status

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

// GetStatusReader is a Reader for the GetStatus structure.
type GetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStatusOK creates a GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {
	return &GetStatusOK{}
}

/*GetStatusOK handles this case with default header values.

Статус инстанса
*/
type GetStatusOK struct {
	Payload *GetStatusOKBody
}

func (o *GetStatusOK) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatusOK  %+v", 200, o.Payload)
}

func (o *GetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusInternalServerError creates a GetStatusInternalServerError with default headers values
func NewGetStatusInternalServerError() *GetStatusInternalServerError {
	return &GetStatusInternalServerError{}
}

/*GetStatusInternalServerError handles this case with default header values.

Internal Server error
*/
type GetStatusInternalServerError struct {
}

func (o *GetStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatusInternalServerError ", 500)
}

func (o *GetStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DataItems0 data items0
swagger:model DataItems0
*/
type DataItems0 struct {
	models.StatusData
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *DataItems0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 models.StatusData
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	o.StatusData = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o DataItems0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(o.StatusData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this data items0
func (o *DataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.StatusData
	if err := o.StatusData.Validate(formats); err != nil {
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

/*GetStatusOKBody get status o k body
swagger:model GetStatusOKBody
*/
type GetStatusOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetStatusOKBody) UnmarshalJSON(raw []byte) error {
	// GetStatusOKBodyAO0
	var getStatusOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &getStatusOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = getStatusOKBodyAO0

	// GetStatusOKBodyAO1
	var dataGetStatusOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetStatusOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataGetStatusOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetStatusOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getStatusOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getStatusOKBodyAO0)

	var dataGetStatusOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}

	dataGetStatusOKBodyAO1.Data = o.Data

	jsonDataGetStatusOKBodyAO1, errGetStatusOKBodyAO1 := swag.WriteJSON(dataGetStatusOKBodyAO1)
	if errGetStatusOKBodyAO1 != nil {
		return nil, errGetStatusOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetStatusOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get status o k body
func (o *GetStatusOKBody) Validate(formats strfmt.Registry) error {
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

func (o *GetStatusOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("getStatusOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStatusOKBody) UnmarshalBinary(b []byte) error {
	var res GetStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
