// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "repo.nefrosovet.ru/maximus-platform/patient/api/models"
)

// PatientProfileUpdateHandlerFunc turns a function with the right signature into a patient profile update handler
type PatientProfileUpdateHandlerFunc func(PatientProfileUpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatientProfileUpdateHandlerFunc) Handle(params PatientProfileUpdateParams) middleware.Responder {
	return fn(params)
}

// PatientProfileUpdateHandler interface for that can handle valid patient profile update params
type PatientProfileUpdateHandler interface {
	Handle(PatientProfileUpdateParams) middleware.Responder
}

// NewPatientProfileUpdate creates a new http.Handler for the patient profile update operation
func NewPatientProfileUpdate(ctx *middleware.Context, handler PatientProfileUpdateHandler) *PatientProfileUpdate {
	return &PatientProfileUpdate{Context: ctx, Handler: handler}
}

/*PatientProfileUpdate swagger:route PATCH /users/{userID} Profile patientProfileUpdate

Редактирование профиля пациента

*/
type PatientProfileUpdate struct {
	Context *middleware.Context
	Handler PatientProfileUpdateHandler
}

func (o *PatientProfileUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatientProfileUpdateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PatientProfileUpdateBadRequestBody patient profile update bad request body
// swagger:model PatientProfileUpdateBadRequestBody
type PatientProfileUpdateBadRequestBody struct {
	PatientProfileUpdateBadRequestBodyAllOf0

	PatientProfileUpdateBadRequestBodyAllOf1

	// errors
	Errors *PatientProfileUpdateBadRequestBodyAO2Errors `json:"errors,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatientProfileUpdateBadRequestBody) UnmarshalJSON(raw []byte) error {
	// PatientProfileUpdateBadRequestBodyAO0
	var patientProfileUpdateBadRequestBodyAO0 PatientProfileUpdateBadRequestBodyAllOf0
	if err := swag.ReadJSON(raw, &patientProfileUpdateBadRequestBodyAO0); err != nil {
		return err
	}
	o.PatientProfileUpdateBadRequestBodyAllOf0 = patientProfileUpdateBadRequestBodyAO0

	// PatientProfileUpdateBadRequestBodyAO1
	var patientProfileUpdateBadRequestBodyAO1 PatientProfileUpdateBadRequestBodyAllOf1
	if err := swag.ReadJSON(raw, &patientProfileUpdateBadRequestBodyAO1); err != nil {
		return err
	}
	o.PatientProfileUpdateBadRequestBodyAllOf1 = patientProfileUpdateBadRequestBodyAO1

	// PatientProfileUpdateBadRequestBodyAO2
	var dataPatientProfileUpdateBadRequestBodyAO2 struct {
		Errors *PatientProfileUpdateBadRequestBodyAO2Errors `json:"errors,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataPatientProfileUpdateBadRequestBodyAO2); err != nil {
		return err
	}

	o.Errors = dataPatientProfileUpdateBadRequestBodyAO2.Errors

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatientProfileUpdateBadRequestBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	patientProfileUpdateBadRequestBodyAO0, err := swag.WriteJSON(o.PatientProfileUpdateBadRequestBodyAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patientProfileUpdateBadRequestBodyAO0)

	patientProfileUpdateBadRequestBodyAO1, err := swag.WriteJSON(o.PatientProfileUpdateBadRequestBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patientProfileUpdateBadRequestBodyAO1)

	var dataPatientProfileUpdateBadRequestBodyAO2 struct {
		Errors *PatientProfileUpdateBadRequestBodyAO2Errors `json:"errors,omitempty"`
	}

	dataPatientProfileUpdateBadRequestBodyAO2.Errors = o.Errors

	jsonDataPatientProfileUpdateBadRequestBodyAO2, errPatientProfileUpdateBadRequestBodyAO2 := swag.WriteJSON(dataPatientProfileUpdateBadRequestBodyAO2)
	if errPatientProfileUpdateBadRequestBodyAO2 != nil {
		return nil, errPatientProfileUpdateBadRequestBodyAO2
	}
	_parts = append(_parts, jsonDataPatientProfileUpdateBadRequestBodyAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patient profile update bad request body
func (o *PatientProfileUpdateBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PatientProfileUpdateBadRequestBodyAllOf0
	// validation for a type composition with PatientProfileUpdateBadRequestBodyAllOf1

	if err := o.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatientProfileUpdateBadRequestBody) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.Errors) { // not required
		return nil
	}

	if o.Errors != nil {
		if err := o.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patientProfileUpdateBadRequest" + "." + "errors")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatientProfileUpdateBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatientProfileUpdateBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PatientProfileUpdateBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatientProfileUpdateBadRequestBodyAO2Errors patient profile update bad request body a o2 errors
// swagger:model PatientProfileUpdateBadRequestBodyAO2Errors
type PatientProfileUpdateBadRequestBodyAO2Errors struct {

	// validation
	Validation *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation `json:"validation,omitempty"`
}

// Validate validates this patient profile update bad request body a o2 errors
func (o *PatientProfileUpdateBadRequestBodyAO2Errors) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValidation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatientProfileUpdateBadRequestBodyAO2Errors) validateValidation(formats strfmt.Registry) error {

	if swag.IsZero(o.Validation) { // not required
		return nil
	}

	if o.Validation != nil {
		if err := o.Validation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patientProfileUpdateBadRequest" + "." + "errors" + "." + "validation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatientProfileUpdateBadRequestBodyAO2Errors) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatientProfileUpdateBadRequestBodyAO2Errors) UnmarshalBinary(b []byte) error {
	var res PatientProfileUpdateBadRequestBodyAO2Errors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatientProfileUpdateBadRequestBodyAO2ErrorsValidation patient profile update bad request body a o2 errors validation
// swagger:model PatientProfileUpdateBadRequestBodyAO2ErrorsValidation
type PatientProfileUpdateBadRequestBodyAO2ErrorsValidation struct {

	// first name
	// Enum: [string]
	FirstName string `json:"firstName,omitempty"`

	// last name
	// Enum: [string]
	LastName string `json:"lastName,omitempty"`

	// locale
	// Enum: [string format]
	Locale string `json:"locale,omitempty"`

	// patronymic
	// Enum: [string]
	Patronymic string `json:"patronymic,omitempty"`

	// theme
	// Enum: [string oneof]
	Theme string `json:"theme,omitempty"`
}

// Validate validates this patient profile update bad request body a o2 errors validation
func (o *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocale(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePatronymic(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTheme(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeFirstNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeFirstNamePropEnum = append(patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeFirstNamePropEnum, v)
	}
}

const (

	// PatientProfileUpdateBadRequestBodyAO2ErrorsValidationFirstNameString captures enum value "string"
	PatientProfileUpdateBadRequestBodyAO2ErrorsValidationFirstNameString string = "string"
)

// prop value enum
func (o *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation) validateFirstNameEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeFirstNamePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation) validateFirstName(formats strfmt.Registry) error {

	if swag.IsZero(o.FirstName) { // not required
		return nil
	}

	// value enum
	if err := o.validateFirstNameEnum("patientProfileUpdateBadRequest"+"."+"errors"+"."+"validation"+"."+"firstName", "body", o.FirstName); err != nil {
		return err
	}

	return nil
}

var patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeLastNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeLastNamePropEnum = append(patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeLastNamePropEnum, v)
	}
}

const (

	// PatientProfileUpdateBadRequestBodyAO2ErrorsValidationLastNameString captures enum value "string"
	PatientProfileUpdateBadRequestBodyAO2ErrorsValidationLastNameString string = "string"
)

// prop value enum
func (o *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation) validateLastNameEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeLastNamePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation) validateLastName(formats strfmt.Registry) error {

	if swag.IsZero(o.LastName) { // not required
		return nil
	}

	// value enum
	if err := o.validateLastNameEnum("patientProfileUpdateBadRequest"+"."+"errors"+"."+"validation"+"."+"lastName", "body", o.LastName); err != nil {
		return err
	}

	return nil
}

var patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeLocalePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string","format"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeLocalePropEnum = append(patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeLocalePropEnum, v)
	}
}

const (

	// PatientProfileUpdateBadRequestBodyAO2ErrorsValidationLocaleString captures enum value "string"
	PatientProfileUpdateBadRequestBodyAO2ErrorsValidationLocaleString string = "string"

	// PatientProfileUpdateBadRequestBodyAO2ErrorsValidationLocaleFormat captures enum value "format"
	PatientProfileUpdateBadRequestBodyAO2ErrorsValidationLocaleFormat string = "format"
)

// prop value enum
func (o *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation) validateLocaleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeLocalePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation) validateLocale(formats strfmt.Registry) error {

	if swag.IsZero(o.Locale) { // not required
		return nil
	}

	// value enum
	if err := o.validateLocaleEnum("patientProfileUpdateBadRequest"+"."+"errors"+"."+"validation"+"."+"locale", "body", o.Locale); err != nil {
		return err
	}

	return nil
}

var patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypePatronymicPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypePatronymicPropEnum = append(patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypePatronymicPropEnum, v)
	}
}

const (

	// PatientProfileUpdateBadRequestBodyAO2ErrorsValidationPatronymicString captures enum value "string"
	PatientProfileUpdateBadRequestBodyAO2ErrorsValidationPatronymicString string = "string"
)

