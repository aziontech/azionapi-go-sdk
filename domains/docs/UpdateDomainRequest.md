# UpdateDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Cnames** | Pointer to **[]string** |  | [optional] 
**CnameAccessOnly** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**EdgeApplicationId** | Pointer to **int64** |  | [optional] 
**DigitalCertificateId** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**IsMtlsEnabled** | Pointer to **bool** |  | [optional] 
**MtlsTrustedCaCertificateId** | Pointer to **NullableInt64** |  | [optional] 
**EdgeFirewallId** | Pointer to **NullableInt64** |  | [optional] 
**MtlsVerification** | Pointer to **string** |  | [optional] 
**CrlList** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewUpdateDomainRequest

`func NewUpdateDomainRequest() *UpdateDomainRequest`

NewUpdateDomainRequest instantiates a new UpdateDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainRequestWithDefaults

`func NewUpdateDomainRequestWithDefaults() *UpdateDomainRequest`

NewUpdateDomainRequestWithDefaults instantiates a new UpdateDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDomainRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDomainRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDomainRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDomainRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCnames

`func (o *UpdateDomainRequest) GetCnames() []string`

GetCnames returns the Cnames field if non-nil, zero value otherwise.

### GetCnamesOk

`func (o *UpdateDomainRequest) GetCnamesOk() (*[]string, bool)`

GetCnamesOk returns a tuple with the Cnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnames

`func (o *UpdateDomainRequest) SetCnames(v []string)`

SetCnames sets Cnames field to given value.

### HasCnames

`func (o *UpdateDomainRequest) HasCnames() bool`

HasCnames returns a boolean if a field has been set.

### GetCnameAccessOnly

`func (o *UpdateDomainRequest) GetCnameAccessOnly() bool`

GetCnameAccessOnly returns the CnameAccessOnly field if non-nil, zero value otherwise.

### GetCnameAccessOnlyOk

`func (o *UpdateDomainRequest) GetCnameAccessOnlyOk() (*bool, bool)`

GetCnameAccessOnlyOk returns a tuple with the CnameAccessOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameAccessOnly

`func (o *UpdateDomainRequest) SetCnameAccessOnly(v bool)`

SetCnameAccessOnly sets CnameAccessOnly field to given value.

### HasCnameAccessOnly

`func (o *UpdateDomainRequest) HasCnameAccessOnly() bool`

HasCnameAccessOnly returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateDomainRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateDomainRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateDomainRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateDomainRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEdgeApplicationId

`func (o *UpdateDomainRequest) GetEdgeApplicationId() int64`

GetEdgeApplicationId returns the EdgeApplicationId field if non-nil, zero value otherwise.

### GetEdgeApplicationIdOk

`func (o *UpdateDomainRequest) GetEdgeApplicationIdOk() (*int64, bool)`

GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplicationId

`func (o *UpdateDomainRequest) SetEdgeApplicationId(v int64)`

SetEdgeApplicationId sets EdgeApplicationId field to given value.

### HasEdgeApplicationId

`func (o *UpdateDomainRequest) HasEdgeApplicationId() bool`

HasEdgeApplicationId returns a boolean if a field has been set.

### GetDigitalCertificateId

`func (o *UpdateDomainRequest) GetDigitalCertificateId() string`

GetDigitalCertificateId returns the DigitalCertificateId field if non-nil, zero value otherwise.

### GetDigitalCertificateIdOk

`func (o *UpdateDomainRequest) GetDigitalCertificateIdOk() (*string, bool)`

GetDigitalCertificateIdOk returns a tuple with the DigitalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalCertificateId

`func (o *UpdateDomainRequest) SetDigitalCertificateId(v string)`

SetDigitalCertificateId sets DigitalCertificateId field to given value.

### HasDigitalCertificateId

`func (o *UpdateDomainRequest) HasDigitalCertificateId() bool`

HasDigitalCertificateId returns a boolean if a field has been set.

### GetEnvironment

