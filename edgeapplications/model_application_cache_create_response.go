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

// checks if the ApplicationCacheCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCacheCreateResponse{}

// ApplicationCacheCreateResponse struct for ApplicationCacheCreateResponse
type ApplicationCacheCreateResponse struct {
	Results *ApplicationCacheCreateResults `json:"results,omitempty"`
	SchemaVersion *int64 `json:"schema_version,omitempty"`
}

// NewApplicationCacheCreateResponse instantiates a new ApplicationCacheCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCacheCreateResponse() *ApplicationCacheCreateResponse {
	this := ApplicationCacheCreateResponse{}
	return &this
}

// NewApplicationCacheCreateResponseWithDefaults instantiates a new ApplicationCacheCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCacheCreateResponseWithDefaults() *ApplicationCacheCreateResponse {
	this := ApplicationCacheCreateResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ApplicationCacheCreateResponse) GetResults() ApplicationCacheCreateResults {
	if o == nil || IsNil(o.Results) {
		var ret ApplicationCacheCreateResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateResponse) GetResultsOk() (*ApplicationCacheCreateResults, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ApplicationCacheCreateResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given ApplicationCacheCreateResults and assigns it to the Results field.
func (o *ApplicationCacheCreateResponse) SetResults(v ApplicationCacheCreateResults) {
	o.Results = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *ApplicationCacheCreateResponse) GetSchemaVersion() int64 {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret int64
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateResponse) GetSchemaVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *ApplicationCacheCreateResponse) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given int64 and assigns it to the SchemaVersion field.
func (o *ApplicationCacheCreateResponse) SetSchemaVersion(v int64) {
	o.SchemaVersion = &v
}

func (o ApplicationCacheCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCacheCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	return toSerialize, nil
}

type NullableApplicationCacheCreateResponse struct {
	value *ApplicationCacheCreateResponse
	isSet bool
}

func (v NullableApplicationCacheCreateResponse) Get() *ApplicationCacheCreateResponse {
	return v.value
}

func (v *NullableApplicationCacheCreateResponse) Set(val *ApplicationCacheCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCacheCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCacheCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCacheCreateResponse(val *ApplicationCacheCreateResponse) *NullableApplicationCacheCreateResponse {
	return &NullableApplicationCacheCreateResponse{value: val, isSet: true}
}

func (v NullableApplicationCacheCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCacheCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


