/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RulesEngineBehaviorObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulesEngineBehaviorObject{}

// RulesEngineBehaviorObject struct for RulesEngineBehaviorObject
type RulesEngineBehaviorObject struct {
	Name string `json:"name"`
	Target RulesEngineBehaviorObjectTarget `json:"target"`
}

type _RulesEngineBehaviorObject RulesEngineBehaviorObject

// NewRulesEngineBehaviorObject instantiates a new RulesEngineBehaviorObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesEngineBehaviorObject(name string, target RulesEngineBehaviorObjectTarget) *RulesEngineBehaviorObject {
	this := RulesEngineBehaviorObject{}
	this.Name = name
	this.Target = target
	return &this
}

// NewRulesEngineBehaviorObjectWithDefaults instantiates a new RulesEngineBehaviorObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesEngineBehaviorObjectWithDefaults() *RulesEngineBehaviorObject {
	this := RulesEngineBehaviorObject{}
	return &this
}

// GetName returns the Name field value
func (o *RulesEngineBehaviorObject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RulesEngineBehaviorObject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RulesEngineBehaviorObject) SetName(v string) {
	o.Name = v
}

// GetTarget returns the Target field value
func (o *RulesEngineBehaviorObject) GetTarget() RulesEngineBehaviorObjectTarget {
	if o == nil {
		var ret RulesEngineBehaviorObjectTarget
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *RulesEngineBehaviorObject) GetTargetOk() (*RulesEngineBehaviorObjectTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *RulesEngineBehaviorObject) SetTarget(v RulesEngineBehaviorObjectTarget) {
	o.Target = v
}

func (o RulesEngineBehaviorObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulesEngineBehaviorObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["target"] = o.Target
	return toSerialize, nil
}

func (o *RulesEngineBehaviorObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"target",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRulesEngineBehaviorObject := _RulesEngineBehaviorObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRulesEngineBehaviorObject)

	if err != nil {
		return err
	}

	*o = RulesEngineBehaviorObject(varRulesEngineBehaviorObject)

	return err
}

type NullableRulesEngineBehaviorObject struct {
	value *RulesEngineBehaviorObject
	isSet bool
}

func (v NullableRulesEngineBehaviorObject) Get() *RulesEngineBehaviorObject {
	return v.value
}

func (v *NullableRulesEngineBehaviorObject) Set(val *RulesEngineBehaviorObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesEngineBehaviorObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesEngineBehaviorObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesEngineBehaviorObject(val *RulesEngineBehaviorObject) *NullableRulesEngineBehaviorObject {
	return &NullableRulesEngineBehaviorObject{value: val, isSet: true}
}

func (v NullableRulesEngineBehaviorObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesEngineBehaviorObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


