/*
Intelligent DNS

Azion Intelligent DNS API

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idns

import (
	"encoding/json"
)

// checks if the GetZonesResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetZonesResponseLinks{}

// GetZonesResponseLinks struct for GetZonesResponseLinks
type GetZonesResponseLinks struct {
	Previous NullableString `json:"previous,omitempty"`
	Next NullableString `json:"next,omitempty"`
}

// NewGetZonesResponseLinks instantiates a new GetZonesResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetZonesResponseLinks() *GetZonesResponseLinks {
	this := GetZonesResponseLinks{}
	return &this
}

// NewGetZonesResponseLinksWithDefaults instantiates a new GetZonesResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetZonesResponseLinksWithDefaults() *GetZonesResponseLinks {
	this := GetZonesResponseLinks{}
	return &this
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetZonesResponseLinks) GetPrevious() string {
	if o == nil || isNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetZonesResponseLinks) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *GetZonesResponseLinks) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *GetZonesResponseLinks) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *GetZonesResponseLinks) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *GetZonesResponseLinks) UnsetPrevious() {
	o.Previous.Unset()
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetZonesResponseLinks) GetNext() string {
	if o == nil || isNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetZonesResponseLinks) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *GetZonesResponseLinks) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *GetZonesResponseLinks) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *GetZonesResponseLinks) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *GetZonesResponseLinks) UnsetNext() {
	o.Next.Unset()
}

func (o GetZonesResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetZonesResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	return toSerialize, nil
}

type NullableGetZonesResponseLinks struct {
	value *GetZonesResponseLinks
	isSet bool
}

func (v NullableGetZonesResponseLinks) Get() *GetZonesResponseLinks {
	return v.value
}

func (v *NullableGetZonesResponseLinks) Set(val *GetZonesResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGetZonesResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGetZonesResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetZonesResponseLinks(val *GetZonesResponseLinks) *NullableGetZonesResponseLinks {
	return &NullableGetZonesResponseLinks{value: val, isSet: true}
}

func (v NullableGetZonesResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetZonesResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

