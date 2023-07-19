# CreateCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Certificate** | **string** |  | 
**PrivateKey** | **string** |  | 
**CertificateType** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateCertificateRequest

`func NewCreateCertificateRequest(name string, certificate string, privateKey string, ) *CreateCertificateRequest`

NewCreateCertificateRequest instantiates a new CreateCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCertificateRequestWithDefaults

`func NewCreateCertificateRequestWithDefaults() *CreateCertificateRequest`

NewCreateCertificateRequestWithDefaults instantiates a new CreateCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCertificateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCertificateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCertificateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCertificate

`func (o *CreateCertificateRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CreateCertificateRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CreateCertificateRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetPrivateKey

`func (o *CreateCertificateRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreateCertificateRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CreateCertificateRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetCertificateType

`func (o *CreateCertificateRequest) GetCertificateType() string`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *CreateCertificateRequest) GetCertificateTypeOk() (*string, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *CreateCertificateRequest) SetCertificateType(v string)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *CreateCertificateRequest) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.

### GetManaged

`func (o *CreateCertificateRequest) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *CreateCertificateRequest) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *CreateCertificateRequest) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *CreateCertificateRequest) HasManaged() bool`

HasManaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


