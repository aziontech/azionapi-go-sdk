/*
Personal Tokens - OpenAPI

The Personal Tokens API allows you to manage your existing personal tokens. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package personal_tokens

import (
	"encoding/json"
)

// checks if the PersonalTokenResponseWithResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonalTokenResponseWithResults{}

// PersonalTokenResponseWithResults struct for PersonalTokenResponseWithResults
type PersonalTokenResponseWithResults struct {
	Count *int64 `json:"count,omitempty"`
	TotalPages *int32 `json:"total_pages,omitempty"`
	SchemaVersion *int64 `json:"schema_version,omitempty"`
	Links *PersonalTokenResponseWithResultsLinks `json:"links,omitempty"`
	Results []PersonalTokenResponseGet `json:"results,omitempty"`
}

// NewPersonalTokenResponseWithResults instantiates a new PersonalTokenResponseWithResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonalTokenResponseWithResults() *PersonalTokenResponseWithResults {
	this := PersonalTokenResponseWithResults{}
	return &this
}

// NewPersonalTokenResponseWithResultsWithDefaults instantiates a new PersonalTokenResponseWithResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonalTokenResponseWithResultsWithDefaults() *PersonalTokenResponseWithResults {
	this := PersonalTokenResponseWithResults{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PersonalTokenResponseWithResults) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalTokenResponseWithResults) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PersonalTokenResponseWithResults) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *PersonalTokenResponseWithResults) SetCount(v int64) {
	o.Count = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *PersonalTokenResponseWithResults) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalTokenResponseWithResults) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *PersonalTokenResponseWithResults) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *PersonalTokenResponseWithResults) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *PersonalTokenResponseWithResults) GetSchemaVersion() int64 {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret int64
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalTokenResponseWithResults) GetSchemaVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *PersonalTokenResponseWithResults) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given int64 and assigns it to the SchemaVersion field.
func (o *PersonalTokenResponseWithResults) SetSchemaVersion(v int64) {
	o.SchemaVersion = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PersonalTokenResponseWithResults) GetLinks() PersonalTokenResponseWithResultsLinks {
	if o == nil || IsNil(o.Links) {
		var ret PersonalTokenResponseWithResultsLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalTokenResponseWithResults) GetLinksOk() (*PersonalTokenResponseWithResultsLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PersonalTokenResponseWithResults) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PersonalTokenResponseWithResultsLinks and assigns it to the Links field.
func (o *PersonalTokenResponseWithResults) SetLinks(v PersonalTokenResponseWithResultsLinks) {
	o.Links = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PersonalTokenResponseWithResults) GetResults() []PersonalTokenResponseGet {
	if o == nil || IsNil(o.Results) {
		var ret []PersonalTokenResponseGet
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonalTokenResponseWithResults) GetResultsOk() ([]PersonalTokenResponseGet, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PersonalTokenResponseWithResults) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []PersonalTokenResponseGet and assigns it to the Results field.
func (o *PersonalTokenResponseWithResults) SetResults(v []PersonalTokenResponseGet) {
	o.Results = v
}

func (o PersonalTokenResponseWithResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonalTokenResponseWithResults) ToMap() (map[string]interface{}, error) {
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

type NullablePersonalTokenResponseWithResults struct {
	value *PersonalTokenResponseWithResults
	isSet bool
}

func (v NullablePersonalTokenResponseWithResults) Get() *PersonalTokenResponseWithResults {
	return v.value
}

func (v *NullablePersonalTokenResponseWithResults) Set(val *PersonalTokenResponseWithResults) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonalTokenResponseWithResults) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonalTokenResponseWithResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonalTokenResponseWithResults(val *PersonalTokenResponseWithResults) *NullablePersonalTokenResponseWithResults {
	return &NullablePersonalTokenResponseWithResults{value: val, isSet: true}
}

func (v NullablePersonalTokenResponseWithResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonalTokenResponseWithResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

