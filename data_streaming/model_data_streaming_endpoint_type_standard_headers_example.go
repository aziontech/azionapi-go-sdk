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

// checks if the DataStreamingEndpointTypeStandardHeadersExample type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataStreamingEndpointTypeStandardHeadersExample{}

// DataStreamingEndpointTypeStandardHeadersExample struct for DataStreamingEndpointTypeStandardHeadersExample
type DataStreamingEndpointTypeStandardHeadersExample struct {
	HeaderName1 NullableString `json:"header-name-1,omitempty"`
	HeaderName2 NullableString `json:"header-name-2,omitempty"`
	HeaderName3 NullableString `json:"header-name-3,omitempty"`
}

// NewDataStreamingEndpointTypeStandardHeadersExample instantiates a new DataStreamingEndpointTypeStandardHeadersExample object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataStreamingEndpointTypeStandardHeadersExample() *DataStreamingEndpointTypeStandardHeadersExample {
	this := DataStreamingEndpointTypeStandardHeadersExample{}
	return &this
}

// NewDataStreamingEndpointTypeStandardHeadersExampleWithDefaults instantiates a new DataStreamingEndpointTypeStandardHeadersExample object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataStreamingEndpointTypeStandardHeadersExampleWithDefaults() *DataStreamingEndpointTypeStandardHeadersExample {
	this := DataStreamingEndpointTypeStandardHeadersExample{}
	return &this
}

// GetHeaderName1 returns the HeaderName1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataStreamingEndpointTypeStandardHeadersExample) GetHeaderName1() string {
	if o == nil || IsNil(o.HeaderName1.Get()) {
		var ret string
		return ret
	}
	return *o.HeaderName1.Get()
}

// GetHeaderName1Ok returns a tuple with the HeaderName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataStreamingEndpointTypeStandardHeadersExample) GetHeaderName1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeaderName1.Get(), o.HeaderName1.IsSet()
}

// HasHeaderName1 returns a boolean if a field has been set.
func (o *DataStreamingEndpointTypeStandardHeadersExample) HasHeaderName1() bool {
	if o != nil && o.HeaderName1.IsSet() {
		return true
	}

	return false
}

// SetHeaderName1 gets a reference to the given NullableString and assigns it to the HeaderName1 field.
func (o *DataStreamingEndpointTypeStandardHeadersExample) SetHeaderName1(v string) {
	o.HeaderName1.Set(&v)
}
// SetHeaderName1Nil sets the value for HeaderName1 to be an explicit nil
func (o *DataStreamingEndpointTypeStandardHeadersExample) SetHeaderName1Nil() {
	o.HeaderName1.Set(nil)
}

// UnsetHeaderName1 ensures that no value is present for HeaderName1, not even an explicit nil
func (o *DataStreamingEndpointTypeStandardHeadersExample) UnsetHeaderName1() {
	o.HeaderName1.Unset()
}

// GetHeaderName2 returns the HeaderName2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataStreamingEndpointTypeStandardHeadersExample) GetHeaderName2() string {
	if o == nil || IsNil(o.HeaderName2.Get()) {
		var ret string
		return ret
	}
	return *o.HeaderName2.Get()
}

// GetHeaderName2Ok returns a tuple with the HeaderName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataStreamingEndpointTypeStandardHeadersExample) GetHeaderName2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeaderName2.Get(), o.HeaderName2.IsSet()
}

// HasHeaderName2 returns a boolean if a field has been set.
func (o *DataStreamingEndpointTypeStandardHeadersExample) HasHeaderName2() bool {
	if o != nil && o.HeaderName2.IsSet() {
		return true
	}

	return false
}

// SetHeaderName2 gets a reference to the given NullableString and assigns it to the HeaderName2 field.
func (o *DataStreamingEndpointTypeStandardHeadersExample) SetHeaderName2(v string) {
	o.HeaderName2.Set(&v)
}
// SetHeaderName2Nil sets the value for HeaderName2 to be an explicit nil
func (o *DataStreamingEndpointTypeStandardHeadersExample) SetHeaderName2Nil() {
	o.HeaderName2.Set(nil)
}

// UnsetHeaderName2 ensures that no value is present for HeaderName2, not even an explicit nil
func (o *DataStreamingEndpointTypeStandardHeadersExample) UnsetHeaderName2() {
	o.HeaderName2.Unset()
}

// GetHeaderName3 returns the HeaderName3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataStreamingEndpointTypeStandardHeadersExample) GetHeaderName3() string {
	if o == nil || IsNil(o.HeaderName3.Get()) {
		var ret string
		return ret
	}
	return *o.HeaderName3.Get()
}

// GetHeaderName3Ok returns a tuple with the HeaderName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataStreamingEndpointTypeStandardHeadersExample) GetHeaderName3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeaderName3.Get(), o.HeaderName3.IsSet()
}

// HasHeaderName3 returns a boolean if a field has been set.
func (o *DataStreamingEndpointTypeStandardHeadersExample) HasHeaderName3() bool {
	if o != nil && o.HeaderName3.IsSet() {
		return true
	}

	return false
}

// SetHeaderName3 gets a reference to the given NullableString and assigns it to the HeaderName3 field.
func (o *DataStreamingEndpointTypeStandardHeadersExample) SetHeaderName3(v string) {
	o.HeaderName3.Set(&v)
}
// SetHeaderName3Nil sets the value for HeaderName3 to be an explicit nil
func (o *DataStreamingEndpointTypeStandardHeadersExample) SetHeaderName3Nil() {
	o.HeaderName3.Set(nil)
}

// UnsetHeaderName3 ensures that no value is present for HeaderName3, not even an explicit nil
func (o *DataStreamingEndpointTypeStandardHeadersExample) UnsetHeaderName3() {
	o.HeaderName3.Unset()
}

func (o DataStreamingEndpointTypeStandardHeadersExample) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataStreamingEndpointTypeStandardHeadersExample) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.HeaderName1.IsSet() {
		toSerialize["header-name-1"] = o.HeaderName1.Get()
	}
	if o.HeaderName2.IsSet() {
		toSerialize["header-name-2"] = o.HeaderName2.Get()
	}
	if o.HeaderName3.IsSet() {
		toSerialize["header-name-3"] = o.HeaderName3.Get()
	}
	return toSerialize, nil
}

type NullableDataStreamingEndpointTypeStandardHeadersExample struct {
	value *DataStreamingEndpointTypeStandardHeadersExample
	isSet bool
}

func (v NullableDataStreamingEndpointTypeStandardHeadersExample) Get() *DataStreamingEndpointTypeStandardHeadersExample {
	return v.value
}

func (v *NullableDataStreamingEndpointTypeStandardHeadersExample) Set(val *DataStreamingEndpointTypeStandardHeadersExample) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStreamingEndpointTypeStandardHeadersExample) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStreamingEndpointTypeStandardHeadersExample) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStreamingEndpointTypeStandardHeadersExample(val *DataStreamingEndpointTypeStandardHeadersExample) *NullableDataStreamingEndpointTypeStandardHeadersExample {
	return &NullableDataStreamingEndpointTypeStandardHeadersExample{value: val, isSet: true}
}

func (v NullableDataStreamingEndpointTypeStandardHeadersExample) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStreamingEndpointTypeStandardHeadersExample) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

