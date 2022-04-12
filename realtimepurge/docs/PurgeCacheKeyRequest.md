# PurgeCacheKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urls** | **[]string** |  | 
**Method** | **string** |  | 
**Layer** | **string** |  | 

## Methods

### NewPurgeCacheKeyRequest

`func NewPurgeCacheKeyRequest(urls []string, method string, layer string, ) *PurgeCacheKeyRequest`

NewPurgeCacheKeyRequest instantiates a new PurgeCacheKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeCacheKeyRequestWithDefaults

`func NewPurgeCacheKeyRequestWithDefaults() *PurgeCacheKeyRequest`

NewPurgeCacheKeyRequestWithDefaults instantiates a new PurgeCacheKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrls

`func (o *PurgeCacheKeyRequest) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *PurgeCacheKeyRequest) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *PurgeCacheKeyRequest) SetUrls(v []string)`

SetUrls sets Urls field to given value.


### GetMethod

`func (o *PurgeCacheKeyRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PurgeCacheKeyRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PurgeCacheKeyRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetLayer

`func (o *PurgeCacheKeyRequest) GetLayer() string`

GetLayer returns the Layer field if non-nil, zero value otherwise.

### GetLayerOk

`func (o *PurgeCacheKeyRequest) GetLayerOk() (*string, bool)`

GetLayerOk returns a tuple with the Layer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer

`func (o *PurgeCacheKeyRequest) SetLayer(v string)`

SetLayer sets Layer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


