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

// checks if the CreateDataStreamingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDataStreamingResponse{}

// CreateDataStreamingResponse struct for CreateDataStreamingResponse
type CreateDataStreamingResponse struct {
	Results []PostDataStreamingResponse `json:"results,omitempty"`
	SchemaVersion *float32 `json:"schema_version,omitempty"`
}

// NewCreateDataStreamingResponse instantiates a new CreateDataStreamingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDataStreamingResponse() *CreateDataStreamingResponse {
	this := CreateDataStreamingResponse{}
	return &this
}

// NewCreateDataStreamingResponseWithDefaults instantiates a new CreateDataStreamingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDataStreamingResponseWithDefaults() *CreateDataStreamingResponse {
	this := CreateDataStreamingResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *CreateDataStreamingResponse) GetResults() []PostDataStreamingResponse {
	if o == nil || IsNil(o.Results) {
		var ret []PostDataStreamingResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataStreamingResponse) GetResultsOk() ([]PostDataStreamingResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CreateDataStreamingResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []PostDataStreamingResponse and assigns it to the Results field.
func (o *CreateDataStreamingResponse) SetResults(v []PostDataStreamingResponse) {
	o.Results = v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *CreateDataStreamingResponse) GetSchemaVersion() float32 {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret float32
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDataStreamingResponse) GetSchemaVersionOk() (*float32, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *CreateDataStreamingResponse) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given float32 and assigns it to the SchemaVersion field.
func (o *CreateDataStreamingResponse) SetSchemaVersion(v float32) {
	o.SchemaVersion = &v
}

func (o CreateDataStreamingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDataStreamingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	return toSerialize, nil
}

type NullableCreateDataStreamingResponse struct {
	value *CreateDataStreamingResponse
	isSet bool
}

func (v NullableCreateDataStreamingResponse) Get() *CreateDataStreamingResponse {
	return v.value
}

func (v *NullableCreateDataStreamingResponse) Set(val *CreateDataStreamingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDataStreamingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDataStreamingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDataStreamingResponse(val *CreateDataStreamingResponse) *NullableCreateDataStreamingResponse {
	return &NullableCreateDataStreamingResponse{value: val, isSet: true}
}

func (v NullableCreateDataStreamingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDataStreamingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


