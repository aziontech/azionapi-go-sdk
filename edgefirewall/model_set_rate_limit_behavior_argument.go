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

// checks if the SetRateLimitBehaviorArgument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetRateLimitBehaviorArgument{}

// SetRateLimitBehaviorArgument struct for SetRateLimitBehaviorArgument
type SetRateLimitBehaviorArgument struct {
	Type *string `json:"type,omitempty"`
	LimitBy *string `json:"limit_by,omitempty"`
	AverageRateLimit *SetRateLimitBehaviorArgumentAverageRateLimit `json:"average_rate_limit,omitempty"`
	MaximumBurstSize *SetRateLimitBehaviorArgumentAverageRateLimit `json:"maximum_burst_size,omitempty"`
}

// NewSetRateLimitBehaviorArgument instantiates a new SetRateLimitBehaviorArgument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetRateLimitBehaviorArgument() *SetRateLimitBehaviorArgument {
	this := SetRateLimitBehaviorArgument{}
	return &this
}

// NewSetRateLimitBehaviorArgumentWithDefaults instantiates a new SetRateLimitBehaviorArgument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetRateLimitBehaviorArgumentWithDefaults() *SetRateLimitBehaviorArgument {
	this := SetRateLimitBehaviorArgument{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SetRateLimitBehaviorArgument) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRateLimitBehaviorArgument) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SetRateLimitBehaviorArgument) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SetRateLimitBehaviorArgument) SetType(v string) {
	o.Type = &v
}

// GetLimitBy returns the LimitBy field value if set, zero value otherwise.
func (o *SetRateLimitBehaviorArgument) GetLimitBy() string {
	if o == nil || IsNil(o.LimitBy) {
		var ret string
		return ret
	}
	return *o.LimitBy
}

// GetLimitByOk returns a tuple with the LimitBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRateLimitBehaviorArgument) GetLimitByOk() (*string, bool) {
	if o == nil || IsNil(o.LimitBy) {
		return nil, false
	}
	return o.LimitBy, true
}

// HasLimitBy returns a boolean if a field has been set.
func (o *SetRateLimitBehaviorArgument) HasLimitBy() bool {
	if o != nil && !IsNil(o.LimitBy) {
		return true
	}

	return false
}

// SetLimitBy gets a reference to the given string and assigns it to the LimitBy field.
func (o *SetRateLimitBehaviorArgument) SetLimitBy(v string) {
	o.LimitBy = &v
}

// GetAverageRateLimit returns the AverageRateLimit field value if set, zero value otherwise.
func (o *SetRateLimitBehaviorArgument) GetAverageRateLimit() SetRateLimitBehaviorArgumentAverageRateLimit {
	if o == nil || IsNil(o.AverageRateLimit) {
		var ret SetRateLimitBehaviorArgumentAverageRateLimit
		return ret
	}
	return *o.AverageRateLimit
}

// GetAverageRateLimitOk returns a tuple with the AverageRateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRateLimitBehaviorArgument) GetAverageRateLimitOk() (*SetRateLimitBehaviorArgumentAverageRateLimit, bool) {
	if o == nil || IsNil(o.AverageRateLimit) {
		return nil, false
	}
	return o.AverageRateLimit, true
}

// HasAverageRateLimit returns a boolean if a field has been set.
func (o *SetRateLimitBehaviorArgument) HasAverageRateLimit() bool {
	if o != nil && !IsNil(o.AverageRateLimit) {
		return true
	}

	return false
}

// SetAverageRateLimit gets a reference to the given SetRateLimitBehaviorArgumentAverageRateLimit and assigns it to the AverageRateLimit field.
func (o *SetRateLimitBehaviorArgument) SetAverageRateLimit(v SetRateLimitBehaviorArgumentAverageRateLimit) {
	o.AverageRateLimit = &v
}

// GetMaximumBurstSize returns the MaximumBurstSize field value if set, zero value otherwise.
func (o *SetRateLimitBehaviorArgument) GetMaximumBurstSize() SetRateLimitBehaviorArgumentAverageRateLimit {
	if o == nil || IsNil(o.MaximumBurstSize) {
		var ret SetRateLimitBehaviorArgumentAverageRateLimit
		return ret
	}
	return *o.MaximumBurstSize
}

// GetMaximumBurstSizeOk returns a tuple with the MaximumBurstSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRateLimitBehaviorArgument) GetMaximumBurstSizeOk() (*SetRateLimitBehaviorArgumentAverageRateLimit, bool) {
	if o == nil || IsNil(o.MaximumBurstSize) {
		return nil, false
	}
	return o.MaximumBurstSize, true
}

// HasMaximumBurstSize returns a boolean if a field has been set.
func (o *SetRateLimitBehaviorArgument) HasMaximumBurstSize() bool {
	if o != nil && !IsNil(o.MaximumBurstSize) {
		return true
	}

	return false
}

// SetMaximumBurstSize gets a reference to the given SetRateLimitBehaviorArgumentAverageRateLimit and assigns it to the MaximumBurstSize field.
func (o *SetRateLimitBehaviorArgument) SetMaximumBurstSize(v SetRateLimitBehaviorArgumentAverageRateLimit) {
	o.MaximumBurstSize = &v
}

func (o SetRateLimitBehaviorArgument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetRateLimitBehaviorArgument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.LimitBy) {
		toSerialize["limit_by"] = o.LimitBy
	}
	if !IsNil(o.AverageRateLimit) {
		toSerialize["average_rate_limit"] = o.AverageRateLimit
	}
	if !IsNil(o.MaximumBurstSize) {
		toSerialize["maximum_burst_size"] = o.MaximumBurstSize
	}
	return toSerialize, nil
}

type NullableSetRateLimitBehaviorArgument struct {
	value *SetRateLimitBehaviorArgument
	isSet bool
}

func (v NullableSetRateLimitBehaviorArgument) Get() *SetRateLimitBehaviorArgument {
	return v.value
}

func (v *NullableSetRateLimitBehaviorArgument) Set(val *SetRateLimitBehaviorArgument) {
	v.value = val
	v.isSet = true
}

func (v NullableSetRateLimitBehaviorArgument) IsSet() bool {
	return v.isSet
}

func (v *NullableSetRateLimitBehaviorArgument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetRateLimitBehaviorArgument(val *SetRateLimitBehaviorArgument) *NullableSetRateLimitBehaviorArgument {
	return &NullableSetRateLimitBehaviorArgument{value: val, isSet: true}
}

func (v NullableSetRateLimitBehaviorArgument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetRateLimitBehaviorArgument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

