// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Error405Data Возвращает ошибку
// swagger:model Error_405_data
type Error405Data struct {
	ErrorData

	// сообщение ответа
	Message string `json:"message,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Error405Data) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ErrorData
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ErrorData = aO0

	// AO1
	var dataAO1 struct {
		Message string `json:"message,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Message = dataAO1.Message

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Error405Data) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ErrorData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Message string `json:"message,omitempty"`
	}

	dataAO1.Message = m.Message

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this error 405 data
func (m *Error405Data) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ErrorData
	if err := m.ErrorData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Error405Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Error405Data) UnmarshalBinary(b []byte) error {
	var res Error405Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
