# DomainResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Cnames** | **[]string** |  | 
**CnameAccessOnly** | **bool** |  | 
**IsActive** | **bool** |  | 
**EdgeApplicationId** | **int64** |  | 
**DigitalCertificateId** | **NullableInt64** |  | 
**DomainName** | **string** |  | 
**Environment** | Pointer to **string** |  | [optional] 
**IsMtlsEnabled** | **bool** |  | 
**MtlsTrustedCaCertificateId** | **NullableString** |  | 
**MtlsVerification** | **NullableString** |  | 

## Methods

### NewDomainResults

`func NewDomainResults(id int64, name string, cnames []string, cnameAccessOnly bool, isActive bool, edgeApplicationId int64, digitalCertificateId NullableInt64, domainName string, isMtlsEnabled bool, mtlsTrustedCaCertificateId NullableString, mtlsVerification NullableString, ) *DomainResults`

NewDomainResults instantiates a new DomainResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainResultsWithDefaults

`func NewDomainResultsWithDefaults() *DomainResults`

NewDomainResultsWithDefaults instantiates a new DomainResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DomainResults) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainResults) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainResults) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *DomainResults) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainResults) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainResults) SetName(v string)`

SetName sets Name field to given value.


### GetCnames

`func (o *DomainResults) GetCnames() []string`

GetCnames returns the Cnames field if non-nil, zero value otherwise.

### GetCnamesOk

`func (o *DomainResults) GetCnamesOk() (*[]string, bool)`

GetCnamesOk returns a tuple with the Cnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnames

`func (o *DomainResults) SetCnames(v []string)`

SetCnames sets Cnames field to given value.


### GetCnameAccessOnly

`func (o *DomainResults) GetCnameAccessOnly() bool`

GetCnameAccessOnly returns the CnameAccessOnly field if non-nil, zero value otherwise.

### GetCnameAccessOnlyOk

`func (o *DomainResults) GetCnameAccessOnlyOk() (*bool, bool)`

GetCnameAccessOnlyOk returns a tuple with the CnameAccessOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameAccessOnly

`func (o *DomainResults) SetCnameAccessOnly(v bool)`

SetCnameAccessOnly sets CnameAccessOnly field to given value.


### GetIsActive

`func (o *DomainResults) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DomainResults) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DomainResults) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetEdgeApplicationId

`func (o *DomainResults) GetEdgeApplicationId() int64`

GetEdgeApplicationId returns the EdgeApplicationId field if non-nil, zero value otherwise.

### GetEdgeApplicationIdOk

`func (o *DomainResults) GetEdgeApplicationIdOk() (*int64, bool)`

GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplicationId

`func (o *DomainResults) SetEdgeApplicationId(v int64)`

SetEdgeApplicationId sets EdgeApplicationId field to given value.


### GetDigitalCertificateId

`func (o *DomainResults) GetDigitalCertificateId() int64`

GetDigitalCertificateId returns the DigitalCertificateId field if non-nil, zero value otherwise.

### GetDigitalCertificateIdOk

`func (o *DomainResults) GetDigitalCertificateIdOk() (*int64, bool)`

GetDigitalCertificateIdOk returns a tuple with the DigitalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalCertificateId

`func (o *DomainResults) SetDigitalCertificateId(v int64)`

SetDigitalCertificateId sets DigitalCertificateId field to given value.


### SetDigitalCertificateIdNil

`func (o *DomainResults) SetDigitalCertificateIdNil(b bool)`

 SetDigitalCertificateIdNil sets the value for DigitalCertificateId to be an explicit nil

### UnsetDigitalCertificateId
`func (o *DomainResults) UnsetDigitalCertificateId()`

UnsetDigitalCertificateId ensures that no value is present for DigitalCertificateId, not even an explicit nil
### GetDomainName

`func (o *DomainResults) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DomainResults) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DomainResults) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetEnvironment

`func (o *DomainResults) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DomainResults) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DomainResults) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DomainResults) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIsMtlsEnabled

`func (o *DomainResults) GetIsMtlsEnabled() bool`

GetIsMtlsEnabled returns the IsMtlsEnabled field if non-nil, zero value otherwise.

### GetIsMtlsEnabledOk

`func (o *DomainResults) GetIsMtlsEnabledOk() (*bool, bool)`

GetIsMtlsEnabledOk returns a tuple with the IsMtlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMtlsEnabled

`func (o *DomainResults) SetIsMtlsEnabled(v bool)`

SetIsMtlsEnabled sets IsMtlsEnabled field to given value.


### GetMtlsTrustedCaCertificateId

`func (o *DomainResults) GetMtlsTrustedCaCertificateId() string`

GetMtlsTrustedCaCertificateId returns the MtlsTrustedCaCertificateId field if non-nil, zero value otherwise.

### GetMtlsTrustedCaCertificateIdOk

`func (o *DomainResults) GetMtlsTrustedCaCertificateIdOk() (*string, bool)`

GetMtlsTrustedCaCertificateIdOk returns a tuple with the MtlsTrustedCaCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsTrustedCaCertificateId

`func (o *DomainResults) SetMtlsTrustedCaCertificateId(v string)`

SetMtlsTrustedCaCertificateId sets MtlsTrustedCaCertificateId field to given value.


### SetMtlsTrustedCaCertificateIdNil

`func (o *DomainResults) SetMtlsTrustedCaCertificateIdNil(b bool)`

 SetMtlsTrustedCaCertificateIdNil sets the value for MtlsTrustedCaCertificateId to be an explicit nil

### UnsetMtlsTrustedCaCertificateId
`func (o *DomainResults) UnsetMtlsTrustedCaCertificateId()`

UnsetMtlsTrustedCaCertificateId ensures that no value is present for MtlsTrustedCaCertificateId, not even an explicit nil
### GetMtlsVerification

`func (o *DomainResults) GetMtlsVerification() string`

GetMtlsVerification returns the MtlsVerification field if non-nil, zero value otherwise.

### GetMtlsVerificationOk

`func (o *DomainResults) GetMtlsVerificationOk() (*string, bool)`

GetMtlsVerificationOk returns a tuple with the MtlsVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsVerification

`func (o *DomainResults) SetMtlsVerification(v string)`

SetMtlsVerification sets MtlsVerification field to given value.


### SetMtlsVerificationNil

`func (o *DomainResults) SetMtlsVerificationNil(b bool)`

 SetMtlsVerificationNil sets the value for MtlsVerification to be an explicit nil

### UnsetMtlsVerification
`func (o *DomainResults) UnsetMtlsVerification()`

UnsetMtlsVerification ensures that no value is present for MtlsVerification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


