// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AppointmentObject Appointment_object
// swagger:model Appointment_object
type AppointmentObject struct {

	// Идентификатор назначения
	// Format: uuid
	ID strfmt.UUID `json:"ID,omitempty"`

	// Комментарий врача
	Comment *string `json:"comment"`

	// Идентификатор врача
	// Format: uuid
	DoctorID *strfmt.UUID `json:"doctorID"`

	// Продолжительность выполнения назанчения
	Duration *int64 `json:"duration"`

	// Фактическая дата выполнения назначения
	// Format: date-time
	Performed *strfmt.DateTime `json:"performed"`

	// Плановая дата назначения
	// Format: date-time
	Planned strfmt.DateTime `json:"planned,omitempty"`

	// Идентификатор программы лечения
	// Format: uuid
	ProgramID strfmt.UUID `json:"programID,omitempty"`

	// Статус выполнения назначения
	// Enum: [PERFORMED NOT_PERFORMED HALF_PERFORMED]
	StatusCode string `json:"statusCode,omitempty"`

	// Кодификатор типа назначения
	TypeCode string `json:"typeCode,omitempty"`
}

// Validate validates this appointment object
func (m *AppointmentObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDoctorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanned(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProgramID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppointmentObject) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("ID", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AppointmentObject) validateDoctorID(formats strfmt.Registry) error {

	if swag.IsZero(m.DoctorID) { // not required
		return nil
	}

	if err := validate.FormatOf("doctorID", "body", "uuid", m.DoctorID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AppointmentObject) validatePerformed(formats strfmt.Registry) error {

	if swag.IsZero(m.Performed) { // not required
		return nil
	}

	if err := validate.FormatOf("performed", "body", "date-time", m.Performed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AppointmentObject) validatePlanned(formats strfmt.Registry) error {

	if swag.IsZero(m.Planned) { // not required
		return nil
	}

	if err := validate.FormatOf("planned", "body", "date-time", m.Planned.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AppointmentObject) validateProgramID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProgramID) { // not required
		return nil
	}

	if err := validate.FormatOf("programID", "body", "uuid", m.ProgramID.String(), formats); err != nil {
		return err
	}

	return nil
}

var appointmentObjectTypeStatusCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PERFORMED","NOT_PERFORMED","HALF_PERFORMED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appointmentObjectTypeStatusCodePropEnum = append(appointmentObjectTypeStatusCodePropEnum, v)
	}
}

const (

	// AppointmentObjectStatusCodePERFORMED captures enum value "PERFORMED"
	AppointmentObjectStatusCodePERFORMED string = "PERFORMED"

	// AppointmentObjectStatusCodeNOTPERFORMED captures enum value "NOT_PERFORMED"
	AppointmentObjectStatusCodeNOTPERFORMED string = "NOT_PERFORMED"

	// AppointmentObjectStatusCodeHALFPERFORMED captures enum value "HALF_PERFORMED"
	AppointmentObjectStatusCodeHALFPERFORMED string = "HALF_PERFORMED"
)

// prop value enum
func (m *AppointmentObject) validateStatusCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, appointmentObjectTypeStatusCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AppointmentObject) validateStatusCode(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusCodeEnum("statusCode", "body", m.StatusCode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppointmentObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppointmentObject) UnmarshalBinary(b []byte) error {
	var res AppointmentObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
