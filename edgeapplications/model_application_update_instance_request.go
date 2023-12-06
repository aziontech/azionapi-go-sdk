/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
	"fmt"
)

// checks if the ApplicationUpdateInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationUpdateInstanceRequest{}

// ApplicationUpdateInstanceRequest struct for ApplicationUpdateInstanceRequest
type ApplicationUpdateInstanceRequest struct {
	Name NullableString `json:"name"`
	EdgeFunctionId NullableInt64 `json:"edge_function_id"`
	Args interface{} `json:"args"`
}

type _ApplicationUpdateInstanceRequest ApplicationUpdateInstanceRequest

// NewApplicationUpdateInstanceRequest instantiates a new ApplicationUpdateInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationUpdateInstanceRequest(name NullableString, edgeFunctionId NullableInt64, args interface{}) *ApplicationUpdateInstanceRequest {
	this := ApplicationUpdateInstanceRequest{}
	this.Name = name
	this.EdgeFunctionId = edgeFunctionId
	this.Args = args
	return &this
}

// NewApplicationUpdateInstanceRequestWithDefaults instantiates a new ApplicationUpdateInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationUpdateInstanceRequestWithDefaults() *ApplicationUpdateInstanceRequest {
	this := ApplicationUpdateInstanceRequest{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApplicationUpdateInstanceRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationUpdateInstanceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *ApplicationUpdateInstanceRequest) SetName(v string) {
	o.Name.Set(&v)
}

// GetEdgeFunctionId returns the EdgeFunctionId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *ApplicationUpdateInstanceRequest) GetEdgeFunctionId() int64 {
	if o == nil || o.EdgeFunctionId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.EdgeFunctionId.Get()
}

// GetEdgeFunctionIdOk returns a tuple with the EdgeFunctionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationUpdateInstanceRequest) GetEdgeFunctionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EdgeFunctionId.Get(), o.EdgeFunctionId.IsSet()
}

// SetEdgeFunctionId sets field value
func (o *ApplicationUpdateInstanceRequest) SetEdgeFunctionId(v int64) {
	o.EdgeFunctionId.Set(&v)
}

// GetArgs returns the Args field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApplicationUpdateInstanceRequest) GetArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Args
}

// GetArgsOk returns a tuple with the Args field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationUpdateInstanceRequest) GetArgsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Args) {
		return nil, false
	}
	return &o.Args, true
}

// SetArgs sets field value
func (o *ApplicationUpdateInstanceRequest) SetArgs(v interface{}) {
	o.Args = v
}

func (o ApplicationUpdateInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationUpdateInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	toSerialize["edge_function_id"] = o.EdgeFunctionId.Get()
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	return toSerialize, nil
}

func (o *ApplicationUpdateInstanceRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"edge_function_id",
		"args",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApplicationUpdateInstanceRequest := _ApplicationUpdateInstanceRequest{}

	err = json.Unmarshal(bytes, &varApplicationUpdateInstanceRequest)

	if err != nil {
		return err
	}

	*o = ApplicationUpdateInstanceRequest(varApplicationUpdateInstanceRequest)

	return err
}

type NullableApplicationUpdateInstanceRequest struct {
	value *ApplicationUpdateInstanceRequest
	isSet bool
}

func (v NullableApplicationUpdateInstanceRequest) Get() *ApplicationUpdateInstanceRequest {
	return v.value
}

func (v *NullableApplicationUpdateInstanceRequest) Set(val *ApplicationUpdateInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationUpdateInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationUpdateInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationUpdateInstanceRequest(val *ApplicationUpdateInstanceRequest) *NullableApplicationUpdateInstanceRequest {
	return &NullableApplicationUpdateInstanceRequest{value: val, isSet: true}
}

func (v NullableApplicationUpdateInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationUpdateInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


