# OriginsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** |  | 
**TotalPages** | **int64** |  | 
**SchemaVersion** | **int64** |  | 
**Links** | [**OriginsResponseLinks**](OriginsResponseLinks.md) |  | 
**Results** | [**[]OriginsResultResponse**](OriginsResultResponse.md) |  | 

## Methods

### NewOriginsResponse

`func NewOriginsResponse(count int64, totalPages int64, schemaVersion int64, links OriginsResponseLinks, results []OriginsResultResponse, ) *OriginsResponse`

NewOriginsResponse instantiates a new OriginsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginsResponseWithDefaults

`func NewOriginsResponseWithDefaults() *OriginsResponse`

NewOriginsResponseWithDefaults instantiates a new OriginsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *OriginsResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OriginsResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OriginsResponse) SetCount(v int64)`

SetCount sets Count field to given value.


### GetTotalPages

`func (o *OriginsResponse) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *OriginsResponse) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *OriginsResponse) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.


### GetSchemaVersion

`func (o *OriginsResponse) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *OriginsResponse) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *OriginsResponse) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetLinks

`func (o *OriginsResponse) GetLinks() OriginsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OriginsResponse) GetLinksOk() (*OriginsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OriginsResponse) SetLinks(v OriginsResponseLinks)`

SetLinks sets Links field to given value.


### GetResults

`func (o *OriginsResponse) GetResults() []OriginsResultResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OriginsResponse) GetResultsOk() (*[]OriginsResultResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OriginsResponse) SetResults(v []OriginsResultResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


