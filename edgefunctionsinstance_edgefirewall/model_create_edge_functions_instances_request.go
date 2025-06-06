/*
Edge Functions Instances API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgefunctionsinstance_edgefirewall

import (
	"encoding/json"
)

// checks if the CreateEdgeFunctionsInstancesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEdgeFunctionsInstancesRequest{}

// CreateEdgeFunctionsInstancesRequest struct for CreateEdgeFunctionsInstancesRequest
type CreateEdgeFunctionsInstancesRequest struct {
	Name *string `json:"name,omitempty"`
	EdgeFunction *int64 `json:"edge_function,omitempty"`
	JsonArgs *CreateEdgeFunctionsInstancesRequestJsonArgs `json:"json_args,omitempty"`
}

// NewCreateEdgeFunctionsInstancesRequest instantiates a new CreateEdgeFunctionsInstancesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEdgeFunctionsInstancesRequest() *CreateEdgeFunctionsInstancesRequest {
	this := CreateEdgeFunctionsInstancesRequest{}
	return &this
}

// NewCreateEdgeFunctionsInstancesRequestWithDefaults instantiates a new CreateEdgeFunctionsInstancesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEdgeFunctionsInstancesRequestWithDefaults() *CreateEdgeFunctionsInstancesRequest {
	this := CreateEdgeFunctionsInstancesRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateEdgeFunctionsInstancesRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEdgeFunctionsInstancesRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateEdgeFunctionsInstancesRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateEdgeFunctionsInstancesRequest) SetName(v string) {
	o.Name = &v
}

// GetEdgeFunction returns the EdgeFunction field value if set, zero value otherwise.
func (o *CreateEdgeFunctionsInstancesRequest) GetEdgeFunction() int64 {
	if o == nil || IsNil(o.EdgeFunction) {
		var ret int64
		return ret
	}
	return *o.EdgeFunction
}

// GetEdgeFunctionOk returns a tuple with the EdgeFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEdgeFunctionsInstancesRequest) GetEdgeFunctionOk() (*int64, bool) {
	if o == nil || IsNil(o.EdgeFunction) {
		return nil, false
	}
	return o.EdgeFunction, true
}

// HasEdgeFunction returns a boolean if a field has been set.
func (o *CreateEdgeFunctionsInstancesRequest) HasEdgeFunction() bool {
	if o != nil && !IsNil(o.EdgeFunction) {
		return true
	}

	return false
}

// SetEdgeFunction gets a reference to the given int64 and assigns it to the EdgeFunction field.
func (o *CreateEdgeFunctionsInstancesRequest) SetEdgeFunction(v int64) {
	o.EdgeFunction = &v
}

// GetJsonArgs returns the JsonArgs field value if set, zero value otherwise.
func (o *CreateEdgeFunctionsInstancesRequest) GetJsonArgs() CreateEdgeFunctionsInstancesRequestJsonArgs {
	if o == nil || IsNil(o.JsonArgs) {
		var ret CreateEdgeFunctionsInstancesRequestJsonArgs
		return ret
	}
	return *o.JsonArgs
}

// GetJsonArgsOk returns a tuple with the JsonArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEdgeFunctionsInstancesRequest) GetJsonArgsOk() (*CreateEdgeFunctionsInstancesRequestJsonArgs, bool) {
	if o == nil || IsNil(o.JsonArgs) {
		return nil, false
	}
	return o.JsonArgs, true
}

// HasJsonArgs returns a boolean if a field has been set.
func (o *CreateEdgeFunctionsInstancesRequest) HasJsonArgs() bool {
	if o != nil && !IsNil(o.JsonArgs) {
		return true
	}

	return false
}

// SetJsonArgs gets a reference to the given CreateEdgeFunctionsInstancesRequestJsonArgs and assigns it to the JsonArgs field.
func (o *CreateEdgeFunctionsInstancesRequest) SetJsonArgs(v CreateEdgeFunctionsInstancesRequestJsonArgs) {
	o.JsonArgs = &v
}

func (o CreateEdgeFunctionsInstancesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEdgeFunctionsInstancesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.EdgeFunction) {
		toSerialize["edge_function"] = o.EdgeFunction
	}
	if !IsNil(o.JsonArgs) {
		toSerialize["json_args"] = o.JsonArgs
	}
	return toSerialize, nil
}

type NullableCreateEdgeFunctionsInstancesRequest struct {
	value *CreateEdgeFunctionsInstancesRequest
	isSet bool
}

func (v NullableCreateEdgeFunctionsInstancesRequest) Get() *CreateEdgeFunctionsInstancesRequest {
	return v.value
}

func (v *NullableCreateEdgeFunctionsInstancesRequest) Set(val *CreateEdgeFunctionsInstancesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEdgeFunctionsInstancesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEdgeFunctionsInstancesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEdgeFunctionsInstancesRequest(val *CreateEdgeFunctionsInstancesRequest) *NullableCreateEdgeFunctionsInstancesRequest {
	return &NullableCreateEdgeFunctionsInstancesRequest{value: val, isSet: true}
}

func (v NullableCreateEdgeFunctionsInstancesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEdgeFunctionsInstancesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


