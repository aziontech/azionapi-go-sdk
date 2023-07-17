# UpdateRulesEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Criteria** | [**[][]RulesEngineCriteria**]([]RulesEngineCriteria.md) |  | 
**Behaviors** | [**[]RulesEngineBehavior**](RulesEngineBehavior.md) |  | 

## Methods

### NewUpdateRulesEngineRequest

`func NewUpdateRulesEngineRequest(name string, criteria [][]RulesEngineCriteria, behaviors []RulesEngineBehavior, ) *UpdateRulesEngineRequest`

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

`func (o *UpdateRulesEngineRequest) GetBehaviors() []RulesEngineBehavior`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *UpdateRulesEngineRequest) GetBehaviorsOk() (*[]RulesEngineBehavior, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *UpdateRulesEngineRequest) SetBehaviors(v []RulesEngineBehavior)`

SetBehaviors sets Behaviors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


