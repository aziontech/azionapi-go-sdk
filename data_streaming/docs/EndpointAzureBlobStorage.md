# EndpointAzureBlobStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointType** | Pointer to **string** |  | [optional] 
**StorageAccount** | Pointer to **string** |  | [optional] 
**ContainerName** | Pointer to **string** |  | [optional] 
**BlobSasToken** | Pointer to **string** |  | [optional] 

## Methods

### NewEndpointAzureBlobStorage

`func NewEndpointAzureBlobStorage() *EndpointAzureBlobStorage`

NewEndpointAzureBlobStorage instantiates a new EndpointAzureBlobStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointAzureBlobStorageWithDefaults

`func NewEndpointAzureBlobStorageWithDefaults() *EndpointAzureBlobStorage`

NewEndpointAzureBlobStorageWithDefaults instantiates a new EndpointAzureBlobStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointType

`func (o *EndpointAzureBlobStorage) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *EndpointAzureBlobStorage) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *EndpointAzureBlobStorage) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *EndpointAzureBlobStorage) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### GetStorageAccount

`func (o *EndpointAzureBlobStorage) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *EndpointAzureBlobStorage) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *EndpointAzureBlobStorage) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *EndpointAzureBlobStorage) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### GetContainerName

`func (o *EndpointAzureBlobStorage) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *EndpointAzureBlobStorage) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *EndpointAzureBlobStorage) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *EndpointAzureBlobStorage) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetBlobSasToken

`func (o *EndpointAzureBlobStorage) GetBlobSasToken() string`

GetBlobSasToken returns the BlobSasToken field if non-nil, zero value otherwise.

### GetBlobSasTokenOk

`func (o *EndpointAzureBlobStorage) GetBlobSasTokenOk() (*string, bool)`

GetBlobSasTokenOk returns a tuple with the BlobSasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobSasToken

`func (o *EndpointAzureBlobStorage) SetBlobSasToken(v string)`

SetBlobSasToken sets BlobSasToken field to given value.

### HasBlobSasToken

`func (o *EndpointAzureBlobStorage) HasBlobSasToken() bool`

HasBlobSasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


