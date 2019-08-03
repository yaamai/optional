// Code generated by go generate
// This file was generated by robots at 2019-08-03 10:39:52.12453387 +0000 UTC

package optional

import (
    "gopkg.in/yaml.v2"
	"encoding/json"
	"errors"
)

// Byte is an optional byte.
type Byte struct {
	value *byte
}

// NewByte creates an optional.Byte from a byte.
func NewByte(v byte) Byte {
	return Byte{&v}
}

// Set sets the byte value.
func (b *Byte) Set(v byte) {
	b.value = &v
}

// Get returns the byte value or an error if not present.
func (b Byte) Get() (byte, error) {
	if !b.Present() {
		var zero byte
		return zero, errors.New("value not present")
	}
	return *b.value, nil
}

// Unwrap returns the byte value or an error if not present.
func (b Byte) Unwrap() byte {
	return *b.value
}

// Present returns whether or not the value is present.
func (b Byte) Present() bool {
	return b.value != nil
}

// OrElse returns the byte value or a default value if the value is not present.
func (b Byte) OrElse(v byte) byte {
	if b.Present() {
		return *b.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (b Byte) If(fn func(byte)) {
	if b.Present() {
		fn(*b.value)
	}
}

func (b Byte) MarshalJSON() ([]byte, error) {
	if b.Present() {
		return json.Marshal(b.value)
	}
	return json.Marshal(nil)
}

func (b *Byte) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		b.value = nil
		return nil
	}

	var value byte

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	b.value = &value
	return nil
}

func (b Byte) MarshalYAML() ([]byte, error) {
	if b.Present() {
		return yaml.Marshal(b.value)
	}
	return yaml.Marshal(nil)
}

func (b *Byte) UnmarshalYAML(data []byte) error {

	if string(data) == "null" {
		b.value = nil
		return nil
	}

	var value byte

	if err := yaml.Unmarshal(data, &value); err != nil {
		return err
	}

	b.value = &value
	return nil
}
