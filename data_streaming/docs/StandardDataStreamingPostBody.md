# StandardDataStreamingPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **int32** | Options:  * &#x60;2&#x60; - Edge Applications Event Collector  * &#x60;4&#x60; - WAF Event Collector  * &#x60;86&#x60; - Edge Functions Event Collector  * &#x60;184&#x60; - Edge Applications + WAF Event Collector  * &#x60;251&#x60; - Activity History Collector  | [optional] 
**DataSource** | Pointer to **NullableString** | Options:  * &#x60;http&#x60; - Edge Applications (default)  * &#x60;waf&#x60; - WAF Events  * &#x60;cells_console&#x60; - Edge Functions  * &#x60;rtm_activity&#x60; - Activity History  | [optional] 
**Active** | Pointer to **NullableBool** |  | [optional] [default to true]
**Endpoint** | Pointer to [**DataStreamingEndpointTypeStandard**](DataStreamingEndpointTypeStandard.md) |  | [optional] 
**DomainsIds** | Pointer to **[]int32** | Note:  * Field not used with the rtm_activity data source.  | [optional] 
**AllDomains** | Pointer to **NullableBool** | Note:  * Field not used with the rtm_activity data source.  | [optional] [default to false]

## Methods

### NewStandardDataStreamingPostBody

`func NewStandardDataStreamingPostBody() *StandardDataStreamingPostBody`

NewStandardDataStreamingPostBody instantiates a new StandardDataStreamingPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardDataStreamingPostBodyWithDefaults

`func NewStandardDataStreamingPostBodyWithDefaults() *StandardDataStreamingPostBody`

NewStandardDataStreamingPostBodyWithDefaults instantiates a new StandardDataStreamingPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StandardDataStreamingPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandardDataStreamingPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandardDataStreamingPostBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StandardDataStreamingPostBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplateId

`func (o *StandardDataStreamingPostBody) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *StandardDataStreamingPostBody) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *StandardDataStreamingPostBody) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *StandardDataStreamingPostBody) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetDataSource

`func (o *StandardDataStreamingPostBody) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *StandardDataStreamingPostBody) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *StandardDataStreamingPostBody) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *StandardDataStreamingPostBody) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### SetDataSourceNil

`func (o *StandardDataStreamingPostBody) SetDataSourceNil(b bool)`

 SetDataSourceNil sets the value for DataSource to be an explicit nil

### UnsetDataSource
`func (o *StandardDataStreamingPostBody) UnsetDataSource()`

UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
### GetActive

`func (o *StandardDataStreamingPostBody) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StandardDataStreamingPostBody) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StandardDataStreamingPostBody) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *StandardDataStreamingPostBody) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *StandardDataStreamingPostBody) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *StandardDataStreamingPostBody) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetEndpoint

`func (o *StandardDataStreamingPostBody) GetEndpoint() DataStreamingEndpointTypeStandard`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *StandardDataStreamingPostBody) GetEndpointOk() (*DataStreamingEndpointTypeStandard, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *StandardDataStreamingPostBody) SetEndpoint(v DataStreamingEndpointTypeStandard)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *StandardDataStreamingPostBody) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetDomainsIds

`func (o *StandardDataStreamingPostBody) GetDomainsIds() []int32`

GetDomainsIds returns the DomainsIds field if non-nil, zero value otherwise.

### GetDomainsIdsOk

`func (o *StandardDataStreamingPostBody) GetDomainsIdsOk() (*[]int32, bool)`

GetDomainsIdsOk returns a tuple with the DomainsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainsIds

`func (o *StandardDataStreamingPostBody) SetDomainsIds(v []int32)`

SetDomainsIds sets DomainsIds field to given value.

### HasDomainsIds

`func (o *StandardDataStreamingPostBody) HasDomainsIds() bool`

HasDomainsIds returns a boolean if a field has been set.

### GetAllDomains

`func (o *StandardDataStreamingPostBody) GetAllDomains() bool`

GetAllDomains returns the AllDomains field if non-nil, zero value otherwise.

### GetAllDomainsOk

`func (o *StandardDataStreamingPostBody) GetAllDomainsOk() (*bool, bool)`

GetAllDomainsOk returns a tuple with the AllDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDomains

`func (o *StandardDataStreamingPostBody) SetAllDomains(v bool)`

SetAllDomains sets AllDomains field to given value.

### HasAllDomains

`func (o *StandardDataStreamingPostBody) HasAllDomains() bool`

HasAllDomains returns a boolean if a field has been set.

### SetAllDomainsNil

`func (o *StandardDataStreamingPostBody) SetAllDomainsNil(b bool)`

 SetAllDomainsNil sets the value for AllDomains to be an explicit nil

### UnsetAllDomains
`func (o *StandardDataStreamingPostBody) UnsetAllDomains()`

UnsetAllDomains ensures that no value is present for AllDomains, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


