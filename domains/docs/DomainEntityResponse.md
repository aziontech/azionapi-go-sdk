# DomainEntityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Cnames** | Pointer to **[]string** |  | [optional] 
**CnameAccessOnly** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**EdgeApplicationId** | Pointer to **int64** |  | [optional] 
**DigitalCertificateId** | Pointer to **int64** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**IsMtlsEnabled** | Pointer to **bool** |  | [optional] 
**MtlsTrustedCaCertificateId** | Pointer to **NullableInt64** |  | [optional] 
**EdgeFirewallId** | Pointer to **NullableInt64** |  | [optional] 
**MtlsVerification** | Pointer to **string** |  | [optional] 
**CrlList** | Pointer to **[]int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainEntityResponse

`func NewDomainEntityResponse() *DomainEntityResponse`

NewDomainEntityResponse instantiates a new DomainEntityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainEntityResponseWithDefaults

`func NewDomainEntityResponseWithDefaults() *DomainEntityResponse`

NewDomainEntityResponseWithDefaults instantiates a new DomainEntityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DomainEntityResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainEntityResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainEntityResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainEntityResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCnames

`func (o *DomainEntityResponse) GetCnames() []string`

GetCnames returns the Cnames field if non-nil, zero value otherwise.

### GetCnamesOk

`func (o *DomainEntityResponse) GetCnamesOk() (*[]string, bool)`

GetCnamesOk returns a tuple with the Cnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnames

`func (o *DomainEntityResponse) SetCnames(v []string)`

SetCnames sets Cnames field to given value.

### HasCnames

`func (o *DomainEntityResponse) HasCnames() bool`

HasCnames returns a boolean if a field has been set.

### GetCnameAccessOnly

`func (o *DomainEntityResponse) GetCnameAccessOnly() bool`

GetCnameAccessOnly returns the CnameAccessOnly field if non-nil, zero value otherwise.

### GetCnameAccessOnlyOk

`func (o *DomainEntityResponse) GetCnameAccessOnlyOk() (*bool, bool)`

GetCnameAccessOnlyOk returns a tuple with the CnameAccessOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameAccessOnly

`func (o *DomainEntityResponse) SetCnameAccessOnly(v bool)`

SetCnameAccessOnly sets CnameAccessOnly field to given value.

### HasCnameAccessOnly

`func (o *DomainEntityResponse) HasCnameAccessOnly() bool`

HasCnameAccessOnly returns a boolean if a field has been set.

### GetIsActive

`func (o *DomainEntityResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DomainEntityResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DomainEntityResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DomainEntityResponse) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEdgeApplicationId

`func (o *DomainEntityResponse) GetEdgeApplicationId() int64`

GetEdgeApplicationId returns the EdgeApplicationId field if non-nil, zero value otherwise.

### GetEdgeApplicationIdOk

`func (o *DomainEntityResponse) GetEdgeApplicationIdOk() (*int64, bool)`

GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplicationId

`func (o *DomainEntityResponse) SetEdgeApplicationId(v int64)`

SetEdgeApplicationId sets EdgeApplicationId field to given value.

### HasEdgeApplicationId

`func (o *DomainEntityResponse) HasEdgeApplicationId() bool`

HasEdgeApplicationId returns a boolean if a field has been set.

### GetDigitalCertificateId

`func (o *DomainEntityResponse) GetDigitalCertificateId() int64`

GetDigitalCertificateId returns the DigitalCertificateId field if non-nil, zero value otherwise.

### GetDigitalCertificateIdOk

`func (o *DomainEntityResponse) GetDigitalCertificateIdOk() (*int64, bool)`

GetDigitalCertificateIdOk returns a tuple with the DigitalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalCertificateId

`func (o *DomainEntityResponse) SetDigitalCertificateId(v int64)`

SetDigitalCertificateId sets DigitalCertificateId field to given value.

### HasDigitalCertificateId

`func (o *DomainEntityResponse) HasDigitalCertificateId() bool`

HasDigitalCertificateId returns a boolean if a field has been set.

### GetEnvironment

`func (o *DomainEntityResponse) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DomainEntityResponse) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DomainEntityResponse) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DomainEntityResponse) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIsMtlsEnabled

`func (o *DomainEntityResponse) GetIsMtlsEnabled() bool`

GetIsMtlsEnabled returns the IsMtlsEnabled field if non-nil, zero value otherwise.

### GetIsMtlsEnabledOk

`func (o *DomainEntityResponse) GetIsMtlsEnabledOk() (*bool, bool)`

GetIsMtlsEnabledOk returns a tuple with the IsMtlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMtlsEnabled

`func (o *DomainEntityResponse) SetIsMtlsEnabled(v bool)`

SetIsMtlsEnabled sets IsMtlsEnabled field to given value.

