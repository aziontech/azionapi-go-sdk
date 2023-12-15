/*
Object Storage

REST API OpenAPI documentation for the Object Storage

API version: 1.0.0 (v1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
	"fmt"
)

// checks if the SuccessObjectOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuccessObjectOperation{}

// SuccessObjectOperation struct for SuccessObjectOperation
type SuccessObjectOperation struct {
	State StateEnum `json:"state"`
	Data ObjectResponseData `json:"data"`
}

type _SuccessObjectOperation SuccessObjectOperation

// NewSuccessObjectOperation instantiates a new SuccessObjectOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessObjectOperation(state StateEnum, data ObjectResponseData) *SuccessObjectOperation {
	this := SuccessObjectOperation{}
	this.State = state
	this.Data = data
	return &this
}

// NewSuccessObjectOperationWithDefaults instantiates a new SuccessObjectOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessObjectOperationWithDefaults() *SuccessObjectOperation {
	this := SuccessObjectOperation{}
	return &this
}

// GetState returns the State field value
func (o *SuccessObjectOperation) GetState() StateEnum {
	if o == nil {
		var ret StateEnum
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SuccessObjectOperation) GetStateOk() (*StateEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SuccessObjectOperation) SetState(v StateEnum) {
	o.State = v
}

// GetData returns the Data field value
func (o *SuccessObjectOperation) GetData() ObjectResponseData {
	if o == nil {
		var ret ObjectResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SuccessObjectOperation) GetDataOk() (*ObjectResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SuccessObjectOperation) SetData(v ObjectResponseData) {
	o.Data = v
}

func (o SuccessObjectOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuccessObjectOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SuccessObjectOperation) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSuccessObjectOperation := _SuccessObjectOperation{}

	err = json.Unmarshal(bytes, &varSuccessObjectOperation)

	if err != nil {
		return err
	}

	*o = SuccessObjectOperation(varSuccessObjectOperation)

	return err
}

type NullableSuccessObjectOperation struct {
	value *SuccessObjectOperation
	isSet bool
}

func (v NullableSuccessObjectOperation) Get() *SuccessObjectOperation {
	return v.value
}

func (v *NullableSuccessObjectOperation) Set(val *SuccessObjectOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessObjectOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessObjectOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessObjectOperation(val *SuccessObjectOperation) *NullableSuccessObjectOperation {
	return &NullableSuccessObjectOperation{value: val, isSet: true}
}

func (v NullableSuccessObjectOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessObjectOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


