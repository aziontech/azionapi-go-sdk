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

// checks if the RulesEngineResultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulesEngineResultResponse{}

// RulesEngineResultResponse struct for RulesEngineResultResponse
type RulesEngineResultResponse struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Phase string `json:"phase"`
	Behaviors []RulesEngineResultResponseBehaviors `json:"behaviors,omitempty"`
	Criteria [][]RulesEngineCriteria `json:"criteria"`
	IsActive bool `json:"is_active"`
	Order int64 `json:"order"`
}

// NewRulesEngineResultResponse instantiates a new RulesEngineResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesEngineResultResponse(id int64, name string, phase string, criteria [][]RulesEngineCriteria, isActive bool, order int64) *RulesEngineResultResponse {
	this := RulesEngineResultResponse{}
	this.Id = id
	this.Name = name
	this.Phase = phase
	this.Criteria = criteria
	this.IsActive = isActive
	this.Order = order
	return &this
}

// NewRulesEngineResultResponseWithDefaults instantiates a new RulesEngineResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesEngineResultResponseWithDefaults() *RulesEngineResultResponse {
	this := RulesEngineResultResponse{}
	return &this
}

// GetId returns the Id field value
func (o *RulesEngineResultResponse) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RulesEngineResultResponse) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RulesEngineResultResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RulesEngineResultResponse) SetName(v string) {
	o.Name = v
}

// GetPhase returns the Phase field value
func (o *RulesEngineResultResponse) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *RulesEngineResultResponse) SetPhase(v string) {
	o.Phase = v
}

// GetBehaviors returns the Behaviors field value if set, zero value otherwise.
func (o *RulesEngineResultResponse) GetBehaviors() []RulesEngineResultResponseBehaviors {
	if o == nil || IsNil(o.Behaviors) {
		var ret []RulesEngineResultResponseBehaviors
		return ret
	}
	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetBehaviorsOk() ([]RulesEngineResultResponseBehaviors, bool) {
	if o == nil || IsNil(o.Behaviors) {
		return nil, false
	}
	return o.Behaviors, true
}

// HasBehaviors returns a boolean if a field has been set.
func (o *RulesEngineResultResponse) HasBehaviors() bool {
	if o != nil && !IsNil(o.Behaviors) {
		return true
	}

	return false
}

// SetBehaviors gets a reference to the given []RulesEngineResultResponseBehaviors and assigns it to the Behaviors field.
func (o *RulesEngineResultResponse) SetBehaviors(v []RulesEngineResultResponseBehaviors) {
	o.Behaviors = v
}

// GetCriteria returns the Criteria field value
func (o *RulesEngineResultResponse) GetCriteria() [][]RulesEngineCriteria {
	if o == nil {
		var ret [][]RulesEngineCriteria
		return ret
	}

	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetCriteriaOk() ([][]RulesEngineCriteria, bool) {
	if o == nil {
		return nil, false
	}
	return o.Criteria, true
}

// SetCriteria sets field value
func (o *RulesEngineResultResponse) SetCriteria(v [][]RulesEngineCriteria) {
	o.Criteria = v
}

// GetIsActive returns the IsActive field value
func (o *RulesEngineResultResponse) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *RulesEngineResultResponse) SetIsActive(v bool) {
	o.IsActive = v
}

// GetOrder returns the Order field value
func (o *RulesEngineResultResponse) GetOrder() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *RulesEngineResultResponse) SetOrder(v int64) {
	o.Order = v
}

func (o RulesEngineResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulesEngineResultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["phase"] = o.Phase
	if !IsNil(o.Behaviors) {
		toSerialize["behaviors"] = o.Behaviors
	}
	toSerialize["criteria"] = o.Criteria
	toSerialize["is_active"] = o.IsActive
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

type NullableRulesEngineResultResponse struct {
	value *RulesEngineResultResponse
	isSet bool
}

func (v NullableRulesEngineResultResponse) Get() *RulesEngineResultResponse {
	return v.value
}

func (v *NullableRulesEngineResultResponse) Set(val *RulesEngineResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesEngineResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesEngineResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesEngineResultResponse(val *RulesEngineResultResponse) *NullableRulesEngineResultResponse {
	return &NullableRulesEngineResultResponse{value: val, isSet: true}
}

func (v NullableRulesEngineResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesEngineResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


