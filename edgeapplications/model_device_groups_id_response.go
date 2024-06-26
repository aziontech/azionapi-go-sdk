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

// checks if the DeviceGroupsIdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceGroupsIdResponse{}

// DeviceGroupsIdResponse struct for DeviceGroupsIdResponse
type DeviceGroupsIdResponse struct {
	Results DeviceGroupsResultResponse `json:"results"`
	SchemaVersion int64 `json:"schema_version"`
}

type _DeviceGroupsIdResponse DeviceGroupsIdResponse

// NewDeviceGroupsIdResponse instantiates a new DeviceGroupsIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceGroupsIdResponse(results DeviceGroupsResultResponse, schemaVersion int64) *DeviceGroupsIdResponse {
	this := DeviceGroupsIdResponse{}
	this.Results = results
	this.SchemaVersion = schemaVersion
	return &this
}

// NewDeviceGroupsIdResponseWithDefaults instantiates a new DeviceGroupsIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceGroupsIdResponseWithDefaults() *DeviceGroupsIdResponse {
	this := DeviceGroupsIdResponse{}
	return &this
}

// GetResults returns the Results field value
func (o *DeviceGroupsIdResponse) GetResults() DeviceGroupsResultResponse {
	if o == nil {
		var ret DeviceGroupsResultResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupsIdResponse) GetResultsOk() (*DeviceGroupsResultResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Results, true
}

// SetResults sets field value
func (o *DeviceGroupsIdResponse) SetResults(v DeviceGroupsResultResponse) {
	o.Results = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *DeviceGroupsIdResponse) GetSchemaVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupsIdResponse) GetSchemaVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *DeviceGroupsIdResponse) SetSchemaVersion(v int64) {
	o.SchemaVersion = v
}

func (o DeviceGroupsIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceGroupsIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	toSerialize["schema_version"] = o.SchemaVersion
	return toSerialize, nil
}

func (o *DeviceGroupsIdResponse) UnmarshalJSON(data []byte) (err error) {
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

	varDeviceGroupsIdResponse := _DeviceGroupsIdResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceGroupsIdResponse)

	if err != nil {
		return err
	}

	*o = DeviceGroupsIdResponse(varDeviceGroupsIdResponse)

	return err
}

type NullableDeviceGroupsIdResponse struct {
	value *DeviceGroupsIdResponse
	isSet bool
}

func (v NullableDeviceGroupsIdResponse) Get() *DeviceGroupsIdResponse {
	return v.value
}

func (v *NullableDeviceGroupsIdResponse) Set(val *DeviceGroupsIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceGroupsIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceGroupsIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceGroupsIdResponse(val *DeviceGroupsIdResponse) *NullableDeviceGroupsIdResponse {
	return &NullableDeviceGroupsIdResponse{value: val, isSet: true}
}

func (v NullableDeviceGroupsIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceGroupsIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


