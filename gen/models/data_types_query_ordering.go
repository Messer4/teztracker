// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataTypesQueryOrdering data types query ordering
// swagger:model DataTypes.QueryOrdering
type DataTypesQueryOrdering struct {

	// direction
	// Required: true
	// Enum: [asc desc]
	Direction *string `json:"direction"`

	// field
	// Required: true
	Field *string `json:"field"`
}

// Validate validates this data types query ordering
func (m *DataTypesQueryOrdering) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateField(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dataTypesQueryOrderingTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["asc","desc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataTypesQueryOrderingTypeDirectionPropEnum = append(dataTypesQueryOrderingTypeDirectionPropEnum, v)
	}
}

const (

	// DataTypesQueryOrderingDirectionAsc captures enum value "asc"
	DataTypesQueryOrderingDirectionAsc string = "asc"

	// DataTypesQueryOrderingDirectionDesc captures enum value "desc"
	DataTypesQueryOrderingDirectionDesc string = "desc"
)

// prop value enum
func (m *DataTypesQueryOrdering) validateDirectionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dataTypesQueryOrderingTypeDirectionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DataTypesQueryOrdering) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", *m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *DataTypesQueryOrdering) validateField(formats strfmt.Registry) error {

	if err := validate.Required("field", "body", m.Field); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataTypesQueryOrdering) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataTypesQueryOrdering) UnmarshalBinary(b []byte) error {
	var res DataTypesQueryOrdering
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
