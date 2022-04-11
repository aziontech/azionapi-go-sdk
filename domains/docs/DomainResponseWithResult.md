# DomainResponseWithResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | **int64** |  | 
**Results** | [**DomainResults**](DomainResults.md) |  | 

## Methods

### NewDomainResponseWithResult

`func NewDomainResponseWithResult(schemaVersion int64, results DomainResults, ) *DomainResponseWithResult`

NewDomainResponseWithResult instantiates a new DomainResponseWithResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainResponseWithResultWithDefaults

`func NewDomainResponseWithResultWithDefaults() *DomainResponseWithResult`

NewDomainResponseWithResultWithDefaults instantiates a new DomainResponseWithResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetResults

`func (o *DomainResponseWithResult) GetResults() DomainResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DomainResponseWithResult) GetResultsOk() (*DomainResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DomainResponseWithResult) SetResults(v DomainResults)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


