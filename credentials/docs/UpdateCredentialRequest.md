# UpdateCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Name** | **string** |  | 
**Status** | **bool** |  | 

## Methods

### NewUpdateCredentialRequest

`func NewUpdateCredentialRequest(description string, name string, status bool, ) *UpdateCredentialRequest`

NewUpdateCredentialRequest instantiates a new UpdateCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCredentialRequestWithDefaults

`func NewUpdateCredentialRequestWithDefaults() *UpdateCredentialRequest`

NewUpdateCredentialRequestWithDefaults instantiates a new UpdateCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateCredentialRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCredentialRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCredentialRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *UpdateCredentialRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCredentialRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCredentialRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *UpdateCredentialRequest) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateCredentialRequest) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateCredentialRequest) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


