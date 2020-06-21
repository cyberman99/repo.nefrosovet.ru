// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClientObject Client_object
// swagger:model Client_object
type ClientObject struct {

	// Пароль клиента
	Password string `json:"password,omitempty"`

	// Логин клиента
	Username string `json:"username,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ClientObject) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Password string `json:"password,omitempty"`

		Username string `json:"username,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Password = dataAO0.Password

	m.Username = dataAO0.Username

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ClientObject) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Password string `json:"password,omitempty"`

		Username string `json:"username,omitempty"`
	}

	dataAO0.Password = m.Password

	dataAO0.Username = m.Username

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this client object
func (m *ClientObject) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClientObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientObject) UnmarshalBinary(b []byte) error {
	var res ClientObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
