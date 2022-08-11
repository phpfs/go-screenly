// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new assets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for assets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AssetsCreate(params *AssetsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetsCreateCreated, error)

	AssetsDelete(params *AssetsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetsDeleteOK, error)

	AssetsList(params *AssetsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetsListOK, error)

	AssetsPartialUpdate(params *AssetsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetsPartialUpdateOK, error)

	AssetsRead(params *AssetsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetsReadOK, error)

	AssetsUpdate(params *AssetsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetsUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AssetsCreate creates asset

  Define a new asset by providing a URL to a web page, an image file or a video file. Assets are the content that can be scheduled on screens, such as images, video and web pages.
*/
func (a *Client) AssetsCreate(params *AssetsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssetsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "assets_create",
		Method:             "POST",
		PathPattern:        "/assets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssetsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AssetsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for assets_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AssetsDelete deletes asset

  Delete a specific asset. Assets are the content that can be scheduled on screens, such as images, video and web pages.
*/
func (a *Client) AssetsDelete(params *AssetsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssetsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "assets_delete",
		Method:             "DELETE",
		PathPattern:        "/assets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssetsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AssetsDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for assets_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AssetsList gets asset list

  Get a list of assets. Assets are the content that can be scheduled on screens, such as images, video and web pages.
*/
func (a *Client) AssetsList(params *AssetsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssetsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "assets_list",
		Method:             "GET",
		PathPattern:        "/assets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssetsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AssetsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for assets_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AssetsPartialUpdate updates asset partially

  Change specific fields of an asset. Assets are the content that can be scheduled on screens, such as images, video and web pages.
*/
func (a *Client) AssetsPartialUpdate(params *AssetsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssetsPartialUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "assets_partial_update",
		Method:             "PATCH",
		PathPattern:        "/assets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssetsPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AssetsPartialUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for assets_partial_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AssetsRead gets asset

  Get the details of a specific asset. Assets are the content that can be scheduled on screens, such as images, video and web pages.
*/
func (a *Client) AssetsRead(params *AssetsReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssetsReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "assets_read",
		Method:             "GET",
		PathPattern:        "/assets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssetsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AssetsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for assets_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AssetsUpdate updates asset

  Change an asset. Assets are the content that can be scheduled on screens, such as images, video and web pages.
*/
func (a *Client) AssetsUpdate(params *AssetsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssetsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "assets_update",
		Method:             "PUT",
		PathPattern:        "/assets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssetsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AssetsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for assets_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}