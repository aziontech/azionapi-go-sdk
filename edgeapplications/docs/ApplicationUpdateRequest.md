# ApplicationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DeliveryProtocol** | Pointer to **string** |  | [optional] 
**HttpPort** | Pointer to **interface{}** |  | [optional] 
**HttpsPort** | Pointer to **interface{}** |  | [optional] 
**MinimumTlsVersion** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DebugRules** | Pointer to **bool** |  | [optional] 
**ApplicationAcceleration** | Pointer to **bool** |  | [optional] 
**Caching** | Pointer to **bool** |  | [optional] 
**DeviceDetection** | Pointer to **bool** |  | [optional] 
**EdgeFirewall** | Pointer to **bool** |  | [optional] 
**EdgeFunctions** | Pointer to **bool** |  | [optional] 
**ImageOptimization** | Pointer to **bool** |  | [optional] 
**L2Caching** | Pointer to **bool** |  | [optional] 
**LoadBalancer** | Pointer to **bool** |  | [optional] 
**RawLogs** | Pointer to **bool** |  | [optional] 
**WebApplicationFirewall** | Pointer to **bool** |  | [optional] 
**Websocket** | Pointer to **bool** |  | [optional] 

## Methods

### NewApplicationUpdateRequest

`func NewApplicationUpdateRequest() *ApplicationUpdateRequest`

NewApplicationUpdateRequest instantiates a new ApplicationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationUpdateRequestWithDefaults

`func NewApplicationUpdateRequestWithDefaults() *ApplicationUpdateRequest`

NewApplicationUpdateRequestWithDefaults instantiates a new ApplicationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeliveryProtocol

`func (o *ApplicationUpdateRequest) GetDeliveryProtocol() string`

GetDeliveryProtocol returns the DeliveryProtocol field if non-nil, zero value otherwise.

### GetDeliveryProtocolOk

`func (o *ApplicationUpdateRequest) GetDeliveryProtocolOk() (*string, bool)`

GetDeliveryProtocolOk returns a tuple with the DeliveryProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryProtocol

`func (o *ApplicationUpdateRequest) SetDeliveryProtocol(v string)`

SetDeliveryProtocol sets DeliveryProtocol field to given value.

### HasDeliveryProtocol

`func (o *ApplicationUpdateRequest) HasDeliveryProtocol() bool`

HasDeliveryProtocol returns a boolean if a field has been set.

### GetHttpPort

`func (o *ApplicationUpdateRequest) GetHttpPort() interface{}`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *ApplicationUpdateRequest) GetHttpPortOk() (*interface{}, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *ApplicationUpdateRequest) SetHttpPort(v interface{})`

SetHttpPort sets HttpPort field to given value.

### HasHttpPort

`func (o *ApplicationUpdateRequest) HasHttpPort() bool`

HasHttpPort returns a boolean if a field has been set.

### SetHttpPortNil

`func (o *ApplicationUpdateRequest) SetHttpPortNil(b bool)`

 SetHttpPortNil sets the value for HttpPort to be an explicit nil

### UnsetHttpPort
`func (o *ApplicationUpdateRequest) UnsetHttpPort()`

UnsetHttpPort ensures that no value is present for HttpPort, not even an explicit nil
### GetHttpsPort

`func (o *ApplicationUpdateRequest) GetHttpsPort() interface{}`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *ApplicationUpdateRequest) GetHttpsPortOk() (*interface{}, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *ApplicationUpdateRequest) SetHttpsPort(v interface{})`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *ApplicationUpdateRequest) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### SetHttpsPortNil

`func (o *ApplicationUpdateRequest) SetHttpsPortNil(b bool)`

 SetHttpsPortNil sets the value for HttpsPort to be an explicit nil

### UnsetHttpsPort
`func (o *ApplicationUpdateRequest) UnsetHttpsPort()`

UnsetHttpsPort ensures that no value is present for HttpsPort, not even an explicit nil
### GetMinimumTlsVersion

`func (o *ApplicationUpdateRequest) GetMinimumTlsVersion() string`

GetMinimumTlsVersion returns the MinimumTlsVersion field if non-nil, zero value otherwise.

### GetMinimumTlsVersionOk

`func (o *ApplicationUpdateRequest) GetMinimumTlsVersionOk() (*string, bool)`

