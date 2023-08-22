# CreateNewDataStreamingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **int32** | Options:  * &#x60;2&#x60; - Edge Applications Event Collector  * &#x60;4&#x60; - WAF Event Collector  * &#x60;86&#x60; - Edge Functions Event Collector  * &#x60;184&#x60; - Edge Applications + WAF Event Collector  * &#x60;251&#x60; - Activity History Collector  | [optional] 
**DataSource** | Pointer to **NullableString** | Options:  * &#x60;http&#x60; - Edge Applications (default)  * &#x60;waf&#x60; - WAF Events  * &#x60;cells_console&#x60; - Edge Functions  * &#x60;rtm_activity&#x60; - Activity History    | [optional] 
**Active** | Pointer to **NullableBool** |  | [optional] [default to true]
**Endpoint** | Pointer to [**DataStreamingEndpointTypeStandard**](DataStreamingEndpointTypeStandard.md) |  | [optional] 
**DomainsIds** | Pointer to **[]int32** | Note:  * Field not used with the rtm_activity data source.  | [optional] 
**AllDomains** | Pointer to **NullableBool** | Note:  * Field not used with the rtm_activity data source.  | [optional] [default to false]
**SamplingPercentage** | Pointer to **NullableInt32** | Note:  * &#x60;Range&#x60; - From 0 to 100.  * &#x60;To use:&#x60; [Contact the sales team](https://www.azion.com/en/contact-sales/) to activate this feature in your account.  | [optional] 
**TemplateModel** | Pointer to **string** | Note:  * Add all variables and values that will be used to stream according to the data source you choose to use.    * All data streaming [variables can be found on the reference documentation](https://www.azion.com/en/documentation/products/data-streaming/#selecting-data-sources).    | [optional] 

## Methods

### NewCreateNewDataStreamingRequest

`func NewCreateNewDataStreamingRequest() *CreateNewDataStreamingRequest`

NewCreateNewDataStreamingRequest instantiates a new CreateNewDataStreamingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNewDataStreamingRequestWithDefaults

`func NewCreateNewDataStreamingRequestWithDefaults() *CreateNewDataStreamingRequest`

NewCreateNewDataStreamingRequestWithDefaults instantiates a new CreateNewDataStreamingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNewDataStreamingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNewDataStreamingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNewDataStreamingRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNewDataStreamingRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplateId

`func (o *CreateNewDataStreamingRequest) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CreateNewDataStreamingRequest) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CreateNewDataStreamingRequest) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CreateNewDataStreamingRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetDataSource

`func (o *CreateNewDataStreamingRequest) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *CreateNewDataStreamingRequest) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *CreateNewDataStreamingRequest) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *CreateNewDataStreamingRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### SetDataSourceNil

`func (o *CreateNewDataStreamingRequest) SetDataSourceNil(b bool)`

 SetDataSourceNil sets the value for DataSource to be an explicit nil

### UnsetDataSource
`func (o *CreateNewDataStreamingRequest) UnsetDataSource()`

UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
### GetActive

`func (o *CreateNewDataStreamingRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateNewDataStreamingRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateNewDataStreamingRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateNewDataStreamingRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *CreateNewDataStreamingRequest) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *CreateNewDataStreamingRequest) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetEndpoint

`func (o *CreateNewDataStreamingRequest) GetEndpoint() DataStreamingEndpointTypeStandard`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CreateNewDataStreamingRequest) GetEndpointOk() (*DataStreamingEndpointTypeStandard, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CreateNewDataStreamingRequest) SetEndpoint(v DataStreamingEndpointTypeStandard)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *CreateNewDataStreamingRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetDomainsIds

`func (o *CreateNewDataStreamingRequest) GetDomainsIds() []int32`

GetDomainsIds returns the DomainsIds field if non-nil, zero value otherwise.

### GetDomainsIdsOk

`func (o *CreateNewDataStreamingRequest) GetDomainsIdsOk() (*[]int32, bool)`

GetDomainsIdsOk returns a tuple with the DomainsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainsIds

`func (o *CreateNewDataStreamingRequest) SetDomainsIds(v []int32)`

SetDomainsIds sets DomainsIds field to given value.

### HasDomainsIds

`func (o *CreateNewDataStreamingRequest) HasDomainsIds() bool`

HasDomainsIds returns a boolean if a field has been set.

### GetAllDomains

`func (o *CreateNewDataStreamingRequest) GetAllDomains() bool`

GetAllDomains returns the AllDomains field if non-nil, zero value otherwise.

### GetAllDomainsOk

`func (o *CreateNewDataStreamingRequest) GetAllDomainsOk() (*bool, bool)`

GetAllDomainsOk returns a tuple with the AllDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDomains

`func (o *CreateNewDataStreamingRequest) SetAllDomains(v bool)`

SetAllDomains sets AllDomains field to given value.

### HasAllDomains

`func (o *CreateNewDataStreamingRequest) HasAllDomains() bool`

HasAllDomains returns a boolean if a field has been set.

### SetAllDomainsNil

`func (o *CreateNewDataStreamingRequest) SetAllDomainsNil(b bool)`

 SetAllDomainsNil sets the value for AllDomains to be an explicit nil

### UnsetAllDomains
`func (o *CreateNewDataStreamingRequest) UnsetAllDomains()`

UnsetAllDomains ensures that no value is present for AllDomains, not even an explicit nil
### GetSamplingPercentage

`func (o *CreateNewDataStreamingRequest) GetSamplingPercentage() int32`

GetSamplingPercentage returns the SamplingPercentage field if non-nil, zero value otherwise.

### GetSamplingPercentageOk

`func (o *CreateNewDataStreamingRequest) GetSamplingPercentageOk() (*int32, bool)`

GetSamplingPercentageOk returns a tuple with the SamplingPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingPercentage

`func (o *CreateNewDataStreamingRequest) SetSamplingPercentage(v int32)`

SetSamplingPercentage sets SamplingPercentage field to given value.

### HasSamplingPercentage

`func (o *CreateNewDataStreamingRequest) HasSamplingPercentage() bool`

HasSamplingPercentage returns a boolean if a field has been set.

### SetSamplingPercentageNil

`func (o *CreateNewDataStreamingRequest) SetSamplingPercentageNil(b bool)`

 SetSamplingPercentageNil sets the value for SamplingPercentage to be an explicit nil

### UnsetSamplingPercentage
`func (o *CreateNewDataStreamingRequest) UnsetSamplingPercentage()`

UnsetSamplingPercentage ensures that no value is present for SamplingPercentage, not even an explicit nil
### GetTemplateModel

`func (o *CreateNewDataStreamingRequest) GetTemplateModel() string`

GetTemplateModel returns the TemplateModel field if non-nil, zero value otherwise.

### GetTemplateModelOk

`func (o *CreateNewDataStreamingRequest) GetTemplateModelOk() (*string, bool)`

GetTemplateModelOk returns a tuple with the TemplateModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateModel

`func (o *CreateNewDataStreamingRequest) SetTemplateModel(v string)`

SetTemplateModel sets TemplateModel field to given value.

### HasTemplateModel

`func (o *CreateNewDataStreamingRequest) HasTemplateModel() bool`

HasTemplateModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


