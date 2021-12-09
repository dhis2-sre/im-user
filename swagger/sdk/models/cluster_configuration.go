// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration ClusterConfiguration cluster configuration
//
// swagger:model ClusterConfiguration
type ClusterConfiguration struct {

	// created at
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"CreatedAt,omitempty"`

	// deleted at
	DeletedAt *DeletedAt `json:"DeletedAt,omitempty"`

	// group ID
	GroupID uint64 `json:"GroupID,omitempty"`

	// ID
	ID uint64 `json:"ID,omitempty"`

	// kubernetes configuration
	KubernetesConfiguration []uint8 `json:"KubernetesConfiguration"`

	// updated at
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"UpdatedAt,omitempty"`
}

// Validate validates this cluster configuration
func (m *ClusterConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterConfiguration) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClusterConfiguration) validateDeletedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.DeletedAt) { // not required
		return nil
	}

	if m.DeletedAt != nil {
		if err := m.DeletedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeletedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeletedAt")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterConfiguration) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster configuration based on the context it is used
func (m *ClusterConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeletedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterConfiguration) contextValidateDeletedAt(ctx context.Context, formats strfmt.Registry) error {

	if m.DeletedAt != nil {
		if err := m.DeletedAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DeletedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DeletedAt")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterConfiguration) UnmarshalBinary(b []byte) error {
	var res ClusterConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
