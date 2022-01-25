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

// UserEntity is an object. This is a short description of a user entity used for permission listing or assignments in other services.
type UserEntity struct {
	// Email:
	Email string `json:"email"`
	// Id:
	Id string `json:"id"`
	// Name:
	Name string `json:"name"`
}

// Validate implements basic validation for this model
func (m UserEntity) Validate() error {
	return validation.Errors{
		"email": validation.Validate(
			m.Email, validation.Required, is.EmailFormat,
		),
		"id": validation.Validate(
			m.Id, validation.Required, is.UUID,
		),
	}.Filter()
}
