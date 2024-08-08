# OriginsResultResponseAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Weight** | **NullableInt64** |  | 
**ServerRole** | **string** |  | 
**IsActive** | **bool** |  | 

## Methods

### NewOriginsResultResponseAddresses

`func NewOriginsResultResponseAddresses(address string, weight NullableInt64, serverRole string, isActive bool, ) *OriginsResultResponseAddresses`

NewOriginsResultResponseAddresses instantiates a new OriginsResultResponseAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginsResultResponseAddressesWithDefaults

`func NewOriginsResultResponseAddressesWithDefaults() *OriginsResultResponseAddresses`

NewOriginsResultResponseAddressesWithDefaults instantiates a new OriginsResultResponseAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *OriginsResultResponseAddresses) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OriginsResultResponseAddresses) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OriginsResultResponseAddresses) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetWeight

`func (o *OriginsResultResponseAddresses) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *OriginsResultResponseAddresses) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *OriginsResultResponseAddresses) SetWeight(v int64)`

SetWeight sets Weight field to given value.


### SetWeightNil

`func (o *OriginsResultResponseAddresses) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *OriginsResultResponseAddresses) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetServerRole

`func (o *OriginsResultResponseAddresses) GetServerRole() string`

GetServerRole returns the ServerRole field if non-nil, zero value otherwise.

### GetServerRoleOk

`func (o *OriginsResultResponseAddresses) GetServerRoleOk() (*string, bool)`

GetServerRoleOk returns a tuple with the ServerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRole

`func (o *OriginsResultResponseAddresses) SetServerRole(v string)`

SetServerRole sets ServerRole field to given value.


### GetIsActive

`func (o *OriginsResultResponseAddresses) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *OriginsResultResponseAddresses) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *OriginsResultResponseAddresses) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


