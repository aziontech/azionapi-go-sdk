# DataStreamingResponseWithResultsResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **int32** |  | [optional] 
**DataSource** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Endpoint** | Pointer to [**DataStreamingEndpointTypeKafka**](DataStreamingEndpointTypeKafka.md) |  | [optional] 
**AllDomains** | Pointer to **bool** |  | [optional] 
**TemplateModel** | Pointer to **string** |  | [optional] 

## Methods

### NewDataStreamingResponseWithResultsResultsInner

`func NewDataStreamingResponseWithResultsResultsInner() *DataStreamingResponseWithResultsResultsInner`

NewDataStreamingResponseWithResultsResultsInner instantiates a new DataStreamingResponseWithResultsResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStreamingResponseWithResultsResultsInnerWithDefaults

`func NewDataStreamingResponseWithResultsResultsInnerWithDefaults() *DataStreamingResponseWithResultsResultsInner`

NewDataStreamingResponseWithResultsResultsInnerWithDefaults instantiates a new DataStreamingResponseWithResultsResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataStreamingResponseWithResultsResultsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataStreamingResponseWithResultsResultsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataStreamingResponseWithResultsResultsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DataStreamingResponseWithResultsResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DataStreamingResponseWithResultsResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataStreamingResponseWithResultsResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataStreamingResponseWithResultsResultsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataStreamingResponseWithResultsResultsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTemplateId

`func (o *DataStreamingResponseWithResultsResultsInner) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *DataStreamingResponseWithResultsResultsInner) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *DataStreamingResponseWithResultsResultsInner) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *DataStreamingResponseWithResultsResultsInner) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetDataSource

`func (o *DataStreamingResponseWithResultsResultsInner) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *DataStreamingResponseWithResultsResultsInner) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *DataStreamingResponseWithResultsResultsInner) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *DataStreamingResponseWithResultsResultsInner) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetActive

`func (o *DataStreamingResponseWithResultsResultsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DataStreamingResponseWithResultsResultsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DataStreamingResponseWithResultsResultsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *DataStreamingResponseWithResultsResultsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEndpoint

`func (o *DataStreamingResponseWithResultsResultsInner) GetEndpoint() DataStreamingEndpointTypeKafka`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DataStreamingResponseWithResultsResultsInner) GetEndpointOk() (*DataStreamingEndpointTypeKafka, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DataStreamingResponseWithResultsResultsInner) SetEndpoint(v DataStreamingEndpointTypeKafka)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *DataStreamingResponseWithResultsResultsInner) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAllDomains

`func (o *DataStreamingResponseWithResultsResultsInner) GetAllDomains() bool`

GetAllDomains returns the AllDomains field if non-nil, zero value otherwise.

### GetAllDomainsOk

`func (o *DataStreamingResponseWithResultsResultsInner) GetAllDomainsOk() (*bool, bool)`

GetAllDomainsOk returns a tuple with the AllDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDomains

`func (o *DataStreamingResponseWithResultsResultsInner) SetAllDomains(v bool)`

SetAllDomains sets AllDomains field to given value.

### HasAllDomains

`func (o *DataStreamingResponseWithResultsResultsInner) HasAllDomains() bool`

HasAllDomains returns a boolean if a field has been set.

### GetTemplateModel

`func (o *DataStreamingResponseWithResultsResultsInner) GetTemplateModel() string`

GetTemplateModel returns the TemplateModel field if non-nil, zero value otherwise.

### GetTemplateModelOk

`func (o *DataStreamingResponseWithResultsResultsInner) GetTemplateModelOk() (*string, bool)`

GetTemplateModelOk returns a tuple with the TemplateModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateModel

`func (o *DataStreamingResponseWithResultsResultsInner) SetTemplateModel(v string)`

SetTemplateModel sets TemplateModel field to given value.

### HasTemplateModel

`func (o *DataStreamingResponseWithResultsResultsInner) HasTemplateModel() bool`

HasTemplateModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


