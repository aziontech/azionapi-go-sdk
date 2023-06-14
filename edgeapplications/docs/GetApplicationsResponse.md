# GetApplicationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** |  | 
**TotalPages** | **int64** |  | 
**SchemaVersion** | **int64** |  | 
**Links** | [**ApplicationLinks**](ApplicationLinks.md) |  | 
**Results** | [**[]ApplicationsResults**](ApplicationsResults.md) |  | 

## Methods

### NewGetApplicationsResponse

`func NewGetApplicationsResponse(count int64, totalPages int64, schemaVersion int64, links ApplicationLinks, results []ApplicationsResults, ) *GetApplicationsResponse`

NewGetApplicationsResponse instantiates a new GetApplicationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApplicationsResponseWithDefaults

`func NewGetApplicationsResponseWithDefaults() *GetApplicationsResponse`

NewGetApplicationsResponseWithDefaults instantiates a new GetApplicationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GetApplicationsResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetApplicationsResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetApplicationsResponse) SetCount(v int64)`

SetCount sets Count field to given value.


### GetTotalPages

`func (o *GetApplicationsResponse) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *GetApplicationsResponse) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *GetApplicationsResponse) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.


### GetSchemaVersion

`func (o *GetApplicationsResponse) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetApplicationsResponse) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetApplicationsResponse) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetLinks

`func (o *GetApplicationsResponse) GetLinks() ApplicationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetApplicationsResponse) GetLinksOk() (*ApplicationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetApplicationsResponse) SetLinks(v ApplicationLinks)`

SetLinks sets Links field to given value.


### GetResults

`func (o *GetApplicationsResponse) GetResults() []ApplicationsResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetApplicationsResponse) GetResultsOk() (*[]ApplicationsResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetApplicationsResponse) SetResults(v []ApplicationsResults)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


