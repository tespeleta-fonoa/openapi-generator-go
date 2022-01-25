// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test oneOf discriminator support
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ValidationError is an object. a validation error for a form field
type ValidationError struct {
	// Message: the user friendly validation error message
	Message string `json:"message,omitempty"`
	// Name: the field name
	Name string `json:"name,omitempty"`
}

// Validate implements basic validation for this model
func (m ValidationError) Validate() error {
	return validation.Errors{}.Filter()
}
