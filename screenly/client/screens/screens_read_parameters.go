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

// NewScreensReadParams creates a new ScreensReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScreensReadParams() *ScreensReadParams {
	return &ScreensReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScreensReadParamsWithTimeout creates a new ScreensReadParams object
// with the ability to set a timeout on a request.
func NewScreensReadParamsWithTimeout(timeout time.Duration) *ScreensReadParams {
	return &ScreensReadParams{
		timeout: timeout,
	}
}

// NewScreensReadParamsWithContext creates a new ScreensReadParams object
// with the ability to set a context for a request.
func NewScreensReadParamsWithContext(ctx context.Context) *ScreensReadParams {
	return &ScreensReadParams{
		Context: ctx,
	}
}

// NewScreensReadParamsWithHTTPClient creates a new ScreensReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewScreensReadParamsWithHTTPClient(client *http.Client) *ScreensReadParams {
	return &ScreensReadParams{
		HTTPClient: client,
	}
}

/* ScreensReadParams contains all the parameters to send to the API endpoint
   for the screens read operation.

   Typically these are written to a http.Request.
*/
type ScreensReadParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the screens read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScreensReadParams) WithDefaults() *ScreensReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the screens read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScreensReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the screens read params
func (o *ScreensReadParams) WithTimeout(timeout time.Duration) *ScreensReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the screens read params
func (o *ScreensReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the screens read params
func (o *ScreensReadParams) WithContext(ctx context.Context) *ScreensReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the screens read params
func (o *ScreensReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the screens read params
func (o *ScreensReadParams) WithHTTPClient(client *http.Client) *ScreensReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the screens read params
func (o *ScreensReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the screens read params
func (o *ScreensReadParams) WithID(id string) *ScreensReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the screens read params
func (o *ScreensReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ScreensReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
