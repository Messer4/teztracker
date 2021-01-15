// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ThirtPartyBakers thirt party bakers
//
// swagger:model ThirtPartyBakers
type ThirtPartyBakers struct {

	// baker
	Baker string `json:"baker,omitempty"`

	// providers
	Providers []*ThirtPartyBakersProvidersItems0 `json:"providers"`
}

// Validate validates this thirt party bakers
func (m *ThirtPartyBakers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProviders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThirtPartyBakers) validateProviders(formats strfmt.Registry) error {

	if swag.IsZero(m.Providers) { // not required
		return nil
	}

	for i := 0; i < len(m.Providers); i++ {
		if swag.IsZero(m.Providers[i]) { // not required
			continue
		}

		if m.Providers[i] != nil {
			if err := m.Providers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("providers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThirtPartyBakers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThirtPartyBakers) UnmarshalBinary(b []byte) error {
	var res ThirtPartyBakers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ThirtPartyBakersProvidersItems0 thirt party bakers providers items0
//
// swagger:model ThirtPartyBakersProvidersItems0
type ThirtPartyBakersProvidersItems0 struct {

	// address
	Address string `json:"address,omitempty"`

	// available capacity
	AvailableCapacity int64 `json:"available_capacity,omitempty"`

	// efficiency
	Efficiency float64 `json:"efficiency,omitempty"`

	// fee
	Fee float64 `json:"fee,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// number
	Number int64 `json:"number,omitempty"`

	// payout accuracy
	PayoutAccuracy string `json:"payout_accuracy,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// staking balance
	StakingBalance int64 `json:"staking_balance,omitempty"`

	// yield
	Yield float64 `json:"yield,omitempty"`
}

// Validate validates this thirt party bakers providers items0
func (m *ThirtPartyBakersProvidersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ThirtPartyBakersProvidersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThirtPartyBakersProvidersItems0) UnmarshalBinary(b []byte) error {
	var res ThirtPartyBakersProvidersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}