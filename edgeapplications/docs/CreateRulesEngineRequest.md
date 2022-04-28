# CreateRulesEngineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Criteria** | [**[][]RulesEngineCriteria**]([]RulesEngineCriteria.md) |  | 
**Behaviors** | [**[]CreateRulesEngineRequestBehaviors**](CreateRulesEngineRequestBehaviors.md) |  | 

## Methods

### NewCreateRulesEngineRequest

`func NewCreateRulesEngineRequest(name string, criteria [][]RulesEngineCriteria, behaviors []CreateRulesEngineRequestBehaviors, ) *CreateRulesEngineRequest`

NewCreateRulesEngineRequest instantiates a new CreateRulesEngineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRulesEngineRequestWithDefaults

`func NewCreateRulesEngineRequestWithDefaults() *CreateRulesEngineRequest`

NewCreateRulesEngineRequestWithDefaults instantiates a new CreateRulesEngineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRulesEngineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRulesEngineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRulesEngineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCriteria

`func (o *CreateRulesEngineRequest) GetCriteria() [][]RulesEngineCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *CreateRulesEngineRequest) GetCriteriaOk() (*[][]RulesEngineCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *CreateRulesEngineRequest) SetCriteria(v [][]RulesEngineCriteria)`

SetCriteria sets Criteria field to given value.


### GetBehaviors

`func (o *CreateRulesEngineRequest) GetBehaviors() []CreateRulesEngineRequestBehaviors`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *CreateRulesEngineRequest) GetBehaviorsOk() (*[]CreateRulesEngineRequestBehaviors, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *CreateRulesEngineRequest) SetBehaviors(v []CreateRulesEngineRequestBehaviors)`

SetBehaviors sets Behaviors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


