# ApplicationInstanceResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **int64** |  | [optional] 
**Results** | Pointer to [**ApplicationInstancesResults**](ApplicationInstancesResults.md) |  | [optional] 

## Methods

### NewApplicationInstanceResults

`func NewApplicationInstanceResults() *ApplicationInstanceResults`

NewApplicationInstanceResults instantiates a new ApplicationInstanceResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstanceResultsWithDefaults

`func NewApplicationInstanceResultsWithDefaults() *ApplicationInstanceResults`

NewApplicationInstanceResultsWithDefaults instantiates a new ApplicationInstanceResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *ApplicationInstanceResults) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ApplicationInstanceResults) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ApplicationInstanceResults) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ApplicationInstanceResults) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetResults

`func (o *ApplicationInstanceResults) GetResults() ApplicationInstancesResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ApplicationInstanceResults) GetResultsOk() (*ApplicationInstancesResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ApplicationInstanceResults) SetResults(v ApplicationInstancesResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *ApplicationInstanceResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


