/*
Intelligent DNS

Azion Intelligent DNS API

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idns

import (
	"encoding/json"
)

// checks if the RecordPostOrPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordPostOrPut{}

// RecordPostOrPut struct for RecordPostOrPut
type RecordPostOrPut struct {
	Id *int32 `json:"id,omitempty"`
	Entry *string `json:"entry,omitempty"`
	Description *string `json:"description,omitempty"`
	AnswersList []string `json:"answers_list,omitempty"`
	Policy *string `json:"policy,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	RecordType *string `json:"record_type,omitempty"`
	Ttl *int32 `json:"ttl,omitempty"`
}

// NewRecordPostOrPut instantiates a new RecordPostOrPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordPostOrPut() *RecordPostOrPut {
	this := RecordPostOrPut{}
	return &this
}

// NewRecordPostOrPutWithDefaults instantiates a new RecordPostOrPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordPostOrPutWithDefaults() *RecordPostOrPut {
	this := RecordPostOrPut{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RecordPostOrPut) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPostOrPut) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RecordPostOrPut) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *RecordPostOrPut) SetId(v int32) {
	o.Id = &v
}

// GetEntry returns the Entry field value if set, zero value otherwise.
func (o *RecordPostOrPut) GetEntry() string {
	if o == nil || IsNil(o.Entry) {
		var ret string
		return ret
	}
	return *o.Entry
}

// GetEntryOk returns a tuple with the Entry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPostOrPut) GetEntryOk() (*string, bool) {
	if o == nil || IsNil(o.Entry) {
		return nil, false
	}
	return o.Entry, true
}

// HasEntry returns a boolean if a field has been set.
func (o *RecordPostOrPut) HasEntry() bool {
	if o != nil && !IsNil(o.Entry) {
		return true
	}

	return false
}

// SetEntry gets a reference to the given string and assigns it to the Entry field.
func (o *RecordPostOrPut) SetEntry(v string) {
	o.Entry = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RecordPostOrPut) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPostOrPut) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RecordPostOrPut) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RecordPostOrPut) SetDescription(v string) {
	o.Description = &v
}

// GetAnswersList returns the AnswersList field value if set, zero value otherwise.
func (o *RecordPostOrPut) GetAnswersList() []string {
	if o == nil || IsNil(o.AnswersList) {
		var ret []string
		return ret
	}
	return o.AnswersList
}

// GetAnswersListOk returns a tuple with the AnswersList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPostOrPut) GetAnswersListOk() ([]string, bool) {
	if o == nil || IsNil(o.AnswersList) {
		return nil, false
	}
	return o.AnswersList, true
}

// HasAnswersList returns a boolean if a field has been set.
func (o *RecordPostOrPut) HasAnswersList() bool {
	if o != nil && !IsNil(o.AnswersList) {
		return true
	}

	return false
}

// SetAnswersList gets a reference to the given []string and assigns it to the AnswersList field.
func (o *RecordPostOrPut) SetAnswersList(v []string) {
	o.AnswersList = v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *RecordPostOrPut) GetPolicy() string {
	if o == nil || IsNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPostOrPut) GetPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *RecordPostOrPut) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *RecordPostOrPut) SetPolicy(v string) {
	o.Policy = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *RecordPostOrPut) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPostOrPut) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *RecordPostOrPut) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *RecordPostOrPut) SetWeight(v int32) {
	o.Weight = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *RecordPostOrPut) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPostOrPut) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *RecordPostOrPut) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *RecordPostOrPut) SetRecordType(v string) {
	o.RecordType = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *RecordPostOrPut) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordPostOrPut) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *RecordPostOrPut) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *RecordPostOrPut) SetTtl(v int32) {
	o.Ttl = &v
}

func (o RecordPostOrPut) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordPostOrPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Entry) {
		toSerialize["entry"] = o.Entry
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AnswersList) {
		toSerialize["answers_list"] = o.AnswersList
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	return toSerialize, nil
}

type NullableRecordPostOrPut struct {
	value *RecordPostOrPut
	isSet bool
}

func (v NullableRecordPostOrPut) Get() *RecordPostOrPut {
	return v.value
}

func (v *NullableRecordPostOrPut) Set(val *RecordPostOrPut) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordPostOrPut) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordPostOrPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordPostOrPut(val *RecordPostOrPut) *NullableRecordPostOrPut {
	return &NullableRecordPostOrPut{value: val, isSet: true}
}

func (v NullableRecordPostOrPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordPostOrPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

