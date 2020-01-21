// Code generated by go-swagger; DO NOT EDIT.

package baking_rights

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBakingRightsParams creates a new GetBakingRightsParams object
// with the default values initialized.
func NewGetBakingRightsParams() *GetBakingRightsParams {
	var ()
	return &GetBakingRightsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBakingRightsParamsWithTimeout creates a new GetBakingRightsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBakingRightsParamsWithTimeout(timeout time.Duration) *GetBakingRightsParams {
	var ()
	return &GetBakingRightsParams{

		timeout: timeout,
	}
}

// NewGetBakingRightsParamsWithContext creates a new GetBakingRightsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBakingRightsParamsWithContext(ctx context.Context) *GetBakingRightsParams {
	var ()
	return &GetBakingRightsParams{

		Context: ctx,
	}
}

// NewGetBakingRightsParamsWithHTTPClient creates a new GetBakingRightsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBakingRightsParamsWithHTTPClient(client *http.Client) *GetBakingRightsParams {
	var ()
	return &GetBakingRightsParams{
		HTTPClient: client,
	}
}

/*GetBakingRightsParams contains all the parameters to send to the API endpoint
for the get baking rights operation typically these are written to a http.Request
*/
type GetBakingRightsParams struct {

	/*All*/
	All *bool
	/*Block*/
	Block string
	/*Level*/
	Level []string
	/*Network*/
	Network string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get baking rights params
func (o *GetBakingRightsParams) WithTimeout(timeout time.Duration) *GetBakingRightsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get baking rights params
func (o *GetBakingRightsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get baking rights params
func (o *GetBakingRightsParams) WithContext(ctx context.Context) *GetBakingRightsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get baking rights params
func (o *GetBakingRightsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get baking rights params
func (o *GetBakingRightsParams) WithHTTPClient(client *http.Client) *GetBakingRightsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get baking rights params
func (o *GetBakingRightsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAll adds the all to the get baking rights params
func (o *GetBakingRightsParams) WithAll(all *bool) *GetBakingRightsParams {
	o.SetAll(all)
	return o
}

// SetAll adds the all to the get baking rights params
func (o *GetBakingRightsParams) SetAll(all *bool) {
	o.All = all
}

// WithBlock adds the block to the get baking rights params
func (o *GetBakingRightsParams) WithBlock(block string) *GetBakingRightsParams {
	o.SetBlock(block)
	return o
}

// SetBlock adds the block to the get baking rights params
func (o *GetBakingRightsParams) SetBlock(block string) {
	o.Block = block
}

// WithLevel adds the level to the get baking rights params
func (o *GetBakingRightsParams) WithLevel(level []string) *GetBakingRightsParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the get baking rights params
func (o *GetBakingRightsParams) SetLevel(level []string) {
	o.Level = level
}

// WithNetwork adds the network to the get baking rights params
func (o *GetBakingRightsParams) WithNetwork(network string) *GetBakingRightsParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the get baking rights params
func (o *GetBakingRightsParams) SetNetwork(network string) {
	o.Network = network
}

// WriteToRequest writes these params to a swagger request
func (o *GetBakingRightsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.All != nil {

		// query param all
		var qrAll bool
		if o.All != nil {
			qrAll = *o.All
		}
		qAll := swag.FormatBool(qrAll)
		if qAll != "" {
			if err := r.SetQueryParam("all", qAll); err != nil {
				return err
			}
		}

	}

	// path param block
	if err := r.SetPathParam("block", o.Block); err != nil {
		return err
	}

	valuesLevel := o.Level

	joinedLevel := swag.JoinByFormat(valuesLevel, "multi")
	// query array param level
	if err := r.SetQueryParam("level", joinedLevel...); err != nil {
		return err
	}

	// path param network
	if err := r.SetPathParam("network", o.Network); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
