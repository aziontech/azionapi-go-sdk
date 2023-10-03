/*
Edge Firewall API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgefirewall

import (
	"encoding/json"
	"time"
)

// checks if the RuleSetResultResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleSetResultResults{}

// RuleSetResultResults struct for RuleSetResultResults
type RuleSetResultResults struct {
	Id *int64 `json:"id,omitempty"`
	LastEditor *string `json:"last_editor,omitempty"`
	LastModified *time.Time `json:"last_modified,omitempty"`
	Name *string `json:"name,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	Description *string `json:"description,omitempty"`
	Behaviors []Behaviors `json:"behaviors,omitempty"`
	Criteria [][]SSLVerificationStatusCriteria `json:"criteria,omitempty"`
	Order *int32 `json:"order,omitempty"`
}

// NewRuleSetResultResults instantiates a new RuleSetResultResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleSetResultResults() *RuleSetResultResults {
	this := RuleSetResultResults{}
	return &this
}

// NewRuleSetResultResultsWithDefaults instantiates a new RuleSetResultResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleSetResultResultsWithDefaults() *RuleSetResultResults {
	this := RuleSetResultResults{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RuleSetResultResults) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSetResultResults) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RuleSetResultResults) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RuleSetResultResults) SetId(v int64) {
	o.Id = &v
}

// GetLastEditor returns the LastEditor field value if set, zero value otherwise.
func (o *RuleSetResultResults) GetLastEditor() string {
	if o == nil || IsNil(o.LastEditor) {
		var ret string
		return ret
	}
	return *o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSetResultResults) GetLastEditorOk() (*string, bool) {
	if o == nil || IsNil(o.LastEditor) {
		return nil, false
	}
	return o.LastEditor, true
}

// HasLastEditor returns a boolean if a field has been set.
func (o *RuleSetResultResults) HasLastEditor() bool {
	if o != nil && !IsNil(o.LastEditor) {
		return true
	}

	return false
}

// SetLastEditor gets a reference to the given string and assigns it to the LastEditor field.
func (o *RuleSetResultResults) SetLastEditor(v string) {
	o.LastEditor = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *RuleSetResultResults) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSetResultResults) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *RuleSetResultResults) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *RuleSetResultResults) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleSetResultResults) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSetResultResults) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleSetResultResults) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleSetResultResults) SetName(v string) {
	o.Name = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *RuleSetResultResults) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSetResultResults) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *RuleSetResultResults) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *RuleSetResultResults) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RuleSetResultResults) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSetResultResults) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RuleSetResultResults) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RuleSetResultResults) SetDescription(v string) {
	o.Description = &v
}

// GetBehaviors returns the Behaviors field value if set, zero value otherwise.
func (o *RuleSetResultResults) GetBehaviors() []Behaviors {
	if o == nil || IsNil(o.Behaviors) {
		var ret []Behaviors
		return ret
	}
	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSetResultResults) GetBehaviorsOk() ([]Behaviors, bool) {
	if o == nil || IsNil(o.Behaviors) {
		return nil, false
	}
	return o.Behaviors, true
}

// HasBehaviors returns a boolean if a field has been set.
func (o *RuleSetResultResults) HasBehaviors() bool {
	if o != nil && !IsNil(o.Behaviors) {
		return true
	}

	return false
}

// SetBehaviors gets a reference to the given []Behaviors and assigns it to the Behaviors field.
func (o *RuleSetResultResults) SetBehaviors(v []Behaviors) {
	o.Behaviors = v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *RuleSetResultResults) GetCriteria() [][]SSLVerificationStatusCriteria {
	if o == nil || IsNil(o.Criteria) {
		var ret [][]SSLVerificationStatusCriteria
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSetResultResults) GetCriteriaOk() ([][]SSLVerificationStatusCriteria, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *RuleSetResultResults) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given [][]SSLVerificationStatusCriteria and assigns it to the Criteria field.
func (o *RuleSetResultResults) SetCriteria(v [][]SSLVerificationStatusCriteria) {
	o.Criteria = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RuleSetResultResults) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleSetResultResults) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RuleSetResultResults) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *RuleSetResultResults) SetOrder(v int32) {
	o.Order = &v
}

func (o RuleSetResultResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleSetResultResults) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Behaviors) {
		toSerialize["behaviors"] = o.Behaviors
	}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullableRuleSetResultResults struct {
	value *RuleSetResultResults
	isSet bool
}

func (v NullableRuleSetResultResults) Get() *RuleSetResultResults {
	return v.value
}

func (v *NullableRuleSetResultResults) Set(val *RuleSetResultResults) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleSetResultResults) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleSetResultResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleSetResultResults(val *RuleSetResultResults) *NullableRuleSetResultResults {
	return &NullableRuleSetResultResults{value: val, isSet: true}
}

func (v NullableRuleSetResultResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleSetResultResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


