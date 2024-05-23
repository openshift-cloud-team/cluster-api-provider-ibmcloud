// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

// NewPcloudPvminstancesConsoleGetParams creates a new PcloudPvminstancesConsoleGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesConsoleGetParams() *PcloudPvminstancesConsoleGetParams {
	return &PcloudPvminstancesConsoleGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesConsoleGetParamsWithTimeout creates a new PcloudPvminstancesConsoleGetParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesConsoleGetParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesConsoleGetParams {
	return &PcloudPvminstancesConsoleGetParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesConsoleGetParamsWithContext creates a new PcloudPvminstancesConsoleGetParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesConsoleGetParamsWithContext(ctx context.Context) *PcloudPvminstancesConsoleGetParams {
	return &PcloudPvminstancesConsoleGetParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesConsoleGetParamsWithHTTPClient creates a new PcloudPvminstancesConsoleGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesConsoleGetParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesConsoleGetParams {
	return &PcloudPvminstancesConsoleGetParams{
		HTTPClient: client,
	}
}

/*
PcloudPvminstancesConsoleGetParams contains all the parameters to send to the API endpoint

	for the pcloud pvminstances console get operation.

	Typically these are written to a http.Request.
*/
type PcloudPvminstancesConsoleGetParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* PvmInstanceID.

	   PCloud PVM Instance ID
	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud pvminstances console get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesConsoleGetParams) WithDefaults() *PcloudPvminstancesConsoleGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances console get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesConsoleGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances console get params
func (o *PcloudPvminstancesConsoleGetParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesConsoleGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances console get params
func (o *PcloudPvminstancesConsoleGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances console get params
func (o *PcloudPvminstancesConsoleGetParams) WithContext(ctx context.Context) *PcloudPvminstancesConsoleGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances console get params
func (o *PcloudPvminstancesConsoleGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances console get params
func (o *PcloudPvminstancesConsoleGetParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesConsoleGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances console get params
func (o *PcloudPvminstancesConsoleGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances console get params
func (o *PcloudPvminstancesConsoleGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesConsoleGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances console get params
func (o *PcloudPvminstancesConsoleGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances console get params
func (o *PcloudPvminstancesConsoleGetParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesConsoleGetParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances console get params
func (o *PcloudPvminstancesConsoleGetParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesConsoleGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
