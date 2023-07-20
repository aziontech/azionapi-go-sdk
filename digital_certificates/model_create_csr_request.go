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

// checks if the CreateCSRRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCSRRequest{}

// CreateCSRRequest struct for CreateCSRRequest
type CreateCSRRequest struct {
	Name *string `json:"name,omitempty"`
	CommonName *string `json:"common_name,omitempty"`
	Country *string `json:"country,omitempty"`
	State *string `json:"state,omitempty"`
	Locality *string `json:"locality,omitempty"`
	Organization *string `json:"organization,omitempty"`
	OrganizationUnity *string `json:"organization_unity,omitempty"`
	Email *string `json:"email,omitempty"`
	PrivateKeyType *string `json:"private_key_type,omitempty"`
	Sans []string `json:"sans,omitempty"`
}

// NewCreateCSRRequest instantiates a new CreateCSRRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCSRRequest() *CreateCSRRequest {
	this := CreateCSRRequest{}
	return &this
}

// NewCreateCSRRequestWithDefaults instantiates a new CreateCSRRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCSRRequestWithDefaults() *CreateCSRRequest {
	this := CreateCSRRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateCSRRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCSRRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateCSRRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateCSRRequest) SetName(v string) {
	o.Name = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *CreateCSRRequest) GetCommonName() string {
	if o == nil || IsNil(o.CommonName) {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCSRRequest) GetCommonNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommonName) {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *CreateCSRRequest) HasCommonName() bool {
	if o != nil && !IsNil(o.CommonName) {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *CreateCSRRequest) SetCommonName(v string) {
	o.CommonName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CreateCSRRequest) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCSRRequest) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CreateCSRRequest) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CreateCSRRequest) SetCountry(v string) {
	o.Country = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CreateCSRRequest) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCSRRequest) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CreateCSRRequest) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CreateCSRRequest) SetState(v string) {
	o.State = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *CreateCSRRequest) GetLocality() string {
	if o == nil || IsNil(o.Locality) {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCSRRequest) GetLocalityOk() (*string, bool) {
	if o == nil || IsNil(o.Locality) {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *CreateCSRRequest) HasLocality() bool {
	if o != nil && !IsNil(o.Locality) {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *CreateCSRRequest) SetLocality(v string) {
	o.Locality = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CreateCSRRequest) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCSRRequest) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CreateCSRRequest) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *CreateCSRRequest) SetOrganization(v string) {
	o.Organization = &v
}

// GetOrganizationUnity returns the OrganizationUnity field value if set, zero value otherwise.
func (o *CreateCSRRequest) GetOrganizationUnity() string {
	if o == nil || IsNil(o.OrganizationUnity) {
		var ret string
		return ret
	}
	return *o.OrganizationUnity
}

// GetOrganizationUnityOk returns a tuple with the OrganizationUnity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCSRRequest) GetOrganizationUnityOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationUnity) {
		return nil, false
	}
	return o.OrganizationUnity, true
}

// HasOrganizationUnity returns a boolean if a field has been set.
func (o *CreateCSRRequest) HasOrganizationUnity() bool {
	if o != nil && !IsNil(o.OrganizationUnity) {
		return true
	}

	return false
}

// SetOrganizationUnity gets a reference to the given string and assigns it to the OrganizationUnity field.
func (o *CreateCSRRequest) SetOrganizationUnity(v string) {
	o.OrganizationUnity = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CreateCSRRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCSRRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateCSRRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CreateCSRRequest) SetEmail(v string) {
	o.Email = &v
}

// GetPrivateKeyType returns the PrivateKeyType field value if set, zero value otherwise.
func (o *CreateCSRRequest) GetPrivateKeyType() string {
	if o == nil || IsNil(o.PrivateKeyType) {
		var ret string
		return ret
	}
	return *o.PrivateKeyType
}

// GetPrivateKeyTypeOk returns a tuple with the PrivateKeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCSRRequest) GetPrivateKeyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKeyType) {
		return nil, false
	}
	return o.PrivateKeyType, true
}

// HasPrivateKeyType returns a boolean if a field has been set.
func (o *CreateCSRRequest) HasPrivateKeyType() bool {
	if o != nil && !IsNil(o.PrivateKeyType) {
		return true
	}

	return false
}

// SetPrivateKeyType gets a reference to the given string and assigns it to the PrivateKeyType field.
func (o *CreateCSRRequest) SetPrivateKeyType(v string) {
	o.PrivateKeyType = &v
}

// GetSans returns the Sans field value if set, zero value otherwise.
func (o *CreateCSRRequest) GetSans() []string {
	if o == nil || IsNil(o.Sans) {
		var ret []string
		return ret
	}
	return o.Sans
}

// GetSansOk returns a tuple with the Sans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCSRRequest) GetSansOk() ([]string, bool) {
	if o == nil || IsNil(o.Sans) {
		return nil, false
	}
	return o.Sans, true
}

// HasSans returns a boolean if a field has been set.
func (o *CreateCSRRequest) HasSans() bool {
	if o != nil && !IsNil(o.Sans) {
		return true
	}

	return false
}

// SetSans gets a reference to the given []string and assigns it to the Sans field.
func (o *CreateCSRRequest) SetSans(v []string) {
	o.Sans = v
}

func (o CreateCSRRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCSRRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CommonName) {
		toSerialize["common_name"] = o.CommonName
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Locality) {
		toSerialize["locality"] = o.Locality
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.OrganizationUnity) {
		toSerialize["organization_unity"] = o.OrganizationUnity
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.PrivateKeyType) {
		toSerialize["private_key_type"] = o.PrivateKeyType
	}
	if !IsNil(o.Sans) {
		toSerialize["sans"] = o.Sans
	}
	return toSerialize, nil
}

type NullableCreateCSRRequest struct {
	value *CreateCSRRequest
	isSet bool
}

func (v NullableCreateCSRRequest) Get() *CreateCSRRequest {
	return v.value
}

func (v *NullableCreateCSRRequest) Set(val *CreateCSRRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCSRRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCSRRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCSRRequest(val *CreateCSRRequest) *NullableCreateCSRRequest {
	return &NullableCreateCSRRequest{value: val, isSet: true}
}

func (v NullableCreateCSRRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCSRRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

