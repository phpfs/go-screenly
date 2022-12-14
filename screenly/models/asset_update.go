// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AssetUpdate asset update
//
// swagger:model AssetUpdate
type AssetUpdate struct {

	// Asset url
	// Read Only: true
	// Min Length: 1
	AssetURL string `json:"asset_url,omitempty"`

	// Disable verification
	//
	// Disables HTTPS verification when set to true.
	DisableVerification bool `json:"disable_verification,omitempty"`

	// Duration
	//
	// Duration of the asset. For videos it is equal to the duration of the whole video. For web and images this field is unused - duration of the playback is controlled by the playlist.
	// Read Only: true
	Duration *float64 `json:"duration,omitempty"`

	// Finalized
	//
	// Indicates whether processing of the asset is finalized or not.
	// Read Only: true
	Finalized *bool `json:"finalized,omitempty"`

	// Folder name
	//
	// Name of the folder the asset is located in.
	FolderName string `json:"folder_name,omitempty"`

	// Headers
	//
	// HTTP headers for web assets.
	Headers map[string]*string `json:"headers,omitempty"`

	// Height
	//
	// Height of the asset.
	// Read Only: true
	Height int64 `json:"height,omitempty"`

	// Id
	//
	// Unique ID of the asset
	// Read Only: true
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// Js injection
	//
	// Custom js code, running when asset is loaded. See examples on: https://github.com/Screenly/playground/tree/master/javascript-injectors/
	// Min Length: 1
	JsInjection string `json:"js_injection,omitempty"`

	// Md5
	//
	// MD5 checksum of the asset.
	// Read Only: true
	// Min Length: 1
	Md5 string `json:"md5,omitempty"`

	// Meta data
	//
	//
	// A key-value user-defined store for an asset.
	// {"meta_1": "meta_1 data", "meta_2": "meta_2 data"}
	//
	MetaData map[string]*string `json:"meta_data,omitempty"`

	// Source md5
	//
	// MD5 checksum of the original source file.
	// Read Only: true
	SourceMd5 *string `json:"source_md5,omitempty"`

	// Source size
	//
	// Size of the original source file in bytes.
	// Read Only: true
	SourceSize int64 `json:"source_size,omitempty"`

	// Source url
	//
	// URL of the original source file.
	// Read Only: true
	// Min Length: 1
	SourceURL *string `json:"source_url,omitempty"`

	// Status
	//
	// Indicates the current status of the asset processing
	// Read Only: true
	// Enum: [none error finished processing]
	Status string `json:"status,omitempty"`

	// Title
	//
	// Title of the asset.
	Title string `json:"title,omitempty"`

	// Type
	//
	// Type of the asset. Web, video and images are supported.
	// Read Only: true
	// Enum: [image video appweb web]
	Type *string `json:"type,omitempty"`

	// Url
	//
	// This is API URL of this asset in the following form: /v3/assets/<id>
	// Read Only: true
	URL string `json:"url,omitempty"`

	// Width
	//
	// Width of the asset.
	// Read Only: true
	Width int64 `json:"width,omitempty"`
}

// Validate validates this asset update
func (m *AssetUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJsInjection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMd5(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *AssetUpdate) validateAssetURL(formats strfmt.Registry) error {
	if swag.IsZero(m.AssetURL) { // not required
		return nil
	}

	if err := validate.MinLength("asset_url", "body", m.AssetURL, 1); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", m.ID, 1); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) validateJsInjection(formats strfmt.Registry) error {
	if swag.IsZero(m.JsInjection) { // not required
		return nil
	}

	if err := validate.MinLength("js_injection", "body", m.JsInjection, 1); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) validateMd5(formats strfmt.Registry) error {
	if swag.IsZero(m.Md5) { // not required
		return nil
	}

	if err := validate.MinLength("md5", "body", m.Md5, 1); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) validateSourceURL(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceURL) { // not required
		return nil
	}

	if err := validate.MinLength("source_url", "body", *m.SourceURL, 1); err != nil {
		return err
	}

	return nil
}

var assetUpdateTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","error","finished","processing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetUpdateTypeStatusPropEnum = append(assetUpdateTypeStatusPropEnum, v)
	}
}

const (

	// AssetUpdateStatusNone captures enum value "none"
	AssetUpdateStatusNone string = "none"

	// AssetUpdateStatusError captures enum value "error"
	AssetUpdateStatusError string = "error"

	// AssetUpdateStatusFinished captures enum value "finished"
	AssetUpdateStatusFinished string = "finished"

	// AssetUpdateStatusProcessing captures enum value "processing"
	AssetUpdateStatusProcessing string = "processing"
)

// prop value enum
func (m *AssetUpdate) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assetUpdateTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AssetUpdate) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var assetUpdateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["image","video","appweb","web"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetUpdateTypeTypePropEnum = append(assetUpdateTypeTypePropEnum, v)
	}
}

const (

	// AssetUpdateTypeImage captures enum value "image"
	AssetUpdateTypeImage string = "image"

	// AssetUpdateTypeVideo captures enum value "video"
	AssetUpdateTypeVideo string = "video"

	// AssetUpdateTypeAppweb captures enum value "appweb"
	AssetUpdateTypeAppweb string = "appweb"

	// AssetUpdateTypeWeb captures enum value "web"
	AssetUpdateTypeWeb string = "web"
)

// prop value enum
func (m *AssetUpdate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assetUpdateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AssetUpdate) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this asset update based on the context it is used
func (m *AssetUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssetURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFinalized(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMd5(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceMd5(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWidth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetUpdate) contextValidateAssetURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "asset_url", "body", string(m.AssetURL)); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) contextValidateFinalized(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "finalized", "body", m.Finalized); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) contextValidateHeight(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "height", "body", int64(m.Height)); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) contextValidateMd5(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "md5", "body", string(m.Md5)); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) contextValidateSourceMd5(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "source_md5", "body", m.SourceMd5); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) contextValidateSourceSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "source_size", "body", int64(m.SourceSize)); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) contextValidateSourceURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "source_url", "body", m.SourceURL); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", string(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *AssetUpdate) contextValidateWidth(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "width", "body", int64(m.Width)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetUpdate) UnmarshalBinary(b []byte) error {
	var res AssetUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
