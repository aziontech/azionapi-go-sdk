/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OriginsIdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OriginsIdResponse{}

// OriginsIdResponse struct for OriginsIdResponse
type OriginsIdResponse struct {
	Results OriginsResultResponse `json:"results"`
	SchemaVersion int64 `json:"schema_version"`
}

type _OriginsIdResponse OriginsIdResponse

// NewOriginsIdResponse instantiates a new OriginsIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginsIdResponse(results OriginsResultResponse, schemaVersion int64) *OriginsIdResponse {
	this := OriginsIdResponse{}
	this.Results = results
	this.SchemaVersion = schemaVersion
	return &this
}

// NewOriginsIdResponseWithDefaults instantiates a new OriginsIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginsIdResponseWithDefaults() *OriginsIdResponse {
	this := OriginsIdResponse{}
	return &this
}

// GetResults returns the Results field value
func (o *OriginsIdResponse) GetResults() OriginsResultResponse {
	if o == nil {
		var ret OriginsResultResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *OriginsIdResponse) GetResultsOk() (*OriginsResultResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *OriginsIdResponse) SetResults(v OriginsResultResponse) {
	o.Results = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *OriginsIdResponse) GetSchemaVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *OriginsIdResponse) GetSchemaVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *OriginsIdResponse) SetSchemaVersion(v int64) {
	o.SchemaVersion = v
}

func (o OriginsIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginsIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	toSerialize["schema_version"] = o.SchemaVersion
	return toSerialize, nil
}

func (o *OriginsIdResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
		"schema_version",
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

	varOriginsIdResponse := _OriginsIdResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOriginsIdResponse)

	if err != nil {
		return err
	}

	*o = OriginsIdResponse(varOriginsIdResponse)

	return err
}

type NullableOriginsIdResponse struct {
	value *OriginsIdResponse
	isSet bool
}

func (v NullableOriginsIdResponse) Get() *OriginsIdResponse {
	return v.value
}

func (v *NullableOriginsIdResponse) Set(val *OriginsIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginsIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginsIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginsIdResponse(val *OriginsIdResponse) *NullableOriginsIdResponse {
	return &NullableOriginsIdResponse{value: val, isSet: true}
}

func (v NullableOriginsIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginsIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


