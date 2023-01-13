# GetRecordsResponseResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneId** | Pointer to **int32** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Records** | Pointer to [**[]Record**](Record.md) | Zone records collection | [optional] 

## Methods

### NewGetRecordsResponseResults

`func NewGetRecordsResponseResults() *GetRecordsResponseResults`

NewGetRecordsResponseResults instantiates a new GetRecordsResponseResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordsResponseResultsWithDefaults

`func NewGetRecordsResponseResultsWithDefaults() *GetRecordsResponseResults`

NewGetRecordsResponseResultsWithDefaults instantiates a new GetRecordsResponseResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneId

`func (o *GetRecordsResponseResults) GetZoneId() int32`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetRecordsResponseResults) GetZoneIdOk() (*int32, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetRecordsResponseResults) SetZoneId(v int32)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetRecordsResponseResults) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDomain

`func (o *GetRecordsResponseResults) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetRecordsResponseResults) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetRecordsResponseResults) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetRecordsResponseResults) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetRecords

`func (o *GetRecordsResponseResults) GetRecords() []Record`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *GetRecordsResponseResults) GetRecordsOk() (*[]Record, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *GetRecordsResponseResults) SetRecords(v []Record)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *GetRecordsResponseResults) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


