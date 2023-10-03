# RuleSetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Behaviors** | Pointer to [**[]Behaviors**](Behaviors.md) |  | [optional] 
**Criteria** | Pointer to [**[][]SSLVerificationStatusCriteria**]([]SSLVerificationStatusCriteria.md) |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**LastEditor** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Order** | Pointer to **int64** |  | [optional] 

## Methods

### NewRuleSetResponse

`func NewRuleSetResponse() *RuleSetResponse`

NewRuleSetResponse instantiates a new RuleSetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleSetResponseWithDefaults

`func NewRuleSetResponseWithDefaults() *RuleSetResponse`

NewRuleSetResponseWithDefaults instantiates a new RuleSetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RuleSetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleSetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleSetResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleSetResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *RuleSetResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *RuleSetResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *RuleSetResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *RuleSetResponse) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetBehaviors

`func (o *RuleSetResponse) GetBehaviors() []Behaviors`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *RuleSetResponse) GetBehaviorsOk() (*[]Behaviors, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *RuleSetResponse) SetBehaviors(v []Behaviors)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *RuleSetResponse) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetCriteria

`func (o *RuleSetResponse) GetCriteria() [][]SSLVerificationStatusCriteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *RuleSetResponse) GetCriteriaOk() (*[][]SSLVerificationStatusCriteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *RuleSetResponse) SetCriteria(v [][]SSLVerificationStatusCriteria)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *RuleSetResponse) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetLastModified

`func (o *RuleSetResponse) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *RuleSetResponse) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *RuleSetResponse) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *RuleSetResponse) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetLastEditor

`func (o *RuleSetResponse) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *RuleSetResponse) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *RuleSetResponse) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *RuleSetResponse) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetId

`func (o *RuleSetResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleSetResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleSetResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RuleSetResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrder

`func (o *RuleSetResponse) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RuleSetResponse) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RuleSetResponse) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RuleSetResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


