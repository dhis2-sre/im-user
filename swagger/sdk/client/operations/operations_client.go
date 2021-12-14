// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetGroupByID(params *GetGroupByIDParams, opts ...ClientOption) (*GetGroupByIDOK, error)

	GetUserByID(params *GetUserByIDParams, opts ...ClientOption) (*GetUserByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetGroupByID Return a group by id
*/
func (a *Client) GetGroupByID(params *GetGroupByIDParams, opts ...ClientOption) (*GetGroupByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGroupById",
		Method:             "GET",
		PathPattern:        "/groups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGroupByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGroupByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGroupById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUserByID Return a user by id
*/
func (a *Client) GetUserByID(params *GetUserByIDParams, opts ...ClientOption) (*GetUserByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserById",
		Method:             "GET",
		PathPattern:        "/findbyid/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}