/*
Credentials API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
)

// checks if the ResponseWithTotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseWithTotal{}

// ResponseWithTotal struct for ResponseWithTotal
type ResponseWithTotal struct {
	Credentials []Response `json:"credentials"`
	Total int64 `json:"total"`
}

// NewResponseWithTotal instantiates a new ResponseWithTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseWithTotal(credentials []Response, total int64) *ResponseWithTotal {
	this := ResponseWithTotal{}
	this.Credentials = credentials
	this.Total = total
	return &this
}

// NewResponseWithTotalWithDefaults instantiates a new ResponseWithTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseWithTotalWithDefaults() *ResponseWithTotal {
	this := ResponseWithTotal{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *ResponseWithTotal) GetCredentials() []Response {
	if o == nil {
		var ret []Response
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *ResponseWithTotal) GetCredentialsOk() ([]Response, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials, true
}

// SetCredentials sets field value
func (o *ResponseWithTotal) SetCredentials(v []Response) {
	o.Credentials = v
}

// GetTotal returns the Total field value
func (o *ResponseWithTotal) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ResponseWithTotal) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ResponseWithTotal) SetTotal(v int64) {
	o.Total = v
}

func (o ResponseWithTotal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseWithTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentials"] = o.Credentials
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

type NullableResponseWithTotal struct {
	value *ResponseWithTotal
	isSet bool
}

func (v NullableResponseWithTotal) Get() *ResponseWithTotal {
	return v.value
}

func (v *NullableResponseWithTotal) Set(val *ResponseWithTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseWithTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseWithTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseWithTotal(val *ResponseWithTotal) *NullableResponseWithTotal {
	return &NullableResponseWithTotal{value: val, isSet: true}
}

func (v NullableResponseWithTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseWithTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


