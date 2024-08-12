# CreateOriginsRequestAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**IsActive** | Pointer to **bool** |  | [optional] 
**Weight** | Pointer to **int64** |  | [optional] 
**ServerRole** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOriginsRequestAddresses

`func NewCreateOriginsRequestAddresses(address string, ) *CreateOriginsRequestAddresses`

NewCreateOriginsRequestAddresses instantiates a new CreateOriginsRequestAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOriginsRequestAddressesWithDefaults

`func NewCreateOriginsRequestAddressesWithDefaults() *CreateOriginsRequestAddresses`

NewCreateOriginsRequestAddressesWithDefaults instantiates a new CreateOriginsRequestAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CreateOriginsRequestAddresses) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateOriginsRequestAddresses) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateOriginsRequestAddresses) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetIsActive

`func (o *CreateOriginsRequestAddresses) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateOriginsRequestAddresses) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateOriginsRequestAddresses) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CreateOriginsRequestAddresses) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetWeight

`func (o *CreateOriginsRequestAddresses) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CreateOriginsRequestAddresses) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CreateOriginsRequestAddresses) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *CreateOriginsRequestAddresses) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetServerRole

`func (o *CreateOriginsRequestAddresses) GetServerRole() string`

GetServerRole returns the ServerRole field if non-nil, zero value otherwise.

### GetServerRoleOk

`func (o *CreateOriginsRequestAddresses) GetServerRoleOk() (*string, bool)`

GetServerRoleOk returns a tuple with the ServerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRole

`func (o *CreateOriginsRequestAddresses) SetServerRole(v string)`

SetServerRole sets ServerRole field to given value.

### HasServerRole

`func (o *CreateOriginsRequestAddresses) HasServerRole() bool`

HasServerRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


