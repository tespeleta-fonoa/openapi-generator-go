// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Top is an object.
type Top struct {
	// Arr:
	Arr []Sub1 `json:"arr,omitempty"`
	// Boo: Type alias for a value type
	Boo Sub3 `json:"boo,omitempty"`
	// Obj:
	Obj Sub1 `json:"obj,omitempty"`
}

// Validate implements basic validation for this model
func (m Top) Validate() error {
	return validation.Errors{
		"arr": validation.Validate(
			m.Arr,
		),
		"boo": validation.Validate(
			m.Boo,
		),
		"obj": validation.Validate(
			m.Obj,
		),
	}.Filter()
}
