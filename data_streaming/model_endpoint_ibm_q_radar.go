/*
Data Streaming - OpenAPI

The Data Streaming API allows you to manage your existing data streamings and templates. Data Streaming allows you to feed your stream processing, SIEM, and big data platforms with the event logs from your applications on Azion in real time. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data_streaming

import (
	"encoding/json"
)

// checks if the EndpointIBMQRadar type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointIBMQRadar{}

// EndpointIBMQRadar struct for EndpointIBMQRadar
type EndpointIBMQRadar struct {
	EndpointType *string `json:"endpoint_type,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewEndpointIBMQRadar instantiates a new EndpointIBMQRadar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointIBMQRadar() *EndpointIBMQRadar {
	this := EndpointIBMQRadar{}
	return &this
}

// NewEndpointIBMQRadarWithDefaults instantiates a new EndpointIBMQRadar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointIBMQRadarWithDefaults() *EndpointIBMQRadar {
	this := EndpointIBMQRadar{}
	return &this
}

// GetEndpointType returns the EndpointType field value if set, zero value otherwise.
func (o *EndpointIBMQRadar) GetEndpointType() string {
	if o == nil || IsNil(o.EndpointType) {
		var ret string
		return ret
	}
	return *o.EndpointType
}

// GetEndpointTypeOk returns a tuple with the EndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointIBMQRadar) GetEndpointTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointType) {
		return nil, false
	}
	return o.EndpointType, true
}

// HasEndpointType returns a boolean if a field has been set.
func (o *EndpointIBMQRadar) HasEndpointType() bool {
	if o != nil && !IsNil(o.EndpointType) {
		return true
	}

	return false
}

// SetEndpointType gets a reference to the given string and assigns it to the EndpointType field.
func (o *EndpointIBMQRadar) SetEndpointType(v string) {
	o.EndpointType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *EndpointIBMQRadar) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointIBMQRadar) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *EndpointIBMQRadar) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *EndpointIBMQRadar) SetUrl(v string) {
	o.Url = &v
}

func (o EndpointIBMQRadar) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointIBMQRadar) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointType) {
		toSerialize["endpoint_type"] = o.EndpointType
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableEndpointIBMQRadar struct {
	value *EndpointIBMQRadar
	isSet bool
}

func (v NullableEndpointIBMQRadar) Get() *EndpointIBMQRadar {
	return v.value
}

func (v *NullableEndpointIBMQRadar) Set(val *EndpointIBMQRadar) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointIBMQRadar) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointIBMQRadar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointIBMQRadar(val *EndpointIBMQRadar) *NullableEndpointIBMQRadar {
	return &NullableEndpointIBMQRadar{value: val, isSet: true}
}

func (v NullableEndpointIBMQRadar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointIBMQRadar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


