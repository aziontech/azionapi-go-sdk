/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
	"fmt"
)

// checks if the ApplicationInstancesGetOneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationInstancesGetOneResponse{}

// ApplicationInstancesGetOneResponse struct for ApplicationInstancesGetOneResponse
type ApplicationInstancesGetOneResponse struct {
	Results ApplicationInstancesResults `json:"results"`
	SchemaVersion int64 `json:"schema_version"`
}

type _ApplicationInstancesGetOneResponse ApplicationInstancesGetOneResponse

// NewApplicationInstancesGetOneResponse instantiates a new ApplicationInstancesGetOneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationInstancesGetOneResponse(results ApplicationInstancesResults, schemaVersion int64) *ApplicationInstancesGetOneResponse {
	this := ApplicationInstancesGetOneResponse{}
	this.Results = results
	this.SchemaVersion = schemaVersion
	return &this
}

// NewApplicationInstancesGetOneResponseWithDefaults instantiates a new ApplicationInstancesGetOneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationInstancesGetOneResponseWithDefaults() *ApplicationInstancesGetOneResponse {
	this := ApplicationInstancesGetOneResponse{}
	return &this
}

// GetResults returns the Results field value
func (o *ApplicationInstancesGetOneResponse) GetResults() ApplicationInstancesResults {
	if o == nil {
		var ret ApplicationInstancesResults
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ApplicationInstancesGetOneResponse) GetResultsOk() (*ApplicationInstancesResults, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *ApplicationInstancesGetOneResponse) SetResults(v ApplicationInstancesResults) {
	o.Results = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *ApplicationInstancesGetOneResponse) GetSchemaVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *ApplicationInstancesGetOneResponse) GetSchemaVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *ApplicationInstancesGetOneResponse) SetSchemaVersion(v int64) {
	o.SchemaVersion = v
}

func (o ApplicationInstancesGetOneResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationInstancesGetOneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	toSerialize["schema_version"] = o.SchemaVersion
	return toSerialize, nil
}

func (o *ApplicationInstancesGetOneResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
		"schema_version",
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

	varApplicationInstancesGetOneResponse := _ApplicationInstancesGetOneResponse{}

	err = json.Unmarshal(bytes, &varApplicationInstancesGetOneResponse)

	if err != nil {
		return err
	}

	*o = ApplicationInstancesGetOneResponse(varApplicationInstancesGetOneResponse)

	return err
}

type NullableApplicationInstancesGetOneResponse struct {
	value *ApplicationInstancesGetOneResponse
	isSet bool
}

func (v NullableApplicationInstancesGetOneResponse) Get() *ApplicationInstancesGetOneResponse {
	return v.value
}

func (v *NullableApplicationInstancesGetOneResponse) Set(val *ApplicationInstancesGetOneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationInstancesGetOneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationInstancesGetOneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationInstancesGetOneResponse(val *ApplicationInstancesGetOneResponse) *NullableApplicationInstancesGetOneResponse {
	return &NullableApplicationInstancesGetOneResponse{value: val, isSet: true}
}

func (v NullableApplicationInstancesGetOneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationInstancesGetOneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


