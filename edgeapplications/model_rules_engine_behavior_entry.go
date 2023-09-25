/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
	"fmt"
)

// RulesEngineBehaviorEntry struct for RulesEngineBehaviorEntry
type RulesEngineBehaviorEntry struct {
	RulesEngineBehaviorObject *RulesEngineBehaviorObject
	RulesEngineBehaviorString *RulesEngineBehaviorString
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RulesEngineBehaviorEntry) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RulesEngineBehaviorObject
	err = json.Unmarshal(data, &dst.RulesEngineBehaviorObject);
	if err == nil {
		jsonRulesEngineBehaviorObject, _ := json.Marshal(dst.RulesEngineBehaviorObject)
		if string(jsonRulesEngineBehaviorObject) == "{}" { // empty struct
			dst.RulesEngineBehaviorObject = nil
		} else {
			return nil // data stored in dst.RulesEngineBehaviorObject, return on the first match
		}
	} else {
		dst.RulesEngineBehaviorObject = nil
	}

	// try to unmarshal JSON data into RulesEngineBehaviorString
	err = json.Unmarshal(data, &dst.RulesEngineBehaviorString);
	if err == nil {
		jsonRulesEngineBehaviorString, _ := json.Marshal(dst.RulesEngineBehaviorString)
		if string(jsonRulesEngineBehaviorString) == "{}" { // empty struct
			dst.RulesEngineBehaviorString = nil
		} else {
			return nil // data stored in dst.RulesEngineBehaviorString, return on the first match
		}
	} else {
		dst.RulesEngineBehaviorString = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RulesEngineBehaviorEntry)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RulesEngineBehaviorEntry) MarshalJSON() ([]byte, error) {
	if src.RulesEngineBehaviorObject != nil {
		return json.Marshal(&src.RulesEngineBehaviorObject)
	}

	if src.RulesEngineBehaviorString != nil {
		return json.Marshal(&src.RulesEngineBehaviorString)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRulesEngineBehaviorEntry struct {
	value *RulesEngineBehaviorEntry
	isSet bool
}

func (v NullableRulesEngineBehaviorEntry) Get() *RulesEngineBehaviorEntry {
	return v.value
}

func (v *NullableRulesEngineBehaviorEntry) Set(val *RulesEngineBehaviorEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesEngineBehaviorEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesEngineBehaviorEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesEngineBehaviorEntry(val *RulesEngineBehaviorEntry) *NullableRulesEngineBehaviorEntry {
	return &NullableRulesEngineBehaviorEntry{value: val, isSet: true}
}

func (v NullableRulesEngineBehaviorEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesEngineBehaviorEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

