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

// checks if the WAFEvents400 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WAFEvents400{}

// WAFEvents400 struct for WAFEvents400
type WAFEvents400 struct {
	Errors []map[string]interface{} `json:"errors,omitempty"`
	SchemaVersion *int32 `json:"schema_version,omitempty"`
}

// NewWAFEvents400 instantiates a new WAFEvents400 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWAFEvents400() *WAFEvents400 {
	this := WAFEvents400{}
	return &this
}

// NewWAFEvents400WithDefaults instantiates a new WAFEvents400 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWAFEvents400WithDefaults() *WAFEvents400 {
	this := WAFEvents400{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *WAFEvents400) GetErrors() []map[string]interface{} {
	if o == nil || IsNil(o.Errors) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFEvents400) GetErrorsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *WAFEvents400) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []map[string]interface{} and assigns it to the Errors field.
func (o *WAFEvents400) SetErrors(v []map[string]interface{}) {
	o.Errors = v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *WAFEvents400) GetSchemaVersion() int32 {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret int32
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WAFEvents400) GetSchemaVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *WAFEvents400) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given int32 and assigns it to the SchemaVersion field.
func (o *WAFEvents400) SetSchemaVersion(v int32) {
	o.SchemaVersion = &v
}

func (o WAFEvents400) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WAFEvents400) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	return toSerialize, nil
}

type NullableWAFEvents400 struct {
	value *WAFEvents400
	isSet bool
}

func (v NullableWAFEvents400) Get() *WAFEvents400 {
	return v.value
}

func (v *NullableWAFEvents400) Set(val *WAFEvents400) {
	v.value = val
	v.isSet = true
}

func (v NullableWAFEvents400) IsSet() bool {
	return v.isSet
}

func (v *NullableWAFEvents400) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWAFEvents400(val *WAFEvents400) *NullableWAFEvents400 {
	return &NullableWAFEvents400{value: val, isSet: true}
}

func (v NullableWAFEvents400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWAFEvents400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


