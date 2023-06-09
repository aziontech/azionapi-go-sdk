# UpdateServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 

## Methods

### NewUpdateServiceRequest

`func NewUpdateServiceRequest() *UpdateServiceRequest`

NewUpdateServiceRequest instantiates a new UpdateServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceRequestWithDefaults

`func NewUpdateServiceRequestWithDefaults() *UpdateServiceRequest`

NewUpdateServiceRequestWithDefaults instantiates a new UpdateServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateServiceRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateServiceRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateServiceRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateServiceRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetName

`func (o *UpdateServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServiceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateServiceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVariables

`func (o *UpdateServiceRequest) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *UpdateServiceRequest) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *UpdateServiceRequest) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *UpdateServiceRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


