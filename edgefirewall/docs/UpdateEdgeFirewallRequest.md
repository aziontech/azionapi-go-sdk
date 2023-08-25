# UpdateEdgeFirewallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]int64** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**EdgeFunctionsEnabled** | Pointer to **bool** |  | [optional] 
**NetworkProtectionEnabled** | Pointer to **bool** |  | [optional] 
**WafEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateEdgeFirewallRequest

`func NewUpdateEdgeFirewallRequest() *UpdateEdgeFirewallRequest`

NewUpdateEdgeFirewallRequest instantiates a new UpdateEdgeFirewallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEdgeFirewallRequestWithDefaults

`func NewUpdateEdgeFirewallRequestWithDefaults() *UpdateEdgeFirewallRequest`

NewUpdateEdgeFirewallRequestWithDefaults instantiates a new UpdateEdgeFirewallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateEdgeFirewallRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEdgeFirewallRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEdgeFirewallRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateEdgeFirewallRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomains

`func (o *UpdateEdgeFirewallRequest) GetDomains() []int64`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *UpdateEdgeFirewallRequest) GetDomainsOk() (*[]int64, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *UpdateEdgeFirewallRequest) SetDomains(v []int64)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *UpdateEdgeFirewallRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateEdgeFirewallRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateEdgeFirewallRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateEdgeFirewallRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateEdgeFirewallRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEdgeFunctionsEnabled

`func (o *UpdateEdgeFirewallRequest) GetEdgeFunctionsEnabled() bool`

GetEdgeFunctionsEnabled returns the EdgeFunctionsEnabled field if non-nil, zero value otherwise.

### GetEdgeFunctionsEnabledOk

`func (o *UpdateEdgeFirewallRequest) GetEdgeFunctionsEnabledOk() (*bool, bool)`

GetEdgeFunctionsEnabledOk returns a tuple with the EdgeFunctionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctionsEnabled

`func (o *UpdateEdgeFirewallRequest) SetEdgeFunctionsEnabled(v bool)`

SetEdgeFunctionsEnabled sets EdgeFunctionsEnabled field to given value.

### HasEdgeFunctionsEnabled

`func (o *UpdateEdgeFirewallRequest) HasEdgeFunctionsEnabled() bool`

HasEdgeFunctionsEnabled returns a boolean if a field has been set.

### GetNetworkProtectionEnabled

`func (o *UpdateEdgeFirewallRequest) GetNetworkProtectionEnabled() bool`

GetNetworkProtectionEnabled returns the NetworkProtectionEnabled field if non-nil, zero value otherwise.

### GetNetworkProtectionEnabledOk

`func (o *UpdateEdgeFirewallRequest) GetNetworkProtectionEnabledOk() (*bool, bool)`

GetNetworkProtectionEnabledOk returns a tuple with the NetworkProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtectionEnabled

`func (o *UpdateEdgeFirewallRequest) SetNetworkProtectionEnabled(v bool)`

SetNetworkProtectionEnabled sets NetworkProtectionEnabled field to given value.

### HasNetworkProtectionEnabled

`func (o *UpdateEdgeFirewallRequest) HasNetworkProtectionEnabled() bool`

HasNetworkProtectionEnabled returns a boolean if a field has been set.

### GetWafEnabled

`func (o *UpdateEdgeFirewallRequest) GetWafEnabled() bool`

GetWafEnabled returns the WafEnabled field if non-nil, zero value otherwise.

### GetWafEnabledOk

`func (o *UpdateEdgeFirewallRequest) GetWafEnabledOk() (*bool, bool)`

GetWafEnabledOk returns a tuple with the WafEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafEnabled

`func (o *UpdateEdgeFirewallRequest) SetWafEnabled(v bool)`

SetWafEnabled sets WafEnabled field to given value.

### HasWafEnabled

`func (o *UpdateEdgeFirewallRequest) HasWafEnabled() bool`

HasWafEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