// prop value enum
func (o *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation) validatePatronymicEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypePatronymicPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation) validatePatronymic(formats strfmt.Registry) error {

	if swag.IsZero(o.Patronymic) { // not required
		return nil
	}

	// value enum
	if err := o.validatePatronymicEnum("patientProfileUpdateBadRequest"+"."+"errors"+"."+"validation"+"."+"patronymic", "body", o.Patronymic); err != nil {
		return err
	}

	return nil
}

var patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeThemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string","oneof"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeThemePropEnum = append(patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeThemePropEnum, v)
	}
}

const (

	// PatientProfileUpdateBadRequestBodyAO2ErrorsValidationThemeString captures enum value "string"
	PatientProfileUpdateBadRequestBodyAO2ErrorsValidationThemeString string = "string"

	// PatientProfileUpdateBadRequestBodyAO2ErrorsValidationThemeOneof captures enum value "oneof"
	PatientProfileUpdateBadRequestBodyAO2ErrorsValidationThemeOneof string = "oneof"
)

// prop value enum
func (o *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation) validateThemeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, patientProfileUpdateBadRequestBodyAO2ErrorsValidationTypeThemePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation) validateTheme(formats strfmt.Registry) error {

	if swag.IsZero(o.Theme) { // not required
		return nil
	}

	// value enum
	if err := o.validateThemeEnum("patientProfileUpdateBadRequest"+"."+"errors"+"."+"validation"+"."+"theme", "body", o.Theme); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatientProfileUpdateBadRequestBodyAO2ErrorsValidation) UnmarshalBinary(b []byte) error {
	var res PatientProfileUpdateBadRequestBodyAO2ErrorsValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatientProfileUpdateBadRequestBodyAllOf0 Ошибка валидации
// swagger:model PatientProfileUpdateBadRequestBodyAllOf0
type PatientProfileUpdateBadRequestBodyAllOf0 interface{}

// PatientProfileUpdateBadRequestBodyAllOf1 patient profile update bad request body all of1
// swagger:model PatientProfileUpdateBadRequestBodyAllOf1
type PatientProfileUpdateBadRequestBodyAllOf1 interface{}

// PatientProfileUpdateBody patient profile update body
// swagger:model PatientProfileUpdateBody
type PatientProfileUpdateBody struct {
	models.ProfileObject
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatientProfileUpdateBody) UnmarshalJSON(raw []byte) error {
	// PatientProfileUpdateParamsBodyAO0
	var patientProfileUpdateParamsBodyAO0 models.ProfileObject
	if err := swag.ReadJSON(raw, &patientProfileUpdateParamsBodyAO0); err != nil {
		return err
	}
	o.ProfileObject = patientProfileUpdateParamsBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatientProfileUpdateBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	patientProfileUpdateParamsBodyAO0, err := swag.WriteJSON(o.ProfileObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patientProfileUpdateParamsBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patient profile update body
func (o *PatientProfileUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ProfileObject
	if err := o.ProfileObject.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PatientProfileUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatientProfileUpdateBody) UnmarshalBinary(b []byte) error {
	var res PatientProfileUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatientProfileUpdateInternalServerErrorBody patient profile update internal server error body
// swagger:model PatientProfileUpdateInternalServerErrorBody
type PatientProfileUpdateInternalServerErrorBody struct {
	models.Error500Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatientProfileUpdateInternalServerErrorBody) UnmarshalJSON(raw []byte) error {
	// PatientProfileUpdateInternalServerErrorBodyAO0
	var patientProfileUpdateInternalServerErrorBodyAO0 models.Error500Data
	if err := swag.ReadJSON(raw, &patientProfileUpdateInternalServerErrorBodyAO0); err != nil {
		return err
	}
	o.Error500Data = patientProfileUpdateInternalServerErrorBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatientProfileUpdateInternalServerErrorBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	patientProfileUpdateInternalServerErrorBodyAO0, err := swag.WriteJSON(o.Error500Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patientProfileUpdateInternalServerErrorBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patient profile update internal server error body
func (o *PatientProfileUpdateInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Error500Data
	if err := o.Error500Data.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PatientProfileUpdateInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatientProfileUpdateInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PatientProfileUpdateInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatientProfileUpdateMethodNotAllowedBody patient profile update method not allowed body
// swagger:model PatientProfileUpdateMethodNotAllowedBody
type PatientProfileUpdateMethodNotAllowedBody struct {
	models.Error405Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatientProfileUpdateMethodNotAllowedBody) UnmarshalJSON(raw []byte) error {
	// PatientProfileUpdateMethodNotAllowedBodyAO0
	var patientProfileUpdateMethodNotAllowedBodyAO0 models.Error405Data
	if err := swag.ReadJSON(raw, &patientProfileUpdateMethodNotAllowedBodyAO0); err != nil {
		return err
	}
	o.Error405Data = patientProfileUpdateMethodNotAllowedBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatientProfileUpdateMethodNotAllowedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	patientProfileUpdateMethodNotAllowedBodyAO0, err := swag.WriteJSON(o.Error405Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patientProfileUpdateMethodNotAllowedBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patient profile update method not allowed body
func (o *PatientProfileUpdateMethodNotAllowedBody) Validate(formats strfmt.Registry) error {
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
func (o *PatientProfileUpdateMethodNotAllowedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatientProfileUpdateMethodNotAllowedBody) UnmarshalBinary(b []byte) error {
	var res PatientProfileUpdateMethodNotAllowedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatientProfileUpdateNotFoundBody patient profile update not found body
// swagger:model PatientProfileUpdateNotFoundBody
type PatientProfileUpdateNotFoundBody struct {
	models.Error404Data
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatientProfileUpdateNotFoundBody) UnmarshalJSON(raw []byte) error {
	// PatientProfileUpdateNotFoundBodyAO0
	var patientProfileUpdateNotFoundBodyAO0 models.Error404Data
	if err := swag.ReadJSON(raw, &patientProfileUpdateNotFoundBodyAO0); err != nil {
		return err
	}
	o.Error404Data = patientProfileUpdateNotFoundBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatientProfileUpdateNotFoundBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	patientProfileUpdateNotFoundBodyAO0, err := swag.WriteJSON(o.Error404Data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patientProfileUpdateNotFoundBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patient profile update not found body
func (o *PatientProfileUpdateNotFoundBody) Validate(formats strfmt.Registry) error {
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
func (o *PatientProfileUpdateNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatientProfileUpdateNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PatientProfileUpdateNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PatientProfileUpdateOKBody patient profile update o k body
// swagger:model PatientProfileUpdateOKBody
type PatientProfileUpdateOKBody struct {
	models.SuccessData

	// data
	// Required: true
	Data []*ProfileDataItem `json:"data"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PatientProfileUpdateOKBody) UnmarshalJSON(raw []byte) error {
	// PatientProfileUpdateOKBodyAO0
	var patientProfileUpdateOKBodyAO0 models.SuccessData
	if err := swag.ReadJSON(raw, &patientProfileUpdateOKBodyAO0); err != nil {
		return err
	}
	o.SuccessData = patientProfileUpdateOKBodyAO0

	// PatientProfileUpdateOKBodyAO1
	var dataPatientProfileUpdateOKBodyAO1 struct {
		Data []*ProfileDataItem `json:"data"`
	}
	if err := swag.ReadJSON(raw, &dataPatientProfileUpdateOKBodyAO1); err != nil {
		return err
	}

	o.Data = dataPatientProfileUpdateOKBodyAO1.Data

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PatientProfileUpdateOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	patientProfileUpdateOKBodyAO0, err := swag.WriteJSON(o.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, patientProfileUpdateOKBodyAO0)

	var dataPatientProfileUpdateOKBodyAO1 struct {
		Data []*ProfileDataItem `json:"data"`
	}

	dataPatientProfileUpdateOKBodyAO1.Data = o.Data

	jsonDataPatientProfileUpdateOKBodyAO1, errPatientProfileUpdateOKBodyAO1 := swag.WriteJSON(dataPatientProfileUpdateOKBodyAO1)
	if errPatientProfileUpdateOKBodyAO1 != nil {
		return nil, errPatientProfileUpdateOKBodyAO1
	}
	_parts = append(_parts, jsonDataPatientProfileUpdateOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this patient profile update o k body
func (o *PatientProfileUpdateOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PatientProfileUpdateOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("patientProfileUpdateOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("patientProfileUpdateOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatientProfileUpdateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatientProfileUpdateOKBody) UnmarshalBinary(b []byte) error {
	var res PatientProfileUpdateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// profile-data-item profile data item
// swagger:model profile-data-item
type ProfileDataItem struct {
	models.ProfileObject
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ProfileDataItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 models.ProfileObject
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	o.ProfileObject = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ProfileDataItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(o.ProfileObject)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this profile data item
func (o *ProfileDataItem) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ProfileObject
	if err := o.ProfileObject.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *ProfileDataItem) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProfileDataItem) UnmarshalBinary(b []byte) error {
	var res ProfileDataItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}