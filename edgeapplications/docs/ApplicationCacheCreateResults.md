# ApplicationCacheCreateResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**BrowserCacheSettings** | **string** |  | 
**BrowserCacheSettingsMaximumTtl** | **int64** |  | 
**CdnCacheSettings** | **string** |  | 
**CdnCacheSettingsMaximumTtl** | **int64** |  | 
**CacheByQueryString** | **string** |  | 
**QueryStringFields** | **[]string** |  | 
**EnableQueryStringSort** | **bool** |  | 
**CacheByCookies** | **string** |  | 
**CookieNames** | **[]string** |  | 
**AdaptiveDeliveryAction** | **string** |  | 
**DeviceGroup** | **[]int32** |  | 
**EnableCachingForPost** | **bool** |  | 
**L2CachingEnabled** | **bool** |  | 
**IsSliceConfigurationEnabled** | Pointer to **bool** |  | [optional] 
**IsSliceEdgeCachingEnabled** | Pointer to **bool** |  | [optional] 
**IsSliceL2CachingEnabled** | Pointer to **bool** |  | [optional] 
**SliceConfigurationRange** | Pointer to **int64** |  | [optional] 
**EnableCachingForOptions** | Pointer to **bool** |  | [optional] 
**EnableStaleCache** | Pointer to **bool** |  | [optional] 
**L2Region** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationCacheCreateResults

`func NewApplicationCacheCreateResults(id int64, name string, browserCacheSettings string, browserCacheSettingsMaximumTtl int64, cdnCacheSettings string, cdnCacheSettingsMaximumTtl int64, cacheByQueryString string, queryStringFields []string, enableQueryStringSort bool, cacheByCookies string, cookieNames []string, adaptiveDeliveryAction string, deviceGroup []int32, enableCachingForPost bool, l2CachingEnabled bool, ) *ApplicationCacheCreateResults`

NewApplicationCacheCreateResults instantiates a new ApplicationCacheCreateResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCacheCreateResultsWithDefaults

`func NewApplicationCacheCreateResultsWithDefaults() *ApplicationCacheCreateResults`

