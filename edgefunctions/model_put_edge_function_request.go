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

// checks if the PutEdgeFunctionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutEdgeFunctionRequest{}

// PutEdgeFunctionRequest struct for PutEdgeFunctionRequest
type PutEdgeFunctionRequest struct {
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`
	JsonArgs interface{} `json:"json_args,omitempty"`
	Active *bool `json:"active,omitempty"`
	InitiatorType *string `json:"initiator_type,omitempty"`
	Language *string `json:"language,omitempty"`
	IsProprietaryCode *bool `json:"is_proprietary_code,omitempty"`
}

// NewPutEdgeFunctionRequest instantiates a new PutEdgeFunctionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutEdgeFunctionRequest() *PutEdgeFunctionRequest {
	this := PutEdgeFunctionRequest{}
	return &this
}

// NewPutEdgeFunctionRequestWithDefaults instantiates a new PutEdgeFunctionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutEdgeFunctionRequestWithDefaults() *PutEdgeFunctionRequest {
	this := PutEdgeFunctionRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PutEdgeFunctionRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutEdgeFunctionRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PutEdgeFunctionRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PutEdgeFunctionRequest) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *PutEdgeFunctionRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutEdgeFunctionRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *PutEdgeFunctionRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *PutEdgeFunctionRequest) SetCode(v string) {
	o.Code = &v
}

// GetJsonArgs returns the JsonArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutEdgeFunctionRequest) GetJsonArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.JsonArgs
}

// GetJsonArgsOk returns a tuple with the JsonArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutEdgeFunctionRequest) GetJsonArgsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.JsonArgs) {
		return nil, false
	}
	return &o.JsonArgs, true
}

// HasJsonArgs returns a boolean if a field has been set.
func (o *PutEdgeFunctionRequest) HasJsonArgs() bool {
	if o != nil && !IsNil(o.JsonArgs) {
		return true
	}

	return false
}

// SetJsonArgs gets a reference to the given interface{} and assigns it to the JsonArgs field.
func (o *PutEdgeFunctionRequest) SetJsonArgs(v interface{}) {
	o.JsonArgs = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PutEdgeFunctionRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutEdgeFunctionRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PutEdgeFunctionRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PutEdgeFunctionRequest) SetActive(v bool) {
	o.Active = &v
}

// GetInitiatorType returns the InitiatorType field value if set, zero value otherwise.
func (o *PutEdgeFunctionRequest) GetInitiatorType() string {
	if o == nil || IsNil(o.InitiatorType) {
		var ret string
		return ret
	}
	return *o.InitiatorType
}

// GetInitiatorTypeOk returns a tuple with the InitiatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutEdgeFunctionRequest) GetInitiatorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatorType) {
		return nil, false
	}
	return o.InitiatorType, true
}

// HasInitiatorType returns a boolean if a field has been set.
func (o *PutEdgeFunctionRequest) HasInitiatorType() bool {
	if o != nil && !IsNil(o.InitiatorType) {
		return true
	}

	return false
}

// SetInitiatorType gets a reference to the given string and assigns it to the InitiatorType field.
func (o *PutEdgeFunctionRequest) SetInitiatorType(v string) {
	o.InitiatorType = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *PutEdgeFunctionRequest) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutEdgeFunctionRequest) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *PutEdgeFunctionRequest) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *PutEdgeFunctionRequest) SetLanguage(v string) {
	o.Language = &v
}

// GetIsProprietaryCode returns the IsProprietaryCode field value if set, zero value otherwise.
func (o *PutEdgeFunctionRequest) GetIsProprietaryCode() bool {
	if o == nil || IsNil(o.IsProprietaryCode) {
		var ret bool
		return ret
	}
	return *o.IsProprietaryCode
}

// GetIsProprietaryCodeOk returns a tuple with the IsProprietaryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutEdgeFunctionRequest) GetIsProprietaryCodeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProprietaryCode) {
		return nil, false
	}
	return o.IsProprietaryCode, true
}

// HasIsProprietaryCode returns a boolean if a field has been set.
func (o *PutEdgeFunctionRequest) HasIsProprietaryCode() bool {
	if o != nil && !IsNil(o.IsProprietaryCode) {
		return true
	}

	return false
}

// SetIsProprietaryCode gets a reference to the given bool and assigns it to the IsProprietaryCode field.
func (o *PutEdgeFunctionRequest) SetIsProprietaryCode(v bool) {
	o.IsProprietaryCode = &v
}

func (o PutEdgeFunctionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutEdgeFunctionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if o.JsonArgs != nil {
		toSerialize["json_args"] = o.JsonArgs
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.InitiatorType) {
		toSerialize["initiator_type"] = o.InitiatorType
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.IsProprietaryCode) {
		toSerialize["is_proprietary_code"] = o.IsProprietaryCode
	}
	return toSerialize, nil
}

type NullablePutEdgeFunctionRequest struct {
	value *PutEdgeFunctionRequest
	isSet bool
}

func (v NullablePutEdgeFunctionRequest) Get() *PutEdgeFunctionRequest {
	return v.value
}

func (v *NullablePutEdgeFunctionRequest) Set(val *PutEdgeFunctionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutEdgeFunctionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutEdgeFunctionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutEdgeFunctionRequest(val *PutEdgeFunctionRequest) *NullablePutEdgeFunctionRequest {
	return &NullablePutEdgeFunctionRequest{value: val, isSet: true}
}

func (v NullablePutEdgeFunctionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutEdgeFunctionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


