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

// checks if the ObjectResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectResponseData{}

// ObjectResponseData struct for ObjectResponseData
type ObjectResponseData struct {
	ObjectKey string `json:"object_key"`
}

type _ObjectResponseData ObjectResponseData

// NewObjectResponseData instantiates a new ObjectResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectResponseData(objectKey string) *ObjectResponseData {
	this := ObjectResponseData{}
	this.ObjectKey = objectKey
	return &this
}

// NewObjectResponseDataWithDefaults instantiates a new ObjectResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectResponseDataWithDefaults() *ObjectResponseData {
	this := ObjectResponseData{}
	return &this
}

// GetObjectKey returns the ObjectKey field value
func (o *ObjectResponseData) GetObjectKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectKey
}

// GetObjectKeyOk returns a tuple with the ObjectKey field value
// and a boolean to check if the value has been set.
func (o *ObjectResponseData) GetObjectKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectKey, true
}

// SetObjectKey sets field value
func (o *ObjectResponseData) SetObjectKey(v string) {
	o.ObjectKey = v
}

func (o ObjectResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_key"] = o.ObjectKey
	return toSerialize, nil
}

func (o *ObjectResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object_key",
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

	varObjectResponseData := _ObjectResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectResponseData)

	if err != nil {
		return err
	}

	*o = ObjectResponseData(varObjectResponseData)

	return err
}

type NullableObjectResponseData struct {
	value *ObjectResponseData
	isSet bool
}

func (v NullableObjectResponseData) Get() *ObjectResponseData {
	return v.value
}

func (v *NullableObjectResponseData) Set(val *ObjectResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectResponseData(val *ObjectResponseData) *NullableObjectResponseData {
	return &NullableObjectResponseData{value: val, isSet: true}
}

func (v NullableObjectResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

