# DomainEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Cnames** | Pointer to **[]string** |  | [optional] 
**CnameAccessOnly** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**EdgeApplicationId** | Pointer to **int64** |  | [optional] 
**DigitalCertificateId** | Pointer to **NullableInt64** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**IsMtlsEnabled** | Pointer to **bool** |  | [optional] 
**MtlsTrustedCaCertificateId** | Pointer to **NullableInt64** |  | [optional] 
**MtlsVerification** | Pointer to **string** |  | [optional] 
**CrlList** | Pointer to **[]int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainEntity

`func NewDomainEntity() *DomainEntity`

NewDomainEntity instantiates a new DomainEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainEntityWithDefaults

`func NewDomainEntityWithDefaults() *DomainEntity`

NewDomainEntityWithDefaults instantiates a new DomainEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DomainEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCnames

`func (o *DomainEntity) GetCnames() []string`

GetCnames returns the Cnames field if non-nil, zero value otherwise.

### GetCnamesOk

`func (o *DomainEntity) GetCnamesOk() (*[]string, bool)`

GetCnamesOk returns a tuple with the Cnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnames

`func (o *DomainEntity) SetCnames(v []string)`

SetCnames sets Cnames field to given value.

### HasCnames

`func (o *DomainEntity) HasCnames() bool`

HasCnames returns a boolean if a field has been set.

### GetCnameAccessOnly

`func (o *DomainEntity) GetCnameAccessOnly() bool`

GetCnameAccessOnly returns the CnameAccessOnly field if non-nil, zero value otherwise.

### GetCnameAccessOnlyOk

`func (o *DomainEntity) GetCnameAccessOnlyOk() (*bool, bool)`

GetCnameAccessOnlyOk returns a tuple with the CnameAccessOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameAccessOnly

`func (o *DomainEntity) SetCnameAccessOnly(v bool)`

SetCnameAccessOnly sets CnameAccessOnly field to given value.

### HasCnameAccessOnly

`func (o *DomainEntity) HasCnameAccessOnly() bool`

HasCnameAccessOnly returns a boolean if a field has been set.

### GetIsActive

`func (o *DomainEntity) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DomainEntity) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DomainEntity) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DomainEntity) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEdgeApplicationId

`func (o *DomainEntity) GetEdgeApplicationId() int64`

GetEdgeApplicationId returns the EdgeApplicationId field if non-nil, zero value otherwise.

### GetEdgeApplicationIdOk

`func (o *DomainEntity) GetEdgeApplicationIdOk() (*int64, bool)`

GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplicationId

`func (o *DomainEntity) SetEdgeApplicationId(v int64)`

SetEdgeApplicationId sets EdgeApplicationId field to given value.

### HasEdgeApplicationId

`func (o *DomainEntity) HasEdgeApplicationId() bool`

HasEdgeApplicationId returns a boolean if a field has been set.

### GetDigitalCertificateId

`func (o *DomainEntity) GetDigitalCertificateId() int64`

GetDigitalCertificateId returns the DigitalCertificateId field if non-nil, zero value otherwise.

### GetDigitalCertificateIdOk

`func (o *DomainEntity) GetDigitalCertificateIdOk() (*int64, bool)`

GetDigitalCertificateIdOk returns a tuple with the DigitalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalCertificateId

`func (o *DomainEntity) SetDigitalCertificateId(v int64)`

SetDigitalCertificateId sets DigitalCertificateId field to given value.

### HasDigitalCertificateId

`func (o *DomainEntity) HasDigitalCertificateId() bool`

HasDigitalCertificateId returns a boolean if a field has been set.

### SetDigitalCertificateIdNil

`func (o *DomainEntity) SetDigitalCertificateIdNil(b bool)`

 SetDigitalCertificateIdNil sets the value for DigitalCertificateId to be an explicit nil

### UnsetDigitalCertificateId
`func (o *DomainEntity) UnsetDigitalCertificateId()`

UnsetDigitalCertificateId ensures that no value is present for DigitalCertificateId, not even an explicit nil
### GetEnvironment

`func (o *DomainEntity) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DomainEntity) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DomainEntity) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DomainEntity) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIsMtlsEnabled

`func (o *DomainEntity) GetIsMtlsEnabled() bool`

GetIsMtlsEnabled returns the IsMtlsEnabled field if non-nil, zero value otherwise.

### GetIsMtlsEnabledOk

`func (o *DomainEntity) GetIsMtlsEnabledOk() (*bool, bool)`

GetIsMtlsEnabledOk returns a tuple with the IsMtlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMtlsEnabled

`func (o *DomainEntity) SetIsMtlsEnabled(v bool)`

SetIsMtlsEnabled sets IsMtlsEnabled field to given value.

### HasIsMtlsEnabled

`func (o *DomainEntity) HasIsMtlsEnabled() bool`

HasIsMtlsEnabled returns a boolean if a field has been set.

### GetMtlsTrustedCaCertificateId

`func (o *DomainEntity) GetMtlsTrustedCaCertificateId() int64`

GetMtlsTrustedCaCertificateId returns the MtlsTrustedCaCertificateId field if non-nil, zero value otherwise.

### GetMtlsTrustedCaCertificateIdOk

`func (o *DomainEntity) GetMtlsTrustedCaCertificateIdOk() (*int64, bool)`

GetMtlsTrustedCaCertificateIdOk returns a tuple with the MtlsTrustedCaCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsTrustedCaCertificateId

`func (o *DomainEntity) SetMtlsTrustedCaCertificateId(v int64)`

SetMtlsTrustedCaCertificateId sets MtlsTrustedCaCertificateId field to given value.

### HasMtlsTrustedCaCertificateId

`func (o *DomainEntity) HasMtlsTrustedCaCertificateId() bool`

HasMtlsTrustedCaCertificateId returns a boolean if a field has been set.

### SetMtlsTrustedCaCertificateIdNil

`func (o *DomainEntity) SetMtlsTrustedCaCertificateIdNil(b bool)`

 SetMtlsTrustedCaCertificateIdNil sets the value for MtlsTrustedCaCertificateId to be an explicit nil

### UnsetMtlsTrustedCaCertificateId
`func (o *DomainEntity) UnsetMtlsTrustedCaCertificateId()`

UnsetMtlsTrustedCaCertificateId ensures that no value is present for MtlsTrustedCaCertificateId, not even an explicit nil
### GetMtlsVerification

`func (o *DomainEntity) GetMtlsVerification() string`

GetMtlsVerification returns the MtlsVerification field if non-nil, zero value otherwise.

### GetMtlsVerificationOk

`func (o *DomainEntity) GetMtlsVerificationOk() (*string, bool)`

GetMtlsVerificationOk returns a tuple with the MtlsVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsVerification

`func (o *DomainEntity) SetMtlsVerification(v string)`

SetMtlsVerification sets MtlsVerification field to given value.

### HasMtlsVerification

`func (o *DomainEntity) HasMtlsVerification() bool`

HasMtlsVerification returns a boolean if a field has been set.

### GetCrlList

`func (o *DomainEntity) GetCrlList() []int64`

GetCrlList returns the CrlList field if non-nil, zero value otherwise.

### GetCrlListOk

`func (o *DomainEntity) GetCrlListOk() (*[]int64, bool)`

GetCrlListOk returns a tuple with the CrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlList

`func (o *DomainEntity) SetCrlList(v []int64)`

SetCrlList sets CrlList field to given value.

### HasCrlList

`func (o *DomainEntity) HasCrlList() bool`

HasCrlList returns a boolean if a field has been set.

### SetCrlListNil

`func (o *DomainEntity) SetCrlListNil(b bool)`

 SetCrlListNil sets the value for CrlList to be an explicit nil

### UnsetCrlList
`func (o *DomainEntity) UnsetCrlList()`

UnsetCrlList ensures that no value is present for CrlList, not even an explicit nil
### GetId

`func (o *DomainEntity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainEntity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainEntity) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DomainEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomainName

`func (o *DomainEntity) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DomainEntity) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DomainEntity) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DomainEntity) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


