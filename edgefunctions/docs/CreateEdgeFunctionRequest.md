# CreateEdgeFunctionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**JsonArgs** | Pointer to **interface{}** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateEdgeFunctionRequest

`func NewCreateEdgeFunctionRequest() *CreateEdgeFunctionRequest`

NewCreateEdgeFunctionRequest instantiates a new CreateEdgeFunctionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEdgeFunctionRequestWithDefaults

`func NewCreateEdgeFunctionRequestWithDefaults() *CreateEdgeFunctionRequest`

NewCreateEdgeFunctionRequestWithDefaults instantiates a new CreateEdgeFunctionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEdgeFunctionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEdgeFunctionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEdgeFunctionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateEdgeFunctionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLanguage

`func (o *CreateEdgeFunctionRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CreateEdgeFunctionRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CreateEdgeFunctionRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CreateEdgeFunctionRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCode

`func (o *CreateEdgeFunctionRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateEdgeFunctionRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateEdgeFunctionRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateEdgeFunctionRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetJsonArgs

`func (o *CreateEdgeFunctionRequest) GetJsonArgs() interface{}`

GetJsonArgs returns the JsonArgs field if non-nil, zero value otherwise.

### GetJsonArgsOk

`func (o *CreateEdgeFunctionRequest) GetJsonArgsOk() (*interface{}, bool)`

GetJsonArgsOk returns a tuple with the JsonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonArgs

`func (o *CreateEdgeFunctionRequest) SetJsonArgs(v interface{})`

SetJsonArgs sets JsonArgs field to given value.

### HasJsonArgs

`func (o *CreateEdgeFunctionRequest) HasJsonArgs() bool`

HasJsonArgs returns a boolean if a field has been set.

### SetJsonArgsNil

`func (o *CreateEdgeFunctionRequest) SetJsonArgsNil(b bool)`

 SetJsonArgsNil sets the value for JsonArgs to be an explicit nil

### UnsetJsonArgs
`func (o *CreateEdgeFunctionRequest) UnsetJsonArgs()`

UnsetJsonArgs ensures that no value is present for JsonArgs, not even an explicit nil
### GetActive

`func (o *CreateEdgeFunctionRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateEdgeFunctionRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateEdgeFunctionRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateEdgeFunctionRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


