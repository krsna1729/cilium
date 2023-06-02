// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CIDRList List of CIDRs
//
// swagger:model CIDRList
type CIDRList struct {

	// list
	List []string `json:"list"`

	// revision
	Revision int64 `json:"revision,omitempty"`
}

// Validate validates this c ID r list
func (m *CIDRList) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this c ID r list based on context it is used
func (m *CIDRList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CIDRList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CIDRList) UnmarshalBinary(b []byte) error {
	var res CIDRList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
