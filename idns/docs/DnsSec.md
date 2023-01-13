# DnsSec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DelegationSigner** | Pointer to [**DnsSecDelegationSigner**](DnsSecDelegationSigner.md) |  | [optional] 
**AlgorithmType** | Pointer to [**DnsSecDelegationSignerDigestType**](DnsSecDelegationSignerDigestType.md) |  | [optional] 
**Digest** | Pointer to **string** |  | [optional] 
**KeyTag** | Pointer to **int32** |  | [optional] 

## Methods

### NewDnsSec

`func NewDnsSec() *DnsSec`

NewDnsSec instantiates a new DnsSec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsSecWithDefaults

`func NewDnsSecWithDefaults() *DnsSec`

NewDnsSecWithDefaults instantiates a new DnsSec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *DnsSec) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DnsSec) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DnsSec) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DnsSec) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *DnsSec) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DnsSec) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DnsSec) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DnsSec) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDelegationSigner

`func (o *DnsSec) GetDelegationSigner() DnsSecDelegationSigner`

GetDelegationSigner returns the DelegationSigner field if non-nil, zero value otherwise.

### GetDelegationSignerOk

`func (o *DnsSec) GetDelegationSignerOk() (*DnsSecDelegationSigner, bool)`

GetDelegationSignerOk returns a tuple with the DelegationSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationSigner

`func (o *DnsSec) SetDelegationSigner(v DnsSecDelegationSigner)`

SetDelegationSigner sets DelegationSigner field to given value.

### HasDelegationSigner

`func (o *DnsSec) HasDelegationSigner() bool`

HasDelegationSigner returns a boolean if a field has been set.

### GetAlgorithmType

`func (o *DnsSec) GetAlgorithmType() DnsSecDelegationSignerDigestType`

GetAlgorithmType returns the AlgorithmType field if non-nil, zero value otherwise.

### GetAlgorithmTypeOk

`func (o *DnsSec) GetAlgorithmTypeOk() (*DnsSecDelegationSignerDigestType, bool)`

GetAlgorithmTypeOk returns a tuple with the AlgorithmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmType

`func (o *DnsSec) SetAlgorithmType(v DnsSecDelegationSignerDigestType)`

SetAlgorithmType sets AlgorithmType field to given value.

### HasAlgorithmType

`func (o *DnsSec) HasAlgorithmType() bool`

HasAlgorithmType returns a boolean if a field has been set.

### GetDigest

`func (o *DnsSec) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DnsSec) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DnsSec) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *DnsSec) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetKeyTag

`func (o *DnsSec) GetKeyTag() int32`

GetKeyTag returns the KeyTag field if non-nil, zero value otherwise.

### GetKeyTagOk

`func (o *DnsSec) GetKeyTagOk() (*int32, bool)`

GetKeyTagOk returns a tuple with the KeyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTag

`func (o *DnsSec) SetKeyTag(v int32)`

SetKeyTag sets KeyTag field to given value.

### HasKeyTag

`func (o *DnsSec) HasKeyTag() bool`

HasKeyTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