GetMinimumTlsVersionOk returns a tuple with the MinimumTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTlsVersion

`func (o *ApplicationUpdateRequest) SetMinimumTlsVersion(v string)`

SetMinimumTlsVersion sets MinimumTlsVersion field to given value.

### HasMinimumTlsVersion

`func (o *ApplicationUpdateRequest) HasMinimumTlsVersion() bool`

HasMinimumTlsVersion returns a boolean if a field has been set.

### GetActive

`func (o *ApplicationUpdateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationUpdateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationUpdateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApplicationUpdateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDebugRules

`func (o *ApplicationUpdateRequest) GetDebugRules() bool`

GetDebugRules returns the DebugRules field if non-nil, zero value otherwise.

### GetDebugRulesOk

`func (o *ApplicationUpdateRequest) GetDebugRulesOk() (*bool, bool)`

GetDebugRulesOk returns a tuple with the DebugRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugRules

`func (o *ApplicationUpdateRequest) SetDebugRules(v bool)`

SetDebugRules sets DebugRules field to given value.

### HasDebugRules

`func (o *ApplicationUpdateRequest) HasDebugRules() bool`

HasDebugRules returns a boolean if a field has been set.

### GetApplicationAcceleration

`func (o *ApplicationUpdateRequest) GetApplicationAcceleration() bool`

GetApplicationAcceleration returns the ApplicationAcceleration field if non-nil, zero value otherwise.

### GetApplicationAccelerationOk

`func (o *ApplicationUpdateRequest) GetApplicationAccelerationOk() (*bool, bool)`

GetApplicationAccelerationOk returns a tuple with the ApplicationAcceleration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAcceleration

`func (o *ApplicationUpdateRequest) SetApplicationAcceleration(v bool)`

SetApplicationAcceleration sets ApplicationAcceleration field to given value.

### HasApplicationAcceleration

`func (o *ApplicationUpdateRequest) HasApplicationAcceleration() bool`

HasApplicationAcceleration returns a boolean if a field has been set.

### GetCaching

`func (o *ApplicationUpdateRequest) GetCaching() bool`

GetCaching returns the Caching field if non-nil, zero value otherwise.

### GetCachingOk

`func (o *ApplicationUpdateRequest) GetCachingOk() (*bool, bool)`

GetCachingOk returns a tuple with the Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaching

`func (o *ApplicationUpdateRequest) SetCaching(v bool)`

SetCaching sets Caching field to given value.

### HasCaching

`func (o *ApplicationUpdateRequest) HasCaching() bool`

HasCaching returns a boolean if a field has been set.

### GetDeviceDetection

`func (o *ApplicationUpdateRequest) GetDeviceDetection() bool`

GetDeviceDetection returns the DeviceDetection field if non-nil, zero value otherwise.

### GetDeviceDetectionOk

`func (o *ApplicationUpdateRequest) GetDeviceDetectionOk() (*bool, bool)`

GetDeviceDetectionOk returns a tuple with the DeviceDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDetection

`func (o *ApplicationUpdateRequest) SetDeviceDetection(v bool)`

SetDeviceDetection sets DeviceDetection field to given value.

### HasDeviceDetection

`func (o *ApplicationUpdateRequest) HasDeviceDetection() bool`

HasDeviceDetection returns a boolean if a field has been set.

### GetEdgeFirewall

`func (o *ApplicationUpdateRequest) GetEdgeFirewall() bool`

GetEdgeFirewall returns the EdgeFirewall field if non-nil, zero value otherwise.

### GetEdgeFirewallOk

`func (o *ApplicationUpdateRequest) GetEdgeFirewallOk() (*bool, bool)`

GetEdgeFirewallOk returns a tuple with the EdgeFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFirewall

`func (o *ApplicationUpdateRequest) SetEdgeFirewall(v bool)`

SetEdgeFirewall sets EdgeFirewall field to given value.

### HasEdgeFirewall

`func (o *ApplicationUpdateRequest) HasEdgeFirewall() bool`

HasEdgeFirewall returns a boolean if a field has been set.

### GetEdgeFunctions

`func (o *ApplicationUpdateRequest) GetEdgeFunctions() bool`

GetEdgeFunctions returns the EdgeFunctions field if non-nil, zero value otherwise.

### GetEdgeFunctionsOk

`func (o *ApplicationUpdateRequest) GetEdgeFunctionsOk() (*bool, bool)`

GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctions

`func (o *ApplicationUpdateRequest) SetEdgeFunctions(v bool)`

SetEdgeFunctions sets EdgeFunctions field to given value.

### HasEdgeFunctions

`func (o *ApplicationUpdateRequest) HasEdgeFunctions() bool`

HasEdgeFunctions returns a boolean if a field has been set.

### GetImageOptimization

`func (o *ApplicationUpdateRequest) GetImageOptimization() bool`

GetImageOptimization returns the ImageOptimization field if non-nil, zero value otherwise.

### GetImageOptimizationOk

`func (o *ApplicationUpdateRequest) GetImageOptimizationOk() (*bool, bool)`

GetImageOptimizationOk returns a tuple with the ImageOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOptimization

`func (o *ApplicationUpdateRequest) SetImageOptimization(v bool)`

SetImageOptimization sets ImageOptimization field to given value.

### HasImageOptimization

`func (o *ApplicationUpdateRequest) HasImageOptimization() bool`

HasImageOptimization returns a boolean if a field has been set.

### GetL2Caching

`func (o *ApplicationUpdateRequest) GetL2Caching() bool`

GetL2Caching returns the L2Caching field if non-nil, zero value otherwise.

### GetL2CachingOk

`func (o *ApplicationUpdateRequest) GetL2CachingOk() (*bool, bool)`

GetL2CachingOk returns a tuple with the L2Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2Caching

`func (o *ApplicationUpdateRequest) SetL2Caching(v bool)`

SetL2Caching sets L2Caching field to given value.

### HasL2Caching

`func (o *ApplicationUpdateRequest) HasL2Caching() bool`

HasL2Caching returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *ApplicationUpdateRequest) GetLoadBalancer() bool`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *ApplicationUpdateRequest) GetLoadBalancerOk() (*bool, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *ApplicationUpdateRequest) SetLoadBalancer(v bool)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *ApplicationUpdateRequest) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetRawLogs

