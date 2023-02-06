# ZonesCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **int32** | The schema version | [optional] 
**Count** | Pointer to **int32** | Number of records | [optional] 
**TotalPages** | Pointer to **int32** | The total pages | [optional] 
**Links** | Pointer to [**ZonesCollectionLinks**](ZonesCollectionLinks.md) |  | [optional] 
**Results** | Pointer to [**[]Zone**](Zone.md) | Hosted zones collection | [optional] 

## Methods

### NewZonesCollection

`func NewZonesCollection() *ZonesCollection`

NewZonesCollection instantiates a new ZonesCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZonesCollectionWithDefaults

`func NewZonesCollectionWithDefaults() *ZonesCollection`

NewZonesCollectionWithDefaults instantiates a new ZonesCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *ZonesCollection) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ZonesCollection) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ZonesCollection) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ZonesCollection) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetCount

`func (o *ZonesCollection) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ZonesCollection) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ZonesCollection) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ZonesCollection) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *ZonesCollection) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ZonesCollection) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ZonesCollection) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *ZonesCollection) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetLinks

`func (o *ZonesCollection) GetLinks() ZonesCollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ZonesCollection) GetLinksOk() (*ZonesCollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ZonesCollection) SetLinks(v ZonesCollectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ZonesCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *ZonesCollection) GetResults() []Zone`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ZonesCollection) GetResultsOk() (*[]Zone, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ZonesCollection) SetResults(v []Zone)`

SetResults sets Results field to given value.

### HasResults

`func (o *ZonesCollection) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


