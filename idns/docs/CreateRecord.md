# CreateRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**Record**](Record.md) |  | [optional] 
**SchemaVersion** | Pointer to **int32** | The schema version | [optional] 

## Methods

### NewCreateRecord

`func NewCreateRecord() *CreateRecord`

NewCreateRecord instantiates a new CreateRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRecordWithDefaults

`func NewCreateRecordWithDefaults() *CreateRecord`

NewCreateRecordWithDefaults instantiates a new CreateRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CreateRecord) GetResults() Record`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CreateRecord) GetResultsOk() (*Record, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CreateRecord) SetResults(v Record)`

SetResults sets Results field to given value.

### HasResults

`func (o *CreateRecord) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *CreateRecord) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CreateRecord) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CreateRecord) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CreateRecord) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


