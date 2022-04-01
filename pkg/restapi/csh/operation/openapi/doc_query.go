// Code generated by go-swagger; DO NOT EDIT.

// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0
//

package openapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DocQuery doc query
//
// swagger:model DocQuery
type DocQuery struct {
	idField string

	// doc ID
	// Required: true
	DocID *string `json:"docID"`

	// path
	Path string `json:"path,omitempty"`

	// upstream auth
	// Required: true
	UpstreamAuth *DocQueryAO1UpstreamAuth `json:"upstreamAuth"`

	// vault ID
	// Required: true
	VaultID *string `json:"vaultID"`
}

// ID gets the id of this subtype
func (m *DocQuery) ID() string {
	return m.idField
}

// SetID sets the id of this subtype
func (m *DocQuery) SetID(val string) {
	m.idField = val
}

// Type gets the type of this subtype
func (m *DocQuery) Type() string {
	return "DocQuery"
}

// SetType sets the type of this subtype
func (m *DocQuery) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *DocQuery) UnmarshalJSON(raw []byte) error {
	var data struct {

		// doc ID
		// Required: true
		DocID *string `json:"docID"`

		// path
		Path string `json:"path,omitempty"`

		// upstream auth
		// Required: true
		UpstreamAuth *DocQueryAO1UpstreamAuth `json:"upstreamAuth"`

		// vault ID
		// Required: true
		VaultID *string `json:"vaultID"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		ID string `json:"id,omitempty"`

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result DocQuery

	result.idField = base.ID

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.DocID = data.DocID
	result.Path = data.Path
	result.UpstreamAuth = data.UpstreamAuth
	result.VaultID = data.VaultID

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m DocQuery) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// doc ID
		// Required: true
		DocID *string `json:"docID"`

		// path
		Path string `json:"path,omitempty"`

		// upstream auth
		// Required: true
		UpstreamAuth *DocQueryAO1UpstreamAuth `json:"upstreamAuth"`

		// vault ID
		// Required: true
		VaultID *string `json:"vaultID"`
	}{

		DocID: m.DocID,

		Path: m.Path,

		UpstreamAuth: m.UpstreamAuth,

		VaultID: m.VaultID,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ID string `json:"id,omitempty"`

		Type string `json:"type"`
	}{

		ID: m.ID(),

		Type: m.Type(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this doc query
func (m *DocQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpstreamAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaultID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocQuery) validateDocID(formats strfmt.Registry) error {

	if err := validate.Required("docID", "body", m.DocID); err != nil {
		return err
	}

	return nil
}

func (m *DocQuery) validateUpstreamAuth(formats strfmt.Registry) error {

	if err := validate.Required("upstreamAuth", "body", m.UpstreamAuth); err != nil {
		return err
	}

	if m.UpstreamAuth != nil {
		if err := m.UpstreamAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upstreamAuth")
			}
			return err
		}
	}

	return nil
}

func (m *DocQuery) validateVaultID(formats strfmt.Registry) error {

	if err := validate.Required("vaultID", "body", m.VaultID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this doc query based on the context it is used
func (m *DocQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUpstreamAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocQuery) contextValidateUpstreamAuth(ctx context.Context, formats strfmt.Registry) error {

	if m.UpstreamAuth != nil {
		if err := m.UpstreamAuth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upstreamAuth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DocQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocQuery) UnmarshalBinary(b []byte) error {
	var res DocQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DocQueryAO1UpstreamAuth doc query a o1 upstream auth
//
// swagger:model DocQueryAO1UpstreamAuth
type DocQueryAO1UpstreamAuth struct {

	// edv
	// Required: true
	Edv *UpstreamAuthorization `json:"edv"`

	// kms
	// Required: true
	Kms *UpstreamAuthorization `json:"kms"`
}

// Validate validates this doc query a o1 upstream auth
func (m *DocQueryAO1UpstreamAuth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocQueryAO1UpstreamAuth) validateEdv(formats strfmt.Registry) error {

	if err := validate.Required("upstreamAuth"+"."+"edv", "body", m.Edv); err != nil {
		return err
	}

	if m.Edv != nil {
		if err := m.Edv.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upstreamAuth" + "." + "edv")
			}
			return err
		}
	}

	return nil
}

func (m *DocQueryAO1UpstreamAuth) validateKms(formats strfmt.Registry) error {

	if err := validate.Required("upstreamAuth"+"."+"kms", "body", m.Kms); err != nil {
		return err
	}

	if m.Kms != nil {
		if err := m.Kms.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upstreamAuth" + "." + "kms")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this doc query a o1 upstream auth based on the context it is used
func (m *DocQueryAO1UpstreamAuth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEdv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DocQueryAO1UpstreamAuth) contextValidateEdv(ctx context.Context, formats strfmt.Registry) error {

	if m.Edv != nil {
		if err := m.Edv.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upstreamAuth" + "." + "edv")
			}
			return err
		}
	}

	return nil
}

func (m *DocQueryAO1UpstreamAuth) contextValidateKms(ctx context.Context, formats strfmt.Registry) error {

	if m.Kms != nil {
		if err := m.Kms.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upstreamAuth" + "." + "kms")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DocQueryAO1UpstreamAuth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocQueryAO1UpstreamAuth) UnmarshalBinary(b []byte) error {
	var res DocQueryAO1UpstreamAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
