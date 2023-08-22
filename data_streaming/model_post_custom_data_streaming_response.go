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

// checks if the PostCustomDataStreamingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCustomDataStreamingResponse{}

// PostCustomDataStreamingResponse struct for PostCustomDataStreamingResponse
type PostCustomDataStreamingResponse struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Options:  * `http` - Edge Applications  * `waf` - WAF Events  * `cells_console` - Edge Functions  * `rtm_activity` - Activity History   
	DataSource *string `json:"data_source,omitempty"`
	// Note:  * Add all variables and values that will be used to stream according to the data source you choose to use.   
	TemplateModel *string `json:"template_model,omitempty"`
	Active *bool `json:"active,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	// Note:  * Field not used with the rtm_activity data source. 
	AllDomains NullableBool `json:"all_domains,omitempty"`
}

// NewPostCustomDataStreamingResponse instantiates a new PostCustomDataStreamingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCustomDataStreamingResponse() *PostCustomDataStreamingResponse {
	this := PostCustomDataStreamingResponse{}
	return &this
}

// NewPostCustomDataStreamingResponseWithDefaults instantiates a new PostCustomDataStreamingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCustomDataStreamingResponseWithDefaults() *PostCustomDataStreamingResponse {
	this := PostCustomDataStreamingResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PostCustomDataStreamingResponse) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCustomDataStreamingResponse) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PostCustomDataStreamingResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PostCustomDataStreamingResponse) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostCustomDataStreamingResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCustomDataStreamingResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostCustomDataStreamingResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostCustomDataStreamingResponse) SetName(v string) {
	o.Name = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *PostCustomDataStreamingResponse) GetDataSource() string {
	if o == nil || IsNil(o.DataSource) {
		var ret string
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCustomDataStreamingResponse) GetDataSourceOk() (*string, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *PostCustomDataStreamingResponse) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given string and assigns it to the DataSource field.
func (o *PostCustomDataStreamingResponse) SetDataSource(v string) {
	o.DataSource = &v
}

// GetTemplateModel returns the TemplateModel field value if set, zero value otherwise.
func (o *PostCustomDataStreamingResponse) GetTemplateModel() string {
	if o == nil || IsNil(o.TemplateModel) {
		var ret string
		return ret
	}
	return *o.TemplateModel
}

// GetTemplateModelOk returns a tuple with the TemplateModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCustomDataStreamingResponse) GetTemplateModelOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateModel) {
		return nil, false
	}
	return o.TemplateModel, true
}

// HasTemplateModel returns a boolean if a field has been set.
func (o *PostCustomDataStreamingResponse) HasTemplateModel() bool {
	if o != nil && !IsNil(o.TemplateModel) {
		return true
	}

	return false
}

// SetTemplateModel gets a reference to the given string and assigns it to the TemplateModel field.
func (o *PostCustomDataStreamingResponse) SetTemplateModel(v string) {
	o.TemplateModel = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PostCustomDataStreamingResponse) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCustomDataStreamingResponse) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PostCustomDataStreamingResponse) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PostCustomDataStreamingResponse) SetActive(v bool) {
	o.Active = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *PostCustomDataStreamingResponse) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCustomDataStreamingResponse) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *PostCustomDataStreamingResponse) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *PostCustomDataStreamingResponse) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetAllDomains returns the AllDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostCustomDataStreamingResponse) GetAllDomains() bool {
	if o == nil || IsNil(o.AllDomains.Get()) {
		var ret bool
		return ret
	}
	return *o.AllDomains.Get()
}

// GetAllDomainsOk returns a tuple with the AllDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostCustomDataStreamingResponse) GetAllDomainsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllDomains.Get(), o.AllDomains.IsSet()
}

// HasAllDomains returns a boolean if a field has been set.
func (o *PostCustomDataStreamingResponse) HasAllDomains() bool {
	if o != nil && o.AllDomains.IsSet() {
		return true
	}

	return false
}

// SetAllDomains gets a reference to the given NullableBool and assigns it to the AllDomains field.
func (o *PostCustomDataStreamingResponse) SetAllDomains(v bool) {
	o.AllDomains.Set(&v)
}
// SetAllDomainsNil sets the value for AllDomains to be an explicit nil
func (o *PostCustomDataStreamingResponse) SetAllDomainsNil() {
	o.AllDomains.Set(nil)
}

// UnsetAllDomains ensures that no value is present for AllDomains, not even an explicit nil
func (o *PostCustomDataStreamingResponse) UnsetAllDomains() {
	o.AllDomains.Unset()
}

func (o PostCustomDataStreamingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCustomDataStreamingResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.TemplateModel) {
		toSerialize["template_model"] = o.TemplateModel
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

type NullablePostCustomDataStreamingResponse struct {
	value *PostCustomDataStreamingResponse
	isSet bool
}

func (v NullablePostCustomDataStreamingResponse) Get() *PostCustomDataStreamingResponse {
	return v.value
}

func (v *NullablePostCustomDataStreamingResponse) Set(val *PostCustomDataStreamingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCustomDataStreamingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCustomDataStreamingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCustomDataStreamingResponse(val *PostCustomDataStreamingResponse) *NullablePostCustomDataStreamingResponse {
	return &NullablePostCustomDataStreamingResponse{value: val, isSet: true}
}

func (v NullablePostCustomDataStreamingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCustomDataStreamingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

