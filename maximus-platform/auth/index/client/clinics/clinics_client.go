// Code generated by go-swagger; DO NOT EDIT.

package clinics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new clinics API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for clinics API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetClinics коллекцияs клиник
*/
func (a *Client) GetClinics(params *GetClinicsParams) (*GetClinicsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClinicsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClinics",
		Method:             "GET",
		PathPattern:        "/clinics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClinicsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClinicsOK), nil

}

/*
GetClinicsClinicGUID информацияs о клинике
*/
func (a *Client) GetClinicsClinicGUID(params *GetClinicsClinicGUIDParams) (*GetClinicsClinicGUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClinicsClinicGUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClinicsClinicGUID",
		Method:             "GET",
		PathPattern:        "/clinics/{clinicGUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClinicsClinicGUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClinicsClinicGUIDOK), nil

}

/*
PostClinics созданиеs клиники
*/
func (a *Client) PostClinics(params *PostClinicsParams) (*PostClinicsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClinicsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostClinics",
		Method:             "POST",
		PathPattern:        "/clinics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostClinicsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostClinicsOK), nil

}

/*
PutClinicsClinicGUID изменениеs клиники
*/
func (a *Client) PutClinicsClinicGUID(params *PutClinicsClinicGUIDParams) (*PutClinicsClinicGUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClinicsClinicGUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutClinicsClinicGUID",
		Method:             "PUT",
		PathPattern:        "/clinics/{clinicGUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutClinicsClinicGUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutClinicsClinicGUIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}