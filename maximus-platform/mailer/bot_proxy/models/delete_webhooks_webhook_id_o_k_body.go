// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DeleteWebhooksWebhookIDOKBody delete webhooks webhook Id o k body
// swagger:model deleteWebhooksWebhookIdOKBody
type DeleteWebhooksWebhookIDOKBody struct {
	SuccessData

	DeleteWebhooksWebhookIDOKBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DeleteWebhooksWebhookIDOKBody) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SuccessData
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SuccessData = aO0

	// AO1
	var aO1 DeleteWebhooksWebhookIDOKBodyAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.DeleteWebhooksWebhookIDOKBodyAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DeleteWebhooksWebhookIDOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.SuccessData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.DeleteWebhooksWebhookIDOKBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this delete webhooks webhook Id o k body
func (m *DeleteWebhooksWebhookIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SuccessData
	if err := m.SuccessData.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with DeleteWebhooksWebhookIDOKBodyAllOf1
	if err := m.DeleteWebhooksWebhookIDOKBodyAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteWebhooksWebhookIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteWebhooksWebhookIDOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteWebhooksWebhookIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
