# OriginsResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginId** | Pointer to **int64** |  | [optional] 
**OriginKey** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**OriginType** | Pointer to **string** |  | [optional] 
**Addresses** | Pointer to [**[]OriginsResultResponseAddresses**](OriginsResultResponseAddresses.md) |  | [optional] 
**OriginProtocolPolicy** | Pointer to **string** |  | [optional] 
**IsOriginRedirectionEnabled** | Pointer to **bool** |  | [optional] 
**HostHeader** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**OriginPath** | Pointer to **string** |  | [optional] 
**ConnectionTimeout** | Pointer to **int64** |  | [optional] 
**TimeoutBetweenBytes** | Pointer to **int64** |  | [optional] 
**HmacAuthentication** | Pointer to **bool** |  | [optional] 
**HmacRegionName** | Pointer to **string** |  | [optional] 
**HmacAccessKey** | Pointer to **string** |  | [optional] 
**HmacSecretKey** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 

## Methods

### NewOriginsResultResponse

`func NewOriginsResultResponse(name string, ) *OriginsResultResponse`

NewOriginsResultResponse instantiates a new OriginsResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginsResultResponseWithDefaults

`func NewOriginsResultResponseWithDefaults() *OriginsResultResponse`

NewOriginsResultResponseWithDefaults instantiates a new OriginsResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginId

`func (o *OriginsResultResponse) GetOriginId() int64`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *OriginsResultResponse) GetOriginIdOk() (*int64, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *OriginsResultResponse) SetOriginId(v int64)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *OriginsResultResponse) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.

### GetOriginKey

`func (o *OriginsResultResponse) GetOriginKey() string`

GetOriginKey returns the OriginKey field if non-nil, zero value otherwise.

### GetOriginKeyOk

`func (o *OriginsResultResponse) GetOriginKeyOk() (*string, bool)`

GetOriginKeyOk returns a tuple with the OriginKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginKey

`func (o *OriginsResultResponse) SetOriginKey(v string)`

SetOriginKey sets OriginKey field to given value.

### HasOriginKey

`func (o *OriginsResultResponse) HasOriginKey() bool`

HasOriginKey returns a boolean if a field has been set.

### GetName

`func (o *OriginsResultResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OriginsResultResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OriginsResultResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOriginType

`func (o *OriginsResultResponse) GetOriginType() string`

GetOriginType returns the OriginType field if non-nil, zero value otherwise.

### GetOriginTypeOk

`func (o *OriginsResultResponse) GetOriginTypeOk() (*string, bool)`

GetOriginTypeOk returns a tuple with the OriginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginType

`func (o *OriginsResultResponse) SetOriginType(v string)`

SetOriginType sets OriginType field to given value.

### HasOriginType

`func (o *OriginsResultResponse) HasOriginType() bool`

HasOriginType returns a boolean if a field has been set.

### GetAddresses

`func (o *OriginsResultResponse) GetAddresses() []OriginsResultResponseAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *OriginsResultResponse) GetAddressesOk() (*[]OriginsResultResponseAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *OriginsResultResponse) SetAddresses(v []OriginsResultResponseAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *OriginsResultResponse) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetOriginProtocolPolicy

`func (o *OriginsResultResponse) GetOriginProtocolPolicy() string`

GetOriginProtocolPolicy returns the OriginProtocolPolicy field if non-nil, zero value otherwise.

### GetOriginProtocolPolicyOk

`func (o *OriginsResultResponse) GetOriginProtocolPolicyOk() (*string, bool)`

GetOriginProtocolPolicyOk returns a tuple with the OriginProtocolPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProtocolPolicy

`func (o *OriginsResultResponse) SetOriginProtocolPolicy(v string)`

SetOriginProtocolPolicy sets OriginProtocolPolicy field to given value.

### HasOriginProtocolPolicy

`func (o *OriginsResultResponse) HasOriginProtocolPolicy() bool`

HasOriginProtocolPolicy returns a boolean if a field has been set.

### GetIsOriginRedirectionEnabled

`func (o *OriginsResultResponse) GetIsOriginRedirectionEnabled() bool`

GetIsOriginRedirectionEnabled returns the IsOriginRedirectionEnabled field if non-nil, zero value otherwise.

### GetIsOriginRedirectionEnabledOk

`func (o *OriginsResultResponse) GetIsOriginRedirectionEnabledOk() (*bool, bool)`

GetIsOriginRedirectionEnabledOk returns a tuple with the IsOriginRedirectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOriginRedirectionEnabled

`func (o *OriginsResultResponse) SetIsOriginRedirectionEnabled(v bool)`

SetIsOriginRedirectionEnabled sets IsOriginRedirectionEnabled field to given value.

### HasIsOriginRedirectionEnabled

`func (o *OriginsResultResponse) HasIsOriginRedirectionEnabled() bool`

HasIsOriginRedirectionEnabled returns a boolean if a field has been set.

### GetHostHeader

`func (o *OriginsResultResponse) GetHostHeader() string`

GetHostHeader returns the HostHeader field if non-nil, zero value otherwise.

### GetHostHeaderOk

`func (o *OriginsResultResponse) GetHostHeaderOk() (*string, bool)`

