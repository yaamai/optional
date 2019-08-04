// Code generated by go generate
// This file was generated by robots at 2019-08-04 02:30:27.408455046 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Int is an optional int.
type Int struct {
	value *int
}

// NewInt creates an optional.Int from a int.
func NewInt(v int) Int {
	return Int{&v}
}

// Set sets the int value.
func (i *Int) Set(v int) {
	i.value = &v
}

// Set sets the int value.
func (i *Int) SetOptional(vOpt Int) {
	if vOpt.Present() {
        v := vOpt.Unwrap()
	    i.value = &v
	}
}

// Get returns the int value or an error if not present.
func (i Int) Get() (int, error) {
	if !i.Present() {
		var zero int
		return zero, errors.New("value not present")
	}
	return *i.value, nil
}

// Unwrap returns the int value or an error if not present.
func (i Int) Unwrap() int {
	return *i.value
}

// Present returns whether or not the value is present.
func (i Int) Present() bool {
	return i.value != nil
}

// OrElse returns the int value or a default value if the value is not present.
func (i Int) OrElse(v int) int {
	if i.Present() {
		return *i.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (i Int) If(fn func(int)) {
	if i.Present() {
		fn(*i.value)
	}
}

func (i Int) MarshalJSON() ([]byte, error) {
	if i.Present() {
		return json.Marshal(i.value)
	}
	return json.Marshal(nil)
}

func (i *Int) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		i.value = nil
		return nil
	}

	var value int

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	i.value = &value
	return nil
}
