// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtoUser dto user
//
// swagger:model dto.User
type DtoUser struct {

	// admin groups
	AdminGroups []*DtoGroup `json:"adminGroups"`

	// email
	Email string `json:"email,omitempty"`

	// groups
	Groups []*DtoGroup `json:"groups"`

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this dto user
func (m *DtoUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtoUser) validateAdminGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.AdminGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.AdminGroups); i++ {
		if swag.IsZero(m.AdminGroups[i]) { // not required
			continue
		}

		if m.AdminGroups[i] != nil {
			if err := m.AdminGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adminGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("adminGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DtoUser) validateGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dto user based on the context it is used
func (m *DtoUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdminGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtoUser) contextValidateAdminGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdminGroups); i++ {

		if m.AdminGroups[i] != nil {
			if err := m.AdminGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adminGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("adminGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DtoUser) contextValidateGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Groups); i++ {

		if m.Groups[i] != nil {
			if err := m.Groups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DtoUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtoUser) UnmarshalBinary(b []byte) error {
	var res DtoUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
