/*
Services API

Azion Services

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package services

import (
	"encoding/json"
)

// checks if the ResourceResponseWithTotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceResponseWithTotal{}

// ResourceResponseWithTotal struct for ResourceResponseWithTotal
type ResourceResponseWithTotal struct {
	Resources []ResourceResponse `json:"resources"`
	Total int64 `json:"total"`
}

// NewResourceResponseWithTotal instantiates a new ResourceResponseWithTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceResponseWithTotal(resources []ResourceResponse, total int64) *ResourceResponseWithTotal {
	this := ResourceResponseWithTotal{}
	this.Resources = resources
	this.Total = total
	return &this
}

// NewResourceResponseWithTotalWithDefaults instantiates a new ResourceResponseWithTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceResponseWithTotalWithDefaults() *ResourceResponseWithTotal {
	this := ResourceResponseWithTotal{}
	return &this
}

// GetResources returns the Resources field value
func (o *ResourceResponseWithTotal) GetResources() []ResourceResponse {
	if o == nil {
		var ret []ResourceResponse
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *ResourceResponseWithTotal) GetResourcesOk() ([]ResourceResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *ResourceResponseWithTotal) SetResources(v []ResourceResponse) {
	o.Resources = v
}

// GetTotal returns the Total field value
func (o *ResourceResponseWithTotal) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ResourceResponseWithTotal) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ResourceResponseWithTotal) SetTotal(v int64) {
	o.Total = v
}

func (o ResourceResponseWithTotal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceResponseWithTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resources"] = o.Resources
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

type NullableResourceResponseWithTotal struct {
	value *ResourceResponseWithTotal
	isSet bool
}

func (v NullableResourceResponseWithTotal) Get() *ResourceResponseWithTotal {
	return v.value
}

func (v *NullableResourceResponseWithTotal) Set(val *ResourceResponseWithTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceResponseWithTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceResponseWithTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceResponseWithTotal(val *ResourceResponseWithTotal) *NullableResourceResponseWithTotal {
	return &NullableResourceResponseWithTotal{value: val, isSet: true}
}

func (v NullableResourceResponseWithTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceResponseWithTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


