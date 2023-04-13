# DeviceGroupsResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**UserAgent** | **string** |  | 

## Methods

### NewDeviceGroupsResultResponse

`func NewDeviceGroupsResultResponse(name string, userAgent string, ) *DeviceGroupsResultResponse`

NewDeviceGroupsResultResponse instantiates a new DeviceGroupsResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceGroupsResultResponseWithDefaults

`func NewDeviceGroupsResultResponseWithDefaults() *DeviceGroupsResultResponse`

NewDeviceGroupsResultResponseWithDefaults instantiates a new DeviceGroupsResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceGroupsResultResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceGroupsResultResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceGroupsResultResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceGroupsResultResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeviceGroupsResultResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceGroupsResultResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceGroupsResultResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUserAgent

`func (o *DeviceGroupsResultResponse) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *DeviceGroupsResultResponse) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *DeviceGroupsResultResponse) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


