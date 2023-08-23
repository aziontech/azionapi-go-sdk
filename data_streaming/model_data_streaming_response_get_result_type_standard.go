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

// checks if the DataStreamingResponseGetResultTypeStandard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataStreamingResponseGetResultTypeStandard{}

// DataStreamingResponseGetResultTypeStandard struct for DataStreamingResponseGetResultTypeStandard
type DataStreamingResponseGetResultTypeStandard struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	TemplateId *int32 `json:"template_id,omitempty"`
	DataSource *string `json:"data_source,omitempty"`
	Active *bool `json:"active,omitempty"`
	Endpoint *DataStreamingEndpointTypeStandard `json:"endpoint,omitempty"`
	AllDomains *bool `json:"all_domains,omitempty"`
}

// NewDataStreamingResponseGetResultTypeStandard instantiates a new DataStreamingResponseGetResultTypeStandard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataStreamingResponseGetResultTypeStandard() *DataStreamingResponseGetResultTypeStandard {
	this := DataStreamingResponseGetResultTypeStandard{}
	return &this
}

// NewDataStreamingResponseGetResultTypeStandardWithDefaults instantiates a new DataStreamingResponseGetResultTypeStandard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataStreamingResponseGetResultTypeStandardWithDefaults() *DataStreamingResponseGetResultTypeStandard {
	this := DataStreamingResponseGetResultTypeStandard{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeStandard) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeStandard) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeStandard) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DataStreamingResponseGetResultTypeStandard) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeStandard) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeStandard) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeStandard) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataStreamingResponseGetResultTypeStandard) SetName(v string) {
	o.Name = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeStandard) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeStandard) GetTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeStandard) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *DataStreamingResponseGetResultTypeStandard) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeStandard) GetDataSource() string {
	if o == nil || IsNil(o.DataSource) {
		var ret string
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeStandard) GetDataSourceOk() (*string, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeStandard) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given string and assigns it to the DataSource field.
func (o *DataStreamingResponseGetResultTypeStandard) SetDataSource(v string) {
	o.DataSource = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeStandard) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeStandard) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeStandard) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *DataStreamingResponseGetResultTypeStandard) SetActive(v bool) {
	o.Active = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeStandard) GetEndpoint() DataStreamingEndpointTypeStandard {
	if o == nil || IsNil(o.Endpoint) {
		var ret DataStreamingEndpointTypeStandard
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeStandard) GetEndpointOk() (*DataStreamingEndpointTypeStandard, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeStandard) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given DataStreamingEndpointTypeStandard and assigns it to the Endpoint field.
func (o *DataStreamingResponseGetResultTypeStandard) SetEndpoint(v DataStreamingEndpointTypeStandard) {
	o.Endpoint = &v
}

// GetAllDomains returns the AllDomains field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeStandard) GetAllDomains() bool {
	if o == nil || IsNil(o.AllDomains) {
		var ret bool
		return ret
	}
	return *o.AllDomains
}

// GetAllDomainsOk returns a tuple with the AllDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeStandard) GetAllDomainsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllDomains) {
		return nil, false
	}
	return o.AllDomains, true
}

// HasAllDomains returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeStandard) HasAllDomains() bool {
	if o != nil && !IsNil(o.AllDomains) {
		return true
	}

	return false
}

// SetAllDomains gets a reference to the given bool and assigns it to the AllDomains field.
func (o *DataStreamingResponseGetResultTypeStandard) SetAllDomains(v bool) {
	o.AllDomains = &v
}

func (o DataStreamingResponseGetResultTypeStandard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataStreamingResponseGetResultTypeStandard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.DataSource) {
		toSerialize["data_source"] = o.DataSource
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.AllDomains) {
		toSerialize["all_domains"] = o.AllDomains
	}
	return toSerialize, nil
}

type NullableDataStreamingResponseGetResultTypeStandard struct {
	value *DataStreamingResponseGetResultTypeStandard
	isSet bool
}

func (v NullableDataStreamingResponseGetResultTypeStandard) Get() *DataStreamingResponseGetResultTypeStandard {
	return v.value
}

func (v *NullableDataStreamingResponseGetResultTypeStandard) Set(val *DataStreamingResponseGetResultTypeStandard) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStreamingResponseGetResultTypeStandard) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStreamingResponseGetResultTypeStandard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStreamingResponseGetResultTypeStandard(val *DataStreamingResponseGetResultTypeStandard) *NullableDataStreamingResponseGetResultTypeStandard {
	return &NullableDataStreamingResponseGetResultTypeStandard{value: val, isSet: true}
}

func (v NullableDataStreamingResponseGetResultTypeStandard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStreamingResponseGetResultTypeStandard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


