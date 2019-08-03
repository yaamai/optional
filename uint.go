// Code generated by go generate
// This file was generated by robots at 2019-08-03 10:39:57.625698558 +0000 UTC

package optional

import (
    "gopkg.in/yaml.v2"
	"encoding/json"
	"errors"
)

// Uint is an optional uint.
type Uint struct {
	value *uint
}

// NewUint creates an optional.Uint from a uint.
func NewUint(v uint) Uint {
	return Uint{&v}
}

// Set sets the uint value.
func (u *Uint) Set(v uint) {
	u.value = &v
}

// Get returns the uint value or an error if not present.
func (u Uint) Get() (uint, error) {
	if !u.Present() {
		var zero uint
		return zero, errors.New("value not present")
	}
	return *u.value, nil
}

// Unwrap returns the uint value or an error if not present.
func (u Uint) Unwrap() uint {
	return *u.value
}

// Present returns whether or not the value is present.
func (u Uint) Present() bool {
	return u.value != nil
}

// OrElse returns the uint value or a default value if the value is not present.
func (u Uint) OrElse(v uint) uint {
	if u.Present() {
		return *u.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (u Uint) If(fn func(uint)) {
	if u.Present() {
		fn(*u.value)
	}
}

func (u Uint) MarshalJSON() ([]byte, error) {
	if u.Present() {
		return json.Marshal(u.value)
	}
	return json.Marshal(nil)
}

func (u *Uint) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		u.value = nil
		return nil
	}

	var value uint

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	u.value = &value
	return nil
}

func (u Uint) MarshalYAML() ([]byte, error) {
	if u.Present() {
		return yaml.Marshal(u.value)
	}
	return yaml.Marshal(nil)
}

func (u *Uint) UnmarshalYAML(data []byte) error {

	if string(data) == "null" {
		u.value = nil
		return nil
	}

	var value uint

	if err := yaml.Unmarshal(data, &value); err != nil {
		return err
	}

	u.value = &value
	return nil
}
