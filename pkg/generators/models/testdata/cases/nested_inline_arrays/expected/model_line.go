// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Line is an object.
type Line struct {
	// Coordinates:
	Coordinates [][]float32 `json:"coordinates,omitempty"`
	// Type:
	Type string `json:"type,omitempty"`
}

// Validate implements basic validation for this model
func (m Line) Validate() error {
	return validation.Errors{
		"coordinates": validation.Validate(
			m.Coordinates,
		),
	}.Filter()
}
