# PostCustomDataStreamingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DataSource** | Pointer to **string** | Options:  * &#x60;http&#x60; - Edge Applications  * &#x60;waf&#x60; - WAF Events  * &#x60;cells_console&#x60; - Edge Functions  * &#x60;rtm_activity&#x60; - Activity History    | [optional] 
**TemplateModel** | Pointer to **string** | Note:  * Add all variables and values that will be used to stream according to the data source you choose to use.    | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**AllDomains** | Pointer to **NullableBool** | Note:  * Field not used with the rtm_activity data source.  | [optional] 

## Methods

### NewPostCustomDataStreamingResponse

`func NewPostCustomDataStreamingResponse() *PostCustomDataStreamingResponse`

NewPostCustomDataStreamingResponse instantiates a new PostCustomDataStreamingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCustomDataStreamingResponseWithDefaults

`func NewPostCustomDataStreamingResponseWithDefaults() *PostCustomDataStreamingResponse`

NewPostCustomDataStreamingResponseWithDefaults instantiates a new PostCustomDataStreamingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostCustomDataStreamingResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCustomDataStreamingResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCustomDataStreamingResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PostCustomDataStreamingResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PostCustomDataStreamingResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostCustomDataStreamingResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostCustomDataStreamingResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostCustomDataStreamingResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataSource

`func (o *PostCustomDataStreamingResponse) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *PostCustomDataStreamingResponse) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *PostCustomDataStreamingResponse) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *PostCustomDataStreamingResponse) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetTemplateModel

`func (o *PostCustomDataStreamingResponse) GetTemplateModel() string`

GetTemplateModel returns the TemplateModel field if non-nil, zero value otherwise.

### GetTemplateModelOk

`func (o *PostCustomDataStreamingResponse) GetTemplateModelOk() (*string, bool)`

GetTemplateModelOk returns a tuple with the TemplateModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateModel

`func (o *PostCustomDataStreamingResponse) SetTemplateModel(v string)`

SetTemplateModel sets TemplateModel field to given value.

### HasTemplateModel

`func (o *PostCustomDataStreamingResponse) HasTemplateModel() bool`

HasTemplateModel returns a boolean if a field has been set.

### GetActive

`func (o *PostCustomDataStreamingResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PostCustomDataStreamingResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PostCustomDataStreamingResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PostCustomDataStreamingResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEndpoint

`func (o *PostCustomDataStreamingResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PostCustomDataStreamingResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PostCustomDataStreamingResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *PostCustomDataStreamingResponse) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAllDomains

`func (o *PostCustomDataStreamingResponse) GetAllDomains() bool`

GetAllDomains returns the AllDomains field if non-nil, zero value otherwise.

### GetAllDomainsOk

`func (o *PostCustomDataStreamingResponse) GetAllDomainsOk() (*bool, bool)`

GetAllDomainsOk returns a tuple with the AllDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDomains

`func (o *PostCustomDataStreamingResponse) SetAllDomains(v bool)`

SetAllDomains sets AllDomains field to given value.

### HasAllDomains

`func (o *PostCustomDataStreamingResponse) HasAllDomains() bool`

HasAllDomains returns a boolean if a field has been set.

### SetAllDomainsNil

`func (o *PostCustomDataStreamingResponse) SetAllDomainsNil(b bool)`

 SetAllDomainsNil sets the value for AllDomains to be an explicit nil

### UnsetAllDomains
`func (o *PostCustomDataStreamingResponse) UnsetAllDomains()`

UnsetAllDomains ensures that no value is present for AllDomains, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