`func (o *ApplicationUpdateRequest) GetRawLogs() bool`

GetRawLogs returns the RawLogs field if non-nil, zero value otherwise.

### GetRawLogsOk

`func (o *ApplicationUpdateRequest) GetRawLogsOk() (*bool, bool)`

GetRawLogsOk returns a tuple with the RawLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawLogs

`func (o *ApplicationUpdateRequest) SetRawLogs(v bool)`

SetRawLogs sets RawLogs field to given value.

### HasRawLogs

`func (o *ApplicationUpdateRequest) HasRawLogs() bool`

HasRawLogs returns a boolean if a field has been set.

### GetWebApplicationFirewall

`func (o *ApplicationUpdateRequest) GetWebApplicationFirewall() bool`

GetWebApplicationFirewall returns the WebApplicationFirewall field if non-nil, zero value otherwise.

### GetWebApplicationFirewallOk

`func (o *ApplicationUpdateRequest) GetWebApplicationFirewallOk() (*bool, bool)`

GetWebApplicationFirewallOk returns a tuple with the WebApplicationFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApplicationFirewall

`func (o *ApplicationUpdateRequest) SetWebApplicationFirewall(v bool)`

SetWebApplicationFirewall sets WebApplicationFirewall field to given value.

### HasWebApplicationFirewall

`func (o *ApplicationUpdateRequest) HasWebApplicationFirewall() bool`

HasWebApplicationFirewall returns a boolean if a field has been set.

### GetWebsocket

`func (o *ApplicationUpdateRequest) GetWebsocket() bool`

GetWebsocket returns the Websocket field if non-nil, zero value otherwise.

### GetWebsocketOk

`func (o *ApplicationUpdateRequest) GetWebsocketOk() (*bool, bool)`

GetWebsocketOk returns a tuple with the Websocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocket

`func (o *ApplicationUpdateRequest) SetWebsocket(v bool)`

SetWebsocket sets Websocket field to given value.

### HasWebsocket

`func (o *ApplicationUpdateRequest) HasWebsocket() bool`

HasWebsocket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


