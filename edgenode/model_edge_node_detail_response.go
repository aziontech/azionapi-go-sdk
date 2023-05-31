/*
Edgenode API

Azion Orchestration

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgenode

import (
	"encoding/json"
)

// checks if the EdgeNodeDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeNodeDetailResponse{}

// EdgeNodeDetailResponse struct for EdgeNodeDetailResponse
type EdgeNodeDetailResponse struct {
	Groups []NodeGroup `json:"groups"`
	HasServices bool `json:"has_services"`
	HashId string `json:"hash_id"`
	Id int64 `json:"id"`
	Modules EdgeNodeModules `json:"modules"`
	Name string `json:"name"`
}

// NewEdgeNodeDetailResponse instantiates a new EdgeNodeDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeNodeDetailResponse(groups []NodeGroup, hasServices bool, hashId string, id int64, modules EdgeNodeModules, name string) *EdgeNodeDetailResponse {
	this := EdgeNodeDetailResponse{}
	this.Groups = groups
	this.HasServices = hasServices
	this.HashId = hashId
	this.Id = id
	this.Modules = modules
	this.Name = name
	return &this
}

// NewEdgeNodeDetailResponseWithDefaults instantiates a new EdgeNodeDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeNodeDetailResponseWithDefaults() *EdgeNodeDetailResponse {
	this := EdgeNodeDetailResponse{}
	return &this
}

// GetGroups returns the Groups field value
func (o *EdgeNodeDetailResponse) GetGroups() []NodeGroup {
	if o == nil {
		var ret []NodeGroup
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *EdgeNodeDetailResponse) GetGroupsOk() ([]NodeGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *EdgeNodeDetailResponse) SetGroups(v []NodeGroup) {
	o.Groups = v
}

// GetHasServices returns the HasServices field value
func (o *EdgeNodeDetailResponse) GetHasServices() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasServices
}

// GetHasServicesOk returns a tuple with the HasServices field value
// and a boolean to check if the value has been set.
func (o *EdgeNodeDetailResponse) GetHasServicesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasServices, true
}

// SetHasServices sets field value
func (o *EdgeNodeDetailResponse) SetHasServices(v bool) {
	o.HasServices = v
}

// GetHashId returns the HashId field value
func (o *EdgeNodeDetailResponse) GetHashId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HashId
}

// GetHashIdOk returns a tuple with the HashId field value
// and a boolean to check if the value has been set.
func (o *EdgeNodeDetailResponse) GetHashIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HashId, true
}

// SetHashId sets field value
func (o *EdgeNodeDetailResponse) SetHashId(v string) {
	o.HashId = v
}

// GetId returns the Id field value
func (o *EdgeNodeDetailResponse) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EdgeNodeDetailResponse) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EdgeNodeDetailResponse) SetId(v int64) {
	o.Id = v
}

// GetModules returns the Modules field value
func (o *EdgeNodeDetailResponse) GetModules() EdgeNodeModules {
	if o == nil {
		var ret EdgeNodeModules
		return ret
	}

	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value
// and a boolean to check if the value has been set.
func (o *EdgeNodeDetailResponse) GetModulesOk() (*EdgeNodeModules, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modules, true
}

// SetModules sets field value
func (o *EdgeNodeDetailResponse) SetModules(v EdgeNodeModules) {
	o.Modules = v
}

// GetName returns the Name field value
func (o *EdgeNodeDetailResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EdgeNodeDetailResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EdgeNodeDetailResponse) SetName(v string) {
	o.Name = v
}

func (o EdgeNodeDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeNodeDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groups"] = o.Groups
	toSerialize["has_services"] = o.HasServices
	toSerialize["hash_id"] = o.HashId
	toSerialize["id"] = o.Id
	toSerialize["modules"] = o.Modules
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableEdgeNodeDetailResponse struct {
	value *EdgeNodeDetailResponse
	isSet bool
}

func (v NullableEdgeNodeDetailResponse) Get() *EdgeNodeDetailResponse {
	return v.value
}

func (v *NullableEdgeNodeDetailResponse) Set(val *EdgeNodeDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeNodeDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeNodeDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeNodeDetailResponse(val *EdgeNodeDetailResponse) *NullableEdgeNodeDetailResponse {
	return &NullableEdgeNodeDetailResponse{value: val, isSet: true}
}

func (v NullableEdgeNodeDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeNodeDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


