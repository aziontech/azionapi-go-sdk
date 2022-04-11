# PurgeUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Urls** | **[]string** |  | 
**Method** | **string** |  | 

## Methods

### NewPurgeUrlRequest

`func NewPurgeUrlRequest(urls []string, method string, ) *PurgeUrlRequest`

NewPurgeUrlRequest instantiates a new PurgeUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurgeUrlRequestWithDefaults

`func NewPurgeUrlRequestWithDefaults() *PurgeUrlRequest`

NewPurgeUrlRequestWithDefaults instantiates a new PurgeUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrls

`func (o *PurgeUrlRequest) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *PurgeUrlRequest) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *PurgeUrlRequest) SetUrls(v []string)`

SetUrls sets Urls field to given value.


### GetMethod

`func (o *PurgeUrlRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PurgeUrlRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PurgeUrlRequest) SetMethod(v string)`

SetMethod sets Method field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


