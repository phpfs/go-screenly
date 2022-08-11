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

// Asset asset
//
// swagger:model Asset
type Asset struct {

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

// Validate validates this asset
func (m *Asset) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssetURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

func (m *Asset) validateAssetURL(formats strfmt.Registry) error {
	if swag.IsZero(m.AssetURL) { // not required
		return nil
	}

	if err := validate.MinLength("asset_url", "body", m.AssetURL, 1); err != nil {
		return err
	}

	return nil
}

func (m *Asset) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", m.ID, 1); err != nil {
		return err
	}

	return nil
}

func (m *Asset) validateMd5(formats strfmt.Registry) error {
	if swag.IsZero(m.Md5) { // not required
		return nil
	}

	if err := validate.MinLength("md5", "body", m.Md5, 1); err != nil {
		return err
	}

	return nil
}

func (m *Asset) validateSourceURL(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceURL) { // not required
		return nil
	}

	if err := validate.MinLength("source_url", "body", *m.SourceURL, 1); err != nil {
		return err
	}

	return nil
}

var assetTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","error","finished","processing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetTypeStatusPropEnum = append(assetTypeStatusPropEnum, v)
	}
}

const (

	// AssetStatusNone captures enum value "none"
	AssetStatusNone string = "none"

	// AssetStatusError captures enum value "error"
	AssetStatusError string = "error"

	// AssetStatusFinished captures enum value "finished"
	AssetStatusFinished string = "finished"

	// AssetStatusProcessing captures enum value "processing"
	AssetStatusProcessing string = "processing"
)

// prop value enum
func (m *Asset) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assetTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Asset) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var assetTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["image","video","appweb","web"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetTypeTypePropEnum = append(assetTypeTypePropEnum, v)
	}
}

const (

	// AssetTypeImage captures enum value "image"
	AssetTypeImage string = "image"

	// AssetTypeVideo captures enum value "video"
	AssetTypeVideo string = "video"

	// AssetTypeAppweb captures enum value "appweb"
	AssetTypeAppweb string = "appweb"

	// AssetTypeWeb captures enum value "web"
	AssetTypeWeb string = "web"
)

// prop value enum
func (m *Asset) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assetTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Asset) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this asset based on the context it is used
func (m *Asset) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *Asset) contextValidateAssetURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "asset_url", "body", string(m.AssetURL)); err != nil {
		return err
	}

	return nil
}

func (m *Asset) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *Asset) contextValidateFinalized(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "finalized", "body", m.Finalized); err != nil {
		return err
	}

	return nil
}

func (m *Asset) contextValidateHeight(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "height", "body", int64(m.Height)); err != nil {
		return err
	}

	return nil
}

func (m *Asset) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Asset) contextValidateMd5(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "md5", "body", string(m.Md5)); err != nil {
		return err
	}

	return nil
}

func (m *Asset) contextValidateSourceMd5(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "source_md5", "body", m.SourceMd5); err != nil {
		return err
	}

	return nil
}

func (m *Asset) contextValidateSourceSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "source_size", "body", int64(m.SourceSize)); err != nil {
		return err
	}

	return nil
}

func (m *Asset) contextValidateSourceURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "source_url", "body", m.SourceURL); err != nil {
		return err
	}

	return nil
}

func (m *Asset) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *Asset) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Asset) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", string(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *Asset) contextValidateWidth(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "width", "body", int64(m.Width)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Asset) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Asset) UnmarshalBinary(b []byte) error {
	var res Asset
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
