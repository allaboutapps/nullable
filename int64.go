// Code generated by go generate; DO NOT EDIT.
// This file was generated by scripts/models/gen_nullable_types.go
package nullable

import (
	"bytes"
	"encoding/json"
	"github.com/go-openapi/strfmt"
)

// Int64 represents a int64 that may be null or not
// present in json at all.
type Int64 struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null and valid int64
	Value   int64
}

// Returns nil if not present or valid. Otherwise it will
// return a pointer to the value.
func (i *Int64) Ptr() *int64 {
	if i.Present && i.Valid {
		return &i.Value
	}

	return nil
}

// UnmarshalJSON implements json.Marshaler interface.
func (i *Int64) UnmarshalJSON(data []byte) error {
	i.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &i.Value); err != nil {
		return err
	}

	i.Valid = true
	return nil
}

// Validate implements runtime.Validateable interface for go-swagger generation.
func (i *Int64) Validate(formats strfmt.Registry) error {
	return nil
}

// Int64Slice represents a []int64 that may be null or not
// present in json at all.
type Int64Slice struct {
	Present bool // Present is true if key is present in json
	Valid   bool // Valid is true if value is not null
	Value   []int64
}

// UnmarshalJSON implements json.Marshaler interface.
func (i *Int64Slice) UnmarshalJSON(data []byte) error {
	i.Present = true

	if bytes.Equal(data, null) {
		return nil
	}

	if err := json.Unmarshal(data, &i.Value); err != nil {
		return err
	}

	i.Valid = true
	return nil
}
