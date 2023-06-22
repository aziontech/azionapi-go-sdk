# CreateDeviceGroupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**UserAgent** | **string** |  | 
**Addresses** | **string** |  | 

## Methods

### NewCreateDeviceGroupsRequest

`func NewCreateDeviceGroupsRequest(userAgent string, addresses string, ) *CreateDeviceGroupsRequest`

NewCreateDeviceGroupsRequest instantiates a new CreateDeviceGroupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceGroupsRequestWithDefaults

`func NewCreateDeviceGroupsRequestWithDefaults() *CreateDeviceGroupsRequest`

NewCreateDeviceGroupsRequestWithDefaults instantiates a new CreateDeviceGroupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDeviceGroupsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDeviceGroupsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDeviceGroupsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateDeviceGroupsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserAgent

`func (o *CreateDeviceGroupsRequest) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *CreateDeviceGroupsRequest) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *CreateDeviceGroupsRequest) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.


### GetAddresses

`func (o *CreateDeviceGroupsRequest) GetAddresses() string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CreateDeviceGroupsRequest) GetAddressesOk() (*string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CreateDeviceGroupsRequest) SetAddresses(v string)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


