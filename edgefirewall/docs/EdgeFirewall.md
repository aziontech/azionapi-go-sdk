# EdgeFirewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**LastEditor** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**EdgeFunctionsEnabled** | Pointer to **bool** |  | [optional] 
**NetworkProtectionEnabled** | Pointer to **bool** |  | [optional] 
**WafEnabled** | Pointer to **bool** |  | [optional] 
**DebugRules** | Pointer to **bool** |  | [optional] 
**Domains** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewEdgeFirewall

`func NewEdgeFirewall() *EdgeFirewall`

NewEdgeFirewall instantiates a new EdgeFirewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFirewallWithDefaults

`func NewEdgeFirewallWithDefaults() *EdgeFirewall`

NewEdgeFirewallWithDefaults instantiates a new EdgeFirewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeFirewall) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeFirewall) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeFirewall) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EdgeFirewall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EdgeFirewall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFirewall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFirewall) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EdgeFirewall) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *EdgeFirewall) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *EdgeFirewall) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *EdgeFirewall) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *EdgeFirewall) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastEditor

`func (o *EdgeFirewall) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *EdgeFirewall) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *EdgeFirewall) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.

### HasLastEditor

`func (o *EdgeFirewall) HasLastEditor() bool`

HasLastEditor returns a boolean if a field has been set.

### GetLastModified

`func (o *EdgeFirewall) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *EdgeFirewall) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *EdgeFirewall) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *EdgeFirewall) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetEdgeFunctionsEnabled

`func (o *EdgeFirewall) GetEdgeFunctionsEnabled() bool`

GetEdgeFunctionsEnabled returns the EdgeFunctionsEnabled field if non-nil, zero value otherwise.

### GetEdgeFunctionsEnabledOk

`func (o *EdgeFirewall) GetEdgeFunctionsEnabledOk() (*bool, bool)`

GetEdgeFunctionsEnabledOk returns a tuple with the EdgeFunctionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctionsEnabled

`func (o *EdgeFirewall) SetEdgeFunctionsEnabled(v bool)`

SetEdgeFunctionsEnabled sets EdgeFunctionsEnabled field to given value.

### HasEdgeFunctionsEnabled

`func (o *EdgeFirewall) HasEdgeFunctionsEnabled() bool`

HasEdgeFunctionsEnabled returns a boolean if a field has been set.

### GetNetworkProtectionEnabled

`func (o *EdgeFirewall) GetNetworkProtectionEnabled() bool`

GetNetworkProtectionEnabled returns the NetworkProtectionEnabled field if non-nil, zero value otherwise.

### GetNetworkProtectionEnabledOk

`func (o *EdgeFirewall) GetNetworkProtectionEnabledOk() (*bool, bool)`

GetNetworkProtectionEnabledOk returns a tuple with the NetworkProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProtectionEnabled

`func (o *EdgeFirewall) SetNetworkProtectionEnabled(v bool)`

SetNetworkProtectionEnabled sets NetworkProtectionEnabled field to given value.

### HasNetworkProtectionEnabled

`func (o *EdgeFirewall) HasNetworkProtectionEnabled() bool`

HasNetworkProtectionEnabled returns a boolean if a field has been set.

### GetWafEnabled

`func (o *EdgeFirewall) GetWafEnabled() bool`

GetWafEnabled returns the WafEnabled field if non-nil, zero value otherwise.

### GetWafEnabledOk

`func (o *EdgeFirewall) GetWafEnabledOk() (*bool, bool)`

GetWafEnabledOk returns a tuple with the WafEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWafEnabled

`func (o *EdgeFirewall) SetWafEnabled(v bool)`

SetWafEnabled sets WafEnabled field to given value.

### HasWafEnabled

`func (o *EdgeFirewall) HasWafEnabled() bool`

HasWafEnabled returns a boolean if a field has been set.

### GetDebugRules

`func (o *EdgeFirewall) GetDebugRules() bool`

GetDebugRules returns the DebugRules field if non-nil, zero value otherwise.

### GetDebugRulesOk

`func (o *EdgeFirewall) GetDebugRulesOk() (*bool, bool)`

GetDebugRulesOk returns a tuple with the DebugRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugRules

`func (o *EdgeFirewall) SetDebugRules(v bool)`

SetDebugRules sets DebugRules field to given value.

### HasDebugRules

`func (o *EdgeFirewall) HasDebugRules() bool`

HasDebugRules returns a boolean if a field has been set.

### GetDomains

`func (o *EdgeFirewall) GetDomains() []int32`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *EdgeFirewall) GetDomainsOk() (*[]int32, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *EdgeFirewall) SetDomains(v []int32)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *EdgeFirewall) HasDomains() bool`

HasDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


