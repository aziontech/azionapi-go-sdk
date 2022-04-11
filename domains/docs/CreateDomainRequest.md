# CreateDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cnames** | **[]string** |  | 
**CnameAccessOnly** | **bool** |  | 
**Name** | **string** |  | 
**IsActive** | **bool** |  | 
**EdgeApplicationId** | **int64** |  | 
**DigitalCertificateId** | **NullableInt64** |  | 

## Methods

### NewCreateDomainRequest

`func NewCreateDomainRequest(cnames []string, cnameAccessOnly bool, name string, isActive bool, edgeApplicationId int64, digitalCertificateId NullableInt64, ) *CreateDomainRequest`

NewCreateDomainRequest instantiates a new CreateDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDomainRequestWithDefaults

`func NewCreateDomainRequestWithDefaults() *CreateDomainRequest`

NewCreateDomainRequestWithDefaults instantiates a new CreateDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCnames

`func (o *CreateDomainRequest) GetCnames() []string`

GetCnames returns the Cnames field if non-nil, zero value otherwise.

### GetCnamesOk

`func (o *CreateDomainRequest) GetCnamesOk() (*[]string, bool)`

GetCnamesOk returns a tuple with the Cnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnames

`func (o *CreateDomainRequest) SetCnames(v []string)`

SetCnames sets Cnames field to given value.


### GetCnameAccessOnly

`func (o *CreateDomainRequest) GetCnameAccessOnly() bool`

GetCnameAccessOnly returns the CnameAccessOnly field if non-nil, zero value otherwise.

### GetCnameAccessOnlyOk

`func (o *CreateDomainRequest) GetCnameAccessOnlyOk() (*bool, bool)`

GetCnameAccessOnlyOk returns a tuple with the CnameAccessOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameAccessOnly

`func (o *CreateDomainRequest) SetCnameAccessOnly(v bool)`

SetCnameAccessOnly sets CnameAccessOnly field to given value.


### GetName

`func (o *CreateDomainRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDomainRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDomainRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *CreateDomainRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateDomainRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateDomainRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetEdgeApplicationId

`func (o *CreateDomainRequest) GetEdgeApplicationId() int64`

GetEdgeApplicationId returns the EdgeApplicationId field if non-nil, zero value otherwise.

### GetEdgeApplicationIdOk

`func (o *CreateDomainRequest) GetEdgeApplicationIdOk() (*int64, bool)`

GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplicationId

`func (o *CreateDomainRequest) SetEdgeApplicationId(v int64)`

SetEdgeApplicationId sets EdgeApplicationId field to given value.


### GetDigitalCertificateId

`func (o *CreateDomainRequest) GetDigitalCertificateId() int64`

GetDigitalCertificateId returns the DigitalCertificateId field if non-nil, zero value otherwise.

### GetDigitalCertificateIdOk

`func (o *CreateDomainRequest) GetDigitalCertificateIdOk() (*int64, bool)`

GetDigitalCertificateIdOk returns a tuple with the DigitalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalCertificateId

`func (o *CreateDomainRequest) SetDigitalCertificateId(v int64)`

SetDigitalCertificateId sets DigitalCertificateId field to given value.


### SetDigitalCertificateIdNil

`func (o *CreateDomainRequest) SetDigitalCertificateIdNil(b bool)`

 SetDigitalCertificateIdNil sets the value for DigitalCertificateId to be an explicit nil

### UnsetDigitalCertificateId
`func (o *CreateDomainRequest) UnsetDigitalCertificateId()`

UnsetDigitalCertificateId ensures that no value is present for DigitalCertificateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


