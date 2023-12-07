# PatchedBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] [readonly] 
**EdgeAccess** | Pointer to [**EdgeAccessEnum**](EdgeAccessEnum.md) |  | [optional] 

## Methods

### NewPatchedBucket

`func NewPatchedBucket() *PatchedBucket`

NewPatchedBucket instantiates a new PatchedBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBucketWithDefaults

`func NewPatchedBucketWithDefaults() *PatchedBucket`

NewPatchedBucketWithDefaults instantiates a new PatchedBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEdgeAccess

`func (o *PatchedBucket) GetEdgeAccess() EdgeAccessEnum`

GetEdgeAccess returns the EdgeAccess field if non-nil, zero value otherwise.

### GetEdgeAccessOk

`func (o *PatchedBucket) GetEdgeAccessOk() (*EdgeAccessEnum, bool)`

GetEdgeAccessOk returns a tuple with the EdgeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeAccess

`func (o *PatchedBucket) SetEdgeAccess(v EdgeAccessEnum)`

SetEdgeAccess sets EdgeAccess field to given value.

### HasEdgeAccess

`func (o *PatchedBucket) HasEdgeAccess() bool`

HasEdgeAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


