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

	"github.com/phpfs/go-screenly/screenly/models"
)

// NewScreensUpdateParams creates a new ScreensUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewScreensUpdateParams() *ScreensUpdateParams {
	return &ScreensUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewScreensUpdateParamsWithTimeout creates a new ScreensUpdateParams object
// with the ability to set a timeout on a request.
func NewScreensUpdateParamsWithTimeout(timeout time.Duration) *ScreensUpdateParams {
	return &ScreensUpdateParams{
		timeout: timeout,
	}
}

// NewScreensUpdateParamsWithContext creates a new ScreensUpdateParams object
// with the ability to set a context for a request.
func NewScreensUpdateParamsWithContext(ctx context.Context) *ScreensUpdateParams {
	return &ScreensUpdateParams{
		Context: ctx,
	}
}

// NewScreensUpdateParamsWithHTTPClient creates a new ScreensUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewScreensUpdateParamsWithHTTPClient(client *http.Client) *ScreensUpdateParams {
	return &ScreensUpdateParams{
		HTTPClient: client,
	}
}

/* ScreensUpdateParams contains all the parameters to send to the API endpoint
   for the screens update operation.

   Typically these are written to a http.Request.
*/
type ScreensUpdateParams struct {

	// Data.
	Data *models.ScreenEndpoint

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the screens update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScreensUpdateParams) WithDefaults() *ScreensUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the screens update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ScreensUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the screens update params
func (o *ScreensUpdateParams) WithTimeout(timeout time.Duration) *ScreensUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the screens update params
func (o *ScreensUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the screens update params
func (o *ScreensUpdateParams) WithContext(ctx context.Context) *ScreensUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the screens update params
func (o *ScreensUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the screens update params
func (o *ScreensUpdateParams) WithHTTPClient(client *http.Client) *ScreensUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the screens update params
func (o *ScreensUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the screens update params
func (o *ScreensUpdateParams) WithData(data *models.ScreenEndpoint) *ScreensUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the screens update params
func (o *ScreensUpdateParams) SetData(data *models.ScreenEndpoint) {
	o.Data = data
}

// WithID adds the id to the screens update params
func (o *ScreensUpdateParams) WithID(id string) *ScreensUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the screens update params
func (o *ScreensUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ScreensUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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