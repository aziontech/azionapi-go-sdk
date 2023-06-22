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

// checks if the CreateDeviceGroupsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceGroupsRequest{}

// CreateDeviceGroupsRequest struct for CreateDeviceGroupsRequest
type CreateDeviceGroupsRequest struct {
	Name *string `json:"name,omitempty"`
	UserAgent string `json:"user_agent"`
	Addresses string `json:"addresses"`
}

// NewCreateDeviceGroupsRequest instantiates a new CreateDeviceGroupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceGroupsRequest(userAgent string, addresses string) *CreateDeviceGroupsRequest {
	this := CreateDeviceGroupsRequest{}
	this.UserAgent = userAgent
	this.Addresses = addresses
	return &this
}

// NewCreateDeviceGroupsRequestWithDefaults instantiates a new CreateDeviceGroupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceGroupsRequestWithDefaults() *CreateDeviceGroupsRequest {
	this := CreateDeviceGroupsRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateDeviceGroupsRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceGroupsRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateDeviceGroupsRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateDeviceGroupsRequest) SetName(v string) {
	o.Name = &v
}

// GetUserAgent returns the UserAgent field value
func (o *CreateDeviceGroupsRequest) GetUserAgent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value
// and a boolean to check if the value has been set.
func (o *CreateDeviceGroupsRequest) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAgent, true
}

// SetUserAgent sets field value
func (o *CreateDeviceGroupsRequest) SetUserAgent(v string) {
	o.UserAgent = v
}

// GetAddresses returns the Addresses field value
func (o *CreateDeviceGroupsRequest) GetAddresses() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *CreateDeviceGroupsRequest) GetAddressesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value
func (o *CreateDeviceGroupsRequest) SetAddresses(v string) {
	o.Addresses = v
}

func (o CreateDeviceGroupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceGroupsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["user_agent"] = o.UserAgent
	toSerialize["addresses"] = o.Addresses
	return toSerialize, nil
}

type NullableCreateDeviceGroupsRequest struct {
	value *CreateDeviceGroupsRequest
	isSet bool
}

func (v NullableCreateDeviceGroupsRequest) Get() *CreateDeviceGroupsRequest {
	return v.value
}

func (v *NullableCreateDeviceGroupsRequest) Set(val *CreateDeviceGroupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceGroupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceGroupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceGroupsRequest(val *CreateDeviceGroupsRequest) *NullableCreateDeviceGroupsRequest {
	return &NullableCreateDeviceGroupsRequest{value: val, isSet: true}
}

func (v NullableCreateDeviceGroupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceGroupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


