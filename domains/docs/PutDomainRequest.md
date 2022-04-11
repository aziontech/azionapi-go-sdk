# PutDomainRequest

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

### NewPutDomainRequest

`func NewPutDomainRequest(cnames []string, cnameAccessOnly bool, name string, isActive bool, edgeApplicationId int64, digitalCertificateId NullableInt64, ) *PutDomainRequest`

NewPutDomainRequest instantiates a new PutDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutDomainRequestWithDefaults

`func NewPutDomainRequestWithDefaults() *PutDomainRequest`

NewPutDomainRequestWithDefaults instantiates a new PutDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCnames

`func (o *PutDomainRequest) GetCnames() []string`

GetCnames returns the Cnames field if non-nil, zero value otherwise.

### GetCnamesOk

`func (o *PutDomainRequest) GetCnamesOk() (*[]string, bool)`

GetCnamesOk returns a tuple with the Cnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnames

`func (o *PutDomainRequest) SetCnames(v []string)`

SetCnames sets Cnames field to given value.


### GetCnameAccessOnly

`func (o *PutDomainRequest) GetCnameAccessOnly() bool`

GetCnameAccessOnly returns the CnameAccessOnly field if non-nil, zero value otherwise.

### GetCnameAccessOnlyOk

`func (o *PutDomainRequest) GetCnameAccessOnlyOk() (*bool, bool)`

GetCnameAccessOnlyOk returns a tuple with the CnameAccessOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameAccessOnly

`func (o *PutDomainRequest) SetCnameAccessOnly(v bool)`

SetCnameAccessOnly sets CnameAccessOnly field to given value.


### GetName

`func (o *PutDomainRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutDomainRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutDomainRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *PutDomainRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PutDomainRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PutDomainRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetEdgeApplicationId

`func (o *PutDomainRequest) GetEdgeApplicationId() int64`

GetEdgeApplicationId returns the EdgeApplicationId field if non-nil, zero value otherwise.

### GetEdgeApplicationIdOk

`func (o *PutDomainRequest) GetEdgeApplicationIdOk() (*int64, bool)`

GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplicationId

`func (o *PutDomainRequest) SetEdgeApplicationId(v int64)`

SetEdgeApplicationId sets EdgeApplicationId field to given value.


### GetDigitalCertificateId

`func (o *PutDomainRequest) GetDigitalCertificateId() int64`

GetDigitalCertificateId returns the DigitalCertificateId field if non-nil, zero value otherwise.

### GetDigitalCertificateIdOk

`func (o *PutDomainRequest) GetDigitalCertificateIdOk() (*int64, bool)`

GetDigitalCertificateIdOk returns a tuple with the DigitalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalCertificateId

`func (o *PutDomainRequest) SetDigitalCertificateId(v int64)`

SetDigitalCertificateId sets DigitalCertificateId field to given value.


### SetDigitalCertificateIdNil

`func (o *PutDomainRequest) SetDigitalCertificateIdNil(b bool)`

 SetDigitalCertificateIdNil sets the value for DigitalCertificateId to be an explicit nil

### UnsetDigitalCertificateId
`func (o *PutDomainRequest) UnsetDigitalCertificateId()`

UnsetDigitalCertificateId ensures that no value is present for DigitalCertificateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


