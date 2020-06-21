// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetClinicsClinicGUIDServices коллекцияs сервисов
*/
func (a *Client) GetClinicsClinicGUIDServices(params *GetClinicsClinicGUIDServicesParams) (*GetClinicsClinicGUIDServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClinicsClinicGUIDServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClinicsClinicGUIDServices",
		Method:             "GET",
		PathPattern:        "/clinics/{clinicGUID}/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClinicsClinicGUIDServicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClinicsClinicGUIDServicesOK), nil

}

/*
GetClinicsClinicGUIDServicesServiceGUID информацияs о приложении
*/
func (a *Client) GetClinicsClinicGUIDServicesServiceGUID(params *GetClinicsClinicGUIDServicesServiceGUIDParams) (*GetClinicsClinicGUIDServicesServiceGUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClinicsClinicGUIDServicesServiceGUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClinicsClinicGUIDServicesServiceGUID",
		Method:             "GET",
		PathPattern:        "/clinics/{clinicGUID}/services/{serviceGUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClinicsClinicGUIDServicesServiceGUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClinicsClinicGUIDServicesServiceGUIDOK), nil

}

/*
PostClinicsClinicGUIDServices созданиеs сервиса
*/
func (a *Client) PostClinicsClinicGUIDServices(params *PostClinicsClinicGUIDServicesParams) (*PostClinicsClinicGUIDServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClinicsClinicGUIDServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostClinicsClinicGUIDServices",
		Method:             "POST",
		PathPattern:        "/clinics/{clinicGUID}/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostClinicsClinicGUIDServicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostClinicsClinicGUIDServicesOK), nil

}

/*
PutClinicsClinicGUIDServicesServiceGUID изменениеs сервиса
*/
func (a *Client) PutClinicsClinicGUIDServicesServiceGUID(params *PutClinicsClinicGUIDServicesServiceGUIDParams) (*PutClinicsClinicGUIDServicesServiceGUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClinicsClinicGUIDServicesServiceGUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClinicsClinicGUIDServicesServiceGUID",
		Method:             "PUT",
		PathPattern:        "/clinics/{clinicGUID}/services/{serviceGUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClinicsClinicGUIDServicesServiceGUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutClinicsClinicGUIDServicesServiceGUIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
