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

// checks if the GenericBehaviorDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericBehaviorDetails{}

// GenericBehaviorDetails struct for GenericBehaviorDetails
type GenericBehaviorDetails struct {
	Name *string `json:"name,omitempty"`
}

// NewGenericBehaviorDetails instantiates a new GenericBehaviorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericBehaviorDetails() *GenericBehaviorDetails {
	this := GenericBehaviorDetails{}
	return &this
}

// NewGenericBehaviorDetailsWithDefaults instantiates a new GenericBehaviorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericBehaviorDetailsWithDefaults() *GenericBehaviorDetails {
	this := GenericBehaviorDetails{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GenericBehaviorDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericBehaviorDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GenericBehaviorDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GenericBehaviorDetails) SetName(v string) {
	o.Name = &v
}

func (o GenericBehaviorDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericBehaviorDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGenericBehaviorDetails struct {
	value *GenericBehaviorDetails
	isSet bool
}

func (v NullableGenericBehaviorDetails) Get() *GenericBehaviorDetails {
	return v.value
}

func (v *NullableGenericBehaviorDetails) Set(val *GenericBehaviorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericBehaviorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericBehaviorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericBehaviorDetails(val *GenericBehaviorDetails) *NullableGenericBehaviorDetails {
	return &NullableGenericBehaviorDetails{value: val, isSet: true}
}

func (v NullableGenericBehaviorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericBehaviorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


