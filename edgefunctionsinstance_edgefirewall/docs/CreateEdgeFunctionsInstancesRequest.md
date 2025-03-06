# CreateEdgeFunctionsInstancesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**EdgeFunction** | Pointer to **int64** |  | [optional] 
**JsonArgs** | Pointer to [**EdgeFunctionsInstanceJsonArgs**](EdgeFunctionsInstanceJsonArgs.md) |  | [optional] 

## Methods

### NewCreateEdgeFunctionsInstancesRequest

`func NewCreateEdgeFunctionsInstancesRequest() *CreateEdgeFunctionsInstancesRequest`

NewCreateEdgeFunctionsInstancesRequest instantiates a new CreateEdgeFunctionsInstancesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEdgeFunctionsInstancesRequestWithDefaults

`func NewCreateEdgeFunctionsInstancesRequestWithDefaults() *CreateEdgeFunctionsInstancesRequest`

NewCreateEdgeFunctionsInstancesRequestWithDefaults instantiates a new CreateEdgeFunctionsInstancesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEdgeFunctionsInstancesRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEdgeFunctionsInstancesRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEdgeFunctionsInstancesRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateEdgeFunctionsInstancesRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEdgeFunction

`func (o *CreateEdgeFunctionsInstancesRequest) GetEdgeFunction() int64`

GetEdgeFunction returns the EdgeFunction field if non-nil, zero value otherwise.

### GetEdgeFunctionOk

`func (o *CreateEdgeFunctionsInstancesRequest) GetEdgeFunctionOk() (*int64, bool)`

GetEdgeFunctionOk returns a tuple with the EdgeFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunction

`func (o *CreateEdgeFunctionsInstancesRequest) SetEdgeFunction(v int64)`

SetEdgeFunction sets EdgeFunction field to given value.

### HasEdgeFunction

`func (o *CreateEdgeFunctionsInstancesRequest) HasEdgeFunction() bool`

HasEdgeFunction returns a boolean if a field has been set.

### GetJsonArgs

`func (o *CreateEdgeFunctionsInstancesRequest) GetJsonArgs() EdgeFunctionsInstanceJsonArgs`

GetJsonArgs returns the JsonArgs field if non-nil, zero value otherwise.

### GetJsonArgsOk

`func (o *CreateEdgeFunctionsInstancesRequest) GetJsonArgsOk() (*EdgeFunctionsInstanceJsonArgs, bool)`

GetJsonArgsOk returns a tuple with the JsonArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonArgs

`func (o *CreateEdgeFunctionsInstancesRequest) SetJsonArgs(v EdgeFunctionsInstanceJsonArgs)`

SetJsonArgs sets JsonArgs field to given value.

### HasJsonArgs

`func (o *CreateEdgeFunctionsInstancesRequest) HasJsonArgs() bool`

HasJsonArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


