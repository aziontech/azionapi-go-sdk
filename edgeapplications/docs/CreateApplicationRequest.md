# CreateApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DeliveryProtocol** | Pointer to **string** |  | [optional] 
**OriginType** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**OriginProtocolPolicy** | Pointer to **string** |  | [optional] 
**HostHeader** | Pointer to **string** |  | [optional] 
**BrowserCacheSettings** | Pointer to **string** |  | [optional] 
**CdnCacheSettings** | Pointer to **string** |  | [optional] 
**BrowserCacheSettingsMaximumTtl** | Pointer to **int64** |  | [optional] 
**CdnCacheSettingsMaximumTtl** | Pointer to **int64** |  | [optional] 

## Methods

### NewCreateApplicationRequest

`func NewCreateApplicationRequest(name string, ) *CreateApplicationRequest`

NewCreateApplicationRequest instantiates a new CreateApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApplicationRequestWithDefaults

`func NewCreateApplicationRequestWithDefaults() *CreateApplicationRequest`

NewCreateApplicationRequestWithDefaults instantiates a new CreateApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateApplicationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApplicationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApplicationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDeliveryProtocol

`func (o *CreateApplicationRequest) GetDeliveryProtocol() string`

GetDeliveryProtocol returns the DeliveryProtocol field if non-nil, zero value otherwise.

### GetDeliveryProtocolOk

`func (o *CreateApplicationRequest) GetDeliveryProtocolOk() (*string, bool)`

GetDeliveryProtocolOk returns a tuple with the DeliveryProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryProtocol

`func (o *CreateApplicationRequest) SetDeliveryProtocol(v string)`

SetDeliveryProtocol sets DeliveryProtocol field to given value.

### HasDeliveryProtocol

`func (o *CreateApplicationRequest) HasDeliveryProtocol() bool`

HasDeliveryProtocol returns a boolean if a field has been set.

### GetOriginType

`func (o *CreateApplicationRequest) GetOriginType() string`

GetOriginType returns the OriginType field if non-nil, zero value otherwise.

### GetOriginTypeOk

`func (o *CreateApplicationRequest) GetOriginTypeOk() (*string, bool)`

GetOriginTypeOk returns a tuple with the OriginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginType

`func (o *CreateApplicationRequest) SetOriginType(v string)`

SetOriginType sets OriginType field to given value.

### HasOriginType

`func (o *CreateApplicationRequest) HasOriginType() bool`

HasOriginType returns a boolean if a field has been set.

### GetAddress

