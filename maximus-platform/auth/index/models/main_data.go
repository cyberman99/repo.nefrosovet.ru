// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MainData Main_data
// swagger:model Main_data
type MainData struct {

	// GUID
	GUID string `json:"GUID,omitempty"`

	// Состояние записи
	Archived *bool `json:"archived,omitempty"`
}

// Validate validates this main data
func (m *MainData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MainData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MainData) UnmarshalBinary(b []byte) error {
	var res MainData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}