`func (o *UpdateDomainRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UpdateDomainRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UpdateDomainRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UpdateDomainRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIsMtlsEnabled

`func (o *UpdateDomainRequest) GetIsMtlsEnabled() bool`

GetIsMtlsEnabled returns the IsMtlsEnabled field if non-nil, zero value otherwise.

### GetIsMtlsEnabledOk

`func (o *UpdateDomainRequest) GetIsMtlsEnabledOk() (*bool, bool)`

GetIsMtlsEnabledOk returns a tuple with the IsMtlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMtlsEnabled

`func (o *UpdateDomainRequest) SetIsMtlsEnabled(v bool)`

SetIsMtlsEnabled sets IsMtlsEnabled field to given value.

### HasIsMtlsEnabled

`func (o *UpdateDomainRequest) HasIsMtlsEnabled() bool`

HasIsMtlsEnabled returns a boolean if a field has been set.

### GetMtlsTrustedCaCertificateId

`func (o *UpdateDomainRequest) GetMtlsTrustedCaCertificateId() int64`

GetMtlsTrustedCaCertificateId returns the MtlsTrustedCaCertificateId field if non-nil, zero value otherwise.

### GetMtlsTrustedCaCertificateIdOk

`func (o *UpdateDomainRequest) GetMtlsTrustedCaCertificateIdOk() (*int64, bool)`

GetMtlsTrustedCaCertificateIdOk returns a tuple with the MtlsTrustedCaCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsTrustedCaCertificateId

`func (o *UpdateDomainRequest) SetMtlsTrustedCaCertificateId(v int64)`

SetMtlsTrustedCaCertificateId sets MtlsTrustedCaCertificateId field to given value.

### HasMtlsTrustedCaCertificateId

`func (o *UpdateDomainRequest) HasMtlsTrustedCaCertificateId() bool`

HasMtlsTrustedCaCertificateId returns a boolean if a field has been set.

### SetMtlsTrustedCaCertificateIdNil

`func (o *UpdateDomainRequest) SetMtlsTrustedCaCertificateIdNil(b bool)`

 SetMtlsTrustedCaCertificateIdNil sets the value for MtlsTrustedCaCertificateId to be an explicit nil

### UnsetMtlsTrustedCaCertificateId
`func (o *UpdateDomainRequest) UnsetMtlsTrustedCaCertificateId()`

UnsetMtlsTrustedCaCertificateId ensures that no value is present for MtlsTrustedCaCertificateId, not even an explicit nil
### GetEdgeFirewallId

`func (o *UpdateDomainRequest) GetEdgeFirewallId() int64`

GetEdgeFirewallId returns the EdgeFirewallId field if non-nil, zero value otherwise.

### GetEdgeFirewallIdOk

`func (o *UpdateDomainRequest) GetEdgeFirewallIdOk() (*int64, bool)`

GetEdgeFirewallIdOk returns a tuple with the EdgeFirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFirewallId

`func (o *UpdateDomainRequest) SetEdgeFirewallId(v int64)`

SetEdgeFirewallId sets EdgeFirewallId field to given value.

### HasEdgeFirewallId

`func (o *UpdateDomainRequest) HasEdgeFirewallId() bool`

HasEdgeFirewallId returns a boolean if a field has been set.

### SetEdgeFirewallIdNil

`func (o *UpdateDomainRequest) SetEdgeFirewallIdNil(b bool)`

 SetEdgeFirewallIdNil sets the value for EdgeFirewallId to be an explicit nil

### UnsetEdgeFirewallId
`func (o *UpdateDomainRequest) UnsetEdgeFirewallId()`

UnsetEdgeFirewallId ensures that no value is present for EdgeFirewallId, not even an explicit nil
### GetMtlsVerification

`func (o *UpdateDomainRequest) GetMtlsVerification() string`

GetMtlsVerification returns the MtlsVerification field if non-nil, zero value otherwise.

### GetMtlsVerificationOk

`func (o *UpdateDomainRequest) GetMtlsVerificationOk() (*string, bool)`

GetMtlsVerificationOk returns a tuple with the MtlsVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsVerification

`func (o *UpdateDomainRequest) SetMtlsVerification(v string)`

SetMtlsVerification sets MtlsVerification field to given value.

### HasMtlsVerification

`func (o *UpdateDomainRequest) HasMtlsVerification() bool`

HasMtlsVerification returns a boolean if a field has been set.

### GetCrlList

`func (o *UpdateDomainRequest) GetCrlList() []int64`

GetCrlList returns the CrlList field if non-nil, zero value otherwise.

### GetCrlListOk

`func (o *UpdateDomainRequest) GetCrlListOk() (*[]int64, bool)`

GetCrlListOk returns a tuple with the CrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlList

`func (o *UpdateDomainRequest) SetCrlList(v []int64)`

SetCrlList sets CrlList field to given value.

### HasCrlList

`func (o *UpdateDomainRequest) HasCrlList() bool`

HasCrlList returns a boolean if a field has been set.

### SetCrlListNil

`func (o *UpdateDomainRequest) SetCrlListNil(b bool)`

 SetCrlListNil sets the value for CrlList to be an explicit nil

### UnsetCrlList
`func (o *UpdateDomainRequest) UnsetCrlList()`

UnsetCrlList ensures that no value is present for CrlList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


