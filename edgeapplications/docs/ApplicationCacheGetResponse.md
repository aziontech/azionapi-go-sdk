# ApplicationCacheGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** |  | 
**TotalPages** | **int64** |  | 
**SchemaVersion** | **int64** |  | 
**Links** | [**ApplicationLinks**](ApplicationLinks.md) |  | 
**Results** | [**[]ApplicationCacheResults**](ApplicationCacheResults.md) |  | 

## Methods

### NewApplicationCacheGetResponse

`func NewApplicationCacheGetResponse(count int64, totalPages int64, schemaVersion int64, links ApplicationLinks, results []ApplicationCacheResults, ) *ApplicationCacheGetResponse`

NewApplicationCacheGetResponse instantiates a new ApplicationCacheGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCacheGetResponseWithDefaults

`func NewApplicationCacheGetResponseWithDefaults() *ApplicationCacheGetResponse`

NewApplicationCacheGetResponseWithDefaults instantiates a new ApplicationCacheGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ApplicationCacheGetResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ApplicationCacheGetResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ApplicationCacheGetResponse) SetCount(v int64)`

SetCount sets Count field to given value.


### GetTotalPages

`func (o *ApplicationCacheGetResponse) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ApplicationCacheGetResponse) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ApplicationCacheGetResponse) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.


### GetSchemaVersion

`func (o *ApplicationCacheGetResponse) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ApplicationCacheGetResponse) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ApplicationCacheGetResponse) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetLinks

`func (o *ApplicationCacheGetResponse) GetLinks() ApplicationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationCacheGetResponse) GetLinksOk() (*ApplicationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationCacheGetResponse) SetLinks(v ApplicationLinks)`

SetLinks sets Links field to given value.


### GetResults

`func (o *ApplicationCacheGetResponse) GetResults() []ApplicationCacheResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ApplicationCacheGetResponse) GetResultsOk() (*[]ApplicationCacheResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ApplicationCacheGetResponse) SetResults(v []ApplicationCacheResults)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


