# ApplicationCachePutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**BrowserCacheSettings** | Pointer to **string** |  | [optional] 
**BrowserCacheSettingsMaximumTtl** | Pointer to **int64** |  | [optional] 
**CdnCacheSettings** | Pointer to **string** |  | [optional] 
**CdnCacheSettingsMaximumTtl** | Pointer to **int64** |  | [optional] 
**CacheByQueryString** | Pointer to **string** |  | [optional] 
**QueryStringFields** | Pointer to **[]string** |  | [optional] 
**EnableQueryStringSort** | Pointer to **bool** |  | [optional] 
**CacheByCookies** | Pointer to **string** |  | [optional] 
**CookieNames** | Pointer to **[]string** |  | [optional] 
**EnableCachingForPost** | Pointer to **bool** |  | [optional] 
**L2CachingEnabled** | Pointer to **bool** |  | [optional] 
**IsSliceConfigurationEnabled** | Pointer to **bool** |  | [optional] 
**IsSliceEdgeCachingEnabled** | Pointer to **bool** |  | [optional] 
**IsSliceL2CachingEnabled** | Pointer to **bool** |  | [optional] 
**SliceConfigurationRange** | Pointer to **int64** |  | [optional] 

## Methods

### NewApplicationCachePutRequest

`func NewApplicationCachePutRequest(name string, ) *ApplicationCachePutRequest`

NewApplicationCachePutRequest instantiates a new ApplicationCachePutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCachePutRequestWithDefaults

`func NewApplicationCachePutRequestWithDefaults() *ApplicationCachePutRequest`

NewApplicationCachePutRequestWithDefaults instantiates a new ApplicationCachePutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationCachePutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCachePutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCachePutRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBrowserCacheSettings

`func (o *ApplicationCachePutRequest) GetBrowserCacheSettings() string`

GetBrowserCacheSettings returns the BrowserCacheSettings field if non-nil, zero value otherwise.

### GetBrowserCacheSettingsOk

`func (o *ApplicationCachePutRequest) GetBrowserCacheSettingsOk() (*string, bool)`

GetBrowserCacheSettingsOk returns a tuple with the BrowserCacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCacheSettings

`func (o *ApplicationCachePutRequest) SetBrowserCacheSettings(v string)`

SetBrowserCacheSettings sets BrowserCacheSettings field to given value.

### HasBrowserCacheSettings

`func (o *ApplicationCachePutRequest) HasBrowserCacheSettings() bool`

HasBrowserCacheSettings returns a boolean if a field has been set.

### GetBrowserCacheSettingsMaximumTtl

`func (o *ApplicationCachePutRequest) GetBrowserCacheSettingsMaximumTtl() int64`

GetBrowserCacheSettingsMaximumTtl returns the BrowserCacheSettingsMaximumTtl field if non-nil, zero value otherwise.

### GetBrowserCacheSettingsMaximumTtlOk

`func (o *ApplicationCachePutRequest) GetBrowserCacheSettingsMaximumTtlOk() (*int64, bool)`

GetBrowserCacheSettingsMaximumTtlOk returns a tuple with the BrowserCacheSettingsMaximumTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCacheSettingsMaximumTtl

`func (o *ApplicationCachePutRequest) SetBrowserCacheSettingsMaximumTtl(v int64)`

SetBrowserCacheSettingsMaximumTtl sets BrowserCacheSettingsMaximumTtl field to given value.

### HasBrowserCacheSettingsMaximumTtl

`func (o *ApplicationCachePutRequest) HasBrowserCacheSettingsMaximumTtl() bool`

HasBrowserCacheSettingsMaximumTtl returns a boolean if a field has been set.

### GetCdnCacheSettings

`func (o *ApplicationCachePutRequest) GetCdnCacheSettings() string`

GetCdnCacheSettings returns the CdnCacheSettings field if non-nil, zero value otherwise.

### GetCdnCacheSettingsOk

`func (o *ApplicationCachePutRequest) GetCdnCacheSettingsOk() (*string, bool)`

GetCdnCacheSettingsOk returns a tuple with the CdnCacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnCacheSettings

`func (o *ApplicationCachePutRequest) SetCdnCacheSettings(v string)`

SetCdnCacheSettings sets CdnCacheSettings field to given value.

### HasCdnCacheSettings

`func (o *ApplicationCachePutRequest) HasCdnCacheSettings() bool`

HasCdnCacheSettings returns a boolean if a field has been set.

### GetCdnCacheSettingsMaximumTtl

`func (o *ApplicationCachePutRequest) GetCdnCacheSettingsMaximumTtl() int64`

GetCdnCacheSettingsMaximumTtl returns the CdnCacheSettingsMaximumTtl field if non-nil, zero value otherwise.

### GetCdnCacheSettingsMaximumTtlOk

`func (o *ApplicationCachePutRequest) GetCdnCacheSettingsMaximumTtlOk() (*int64, bool)`

