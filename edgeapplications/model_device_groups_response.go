/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
)

// checks if the DeviceGroupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceGroupsResponse{}

// DeviceGroupsResponse struct for DeviceGroupsResponse
type DeviceGroupsResponse struct {
	Count int64 `json:"count"`
	TotalPages int64 `json:"total_pages"`
	SchemaVersion int64 `json:"schema_version"`
	Links DeviceGroupsResponseLinks `json:"links"`
	Results []DeviceGroupsResultResponse `json:"results"`
}

// NewDeviceGroupsResponse instantiates a new DeviceGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceGroupsResponse(count int64, totalPages int64, schemaVersion int64, links DeviceGroupsResponseLinks, results []DeviceGroupsResultResponse) *DeviceGroupsResponse {
	this := DeviceGroupsResponse{}
	this.Count = count
	this.TotalPages = totalPages
	this.SchemaVersion = schemaVersion
	this.Links = links
	this.Results = results
	return &this
}

// NewDeviceGroupsResponseWithDefaults instantiates a new DeviceGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceGroupsResponseWithDefaults() *DeviceGroupsResponse {
	this := DeviceGroupsResponse{}
	return &this
}

// GetCount returns the Count field value
func (o *DeviceGroupsResponse) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupsResponse) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *DeviceGroupsResponse) SetCount(v int64) {
	o.Count = v
}

// GetTotalPages returns the TotalPages field value
func (o *DeviceGroupsResponse) GetTotalPages() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupsResponse) GetTotalPagesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPages, true
}

// SetTotalPages sets field value
func (o *DeviceGroupsResponse) SetTotalPages(v int64) {
	o.TotalPages = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *DeviceGroupsResponse) GetSchemaVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupsResponse) GetSchemaVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *DeviceGroupsResponse) SetSchemaVersion(v int64) {
	o.SchemaVersion = v
}

// GetLinks returns the Links field value
func (o *DeviceGroupsResponse) GetLinks() DeviceGroupsResponseLinks {
	if o == nil {
		var ret DeviceGroupsResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupsResponse) GetLinksOk() (*DeviceGroupsResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *DeviceGroupsResponse) SetLinks(v DeviceGroupsResponseLinks) {
	o.Links = v
}

// GetResults returns the Results field value
func (o *DeviceGroupsResponse) GetResults() []DeviceGroupsResultResponse {
	if o == nil {
		var ret []DeviceGroupsResultResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *DeviceGroupsResponse) GetResultsOk() ([]DeviceGroupsResultResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *DeviceGroupsResponse) SetResults(v []DeviceGroupsResultResponse) {
	o.Results = v
}

func (o DeviceGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceGroupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["total_pages"] = o.TotalPages
	toSerialize["schema_version"] = o.SchemaVersion
	toSerialize["links"] = o.Links
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableDeviceGroupsResponse struct {
	value *DeviceGroupsResponse
	isSet bool
}

func (v NullableDeviceGroupsResponse) Get() *DeviceGroupsResponse {
	return v.value
}

func (v *NullableDeviceGroupsResponse) Set(val *DeviceGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceGroupsResponse(val *DeviceGroupsResponse) *NullableDeviceGroupsResponse {
	return &NullableDeviceGroupsResponse{value: val, isSet: true}
}

func (v NullableDeviceGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


