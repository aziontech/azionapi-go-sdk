# UpdateOriginsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**OriginType** | Pointer to **string** |  | [optional] 
**Addresses** | [**[]CreateOriginsRequestAddresses**](CreateOriginsRequestAddresses.md) |  | 
**OriginProtocolPolicy** | Pointer to **string** |  | [optional] 
**HostHeader** | **string** |  | 
**OriginPath** | Pointer to **string** |  | [optional] 
**HmacAuthentication** | Pointer to **bool** |  | [optional] 
**HmacRegionName** | Pointer to **string** |  | [optional] 
**HmacAccessKey** | Pointer to **string** |  | [optional] 
**HmacSecretKey** | Pointer to **string** |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateOriginsRequest

`func NewUpdateOriginsRequest(name string, addresses []CreateOriginsRequestAddresses, hostHeader string, ) *UpdateOriginsRequest`

NewUpdateOriginsRequest instantiates a new UpdateOriginsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOriginsRequestWithDefaults

`func NewUpdateOriginsRequestWithDefaults() *UpdateOriginsRequest`

NewUpdateOriginsRequestWithDefaults instantiates a new UpdateOriginsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateOriginsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOriginsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOriginsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOriginType

`func (o *UpdateOriginsRequest) GetOriginType() string`

GetOriginType returns the OriginType field if non-nil, zero value otherwise.

### GetOriginTypeOk

`func (o *UpdateOriginsRequest) GetOriginTypeOk() (*string, bool)`

GetOriginTypeOk returns a tuple with the OriginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginType

`func (o *UpdateOriginsRequest) SetOriginType(v string)`

SetOriginType sets OriginType field to given value.

### HasOriginType

`func (o *UpdateOriginsRequest) HasOriginType() bool`

HasOriginType returns a boolean if a field has been set.

### GetAddresses

`func (o *UpdateOriginsRequest) GetAddresses() []CreateOriginsRequestAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *UpdateOriginsRequest) GetAddressesOk() (*[]CreateOriginsRequestAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *UpdateOriginsRequest) SetAddresses(v []CreateOriginsRequestAddresses)`

SetAddresses sets Addresses field to given value.


### GetOriginProtocolPolicy

`func (o *UpdateOriginsRequest) GetOriginProtocolPolicy() string`

GetOriginProtocolPolicy returns the OriginProtocolPolicy field if non-nil, zero value otherwise.

### GetOriginProtocolPolicyOk

`func (o *UpdateOriginsRequest) GetOriginProtocolPolicyOk() (*string, bool)`

GetOriginProtocolPolicyOk returns a tuple with the OriginProtocolPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProtocolPolicy

`func (o *UpdateOriginsRequest) SetOriginProtocolPolicy(v string)`

SetOriginProtocolPolicy sets OriginProtocolPolicy field to given value.

### HasOriginProtocolPolicy

`func (o *UpdateOriginsRequest) HasOriginProtocolPolicy() bool`

HasOriginProtocolPolicy returns a boolean if a field has been set.

### GetHostHeader

`func (o *UpdateOriginsRequest) GetHostHeader() string`

GetHostHeader returns the HostHeader field if non-nil, zero value otherwise.

### GetHostHeaderOk

`func (o *UpdateOriginsRequest) GetHostHeaderOk() (*string, bool)`

GetHostHeaderOk returns a tuple with the HostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostHeader

`func (o *UpdateOriginsRequest) SetHostHeader(v string)`

SetHostHeader sets HostHeader field to given value.


### GetOriginPath

`func (o *UpdateOriginsRequest) GetOriginPath() string`

GetOriginPath returns the OriginPath field if non-nil, zero value otherwise.

### GetOriginPathOk

`func (o *UpdateOriginsRequest) GetOriginPathOk() (*string, bool)`

GetOriginPathOk returns a tuple with the OriginPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPath

`func (o *UpdateOriginsRequest) SetOriginPath(v string)`

SetOriginPath sets OriginPath field to given value.

### HasOriginPath

`func (o *UpdateOriginsRequest) HasOriginPath() bool`

HasOriginPath returns a boolean if a field has been set.

### GetHmacAuthentication

`func (o *UpdateOriginsRequest) GetHmacAuthentication() bool`

GetHmacAuthentication returns the HmacAuthentication field if non-nil, zero value otherwise.

### GetHmacAuthenticationOk

`func (o *UpdateOriginsRequest) GetHmacAuthenticationOk() (*bool, bool)`

GetHmacAuthenticationOk returns a tuple with the HmacAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAuthentication

`func (o *UpdateOriginsRequest) SetHmacAuthentication(v bool)`

SetHmacAuthentication sets HmacAuthentication field to given value.

### HasHmacAuthentication

`func (o *UpdateOriginsRequest) HasHmacAuthentication() bool`

HasHmacAuthentication returns a boolean if a field has been set.

### GetHmacRegionName

`func (o *UpdateOriginsRequest) GetHmacRegionName() string`

GetHmacRegionName returns the HmacRegionName field if non-nil, zero value otherwise.

### GetHmacRegionNameOk

`func (o *UpdateOriginsRequest) GetHmacRegionNameOk() (*string, bool)`

GetHmacRegionNameOk returns a tuple with the HmacRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacRegionName

`func (o *UpdateOriginsRequest) SetHmacRegionName(v string)`

SetHmacRegionName sets HmacRegionName field to given value.

### HasHmacRegionName

`func (o *UpdateOriginsRequest) HasHmacRegionName() bool`

HasHmacRegionName returns a boolean if a field has been set.

### GetHmacAccessKey

`func (o *UpdateOriginsRequest) GetHmacAccessKey() string`

GetHmacAccessKey returns the HmacAccessKey field if non-nil, zero value otherwise.

### GetHmacAccessKeyOk

`func (o *UpdateOriginsRequest) GetHmacAccessKeyOk() (*string, bool)`

GetHmacAccessKeyOk returns a tuple with the HmacAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAccessKey

`func (o *UpdateOriginsRequest) SetHmacAccessKey(v string)`

SetHmacAccessKey sets HmacAccessKey field to given value.

### HasHmacAccessKey

`func (o *UpdateOriginsRequest) HasHmacAccessKey() bool`

HasHmacAccessKey returns a boolean if a field has been set.

### GetHmacSecretKey

`func (o *UpdateOriginsRequest) GetHmacSecretKey() string`

GetHmacSecretKey returns the HmacSecretKey field if non-nil, zero value otherwise.

### GetHmacSecretKeyOk

`func (o *UpdateOriginsRequest) GetHmacSecretKeyOk() (*string, bool)`

GetHmacSecretKeyOk returns a tuple with the HmacSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacSecretKey

`func (o *UpdateOriginsRequest) SetHmacSecretKey(v string)`

SetHmacSecretKey sets HmacSecretKey field to given value.

### HasHmacSecretKey

`func (o *UpdateOriginsRequest) HasHmacSecretKey() bool`

HasHmacSecretKey returns a boolean if a field has been set.

### GetBucket

`func (o *UpdateOriginsRequest) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *UpdateOriginsRequest) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *UpdateOriginsRequest) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *UpdateOriginsRequest) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetPrefix

`func (o *UpdateOriginsRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *UpdateOriginsRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *UpdateOriginsRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *UpdateOriginsRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


