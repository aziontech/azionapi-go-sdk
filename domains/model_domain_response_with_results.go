/*
Domain API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package domains

import (
	"encoding/json"
)

// checks if the DomainResponseWithResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainResponseWithResults{}

// DomainResponseWithResults struct for DomainResponseWithResults
type DomainResponseWithResults struct {
	Count int64 `json:"count"`
	TotalPages int64 `json:"total_pages"`
	SchemaVersion int64 `json:"schema_version"`
	Links DomainLinks `json:"links"`
	Results []DomainResults `json:"results"`
}

// NewDomainResponseWithResults instantiates a new DomainResponseWithResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainResponseWithResults(count int64, totalPages int64, schemaVersion int64, links DomainLinks, results []DomainResults) *DomainResponseWithResults {
	this := DomainResponseWithResults{}
	this.Count = count
	this.TotalPages = totalPages
	this.SchemaVersion = schemaVersion
	this.Links = links
	this.Results = results
	return &this
}

// NewDomainResponseWithResultsWithDefaults instantiates a new DomainResponseWithResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainResponseWithResultsWithDefaults() *DomainResponseWithResults {
	this := DomainResponseWithResults{}
	return &this
}

// GetCount returns the Count field value
func (o *DomainResponseWithResults) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *DomainResponseWithResults) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *DomainResponseWithResults) SetCount(v int64) {
	o.Count = v
}

// GetTotalPages returns the TotalPages field value
func (o *DomainResponseWithResults) GetTotalPages() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value
// and a boolean to check if the value has been set.
func (o *DomainResponseWithResults) GetTotalPagesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPages, true
}

// SetTotalPages sets field value
func (o *DomainResponseWithResults) SetTotalPages(v int64) {
	o.TotalPages = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *DomainResponseWithResults) GetSchemaVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *DomainResponseWithResults) GetSchemaVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *DomainResponseWithResults) SetSchemaVersion(v int64) {
	o.SchemaVersion = v
}

// GetLinks returns the Links field value
func (o *DomainResponseWithResults) GetLinks() DomainLinks {
	if o == nil {
		var ret DomainLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *DomainResponseWithResults) GetLinksOk() (*DomainLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *DomainResponseWithResults) SetLinks(v DomainLinks) {
	o.Links = v
}

// GetResults returns the Results field value
func (o *DomainResponseWithResults) GetResults() []DomainResults {
	if o == nil {
		var ret []DomainResults
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *DomainResponseWithResults) GetResultsOk() ([]DomainResults, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *DomainResponseWithResults) SetResults(v []DomainResults) {
	o.Results = v
}

func (o DomainResponseWithResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainResponseWithResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["total_pages"] = o.TotalPages
	toSerialize["schema_version"] = o.SchemaVersion
	toSerialize["links"] = o.Links
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableDomainResponseWithResults struct {
	value *DomainResponseWithResults
	isSet bool
}

func (v NullableDomainResponseWithResults) Get() *DomainResponseWithResults {
	return v.value
}

func (v *NullableDomainResponseWithResults) Set(val *DomainResponseWithResults) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainResponseWithResults) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainResponseWithResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainResponseWithResults(val *DomainResponseWithResults) *NullableDomainResponseWithResults {
	return &NullableDomainResponseWithResults{value: val, isSet: true}
}

func (v NullableDomainResponseWithResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainResponseWithResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


