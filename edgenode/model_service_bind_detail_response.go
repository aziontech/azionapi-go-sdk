/*
Edge Node API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgenode

import (
	"encoding/json"
)

// checks if the ServiceBindDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceBindDetailResponse{}

// ServiceBindDetailResponse struct for ServiceBindDetailResponse
type ServiceBindDetailResponse struct {
	Id int64 `json:"id"`
	ServiceId int64 `json:"service_id"`
	ServiceName string `json:"service_name"`
	Variables []Variable `json:"variables"`
}

// NewServiceBindDetailResponse instantiates a new ServiceBindDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBindDetailResponse(id int64, serviceId int64, serviceName string, variables []Variable) *ServiceBindDetailResponse {
	this := ServiceBindDetailResponse{}
	this.Id = id
	this.ServiceId = serviceId
	this.ServiceName = serviceName
	this.Variables = variables
	return &this
}

// NewServiceBindDetailResponseWithDefaults instantiates a new ServiceBindDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBindDetailResponseWithDefaults() *ServiceBindDetailResponse {
	this := ServiceBindDetailResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceBindDetailResponse) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceBindDetailResponse) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceBindDetailResponse) SetId(v int64) {
	o.Id = v
}

// GetServiceId returns the ServiceId field value
func (o *ServiceBindDetailResponse) GetServiceId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ServiceBindDetailResponse) GetServiceIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *ServiceBindDetailResponse) SetServiceId(v int64) {
	o.ServiceId = v
}

// GetServiceName returns the ServiceName field value
func (o *ServiceBindDetailResponse) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *ServiceBindDetailResponse) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *ServiceBindDetailResponse) SetServiceName(v string) {
	o.ServiceName = v
}

// GetVariables returns the Variables field value
func (o *ServiceBindDetailResponse) GetVariables() []Variable {
	if o == nil {
		var ret []Variable
		return ret
	}

	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *ServiceBindDetailResponse) GetVariablesOk() ([]Variable, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variables, true
}

// SetVariables sets field value
func (o *ServiceBindDetailResponse) SetVariables(v []Variable) {
	o.Variables = v
}

func (o ServiceBindDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceBindDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["service_id"] = o.ServiceId
	toSerialize["service_name"] = o.ServiceName
	toSerialize["variables"] = o.Variables
	return toSerialize, nil
}

type NullableServiceBindDetailResponse struct {
	value *ServiceBindDetailResponse
	isSet bool
}

func (v NullableServiceBindDetailResponse) Get() *ServiceBindDetailResponse {
	return v.value
}

func (v *NullableServiceBindDetailResponse) Set(val *ServiceBindDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBindDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBindDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBindDetailResponse(val *ServiceBindDetailResponse) *NullableServiceBindDetailResponse {
	return &NullableServiceBindDetailResponse{value: val, isSet: true}
}

func (v NullableServiceBindDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBindDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


