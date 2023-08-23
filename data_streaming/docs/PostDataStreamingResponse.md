# PostDataStreamingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **int32** | Options:  * &#x60;2&#x60; - Edge Applications Event Collector  * &#x60;4&#x60; - WAF Event Collector  * &#x60;86&#x60; - Edge Functions Event Collector  * &#x60;184&#x60; - Edge Applications + WAF Event Collector  * &#x60;251&#x60; - Activity History Collector  | [optional] 
**DataSource** | Pointer to **string** | Options:  * &#x60;http&#x60; - Edge Applications  * &#x60;waf&#x60; - WAF Events  * &#x60;cells_console&#x60; - Edge Functions  * &#x60;rtm_activity&#x60; - Activity History  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Endpoint** | Pointer to [**[]PostDataStreamingResponseEndpointInner**](PostDataStreamingResponseEndpointInner.md) |  | [optional] 
**AllDomains** | Pointer to **NullableBool** | Note:  * Field not used with the rtm_activity data source.  | [optional] 

## Methods

### NewPostDataStreamingResponse

`func NewPostDataStreamingResponse() *PostDataStreamingResponse`

NewPostDataStreamingResponse instantiates a new PostDataStreamingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDataStreamingResponseWithDefaults

`func NewPostDataStreamingResponseWithDefaults() *PostDataStreamingResponse`

NewPostDataStreamingResponseWithDefaults instantiates a new PostDataStreamingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostDataStreamingResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostDataStreamingResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostDataStreamingResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PostDataStreamingResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PostDataStreamingResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostDataStreamingResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostDataStreamingResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostDataStreamingResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplateId

`func (o *PostDataStreamingResponse) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *PostDataStreamingResponse) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *PostDataStreamingResponse) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *PostDataStreamingResponse) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetDataSource

`func (o *PostDataStreamingResponse) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *PostDataStreamingResponse) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *PostDataStreamingResponse) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *PostDataStreamingResponse) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetActive

`func (o *PostDataStreamingResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PostDataStreamingResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PostDataStreamingResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PostDataStreamingResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEndpoint

`func (o *PostDataStreamingResponse) GetEndpoint() []PostDataStreamingResponseEndpointInner`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PostDataStreamingResponse) GetEndpointOk() (*[]PostDataStreamingResponseEndpointInner, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PostDataStreamingResponse) SetEndpoint(v []PostDataStreamingResponseEndpointInner)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *PostDataStreamingResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAllDomains

`func (o *PostDataStreamingResponse) GetAllDomains() bool`

GetAllDomains returns the AllDomains field if non-nil, zero value otherwise.

### GetAllDomainsOk

`func (o *PostDataStreamingResponse) GetAllDomainsOk() (*bool, bool)`

GetAllDomainsOk returns a tuple with the AllDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDomains

`func (o *PostDataStreamingResponse) SetAllDomains(v bool)`

SetAllDomains sets AllDomains field to given value.

### HasAllDomains

`func (o *PostDataStreamingResponse) HasAllDomains() bool`

HasAllDomains returns a boolean if a field has been set.

### SetAllDomainsNil

`func (o *PostDataStreamingResponse) SetAllDomainsNil(b bool)`

 SetAllDomainsNil sets the value for AllDomains to be an explicit nil

### UnsetAllDomains
`func (o *PostDataStreamingResponse) UnsetAllDomains()`

UnsetAllDomains ensures that no value is present for AllDomains, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


