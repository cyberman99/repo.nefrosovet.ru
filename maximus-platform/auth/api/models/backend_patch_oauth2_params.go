// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackendPatchOauth2Params Backend_patch_oauth2_params
// swagger:model Backend_patch_oauth2_params
type BackendPatchOauth2Params struct {

	// Идентификатор приложения
	ClientID *string `json:"clientID,omitempty"`

	// Секретный ключ приложения
	ClientSecret *string `json:"clientSecret,omitempty"`

	// Описание бэкенда
	Description *string `json:"description,omitempty"`

	// Название сервиса, на котором авторизуемся
	// Enum: [GITHUB GOOGLE EMPLOYEE YANDEX ESIA]
	Provider *string `json:"provider,omitempty"`

	// Сущность синхронизации
	// Enum: [PATIENT EMPLOYEE]
	Sync *string `json:"sync,omitempty"`
}

// Validate validates this backend patch oauth2 params
func (m *BackendPatchOauth2Params) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSync(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var backendPatchOauth2ParamsTypeProviderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["GITHUB","GOOGLE","EMPLOYEE","YANDEX","ESIA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendPatchOauth2ParamsTypeProviderPropEnum = append(backendPatchOauth2ParamsTypeProviderPropEnum, v)
	}
}

const (

	// BackendPatchOauth2ParamsProviderGITHUB captures enum value "GITHUB"
	BackendPatchOauth2ParamsProviderGITHUB string = "GITHUB"

	// BackendPatchOauth2ParamsProviderGOOGLE captures enum value "GOOGLE"
	BackendPatchOauth2ParamsProviderGOOGLE string = "GOOGLE"

	// BackendPatchOauth2ParamsProviderEMPLOYEE captures enum value "EMPLOYEE"
	BackendPatchOauth2ParamsProviderEMPLOYEE string = "EMPLOYEE"

	// BackendPatchOauth2ParamsProviderYANDEX captures enum value "YANDEX"
	BackendPatchOauth2ParamsProviderYANDEX string = "YANDEX"

	// BackendPatchOauth2ParamsProviderESIA captures enum value "ESIA"
	BackendPatchOauth2ParamsProviderESIA string = "ESIA"
)

// prop value enum
func (m *BackendPatchOauth2Params) validateProviderEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendPatchOauth2ParamsTypeProviderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BackendPatchOauth2Params) validateProvider(formats strfmt.Registry) error {

	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	// value enum
	if err := m.validateProviderEnum("provider", "body", *m.Provider); err != nil {
		return err
	}

	return nil
}

var backendPatchOauth2ParamsTypeSyncPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PATIENT","EMPLOYEE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backendPatchOauth2ParamsTypeSyncPropEnum = append(backendPatchOauth2ParamsTypeSyncPropEnum, v)
	}
}

const (

	// BackendPatchOauth2ParamsSyncPATIENT captures enum value "PATIENT"
	BackendPatchOauth2ParamsSyncPATIENT string = "PATIENT"

	// BackendPatchOauth2ParamsSyncEMPLOYEE captures enum value "EMPLOYEE"
	BackendPatchOauth2ParamsSyncEMPLOYEE string = "EMPLOYEE"
)

// prop value enum
func (m *BackendPatchOauth2Params) validateSyncEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, backendPatchOauth2ParamsTypeSyncPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BackendPatchOauth2Params) validateSync(formats strfmt.Registry) error {

	if swag.IsZero(m.Sync) { // not required
		return nil
	}

	// value enum
	if err := m.validateSyncEnum("sync", "body", *m.Sync); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackendPatchOauth2Params) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackendPatchOauth2Params) UnmarshalBinary(b []byte) error {
	var res BackendPatchOauth2Params
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
