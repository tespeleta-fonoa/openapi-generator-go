// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Foo is an object.
type Foo struct {
	// Bar:
	Bar struct {
		Bar string `json:"bar,omitempty"`
		Foo string `json:"foo,omitempty"`
	} `json:"bar,omitempty"`
}

// Validate implements basic validation for this model
func (m Foo) Validate() error {
	return validation.Errors{}.Filter()
}
