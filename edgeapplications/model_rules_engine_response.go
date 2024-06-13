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

// checks if the RulesEngineResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulesEngineResponse{}

// RulesEngineResponse struct for RulesEngineResponse
type RulesEngineResponse struct {
	Count int64 `json:"count"`
	TotalPages int64 `json:"total_pages"`
	SchemaVersion int64 `json:"schema_version"`
	Links OriginsResponseLinks `json:"links"`
	Results []RulesEngineResultResponse `json:"results"`
}

type _RulesEngineResponse RulesEngineResponse

// NewRulesEngineResponse instantiates a new RulesEngineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesEngineResponse(count int64, totalPages int64, schemaVersion int64, links OriginsResponseLinks, results []RulesEngineResultResponse) *RulesEngineResponse {
	this := RulesEngineResponse{}
	this.Count = count
	this.TotalPages = totalPages
	this.SchemaVersion = schemaVersion
	this.Links = links
	this.Results = results
	return &this
}

// NewRulesEngineResponseWithDefaults instantiates a new RulesEngineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesEngineResponseWithDefaults() *RulesEngineResponse {
	this := RulesEngineResponse{}
	return &this
}

// GetCount returns the Count field value
func (o *RulesEngineResponse) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResponse) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *RulesEngineResponse) SetCount(v int64) {
	o.Count = v
}

// GetTotalPages returns the TotalPages field value
func (o *RulesEngineResponse) GetTotalPages() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResponse) GetTotalPagesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPages, true
}

// SetTotalPages sets field value
func (o *RulesEngineResponse) SetTotalPages(v int64) {
	o.TotalPages = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *RulesEngineResponse) GetSchemaVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResponse) GetSchemaVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *RulesEngineResponse) SetSchemaVersion(v int64) {
	o.SchemaVersion = v
}

// GetLinks returns the Links field value
func (o *RulesEngineResponse) GetLinks() OriginsResponseLinks {
	if o == nil {
		var ret OriginsResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResponse) GetLinksOk() (*OriginsResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RulesEngineResponse) SetLinks(v OriginsResponseLinks) {
	o.Links = v
}

// GetResults returns the Results field value
func (o *RulesEngineResponse) GetResults() []RulesEngineResultResponse {
	if o == nil {
		var ret []RulesEngineResultResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResponse) GetResultsOk() ([]RulesEngineResultResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *RulesEngineResponse) SetResults(v []RulesEngineResultResponse) {
	o.Results = v
}

func (o RulesEngineResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulesEngineResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["total_pages"] = o.TotalPages
	toSerialize["schema_version"] = o.SchemaVersion
	toSerialize["links"] = o.Links
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *RulesEngineResponse) UnmarshalJSON(data []byte) (err error) {
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

	varRulesEngineResponse := _RulesEngineResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRulesEngineResponse)

	if err != nil {
		return err
	}

	*o = RulesEngineResponse(varRulesEngineResponse)

	return err
}

type NullableRulesEngineResponse struct {
	value *RulesEngineResponse
	isSet bool
}

func (v NullableRulesEngineResponse) Get() *RulesEngineResponse {
	return v.value
}

func (v *NullableRulesEngineResponse) Set(val *RulesEngineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesEngineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesEngineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesEngineResponse(val *RulesEngineResponse) *NullableRulesEngineResponse {
	return &NullableRulesEngineResponse{value: val, isSet: true}
}

func (v NullableRulesEngineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesEngineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


