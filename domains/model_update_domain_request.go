/*
Domain API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package domains

import (
	"encoding/json"
)

// checks if the UpdateDomainRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDomainRequest{}

// UpdateDomainRequest struct for UpdateDomainRequest
type UpdateDomainRequest struct {
	Name *string `json:"name,omitempty"`
	Cnames []string `json:"cnames,omitempty"`
	CnameAccessOnly *bool `json:"cname_access_only,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	EdgeApplicationId *int64 `json:"edge_application_id,omitempty"`
	DigitalCertificateId *DomainDataDigitalCertificateId `json:"digital_certificate_id,omitempty"`
	Environment *string `json:"environment,omitempty"`
	IsMtlsEnabled *bool `json:"is_mtls_enabled,omitempty"`
	MtlsTrustedCaCertificateId NullableInt64 `json:"mtls_trusted_ca_certificate_id,omitempty"`
	EdgeFirewallId NullableInt64 `json:"edge_firewall_id,omitempty"`
	MtlsVerification *string `json:"mtls_verification,omitempty"`
	CrlList []int64 `json:"crl_list,omitempty"`
}

// NewUpdateDomainRequest instantiates a new UpdateDomainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDomainRequest() *UpdateDomainRequest {
	this := UpdateDomainRequest{}
	return &this
}

// NewUpdateDomainRequestWithDefaults instantiates a new UpdateDomainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDomainRequestWithDefaults() *UpdateDomainRequest {
	this := UpdateDomainRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateDomainRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateDomainRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateDomainRequest) SetName(v string) {
	o.Name = &v
}

// GetCnames returns the Cnames field value if set, zero value otherwise.
func (o *UpdateDomainRequest) GetCnames() []string {
	if o == nil || IsNil(o.Cnames) {
		var ret []string
		return ret
	}
	return o.Cnames
}

// GetCnamesOk returns a tuple with the Cnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequest) GetCnamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Cnames) {
		return nil, false
	}
	return o.Cnames, true
}

// HasCnames returns a boolean if a field has been set.
func (o *UpdateDomainRequest) HasCnames() bool {
	if o != nil && !IsNil(o.Cnames) {
		return true
	}

	return false
}

// SetCnames gets a reference to the given []string and assigns it to the Cnames field.
func (o *UpdateDomainRequest) SetCnames(v []string) {
	o.Cnames = v
}

// GetCnameAccessOnly returns the CnameAccessOnly field value if set, zero value otherwise.
func (o *UpdateDomainRequest) GetCnameAccessOnly() bool {
	if o == nil || IsNil(o.CnameAccessOnly) {
		var ret bool
		return ret
	}
	return *o.CnameAccessOnly
}

// GetCnameAccessOnlyOk returns a tuple with the CnameAccessOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequest) GetCnameAccessOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.CnameAccessOnly) {
		return nil, false
	}
	return o.CnameAccessOnly, true
}

// HasCnameAccessOnly returns a boolean if a field has been set.
func (o *UpdateDomainRequest) HasCnameAccessOnly() bool {
	if o != nil && !IsNil(o.CnameAccessOnly) {
		return true
	}

	return false
}

// SetCnameAccessOnly gets a reference to the given bool and assigns it to the CnameAccessOnly field.
func (o *UpdateDomainRequest) SetCnameAccessOnly(v bool) {
	o.CnameAccessOnly = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UpdateDomainRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UpdateDomainRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UpdateDomainRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetEdgeApplicationId returns the EdgeApplicationId field value if set, zero value otherwise.
func (o *UpdateDomainRequest) GetEdgeApplicationId() int64 {
	if o == nil || IsNil(o.EdgeApplicationId) {
		var ret int64
		return ret
	}
	return *o.EdgeApplicationId
}

// GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequest) GetEdgeApplicationIdOk() (*int64, bool) {
	if o == nil || IsNil(o.EdgeApplicationId) {
		return nil, false
	}
	return o.EdgeApplicationId, true
}

// HasEdgeApplicationId returns a boolean if a field has been set.
func (o *UpdateDomainRequest) HasEdgeApplicationId() bool {
	if o != nil && !IsNil(o.EdgeApplicationId) {
		return true
	}

	return false
}

// SetEdgeApplicationId gets a reference to the given int64 and assigns it to the EdgeApplicationId field.
func (o *UpdateDomainRequest) SetEdgeApplicationId(v int64) {
	o.EdgeApplicationId = &v
}

// GetDigitalCertificateId returns the DigitalCertificateId field value if set, zero value otherwise.
func (o *UpdateDomainRequest) GetDigitalCertificateId() DomainDataDigitalCertificateId {
	if o == nil || IsNil(o.DigitalCertificateId) {
		var ret DomainDataDigitalCertificateId
		return ret
	}
	return *o.DigitalCertificateId
}

// GetDigitalCertificateIdOk returns a tuple with the DigitalCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequest) GetDigitalCertificateIdOk() (*DomainDataDigitalCertificateId, bool) {
	if o == nil || IsNil(o.DigitalCertificateId) {
		return nil, false
	}
	return o.DigitalCertificateId, true
}

// HasDigitalCertificateId returns a boolean if a field has been set.
func (o *UpdateDomainRequest) HasDigitalCertificateId() bool {
	if o != nil && !IsNil(o.DigitalCertificateId) {
		return true
	}

	return false
}

// SetDigitalCertificateId gets a reference to the given DomainDataDigitalCertificateId and assigns it to the DigitalCertificateId field.
func (o *UpdateDomainRequest) SetDigitalCertificateId(v DomainDataDigitalCertificateId) {
	o.DigitalCertificateId = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *UpdateDomainRequest) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *UpdateDomainRequest) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *UpdateDomainRequest) SetEnvironment(v string) {
	o.Environment = &v
}

// GetIsMtlsEnabled returns the IsMtlsEnabled field value if set, zero value otherwise.
func (o *UpdateDomainRequest) GetIsMtlsEnabled() bool {
	if o == nil || IsNil(o.IsMtlsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsMtlsEnabled
}

// GetIsMtlsEnabledOk returns a tuple with the IsMtlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequest) GetIsMtlsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMtlsEnabled) {
		return nil, false
	}
	return o.IsMtlsEnabled, true
}

// HasIsMtlsEnabled returns a boolean if a field has been set.
func (o *UpdateDomainRequest) HasIsMtlsEnabled() bool {
	if o != nil && !IsNil(o.IsMtlsEnabled) {
		return true
	}

	return false
}

// SetIsMtlsEnabled gets a reference to the given bool and assigns it to the IsMtlsEnabled field.
func (o *UpdateDomainRequest) SetIsMtlsEnabled(v bool) {
	o.IsMtlsEnabled = &v
}

// GetMtlsTrustedCaCertificateId returns the MtlsTrustedCaCertificateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDomainRequest) GetMtlsTrustedCaCertificateId() int64 {
	if o == nil || IsNil(o.MtlsTrustedCaCertificateId.Get()) {
		var ret int64
		return ret
	}
	return *o.MtlsTrustedCaCertificateId.Get()
}

// GetMtlsTrustedCaCertificateIdOk returns a tuple with the MtlsTrustedCaCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDomainRequest) GetMtlsTrustedCaCertificateIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MtlsTrustedCaCertificateId.Get(), o.MtlsTrustedCaCertificateId.IsSet()
}

// HasMtlsTrustedCaCertificateId returns a boolean if a field has been set.
func (o *UpdateDomainRequest) HasMtlsTrustedCaCertificateId() bool {
	if o != nil && o.MtlsTrustedCaCertificateId.IsSet() {
		return true
	}

	return false
}

// SetMtlsTrustedCaCertificateId gets a reference to the given NullableInt64 and assigns it to the MtlsTrustedCaCertificateId field.
func (o *UpdateDomainRequest) SetMtlsTrustedCaCertificateId(v int64) {
	o.MtlsTrustedCaCertificateId.Set(&v)
}
// SetMtlsTrustedCaCertificateIdNil sets the value for MtlsTrustedCaCertificateId to be an explicit nil
func (o *UpdateDomainRequest) SetMtlsTrustedCaCertificateIdNil() {
	o.MtlsTrustedCaCertificateId.Set(nil)
}

// UnsetMtlsTrustedCaCertificateId ensures that no value is present for MtlsTrustedCaCertificateId, not even an explicit nil
func (o *UpdateDomainRequest) UnsetMtlsTrustedCaCertificateId() {
	o.MtlsTrustedCaCertificateId.Unset()
}

// GetEdgeFirewallId returns the EdgeFirewallId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDomainRequest) GetEdgeFirewallId() int64 {
	if o == nil || IsNil(o.EdgeFirewallId.Get()) {
		var ret int64
		return ret
	}
	return *o.EdgeFirewallId.Get()
}

// GetEdgeFirewallIdOk returns a tuple with the EdgeFirewallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDomainRequest) GetEdgeFirewallIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EdgeFirewallId.Get(), o.EdgeFirewallId.IsSet()
}

// HasEdgeFirewallId returns a boolean if a field has been set.
func (o *UpdateDomainRequest) HasEdgeFirewallId() bool {
	if o != nil && o.EdgeFirewallId.IsSet() {
		return true
	}

	return false
}

// SetEdgeFirewallId gets a reference to the given NullableInt64 and assigns it to the EdgeFirewallId field.
func (o *UpdateDomainRequest) SetEdgeFirewallId(v int64) {
	o.EdgeFirewallId.Set(&v)
}
// SetEdgeFirewallIdNil sets the value for EdgeFirewallId to be an explicit nil
func (o *UpdateDomainRequest) SetEdgeFirewallIdNil() {
	o.EdgeFirewallId.Set(nil)
}

// UnsetEdgeFirewallId ensures that no value is present for EdgeFirewallId, not even an explicit nil
func (o *UpdateDomainRequest) UnsetEdgeFirewallId() {
	o.EdgeFirewallId.Unset()
}

// GetMtlsVerification returns the MtlsVerification field value if set, zero value otherwise.
func (o *UpdateDomainRequest) GetMtlsVerification() string {
	if o == nil || IsNil(o.MtlsVerification) {
		var ret string
		return ret
	}
	return *o.MtlsVerification
}

// GetMtlsVerificationOk returns a tuple with the MtlsVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDomainRequest) GetMtlsVerificationOk() (*string, bool) {
	if o == nil || IsNil(o.MtlsVerification) {
		return nil, false
	}
	return o.MtlsVerification, true
}

// HasMtlsVerification returns a boolean if a field has been set.
func (o *UpdateDomainRequest) HasMtlsVerification() bool {
	if o != nil && !IsNil(o.MtlsVerification) {
		return true
	}

	return false
}

// SetMtlsVerification gets a reference to the given string and assigns it to the MtlsVerification field.
func (o *UpdateDomainRequest) SetMtlsVerification(v string) {
	o.MtlsVerification = &v
}

// GetCrlList returns the CrlList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDomainRequest) GetCrlList() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.CrlList
}

// GetCrlListOk returns a tuple with the CrlList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDomainRequest) GetCrlListOk() ([]int64, bool) {
	if o == nil || IsNil(o.CrlList) {
		return nil, false
	}
	return o.CrlList, true
}

// HasCrlList returns a boolean if a field has been set.
func (o *UpdateDomainRequest) HasCrlList() bool {
	if o != nil && !IsNil(o.CrlList) {
		return true
	}

	return false
}

// SetCrlList gets a reference to the given []int64 and assigns it to the CrlList field.
func (o *UpdateDomainRequest) SetCrlList(v []int64) {
	o.CrlList = v
}

func (o UpdateDomainRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDomainRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Cnames) {
		toSerialize["cnames"] = o.Cnames
	}
	if !IsNil(o.CnameAccessOnly) {
		toSerialize["cname_access_only"] = o.CnameAccessOnly
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.EdgeApplicationId) {
		toSerialize["edge_application_id"] = o.EdgeApplicationId
	}
	if !IsNil(o.DigitalCertificateId) {
		toSerialize["digital_certificate_id"] = o.DigitalCertificateId
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.IsMtlsEnabled) {
		toSerialize["is_mtls_enabled"] = o.IsMtlsEnabled
	}
	if o.MtlsTrustedCaCertificateId.IsSet() {
		toSerialize["mtls_trusted_ca_certificate_id"] = o.MtlsTrustedCaCertificateId.Get()
	}
	if o.EdgeFirewallId.IsSet() {
		toSerialize["edge_firewall_id"] = o.EdgeFirewallId.Get()
	}
	if !IsNil(o.MtlsVerification) {
		toSerialize["mtls_verification"] = o.MtlsVerification
	}
	if o.CrlList != nil {
		toSerialize["crl_list"] = o.CrlList
	}
	return toSerialize, nil
}

type NullableUpdateDomainRequest struct {
	value *UpdateDomainRequest
	isSet bool
}

func (v NullableUpdateDomainRequest) Get() *UpdateDomainRequest {
	return v.value
}

func (v *NullableUpdateDomainRequest) Set(val *UpdateDomainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDomainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDomainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDomainRequest(val *UpdateDomainRequest) *NullableUpdateDomainRequest {
	return &NullableUpdateDomainRequest{value: val, isSet: true}
}

func (v NullableUpdateDomainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDomainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


