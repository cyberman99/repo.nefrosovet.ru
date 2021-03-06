// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Error403Data error 403 data
// swagger:model Error_403_data
type Error403Data struct {
	ErrorData

	Error403DataAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Error403Data) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ErrorData
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ErrorData = aO0

	// AO1
	var aO1 Error403DataAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.Error403DataAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Error403Data) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ErrorData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.Error403DataAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this error 403 data
func (m *Error403Data) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ErrorData
	if err := m.ErrorData.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with Error403DataAllOf1
	if err := m.Error403DataAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Error403Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Error403Data) UnmarshalBinary(b []byte) error {
	var res Error403Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
