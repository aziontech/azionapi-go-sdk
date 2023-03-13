/*
Intelligent DNS

Azion Intelligent DNS API

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package idns

import (
	"encoding/json"
)

// checks if the GetRecordsResponseResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecordsResponseResults{}

// GetRecordsResponseResults struct for GetRecordsResponseResults
type GetRecordsResponseResults struct {
	ZoneId     *int32  `json:"zone_id,omitempty"`
	ZoneDomain *string `json:"zone_domain,omitempty"`
	// Zone records collection
	Records []Record `json:"records,omitempty"`
}

// NewGetRecordsResponseResults instantiates a new GetRecordsResponseResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecordsResponseResults() *GetRecordsResponseResults {
	this := GetRecordsResponseResults{}
	return &this
}

// NewGetRecordsResponseResultsWithDefaults instantiates a new GetRecordsResponseResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecordsResponseResultsWithDefaults() *GetRecordsResponseResults {
	this := GetRecordsResponseResults{}
	return &this
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *GetRecordsResponseResults) GetZoneId() int32 {
	if o == nil || IsNil(o.ZoneId) {
		var ret int32
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsResponseResults) GetZoneIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ZoneId) {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *GetRecordsResponseResults) HasZoneId() bool {
	if o != nil && !IsNil(o.ZoneId) {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given int32 and assigns it to the ZoneId field.
func (o *GetRecordsResponseResults) SetZoneId(v int32) {
	o.ZoneId = &v
}

// GetZoneDomain returns the ZoneDomain field value if set, zero value otherwise.
func (o *GetRecordsResponseResults) GetZoneDomain() string {
	if o == nil || IsNil(o.ZoneDomain) {
		var ret string
		return ret
	}
	return *o.ZoneDomain
}

// GetZoneDomainOk returns a tuple with the ZoneDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsResponseResults) GetZoneDomainOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneDomain) {
		return nil, false
	}
	return o.ZoneDomain, true
}

// HasZoneDomain returns a boolean if a field has been set.
func (o *GetRecordsResponseResults) HasZoneDomain() bool {
	if o != nil && !IsNil(o.ZoneDomain) {
		return true
	}

	return false
}

// SetZoneDomain gets a reference to the given string and assigns it to the ZoneDomain field.
func (o *GetRecordsResponseResults) SetZoneDomain(v string) {
	o.ZoneDomain = &v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *GetRecordsResponseResults) GetRecords() []Record {
	if o == nil || IsNil(o.Records) {
		var ret []Record
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordsResponseResults) GetRecordsOk() ([]Record, bool) {
	if o == nil || IsNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *GetRecordsResponseResults) HasRecords() bool {
	if o != nil && !IsNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []Record and assigns it to the Records field.
func (o *GetRecordsResponseResults) SetRecords(v []Record) {
	o.Records = v
}

func (o GetRecordsResponseResults) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecordsResponseResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ZoneId) {
		toSerialize["zone_id"] = o.ZoneId
	}
	if !IsNil(o.ZoneDomain) {
		toSerialize["zone_domain"] = o.ZoneDomain
	}
	if !IsNil(o.Records) {
		toSerialize["records"] = o.Records
	}
	return toSerialize, nil
}

type NullableGetRecordsResponseResults struct {
	value *GetRecordsResponseResults
	isSet bool
}

func (v NullableGetRecordsResponseResults) Get() *GetRecordsResponseResults {
	return v.value
}

func (v *NullableGetRecordsResponseResults) Set(val *GetRecordsResponseResults) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecordsResponseResults) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecordsResponseResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecordsResponseResults(val *GetRecordsResponseResults) *NullableGetRecordsResponseResults {
	return &NullableGetRecordsResponseResults{value: val, isSet: true}
}

func (v NullableGetRecordsResponseResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecordsResponseResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
