// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Address is an object.
type Address struct {
	// Name:
	Name *string `json:"name"`
	// Number:
	Number int32 `json:"number"`
	// Street:
	Street string `json:"street"`
}

// Validate implements basic validation for this model
func (m Address) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.Length(2, 0),
		),
		"street": validation.Validate(
			m.Street, validation.Required, validation.Length(2, 0),
		),
	}.Filter()
}
