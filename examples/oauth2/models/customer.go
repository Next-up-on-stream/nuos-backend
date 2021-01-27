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

// Customer customer
//
// swagger:model customer
type Customer struct {

	// agent associated with this customer
	AgentID int32 `json:"agentId,omitempty"`

	// internal identifier of a customer
	// Required: true
	// Read Only: true
	CustomerID int64 `json:"customerId"`

	// fips code
	// Required: true
	// Min Length: 1
	FipsCode *string `json:"fipsCode"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// Lookup identifier to find a customer in the system
	// Required: true
	// Min Length: 11
	Ssn *string `json:"ssn"`

	// surname
	// Required: true
	// Min Length: 1
	Surname *string `json:"surname"`
}

// Validate validates this customer
func (m *Customer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFipsCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSurname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Customer) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customerId", "body", int64(m.CustomerID)); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateFipsCode(formats strfmt.Registry) error {

	if err := validate.Required("fipsCode", "body", m.FipsCode); err != nil {
		return err
	}

	if err := validate.MinLength("fipsCode", "body", *m.FipsCode, 1); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateSsn(formats strfmt.Registry) error {

	if err := validate.Required("ssn", "body", m.Ssn); err != nil {
		return err
	}

	if err := validate.MinLength("ssn", "body", *m.Ssn, 11); err != nil {
		return err
	}

	return nil
}

func (m *Customer) validateSurname(formats strfmt.Registry) error {

	if err := validate.Required("surname", "body", m.Surname); err != nil {
		return err
	}

	if err := validate.MinLength("surname", "body", *m.Surname, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this customer based on the context it is used
func (m *Customer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomerID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Customer) contextValidateCustomerID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "customerId", "body", int64(m.CustomerID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Customer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Customer) UnmarshalBinary(b []byte) error {
	var res Customer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
