# CreateApplicationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**ApplicationResultsCreate**](ApplicationResultsCreate.md) |  | 
**SchemaVersion** | **int64** |  | 

## Methods

### NewCreateApplicationResult

`func NewCreateApplicationResult(results ApplicationResultsCreate, schemaVersion int64, ) *CreateApplicationResult`

NewCreateApplicationResult instantiates a new CreateApplicationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationResultWithDefaults

`func NewCreateApplicationResultWithDefaults() *CreateApplicationResult`

NewCreateApplicationResultWithDefaults instantiates a new CreateApplicationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CreateApplicationResult) GetResults() ApplicationResultsCreate`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CreateApplicationResult) GetResultsOk() (*ApplicationResultsCreate, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CreateApplicationResult) SetResults(v ApplicationResultsCreate)`

SetResults sets Results field to given value.


### GetSchemaVersion

`func (o *CreateApplicationResult) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CreateApplicationResult) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CreateApplicationResult) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


