/*
Data Streaming - OpenAPI

The Data Streaming API allows you to manage your existing data streamings and templates. Data Streaming allows you to feed your stream processing, SIEM, and big data platforms with the event logs from your applications on Azion in real time. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data_streaming

import (
	"encoding/json"
)

// checks if the DataStreamingsDomainResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataStreamingsDomainResponse{}

// DataStreamingsDomainResponse struct for DataStreamingsDomainResponse
type DataStreamingsDomainResponse struct {
	Count *int32 `json:"count,omitempty"`
	TotalPages *int32 `json:"total_pages,omitempty"`
	SchemaVersion *float32 `json:"schema_version,omitempty"`
	Links *DataStreamingsDomainResponseLinks `json:"links,omitempty"`
	Results []DataStreamingsDomainResult `json:"results,omitempty"`
}

// NewDataStreamingsDomainResponse instantiates a new DataStreamingsDomainResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataStreamingsDomainResponse() *DataStreamingsDomainResponse {
	this := DataStreamingsDomainResponse{}
	return &this
}

// NewDataStreamingsDomainResponseWithDefaults instantiates a new DataStreamingsDomainResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataStreamingsDomainResponseWithDefaults() *DataStreamingsDomainResponse {
	this := DataStreamingsDomainResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *DataStreamingsDomainResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingsDomainResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *DataStreamingsDomainResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *DataStreamingsDomainResponse) SetCount(v int32) {
	o.Count = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *DataStreamingsDomainResponse) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingsDomainResponse) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *DataStreamingsDomainResponse) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *DataStreamingsDomainResponse) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *DataStreamingsDomainResponse) GetSchemaVersion() float32 {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret float32
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingsDomainResponse) GetSchemaVersionOk() (*float32, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *DataStreamingsDomainResponse) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given float32 and assigns it to the SchemaVersion field.
func (o *DataStreamingsDomainResponse) SetSchemaVersion(v float32) {
	o.SchemaVersion = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DataStreamingsDomainResponse) GetLinks() DataStreamingsDomainResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret DataStreamingsDomainResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingsDomainResponse) GetLinksOk() (*DataStreamingsDomainResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DataStreamingsDomainResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given DataStreamingsDomainResponseLinks and assigns it to the Links field.
func (o *DataStreamingsDomainResponse) SetLinks(v DataStreamingsDomainResponseLinks) {
	o.Links = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DataStreamingsDomainResponse) GetResults() []DataStreamingsDomainResult {
	if o == nil || IsNil(o.Results) {
		var ret []DataStreamingsDomainResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingsDomainResponse) GetResultsOk() ([]DataStreamingsDomainResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DataStreamingsDomainResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DataStreamingsDomainResult and assigns it to the Results field.
func (o *DataStreamingsDomainResponse) SetResults(v []DataStreamingsDomainResult) {
	o.Results = v
}

func (o DataStreamingsDomainResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataStreamingsDomainResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableDataStreamingsDomainResponse struct {
	value *DataStreamingsDomainResponse
	isSet bool
}

func (v NullableDataStreamingsDomainResponse) Get() *DataStreamingsDomainResponse {
	return v.value
}

func (v *NullableDataStreamingsDomainResponse) Set(val *DataStreamingsDomainResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStreamingsDomainResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStreamingsDomainResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStreamingsDomainResponse(val *DataStreamingsDomainResponse) *NullableDataStreamingsDomainResponse {
	return &NullableDataStreamingsDomainResponse{value: val, isSet: true}
}

func (v NullableDataStreamingsDomainResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStreamingsDomainResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


