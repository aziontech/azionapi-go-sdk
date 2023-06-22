# ListEdgeFirewallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**TotalPages** | Pointer to **int64** |  | [optional] 
**SchemaVersion** | Pointer to **int64** |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Results** | Pointer to [**[]EdgeFirewall**](EdgeFirewall.md) |  | [optional] 

## Methods

### NewListEdgeFirewallResponse

`func NewListEdgeFirewallResponse() *ListEdgeFirewallResponse`

NewListEdgeFirewallResponse instantiates a new ListEdgeFirewallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEdgeFirewallResponseWithDefaults

`func NewListEdgeFirewallResponseWithDefaults() *ListEdgeFirewallResponse`

NewListEdgeFirewallResponseWithDefaults instantiates a new ListEdgeFirewallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListEdgeFirewallResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListEdgeFirewallResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListEdgeFirewallResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListEdgeFirewallResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *ListEdgeFirewallResponse) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ListEdgeFirewallResponse) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ListEdgeFirewallResponse) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *ListEdgeFirewallResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ListEdgeFirewallResponse) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ListEdgeFirewallResponse) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ListEdgeFirewallResponse) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ListEdgeFirewallResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetLinks

`func (o *ListEdgeFirewallResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListEdgeFirewallResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListEdgeFirewallResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListEdgeFirewallResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *ListEdgeFirewallResponse) GetResults() []EdgeFirewall`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListEdgeFirewallResponse) GetResultsOk() (*[]EdgeFirewall, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListEdgeFirewallResponse) SetResults(v []EdgeFirewall)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListEdgeFirewallResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


