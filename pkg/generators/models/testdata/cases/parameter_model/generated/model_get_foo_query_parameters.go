// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// GetFooQueryParameters is an object.
type GetFooQueryParameters struct {
	// Param1:
	Param1 string `json:"param1,omitempty"`
	// Id:
	Id string `json:"id"`
	// Param2:
	Param2 int32 `json:"param2,omitempty"`
	// Param3:
	Param3 []string `json:"param3,omitempty"`
	// Page: The current set of paged results to display, based on a 1-based array index
	Page int32 `json:"page,omitempty"`
}

// Validate implements basic validation for this model
func (m GetFooQueryParameters) Validate() error {
	return validation.Errors{
		"param1": validation.Validate(
			m.Param1, is.UUID,
		),
		"id": validation.Validate(
			m.Id, validation.Required, is.UUID,
		),
		"param2": validation.Validate(
			m.Param2, validation.Min(int32(0)), validation.Max(int32(10)),
		),
		"param3": validation.Validate(
			m.Param3,
		),
		"page": validation.Validate(
			m.Page, validation.Min(int32(1)),
		),
	}.Filter()
}
