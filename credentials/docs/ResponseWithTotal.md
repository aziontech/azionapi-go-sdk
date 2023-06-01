# ResponseWithTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**[]Response**](Response.md) |  | 
**Total** | **int64** |  | 

## Methods

### NewResponseWithTotal

`func NewResponseWithTotal(credentials []Response, total int64, ) *ResponseWithTotal`

NewResponseWithTotal instantiates a new ResponseWithTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithTotalWithDefaults

`func NewResponseWithTotalWithDefaults() *ResponseWithTotal`

NewResponseWithTotalWithDefaults instantiates a new ResponseWithTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ResponseWithTotal) GetCredentials() []Response`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ResponseWithTotal) GetCredentialsOk() (*[]Response, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ResponseWithTotal) SetCredentials(v []Response)`

SetCredentials sets Credentials field to given value.


### GetTotal

`func (o *ResponseWithTotal) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseWithTotal) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseWithTotal) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


