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

// ScreenDetailed screen detailed
//
// swagger:model ScreenDetailed
type ScreenDetailed struct {

	// Lat/Long tuple of device geolocation.
	Coords []float64 `json:"coords"`

	// Groups to which this screen is attached
	Groups []*Group `json:"groups"`

	// Hardware version
	//
	// Hardware version of the device. Currently only Raspberry Pi versions are detected
	HardwareVersion string `json:"hardware_version,omitempty"`

	// Hostname
	//
	// Unique hostname of the device, assigned by screenly
	// Min Length: 1
	Hostname string `json:"hostname,omitempty"`

	// Id
	//
	// Unique ID of the screen
	// Read Only: true
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// In sync
	//
	// Shows if the screen is in sync
	// Required: true
	InSync *bool `json:"in_sync"`

	// Is enabled
	//
	// Set to enable/disable screen
	// Required: true
	IsEnabled *bool `json:"is_enabled"`

	// Last ip
	//
	// IP used used by the device
	LastIP string `json:"last_ip,omitempty"`

	// Last ping
	//
	// The last time screen has pinged the server
	LastPing *string `json:"last_ping,omitempty"`

	// Last screenshot
	//
	// The last screenshot captured from the device
	LastScreenshot string `json:"last_screenshot,omitempty"`

	// Last screenshot time
	//
	// Time when last screenshot was taken
	LastScreenshotTime *string `json:"last_screenshot_time,omitempty"`

	// Load avg
	//
	// Average load of the device
	LoadAvg string `json:"load_avg,omitempty"`

	// Local ip
	//
	// Device IP in the local network
	LocalIP string `json:"local_ip,omitempty"`

	// Location
	//
	// Geographic location auto-detected for device
	Location string `json:"location,omitempty"`

	// Mac
	//
	// MAC address of the device
	Mac string `json:"mac,omitempty"`

	// Name
	//
	// Name of the screen
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// Software version
	//
	// Software version of the device
	SoftwareVersion string `json:"software_version,omitempty"`

	// Status
	//
	// Current status of the screen. ONLINE/OUT OF SYNC/OFFLINE/DISABLED
	// Required: true
	// Min Length: 1
	Status *string `json:"status"`

	// Team
	//
	// Team ID of the device
	// Min Length: 1
	Team string `json:"team,omitempty"`

	// Timezone
	//
	// Timezone detected for the device
	Timezone string `json:"timezone,omitempty"`

	// Type
	//
	// Type of the screen. HARDWARE/VIRTUAL
	// Required: true
	// Min Length: 1
	Type *string `json:"type"`

	// Uptime
	//
	// Uptime of the device
	Uptime string `json:"uptime,omitempty"`

	// Ws open
	//
	// Shows if device has a websocket connection
	WsOpen bool `json:"ws_open,omitempty"`
}

// Validate validates this screen detailed
func (m *ScreenDetailed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInSync(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScreenDetailed) validateGroups(formats strfmt.Registry) error {
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

func (m *ScreenDetailed) validateHostname(formats strfmt.Registry) error {
	if swag.IsZero(m.Hostname) { // not required
		return nil
	}

	if err := validate.MinLength("hostname", "body", m.Hostname, 1); err != nil {
		return err
	}

	return nil
}

func (m *ScreenDetailed) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", m.ID, 1); err != nil {
		return err
	}

	return nil
}

func (m *ScreenDetailed) validateInSync(formats strfmt.Registry) error {

	if err := validate.Required("in_sync", "body", m.InSync); err != nil {
		return err
	}

	return nil
}

func (m *ScreenDetailed) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("is_enabled", "body", m.IsEnabled); err != nil {
		return err
	}

	return nil
}

func (m *ScreenDetailed) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *ScreenDetailed) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.MinLength("status", "body", *m.Status, 1); err != nil {
		return err
	}

	return nil
}

func (m *ScreenDetailed) validateTeam(formats strfmt.Registry) error {
	if swag.IsZero(m.Team) { // not required
		return nil
	}

	if err := validate.MinLength("team", "body", m.Team, 1); err != nil {
		return err
	}

	return nil
}

func (m *ScreenDetailed) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.MinLength("type", "body", *m.Type, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this screen detailed based on the context it is used
func (m *ScreenDetailed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScreenDetailed) contextValidateGroups(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ScreenDetailed) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScreenDetailed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScreenDetailed) UnmarshalBinary(b []byte) error {
	var res ScreenDetailed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
