// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ColumnMetadata is an object. Metadata for single column
type ColumnMetadata struct {
	// Comment: Column description
	Comment string `json:"comment,omitempty"`
	// Name: Column name
	Name string `json:"name"`
	// Type: Type metadata
	Type ColumnTypeMetadata `json:"type"`
}

// Validate implements basic validation for this model
func (m ColumnMetadata) Validate() error {
	return validation.Errors{
		"type": validation.Validate(
			m.Type, validation.NotNil,
		),
	}.Filter()
}
