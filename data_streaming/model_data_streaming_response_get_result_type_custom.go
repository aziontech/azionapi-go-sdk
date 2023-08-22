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

// checks if the DataStreamingResponseGetResultTypeCustom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataStreamingResponseGetResultTypeCustom{}

// DataStreamingResponseGetResultTypeCustom struct for DataStreamingResponseGetResultTypeCustom
type DataStreamingResponseGetResultTypeCustom struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	DataSource *string `json:"data_source,omitempty"`
	Active *bool `json:"active,omitempty"`
	Endpoint *DataStreamingEndpointTypeKafka `json:"endpoint,omitempty"`
	AllDomains *bool `json:"all_domains,omitempty"`
	TemplateModel *string `json:"template_model,omitempty"`
}

// NewDataStreamingResponseGetResultTypeCustom instantiates a new DataStreamingResponseGetResultTypeCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataStreamingResponseGetResultTypeCustom() *DataStreamingResponseGetResultTypeCustom {
	this := DataStreamingResponseGetResultTypeCustom{}
	return &this
}

// NewDataStreamingResponseGetResultTypeCustomWithDefaults instantiates a new DataStreamingResponseGetResultTypeCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataStreamingResponseGetResultTypeCustomWithDefaults() *DataStreamingResponseGetResultTypeCustom {
	this := DataStreamingResponseGetResultTypeCustom{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeCustom) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeCustom) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeCustom) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DataStreamingResponseGetResultTypeCustom) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeCustom) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeCustom) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeCustom) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataStreamingResponseGetResultTypeCustom) SetName(v string) {
	o.Name = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeCustom) GetDataSource() string {
	if o == nil || IsNil(o.DataSource) {
		var ret string
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeCustom) GetDataSourceOk() (*string, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeCustom) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given string and assigns it to the DataSource field.
func (o *DataStreamingResponseGetResultTypeCustom) SetDataSource(v string) {
	o.DataSource = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeCustom) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeCustom) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeCustom) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *DataStreamingResponseGetResultTypeCustom) SetActive(v bool) {
	o.Active = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeCustom) GetEndpoint() DataStreamingEndpointTypeKafka {
	if o == nil || IsNil(o.Endpoint) {
		var ret DataStreamingEndpointTypeKafka
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeCustom) GetEndpointOk() (*DataStreamingEndpointTypeKafka, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeCustom) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given DataStreamingEndpointTypeKafka and assigns it to the Endpoint field.
func (o *DataStreamingResponseGetResultTypeCustom) SetEndpoint(v DataStreamingEndpointTypeKafka) {
	o.Endpoint = &v
}

// GetAllDomains returns the AllDomains field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeCustom) GetAllDomains() bool {
	if o == nil || IsNil(o.AllDomains) {
		var ret bool
		return ret
	}
	return *o.AllDomains
}

// GetAllDomainsOk returns a tuple with the AllDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeCustom) GetAllDomainsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllDomains) {
		return nil, false
	}
	return o.AllDomains, true
}

// HasAllDomains returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeCustom) HasAllDomains() bool {
	if o != nil && !IsNil(o.AllDomains) {
		return true
	}

	return false
}

// SetAllDomains gets a reference to the given bool and assigns it to the AllDomains field.
func (o *DataStreamingResponseGetResultTypeCustom) SetAllDomains(v bool) {
	o.AllDomains = &v
}

// GetTemplateModel returns the TemplateModel field value if set, zero value otherwise.
func (o *DataStreamingResponseGetResultTypeCustom) GetTemplateModel() string {
	if o == nil || IsNil(o.TemplateModel) {
		var ret string
		return ret
	}
	return *o.TemplateModel
}

// GetTemplateModelOk returns a tuple with the TemplateModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingResponseGetResultTypeCustom) GetTemplateModelOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateModel) {
		return nil, false
	}
	return o.TemplateModel, true
}

// HasTemplateModel returns a boolean if a field has been set.
func (o *DataStreamingResponseGetResultTypeCustom) HasTemplateModel() bool {
	if o != nil && !IsNil(o.TemplateModel) {
		return true
	}

	return false
}

// SetTemplateModel gets a reference to the given string and assigns it to the TemplateModel field.
func (o *DataStreamingResponseGetResultTypeCustom) SetTemplateModel(v string) {
	o.TemplateModel = &v
}

func (o DataStreamingResponseGetResultTypeCustom) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataStreamingResponseGetResultTypeCustom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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
	if !IsNil(o.TemplateModel) {
		toSerialize["template_model"] = o.TemplateModel
	}
	return toSerialize, nil
}

type NullableDataStreamingResponseGetResultTypeCustom struct {
	value *DataStreamingResponseGetResultTypeCustom
	isSet bool
}

func (v NullableDataStreamingResponseGetResultTypeCustom) Get() *DataStreamingResponseGetResultTypeCustom {
	return v.value
}

func (v *NullableDataStreamingResponseGetResultTypeCustom) Set(val *DataStreamingResponseGetResultTypeCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStreamingResponseGetResultTypeCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStreamingResponseGetResultTypeCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStreamingResponseGetResultTypeCustom(val *DataStreamingResponseGetResultTypeCustom) *NullableDataStreamingResponseGetResultTypeCustom {
	return &NullableDataStreamingResponseGetResultTypeCustom{value: val, isSet: true}
}

func (v NullableDataStreamingResponseGetResultTypeCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStreamingResponseGetResultTypeCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


