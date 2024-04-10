# DomainResponseWithResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Links** | Pointer to [**DomainLinks**](DomainLinks.md) |  | [optional] 
**Results** | [**DomainEntity**](DomainEntity.md) |  | 
**TotalPages** | Pointer to **int64** |  | [optional] 
**SchemaVersion** | **int64** |  | 

## Methods

### NewDomainResponseWithResult

`func NewDomainResponseWithResult(results DomainEntity, schemaVersion int64, ) *DomainResponseWithResult`

NewDomainResponseWithResult instantiates a new DomainResponseWithResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainResponseWithResultWithDefaults

`func NewDomainResponseWithResultWithDefaults() *DomainResponseWithResult`

NewDomainResponseWithResultWithDefaults instantiates a new DomainResponseWithResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DomainResponseWithResult) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DomainResponseWithResult) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DomainResponseWithResult) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *DomainResponseWithResult) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLinks

`func (o *DomainResponseWithResult) GetLinks() DomainLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DomainResponseWithResult) GetLinksOk() (*DomainLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DomainResponseWithResult) SetLinks(v DomainLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DomainResponseWithResult) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *DomainResponseWithResult) GetResults() DomainEntity`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DomainResponseWithResult) GetResultsOk() (*DomainEntity, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DomainResponseWithResult) SetResults(v DomainEntity)`

SetResults sets Results field to given value.


### GetTotalPages

`func (o *DomainResponseWithResult) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *DomainResponseWithResult) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *DomainResponseWithResult) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *DomainResponseWithResult) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *DomainResponseWithResult) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *DomainResponseWithResult) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *DomainResponseWithResult) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


