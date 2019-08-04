// Code generated by go generate
// This file was generated by robots at 2019-08-04 01:30:01.210612088 +0000 UTC

package optional

import (
    "gopkg.in/yaml.v2"
	"encoding/json"
	"errors"
)

// Complex128 is an optional complex128.
type Complex128 struct {
	value *complex128
}

// NewComplex128 creates an optional.Complex128 from a complex128.
func NewComplex128(v complex128) Complex128 {
	return Complex128{&v}
}

// Set sets the complex128 value.
func (c *Complex128) Set(v complex128) {
	c.value = &v
}

// Get returns the complex128 value or an error if not present.
func (c Complex128) Get() (complex128, error) {
	if !c.Present() {
		var zero complex128
		return zero, errors.New("value not present")
	}
	return *c.value, nil
}

// Unwrap returns the complex128 value or an error if not present.
func (c Complex128) Unwrap() complex128 {
	return *c.value
}

// Present returns whether or not the value is present.
func (c Complex128) Present() bool {
	return c.value != nil
}

// OrElse returns the complex128 value or a default value if the value is not present.
func (c Complex128) OrElse(v complex128) complex128 {
	if c.Present() {
		return *c.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (c Complex128) If(fn func(complex128)) {
	if c.Present() {
		fn(*c.value)
	}
}

func (c Complex128) MarshalJSON() ([]byte, error) {
	if c.Present() {
		return json.Marshal(c.value)
	}
	return json.Marshal(nil)
}

func (c *Complex128) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		c.value = nil
		return nil
	}

	var value complex128

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	c.value = &value
	return nil
}
