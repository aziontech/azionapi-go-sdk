# DomainData

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
**EdgeFirewallId** | Pointer to **NullableInt64** |  | [optional] 
**MtlsVerification** | Pointer to **string** |  | [optional] 
**CrlList** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewDomainData

`func NewDomainData() *DomainData`

NewDomainData instantiates a new DomainData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainDataWithDefaults

`func NewDomainDataWithDefaults() *DomainData`

NewDomainDataWithDefaults instantiates a new DomainData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DomainData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCnames

`func (o *DomainData) GetCnames() []string`

GetCnames returns the Cnames field if non-nil, zero value otherwise.

### GetCnamesOk

`func (o *DomainData) GetCnamesOk() (*[]string, bool)`

GetCnamesOk returns a tuple with the Cnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnames

`func (o *DomainData) SetCnames(v []string)`

SetCnames sets Cnames field to given value.

### HasCnames

`func (o *DomainData) HasCnames() bool`

HasCnames returns a boolean if a field has been set.

### GetCnameAccessOnly

`func (o *DomainData) GetCnameAccessOnly() bool`

GetCnameAccessOnly returns the CnameAccessOnly field if non-nil, zero value otherwise.

### GetCnameAccessOnlyOk

`func (o *DomainData) GetCnameAccessOnlyOk() (*bool, bool)`

GetCnameAccessOnlyOk returns a tuple with the CnameAccessOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameAccessOnly

`func (o *DomainData) SetCnameAccessOnly(v bool)`

SetCnameAccessOnly sets CnameAccessOnly field to given value.

### HasCnameAccessOnly

`func (o *DomainData) HasCnameAccessOnly() bool`

HasCnameAccessOnly returns a boolean if a field has been set.

### GetIsActive

`func (o *DomainData) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DomainData) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DomainData) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DomainData) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEdgeApplicationId

`func (o *DomainData) GetEdgeApplicationId() int64`

GetEdgeApplicationId returns the EdgeApplicationId field if non-nil, zero value otherwise.

### GetEdgeApplicationIdOk

`func (o *DomainData) GetEdgeApplicationIdOk() (*int64, bool)`

GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplicationId

`func (o *DomainData) SetEdgeApplicationId(v int64)`

SetEdgeApplicationId sets EdgeApplicationId field to given value.

### HasEdgeApplicationId

`func (o *DomainData) HasEdgeApplicationId() bool`

HasEdgeApplicationId returns a boolean if a field has been set.

### GetDigitalCertificateId

`func (o *DomainData) GetDigitalCertificateId() int64`

GetDigitalCertificateId returns the DigitalCertificateId field if non-nil, zero value otherwise.

### GetDigitalCertificateIdOk

`func (o *DomainData) GetDigitalCertificateIdOk() (*int64, bool)`

GetDigitalCertificateIdOk returns a tuple with the DigitalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalCertificateId

`func (o *DomainData) SetDigitalCertificateId(v int64)`

SetDigitalCertificateId sets DigitalCertificateId field to given value.

### HasDigitalCertificateId

`func (o *DomainData) HasDigitalCertificateId() bool`

HasDigitalCertificateId returns a boolean if a field has been set.

### SetDigitalCertificateIdNil

`func (o *DomainData) SetDigitalCertificateIdNil(b bool)`

 SetDigitalCertificateIdNil sets the value for DigitalCertificateId to be an explicit nil

### UnsetDigitalCertificateId
`func (o *DomainData) UnsetDigitalCertificateId()`

UnsetDigitalCertificateId ensures that no value is present for DigitalCertificateId, not even an explicit nil
### GetEnvironment

`func (o *DomainData) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DomainData) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DomainData) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DomainData) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIsMtlsEnabled

`func (o *DomainData) GetIsMtlsEnabled() bool`

GetIsMtlsEnabled returns the IsMtlsEnabled field if non-nil, zero value otherwise.

### GetIsMtlsEnabledOk

`func (o *DomainData) GetIsMtlsEnabledOk() (*bool, bool)`

GetIsMtlsEnabledOk returns a tuple with the IsMtlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMtlsEnabled

`func (o *DomainData) SetIsMtlsEnabled(v bool)`

SetIsMtlsEnabled sets IsMtlsEnabled field to given value.

