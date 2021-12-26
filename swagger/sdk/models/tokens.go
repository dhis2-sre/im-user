// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Tokens Tokens domain object defining user tokens
//
// swagger:model Tokens
type Tokens struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// expires in
	ExpiresIn uint64 `json:"expires_in,omitempty"`

	// refresh token
	RefreshToken string `json:"refresh_token,omitempty"`

	// token type
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this tokens
func (m *Tokens) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tokens based on context it is used
func (m *Tokens) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Tokens) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tokens) UnmarshalBinary(b []byte) error {
	var res Tokens
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}