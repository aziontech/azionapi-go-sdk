# PutDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Cnames** | **[]string** |  | 
**CnameAccessOnly** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**EdgeApplicationId** | **int64** |  | 
**DigitalCertificateId** | Pointer to **NullableInt64** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**IsMtlsEnabled** | Pointer to **bool** |  | [optional] 
**MtlsTrustedCaCertificateId** | Pointer to **NullableInt64** |  | [optional] 
**MtlsVerification** | Pointer to **string** |  | [optional] 
**CrlList** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewPutDomainRequest

`func NewPutDomainRequest(name string, cnames []string, edgeApplicationId int64, ) *PutDomainRequest`

NewPutDomainRequest instantiates a new PutDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutDomainRequestWithDefaults

`func NewPutDomainRequestWithDefaults() *PutDomainRequest`

NewPutDomainRequestWithDefaults instantiates a new PutDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### HasCnameAccessOnly

`func (o *PutDomainRequest) HasCnameAccessOnly() bool`

HasCnameAccessOnly returns a boolean if a field has been set.

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

### HasIsActive

`func (o *PutDomainRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

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

### HasDigitalCertificateId

`func (o *PutDomainRequest) HasDigitalCertificateId() bool`

HasDigitalCertificateId returns a boolean if a field has been set.

### SetDigitalCertificateIdNil

`func (o *PutDomainRequest) SetDigitalCertificateIdNil(b bool)`

 SetDigitalCertificateIdNil sets the value for DigitalCertificateId to be an explicit nil

### UnsetDigitalCertificateId
`func (o *PutDomainRequest) UnsetDigitalCertificateId()`

UnsetDigitalCertificateId ensures that no value is present for DigitalCertificateId, not even an explicit nil
### GetEnvironment

`func (o *PutDomainRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PutDomainRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PutDomainRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PutDomainRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIsMtlsEnabled

`func (o *PutDomainRequest) GetIsMtlsEnabled() bool`

GetIsMtlsEnabled returns the IsMtlsEnabled field if non-nil, zero value otherwise.

### GetIsMtlsEnabledOk

`func (o *PutDomainRequest) GetIsMtlsEnabledOk() (*bool, bool)`

GetIsMtlsEnabledOk returns a tuple with the IsMtlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMtlsEnabled

`func (o *PutDomainRequest) SetIsMtlsEnabled(v bool)`

SetIsMtlsEnabled sets IsMtlsEnabled field to given value.

### HasIsMtlsEnabled

`func (o *PutDomainRequest) HasIsMtlsEnabled() bool`

HasIsMtlsEnabled returns a boolean if a field has been set.

### GetMtlsTrustedCaCertificateId

`func (o *PutDomainRequest) GetMtlsTrustedCaCertificateId() int64`

GetMtlsTrustedCaCertificateId returns the MtlsTrustedCaCertificateId field if non-nil, zero value otherwise.

### GetMtlsTrustedCaCertificateIdOk

`func (o *PutDomainRequest) GetMtlsTrustedCaCertificateIdOk() (*int64, bool)`

GetMtlsTrustedCaCertificateIdOk returns a tuple with the MtlsTrustedCaCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsTrustedCaCertificateId

`func (o *PutDomainRequest) SetMtlsTrustedCaCertificateId(v int64)`

SetMtlsTrustedCaCertificateId sets MtlsTrustedCaCertificateId field to given value.

### HasMtlsTrustedCaCertificateId

`func (o *PutDomainRequest) HasMtlsTrustedCaCertificateId() bool`

HasMtlsTrustedCaCertificateId returns a boolean if a field has been set.

### SetMtlsTrustedCaCertificateIdNil

`func (o *PutDomainRequest) SetMtlsTrustedCaCertificateIdNil(b bool)`

 SetMtlsTrustedCaCertificateIdNil sets the value for MtlsTrustedCaCertificateId to be an explicit nil

### UnsetMtlsTrustedCaCertificateId
`func (o *PutDomainRequest) UnsetMtlsTrustedCaCertificateId()`

UnsetMtlsTrustedCaCertificateId ensures that no value is present for MtlsTrustedCaCertificateId, not even an explicit nil
### GetMtlsVerification

`func (o *PutDomainRequest) GetMtlsVerification() string`

GetMtlsVerification returns the MtlsVerification field if non-nil, zero value otherwise.

### GetMtlsVerificationOk

`func (o *PutDomainRequest) GetMtlsVerificationOk() (*string, bool)`

GetMtlsVerificationOk returns a tuple with the MtlsVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsVerification

`func (o *PutDomainRequest) SetMtlsVerification(v string)`

SetMtlsVerification sets MtlsVerification field to given value.

### HasMtlsVerification

`func (o *PutDomainRequest) HasMtlsVerification() bool`

HasMtlsVerification returns a boolean if a field has been set.

### GetCrlList

`func (o *PutDomainRequest) GetCrlList() []int64`

GetCrlList returns the CrlList field if non-nil, zero value otherwise.

### GetCrlListOk

`func (o *PutDomainRequest) GetCrlListOk() (*[]int64, bool)`

GetCrlListOk returns a tuple with the CrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlList

`func (o *PutDomainRequest) SetCrlList(v []int64)`

SetCrlList sets CrlList field to given value.

### HasCrlList

`func (o *PutDomainRequest) HasCrlList() bool`

HasCrlList returns a boolean if a field has been set.

### SetCrlListNil

`func (o *PutDomainRequest) SetCrlListNil(b bool)`

 SetCrlListNil sets the value for CrlList to be an explicit nil

### UnsetCrlList
`func (o *PutDomainRequest) UnsetCrlList()`

UnsetCrlList ensures that no value is present for CrlList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