### HasIsMtlsEnabled

`func (o *DomainData) HasIsMtlsEnabled() bool`

HasIsMtlsEnabled returns a boolean if a field has been set.

### GetMtlsTrustedCaCertificateId

`func (o *DomainData) GetMtlsTrustedCaCertificateId() int64`

GetMtlsTrustedCaCertificateId returns the MtlsTrustedCaCertificateId field if non-nil, zero value otherwise.

### GetMtlsTrustedCaCertificateIdOk

`func (o *DomainData) GetMtlsTrustedCaCertificateIdOk() (*int64, bool)`

GetMtlsTrustedCaCertificateIdOk returns a tuple with the MtlsTrustedCaCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsTrustedCaCertificateId

`func (o *DomainData) SetMtlsTrustedCaCertificateId(v int64)`

SetMtlsTrustedCaCertificateId sets MtlsTrustedCaCertificateId field to given value.

### HasMtlsTrustedCaCertificateId

`func (o *DomainData) HasMtlsTrustedCaCertificateId() bool`

HasMtlsTrustedCaCertificateId returns a boolean if a field has been set.

### SetMtlsTrustedCaCertificateIdNil

`func (o *DomainData) SetMtlsTrustedCaCertificateIdNil(b bool)`

 SetMtlsTrustedCaCertificateIdNil sets the value for MtlsTrustedCaCertificateId to be an explicit nil

### UnsetMtlsTrustedCaCertificateId
`func (o *DomainData) UnsetMtlsTrustedCaCertificateId()`

UnsetMtlsTrustedCaCertificateId ensures that no value is present for MtlsTrustedCaCertificateId, not even an explicit nil
### GetEdgeFirewallId

`func (o *DomainData) GetEdgeFirewallId() int64`

GetEdgeFirewallId returns the EdgeFirewallId field if non-nil, zero value otherwise.

### GetEdgeFirewallIdOk

`func (o *DomainData) GetEdgeFirewallIdOk() (*int64, bool)`

GetEdgeFirewallIdOk returns a tuple with the EdgeFirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFirewallId

`func (o *DomainData) SetEdgeFirewallId(v int64)`

SetEdgeFirewallId sets EdgeFirewallId field to given value.

### HasEdgeFirewallId

`func (o *DomainData) HasEdgeFirewallId() bool`

HasEdgeFirewallId returns a boolean if a field has been set.

### SetEdgeFirewallIdNil

`func (o *DomainData) SetEdgeFirewallIdNil(b bool)`

 SetEdgeFirewallIdNil sets the value for EdgeFirewallId to be an explicit nil

### UnsetEdgeFirewallId
`func (o *DomainData) UnsetEdgeFirewallId()`

UnsetEdgeFirewallId ensures that no value is present for EdgeFirewallId, not even an explicit nil
### GetMtlsVerification

`func (o *DomainData) GetMtlsVerification() string`

GetMtlsVerification returns the MtlsVerification field if non-nil, zero value otherwise.

### GetMtlsVerificationOk

`func (o *DomainData) GetMtlsVerificationOk() (*string, bool)`

GetMtlsVerificationOk returns a tuple with the MtlsVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsVerification

`func (o *DomainData) SetMtlsVerification(v string)`

SetMtlsVerification sets MtlsVerification field to given value.

### HasMtlsVerification

`func (o *DomainData) HasMtlsVerification() bool`

HasMtlsVerification returns a boolean if a field has been set.

### GetCrlList

`func (o *DomainData) GetCrlList() []int64`

GetCrlList returns the CrlList field if non-nil, zero value otherwise.

### GetCrlListOk

`func (o *DomainData) GetCrlListOk() (*[]int64, bool)`

GetCrlListOk returns a tuple with the CrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlList

`func (o *DomainData) SetCrlList(v []int64)`

SetCrlList sets CrlList field to given value.

### HasCrlList

`func (o *DomainData) HasCrlList() bool`

HasCrlList returns a boolean if a field has been set.

### SetCrlListNil

`func (o *DomainData) SetCrlListNil(b bool)`

 SetCrlListNil sets the value for CrlList to be an explicit nil

### UnsetCrlList
`func (o *DomainData) UnsetCrlList()`

UnsetCrlList ensures that no value is present for CrlList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


