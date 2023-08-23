# CreateCustomDataStreamingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]PostCustomDataStreamingResponse**](PostCustomDataStreamingResponse.md) |  | [optional] 
**SchemaVersion** | Pointer to **float32** |  | [optional] 

## Methods

### NewCreateCustomDataStreamingResponse

`func NewCreateCustomDataStreamingResponse() *CreateCustomDataStreamingResponse`

NewCreateCustomDataStreamingResponse instantiates a new CreateCustomDataStreamingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomDataStreamingResponseWithDefaults

`func NewCreateCustomDataStreamingResponseWithDefaults() *CreateCustomDataStreamingResponse`

NewCreateCustomDataStreamingResponseWithDefaults instantiates a new CreateCustomDataStreamingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CreateCustomDataStreamingResponse) GetResults() []PostCustomDataStreamingResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CreateCustomDataStreamingResponse) GetResultsOk() (*[]PostCustomDataStreamingResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CreateCustomDataStreamingResponse) SetResults(v []PostCustomDataStreamingResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *CreateCustomDataStreamingResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *CreateCustomDataStreamingResponse) GetSchemaVersion() float32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CreateCustomDataStreamingResponse) GetSchemaVersionOk() (*float32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CreateCustomDataStreamingResponse) SetSchemaVersion(v float32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CreateCustomDataStreamingResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


