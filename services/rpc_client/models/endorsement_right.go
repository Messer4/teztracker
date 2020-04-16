// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EndorsementRight endorsement right
// swagger:model EndorsementRight
type EndorsementRight struct {

	// delegate
	Delegate string `json:"delegate,omitempty"`

	// level
	Level int64 `json:"level,omitempty"`

	// slots
	Slots []int64 `json:"slots"`
}

// Validate validates this endorsement right
func (m *EndorsementRight) Validate(formats strfmt.Registry) error {
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
