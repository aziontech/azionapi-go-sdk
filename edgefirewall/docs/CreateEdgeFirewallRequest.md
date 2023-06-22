# CreateEdgeFirewallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Domains** | Pointer to **[]int32** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**EdgeFunctionsEnabled** | Pointer to **bool** |  | [optional] 
**NetworkProtectionEnabled** | Pointer to **bool** |  | [optional] 
**WafEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateEdgeFirewallRequest

`func NewCreateEdgeFirewallRequest() *CreateEdgeFirewallRequest`

NewCreateEdgeFirewallRequest instantiates a new CreateEdgeFirewallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEdgeFirewallRequestWithDefaults

`func NewCreateEdgeFirewallRequestWithDefaults() *CreateEdgeFirewallRequest`

NewCreateEdgeFirewallRequestWithDefaults instantiates a new CreateEdgeFirewallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEdgeFirewallRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEdgeFirewallRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEdgeFirewallRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateEdgeFirewallRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomains

`func (o *CreateEdgeFirewallRequest) GetDomains() []int32`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CreateEdgeFirewallRequest) GetDomainsOk() (*[]int32, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CreateEdgeFirewallRequest) SetDomains(v []int32)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *CreateEdgeFirewallRequest) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetIsActive

`func (o *CreateEdgeFirewallRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateEdgeFirewallRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateEdgeFirewallRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CreateEdgeFirewallRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEdgeFunctionsEnabled

`func (o *CreateEdgeFirewallRequest) GetEdgeFunctionsEnabled() bool`

GetEdgeFunctionsEnabled returns the EdgeFunctionsEnabled field if non-nil, zero value otherwise.

### GetEdgeFunctionsEnabledOk

`func (o *CreateEdgeFirewallRequest) GetEdgeFunctionsEnabledOk() (*bool, bool)`

GetEdgeFunctionsEnabledOk returns a tuple with the EdgeFunctionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctionsEnabled

`func (o *CreateEdgeFirewallRequest) SetEdgeFunctionsEnabled(v bool)`

SetEdgeFunctionsEnabled sets EdgeFunctionsEnabled field to given value.

### HasEdgeFunctionsEnabled

`func (o *CreateEdgeFirewallRequest) HasEdgeFunctionsEnabled() bool`

HasEdgeFunctionsEnabled returns a boolean if a field has been set.

### GetNetworkProtectionEnabled

`func (o *CreateEdgeFirewallRequest) GetNetworkProtectionEnabled() bool`

GetNetworkProtectionEnabled returns the NetworkProtectionEnabled field if non-nil, zero value otherwise.

### GetNetworkProtectionEnabledOk

`func (o *CreateEdgeFirewallRequest) GetNetworkProtectionEnabledOk() (*bool, bool)`

GetNetworkProtectionEnabledOk returns a tuple with the NetworkProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtectionEnabled

`func (o *CreateEdgeFirewallRequest) SetNetworkProtectionEnabled(v bool)`

SetNetworkProtectionEnabled sets NetworkProtectionEnabled field to given value.

### HasNetworkProtectionEnabled

`func (o *CreateEdgeFirewallRequest) HasNetworkProtectionEnabled() bool`

HasNetworkProtectionEnabled returns a boolean if a field has been set.

### GetWafEnabled

`func (o *CreateEdgeFirewallRequest) GetWafEnabled() bool`

GetWafEnabled returns the WafEnabled field if non-nil, zero value otherwise.

### GetWafEnabledOk

`func (o *CreateEdgeFirewallRequest) GetWafEnabledOk() (*bool, bool)`

GetWafEnabledOk returns a tuple with the WafEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafEnabled

`func (o *CreateEdgeFirewallRequest) SetWafEnabled(v bool)`

SetWafEnabled sets WafEnabled field to given value.

### HasWafEnabled

`func (o *CreateEdgeFirewallRequest) HasWafEnabled() bool`

HasWafEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


