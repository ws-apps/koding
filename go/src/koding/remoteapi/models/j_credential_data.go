package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// JCredentialData j credential data
// swagger:model JCredentialData
type JCredentialData struct {

	// id
	ID string `json:"_id,omitempty"`

	// identifier
	Identifier string `json:"identifier,omitempty"`

	// meta
	// Required: true
	Meta interface{} `json:"meta"`

	// origin Id
	// Required: true
	OriginID *string `json:"originId"`
}

// Validate validates this j credential data
func (m *JCredentialData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMeta(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOriginID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JCredentialData) validateMeta(formats strfmt.Registry) error {

	return nil
}

func (m *JCredentialData) validateOriginID(formats strfmt.Registry) error {

	if err := validate.Required("originId", "body", m.OriginID); err != nil {
		return err
	}

	return nil
}
