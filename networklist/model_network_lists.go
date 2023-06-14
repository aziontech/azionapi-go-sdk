/*
Azion API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networklist

import (
	"encoding/json"
)

// checks if the NetworkLists type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkLists{}

// NetworkLists struct for NetworkLists
type NetworkLists struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ItemsValues []string `json:"items_values,omitempty"`
	ListType *string `json:"list_type,omitempty"`
}

// NewNetworkLists instantiates a new NetworkLists object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkLists() *NetworkLists {
	this := NetworkLists{}
	return &this
}

// NewNetworkListsWithDefaults instantiates a new NetworkLists object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkListsWithDefaults() *NetworkLists {
	this := NetworkLists{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkLists) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLists) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkLists) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NetworkLists) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkLists) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLists) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkLists) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkLists) SetName(v string) {
	o.Name = &v
}

// GetItemsValues returns the ItemsValues field value if set, zero value otherwise.
func (o *NetworkLists) GetItemsValues() []string {
	if o == nil || IsNil(o.ItemsValues) {
		var ret []string
		return ret
	}
	return o.ItemsValues
}

// GetItemsValuesOk returns a tuple with the ItemsValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLists) GetItemsValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.ItemsValues) {
		return nil, false
	}
	return o.ItemsValues, true
}

// HasItemsValues returns a boolean if a field has been set.
func (o *NetworkLists) HasItemsValues() bool {
	if o != nil && !IsNil(o.ItemsValues) {
		return true
	}

	return false
}

// SetItemsValues gets a reference to the given []string and assigns it to the ItemsValues field.
func (o *NetworkLists) SetItemsValues(v []string) {
	o.ItemsValues = v
}

// GetListType returns the ListType field value if set, zero value otherwise.
func (o *NetworkLists) GetListType() string {
	if o == nil || IsNil(o.ListType) {
		var ret string
		return ret
	}
	return *o.ListType
}

// GetListTypeOk returns a tuple with the ListType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLists) GetListTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ListType) {
		return nil, false
	}
	return o.ListType, true
}

// HasListType returns a boolean if a field has been set.
func (o *NetworkLists) HasListType() bool {
	if o != nil && !IsNil(o.ListType) {
		return true
	}

	return false
}

// SetListType gets a reference to the given string and assigns it to the ListType field.
func (o *NetworkLists) SetListType(v string) {
	o.ListType = &v
}

func (o NetworkLists) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkLists) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ItemsValues) {
		toSerialize["items_values"] = o.ItemsValues
	}
	if !IsNil(o.ListType) {
		toSerialize["list_type"] = o.ListType
	}
	return toSerialize, nil
}

type NullableNetworkLists struct {
	value *NetworkLists
	isSet bool
}

func (v NullableNetworkLists) Get() *NetworkLists {
	return v.value
}

func (v *NullableNetworkLists) Set(val *NetworkLists) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkLists) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkLists) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkLists(val *NetworkLists) *NullableNetworkLists {
	return &NullableNetworkLists{value: val, isSet: true}
}

func (v NullableNetworkLists) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkLists) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


