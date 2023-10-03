# RuleSetResultResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LastEditor** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Behaviors** | Pointer to [**[]Behaviors**](Behaviors.md) |  | [optional] 
**Criteria** | Pointer to [**[][]SSLVerificationStatusCriteria**]([]SSLVerificationStatusCriteria.md) |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 

## Methods

### NewRuleSetResultResults

`func NewRuleSetResultResults() *RuleSetResultResults`

NewRuleSetResultResults instantiates a new RuleSetResultResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleSetResultResultsWithDefaults

`func NewRuleSetResultResultsWithDefaults() *RuleSetResultResults`

NewRuleSetResultResultsWithDefaults instantiates a new RuleSetResultResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleSetResultResults) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleSetResultResults) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleSetResultResults) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RuleSetResultResults) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastEditor

`func (o *RuleSetResultResults) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *RuleSetResultResults) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *RuleSetResultResults) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *RuleSetResultResults) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetLastModified

`func (o *RuleSetResultResults) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *RuleSetResultResults) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *RuleSetResultResults) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *RuleSetResultResults) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetName

`func (o *RuleSetResultResults) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleSetResultResults) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleSetResultResults) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleSetResultResults) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *RuleSetResultResults) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RuleSetResultResults) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RuleSetResultResults) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RuleSetResultResults) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDescription

`func (o *RuleSetResultResults) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleSetResultResults) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleSetResultResults) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleSetResultResults) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBehaviors

`func (o *RuleSetResultResults) GetBehaviors() []Behaviors`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *RuleSetResultResults) GetBehaviorsOk() (*[]Behaviors, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *RuleSetResultResults) SetBehaviors(v []Behaviors)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *RuleSetResultResults) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetCriteria

`func (o *RuleSetResultResults) GetCriteria() [][]SSLVerificationStatusCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *RuleSetResultResults) GetCriteriaOk() (*[][]SSLVerificationStatusCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *RuleSetResultResults) SetCriteria(v [][]SSLVerificationStatusCriteria)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *RuleSetResultResults) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetOrder

`func (o *RuleSetResultResults) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RuleSetResultResults) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RuleSetResultResults) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RuleSetResultResults) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


