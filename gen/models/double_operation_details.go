// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DoubleOperationDetails double operation details
// swagger:model DoubleOperationDetails
type DoubleOperationDetails struct {

	// baker reward
	BakerReward int64 `json:"baker_reward,omitempty"`

	// denounced level
	DenouncedLevel int64 `json:"denounced_level,omitempty"`

	// evidence baker
	EvidenceBaker string `json:"evidence_baker,omitempty"`

	// evidence baker name
	EvidenceBakerName string `json:"evidence_baker_name,omitempty"`

	// lost deposits
	LostDeposits int64 `json:"lost_deposits,omitempty"`

	// lost fees
	LostFees int64 `json:"lost_fees,omitempty"`

	// lost rewards
	LostRewards int64 `json:"lost_rewards,omitempty"`

	// offender
	Offender string `json:"offender,omitempty"`

	// offender name
	OffenderName string `json:"offender_name,omitempty"`

	// priority
	Priority int64 `json:"priority,omitempty"`
}

// Validate validates this double operation details
func (m *DoubleOperationDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DoubleOperationDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DoubleOperationDetails) UnmarshalBinary(b []byte) error {
	var res DoubleOperationDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
