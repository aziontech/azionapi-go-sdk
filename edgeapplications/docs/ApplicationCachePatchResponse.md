# ApplicationCachePatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**ApplicationCacheResponseDetails**](ApplicationCacheResponseDetails.md) |  | [optional] 
**SchemaVersion** | Pointer to **int64** |  | [optional] 

## Methods

### NewApplicationCachePatchResponse

`func NewApplicationCachePatchResponse() *ApplicationCachePatchResponse`

NewApplicationCachePatchResponse instantiates a new ApplicationCachePatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCachePatchResponseWithDefaults

`func NewApplicationCachePatchResponseWithDefaults() *ApplicationCachePatchResponse`

NewApplicationCachePatchResponseWithDefaults instantiates a new ApplicationCachePatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ApplicationCachePatchResponse) GetResults() ApplicationCacheResponseDetails`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ApplicationCachePatchResponse) GetResultsOk() (*ApplicationCacheResponseDetails, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ApplicationCachePatchResponse) SetResults(v ApplicationCacheResponseDetails)`

SetResults sets Results field to given value.

### HasResults

`func (o *ApplicationCachePatchResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ApplicationCachePatchResponse) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ApplicationCachePatchResponse) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ApplicationCachePatchResponse) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ApplicationCachePatchResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


