# RulesEngineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** |  | 
**TotalPages** | **int64** |  | 
**SchemaVersion** | **int64** |  | 
**Links** | [**OriginsResponseLinks**](OriginsResponseLinks.md) |  | 
**Results** | [**[]RulesEngineResultResponse**](RulesEngineResultResponse.md) |  | 

## Methods

### NewRulesEngineResponse

`func NewRulesEngineResponse(count int64, totalPages int64, schemaVersion int64, links OriginsResponseLinks, results []RulesEngineResultResponse, ) *RulesEngineResponse`

NewRulesEngineResponse instantiates a new RulesEngineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesEngineResponseWithDefaults

`func NewRulesEngineResponseWithDefaults() *RulesEngineResponse`

NewRulesEngineResponseWithDefaults instantiates a new RulesEngineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RulesEngineResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RulesEngineResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RulesEngineResponse) SetCount(v int64)`

SetCount sets Count field to given value.


### GetTotalPages

`func (o *RulesEngineResponse) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *RulesEngineResponse) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *RulesEngineResponse) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.


### GetSchemaVersion

`func (o *RulesEngineResponse) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RulesEngineResponse) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RulesEngineResponse) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetLinks

`func (o *RulesEngineResponse) GetLinks() OriginsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RulesEngineResponse) GetLinksOk() (*OriginsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RulesEngineResponse) SetLinks(v OriginsResponseLinks)`

SetLinks sets Links field to given value.


### GetResults

`func (o *RulesEngineResponse) GetResults() []RulesEngineResultResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RulesEngineResponse) GetResultsOk() (*[]RulesEngineResultResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RulesEngineResponse) SetResults(v []RulesEngineResultResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


