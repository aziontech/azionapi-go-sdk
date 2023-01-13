# CreateZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Zone**](Zone.md) | The created hosted zone | [optional] 
**SchemaVersion** | Pointer to **int32** | The schema version | [optional] 

## Methods

### NewCreateZone

`func NewCreateZone() *CreateZone`

NewCreateZone instantiates a new CreateZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateZoneWithDefaults

`func NewCreateZoneWithDefaults() *CreateZone`

NewCreateZoneWithDefaults instantiates a new CreateZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CreateZone) GetResults() []Zone`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CreateZone) GetResultsOk() (*[]Zone, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CreateZone) SetResults(v []Zone)`

SetResults sets Results field to given value.

### HasResults

`func (o *CreateZone) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *CreateZone) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CreateZone) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CreateZone) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CreateZone) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


