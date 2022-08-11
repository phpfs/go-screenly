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

// PlaylistRead playlist read
//
// swagger:model PlaylistRead
type PlaylistRead struct {

	// Assets attached to this playlist
	Assets []*PlaylistItem `json:"assets"`

	// Duration
	//
	// Duration of the playlist in seconds. It is equal to the sum of the duration of assets inside.
	// Read Only: true
	Duration float64 `json:"duration,omitempty"`

	// Screen groups to which this playlist is attached
	Groups []*Group `json:"groups"`

	// Id
	//
	// Unique ID of the playlist
	// Read Only: true
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// Is enabled
	//
	// Set to enable/disable playlist
	// Required: true
	IsEnabled *bool `json:"is_enabled"`

	// Predicate
	//
	// Predicate in the screenly pro format
	// Required: true
	// Min Length: 1
	Predicate *string `json:"predicate"`

	// Priority
	//
	// Playlists with a priority of 1 suppress any scheduled playlists with a lower priority.
	// Maximum: 1
	// Minimum: 0
	Priority *int64 `json:"priority,omitempty"`

	// Title
	//
	// Title of the playlist
	// Required: true
	// Min Length: 1
	Title *string `json:"title"`

	// Url
	// Read Only: true
	URL string `json:"url,omitempty"`
}

// Validate validates this playlist read
func (m *PlaylistRead) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePredicate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlaylistRead) validateAssets(formats strfmt.Registry) error {
	if swag.IsZero(m.Assets) { // not required
		return nil
	}

	for i := 0; i < len(m.Assets); i++ {
		if swag.IsZero(m.Assets[i]) { // not required
			continue
		}

		if m.Assets[i] != nil {
			if err := m.Assets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlaylistRead) validateGroups(formats strfmt.Registry) error {
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

func (m *PlaylistRead) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", m.ID, 1); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistRead) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("is_enabled", "body", m.IsEnabled); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistRead) validatePredicate(formats strfmt.Registry) error {

	if err := validate.Required("predicate", "body", m.Predicate); err != nil {
		return err
	}

	if err := validate.MinLength("predicate", "body", *m.Predicate, 1); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistRead) validatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	if err := validate.MinimumInt("priority", "body", *m.Priority, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("priority", "body", *m.Priority, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistRead) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this playlist read based on the context it is used
func (m *PlaylistRead) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *PlaylistRead) contextValidateAssets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Assets); i++ {

		if m.Assets[i] != nil {
			if err := m.Assets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PlaylistRead) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", float64(m.Duration)); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistRead) contextValidateGroups(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PlaylistRead) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistRead) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", string(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlaylistRead) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlaylistRead) UnmarshalBinary(b []byte) error {
	var res PlaylistRead
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
