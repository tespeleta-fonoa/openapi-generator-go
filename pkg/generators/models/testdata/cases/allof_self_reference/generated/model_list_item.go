// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ListItem is an object.
type ListItem struct {
	// Next: the next item
	Next *ListItem `json:"next,omitempty"`
	// Value:
	Value string `json:"value,omitempty"`
}

// Validate implements basic validation for this model
func (m ListItem) Validate() error {
	return validation.Errors{
		"next": validation.Validate(
			m.Next,
		),
	}.Filter()
}
