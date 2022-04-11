/*
Purge API

Azion Real-Time Purge

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package realtimepurge

import (
	"encoding/json"
)

// PurgeCacheKeyRequest struct for PurgeCacheKeyRequest
type PurgeCacheKeyRequest struct {
	Urls []string `json:"urls"`
	Method string `json:"method"`
	Layer string `json:"layer"`
}

// NewPurgeCacheKeyRequest instantiates a new PurgeCacheKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurgeCacheKeyRequest(urls []string, method string, layer string) *PurgeCacheKeyRequest {
	this := PurgeCacheKeyRequest{}
	this.Urls = urls
	this.Method = method
	this.Layer = layer
	return &this
}

// NewPurgeCacheKeyRequestWithDefaults instantiates a new PurgeCacheKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurgeCacheKeyRequestWithDefaults() *PurgeCacheKeyRequest {
	this := PurgeCacheKeyRequest{}
	return &this
}

// GetUrls returns the Urls field value
func (o *PurgeCacheKeyRequest) GetUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *PurgeCacheKeyRequest) GetUrlsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *PurgeCacheKeyRequest) SetUrls(v []string) {
	o.Urls = v
}

// GetMethod returns the Method field value
func (o *PurgeCacheKeyRequest) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *PurgeCacheKeyRequest) GetMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *PurgeCacheKeyRequest) SetMethod(v string) {
	o.Method = v
}

// GetLayer returns the Layer field value
func (o *PurgeCacheKeyRequest) GetLayer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Layer
}

// GetLayerOk returns a tuple with the Layer field value
// and a boolean to check if the value has been set.
func (o *PurgeCacheKeyRequest) GetLayerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Layer, true
}

// SetLayer sets field value
func (o *PurgeCacheKeyRequest) SetLayer(v string) {
	o.Layer = v
}

func (o PurgeCacheKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["urls"] = o.Urls
	}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["layer"] = o.Layer
	}
	return json.Marshal(toSerialize)
}

type NullablePurgeCacheKeyRequest struct {
	value *PurgeCacheKeyRequest
	isSet bool
}

func (v NullablePurgeCacheKeyRequest) Get() *PurgeCacheKeyRequest {
	return v.value
}

func (v *NullablePurgeCacheKeyRequest) Set(val *PurgeCacheKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePurgeCacheKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePurgeCacheKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurgeCacheKeyRequest(val *PurgeCacheKeyRequest) *NullablePurgeCacheKeyRequest {
	return &NullablePurgeCacheKeyRequest{value: val, isSet: true}
}

func (v NullablePurgeCacheKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurgeCacheKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


