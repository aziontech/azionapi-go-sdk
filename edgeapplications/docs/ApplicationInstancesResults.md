# ApplicationInstancesResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**EdgeFunctionId** | **int64** |  | 
**Name** | **string** |  | 
**Args** | [**ApplicationInstancesResultsArgs**](ApplicationInstancesResultsArgs.md) |  | 

## Methods

### NewApplicationInstancesResults

`func NewApplicationInstancesResults(id int64, edgeFunctionId int64, name string, args ApplicationInstancesResultsArgs, ) *ApplicationInstancesResults`

NewApplicationInstancesResults instantiates a new ApplicationInstancesResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationInstancesResultsWithDefaults

`func NewApplicationInstancesResultsWithDefaults() *ApplicationInstancesResults`

NewApplicationInstancesResultsWithDefaults instantiates a new ApplicationInstancesResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationInstancesResults) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationInstancesResults) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationInstancesResults) SetId(v int64)`

SetId sets Id field to given value.


### GetEdgeFunctionId

`func (o *ApplicationInstancesResults) GetEdgeFunctionId() int64`

GetEdgeFunctionId returns the EdgeFunctionId field if non-nil, zero value otherwise.

### GetEdgeFunctionIdOk

`func (o *ApplicationInstancesResults) GetEdgeFunctionIdOk() (*int64, bool)`

GetEdgeFunctionIdOk returns a tuple with the EdgeFunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctionId

`func (o *ApplicationInstancesResults) SetEdgeFunctionId(v int64)`

SetEdgeFunctionId sets EdgeFunctionId field to given value.


### GetName

`func (o *ApplicationInstancesResults) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationInstancesResults) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationInstancesResults) SetName(v string)`

SetName sets Name field to given value.


### GetArgs

`func (o *ApplicationInstancesResults) GetArgs() ApplicationInstancesResultsArgs`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ApplicationInstancesResults) GetArgsOk() (*ApplicationInstancesResultsArgs, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ApplicationInstancesResults) SetArgs(v ApplicationInstancesResultsArgs)`

SetArgs sets Args field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


