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

// checks if the EndpointAWSKinesisFirehose type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointAWSKinesisFirehose{}

// EndpointAWSKinesisFirehose struct for EndpointAWSKinesisFirehose
type EndpointAWSKinesisFirehose struct {
	EndpointType *string `json:"endpoint_type,omitempty"`
	AccessKey *string `json:"access_key,omitempty"`
	StreamName *string `json:"stream_name,omitempty"`
	Region *string `json:"region,omitempty"`
	SecretKey *string `json:"secret_key,omitempty"`
}

// NewEndpointAWSKinesisFirehose instantiates a new EndpointAWSKinesisFirehose object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointAWSKinesisFirehose() *EndpointAWSKinesisFirehose {
	this := EndpointAWSKinesisFirehose{}
	return &this
}

// NewEndpointAWSKinesisFirehoseWithDefaults instantiates a new EndpointAWSKinesisFirehose object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointAWSKinesisFirehoseWithDefaults() *EndpointAWSKinesisFirehose {
	this := EndpointAWSKinesisFirehose{}
	return &this
}

// GetEndpointType returns the EndpointType field value if set, zero value otherwise.
func (o *EndpointAWSKinesisFirehose) GetEndpointType() string {
	if o == nil || IsNil(o.EndpointType) {
		var ret string
		return ret
	}
	return *o.EndpointType
}

// GetEndpointTypeOk returns a tuple with the EndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAWSKinesisFirehose) GetEndpointTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointType) {
		return nil, false
	}
	return o.EndpointType, true
}

// HasEndpointType returns a boolean if a field has been set.
func (o *EndpointAWSKinesisFirehose) HasEndpointType() bool {
	if o != nil && !IsNil(o.EndpointType) {
		return true
	}

	return false
}

// SetEndpointType gets a reference to the given string and assigns it to the EndpointType field.
func (o *EndpointAWSKinesisFirehose) SetEndpointType(v string) {
	o.EndpointType = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *EndpointAWSKinesisFirehose) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAWSKinesisFirehose) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *EndpointAWSKinesisFirehose) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *EndpointAWSKinesisFirehose) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetStreamName returns the StreamName field value if set, zero value otherwise.
func (o *EndpointAWSKinesisFirehose) GetStreamName() string {
	if o == nil || IsNil(o.StreamName) {
		var ret string
		return ret
	}
	return *o.StreamName
}

// GetStreamNameOk returns a tuple with the StreamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAWSKinesisFirehose) GetStreamNameOk() (*string, bool) {
	if o == nil || IsNil(o.StreamName) {
		return nil, false
	}
	return o.StreamName, true
}

// HasStreamName returns a boolean if a field has been set.
func (o *EndpointAWSKinesisFirehose) HasStreamName() bool {
	if o != nil && !IsNil(o.StreamName) {
		return true
	}

	return false
}

// SetStreamName gets a reference to the given string and assigns it to the StreamName field.
func (o *EndpointAWSKinesisFirehose) SetStreamName(v string) {
	o.StreamName = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *EndpointAWSKinesisFirehose) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAWSKinesisFirehose) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *EndpointAWSKinesisFirehose) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *EndpointAWSKinesisFirehose) SetRegion(v string) {
	o.Region = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *EndpointAWSKinesisFirehose) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAWSKinesisFirehose) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *EndpointAWSKinesisFirehose) HasSecretKey() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *EndpointAWSKinesisFirehose) SetSecretKey(v string) {
	o.SecretKey = &v
}

func (o EndpointAWSKinesisFirehose) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointAWSKinesisFirehose) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointType) {
		toSerialize["endpoint_type"] = o.EndpointType
	}
	if !IsNil(o.AccessKey) {
		toSerialize["access_key"] = o.AccessKey
	}
	if !IsNil(o.StreamName) {
		toSerialize["stream_name"] = o.StreamName
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.SecretKey) {
		toSerialize["secret_key"] = o.SecretKey
	}
	return toSerialize, nil
}

type NullableEndpointAWSKinesisFirehose struct {
	value *EndpointAWSKinesisFirehose
	isSet bool
}

func (v NullableEndpointAWSKinesisFirehose) Get() *EndpointAWSKinesisFirehose {
	return v.value
}

func (v *NullableEndpointAWSKinesisFirehose) Set(val *EndpointAWSKinesisFirehose) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointAWSKinesisFirehose) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointAWSKinesisFirehose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointAWSKinesisFirehose(val *EndpointAWSKinesisFirehose) *NullableEndpointAWSKinesisFirehose {
	return &NullableEndpointAWSKinesisFirehose{value: val, isSet: true}
}

func (v NullableEndpointAWSKinesisFirehose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointAWSKinesisFirehose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


