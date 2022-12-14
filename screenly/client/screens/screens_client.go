// Code generated by go-swagger; DO NOT EDIT.

package screens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new screens API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for screens API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ScreensCreate(params *ScreensCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ScreensCreateCreated, error)

	ScreensDelete(params *ScreensDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ScreensDeleteOK, error)

	ScreensList(params *ScreensListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ScreensListOK, error)

	ScreensRead(params *ScreensReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ScreensReadOK, error)

	ScreensUpdate(params *ScreensUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ScreensUpdateOK, error)

	ScreensUpload(params *ScreensUploadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ScreensUploadCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ScreensCreate creates screen

  Create new screen. Screens are a representation of the physical device running Screenly software.
*/
func (a *Client) ScreensCreate(params *ScreensCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ScreensCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScreensCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "screens_create",
		Method:             "POST",
		PathPattern:        "/screens/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScreensCreateReader{formats: a.formats},
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
	success, ok := result.(*ScreensCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for screens_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScreensDelete deletes screen

  Delete a specific screen. Screens are a representation of the physical device running Screenly software.
*/
func (a *Client) ScreensDelete(params *ScreensDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ScreensDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScreensDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "screens_delete",
		Method:             "DELETE",
		PathPattern:        "/screens/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScreensDeleteReader{formats: a.formats},
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
	success, ok := result.(*ScreensDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for screens_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScreensList gets screen list

  Get a list of screens. Screens are a representation of the physical device running Screenly software.
*/
func (a *Client) ScreensList(params *ScreensListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ScreensListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScreensListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "screens_list",
		Method:             "GET",
		PathPattern:        "/screens/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScreensListReader{formats: a.formats},
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
	success, ok := result.(*ScreensListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for screens_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScreensRead gets screen

  Get the details of a specific screen. Screens are a representation of the physical device running Screenly software.
*/
func (a *Client) ScreensRead(params *ScreensReadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ScreensReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScreensReadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "screens_read",
		Method:             "GET",
		PathPattern:        "/screens/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScreensReadReader{formats: a.formats},
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
	success, ok := result.(*ScreensReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for screens_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScreensUpdate updates screen

  Change a screen. Screens are a representation of the physical device running Screenly software.
*/
func (a *Client) ScreensUpdate(params *ScreensUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ScreensUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScreensUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "screens_update",
		Method:             "PUT",
		PathPattern:        "/screens/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScreensUpdateReader{formats: a.formats},
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
	success, ok := result.(*ScreensUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for screens_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ScreensUpload screens upload API
*/
func (a *Client) ScreensUpload(params *ScreensUploadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ScreensUploadCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScreensUploadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "screens_upload",
		Method:             "POST",
		PathPattern:        "/screens/upload/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScreensUploadReader{formats: a.formats},
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
	success, ok := result.(*ScreensUploadCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for screens_upload: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
