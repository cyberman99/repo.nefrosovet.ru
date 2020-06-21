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

// AppointmentProgramObject Appointment_program_object
// swagger:model Appointment_program_object
type AppointmentProgramObject struct {

	// Идентификатор программы назначений
	// Format: uuid
	ID strfmt.UUID `json:"ID,omitempty"`

	// Дата начала программы назанчения
	// Format: date-time
	Begin strfmt.DateTime `json:"begin,omitempty"`

	// Комментарий врача
	Comment *string `json:"comment"`

	// Идентификатор врача
	// Format: uuid
	DoctorID *strfmt.UUID `json:"doctorID"`

	// Дата окончания программы назанчения
	// Format: date-time
	End *strfmt.DateTime `json:"end"`

	// Периодичность
	Periodicity *string `json:"periodicity"`

	// Статус программы назначений
	// Enum: [OPEN CLOSED]
	StatusCode string `json:"statusCode,omitempty"`

	// Тип программы назначений
	TypeCode string `json:"typeCode,omitempty"`
}

// Validate validates this appointment program object
func (m *AppointmentProgramObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBegin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDoctorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
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

func (m *AppointmentProgramObject) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("ID", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AppointmentProgramObject) validateBegin(formats strfmt.Registry) error {

	if swag.IsZero(m.Begin) { // not required
		return nil
	}

	if err := validate.FormatOf("begin", "body", "date-time", m.Begin.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AppointmentProgramObject) validateDoctorID(formats strfmt.Registry) error {

	if swag.IsZero(m.DoctorID) { // not required
		return nil
	}

	if err := validate.FormatOf("doctorID", "body", "uuid", m.DoctorID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AppointmentProgramObject) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

var appointmentProgramObjectTypeStatusCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN","CLOSED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appointmentProgramObjectTypeStatusCodePropEnum = append(appointmentProgramObjectTypeStatusCodePropEnum, v)
	}
}

const (

	// AppointmentProgramObjectStatusCodeOPEN captures enum value "OPEN"
	AppointmentProgramObjectStatusCodeOPEN string = "OPEN"

	// AppointmentProgramObjectStatusCodeCLOSED captures enum value "CLOSED"
	AppointmentProgramObjectStatusCodeCLOSED string = "CLOSED"
)

// prop value enum
func (m *AppointmentProgramObject) validateStatusCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, appointmentProgramObjectTypeStatusCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AppointmentProgramObject) validateStatusCode(formats strfmt.Registry) error {

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
func (m *AppointmentProgramObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppointmentProgramObject) UnmarshalBinary(b []byte) error {
	var res AppointmentProgramObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
