// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Foo is an object. this is a foo object
type Foo struct {
	// Mixin:
	Mixin string `json:"mixin,omitempty"`
	// SubField:
	SubField string `json:"subField,omitempty"`
}

// Validate implements basic validation for this model
func (m Foo) Validate() error {
	return validation.Errors{}.Filter()
}
