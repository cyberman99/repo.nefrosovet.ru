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

// PatchPatientsPatientGUIDReader is a Reader for the PatchPatientsPatientGUID structure.
type PatchPatientsPatientGUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPatientsPatientGUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchPatientsPatientGUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchPatientsPatientGUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPatchPatientsPatientGUIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchPatientsPatientGUIDOK creates a PatchPatientsPatientGUIDOK with default headers values
func NewPatchPatientsPatientGUIDOK() *PatchPatientsPatientGUIDOK {
	return &PatchPatientsPatientGUIDOK{}
}

/*PatchPatientsPatientGUIDOK handles this case with default header values.

Коллекция пациентов
*/
type PatchPatientsPatientGUIDOK struct {
	Payload *PatchPatientsPatientGUIDOKBody
}

func (o *PatchPatientsPatientGUIDOK) Error() string {
	return fmt.Sprintf("[PATCH /patients/{patientGUID}][%d] patchPatientsPatientGuidOK  %+v", 200, o.Payload)
}

func (o *PatchPatientsPatientGUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PatchPatientsPatientGUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPatientsPatientGUIDBadRequest creates a PatchPatientsPatientGUIDBadRequest with default headers values
func NewPatchPatientsPatientGUIDBadRequest() *PatchPatientsPatientGUIDBadRequest {
	return &PatchPatientsPatientGUIDBadRequest{}
}

/*PatchPatientsPatientGUIDBadRequest handles this case with default header values.

Validation error
*/
type PatchPatientsPatientGUIDBadRequest struct {
	Payload *PatchPatientsPatientGUIDBadRequestBody
}

func (o *PatchPatientsPatientGUIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /patients/{patientGUID}][%d] patchPatientsPatientGuidBadRequest  %+v", 400, o.Payload)
}

