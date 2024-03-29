/*
Data Streaming - OpenAPI

The Data Streaming API allows you to manage your existing data streamings and templates. Data Streaming allows you to feed your stream processing, SIEM, and big data platforms with the event logs from your applications on Azion in real time. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data_streaming

import (
	"encoding/json"
)

// checks if the StandardDataStreamingPostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardDataStreamingPostBody{}

// StandardDataStreamingPostBody struct for StandardDataStreamingPostBody
type StandardDataStreamingPostBody struct {
	Name *string `json:"name,omitempty"`
	// Options:  * `2` - Edge Applications Event Collector  * `4` - WAF Event Collector  * `86` - Edge Functions Event Collector  * `184` - Edge Applications + WAF Event Collector  * `251` - Activity History Collector 
	TemplateId *int32 `json:"template_id,omitempty"`
	// Options:  * `http` - Edge Applications (default)  * `waf` - WAF Events  * `cells_console` - Edge Functions  * `rtm_activity` - Activity History 
	DataSource NullableString `json:"data_source,omitempty"`
	Active NullableBool `json:"active,omitempty"`
	Endpoint *DataStreamingEndpointTypeStandard `json:"endpoint,omitempty"`
	// Note:  * Field not used with the rtm_activity data source. 
	DomainsIds []int32 `json:"domains_ids,omitempty"`
	// Note:  * Field not used with the rtm_activity data source. 
	AllDomains NullableBool `json:"all_domains,omitempty"`
}

// NewStandardDataStreamingPostBody instantiates a new StandardDataStreamingPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardDataStreamingPostBody() *StandardDataStreamingPostBody {
	this := StandardDataStreamingPostBody{}
	var active bool = true
	this.Active = *NewNullableBool(&active)
	var allDomains bool = false
	this.AllDomains = *NewNullableBool(&allDomains)
	return &this
}

// NewStandardDataStreamingPostBodyWithDefaults instantiates a new StandardDataStreamingPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardDataStreamingPostBodyWithDefaults() *StandardDataStreamingPostBody {
	this := StandardDataStreamingPostBody{}
	var active bool = true
	this.Active = *NewNullableBool(&active)
	var allDomains bool = false
	this.AllDomains = *NewNullableBool(&allDomains)
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StandardDataStreamingPostBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardDataStreamingPostBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StandardDataStreamingPostBody) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StandardDataStreamingPostBody) SetName(v string) {
	o.Name = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *StandardDataStreamingPostBody) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardDataStreamingPostBody) GetTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *StandardDataStreamingPostBody) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *StandardDataStreamingPostBody) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandardDataStreamingPostBody) GetDataSource() string {
	if o == nil || IsNil(o.DataSource.Get()) {
		var ret string
		return ret
	}
	return *o.DataSource.Get()
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandardDataStreamingPostBody) GetDataSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSource.Get(), o.DataSource.IsSet()
}

// HasDataSource returns a boolean if a field has been set.
func (o *StandardDataStreamingPostBody) HasDataSource() bool {
	if o != nil && o.DataSource.IsSet() {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given NullableString and assigns it to the DataSource field.
func (o *StandardDataStreamingPostBody) SetDataSource(v string) {
	o.DataSource.Set(&v)
}
// SetDataSourceNil sets the value for DataSource to be an explicit nil
func (o *StandardDataStreamingPostBody) SetDataSourceNil() {
	o.DataSource.Set(nil)
}

// UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
func (o *StandardDataStreamingPostBody) UnsetDataSource() {
	o.DataSource.Unset()
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandardDataStreamingPostBody) GetActive() bool {
	if o == nil || IsNil(o.Active.Get()) {
		var ret bool
		return ret
	}
	return *o.Active.Get()
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandardDataStreamingPostBody) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active.Get(), o.Active.IsSet()
}

// HasActive returns a boolean if a field has been set.
func (o *StandardDataStreamingPostBody) HasActive() bool {
	if o != nil && o.Active.IsSet() {
		return true
	}

	return false
}

// SetActive gets a reference to the given NullableBool and assigns it to the Active field.
func (o *StandardDataStreamingPostBody) SetActive(v bool) {
	o.Active.Set(&v)
}
// SetActiveNil sets the value for Active to be an explicit nil
func (o *StandardDataStreamingPostBody) SetActiveNil() {
	o.Active.Set(nil)
}

// UnsetActive ensures that no value is present for Active, not even an explicit nil
func (o *StandardDataStreamingPostBody) UnsetActive() {
	o.Active.Unset()
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *StandardDataStreamingPostBody) GetEndpoint() DataStreamingEndpointTypeStandard {
	if o == nil || IsNil(o.Endpoint) {
		var ret DataStreamingEndpointTypeStandard
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardDataStreamingPostBody) GetEndpointOk() (*DataStreamingEndpointTypeStandard, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *StandardDataStreamingPostBody) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given DataStreamingEndpointTypeStandard and assigns it to the Endpoint field.
func (o *StandardDataStreamingPostBody) SetEndpoint(v DataStreamingEndpointTypeStandard) {
	o.Endpoint = &v
}

// GetDomainsIds returns the DomainsIds field value if set, zero value otherwise.
func (o *StandardDataStreamingPostBody) GetDomainsIds() []int32 {
	if o == nil || IsNil(o.DomainsIds) {
		var ret []int32
		return ret
	}
	return o.DomainsIds
}

// GetDomainsIdsOk returns a tuple with the DomainsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardDataStreamingPostBody) GetDomainsIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.DomainsIds) {
		return nil, false
	}
	return o.DomainsIds, true
}

// HasDomainsIds returns a boolean if a field has been set.
func (o *StandardDataStreamingPostBody) HasDomainsIds() bool {
	if o != nil && !IsNil(o.DomainsIds) {
		return true
	}

	return false
}

// SetDomainsIds gets a reference to the given []int32 and assigns it to the DomainsIds field.
func (o *StandardDataStreamingPostBody) SetDomainsIds(v []int32) {
	o.DomainsIds = v
}

// GetAllDomains returns the AllDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandardDataStreamingPostBody) GetAllDomains() bool {
	if o == nil || IsNil(o.AllDomains.Get()) {
		var ret bool
		return ret
	}
	return *o.AllDomains.Get()
}

// GetAllDomainsOk returns a tuple with the AllDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandardDataStreamingPostBody) GetAllDomainsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllDomains.Get(), o.AllDomains.IsSet()
}

// HasAllDomains returns a boolean if a field has been set.
func (o *StandardDataStreamingPostBody) HasAllDomains() bool {
	if o != nil && o.AllDomains.IsSet() {
		return true
	}

	return false
}

// SetAllDomains gets a reference to the given NullableBool and assigns it to the AllDomains field.
func (o *StandardDataStreamingPostBody) SetAllDomains(v bool) {
	o.AllDomains.Set(&v)
}
// SetAllDomainsNil sets the value for AllDomains to be an explicit nil
func (o *StandardDataStreamingPostBody) SetAllDomainsNil() {
	o.AllDomains.Set(nil)
}

// UnsetAllDomains ensures that no value is present for AllDomains, not even an explicit nil
func (o *StandardDataStreamingPostBody) UnsetAllDomains() {
	o.AllDomains.Unset()
}

func (o StandardDataStreamingPostBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandardDataStreamingPostBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if o.DataSource.IsSet() {
		toSerialize["data_source"] = o.DataSource.Get()
	}
	if o.Active.IsSet() {
		toSerialize["active"] = o.Active.Get()
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.DomainsIds) {
		toSerialize["domains_ids"] = o.DomainsIds
	}
	if o.AllDomains.IsSet() {
		toSerialize["all_domains"] = o.AllDomains.Get()
	}
	return toSerialize, nil
}

type NullableStandardDataStreamingPostBody struct {
	value *StandardDataStreamingPostBody
	isSet bool
}

func (v NullableStandardDataStreamingPostBody) Get() *StandardDataStreamingPostBody {
	return v.value
}

func (v *NullableStandardDataStreamingPostBody) Set(val *StandardDataStreamingPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardDataStreamingPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardDataStreamingPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardDataStreamingPostBody(val *StandardDataStreamingPostBody) *NullableStandardDataStreamingPostBody {
	return &NullableStandardDataStreamingPostBody{value: val, isSet: true}
}

func (v NullableStandardDataStreamingPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardDataStreamingPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


