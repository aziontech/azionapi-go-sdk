# RulesEngineResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Phase** | **string** |  | 
**Behaviors** | Pointer to [**[]CreateRulesEngineRequestBehaviors**](CreateRulesEngineRequestBehaviors.md) |  | [optional] 
**Criteria** | [**[][]RulesEngineCriteria**]([]RulesEngineCriteria.md) |  | 
**IsActive** | **bool** |  | 
**Order** | **int64** |  | 

## Methods

### NewRulesEngineResultResponse

`func NewRulesEngineResultResponse(id int64, name string, phase string, criteria [][]RulesEngineCriteria, isActive bool, order int64, ) *RulesEngineResultResponse`

NewRulesEngineResultResponse instantiates a new RulesEngineResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesEngineResultResponseWithDefaults

`func NewRulesEngineResultResponseWithDefaults() *RulesEngineResultResponse`

NewRulesEngineResultResponseWithDefaults instantiates a new RulesEngineResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RulesEngineResultResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RulesEngineResultResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RulesEngineResultResponse) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *RulesEngineResultResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RulesEngineResultResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RulesEngineResultResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPhase

`func (o *RulesEngineResultResponse) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *RulesEngineResultResponse) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *RulesEngineResultResponse) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetBehaviors

`func (o *RulesEngineResultResponse) GetBehaviors() []CreateRulesEngineRequestBehaviors`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *RulesEngineResultResponse) GetBehaviorsOk() (*[]CreateRulesEngineRequestBehaviors, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *RulesEngineResultResponse) SetBehaviors(v []CreateRulesEngineRequestBehaviors)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *RulesEngineResultResponse) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetCriteria

`func (o *RulesEngineResultResponse) GetCriteria() [][]RulesEngineCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *RulesEngineResultResponse) GetCriteriaOk() (*[][]RulesEngineCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *RulesEngineResultResponse) SetCriteria(v [][]RulesEngineCriteria)`

SetCriteria sets Criteria field to given value.


### GetIsActive

`func (o *RulesEngineResultResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RulesEngineResultResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RulesEngineResultResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetOrder

`func (o *RulesEngineResultResponse) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RulesEngineResultResponse) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RulesEngineResultResponse) SetOrder(v int64)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


