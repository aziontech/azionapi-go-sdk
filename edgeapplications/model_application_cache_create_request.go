/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApplicationCacheCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCacheCreateRequest{}

// ApplicationCacheCreateRequest struct for ApplicationCacheCreateRequest
type ApplicationCacheCreateRequest struct {
	Name string `json:"name"`
	BrowserCacheSettings *string `json:"browser_cache_settings,omitempty"`
	BrowserCacheSettingsMaximumTtl *int64 `json:"browser_cache_settings_maximum_ttl,omitempty"`
	CdnCacheSettings *string `json:"cdn_cache_settings,omitempty"`
	CdnCacheSettingsMaximumTtl *int64 `json:"cdn_cache_settings_maximum_ttl,omitempty"`
	CacheByQueryString *string `json:"cache_by_query_string,omitempty"`
	QueryStringFields []string `json:"query_string_fields,omitempty"`
	EnableQueryStringSort *bool `json:"enable_query_string_sort,omitempty"`
	CacheByCookies *string `json:"cache_by_cookies,omitempty"`
	CookieNames []string `json:"cookie_names,omitempty"`
	AdaptiveDeliveryAction *string `json:"adaptive_delivery_action,omitempty"`
	DeviceGroup []int32 `json:"device_group,omitempty"`
	EnableCachingForPost *bool `json:"enable_caching_for_post,omitempty"`
	L2CachingEnabled *bool `json:"l2_caching_enabled,omitempty"`
	IsSliceConfigurationEnabled *bool `json:"is_slice_configuration_enabled,omitempty"`
	IsSliceEdgeCachingEnabled *bool `json:"is_slice_edge_caching_enabled,omitempty"`
	IsSliceL2CachingEnabled *bool `json:"is_slice_l2_caching_enabled,omitempty"`
	SliceConfigurationRange *int64 `json:"slice_configuration_range,omitempty"`
	EnableCachingForOptions *bool `json:"enable_caching_for_options,omitempty"`
	EnableStaleCache *bool `json:"enable_stale_cache,omitempty"`
	L2Region *string `json:"l2_region,omitempty"`
}

type _ApplicationCacheCreateRequest ApplicationCacheCreateRequest

// NewApplicationCacheCreateRequest instantiates a new ApplicationCacheCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCacheCreateRequest(name string) *ApplicationCacheCreateRequest {
	this := ApplicationCacheCreateRequest{}
	this.Name = name
	return &this
}

// NewApplicationCacheCreateRequestWithDefaults instantiates a new ApplicationCacheCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCacheCreateRequestWithDefaults() *ApplicationCacheCreateRequest {
	this := ApplicationCacheCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationCacheCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCacheCreateRequest) SetName(v string) {
	o.Name = v
}

// GetBrowserCacheSettings returns the BrowserCacheSettings field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetBrowserCacheSettings() string {
	if o == nil || IsNil(o.BrowserCacheSettings) {
		var ret string
		return ret
	}
	return *o.BrowserCacheSettings
}

// GetBrowserCacheSettingsOk returns a tuple with the BrowserCacheSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetBrowserCacheSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.BrowserCacheSettings) {
		return nil, false
	}
	return o.BrowserCacheSettings, true
}

// HasBrowserCacheSettings returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasBrowserCacheSettings() bool {
	if o != nil && !IsNil(o.BrowserCacheSettings) {
		return true
	}

	return false
}

// SetBrowserCacheSettings gets a reference to the given string and assigns it to the BrowserCacheSettings field.
func (o *ApplicationCacheCreateRequest) SetBrowserCacheSettings(v string) {
	o.BrowserCacheSettings = &v
}

// GetBrowserCacheSettingsMaximumTtl returns the BrowserCacheSettingsMaximumTtl field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetBrowserCacheSettingsMaximumTtl() int64 {
	if o == nil || IsNil(o.BrowserCacheSettingsMaximumTtl) {
		var ret int64
		return ret
	}
	return *o.BrowserCacheSettingsMaximumTtl
}

// GetBrowserCacheSettingsMaximumTtlOk returns a tuple with the BrowserCacheSettingsMaximumTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetBrowserCacheSettingsMaximumTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.BrowserCacheSettingsMaximumTtl) {
		return nil, false
	}
	return o.BrowserCacheSettingsMaximumTtl, true
}

