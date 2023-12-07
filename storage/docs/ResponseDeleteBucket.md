# ResponseDeleteBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**StateEnum**](StateEnum.md) |  | 
**Data** | [**NullableResponseDeleteBucketData**](ResponseDeleteBucketData.md) |  | 

## Methods

### NewResponseDeleteBucket

`func NewResponseDeleteBucket(state StateEnum, data NullableResponseDeleteBucketData, ) *ResponseDeleteBucket`

NewResponseDeleteBucket instantiates a new ResponseDeleteBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseDeleteBucketWithDefaults

`func NewResponseDeleteBucketWithDefaults() *ResponseDeleteBucket`

NewResponseDeleteBucketWithDefaults instantiates a new ResponseDeleteBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseDeleteBucket) GetState() StateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseDeleteBucket) GetStateOk() (*StateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseDeleteBucket) SetState(v StateEnum)`

SetState sets State field to given value.


### GetData

`func (o *ResponseDeleteBucket) GetData() ResponseDeleteBucketData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseDeleteBucket) GetDataOk() (*ResponseDeleteBucketData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseDeleteBucket) SetData(v ResponseDeleteBucketData)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ResponseDeleteBucket) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ResponseDeleteBucket) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


