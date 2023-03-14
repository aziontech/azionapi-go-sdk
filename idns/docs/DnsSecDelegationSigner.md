# DnsSecDelegationSigner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DigestType** | Pointer to [**DnsSecDelegationSignerDigestType**](DnsSecDelegationSignerDigestType.md) |  | [optional] 
**AlgorithmType** | Pointer to [**DnsSecDelegationSignerDigestType**](DnsSecDelegationSignerDigestType.md) |  | [optional] 
**Digest** | Pointer to **string** |  | [optional] 
**KeyTag** | Pointer to **int32** |  | [optional] 

## Methods

### NewDnsSecDelegationSigner

`func NewDnsSecDelegationSigner() *DnsSecDelegationSigner`

NewDnsSecDelegationSigner instantiates a new DnsSecDelegationSigner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsSecDelegationSignerWithDefaults

`func NewDnsSecDelegationSignerWithDefaults() *DnsSecDelegationSigner`

NewDnsSecDelegationSignerWithDefaults instantiates a new DnsSecDelegationSigner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigestType

`func (o *DnsSecDelegationSigner) GetDigestType() DnsSecDelegationSignerDigestType`

GetDigestType returns the DigestType field if non-nil, zero value otherwise.

### GetDigestTypeOk

`func (o *DnsSecDelegationSigner) GetDigestTypeOk() (*DnsSecDelegationSignerDigestType, bool)`

GetDigestTypeOk returns a tuple with the DigestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestType

`func (o *DnsSecDelegationSigner) SetDigestType(v DnsSecDelegationSignerDigestType)`

SetDigestType sets DigestType field to given value.

### HasDigestType

`func (o *DnsSecDelegationSigner) HasDigestType() bool`

HasDigestType returns a boolean if a field has been set.

### GetAlgorithmType

`func (o *DnsSecDelegationSigner) GetAlgorithmType() DnsSecDelegationSignerDigestType`

GetAlgorithmType returns the AlgorithmType field if non-nil, zero value otherwise.

### GetAlgorithmTypeOk

`func (o *DnsSecDelegationSigner) GetAlgorithmTypeOk() (*DnsSecDelegationSignerDigestType, bool)`

GetAlgorithmTypeOk returns a tuple with the AlgorithmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmType

`func (o *DnsSecDelegationSigner) SetAlgorithmType(v DnsSecDelegationSignerDigestType)`

SetAlgorithmType sets AlgorithmType field to given value.

### HasAlgorithmType

`func (o *DnsSecDelegationSigner) HasAlgorithmType() bool`

HasAlgorithmType returns a boolean if a field has been set.

### GetDigest

`func (o *DnsSecDelegationSigner) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DnsSecDelegationSigner) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DnsSecDelegationSigner) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *DnsSecDelegationSigner) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetKeyTag

`func (o *DnsSecDelegationSigner) GetKeyTag() int32`

GetKeyTag returns the KeyTag field if non-nil, zero value otherwise.

### GetKeyTagOk

`func (o *DnsSecDelegationSigner) GetKeyTagOk() (*int32, bool)`

GetKeyTagOk returns a tuple with the KeyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTag

`func (o *DnsSecDelegationSigner) SetKeyTag(v int32)`

SetKeyTag sets KeyTag field to given value.

### HasKeyTag

`func (o *DnsSecDelegationSigner) HasKeyTag() bool`

HasKeyTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


