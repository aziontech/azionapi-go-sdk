/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
	"fmt"
)

// checks if the CreateOriginsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOriginsRequest{}

// CreateOriginsRequest struct for CreateOriginsRequest
type CreateOriginsRequest struct {
	Name string `json:"name"`
	OriginType *string `json:"origin_type,omitempty"`
	Addresses []CreateOriginsRequestAddresses `json:"addresses"`
	OriginProtocolPolicy *string `json:"origin_protocol_policy,omitempty"`
	HostHeader string `json:"host_header"`
	OriginPath *string `json:"origin_path,omitempty"`
	HmacAuthentication *bool `json:"hmac_authentication,omitempty"`
	HmacRegionName *string `json:"hmac_region_name,omitempty"`
	HmacAccessKey *string `json:"hmac_access_key,omitempty"`
	HmacSecretKey *string `json:"hmac_secret_key,omitempty"`
	Bucket *string `json:"bucket,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
}

type _CreateOriginsRequest CreateOriginsRequest

// NewCreateOriginsRequest instantiates a new CreateOriginsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOriginsRequest(name string, addresses []CreateOriginsRequestAddresses, hostHeader string) *CreateOriginsRequest {
	this := CreateOriginsRequest{}
	this.Name = name
	this.Addresses = addresses
	this.HostHeader = hostHeader
	return &this
}

// NewCreateOriginsRequestWithDefaults instantiates a new CreateOriginsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOriginsRequestWithDefaults() *CreateOriginsRequest {
	this := CreateOriginsRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOriginsRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOriginsRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOriginsRequest) SetName(v string) {
	o.Name = v
}

// GetOriginType returns the OriginType field value if set, zero value otherwise.
func (o *CreateOriginsRequest) GetOriginType() string {
	if o == nil || IsNil(o.OriginType) {
		var ret string
		return ret
	}
	return *o.OriginType
}

// GetOriginTypeOk returns a tuple with the OriginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOriginsRequest) GetOriginTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OriginType) {
		return nil, false
	}
	return o.OriginType, true
}

// HasOriginType returns a boolean if a field has been set.
func (o *CreateOriginsRequest) HasOriginType() bool {
	if o != nil && !IsNil(o.OriginType) {
		return true
	}

	return false
}

// SetOriginType gets a reference to the given string and assigns it to the OriginType field.
func (o *CreateOriginsRequest) SetOriginType(v string) {
	o.OriginType = &v
}

// GetAddresses returns the Addresses field value
func (o *CreateOriginsRequest) GetAddresses() []CreateOriginsRequestAddresses {
	if o == nil {
		var ret []CreateOriginsRequestAddresses
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *CreateOriginsRequest) GetAddressesOk() ([]CreateOriginsRequestAddresses, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *CreateOriginsRequest) SetAddresses(v []CreateOriginsRequestAddresses) {
	o.Addresses = v
}

// GetOriginProtocolPolicy returns the OriginProtocolPolicy field value if set, zero value otherwise.
func (o *CreateOriginsRequest) GetOriginProtocolPolicy() string {
	if o == nil || IsNil(o.OriginProtocolPolicy) {
		var ret string
		return ret
	}
	return *o.OriginProtocolPolicy
}

// GetOriginProtocolPolicyOk returns a tuple with the OriginProtocolPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOriginsRequest) GetOriginProtocolPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.OriginProtocolPolicy) {
		return nil, false
	}
	return o.OriginProtocolPolicy, true
}

// HasOriginProtocolPolicy returns a boolean if a field has been set.
func (o *CreateOriginsRequest) HasOriginProtocolPolicy() bool {
	if o != nil && !IsNil(o.OriginProtocolPolicy) {
		return true
	}

	return false
}

// SetOriginProtocolPolicy gets a reference to the given string and assigns it to the OriginProtocolPolicy field.
func (o *CreateOriginsRequest) SetOriginProtocolPolicy(v string) {
	o.OriginProtocolPolicy = &v
}

// GetHostHeader returns the HostHeader field value
func (o *CreateOriginsRequest) GetHostHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostHeader
}

// GetHostHeaderOk returns a tuple with the HostHeader field value
// and a boolean to check if the value has been set.
func (o *CreateOriginsRequest) GetHostHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostHeader, true
}

// SetHostHeader sets field value
func (o *CreateOriginsRequest) SetHostHeader(v string) {
	o.HostHeader = v
}

// GetOriginPath returns the OriginPath field value if set, zero value otherwise.
func (o *CreateOriginsRequest) GetOriginPath() string {
	if o == nil || IsNil(o.OriginPath) {
		var ret string
		return ret
	}
	return *o.OriginPath
}

// GetOriginPathOk returns a tuple with the OriginPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOriginsRequest) GetOriginPathOk() (*string, bool) {
	if o == nil || IsNil(o.OriginPath) {
		return nil, false
	}
	return o.OriginPath, true
}

// HasOriginPath returns a boolean if a field has been set.
func (o *CreateOriginsRequest) HasOriginPath() bool {
	if o != nil && !IsNil(o.OriginPath) {
		return true
	}

	return false
}

// SetOriginPath gets a reference to the given string and assigns it to the OriginPath field.
func (o *CreateOriginsRequest) SetOriginPath(v string) {
	o.OriginPath = &v
}

// GetHmacAuthentication returns the HmacAuthentication field value if set, zero value otherwise.
func (o *CreateOriginsRequest) GetHmacAuthentication() bool {
	if o == nil || IsNil(o.HmacAuthentication) {
		var ret bool
		return ret
	}
	return *o.HmacAuthentication
}

// GetHmacAuthenticationOk returns a tuple with the HmacAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOriginsRequest) GetHmacAuthenticationOk() (*bool, bool) {
	if o == nil || IsNil(o.HmacAuthentication) {
		return nil, false
	}
	return o.HmacAuthentication, true
}

// HasHmacAuthentication returns a boolean if a field has been set.
func (o *CreateOriginsRequest) HasHmacAuthentication() bool {
	if o != nil && !IsNil(o.HmacAuthentication) {
		return true
	}

	return false
}

// SetHmacAuthentication gets a reference to the given bool and assigns it to the HmacAuthentication field.
func (o *CreateOriginsRequest) SetHmacAuthentication(v bool) {
	o.HmacAuthentication = &v
}

// GetHmacRegionName returns the HmacRegionName field value if set, zero value otherwise.
func (o *CreateOriginsRequest) GetHmacRegionName() string {
	if o == nil || IsNil(o.HmacRegionName) {
		var ret string
		return ret
	}
	return *o.HmacRegionName
}

// GetHmacRegionNameOk returns a tuple with the HmacRegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOriginsRequest) GetHmacRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.HmacRegionName) {
		return nil, false
	}
	return o.HmacRegionName, true
}

// HasHmacRegionName returns a boolean if a field has been set.
func (o *CreateOriginsRequest) HasHmacRegionName() bool {
	if o != nil && !IsNil(o.HmacRegionName) {
		return true
	}

	return false
}

// SetHmacRegionName gets a reference to the given string and assigns it to the HmacRegionName field.
func (o *CreateOriginsRequest) SetHmacRegionName(v string) {
	o.HmacRegionName = &v
}

// GetHmacAccessKey returns the HmacAccessKey field value if set, zero value otherwise.
func (o *CreateOriginsRequest) GetHmacAccessKey() string {
	if o == nil || IsNil(o.HmacAccessKey) {
		var ret string
		return ret
	}
	return *o.HmacAccessKey
}

// GetHmacAccessKeyOk returns a tuple with the HmacAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOriginsRequest) GetHmacAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.HmacAccessKey) {
		return nil, false
	}
	return o.HmacAccessKey, true
}

// HasHmacAccessKey returns a boolean if a field has been set.
func (o *CreateOriginsRequest) HasHmacAccessKey() bool {
	if o != nil && !IsNil(o.HmacAccessKey) {
		return true
	}

	return false
}

// SetHmacAccessKey gets a reference to the given string and assigns it to the HmacAccessKey field.
func (o *CreateOriginsRequest) SetHmacAccessKey(v string) {
	o.HmacAccessKey = &v
}

// GetHmacSecretKey returns the HmacSecretKey field value if set, zero value otherwise.
func (o *CreateOriginsRequest) GetHmacSecretKey() string {
	if o == nil || IsNil(o.HmacSecretKey) {
		var ret string
		return ret
	}
	return *o.HmacSecretKey
}

// GetHmacSecretKeyOk returns a tuple with the HmacSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOriginsRequest) GetHmacSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.HmacSecretKey) {
		return nil, false
	}
	return o.HmacSecretKey, true
}

// HasHmacSecretKey returns a boolean if a field has been set.
func (o *CreateOriginsRequest) HasHmacSecretKey() bool {
	if o != nil && !IsNil(o.HmacSecretKey) {
		return true
	}

	return false
}

// SetHmacSecretKey gets a reference to the given string and assigns it to the HmacSecretKey field.
func (o *CreateOriginsRequest) SetHmacSecretKey(v string) {
	o.HmacSecretKey = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *CreateOriginsRequest) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOriginsRequest) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *CreateOriginsRequest) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *CreateOriginsRequest) SetBucket(v string) {
	o.Bucket = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *CreateOriginsRequest) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOriginsRequest) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *CreateOriginsRequest) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *CreateOriginsRequest) SetPrefix(v string) {
	o.Prefix = &v
}

func (o CreateOriginsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOriginsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.OriginType) {
		toSerialize["origin_type"] = o.OriginType
	}
	toSerialize["addresses"] = o.Addresses
	if !IsNil(o.OriginProtocolPolicy) {
		toSerialize["origin_protocol_policy"] = o.OriginProtocolPolicy
	}
	toSerialize["host_header"] = o.HostHeader
	if !IsNil(o.OriginPath) {
		toSerialize["origin_path"] = o.OriginPath
	}
	if !IsNil(o.HmacAuthentication) {
		toSerialize["hmac_authentication"] = o.HmacAuthentication
	}
	if !IsNil(o.HmacRegionName) {
		toSerialize["hmac_region_name"] = o.HmacRegionName
	}
	if !IsNil(o.HmacAccessKey) {
		toSerialize["hmac_access_key"] = o.HmacAccessKey
	}
	if !IsNil(o.HmacSecretKey) {
		toSerialize["hmac_secret_key"] = o.HmacSecretKey
	}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	return toSerialize, nil
}

func (o *CreateOriginsRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"addresses",
		"host_header",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateOriginsRequest := _CreateOriginsRequest{}

	err = json.Unmarshal(bytes, &varCreateOriginsRequest)

	if err != nil {
		return err
	}

	*o = CreateOriginsRequest(varCreateOriginsRequest)

	return err
}

type NullableCreateOriginsRequest struct {
	value *CreateOriginsRequest
	isSet bool
}

func (v NullableCreateOriginsRequest) Get() *CreateOriginsRequest {
	return v.value
}

func (v *NullableCreateOriginsRequest) Set(val *CreateOriginsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOriginsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOriginsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOriginsRequest(val *CreateOriginsRequest) *NullableCreateOriginsRequest {
	return &NullableCreateOriginsRequest{value: val, isSet: true}
}

func (v NullableCreateOriginsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOriginsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


