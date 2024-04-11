/*
Domain API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package domains

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateDomainRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDomainRequest{}

// CreateDomainRequest struct for CreateDomainRequest
type CreateDomainRequest struct {
	Name string `json:"name"`
	Cnames []string `json:"cnames"`
	CnameAccessOnly *bool `json:"cname_access_only,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	EdgeApplicationId int64 `json:"edge_application_id"`
	DigitalCertificateId NullableInt64 `json:"digital_certificate_id,omitempty"`
	Environment *string `json:"environment,omitempty"`
	IsMtlsEnabled *bool `json:"is_mtls_enabled,omitempty"`
	MtlsTrustedCaCertificateId NullableInt64 `json:"mtls_trusted_ca_certificate_id,omitempty"`
	EdgeFirewallId NullableInt64 `json:"edge_firewall_id,omitempty"`
	MtlsVerification *string `json:"mtls_verification,omitempty"`
	CrlList []int64 `json:"crl_list,omitempty"`
}

type _CreateDomainRequest CreateDomainRequest

// NewCreateDomainRequest instantiates a new CreateDomainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDomainRequest(name string, cnames []string, edgeApplicationId int64) *CreateDomainRequest {
	this := CreateDomainRequest{}
	this.Name = name
	this.Cnames = cnames
	this.EdgeApplicationId = edgeApplicationId
	return &this
}

// NewCreateDomainRequestWithDefaults instantiates a new CreateDomainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDomainRequestWithDefaults() *CreateDomainRequest {
	this := CreateDomainRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateDomainRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDomainRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateDomainRequest) SetName(v string) {
	o.Name = v
}

// GetCnames returns the Cnames field value
func (o *CreateDomainRequest) GetCnames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Cnames
}

// GetCnamesOk returns a tuple with the Cnames field value
// and a boolean to check if the value has been set.
func (o *CreateDomainRequest) GetCnamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cnames, true
}

// SetCnames sets field value
func (o *CreateDomainRequest) SetCnames(v []string) {
	o.Cnames = v
}

// GetCnameAccessOnly returns the CnameAccessOnly field value if set, zero value otherwise.
func (o *CreateDomainRequest) GetCnameAccessOnly() bool {
	if o == nil || IsNil(o.CnameAccessOnly) {
		var ret bool
		return ret
	}
	return *o.CnameAccessOnly
}

// GetCnameAccessOnlyOk returns a tuple with the CnameAccessOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomainRequest) GetCnameAccessOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.CnameAccessOnly) {
		return nil, false
	}
	return o.CnameAccessOnly, true
}

// HasCnameAccessOnly returns a boolean if a field has been set.
func (o *CreateDomainRequest) HasCnameAccessOnly() bool {
	if o != nil && !IsNil(o.CnameAccessOnly) {
		return true
	}

	return false
}

// SetCnameAccessOnly gets a reference to the given bool and assigns it to the CnameAccessOnly field.
func (o *CreateDomainRequest) SetCnameAccessOnly(v bool) {
	o.CnameAccessOnly = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *CreateDomainRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomainRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *CreateDomainRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *CreateDomainRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetEdgeApplicationId returns the EdgeApplicationId field value
func (o *CreateDomainRequest) GetEdgeApplicationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EdgeApplicationId
}

// GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field value
// and a boolean to check if the value has been set.
func (o *CreateDomainRequest) GetEdgeApplicationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeApplicationId, true
}

// SetEdgeApplicationId sets field value
func (o *CreateDomainRequest) SetEdgeApplicationId(v int64) {
	o.EdgeApplicationId = v
}

// GetDigitalCertificateId returns the DigitalCertificateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateDomainRequest) GetDigitalCertificateId() int64 {
	if o == nil || IsNil(o.DigitalCertificateId.Get()) {
		var ret int64
		return ret
	}
	return *o.DigitalCertificateId.Get()
}

// GetDigitalCertificateIdOk returns a tuple with the DigitalCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateDomainRequest) GetDigitalCertificateIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DigitalCertificateId.Get(), o.DigitalCertificateId.IsSet()
}

// HasDigitalCertificateId returns a boolean if a field has been set.
func (o *CreateDomainRequest) HasDigitalCertificateId() bool {
	if o != nil && o.DigitalCertificateId.IsSet() {
		return true
	}

	return false
}

// SetDigitalCertificateId gets a reference to the given NullableInt64 and assigns it to the DigitalCertificateId field.
func (o *CreateDomainRequest) SetDigitalCertificateId(v int64) {
	o.DigitalCertificateId.Set(&v)
}
// SetDigitalCertificateIdNil sets the value for DigitalCertificateId to be an explicit nil
func (o *CreateDomainRequest) SetDigitalCertificateIdNil() {
	o.DigitalCertificateId.Set(nil)
}

// UnsetDigitalCertificateId ensures that no value is present for DigitalCertificateId, not even an explicit nil
func (o *CreateDomainRequest) UnsetDigitalCertificateId() {
	o.DigitalCertificateId.Unset()
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *CreateDomainRequest) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomainRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CreateDomainRequest) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *CreateDomainRequest) SetEnvironment(v string) {
	o.Environment = &v
}

// GetIsMtlsEnabled returns the IsMtlsEnabled field value if set, zero value otherwise.
func (o *CreateDomainRequest) GetIsMtlsEnabled() bool {
	if o == nil || IsNil(o.IsMtlsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsMtlsEnabled
}

// GetIsMtlsEnabledOk returns a tuple with the IsMtlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomainRequest) GetIsMtlsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMtlsEnabled) {
		return nil, false
	}
	return o.IsMtlsEnabled, true
}

// HasIsMtlsEnabled returns a boolean if a field has been set.
func (o *CreateDomainRequest) HasIsMtlsEnabled() bool {
	if o != nil && !IsNil(o.IsMtlsEnabled) {
		return true
	}

	return false
}

// SetIsMtlsEnabled gets a reference to the given bool and assigns it to the IsMtlsEnabled field.
func (o *CreateDomainRequest) SetIsMtlsEnabled(v bool) {
	o.IsMtlsEnabled = &v
}

// GetMtlsTrustedCaCertificateId returns the MtlsTrustedCaCertificateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateDomainRequest) GetMtlsTrustedCaCertificateId() int64 {
	if o == nil || IsNil(o.MtlsTrustedCaCertificateId.Get()) {
		var ret int64
		return ret
	}
	return *o.MtlsTrustedCaCertificateId.Get()
}

// GetMtlsTrustedCaCertificateIdOk returns a tuple with the MtlsTrustedCaCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateDomainRequest) GetMtlsTrustedCaCertificateIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MtlsTrustedCaCertificateId.Get(), o.MtlsTrustedCaCertificateId.IsSet()
}

// HasMtlsTrustedCaCertificateId returns a boolean if a field has been set.
func (o *CreateDomainRequest) HasMtlsTrustedCaCertificateId() bool {
	if o != nil && o.MtlsTrustedCaCertificateId.IsSet() {
		return true
	}

	return false
}

// SetMtlsTrustedCaCertificateId gets a reference to the given NullableInt64 and assigns it to the MtlsTrustedCaCertificateId field.
func (o *CreateDomainRequest) SetMtlsTrustedCaCertificateId(v int64) {
	o.MtlsTrustedCaCertificateId.Set(&v)
}
// SetMtlsTrustedCaCertificateIdNil sets the value for MtlsTrustedCaCertificateId to be an explicit nil
func (o *CreateDomainRequest) SetMtlsTrustedCaCertificateIdNil() {
	o.MtlsTrustedCaCertificateId.Set(nil)
}

// UnsetMtlsTrustedCaCertificateId ensures that no value is present for MtlsTrustedCaCertificateId, not even an explicit nil
func (o *CreateDomainRequest) UnsetMtlsTrustedCaCertificateId() {
	o.MtlsTrustedCaCertificateId.Unset()
}

// GetEdgeFirewallId returns the EdgeFirewallId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateDomainRequest) GetEdgeFirewallId() int64 {
	if o == nil || IsNil(o.EdgeFirewallId.Get()) {
		var ret int64
		return ret
	}
	return *o.EdgeFirewallId.Get()
}

// GetEdgeFirewallIdOk returns a tuple with the EdgeFirewallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateDomainRequest) GetEdgeFirewallIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EdgeFirewallId.Get(), o.EdgeFirewallId.IsSet()
}

// HasEdgeFirewallId returns a boolean if a field has been set.
func (o *CreateDomainRequest) HasEdgeFirewallId() bool {
	if o != nil && o.EdgeFirewallId.IsSet() {
		return true
	}

	return false
}

// SetEdgeFirewallId gets a reference to the given NullableInt64 and assigns it to the EdgeFirewallId field.
func (o *CreateDomainRequest) SetEdgeFirewallId(v int64) {
	o.EdgeFirewallId.Set(&v)
}
// SetEdgeFirewallIdNil sets the value for EdgeFirewallId to be an explicit nil
func (o *CreateDomainRequest) SetEdgeFirewallIdNil() {
	o.EdgeFirewallId.Set(nil)
}

// UnsetEdgeFirewallId ensures that no value is present for EdgeFirewallId, not even an explicit nil
func (o *CreateDomainRequest) UnsetEdgeFirewallId() {
	o.EdgeFirewallId.Unset()
}

// GetMtlsVerification returns the MtlsVerification field value if set, zero value otherwise.
func (o *CreateDomainRequest) GetMtlsVerification() string {
	if o == nil || IsNil(o.MtlsVerification) {
		var ret string
		return ret
	}
	return *o.MtlsVerification
}

// GetMtlsVerificationOk returns a tuple with the MtlsVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDomainRequest) GetMtlsVerificationOk() (*string, bool) {
	if o == nil || IsNil(o.MtlsVerification) {
		return nil, false
	}
	return o.MtlsVerification, true
}

// HasMtlsVerification returns a boolean if a field has been set.
func (o *CreateDomainRequest) HasMtlsVerification() bool {
	if o != nil && !IsNil(o.MtlsVerification) {
		return true
	}

	return false
}

// SetMtlsVerification gets a reference to the given string and assigns it to the MtlsVerification field.
func (o *CreateDomainRequest) SetMtlsVerification(v string) {
	o.MtlsVerification = &v
}

// GetCrlList returns the CrlList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateDomainRequest) GetCrlList() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.CrlList
}

// GetCrlListOk returns a tuple with the CrlList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateDomainRequest) GetCrlListOk() ([]int64, bool) {
	if o == nil || IsNil(o.CrlList) {
		return nil, false
	}
	return o.CrlList, true
}

// HasCrlList returns a boolean if a field has been set.
func (o *CreateDomainRequest) HasCrlList() bool {
	if o != nil && !IsNil(o.CrlList) {
		return true
	}

	return false
}

// SetCrlList gets a reference to the given []int64 and assigns it to the CrlList field.
func (o *CreateDomainRequest) SetCrlList(v []int64) {
	o.CrlList = v
}

func (o CreateDomainRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDomainRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["cnames"] = o.Cnames
	if !IsNil(o.CnameAccessOnly) {
		toSerialize["cname_access_only"] = o.CnameAccessOnly
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	toSerialize["edge_application_id"] = o.EdgeApplicationId
	if o.DigitalCertificateId.IsSet() {
		toSerialize["digital_certificate_id"] = o.DigitalCertificateId.Get()
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

func (o *CreateDomainRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"cnames",
		"edge_application_id",
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

	varCreateDomainRequest := _CreateDomainRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateDomainRequest)

	if err != nil {
		return err
	}

	*o = CreateDomainRequest(varCreateDomainRequest)

	return err
}

type NullableCreateDomainRequest struct {
	value *CreateDomainRequest
	isSet bool
}

func (v NullableCreateDomainRequest) Get() *CreateDomainRequest {
	return v.value
}

func (v *NullableCreateDomainRequest) Set(val *CreateDomainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDomainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDomainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDomainRequest(val *CreateDomainRequest) *NullableCreateDomainRequest {
	return &NullableCreateDomainRequest{value: val, isSet: true}
}

func (v NullableCreateDomainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDomainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


