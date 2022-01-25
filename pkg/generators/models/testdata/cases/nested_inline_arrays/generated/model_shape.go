// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Shape is an object.
type Shape struct {
	// Coordinates:
	Coordinates []Line `json:"coordinates,omitempty"`
	// Type:
	Type string `json:"type,omitempty"`
}

// Validate implements basic validation for this model
func (m Shape) Validate() error {
	return validation.Errors{
		"coordinates": validation.Validate(
			m.Coordinates,
		),
	}.Filter()
}
