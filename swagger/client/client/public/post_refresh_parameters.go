// Code generated by go-swagger; DO NOT EDIT.

package public

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

	"github.com/dhis2-sre/im-users/swagger/client/models"
)

// NewPostRefreshParams creates a new PostRefreshParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostRefreshParams() *PostRefreshParams {
	return &PostRefreshParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostRefreshParamsWithTimeout creates a new PostRefreshParams object
// with the ability to set a timeout on a request.
func NewPostRefreshParamsWithTimeout(timeout time.Duration) *PostRefreshParams {
	return &PostRefreshParams{
		timeout: timeout,
	}
}

// NewPostRefreshParamsWithContext creates a new PostRefreshParams object
// with the ability to set a context for a request.
func NewPostRefreshParamsWithContext(ctx context.Context) *PostRefreshParams {
	return &PostRefreshParams{
		Context: ctx,
	}
}

// NewPostRefreshParamsWithHTTPClient creates a new PostRefreshParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostRefreshParamsWithHTTPClient(client *http.Client) *PostRefreshParams {
	return &PostRefreshParams{
		HTTPClient: client,
	}
}

/* PostRefreshParams contains all the parameters to send to the API endpoint
   for the post refresh operation.

   Typically these are written to a http.Request.
*/
type PostRefreshParams struct {

	/* RefreshTokenRequest.

	   Refresh token request
	*/
	RefreshTokenRequest *models.GithubComDhis2SreImUsersPgkUserRefreshTokenRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post refresh params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRefreshParams) WithDefaults() *PostRefreshParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post refresh params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostRefreshParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post refresh params
func (o *PostRefreshParams) WithTimeout(timeout time.Duration) *PostRefreshParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post refresh params
func (o *PostRefreshParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post refresh params
func (o *PostRefreshParams) WithContext(ctx context.Context) *PostRefreshParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post refresh params
func (o *PostRefreshParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post refresh params
func (o *PostRefreshParams) WithHTTPClient(client *http.Client) *PostRefreshParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post refresh params
func (o *PostRefreshParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRefreshTokenRequest adds the refreshTokenRequest to the post refresh params
func (o *PostRefreshParams) WithRefreshTokenRequest(refreshTokenRequest *models.GithubComDhis2SreImUsersPgkUserRefreshTokenRequest) *PostRefreshParams {
	o.SetRefreshTokenRequest(refreshTokenRequest)
	return o
}

// SetRefreshTokenRequest adds the refreshTokenRequest to the post refresh params
func (o *PostRefreshParams) SetRefreshTokenRequest(refreshTokenRequest *models.GithubComDhis2SreImUsersPgkUserRefreshTokenRequest) {
	o.RefreshTokenRequest = refreshTokenRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostRefreshParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.RefreshTokenRequest != nil {
		if err := r.SetBodyParam(o.RefreshTokenRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
