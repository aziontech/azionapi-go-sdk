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

// checks if the EndpointGoogleBigQueryServiceAccountKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointGoogleBigQueryServiceAccountKey{}

// EndpointGoogleBigQueryServiceAccountKey struct for EndpointGoogleBigQueryServiceAccountKey
type EndpointGoogleBigQueryServiceAccountKey struct {
	ServiceAccountKey *string `json:"service_account_key,omitempty"`
}

// NewEndpointGoogleBigQueryServiceAccountKey instantiates a new EndpointGoogleBigQueryServiceAccountKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointGoogleBigQueryServiceAccountKey() *EndpointGoogleBigQueryServiceAccountKey {
	this := EndpointGoogleBigQueryServiceAccountKey{}
	return &this
}

// NewEndpointGoogleBigQueryServiceAccountKeyWithDefaults instantiates a new EndpointGoogleBigQueryServiceAccountKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointGoogleBigQueryServiceAccountKeyWithDefaults() *EndpointGoogleBigQueryServiceAccountKey {
	this := EndpointGoogleBigQueryServiceAccountKey{}
	return &this
}

// GetServiceAccountKey returns the ServiceAccountKey field value if set, zero value otherwise.
func (o *EndpointGoogleBigQueryServiceAccountKey) GetServiceAccountKey() string {
	if o == nil || IsNil(o.ServiceAccountKey) {
		var ret string
		return ret
	}
	return *o.ServiceAccountKey
}

// GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointGoogleBigQueryServiceAccountKey) GetServiceAccountKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAccountKey) {
		return nil, false
	}
	return o.ServiceAccountKey, true
}

// HasServiceAccountKey returns a boolean if a field has been set.
func (o *EndpointGoogleBigQueryServiceAccountKey) HasServiceAccountKey() bool {
	if o != nil && !IsNil(o.ServiceAccountKey) {
		return true
	}

	return false
}

// SetServiceAccountKey gets a reference to the given string and assigns it to the ServiceAccountKey field.
func (o *EndpointGoogleBigQueryServiceAccountKey) SetServiceAccountKey(v string) {
	o.ServiceAccountKey = &v
}

func (o EndpointGoogleBigQueryServiceAccountKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointGoogleBigQueryServiceAccountKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceAccountKey) {
		toSerialize["service_account_key"] = o.ServiceAccountKey
	}
	return toSerialize, nil
}

type NullableEndpointGoogleBigQueryServiceAccountKey struct {
	value *EndpointGoogleBigQueryServiceAccountKey
	isSet bool
}

func (v NullableEndpointGoogleBigQueryServiceAccountKey) Get() *EndpointGoogleBigQueryServiceAccountKey {
	return v.value
}

func (v *NullableEndpointGoogleBigQueryServiceAccountKey) Set(val *EndpointGoogleBigQueryServiceAccountKey) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointGoogleBigQueryServiceAccountKey) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointGoogleBigQueryServiceAccountKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointGoogleBigQueryServiceAccountKey(val *EndpointGoogleBigQueryServiceAccountKey) *NullableEndpointGoogleBigQueryServiceAccountKey {
	return &NullableEndpointGoogleBigQueryServiceAccountKey{value: val, isSet: true}
}

func (v NullableEndpointGoogleBigQueryServiceAccountKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointGoogleBigQueryServiceAccountKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


