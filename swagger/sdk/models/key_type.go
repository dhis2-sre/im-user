// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// KeyType KeyType represents the key type ("kty") that are supported
//
// swagger:model KeyType
type KeyType string

// Validate validates this key type
func (m KeyType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this key type based on context it is used
func (m KeyType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
