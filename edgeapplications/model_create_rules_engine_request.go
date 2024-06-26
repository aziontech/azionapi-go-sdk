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

// checks if the CreateRulesEngineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRulesEngineRequest{}

// CreateRulesEngineRequest struct for CreateRulesEngineRequest
type CreateRulesEngineRequest struct {
	Name string `json:"name"`
	Order *int64 `json:"order,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	Description *string `json:"description,omitempty"`
	Criteria [][]RulesEngineCriteria `json:"criteria"`
	Behaviors []RulesEngineBehaviorEntry `json:"behaviors"`
}

type _CreateRulesEngineRequest CreateRulesEngineRequest

// NewCreateRulesEngineRequest instantiates a new CreateRulesEngineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRulesEngineRequest(name string, criteria [][]RulesEngineCriteria, behaviors []RulesEngineBehaviorEntry) *CreateRulesEngineRequest {
	this := CreateRulesEngineRequest{}
	this.Name = name
	this.Criteria = criteria
	this.Behaviors = behaviors
	return &this
}

// NewCreateRulesEngineRequestWithDefaults instantiates a new CreateRulesEngineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRulesEngineRequestWithDefaults() *CreateRulesEngineRequest {
	this := CreateRulesEngineRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateRulesEngineRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRulesEngineRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRulesEngineRequest) SetName(v string) {
	o.Name = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *CreateRulesEngineRequest) GetOrder() int64 {
	if o == nil || IsNil(o.Order) {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRulesEngineRequest) GetOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *CreateRulesEngineRequest) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *CreateRulesEngineRequest) SetOrder(v int64) {
	o.Order = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *CreateRulesEngineRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRulesEngineRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *CreateRulesEngineRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *CreateRulesEngineRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateRulesEngineRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRulesEngineRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateRulesEngineRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateRulesEngineRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCriteria returns the Criteria field value
func (o *CreateRulesEngineRequest) GetCriteria() [][]RulesEngineCriteria {
	if o == nil {
		var ret [][]RulesEngineCriteria
		return ret
	}

	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value
// and a boolean to check if the value has been set.
func (o *CreateRulesEngineRequest) GetCriteriaOk() ([][]RulesEngineCriteria, bool) {
	if o == nil {
		return nil, false
	}
	return o.Criteria, true
}

// SetCriteria sets field value
func (o *CreateRulesEngineRequest) SetCriteria(v [][]RulesEngineCriteria) {
	o.Criteria = v
}

// GetBehaviors returns the Behaviors field value
func (o *CreateRulesEngineRequest) GetBehaviors() []RulesEngineBehaviorEntry {
	if o == nil {
		var ret []RulesEngineBehaviorEntry
		return ret
	}

	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value
// and a boolean to check if the value has been set.
func (o *CreateRulesEngineRequest) GetBehaviorsOk() ([]RulesEngineBehaviorEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Behaviors, true
}

// SetBehaviors sets field value
func (o *CreateRulesEngineRequest) SetBehaviors(v []RulesEngineBehaviorEntry) {
	o.Behaviors = v
}

func (o CreateRulesEngineRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRulesEngineRequest) ToMap() (map[string]interface{}, error) {
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

func (o *CreateRulesEngineRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCreateRulesEngineRequest := _CreateRulesEngineRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRulesEngineRequest)

	if err != nil {
		return err
	}

	*o = CreateRulesEngineRequest(varCreateRulesEngineRequest)

	return err
}

type NullableCreateRulesEngineRequest struct {
	value *CreateRulesEngineRequest
	isSet bool
}

func (v NullableCreateRulesEngineRequest) Get() *CreateRulesEngineRequest {
	return v.value
}

func (v *NullableCreateRulesEngineRequest) Set(val *CreateRulesEngineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRulesEngineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRulesEngineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRulesEngineRequest(val *CreateRulesEngineRequest) *NullableCreateRulesEngineRequest {
	return &NullableCreateRulesEngineRequest{value: val, isSet: true}
}

func (v NullableCreateRulesEngineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRulesEngineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