GetCdnCacheSettingsMaximumTtlOk returns a tuple with the CdnCacheSettingsMaximumTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnCacheSettingsMaximumTtl

`func (o *ApplicationCachePutRequest) SetCdnCacheSettingsMaximumTtl(v int64)`

SetCdnCacheSettingsMaximumTtl sets CdnCacheSettingsMaximumTtl field to given value.

### HasCdnCacheSettingsMaximumTtl

`func (o *ApplicationCachePutRequest) HasCdnCacheSettingsMaximumTtl() bool`

HasCdnCacheSettingsMaximumTtl returns a boolean if a field has been set.

### GetCacheByQueryString

`func (o *ApplicationCachePutRequest) GetCacheByQueryString() string`

GetCacheByQueryString returns the CacheByQueryString field if non-nil, zero value otherwise.

### GetCacheByQueryStringOk

`func (o *ApplicationCachePutRequest) GetCacheByQueryStringOk() (*string, bool)`

GetCacheByQueryStringOk returns a tuple with the CacheByQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByQueryString

`func (o *ApplicationCachePutRequest) SetCacheByQueryString(v string)`

SetCacheByQueryString sets CacheByQueryString field to given value.

### HasCacheByQueryString

`func (o *ApplicationCachePutRequest) HasCacheByQueryString() bool`

HasCacheByQueryString returns a boolean if a field has been set.

### GetQueryStringFields

`func (o *ApplicationCachePutRequest) GetQueryStringFields() []string`

GetQueryStringFields returns the QueryStringFields field if non-nil, zero value otherwise.

### GetQueryStringFieldsOk

`func (o *ApplicationCachePutRequest) GetQueryStringFieldsOk() (*[]string, bool)`

GetQueryStringFieldsOk returns a tuple with the QueryStringFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStringFields

`func (o *ApplicationCachePutRequest) SetQueryStringFields(v []string)`

SetQueryStringFields sets QueryStringFields field to given value.

### HasQueryStringFields

`func (o *ApplicationCachePutRequest) HasQueryStringFields() bool`

HasQueryStringFields returns a boolean if a field has been set.

### GetEnableQueryStringSort

`func (o *ApplicationCachePutRequest) GetEnableQueryStringSort() bool`

GetEnableQueryStringSort returns the EnableQueryStringSort field if non-nil, zero value otherwise.

### GetEnableQueryStringSortOk

`func (o *ApplicationCachePutRequest) GetEnableQueryStringSortOk() (*bool, bool)`

GetEnableQueryStringSortOk returns a tuple with the EnableQueryStringSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQueryStringSort

`func (o *ApplicationCachePutRequest) SetEnableQueryStringSort(v bool)`

SetEnableQueryStringSort sets EnableQueryStringSort field to given value.

### HasEnableQueryStringSort

`func (o *ApplicationCachePutRequest) HasEnableQueryStringSort() bool`

HasEnableQueryStringSort returns a boolean if a field has been set.

### GetCacheByCookies

`func (o *ApplicationCachePutRequest) GetCacheByCookies() string`

GetCacheByCookies returns the CacheByCookies field if non-nil, zero value otherwise.

### GetCacheByCookiesOk

`func (o *ApplicationCachePutRequest) GetCacheByCookiesOk() (*string, bool)`

GetCacheByCookiesOk returns a tuple with the CacheByCookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByCookies

`func (o *ApplicationCachePutRequest) SetCacheByCookies(v string)`

SetCacheByCookies sets CacheByCookies field to given value.

### HasCacheByCookies

`func (o *ApplicationCachePutRequest) HasCacheByCookies() bool`

HasCacheByCookies returns a boolean if a field has been set.

### GetCookieNames

`func (o *ApplicationCachePutRequest) GetCookieNames() []string`

GetCookieNames returns the CookieNames field if non-nil, zero value otherwise.

### GetCookieNamesOk

`func (o *ApplicationCachePutRequest) GetCookieNamesOk() (*[]string, bool)`

GetCookieNamesOk returns a tuple with the CookieNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieNames

`func (o *ApplicationCachePutRequest) SetCookieNames(v []string)`

SetCookieNames sets CookieNames field to given value.

### HasCookieNames

`func (o *ApplicationCachePutRequest) HasCookieNames() bool`

HasCookieNames returns a boolean if a field has been set.

### GetEnableCachingForPost

`func (o *ApplicationCachePutRequest) GetEnableCachingForPost() bool`

GetEnableCachingForPost returns the EnableCachingForPost field if non-nil, zero value otherwise.

### GetEnableCachingForPostOk

`func (o *ApplicationCachePutRequest) GetEnableCachingForPostOk() (*bool, bool)`

GetEnableCachingForPostOk returns a tuple with the EnableCachingForPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCachingForPost

`func (o *ApplicationCachePutRequest) SetEnableCachingForPost(v bool)`