`func (o *CreateApplicationRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateApplicationRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateApplicationRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CreateApplicationRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetOriginProtocolPolicy

`func (o *CreateApplicationRequest) GetOriginProtocolPolicy() string`

GetOriginProtocolPolicy returns the OriginProtocolPolicy field if non-nil, zero value otherwise.

### GetOriginProtocolPolicyOk

`func (o *CreateApplicationRequest) GetOriginProtocolPolicyOk() (*string, bool)`

GetOriginProtocolPolicyOk returns a tuple with the OriginProtocolPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProtocolPolicy

`func (o *CreateApplicationRequest) SetOriginProtocolPolicy(v string)`

SetOriginProtocolPolicy sets OriginProtocolPolicy field to given value.

### HasOriginProtocolPolicy

`func (o *CreateApplicationRequest) HasOriginProtocolPolicy() bool`

HasOriginProtocolPolicy returns a boolean if a field has been set.

### GetHostHeader

`func (o *CreateApplicationRequest) GetHostHeader() string`

GetHostHeader returns the HostHeader field if non-nil, zero value otherwise.

### GetHostHeaderOk

`func (o *CreateApplicationRequest) GetHostHeaderOk() (*string, bool)`

GetHostHeaderOk returns a tuple with the HostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostHeader

`func (o *CreateApplicationRequest) SetHostHeader(v string)`

SetHostHeader sets HostHeader field to given value.

### HasHostHeader

`func (o *CreateApplicationRequest) HasHostHeader() bool`

HasHostHeader returns a boolean if a field has been set.

### GetBrowserCacheSettings

`func (o *CreateApplicationRequest) GetBrowserCacheSettings() string`

GetBrowserCacheSettings returns the BrowserCacheSettings field if non-nil, zero value otherwise.

### GetBrowserCacheSettingsOk

`func (o *CreateApplicationRequest) GetBrowserCacheSettingsOk() (*string, bool)`

GetBrowserCacheSettingsOk returns a tuple with the BrowserCacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCacheSettings

`func (o *CreateApplicationRequest) SetBrowserCacheSettings(v string)`

SetBrowserCacheSettings sets BrowserCacheSettings field to given value.

### HasBrowserCacheSettings

`func (o *CreateApplicationRequest) HasBrowserCacheSettings() bool`

HasBrowserCacheSettings returns a boolean if a field has been set.

### GetCdnCacheSettings

`func (o *CreateApplicationRequest) GetCdnCacheSettings() string`

GetCdnCacheSettings returns the CdnCacheSettings field if non-nil, zero value otherwise.

### GetCdnCacheSettingsOk

`func (o *CreateApplicationRequest) GetCdnCacheSettingsOk() (*string, bool)`

GetCdnCacheSettingsOk returns a tuple with the CdnCacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnCacheSettings

`func (o *CreateApplicationRequest) SetCdnCacheSettings(v string)`

SetCdnCacheSettings sets CdnCacheSettings field to given value.

### HasCdnCacheSettings

`func (o *CreateApplicationRequest) HasCdnCacheSettings() bool`

HasCdnCacheSettings returns a boolean if a field has been set.

### GetBrowserCacheSettingsMaximumTtl

`func (o *CreateApplicationRequest) GetBrowserCacheSettingsMaximumTtl() int64`

GetBrowserCacheSettingsMaximumTtl returns the BrowserCacheSettingsMaximumTtl field if non-nil, zero value otherwise.

### GetBrowserCacheSettingsMaximumTtlOk

`func (o *CreateApplicationRequest) GetBrowserCacheSettingsMaximumTtlOk() (*int64, bool)`

GetBrowserCacheSettingsMaximumTtlOk returns a tuple with the BrowserCacheSettingsMaximumTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCacheSettingsMaximumTtl

`func (o *CreateApplicationRequest) SetBrowserCacheSettingsMaximumTtl(v int64)`

SetBrowserCacheSettingsMaximumTtl sets BrowserCacheSettingsMaximumTtl field to given value.

### HasBrowserCacheSettingsMaximumTtl

`func (o *CreateApplicationRequest) HasBrowserCacheSettingsMaximumTtl() bool`

HasBrowserCacheSettingsMaximumTtl returns a boolean if a field has been set.

### GetCdnCacheSettingsMaximumTtl

`func (o *CreateApplicationRequest) GetCdnCacheSettingsMaximumTtl() int64`

GetCdnCacheSettingsMaximumTtl returns the CdnCacheSettingsMaximumTtl field if non-nil, zero value otherwise.

### GetCdnCacheSettingsMaximumTtlOk

`func (o *CreateApplicationRequest) GetCdnCacheSettingsMaximumTtlOk() (*int64, bool)`

GetCdnCacheSettingsMaximumTtlOk returns a tuple with the CdnCacheSettingsMaximumTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnCacheSettingsMaximumTtl

`func (o *CreateApplicationRequest) SetCdnCacheSettingsMaximumTtl(v int64)`

SetCdnCacheSettingsMaximumTtl sets CdnCacheSettingsMaximumTtl field to given value.

### HasCdnCacheSettingsMaximumTtl

`func (o *CreateApplicationRequest) HasCdnCacheSettingsMaximumTtl() bool`

HasCdnCacheSettingsMaximumTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


