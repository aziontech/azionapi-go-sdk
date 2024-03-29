/*
Services API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package services

import (
	"encoding/json"
)

// checks if the ResourceDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceDetail{}

// ResourceDetail struct for ResourceDetail
type ResourceDetail struct {
	Content string `json:"content"`
	ContentType string `json:"content_type"`
	Id int64 `json:"id"`
	Name string `json:"name"`
	Trigger string `json:"trigger"`
}

// NewResourceDetail instantiates a new ResourceDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceDetail(content string, contentType string, id int64, name string, trigger string) *ResourceDetail {
	this := ResourceDetail{}
	this.Content = content
	this.ContentType = contentType
	this.Id = id
	this.Name = name
	this.Trigger = trigger
	return &this
}

// NewResourceDetailWithDefaults instantiates a new ResourceDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceDetailWithDefaults() *ResourceDetail {
	this := ResourceDetail{}
	return &this
}

// GetContent returns the Content field value
func (o *ResourceDetail) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ResourceDetail) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ResourceDetail) SetContent(v string) {
	o.Content = v
}

// GetContentType returns the ContentType field value
func (o *ResourceDetail) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *ResourceDetail) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *ResourceDetail) SetContentType(v string) {
	o.ContentType = v
}

// GetId returns the Id field value
func (o *ResourceDetail) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceDetail) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceDetail) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ResourceDetail) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceDetail) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceDetail) SetName(v string) {
	o.Name = v
}

// GetTrigger returns the Trigger field value
func (o *ResourceDetail) GetTrigger() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *ResourceDetail) GetTriggerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value
func (o *ResourceDetail) SetTrigger(v string) {
	o.Trigger = v
}

func (o ResourceDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	toSerialize["content_type"] = o.ContentType
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["trigger"] = o.Trigger
	return toSerialize, nil
}

type NullableResourceDetail struct {
	value *ResourceDetail
	isSet bool
}

func (v NullableResourceDetail) Get() *ResourceDetail {
	return v.value
}

func (v *NullableResourceDetail) Set(val *ResourceDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceDetail(val *ResourceDetail) *NullableResourceDetail {
	return &NullableResourceDetail{value: val, isSet: true}
}

func (v NullableResourceDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


