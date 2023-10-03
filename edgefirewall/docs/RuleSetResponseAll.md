# RuleSetResponseAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**TotalPages** | Pointer to **int64** |  | [optional] 
**SchemaVersion** | Pointer to **int32** |  | [optional] [default to 3]
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Results** | Pointer to [**[]RuleSetResultAll**](RuleSetResultAll.md) |  | [optional] 

## Methods

### NewRuleSetResponseAll

`func NewRuleSetResponseAll() *RuleSetResponseAll`

NewRuleSetResponseAll instantiates a new RuleSetResponseAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleSetResponseAllWithDefaults

`func NewRuleSetResponseAllWithDefaults() *RuleSetResponseAll`

NewRuleSetResponseAllWithDefaults instantiates a new RuleSetResponseAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RuleSetResponseAll) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RuleSetResponseAll) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RuleSetResponseAll) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *RuleSetResponseAll) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *RuleSetResponseAll) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *RuleSetResponseAll) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *RuleSetResponseAll) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *RuleSetResponseAll) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *RuleSetResponseAll) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RuleSetResponseAll) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RuleSetResponseAll) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *RuleSetResponseAll) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetLinks

`func (o *RuleSetResponseAll) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RuleSetResponseAll) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RuleSetResponseAll) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RuleSetResponseAll) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *RuleSetResponseAll) GetResults() []RuleSetResultAll`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RuleSetResponseAll) GetResultsOk() (*[]RuleSetResultAll, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RuleSetResponseAll) SetResults(v []RuleSetResultAll)`

SetResults sets Results field to given value.

### HasResults

`func (o *RuleSetResponseAll) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


