# BucketObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | [readonly] 
**LastModified** | **time.Time** |  | [readonly] 
**Size** | **int32** |  | [readonly] 
**Etag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewBucketObject

`func NewBucketObject(key string, lastModified time.Time, size int32, ) *BucketObject`

NewBucketObject instantiates a new BucketObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketObjectWithDefaults

`func NewBucketObjectWithDefaults() *BucketObject`

NewBucketObjectWithDefaults instantiates a new BucketObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *BucketObject) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BucketObject) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BucketObject) SetKey(v string)`

SetKey sets Key field to given value.


### GetLastModified

`func (o *BucketObject) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *BucketObject) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *BucketObject) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetSize

`func (o *BucketObject) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BucketObject) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BucketObject) SetSize(v int32)`

SetSize sets Size field to given value.


### GetEtag

`func (o *BucketObject) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *BucketObject) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *BucketObject) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *BucketObject) HasEtag() bool`

HasEtag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


