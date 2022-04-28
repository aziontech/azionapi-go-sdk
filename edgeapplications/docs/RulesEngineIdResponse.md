# RulesEngineIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**RulesEngineResultResponse**](RulesEngineResultResponse.md) |  | 
**SchemaVersion** | **int64** |  | 

## Methods

### NewRulesEngineIdResponse

`func NewRulesEngineIdResponse(results RulesEngineResultResponse, schemaVersion int64, ) *RulesEngineIdResponse`

NewRulesEngineIdResponse instantiates a new RulesEngineIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesEngineIdResponseWithDefaults

`func NewRulesEngineIdResponseWithDefaults() *RulesEngineIdResponse`

NewRulesEngineIdResponseWithDefaults instantiates a new RulesEngineIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *RulesEngineIdResponse) GetResults() RulesEngineResultResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RulesEngineIdResponse) GetResultsOk() (*RulesEngineResultResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RulesEngineIdResponse) SetResults(v RulesEngineResultResponse)`

SetResults sets Results field to given value.


### GetSchemaVersion

`func (o *RulesEngineIdResponse) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RulesEngineIdResponse) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RulesEngineIdResponse) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


