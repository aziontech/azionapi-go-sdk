/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateRulesEngineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRulesEngineRequest{}

// UpdateRulesEngineRequest struct for UpdateRulesEngineRequest
type UpdateRulesEngineRequest struct {
	Name string `json:"name"`
	Order *int64 `json:"order,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	Description *string `json:"description,omitempty"`
	Criteria [][]RulesEngineCriteria `json:"criteria"`
	Behaviors []RulesEngineBehaviorEntry `json:"behaviors"`
}

type _UpdateRulesEngineRequest UpdateRulesEngineRequest

// NewUpdateRulesEngineRequest instantiates a new UpdateRulesEngineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRulesEngineRequest(name string, criteria [][]RulesEngineCriteria, behaviors []RulesEngineBehaviorEntry) *UpdateRulesEngineRequest {
	this := UpdateRulesEngineRequest{}
	this.Name = name
	this.Criteria = criteria
	this.Behaviors = behaviors
	return &this
}

// NewUpdateRulesEngineRequestWithDefaults instantiates a new UpdateRulesEngineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRulesEngineRequestWithDefaults() *UpdateRulesEngineRequest {
	this := UpdateRulesEngineRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateRulesEngineRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesEngineRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateRulesEngineRequest) SetName(v string) {
	o.Name = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *UpdateRulesEngineRequest) GetOrder() int64 {
	if o == nil || IsNil(o.Order) {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRulesEngineRequest) GetOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *UpdateRulesEngineRequest) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *UpdateRulesEngineRequest) SetOrder(v int64) {
	o.Order = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UpdateRulesEngineRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRulesEngineRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UpdateRulesEngineRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UpdateRulesEngineRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateRulesEngineRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRulesEngineRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateRulesEngineRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateRulesEngineRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCriteria returns the Criteria field value
func (o *UpdateRulesEngineRequest) GetCriteria() [][]RulesEngineCriteria {
	if o == nil {
		var ret [][]RulesEngineCriteria
		return ret
	}

	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesEngineRequest) GetCriteriaOk() ([][]RulesEngineCriteria, bool) {
	if o == nil {
		return nil, false
	}
	return o.Criteria, true
}

// SetCriteria sets field value
func (o *UpdateRulesEngineRequest) SetCriteria(v [][]RulesEngineCriteria) {
	o.Criteria = v
}

// GetBehaviors returns the Behaviors field value
func (o *UpdateRulesEngineRequest) GetBehaviors() []RulesEngineBehaviorEntry {
	if o == nil {
		var ret []RulesEngineBehaviorEntry
		return ret
	}

	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesEngineRequest) GetBehaviorsOk() ([]RulesEngineBehaviorEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Behaviors, true
}

// SetBehaviors sets field value
func (o *UpdateRulesEngineRequest) SetBehaviors(v []RulesEngineBehaviorEntry) {
	o.Behaviors = v
}

func (o UpdateRulesEngineRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRulesEngineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["criteria"] = o.Criteria
	toSerialize["behaviors"] = o.Behaviors
	return toSerialize, nil
}

func (o *UpdateRulesEngineRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"criteria",
		"behaviors",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateRulesEngineRequest := _UpdateRulesEngineRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateRulesEngineRequest)

	if err != nil {
		return err
	}

	*o = UpdateRulesEngineRequest(varUpdateRulesEngineRequest)

	return err
}

type NullableUpdateRulesEngineRequest struct {
	value *UpdateRulesEngineRequest
	isSet bool
}

func (v NullableUpdateRulesEngineRequest) Get() *UpdateRulesEngineRequest {
	return v.value
}

func (v *NullableUpdateRulesEngineRequest) Set(val *UpdateRulesEngineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRulesEngineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRulesEngineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRulesEngineRequest(val *UpdateRulesEngineRequest) *NullableUpdateRulesEngineRequest {
	return &NullableUpdateRulesEngineRequest{value: val, isSet: true}
}

func (v NullableUpdateRulesEngineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRulesEngineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


