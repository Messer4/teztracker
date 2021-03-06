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

// EndorsementRight endorsement right
// swagger:model EndorsementRight
type EndorsementRight struct {

	// delegate
	Delegate string `json:"delegate,omitempty"`

	// estimated time
	// Format: date-time
	EstimatedTime strfmt.DateTime `json:"estimated_time,omitempty"`

	// level
	Level int64 `json:"level,omitempty"`

	// slots
	Slots []int64 `json:"slots"`
}

// Validate validates this endorsement right
func (m *EndorsementRight) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEstimatedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndorsementRight) validateEstimatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EstimatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("estimated_time", "body", "date-time", m.EstimatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndorsementRight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndorsementRight) UnmarshalBinary(b []byte) error {
	var res EndorsementRight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
