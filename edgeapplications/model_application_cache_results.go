/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
)

// checks if the ApplicationCacheResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCacheResults{}

// ApplicationCacheResults struct for ApplicationCacheResults
type ApplicationCacheResults struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	BrowserCacheSettings string `json:"browser_cache_settings"`
	BrowserCacheSettingsMaximumTtl int64 `json:"browser_cache_settings_maximum_ttl"`
	CdnCacheSettings string `json:"cdn_cache_settings"`
	CdnCacheSettingsMaximumTtl int64 `json:"cdn_cache_settings_maximum_ttl"`
	CacheByQueryString string `json:"cache_by_query_string"`
	QueryStringFields []string `json:"query_string_fields"`
	EnableQueryStringSort bool `json:"enable_query_string_sort"`
	CacheByCookies string `json:"cache_by_cookies"`
	CookieNames []string `json:"cookie_names"`
	AdaptiveDeliveryAction string `json:"adaptive_delivery_action"`
	DeviceGroup []string `json:"device_group"`
	EnableCachingForPost bool `json:"enable_caching_for_post"`
	L2CachingEnabled bool `json:"l2_caching_enabled"`
	IsSliceConfigurationEnabled *bool `json:"is_slice_configuration_enabled,omitempty"`
	IsSliceEdgeCachingEnabled *bool `json:"is_slice_edge_caching_enabled,omitempty"`
	IsSliceL2CachingEnabled *bool `json:"is_slice_l2_caching_enabled,omitempty"`
	SliceConfigurationRange *bool `json:"slice_configuration_range,omitempty"`
	EnableCachingForOptions bool `json:"enable_caching_for_options"`
	EnableStaleCache bool `json:"enable_stale_cache"`
	L2Region string `json:"l2_region"`
}

// NewApplicationCacheResults instantiates a new ApplicationCacheResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCacheResults(id int64, name string, browserCacheSettings string, browserCacheSettingsMaximumTtl int64, cdnCacheSettings string, cdnCacheSettingsMaximumTtl int64, cacheByQueryString string, queryStringFields []string, enableQueryStringSort bool, cacheByCookies string, cookieNames []string, adaptiveDeliveryAction string, deviceGroup []string, enableCachingForPost bool, l2CachingEnabled bool, enableCachingForOptions bool, enableStaleCache bool, l2Region string) *ApplicationCacheResults {
	this := ApplicationCacheResults{}
	this.Id = id
	this.Name = name
	this.BrowserCacheSettings = browserCacheSettings
	this.BrowserCacheSettingsMaximumTtl = browserCacheSettingsMaximumTtl
	this.CdnCacheSettings = cdnCacheSettings
	this.CdnCacheSettingsMaximumTtl = cdnCacheSettingsMaximumTtl
	this.CacheByQueryString = cacheByQueryString
	this.QueryStringFields = queryStringFields
	this.EnableQueryStringSort = enableQueryStringSort
	this.CacheByCookies = cacheByCookies
	this.CookieNames = cookieNames
	this.AdaptiveDeliveryAction = adaptiveDeliveryAction
	this.DeviceGroup = deviceGroup
	this.EnableCachingForPost = enableCachingForPost
	this.L2CachingEnabled = l2CachingEnabled
	this.EnableCachingForOptions = enableCachingForOptions
	this.EnableStaleCache = enableStaleCache
	this.L2Region = l2Region
	return &this
}

// NewApplicationCacheResultsWithDefaults instantiates a new ApplicationCacheResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCacheResultsWithDefaults() *ApplicationCacheResults {
	this := ApplicationCacheResults{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationCacheResults) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationCacheResults) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ApplicationCacheResults) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCacheResults) SetName(v string) {
	o.Name = v
}

// GetBrowserCacheSettings returns the BrowserCacheSettings field value
func (o *ApplicationCacheResults) GetBrowserCacheSettings() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrowserCacheSettings
}

// GetBrowserCacheSettingsOk returns a tuple with the BrowserCacheSettings field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetBrowserCacheSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrowserCacheSettings, true
}

// SetBrowserCacheSettings sets field value
func (o *ApplicationCacheResults) SetBrowserCacheSettings(v string) {
	o.BrowserCacheSettings = v
}

// GetBrowserCacheSettingsMaximumTtl returns the BrowserCacheSettingsMaximumTtl field value
func (o *ApplicationCacheResults) GetBrowserCacheSettingsMaximumTtl() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BrowserCacheSettingsMaximumTtl
}

