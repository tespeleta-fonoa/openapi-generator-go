// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ColumnTypeMetadata is an object. Type metadata
type ColumnTypeMetadata struct {
	// Columns: List of columns if this type is structural
	Columns []ColumnMetadata `json:"columns,omitempty"`
	// ItemType: Type metadata
	ItemType *ColumnTypeMetadata `json:"itemType,omitempty"`
	// Nullable: Column nullability
	Nullable Nullability `json:"nullable"`
	// OriginalName: Original column type as given by data source
	OriginalName string `json:"originalName"`
	// Type: Normalized column type. If type cannot be determined or is not compatible, then 'other'.
	Type ColumnType `json:"type"`
}

// Validate implements basic validation for this model
func (m ColumnTypeMetadata) Validate() error {
	return validation.Errors{
		"columns": validation.Validate(
			m.Columns,
		),
		"itemType": validation.Validate(
			m.ItemType,
		),
		"nullable": validation.Validate(
			m.Nullable, validation.Required,
		),
		"type": validation.Validate(
			m.Type, validation.Required,
		),
	}.Filter()
}
