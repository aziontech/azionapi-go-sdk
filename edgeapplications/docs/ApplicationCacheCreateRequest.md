# ApplicationCacheCreateRequest

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
**UserEmail** | Pointer to **string** |  | [optional] 
**L2CachingEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewApplicationCacheCreateRequest

`func NewApplicationCacheCreateRequest(name string, ) *ApplicationCacheCreateRequest`

NewApplicationCacheCreateRequest instantiates a new ApplicationCacheCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCacheCreateRequestWithDefaults

`func NewApplicationCacheCreateRequestWithDefaults() *ApplicationCacheCreateRequest`

NewApplicationCacheCreateRequestWithDefaults instantiates a new ApplicationCacheCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationCacheCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCacheCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCacheCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBrowserCacheSettings

`func (o *ApplicationCacheCreateRequest) GetBrowserCacheSettings() string`

GetBrowserCacheSettings returns the BrowserCacheSettings field if non-nil, zero value otherwise.

### GetBrowserCacheSettingsOk

`func (o *ApplicationCacheCreateRequest) GetBrowserCacheSettingsOk() (*string, bool)`

GetBrowserCacheSettingsOk returns a tuple with the BrowserCacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCacheSettings

`func (o *ApplicationCacheCreateRequest) SetBrowserCacheSettings(v string)`

SetBrowserCacheSettings sets BrowserCacheSettings field to given value.

### HasBrowserCacheSettings

`func (o *ApplicationCacheCreateRequest) HasBrowserCacheSettings() bool`

HasBrowserCacheSettings returns a boolean if a field has been set.

### GetBrowserCacheSettingsMaximumTtl

`func (o *ApplicationCacheCreateRequest) GetBrowserCacheSettingsMaximumTtl() int64`

GetBrowserCacheSettingsMaximumTtl returns the BrowserCacheSettingsMaximumTtl field if non-nil, zero value otherwise.

### GetBrowserCacheSettingsMaximumTtlOk

`func (o *ApplicationCacheCreateRequest) GetBrowserCacheSettingsMaximumTtlOk() (*int64, bool)`

GetBrowserCacheSettingsMaximumTtlOk returns a tuple with the BrowserCacheSettingsMaximumTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserCacheSettingsMaximumTtl

`func (o *ApplicationCacheCreateRequest) SetBrowserCacheSettingsMaximumTtl(v int64)`

SetBrowserCacheSettingsMaximumTtl sets BrowserCacheSettingsMaximumTtl field to given value.

### HasBrowserCacheSettingsMaximumTtl

`func (o *ApplicationCacheCreateRequest) HasBrowserCacheSettingsMaximumTtl() bool`

HasBrowserCacheSettingsMaximumTtl returns a boolean if a field has been set.

### GetCdnCacheSettings

`func (o *ApplicationCacheCreateRequest) GetCdnCacheSettings() string`

GetCdnCacheSettings returns the CdnCacheSettings field if non-nil, zero value otherwise.

### GetCdnCacheSettingsOk

`func (o *ApplicationCacheCreateRequest) GetCdnCacheSettingsOk() (*string, bool)`

GetCdnCacheSettingsOk returns a tuple with the CdnCacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnCacheSettings

`func (o *ApplicationCacheCreateRequest) SetCdnCacheSettings(v string)`

SetCdnCacheSettings sets CdnCacheSettings field to given value.

### HasCdnCacheSettings

`func (o *ApplicationCacheCreateRequest) HasCdnCacheSettings() bool`

HasCdnCacheSettings returns a boolean if a field has been set.

### GetCdnCacheSettingsMaximumTtl

`func (o *ApplicationCacheCreateRequest) GetCdnCacheSettingsMaximumTtl() int64`

GetCdnCacheSettingsMaximumTtl returns the CdnCacheSettingsMaximumTtl field if non-nil, zero value otherwise.

### GetCdnCacheSettingsMaximumTtlOk

`func (o *ApplicationCacheCreateRequest) GetCdnCacheSettingsMaximumTtlOk() (*int64, bool)`

GetCdnCacheSettingsMaximumTtlOk returns a tuple with the CdnCacheSettingsMaximumTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnCacheSettingsMaximumTtl

`func (o *ApplicationCacheCreateRequest) SetCdnCacheSettingsMaximumTtl(v int64)`

SetCdnCacheSettingsMaximumTtl sets CdnCacheSettingsMaximumTtl field to given value.

### HasCdnCacheSettingsMaximumTtl

`func (o *ApplicationCacheCreateRequest) HasCdnCacheSettingsMaximumTtl() bool`

HasCdnCacheSettingsMaximumTtl returns a boolean if a field has been set.

### GetCacheByQueryString

`func (o *ApplicationCacheCreateRequest) GetCacheByQueryString() string`

GetCacheByQueryString returns the CacheByQueryString field if non-nil, zero value otherwise.

### GetCacheByQueryStringOk

`func (o *ApplicationCacheCreateRequest) GetCacheByQueryStringOk() (*string, bool)`

GetCacheByQueryStringOk returns a tuple with the CacheByQueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByQueryString

`func (o *ApplicationCacheCreateRequest) SetCacheByQueryString(v string)`

SetCacheByQueryString sets CacheByQueryString field to given value.

### HasCacheByQueryString

`func (o *ApplicationCacheCreateRequest) HasCacheByQueryString() bool`

HasCacheByQueryString returns a boolean if a field has been set.

### GetQueryStringFields

`func (o *ApplicationCacheCreateRequest) GetQueryStringFields() []string`

GetQueryStringFields returns the QueryStringFields field if non-nil, zero value otherwise.

### GetQueryStringFieldsOk

`func (o *ApplicationCacheCreateRequest) GetQueryStringFieldsOk() (*[]string, bool)`

GetQueryStringFieldsOk returns a tuple with the QueryStringFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStringFields

`func (o *ApplicationCacheCreateRequest) SetQueryStringFields(v []string)`

SetQueryStringFields sets QueryStringFields field to given value.

### HasQueryStringFields

`func (o *ApplicationCacheCreateRequest) HasQueryStringFields() bool`

HasQueryStringFields returns a boolean if a field has been set.

### SetQueryStringFieldsNil

`func (o *ApplicationCacheCreateRequest) SetQueryStringFieldsNil(b bool)`

 SetQueryStringFieldsNil sets the value for QueryStringFields to be an explicit nil

### UnsetQueryStringFields
`func (o *ApplicationCacheCreateRequest) UnsetQueryStringFields()`

UnsetQueryStringFields ensures that no value is present for QueryStringFields, not even an explicit nil
### GetEnableQueryStringSort

`func (o *ApplicationCacheCreateRequest) GetEnableQueryStringSort() bool`

GetEnableQueryStringSort returns the EnableQueryStringSort field if non-nil, zero value otherwise.

### GetEnableQueryStringSortOk

`func (o *ApplicationCacheCreateRequest) GetEnableQueryStringSortOk() (*bool, bool)`

GetEnableQueryStringSortOk returns a tuple with the EnableQueryStringSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQueryStringSort

`func (o *ApplicationCacheCreateRequest) SetEnableQueryStringSort(v bool)`

SetEnableQueryStringSort sets EnableQueryStringSort field to given value.

### HasEnableQueryStringSort

`func (o *ApplicationCacheCreateRequest) HasEnableQueryStringSort() bool`

HasEnableQueryStringSort returns a boolean if a field has been set.

### GetCacheByCookies

`func (o *ApplicationCacheCreateRequest) GetCacheByCookies() string`

GetCacheByCookies returns the CacheByCookies field if non-nil, zero value otherwise.

### GetCacheByCookiesOk

`func (o *ApplicationCacheCreateRequest) GetCacheByCookiesOk() (*string, bool)`

GetCacheByCookiesOk returns a tuple with the CacheByCookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheByCookies

`func (o *ApplicationCacheCreateRequest) SetCacheByCookies(v string)`

SetCacheByCookies sets CacheByCookies field to given value.

### HasCacheByCookies

`func (o *ApplicationCacheCreateRequest) HasCacheByCookies() bool`

HasCacheByCookies returns a boolean if a field has been set.

### GetCookieNames

`func (o *ApplicationCacheCreateRequest) GetCookieNames() []string`

GetCookieNames returns the CookieNames field if non-nil, zero value otherwise.

### GetCookieNamesOk

`func (o *ApplicationCacheCreateRequest) GetCookieNamesOk() (*[]string, bool)`

GetCookieNamesOk returns a tuple with the CookieNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieNames

`func (o *ApplicationCacheCreateRequest) SetCookieNames(v []string)`

SetCookieNames sets CookieNames field to given value.

### HasCookieNames

`func (o *ApplicationCacheCreateRequest) HasCookieNames() bool`

HasCookieNames returns a boolean if a field has been set.

### SetCookieNamesNil

`func (o *ApplicationCacheCreateRequest) SetCookieNamesNil(b bool)`

 SetCookieNamesNil sets the value for CookieNames to be an explicit nil

### UnsetCookieNames
`func (o *ApplicationCacheCreateRequest) UnsetCookieNames()`

UnsetCookieNames ensures that no value is present for CookieNames, not even an explicit nil
### GetUserEmail

`func (o *ApplicationCacheCreateRequest) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *ApplicationCacheCreateRequest) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *ApplicationCacheCreateRequest) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *ApplicationCacheCreateRequest) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetL2CachingEnabled

`func (o *ApplicationCacheCreateRequest) GetL2CachingEnabled() bool`

GetL2CachingEnabled returns the L2CachingEnabled field if non-nil, zero value otherwise.

### GetL2CachingEnabledOk

`func (o *ApplicationCacheCreateRequest) GetL2CachingEnabledOk() (*bool, bool)`

GetL2CachingEnabledOk returns a tuple with the L2CachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2CachingEnabled

`func (o *ApplicationCacheCreateRequest) SetL2CachingEnabled(v bool)`

SetL2CachingEnabled sets L2CachingEnabled field to given value.

### HasL2CachingEnabled

`func (o *ApplicationCacheCreateRequest) HasL2CachingEnabled() bool`

HasL2CachingEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