// GetBrowserCacheSettingsMaximumTtlOk returns a tuple with the BrowserCacheSettingsMaximumTtl field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetBrowserCacheSettingsMaximumTtlOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrowserCacheSettingsMaximumTtl, true
}

// SetBrowserCacheSettingsMaximumTtl sets field value
func (o *ApplicationCacheResults) SetBrowserCacheSettingsMaximumTtl(v int64) {
	o.BrowserCacheSettingsMaximumTtl = v
}

// GetCdnCacheSettings returns the CdnCacheSettings field value
func (o *ApplicationCacheResults) GetCdnCacheSettings() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CdnCacheSettings
}

// GetCdnCacheSettingsOk returns a tuple with the CdnCacheSettings field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetCdnCacheSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CdnCacheSettings, true
}

// SetCdnCacheSettings sets field value
func (o *ApplicationCacheResults) SetCdnCacheSettings(v string) {
	o.CdnCacheSettings = v
}

// GetCdnCacheSettingsMaximumTtl returns the CdnCacheSettingsMaximumTtl field value
func (o *ApplicationCacheResults) GetCdnCacheSettingsMaximumTtl() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CdnCacheSettingsMaximumTtl
}

// GetCdnCacheSettingsMaximumTtlOk returns a tuple with the CdnCacheSettingsMaximumTtl field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetCdnCacheSettingsMaximumTtlOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CdnCacheSettingsMaximumTtl, true
}

// SetCdnCacheSettingsMaximumTtl sets field value
func (o *ApplicationCacheResults) SetCdnCacheSettingsMaximumTtl(v int64) {
	o.CdnCacheSettingsMaximumTtl = v
}

// GetCacheByQueryString returns the CacheByQueryString field value
func (o *ApplicationCacheResults) GetCacheByQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CacheByQueryString
}

// GetCacheByQueryStringOk returns a tuple with the CacheByQueryString field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetCacheByQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheByQueryString, true
}

// SetCacheByQueryString sets field value
func (o *ApplicationCacheResults) SetCacheByQueryString(v string) {
	o.CacheByQueryString = v
}

// GetQueryStringFields returns the QueryStringFields field value
func (o *ApplicationCacheResults) GetQueryStringFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.QueryStringFields
}

// GetQueryStringFieldsOk returns a tuple with the QueryStringFields field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetQueryStringFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueryStringFields, true
}

// SetQueryStringFields sets field value
func (o *ApplicationCacheResults) SetQueryStringFields(v []string) {
	o.QueryStringFields = v
}

// GetEnableQueryStringSort returns the EnableQueryStringSort field value
func (o *ApplicationCacheResults) GetEnableQueryStringSort() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableQueryStringSort
}

// GetEnableQueryStringSortOk returns a tuple with the EnableQueryStringSort field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetEnableQueryStringSortOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableQueryStringSort, true
}

// SetEnableQueryStringSort sets field value
func (o *ApplicationCacheResults) SetEnableQueryStringSort(v bool) {
	o.EnableQueryStringSort = v
}

// GetCacheByCookies returns the CacheByCookies field value
func (o *ApplicationCacheResults) GetCacheByCookies() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CacheByCookies
}

// GetCacheByCookiesOk returns a tuple with the CacheByCookies field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetCacheByCookiesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheByCookies, true
}

// SetCacheByCookies sets field value
func (o *ApplicationCacheResults) SetCacheByCookies(v string) {
	o.CacheByCookies = v
}

// GetCookieNames returns the CookieNames field value
func (o *ApplicationCacheResults) GetCookieNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CookieNames
}

// GetCookieNamesOk returns a tuple with the CookieNames field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetCookieNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CookieNames, true
}

// SetCookieNames sets field value
func (o *ApplicationCacheResults) SetCookieNames(v []string) {
	o.CookieNames = v
}

// GetAdaptiveDeliveryAction returns the AdaptiveDeliveryAction field value
func (o *ApplicationCacheResults) GetAdaptiveDeliveryAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdaptiveDeliveryAction
}

// GetAdaptiveDeliveryActionOk returns a tuple with the AdaptiveDeliveryAction field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetAdaptiveDeliveryActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdaptiveDeliveryAction, true
}

// SetAdaptiveDeliveryAction sets field value
func (o *ApplicationCacheResults) SetAdaptiveDeliveryAction(v string) {
	o.AdaptiveDeliveryAction = v
}

// GetDeviceGroup returns the DeviceGroup field value
func (o *ApplicationCacheResults) GetDeviceGroup() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DeviceGroup
}

