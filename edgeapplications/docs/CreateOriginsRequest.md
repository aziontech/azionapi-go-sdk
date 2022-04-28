# CreateOriginsRequest

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

## Methods

### NewCreateOriginsRequest

`func NewCreateOriginsRequest(name string, addresses []CreateOriginsRequestAddresses, hostHeader string, ) *CreateOriginsRequest`

NewCreateOriginsRequest instantiates a new CreateOriginsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOriginsRequestWithDefaults

`func NewCreateOriginsRequestWithDefaults() *CreateOriginsRequest`

NewCreateOriginsRequestWithDefaults instantiates a new CreateOriginsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOriginsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOriginsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOriginsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOriginType

`func (o *CreateOriginsRequest) GetOriginType() string`

GetOriginType returns the OriginType field if non-nil, zero value otherwise.

### GetOriginTypeOk

`func (o *CreateOriginsRequest) GetOriginTypeOk() (*string, bool)`

GetOriginTypeOk returns a tuple with the OriginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginType

`func (o *CreateOriginsRequest) SetOriginType(v string)`

SetOriginType sets OriginType field to given value.

### HasOriginType

`func (o *CreateOriginsRequest) HasOriginType() bool`

HasOriginType returns a boolean if a field has been set.

### GetAddresses

`func (o *CreateOriginsRequest) GetAddresses() []CreateOriginsRequestAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CreateOriginsRequest) GetAddressesOk() (*[]CreateOriginsRequestAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CreateOriginsRequest) SetAddresses(v []CreateOriginsRequestAddresses)`

SetAddresses sets Addresses field to given value.


### GetOriginProtocolPolicy

`func (o *CreateOriginsRequest) GetOriginProtocolPolicy() string`

GetOriginProtocolPolicy returns the OriginProtocolPolicy field if non-nil, zero value otherwise.

### GetOriginProtocolPolicyOk

`func (o *CreateOriginsRequest) GetOriginProtocolPolicyOk() (*string, bool)`

GetOriginProtocolPolicyOk returns a tuple with the OriginProtocolPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProtocolPolicy

`func (o *CreateOriginsRequest) SetOriginProtocolPolicy(v string)`

SetOriginProtocolPolicy sets OriginProtocolPolicy field to given value.

### HasOriginProtocolPolicy

`func (o *CreateOriginsRequest) HasOriginProtocolPolicy() bool`

HasOriginProtocolPolicy returns a boolean if a field has been set.

### GetHostHeader

`func (o *CreateOriginsRequest) GetHostHeader() string`

GetHostHeader returns the HostHeader field if non-nil, zero value otherwise.

### GetHostHeaderOk

`func (o *CreateOriginsRequest) GetHostHeaderOk() (*string, bool)`

GetHostHeaderOk returns a tuple with the HostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostHeader

`func (o *CreateOriginsRequest) SetHostHeader(v string)`

SetHostHeader sets HostHeader field to given value.


### GetOriginPath

`func (o *CreateOriginsRequest) GetOriginPath() string`

GetOriginPath returns the OriginPath field if non-nil, zero value otherwise.

### GetOriginPathOk

`func (o *CreateOriginsRequest) GetOriginPathOk() (*string, bool)`

GetOriginPathOk returns a tuple with the OriginPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPath

`func (o *CreateOriginsRequest) SetOriginPath(v string)`

SetOriginPath sets OriginPath field to given value.

### HasOriginPath

`func (o *CreateOriginsRequest) HasOriginPath() bool`

HasOriginPath returns a boolean if a field has been set.

### GetHmacAuthentication

`func (o *CreateOriginsRequest) GetHmacAuthentication() bool`

GetHmacAuthentication returns the HmacAuthentication field if non-nil, zero value otherwise.

### GetHmacAuthenticationOk

`func (o *CreateOriginsRequest) GetHmacAuthenticationOk() (*bool, bool)`

GetHmacAuthenticationOk returns a tuple with the HmacAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAuthentication

`func (o *CreateOriginsRequest) SetHmacAuthentication(v bool)`

SetHmacAuthentication sets HmacAuthentication field to given value.

### HasHmacAuthentication

`func (o *CreateOriginsRequest) HasHmacAuthentication() bool`

HasHmacAuthentication returns a boolean if a field has been set.

### GetHmacRegionName

`func (o *CreateOriginsRequest) GetHmacRegionName() string`

GetHmacRegionName returns the HmacRegionName field if non-nil, zero value otherwise.

### GetHmacRegionNameOk

`func (o *CreateOriginsRequest) GetHmacRegionNameOk() (*string, bool)`

GetHmacRegionNameOk returns a tuple with the HmacRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacRegionName

`func (o *CreateOriginsRequest) SetHmacRegionName(v string)`

SetHmacRegionName sets HmacRegionName field to given value.

### HasHmacRegionName

`func (o *CreateOriginsRequest) HasHmacRegionName() bool`

HasHmacRegionName returns a boolean if a field has been set.

### GetHmacAccessKey

`func (o *CreateOriginsRequest) GetHmacAccessKey() string`

GetHmacAccessKey returns the HmacAccessKey field if non-nil, zero value otherwise.

### GetHmacAccessKeyOk

`func (o *CreateOriginsRequest) GetHmacAccessKeyOk() (*string, bool)`

GetHmacAccessKeyOk returns a tuple with the HmacAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAccessKey

`func (o *CreateOriginsRequest) SetHmacAccessKey(v string)`

SetHmacAccessKey sets HmacAccessKey field to given value.

### HasHmacAccessKey

`func (o *CreateOriginsRequest) HasHmacAccessKey() bool`

HasHmacAccessKey returns a boolean if a field has been set.

### GetHmacSecretKey

`func (o *CreateOriginsRequest) GetHmacSecretKey() string`

GetHmacSecretKey returns the HmacSecretKey field if non-nil, zero value otherwise.

### GetHmacSecretKeyOk

`func (o *CreateOriginsRequest) GetHmacSecretKeyOk() (*string, bool)`

GetHmacSecretKeyOk returns a tuple with the HmacSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacSecretKey

`func (o *CreateOriginsRequest) SetHmacSecretKey(v string)`

SetHmacSecretKey sets HmacSecretKey field to given value.

### HasHmacSecretKey

`func (o *CreateOriginsRequest) HasHmacSecretKey() bool`

HasHmacSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


