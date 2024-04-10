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

// checks if the DomainData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainData{}

// DomainData struct for DomainData
type DomainData struct {
	Name *string `json:"name,omitempty"`
	Cnames []string `json:"cnames,omitempty"`
	CnameAccessOnly *bool `json:"cname_access_only,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	EdgeApplicationId *int64 `json:"edge_application_id,omitempty"`
	DigitalCertificateId NullableInt64 `json:"digital_certificate_id,omitempty"`
	Environment *string `json:"environment,omitempty"`
	IsMtlsEnabled *bool `json:"is_mtls_enabled,omitempty"`
	MtlsTrustedCaCertificateId NullableInt64 `json:"mtls_trusted_ca_certificate_id,omitempty"`
	MtlsVerification *string `json:"mtls_verification,omitempty"`
	CrlList []int64 `json:"crl_list,omitempty"`
}

// NewDomainData instantiates a new DomainData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainData() *DomainData {
	this := DomainData{}
	return &this
}

// NewDomainDataWithDefaults instantiates a new DomainData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainDataWithDefaults() *DomainData {
	this := DomainData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DomainData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DomainData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DomainData) SetName(v string) {
	o.Name = &v
}

// GetCnames returns the Cnames field value if set, zero value otherwise.
func (o *DomainData) GetCnames() []string {
	if o == nil || IsNil(o.Cnames) {
		var ret []string
		return ret
	}
	return o.Cnames
}

// GetCnamesOk returns a tuple with the Cnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainData) GetCnamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Cnames) {
		return nil, false
	}
	return o.Cnames, true
}

// HasCnames returns a boolean if a field has been set.
func (o *DomainData) HasCnames() bool {
	if o != nil && !IsNil(o.Cnames) {
		return true
	}

	return false
}

// SetCnames gets a reference to the given []string and assigns it to the Cnames field.
func (o *DomainData) SetCnames(v []string) {
	o.Cnames = v
}

// GetCnameAccessOnly returns the CnameAccessOnly field value if set, zero value otherwise.
func (o *DomainData) GetCnameAccessOnly() bool {
	if o == nil || IsNil(o.CnameAccessOnly) {
		var ret bool
		return ret
	}
	return *o.CnameAccessOnly
}

// GetCnameAccessOnlyOk returns a tuple with the CnameAccessOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainData) GetCnameAccessOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.CnameAccessOnly) {
		return nil, false
	}
	return o.CnameAccessOnly, true
}

// HasCnameAccessOnly returns a boolean if a field has been set.
func (o *DomainData) HasCnameAccessOnly() bool {
	if o != nil && !IsNil(o.CnameAccessOnly) {
		return true
	}

	return false
}

// SetCnameAccessOnly gets a reference to the given bool and assigns it to the CnameAccessOnly field.
func (o *DomainData) SetCnameAccessOnly(v bool) {
	o.CnameAccessOnly = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *DomainData) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainData) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *DomainData) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *DomainData) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetEdgeApplicationId returns the EdgeApplicationId field value if set, zero value otherwise.
func (o *DomainData) GetEdgeApplicationId() int64 {
	if o == nil || IsNil(o.EdgeApplicationId) {
		var ret int64
		return ret
	}
	return *o.EdgeApplicationId
}

// GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainData) GetEdgeApplicationIdOk() (*int64, bool) {
	if o == nil || IsNil(o.EdgeApplicationId) {
		return nil, false
	}
	return o.EdgeApplicationId, true
}

// HasEdgeApplicationId returns a boolean if a field has been set.
func (o *DomainData) HasEdgeApplicationId() bool {
	if o != nil && !IsNil(o.EdgeApplicationId) {
		return true
	}

	return false
}

// SetEdgeApplicationId gets a reference to the given int64 and assigns it to the EdgeApplicationId field.
func (o *DomainData) SetEdgeApplicationId(v int64) {
	o.EdgeApplicationId = &v
}

// GetDigitalCertificateId returns the DigitalCertificateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainData) GetDigitalCertificateId() int64 {
	if o == nil || IsNil(o.DigitalCertificateId.Get()) {
		var ret int64
		return ret
	}
	return *o.DigitalCertificateId.Get()
}

// GetDigitalCertificateIdOk returns a tuple with the DigitalCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainData) GetDigitalCertificateIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DigitalCertificateId.Get(), o.DigitalCertificateId.IsSet()
}

// HasDigitalCertificateId returns a boolean if a field has been set.
func (o *DomainData) HasDigitalCertificateId() bool {
	if o != nil && o.DigitalCertificateId.IsSet() {
		return true
	}

	return false
}

// SetDigitalCertificateId gets a reference to the given NullableInt64 and assigns it to the DigitalCertificateId field.
func (o *DomainData) SetDigitalCertificateId(v int64) {
	o.DigitalCertificateId.Set(&v)
}
// SetDigitalCertificateIdNil sets the value for DigitalCertificateId to be an explicit nil
func (o *DomainData) SetDigitalCertificateIdNil() {
	o.DigitalCertificateId.Set(nil)
}

// UnsetDigitalCertificateId ensures that no value is present for DigitalCertificateId, not even an explicit nil
func (o *DomainData) UnsetDigitalCertificateId() {
	o.DigitalCertificateId.Unset()
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *DomainData) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainData) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *DomainData) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *DomainData) SetEnvironment(v string) {
	o.Environment = &v
}

// GetIsMtlsEnabled returns the IsMtlsEnabled field value if set, zero value otherwise.
func (o *DomainData) GetIsMtlsEnabled() bool {
	if o == nil || IsNil(o.IsMtlsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsMtlsEnabled
}

// GetIsMtlsEnabledOk returns a tuple with the IsMtlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainData) GetIsMtlsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMtlsEnabled) {
		return nil, false
	}
	return o.IsMtlsEnabled, true
}

// HasIsMtlsEnabled returns a boolean if a field has been set.
func (o *DomainData) HasIsMtlsEnabled() bool {
	if o != nil && !IsNil(o.IsMtlsEnabled) {
		return true
	}

	return false
}

// SetIsMtlsEnabled gets a reference to the given bool and assigns it to the IsMtlsEnabled field.
func (o *DomainData) SetIsMtlsEnabled(v bool) {
	o.IsMtlsEnabled = &v
}

// GetMtlsTrustedCaCertificateId returns the MtlsTrustedCaCertificateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainData) GetMtlsTrustedCaCertificateId() int64 {
	if o == nil || IsNil(o.MtlsTrustedCaCertificateId.Get()) {
		var ret int64
		return ret
	}
	return *o.MtlsTrustedCaCertificateId.Get()
}

// GetMtlsTrustedCaCertificateIdOk returns a tuple with the MtlsTrustedCaCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainData) GetMtlsTrustedCaCertificateIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MtlsTrustedCaCertificateId.Get(), o.MtlsTrustedCaCertificateId.IsSet()
}

// HasMtlsTrustedCaCertificateId returns a boolean if a field has been set.
func (o *DomainData) HasMtlsTrustedCaCertificateId() bool {
	if o != nil && o.MtlsTrustedCaCertificateId.IsSet() {
		return true
	}

	return false
}

// SetMtlsTrustedCaCertificateId gets a reference to the given NullableInt64 and assigns it to the MtlsTrustedCaCertificateId field.
func (o *DomainData) SetMtlsTrustedCaCertificateId(v int64) {
	o.MtlsTrustedCaCertificateId.Set(&v)
}
// SetMtlsTrustedCaCertificateIdNil sets the value for MtlsTrustedCaCertificateId to be an explicit nil
func (o *DomainData) SetMtlsTrustedCaCertificateIdNil() {
	o.MtlsTrustedCaCertificateId.Set(nil)
}

// UnsetMtlsTrustedCaCertificateId ensures that no value is present for MtlsTrustedCaCertificateId, not even an explicit nil
func (o *DomainData) UnsetMtlsTrustedCaCertificateId() {
	o.MtlsTrustedCaCertificateId.Unset()
}

// GetMtlsVerification returns the MtlsVerification field value if set, zero value otherwise.
func (o *DomainData) GetMtlsVerification() string {
	if o == nil || IsNil(o.MtlsVerification) {
		var ret string
		return ret
	}
	return *o.MtlsVerification
}

// GetMtlsVerificationOk returns a tuple with the MtlsVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainData) GetMtlsVerificationOk() (*string, bool) {
	if o == nil || IsNil(o.MtlsVerification) {
		return nil, false
	}
	return o.MtlsVerification, true
}

// HasMtlsVerification returns a boolean if a field has been set.
func (o *DomainData) HasMtlsVerification() bool {
	if o != nil && !IsNil(o.MtlsVerification) {
		return true
	}

	return false
}

// SetMtlsVerification gets a reference to the given string and assigns it to the MtlsVerification field.
func (o *DomainData) SetMtlsVerification(v string) {
	o.MtlsVerification = &v
}

// GetCrlList returns the CrlList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainData) GetCrlList() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.CrlList
}

// GetCrlListOk returns a tuple with the CrlList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainData) GetCrlListOk() ([]int64, bool) {
	if o == nil || IsNil(o.CrlList) {
		return nil, false
	}
	return o.CrlList, true
}

// HasCrlList returns a boolean if a field has been set.
func (o *DomainData) HasCrlList() bool {
	if o != nil && !IsNil(o.CrlList) {
		return true
	}

	return false
}

// SetCrlList gets a reference to the given []int64 and assigns it to the CrlList field.
func (o *DomainData) SetCrlList(v []int64) {
	o.CrlList = v
}

func (o DomainData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainData) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.MtlsVerification) {
		toSerialize["mtls_verification"] = o.MtlsVerification
	}
	if o.CrlList != nil {
		toSerialize["crl_list"] = o.CrlList
	}
	return toSerialize, nil
}

type NullableDomainData struct {
	value *DomainData
	isSet bool
}

func (v NullableDomainData) Get() *DomainData {
	return v.value
}

func (v *NullableDomainData) Set(val *DomainData) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainData) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainData(val *DomainData) *NullableDomainData {
	return &NullableDomainData{value: val, isSet: true}
}

func (v NullableDomainData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


