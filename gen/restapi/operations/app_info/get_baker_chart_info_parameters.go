// Code generated by go-swagger; DO NOT EDIT.

package app_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBakerChartInfoParams creates a new GetBakerChartInfoParams object
// with the default values initialized.
func NewGetBakerChartInfoParams() GetBakerChartInfoParams {

	var (
		// initialize parameters with default values

		limitDefault = int64(20)
	)

	return GetBakerChartInfoParams{
		Limit: &limitDefault,
	}
}

// GetBakerChartInfoParams contains all the bound params for the get baker chart info operation
// typically these are obtained from a http.Request
//
// swagger:parameters getBakerChartInfo
type GetBakerChartInfoParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	  Default: 20
	*/
	Limit *int64
	/*Not used
	  Required: true
	  In: path
	*/
	Network string
	/*Not used
	  Required: true
	  In: path
	*/
	Platform string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBakerChartInfoParams() beforehand.
func (o *GetBakerChartInfoParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	rNetwork, rhkNetwork, _ := route.Params.GetOK("network")
	if err := o.bindNetwork(rNetwork, rhkNetwork, route.Formats); err != nil {
		res = append(res, err)
	}

	rPlatform, rhkPlatform, _ := route.Params.GetOK("platform")
	if err := o.bindPlatform(rPlatform, rhkPlatform, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetBakerChartInfoParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetBakerChartInfoParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

// bindNetwork binds and validates parameter Network from path.
func (o *GetBakerChartInfoParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Network = raw

	if err := o.validateNetwork(formats); err != nil {
		return err
	}

	return nil
}

// validateNetwork carries on validations for parameter Network
func (o *GetBakerChartInfoParams) validateNetwork(formats strfmt.Registry) error {

	if err := validate.Enum("network", "path", o.Network, []interface{}{"mainnet", "delphi"}); err != nil {
		return err
	}

	return nil
}

// bindPlatform binds and validates parameter Platform from path.
func (o *GetBakerChartInfoParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Platform = raw

	if err := o.validatePlatform(formats); err != nil {
		return err
	}

	return nil
}

// validatePlatform carries on validations for parameter Platform
func (o *GetBakerChartInfoParams) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Enum("platform", "path", o.Platform, []interface{}{"tezos"}); err != nil {
		return err
	}

	return nil
}
