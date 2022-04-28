# ApplicationCreateInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EdgeFunctionId** | **int64** |  | 
**Args** | **interface{}** |  | 

## Methods

### NewApplicationCreateInstanceRequest

`func NewApplicationCreateInstanceRequest(name string, edgeFunctionId int64, args interface{}, ) *ApplicationCreateInstanceRequest`

NewApplicationCreateInstanceRequest instantiates a new ApplicationCreateInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCreateInstanceRequestWithDefaults

`func NewApplicationCreateInstanceRequestWithDefaults() *ApplicationCreateInstanceRequest`

NewApplicationCreateInstanceRequestWithDefaults instantiates a new ApplicationCreateInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationCreateInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCreateInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCreateInstanceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEdgeFunctionId

`func (o *ApplicationCreateInstanceRequest) GetEdgeFunctionId() int64`

GetEdgeFunctionId returns the EdgeFunctionId field if non-nil, zero value otherwise.

### GetEdgeFunctionIdOk

`func (o *ApplicationCreateInstanceRequest) GetEdgeFunctionIdOk() (*int64, bool)`

GetEdgeFunctionIdOk returns a tuple with the EdgeFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctionId

`func (o *ApplicationCreateInstanceRequest) SetEdgeFunctionId(v int64)`

SetEdgeFunctionId sets EdgeFunctionId field to given value.


### GetArgs

`func (o *ApplicationCreateInstanceRequest) GetArgs() interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ApplicationCreateInstanceRequest) GetArgsOk() (*interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ApplicationCreateInstanceRequest) SetArgs(v interface{})`

SetArgs sets Args field to given value.


### SetArgsNil

`func (o *ApplicationCreateInstanceRequest) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *ApplicationCreateInstanceRequest) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