### HasIsMtlsEnabled

`func (o *DomainEntityResponse) HasIsMtlsEnabled() bool`

HasIsMtlsEnabled returns a boolean if a field has been set.

### GetMtlsTrustedCaCertificateId

`func (o *DomainEntityResponse) GetMtlsTrustedCaCertificateId() int64`

GetMtlsTrustedCaCertificateId returns the MtlsTrustedCaCertificateId field if non-nil, zero value otherwise.

### GetMtlsTrustedCaCertificateIdOk

`func (o *DomainEntityResponse) GetMtlsTrustedCaCertificateIdOk() (*int64, bool)`

GetMtlsTrustedCaCertificateIdOk returns a tuple with the MtlsTrustedCaCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsTrustedCaCertificateId

`func (o *DomainEntityResponse) SetMtlsTrustedCaCertificateId(v int64)`

SetMtlsTrustedCaCertificateId sets MtlsTrustedCaCertificateId field to given value.

### HasMtlsTrustedCaCertificateId

`func (o *DomainEntityResponse) HasMtlsTrustedCaCertificateId() bool`

HasMtlsTrustedCaCertificateId returns a boolean if a field has been set.

### SetMtlsTrustedCaCertificateIdNil

`func (o *DomainEntityResponse) SetMtlsTrustedCaCertificateIdNil(b bool)`

 SetMtlsTrustedCaCertificateIdNil sets the value for MtlsTrustedCaCertificateId to be an explicit nil

### UnsetMtlsTrustedCaCertificateId
`func (o *DomainEntityResponse) UnsetMtlsTrustedCaCertificateId()`

UnsetMtlsTrustedCaCertificateId ensures that no value is present for MtlsTrustedCaCertificateId, not even an explicit nil
### GetEdgeFirewallId

`func (o *DomainEntityResponse) GetEdgeFirewallId() int64`

GetEdgeFirewallId returns the EdgeFirewallId field if non-nil, zero value otherwise.

### GetEdgeFirewallIdOk

`func (o *DomainEntityResponse) GetEdgeFirewallIdOk() (*int64, bool)`

GetEdgeFirewallIdOk returns a tuple with the EdgeFirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFirewallId

`func (o *DomainEntityResponse) SetEdgeFirewallId(v int64)`

SetEdgeFirewallId sets EdgeFirewallId field to given value.

### HasEdgeFirewallId

`func (o *DomainEntityResponse) HasEdgeFirewallId() bool`

HasEdgeFirewallId returns a boolean if a field has been set.

### SetEdgeFirewallIdNil

`func (o *DomainEntityResponse) SetEdgeFirewallIdNil(b bool)`

 SetEdgeFirewallIdNil sets the value for EdgeFirewallId to be an explicit nil

### UnsetEdgeFirewallId
`func (o *DomainEntityResponse) UnsetEdgeFirewallId()`

UnsetEdgeFirewallId ensures that no value is present for EdgeFirewallId, not even an explicit nil
### GetMtlsVerification

`func (o *DomainEntityResponse) GetMtlsVerification() string`

GetMtlsVerification returns the MtlsVerification field if non-nil, zero value otherwise.

### GetMtlsVerificationOk

`func (o *DomainEntityResponse) GetMtlsVerificationOk() (*string, bool)`

GetMtlsVerificationOk returns a tuple with the MtlsVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtlsVerification

`func (o *DomainEntityResponse) SetMtlsVerification(v string)`

SetMtlsVerification sets MtlsVerification field to given value.

### HasMtlsVerification

`func (o *DomainEntityResponse) HasMtlsVerification() bool`

HasMtlsVerification returns a boolean if a field has been set.

### GetCrlList

`func (o *DomainEntityResponse) GetCrlList() []int64`

GetCrlList returns the CrlList field if non-nil, zero value otherwise.

### GetCrlListOk

`func (o *DomainEntityResponse) GetCrlListOk() (*[]int64, bool)`

GetCrlListOk returns a tuple with the CrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrlList

`func (o *DomainEntityResponse) SetCrlList(v []int64)`

SetCrlList sets CrlList field to given value.

### HasCrlList

`func (o *DomainEntityResponse) HasCrlList() bool`

HasCrlList returns a boolean if a field has been set.

### SetCrlListNil

`func (o *DomainEntityResponse) SetCrlListNil(b bool)`

 SetCrlListNil sets the value for CrlList to be an explicit nil

### UnsetCrlList
`func (o *DomainEntityResponse) UnsetCrlList()`

UnsetCrlList ensures that no value is present for CrlList, not even an explicit nil
### GetId

`func (o *DomainEntityResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainEntityResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainEntityResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DomainEntityResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomainName

`func (o *DomainEntityResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DomainEntityResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DomainEntityResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DomainEntityResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


