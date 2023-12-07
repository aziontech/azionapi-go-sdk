# ResponseDeleteBucketData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [readonly] 
**EdgeAccess** | [**EdgeAccessEnum**](EdgeAccessEnum.md) |  | 

## Methods

### NewResponseDeleteBucketData

`func NewResponseDeleteBucketData(name string, edgeAccess EdgeAccessEnum, ) *ResponseDeleteBucketData`

NewResponseDeleteBucketData instantiates a new ResponseDeleteBucketData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeleteBucketDataWithDefaults

`func NewResponseDeleteBucketDataWithDefaults() *ResponseDeleteBucketData`

NewResponseDeleteBucketDataWithDefaults instantiates a new ResponseDeleteBucketData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResponseDeleteBucketData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponseDeleteBucketData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponseDeleteBucketData) SetName(v string)`

SetName sets Name field to given value.


### GetEdgeAccess

`func (o *ResponseDeleteBucketData) GetEdgeAccess() EdgeAccessEnum`

GetEdgeAccess returns the EdgeAccess field if non-nil, zero value otherwise.

### GetEdgeAccessOk

`func (o *ResponseDeleteBucketData) GetEdgeAccessOk() (*EdgeAccessEnum, bool)`

GetEdgeAccessOk returns a tuple with the EdgeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeAccess

`func (o *ResponseDeleteBucketData) SetEdgeAccess(v EdgeAccessEnum)`

SetEdgeAccess sets EdgeAccess field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


