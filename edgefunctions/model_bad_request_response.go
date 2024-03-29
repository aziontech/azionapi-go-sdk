/*
Edge Function API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgefunctions

import (
	"encoding/json"
)

// checks if the BadRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BadRequestResponse{}

// BadRequestResponse struct for BadRequestResponse
type BadRequestResponse struct {
	Name []string `json:"name,omitempty"`
	Active []string `json:"active,omitempty"`
	Code []string `json:"code,omitempty"`
	Language []string `json:"language,omitempty"`
}

// NewBadRequestResponse instantiates a new BadRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequestResponse() *BadRequestResponse {
	this := BadRequestResponse{}
	return &this
}

// NewBadRequestResponseWithDefaults instantiates a new BadRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestResponseWithDefaults() *BadRequestResponse {
	this := BadRequestResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BadRequestResponse) GetName() []string {
	if o == nil || IsNil(o.Name) {
		var ret []string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestResponse) GetNameOk() ([]string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BadRequestResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *BadRequestResponse) SetName(v []string) {
	o.Name = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *BadRequestResponse) GetActive() []string {
	if o == nil || IsNil(o.Active) {
		var ret []string
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestResponse) GetActiveOk() ([]string, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *BadRequestResponse) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given []string and assigns it to the Active field.
func (o *BadRequestResponse) SetActive(v []string) {
	o.Active = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *BadRequestResponse) GetCode() []string {
	if o == nil || IsNil(o.Code) {
		var ret []string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestResponse) GetCodeOk() ([]string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *BadRequestResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given []string and assigns it to the Code field.
func (o *BadRequestResponse) SetCode(v []string) {
	o.Code = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *BadRequestResponse) GetLanguage() []string {
	if o == nil || IsNil(o.Language) {
		var ret []string
		return ret
	}
	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestResponse) GetLanguageOk() ([]string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *BadRequestResponse) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given []string and assigns it to the Language field.
func (o *BadRequestResponse) SetLanguage(v []string) {
	o.Language = v
}

func (o BadRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BadRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	return toSerialize, nil
}

type NullableBadRequestResponse struct {
	value *BadRequestResponse
	isSet bool
}

func (v NullableBadRequestResponse) Get() *BadRequestResponse {
	return v.value
}

func (v *NullableBadRequestResponse) Set(val *BadRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestResponse(val *BadRequestResponse) *NullableBadRequestResponse {
	return &NullableBadRequestResponse{value: val, isSet: true}
}

func (v NullableBadRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


