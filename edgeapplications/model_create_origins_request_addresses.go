/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
)

// checks if the CreateOriginsRequestAddresses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOriginsRequestAddresses{}

// CreateOriginsRequestAddresses struct for CreateOriginsRequestAddresses
type CreateOriginsRequestAddresses struct {
	Address string `json:"address"`
}

// NewCreateOriginsRequestAddresses instantiates a new CreateOriginsRequestAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOriginsRequestAddresses(address string) *CreateOriginsRequestAddresses {
	this := CreateOriginsRequestAddresses{}
	this.Address = address
	return &this
}

// NewCreateOriginsRequestAddressesWithDefaults instantiates a new CreateOriginsRequestAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOriginsRequestAddressesWithDefaults() *CreateOriginsRequestAddresses {
	this := CreateOriginsRequestAddresses{}
	return &this
}

// GetAddress returns the Address field value
func (o *CreateOriginsRequestAddresses) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CreateOriginsRequestAddresses) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CreateOriginsRequestAddresses) SetAddress(v string) {
	o.Address = v
}

func (o CreateOriginsRequestAddresses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOriginsRequestAddresses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

type NullableCreateOriginsRequestAddresses struct {
	value *CreateOriginsRequestAddresses
	isSet bool
}

func (v NullableCreateOriginsRequestAddresses) Get() *CreateOriginsRequestAddresses {
	return v.value
}

func (v *NullableCreateOriginsRequestAddresses) Set(val *CreateOriginsRequestAddresses) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOriginsRequestAddresses) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOriginsRequestAddresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOriginsRequestAddresses(val *CreateOriginsRequestAddresses) *NullableCreateOriginsRequestAddresses {
	return &NullableCreateOriginsRequestAddresses{value: val, isSet: true}
}

func (v NullableCreateOriginsRequestAddresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOriginsRequestAddresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


