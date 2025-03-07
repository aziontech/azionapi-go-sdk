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

// checks if the ApplicationPutInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationPutInstanceRequest{}

// ApplicationPutInstanceRequest struct for ApplicationPutInstanceRequest
type ApplicationPutInstanceRequest struct {
	Name string `json:"name"`
	EdgeFunctionId int64 `json:"edge_function_id"`
	Args ApplicationCreateInstanceRequestArgs `json:"args"`
}

type _ApplicationPutInstanceRequest ApplicationPutInstanceRequest

// NewApplicationPutInstanceRequest instantiates a new ApplicationPutInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationPutInstanceRequest(name string, edgeFunctionId int64, args ApplicationCreateInstanceRequestArgs) *ApplicationPutInstanceRequest {
	this := ApplicationPutInstanceRequest{}
	this.Name = name
	this.EdgeFunctionId = edgeFunctionId
	this.Args = args
	return &this
}

// NewApplicationPutInstanceRequestWithDefaults instantiates a new ApplicationPutInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationPutInstanceRequestWithDefaults() *ApplicationPutInstanceRequest {
	this := ApplicationPutInstanceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationPutInstanceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationPutInstanceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationPutInstanceRequest) SetName(v string) {
	o.Name = v
}

// GetEdgeFunctionId returns the EdgeFunctionId field value
func (o *ApplicationPutInstanceRequest) GetEdgeFunctionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EdgeFunctionId
}

// GetEdgeFunctionIdOk returns a tuple with the EdgeFunctionId field value
// and a boolean to check if the value has been set.
func (o *ApplicationPutInstanceRequest) GetEdgeFunctionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFunctionId, true
}

// SetEdgeFunctionId sets field value
func (o *ApplicationPutInstanceRequest) SetEdgeFunctionId(v int64) {
	o.EdgeFunctionId = v
}

// GetArgs returns the Args field value
func (o *ApplicationPutInstanceRequest) GetArgs() ApplicationCreateInstanceRequestArgs {
	if o == nil {
		var ret ApplicationCreateInstanceRequestArgs
		return ret
	}

	return o.Args
}

// GetArgsOk returns a tuple with the Args field value
// and a boolean to check if the value has been set.
func (o *ApplicationPutInstanceRequest) GetArgsOk() (*ApplicationCreateInstanceRequestArgs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Args, true
}

// SetArgs sets field value
func (o *ApplicationPutInstanceRequest) SetArgs(v ApplicationCreateInstanceRequestArgs) {
	o.Args = v
}

func (o ApplicationPutInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationPutInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["edge_function_id"] = o.EdgeFunctionId
	toSerialize["args"] = o.Args
	return toSerialize, nil
}

func (o *ApplicationPutInstanceRequest) UnmarshalJSON(data []byte) (err error) {
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

	varApplicationPutInstanceRequest := _ApplicationPutInstanceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationPutInstanceRequest)

	if err != nil {
		return err
	}

	*o = ApplicationPutInstanceRequest(varApplicationPutInstanceRequest)

	return err
}

type NullableApplicationPutInstanceRequest struct {
	value *ApplicationPutInstanceRequest
	isSet bool
}

func (v NullableApplicationPutInstanceRequest) Get() *ApplicationPutInstanceRequest {
	return v.value
}

func (v *NullableApplicationPutInstanceRequest) Set(val *ApplicationPutInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationPutInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationPutInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationPutInstanceRequest(val *ApplicationPutInstanceRequest) *NullableApplicationPutInstanceRequest {
	return &NullableApplicationPutInstanceRequest{value: val, isSet: true}
}

func (v NullableApplicationPutInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationPutInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


