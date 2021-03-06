// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EmployeeObject Employee_object
// swagger:model Employee_object
type EmployeeObject struct {

	// E-mail адрес
	Email string `json:"email,omitempty"`

	// Имя
	FirstName string `json:"firstName,omitempty"`

	// Фамилия
	LastName string `json:"lastName,omitempty"`

	// Номер мобильного телефона
	Mobile string `json:"mobile,omitempty"`

	// Отчество
	Patronymic string `json:"patronymic,omitempty"`

	// Номер смарт-карты
	SmartCardNumber string `json:"smartCardNumber,omitempty"`

	// СНИЛС
	Snils string `json:"snils,omitempty"`

	// Имя пользователя
	Username string `json:"username,omitempty"`
}

// Validate validates this employee object
func (m *EmployeeObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmployeeObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmployeeObject) UnmarshalBinary(b []byte) error {
	var res EmployeeObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
