/*
Services API

Azion Services

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package services

import (
	"encoding/json"
)

// checks if the CreateResourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateResourceRequest{}

// CreateResourceRequest struct for CreateResourceRequest
type CreateResourceRequest struct {
	Content string `json:"content"`
	ContentType string `json:"content_type"`
	Name string `json:"name"`
	Trigger string `json:"trigger"`
}

// NewCreateResourceRequest instantiates a new CreateResourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateResourceRequest(content string, contentType string, name string, trigger string) *CreateResourceRequest {
	this := CreateResourceRequest{}
	this.Content = content
	this.ContentType = contentType
	this.Name = name
	this.Trigger = trigger
	return &this
}

// NewCreateResourceRequestWithDefaults instantiates a new CreateResourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateResourceRequestWithDefaults() *CreateResourceRequest {
	this := CreateResourceRequest{}
	return &this
}

// GetContent returns the Content field value
func (o *CreateResourceRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CreateResourceRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CreateResourceRequest) SetContent(v string) {
	o.Content = v
}

// GetContentType returns the ContentType field value
func (o *CreateResourceRequest) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *CreateResourceRequest) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *CreateResourceRequest) SetContentType(v string) {
	o.ContentType = v
}

// GetName returns the Name field value
func (o *CreateResourceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateResourceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateResourceRequest) SetName(v string) {
	o.Name = v
}

// GetTrigger returns the Trigger field value
func (o *CreateResourceRequest) GetTrigger() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *CreateResourceRequest) GetTriggerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value
func (o *CreateResourceRequest) SetTrigger(v string) {
	o.Trigger = v
}

func (o CreateResourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateResourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["content"] = o.Content
	toSerialize["content_type"] = o.ContentType
	toSerialize["name"] = o.Name
	toSerialize["trigger"] = o.Trigger
	return toSerialize, nil
}

type NullableCreateResourceRequest struct {
	value *CreateResourceRequest
	isSet bool
}

func (v NullableCreateResourceRequest) Get() *CreateResourceRequest {
	return v.value
}

func (v *NullableCreateResourceRequest) Set(val *CreateResourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateResourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateResourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateResourceRequest(val *CreateResourceRequest) *NullableCreateResourceRequest {
	return &NullableCreateResourceRequest{value: val, isSet: true}
}

func (v NullableCreateResourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateResourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


