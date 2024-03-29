/*
Personal Tokens - OpenAPI

The Personal Tokens API allows you to manage your existing personal tokens. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package personal_tokens

import (
	"encoding/json"
	"time"
)

// checks if the CreatePersonalTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePersonalTokenResponse{}

// CreatePersonalTokenResponse struct for CreatePersonalTokenResponse
type CreatePersonalTokenResponse struct {
	Uuid *string `json:"uuid,omitempty"`
	Name *string `json:"name,omitempty"`
	Key *string `json:"key,omitempty"`
	UserId *float32 `json:"user_id,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	Description NullableString `json:"description,omitempty"`
}

// NewCreatePersonalTokenResponse instantiates a new CreatePersonalTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePersonalTokenResponse() *CreatePersonalTokenResponse {
	this := CreatePersonalTokenResponse{}
	return &this
}

// NewCreatePersonalTokenResponseWithDefaults instantiates a new CreatePersonalTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePersonalTokenResponseWithDefaults() *CreatePersonalTokenResponse {
	this := CreatePersonalTokenResponse{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *CreatePersonalTokenResponse) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePersonalTokenResponse) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *CreatePersonalTokenResponse) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *CreatePersonalTokenResponse) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreatePersonalTokenResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePersonalTokenResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreatePersonalTokenResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreatePersonalTokenResponse) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CreatePersonalTokenResponse) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePersonalTokenResponse) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CreatePersonalTokenResponse) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CreatePersonalTokenResponse) SetKey(v string) {
	o.Key = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CreatePersonalTokenResponse) GetUserId() float32 {
	if o == nil || IsNil(o.UserId) {
		var ret float32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePersonalTokenResponse) GetUserIdOk() (*float32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CreatePersonalTokenResponse) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given float32 and assigns it to the UserId field.
func (o *CreatePersonalTokenResponse) SetUserId(v float32) {
	o.UserId = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CreatePersonalTokenResponse) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePersonalTokenResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CreatePersonalTokenResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CreatePersonalTokenResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CreatePersonalTokenResponse) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePersonalTokenResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CreatePersonalTokenResponse) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *CreatePersonalTokenResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePersonalTokenResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePersonalTokenResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreatePersonalTokenResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreatePersonalTokenResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreatePersonalTokenResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreatePersonalTokenResponse) UnsetDescription() {
	o.Description.Unset()
}

func (o CreatePersonalTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePersonalTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableCreatePersonalTokenResponse struct {
	value *CreatePersonalTokenResponse
	isSet bool
}

func (v NullableCreatePersonalTokenResponse) Get() *CreatePersonalTokenResponse {
	return v.value
}

func (v *NullableCreatePersonalTokenResponse) Set(val *CreatePersonalTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePersonalTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePersonalTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePersonalTokenResponse(val *CreatePersonalTokenResponse) *NullableCreatePersonalTokenResponse {
	return &NullableCreatePersonalTokenResponse{value: val, isSet: true}
}

func (v NullableCreatePersonalTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePersonalTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


