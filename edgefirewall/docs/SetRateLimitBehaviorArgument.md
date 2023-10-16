# SetRateLimitBehaviorArgument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**LimitBy** | Pointer to **string** |  | [optional] 
**AverageRateLimit** | Pointer to [**SetRateLimitBehaviorArgumentAverageRateLimit**](SetRateLimitBehaviorArgumentAverageRateLimit.md) |  | [optional] 
**MaximumBurstSize** | Pointer to [**SetRateLimitBehaviorArgumentAverageRateLimit**](SetRateLimitBehaviorArgumentAverageRateLimit.md) |  | [optional] 

## Methods

### NewSetRateLimitBehaviorArgument

`func NewSetRateLimitBehaviorArgument() *SetRateLimitBehaviorArgument`

NewSetRateLimitBehaviorArgument instantiates a new SetRateLimitBehaviorArgument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRateLimitBehaviorArgumentWithDefaults

`func NewSetRateLimitBehaviorArgumentWithDefaults() *SetRateLimitBehaviorArgument`

NewSetRateLimitBehaviorArgumentWithDefaults instantiates a new SetRateLimitBehaviorArgument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SetRateLimitBehaviorArgument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetRateLimitBehaviorArgument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetRateLimitBehaviorArgument) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SetRateLimitBehaviorArgument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLimitBy

`func (o *SetRateLimitBehaviorArgument) GetLimitBy() string`

GetLimitBy returns the LimitBy field if non-nil, zero value otherwise.

### GetLimitByOk

`func (o *SetRateLimitBehaviorArgument) GetLimitByOk() (*string, bool)`

GetLimitByOk returns a tuple with the LimitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitBy

`func (o *SetRateLimitBehaviorArgument) SetLimitBy(v string)`

SetLimitBy sets LimitBy field to given value.

### HasLimitBy

`func (o *SetRateLimitBehaviorArgument) HasLimitBy() bool`

HasLimitBy returns a boolean if a field has been set.

### GetAverageRateLimit

`func (o *SetRateLimitBehaviorArgument) GetAverageRateLimit() SetRateLimitBehaviorArgumentAverageRateLimit`

GetAverageRateLimit returns the AverageRateLimit field if non-nil, zero value otherwise.

### GetAverageRateLimitOk

`func (o *SetRateLimitBehaviorArgument) GetAverageRateLimitOk() (*SetRateLimitBehaviorArgumentAverageRateLimit, bool)`

GetAverageRateLimitOk returns a tuple with the AverageRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRateLimit

`func (o *SetRateLimitBehaviorArgument) SetAverageRateLimit(v SetRateLimitBehaviorArgumentAverageRateLimit)`

SetAverageRateLimit sets AverageRateLimit field to given value.

### HasAverageRateLimit

`func (o *SetRateLimitBehaviorArgument) HasAverageRateLimit() bool`

HasAverageRateLimit returns a boolean if a field has been set.

### GetMaximumBurstSize

`func (o *SetRateLimitBehaviorArgument) GetMaximumBurstSize() SetRateLimitBehaviorArgumentAverageRateLimit`

GetMaximumBurstSize returns the MaximumBurstSize field if non-nil, zero value otherwise.

### GetMaximumBurstSizeOk

`func (o *SetRateLimitBehaviorArgument) GetMaximumBurstSizeOk() (*SetRateLimitBehaviorArgumentAverageRateLimit, bool)`

GetMaximumBurstSizeOk returns a tuple with the MaximumBurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumBurstSize

`func (o *SetRateLimitBehaviorArgument) SetMaximumBurstSize(v SetRateLimitBehaviorArgumentAverageRateLimit)`

SetMaximumBurstSize sets MaximumBurstSize field to given value.

### HasMaximumBurstSize

`func (o *SetRateLimitBehaviorArgument) HasMaximumBurstSize() bool`

HasMaximumBurstSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


