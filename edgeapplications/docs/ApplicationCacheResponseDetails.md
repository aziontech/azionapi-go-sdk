# ApplicationCacheResponseDetails

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
**AdaptiveDeliveryAction** | Pointer to **string** |  | [optional] 
**DeviceGroup** | Pointer to **[]int32** |  | [optional] 
**EnableCachingForPost** | **bool** |  | 
**EnableCachingForOptions** | Pointer to **bool** |  | [optional] 
**L2CachingEnabled** | **bool** |  | 
**IsSliceConfigurationEnabled** | Pointer to **bool** |  | [optional] 
**IsSliceEdgeCachingEnabled** | Pointer to **bool** |  | [optional] 
**IsSliceL2CachingEnabled** | Pointer to **bool** |  | [optional] 
**SliceConfigurationRange** | Pointer to **int64** |  | [optional] 
**EnableStaleCache** | Pointer to **bool** |  | [optional] 
**L2Region** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApplicationCacheResponseDetails

`func NewApplicationCacheResponseDetails(id int64, name string, browserCacheSettings string, browserCacheSettingsMaximumTtl int64, cdnCacheSettings string, cdnCacheSettingsMaximumTtl int64, cacheByQueryString string, queryStringFields []string, enableQueryStringSort bool, cacheByCookies string, cookieNames []string, enableCachingForPost bool, l2CachingEnabled bool, ) *ApplicationCacheResponseDetails`

NewApplicationCacheResponseDetails instantiates a new ApplicationCacheResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCacheResponseDetailsWithDefaults

`func NewApplicationCacheResponseDetailsWithDefaults() *ApplicationCacheResponseDetails`

