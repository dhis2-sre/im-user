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
)

// NewSignOutParams creates a new SignOutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSignOutParams() *SignOutParams {
	return &SignOutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSignOutParamsWithTimeout creates a new SignOutParams object
// with the ability to set a timeout on a request.
func NewSignOutParamsWithTimeout(timeout time.Duration) *SignOutParams {
	return &SignOutParams{
		timeout: timeout,
	}
}

// NewSignOutParamsWithContext creates a new SignOutParams object
// with the ability to set a context for a request.
func NewSignOutParamsWithContext(ctx context.Context) *SignOutParams {
	return &SignOutParams{
		Context: ctx,
	}
}

// NewSignOutParamsWithHTTPClient creates a new SignOutParams object
// with the ability to set a custom HTTPClient for a request.
func NewSignOutParamsWithHTTPClient(client *http.Client) *SignOutParams {
	return &SignOutParams{
		HTTPClient: client,
	}
}

/* SignOutParams contains all the parameters to send to the API endpoint
   for the sign out operation.

   Typically these are written to a http.Request.
*/
type SignOutParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the sign out params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SignOutParams) WithDefaults() *SignOutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the sign out params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SignOutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the sign out params
func (o *SignOutParams) WithTimeout(timeout time.Duration) *SignOutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sign out params
func (o *SignOutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sign out params
func (o *SignOutParams) WithContext(ctx context.Context) *SignOutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sign out params
func (o *SignOutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sign out params
func (o *SignOutParams) WithHTTPClient(client *http.Client) *SignOutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sign out params
func (o *SignOutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SignOutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}