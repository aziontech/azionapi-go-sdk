# ApplicationPutInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EdgeFunctionId** | **int64** |  | 
**Args** | [**ApplicationInstancesResultsArgs**](ApplicationInstancesResultsArgs.md) |  | 

## Methods

### NewApplicationPutInstanceRequest

`func NewApplicationPutInstanceRequest(name string, edgeFunctionId int64, args ApplicationInstancesResultsArgs, ) *ApplicationPutInstanceRequest`

NewApplicationPutInstanceRequest instantiates a new ApplicationPutInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPutInstanceRequestWithDefaults

`func NewApplicationPutInstanceRequestWithDefaults() *ApplicationPutInstanceRequest`

NewApplicationPutInstanceRequestWithDefaults instantiates a new ApplicationPutInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationPutInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationPutInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationPutInstanceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEdgeFunctionId

`func (o *ApplicationPutInstanceRequest) GetEdgeFunctionId() int64`

GetEdgeFunctionId returns the EdgeFunctionId field if non-nil, zero value otherwise.

### GetEdgeFunctionIdOk

`func (o *ApplicationPutInstanceRequest) GetEdgeFunctionIdOk() (*int64, bool)`

GetEdgeFunctionIdOk returns a tuple with the EdgeFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctionId

`func (o *ApplicationPutInstanceRequest) SetEdgeFunctionId(v int64)`

SetEdgeFunctionId sets EdgeFunctionId field to given value.


### GetArgs

`func (o *ApplicationPutInstanceRequest) GetArgs() ApplicationInstancesResultsArgs`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ApplicationPutInstanceRequest) GetArgsOk() (*ApplicationInstancesResultsArgs, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ApplicationPutInstanceRequest) SetArgs(v ApplicationInstancesResultsArgs)`

SetArgs sets Args field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


