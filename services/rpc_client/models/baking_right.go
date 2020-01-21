// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BakingRight baking right
// swagger:model BakingRight
type BakingRight struct {

	// delegate
	Delegate string `json:"delegate,omitempty"`

	// estimated time
	// Format: date-time
	EstimatedTime strfmt.DateTime `json:"estimated_time,omitempty"`

	// level
	Level int64 `json:"level,omitempty"`

	// priority
	Priority int64 `json:"priority,omitempty"`
}

// Validate validates this baking right
func (m *BakingRight) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEstimatedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BakingRight) validateEstimatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EstimatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("estimated_time", "body", "date-time", m.EstimatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BakingRight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BakingRight) UnmarshalBinary(b []byte) error {
	var res BakingRight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
