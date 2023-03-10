/*
Intelligent DNS

Azion Intelligent DNS API

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idns

import (
	"encoding/json"
)

// GeneralError struct for GeneralError
type GeneralError struct {
	Detail *string `json:"detail,omitempty"`
}

// NewGeneralError instantiates a new GeneralError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneralError() *GeneralError {
	this := GeneralError{}
	return &this
}

// NewGeneralErrorWithDefaults instantiates a new GeneralError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralErrorWithDefaults() *GeneralError {
	this := GeneralError{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *GeneralError) GetDetail() string {
	if o == nil || isNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeneralError) GetDetailOk() (*string, bool) {
	if o == nil || isNil(o.Detail) {
    return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *GeneralError) HasDetail() bool {
	if o != nil && !isNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *GeneralError) SetDetail(v string) {
	o.Detail = &v
}

func (o GeneralError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return json.Marshal(toSerialize)
}

type NullableGeneralError struct {
	value *GeneralError
	isSet bool
}

func (v NullableGeneralError) Get() *GeneralError {
	return v.value
}

func (v *NullableGeneralError) Set(val *GeneralError) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralError) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralError(val *GeneralError) *NullableGeneralError {
	return &NullableGeneralError{value: val, isSet: true}
}

func (v NullableGeneralError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