// GetDeviceGroupOk returns a tuple with the DeviceGroup field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetDeviceGroupOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceGroup, true
}

// SetDeviceGroup sets field value
func (o *ApplicationCacheResults) SetDeviceGroup(v []string) {
	o.DeviceGroup = v
}

// GetEnableCachingForPost returns the EnableCachingForPost field value
func (o *ApplicationCacheResults) GetEnableCachingForPost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableCachingForPost
}

// GetEnableCachingForPostOk returns a tuple with the EnableCachingForPost field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetEnableCachingForPostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableCachingForPost, true
}

// SetEnableCachingForPost sets field value
func (o *ApplicationCacheResults) SetEnableCachingForPost(v bool) {
	o.EnableCachingForPost = v
}

// GetL2CachingEnabled returns the L2CachingEnabled field value
func (o *ApplicationCacheResults) GetL2CachingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.L2CachingEnabled
}

// GetL2CachingEnabledOk returns a tuple with the L2CachingEnabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetL2CachingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.L2CachingEnabled, true
}

// SetL2CachingEnabled sets field value
func (o *ApplicationCacheResults) SetL2CachingEnabled(v bool) {
	o.L2CachingEnabled = v
}

// GetIsSliceConfigurationEnabled returns the IsSliceConfigurationEnabled field value if set, zero value otherwise.
func (o *ApplicationCacheResults) GetIsSliceConfigurationEnabled() bool {
	if o == nil || IsNil(o.IsSliceConfigurationEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSliceConfigurationEnabled
}

// GetIsSliceConfigurationEnabledOk returns a tuple with the IsSliceConfigurationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetIsSliceConfigurationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSliceConfigurationEnabled) {
		return nil, false
	}
	return o.IsSliceConfigurationEnabled, true
}

// HasIsSliceConfigurationEnabled returns a boolean if a field has been set.
func (o *ApplicationCacheResults) HasIsSliceConfigurationEnabled() bool {
	if o != nil && !IsNil(o.IsSliceConfigurationEnabled) {
		return true
	}

	return false
}

// SetIsSliceConfigurationEnabled gets a reference to the given bool and assigns it to the IsSliceConfigurationEnabled field.
func (o *ApplicationCacheResults) SetIsSliceConfigurationEnabled(v bool) {
	o.IsSliceConfigurationEnabled = &v
}

// GetIsSliceEdgeCachingEnabled returns the IsSliceEdgeCachingEnabled field value if set, zero value otherwise.
func (o *ApplicationCacheResults) GetIsSliceEdgeCachingEnabled() bool {
	if o == nil || IsNil(o.IsSliceEdgeCachingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSliceEdgeCachingEnabled
}

// GetIsSliceEdgeCachingEnabledOk returns a tuple with the IsSliceEdgeCachingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetIsSliceEdgeCachingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSliceEdgeCachingEnabled) {
		return nil, false
	}
	return o.IsSliceEdgeCachingEnabled, true
}

// HasIsSliceEdgeCachingEnabled returns a boolean if a field has been set.
func (o *ApplicationCacheResults) HasIsSliceEdgeCachingEnabled() bool {
	if o != nil && !IsNil(o.IsSliceEdgeCachingEnabled) {
		return true
	}

	return false
}

// SetIsSliceEdgeCachingEnabled gets a reference to the given bool and assigns it to the IsSliceEdgeCachingEnabled field.
func (o *ApplicationCacheResults) SetIsSliceEdgeCachingEnabled(v bool) {
	o.IsSliceEdgeCachingEnabled = &v
}

// GetIsSliceL2CachingEnabled returns the IsSliceL2CachingEnabled field value if set, zero value otherwise.
func (o *ApplicationCacheResults) GetIsSliceL2CachingEnabled() bool {
	if o == nil || IsNil(o.IsSliceL2CachingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSliceL2CachingEnabled
}

// GetIsSliceL2CachingEnabledOk returns a tuple with the IsSliceL2CachingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetIsSliceL2CachingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSliceL2CachingEnabled) {
		return nil, false
	}
	return o.IsSliceL2CachingEnabled, true
}

// HasIsSliceL2CachingEnabled returns a boolean if a field has been set.
func (o *ApplicationCacheResults) HasIsSliceL2CachingEnabled() bool {
	if o != nil && !IsNil(o.IsSliceL2CachingEnabled) {
		return true
	}

	return false
}

