// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BackendOauth2IDParam Backend_oauth2_ID_param
// swagger:model Backend_oauth2_ID_param
type BackendOauth2IDParam struct {

	// Кастомный ID бэкенда (Для OAuth)
	ID string `json:"ID,omitempty"`
}

// Validate validates this backend oauth2 ID param
func (m *BackendOauth2IDParam) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackendOauth2IDParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackendOauth2IDParam) UnmarshalBinary(b []byte) error {
	var res BackendOauth2IDParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
