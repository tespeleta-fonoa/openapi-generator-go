// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Top is an object.
type Top struct {
	// Arr:
	Arr []Sub1 `json:"arr,omitempty"`
	// Boo: Type alias for a value type
	Boo bool `json:"boo,omitempty"`
	// Obj:
	Obj Sub1 `json:"obj,omitempty"`
}

// Validate implements basic validation for this model
func (m Top) Validate() error {
	return validation.Errors{
		"arr": validation.Validate(
			m.Arr,
		),
		"obj": validation.Validate(
			m.Obj,
		),
	}.Filter()
}

// GetArr returns the Arr property
func (m Top) GetArr() []Sub1 {
	return m.Arr
}

// SetArr sets the Arr property
func (m Top) SetArr(val []Sub1) {
	m.Arr = val
}

// GetBoo returns the Boo property
func (m Top) GetBoo() bool {
	return m.Boo
}

// SetBoo sets the Boo property
func (m Top) SetBoo(val bool) {
	m.Boo = val
}

// GetObj returns the Obj property
func (m Top) GetObj() Sub1 {
	return m.Obj
}

// SetObj sets the Obj property
func (m Top) SetObj(val Sub1) {
	m.Obj = val
}
