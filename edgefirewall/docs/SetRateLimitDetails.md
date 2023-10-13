# SetRateLimitDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**LimitBy** | Pointer to **string** |  | [optional] 
**AverageRateLimit** | Pointer to **int32** |  | [optional] 
**MaximumBurstSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewSetRateLimitDetails

`func NewSetRateLimitDetails() *SetRateLimitDetails`

NewSetRateLimitDetails instantiates a new SetRateLimitDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRateLimitDetailsWithDefaults

`func NewSetRateLimitDetailsWithDefaults() *SetRateLimitDetails`

NewSetRateLimitDetailsWithDefaults instantiates a new SetRateLimitDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SetRateLimitDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetRateLimitDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetRateLimitDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SetRateLimitDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLimitBy

`func (o *SetRateLimitDetails) GetLimitBy() string`

GetLimitBy returns the LimitBy field if non-nil, zero value otherwise.

### GetLimitByOk

`func (o *SetRateLimitDetails) GetLimitByOk() (*string, bool)`

GetLimitByOk returns a tuple with the LimitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBy

`func (o *SetRateLimitDetails) SetLimitBy(v string)`

SetLimitBy sets LimitBy field to given value.

### HasLimitBy

`func (o *SetRateLimitDetails) HasLimitBy() bool`

HasLimitBy returns a boolean if a field has been set.

### GetAverageRateLimit

`func (o *SetRateLimitDetails) GetAverageRateLimit() int32`

GetAverageRateLimit returns the AverageRateLimit field if non-nil, zero value otherwise.

### GetAverageRateLimitOk

`func (o *SetRateLimitDetails) GetAverageRateLimitOk() (*int32, bool)`

GetAverageRateLimitOk returns a tuple with the AverageRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRateLimit

`func (o *SetRateLimitDetails) SetAverageRateLimit(v int32)`

SetAverageRateLimit sets AverageRateLimit field to given value.

### HasAverageRateLimit

`func (o *SetRateLimitDetails) HasAverageRateLimit() bool`

HasAverageRateLimit returns a boolean if a field has been set.

### GetMaximumBurstSize

`func (o *SetRateLimitDetails) GetMaximumBurstSize() int32`

GetMaximumBurstSize returns the MaximumBurstSize field if non-nil, zero value otherwise.

### GetMaximumBurstSizeOk

`func (o *SetRateLimitDetails) GetMaximumBurstSizeOk() (*int32, bool)`

GetMaximumBurstSizeOk returns a tuple with the MaximumBurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBurstSize

`func (o *SetRateLimitDetails) SetMaximumBurstSize(v int32)`

SetMaximumBurstSize sets MaximumBurstSize field to given value.

### HasMaximumBurstSize

`func (o *SetRateLimitDetails) HasMaximumBurstSize() bool`

HasMaximumBurstSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


