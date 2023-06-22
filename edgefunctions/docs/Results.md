# Results

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**JsonArgs** | Pointer to **interface{}** |  | [optional] 
**FunctionToRun** | Pointer to **string** |  | [optional] 
**InitiatorType** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**LastEditor** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **string** |  | [optional] 
**ReferenceCount** | Pointer to **int64** |  | [optional] 
**IsProprietaryCode** | Pointer to **bool** |  | [optional] 

## Methods

### NewResults

`func NewResults() *Results`

NewResults instantiates a new Results object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultsWithDefaults

`func NewResultsWithDefaults() *Results`

NewResultsWithDefaults instantiates a new Results object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Results) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Results) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Results) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Results) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Results) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Results) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Results) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Results) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLanguage

`func (o *Results) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Results) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Results) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Results) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCode

`func (o *Results) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Results) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Results) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Results) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetJsonArgs

`func (o *Results) GetJsonArgs() interface{}`

GetJsonArgs returns the JsonArgs field if non-nil, zero value otherwise.

### GetJsonArgsOk

`func (o *Results) GetJsonArgsOk() (*interface{}, bool)`

GetJsonArgsOk returns a tuple with the JsonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonArgs

`func (o *Results) SetJsonArgs(v interface{})`

SetJsonArgs sets JsonArgs field to given value.

### HasJsonArgs

`func (o *Results) HasJsonArgs() bool`

HasJsonArgs returns a boolean if a field has been set.

### SetJsonArgsNil

`func (o *Results) SetJsonArgsNil(b bool)`

 SetJsonArgsNil sets the value for JsonArgs to be an explicit nil

### UnsetJsonArgs
`func (o *Results) UnsetJsonArgs()`

UnsetJsonArgs ensures that no value is present for JsonArgs, not even an explicit nil
### GetFunctionToRun

`func (o *Results) GetFunctionToRun() string`

GetFunctionToRun returns the FunctionToRun field if non-nil, zero value otherwise.

### GetFunctionToRunOk

`func (o *Results) GetFunctionToRunOk() (*string, bool)`

GetFunctionToRunOk returns a tuple with the FunctionToRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionToRun

`func (o *Results) SetFunctionToRun(v string)`

SetFunctionToRun sets FunctionToRun field to given value.

### HasFunctionToRun

`func (o *Results) HasFunctionToRun() bool`

HasFunctionToRun returns a boolean if a field has been set.

### GetInitiatorType

`func (o *Results) GetInitiatorType() string`

GetInitiatorType returns the InitiatorType field if non-nil, zero value otherwise.

### GetInitiatorTypeOk

`func (o *Results) GetInitiatorTypeOk() (*string, bool)`

GetInitiatorTypeOk returns a tuple with the InitiatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorType

`func (o *Results) SetInitiatorType(v string)`

SetInitiatorType sets InitiatorType field to given value.

### HasInitiatorType

`func (o *Results) HasInitiatorType() bool`

HasInitiatorType returns a boolean if a field has been set.

### GetActive

`func (o *Results) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Results) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Results) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Results) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLastEditor

`func (o *Results) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *Results) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *Results) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *Results) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetModified

`func (o *Results) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Results) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Results) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Results) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetReferenceCount

`func (o *Results) GetReferenceCount() int64`

GetReferenceCount returns the ReferenceCount field if non-nil, zero value otherwise.

### GetReferenceCountOk

`func (o *Results) GetReferenceCountOk() (*int64, bool)`

GetReferenceCountOk returns a tuple with the ReferenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCount

`func (o *Results) SetReferenceCount(v int64)`

SetReferenceCount sets ReferenceCount field to given value.

### HasReferenceCount

`func (o *Results) HasReferenceCount() bool`

HasReferenceCount returns a boolean if a field has been set.

### GetIsProprietaryCode

`func (o *Results) GetIsProprietaryCode() bool`

GetIsProprietaryCode returns the IsProprietaryCode field if non-nil, zero value otherwise.

### GetIsProprietaryCodeOk

`func (o *Results) GetIsProprietaryCodeOk() (*bool, bool)`

GetIsProprietaryCodeOk returns a tuple with the IsProprietaryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProprietaryCode

`func (o *Results) SetIsProprietaryCode(v bool)`

SetIsProprietaryCode sets IsProprietaryCode field to given value.

### HasIsProprietaryCode

`func (o *Results) HasIsProprietaryCode() bool`

HasIsProprietaryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