NewApplicationCacheCreateResultsWithDefaults instantiates a new ApplicationCacheCreateResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationCacheCreateResults) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationCacheCreateResults) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationCacheCreateResults) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationCacheCreateResults) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCacheCreateResults) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCacheCreateResults) SetName(v string)`

SetName sets Name field to given value.


### GetBrowserCacheSettings

`func (o *ApplicationCacheCreateResults) GetBrowserCacheSettings() string`

GetBrowserCacheSettings returns the BrowserCacheSettings field if non-nil, zero value otherwise.

### GetBrowserCacheSettingsOk

`func (o *ApplicationCacheCreateResults) GetBrowserCacheSettingsOk() (*string, bool)`

GetBrowserCacheSettingsOk returns a tuple with the BrowserCacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCacheSettings

`func (o *ApplicationCacheCreateResults) SetBrowserCacheSettings(v string)`

SetBrowserCacheSettings sets BrowserCacheSettings field to given value.


### GetBrowserCacheSettingsMaximumTtl

`func (o *ApplicationCacheCreateResults) GetBrowserCacheSettingsMaximumTtl() int64`

GetBrowserCacheSettingsMaximumTtl returns the BrowserCacheSettingsMaximumTtl field if non-nil, zero value otherwise.

### GetBrowserCacheSettingsMaximumTtlOk

`func (o *ApplicationCacheCreateResults) GetBrowserCacheSettingsMaximumTtlOk() (*int64, bool)`

GetBrowserCacheSettingsMaximumTtlOk returns a tuple with the BrowserCacheSettingsMaximumTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCacheSettingsMaximumTtl

`func (o *ApplicationCacheCreateResults) SetBrowserCacheSettingsMaximumTtl(v int64)`

SetBrowserCacheSettingsMaximumTtl sets BrowserCacheSettingsMaximumTtl field to given value.


### GetCdnCacheSettings

`func (o *ApplicationCacheCreateResults) GetCdnCacheSettings() string`

GetCdnCacheSettings returns the CdnCacheSettings field if non-nil, zero value otherwise.

### GetCdnCacheSettingsOk

`func (o *ApplicationCacheCreateResults) GetCdnCacheSettingsOk() (*string, bool)`

GetCdnCacheSettingsOk returns a tuple with the CdnCacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnCacheSettings

`func (o *ApplicationCacheCreateResults) SetCdnCacheSettings(v string)`

SetCdnCacheSettings sets CdnCacheSettings field to given value.


### GetCdnCacheSettingsMaximumTtl

`func (o *ApplicationCacheCreateResults) GetCdnCacheSettingsMaximumTtl() int64`

GetCdnCacheSettingsMaximumTtl returns the CdnCacheSettingsMaximumTtl field if non-nil, zero value otherwise.

### GetCdnCacheSettingsMaximumTtlOk

`func (o *ApplicationCacheCreateResults) GetCdnCacheSettingsMaximumTtlOk() (*int64, bool)`

GetCdnCacheSettingsMaximumTtlOk returns a tuple with the CdnCacheSettingsMaximumTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnCacheSettingsMaximumTtl

`func (o *ApplicationCacheCreateResults) SetCdnCacheSettingsMaximumTtl(v int64)`

SetCdnCacheSettingsMaximumTtl sets CdnCacheSettingsMaximumTtl field to given value.


### GetCacheByQueryString

`func (o *ApplicationCacheCreateResults) GetCacheByQueryString() string`

GetCacheByQueryString returns the CacheByQueryString field if non-nil, zero value otherwise.

### GetCacheByQueryStringOk

`func (o *ApplicationCacheCreateResults) GetCacheByQueryStringOk() (*string, bool)`

GetCacheByQueryStringOk returns a tuple with the CacheByQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByQueryString

`func (o *ApplicationCacheCreateResults) SetCacheByQueryString(v string)`

SetCacheByQueryString sets CacheByQueryString field to given value.


### GetQueryStringFields

`func (o *ApplicationCacheCreateResults) GetQueryStringFields() []string`

GetQueryStringFields returns the QueryStringFields field if non-nil, zero value otherwise.

### GetQueryStringFieldsOk

`func (o *ApplicationCacheCreateResults) GetQueryStringFieldsOk() (*[]string, bool)`

GetQueryStringFieldsOk returns a tuple with the QueryStringFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStringFields

`func (o *ApplicationCacheCreateResults) SetQueryStringFields(v []string)`

SetQueryStringFields sets QueryStringFields field to given value.


### GetEnableQueryStringSort

`func (o *ApplicationCacheCreateResults) GetEnableQueryStringSort() bool`

GetEnableQueryStringSort returns the EnableQueryStringSort field if non-nil, zero value otherwise.

### GetEnableQueryStringSortOk

`func (o *ApplicationCacheCreateResults) GetEnableQueryStringSortOk() (*bool, bool)`

GetEnableQueryStringSortOk returns a tuple with the EnableQueryStringSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQueryStringSort

`func (o *ApplicationCacheCreateResults) SetEnableQueryStringSort(v bool)`

SetEnableQueryStringSort sets EnableQueryStringSort field to given value.


### GetCacheByCookies

`func (o *ApplicationCacheCreateResults) GetCacheByCookies() string`

GetCacheByCookies returns the CacheByCookies field if non-nil, zero value otherwise.

### GetCacheByCookiesOk

`func (o *ApplicationCacheCreateResults) GetCacheByCookiesOk() (*string, bool)`

GetCacheByCookiesOk returns a tuple with the CacheByCookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByCookies

`func (o *ApplicationCacheCreateResults) SetCacheByCookies(v string)`

SetCacheByCookies sets CacheByCookies field to given value.


### GetCookieNames

`func (o *ApplicationCacheCreateResults) GetCookieNames() []string`

GetCookieNames returns the CookieNames field if non-nil, zero value otherwise.

### GetCookieNamesOk

`func (o *ApplicationCacheCreateResults) GetCookieNamesOk() (*[]string, bool)`

GetCookieNamesOk returns a tuple with the CookieNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieNames

`func (o *ApplicationCacheCreateResults) SetCookieNames(v []string)`

SetCookieNames sets CookieNames field to given value.


### GetAdaptiveDeliveryAction

`func (o *ApplicationCacheCreateResults) GetAdaptiveDeliveryAction() string`

GetAdaptiveDeliveryAction returns the AdaptiveDeliveryAction field if non-nil, zero value otherwise.

### GetAdaptiveDeliveryActionOk

`func (o *ApplicationCacheCreateResults) GetAdaptiveDeliveryActionOk() (*string, bool)`

GetAdaptiveDeliveryActionOk returns a tuple with the AdaptiveDeliveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptiveDeliveryAction

`func (o *ApplicationCacheCreateResults) SetAdaptiveDeliveryAction(v string)`

SetAdaptiveDeliveryAction sets AdaptiveDeliveryAction field to given value.


### GetDeviceGroup

`func (o *ApplicationCacheCreateResults) GetDeviceGroup() []int32`

GetDeviceGroup returns the DeviceGroup field if non-nil, zero value otherwise.

### GetDeviceGroupOk

`func (o *ApplicationCacheCreateResults) GetDeviceGroupOk() (*[]int32, bool)`

GetDeviceGroupOk returns a tuple with the DeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroup

`func (o *ApplicationCacheCreateResults) SetDeviceGroup(v []int32)`

SetDeviceGroup sets DeviceGroup field to given value.


### GetEnableCachingForPost

`func (o *ApplicationCacheCreateResults) GetEnableCachingForPost() bool`

GetEnableCachingForPost returns the EnableCachingForPost field if non-nil, zero value otherwise.

### GetEnableCachingForPostOk

`func (o *ApplicationCacheCreateResults) GetEnableCachingForPostOk() (*bool, bool)`

GetEnableCachingForPostOk returns a tuple with the EnableCachingForPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCachingForPost

`func (o *ApplicationCacheCreateResults) SetEnableCachingForPost(v bool)`

SetEnableCachingForPost sets EnableCachingForPost field to given value.


### GetL2CachingEnabled

`func (o *ApplicationCacheCreateResults) GetL2CachingEnabled() bool`

GetL2CachingEnabled returns the L2CachingEnabled field if non-nil, zero value otherwise.

### GetL2CachingEnabledOk

`func (o *ApplicationCacheCreateResults) GetL2CachingEnabledOk() (*bool, bool)`

GetL2CachingEnabledOk returns a tuple with the L2CachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2CachingEnabled

`func (o *ApplicationCacheCreateResults) SetL2CachingEnabled(v bool)`

SetL2CachingEnabled sets L2CachingEnabled field to given value.


### GetIsSliceConfigurationEnabled

`func (o *ApplicationCacheCreateResults) GetIsSliceConfigurationEnabled() bool`

GetIsSliceConfigurationEnabled returns the IsSliceConfigurationEnabled field if non-nil, zero value otherwise.

### GetIsSliceConfigurationEnabledOk

`func (o *ApplicationCacheCreateResults) GetIsSliceConfigurationEnabledOk() (*bool, bool)`

GetIsSliceConfigurationEnabledOk returns a tuple with the IsSliceConfigurationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSliceConfigurationEnabled

`func (o *ApplicationCacheCreateResults) SetIsSliceConfigurationEnabled(v bool)`

SetIsSliceConfigurationEnabled sets IsSliceConfigurationEnabled field to given value.

### HasIsSliceConfigurationEnabled

`func (o *ApplicationCacheCreateResults) HasIsSliceConfigurationEnabled() bool`

HasIsSliceConfigurationEnabled returns a boolean if a field has been set.

### GetIsSliceEdgeCachingEnabled

`func (o *ApplicationCacheCreateResults) GetIsSliceEdgeCachingEnabled() bool`

GetIsSliceEdgeCachingEnabled returns the IsSliceEdgeCachingEnabled field if non-nil, zero value otherwise.

### GetIsSliceEdgeCachingEnabledOk

`func (o *ApplicationCacheCreateResults) GetIsSliceEdgeCachingEnabledOk() (*bool, bool)`

GetIsSliceEdgeCachingEnabledOk returns a tuple with the IsSliceEdgeCachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSliceEdgeCachingEnabled

`func (o *ApplicationCacheCreateResults) SetIsSliceEdgeCachingEnabled(v bool)`

SetIsSliceEdgeCachingEnabled sets IsSliceEdgeCachingEnabled field to given value.

### HasIsSliceEdgeCachingEnabled

`func (o *ApplicationCacheCreateResults) HasIsSliceEdgeCachingEnabled() bool`

HasIsSliceEdgeCachingEnabled returns a boolean if a field has been set.

### GetIsSliceL2CachingEnabled

`func (o *ApplicationCacheCreateResults) GetIsSliceL2CachingEnabled() bool`

GetIsSliceL2CachingEnabled returns the IsSliceL2CachingEnabled field if non-nil, zero value otherwise.

### GetIsSliceL2CachingEnabledOk

`func (o *ApplicationCacheCreateResults) GetIsSliceL2CachingEnabledOk() (*bool, bool)`

GetIsSliceL2CachingEnabledOk returns a tuple with the IsSliceL2CachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSliceL2CachingEnabled

`func (o *ApplicationCacheCreateResults) SetIsSliceL2CachingEnabled(v bool)`

SetIsSliceL2CachingEnabled sets IsSliceL2CachingEnabled field to given value.

### HasIsSliceL2CachingEnabled

`func (o *ApplicationCacheCreateResults) HasIsSliceL2CachingEnabled() bool`

HasIsSliceL2CachingEnabled returns a boolean if a field has been set.

### GetSliceConfigurationRange

`func (o *ApplicationCacheCreateResults) GetSliceConfigurationRange() int64`

GetSliceConfigurationRange returns the SliceConfigurationRange field if non-nil, zero value otherwise.

### GetSliceConfigurationRangeOk

`func (o *ApplicationCacheCreateResults) GetSliceConfigurationRangeOk() (*int64, bool)`

GetSliceConfigurationRangeOk returns a tuple with the SliceConfigurationRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceConfigurationRange

`func (o *ApplicationCacheCreateResults) SetSliceConfigurationRange(v int64)`

SetSliceConfigurationRange sets SliceConfigurationRange field to given value.

### HasSliceConfigurationRange

`func (o *ApplicationCacheCreateResults) HasSliceConfigurationRange() bool`

HasSliceConfigurationRange returns a boolean if a field has been set.

### GetEnableCachingForOptions

`func (o *ApplicationCacheCreateResults) GetEnableCachingForOptions() bool`

GetEnableCachingForOptions returns the EnableCachingForOptions field if non-nil, zero value otherwise.

### GetEnableCachingForOptionsOk

`func (o *ApplicationCacheCreateResults) GetEnableCachingForOptionsOk() (*bool, bool)`

GetEnableCachingForOptionsOk returns a tuple with the EnableCachingForOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCachingForOptions

`func (o *ApplicationCacheCreateResults) SetEnableCachingForOptions(v bool)`

SetEnableCachingForOptions sets EnableCachingForOptions field to given value.

### HasEnableCachingForOptions

`func (o *ApplicationCacheCreateResults) HasEnableCachingForOptions() bool`

HasEnableCachingForOptions returns a boolean if a field has been set.

### GetEnableStaleCache

`func (o *ApplicationCacheCreateResults) GetEnableStaleCache() bool`

GetEnableStaleCache returns the EnableStaleCache field if non-nil, zero value otherwise.

### GetEnableStaleCacheOk

`func (o *ApplicationCacheCreateResults) GetEnableStaleCacheOk() (*bool, bool)`

GetEnableStaleCacheOk returns a tuple with the EnableStaleCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStaleCache

`func (o *ApplicationCacheCreateResults) SetEnableStaleCache(v bool)`

SetEnableStaleCache sets EnableStaleCache field to given value.

### HasEnableStaleCache

`func (o *ApplicationCacheCreateResults) HasEnableStaleCache() bool`

HasEnableStaleCache returns a boolean if a field has been set.

### GetL2Region

`func (o *ApplicationCacheCreateResults) GetL2Region() string`

GetL2Region returns the L2Region field if non-nil, zero value otherwise.

### GetL2RegionOk

`func (o *ApplicationCacheCreateResults) GetL2RegionOk() (*string, bool)`

GetL2RegionOk returns a tuple with the L2Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2Region

`func (o *ApplicationCacheCreateResults) SetL2Region(v string)`

SetL2Region sets L2Region field to given value.

### HasL2Region

`func (o *ApplicationCacheCreateResults) HasL2Region() bool`

HasL2Region returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


