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

// checks if the DataStreamingPostBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataStreamingPostBody{}

// DataStreamingPostBody struct for DataStreamingPostBody
type DataStreamingPostBody struct {
	Name *string `json:"name,omitempty"`
	// Options:  * `2` - Edge Applications Event Collector  * `4` - WAF Event Collector  * `86` - Edge Functions Event Collector  * `184` - Edge Applications + WAF Event Collector  * `251` - Activity History Collector 
	TemplateId *int32 `json:"template_id,omitempty"`
	// Options:  * `http` - Edge Applications (default)  * `waf` - WAF Events  * `cells_console` - Edge Functions  * `rtm_activity` - Activity History 
	DataSource NullableString `json:"data_source,omitempty"`
	Active NullableBool `json:"active,omitempty"`
	// Options' examples:  - `Standard HTTP/HTTPS POST` - { \"endpoint_type\": \"standard\", \"url\": \"http://example.com\", \"log_line_separator\": \"\\n\", \"payload_format\": \"$dataset\", \"max_size\": 1000024 }  - `Apache Kafka` - { \"endpoint_type\": \"kafka\", \"kafka_topic\": \"example_topic\", \"bootstrap_servers\": \"kafka-server.com:9092,kafka-server-2.com:9092\", \"use_tls\":true }  - `Simple Storage Service (S3)` - { \"endpoint_type\": \"s3\", \"access_key\": \"MYACCESSKEY\", \"region\": \"us-east-1\", \"object_key_prefix\": \"my_prefix_\", \"bucket_name\": \"bucket_example\", \"content_type\": \"plain/text\", \"host_url\": \"http://aws-host.com\", \"secret_key\": \"MYSECRETKEY\" }  - `Google BigQuery` - { \"endpoint_type\": \"big_query\", \"dataset_id\": \"my_dataset\", \"project_id\": \"my_project\", \"table_id\": \"my_table\", \"service_account_key\": \"{ \"service_account_key\": \"key_content\" }\" }  - `Elasticsearch` - { “endpoint_type”: \"elasticsearch\", “url”: “http://elasticsearch.com”, “api_key”: “XYZ_API_KEY” }  - `AWS Kinesis Data Firehose` -  { \"endpoint_type\": \"aws_kinesis_firehose\", \"access_key\": \"MYACCESSKEY\", \"stream_name\": \"my_stream_name\", \"region\": \"us-east-1\", \"secret_key\": \"MYSECRETKEY\" }  - `Datadog` - { \"endpoint_type\": \"datadog\", \"url\": \"https://http-intake.logs.datadoghq.com/v1/input\", \"api_key\": \"MYAPIKEY\" }  - `IBM QRadar` - { \"endpoint_type\": \"qradar\", \"url\": \"http://137.15.824.10:14440” }  - `Azure Monitor` - { \"endpoint_type\": \"azure_monitor\", \"log_type\": \"myLogType\", \"shared_key\": \"mysharedkey\", \"time_generated_field\": \"timeGeneratedField\", \"workspace_id\": \"anfhw-123sd-466gcs\"}  - `Azure Blob Storage` - { \"endpoint_type\": \"azure_blob_storage\", \"storage_account\": \"mystorageaccount\", \"container_name\": \"log_container\", \"blob_sas_token\": \"fd56e23e1f12efe\" }  - `Splunk` - { \"endpoint_type\": \"splunk\", \"url\": \"https://inputs.splunk-client.splunkcloud.com:1337/services/collector\", \"api_key\": \"MYAPIKEY\" } 
	Endpoint *string `json:"endpoint,omitempty"`
	// Note:  * Field not used with the rtm_activity data source. 
	DomainsIds []int32 `json:"domains_ids,omitempty"`
	// Note:  * Field not used with the rtm_activity data source. 
	AllDomains NullableBool `json:"all_domains,omitempty"`
	// Note:  * `Range` - From 0 to 100.  * `To use:` [Contact the sales team](https://www.azion.com/en/contact-sales/) to activate this feature in your account. 
	SamplingPercentage NullableInt32 `json:"sampling_percentage,omitempty"`
}

// NewDataStreamingPostBody instantiates a new DataStreamingPostBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataStreamingPostBody() *DataStreamingPostBody {
	this := DataStreamingPostBody{}
	var active bool = true
	this.Active = *NewNullableBool(&active)
	var allDomains bool = false
	this.AllDomains = *NewNullableBool(&allDomains)
	return &this
}

// NewDataStreamingPostBodyWithDefaults instantiates a new DataStreamingPostBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataStreamingPostBodyWithDefaults() *DataStreamingPostBody {
	this := DataStreamingPostBody{}
	var active bool = true
	this.Active = *NewNullableBool(&active)
	var allDomains bool = false
	this.AllDomains = *NewNullableBool(&allDomains)
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataStreamingPostBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingPostBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataStreamingPostBody) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataStreamingPostBody) SetName(v string) {
	o.Name = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *DataStreamingPostBody) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingPostBody) GetTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *DataStreamingPostBody) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *DataStreamingPostBody) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataStreamingPostBody) GetDataSource() string {
	if o == nil || IsNil(o.DataSource.Get()) {
		var ret string
		return ret
	}
	return *o.DataSource.Get()
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataStreamingPostBody) GetDataSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSource.Get(), o.DataSource.IsSet()
}

