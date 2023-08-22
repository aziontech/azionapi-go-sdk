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

// checks if the EndpointKafka type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointKafka{}

// EndpointKafka struct for EndpointKafka
type EndpointKafka struct {
	EndpointType *string `json:"endpoint_type,omitempty"`
	KafkaTopic *string `json:"kafka_topic,omitempty"`
	BootstrapServers *string `json:"bootstrap_servers,omitempty"`
	UseTls *bool `json:"use_tls,omitempty"`
}

// NewEndpointKafka instantiates a new EndpointKafka object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointKafka() *EndpointKafka {
	this := EndpointKafka{}
	return &this
}

// NewEndpointKafkaWithDefaults instantiates a new EndpointKafka object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointKafkaWithDefaults() *EndpointKafka {
	this := EndpointKafka{}
	return &this
}

// GetEndpointType returns the EndpointType field value if set, zero value otherwise.
func (o *EndpointKafka) GetEndpointType() string {
	if o == nil || IsNil(o.EndpointType) {
		var ret string
		return ret
	}
	return *o.EndpointType
}

// GetEndpointTypeOk returns a tuple with the EndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointKafka) GetEndpointTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointType) {
		return nil, false
	}
	return o.EndpointType, true
}

// HasEndpointType returns a boolean if a field has been set.
func (o *EndpointKafka) HasEndpointType() bool {
	if o != nil && !IsNil(o.EndpointType) {
		return true
	}

	return false
}

// SetEndpointType gets a reference to the given string and assigns it to the EndpointType field.
func (o *EndpointKafka) SetEndpointType(v string) {
	o.EndpointType = &v
}

// GetKafkaTopic returns the KafkaTopic field value if set, zero value otherwise.
func (o *EndpointKafka) GetKafkaTopic() string {
	if o == nil || IsNil(o.KafkaTopic) {
		var ret string
		return ret
	}
	return *o.KafkaTopic
}

// GetKafkaTopicOk returns a tuple with the KafkaTopic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointKafka) GetKafkaTopicOk() (*string, bool) {
	if o == nil || IsNil(o.KafkaTopic) {
		return nil, false
	}
	return o.KafkaTopic, true
}

// HasKafkaTopic returns a boolean if a field has been set.
func (o *EndpointKafka) HasKafkaTopic() bool {
	if o != nil && !IsNil(o.KafkaTopic) {
		return true
	}

	return false
}

// SetKafkaTopic gets a reference to the given string and assigns it to the KafkaTopic field.
func (o *EndpointKafka) SetKafkaTopic(v string) {
	o.KafkaTopic = &v
}

// GetBootstrapServers returns the BootstrapServers field value if set, zero value otherwise.
func (o *EndpointKafka) GetBootstrapServers() string {
	if o == nil || IsNil(o.BootstrapServers) {
		var ret string
		return ret
	}
	return *o.BootstrapServers
}

// GetBootstrapServersOk returns a tuple with the BootstrapServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointKafka) GetBootstrapServersOk() (*string, bool) {
	if o == nil || IsNil(o.BootstrapServers) {
		return nil, false
	}
	return o.BootstrapServers, true
}

// HasBootstrapServers returns a boolean if a field has been set.
func (o *EndpointKafka) HasBootstrapServers() bool {
	if o != nil && !IsNil(o.BootstrapServers) {
		return true
	}

	return false
}

// SetBootstrapServers gets a reference to the given string and assigns it to the BootstrapServers field.
func (o *EndpointKafka) SetBootstrapServers(v string) {
	o.BootstrapServers = &v
}

// GetUseTls returns the UseTls field value if set, zero value otherwise.
func (o *EndpointKafka) GetUseTls() bool {
	if o == nil || IsNil(o.UseTls) {
		var ret bool
		return ret
	}
	return *o.UseTls
}

// GetUseTlsOk returns a tuple with the UseTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointKafka) GetUseTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseTls) {
		return nil, false
	}
	return o.UseTls, true
}

// HasUseTls returns a boolean if a field has been set.
func (o *EndpointKafka) HasUseTls() bool {
	if o != nil && !IsNil(o.UseTls) {
		return true
	}

	return false
}

// SetUseTls gets a reference to the given bool and assigns it to the UseTls field.
func (o *EndpointKafka) SetUseTls(v bool) {
	o.UseTls = &v
}

func (o EndpointKafka) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointKafka) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointType) {
		toSerialize["endpoint_type"] = o.EndpointType
	}
	if !IsNil(o.KafkaTopic) {
		toSerialize["kafka_topic"] = o.KafkaTopic
	}
	if !IsNil(o.BootstrapServers) {
		toSerialize["bootstrap_servers"] = o.BootstrapServers
	}
	if !IsNil(o.UseTls) {
		toSerialize["use_tls"] = o.UseTls
	}
	return toSerialize, nil
}

type NullableEndpointKafka struct {
	value *EndpointKafka
	isSet bool
}

func (v NullableEndpointKafka) Get() *EndpointKafka {
	return v.value
}

func (v *NullableEndpointKafka) Set(val *EndpointKafka) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointKafka) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointKafka) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointKafka(val *EndpointKafka) *NullableEndpointKafka {
	return &NullableEndpointKafka{value: val, isSet: true}
}

func (v NullableEndpointKafka) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointKafka) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

