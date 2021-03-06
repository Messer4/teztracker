// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BakingRightsPerBlock baking rights per block
// swagger:model BakingRightsPerBlock
type BakingRightsPerBlock struct {

	// baker
	Baker string `json:"baker,omitempty"`

	// baker priority
	// Required: true
	BakerPriority *int64 `json:"baker_priority"`

	// block hash
	BlockHash string `json:"block_hash,omitempty"`

	// level
	Level int64 `json:"level,omitempty"`

	// rights
	Rights []*BakingRightsRow `json:"rights"`
}

// Validate validates this baking rights per block
func (m *BakingRightsPerBlock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBakerPriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BakingRightsPerBlock) validateBakerPriority(formats strfmt.Registry) error {

	if err := validate.Required("baker_priority", "body", m.BakerPriority); err != nil {
		return err
	}

	return nil
}

func (m *BakingRightsPerBlock) validateRights(formats strfmt.Registry) error {

	if swag.IsZero(m.Rights) { // not required
		return nil
	}

	for i := 0; i < len(m.Rights); i++ {
		if swag.IsZero(m.Rights[i]) { // not required
			continue
		}

		if m.Rights[i] != nil {
			if err := m.Rights[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rights" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BakingRightsPerBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BakingRightsPerBlock) UnmarshalBinary(b []byte) error {
	var res BakingRightsPerBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
