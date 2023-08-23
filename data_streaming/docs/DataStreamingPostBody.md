# DataStreamingPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **int32** | Options:  * &#x60;2&#x60; - Edge Applications Event Collector  * &#x60;4&#x60; - WAF Event Collector  * &#x60;86&#x60; - Edge Functions Event Collector  * &#x60;184&#x60; - Edge Applications + WAF Event Collector  * &#x60;251&#x60; - Activity History Collector  | [optional] 
**DataSource** | Pointer to **NullableString** | Options:  * &#x60;http&#x60; - Edge Applications (default)  * &#x60;waf&#x60; - WAF Events  * &#x60;cells_console&#x60; - Edge Functions  * &#x60;rtm_activity&#x60; - Activity History  | [optional] 
**Active** | Pointer to **NullableBool** |  | [optional] [default to true]
**Endpoint** | Pointer to **string** | Options&#39; examples:  - &#x60;Standard HTTP/HTTPS POST&#x60; - { \&quot;endpoint_type\&quot;: \&quot;standard\&quot;, \&quot;url\&quot;: \&quot;http://example.com\&quot;, \&quot;log_line_separator\&quot;: \&quot;\\n\&quot;, \&quot;payload_format\&quot;: \&quot;$dataset\&quot;, \&quot;max_size\&quot;: 1000024 }  - &#x60;Apache Kafka&#x60; - { \&quot;endpoint_type\&quot;: \&quot;kafka\&quot;, \&quot;kafka_topic\&quot;: \&quot;example_topic\&quot;, \&quot;bootstrap_servers\&quot;: \&quot;kafka-server.com:9092,kafka-server-2.com:9092\&quot;, \&quot;use_tls\&quot;:true }  - &#x60;Simple Storage Service (S3)&#x60; - { \&quot;endpoint_type\&quot;: \&quot;s3\&quot;, \&quot;access_key\&quot;: \&quot;MYACCESSKEY\&quot;, \&quot;region\&quot;: \&quot;us-east-1\&quot;, \&quot;object_key_prefix\&quot;: \&quot;my_prefix_\&quot;, \&quot;bucket_name\&quot;: \&quot;bucket_example\&quot;, \&quot;content_type\&quot;: \&quot;plain/text\&quot;, \&quot;host_url\&quot;: \&quot;http://aws-host.com\&quot;, \&quot;secret_key\&quot;: \&quot;MYSECRETKEY\&quot; }  - &#x60;Google BigQuery&#x60; - { \&quot;endpoint_type\&quot;: \&quot;big_query\&quot;, \&quot;dataset_id\&quot;: \&quot;my_dataset\&quot;, \&quot;project_id\&quot;: \&quot;my_project\&quot;, \&quot;table_id\&quot;: \&quot;my_table\&quot;, \&quot;service_account_key\&quot;: \&quot;{ \&quot;service_account_key\&quot;: \&quot;key_content\&quot; }\&quot; }  - &#x60;Elasticsearch&#x60; - { “endpoint_type”: \&quot;elasticsearch\&quot;, “url”: “http://elasticsearch.com”, “api_key”: “XYZ_API_KEY” }  - &#x60;AWS Kinesis Data Firehose&#x60; -  { \&quot;endpoint_type\&quot;: \&quot;aws_kinesis_firehose\&quot;, \&quot;access_key\&quot;: \&quot;MYACCESSKEY\&quot;, \&quot;stream_name\&quot;: \&quot;my_stream_name\&quot;, \&quot;region\&quot;: \&quot;us-east-1\&quot;, \&quot;secret_key\&quot;: \&quot;MYSECRETKEY\&quot; }  - &#x60;Datadog&#x60; - { \&quot;endpoint_type\&quot;: \&quot;datadog\&quot;, \&quot;url\&quot;: \&quot;https://http-intake.logs.datadoghq.com/v1/input\&quot;, \&quot;api_key\&quot;: \&quot;MYAPIKEY\&quot; }  - &#x60;IBM QRadar&#x60; - { \&quot;endpoint_type\&quot;: \&quot;qradar\&quot;, \&quot;url\&quot;: \&quot;http://137.15.824.10:14440” }  - &#x60;Azure Monitor&#x60; - { \&quot;endpoint_type\&quot;: \&quot;azure_monitor\&quot;, \&quot;log_type\&quot;: \&quot;myLogType\&quot;, \&quot;shared_key\&quot;: \&quot;mysharedkey\&quot;, \&quot;time_generated_field\&quot;: \&quot;timeGeneratedField\&quot;, \&quot;workspace_id\&quot;: \&quot;anfhw-123sd-466gcs\&quot;}  - &#x60;Azure Blob Storage&#x60; - { \&quot;endpoint_type\&quot;: \&quot;azure_blob_storage\&quot;, \&quot;storage_account\&quot;: \&quot;mystorageaccount\&quot;, \&quot;container_name\&quot;: \&quot;log_container\&quot;, \&quot;blob_sas_token\&quot;: \&quot;fd56e23e1f12efe\&quot; }  - &#x60;Splunk&#x60; - { \&quot;endpoint_type\&quot;: \&quot;splunk\&quot;, \&quot;url\&quot;: \&quot;https://inputs.splunk-client.splunkcloud.com:1337/services/collector\&quot;, \&quot;api_key\&quot;: \&quot;MYAPIKEY\&quot; }  | [optional] 
**DomainsIds** | Pointer to **[]int32** | Note:  * Field not used with the rtm_activity data source.  | [optional] 
**AllDomains** | Pointer to **NullableBool** | Note:  * Field not used with the rtm_activity data source.  | [optional] [default to false]
**SamplingPercentage** | Pointer to **NullableInt32** | Note:  * &#x60;Range&#x60; - From 0 to 100.  * &#x60;To use:&#x60; [Contact the sales team](https://www.azion.com/en/contact-sales/) to activate this feature in your account.  | [optional] 

