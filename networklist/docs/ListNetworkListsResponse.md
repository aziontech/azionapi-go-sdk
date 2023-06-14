# ListNetworkListsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**TotalPages** | Pointer to **int64** |  | [optional] 
**SchemaVersion** | Pointer to **int64** |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Results** | Pointer to [**[]NetworkLists**](NetworkLists.md) |  | [optional] 

## Methods

### NewListNetworkListsResponse

`func NewListNetworkListsResponse() *ListNetworkListsResponse`

NewListNetworkListsResponse instantiates a new ListNetworkListsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkListsResponseWithDefaults

`func NewListNetworkListsResponseWithDefaults() *ListNetworkListsResponse`

NewListNetworkListsResponseWithDefaults instantiates a new ListNetworkListsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListNetworkListsResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListNetworkListsResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListNetworkListsResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListNetworkListsResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *ListNetworkListsResponse) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ListNetworkListsResponse) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ListNetworkListsResponse) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *ListNetworkListsResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ListNetworkListsResponse) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ListNetworkListsResponse) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ListNetworkListsResponse) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ListNetworkListsResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetLinks

`func (o *ListNetworkListsResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListNetworkListsResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListNetworkListsResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListNetworkListsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *ListNetworkListsResponse) GetResults() []NetworkLists`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListNetworkListsResponse) GetResultsOk() (*[]NetworkLists, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListNetworkListsResponse) SetResults(v []NetworkLists)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListNetworkListsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


