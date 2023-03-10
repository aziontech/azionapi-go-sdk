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

// checks if the DnsSecDelegationSignerDigestType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsSecDelegationSignerDigestType{}

// DnsSecDelegationSignerDigestType struct for DnsSecDelegationSignerDigestType
type DnsSecDelegationSignerDigestType struct {
	Id *int32 `json:"id,omitempty"`
	Slug *string `json:"slug,omitempty"`
}

// NewDnsSecDelegationSignerDigestType instantiates a new DnsSecDelegationSignerDigestType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsSecDelegationSignerDigestType() *DnsSecDelegationSignerDigestType {
	this := DnsSecDelegationSignerDigestType{}
	return &this
}

// NewDnsSecDelegationSignerDigestTypeWithDefaults instantiates a new DnsSecDelegationSignerDigestType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsSecDelegationSignerDigestTypeWithDefaults() *DnsSecDelegationSignerDigestType {
	this := DnsSecDelegationSignerDigestType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DnsSecDelegationSignerDigestType) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsSecDelegationSignerDigestType) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DnsSecDelegationSignerDigestType) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DnsSecDelegationSignerDigestType) SetId(v int32) {
	o.Id = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *DnsSecDelegationSignerDigestType) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsSecDelegationSignerDigestType) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *DnsSecDelegationSignerDigestType) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *DnsSecDelegationSignerDigestType) SetSlug(v string) {
	o.Slug = &v
}

func (o DnsSecDelegationSignerDigestType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsSecDelegationSignerDigestType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	return toSerialize, nil
}

type NullableDnsSecDelegationSignerDigestType struct {
	value *DnsSecDelegationSignerDigestType
	isSet bool
}

func (v NullableDnsSecDelegationSignerDigestType) Get() *DnsSecDelegationSignerDigestType {
	return v.value
}

func (v *NullableDnsSecDelegationSignerDigestType) Set(val *DnsSecDelegationSignerDigestType) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsSecDelegationSignerDigestType) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsSecDelegationSignerDigestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsSecDelegationSignerDigestType(val *DnsSecDelegationSignerDigestType) *NullableDnsSecDelegationSignerDigestType {
	return &NullableDnsSecDelegationSignerDigestType{value: val, isSet: true}
}

func (v NullableDnsSecDelegationSignerDigestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsSecDelegationSignerDigestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

