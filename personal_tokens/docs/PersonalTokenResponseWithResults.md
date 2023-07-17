# PersonalTokenResponseWithResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**SchemaVersion** | Pointer to **int64** |  | [optional] 
**Links** | Pointer to [**PersonalTokenResponseWithResultsLinks**](PersonalTokenResponseWithResultsLinks.md) |  | [optional] 
**Results** | Pointer to [**[]PersonalTokenResponseGet**](PersonalTokenResponseGet.md) |  | [optional] 

## Methods

### NewPersonalTokenResponseWithResults

`func NewPersonalTokenResponseWithResults() *PersonalTokenResponseWithResults`

NewPersonalTokenResponseWithResults instantiates a new PersonalTokenResponseWithResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalTokenResponseWithResultsWithDefaults

`func NewPersonalTokenResponseWithResultsWithDefaults() *PersonalTokenResponseWithResults`

NewPersonalTokenResponseWithResultsWithDefaults instantiates a new PersonalTokenResponseWithResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PersonalTokenResponseWithResults) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PersonalTokenResponseWithResults) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PersonalTokenResponseWithResults) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *PersonalTokenResponseWithResults) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *PersonalTokenResponseWithResults) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *PersonalTokenResponseWithResults) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *PersonalTokenResponseWithResults) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *PersonalTokenResponseWithResults) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *PersonalTokenResponseWithResults) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *PersonalTokenResponseWithResults) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *PersonalTokenResponseWithResults) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *PersonalTokenResponseWithResults) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetLinks

`func (o *PersonalTokenResponseWithResults) GetLinks() PersonalTokenResponseWithResultsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PersonalTokenResponseWithResults) GetLinksOk() (*PersonalTokenResponseWithResultsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PersonalTokenResponseWithResults) SetLinks(v PersonalTokenResponseWithResultsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PersonalTokenResponseWithResults) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *PersonalTokenResponseWithResults) GetResults() []PersonalTokenResponseGet`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PersonalTokenResponseWithResults) GetResultsOk() (*[]PersonalTokenResponseGet, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PersonalTokenResponseWithResults) SetResults(v []PersonalTokenResponseGet)`

SetResults sets Results field to given value.

### HasResults

`func (o *PersonalTokenResponseWithResults) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


