# VariableCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Value** | **string** |  | 
**Secret** | Pointer to **bool** |  | [optional] 

## Methods

### NewVariableCreate

`func NewVariableCreate(key string, value string, ) *VariableCreate`

NewVariableCreate instantiates a new VariableCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableCreateWithDefaults

`func NewVariableCreateWithDefaults() *VariableCreate`

NewVariableCreateWithDefaults instantiates a new VariableCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VariableCreate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VariableCreate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VariableCreate) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *VariableCreate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableCreate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableCreate) SetValue(v string)`

SetValue sets Value field to given value.


### GetSecret

`func (o *VariableCreate) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *VariableCreate) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *VariableCreate) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *VariableCreate) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


