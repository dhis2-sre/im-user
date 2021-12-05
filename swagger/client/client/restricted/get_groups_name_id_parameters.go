// Code generated by go-swagger; DO NOT EDIT.

package restricted

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
)

// NewGetGroupsNameIDParams creates a new GetGroupsNameIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGroupsNameIDParams() *GetGroupsNameIDParams {
	return &GetGroupsNameIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupsNameIDParamsWithTimeout creates a new GetGroupsNameIDParams object
// with the ability to set a timeout on a request.
func NewGetGroupsNameIDParamsWithTimeout(timeout time.Duration) *GetGroupsNameIDParams {
	return &GetGroupsNameIDParams{
		timeout: timeout,
	}
}

// NewGetGroupsNameIDParamsWithContext creates a new GetGroupsNameIDParams object
// with the ability to set a context for a request.
func NewGetGroupsNameIDParamsWithContext(ctx context.Context) *GetGroupsNameIDParams {
	return &GetGroupsNameIDParams{
		Context: ctx,
	}
}

// NewGetGroupsNameIDParamsWithHTTPClient creates a new GetGroupsNameIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGroupsNameIDParamsWithHTTPClient(client *http.Client) *GetGroupsNameIDParams {
	return &GetGroupsNameIDParams{
		HTTPClient: client,
	}
}

/* GetGroupsNameIDParams contains all the parameters to send to the API endpoint
   for the get groups name ID operation.

   Typically these are written to a http.Request.
*/
type GetGroupsNameIDParams struct {

	/* Name.

	   Group name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get groups name ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGroupsNameIDParams) WithDefaults() *GetGroupsNameIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get groups name ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGroupsNameIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get groups name ID params
func (o *GetGroupsNameIDParams) WithTimeout(timeout time.Duration) *GetGroupsNameIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get groups name ID params
func (o *GetGroupsNameIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get groups name ID params
func (o *GetGroupsNameIDParams) WithContext(ctx context.Context) *GetGroupsNameIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get groups name ID params
func (o *GetGroupsNameIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get groups name ID params
func (o *GetGroupsNameIDParams) WithHTTPClient(client *http.Client) *GetGroupsNameIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get groups name ID params
func (o *GetGroupsNameIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get groups name ID params
func (o *GetGroupsNameIDParams) WithName(name string) *GetGroupsNameIDParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get groups name ID params
func (o *GetGroupsNameIDParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupsNameIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
