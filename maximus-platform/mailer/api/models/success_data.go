// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SuccessData Все хорошо
// swagger:model Success_data
type SuccessData struct {
	BaseData

	// errors
	// Required: true
	Errors interface{} `json:"errors"`

	// сообщение ответа
	// Required: true
	Message *string `json:"message"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SuccessData) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseData
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseData = aO0

	// AO1
	var dataAO1 struct {
		Errors interface{} `json:"errors"`

		Message *string `json:"message"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Errors = dataAO1.Errors

	m.Message = dataAO1.Message

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SuccessData) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BaseData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Errors interface{} `json:"errors"`

		Message *string `json:"message"`
	}

	dataAO1.Errors = m.Errors

	dataAO1.Message = m.Message

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this success data
func (m *SuccessData) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseData
	if err := m.BaseData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SuccessData) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("errors", "body", m.Errors); err != nil {
		return err
	}

	return nil
}

func (m *SuccessData) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SuccessData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuccessData) UnmarshalBinary(b []byte) error {
	var res SuccessData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
