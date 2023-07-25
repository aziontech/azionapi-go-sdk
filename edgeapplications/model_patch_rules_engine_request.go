/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
)

// checks if the PatchRulesEngineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchRulesEngineRequest{}

// PatchRulesEngineRequest struct for PatchRulesEngineRequest
type PatchRulesEngineRequest struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Criteria [][]RulesEngineCriteria `json:"criteria,omitempty"`
	Behaviors []RulesEngineBehavior `json:"behaviors,omitempty"`
}

// NewPatchRulesEngineRequest instantiates a new PatchRulesEngineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchRulesEngineRequest() *PatchRulesEngineRequest {
	this := PatchRulesEngineRequest{}
	return &this
}

// NewPatchRulesEngineRequestWithDefaults instantiates a new PatchRulesEngineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchRulesEngineRequestWithDefaults() *PatchRulesEngineRequest {
	this := PatchRulesEngineRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchRulesEngineRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchRulesEngineRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchRulesEngineRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchRulesEngineRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchRulesEngineRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchRulesEngineRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchRulesEngineRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchRulesEngineRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *PatchRulesEngineRequest) GetCriteria() [][]RulesEngineCriteria {
	if o == nil || IsNil(o.Criteria) {
		var ret [][]RulesEngineCriteria
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchRulesEngineRequest) GetCriteriaOk() ([][]RulesEngineCriteria, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *PatchRulesEngineRequest) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given [][]RulesEngineCriteria and assigns it to the Criteria field.
func (o *PatchRulesEngineRequest) SetCriteria(v [][]RulesEngineCriteria) {
	o.Criteria = v
}

// GetBehaviors returns the Behaviors field value if set, zero value otherwise.
func (o *PatchRulesEngineRequest) GetBehaviors() []RulesEngineBehavior {
	if o == nil || IsNil(o.Behaviors) {
		var ret []RulesEngineBehavior
		return ret
	}
	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchRulesEngineRequest) GetBehaviorsOk() ([]RulesEngineBehavior, bool) {
	if o == nil || IsNil(o.Behaviors) {
		return nil, false
	}
	return o.Behaviors, true
}

// HasBehaviors returns a boolean if a field has been set.
func (o *PatchRulesEngineRequest) HasBehaviors() bool {
	if o != nil && !IsNil(o.Behaviors) {
		return true
	}

	return false
}

// SetBehaviors gets a reference to the given []RulesEngineBehavior and assigns it to the Behaviors field.
func (o *PatchRulesEngineRequest) SetBehaviors(v []RulesEngineBehavior) {
	o.Behaviors = v
}

func (o PatchRulesEngineRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchRulesEngineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Behaviors) {
		toSerialize["behaviors"] = o.Behaviors
	}
	return toSerialize, nil
}

type NullablePatchRulesEngineRequest struct {
	value *PatchRulesEngineRequest
	isSet bool
}

func (v NullablePatchRulesEngineRequest) Get() *PatchRulesEngineRequest {
	return v.value
}

func (v *NullablePatchRulesEngineRequest) Set(val *PatchRulesEngineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchRulesEngineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchRulesEngineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchRulesEngineRequest(val *PatchRulesEngineRequest) *NullablePatchRulesEngineRequest {
	return &NullablePatchRulesEngineRequest{value: val, isSet: true}
}

func (v NullablePatchRulesEngineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchRulesEngineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


