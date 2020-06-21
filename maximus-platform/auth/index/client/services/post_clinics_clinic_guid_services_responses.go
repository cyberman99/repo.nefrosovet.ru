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

// PostClinicsClinicGUIDServicesReader is a Reader for the PostClinicsClinicGUIDServices structure.
type PostClinicsClinicGUIDServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostClinicsClinicGUIDServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostClinicsClinicGUIDServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostClinicsClinicGUIDServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostClinicsClinicGUIDServicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPostClinicsClinicGUIDServicesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostClinicsClinicGUIDServicesOK creates a PostClinicsClinicGUIDServicesOK with default headers values
func NewPostClinicsClinicGUIDServicesOK() *PostClinicsClinicGUIDServicesOK {
	return &PostClinicsClinicGUIDServicesOK{}
}

/*PostClinicsClinicGUIDServicesOK handles this case with default header values.

Коллекция сервисов
*/
type PostClinicsClinicGUIDServicesOK struct {
	Payload *PostClinicsClinicGUIDServicesOKBody
}

func (o *PostClinicsClinicGUIDServicesOK) Error() string {
	return fmt.Sprintf("[POST /clinics/{clinicGUID}/services][%d] postClinicsClinicGuidServicesOK  %+v", 200, o.Payload)
}

func (o *PostClinicsClinicGUIDServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostClinicsClinicGUIDServicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostClinicsClinicGUIDServicesBadRequest creates a PostClinicsClinicGUIDServicesBadRequest with default headers values
func NewPostClinicsClinicGUIDServicesBadRequest() *PostClinicsClinicGUIDServicesBadRequest {
	return &PostClinicsClinicGUIDServicesBadRequest{}
}

/*PostClinicsClinicGUIDServicesBadRequest handles this case with default header values.

Validation error
*/
type PostClinicsClinicGUIDServicesBadRequest struct {
	Payload *PostClinicsClinicGUIDServicesBadRequestBody
}

func (o *PostClinicsClinicGUIDServicesBadRequest) Error() string {
	return fmt.Sprintf("[POST /clinics/{clinicGUID}/services][%d] postClinicsClinicGuidServicesBadRequest  %+v", 400, o.Payload)
}

func (o *PostClinicsClinicGUIDServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostClinicsClinicGUIDServicesBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostClinicsClinicGUIDServicesNotFound creates a PostClinicsClinicGUIDServicesNotFound with default headers values
func NewPostClinicsClinicGUIDServicesNotFound() *PostClinicsClinicGUIDServicesNotFound {
	return &PostClinicsClinicGUIDServicesNotFound{}
}

/*PostClinicsClinicGUIDServicesNotFound handles this case with default header values.

Not found
*/
type PostClinicsClinicGUIDServicesNotFound struct {
	Payload *PostClinicsClinicGUIDServicesNotFoundBody
}

func (o *PostClinicsClinicGUIDServicesNotFound) Error() string {
	return fmt.Sprintf("[POST /clinics/{clinicGUID}/services][%d] postClinicsClinicGuidServicesNotFound  %+v", 404, o.Payload)
}

func (o *PostClinicsClinicGUIDServicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostClinicsClinicGUIDServicesNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostClinicsClinicGUIDServicesMethodNotAllowed creates a PostClinicsClinicGUIDServicesMethodNotAllowed with default headers values
func NewPostClinicsClinicGUIDServicesMethodNotAllowed() *PostClinicsClinicGUIDServicesMethodNotAllowed {
	return &PostClinicsClinicGUIDServicesMethodNotAllowed{}
}

/*PostClinicsClinicGUIDServicesMethodNotAllowed handles this case with default header values.

Invalid Method
*/
type PostClinicsClinicGUIDServicesMethodNotAllowed struct {
	Payload *PostClinicsClinicGUIDServicesMethodNotAllowedBody
}

func (o *PostClinicsClinicGUIDServicesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /clinics/{clinicGUID}/services][%d] postClinicsClinicGuidServicesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostClinicsClinicGUIDServicesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostClinicsClinicGUIDServicesMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostClinicsClinicGUIDServicesBadRequestBody post clinics clinic GUID services bad request body
swagger:model PostClinicsClinicGUIDServicesBadRequestBody
*/
type PostClinicsClinicGUIDServicesBadRequestBody struct {
	models.Error400Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostClinicsClinicGUIDServicesBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PostClinicsClinicGUIDServicesBadRequestBodyAO0
	var postClinicsClinicGUIDServicesBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &postClinicsClinicGUIDServicesBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = postClinicsClinicGUIDServicesBadRequestBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostClinicsClinicGUIDServicesBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postClinicsClinicGUIDServicesBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postClinicsClinicGUIDServicesBadRequestBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post clinics clinic GUID services bad request body
func (o *PostClinicsClinicGUIDServicesBadRequestBody) Validate(formats strfmt.Registry) error {
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
func (o *PostClinicsClinicGUIDServicesBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostClinicsClinicGUIDServicesBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostClinicsClinicGUIDServicesBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostClinicsClinicGUIDServicesBody post clinics clinic GUID services body
swagger:model PostClinicsClinicGUIDServicesBody
*/
type PostClinicsClinicGUIDServicesBody struct {
	models.MainData

	models.ServiceObject

	PostClinicsClinicGUIDServicesParamsBodyAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostClinicsClinicGUIDServicesBody) UnmarshalJSON(raw []byte) error {
	// PostClinicsClinicGUIDServicesParamsBodyAO0
	var postClinicsClinicGUIDServicesParamsBodyAO0 models.MainData
	if err := swag.ReadJSON(raw, &postClinicsClinicGUIDServicesParamsBodyAO0); err != nil {
		return err
	}
	o.MainData = postClinicsClinicGUIDServicesParamsBodyAO0

	// PostClinicsClinicGUIDServicesParamsBodyAO1
	var postClinicsClinicGUIDServicesParamsBodyAO1 models.ServiceObject
	if err := swag.ReadJSON(raw, &postClinicsClinicGUIDServicesParamsBodyAO1); err != nil {
		return err
	}
	o.ServiceObject = postClinicsClinicGUIDServicesParamsBodyAO1

	// PostClinicsClinicGUIDServicesParamsBodyAO2
	var postClinicsClinicGUIDServicesParamsBodyAO2 PostClinicsClinicGUIDServicesParamsBodyAllOf2
	if err := swag.ReadJSON(raw, &postClinicsClinicGUIDServicesParamsBodyAO2); err != nil {
		return err
	}
	o.PostClinicsClinicGUIDServicesParamsBodyAllOf2 = postClinicsClinicGUIDServicesParamsBodyAO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostClinicsClinicGUIDServicesBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	postClinicsClinicGUIDServicesParamsBodyAO0, err := swag.WriteJSON(o.MainData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postClinicsClinicGUIDServicesParamsBodyAO0)

	postClinicsClinicGUIDServicesParamsBodyAO1, err := swag.WriteJSON(o.ServiceObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postClinicsClinicGUIDServicesParamsBodyAO1)

	postClinicsClinicGUIDServicesParamsBodyAO2, err := swag.WriteJSON(o.PostClinicsClinicGUIDServicesParamsBodyAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postClinicsClinicGUIDServicesParamsBodyAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post clinics clinic GUID services body
func (o *PostClinicsClinicGUIDServicesBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.MainData
	if err := o.MainData.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with models.ServiceObject
	if err := o.ServiceObject.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PostClinicsClinicGUIDServicesParamsBodyAllOf2

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostClinicsClinicGUIDServicesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostClinicsClinicGUIDServicesBody) UnmarshalBinary(b []byte) error {
	var res PostClinicsClinicGUIDServicesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostClinicsClinicGUIDServicesMethodNotAllowedBody post clinics clinic GUID services method not allowed body
swagger:model PostClinicsClinicGUIDServicesMethodNotAllowedBody
*/
type PostClinicsClinicGUIDServicesMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostClinicsClinicGUIDServicesMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PostClinicsClinicGUIDServicesMethodNotAllowedBodyAO0
	var postClinicsClinicGUIDServicesMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &postClinicsClinicGUIDServicesMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = postClinicsClinicGUIDServicesMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostClinicsClinicGUIDServicesMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postClinicsClinicGUIDServicesMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postClinicsClinicGUIDServicesMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post clinics clinic GUID services method not allowed body
func (o *PostClinicsClinicGUIDServicesMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PostClinicsClinicGUIDServicesMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostClinicsClinicGUIDServicesMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PostClinicsClinicGUIDServicesMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostClinicsClinicGUIDServicesNotFoundBody post clinics clinic GUID services not found body
swagger:model PostClinicsClinicGUIDServicesNotFoundBody
*/
type PostClinicsClinicGUIDServicesNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostClinicsClinicGUIDServicesNotFoundBody) UnmarshalJSON(raw []byte) error {
	// PostClinicsClinicGUIDServicesNotFoundBodyAO0
	var postClinicsClinicGUIDServicesNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &postClinicsClinicGUIDServicesNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = postClinicsClinicGUIDServicesNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostClinicsClinicGUIDServicesNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postClinicsClinicGUIDServicesNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postClinicsClinicGUIDServicesNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post clinics clinic GUID services not found body
func (o *PostClinicsClinicGUIDServicesNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *PostClinicsClinicGUIDServicesNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostClinicsClinicGUIDServicesNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PostClinicsClinicGUIDServicesNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostClinicsClinicGUIDServicesOKBody post clinics clinic GUID services o k body
swagger:model PostClinicsClinicGUIDServicesOKBody
*/
type PostClinicsClinicGUIDServicesOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostClinicsClinicGUIDServicesOKBody) UnmarshalJSON(raw []byte) error {
	// PostClinicsClinicGUIDServicesOKBodyAO0
	var postClinicsClinicGUIDServicesOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &postClinicsClinicGUIDServicesOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = postClinicsClinicGUIDServicesOKBodyAO0

	// PostClinicsClinicGUIDServicesOKBodyAO1
	var dataPostClinicsClinicGUIDServicesOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPostClinicsClinicGUIDServicesOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPostClinicsClinicGUIDServicesOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostClinicsClinicGUIDServicesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	postClinicsClinicGUIDServicesOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postClinicsClinicGUIDServicesOKBodyAO0)

	var dataPostClinicsClinicGUIDServicesOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}

	dataPostClinicsClinicGUIDServicesOKBodyAO1.Data = o.Data

	jsonDataPostClinicsClinicGUIDServicesOKBodyAO1, errPostClinicsClinicGUIDServicesOKBodyAO1 := swag.WriteJSON(dataPostClinicsClinicGUIDServicesOKBodyAO1)
	if errPostClinicsClinicGUIDServicesOKBodyAO1 != nil {
		return nil, errPostClinicsClinicGUIDServicesOKBodyAO1
	}
	_parts = append(_parts, jsonDataPostClinicsClinicGUIDServicesOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post clinics clinic GUID services o k body
func (o *PostClinicsClinicGUIDServicesOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostClinicsClinicGUIDServicesOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("postClinicsClinicGuidServicesOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostClinicsClinicGUIDServicesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostClinicsClinicGUIDServicesOKBody) UnmarshalBinary(b []byte) error {
	var res PostClinicsClinicGUIDServicesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostClinicsClinicGUIDServicesParamsBodyAllOf2 post clinics clinic GUID services params body all of2
swagger:model PostClinicsClinicGUIDServicesParamsBodyAllOf2
*/
type PostClinicsClinicGUIDServicesParamsBodyAllOf2 interface{}
