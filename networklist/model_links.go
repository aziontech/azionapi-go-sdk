/*
Azion API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networklist

import (
	"encoding/json"
)

// checks if the Links type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Links{}

// Links struct for Links
type Links struct {
	Previous *string `json:"previous,omitempty"`
	Next *string `json:"next,omitempty"`
}

// NewLinks instantiates a new Links object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinks() *Links {
	this := Links{}
	return &this
}

// NewLinksWithDefaults instantiates a new Links object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksWithDefaults() *Links {
	this := Links{}
	return &this
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *Links) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *Links) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *Links) SetPrevious(v string) {
	o.Previous = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *Links) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *Links) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *Links) SetNext(v string) {
	o.Next = &v
}

func (o Links) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Links) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

type NullableLinks struct {
	value *Links
	isSet bool
}

func (v NullableLinks) Get() *Links {
	return v.value
}

func (v *NullableLinks) Set(val *Links) {
	v.value = val
	v.isSet = true
}

func (v NullableLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinks(val *Links) *NullableLinks {
	return &NullableLinks{value: val, isSet: true}
}

func (v NullableLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


