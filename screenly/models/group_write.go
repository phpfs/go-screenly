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

// GroupWrite group write
//
// swagger:model GroupWrite
type GroupWrite struct {

	// Id
	// Read Only: true
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// Name
	//
	// Group name.
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// Screens
	//
	//
	// Screens in this group
	// in form [
	//   {"id": "screen1_id"},
	//   {"id": "screen2_id"},
	// ]
	//
	Screens []interface{} `json:"screens"`

	// Url
	// Read Only: true
	URL string `json:"url,omitempty"`
}

// Validate validates this group write
func (m *GroupWrite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupWrite) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", m.ID, 1); err != nil {
		return err
	}

	return nil
}

func (m *GroupWrite) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this group write based on the context it is used
func (m *GroupWrite) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupWrite) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *GroupWrite) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", string(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupWrite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupWrite) UnmarshalBinary(b []byte) error {
	var res GroupWrite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
