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

// checks if the BadRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BadRequestResponse{}

// BadRequestResponse struct for BadRequestResponse
type BadRequestResponse struct {
	Name []string `json:"name,omitempty"`
	ItemsValues []string `json:"items_values,omitempty"`
	ListType []string `json:"list_type,omitempty"`
}

// NewBadRequestResponse instantiates a new BadRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadRequestResponse() *BadRequestResponse {
	this := BadRequestResponse{}
	return &this
}

// NewBadRequestResponseWithDefaults instantiates a new BadRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadRequestResponseWithDefaults() *BadRequestResponse {
	this := BadRequestResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BadRequestResponse) GetName() []string {
	if o == nil || IsNil(o.Name) {
		var ret []string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestResponse) GetNameOk() ([]string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BadRequestResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *BadRequestResponse) SetName(v []string) {
	o.Name = v
}

// GetItemsValues returns the ItemsValues field value if set, zero value otherwise.
func (o *BadRequestResponse) GetItemsValues() []string {
	if o == nil || IsNil(o.ItemsValues) {
		var ret []string
		return ret
	}
	return o.ItemsValues
}

// GetItemsValuesOk returns a tuple with the ItemsValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestResponse) GetItemsValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.ItemsValues) {
		return nil, false
	}
	return o.ItemsValues, true
}

// HasItemsValues returns a boolean if a field has been set.
func (o *BadRequestResponse) HasItemsValues() bool {
	if o != nil && !IsNil(o.ItemsValues) {
		return true
	}

	return false
}

// SetItemsValues gets a reference to the given []string and assigns it to the ItemsValues field.
func (o *BadRequestResponse) SetItemsValues(v []string) {
	o.ItemsValues = v
}

// GetListType returns the ListType field value if set, zero value otherwise.
func (o *BadRequestResponse) GetListType() []string {
	if o == nil || IsNil(o.ListType) {
		var ret []string
		return ret
	}
	return o.ListType
}

// GetListTypeOk returns a tuple with the ListType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BadRequestResponse) GetListTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.ListType) {
		return nil, false
	}
	return o.ListType, true
}

// HasListType returns a boolean if a field has been set.
func (o *BadRequestResponse) HasListType() bool {
	if o != nil && !IsNil(o.ListType) {
		return true
	}

	return false
}

// SetListType gets a reference to the given []string and assigns it to the ListType field.
func (o *BadRequestResponse) SetListType(v []string) {
	o.ListType = v
}

func (o BadRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BadRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableBadRequestResponse struct {
	value *BadRequestResponse
	isSet bool
}

func (v NullableBadRequestResponse) Get() *BadRequestResponse {
	return v.value
}

func (v *NullableBadRequestResponse) Set(val *BadRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBadRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBadRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadRequestResponse(val *BadRequestResponse) *NullableBadRequestResponse {
	return &NullableBadRequestResponse{value: val, isSet: true}
}

func (v NullableBadRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


