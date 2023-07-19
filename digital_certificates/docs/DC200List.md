# DC200List

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**SchemaVersion** | Pointer to **int32** |  | [optional] 
**Links** | Pointer to [**DC200ListLinks**](DC200ListLinks.md) |  | [optional] 
**Results** | Pointer to [**[][]ResultsInner**]([]ResultsInner.md) |  | [optional] 

## Methods

### NewDC200List

`func NewDC200List() *DC200List`

NewDC200List instantiates a new DC200List object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDC200ListWithDefaults

`func NewDC200ListWithDefaults() *DC200List`

NewDC200ListWithDefaults instantiates a new DC200List object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DC200List) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DC200List) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DC200List) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DC200List) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *DC200List) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *DC200List) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *DC200List) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *DC200List) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *DC200List) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *DC200List) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *DC200List) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *DC200List) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetLinks

`func (o *DC200List) GetLinks() DC200ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DC200List) GetLinksOk() (*DC200ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DC200List) SetLinks(v DC200ListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DC200List) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *DC200List) GetResults() [][]ResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DC200List) GetResultsOk() (*[][]ResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DC200List) SetResults(v [][]ResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *DC200List) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


