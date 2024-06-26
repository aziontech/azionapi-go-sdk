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

// checks if the ApplicationInstancesGetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationInstancesGetResponse{}

// ApplicationInstancesGetResponse struct for ApplicationInstancesGetResponse
type ApplicationInstancesGetResponse struct {
	Count int64 `json:"count"`
	TotalPages int64 `json:"total_pages"`
	SchemaVersion int64 `json:"schema_version"`
	Links ApplicationLinks `json:"links"`
	Results []ApplicationInstancesResults `json:"results"`
}

type _ApplicationInstancesGetResponse ApplicationInstancesGetResponse

// NewApplicationInstancesGetResponse instantiates a new ApplicationInstancesGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationInstancesGetResponse(count int64, totalPages int64, schemaVersion int64, links ApplicationLinks, results []ApplicationInstancesResults) *ApplicationInstancesGetResponse {
	this := ApplicationInstancesGetResponse{}
	this.Count = count
	this.TotalPages = totalPages
	this.SchemaVersion = schemaVersion
	this.Links = links
	this.Results = results
	return &this
}

// NewApplicationInstancesGetResponseWithDefaults instantiates a new ApplicationInstancesGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationInstancesGetResponseWithDefaults() *ApplicationInstancesGetResponse {
	this := ApplicationInstancesGetResponse{}
	return &this
}

// GetCount returns the Count field value
func (o *ApplicationInstancesGetResponse) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ApplicationInstancesGetResponse) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ApplicationInstancesGetResponse) SetCount(v int64) {
	o.Count = v
}

// GetTotalPages returns the TotalPages field value
func (o *ApplicationInstancesGetResponse) GetTotalPages() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value
// and a boolean to check if the value has been set.
func (o *ApplicationInstancesGetResponse) GetTotalPagesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPages, true
}

// SetTotalPages sets field value
func (o *ApplicationInstancesGetResponse) SetTotalPages(v int64) {
	o.TotalPages = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *ApplicationInstancesGetResponse) GetSchemaVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *ApplicationInstancesGetResponse) GetSchemaVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *ApplicationInstancesGetResponse) SetSchemaVersion(v int64) {
	o.SchemaVersion = v
}

// GetLinks returns the Links field value
func (o *ApplicationInstancesGetResponse) GetLinks() ApplicationLinks {
	if o == nil {
		var ret ApplicationLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ApplicationInstancesGetResponse) GetLinksOk() (*ApplicationLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ApplicationInstancesGetResponse) SetLinks(v ApplicationLinks) {
	o.Links = v
}

// GetResults returns the Results field value
func (o *ApplicationInstancesGetResponse) GetResults() []ApplicationInstancesResults {
	if o == nil {
		var ret []ApplicationInstancesResults
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *ApplicationInstancesGetResponse) GetResultsOk() ([]ApplicationInstancesResults, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *ApplicationInstancesGetResponse) SetResults(v []ApplicationInstancesResults) {
	o.Results = v
}

func (o ApplicationInstancesGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationInstancesGetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["total_pages"] = o.TotalPages
	toSerialize["schema_version"] = o.SchemaVersion
	toSerialize["links"] = o.Links
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *ApplicationInstancesGetResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"total_pages",
		"schema_version",
		"links",
		"results",
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

	varApplicationInstancesGetResponse := _ApplicationInstancesGetResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationInstancesGetResponse)

	if err != nil {
		return err
	}

	*o = ApplicationInstancesGetResponse(varApplicationInstancesGetResponse)

	return err
}

type NullableApplicationInstancesGetResponse struct {
	value *ApplicationInstancesGetResponse
	isSet bool
}

func (v NullableApplicationInstancesGetResponse) Get() *ApplicationInstancesGetResponse {
	return v.value
}

func (v *NullableApplicationInstancesGetResponse) Set(val *ApplicationInstancesGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationInstancesGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationInstancesGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationInstancesGetResponse(val *ApplicationInstancesGetResponse) *NullableApplicationInstancesGetResponse {
	return &NullableApplicationInstancesGetResponse{value: val, isSet: true}
}

func (v NullableApplicationInstancesGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationInstancesGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


