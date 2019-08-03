// Code generated by go generate
// This file was generated by robots at 2019-08-03 10:39:55.644831333 +0000 UTC

package optional

import (
    "gopkg.in/yaml.v2"
	"encoding/json"
	"errors"
)

// Int32 is an optional int32.
type Int32 struct {
	value *int32
}

// NewInt32 creates an optional.Int32 from a int32.
func NewInt32(v int32) Int32 {
	return Int32{&v}
}

// Set sets the int32 value.
func (i *Int32) Set(v int32) {
	i.value = &v
}

// Get returns the int32 value or an error if not present.
func (i Int32) Get() (int32, error) {
	if !i.Present() {
		var zero int32
		return zero, errors.New("value not present")
	}
	return *i.value, nil
}

// Unwrap returns the int32 value or an error if not present.
func (i Int32) Unwrap() int32 {
	return *i.value
}

// Present returns whether or not the value is present.
func (i Int32) Present() bool {
	return i.value != nil
}

// OrElse returns the int32 value or a default value if the value is not present.
func (i Int32) OrElse(v int32) int32 {
	if i.Present() {
		return *i.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (i Int32) If(fn func(int32)) {
	if i.Present() {
		fn(*i.value)
	}
}

func (i Int32) MarshalJSON() ([]byte, error) {
	if i.Present() {
		return json.Marshal(i.value)
	}
	return json.Marshal(nil)
}

func (i *Int32) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		i.value = nil
		return nil
	}

	var value int32

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.value = &value
	return nil
}

func (i Int32) MarshalYAML() ([]byte, error) {
	if i.Present() {
		return yaml.Marshal(i.value)
	}
	return yaml.Marshal(nil)
}

func (i *Int32) UnmarshalYAML(data []byte) error {

	if string(data) == "null" {
		i.value = nil
		return nil
	}

	var value int32

	if err := yaml.Unmarshal(data, &value); err != nil {
		return err
	}

	i.value = &value
	return nil
}