// HasDataSource returns a boolean if a field has been set.
func (o *DataStreamingPostBody) HasDataSource() bool {
	if o != nil && o.DataSource.IsSet() {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given NullableString and assigns it to the DataSource field.
func (o *DataStreamingPostBody) SetDataSource(v string) {
	o.DataSource.Set(&v)
}
// SetDataSourceNil sets the value for DataSource to be an explicit nil
func (o *DataStreamingPostBody) SetDataSourceNil() {
	o.DataSource.Set(nil)
}

// UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
func (o *DataStreamingPostBody) UnsetDataSource() {
	o.DataSource.Unset()
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataStreamingPostBody) GetActive() bool {
	if o == nil || IsNil(o.Active.Get()) {
		var ret bool
		return ret
	}
	return *o.Active.Get()
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataStreamingPostBody) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active.Get(), o.Active.IsSet()
}

// HasActive returns a boolean if a field has been set.
func (o *DataStreamingPostBody) HasActive() bool {
	if o != nil && o.Active.IsSet() {
		return true
	}

	return false
}

// SetActive gets a reference to the given NullableBool and assigns it to the Active field.
func (o *DataStreamingPostBody) SetActive(v bool) {
	o.Active.Set(&v)
}
// SetActiveNil sets the value for Active to be an explicit nil
func (o *DataStreamingPostBody) SetActiveNil() {
	o.Active.Set(nil)
}

// UnsetActive ensures that no value is present for Active, not even an explicit nil
func (o *DataStreamingPostBody) UnsetActive() {
	o.Active.Unset()
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *DataStreamingPostBody) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingPostBody) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *DataStreamingPostBody) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *DataStreamingPostBody) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetDomainsIds returns the DomainsIds field value if set, zero value otherwise.
func (o *DataStreamingPostBody) GetDomainsIds() []int32 {
	if o == nil || IsNil(o.DomainsIds) {
		var ret []int32
		return ret
	}
	return o.DomainsIds
}

// GetDomainsIdsOk returns a tuple with the DomainsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStreamingPostBody) GetDomainsIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.DomainsIds) {
		return nil, false
	}
	return o.DomainsIds, true
}

// HasDomainsIds returns a boolean if a field has been set.
func (o *DataStreamingPostBody) HasDomainsIds() bool {
	if o != nil && !IsNil(o.DomainsIds) {
		return true
	}

	return false
}

// SetDomainsIds gets a reference to the given []int32 and assigns it to the DomainsIds field.
func (o *DataStreamingPostBody) SetDomainsIds(v []int32) {
	o.DomainsIds = v
}

// GetAllDomains returns the AllDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataStreamingPostBody) GetAllDomains() bool {
	if o == nil || IsNil(o.AllDomains.Get()) {
		var ret bool
		return ret
	}
	return *o.AllDomains.Get()
}

// GetAllDomainsOk returns a tuple with the AllDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataStreamingPostBody) GetAllDomainsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllDomains.Get(), o.AllDomains.IsSet()
}

// HasAllDomains returns a boolean if a field has been set.
func (o *DataStreamingPostBody) HasAllDomains() bool {
	if o != nil && o.AllDomains.IsSet() {
		return true
	}

	return false
}

// SetAllDomains gets a reference to the given NullableBool and assigns it to the AllDomains field.
func (o *DataStreamingPostBody) SetAllDomains(v bool) {
	o.AllDomains.Set(&v)
}
// SetAllDomainsNil sets the value for AllDomains to be an explicit nil
func (o *DataStreamingPostBody) SetAllDomainsNil() {
	o.AllDomains.Set(nil)
}

// UnsetAllDomains ensures that no value is present for AllDomains, not even an explicit nil
func (o *DataStreamingPostBody) UnsetAllDomains() {
	o.AllDomains.Unset()
}

// GetSamplingPercentage returns the SamplingPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataStreamingPostBody) GetSamplingPercentage() int32 {
	if o == nil || IsNil(o.SamplingPercentage.Get()) {
		var ret int32
		return ret
	}
	return *o.SamplingPercentage.Get()
}

// GetSamplingPercentageOk returns a tuple with the SamplingPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataStreamingPostBody) GetSamplingPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SamplingPercentage.Get(), o.SamplingPercentage.IsSet()
}

// HasSamplingPercentage returns a boolean if a field has been set.
func (o *DataStreamingPostBody) HasSamplingPercentage() bool {
	if o != nil && o.SamplingPercentage.IsSet() {
		return true
	}

	return false
}

// SetSamplingPercentage gets a reference to the given NullableInt32 and assigns it to the SamplingPercentage field.
func (o *DataStreamingPostBody) SetSamplingPercentage(v int32) {
	o.SamplingPercentage.Set(&v)
}
// SetSamplingPercentageNil sets the value for SamplingPercentage to be an explicit nil
func (o *DataStreamingPostBody) SetSamplingPercentageNil() {
	o.SamplingPercentage.Set(nil)
}

// UnsetSamplingPercentage ensures that no value is present for SamplingPercentage, not even an explicit nil
func (o *DataStreamingPostBody) UnsetSamplingPercentage() {
	o.SamplingPercentage.Unset()
}

func (o DataStreamingPostBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataStreamingPostBody) ToMap() (map[string]interface{}, error) {
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
	if o.SamplingPercentage.IsSet() {
		toSerialize["sampling_percentage"] = o.SamplingPercentage.Get()
	}
	return toSerialize, nil
}

type NullableDataStreamingPostBody struct {
	value *DataStreamingPostBody
	isSet bool
}

func (v NullableDataStreamingPostBody) Get() *DataStreamingPostBody {
	return v.value
}

func (v *NullableDataStreamingPostBody) Set(val *DataStreamingPostBody) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStreamingPostBody) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStreamingPostBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStreamingPostBody(val *DataStreamingPostBody) *NullableDataStreamingPostBody {
	return &NullableDataStreamingPostBody{value: val, isSet: true}
}

func (v NullableDataStreamingPostBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStreamingPostBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


