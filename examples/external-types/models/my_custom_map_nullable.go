// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	alternate "github.com/go-swagger/go-swagger/examples/external-types/fred"
)

// MyCustomMapNullable This generate a map type in models, based on the external type.
// Notice the impact of the x-nullable directive, to produce a map of pointers.
//
// MapNullable map[string]map[string]*alternate.MyAlternateType
//
//
// swagger:model MyCustomMapNullable
type MyCustomMapNullable map[string]map[string]*alternate.MyAlternateType

// Validate validates this my custom map nullable
func (m MyCustomMapNullable) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		for kk := range m[k] {

			if swag.IsZero(m[k][kk]) { // not required
				continue
			}
			if val, ok := m[k][kk]; ok {
				if err := val.Validate(formats); err != nil {
					return err
				}
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this my custom map nullable based on context it is used
func (m MyCustomMapNullable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