NewApplicationCacheResponseDetailsWithDefaults instantiates a new ApplicationCacheResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationCacheResponseDetails) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationCacheResponseDetails) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationCacheResponseDetails) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationCacheResponseDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCacheResponseDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCacheResponseDetails) SetName(v string)`

SetName sets Name field to given value.


### GetBrowserCacheSettings

`func (o *ApplicationCacheResponseDetails) GetBrowserCacheSettings() string`

GetBrowserCacheSettings returns the BrowserCacheSettings field if non-nil, zero value otherwise.

### GetBrowserCacheSettingsOk

`func (o *ApplicationCacheResponseDetails) GetBrowserCacheSettingsOk() (*string, bool)`

GetBrowserCacheSettingsOk returns a tuple with the BrowserCacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCacheSettings

`func (o *ApplicationCacheResponseDetails) SetBrowserCacheSettings(v string)`

SetBrowserCacheSettings sets BrowserCacheSettings field to given value.


### GetBrowserCacheSettingsMaximumTtl

`func (o *ApplicationCacheResponseDetails) GetBrowserCacheSettingsMaximumTtl() int64`

GetBrowserCacheSettingsMaximumTtl returns the BrowserCacheSettingsMaximumTtl field if non-nil, zero value otherwise.

### GetBrowserCacheSettingsMaximumTtlOk

`func (o *ApplicationCacheResponseDetails) GetBrowserCacheSettingsMaximumTtlOk() (*int64, bool)`

GetBrowserCacheSettingsMaximumTtlOk returns a tuple with the BrowserCacheSettingsMaximumTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCacheSettingsMaximumTtl

`func (o *ApplicationCacheResponseDetails) SetBrowserCacheSettingsMaximumTtl(v int64)`

SetBrowserCacheSettingsMaximumTtl sets BrowserCacheSettingsMaximumTtl field to given value.


### GetCdnCacheSettings

`func (o *ApplicationCacheResponseDetails) GetCdnCacheSettings() string`

GetCdnCacheSettings returns the CdnCacheSettings field if non-nil, zero value otherwise.

### GetCdnCacheSettingsOk

`func (o *ApplicationCacheResponseDetails) GetCdnCacheSettingsOk() (*string, bool)`

GetCdnCacheSettingsOk returns a tuple with the CdnCacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnCacheSettings

`func (o *ApplicationCacheResponseDetails) SetCdnCacheSettings(v string)`

SetCdnCacheSettings sets CdnCacheSettings field to given value.


### GetCdnCacheSettingsMaximumTtl

`func (o *ApplicationCacheResponseDetails) GetCdnCacheSettingsMaximumTtl() int64`

GetCdnCacheSettingsMaximumTtl returns the CdnCacheSettingsMaximumTtl field if non-nil, zero value otherwise.

### GetCdnCacheSettingsMaximumTtlOk

`func (o *ApplicationCacheResponseDetails) GetCdnCacheSettingsMaximumTtlOk() (*int64, bool)`

GetCdnCacheSettingsMaximumTtlOk returns a tuple with the CdnCacheSettingsMaximumTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnCacheSettingsMaximumTtl

`func (o *ApplicationCacheResponseDetails) SetCdnCacheSettingsMaximumTtl(v int64)`

SetCdnCacheSettingsMaximumTtl sets CdnCacheSettingsMaximumTtl field to given value.


### GetCacheByQueryString

`func (o *ApplicationCacheResponseDetails) GetCacheByQueryString() string`

GetCacheByQueryString returns the CacheByQueryString field if non-nil, zero value otherwise.

### GetCacheByQueryStringOk

`func (o *ApplicationCacheResponseDetails) GetCacheByQueryStringOk() (*string, bool)`

GetCacheByQueryStringOk returns a tuple with the CacheByQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByQueryString

`func (o *ApplicationCacheResponseDetails) SetCacheByQueryString(v string)`

SetCacheByQueryString sets CacheByQueryString field to given value.


### GetQueryStringFields

`func (o *ApplicationCacheResponseDetails) GetQueryStringFields() []string`

GetQueryStringFields returns the QueryStringFields field if non-nil, zero value otherwise.

### GetQueryStringFieldsOk

`func (o *ApplicationCacheResponseDetails) GetQueryStringFieldsOk() (*[]string, bool)`

GetQueryStringFieldsOk returns a tuple with the QueryStringFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStringFields

`func (o *ApplicationCacheResponseDetails) SetQueryStringFields(v []string)`

SetQueryStringFields sets QueryStringFields field to given value.


### SetQueryStringFieldsNil

`func (o *ApplicationCacheResponseDetails) SetQueryStringFieldsNil(b bool)`

 SetQueryStringFieldsNil sets the value for QueryStringFields to be an explicit nil

### UnsetQueryStringFields
`func (o *ApplicationCacheResponseDetails) UnsetQueryStringFields()`

UnsetQueryStringFields ensures that no value is present for QueryStringFields, not even an explicit nil
### GetEnableQueryStringSort

`func (o *ApplicationCacheResponseDetails) GetEnableQueryStringSort() bool`

GetEnableQueryStringSort returns the EnableQueryStringSort field if non-nil, zero value otherwise.

### GetEnableQueryStringSortOk

`func (o *ApplicationCacheResponseDetails) GetEnableQueryStringSortOk() (*bool, bool)`

GetEnableQueryStringSortOk returns a tuple with the EnableQueryStringSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQueryStringSort

`func (o *ApplicationCacheResponseDetails) SetEnableQueryStringSort(v bool)`

SetEnableQueryStringSort sets EnableQueryStringSort field to given value.


### GetCacheByCookies

`func (o *ApplicationCacheResponseDetails) GetCacheByCookies() string`

GetCacheByCookies returns the CacheByCookies field if non-nil, zero value otherwise.

### GetCacheByCookiesOk

`func (o *ApplicationCacheResponseDetails) GetCacheByCookiesOk() (*string, bool)`

GetCacheByCookiesOk returns a tuple with the CacheByCookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByCookies

`func (o *ApplicationCacheResponseDetails) SetCacheByCookies(v string)`

SetCacheByCookies sets CacheByCookies field to given value.


### GetCookieNames

`func (o *ApplicationCacheResponseDetails) GetCookieNames() []string`

GetCookieNames returns the CookieNames field if non-nil, zero value otherwise.

### GetCookieNamesOk

`func (o *ApplicationCacheResponseDetails) GetCookieNamesOk() (*[]string, bool)`

GetCookieNamesOk returns a tuple with the CookieNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieNames

`func (o *ApplicationCacheResponseDetails) SetCookieNames(v []string)`

SetCookieNames sets CookieNames field to given value.


### SetCookieNamesNil

`func (o *ApplicationCacheResponseDetails) SetCookieNamesNil(b bool)`

 SetCookieNamesNil sets the value for CookieNames to be an explicit nil

### UnsetCookieNames
`func (o *ApplicationCacheResponseDetails) UnsetCookieNames()`

UnsetCookieNames ensures that no value is present for CookieNames, not even an explicit nil
### GetAdaptiveDeliveryAction

`func (o *ApplicationCacheResponseDetails) GetAdaptiveDeliveryAction() string`

GetAdaptiveDeliveryAction returns the AdaptiveDeliveryAction field if non-nil, zero value otherwise.

### GetAdaptiveDeliveryActionOk

`func (o *ApplicationCacheResponseDetails) GetAdaptiveDeliveryActionOk() (*string, bool)`

GetAdaptiveDeliveryActionOk returns a tuple with the AdaptiveDeliveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptiveDeliveryAction

`func (o *ApplicationCacheResponseDetails) SetAdaptiveDeliveryAction(v string)`

SetAdaptiveDeliveryAction sets AdaptiveDeliveryAction field to given value.

### HasAdaptiveDeliveryAction

`func (o *ApplicationCacheResponseDetails) HasAdaptiveDeliveryAction() bool`

HasAdaptiveDeliveryAction returns a boolean if a field has been set.

### GetDeviceGroup

`func (o *ApplicationCacheResponseDetails) GetDeviceGroup() []int32`

GetDeviceGroup returns the DeviceGroup field if non-nil, zero value otherwise.

### GetDeviceGroupOk

`func (o *ApplicationCacheResponseDetails) GetDeviceGroupOk() (*[]int32, bool)`

GetDeviceGroupOk returns a tuple with the DeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceGroup

`func (o *ApplicationCacheResponseDetails) SetDeviceGroup(v []int32)`

SetDeviceGroup sets DeviceGroup field to given value.

### HasDeviceGroup

`func (o *ApplicationCacheResponseDetails) HasDeviceGroup() bool`

HasDeviceGroup returns a boolean if a field has been set.

### GetEnableCachingForPost

`func (o *ApplicationCacheResponseDetails) GetEnableCachingForPost() bool`

GetEnableCachingForPost returns the EnableCachingForPost field if non-nil, zero value otherwise.

### GetEnableCachingForPostOk

`func (o *ApplicationCacheResponseDetails) GetEnableCachingForPostOk() (*bool, bool)`

GetEnableCachingForPostOk returns a tuple with the EnableCachingForPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCachingForPost

`func (o *ApplicationCacheResponseDetails) SetEnableCachingForPost(v bool)`

SetEnableCachingForPost sets EnableCachingForPost field to given value.


### GetEnableCachingForOptions

`func (o *ApplicationCacheResponseDetails) GetEnableCachingForOptions() bool`

GetEnableCachingForOptions returns the EnableCachingForOptions field if non-nil, zero value otherwise.

### GetEnableCachingForOptionsOk

`func (o *ApplicationCacheResponseDetails) GetEnableCachingForOptionsOk() (*bool, bool)`

GetEnableCachingForOptionsOk returns a tuple with the EnableCachingForOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCachingForOptions

`func (o *ApplicationCacheResponseDetails) SetEnableCachingForOptions(v bool)`

SetEnableCachingForOptions sets EnableCachingForOptions field to given value.

### HasEnableCachingForOptions

`func (o *ApplicationCacheResponseDetails) HasEnableCachingForOptions() bool`

HasEnableCachingForOptions returns a boolean if a field has been set.

### GetL2CachingEnabled

`func (o *ApplicationCacheResponseDetails) GetL2CachingEnabled() bool`

GetL2CachingEnabled returns the L2CachingEnabled field if non-nil, zero value otherwise.

### GetL2CachingEnabledOk

`func (o *ApplicationCacheResponseDetails) GetL2CachingEnabledOk() (*bool, bool)`

GetL2CachingEnabledOk returns a tuple with the L2CachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2CachingEnabled

`func (o *ApplicationCacheResponseDetails) SetL2CachingEnabled(v bool)`

SetL2CachingEnabled sets L2CachingEnabled field to given value.


### GetIsSliceConfigurationEnabled

`func (o *ApplicationCacheResponseDetails) GetIsSliceConfigurationEnabled() bool`

GetIsSliceConfigurationEnabled returns the IsSliceConfigurationEnabled field if non-nil, zero value otherwise.

### GetIsSliceConfigurationEnabledOk

`func (o *ApplicationCacheResponseDetails) GetIsSliceConfigurationEnabledOk() (*bool, bool)`

GetIsSliceConfigurationEnabledOk returns a tuple with the IsSliceConfigurationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSliceConfigurationEnabled

`func (o *ApplicationCacheResponseDetails) SetIsSliceConfigurationEnabled(v bool)`

SetIsSliceConfigurationEnabled sets IsSliceConfigurationEnabled field to given value.

### HasIsSliceConfigurationEnabled

`func (o *ApplicationCacheResponseDetails) HasIsSliceConfigurationEnabled() bool`

HasIsSliceConfigurationEnabled returns a boolean if a field has been set.

### GetIsSliceEdgeCachingEnabled

`func (o *ApplicationCacheResponseDetails) GetIsSliceEdgeCachingEnabled() bool`

GetIsSliceEdgeCachingEnabled returns the IsSliceEdgeCachingEnabled field if non-nil, zero value otherwise.

### GetIsSliceEdgeCachingEnabledOk

`func (o *ApplicationCacheResponseDetails) GetIsSliceEdgeCachingEnabledOk() (*bool, bool)`

GetIsSliceEdgeCachingEnabledOk returns a tuple with the IsSliceEdgeCachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSliceEdgeCachingEnabled

`func (o *ApplicationCacheResponseDetails) SetIsSliceEdgeCachingEnabled(v bool)`

SetIsSliceEdgeCachingEnabled sets IsSliceEdgeCachingEnabled field to given value.

### HasIsSliceEdgeCachingEnabled

`func (o *ApplicationCacheResponseDetails) HasIsSliceEdgeCachingEnabled() bool`

HasIsSliceEdgeCachingEnabled returns a boolean if a field has been set.

### GetIsSliceL2CachingEnabled

`func (o *ApplicationCacheResponseDetails) GetIsSliceL2CachingEnabled() bool`

GetIsSliceL2CachingEnabled returns the IsSliceL2CachingEnabled field if non-nil, zero value otherwise.

### GetIsSliceL2CachingEnabledOk

`func (o *ApplicationCacheResponseDetails) GetIsSliceL2CachingEnabledOk() (*bool, bool)`

GetIsSliceL2CachingEnabledOk returns a tuple with the IsSliceL2CachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSliceL2CachingEnabled

`func (o *ApplicationCacheResponseDetails) SetIsSliceL2CachingEnabled(v bool)`

SetIsSliceL2CachingEnabled sets IsSliceL2CachingEnabled field to given value.

### HasIsSliceL2CachingEnabled

`func (o *ApplicationCacheResponseDetails) HasIsSliceL2CachingEnabled() bool`

HasIsSliceL2CachingEnabled returns a boolean if a field has been set.

### GetSliceConfigurationRange

`func (o *ApplicationCacheResponseDetails) GetSliceConfigurationRange() int64`

GetSliceConfigurationRange returns the SliceConfigurationRange field if non-nil, zero value otherwise.

### GetSliceConfigurationRangeOk

`func (o *ApplicationCacheResponseDetails) GetSliceConfigurationRangeOk() (*int64, bool)`

GetSliceConfigurationRangeOk returns a tuple with the SliceConfigurationRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceConfigurationRange

`func (o *ApplicationCacheResponseDetails) SetSliceConfigurationRange(v int64)`

SetSliceConfigurationRange sets SliceConfigurationRange field to given value.

### HasSliceConfigurationRange

`func (o *ApplicationCacheResponseDetails) HasSliceConfigurationRange() bool`

HasSliceConfigurationRange returns a boolean if a field has been set.

### GetEnableStaleCache

`func (o *ApplicationCacheResponseDetails) GetEnableStaleCache() bool`

GetEnableStaleCache returns the EnableStaleCache field if non-nil, zero value otherwise.

### GetEnableStaleCacheOk

`func (o *ApplicationCacheResponseDetails) GetEnableStaleCacheOk() (*bool, bool)`

GetEnableStaleCacheOk returns a tuple with the EnableStaleCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStaleCache

`func (o *ApplicationCacheResponseDetails) SetEnableStaleCache(v bool)`

SetEnableStaleCache sets EnableStaleCache field to given value.

### HasEnableStaleCache

`func (o *ApplicationCacheResponseDetails) HasEnableStaleCache() bool`

HasEnableStaleCache returns a boolean if a field has been set.

### GetL2Region

`func (o *ApplicationCacheResponseDetails) GetL2Region() string`

GetL2Region returns the L2Region field if non-nil, zero value otherwise.

### GetL2RegionOk

`func (o *ApplicationCacheResponseDetails) GetL2RegionOk() (*string, bool)`

GetL2RegionOk returns a tuple with the L2Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2Region

`func (o *ApplicationCacheResponseDetails) SetL2Region(v string)`

SetL2Region sets L2Region field to given value.

### HasL2Region

`func (o *ApplicationCacheResponseDetails) HasL2Region() bool`

HasL2Region returns a boolean if a field has been set.

### SetL2RegionNil

`func (o *ApplicationCacheResponseDetails) SetL2RegionNil(b bool)`

 SetL2RegionNil sets the value for L2Region to be an explicit nil

### UnsetL2Region
`func (o *ApplicationCacheResponseDetails) UnsetL2Region()`

UnsetL2Region ensures that no value is present for L2Region, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


