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
	Jwks(params *JwksParams, opts ...ClientOption) (*JwksOK, error)

	AddClusterConfigurationToGroup(params *AddClusterConfigurationToGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddClusterConfigurationToGroupCreated, error)

	AddUserToGroup(params *AddUserToGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddUserToGroupCreated, error)

	FindGroupByName(params *FindGroupByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindGroupByNameOK, error)

	FindUserByID(params *FindUserByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUserByIDOK, error)

	GroupCreate(params *GroupCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GroupCreateCreated, error)

	Health(params *HealthParams, opts ...ClientOption) (*HealthOK, error)

	Me(params *MeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MeOK, error)

	RefreshToken(params *RefreshTokenParams, opts ...ClientOption) (*RefreshTokenCreated, error)

	SignIn(params *SignInParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SignInCreated, error)

	SignOut(params *SignOutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SignOutOK, error)

	SignUp(params *SignUpParams, opts ...ClientOption) (*SignUpCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Jwks js w k s

Return a JWKS containing the public key which can be used to validate the JWT's dispensed at /signin
*/
func (a *Client) Jwks(params *JwksParams, opts ...ClientOption) (*JwksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJwksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Jwks",
		Method:             "GET",
		PathPattern:        "/jwks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &JwksReader{formats: a.formats},
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
	success, ok := result.(*JwksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Jwks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	AddClusterConfigurationToGroup adds cluster configuration to group

	Add a cluster configuration to a group. This will allow deploying to a remote cluster.

Currently only configurations with embedded access tokens are support.
The configuration needs to be encrypted using Mozilla Sops. Please see ./scripts/addClusterConfigToGroup.sh for an example of how this can be done.
*/
func (a *Client) AddClusterConfigurationToGroup(params *AddClusterConfigurationToGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddClusterConfigurationToGroupCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddClusterConfigurationToGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addClusterConfigurationToGroup",
		Method:             "POST",
		PathPattern:        "/groups/{group}/cluster-configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddClusterConfigurationToGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*AddClusterConfigurationToGroupCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addClusterConfigurationToGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AddUserToGroup adds user to group

Add a user to a group...
*/
func (a *Client) AddUserToGroup(params *AddUserToGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddUserToGroupCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddUserToGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addUserToGroup",
		Method:             "POST",
		PathPattern:        "/groups/{group}/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddUserToGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*AddUserToGroupCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addUserToGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindGroupByName finds group

Find a group by its name
*/
func (a *Client) FindGroupByName(params *FindGroupByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindGroupByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindGroupByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findGroupByName",
		Method:             "GET",
		PathPattern:        "/groups/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindGroupByNameReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*FindGroupByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findGroupByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FindUserByID finds user

Find a user by its id
*/
func (a *Client) FindUserByID(params *FindUserByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindUserByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindUserByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findUserById",
		Method:             "GET",
		PathPattern:        "/users/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindUserByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*FindUserByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findUserById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GroupCreate creates group

Create a group...
*/
func (a *Client) GroupCreate(params *GroupCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GroupCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGroupCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "groupCreate",
		Method:             "POST",
		PathPattern:        "/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GroupCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*GroupCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for groupCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Health healths status

Show service health status
*/
func (a *Client) Health(params *HealthParams, opts ...ClientOption) (*HealthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHealthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "health",
		Method:             "GET",
		PathPattern:        "/health",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HealthReader{formats: a.formats},
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
	success, ok := result.(*HealthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for health: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Me users details

Current user details
*/
func (a *Client) Me(params *MeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "me",
		Method:             "GET",
		PathPattern:        "/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MeReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*MeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for me: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RefreshToken refreshes tokens

Refresh user tokens
*/
func (a *Client) RefreshToken(params *RefreshTokenParams, opts ...ClientOption) (*RefreshTokenCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRefreshTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "refreshToken",
		Method:             "POST",
		PathPattern:        "/refresh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RefreshTokenReader{formats: a.formats},
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
	success, ok := result.(*RefreshTokenCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for refreshToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SignIn signs in

Sign in... And get tokens
*/
func (a *Client) SignIn(params *SignInParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SignInCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSignInParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "signIn",
		Method:             "POST",
		PathPattern:        "/tokens",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SignInReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*SignInCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for signIn: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SignOut signs out

Sign out user... The authentication is done using oauth and JWT. A JWT can't easily be invalidated so even after calling this endpoint a user can still sign in assuming the JWT isn't expired. However, the token can't be refreshed using the refresh token supplied upon signin
*/
func (a *Client) SignOut(params *SignOutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SignOutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSignOutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "signOut",
		Method:             "DELETE",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SignOutReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*SignOutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for signOut: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SignUp signs up user

Sign up a user. This endpoint is publicly accessible and therefor anyone can sign up. However, before being able to perform any actions, users needs to be a member of a group. And only administrators can add users to groups.
*/
func (a *Client) SignUp(params *SignUpParams, opts ...ClientOption) (*SignUpCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSignUpParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "signUp",
		Method:             "POST",
		PathPattern:        "/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SignUpReader{formats: a.formats},
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
	success, ok := result.(*SignUpCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for signUp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
