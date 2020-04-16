// Code generated by go-swagger; DO NOT EDIT.

package endorsing_rights

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

// NewGetEndorsingRightsParams creates a new GetEndorsingRightsParams object
// with the default values initialized.
func NewGetEndorsingRightsParams() *GetEndorsingRightsParams {
	var ()
	return &GetEndorsingRightsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEndorsingRightsParamsWithTimeout creates a new GetEndorsingRightsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEndorsingRightsParamsWithTimeout(timeout time.Duration) *GetEndorsingRightsParams {
	var ()
	return &GetEndorsingRightsParams{

		timeout: timeout,
	}
}

// NewGetEndorsingRightsParamsWithContext creates a new GetEndorsingRightsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEndorsingRightsParamsWithContext(ctx context.Context) *GetEndorsingRightsParams {
	var ()
	return &GetEndorsingRightsParams{

		Context: ctx,
	}
}

// NewGetEndorsingRightsParamsWithHTTPClient creates a new GetEndorsingRightsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEndorsingRightsParamsWithHTTPClient(client *http.Client) *GetEndorsingRightsParams {
	var ()
	return &GetEndorsingRightsParams{
		HTTPClient: client,
	}
}

/*GetEndorsingRightsParams contains all the parameters to send to the API endpoint
for the get endorsing rights operation typically these are written to a http.Request
*/
type GetEndorsingRightsParams struct {

	/*Block*/
	Block string
	/*Cycle*/
	Cycle *string
	/*Delegate*/
	Delegate *string
	/*Level*/
	Level []string
	/*Network*/
	Network string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get endorsing rights params
func (o *GetEndorsingRightsParams) WithTimeout(timeout time.Duration) *GetEndorsingRightsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get endorsing rights params
func (o *GetEndorsingRightsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get endorsing rights params
func (o *GetEndorsingRightsParams) WithContext(ctx context.Context) *GetEndorsingRightsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get endorsing rights params
func (o *GetEndorsingRightsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get endorsing rights params
func (o *GetEndorsingRightsParams) WithHTTPClient(client *http.Client) *GetEndorsingRightsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get endorsing rights params
func (o *GetEndorsingRightsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBlock adds the block to the get endorsing rights params
func (o *GetEndorsingRightsParams) WithBlock(block string) *GetEndorsingRightsParams {
	o.SetBlock(block)
	return o
}

// SetBlock adds the block to the get endorsing rights params
func (o *GetEndorsingRightsParams) SetBlock(block string) {
	o.Block = block
}

// WithCycle adds the cycle to the get endorsing rights params
func (o *GetEndorsingRightsParams) WithCycle(cycle *string) *GetEndorsingRightsParams {
	o.SetCycle(cycle)
	return o
}

// SetCycle adds the cycle to the get endorsing rights params
func (o *GetEndorsingRightsParams) SetCycle(cycle *string) {
	o.Cycle = cycle
}

// WithDelegate adds the delegate to the get endorsing rights params
func (o *GetEndorsingRightsParams) WithDelegate(delegate *string) *GetEndorsingRightsParams {
	o.SetDelegate(delegate)
	return o
}

// SetDelegate adds the delegate to the get endorsing rights params
func (o *GetEndorsingRightsParams) SetDelegate(delegate *string) {
	o.Delegate = delegate
}

// WithLevel adds the level to the get endorsing rights params
func (o *GetEndorsingRightsParams) WithLevel(level []string) *GetEndorsingRightsParams {
	o.SetLevel(level)
	return o
}

// SetLevel adds the level to the get endorsing rights params
func (o *GetEndorsingRightsParams) SetLevel(level []string) {
	o.Level = level
}

// WithNetwork adds the network to the get endorsing rights params
func (o *GetEndorsingRightsParams) WithNetwork(network string) *GetEndorsingRightsParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the get endorsing rights params
func (o *GetEndorsingRightsParams) SetNetwork(network string) {
	o.Network = network
}

// WriteToRequest writes these params to a swagger request
func (o *GetEndorsingRightsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param block
	if err := r.SetPathParam("block", o.Block); err != nil {
		return err
	}

	if o.Cycle != nil {

		// query param cycle
		var qrCycle string
		if o.Cycle != nil {
			qrCycle = *o.Cycle
		}
		qCycle := qrCycle
		if qCycle != "" {
			if err := r.SetQueryParam("cycle", qCycle); err != nil {
				return err
			}
		}

	}

	if o.Delegate != nil {

		// query param delegate
		var qrDelegate string
		if o.Delegate != nil {
			qrDelegate = *o.Delegate
		}
		qDelegate := qrDelegate
		if qDelegate != "" {
			if err := r.SetQueryParam("delegate", qDelegate); err != nil {
				return err
			}
		}

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