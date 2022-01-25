// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test oneOf discriminator support
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// GenericError is an object. Represents and unknown error. The system may be unstable, it is unknown if retrying the request will succeed.
type GenericError struct {
	// Code: machine friendly error code.
	Code string `json:"code,omitempty"`
	// Kind:
	Kind string `json:"kind"`
	// Message: the user friendly error message.
	Message string `json:"message"`
	// TraceId: the request tracing id, this can be submitted during bug reports to help with debugging the underlying cause.
	TraceId string `json:"traceId,omitempty"`
}

// Validate implements basic validation for this model
func (m GenericError) Validate() error {
	return validation.Errors{}.Filter()
}
