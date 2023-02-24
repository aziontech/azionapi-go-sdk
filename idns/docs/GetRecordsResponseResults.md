# GetRecordsResponseResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneId** | Pointer to **int32** |  | [optional] 
**ZoneDomain** | Pointer to **string** |  | [optional] 
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

### GetZoneDomain

`func (o *GetRecordsResponseResults) GetZoneDomain() string`

GetZoneDomain returns the ZoneDomain field if non-nil, zero value otherwise.

### GetZoneDomainOk

`func (o *GetRecordsResponseResults) GetZoneDomainOk() (*string, bool)`

GetZoneDomainOk returns a tuple with the ZoneDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneDomain

`func (o *GetRecordsResponseResults) SetZoneDomain(v string)`

SetZoneDomain sets ZoneDomain field to given value.

### HasZoneDomain

`func (o *GetRecordsResponseResults) HasZoneDomain() bool`

HasZoneDomain returns a boolean if a field has been set.

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


