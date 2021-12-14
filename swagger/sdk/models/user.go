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
	"github.com/go-openapi/validate"
)

// User User domain object defining a user
//
// swagger:model User
type User struct {

	// admin groups
	AdminGroups []*Group `json:"AdminGroups"`

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

	// email
	Email string `json:"Email,omitempty"`

	// groups
	Groups []*Group `json:"Groups"`

	// ID
	ID uint64 `json:"ID,omitempty"`

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

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
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

func (m *User) validateAdminGroups(formats strfmt.Registry) error {
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
					return ve.ValidateName("AdminGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AdminGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("CreatedAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateDeletedAt(formats strfmt.Registry) error {
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

func (m *User) validateGroups(formats strfmt.Registry) error {
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
					return ve.ValidateName("Groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("UpdatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user based on the context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdminGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeletedAt(ctx, formats); err != nil {
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

func (m *User) contextValidateAdminGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdminGroups); i++ {

		if m.AdminGroups[i] != nil {
			if err := m.AdminGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AdminGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AdminGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) contextValidateDeletedAt(ctx context.Context, formats strfmt.Registry) error {

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

func (m *User) contextValidateGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Groups); i++ {

		if m.Groups[i] != nil {
			if err := m.Groups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}