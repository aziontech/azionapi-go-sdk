/*
Web Application Firewall API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package waf

import (
	"encoding/json"
)

// checks if the WAFEvents200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WAFEvents200{}

// WAFEvents200 struct for WAFEvents200
type WAFEvents200 struct {
	Results []WAFEvents200ResultsInner `json:"results,omitempty"`
	SchemaVersion *int64 `json:"schema_version,omitempty"`
}

// NewWAFEvents200 instantiates a new WAFEvents200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWAFEvents200() *WAFEvents200 {
	this := WAFEvents200{}
	return &this
}

// NewWAFEvents200WithDefaults instantiates a new WAFEvents200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWAFEvents200WithDefaults() *WAFEvents200 {
	this := WAFEvents200{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *WAFEvents200) GetResults() []WAFEvents200ResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret []WAFEvents200ResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFEvents200) GetResultsOk() ([]WAFEvents200ResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *WAFEvents200) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []WAFEvents200ResultsInner and assigns it to the Results field.
func (o *WAFEvents200) SetResults(v []WAFEvents200ResultsInner) {
	o.Results = v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *WAFEvents200) GetSchemaVersion() int64 {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret int64
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFEvents200) GetSchemaVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *WAFEvents200) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given int64 and assigns it to the SchemaVersion field.
func (o *WAFEvents200) SetSchemaVersion(v int64) {
	o.SchemaVersion = &v
}

func (o WAFEvents200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WAFEvents200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	return toSerialize, nil
}

type NullableWAFEvents200 struct {
	value *WAFEvents200
	isSet bool
}

func (v NullableWAFEvents200) Get() *WAFEvents200 {
	return v.value
}

func (v *NullableWAFEvents200) Set(val *WAFEvents200) {
	v.value = val
	v.isSet = true
}

func (v NullableWAFEvents200) IsSet() bool {
	return v.isSet
}

func (v *NullableWAFEvents200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWAFEvents200(val *WAFEvents200) *NullableWAFEvents200 {
	return &NullableWAFEvents200{value: val, isSet: true}
}

func (v NullableWAFEvents200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWAFEvents200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