func (o *PatchPatientsPatientGUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PatchPatientsPatientGUIDBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPatientsPatientGUIDMethodNotAllowed creates a PatchPatientsPatientGUIDMethodNotAllowed with default headers values
func NewPatchPatientsPatientGUIDMethodNotAllowed() *PatchPatientsPatientGUIDMethodNotAllowed {
	return &PatchPatientsPatientGUIDMethodNotAllowed{}
}

/*PatchPatientsPatientGUIDMethodNotAllowed handles this case with default header values.

Invalid Method
*/
type PatchPatientsPatientGUIDMethodNotAllowed struct {
	Payload *PatchPatientsPatientGUIDMethodNotAllowedBody
}

func (o *PatchPatientsPatientGUIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /patients/{patientGUID}][%d] patchPatientsPatientGuidMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PatchPatientsPatientGUIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PatchPatientsPatientGUIDMethodNotAllowedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PatchPatientsPatientGUIDBadRequestBody patch patients patient GUID bad request body
swagger:model PatchPatientsPatientGUIDBadRequestBody
*/
type PatchPatientsPatientGUIDBadRequestBody struct {
	models.Error400Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchPatientsPatientGUIDBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PatchPatientsPatientGUIDBadRequestBodyAO0
	var patchPatientsPatientGUIDBadRequestBodyAO0 models.Error400Data
	if err := swag.ReadJSON(raw, &patchPatientsPatientGUIDBadRequestBodyAO0); err != nil {
		return err
	}
	o.Error400Data = patchPatientsPatientGUIDBadRequestBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchPatientsPatientGUIDBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	patchPatientsPatientGUIDBadRequestBodyAO0, err := swag.WriteJSON(o.Error400Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchPatientsPatientGUIDBadRequestBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch patients patient GUID bad request body
func (o *PatchPatientsPatientGUIDBadRequestBody) Validate(formats strfmt.Registry) error {
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
func (o *PatchPatientsPatientGUIDBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchPatientsPatientGUIDBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PatchPatientsPatientGUIDBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchPatientsPatientGUIDBody patch patients patient GUID body
swagger:model PatchPatientsPatientGUIDBody
*/
type PatchPatientsPatientGUIDBody struct {
	models.PatientObject

	models.PasswordObject
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchPatientsPatientGUIDBody) UnmarshalJSON(raw []byte) error {
	// PatchPatientsPatientGUIDParamsBodyAO0
	var patchPatientsPatientGUIDParamsBodyAO0 models.PatientObject
	if err := swag.ReadJSON(raw, &patchPatientsPatientGUIDParamsBodyAO0); err != nil {
		return err
	}
	o.PatientObject = patchPatientsPatientGUIDParamsBodyAO0

	// PatchPatientsPatientGUIDParamsBodyAO1
	var patchPatientsPatientGUIDParamsBodyAO1 models.PasswordObject
	if err := swag.ReadJSON(raw, &patchPatientsPatientGUIDParamsBodyAO1); err != nil {
		return err
	}
	o.PasswordObject = patchPatientsPatientGUIDParamsBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchPatientsPatientGUIDBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	patchPatientsPatientGUIDParamsBodyAO0, err := swag.WriteJSON(o.PatientObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchPatientsPatientGUIDParamsBodyAO0)

	patchPatientsPatientGUIDParamsBodyAO1, err := swag.WriteJSON(o.PasswordObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchPatientsPatientGUIDParamsBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch patients patient GUID body
func (o *PatchPatientsPatientGUIDBody) Validate(formats strfmt.Registry) error {
	var res []error

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
func (o *PatchPatientsPatientGUIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchPatientsPatientGUIDBody) UnmarshalBinary(b []byte) error {
	var res PatchPatientsPatientGUIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchPatientsPatientGUIDMethodNotAllowedBody patch patients patient GUID method not allowed body
swagger:model PatchPatientsPatientGUIDMethodNotAllowedBody
*/
type PatchPatientsPatientGUIDMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchPatientsPatientGUIDMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PatchPatientsPatientGUIDMethodNotAllowedBodyAO0
	var patchPatientsPatientGUIDMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &patchPatientsPatientGUIDMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = patchPatientsPatientGUIDMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchPatientsPatientGUIDMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	patchPatientsPatientGUIDMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchPatientsPatientGUIDMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch patients patient GUID method not allowed body
func (o *PatchPatientsPatientGUIDMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PatchPatientsPatientGUIDMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchPatientsPatientGUIDMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PatchPatientsPatientGUIDMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchPatientsPatientGUIDOKBody patch patients patient GUID o k body
swagger:model PatchPatientsPatientGUIDOKBody
*/
type PatchPatientsPatientGUIDOKBody struct {
	models.SuccessData

	// data
	Data []*DataItems0 `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatchPatientsPatientGUIDOKBody) UnmarshalJSON(raw []byte) error {
	// PatchPatientsPatientGUIDOKBodyAO0
	var patchPatientsPatientGUIDOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &patchPatientsPatientGUIDOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = patchPatientsPatientGUIDOKBodyAO0

	// PatchPatientsPatientGUIDOKBodyAO1
	var dataPatchPatientsPatientGUIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPatchPatientsPatientGUIDOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPatchPatientsPatientGUIDOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatchPatientsPatientGUIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	patchPatientsPatientGUIDOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patchPatientsPatientGUIDOKBodyAO0)

	var dataPatchPatientsPatientGUIDOKBodyAO1 struct {
		Data []*DataItems0 `json:"data,omitempty"`
	}

	dataPatchPatientsPatientGUIDOKBodyAO1.Data = o.Data

	jsonDataPatchPatientsPatientGUIDOKBodyAO1, errPatchPatientsPatientGUIDOKBodyAO1 := swag.WriteJSON(dataPatchPatientsPatientGUIDOKBodyAO1)
	if errPatchPatientsPatientGUIDOKBodyAO1 != nil {
		return nil, errPatchPatientsPatientGUIDOKBodyAO1
	}
	_parts = append(_parts, jsonDataPatchPatientsPatientGUIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patch patients patient GUID o k body
func (o *PatchPatientsPatientGUIDOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PatchPatientsPatientGUIDOKBody) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("patchPatientsPatientGuidOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchPatientsPatientGUIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchPatientsPatientGUIDOKBody) UnmarshalBinary(b []byte) error {
	var res PatchPatientsPatientGUIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
