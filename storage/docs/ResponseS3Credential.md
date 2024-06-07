# ResponseS3Credential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**S3Credential**](S3Credential.md) |  | [optional] 

## Methods

### NewResponseS3Credential

`func NewResponseS3Credential() *ResponseS3Credential`

NewResponseS3Credential instantiates a new ResponseS3Credential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseS3CredentialWithDefaults

`func NewResponseS3CredentialWithDefaults() *ResponseS3Credential`

NewResponseS3CredentialWithDefaults instantiates a new ResponseS3Credential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResponseS3Credential) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponseS3Credential) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponseS3Credential) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponseS3Credential) HasState() bool`

HasState returns a boolean if a field has been set.

### GetData

`func (o *ResponseS3Credential) GetData() S3Credential`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseS3Credential) GetDataOk() (*S3Credential, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseS3Credential) SetData(v S3Credential)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseS3Credential) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


