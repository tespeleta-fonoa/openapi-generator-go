// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Person is an object.
type Person struct {
	// Age:
	Age int32 `json:"age,omitempty"`
	// Name:
	Name string `json:"name,omitempty"`
}

// Validate implements basic validation for this model
func (m Person) Validate() error {
	return validation.Errors{}.Filter()
}
