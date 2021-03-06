// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PublicBakerSearch public baker search
// swagger:model PublicBakerSearch
type PublicBakerSearch struct {

	// account Id
	AccountID string `json:"accountId,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this public baker search
func (m *PublicBakerSearch) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PublicBakerSearch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicBakerSearch) UnmarshalBinary(b []byte) error {
	var res PublicBakerSearch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
