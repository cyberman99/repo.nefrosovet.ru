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

// PostPatientsReader is a Reader for the PostPatients structure.
type PostPatientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPatientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPatientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPatientsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPostPatientsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPatientsOK creates a PostPatientsOK with default headers values
func NewPostPatientsOK() *PostPatientsOK {
	return &PostPatientsOK{}
}

/*PostPatientsOK handles this case with default header values.

Коллекция пациентов
*/
type PostPatientsOK struct {
	Payload *PostPatientsOKBody
}

func (o *PostPatientsOK) Error() string {
	return fmt.Sprintf("[POST /patients][%d] postPatientsOK  %+v", 200, o.Payload)
}

func (o *PostPatientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostPatientsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPatientsBadRequest creates a PostPatientsBadRequest with default headers values
func NewPostPatientsBadRequest() *PostPatientsBadRequest {
	return &PostPatientsBadRequest{}
}

/*PostPatientsBadRequest handles this case with default header values.

Validation error
*/
type PostPatientsBadRequest struct {
	Payload *PostPatientsBadRequestBody
}

func (o *PostPatientsBadRequest) Error() string {
	return fmt.Sprintf("[POST /patients][%d] postPatientsBadRequest  %+v", 400, o.Payload)
}

func (o *PostPatientsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostPatientsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPatientsMethodNotAllowed creates a PostPatientsMethodNotAllowed with default headers values
func NewPostPatientsMethodNotAllowed() *PostPatientsMethodNotAllowed {
	return &PostPatientsMethodNotAllowed{}
}

/*PostPatientsMethodNotAllowed handles this case with default header values.

Invalid Method
*/
type PostPatientsMethodNotAllowed struct {
	Payload *PostPatientsMethodNotAllowedBody
}

func (o *PostPatientsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /patients][%d] postPatientsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostPatientsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostPatientsMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostPatientsBadRequestBody post patients bad request body
swagger:model PostPatientsBadRequestBody
*/
type PostPatientsBadRequestBody struct {
	models.Error400Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostPatientsBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PostPatientsBadRequestBodyAO0
	var postPatientsBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &postPatientsBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = postPatientsBadRequestBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostPatientsBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postPatientsBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postPatientsBadRequestBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post patients bad request body
func (o *PostPatientsBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error400Data
	if err := o.Error400Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostPatientsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPatientsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostPatientsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostPatientsBody post patients body
swagger:model PostPatientsBody
*/
type PostPatientsBody struct {
	models.MainData

	models.PatientObject

	models.PasswordObject
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostPatientsBody) UnmarshalJSON(raw []byte) error {
	// PostPatientsParamsBodyAO0
	var postPatientsParamsBodyAO0 models.MainData
	if err := swag.ReadJSON(raw, &postPatientsParamsBodyAO0); err != nil {
		return err
	}
	o.MainData = postPatientsParamsBodyAO0

	// PostPatientsParamsBodyAO1
	var postPatientsParamsBodyAO1 models.PatientObject
	if err := swag.ReadJSON(raw, &postPatientsParamsBodyAO1); err != nil {
		return err
	}
	o.PatientObject = postPatientsParamsBodyAO1

	// PostPatientsParamsBodyAO2
	var postPatientsParamsBodyAO2 models.PasswordObject
	if err := swag.ReadJSON(raw, &postPatientsParamsBodyAO2); err != nil {
		return err
	}
	o.PasswordObject = postPatientsParamsBodyAO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostPatientsBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	postPatientsParamsBodyAO0, err := swag.WriteJSON(o.MainData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postPatientsParamsBodyAO0)

	postPatientsParamsBodyAO1, err := swag.WriteJSON(o.PatientObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postPatientsParamsBodyAO1)

	postPatientsParamsBodyAO2, err := swag.WriteJSON(o.PasswordObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postPatientsParamsBodyAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post patients body
func (o *PostPatientsBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.MainData
	if err := o.MainData.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.PatientObject
	if err := o.PatientObject.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.PasswordObject
	if err := o.PasswordObject.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostPatientsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPatientsBody) UnmarshalBinary(b []byte) error {
	var res PostPatientsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostPatientsMethodNotAllowedBody post patients method not allowed body
swagger:model PostPatientsMethodNotAllowedBody
*/
type PostPatientsMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostPatientsMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PostPatientsMethodNotAllowedBodyAO0
	var postPatientsMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &postPatientsMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = postPatientsMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostPatientsMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postPatientsMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postPatientsMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post patients method not allowed body
func (o *PostPatientsMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error405Data
	if err := o.Error405Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostPatientsMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPatientsMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PostPatientsMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostPatientsOKBody post patients o k body
swagger:model PostPatientsOKBody
*/
type PostPatientsOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostPatientsOKBody) UnmarshalJSON(raw []byte) error {
	// PostPatientsOKBodyAO0
	var postPatientsOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postPatientsOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postPatientsOKBodyAO0

	// PostPatientsOKBodyAO1
	var dataPostPatientsOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostPatientsOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostPatientsOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostPatientsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postPatientsOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postPatientsOKBodyAO0)

	var dataPostPatientsOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}

	dataPostPatientsOKBodyAO1.Data = o.Data

	jsonDataPostPatientsOKBodyAO1, errPostPatientsOKBodyAO1 := swag.WriteJSON(dataPostPatientsOKBodyAO1)
	if errPostPatientsOKBodyAO1 != nil {
		return nil, errPostPatientsOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostPatientsOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post patients o k body
func (o *PostPatientsOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostPatientsOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("postPatientsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostPatientsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostPatientsOKBody) UnmarshalBinary(b []byte) error {
	var res PostPatientsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
