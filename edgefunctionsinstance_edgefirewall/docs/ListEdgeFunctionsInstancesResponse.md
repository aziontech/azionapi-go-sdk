# ListEdgeFunctionsInstancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**TotalPages** | Pointer to **int64** |  | [optional] 
**SchemaVersion** | Pointer to **int64** |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Results** | Pointer to [**[]EdgeFunctionsInstance**](EdgeFunctionsInstance.md) |  | [optional] 

## Methods

### NewListEdgeFunctionsInstancesResponse

`func NewListEdgeFunctionsInstancesResponse() *ListEdgeFunctionsInstancesResponse`

NewListEdgeFunctionsInstancesResponse instantiates a new ListEdgeFunctionsInstancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEdgeFunctionsInstancesResponseWithDefaults

`func NewListEdgeFunctionsInstancesResponseWithDefaults() *ListEdgeFunctionsInstancesResponse`

NewListEdgeFunctionsInstancesResponseWithDefaults instantiates a new ListEdgeFunctionsInstancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListEdgeFunctionsInstancesResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListEdgeFunctionsInstancesResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListEdgeFunctionsInstancesResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListEdgeFunctionsInstancesResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *ListEdgeFunctionsInstancesResponse) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ListEdgeFunctionsInstancesResponse) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ListEdgeFunctionsInstancesResponse) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *ListEdgeFunctionsInstancesResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ListEdgeFunctionsInstancesResponse) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ListEdgeFunctionsInstancesResponse) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ListEdgeFunctionsInstancesResponse) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ListEdgeFunctionsInstancesResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetLinks

`func (o *ListEdgeFunctionsInstancesResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListEdgeFunctionsInstancesResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListEdgeFunctionsInstancesResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListEdgeFunctionsInstancesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *ListEdgeFunctionsInstancesResponse) GetResults() []EdgeFunctionsInstance`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListEdgeFunctionsInstancesResponse) GetResultsOk() (*[]EdgeFunctionsInstance, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListEdgeFunctionsInstancesResponse) SetResults(v []EdgeFunctionsInstance)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListEdgeFunctionsInstancesResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


