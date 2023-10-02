/*
Edge Firewall API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgefirewall

import (
	"encoding/json"
)

// checks if the SSLVerificationStatusCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SSLVerificationStatusCriteria{}

// SSLVerificationStatusCriteria struct for SSLVerificationStatusCriteria
type SSLVerificationStatusCriteria struct {
	Variable *Variables `json:"variable,omitempty"`
	Operator *SSLVerificationStatusOperators `json:"operator,omitempty"`
	Conditional *Conditionals `json:"conditional,omitempty"`
	InputValue *SSLVerificationStatusArguments `json:"input_value,omitempty"`
}

// NewSSLVerificationStatusCriteria instantiates a new SSLVerificationStatusCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSLVerificationStatusCriteria() *SSLVerificationStatusCriteria {
	this := SSLVerificationStatusCriteria{}
	return &this
}

// NewSSLVerificationStatusCriteriaWithDefaults instantiates a new SSLVerificationStatusCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSLVerificationStatusCriteriaWithDefaults() *SSLVerificationStatusCriteria {
	this := SSLVerificationStatusCriteria{}
	return &this
}

// GetVariable returns the Variable field value if set, zero value otherwise.
func (o *SSLVerificationStatusCriteria) GetVariable() Variables {
	if o == nil || IsNil(o.Variable) {
		var ret Variables
		return ret
	}
	return *o.Variable
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLVerificationStatusCriteria) GetVariableOk() (*Variables, bool) {
	if o == nil || IsNil(o.Variable) {
		return nil, false
	}
	return o.Variable, true
}

// HasVariable returns a boolean if a field has been set.
func (o *SSLVerificationStatusCriteria) HasVariable() bool {
	if o != nil && !IsNil(o.Variable) {
		return true
	}

	return false
}

// SetVariable gets a reference to the given Variables and assigns it to the Variable field.
func (o *SSLVerificationStatusCriteria) SetVariable(v Variables) {
	o.Variable = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *SSLVerificationStatusCriteria) GetOperator() SSLVerificationStatusOperators {
	if o == nil || IsNil(o.Operator) {
		var ret SSLVerificationStatusOperators
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLVerificationStatusCriteria) GetOperatorOk() (*SSLVerificationStatusOperators, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *SSLVerificationStatusCriteria) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given SSLVerificationStatusOperators and assigns it to the Operator field.
func (o *SSLVerificationStatusCriteria) SetOperator(v SSLVerificationStatusOperators) {
	o.Operator = &v
}

// GetConditional returns the Conditional field value if set, zero value otherwise.
func (o *SSLVerificationStatusCriteria) GetConditional() Conditionals {
	if o == nil || IsNil(o.Conditional) {
		var ret Conditionals
		return ret
	}
	return *o.Conditional
}

// GetConditionalOk returns a tuple with the Conditional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLVerificationStatusCriteria) GetConditionalOk() (*Conditionals, bool) {
	if o == nil || IsNil(o.Conditional) {
		return nil, false
	}
	return o.Conditional, true
}

// HasConditional returns a boolean if a field has been set.
func (o *SSLVerificationStatusCriteria) HasConditional() bool {
	if o != nil && !IsNil(o.Conditional) {
		return true
	}

	return false
}

// SetConditional gets a reference to the given Conditionals and assigns it to the Conditional field.
func (o *SSLVerificationStatusCriteria) SetConditional(v Conditionals) {
	o.Conditional = &v
}

// GetInputValue returns the InputValue field value if set, zero value otherwise.
func (o *SSLVerificationStatusCriteria) GetInputValue() SSLVerificationStatusArguments {
	if o == nil || IsNil(o.InputValue) {
		var ret SSLVerificationStatusArguments
		return ret
	}
	return *o.InputValue
}

// GetInputValueOk returns a tuple with the InputValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSLVerificationStatusCriteria) GetInputValueOk() (*SSLVerificationStatusArguments, bool) {
	if o == nil || IsNil(o.InputValue) {
		return nil, false
	}
	return o.InputValue, true
}

// HasInputValue returns a boolean if a field has been set.
func (o *SSLVerificationStatusCriteria) HasInputValue() bool {
	if o != nil && !IsNil(o.InputValue) {
		return true
	}

	return false
}

// SetInputValue gets a reference to the given SSLVerificationStatusArguments and assigns it to the InputValue field.
func (o *SSLVerificationStatusCriteria) SetInputValue(v SSLVerificationStatusArguments) {
	o.InputValue = &v
}

func (o SSLVerificationStatusCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SSLVerificationStatusCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Variable) {
		toSerialize["variable"] = o.Variable
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Conditional) {
		toSerialize["conditional"] = o.Conditional
	}
	if !IsNil(o.InputValue) {
		toSerialize["input_value"] = o.InputValue
	}
	return toSerialize, nil
}

type NullableSSLVerificationStatusCriteria struct {
	value *SSLVerificationStatusCriteria
	isSet bool
}

func (v NullableSSLVerificationStatusCriteria) Get() *SSLVerificationStatusCriteria {
	return v.value
}

func (v *NullableSSLVerificationStatusCriteria) Set(val *SSLVerificationStatusCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableSSLVerificationStatusCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableSSLVerificationStatusCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSLVerificationStatusCriteria(val *SSLVerificationStatusCriteria) *NullableSSLVerificationStatusCriteria {
	return &NullableSSLVerificationStatusCriteria{value: val, isSet: true}
}

func (v NullableSSLVerificationStatusCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSLVerificationStatusCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


