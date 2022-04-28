# ApplicationUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**ApplicationUpdateResults**](ApplicationUpdateResults.md) |  | 
**SchemaVersion** | **int64** |  | 

## Methods

### NewApplicationUpdateResponse

`func NewApplicationUpdateResponse(results ApplicationUpdateResults, schemaVersion int64, ) *ApplicationUpdateResponse`

NewApplicationUpdateResponse instantiates a new ApplicationUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationUpdateResponseWithDefaults

`func NewApplicationUpdateResponseWithDefaults() *ApplicationUpdateResponse`

NewApplicationUpdateResponseWithDefaults instantiates a new ApplicationUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ApplicationUpdateResponse) GetResults() ApplicationUpdateResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ApplicationUpdateResponse) GetResultsOk() (*ApplicationUpdateResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ApplicationUpdateResponse) SetResults(v ApplicationUpdateResults)`

SetResults sets Results field to given value.


### GetSchemaVersion

`func (o *ApplicationUpdateResponse) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ApplicationUpdateResponse) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ApplicationUpdateResponse) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


