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

// EdgeFunctionResponse struct for EdgeFunctionResponse
type EdgeFunctionResponse struct {
	Results *Results `json:"results,omitempty"`
	SchemaVersion *float32 `json:"schema_version,omitempty"`
}

// NewEdgeFunctionResponse instantiates a new EdgeFunctionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFunctionResponse() *EdgeFunctionResponse {
	this := EdgeFunctionResponse{}
	return &this
}

// NewEdgeFunctionResponseWithDefaults instantiates a new EdgeFunctionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFunctionResponseWithDefaults() *EdgeFunctionResponse {
	this := EdgeFunctionResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *EdgeFunctionResponse) GetResults() Results {
	if o == nil || o.Results == nil {
		var ret Results
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFunctionResponse) GetResultsOk() (*Results, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *EdgeFunctionResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given Results and assigns it to the Results field.
func (o *EdgeFunctionResponse) SetResults(v Results) {
	o.Results = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *EdgeFunctionResponse) GetSchemaVersion() float32 {
	if o == nil || o.SchemaVersion == nil {
		var ret float32
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFunctionResponse) GetSchemaVersionOk() (*float32, bool) {
	if o == nil || o.SchemaVersion == nil {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *EdgeFunctionResponse) HasSchemaVersion() bool {
	if o != nil && o.SchemaVersion != nil {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given float32 and assigns it to the SchemaVersion field.
func (o *EdgeFunctionResponse) SetSchemaVersion(v float32) {
	o.SchemaVersion = &v
}

func (o EdgeFunctionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if o.SchemaVersion != nil {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	return json.Marshal(toSerialize)
}

type NullableEdgeFunctionResponse struct {
	value *EdgeFunctionResponse
	isSet bool
}

func (v NullableEdgeFunctionResponse) Get() *EdgeFunctionResponse {
	return v.value
}

func (v *NullableEdgeFunctionResponse) Set(val *EdgeFunctionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFunctionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFunctionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFunctionResponse(val *EdgeFunctionResponse) *NullableEdgeFunctionResponse {
	return &NullableEdgeFunctionResponse{value: val, isSet: true}
}

func (v NullableEdgeFunctionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFunctionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