// HasBrowserCacheSettingsMaximumTtl returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasBrowserCacheSettingsMaximumTtl() bool {
	if o != nil && !IsNil(o.BrowserCacheSettingsMaximumTtl) {
		return true
	}

	return false
}

// SetBrowserCacheSettingsMaximumTtl gets a reference to the given int64 and assigns it to the BrowserCacheSettingsMaximumTtl field.
func (o *ApplicationCacheCreateRequest) SetBrowserCacheSettingsMaximumTtl(v int64) {
	o.BrowserCacheSettingsMaximumTtl = &v
}

// GetCdnCacheSettings returns the CdnCacheSettings field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetCdnCacheSettings() string {
	if o == nil || IsNil(o.CdnCacheSettings) {
		var ret string
		return ret
	}
	return *o.CdnCacheSettings
}

// GetCdnCacheSettingsOk returns a tuple with the CdnCacheSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetCdnCacheSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.CdnCacheSettings) {
		return nil, false
	}
	return o.CdnCacheSettings, true
}

// HasCdnCacheSettings returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasCdnCacheSettings() bool {
	if o != nil && !IsNil(o.CdnCacheSettings) {
		return true
	}

	return false
}

// SetCdnCacheSettings gets a reference to the given string and assigns it to the CdnCacheSettings field.
func (o *ApplicationCacheCreateRequest) SetCdnCacheSettings(v string) {
	o.CdnCacheSettings = &v
}

// GetCdnCacheSettingsMaximumTtl returns the CdnCacheSettingsMaximumTtl field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetCdnCacheSettingsMaximumTtl() int64 {
	if o == nil || IsNil(o.CdnCacheSettingsMaximumTtl) {
		var ret int64
		return ret
	}
	return *o.CdnCacheSettingsMaximumTtl
}

// GetCdnCacheSettingsMaximumTtlOk returns a tuple with the CdnCacheSettingsMaximumTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetCdnCacheSettingsMaximumTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.CdnCacheSettingsMaximumTtl) {
		return nil, false
	}
	return o.CdnCacheSettingsMaximumTtl, true
}

// HasCdnCacheSettingsMaximumTtl returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasCdnCacheSettingsMaximumTtl() bool {
	if o != nil && !IsNil(o.CdnCacheSettingsMaximumTtl) {
		return true
	}

	return false
}

// SetCdnCacheSettingsMaximumTtl gets a reference to the given int64 and assigns it to the CdnCacheSettingsMaximumTtl field.
func (o *ApplicationCacheCreateRequest) SetCdnCacheSettingsMaximumTtl(v int64) {
	o.CdnCacheSettingsMaximumTtl = &v
}

// GetCacheByQueryString returns the CacheByQueryString field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetCacheByQueryString() string {
	if o == nil || IsNil(o.CacheByQueryString) {
		var ret string
		return ret
	}
	return *o.CacheByQueryString
}

// GetCacheByQueryStringOk returns a tuple with the CacheByQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetCacheByQueryStringOk() (*string, bool) {
	if o == nil || IsNil(o.CacheByQueryString) {
		return nil, false
	}
	return o.CacheByQueryString, true
}

// HasCacheByQueryString returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasCacheByQueryString() bool {
	if o != nil && !IsNil(o.CacheByQueryString) {
		return true
	}

	return false
}

// SetCacheByQueryString gets a reference to the given string and assigns it to the CacheByQueryString field.
func (o *ApplicationCacheCreateRequest) SetCacheByQueryString(v string) {
	o.CacheByQueryString = &v
}

// GetQueryStringFields returns the QueryStringFields field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetQueryStringFields() []string {
	if o == nil || IsNil(o.QueryStringFields) {
		var ret []string
		return ret
	}
	return o.QueryStringFields
}

// GetQueryStringFieldsOk returns a tuple with the QueryStringFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetQueryStringFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.QueryStringFields) {
		return nil, false
	}
	return o.QueryStringFields, true
}

// HasQueryStringFields returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasQueryStringFields() bool {
	if o != nil && !IsNil(o.QueryStringFields) {
		return true
	}

	return false
}

