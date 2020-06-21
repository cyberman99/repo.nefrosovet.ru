// Code generated by go-swagger; DO NOT EDIT.

package utils

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

// PostUnlinkPatientGUIDReader is a Reader for the PostUnlinkPatientGUID structure.
type PostUnlinkPatientGUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUnlinkPatientGUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostUnlinkPatientGUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostUnlinkPatientGUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostUnlinkPatientGUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPostUnlinkPatientGUIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUnlinkPatientGUIDOK creates a PostUnlinkPatientGUIDOK with default headers values
func NewPostUnlinkPatientGUIDOK() *PostUnlinkPatientGUIDOK {
	return &PostUnlinkPatientGUIDOK{}
}

/*PostUnlinkPatientGUIDOK handles this case with default header values.

Коллекция пациентов
*/
type PostUnlinkPatientGUIDOK struct {
	Payload *PostUnlinkPatientGUIDOKBody
}

func (o *PostUnlinkPatientGUIDOK) Error() string {
	return fmt.Sprintf("[POST /unlink/{patientGUID}][%d] postUnlinkPatientGuidOK  %+v", 200, o.Payload)
}

func (o *PostUnlinkPatientGUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostUnlinkPatientGUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUnlinkPatientGUIDBadRequest creates a PostUnlinkPatientGUIDBadRequest with default headers values
func NewPostUnlinkPatientGUIDBadRequest() *PostUnlinkPatientGUIDBadRequest {
	return &PostUnlinkPatientGUIDBadRequest{}
}

/*PostUnlinkPatientGUIDBadRequest handles this case with default header values.

Validation error
*/
type PostUnlinkPatientGUIDBadRequest struct {
	Payload *PostUnlinkPatientGUIDBadRequestBody
}

func (o *PostUnlinkPatientGUIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /unlink/{patientGUID}][%d] postUnlinkPatientGuidBadRequest  %+v", 400, o.Payload)
}

func (o *PostUnlinkPatientGUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostUnlinkPatientGUIDBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUnlinkPatientGUIDNotFound creates a PostUnlinkPatientGUIDNotFound with default headers values
func NewPostUnlinkPatientGUIDNotFound() *PostUnlinkPatientGUIDNotFound {
	return &PostUnlinkPatientGUIDNotFound{}
}

/*PostUnlinkPatientGUIDNotFound handles this case with default header values.

Not found
*/
type PostUnlinkPatientGUIDNotFound struct {
	Payload *PostUnlinkPatientGUIDNotFoundBody
}

func (o *PostUnlinkPatientGUIDNotFound) Error() string {
	return fmt.Sprintf("[POST /unlink/{patientGUID}][%d] postUnlinkPatientGuidNotFound  %+v", 404, o.Payload)
}

func (o *PostUnlinkPatientGUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostUnlinkPatientGUIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUnlinkPatientGUIDMethodNotAllowed creates a PostUnlinkPatientGUIDMethodNotAllowed with default headers values
func NewPostUnlinkPatientGUIDMethodNotAllowed() *PostUnlinkPatientGUIDMethodNotAllowed {
	return &PostUnlinkPatientGUIDMethodNotAllowed{}
}

/*PostUnlinkPatientGUIDMethodNotAllowed handles this case with default header values.

Invalid Method
*/
type PostUnlinkPatientGUIDMethodNotAllowed struct {
	Payload *PostUnlinkPatientGUIDMethodNotAllowedBody
}

func (o *PostUnlinkPatientGUIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /unlink/{patientGUID}][%d] postUnlinkPatientGuidMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostUnlinkPatientGUIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostUnlinkPatientGUIDMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostUnlinkPatientGUIDBadRequestBody post unlink patient GUID bad request body
swagger:model PostUnlinkPatientGUIDBadRequestBody
*/
type PostUnlinkPatientGUIDBadRequestBody struct {
	models.Error400Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostUnlinkPatientGUIDBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PostUnlinkPatientGUIDBadRequestBodyAO0
	var postUnlinkPatientGUIDBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &postUnlinkPatientGUIDBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = postUnlinkPatientGUIDBadRequestBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostUnlinkPatientGUIDBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postUnlinkPatientGUIDBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postUnlinkPatientGUIDBadRequestBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post unlink patient GUID bad request body
func (o *PostUnlinkPatientGUIDBadRequestBody) Validate(formats strfmt.Registry) error {
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
func (o *PostUnlinkPatientGUIDBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostUnlinkPatientGUIDBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostUnlinkPatientGUIDBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostUnlinkPatientGUIDMethodNotAllowedBody post unlink patient GUID method not allowed body
swagger:model PostUnlinkPatientGUIDMethodNotAllowedBody
*/
type PostUnlinkPatientGUIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostUnlinkPatientGUIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PostUnlinkPatientGUIDMethodNotAllowedBodyAO0
	var postUnlinkPatientGUIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &postUnlinkPatientGUIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = postUnlinkPatientGUIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostUnlinkPatientGUIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postUnlinkPatientGUIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postUnlinkPatientGUIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post unlink patient GUID method not allowed body
func (o *PostUnlinkPatientGUIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PostUnlinkPatientGUIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostUnlinkPatientGUIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PostUnlinkPatientGUIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostUnlinkPatientGUIDNotFoundBody post unlink patient GUID not found body
swagger:model PostUnlinkPatientGUIDNotFoundBody
*/
type PostUnlinkPatientGUIDNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostUnlinkPatientGUIDNotFoundBody) UnmarshalJSON(raw []byte) error {
	// PostUnlinkPatientGUIDNotFoundBodyAO0
	var postUnlinkPatientGUIDNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &postUnlinkPatientGUIDNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = postUnlinkPatientGUIDNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostUnlinkPatientGUIDNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postUnlinkPatientGUIDNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postUnlinkPatientGUIDNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post unlink patient GUID not found body
func (o *PostUnlinkPatientGUIDNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *PostUnlinkPatientGUIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostUnlinkPatientGUIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PostUnlinkPatientGUIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostUnlinkPatientGUIDOKBody post unlink patient GUID o k body
swagger:model PostUnlinkPatientGUIDOKBody
*/
type PostUnlinkPatientGUIDOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostUnlinkPatientGUIDOKBody) UnmarshalJSON(raw []byte) error {
	// PostUnlinkPatientGUIDOKBodyAO0
	var postUnlinkPatientGUIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postUnlinkPatientGUIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postUnlinkPatientGUIDOKBodyAO0

	// PostUnlinkPatientGUIDOKBodyAO1
	var dataPostUnlinkPatientGUIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostUnlinkPatientGUIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostUnlinkPatientGUIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostUnlinkPatientGUIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postUnlinkPatientGUIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postUnlinkPatientGUIDOKBodyAO0)

	var dataPostUnlinkPatientGUIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}

	dataPostUnlinkPatientGUIDOKBodyAO1.Data = o.Data

	jsonDataPostUnlinkPatientGUIDOKBodyAO1, errPostUnlinkPatientGUIDOKBodyAO1 := swag.WriteJSON(dataPostUnlinkPatientGUIDOKBodyAO1)
	if errPostUnlinkPatientGUIDOKBodyAO1 != nil {
		return nil, errPostUnlinkPatientGUIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostUnlinkPatientGUIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post unlink patient GUID o k body
func (o *PostUnlinkPatientGUIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostUnlinkPatientGUIDOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("postUnlinkPatientGuidOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostUnlinkPatientGUIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostUnlinkPatientGUIDOKBody) UnmarshalBinary(b []byte) error {
	var res PostUnlinkPatientGUIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}