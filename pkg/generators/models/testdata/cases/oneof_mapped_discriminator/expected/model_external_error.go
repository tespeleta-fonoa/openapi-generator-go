// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test oneOf discriminator support
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ExternalError is an object. Represents an error from an external service. The system is stable, but can not complete the request. Try again later.
type ExternalError struct {
	// Kind:
	Kind string `json:"kind"`
	// Service: name of the service that is returning errors
	Service string `json:"service,omitempty"`
	// TraceId: the request tracing id, this can be submitted during bug reports to help with debugging the underlying cause.
	TraceId string `json:"traceId"`
}

// Validate implements basic validation for this model
func (m ExternalError) Validate() error {
	return validation.Errors{
		"service": validation.Validate(
			m.Service, validation.Length(1, 0),
		),
		"traceId": validation.Validate(
			m.TraceId, validation.Required, validation.Length(1, 0),
		),
	}.Filter()
}