// SetQueryStringFields gets a reference to the given []string and assigns it to the QueryStringFields field.
func (o *ApplicationCacheCreateRequest) SetQueryStringFields(v []string) {
	o.QueryStringFields = v
}

// GetEnableQueryStringSort returns the EnableQueryStringSort field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetEnableQueryStringSort() bool {
	if o == nil || IsNil(o.EnableQueryStringSort) {
		var ret bool
		return ret
	}
	return *o.EnableQueryStringSort
}

// GetEnableQueryStringSortOk returns a tuple with the EnableQueryStringSort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetEnableQueryStringSortOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableQueryStringSort) {
		return nil, false
	}
	return o.EnableQueryStringSort, true
}

// HasEnableQueryStringSort returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasEnableQueryStringSort() bool {
	if o != nil && !IsNil(o.EnableQueryStringSort) {
		return true
	}

	return false
}

// SetEnableQueryStringSort gets a reference to the given bool and assigns it to the EnableQueryStringSort field.
func (o *ApplicationCacheCreateRequest) SetEnableQueryStringSort(v bool) {
	o.EnableQueryStringSort = &v
}

// GetCacheByCookies returns the CacheByCookies field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetCacheByCookies() string {
	if o == nil || IsNil(o.CacheByCookies) {
		var ret string
		return ret
	}
	return *o.CacheByCookies
}

// GetCacheByCookiesOk returns a tuple with the CacheByCookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetCacheByCookiesOk() (*string, bool) {
	if o == nil || IsNil(o.CacheByCookies) {
		return nil, false
	}
	return o.CacheByCookies, true
}

// HasCacheByCookies returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasCacheByCookies() bool {
	if o != nil && !IsNil(o.CacheByCookies) {
		return true
	}

	return false
}

// SetCacheByCookies gets a reference to the given string and assigns it to the CacheByCookies field.
func (o *ApplicationCacheCreateRequest) SetCacheByCookies(v string) {
	o.CacheByCookies = &v
}

// GetCookieNames returns the CookieNames field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetCookieNames() []string {
	if o == nil || IsNil(o.CookieNames) {
		var ret []string
		return ret
	}
	return o.CookieNames
}

// GetCookieNamesOk returns a tuple with the CookieNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetCookieNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.CookieNames) {
		return nil, false
	}
	return o.CookieNames, true
}

// HasCookieNames returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasCookieNames() bool {
	if o != nil && !IsNil(o.CookieNames) {
		return true
	}

	return false
}

// SetCookieNames gets a reference to the given []string and assigns it to the CookieNames field.
func (o *ApplicationCacheCreateRequest) SetCookieNames(v []string) {
	o.CookieNames = v
}

// GetAdaptiveDeliveryAction returns the AdaptiveDeliveryAction field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetAdaptiveDeliveryAction() string {
	if o == nil || IsNil(o.AdaptiveDeliveryAction) {
		var ret string
		return ret
	}
	return *o.AdaptiveDeliveryAction
}

// GetAdaptiveDeliveryActionOk returns a tuple with the AdaptiveDeliveryAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetAdaptiveDeliveryActionOk() (*string, bool) {
	if o == nil || IsNil(o.AdaptiveDeliveryAction) {
		return nil, false
	}
	return o.AdaptiveDeliveryAction, true
}

// HasAdaptiveDeliveryAction returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasAdaptiveDeliveryAction() bool {
	if o != nil && !IsNil(o.AdaptiveDeliveryAction) {
		return true
	}

	return false
}

// SetAdaptiveDeliveryAction gets a reference to the given string and assigns it to the AdaptiveDeliveryAction field.
func (o *ApplicationCacheCreateRequest) SetAdaptiveDeliveryAction(v string) {
	o.AdaptiveDeliveryAction = &v
}

// GetDeviceGroup returns the DeviceGroup field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetDeviceGroup() []int32 {
	if o == nil || IsNil(o.DeviceGroup) {
		var ret []int32
		return ret
	}
	return o.DeviceGroup
}

// GetDeviceGroupOk returns a tuple with the DeviceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetDeviceGroupOk() ([]int32, bool) {
	if o == nil || IsNil(o.DeviceGroup) {
		return nil, false
	}
	return o.DeviceGroup, true
}

