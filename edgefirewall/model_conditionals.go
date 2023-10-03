/*
Edge Firewall API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgefirewall

import (
	"encoding/json"
	"fmt"
)

// Conditionals the model 'Conditionals'
type Conditionals string

// List of Conditionals
const (
	IF Conditionals = "if"
	AND Conditionals = "and"
	OR Conditionals = "or"
)

// All allowed values of Conditionals enum
var AllowedConditionalsEnumValues = []Conditionals{
	"if",
	"and",
	"or",
}

func (v *Conditionals) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Conditionals(value)
	for _, existing := range AllowedConditionalsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Conditionals", value)
}

// NewConditionalsFromValue returns a pointer to a valid Conditionals
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConditionalsFromValue(v string) (*Conditionals, error) {
	ev := Conditionals(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Conditionals: valid values are %v", v, AllowedConditionalsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Conditionals) IsValid() bool {
	for _, existing := range AllowedConditionalsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Conditionals value
func (v Conditionals) Ptr() *Conditionals {
	return &v
}

type NullableConditionals struct {
	value *Conditionals
	isSet bool
}

func (v NullableConditionals) Get() *Conditionals {
	return v.value
}

func (v *NullableConditionals) Set(val *Conditionals) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionals) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionals(val *Conditionals) *NullableConditionals {
	return &NullableConditionals{value: val, isSet: true}
}

func (v NullableConditionals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

