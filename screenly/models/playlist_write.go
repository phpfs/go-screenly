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

// PlaylistWrite playlist write
//
// swagger:model PlaylistWrite
type PlaylistWrite struct {

	// Assets
	//
	//
	// Assets attached to this playlist
	// in form
	// [
	//  {"id": "asset_id", "duration": 5},
	//  {"id": "asset2_id"}
	// ]
	// Duration is optional and set to 10 seconds by default.
	//
	Assets interface{} `json:"assets,omitempty"`

	// Duration
	//
	// Duration of the playlist in seconds. It is equal to the sum of the duration of assets inside.
	// Read Only: true
	Duration float64 `json:"duration,omitempty"`

	// Groups
	//
	//
	// Screen groups
	// to which this playlist is attached
	// in the following form:
	// [
	//  {"id": "group_id"},
	//  {"id": "group2_id"}
	// ]
	//
	Groups interface{} `json:"groups,omitempty"`

	// Id
	//
	// Identifier of the playlist
	// Read Only: true
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// Is enabled
	//
	// Set to enable/disable playlist
	IsEnabled bool `json:"is_enabled,omitempty"`

	// Predicate
	//
	// Predicate in the screenly pro format. Defaults to 'TRUE'
	// Min Length: 1
	Predicate string `json:"predicate,omitempty"`

	// Priority
	//
	// Playlists with a priority of 1 suppress any scheduled playlists with a lower priority.
	// Maximum: 1
	// Minimum: 0
	Priority *int64 `json:"priority,omitempty"`

	// Title
	//
	// Title of the playlist
	// Min Length: 1
	Title string `json:"title,omitempty"`

	// Url
	// Read Only: true
	URL string `json:"url,omitempty"`
}

// Validate validates this playlist write
func (m *PlaylistWrite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
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

func (m *PlaylistWrite) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", m.ID, 1); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistWrite) validatePredicate(formats strfmt.Registry) error {
	if swag.IsZero(m.Predicate) { // not required
		return nil
	}

	if err := validate.MinLength("predicate", "body", m.Predicate, 1); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistWrite) validatePriority(formats strfmt.Registry) error {
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

func (m *PlaylistWrite) validateTitle(formats strfmt.Registry) error {
	if swag.IsZero(m.Title) { // not required
		return nil
	}

	if err := validate.MinLength("title", "body", m.Title, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this playlist write based on the context it is used
func (m *PlaylistWrite) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDuration(ctx, formats); err != nil {
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

func (m *PlaylistWrite) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", float64(m.Duration)); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistWrite) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *PlaylistWrite) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", string(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlaylistWrite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlaylistWrite) UnmarshalBinary(b []byte) error {
	var res PlaylistWrite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}