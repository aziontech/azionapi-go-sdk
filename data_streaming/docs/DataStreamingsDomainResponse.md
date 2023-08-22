# DataStreamingsDomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**SchemaVersion** | Pointer to **float32** |  | [optional] 
**Links** | Pointer to [**DataStreamingsDomainResponseLinks**](DataStreamingsDomainResponseLinks.md) |  | [optional] 
**Results** | Pointer to [**[]DataStreamingsDomainResult**](DataStreamingsDomainResult.md) |  | [optional] 

## Methods

### NewDataStreamingsDomainResponse

`func NewDataStreamingsDomainResponse() *DataStreamingsDomainResponse`

NewDataStreamingsDomainResponse instantiates a new DataStreamingsDomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStreamingsDomainResponseWithDefaults

`func NewDataStreamingsDomainResponseWithDefaults() *DataStreamingsDomainResponse`

NewDataStreamingsDomainResponseWithDefaults instantiates a new DataStreamingsDomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DataStreamingsDomainResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DataStreamingsDomainResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DataStreamingsDomainResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DataStreamingsDomainResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *DataStreamingsDomainResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *DataStreamingsDomainResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *DataStreamingsDomainResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *DataStreamingsDomainResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *DataStreamingsDomainResponse) GetSchemaVersion() float32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *DataStreamingsDomainResponse) GetSchemaVersionOk() (*float32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *DataStreamingsDomainResponse) SetSchemaVersion(v float32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *DataStreamingsDomainResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetLinks

`func (o *DataStreamingsDomainResponse) GetLinks() DataStreamingsDomainResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DataStreamingsDomainResponse) GetLinksOk() (*DataStreamingsDomainResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DataStreamingsDomainResponse) SetLinks(v DataStreamingsDomainResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DataStreamingsDomainResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *DataStreamingsDomainResponse) GetResults() []DataStreamingsDomainResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DataStreamingsDomainResponse) GetResultsOk() (*[]DataStreamingsDomainResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DataStreamingsDomainResponse) SetResults(v []DataStreamingsDomainResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *DataStreamingsDomainResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