// HasDeviceGroup returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasDeviceGroup() bool {
	if o != nil && !IsNil(o.DeviceGroup) {
		return true
	}

	return false
}

// SetDeviceGroup gets a reference to the given []int32 and assigns it to the DeviceGroup field.
func (o *ApplicationCacheCreateRequest) SetDeviceGroup(v []int32) {
	o.DeviceGroup = v
}

// GetEnableCachingForPost returns the EnableCachingForPost field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetEnableCachingForPost() bool {
	if o == nil || IsNil(o.EnableCachingForPost) {
		var ret bool
		return ret
	}
	return *o.EnableCachingForPost
}

// GetEnableCachingForPostOk returns a tuple with the EnableCachingForPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetEnableCachingForPostOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCachingForPost) {
		return nil, false
	}
	return o.EnableCachingForPost, true
}

// HasEnableCachingForPost returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasEnableCachingForPost() bool {
	if o != nil && !IsNil(o.EnableCachingForPost) {
		return true
	}

	return false
}

// SetEnableCachingForPost gets a reference to the given bool and assigns it to the EnableCachingForPost field.
func (o *ApplicationCacheCreateRequest) SetEnableCachingForPost(v bool) {
	o.EnableCachingForPost = &v
}

// GetL2CachingEnabled returns the L2CachingEnabled field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetL2CachingEnabled() bool {
	if o == nil || IsNil(o.L2CachingEnabled) {
		var ret bool
		return ret
	}
	return *o.L2CachingEnabled
}

// GetL2CachingEnabledOk returns a tuple with the L2CachingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetL2CachingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.L2CachingEnabled) {
		return nil, false
	}
	return o.L2CachingEnabled, true
}

// HasL2CachingEnabled returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasL2CachingEnabled() bool {
	if o != nil && !IsNil(o.L2CachingEnabled) {
		return true
	}

	return false
}

// SetL2CachingEnabled gets a reference to the given bool and assigns it to the L2CachingEnabled field.
func (o *ApplicationCacheCreateRequest) SetL2CachingEnabled(v bool) {
	o.L2CachingEnabled = &v
}

// GetIsSliceConfigurationEnabled returns the IsSliceConfigurationEnabled field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetIsSliceConfigurationEnabled() bool {
	if o == nil || IsNil(o.IsSliceConfigurationEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSliceConfigurationEnabled
}

// GetIsSliceConfigurationEnabledOk returns a tuple with the IsSliceConfigurationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetIsSliceConfigurationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSliceConfigurationEnabled) {
		return nil, false
	}
	return o.IsSliceConfigurationEnabled, true
}

// HasIsSliceConfigurationEnabled returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasIsSliceConfigurationEnabled() bool {
	if o != nil && !IsNil(o.IsSliceConfigurationEnabled) {
		return true
	}

	return false
}

// SetIsSliceConfigurationEnabled gets a reference to the given bool and assigns it to the IsSliceConfigurationEnabled field.
func (o *ApplicationCacheCreateRequest) SetIsSliceConfigurationEnabled(v bool) {
	o.IsSliceConfigurationEnabled = &v
}

// GetIsSliceEdgeCachingEnabled returns the IsSliceEdgeCachingEnabled field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetIsSliceEdgeCachingEnabled() bool {
	if o == nil || IsNil(o.IsSliceEdgeCachingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSliceEdgeCachingEnabled
}

// GetIsSliceEdgeCachingEnabledOk returns a tuple with the IsSliceEdgeCachingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetIsSliceEdgeCachingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSliceEdgeCachingEnabled) {
		return nil, false
	}
	return o.IsSliceEdgeCachingEnabled, true
}

// HasIsSliceEdgeCachingEnabled returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasIsSliceEdgeCachingEnabled() bool {
	if o != nil && !IsNil(o.IsSliceEdgeCachingEnabled) {
		return true
	}

	return false
}

// SetIsSliceEdgeCachingEnabled gets a reference to the given bool and assigns it to the IsSliceEdgeCachingEnabled field.
func (o *ApplicationCacheCreateRequest) SetIsSliceEdgeCachingEnabled(v bool) {
	o.IsSliceEdgeCachingEnabled = &v
}

