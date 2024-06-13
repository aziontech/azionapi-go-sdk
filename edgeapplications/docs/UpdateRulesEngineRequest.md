# UpdateRulesEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Order** | Pointer to **int64** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Criteria** | [**[][]RulesEngineCriteria**]([]RulesEngineCriteria.md) |  | 
**Behaviors** | [**[]RulesEngineBehaviorEntry**](RulesEngineBehaviorEntry.md) |  | 

## Methods

### NewUpdateRulesEngineRequest

`func NewUpdateRulesEngineRequest(name string, criteria [][]RulesEngineCriteria, behaviors []RulesEngineBehaviorEntry, ) *UpdateRulesEngineRequest`

NewUpdateRulesEngineRequest instantiates a new UpdateRulesEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRulesEngineRequestWithDefaults

`func NewUpdateRulesEngineRequestWithDefaults() *UpdateRulesEngineRequest`

NewUpdateRulesEngineRequestWithDefaults instantiates a new UpdateRulesEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRulesEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRulesEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRulesEngineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOrder

`func (o *UpdateRulesEngineRequest) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *UpdateRulesEngineRequest) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *UpdateRulesEngineRequest) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *UpdateRulesEngineRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateRulesEngineRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateRulesEngineRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateRulesEngineRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateRulesEngineRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateRulesEngineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateRulesEngineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateRulesEngineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateRulesEngineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCriteria

`func (o *UpdateRulesEngineRequest) GetCriteria() [][]RulesEngineCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *UpdateRulesEngineRequest) GetCriteriaOk() (*[][]RulesEngineCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *UpdateRulesEngineRequest) SetCriteria(v [][]RulesEngineCriteria)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *UpdateRulesEngineRequest) GetBehaviors() []RulesEngineBehaviorEntry`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *UpdateRulesEngineRequest) GetBehaviorsOk() (*[]RulesEngineBehaviorEntry, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *UpdateRulesEngineRequest) SetBehaviors(v []RulesEngineBehaviorEntry)`

SetBehaviors sets Behaviors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


