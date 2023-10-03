# RuleSetResultAll

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

### NewRuleSetResultAll

`func NewRuleSetResultAll() *RuleSetResultAll`

NewRuleSetResultAll instantiates a new RuleSetResultAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleSetResultAllWithDefaults

`func NewRuleSetResultAllWithDefaults() *RuleSetResultAll`

NewRuleSetResultAllWithDefaults instantiates a new RuleSetResultAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RuleSetResultAll) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleSetResultAll) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleSetResultAll) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RuleSetResultAll) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastEditor

`func (o *RuleSetResultAll) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *RuleSetResultAll) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *RuleSetResultAll) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *RuleSetResultAll) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetLastModified

`func (o *RuleSetResultAll) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *RuleSetResultAll) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *RuleSetResultAll) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *RuleSetResultAll) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetName

`func (o *RuleSetResultAll) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleSetResultAll) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleSetResultAll) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleSetResultAll) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *RuleSetResultAll) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RuleSetResultAll) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RuleSetResultAll) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RuleSetResultAll) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDescription

`func (o *RuleSetResultAll) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleSetResultAll) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleSetResultAll) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleSetResultAll) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBehaviors

`func (o *RuleSetResultAll) GetBehaviors() []Behaviors`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *RuleSetResultAll) GetBehaviorsOk() (*[]Behaviors, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *RuleSetResultAll) SetBehaviors(v []Behaviors)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *RuleSetResultAll) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetCriteria

`func (o *RuleSetResultAll) GetCriteria() [][]SSLVerificationStatusCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *RuleSetResultAll) GetCriteriaOk() (*[][]SSLVerificationStatusCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *RuleSetResultAll) SetCriteria(v [][]SSLVerificationStatusCriteria)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *RuleSetResultAll) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetOrder

`func (o *RuleSetResultAll) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RuleSetResultAll) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RuleSetResultAll) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RuleSetResultAll) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


