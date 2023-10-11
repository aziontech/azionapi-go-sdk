# WAFEvents200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]WAFEvents200ResultsInner**](WAFEvents200ResultsInner.md) |  | [optional] 
**SchemaVersion** | Pointer to **int64** |  | [optional] 

## Methods

### NewWAFEvents200

`func NewWAFEvents200() *WAFEvents200`

NewWAFEvents200 instantiates a new WAFEvents200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFEvents200WithDefaults

`func NewWAFEvents200WithDefaults() *WAFEvents200`

NewWAFEvents200WithDefaults instantiates a new WAFEvents200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *WAFEvents200) GetResults() []WAFEvents200ResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *WAFEvents200) GetResultsOk() (*[]WAFEvents200ResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *WAFEvents200) SetResults(v []WAFEvents200ResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *WAFEvents200) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *WAFEvents200) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *WAFEvents200) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *WAFEvents200) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *WAFEvents200) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


