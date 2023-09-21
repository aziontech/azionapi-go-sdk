# PatchRulesEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Criteria** | Pointer to [**[][]RulesEngineCriteria**]([]RulesEngineCriteria.md) |  | [optional] 
**Behaviors** | Pointer to [**[]RulesEngineBehaviorEntry**](RulesEngineBehaviorEntry.md) |  | [optional] 

## Methods

### NewPatchRulesEngineRequest

`func NewPatchRulesEngineRequest() *PatchRulesEngineRequest`

NewPatchRulesEngineRequest instantiates a new PatchRulesEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchRulesEngineRequestWithDefaults

`func NewPatchRulesEngineRequestWithDefaults() *PatchRulesEngineRequest`

NewPatchRulesEngineRequestWithDefaults instantiates a new PatchRulesEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchRulesEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchRulesEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchRulesEngineRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchRulesEngineRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchRulesEngineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchRulesEngineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchRulesEngineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchRulesEngineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCriteria

`func (o *PatchRulesEngineRequest) GetCriteria() [][]RulesEngineCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *PatchRulesEngineRequest) GetCriteriaOk() (*[][]RulesEngineCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *PatchRulesEngineRequest) SetCriteria(v [][]RulesEngineCriteria)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *PatchRulesEngineRequest) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetBehaviors

`func (o *PatchRulesEngineRequest) GetBehaviors() []RulesEngineBehaviorEntry`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *PatchRulesEngineRequest) GetBehaviorsOk() (*[]RulesEngineBehaviorEntry, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *PatchRulesEngineRequest) SetBehaviors(v []RulesEngineBehaviorEntry)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *PatchRulesEngineRequest) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


