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

// checks if the CreateEdgeFirewallRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEdgeFirewallRequest{}

// CreateEdgeFirewallRequest struct for CreateEdgeFirewallRequest
type CreateEdgeFirewallRequest struct {
	Name *string `json:"name,omitempty"`
	Domains []int64 `json:"domains,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	EdgeFunctionsEnabled *bool `json:"edge_functions_enabled,omitempty"`
	NetworkProtectionEnabled *bool `json:"network_protection_enabled,omitempty"`
	WafEnabled *bool `json:"waf_enabled,omitempty"`
}

// NewCreateEdgeFirewallRequest instantiates a new CreateEdgeFirewallRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEdgeFirewallRequest() *CreateEdgeFirewallRequest {
	this := CreateEdgeFirewallRequest{}
	return &this
}

// NewCreateEdgeFirewallRequestWithDefaults instantiates a new CreateEdgeFirewallRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEdgeFirewallRequestWithDefaults() *CreateEdgeFirewallRequest {
	this := CreateEdgeFirewallRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateEdgeFirewallRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEdgeFirewallRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateEdgeFirewallRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateEdgeFirewallRequest) SetName(v string) {
	o.Name = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *CreateEdgeFirewallRequest) GetDomains() []int64 {
	if o == nil || IsNil(o.Domains) {
		var ret []int64
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEdgeFirewallRequest) GetDomainsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *CreateEdgeFirewallRequest) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []int64 and assigns it to the Domains field.
func (o *CreateEdgeFirewallRequest) SetDomains(v []int64) {
	o.Domains = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *CreateEdgeFirewallRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEdgeFirewallRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *CreateEdgeFirewallRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *CreateEdgeFirewallRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetEdgeFunctionsEnabled returns the EdgeFunctionsEnabled field value if set, zero value otherwise.
func (o *CreateEdgeFirewallRequest) GetEdgeFunctionsEnabled() bool {
	if o == nil || IsNil(o.EdgeFunctionsEnabled) {
		var ret bool
		return ret
	}
	return *o.EdgeFunctionsEnabled
}

// GetEdgeFunctionsEnabledOk returns a tuple with the EdgeFunctionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEdgeFirewallRequest) GetEdgeFunctionsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EdgeFunctionsEnabled) {
		return nil, false
	}
	return o.EdgeFunctionsEnabled, true
}

// HasEdgeFunctionsEnabled returns a boolean if a field has been set.
func (o *CreateEdgeFirewallRequest) HasEdgeFunctionsEnabled() bool {
	if o != nil && !IsNil(o.EdgeFunctionsEnabled) {
		return true
	}

	return false
}

// SetEdgeFunctionsEnabled gets a reference to the given bool and assigns it to the EdgeFunctionsEnabled field.
func (o *CreateEdgeFirewallRequest) SetEdgeFunctionsEnabled(v bool) {
	o.EdgeFunctionsEnabled = &v
}

// GetNetworkProtectionEnabled returns the NetworkProtectionEnabled field value if set, zero value otherwise.
func (o *CreateEdgeFirewallRequest) GetNetworkProtectionEnabled() bool {
	if o == nil || IsNil(o.NetworkProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.NetworkProtectionEnabled
}

// GetNetworkProtectionEnabledOk returns a tuple with the NetworkProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEdgeFirewallRequest) GetNetworkProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NetworkProtectionEnabled) {
		return nil, false
	}
	return o.NetworkProtectionEnabled, true
}

// HasNetworkProtectionEnabled returns a boolean if a field has been set.
func (o *CreateEdgeFirewallRequest) HasNetworkProtectionEnabled() bool {
	if o != nil && !IsNil(o.NetworkProtectionEnabled) {
		return true
	}

	return false
}

// SetNetworkProtectionEnabled gets a reference to the given bool and assigns it to the NetworkProtectionEnabled field.
func (o *CreateEdgeFirewallRequest) SetNetworkProtectionEnabled(v bool) {
	o.NetworkProtectionEnabled = &v
}

// GetWafEnabled returns the WafEnabled field value if set, zero value otherwise.
func (o *CreateEdgeFirewallRequest) GetWafEnabled() bool {
	if o == nil || IsNil(o.WafEnabled) {
		var ret bool
		return ret
	}
	return *o.WafEnabled
}

// GetWafEnabledOk returns a tuple with the WafEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEdgeFirewallRequest) GetWafEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WafEnabled) {
		return nil, false
	}
	return o.WafEnabled, true
}

// HasWafEnabled returns a boolean if a field has been set.
func (o *CreateEdgeFirewallRequest) HasWafEnabled() bool {
	if o != nil && !IsNil(o.WafEnabled) {
		return true
	}

	return false
}

// SetWafEnabled gets a reference to the given bool and assigns it to the WafEnabled field.
func (o *CreateEdgeFirewallRequest) SetWafEnabled(v bool) {
	o.WafEnabled = &v
}

func (o CreateEdgeFirewallRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEdgeFirewallRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.EdgeFunctionsEnabled) {
		toSerialize["edge_functions_enabled"] = o.EdgeFunctionsEnabled
	}
	if !IsNil(o.NetworkProtectionEnabled) {
		toSerialize["network_protection_enabled"] = o.NetworkProtectionEnabled
	}
	if !IsNil(o.WafEnabled) {
		toSerialize["waf_enabled"] = o.WafEnabled
	}
	return toSerialize, nil
}

type NullableCreateEdgeFirewallRequest struct {
	value *CreateEdgeFirewallRequest
	isSet bool
}

func (v NullableCreateEdgeFirewallRequest) Get() *CreateEdgeFirewallRequest {
	return v.value
}

func (v *NullableCreateEdgeFirewallRequest) Set(val *CreateEdgeFirewallRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEdgeFirewallRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEdgeFirewallRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEdgeFirewallRequest(val *CreateEdgeFirewallRequest) *NullableCreateEdgeFirewallRequest {
	return &NullableCreateEdgeFirewallRequest{value: val, isSet: true}
}

func (v NullableCreateEdgeFirewallRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEdgeFirewallRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