// GetIsSliceL2CachingEnabled returns the IsSliceL2CachingEnabled field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetIsSliceL2CachingEnabled() bool {
	if o == nil || IsNil(o.IsSliceL2CachingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSliceL2CachingEnabled
}

// GetIsSliceL2CachingEnabledOk returns a tuple with the IsSliceL2CachingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetIsSliceL2CachingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSliceL2CachingEnabled) {
		return nil, false
	}
	return o.IsSliceL2CachingEnabled, true
}

// HasIsSliceL2CachingEnabled returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasIsSliceL2CachingEnabled() bool {
	if o != nil && !IsNil(o.IsSliceL2CachingEnabled) {
		return true
	}

	return false
}

// SetIsSliceL2CachingEnabled gets a reference to the given bool and assigns it to the IsSliceL2CachingEnabled field.
func (o *ApplicationCacheCreateRequest) SetIsSliceL2CachingEnabled(v bool) {
	o.IsSliceL2CachingEnabled = &v
}

// GetSliceConfigurationRange returns the SliceConfigurationRange field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetSliceConfigurationRange() int64 {
	if o == nil || IsNil(o.SliceConfigurationRange) {
		var ret int64
		return ret
	}
	return *o.SliceConfigurationRange
}

// GetSliceConfigurationRangeOk returns a tuple with the SliceConfigurationRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetSliceConfigurationRangeOk() (*int64, bool) {
	if o == nil || IsNil(o.SliceConfigurationRange) {
		return nil, false
	}
	return o.SliceConfigurationRange, true
}

// HasSliceConfigurationRange returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasSliceConfigurationRange() bool {
	if o != nil && !IsNil(o.SliceConfigurationRange) {
		return true
	}

	return false
}

// SetSliceConfigurationRange gets a reference to the given int64 and assigns it to the SliceConfigurationRange field.
func (o *ApplicationCacheCreateRequest) SetSliceConfigurationRange(v int64) {
	o.SliceConfigurationRange = &v
}

// GetEnableCachingForOptions returns the EnableCachingForOptions field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetEnableCachingForOptions() bool {
	if o == nil || IsNil(o.EnableCachingForOptions) {
		var ret bool
		return ret
	}
	return *o.EnableCachingForOptions
}

// GetEnableCachingForOptionsOk returns a tuple with the EnableCachingForOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetEnableCachingForOptionsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableCachingForOptions) {
		return nil, false
	}
	return o.EnableCachingForOptions, true
}

// HasEnableCachingForOptions returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasEnableCachingForOptions() bool {
	if o != nil && !IsNil(o.EnableCachingForOptions) {
		return true
	}

	return false
}

// SetEnableCachingForOptions gets a reference to the given bool and assigns it to the EnableCachingForOptions field.
func (o *ApplicationCacheCreateRequest) SetEnableCachingForOptions(v bool) {
	o.EnableCachingForOptions = &v
}

// GetEnableStaleCache returns the EnableStaleCache field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetEnableStaleCache() bool {
	if o == nil || IsNil(o.EnableStaleCache) {
		var ret bool
		return ret
	}
	return *o.EnableStaleCache
}

// GetEnableStaleCacheOk returns a tuple with the EnableStaleCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetEnableStaleCacheOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableStaleCache) {
		return nil, false
	}
	return o.EnableStaleCache, true
}

// HasEnableStaleCache returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasEnableStaleCache() bool {
	if o != nil && !IsNil(o.EnableStaleCache) {
		return true
	}

	return false
}

// SetEnableStaleCache gets a reference to the given bool and assigns it to the EnableStaleCache field.
func (o *ApplicationCacheCreateRequest) SetEnableStaleCache(v bool) {
	o.EnableStaleCache = &v
}

// GetL2Region returns the L2Region field value if set, zero value otherwise.
func (o *ApplicationCacheCreateRequest) GetL2Region() string {
	if o == nil || IsNil(o.L2Region) {
		var ret string
		return ret
	}
	return *o.L2Region
}

// GetL2RegionOk returns a tuple with the L2Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheCreateRequest) GetL2RegionOk() (*string, bool) {
	if o == nil || IsNil(o.L2Region) {
		return nil, false
	}
	return o.L2Region, true
}

