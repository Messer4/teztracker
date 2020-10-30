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

// AccountBakingRow account baking row
// swagger:model AccountBakingRow
type AccountBakingRow struct {

	// avg priority
	// Required: true
	AvgPriority *float32 `json:"avgPriority"`

	// blocks
	// Required: true
	Blocks *int64 `json:"blocks"`

	// cycle
	// Required: true
	Cycle *int64 `json:"cycle"`

	// cycle end
	CycleEnd int64 `json:"cycleEnd,omitempty"`

	// cycle start
	CycleStart int64 `json:"cycleStart,omitempty"`

	// missed
	// Required: true
	Missed *int64 `json:"missed"`

	// rewards
	// Required: true
	Rewards *int64 `json:"rewards"`

	// status
	Status string `json:"status,omitempty"`

	// stolen
	// Required: true
	Stolen *int64 `json:"stolen"`

	// total deposit
	// Required: true
	TotalDeposit *int64 `json:"totalDeposit"`
}

// Validate validates this account baking row
func (m *AccountBakingRow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvgPriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlocks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMissed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStolen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalDeposit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountBakingRow) validateAvgPriority(formats strfmt.Registry) error {

	if err := validate.Required("avgPriority", "body", m.AvgPriority); err != nil {
		return err
	}

	return nil
}

func (m *AccountBakingRow) validateBlocks(formats strfmt.Registry) error {

	if err := validate.Required("blocks", "body", m.Blocks); err != nil {
		return err
	}

	return nil
}

func (m *AccountBakingRow) validateCycle(formats strfmt.Registry) error {

	if err := validate.Required("cycle", "body", m.Cycle); err != nil {
		return err
	}

	return nil
}

func (m *AccountBakingRow) validateMissed(formats strfmt.Registry) error {

	if err := validate.Required("missed", "body", m.Missed); err != nil {
		return err
	}

	return nil
}

func (m *AccountBakingRow) validateRewards(formats strfmt.Registry) error {

	if err := validate.Required("rewards", "body", m.Rewards); err != nil {
		return err
	}

	return nil
}

func (m *AccountBakingRow) validateStolen(formats strfmt.Registry) error {

	if err := validate.Required("stolen", "body", m.Stolen); err != nil {
		return err
	}

	return nil
}

func (m *AccountBakingRow) validateTotalDeposit(formats strfmt.Registry) error {

	if err := validate.Required("totalDeposit", "body", m.TotalDeposit); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountBakingRow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountBakingRow) UnmarshalBinary(b []byte) error {
	var res AccountBakingRow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