GetHostHeaderOk returns a tuple with the HostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostHeader

`func (o *OriginsResultResponse) SetHostHeader(v string)`

SetHostHeader sets HostHeader field to given value.

### HasHostHeader

`func (o *OriginsResultResponse) HasHostHeader() bool`

HasHostHeader returns a boolean if a field has been set.

### GetMethod

`func (o *OriginsResultResponse) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *OriginsResultResponse) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *OriginsResultResponse) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *OriginsResultResponse) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetOriginPath

`func (o *OriginsResultResponse) GetOriginPath() string`

GetOriginPath returns the OriginPath field if non-nil, zero value otherwise.

### GetOriginPathOk

`func (o *OriginsResultResponse) GetOriginPathOk() (*string, bool)`

GetOriginPathOk returns a tuple with the OriginPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPath

`func (o *OriginsResultResponse) SetOriginPath(v string)`

SetOriginPath sets OriginPath field to given value.

### HasOriginPath

`func (o *OriginsResultResponse) HasOriginPath() bool`

HasOriginPath returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *OriginsResultResponse) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *OriginsResultResponse) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *OriginsResultResponse) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *OriginsResultResponse) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetTimeoutBetweenBytes

`func (o *OriginsResultResponse) GetTimeoutBetweenBytes() int64`

GetTimeoutBetweenBytes returns the TimeoutBetweenBytes field if non-nil, zero value otherwise.

### GetTimeoutBetweenBytesOk

`func (o *OriginsResultResponse) GetTimeoutBetweenBytesOk() (*int64, bool)`

GetTimeoutBetweenBytesOk returns a tuple with the TimeoutBetweenBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutBetweenBytes

`func (o *OriginsResultResponse) SetTimeoutBetweenBytes(v int64)`

SetTimeoutBetweenBytes sets TimeoutBetweenBytes field to given value.

### HasTimeoutBetweenBytes

`func (o *OriginsResultResponse) HasTimeoutBetweenBytes() bool`

HasTimeoutBetweenBytes returns a boolean if a field has been set.

### GetHmacAuthentication

`func (o *OriginsResultResponse) GetHmacAuthentication() bool`

GetHmacAuthentication returns the HmacAuthentication field if non-nil, zero value otherwise.

### GetHmacAuthenticationOk

`func (o *OriginsResultResponse) GetHmacAuthenticationOk() (*bool, bool)`

GetHmacAuthenticationOk returns a tuple with the HmacAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAuthentication

`func (o *OriginsResultResponse) SetHmacAuthentication(v bool)`

SetHmacAuthentication sets HmacAuthentication field to given value.

### HasHmacAuthentication

`func (o *OriginsResultResponse) HasHmacAuthentication() bool`

HasHmacAuthentication returns a boolean if a field has been set.

### GetHmacRegionName

`func (o *OriginsResultResponse) GetHmacRegionName() string`

GetHmacRegionName returns the HmacRegionName field if non-nil, zero value otherwise.

### GetHmacRegionNameOk

`func (o *OriginsResultResponse) GetHmacRegionNameOk() (*string, bool)`

GetHmacRegionNameOk returns a tuple with the HmacRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacRegionName

`func (o *OriginsResultResponse) SetHmacRegionName(v string)`

SetHmacRegionName sets HmacRegionName field to given value.

### HasHmacRegionName

`func (o *OriginsResultResponse) HasHmacRegionName() bool`

HasHmacRegionName returns a boolean if a field has been set.

### GetHmacAccessKey

`func (o *OriginsResultResponse) GetHmacAccessKey() string`

GetHmacAccessKey returns the HmacAccessKey field if non-nil, zero value otherwise.

### GetHmacAccessKeyOk

`func (o *OriginsResultResponse) GetHmacAccessKeyOk() (*string, bool)`

GetHmacAccessKeyOk returns a tuple with the HmacAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAccessKey

`func (o *OriginsResultResponse) SetHmacAccessKey(v string)`

SetHmacAccessKey sets HmacAccessKey field to given value.

### HasHmacAccessKey

`func (o *OriginsResultResponse) HasHmacAccessKey() bool`

HasHmacAccessKey returns a boolean if a field has been set.

### GetHmacSecretKey

`func (o *OriginsResultResponse) GetHmacSecretKey() string`

GetHmacSecretKey returns the HmacSecretKey field if non-nil, zero value otherwise.

### GetHmacSecretKeyOk

`func (o *OriginsResultResponse) GetHmacSecretKeyOk() (*string, bool)`

GetHmacSecretKeyOk returns a tuple with the HmacSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacSecretKey

`func (o *OriginsResultResponse) SetHmacSecretKey(v string)`

SetHmacSecretKey sets HmacSecretKey field to given value.

### HasHmacSecretKey

`func (o *OriginsResultResponse) HasHmacSecretKey() bool`

HasHmacSecretKey returns a boolean if a field has been set.

### GetBucket

`func (o *OriginsResultResponse) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *OriginsResultResponse) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *OriginsResultResponse) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *OriginsResultResponse) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetPrefix

`func (o *OriginsResultResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *OriginsResultResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *OriginsResultResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *OriginsResultResponse) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


