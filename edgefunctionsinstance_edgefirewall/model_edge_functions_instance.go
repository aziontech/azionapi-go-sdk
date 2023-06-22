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

// checks if the EdgeFunctionsInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeFunctionsInstance{}

// EdgeFunctionsInstance struct for EdgeFunctionsInstance
type EdgeFunctionsInstance struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	JsonArgs map[string]interface{} `json:"json_args,omitempty"`
	EdgeFunction *int32 `json:"edge_function,omitempty"`
}

// NewEdgeFunctionsInstance instantiates a new EdgeFunctionsInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeFunctionsInstance() *EdgeFunctionsInstance {
	this := EdgeFunctionsInstance{}
	return &this
}

// NewEdgeFunctionsInstanceWithDefaults instantiates a new EdgeFunctionsInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeFunctionsInstanceWithDefaults() *EdgeFunctionsInstance {
	this := EdgeFunctionsInstance{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EdgeFunctionsInstance) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsInstance) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EdgeFunctionsInstance) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EdgeFunctionsInstance) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EdgeFunctionsInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EdgeFunctionsInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EdgeFunctionsInstance) SetName(v string) {
	o.Name = &v
}

// GetJsonArgs returns the JsonArgs field value if set, zero value otherwise.
func (o *EdgeFunctionsInstance) GetJsonArgs() map[string]interface{} {
	if o == nil || IsNil(o.JsonArgs) {
		var ret map[string]interface{}
		return ret
	}
	return o.JsonArgs
}

// GetJsonArgsOk returns a tuple with the JsonArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsInstance) GetJsonArgsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.JsonArgs) {
		return map[string]interface{}{}, false
	}
	return o.JsonArgs, true
}

// HasJsonArgs returns a boolean if a field has been set.
func (o *EdgeFunctionsInstance) HasJsonArgs() bool {
	if o != nil && !IsNil(o.JsonArgs) {
		return true
	}

	return false
}

// SetJsonArgs gets a reference to the given map[string]interface{} and assigns it to the JsonArgs field.
func (o *EdgeFunctionsInstance) SetJsonArgs(v map[string]interface{}) {
	o.JsonArgs = v
}

// GetEdgeFunction returns the EdgeFunction field value if set, zero value otherwise.
func (o *EdgeFunctionsInstance) GetEdgeFunction() int32 {
	if o == nil || IsNil(o.EdgeFunction) {
		var ret int32
		return ret
	}
	return *o.EdgeFunction
}

// GetEdgeFunctionOk returns a tuple with the EdgeFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeFunctionsInstance) GetEdgeFunctionOk() (*int32, bool) {
	if o == nil || IsNil(o.EdgeFunction) {
		return nil, false
	}
	return o.EdgeFunction, true
}

// HasEdgeFunction returns a boolean if a field has been set.
func (o *EdgeFunctionsInstance) HasEdgeFunction() bool {
	if o != nil && !IsNil(o.EdgeFunction) {
		return true
	}

	return false
}

// SetEdgeFunction gets a reference to the given int32 and assigns it to the EdgeFunction field.
func (o *EdgeFunctionsInstance) SetEdgeFunction(v int32) {
	o.EdgeFunction = &v
}

func (o EdgeFunctionsInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeFunctionsInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.JsonArgs) {
		toSerialize["json_args"] = o.JsonArgs
	}
	if !IsNil(o.EdgeFunction) {
		toSerialize["edge_function"] = o.EdgeFunction
	}
	return toSerialize, nil
}

type NullableEdgeFunctionsInstance struct {
	value *EdgeFunctionsInstance
	isSet bool
}

func (v NullableEdgeFunctionsInstance) Get() *EdgeFunctionsInstance {
	return v.value
}

func (v *NullableEdgeFunctionsInstance) Set(val *EdgeFunctionsInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeFunctionsInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeFunctionsInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeFunctionsInstance(val *EdgeFunctionsInstance) *NullableEdgeFunctionsInstance {
	return &NullableEdgeFunctionsInstance{value: val, isSet: true}
}

func (v NullableEdgeFunctionsInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeFunctionsInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


