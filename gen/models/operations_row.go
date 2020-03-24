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

// OperationsRow operations row
// swagger:model OperationsRow
type OperationsRow struct {

	// amount
	Amount int64 `json:"amount,omitempty"`

	// balance
	Balance int64 `json:"balance,omitempty"`

	// ballot
	Ballot string `json:"ballot,omitempty"`

	// block hash
	// Required: true
	BlockHash *string `json:"blockHash"`

	// block level
	// Required: true
	BlockLevel *int64 `json:"blockLevel"`

	// consumed gas
	ConsumedGas int64 `json:"consumedGas,omitempty"`

	// counter
	Counter int64 `json:"counter,omitempty"`

	// delegatable
	Delegatable bool `json:"delegatable,omitempty"`

	// delegate
	Delegate string `json:"delegate,omitempty"`

	// delegate name
	DelegateName string `json:"delegateName,omitempty"`

	// destination
	Destination string `json:"destination,omitempty"`

	// destination name
	DestinationName string `json:"destinationName,omitempty"`

	// double bake
	DoubleBake *DoubleBakingDetails `json:"doubleBake,omitempty"`

	// fee
	Fee int64 `json:"fee,omitempty"`

	// gas limit
	GasLimit int64 `json:"gasLimit,omitempty"`

	// kind
	// Required: true
	Kind *string `json:"kind"`

	// level
	Level int64 `json:"level,omitempty"`

	// manager pubkey
	ManagerPubkey string `json:"managerPubkey,omitempty"`

	// nonce
	Nonce string `json:"nonce,omitempty"`

	// operation group hash
	// Required: true
	OperationGroupHash *string `json:"operationGroupHash"`

	// operation Id
	// Required: true
	OperationID *int64 `json:"operationId"`

	// originated contracts
	OriginatedContracts string `json:"originatedContracts,omitempty"`

	// paid storage size diff
	PaidStorageSizeDiff int64 `json:"paidStorageSizeDiff,omitempty"`

	// parameters
	Parameters string `json:"parameters,omitempty"`

	// pkh
	Pkh string `json:"pkh,omitempty"`

	// proposal
	Proposal string `json:"proposal,omitempty"`

	// public key
	PublicKey string `json:"publicKey,omitempty"`

	// script
	Script string `json:"script,omitempty"`

	// secret
	Secret string `json:"secret,omitempty"`

	// slots
	Slots string `json:"slots,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// source name
	SourceName string `json:"sourceName,omitempty"`

	// spendable
	Spendable bool `json:"spendable,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// storage
	Storage string `json:"storage,omitempty"`

	// storage limit
	StorageLimit int64 `json:"storageLimit,omitempty"`

	// storage size
	StorageSize int64 `json:"storageSize,omitempty"`

	// timestamp
	// Required: true
	Timestamp *int64 `json:"timestamp"`
}

// Validate validates this operations row
func (m *OperationsRow) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlockLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDoubleBake(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationGroupHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperationsRow) validateBlockHash(formats strfmt.Registry) error {

	if err := validate.Required("blockHash", "body", m.BlockHash); err != nil {
		return err
	}

	return nil
}

func (m *OperationsRow) validateBlockLevel(formats strfmt.Registry) error {

	if err := validate.Required("blockLevel", "body", m.BlockLevel); err != nil {
		return err
	}

	return nil
}

func (m *OperationsRow) validateDoubleBake(formats strfmt.Registry) error {

	if swag.IsZero(m.DoubleBake) { // not required
		return nil
	}

	if m.DoubleBake != nil {
		if err := m.DoubleBake.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("doubleBake")
			}
			return err
		}
	}

	return nil
}

func (m *OperationsRow) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *OperationsRow) validateOperationGroupHash(formats strfmt.Registry) error {

	if err := validate.Required("operationGroupHash", "body", m.OperationGroupHash); err != nil {
		return err
	}

	return nil
}

func (m *OperationsRow) validateOperationID(formats strfmt.Registry) error {

	if err := validate.Required("operationId", "body", m.OperationID); err != nil {
		return err
	}

	return nil
}

func (m *OperationsRow) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OperationsRow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationsRow) UnmarshalBinary(b []byte) error {
	var res OperationsRow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
