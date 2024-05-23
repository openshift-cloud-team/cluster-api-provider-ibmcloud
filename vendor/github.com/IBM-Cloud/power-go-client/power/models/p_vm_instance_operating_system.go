// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PVMInstanceOperatingSystem p VM instance operating system
//
// swagger:model PVMInstanceOperatingSystem
type PVMInstanceOperatingSystem struct {

	// Type of the OS [aix, ibmi, rhel, sles, vtl, rhcos]
	// Required: true
	Type *string `json:"type"`

	// OS system information (usually version and build)
	Version string `json:"version,omitempty"`
}

// Validate validates this p VM instance operating system
func (m *PVMInstanceOperatingSystem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceOperatingSystem) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p VM instance operating system based on context it is used
func (m *PVMInstanceOperatingSystem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstanceOperatingSystem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstanceOperatingSystem) UnmarshalBinary(b []byte) error {
	var res PVMInstanceOperatingSystem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
