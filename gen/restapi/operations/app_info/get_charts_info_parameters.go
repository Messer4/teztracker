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

// NewGetChartsInfoParams creates a new GetChartsInfoParams object
// no default values defined in spec.
func NewGetChartsInfoParams() GetChartsInfoParams {

	return GetChartsInfoParams{}
}

// GetChartsInfoParams contains all the bound params for the get charts info operation
// typically these are obtained from a http.Request
//
// swagger:parameters getChartsInfo
type GetChartsInfoParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	Columns []string
	/*
	  Required: true
	  In: query
	*/
	From int64
	/*Not used
	  Required: true
	  In: path
	*/
	Network string
	/*
	  Required: true
	  In: query
	*/
	Period string
	/*Not used
	  Required: true
	  In: path
	*/
	Platform string
	/*
	  Required: true
	  In: query
	*/
	To int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetChartsInfoParams() beforehand.
func (o *GetChartsInfoParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qColumns, qhkColumns, _ := qs.GetOK("columns")
	if err := o.bindColumns(qColumns, qhkColumns, route.Formats); err != nil {
		res = append(res, err)
	}

	qFrom, qhkFrom, _ := qs.GetOK("from")
	if err := o.bindFrom(qFrom, qhkFrom, route.Formats); err != nil {
		res = append(res, err)
	}

	rNetwork, rhkNetwork, _ := route.Params.GetOK("network")
	if err := o.bindNetwork(rNetwork, rhkNetwork, route.Formats); err != nil {
		res = append(res, err)
	}

	qPeriod, qhkPeriod, _ := qs.GetOK("period")
	if err := o.bindPeriod(qPeriod, qhkPeriod, route.Formats); err != nil {
		res = append(res, err)
	}

	rPlatform, rhkPlatform, _ := route.Params.GetOK("platform")
	if err := o.bindPlatform(rPlatform, rhkPlatform, route.Formats); err != nil {
		res = append(res, err)
	}

	qTo, qhkTo, _ := qs.GetOK("to")
	if err := o.bindTo(qTo, qhkTo, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindColumns binds and validates array parameter Columns from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *GetChartsInfoParams) bindColumns(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("columns", "query")
	}

	var qvColumns string
	if len(rawData) > 0 {
		qvColumns = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	columnsIC := swag.SplitByFormat(qvColumns, "")

	if len(columnsIC) == 0 {
		return errors.Required("columns", "query")
	}

	var columnsIR []string
	for _, columnsIV := range columnsIC {
		columnsI := columnsIV

		columnsIR = append(columnsIR, columnsI)
	}

	o.Columns = columnsIR

	return nil
}

// bindFrom binds and validates parameter From from query.
func (o *GetChartsInfoParams) bindFrom(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("from", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("from", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("from", "query", "int64", raw)
	}
	o.From = value

	return nil
}

// bindNetwork binds and validates parameter Network from path.
func (o *GetChartsInfoParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetChartsInfoParams) validateNetwork(formats strfmt.Registry) error {

	if err := validate.Enum("network", "path", o.Network, []interface{}{"mainnet", "babylonnet", "carthagenet"}); err != nil {
		return err
	}

	return nil
}

// bindPeriod binds and validates parameter Period from query.
func (o *GetChartsInfoParams) bindPeriod(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("period", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("period", "query", raw); err != nil {
		return err
	}

	o.Period = raw

	if err := o.validatePeriod(formats); err != nil {
		return err
	}

	return nil
}

// validatePeriod carries on validations for parameter Period
func (o *GetChartsInfoParams) validatePeriod(formats strfmt.Registry) error {

	if err := validate.Enum("period", "query", o.Period, []interface{}{"D"}); err != nil {
		return err
	}

	return nil
}

// bindPlatform binds and validates parameter Platform from path.
func (o *GetChartsInfoParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetChartsInfoParams) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Enum("platform", "path", o.Platform, []interface{}{"tezos"}); err != nil {
		return err
	}

	return nil
}

// bindTo binds and validates parameter To from query.
func (o *GetChartsInfoParams) bindTo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("to", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("to", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("to", "query", "int64", raw)
	}
	o.To = value

	return nil
}
