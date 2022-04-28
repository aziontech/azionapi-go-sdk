# OriginsResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginId** | **int64** |  | 
**OriginKey** | **string** |  | 
**Name** | **string** |  | 
**OriginType** | **string** |  | 
**Addresses** | [**[]OriginsResultResponseAddresses**](OriginsResultResponseAddresses.md) |  | 
**OriginProtocolPolicy** | **string** |  | 
**IsOriginRedirectionEnabled** | **bool** |  | 
**HostHeader** | **string** |  | 
**Method** | **string** |  | 
**OriginPath** | **string** |  | 
**ConnectionTimeout** | **int64** |  | 
**TimeoutBetweenBytes** | **int64** |  | 
**HmacAuthentication** | **bool** |  | 
**HmacRegionName** | **string** |  | 
**HmacAccessKey** | **string** |  | 
**HmacSecretKey** | **string** |  | 

## Methods

### NewOriginsResultResponse

`func NewOriginsResultResponse(originId int64, originKey string, name string, originType string, addresses []OriginsResultResponseAddresses, originProtocolPolicy string, isOriginRedirectionEnabled bool, hostHeader string, method string, originPath string, connectionTimeout int64, timeoutBetweenBytes int64, hmacAuthentication bool, hmacRegionName string, hmacAccessKey string, hmacSecretKey string, ) *OriginsResultResponse`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


