// Code generated by go-swagger; DO NOT EDIT.

package employees

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new employees API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for employees API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetEmployees коллекцияs сотрудников
*/
func (a *Client) GetEmployees(params *GetEmployeesParams) (*GetEmployeesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmployeesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEmployees",
		Method:             "GET",
		PathPattern:        "/employees",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEmployeesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEmployeesOK), nil

}

/*
GetEmployeesEmployeeGUID информацияs о сотруднике
*/
func (a *Client) GetEmployeesEmployeeGUID(params *GetEmployeesEmployeeGUIDParams) (*GetEmployeesEmployeeGUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEmployeesEmployeeGUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetEmployeesEmployeeGUID",
		Method:             "GET",
		PathPattern:        "/employees/{employeeGUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEmployeesEmployeeGUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEmployeesEmployeeGUIDOK), nil

}

/*
PatchEmployeesEmployeeGUID изменениеs сотрудника
*/
func (a *Client) PatchEmployeesEmployeeGUID(params *PatchEmployeesEmployeeGUIDParams) (*PatchEmployeesEmployeeGUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchEmployeesEmployeeGUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchEmployeesEmployeeGUID",
		Method:             "PATCH",
		PathPattern:        "/employees/{employeeGUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchEmployeesEmployeeGUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchEmployeesEmployeeGUIDOK), nil

}

/*
PostEmployees созданиеs сотрудника
*/
func (a *Client) PostEmployees(params *PostEmployeesParams) (*PostEmployeesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostEmployeesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostEmployees",
		Method:             "POST",
		PathPattern:        "/employees",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostEmployeesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostEmployeesOK), nil

}

/*
PutEmployeesEmployeeGUID изменениеs сотрудника
*/
func (a *Client) PutEmployeesEmployeeGUID(params *PutEmployeesEmployeeGUIDParams) (*PutEmployeesEmployeeGUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutEmployeesEmployeeGUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutEmployeesEmployeeGUID",
		Method:             "PUT",
		PathPattern:        "/employees/{employeeGUID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutEmployeesEmployeeGUIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutEmployeesEmployeeGUIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}