SetEnableCachingForPost sets EnableCachingForPost field to given value.

### HasEnableCachingForPost

`func (o *ApplicationCachePutRequest) HasEnableCachingForPost() bool`

HasEnableCachingForPost returns a boolean if a field has been set.

### GetL2CachingEnabled

`func (o *ApplicationCachePutRequest) GetL2CachingEnabled() bool`

GetL2CachingEnabled returns the L2CachingEnabled field if non-nil, zero value otherwise.

### GetL2CachingEnabledOk

`func (o *ApplicationCachePutRequest) GetL2CachingEnabledOk() (*bool, bool)`

GetL2CachingEnabledOk returns a tuple with the L2CachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2CachingEnabled

`func (o *ApplicationCachePutRequest) SetL2CachingEnabled(v bool)`

SetL2CachingEnabled sets L2CachingEnabled field to given value.

### HasL2CachingEnabled

`func (o *ApplicationCachePutRequest) HasL2CachingEnabled() bool`

HasL2CachingEnabled returns a boolean if a field has been set.

### GetIsSliceConfigurationEnabled

`func (o *ApplicationCachePutRequest) GetIsSliceConfigurationEnabled() bool`

GetIsSliceConfigurationEnabled returns the IsSliceConfigurationEnabled field if non-nil, zero value otherwise.

### GetIsSliceConfigurationEnabledOk

`func (o *ApplicationCachePutRequest) GetIsSliceConfigurationEnabledOk() (*bool, bool)`

GetIsSliceConfigurationEnabledOk returns a tuple with the IsSliceConfigurationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSliceConfigurationEnabled

`func (o *ApplicationCachePutRequest) SetIsSliceConfigurationEnabled(v bool)`

SetIsSliceConfigurationEnabled sets IsSliceConfigurationEnabled field to given value.

### HasIsSliceConfigurationEnabled

`func (o *ApplicationCachePutRequest) HasIsSliceConfigurationEnabled() bool`

HasIsSliceConfigurationEnabled returns a boolean if a field has been set.

### GetIsSliceEdgeCachingEnabled

`func (o *ApplicationCachePutRequest) GetIsSliceEdgeCachingEnabled() bool`

GetIsSliceEdgeCachingEnabled returns the IsSliceEdgeCachingEnabled field if non-nil, zero value otherwise.

### GetIsSliceEdgeCachingEnabledOk

`func (o *ApplicationCachePutRequest) GetIsSliceEdgeCachingEnabledOk() (*bool, bool)`

GetIsSliceEdgeCachingEnabledOk returns a tuple with the IsSliceEdgeCachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSliceEdgeCachingEnabled

`func (o *ApplicationCachePutRequest) SetIsSliceEdgeCachingEnabled(v bool)`

SetIsSliceEdgeCachingEnabled sets IsSliceEdgeCachingEnabled field to given value.

### HasIsSliceEdgeCachingEnabled

`func (o *ApplicationCachePutRequest) HasIsSliceEdgeCachingEnabled() bool`

HasIsSliceEdgeCachingEnabled returns a boolean if a field has been set.

### GetIsSliceL2CachingEnabled

`func (o *ApplicationCachePutRequest) GetIsSliceL2CachingEnabled() bool`

GetIsSliceL2CachingEnabled returns the IsSliceL2CachingEnabled field if non-nil, zero value otherwise.

### GetIsSliceL2CachingEnabledOk

`func (o *ApplicationCachePutRequest) GetIsSliceL2CachingEnabledOk() (*bool, bool)`

GetIsSliceL2CachingEnabledOk returns a tuple with the IsSliceL2CachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSliceL2CachingEnabled

`func (o *ApplicationCachePutRequest) SetIsSliceL2CachingEnabled(v bool)`

SetIsSliceL2CachingEnabled sets IsSliceL2CachingEnabled field to given value.

### HasIsSliceL2CachingEnabled

`func (o *ApplicationCachePutRequest) HasIsSliceL2CachingEnabled() bool`

HasIsSliceL2CachingEnabled returns a boolean if a field has been set.

### GetSliceConfigurationRange

`func (o *ApplicationCachePutRequest) GetSliceConfigurationRange() int64`

GetSliceConfigurationRange returns the SliceConfigurationRange field if non-nil, zero value otherwise.

### GetSliceConfigurationRangeOk

`func (o *ApplicationCachePutRequest) GetSliceConfigurationRangeOk() (*int64, bool)`

GetSliceConfigurationRangeOk returns a tuple with the SliceConfigurationRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceConfigurationRange

`func (o *ApplicationCachePutRequest) SetSliceConfigurationRange(v int64)`

SetSliceConfigurationRange sets SliceConfigurationRange field to given value.

### HasSliceConfigurationRange

`func (o *ApplicationCachePutRequest) HasSliceConfigurationRange() bool`

HasSliceConfigurationRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


