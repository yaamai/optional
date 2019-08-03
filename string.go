// Code generated by go generate
// This file was generated by robots at 2019-08-03 10:39:57.248406214 +0000 UTC

package optional

import (
    "gopkg.in/yaml.v2"
	"encoding/json"
	"errors"
)

// String is an optional string.
type String struct {
	value *string
}

// NewString creates an optional.String from a string.
func NewString(v string) String {
	return String{&v}
}

// Set sets the string value.
func (s *String) Set(v string) {
	s.value = &v
}

// Get returns the string value or an error if not present.
func (s String) Get() (string, error) {
	if !s.Present() {
		var zero string
		return zero, errors.New("value not present")
	}
	return *s.value, nil
}

// Unwrap returns the string value or an error if not present.
func (s String) Unwrap() string {
	return *s.value
}

// Present returns whether or not the value is present.
func (s String) Present() bool {
	return s.value != nil
}

// OrElse returns the string value or a default value if the value is not present.
func (s String) OrElse(v string) string {
	if s.Present() {
		return *s.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (s String) If(fn func(string)) {
	if s.Present() {
		fn(*s.value)
	}
}

func (s String) MarshalJSON() ([]byte, error) {
	if s.Present() {
		return json.Marshal(s.value)
	}
	return json.Marshal(nil)
}

func (s *String) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		s.value = nil
		return nil
	}

	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	s.value = &value
	return nil
}

func (s String) MarshalYAML() ([]byte, error) {
	if s.Present() {
		return yaml.Marshal(s.value)
	}
	return yaml.Marshal(nil)
}

func (s *String) UnmarshalYAML(data []byte) error {

	if string(data) == "null" {
		s.value = nil
		return nil
	}

	var value string

	if err := yaml.Unmarshal(data, &value); err != nil {
		return err
	}

	s.value = &value
	return nil
}