// SetIsSliceL2CachingEnabled gets a reference to the given bool and assigns it to the IsSliceL2CachingEnabled field.
func (o *ApplicationCacheResults) SetIsSliceL2CachingEnabled(v bool) {
	o.IsSliceL2CachingEnabled = &v
}

// GetSliceConfigurationRange returns the SliceConfigurationRange field value if set, zero value otherwise.
func (o *ApplicationCacheResults) GetSliceConfigurationRange() bool {
	if o == nil || IsNil(o.SliceConfigurationRange) {
		var ret bool
		return ret
	}
	return *o.SliceConfigurationRange
}

// GetSliceConfigurationRangeOk returns a tuple with the SliceConfigurationRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetSliceConfigurationRangeOk() (*bool, bool) {
	if o == nil || IsNil(o.SliceConfigurationRange) {
		return nil, false
	}
	return o.SliceConfigurationRange, true
}

// HasSliceConfigurationRange returns a boolean if a field has been set.
func (o *ApplicationCacheResults) HasSliceConfigurationRange() bool {
	if o != nil && !IsNil(o.SliceConfigurationRange) {
		return true
	}

	return false
}

// SetSliceConfigurationRange gets a reference to the given bool and assigns it to the SliceConfigurationRange field.
func (o *ApplicationCacheResults) SetSliceConfigurationRange(v bool) {
	o.SliceConfigurationRange = &v
}

// GetEnableCachingForOptions returns the EnableCachingForOptions field value
func (o *ApplicationCacheResults) GetEnableCachingForOptions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableCachingForOptions
}

// GetEnableCachingForOptionsOk returns a tuple with the EnableCachingForOptions field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetEnableCachingForOptionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableCachingForOptions, true
}

// SetEnableCachingForOptions sets field value
func (o *ApplicationCacheResults) SetEnableCachingForOptions(v bool) {
	o.EnableCachingForOptions = v
}

// GetEnableStaleCache returns the EnableStaleCache field value
func (o *ApplicationCacheResults) GetEnableStaleCache() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableStaleCache
}

// GetEnableStaleCacheOk returns a tuple with the EnableStaleCache field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetEnableStaleCacheOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableStaleCache, true
}

// SetEnableStaleCache sets field value
func (o *ApplicationCacheResults) SetEnableStaleCache(v bool) {
	o.EnableStaleCache = v
}

// GetL2Region returns the L2Region field value
func (o *ApplicationCacheResults) GetL2Region() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.L2Region
}

// GetL2RegionOk returns a tuple with the L2Region field value
// and a boolean to check if the value has been set.
func (o *ApplicationCacheResults) GetL2RegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.L2Region, true
}

// SetL2Region sets field value
func (o *ApplicationCacheResults) SetL2Region(v string) {
	o.L2Region = v
}

func (o ApplicationCacheResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCacheResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["browser_cache_settings"] = o.BrowserCacheSettings
	toSerialize["browser_cache_settings_maximum_ttl"] = o.BrowserCacheSettingsMaximumTtl
	toSerialize["cdn_cache_settings"] = o.CdnCacheSettings
	toSerialize["cdn_cache_settings_maximum_ttl"] = o.CdnCacheSettingsMaximumTtl
	toSerialize["cache_by_query_string"] = o.CacheByQueryString
	toSerialize["query_string_fields"] = o.QueryStringFields
	toSerialize["enable_query_string_sort"] = o.EnableQueryStringSort
	toSerialize["cache_by_cookies"] = o.CacheByCookies
	toSerialize["cookie_names"] = o.CookieNames
	toSerialize["adaptive_delivery_action"] = o.AdaptiveDeliveryAction
	toSerialize["device_group"] = o.DeviceGroup
	toSerialize["enable_caching_for_post"] = o.EnableCachingForPost
	toSerialize["l2_caching_enabled"] = o.L2CachingEnabled
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
	toSerialize["enable_caching_for_options"] = o.EnableCachingForOptions
	toSerialize["enable_stale_cache"] = o.EnableStaleCache
	toSerialize["l2_region"] = o.L2Region
	return toSerialize, nil
}

type NullableApplicationCacheResults struct {
	value *ApplicationCacheResults
	isSet bool
}

func (v NullableApplicationCacheResults) Get() *ApplicationCacheResults {
	return v.value
}

func (v *NullableApplicationCacheResults) Set(val *ApplicationCacheResults) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCacheResults) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCacheResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCacheResults(val *ApplicationCacheResults) *NullableApplicationCacheResults {
	return &NullableApplicationCacheResults{value: val, isSet: true}
}

func (v NullableApplicationCacheResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCacheResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


