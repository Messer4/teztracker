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

// BlockPriorityCounter block priority counter
// swagger:model BlockPriorityCounter
type BlockPriorityCounter struct {

	// first priority
	// Required: true
	FirstPriority *int64 `json:"firstPriority"`

	// second priority
	// Required: true
	SecondPriority *int64 `json:"secondPriority"`

	// zero priority
	// Required: true
	ZeroPriority *int64 `json:"zeroPriority"`
}

// Validate validates this block priority counter
func (m *BlockPriorityCounter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirstPriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondPriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZeroPriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockPriorityCounter) validateFirstPriority(formats strfmt.Registry) error {

	if err := validate.Required("firstPriority", "body", m.FirstPriority); err != nil {
		return err
	}

	return nil
}

func (m *BlockPriorityCounter) validateSecondPriority(formats strfmt.Registry) error {

	if err := validate.Required("secondPriority", "body", m.SecondPriority); err != nil {
		return err
	}

	return nil
}

func (m *BlockPriorityCounter) validateZeroPriority(formats strfmt.Registry) error {

	if err := validate.Required("zeroPriority", "body", m.ZeroPriority); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlockPriorityCounter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockPriorityCounter) UnmarshalBinary(b []byte) error {
	var res BlockPriorityCounter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
