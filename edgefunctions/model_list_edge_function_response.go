/*
Edge Function

Azion Edge Function API

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgefunctions

import (
	"encoding/json"
)

// ListEdgeFunctionResponse struct for ListEdgeFunctionResponse
type ListEdgeFunctionResponse struct {
	Count *int64 `json:"count,omitempty"`
	TotalPages *int64 `json:"total_pages,omitempty"`
	SchemaVersion *int64 `json:"schema_version,omitempty"`
	Links *Links `json:"links,omitempty"`
	Results *[]Results `json:"results,omitempty"`
}

// NewListEdgeFunctionResponse instantiates a new ListEdgeFunctionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEdgeFunctionResponse() *ListEdgeFunctionResponse {
	this := ListEdgeFunctionResponse{}
	return &this
}

// NewListEdgeFunctionResponseWithDefaults instantiates a new ListEdgeFunctionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEdgeFunctionResponseWithDefaults() *ListEdgeFunctionResponse {
	this := ListEdgeFunctionResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListEdgeFunctionResponse) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEdgeFunctionResponse) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListEdgeFunctionResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ListEdgeFunctionResponse) SetCount(v int64) {
	o.Count = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *ListEdgeFunctionResponse) GetTotalPages() int64 {
	if o == nil || o.TotalPages == nil {
		var ret int64
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEdgeFunctionResponse) GetTotalPagesOk() (*int64, bool) {
	if o == nil || o.TotalPages == nil {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *ListEdgeFunctionResponse) HasTotalPages() bool {
	if o != nil && o.TotalPages != nil {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int64 and assigns it to the TotalPages field.
func (o *ListEdgeFunctionResponse) SetTotalPages(v int64) {
	o.TotalPages = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *ListEdgeFunctionResponse) GetSchemaVersion() int64 {
	if o == nil || o.SchemaVersion == nil {
		var ret int64
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEdgeFunctionResponse) GetSchemaVersionOk() (*int64, bool) {
	if o == nil || o.SchemaVersion == nil {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *ListEdgeFunctionResponse) HasSchemaVersion() bool {
	if o != nil && o.SchemaVersion != nil {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given int64 and assigns it to the SchemaVersion field.
func (o *ListEdgeFunctionResponse) SetSchemaVersion(v int64) {
	o.SchemaVersion = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListEdgeFunctionResponse) GetLinks() Links {
	if o == nil || o.Links == nil {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEdgeFunctionResponse) GetLinksOk() (*Links, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListEdgeFunctionResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *ListEdgeFunctionResponse) SetLinks(v Links) {
	o.Links = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListEdgeFunctionResponse) GetResults() []Results {
	if o == nil || o.Results == nil {
		var ret []Results
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEdgeFunctionResponse) GetResultsOk() (*[]Results, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListEdgeFunctionResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Results and assigns it to the Results field.
func (o *ListEdgeFunctionResponse) SetResults(v []Results) {
	o.Results = &v
}

func (o ListEdgeFunctionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.TotalPages != nil {
		toSerialize["total_pages"] = o.TotalPages
	}
	if o.SchemaVersion != nil {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableListEdgeFunctionResponse struct {
	value *ListEdgeFunctionResponse
	isSet bool
}

func (v NullableListEdgeFunctionResponse) Get() *ListEdgeFunctionResponse {
	return v.value
}

func (v *NullableListEdgeFunctionResponse) Set(val *ListEdgeFunctionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListEdgeFunctionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListEdgeFunctionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEdgeFunctionResponse(val *ListEdgeFunctionResponse) *NullableListEdgeFunctionResponse {
	return &NullableListEdgeFunctionResponse{value: val, isSet: true}
}

func (v NullableListEdgeFunctionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEdgeFunctionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


