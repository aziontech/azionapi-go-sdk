# PostOrPutRecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **int32** | The schema version | [optional] 
**Results** | Pointer to [**Record**](Record.md) |  | [optional] 

## Methods

### NewPostOrPutRecordResponse

`func NewPostOrPutRecordResponse() *PostOrPutRecordResponse`

NewPostOrPutRecordResponse instantiates a new PostOrPutRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostOrPutRecordResponseWithDefaults

`func NewPostOrPutRecordResponseWithDefaults() *PostOrPutRecordResponse`

NewPostOrPutRecordResponseWithDefaults instantiates a new PostOrPutRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *PostOrPutRecordResponse) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *PostOrPutRecordResponse) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *PostOrPutRecordResponse) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *PostOrPutRecordResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetResults

`func (o *PostOrPutRecordResponse) GetResults() Record`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PostOrPutRecordResponse) GetResultsOk() (*Record, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PostOrPutRecordResponse) SetResults(v Record)`

SetResults sets Results field to given value.

### HasResults

`func (o *PostOrPutRecordResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

