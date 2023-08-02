# CreateApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ApplicationAcceleration** | Pointer to **bool** |  | [optional] 
**DeliveryProtocol** | Pointer to **string** |  | [optional] 
**OriginType** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**OriginProtocolPolicy** | Pointer to **string** |  | [optional] 
**HostHeader** | Pointer to **string** |  | [optional] 
**BrowserCacheSettings** | Pointer to **string** |  | [optional] 
**CdnCacheSettings** | Pointer to **string** |  | [optional] 
**BrowserCacheSettingsMaximumTtl** | Pointer to **int64** |  | [optional] 
**CdnCacheSettingsMaximumTtl** | Pointer to **int64** |  | [optional] 
**DebugRules** | Pointer to **bool** |  | [optional] 
**SupportedCiphers** | Pointer to **string** |  | [optional] 
**HttpPort** | Pointer to **interface{}** |  | [optional] 
**HttpsPort** | Pointer to **interface{}** |  | [optional] 
**L2Caching** | Pointer to **bool** |  | [optional] 
**Http3** | Pointer to **bool** |  | [optional] 
**MinimumTlsVersion** | Pointer to **string** |  | [optional] 

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


### GetApplicationAcceleration

`func (o *CreateApplicationRequest) GetApplicationAcceleration() bool`

GetApplicationAcceleration returns the ApplicationAcceleration field if non-nil, zero value otherwise.

### GetApplicationAccelerationOk

`func (o *CreateApplicationRequest) GetApplicationAccelerationOk() (*bool, bool)`

GetApplicationAccelerationOk returns a tuple with the ApplicationAcceleration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAcceleration

`func (o *CreateApplicationRequest) SetApplicationAcceleration(v bool)`

SetApplicationAcceleration sets ApplicationAcceleration field to given value.

### HasApplicationAcceleration

`func (o *CreateApplicationRequest) HasApplicationAcceleration() bool`

HasApplicationAcceleration returns a boolean if a field has been set.

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

### GetDebugRules

`func (o *CreateApplicationRequest) GetDebugRules() bool`

GetDebugRules returns the DebugRules field if non-nil, zero value otherwise.

### GetDebugRulesOk

`func (o *CreateApplicationRequest) GetDebugRulesOk() (*bool, bool)`

GetDebugRulesOk returns a tuple with the DebugRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugRules

`func (o *CreateApplicationRequest) SetDebugRules(v bool)`

SetDebugRules sets DebugRules field to given value.

### HasDebugRules

`func (o *CreateApplicationRequest) HasDebugRules() bool`

HasDebugRules returns a boolean if a field has been set.

### GetSupportedCiphers

`func (o *CreateApplicationRequest) GetSupportedCiphers() string`

GetSupportedCiphers returns the SupportedCiphers field if non-nil, zero value otherwise.

### GetSupportedCiphersOk

`func (o *CreateApplicationRequest) GetSupportedCiphersOk() (*string, bool)`

GetSupportedCiphersOk returns a tuple with the SupportedCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCiphers

`func (o *CreateApplicationRequest) SetSupportedCiphers(v string)`

SetSupportedCiphers sets SupportedCiphers field to given value.

### HasSupportedCiphers

`func (o *CreateApplicationRequest) HasSupportedCiphers() bool`

HasSupportedCiphers returns a boolean if a field has been set.

### GetHttpPort

`func (o *CreateApplicationRequest) GetHttpPort() interface{}`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *CreateApplicationRequest) GetHttpPortOk() (*interface{}, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *CreateApplicationRequest) SetHttpPort(v interface{})`

SetHttpPort sets HttpPort field to given value.

### HasHttpPort

`func (o *CreateApplicationRequest) HasHttpPort() bool`

HasHttpPort returns a boolean if a field has been set.

### SetHttpPortNil

`func (o *CreateApplicationRequest) SetHttpPortNil(b bool)`

 SetHttpPortNil sets the value for HttpPort to be an explicit nil

### UnsetHttpPort
`func (o *CreateApplicationRequest) UnsetHttpPort()`

UnsetHttpPort ensures that no value is present for HttpPort, not even an explicit nil
### GetHttpsPort

`func (o *CreateApplicationRequest) GetHttpsPort() interface{}`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *CreateApplicationRequest) GetHttpsPortOk() (*interface{}, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *CreateApplicationRequest) SetHttpsPort(v interface{})`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *CreateApplicationRequest) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### SetHttpsPortNil

`func (o *CreateApplicationRequest) SetHttpsPortNil(b bool)`

 SetHttpsPortNil sets the value for HttpsPort to be an explicit nil

### UnsetHttpsPort
`func (o *CreateApplicationRequest) UnsetHttpsPort()`

UnsetHttpsPort ensures that no value is present for HttpsPort, not even an explicit nil
### GetL2Caching

`func (o *CreateApplicationRequest) GetL2Caching() bool`

GetL2Caching returns the L2Caching field if non-nil, zero value otherwise.

### GetL2CachingOk

`func (o *CreateApplicationRequest) GetL2CachingOk() (*bool, bool)`

GetL2CachingOk returns a tuple with the L2Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2Caching

`func (o *CreateApplicationRequest) SetL2Caching(v bool)`

SetL2Caching sets L2Caching field to given value.

### HasL2Caching

`func (o *CreateApplicationRequest) HasL2Caching() bool`

HasL2Caching returns a boolean if a field has been set.

### GetHttp3

`func (o *CreateApplicationRequest) GetHttp3() bool`

GetHttp3 returns the Http3 field if non-nil, zero value otherwise.

### GetHttp3Ok

`func (o *CreateApplicationRequest) GetHttp3Ok() (*bool, bool)`

GetHttp3Ok returns a tuple with the Http3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp3

`func (o *CreateApplicationRequest) SetHttp3(v bool)`

SetHttp3 sets Http3 field to given value.

### HasHttp3

`func (o *CreateApplicationRequest) HasHttp3() bool`

HasHttp3 returns a boolean if a field has been set.

### GetMinimumTlsVersion

`func (o *CreateApplicationRequest) GetMinimumTlsVersion() string`

GetMinimumTlsVersion returns the MinimumTlsVersion field if non-nil, zero value otherwise.

### GetMinimumTlsVersionOk

`func (o *CreateApplicationRequest) GetMinimumTlsVersionOk() (*string, bool)`

GetMinimumTlsVersionOk returns a tuple with the MinimumTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTlsVersion

`func (o *CreateApplicationRequest) SetMinimumTlsVersion(v string)`

SetMinimumTlsVersion sets MinimumTlsVersion field to given value.

### HasMinimumTlsVersion

`func (o *CreateApplicationRequest) HasMinimumTlsVersion() bool`

HasMinimumTlsVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


