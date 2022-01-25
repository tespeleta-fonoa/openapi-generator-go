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

// BaseEntity is an object. Contains shared properties for all the entities
type BaseEntity struct {
	// Id:
	Id string `json:"id"`
	// Name:
	Name string `json:"name"`
}

// Validate implements basic validation for this model
func (m BaseEntity) Validate() error {
	return validation.Errors{
		"id": validation.Validate(
			m.Id, validation.Required, is.UUID,
		),
	}.Filter()
}
