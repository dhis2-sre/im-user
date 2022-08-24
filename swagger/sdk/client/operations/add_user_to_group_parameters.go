// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewAddUserToGroupParams creates a new AddUserToGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddUserToGroupParams() *AddUserToGroupParams {
	return &AddUserToGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddUserToGroupParamsWithTimeout creates a new AddUserToGroupParams object
// with the ability to set a timeout on a request.
func NewAddUserToGroupParamsWithTimeout(timeout time.Duration) *AddUserToGroupParams {
	return &AddUserToGroupParams{
		timeout: timeout,
	}
}

// NewAddUserToGroupParamsWithContext creates a new AddUserToGroupParams object
// with the ability to set a context for a request.
func NewAddUserToGroupParamsWithContext(ctx context.Context) *AddUserToGroupParams {
	return &AddUserToGroupParams{
		Context: ctx,
	}
}

// NewAddUserToGroupParamsWithHTTPClient creates a new AddUserToGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddUserToGroupParamsWithHTTPClient(client *http.Client) *AddUserToGroupParams {
	return &AddUserToGroupParams{
		HTTPClient: client,
	}
}

/* AddUserToGroupParams contains all the parameters to send to the API endpoint
   for the add user to group operation.

   Typically these are written to a http.Request.
*/
type AddUserToGroupParams struct {

	// Group.
	Group string

	// UserID.
	//
	// Format: uint64
	UserID uint64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add user to group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddUserToGroupParams) WithDefaults() *AddUserToGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add user to group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddUserToGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add user to group params
func (o *AddUserToGroupParams) WithTimeout(timeout time.Duration) *AddUserToGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add user to group params
func (o *AddUserToGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add user to group params
func (o *AddUserToGroupParams) WithContext(ctx context.Context) *AddUserToGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add user to group params
func (o *AddUserToGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add user to group params
func (o *AddUserToGroupParams) WithHTTPClient(client *http.Client) *AddUserToGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add user to group params
func (o *AddUserToGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroup adds the group to the add user to group params
func (o *AddUserToGroupParams) WithGroup(group string) *AddUserToGroupParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the add user to group params
func (o *AddUserToGroupParams) SetGroup(group string) {
	o.Group = group
}

// WithUserID adds the userID to the add user to group params
func (o *AddUserToGroupParams) WithUserID(userID uint64) *AddUserToGroupParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the add user to group params
func (o *AddUserToGroupParams) SetUserID(userID uint64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AddUserToGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param group
	if err := r.SetPathParam("group", o.Group); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", swag.FormatUint64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
