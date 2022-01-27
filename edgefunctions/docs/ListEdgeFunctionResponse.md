# ListEdgeFunctionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**TotalPages** | Pointer to **int64** |  | [optional] 
**SchemaVersion** | Pointer to **int64** |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Results** | Pointer to [**[]Results**](Results.md) |  | [optional] 

## Methods

### NewListEdgeFunctionResponse

`func NewListEdgeFunctionResponse() *ListEdgeFunctionResponse`

NewListEdgeFunctionResponse instantiates a new ListEdgeFunctionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEdgeFunctionResponseWithDefaults

`func NewListEdgeFunctionResponseWithDefaults() *ListEdgeFunctionResponse`

NewListEdgeFunctionResponseWithDefaults instantiates a new ListEdgeFunctionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListEdgeFunctionResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListEdgeFunctionResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListEdgeFunctionResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListEdgeFunctionResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *ListEdgeFunctionResponse) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ListEdgeFunctionResponse) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ListEdgeFunctionResponse) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *ListEdgeFunctionResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ListEdgeFunctionResponse) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ListEdgeFunctionResponse) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ListEdgeFunctionResponse) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ListEdgeFunctionResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetLinks

`func (o *ListEdgeFunctionResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListEdgeFunctionResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListEdgeFunctionResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListEdgeFunctionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *ListEdgeFunctionResponse) GetResults() []Results`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListEdgeFunctionResponse) GetResultsOk() (*[]Results, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListEdgeFunctionResponse) SetResults(v []Results)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListEdgeFunctionResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


