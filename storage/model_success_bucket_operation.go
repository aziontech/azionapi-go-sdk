/*
Object Storage

REST API OpenAPI documentation for the Object Storage

API version: 1.0.0 (v1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SuccessBucketOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuccessBucketOperation{}

// SuccessBucketOperation struct for SuccessBucketOperation
type SuccessBucketOperation struct {
	State StateEnum `json:"state"`
	Data Bucket `json:"data"`
}

type _SuccessBucketOperation SuccessBucketOperation

// NewSuccessBucketOperation instantiates a new SuccessBucketOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessBucketOperation(state StateEnum, data Bucket) *SuccessBucketOperation {
	this := SuccessBucketOperation{}
	this.State = state
	this.Data = data
	return &this
}

// NewSuccessBucketOperationWithDefaults instantiates a new SuccessBucketOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessBucketOperationWithDefaults() *SuccessBucketOperation {
	this := SuccessBucketOperation{}
	return &this
}

// GetState returns the State field value
func (o *SuccessBucketOperation) GetState() StateEnum {
	if o == nil {
		var ret StateEnum
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SuccessBucketOperation) GetStateOk() (*StateEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SuccessBucketOperation) SetState(v StateEnum) {
	o.State = v
}

// GetData returns the Data field value
func (o *SuccessBucketOperation) GetData() Bucket {
	if o == nil {
		var ret Bucket
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SuccessBucketOperation) GetDataOk() (*Bucket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SuccessBucketOperation) SetData(v Bucket) {
	o.Data = v
}

func (o SuccessBucketOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuccessBucketOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SuccessBucketOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSuccessBucketOperation := _SuccessBucketOperation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuccessBucketOperation)

	if err != nil {
		return err
	}

	*o = SuccessBucketOperation(varSuccessBucketOperation)

	return err
}

type NullableSuccessBucketOperation struct {
	value *SuccessBucketOperation
	isSet bool
}

func (v NullableSuccessBucketOperation) Get() *SuccessBucketOperation {
	return v.value
}

func (v *NullableSuccessBucketOperation) Set(val *SuccessBucketOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessBucketOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessBucketOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessBucketOperation(val *SuccessBucketOperation) *NullableSuccessBucketOperation {
	return &NullableSuccessBucketOperation{value: val, isSet: true}
}

func (v NullableSuccessBucketOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessBucketOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


