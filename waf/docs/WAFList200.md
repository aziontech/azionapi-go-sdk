# WAFList200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**TotalPages** | Pointer to **int64** |  | [optional] 
**Links** | Pointer to [**SingleWAF**](SingleWAF.md) |  | [optional] 
**SchemaVersion** | Pointer to **int64** |  | [optional] 

## Methods

### NewWAFList200

`func NewWAFList200() *WAFList200`

NewWAFList200 instantiates a new WAFList200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWAFList200WithDefaults

`func NewWAFList200WithDefaults() *WAFList200`

NewWAFList200WithDefaults instantiates a new WAFList200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *WAFList200) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *WAFList200) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *WAFList200) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *WAFList200) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *WAFList200) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *WAFList200) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *WAFList200) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *WAFList200) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetLinks

`func (o *WAFList200) GetLinks() SingleWAF`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WAFList200) GetLinksOk() (*SingleWAF, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WAFList200) SetLinks(v SingleWAF)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WAFList200) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *WAFList200) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *WAFList200) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *WAFList200) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *WAFList200) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


