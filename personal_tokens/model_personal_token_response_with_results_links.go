/*
Personal Tokens - OpenAPI

The Personal Tokens API allows you to manage your existing personal tokens. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package personal_tokens

import (
	"encoding/json"
)

// checks if the PersonalTokenResponseWithResultsLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonalTokenResponseWithResultsLinks{}

// PersonalTokenResponseWithResultsLinks struct for PersonalTokenResponseWithResultsLinks
type PersonalTokenResponseWithResultsLinks struct {
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
}

// NewPersonalTokenResponseWithResultsLinks instantiates a new PersonalTokenResponseWithResultsLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonalTokenResponseWithResultsLinks() *PersonalTokenResponseWithResultsLinks {
	this := PersonalTokenResponseWithResultsLinks{}
	return &this
}

// NewPersonalTokenResponseWithResultsLinksWithDefaults instantiates a new PersonalTokenResponseWithResultsLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonalTokenResponseWithResultsLinksWithDefaults() *PersonalTokenResponseWithResultsLinks {
	this := PersonalTokenResponseWithResultsLinks{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersonalTokenResponseWithResultsLinks) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PersonalTokenResponseWithResultsLinks) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PersonalTokenResponseWithResultsLinks) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PersonalTokenResponseWithResultsLinks) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PersonalTokenResponseWithResultsLinks) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PersonalTokenResponseWithResultsLinks) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersonalTokenResponseWithResultsLinks) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PersonalTokenResponseWithResultsLinks) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PersonalTokenResponseWithResultsLinks) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PersonalTokenResponseWithResultsLinks) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PersonalTokenResponseWithResultsLinks) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PersonalTokenResponseWithResultsLinks) UnsetPrevious() {
	o.Previous.Unset()
}

func (o PersonalTokenResponseWithResultsLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonalTokenResponseWithResultsLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	return toSerialize, nil
}

type NullablePersonalTokenResponseWithResultsLinks struct {
	value *PersonalTokenResponseWithResultsLinks
	isSet bool
}

func (v NullablePersonalTokenResponseWithResultsLinks) Get() *PersonalTokenResponseWithResultsLinks {
	return v.value
}

func (v *NullablePersonalTokenResponseWithResultsLinks) Set(val *PersonalTokenResponseWithResultsLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonalTokenResponseWithResultsLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonalTokenResponseWithResultsLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonalTokenResponseWithResultsLinks(val *PersonalTokenResponseWithResultsLinks) *NullablePersonalTokenResponseWithResultsLinks {
	return &NullablePersonalTokenResponseWithResultsLinks{value: val, isSet: true}
}

func (v NullablePersonalTokenResponseWithResultsLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonalTokenResponseWithResultsLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


