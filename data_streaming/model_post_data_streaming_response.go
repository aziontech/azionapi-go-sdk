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

// checks if the PostDataStreamingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostDataStreamingResponse{}

// PostDataStreamingResponse struct for PostDataStreamingResponse
type PostDataStreamingResponse struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Options:  * `2` - Edge Applications Event Collector  * `4` - WAF Event Collector  * `86` - Edge Functions Event Collector  * `184` - Edge Applications + WAF Event Collector  * `251` - Activity History Collector 
	TemplateId *int32 `json:"template_id,omitempty"`
	// Options:  * `http` - Edge Applications  * `waf` - WAF Events  * `cells_console` - Edge Functions  * `rtm_activity` - Activity History 
	DataSource *string `json:"data_source,omitempty"`
	Active *bool `json:"active,omitempty"`
	Endpoint []PostDataStreamingResponseEndpointInner `json:"endpoint,omitempty"`
	// Note:  * Field not used with the rtm_activity data source. 
	AllDomains NullableBool `json:"all_domains,omitempty"`
}

// NewPostDataStreamingResponse instantiates a new PostDataStreamingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostDataStreamingResponse() *PostDataStreamingResponse {
	this := PostDataStreamingResponse{}
	return &this
}

// NewPostDataStreamingResponseWithDefaults instantiates a new PostDataStreamingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostDataStreamingResponseWithDefaults() *PostDataStreamingResponse {
	this := PostDataStreamingResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PostDataStreamingResponse) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDataStreamingResponse) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PostDataStreamingResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PostDataStreamingResponse) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostDataStreamingResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDataStreamingResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostDataStreamingResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostDataStreamingResponse) SetName(v string) {
	o.Name = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *PostDataStreamingResponse) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDataStreamingResponse) GetTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *PostDataStreamingResponse) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *PostDataStreamingResponse) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *PostDataStreamingResponse) GetDataSource() string {
	if o == nil || IsNil(o.DataSource) {
		var ret string
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDataStreamingResponse) GetDataSourceOk() (*string, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *PostDataStreamingResponse) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given string and assigns it to the DataSource field.
func (o *PostDataStreamingResponse) SetDataSource(v string) {
	o.DataSource = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PostDataStreamingResponse) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDataStreamingResponse) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PostDataStreamingResponse) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PostDataStreamingResponse) SetActive(v bool) {
	o.Active = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *PostDataStreamingResponse) GetEndpoint() []PostDataStreamingResponseEndpointInner {
	if o == nil || IsNil(o.Endpoint) {
		var ret []PostDataStreamingResponseEndpointInner
		return ret
	}
	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDataStreamingResponse) GetEndpointOk() ([]PostDataStreamingResponseEndpointInner, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *PostDataStreamingResponse) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given []PostDataStreamingResponseEndpointInner and assigns it to the Endpoint field.
func (o *PostDataStreamingResponse) SetEndpoint(v []PostDataStreamingResponseEndpointInner) {
	o.Endpoint = v
}

// GetAllDomains returns the AllDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostDataStreamingResponse) GetAllDomains() bool {
	if o == nil || IsNil(o.AllDomains.Get()) {
		var ret bool
		return ret
	}
	return *o.AllDomains.Get()
}

// GetAllDomainsOk returns a tuple with the AllDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostDataStreamingResponse) GetAllDomainsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllDomains.Get(), o.AllDomains.IsSet()
}

// HasAllDomains returns a boolean if a field has been set.
func (o *PostDataStreamingResponse) HasAllDomains() bool {
	if o != nil && o.AllDomains.IsSet() {
		return true
	}

	return false
}

// SetAllDomains gets a reference to the given NullableBool and assigns it to the AllDomains field.
func (o *PostDataStreamingResponse) SetAllDomains(v bool) {
	o.AllDomains.Set(&v)
}
// SetAllDomainsNil sets the value for AllDomains to be an explicit nil
func (o *PostDataStreamingResponse) SetAllDomainsNil() {
	o.AllDomains.Set(nil)
}

// UnsetAllDomains ensures that no value is present for AllDomains, not even an explicit nil
func (o *PostDataStreamingResponse) UnsetAllDomains() {
	o.AllDomains.Unset()
}

func (o PostDataStreamingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostDataStreamingResponse) ToMap() (map[string]interface{}, error) {
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
	if o.AllDomains.IsSet() {
		toSerialize["all_domains"] = o.AllDomains.Get()
	}
	return toSerialize, nil
}

type NullablePostDataStreamingResponse struct {
	value *PostDataStreamingResponse
	isSet bool
}

func (v NullablePostDataStreamingResponse) Get() *PostDataStreamingResponse {
	return v.value
}

func (v *NullablePostDataStreamingResponse) Set(val *PostDataStreamingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostDataStreamingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostDataStreamingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostDataStreamingResponse(val *PostDataStreamingResponse) *NullablePostDataStreamingResponse {
	return &NullablePostDataStreamingResponse{value: val, isSet: true}
}

func (v NullablePostDataStreamingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostDataStreamingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