// HasL2Region returns a boolean if a field has been set.
func (o *ApplicationCacheCreateRequest) HasL2Region() bool {
	if o != nil && !IsNil(o.L2Region) {
		return true
	}

	return false
}

// SetL2Region gets a reference to the given string and assigns it to the L2Region field.
func (o *ApplicationCacheCreateRequest) SetL2Region(v string) {
	o.L2Region = &v
}

func (o ApplicationCacheCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCacheCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.BrowserCacheSettings) {
		toSerialize["browser_cache_settings"] = o.BrowserCacheSettings
	}
	if !IsNil(o.BrowserCacheSettingsMaximumTtl) {
		toSerialize["browser_cache_settings_maximum_ttl"] = o.BrowserCacheSettingsMaximumTtl
	}
	if !IsNil(o.CdnCacheSettings) {
		toSerialize["cdn_cache_settings"] = o.CdnCacheSettings
	}
	if !IsNil(o.CdnCacheSettingsMaximumTtl) {
		toSerialize["cdn_cache_settings_maximum_ttl"] = o.CdnCacheSettingsMaximumTtl
	}
	if !IsNil(o.CacheByQueryString) {
		toSerialize["cache_by_query_string"] = o.CacheByQueryString
	}
	if !IsNil(o.QueryStringFields) {
		toSerialize["query_string_fields"] = o.QueryStringFields
	}
	if !IsNil(o.EnableQueryStringSort) {
		toSerialize["enable_query_string_sort"] = o.EnableQueryStringSort
	}
	if !IsNil(o.CacheByCookies) {
		toSerialize["cache_by_cookies"] = o.CacheByCookies
	}
	if !IsNil(o.CookieNames) {
		toSerialize["cookie_names"] = o.CookieNames
	}
	if !IsNil(o.AdaptiveDeliveryAction) {
		toSerialize["adaptive_delivery_action"] = o.AdaptiveDeliveryAction
	}
	if !IsNil(o.DeviceGroup) {
		toSerialize["device_group"] = o.DeviceGroup
	}
	if !IsNil(o.EnableCachingForPost) {
		toSerialize["enable_caching_for_post"] = o.EnableCachingForPost
	}
	if !IsNil(o.L2CachingEnabled) {
		toSerialize["l2_caching_enabled"] = o.L2CachingEnabled
	}
	if !IsNil(o.IsSliceConfigurationEnabled) {
		toSerialize["is_slice_configuration_enabled"] = o.IsSliceConfigurationEnabled
	}
	if !IsNil(o.IsSliceEdgeCachingEnabled) {
		toSerialize["is_slice_edge_caching_enabled"] = o.IsSliceEdgeCachingEnabled
	}
	if !IsNil(o.IsSliceL2CachingEnabled) {
		toSerialize["is_slice_l2_caching_enabled"] = o.IsSliceL2CachingEnabled
	}
	if !IsNil(o.SliceConfigurationRange) {
		toSerialize["slice_configuration_range"] = o.SliceConfigurationRange
	}
	if !IsNil(o.EnableCachingForOptions) {
		toSerialize["enable_caching_for_options"] = o.EnableCachingForOptions
	}
	if !IsNil(o.EnableStaleCache) {
		toSerialize["enable_stale_cache"] = o.EnableStaleCache
	}
	if !IsNil(o.L2Region) {
		toSerialize["l2_region"] = o.L2Region
	}
	return toSerialize, nil
}

func (o *ApplicationCacheCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApplicationCacheCreateRequest := _ApplicationCacheCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCacheCreateRequest)

	if err != nil {
		return err
	}

	*o = ApplicationCacheCreateRequest(varApplicationCacheCreateRequest)

	return err
}

type NullableApplicationCacheCreateRequest struct {
	value *ApplicationCacheCreateRequest
	isSet bool
}

func (v NullableApplicationCacheCreateRequest) Get() *ApplicationCacheCreateRequest {
	return v.value
}

func (v *NullableApplicationCacheCreateRequest) Set(val *ApplicationCacheCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCacheCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCacheCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCacheCreateRequest(val *ApplicationCacheCreateRequest) *NullableApplicationCacheCreateRequest {
	return &NullableApplicationCacheCreateRequest{value: val, isSet: true}
}

func (v NullableApplicationCacheCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCacheCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


