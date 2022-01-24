/*
Edge Function

Azion Edge Function API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgefunctions

import (
	"encoding/json"
)

// InlineResponse200Links struct for InlineResponse200Links
type InlineResponse200Links struct {
	Previous *string `json:"previous,omitempty"`
	Next *string `json:"next,omitempty"`
}

// NewInlineResponse200Links instantiates a new InlineResponse200Links object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200Links() *InlineResponse200Links {
	this := InlineResponse200Links{}
	return &this
}

// NewInlineResponse200LinksWithDefaults instantiates a new InlineResponse200Links object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200LinksWithDefaults() *InlineResponse200Links {
	this := InlineResponse200Links{}
	return &this
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *InlineResponse200Links) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200Links) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *InlineResponse200Links) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *InlineResponse200Links) SetPrevious(v string) {
	o.Previous = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *InlineResponse200Links) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200Links) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *InlineResponse200Links) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *InlineResponse200Links) SetNext(v string) {
	o.Next = &v
}

func (o InlineResponse200Links) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200Links struct {
	value *InlineResponse200Links
	isSet bool
}

func (v NullableInlineResponse200Links) Get() *InlineResponse200Links {
	return v.value
}

func (v *NullableInlineResponse200Links) Set(val *InlineResponse200Links) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200Links) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200Links) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200Links(val *InlineResponse200Links) *NullableInlineResponse200Links {
	return &NullableInlineResponse200Links{value: val, isSet: true}
}

func (v NullableInlineResponse200Links) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200Links) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


