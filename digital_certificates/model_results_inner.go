/*
Digital Certificates API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package digital_certificates

import (
	"encoding/json"
)

// checks if the ResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultsInner{}

// ResultsInner struct for ResultsInner
type ResultsInner struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	SubjectName []string `json:"subject_name,omitempty"`
	Validity NullableString `json:"validity,omitempty"`
	Status *string `json:"status,omitempty"`
	CertificateType *string `json:"certificate_type,omitempty"`
	Managed *bool `json:"managed,omitempty"`
	Issuer NullableString `json:"issuer,omitempty"`
	AzionInformation *string `json:"azion_information,omitempty"`
}

// NewResultsInner instantiates a new ResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultsInner() *ResultsInner {
	this := ResultsInner{}
	return &this
}

// NewResultsInnerWithDefaults instantiates a new ResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultsInnerWithDefaults() *ResultsInner {
	this := ResultsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResultsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResultsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ResultsInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResultsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResultsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResultsInner) SetName(v string) {
	o.Name = &v
}

// GetSubjectName returns the SubjectName field value if set, zero value otherwise.
func (o *ResultsInner) GetSubjectName() []string {
	if o == nil || IsNil(o.SubjectName) {
		var ret []string
		return ret
	}
	return o.SubjectName
}

// GetSubjectNameOk returns a tuple with the SubjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsInner) GetSubjectNameOk() ([]string, bool) {
	if o == nil || IsNil(o.SubjectName) {
		return nil, false
	}
	return o.SubjectName, true
}

// HasSubjectName returns a boolean if a field has been set.
func (o *ResultsInner) HasSubjectName() bool {
	if o != nil && !IsNil(o.SubjectName) {
		return true
	}

	return false
}

// SetSubjectName gets a reference to the given []string and assigns it to the SubjectName field.
func (o *ResultsInner) SetSubjectName(v []string) {
	o.SubjectName = v
}

// GetValidity returns the Validity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultsInner) GetValidity() string {
	if o == nil || IsNil(o.Validity.Get()) {
		var ret string
		return ret
	}
	return *o.Validity.Get()
}

// GetValidityOk returns a tuple with the Validity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultsInner) GetValidityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Validity.Get(), o.Validity.IsSet()
}

// HasValidity returns a boolean if a field has been set.
func (o *ResultsInner) HasValidity() bool {
	if o != nil && o.Validity.IsSet() {
		return true
	}

	return false
}

// SetValidity gets a reference to the given NullableString and assigns it to the Validity field.
func (o *ResultsInner) SetValidity(v string) {
	o.Validity.Set(&v)
}
// SetValidityNil sets the value for Validity to be an explicit nil
func (o *ResultsInner) SetValidityNil() {
	o.Validity.Set(nil)
}

// UnsetValidity ensures that no value is present for Validity, not even an explicit nil
func (o *ResultsInner) UnsetValidity() {
	o.Validity.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResultsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResultsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ResultsInner) SetStatus(v string) {
	o.Status = &v
}

// GetCertificateType returns the CertificateType field value if set, zero value otherwise.
func (o *ResultsInner) GetCertificateType() string {
	if o == nil || IsNil(o.CertificateType) {
		var ret string
		return ret
	}
	return *o.CertificateType
}

// GetCertificateTypeOk returns a tuple with the CertificateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsInner) GetCertificateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateType) {
		return nil, false
	}
	return o.CertificateType, true
}

// HasCertificateType returns a boolean if a field has been set.
func (o *ResultsInner) HasCertificateType() bool {
	if o != nil && !IsNil(o.CertificateType) {
		return true
	}

	return false
}

// SetCertificateType gets a reference to the given string and assigns it to the CertificateType field.
func (o *ResultsInner) SetCertificateType(v string) {
	o.CertificateType = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *ResultsInner) GetManaged() bool {
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsInner) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *ResultsInner) HasManaged() bool {
	if o != nil && !IsNil(o.Managed) {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *ResultsInner) SetManaged(v bool) {
	o.Managed = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResultsInner) GetIssuer() string {
	if o == nil || IsNil(o.Issuer.Get()) {
		var ret string
		return ret
	}
	return *o.Issuer.Get()
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResultsInner) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issuer.Get(), o.Issuer.IsSet()
}

// HasIssuer returns a boolean if a field has been set.
func (o *ResultsInner) HasIssuer() bool {
	if o != nil && o.Issuer.IsSet() {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given NullableString and assigns it to the Issuer field.
func (o *ResultsInner) SetIssuer(v string) {
	o.Issuer.Set(&v)
}
// SetIssuerNil sets the value for Issuer to be an explicit nil
func (o *ResultsInner) SetIssuerNil() {
	o.Issuer.Set(nil)
}

// UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
func (o *ResultsInner) UnsetIssuer() {
	o.Issuer.Unset()
}

// GetAzionInformation returns the AzionInformation field value if set, zero value otherwise.
func (o *ResultsInner) GetAzionInformation() string {
	if o == nil || IsNil(o.AzionInformation) {
		var ret string
		return ret
	}
	return *o.AzionInformation
}

// GetAzionInformationOk returns a tuple with the AzionInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsInner) GetAzionInformationOk() (*string, bool) {
	if o == nil || IsNil(o.AzionInformation) {
		return nil, false
	}
	return o.AzionInformation, true
}

// HasAzionInformation returns a boolean if a field has been set.
func (o *ResultsInner) HasAzionInformation() bool {
	if o != nil && !IsNil(o.AzionInformation) {
		return true
	}

	return false
}

// SetAzionInformation gets a reference to the given string and assigns it to the AzionInformation field.
func (o *ResultsInner) SetAzionInformation(v string) {
	o.AzionInformation = &v
}

func (o ResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SubjectName) {
		toSerialize["subject_name"] = o.SubjectName
	}
	if o.Validity.IsSet() {
		toSerialize["validity"] = o.Validity.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CertificateType) {
		toSerialize["certificate_type"] = o.CertificateType
	}
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}
	if o.Issuer.IsSet() {
		toSerialize["issuer"] = o.Issuer.Get()
	}
	if !IsNil(o.AzionInformation) {
		toSerialize["azion_information"] = o.AzionInformation
	}
	return toSerialize, nil
}

type NullableResultsInner struct {
	value *ResultsInner
	isSet bool
}

func (v NullableResultsInner) Get() *ResultsInner {
	return v.value
}

func (v *NullableResultsInner) Set(val *ResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultsInner(val *ResultsInner) *NullableResultsInner {
	return &NullableResultsInner{value: val, isSet: true}
}

func (v NullableResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


