# RecordsCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **int32** | The schema version | [optional] 
**Count** | Pointer to **int32** | Number of records | [optional] 
**TotalPages** | Pointer to **int32** | The total pages | [optional] 
**Links** | Pointer to [**ZonesCollectionLinks**](ZonesCollectionLinks.md) |  | [optional] 
**Results** | Pointer to [**RecordsCollectionResults**](RecordsCollectionResults.md) |  | [optional] 

## Methods

### NewRecordsCollection

`func NewRecordsCollection() *RecordsCollection`

NewRecordsCollection instantiates a new RecordsCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordsCollectionWithDefaults

`func NewRecordsCollectionWithDefaults() *RecordsCollection`

NewRecordsCollectionWithDefaults instantiates a new RecordsCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *RecordsCollection) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RecordsCollection) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RecordsCollection) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *RecordsCollection) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetCount

`func (o *RecordsCollection) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RecordsCollection) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RecordsCollection) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *RecordsCollection) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *RecordsCollection) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *RecordsCollection) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *RecordsCollection) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *RecordsCollection) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetLinks

`func (o *RecordsCollection) GetLinks() ZonesCollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RecordsCollection) GetLinksOk() (*ZonesCollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RecordsCollection) SetLinks(v ZonesCollectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RecordsCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *RecordsCollection) GetResults() RecordsCollectionResults`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RecordsCollection) GetResultsOk() (*RecordsCollectionResults, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RecordsCollection) SetResults(v RecordsCollectionResults)`

SetResults sets Results field to given value.

### HasResults

`func (o *RecordsCollection) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


