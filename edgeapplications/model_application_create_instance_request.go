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

// checks if the ApplicationCreateInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCreateInstanceRequest{}

// ApplicationCreateInstanceRequest struct for ApplicationCreateInstanceRequest
type ApplicationCreateInstanceRequest struct {
	Name string `json:"name"`
	EdgeFunctionId int64 `json:"edge_function_id"`
	Args interface{} `json:"args"`
}

type _ApplicationCreateInstanceRequest ApplicationCreateInstanceRequest

// NewApplicationCreateInstanceRequest instantiates a new ApplicationCreateInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCreateInstanceRequest(name string, edgeFunctionId int64, args interface{}) *ApplicationCreateInstanceRequest {
	this := ApplicationCreateInstanceRequest{}
	this.Name = name
	this.EdgeFunctionId = edgeFunctionId
	this.Args = args
	return &this
}

// NewApplicationCreateInstanceRequestWithDefaults instantiates a new ApplicationCreateInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCreateInstanceRequestWithDefaults() *ApplicationCreateInstanceRequest {
	this := ApplicationCreateInstanceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationCreateInstanceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCreateInstanceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCreateInstanceRequest) SetName(v string) {
	o.Name = v
}

// GetEdgeFunctionId returns the EdgeFunctionId field value
func (o *ApplicationCreateInstanceRequest) GetEdgeFunctionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EdgeFunctionId
}

// GetEdgeFunctionIdOk returns a tuple with the EdgeFunctionId field value
// and a boolean to check if the value has been set.
func (o *ApplicationCreateInstanceRequest) GetEdgeFunctionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFunctionId, true
}

// SetEdgeFunctionId sets field value
func (o *ApplicationCreateInstanceRequest) SetEdgeFunctionId(v int64) {
	o.EdgeFunctionId = v
}

// GetArgs returns the Args field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApplicationCreateInstanceRequest) GetArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Args
}

// GetArgsOk returns a tuple with the Args field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCreateInstanceRequest) GetArgsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Args) {
		return nil, false
	}
	return &o.Args, true
}

// SetArgs sets field value
func (o *ApplicationCreateInstanceRequest) SetArgs(v interface{}) {
	o.Args = v
}

func (o ApplicationCreateInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCreateInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["edge_function_id"] = o.EdgeFunctionId
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	return toSerialize, nil
}

func (o *ApplicationCreateInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"edge_function_id",
		"args",
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

	varApplicationCreateInstanceRequest := _ApplicationCreateInstanceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCreateInstanceRequest)

	if err != nil {
		return err
	}

	*o = ApplicationCreateInstanceRequest(varApplicationCreateInstanceRequest)

	return err
}

type NullableApplicationCreateInstanceRequest struct {
	value *ApplicationCreateInstanceRequest
	isSet bool
}

func (v NullableApplicationCreateInstanceRequest) Get() *ApplicationCreateInstanceRequest {
	return v.value
}

func (v *NullableApplicationCreateInstanceRequest) Set(val *ApplicationCreateInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCreateInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCreateInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCreateInstanceRequest(val *ApplicationCreateInstanceRequest) *NullableApplicationCreateInstanceRequest {
	return &NullableApplicationCreateInstanceRequest{value: val, isSet: true}
}

func (v NullableApplicationCreateInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCreateInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


