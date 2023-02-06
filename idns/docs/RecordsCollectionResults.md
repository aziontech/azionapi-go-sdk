# RecordsCollectionResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneId** | Pointer to **int32** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Records** | Pointer to [**[]Record**](Record.md) | Zone records collection | [optional] 

## Methods

### NewRecordsCollectionResults

`func NewRecordsCollectionResults() *RecordsCollectionResults`

NewRecordsCollectionResults instantiates a new RecordsCollectionResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordsCollectionResultsWithDefaults

`func NewRecordsCollectionResultsWithDefaults() *RecordsCollectionResults`

NewRecordsCollectionResultsWithDefaults instantiates a new RecordsCollectionResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneId

`func (o *RecordsCollectionResults) GetZoneId() int32`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *RecordsCollectionResults) GetZoneIdOk() (*int32, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *RecordsCollectionResults) SetZoneId(v int32)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *RecordsCollectionResults) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDomain

`func (o *RecordsCollectionResults) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RecordsCollectionResults) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RecordsCollectionResults) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *RecordsCollectionResults) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetRecords

`func (o *RecordsCollectionResults) GetRecords() []Record`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *RecordsCollectionResults) GetRecordsOk() (*[]Record, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *RecordsCollectionResults) SetRecords(v []Record)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *RecordsCollectionResults) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


