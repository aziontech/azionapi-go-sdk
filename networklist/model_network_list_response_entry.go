/*
Network Lists API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networklist

import (
	"encoding/json"
)

// checks if the NetworkListResponseEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkListResponseEntry{}

// NetworkListResponseEntry struct for NetworkListResponseEntry
type NetworkListResponseEntry struct {
	Id *int64 `json:"id,omitempty"`
	LastEditor *string `json:"last_editor,omitempty"`
	LastModified *string `json:"last_modified,omitempty"`
	ListType *string `json:"list_type,omitempty"`
	Name *string `json:"name,omitempty"`
	ItemsValues []string `json:"items_values,omitempty"`
}

// NewNetworkListResponseEntry instantiates a new NetworkListResponseEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkListResponseEntry() *NetworkListResponseEntry {
	this := NetworkListResponseEntry{}
	return &this
}

// NewNetworkListResponseEntryWithDefaults instantiates a new NetworkListResponseEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkListResponseEntryWithDefaults() *NetworkListResponseEntry {
	this := NetworkListResponseEntry{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkListResponseEntry) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkListResponseEntry) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkListResponseEntry) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *NetworkListResponseEntry) SetId(v int64) {
	o.Id = &v
}

// GetLastEditor returns the LastEditor field value if set, zero value otherwise.
func (o *NetworkListResponseEntry) GetLastEditor() string {
	if o == nil || IsNil(o.LastEditor) {
		var ret string
		return ret
	}
	return *o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkListResponseEntry) GetLastEditorOk() (*string, bool) {
	if o == nil || IsNil(o.LastEditor) {
		return nil, false
	}
	return o.LastEditor, true
}

// HasLastEditor returns a boolean if a field has been set.
func (o *NetworkListResponseEntry) HasLastEditor() bool {
	if o != nil && !IsNil(o.LastEditor) {
		return true
	}

	return false
}

// SetLastEditor gets a reference to the given string and assigns it to the LastEditor field.
func (o *NetworkListResponseEntry) SetLastEditor(v string) {
	o.LastEditor = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *NetworkListResponseEntry) GetLastModified() string {
	if o == nil || IsNil(o.LastModified) {
		var ret string
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkListResponseEntry) GetLastModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *NetworkListResponseEntry) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given string and assigns it to the LastModified field.
func (o *NetworkListResponseEntry) SetLastModified(v string) {
	o.LastModified = &v
}

// GetListType returns the ListType field value if set, zero value otherwise.
func (o *NetworkListResponseEntry) GetListType() string {
	if o == nil || IsNil(o.ListType) {
		var ret string
		return ret
	}
	return *o.ListType
}

// GetListTypeOk returns a tuple with the ListType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkListResponseEntry) GetListTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ListType) {
		return nil, false
	}
	return o.ListType, true
}

// HasListType returns a boolean if a field has been set.
func (o *NetworkListResponseEntry) HasListType() bool {
	if o != nil && !IsNil(o.ListType) {
		return true
	}

	return false
}

// SetListType gets a reference to the given string and assigns it to the ListType field.
func (o *NetworkListResponseEntry) SetListType(v string) {
	o.ListType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkListResponseEntry) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkListResponseEntry) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkListResponseEntry) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkListResponseEntry) SetName(v string) {
	o.Name = &v
}

// GetItemsValues returns the ItemsValues field value if set, zero value otherwise.
func (o *NetworkListResponseEntry) GetItemsValues() []string {
	if o == nil || IsNil(o.ItemsValues) {
		var ret []string
		return ret
	}
	return o.ItemsValues
}

// GetItemsValuesOk returns a tuple with the ItemsValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkListResponseEntry) GetItemsValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.ItemsValues) {
		return nil, false
	}
	return o.ItemsValues, true
}

// HasItemsValues returns a boolean if a field has been set.
func (o *NetworkListResponseEntry) HasItemsValues() bool {
	if o != nil && !IsNil(o.ItemsValues) {
		return true
	}

	return false
}

// SetItemsValues gets a reference to the given []string and assigns it to the ItemsValues field.
func (o *NetworkListResponseEntry) SetItemsValues(v []string) {
	o.ItemsValues = v
}

func (o NetworkListResponseEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkListResponseEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastEditor) {
		toSerialize["last_editor"] = o.LastEditor
	}
	if !IsNil(o.LastModified) {
		toSerialize["last_modified"] = o.LastModified
	}
	if !IsNil(o.ListType) {
		toSerialize["list_type"] = o.ListType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ItemsValues) {
		toSerialize["items_values"] = o.ItemsValues
	}
	return toSerialize, nil
}

type NullableNetworkListResponseEntry struct {
	value *NetworkListResponseEntry
	isSet bool
}

func (v NullableNetworkListResponseEntry) Get() *NetworkListResponseEntry {
	return v.value
}

func (v *NullableNetworkListResponseEntry) Set(val *NetworkListResponseEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkListResponseEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkListResponseEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkListResponseEntry(val *NetworkListResponseEntry) *NullableNetworkListResponseEntry {
	return &NullableNetworkListResponseEntry{value: val, isSet: true}
}

func (v NullableNetworkListResponseEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkListResponseEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

