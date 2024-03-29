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

// checks if the EdgeNodeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeNodeResponse{}

// EdgeNodeResponse struct for EdgeNodeResponse
type EdgeNodeResponse struct {
	Groups []NodeGroup `json:"groups,omitempty"`
	HashId *string `json:"hash_id,omitempty"`
	Id int64 `json:"id"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewEdgeNodeResponse instantiates a new EdgeNodeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeNodeResponse(id int64) *EdgeNodeResponse {
	this := EdgeNodeResponse{}
	this.Id = id
	return &this
}

// NewEdgeNodeResponseWithDefaults instantiates a new EdgeNodeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeNodeResponseWithDefaults() *EdgeNodeResponse {
	this := EdgeNodeResponse{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *EdgeNodeResponse) GetGroups() []NodeGroup {
	if o == nil || IsNil(o.Groups) {
		var ret []NodeGroup
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeNodeResponse) GetGroupsOk() ([]NodeGroup, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *EdgeNodeResponse) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []NodeGroup and assigns it to the Groups field.
func (o *EdgeNodeResponse) SetGroups(v []NodeGroup) {
	o.Groups = v
}

// GetHashId returns the HashId field value if set, zero value otherwise.
func (o *EdgeNodeResponse) GetHashId() string {
	if o == nil || IsNil(o.HashId) {
		var ret string
		return ret
	}
	return *o.HashId
}

// GetHashIdOk returns a tuple with the HashId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeNodeResponse) GetHashIdOk() (*string, bool) {
	if o == nil || IsNil(o.HashId) {
		return nil, false
	}
	return o.HashId, true
}

// HasHashId returns a boolean if a field has been set.
func (o *EdgeNodeResponse) HasHashId() bool {
	if o != nil && !IsNil(o.HashId) {
		return true
	}

	return false
}

// SetHashId gets a reference to the given string and assigns it to the HashId field.
func (o *EdgeNodeResponse) SetHashId(v string) {
	o.HashId = &v
}

// GetId returns the Id field value
func (o *EdgeNodeResponse) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EdgeNodeResponse) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EdgeNodeResponse) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EdgeNodeResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeNodeResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EdgeNodeResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EdgeNodeResponse) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EdgeNodeResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeNodeResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EdgeNodeResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EdgeNodeResponse) SetStatus(v string) {
	o.Status = &v
}

func (o EdgeNodeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeNodeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.HashId) {
		toSerialize["hash_id"] = o.HashId
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableEdgeNodeResponse struct {
	value *EdgeNodeResponse
	isSet bool
}

func (v NullableEdgeNodeResponse) Get() *EdgeNodeResponse {
	return v.value
}

func (v *NullableEdgeNodeResponse) Set(val *EdgeNodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeNodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeNodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeNodeResponse(val *EdgeNodeResponse) *NullableEdgeNodeResponse {
	return &NullableEdgeNodeResponse{value: val, isSet: true}
}

func (v NullableEdgeNodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeNodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


