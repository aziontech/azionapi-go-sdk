# Behaviors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Argument** | Pointer to [**BehaviorsArgument**](BehaviorsArgument.md) |  | [optional] 

## Methods

### NewBehaviors

`func NewBehaviors() *Behaviors`

NewBehaviors instantiates a new Behaviors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorsWithDefaults

`func NewBehaviorsWithDefaults() *Behaviors`

NewBehaviorsWithDefaults instantiates a new Behaviors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Behaviors) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Behaviors) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Behaviors) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Behaviors) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArgument

`func (o *Behaviors) GetArgument() BehaviorsArgument`

GetArgument returns the Argument field if non-nil, zero value otherwise.

### GetArgumentOk

`func (o *Behaviors) GetArgumentOk() (*BehaviorsArgument, bool)`

GetArgumentOk returns a tuple with the Argument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgument

`func (o *Behaviors) SetArgument(v BehaviorsArgument)`

SetArgument sets Argument field to given value.

### HasArgument

`func (o *Behaviors) HasArgument() bool`

HasArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


