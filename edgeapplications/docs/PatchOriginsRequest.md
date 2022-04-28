# PatchOriginsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**OriginType** | Pointer to **string** |  | [optional] 
**Addresses** | Pointer to [**[]CreateOriginsRequestAddresses**](CreateOriginsRequestAddresses.md) |  | [optional] 
**OriginProtocolPolicy** | Pointer to **string** |  | [optional] 
**HostHeader** | Pointer to **string** |  | [optional] 
**OriginPath** | Pointer to **string** |  | [optional] 
**HmacAuthentication** | Pointer to **bool** |  | [optional] 
**HmacRegionName** | Pointer to **string** |  | [optional] 
**HmacAccessKey** | Pointer to **string** |  | [optional] 
**HmacSecretKey** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchOriginsRequest

`func NewPatchOriginsRequest() *PatchOriginsRequest`

NewPatchOriginsRequest instantiates a new PatchOriginsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchOriginsRequestWithDefaults

`func NewPatchOriginsRequestWithDefaults() *PatchOriginsRequest`

NewPatchOriginsRequestWithDefaults instantiates a new PatchOriginsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchOriginsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchOriginsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchOriginsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchOriginsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginType

`func (o *PatchOriginsRequest) GetOriginType() string`

GetOriginType returns the OriginType field if non-nil, zero value otherwise.

### GetOriginTypeOk

`func (o *PatchOriginsRequest) GetOriginTypeOk() (*string, bool)`

GetOriginTypeOk returns a tuple with the OriginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginType

`func (o *PatchOriginsRequest) SetOriginType(v string)`

SetOriginType sets OriginType field to given value.

### HasOriginType

`func (o *PatchOriginsRequest) HasOriginType() bool`

HasOriginType returns a boolean if a field has been set.

### GetAddresses

`func (o *PatchOriginsRequest) GetAddresses() []CreateOriginsRequestAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PatchOriginsRequest) GetAddressesOk() (*[]CreateOriginsRequestAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PatchOriginsRequest) SetAddresses(v []CreateOriginsRequestAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PatchOriginsRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetOriginProtocolPolicy

`func (o *PatchOriginsRequest) GetOriginProtocolPolicy() string`

GetOriginProtocolPolicy returns the OriginProtocolPolicy field if non-nil, zero value otherwise.

### GetOriginProtocolPolicyOk

`func (o *PatchOriginsRequest) GetOriginProtocolPolicyOk() (*string, bool)`

GetOriginProtocolPolicyOk returns a tuple with the OriginProtocolPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProtocolPolicy

`func (o *PatchOriginsRequest) SetOriginProtocolPolicy(v string)`

SetOriginProtocolPolicy sets OriginProtocolPolicy field to given value.

### HasOriginProtocolPolicy

`func (o *PatchOriginsRequest) HasOriginProtocolPolicy() bool`

HasOriginProtocolPolicy returns a boolean if a field has been set.

### GetHostHeader

`func (o *PatchOriginsRequest) GetHostHeader() string`

GetHostHeader returns the HostHeader field if non-nil, zero value otherwise.

### GetHostHeaderOk

`func (o *PatchOriginsRequest) GetHostHeaderOk() (*string, bool)`

GetHostHeaderOk returns a tuple with the HostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostHeader

`func (o *PatchOriginsRequest) SetHostHeader(v string)`

SetHostHeader sets HostHeader field to given value.

### HasHostHeader

`func (o *PatchOriginsRequest) HasHostHeader() bool`

HasHostHeader returns a boolean if a field has been set.

### GetOriginPath

`func (o *PatchOriginsRequest) GetOriginPath() string`

GetOriginPath returns the OriginPath field if non-nil, zero value otherwise.

### GetOriginPathOk

`func (o *PatchOriginsRequest) GetOriginPathOk() (*string, bool)`

GetOriginPathOk returns a tuple with the OriginPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPath

`func (o *PatchOriginsRequest) SetOriginPath(v string)`

SetOriginPath sets OriginPath field to given value.

### HasOriginPath

`func (o *PatchOriginsRequest) HasOriginPath() bool`

HasOriginPath returns a boolean if a field has been set.

### GetHmacAuthentication

`func (o *PatchOriginsRequest) GetHmacAuthentication() bool`

GetHmacAuthentication returns the HmacAuthentication field if non-nil, zero value otherwise.

### GetHmacAuthenticationOk

`func (o *PatchOriginsRequest) GetHmacAuthenticationOk() (*bool, bool)`

GetHmacAuthenticationOk returns a tuple with the HmacAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAuthentication

`func (o *PatchOriginsRequest) SetHmacAuthentication(v bool)`

SetHmacAuthentication sets HmacAuthentication field to given value.

### HasHmacAuthentication

`func (o *PatchOriginsRequest) HasHmacAuthentication() bool`

HasHmacAuthentication returns a boolean if a field has been set.

### GetHmacRegionName

`func (o *PatchOriginsRequest) GetHmacRegionName() string`

GetHmacRegionName returns the HmacRegionName field if non-nil, zero value otherwise.

### GetHmacRegionNameOk

`func (o *PatchOriginsRequest) GetHmacRegionNameOk() (*string, bool)`

GetHmacRegionNameOk returns a tuple with the HmacRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacRegionName

`func (o *PatchOriginsRequest) SetHmacRegionName(v string)`

SetHmacRegionName sets HmacRegionName field to given value.

### HasHmacRegionName

`func (o *PatchOriginsRequest) HasHmacRegionName() bool`

HasHmacRegionName returns a boolean if a field has been set.

### GetHmacAccessKey

`func (o *PatchOriginsRequest) GetHmacAccessKey() string`

GetHmacAccessKey returns the HmacAccessKey field if non-nil, zero value otherwise.

### GetHmacAccessKeyOk

`func (o *PatchOriginsRequest) GetHmacAccessKeyOk() (*string, bool)`

GetHmacAccessKeyOk returns a tuple with the HmacAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacAccessKey

`func (o *PatchOriginsRequest) SetHmacAccessKey(v string)`

SetHmacAccessKey sets HmacAccessKey field to given value.

### HasHmacAccessKey

`func (o *PatchOriginsRequest) HasHmacAccessKey() bool`

HasHmacAccessKey returns a boolean if a field has been set.

### GetHmacSecretKey

`func (o *PatchOriginsRequest) GetHmacSecretKey() string`

GetHmacSecretKey returns the HmacSecretKey field if non-nil, zero value otherwise.

### GetHmacSecretKeyOk

`func (o *PatchOriginsRequest) GetHmacSecretKeyOk() (*string, bool)`

GetHmacSecretKeyOk returns a tuple with the HmacSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmacSecretKey

`func (o *PatchOriginsRequest) SetHmacSecretKey(v string)`

SetHmacSecretKey sets HmacSecretKey field to given value.

### HasHmacSecretKey

`func (o *PatchOriginsRequest) HasHmacSecretKey() bool`

HasHmacSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


