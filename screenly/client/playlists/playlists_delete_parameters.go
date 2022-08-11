// Code generated by go-swagger; DO NOT EDIT.

package playlists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPlaylistsDeleteParams creates a new PlaylistsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPlaylistsDeleteParams() *PlaylistsDeleteParams {
	return &PlaylistsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPlaylistsDeleteParamsWithTimeout creates a new PlaylistsDeleteParams object
// with the ability to set a timeout on a request.
func NewPlaylistsDeleteParamsWithTimeout(timeout time.Duration) *PlaylistsDeleteParams {
	return &PlaylistsDeleteParams{
		timeout: timeout,
	}
}

// NewPlaylistsDeleteParamsWithContext creates a new PlaylistsDeleteParams object
// with the ability to set a context for a request.
func NewPlaylistsDeleteParamsWithContext(ctx context.Context) *PlaylistsDeleteParams {
	return &PlaylistsDeleteParams{
		Context: ctx,
	}
}

// NewPlaylistsDeleteParamsWithHTTPClient creates a new PlaylistsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPlaylistsDeleteParamsWithHTTPClient(client *http.Client) *PlaylistsDeleteParams {
	return &PlaylistsDeleteParams{
		HTTPClient: client,
	}
}

/* PlaylistsDeleteParams contains all the parameters to send to the API endpoint
   for the playlists delete operation.

   Typically these are written to a http.Request.
*/
type PlaylistsDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the playlists delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlaylistsDeleteParams) WithDefaults() *PlaylistsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the playlists delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlaylistsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the playlists delete params
func (o *PlaylistsDeleteParams) WithTimeout(timeout time.Duration) *PlaylistsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the playlists delete params
func (o *PlaylistsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the playlists delete params
func (o *PlaylistsDeleteParams) WithContext(ctx context.Context) *PlaylistsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the playlists delete params
func (o *PlaylistsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the playlists delete params
func (o *PlaylistsDeleteParams) WithHTTPClient(client *http.Client) *PlaylistsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the playlists delete params
func (o *PlaylistsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the playlists delete params
func (o *PlaylistsDeleteParams) WithID(id string) *PlaylistsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the playlists delete params
func (o *PlaylistsDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PlaylistsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
