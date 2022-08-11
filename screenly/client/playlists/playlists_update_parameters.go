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

	"github.com/phpfs/go-screenly/screenly/models"
)

// NewPlaylistsUpdateParams creates a new PlaylistsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPlaylistsUpdateParams() *PlaylistsUpdateParams {
	return &PlaylistsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPlaylistsUpdateParamsWithTimeout creates a new PlaylistsUpdateParams object
// with the ability to set a timeout on a request.
func NewPlaylistsUpdateParamsWithTimeout(timeout time.Duration) *PlaylistsUpdateParams {
	return &PlaylistsUpdateParams{
		timeout: timeout,
	}
}

// NewPlaylistsUpdateParamsWithContext creates a new PlaylistsUpdateParams object
// with the ability to set a context for a request.
func NewPlaylistsUpdateParamsWithContext(ctx context.Context) *PlaylistsUpdateParams {
	return &PlaylistsUpdateParams{
		Context: ctx,
	}
}

// NewPlaylistsUpdateParamsWithHTTPClient creates a new PlaylistsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPlaylistsUpdateParamsWithHTTPClient(client *http.Client) *PlaylistsUpdateParams {
	return &PlaylistsUpdateParams{
		HTTPClient: client,
	}
}

/* PlaylistsUpdateParams contains all the parameters to send to the API endpoint
   for the playlists update operation.

   Typically these are written to a http.Request.
*/
type PlaylistsUpdateParams struct {

	// Data.
	Data *models.PlaylistWrite

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the playlists update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlaylistsUpdateParams) WithDefaults() *PlaylistsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the playlists update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PlaylistsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the playlists update params
func (o *PlaylistsUpdateParams) WithTimeout(timeout time.Duration) *PlaylistsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the playlists update params
func (o *PlaylistsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the playlists update params
func (o *PlaylistsUpdateParams) WithContext(ctx context.Context) *PlaylistsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the playlists update params
func (o *PlaylistsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the playlists update params
func (o *PlaylistsUpdateParams) WithHTTPClient(client *http.Client) *PlaylistsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the playlists update params
func (o *PlaylistsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the playlists update params
func (o *PlaylistsUpdateParams) WithData(data *models.PlaylistWrite) *PlaylistsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the playlists update params
func (o *PlaylistsUpdateParams) SetData(data *models.PlaylistWrite) {
	o.Data = data
}

// WithID adds the id to the playlists update params
func (o *PlaylistsUpdateParams) WithID(id string) *PlaylistsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the playlists update params
func (o *PlaylistsUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PlaylistsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
