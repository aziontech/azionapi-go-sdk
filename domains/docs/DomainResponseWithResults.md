# DomainResponseWithResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** |  | 
**TotalPages** | **int64** |  | 
**SchemaVersion** | **int64** |  | 
**Links** | [**DomainLinks**](DomainLinks.md) |  | 
**Results** | [**[]DomainEntityResponse**](DomainEntityResponse.md) |  | 

## Methods

### NewDomainResponseWithResults

`func NewDomainResponseWithResults(count int64, totalPages int64, schemaVersion int64, links DomainLinks, results []DomainEntityResponse, ) *DomainResponseWithResults`

NewDomainResponseWithResults instantiates a new DomainResponseWithResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainResponseWithResultsWithDefaults

`func NewDomainResponseWithResultsWithDefaults() *DomainResponseWithResults`

NewDomainResponseWithResultsWithDefaults instantiates a new DomainResponseWithResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DomainResponseWithResults) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DomainResponseWithResults) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DomainResponseWithResults) SetCount(v int64)`

SetCount sets Count field to given value.


### GetTotalPages

`func (o *DomainResponseWithResults) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *DomainResponseWithResults) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *DomainResponseWithResults) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.


### GetSchemaVersion

`func (o *DomainResponseWithResults) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *DomainResponseWithResults) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *DomainResponseWithResults) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetLinks

`func (o *DomainResponseWithResults) GetLinks() DomainLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DomainResponseWithResults) GetLinksOk() (*DomainLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DomainResponseWithResults) SetLinks(v DomainLinks)`

SetLinks sets Links field to given value.


### GetResults

`func (o *DomainResponseWithResults) GetResults() []DomainEntityResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DomainResponseWithResults) GetResultsOk() (*[]DomainEntityResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DomainResponseWithResults) SetResults(v []DomainEntityResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


