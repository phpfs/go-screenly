// Code generated by go-swagger; DO NOT EDIT.

package groups

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

// NewGroupsCreateParams creates a new GroupsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupsCreateParams() *GroupsCreateParams {
	return &GroupsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupsCreateParamsWithTimeout creates a new GroupsCreateParams object
// with the ability to set a timeout on a request.
func NewGroupsCreateParamsWithTimeout(timeout time.Duration) *GroupsCreateParams {
	return &GroupsCreateParams{
		timeout: timeout,
	}
}

// NewGroupsCreateParamsWithContext creates a new GroupsCreateParams object
// with the ability to set a context for a request.
func NewGroupsCreateParamsWithContext(ctx context.Context) *GroupsCreateParams {
	return &GroupsCreateParams{
		Context: ctx,
	}
}

// NewGroupsCreateParamsWithHTTPClient creates a new GroupsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupsCreateParamsWithHTTPClient(client *http.Client) *GroupsCreateParams {
	return &GroupsCreateParams{
		HTTPClient: client,
	}
}

/* GroupsCreateParams contains all the parameters to send to the API endpoint
   for the groups create operation.

   Typically these are written to a http.Request.
*/
type GroupsCreateParams struct {

	// Data.
	Data *models.GroupWrite

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsCreateParams) WithDefaults() *GroupsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the groups create params
func (o *GroupsCreateParams) WithTimeout(timeout time.Duration) *GroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the groups create params
func (o *GroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the groups create params
func (o *GroupsCreateParams) WithContext(ctx context.Context) *GroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the groups create params
func (o *GroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the groups create params
func (o *GroupsCreateParams) WithHTTPClient(client *http.Client) *GroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the groups create params
func (o *GroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the groups create params
func (o *GroupsCreateParams) WithData(data *models.GroupWrite) *GroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the groups create params
func (o *GroupsCreateParams) SetData(data *models.GroupWrite) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *GroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
