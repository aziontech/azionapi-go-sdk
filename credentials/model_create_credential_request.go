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

// checks if the CreateCredentialRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCredentialRequest{}

// CreateCredentialRequest struct for CreateCredentialRequest
type CreateCredentialRequest struct {
	Description string `json:"description"`
	Name string `json:"name"`
	Status bool `json:"status"`
}

// NewCreateCredentialRequest instantiates a new CreateCredentialRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCredentialRequest(description string, name string, status bool) *CreateCredentialRequest {
	this := CreateCredentialRequest{}
	this.Description = description
	this.Name = name
	this.Status = status
	return &this
}

// NewCreateCredentialRequestWithDefaults instantiates a new CreateCredentialRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCredentialRequestWithDefaults() *CreateCredentialRequest {
	this := CreateCredentialRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateCredentialRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateCredentialRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateCredentialRequest) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *CreateCredentialRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCredentialRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCredentialRequest) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *CreateCredentialRequest) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateCredentialRequest) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateCredentialRequest) SetStatus(v bool) {
	o.Status = v
}

func (o CreateCredentialRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCredentialRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableCreateCredentialRequest struct {
	value *CreateCredentialRequest
	isSet bool
}

func (v NullableCreateCredentialRequest) Get() *CreateCredentialRequest {
	return v.value
}

func (v *NullableCreateCredentialRequest) Set(val *CreateCredentialRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCredentialRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCredentialRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCredentialRequest(val *CreateCredentialRequest) *NullableCreateCredentialRequest {
	return &NullableCreateCredentialRequest{value: val, isSet: true}
}

func (v NullableCreateCredentialRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCredentialRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


