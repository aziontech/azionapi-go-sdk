/*
Edge Firewall API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgefirewall

import (
	"encoding/json"
)

// checks if the SetCustomResponseArgument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetCustomResponseArgument{}

// SetCustomResponseArgument struct for SetCustomResponseArgument
type SetCustomResponseArgument struct {
	StatusCode SetCustomResponseArgumentStatusCode `json:"status_code"`
	ContentType string `json:"content_type"`
	ContentBody string `json:"content_body"`
}

// NewSetCustomResponseArgument instantiates a new SetCustomResponseArgument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetCustomResponseArgument(statusCode SetCustomResponseArgumentStatusCode, contentType string, contentBody string) *SetCustomResponseArgument {
	this := SetCustomResponseArgument{}
	this.StatusCode = statusCode
	this.ContentType = contentType
	this.ContentBody = contentBody
	return &this
}

// NewSetCustomResponseArgumentWithDefaults instantiates a new SetCustomResponseArgument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetCustomResponseArgumentWithDefaults() *SetCustomResponseArgument {
	this := SetCustomResponseArgument{}
	return &this
}

// GetStatusCode returns the StatusCode field value
func (o *SetCustomResponseArgument) GetStatusCode() SetCustomResponseArgumentStatusCode {
	if o == nil {
		var ret SetCustomResponseArgumentStatusCode
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *SetCustomResponseArgument) GetStatusCodeOk() (*SetCustomResponseArgumentStatusCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *SetCustomResponseArgument) SetStatusCode(v SetCustomResponseArgumentStatusCode) {
	o.StatusCode = v
}

// GetContentType returns the ContentType field value
func (o *SetCustomResponseArgument) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *SetCustomResponseArgument) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *SetCustomResponseArgument) SetContentType(v string) {
	o.ContentType = v
}

// GetContentBody returns the ContentBody field value
func (o *SetCustomResponseArgument) GetContentBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentBody
}

// GetContentBodyOk returns a tuple with the ContentBody field value
// and a boolean to check if the value has been set.
func (o *SetCustomResponseArgument) GetContentBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentBody, true
}

// SetContentBody sets field value
func (o *SetCustomResponseArgument) SetContentBody(v string) {
	o.ContentBody = v
}

func (o SetCustomResponseArgument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetCustomResponseArgument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status_code"] = o.StatusCode
	toSerialize["content_type"] = o.ContentType
	toSerialize["content_body"] = o.ContentBody
	return toSerialize, nil
}

type NullableSetCustomResponseArgument struct {
	value *SetCustomResponseArgument
	isSet bool
}

func (v NullableSetCustomResponseArgument) Get() *SetCustomResponseArgument {
	return v.value
}

func (v *NullableSetCustomResponseArgument) Set(val *SetCustomResponseArgument) {
	v.value = val
	v.isSet = true
}

func (v NullableSetCustomResponseArgument) IsSet() bool {
	return v.isSet
}

func (v *NullableSetCustomResponseArgument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetCustomResponseArgument(val *SetCustomResponseArgument) *NullableSetCustomResponseArgument {
	return &NullableSetCustomResponseArgument{value: val, isSet: true}
}

func (v NullableSetCustomResponseArgument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetCustomResponseArgument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