## Methods

### NewDataStreamingPostBody

`func NewDataStreamingPostBody() *DataStreamingPostBody`

NewDataStreamingPostBody instantiates a new DataStreamingPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStreamingPostBodyWithDefaults

`func NewDataStreamingPostBodyWithDefaults() *DataStreamingPostBody`

NewDataStreamingPostBodyWithDefaults instantiates a new DataStreamingPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataStreamingPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataStreamingPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataStreamingPostBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataStreamingPostBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplateId

`func (o *DataStreamingPostBody) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DataStreamingPostBody) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DataStreamingPostBody) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DataStreamingPostBody) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetDataSource

`func (o *DataStreamingPostBody) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *DataStreamingPostBody) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *DataStreamingPostBody) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *DataStreamingPostBody) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### SetDataSourceNil

`func (o *DataStreamingPostBody) SetDataSourceNil(b bool)`

 SetDataSourceNil sets the value for DataSource to be an explicit nil

### UnsetDataSource
`func (o *DataStreamingPostBody) UnsetDataSource()`

UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
### GetActive

`func (o *DataStreamingPostBody) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DataStreamingPostBody) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DataStreamingPostBody) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DataStreamingPostBody) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *DataStreamingPostBody) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *DataStreamingPostBody) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetEndpoint

`func (o *DataStreamingPostBody) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DataStreamingPostBody) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DataStreamingPostBody) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *DataStreamingPostBody) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetDomainsIds

`func (o *DataStreamingPostBody) GetDomainsIds() []int32`

GetDomainsIds returns the DomainsIds field if non-nil, zero value otherwise.

### GetDomainsIdsOk

`func (o *DataStreamingPostBody) GetDomainsIdsOk() (*[]int32, bool)`

GetDomainsIdsOk returns a tuple with the DomainsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainsIds

`func (o *DataStreamingPostBody) SetDomainsIds(v []int32)`

SetDomainsIds sets DomainsIds field to given value.

### HasDomainsIds

`func (o *DataStreamingPostBody) HasDomainsIds() bool`

HasDomainsIds returns a boolean if a field has been set.

### GetAllDomains

`func (o *DataStreamingPostBody) GetAllDomains() bool`

GetAllDomains returns the AllDomains field if non-nil, zero value otherwise.

### GetAllDomainsOk

`func (o *DataStreamingPostBody) GetAllDomainsOk() (*bool, bool)`

GetAllDomainsOk returns a tuple with the AllDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDomains

`func (o *DataStreamingPostBody) SetAllDomains(v bool)`

SetAllDomains sets AllDomains field to given value.

### HasAllDomains

`func (o *DataStreamingPostBody) HasAllDomains() bool`

HasAllDomains returns a boolean if a field has been set.

### SetAllDomainsNil

`func (o *DataStreamingPostBody) SetAllDomainsNil(b bool)`

 SetAllDomainsNil sets the value for AllDomains to be an explicit nil

### UnsetAllDomains
`func (o *DataStreamingPostBody) UnsetAllDomains()`

UnsetAllDomains ensures that no value is present for AllDomains, not even an explicit nil
### GetSamplingPercentage

`func (o *DataStreamingPostBody) GetSamplingPercentage() int32`

GetSamplingPercentage returns the SamplingPercentage field if non-nil, zero value otherwise.

### GetSamplingPercentageOk

`func (o *DataStreamingPostBody) GetSamplingPercentageOk() (*int32, bool)`

GetSamplingPercentageOk returns a tuple with the SamplingPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingPercentage

`func (o *DataStreamingPostBody) SetSamplingPercentage(v int32)`

SetSamplingPercentage sets SamplingPercentage field to given value.

### HasSamplingPercentage

`func (o *DataStreamingPostBody) HasSamplingPercentage() bool`

HasSamplingPercentage returns a boolean if a field has been set.

### SetSamplingPercentageNil

`func (o *DataStreamingPostBody) SetSamplingPercentageNil(b bool)`

 SetSamplingPercentageNil sets the value for SamplingPercentage to be an explicit nil

### UnsetSamplingPercentage
`func (o *DataStreamingPostBody) UnsetSamplingPercentage()`

UnsetSamplingPercentage ensures that no value is present for SamplingPercentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


