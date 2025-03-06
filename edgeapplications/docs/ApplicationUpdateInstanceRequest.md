# ApplicationUpdateInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**EdgeFunctionId** | **NullableInt64** |  | 
**Args** | [**ApplicationInstancesResultsArgs**](ApplicationInstancesResultsArgs.md) |  | 

## Methods

### NewApplicationUpdateInstanceRequest

`func NewApplicationUpdateInstanceRequest(name NullableString, edgeFunctionId NullableInt64, args ApplicationInstancesResultsArgs, ) *ApplicationUpdateInstanceRequest`

NewApplicationUpdateInstanceRequest instantiates a new ApplicationUpdateInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationUpdateInstanceRequestWithDefaults

`func NewApplicationUpdateInstanceRequestWithDefaults() *ApplicationUpdateInstanceRequest`

NewApplicationUpdateInstanceRequestWithDefaults instantiates a new ApplicationUpdateInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationUpdateInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationUpdateInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationUpdateInstanceRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ApplicationUpdateInstanceRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ApplicationUpdateInstanceRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEdgeFunctionId

`func (o *ApplicationUpdateInstanceRequest) GetEdgeFunctionId() int64`

GetEdgeFunctionId returns the EdgeFunctionId field if non-nil, zero value otherwise.

### GetEdgeFunctionIdOk

`func (o *ApplicationUpdateInstanceRequest) GetEdgeFunctionIdOk() (*int64, bool)`

GetEdgeFunctionIdOk returns a tuple with the EdgeFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctionId

`func (o *ApplicationUpdateInstanceRequest) SetEdgeFunctionId(v int64)`

SetEdgeFunctionId sets EdgeFunctionId field to given value.


### SetEdgeFunctionIdNil

`func (o *ApplicationUpdateInstanceRequest) SetEdgeFunctionIdNil(b bool)`

 SetEdgeFunctionIdNil sets the value for EdgeFunctionId to be an explicit nil

### UnsetEdgeFunctionId
`func (o *ApplicationUpdateInstanceRequest) UnsetEdgeFunctionId()`

UnsetEdgeFunctionId ensures that no value is present for EdgeFunctionId, not even an explicit nil
### GetArgs

`func (o *ApplicationUpdateInstanceRequest) GetArgs() ApplicationInstancesResultsArgs`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ApplicationUpdateInstanceRequest) GetArgsOk() (*ApplicationInstancesResultsArgs, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ApplicationUpdateInstanceRequest) SetArgs(v ApplicationInstancesResultsArgs)`

SetArgs sets Args field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


