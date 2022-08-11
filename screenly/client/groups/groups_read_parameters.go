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
)

// NewGroupsReadParams creates a new GroupsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupsReadParams() *GroupsReadParams {
	return &GroupsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupsReadParamsWithTimeout creates a new GroupsReadParams object
// with the ability to set a timeout on a request.
func NewGroupsReadParamsWithTimeout(timeout time.Duration) *GroupsReadParams {
	return &GroupsReadParams{
		timeout: timeout,
	}
}

// NewGroupsReadParamsWithContext creates a new GroupsReadParams object
// with the ability to set a context for a request.
func NewGroupsReadParamsWithContext(ctx context.Context) *GroupsReadParams {
	return &GroupsReadParams{
		Context: ctx,
	}
}

// NewGroupsReadParamsWithHTTPClient creates a new GroupsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupsReadParamsWithHTTPClient(client *http.Client) *GroupsReadParams {
	return &GroupsReadParams{
		HTTPClient: client,
	}
}

/* GroupsReadParams contains all the parameters to send to the API endpoint
   for the groups read operation.

   Typically these are written to a http.Request.
*/
type GroupsReadParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the groups read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsReadParams) WithDefaults() *GroupsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the groups read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the groups read params
func (o *GroupsReadParams) WithTimeout(timeout time.Duration) *GroupsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the groups read params
func (o *GroupsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the groups read params
func (o *GroupsReadParams) WithContext(ctx context.Context) *GroupsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the groups read params
func (o *GroupsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the groups read params
func (o *GroupsReadParams) WithHTTPClient(client *http.Client) *GroupsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the groups read params
func (o *GroupsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the groups read params
func (o *GroupsReadParams) WithID(id string) *GroupsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the groups read params
func (o *GroupsReadParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GroupsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
