# RuleSetResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**RuleSetResultResults**](RuleSetResultResults.md) |  | [optional] 
**SchemaVersion** | Pointer to **int32** |  | [optional] 

## Methods

### NewRuleSetResult

`func NewRuleSetResult() *RuleSetResult`

NewRuleSetResult instantiates a new RuleSetResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleSetResultWithDefaults

`func NewRuleSetResultWithDefaults() *RuleSetResult`

NewRuleSetResultWithDefaults instantiates a new RuleSetResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *RuleSetResult) GetResults() RuleSetResultResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RuleSetResult) GetResultsOk() (*RuleSetResultResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RuleSetResult) SetResults(v RuleSetResultResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *RuleSetResult) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *RuleSetResult) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RuleSetResult) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RuleSetResult) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *RuleSetResult) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


