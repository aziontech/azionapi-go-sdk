# DnsSec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DelegationSigner** | Pointer to [**DnsSecDelegationSigner**](DnsSecDelegationSigner.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


