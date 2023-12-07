# BucketCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EdgeAccess** | [**EdgeAccessEnum**](EdgeAccessEnum.md) |  | 

## Methods

### NewBucketCreate

`func NewBucketCreate(name string, edgeAccess EdgeAccessEnum, ) *BucketCreate`

NewBucketCreate instantiates a new BucketCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketCreateWithDefaults

`func NewBucketCreateWithDefaults() *BucketCreate`

NewBucketCreateWithDefaults instantiates a new BucketCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BucketCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BucketCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BucketCreate) SetName(v string)`

SetName sets Name field to given value.


### GetEdgeAccess

`func (o *BucketCreate) GetEdgeAccess() EdgeAccessEnum`

GetEdgeAccess returns the EdgeAccess field if non-nil, zero value otherwise.

### GetEdgeAccessOk

`func (o *BucketCreate) GetEdgeAccessOk() (*EdgeAccessEnum, bool)`

GetEdgeAccessOk returns a tuple with the EdgeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeAccess

`func (o *BucketCreate) SetEdgeAccess(v EdgeAccessEnum)`

SetEdgeAccess sets EdgeAccess field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


