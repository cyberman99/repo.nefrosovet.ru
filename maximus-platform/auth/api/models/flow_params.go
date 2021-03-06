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

// FlowParams Flow_params
// swagger:model Flow_params
type FlowParams struct {

	// Идентификатор бэкенда
	// Required: true
	BackendID *string `json:"backendID"`

	// Порядковый номер
	// Required: true
	Order *int64 `json:"order"`
}

// Validate validates this flow params
func (m *FlowParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackendID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrder(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlowParams) validateBackendID(formats strfmt.Registry) error {

	if err := validate.Required("backendID", "body", m.BackendID); err != nil {
		return err
	}

	return nil
}

func (m *FlowParams) validateOrder(formats strfmt.Registry) error {

	if err := validate.Required("order", "body", m.Order); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowParams) UnmarshalBinary(b []byte) error {
	var res FlowParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
