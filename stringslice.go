// Code generated by go generate
// This file was generated by robots at 2019-08-04 01:30:07.645544933 +0000 UTC

package optional

import (
    "gopkg.in/yaml.v2"
	"encoding/json"
	"errors"
)

// StringSlice is an optional []string.
type StringSlice struct {
	value *[]string
}

// NewStringSlice creates an optional.StringSlice from a []string.
func NewStringSlice(v []string) StringSlice {
	return StringSlice{&v}
}

// Set sets the []string value.
func (s *StringSlice) Set(v []string) {
	s.value = &v
}

// Get returns the []string value or an error if not present.
func (s StringSlice) Get() ([]string, error) {
	if !s.Present() {
		var zero []string
		return zero, errors.New("value not present")
	}
	return *s.value, nil
}

// Unwrap returns the []string value or an error if not present.
func (s StringSlice) Unwrap() []string {
	return *s.value
}

// Present returns whether or not the value is present.
func (s StringSlice) Present() bool {
	return s.value != nil
}

// OrElse returns the []string value or a default value if the value is not present.
func (s StringSlice) OrElse(v []string) []string {
	if s.Present() {
		return *s.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (s StringSlice) If(fn func([]string)) {
	if s.Present() {
		fn(*s.value)
	}
}

func (s StringSlice) MarshalJSON() ([]byte, error) {
	if s.Present() {
		return json.Marshal(s.value)
	}
	return json.Marshal(nil)
}

func (s *StringSlice) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		s.value = nil
		return nil
	}

	var value []string

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	s.value = &value
	return nil
}
