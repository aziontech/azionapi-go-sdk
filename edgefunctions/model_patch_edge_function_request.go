/*
Edge Function

Azion Edge Function API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgefunctions

import (
	"encoding/json"
)

// PatchEdgeFunctionRequest struct for PatchEdgeFunctionRequest
type PatchEdgeFunctionRequest struct {
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`
	JsonArgs *map[string]interface{} `json:"json_args,omitempty"`
	Active *bool `json:"active,omitempty"`
}

// NewPatchEdgeFunctionRequest instantiates a new PatchEdgeFunctionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchEdgeFunctionRequest() *PatchEdgeFunctionRequest {
	this := PatchEdgeFunctionRequest{}
	return &this
}

// NewPatchEdgeFunctionRequestWithDefaults instantiates a new PatchEdgeFunctionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchEdgeFunctionRequestWithDefaults() *PatchEdgeFunctionRequest {
	this := PatchEdgeFunctionRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchEdgeFunctionRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchEdgeFunctionRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchEdgeFunctionRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchEdgeFunctionRequest) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PatchEdgeFunctionRequest) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchEdgeFunctionRequest) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PatchEdgeFunctionRequest) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *PatchEdgeFunctionRequest) SetCode(v string) {
	o.Code = &v
}

// GetJsonArgs returns the JsonArgs field value if set, zero value otherwise.
func (o *PatchEdgeFunctionRequest) GetJsonArgs() map[string]interface{} {
	if o == nil || o.JsonArgs == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.JsonArgs
}

// GetJsonArgsOk returns a tuple with the JsonArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchEdgeFunctionRequest) GetJsonArgsOk() (*map[string]interface{}, bool) {
	if o == nil || o.JsonArgs == nil {
		return nil, false
	}
	return o.JsonArgs, true
}

// HasJsonArgs returns a boolean if a field has been set.
func (o *PatchEdgeFunctionRequest) HasJsonArgs() bool {
	if o != nil && o.JsonArgs != nil {
		return true
	}

	return false
}

// SetJsonArgs gets a reference to the given map[string]interface{} and assigns it to the JsonArgs field.
func (o *PatchEdgeFunctionRequest) SetJsonArgs(v map[string]interface{}) {
	o.JsonArgs = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PatchEdgeFunctionRequest) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchEdgeFunctionRequest) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PatchEdgeFunctionRequest) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PatchEdgeFunctionRequest) SetActive(v bool) {
	o.Active = &v
}

func (o PatchEdgeFunctionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.JsonArgs != nil {
		toSerialize["json_args"] = o.JsonArgs
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	return json.Marshal(toSerialize)
}

type NullablePatchEdgeFunctionRequest struct {
	value *PatchEdgeFunctionRequest
	isSet bool
}

func (v NullablePatchEdgeFunctionRequest) Get() *PatchEdgeFunctionRequest {
	return v.value
}

func (v *NullablePatchEdgeFunctionRequest) Set(val *PatchEdgeFunctionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchEdgeFunctionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchEdgeFunctionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchEdgeFunctionRequest(val *PatchEdgeFunctionRequest) *NullablePatchEdgeFunctionRequest {
	return &NullablePatchEdgeFunctionRequest{value: val, isSet: true}
}

func (v NullablePatchEdgeFunctionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchEdgeFunctionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

