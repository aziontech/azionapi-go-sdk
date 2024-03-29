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

// checks if the NullArgumentBehavior type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NullArgumentBehavior{}

// NullArgumentBehavior struct for NullArgumentBehavior
type NullArgumentBehavior struct {
	Name *string `json:"name,omitempty"`
	Argument NullableInt32 `json:"argument,omitempty"`
}

// NewNullArgumentBehavior instantiates a new NullArgumentBehavior object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNullArgumentBehavior() *NullArgumentBehavior {
	this := NullArgumentBehavior{}
	return &this
}

// NewNullArgumentBehaviorWithDefaults instantiates a new NullArgumentBehavior object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNullArgumentBehaviorWithDefaults() *NullArgumentBehavior {
	this := NullArgumentBehavior{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NullArgumentBehavior) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullArgumentBehavior) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NullArgumentBehavior) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NullArgumentBehavior) SetName(v string) {
	o.Name = &v
}

// GetArgument returns the Argument field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NullArgumentBehavior) GetArgument() int32 {
	if o == nil || IsNil(o.Argument.Get()) {
		var ret int32
		return ret
	}
	return *o.Argument.Get()
}

// GetArgumentOk returns a tuple with the Argument field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullArgumentBehavior) GetArgumentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Argument.Get(), o.Argument.IsSet()
}

// HasArgument returns a boolean if a field has been set.
func (o *NullArgumentBehavior) HasArgument() bool {
	if o != nil && o.Argument.IsSet() {
		return true
	}

	return false
}

// SetArgument gets a reference to the given NullableInt32 and assigns it to the Argument field.
func (o *NullArgumentBehavior) SetArgument(v int32) {
	o.Argument.Set(&v)
}
// SetArgumentNil sets the value for Argument to be an explicit nil
func (o *NullArgumentBehavior) SetArgumentNil() {
	o.Argument.Set(nil)
}

// UnsetArgument ensures that no value is present for Argument, not even an explicit nil
func (o *NullArgumentBehavior) UnsetArgument() {
	o.Argument.Unset()
}

func (o NullArgumentBehavior) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NullArgumentBehavior) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Argument.IsSet() {
		toSerialize["argument"] = o.Argument.Get()
	}
	return toSerialize, nil
}

type NullableNullArgumentBehavior struct {
	value *NullArgumentBehavior
	isSet bool
}

func (v NullableNullArgumentBehavior) Get() *NullArgumentBehavior {
	return v.value
}

func (v *NullableNullArgumentBehavior) Set(val *NullArgumentBehavior) {
	v.value = val
	v.isSet = true
}

func (v NullableNullArgumentBehavior) IsSet() bool {
	return v.isSet
}

func (v *NullableNullArgumentBehavior) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullArgumentBehavior(val *NullArgumentBehavior) *NullableNullArgumentBehavior {
	return &NullableNullArgumentBehavior{value: val, isSet: true}
}

func (v NullableNullArgumentBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullArgumentBehavior) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


