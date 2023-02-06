# GetZonesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **int32** | The schema version | [optional] 
**Count** | Pointer to **int32** | Number of records | [optional] 
**TotalPages** | Pointer to **int32** | The total pages | [optional] 
**Links** | Pointer to [**GetZonesResponseLinks**](GetZonesResponseLinks.md) |  | [optional] 
**Results** | Pointer to [**[]Zone**](Zone.md) | Hosted zones collection | [optional] 

## Methods

### NewGetZonesResponse

`func NewGetZonesResponse() *GetZonesResponse`

NewGetZonesResponse instantiates a new GetZonesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZonesResponseWithDefaults

`func NewGetZonesResponseWithDefaults() *GetZonesResponse`

NewGetZonesResponseWithDefaults instantiates a new GetZonesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *GetZonesResponse) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetZonesResponse) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetZonesResponse) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *GetZonesResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetCount

`func (o *GetZonesResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetZonesResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetZonesResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetZonesResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalPages

`func (o *GetZonesResponse) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *GetZonesResponse) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *GetZonesResponse) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *GetZonesResponse) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetLinks

`func (o *GetZonesResponse) GetLinks() GetZonesResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetZonesResponse) GetLinksOk() (*GetZonesResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetZonesResponse) SetLinks(v GetZonesResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetZonesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *GetZonesResponse) GetResults() []Zone`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetZonesResponse) GetResultsOk() (*[]Zone, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetZonesResponse) SetResults(v []Zone)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetZonesResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


