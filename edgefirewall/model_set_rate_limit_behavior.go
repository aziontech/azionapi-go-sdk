/*
Edge Firewall API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgefirewall

import (
	"encoding/json"
)

// checks if the SetRateLimitBehavior type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetRateLimitBehavior{}

// SetRateLimitBehavior struct for SetRateLimitBehavior
type SetRateLimitBehavior struct {
	Name *string `json:"name,omitempty"`
	Argument *SetRateLimitBehaviorArgument `json:"argument,omitempty"`
}

// NewSetRateLimitBehavior instantiates a new SetRateLimitBehavior object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetRateLimitBehavior() *SetRateLimitBehavior {
	this := SetRateLimitBehavior{}
	return &this
}

// NewSetRateLimitBehaviorWithDefaults instantiates a new SetRateLimitBehavior object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetRateLimitBehaviorWithDefaults() *SetRateLimitBehavior {
	this := SetRateLimitBehavior{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SetRateLimitBehavior) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRateLimitBehavior) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SetRateLimitBehavior) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SetRateLimitBehavior) SetName(v string) {
	o.Name = &v
}

// GetArgument returns the Argument field value if set, zero value otherwise.
func (o *SetRateLimitBehavior) GetArgument() SetRateLimitBehaviorArgument {
	if o == nil || IsNil(o.Argument) {
		var ret SetRateLimitBehaviorArgument
		return ret
	}
	return *o.Argument
}

// GetArgumentOk returns a tuple with the Argument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRateLimitBehavior) GetArgumentOk() (*SetRateLimitBehaviorArgument, bool) {
	if o == nil || IsNil(o.Argument) {
		return nil, false
	}
	return o.Argument, true
}

// HasArgument returns a boolean if a field has been set.
func (o *SetRateLimitBehavior) HasArgument() bool {
	if o != nil && !IsNil(o.Argument) {
		return true
	}

	return false
}

// SetArgument gets a reference to the given SetRateLimitBehaviorArgument and assigns it to the Argument field.
func (o *SetRateLimitBehavior) SetArgument(v SetRateLimitBehaviorArgument) {
	o.Argument = &v
}

func (o SetRateLimitBehavior) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetRateLimitBehavior) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Argument) {
		toSerialize["argument"] = o.Argument
	}
	return toSerialize, nil
}

type NullableSetRateLimitBehavior struct {
	value *SetRateLimitBehavior
	isSet bool
}

func (v NullableSetRateLimitBehavior) Get() *SetRateLimitBehavior {
	return v.value
}

func (v *NullableSetRateLimitBehavior) Set(val *SetRateLimitBehavior) {
	v.value = val
	v.isSet = true
}

func (v NullableSetRateLimitBehavior) IsSet() bool {
	return v.isSet
}

func (v *NullableSetRateLimitBehavior) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetRateLimitBehavior(val *SetRateLimitBehavior) *NullableSetRateLimitBehavior {
	return &NullableSetRateLimitBehavior{value: val, isSet: true}
}

func (v NullableSetRateLimitBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetRateLimitBehavior) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

