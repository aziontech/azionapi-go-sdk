# CreateRuleSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Behaviors** | Pointer to [**[]Behaviors**](Behaviors.md) |  | [optional] 
**Criteria** | Pointer to [**[][]SSLVerificationStatusCriteria**]([]SSLVerificationStatusCriteria.md) |  | [optional] 

## Methods

### NewCreateRuleSetRequest

`func NewCreateRuleSetRequest() *CreateRuleSetRequest`

NewCreateRuleSetRequest instantiates a new CreateRuleSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRuleSetRequestWithDefaults

`func NewCreateRuleSetRequestWithDefaults() *CreateRuleSetRequest`

NewCreateRuleSetRequestWithDefaults instantiates a new CreateRuleSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRuleSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRuleSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRuleSetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateRuleSetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *CreateRuleSetRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateRuleSetRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateRuleSetRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CreateRuleSetRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetBehaviors

`func (o *CreateRuleSetRequest) GetBehaviors() []Behaviors`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *CreateRuleSetRequest) GetBehaviorsOk() (*[]Behaviors, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *CreateRuleSetRequest) SetBehaviors(v []Behaviors)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *CreateRuleSetRequest) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetCriteria

`func (o *CreateRuleSetRequest) GetCriteria() [][]SSLVerificationStatusCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *CreateRuleSetRequest) GetCriteriaOk() (*[][]SSLVerificationStatusCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *CreateRuleSetRequest) SetCriteria(v [][]SSLVerificationStatusCriteria)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *CreateRuleSetRequest) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


