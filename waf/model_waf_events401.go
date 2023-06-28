/*
Web Application Firewall API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package waf

import (
	"encoding/json"
)

// checks if the WAFEvents401 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WAFEvents401{}

// WAFEvents401 struct for WAFEvents401
type WAFEvents401 struct {
	Detail *string `json:"detail,omitempty"`
}

// NewWAFEvents401 instantiates a new WAFEvents401 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWAFEvents401() *WAFEvents401 {
	this := WAFEvents401{}
	return &this
}

// NewWAFEvents401WithDefaults instantiates a new WAFEvents401 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWAFEvents401WithDefaults() *WAFEvents401 {
	this := WAFEvents401{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *WAFEvents401) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFEvents401) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *WAFEvents401) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *WAFEvents401) SetDetail(v string) {
	o.Detail = &v
}

func (o WAFEvents401) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WAFEvents401) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableWAFEvents401 struct {
	value *WAFEvents401
	isSet bool
}

func (v NullableWAFEvents401) Get() *WAFEvents401 {
	return v.value
}

func (v *NullableWAFEvents401) Set(val *WAFEvents401) {
	v.value = val
	v.isSet = true
}

func (v NullableWAFEvents401) IsSet() bool {
	return v.isSet
}

func (v *NullableWAFEvents401) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWAFEvents401(val *WAFEvents401) *NullableWAFEvents401 {
	return &NullableWAFEvents401{value: val, isSet: true}
}

func (v NullableWAFEvents401) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWAFEvents401) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

