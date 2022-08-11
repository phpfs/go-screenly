// Code generated by go-swagger; DO NOT EDIT.

package screens

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

// NewScreensDeleteParams creates a new ScreensDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScreensDeleteParams() *ScreensDeleteParams {
	return &ScreensDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScreensDeleteParamsWithTimeout creates a new ScreensDeleteParams object
// with the ability to set a timeout on a request.
func NewScreensDeleteParamsWithTimeout(timeout time.Duration) *ScreensDeleteParams {
	return &ScreensDeleteParams{
		timeout: timeout,
	}
}

// NewScreensDeleteParamsWithContext creates a new ScreensDeleteParams object
// with the ability to set a context for a request.
func NewScreensDeleteParamsWithContext(ctx context.Context) *ScreensDeleteParams {
	return &ScreensDeleteParams{
		Context: ctx,
	}
}

// NewScreensDeleteParamsWithHTTPClient creates a new ScreensDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewScreensDeleteParamsWithHTTPClient(client *http.Client) *ScreensDeleteParams {
	return &ScreensDeleteParams{
		HTTPClient: client,
	}
}

/* ScreensDeleteParams contains all the parameters to send to the API endpoint
   for the screens delete operation.

   Typically these are written to a http.Request.
*/
type ScreensDeleteParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the screens delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScreensDeleteParams) WithDefaults() *ScreensDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the screens delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScreensDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the screens delete params
func (o *ScreensDeleteParams) WithTimeout(timeout time.Duration) *ScreensDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the screens delete params
func (o *ScreensDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the screens delete params
func (o *ScreensDeleteParams) WithContext(ctx context.Context) *ScreensDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the screens delete params
func (o *ScreensDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the screens delete params
func (o *ScreensDeleteParams) WithHTTPClient(client *http.Client) *ScreensDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the screens delete params
func (o *ScreensDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the screens delete params
func (o *ScreensDeleteParams) WithID(id string) *ScreensDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the screens delete params
func (o *ScreensDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ScreensDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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