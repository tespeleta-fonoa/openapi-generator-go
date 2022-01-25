// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Artist is an object.
type Artist struct {
	// LeftHand:
	LeftHand *AnyThing `json:"leftHand,omitempty"`
	// RightHand:
	RightHand *Color `json:"rightHand,omitempty"`
}

// Validate implements basic validation for this model
func (m Artist) Validate() error {
	return validation.Errors{
		"leftHand": validation.Validate(
			m.LeftHand,
		),
		"rightHand": validation.Validate(
			m.RightHand,
		),
	}.Filter()
}
