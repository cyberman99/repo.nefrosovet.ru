// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BackendIDObject BackendID_object
// swagger:model BackendID_object
type BackendIDObject struct {

	// Идентификатор бэкенда
	ID string `json:"ID,omitempty"`
}

// Validate validates this backend ID object
func (m *BackendIDObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackendIDObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackendIDObject) UnmarshalBinary(b []byte) error {
	var res BackendIDObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}