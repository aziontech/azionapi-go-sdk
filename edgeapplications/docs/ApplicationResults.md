# ApplicationResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Active** | **bool** |  | 
**DebugRules** | **bool** |  | 
**Http3** | **bool** |  | 
**SupportedCiphers** | **string** |  | 
**DeliveryProtocol** | **string** |  | 
**HttpPort** | **interface{}** |  | 
**HttpsPort** | **interface{}** |  | 
**MinimumTlsVersion** | **string** |  | 
**ApplicationAcceleration** | **bool** |  | 
**Caching** | **bool** |  | 
**DeviceDetection** | **bool** |  | 
**EdgeFirewall** | **bool** |  | 
**EdgeFunctions** | **bool** |  | 
**ImageOptimization** | **bool** |  | 
**L2Caching** | **bool** |  | 
**LoadBalancer** | **bool** |  | 
**RawLogs** | **bool** |  | 
**WebApplicationFirewall** | **bool** |  | 
**Websocket** | Pointer to **bool** |  | [optional] 

## Methods

### NewApplicationResults

`func NewApplicationResults(id int64, name string, active bool, debugRules bool, http3 bool, supportedCiphers string, deliveryProtocol string, httpPort interface{}, httpsPort interface{}, minimumTlsVersion string, applicationAcceleration bool, caching bool, deviceDetection bool, edgeFirewall bool, edgeFunctions bool, imageOptimization bool, l2Caching bool, loadBalancer bool, rawLogs bool, webApplicationFirewall bool, ) *ApplicationResults`

NewApplicationResults instantiates a new ApplicationResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResultsWithDefaults

`func NewApplicationResultsWithDefaults() *ApplicationResults`

NewApplicationResultsWithDefaults instantiates a new ApplicationResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationResults) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResults) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResults) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationResults) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationResults) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationResults) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *ApplicationResults) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationResults) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationResults) SetActive(v bool)`

SetActive sets Active field to given value.


### GetDebugRules

`func (o *ApplicationResults) GetDebugRules() bool`

GetDebugRules returns the DebugRules field if non-nil, zero value otherwise.

### GetDebugRulesOk

`func (o *ApplicationResults) GetDebugRulesOk() (*bool, bool)`

GetDebugRulesOk returns a tuple with the DebugRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugRules

`func (o *ApplicationResults) SetDebugRules(v bool)`

SetDebugRules sets DebugRules field to given value.


### GetHttp3

`func (o *ApplicationResults) GetHttp3() bool`

GetHttp3 returns the Http3 field if non-nil, zero value otherwise.

### GetHttp3Ok

`func (o *ApplicationResults) GetHttp3Ok() (*bool, bool)`

GetHttp3Ok returns a tuple with the Http3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp3

`func (o *ApplicationResults) SetHttp3(v bool)`

SetHttp3 sets Http3 field to given value.


### GetSupportedCiphers

`func (o *ApplicationResults) GetSupportedCiphers() string`

GetSupportedCiphers returns the SupportedCiphers field if non-nil, zero value otherwise.

### GetSupportedCiphersOk

`func (o *ApplicationResults) GetSupportedCiphersOk() (*string, bool)`

GetSupportedCiphersOk returns a tuple with the SupportedCiphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCiphers

`func (o *ApplicationResults) SetSupportedCiphers(v string)`

SetSupportedCiphers sets SupportedCiphers field to given value.


### GetDeliveryProtocol

`func (o *ApplicationResults) GetDeliveryProtocol() string`

GetDeliveryProtocol returns the DeliveryProtocol field if non-nil, zero value otherwise.

### GetDeliveryProtocolOk

`func (o *ApplicationResults) GetDeliveryProtocolOk() (*string, bool)`

GetDeliveryProtocolOk returns a tuple with the DeliveryProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryProtocol

`func (o *ApplicationResults) SetDeliveryProtocol(v string)`

SetDeliveryProtocol sets DeliveryProtocol field to given value.


### GetHttpPort

`func (o *ApplicationResults) GetHttpPort() interface{}`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *ApplicationResults) GetHttpPortOk() (*interface{}, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *ApplicationResults) SetHttpPort(v interface{})`

SetHttpPort sets HttpPort field to given value.


### SetHttpPortNil

`func (o *ApplicationResults) SetHttpPortNil(b bool)`

 SetHttpPortNil sets the value for HttpPort to be an explicit nil

### UnsetHttpPort
`func (o *ApplicationResults) UnsetHttpPort()`

UnsetHttpPort ensures that no value is present for HttpPort, not even an explicit nil
### GetHttpsPort

`func (o *ApplicationResults) GetHttpsPort() interface{}`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *ApplicationResults) GetHttpsPortOk() (*interface{}, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *ApplicationResults) SetHttpsPort(v interface{})`

SetHttpsPort sets HttpsPort field to given value.


### SetHttpsPortNil

`func (o *ApplicationResults) SetHttpsPortNil(b bool)`

 SetHttpsPortNil sets the value for HttpsPort to be an explicit nil

### UnsetHttpsPort
`func (o *ApplicationResults) UnsetHttpsPort()`

UnsetHttpsPort ensures that no value is present for HttpsPort, not even an explicit nil
### GetMinimumTlsVersion

`func (o *ApplicationResults) GetMinimumTlsVersion() string`

GetMinimumTlsVersion returns the MinimumTlsVersion field if non-nil, zero value otherwise.

### GetMinimumTlsVersionOk

`func (o *ApplicationResults) GetMinimumTlsVersionOk() (*string, bool)`

GetMinimumTlsVersionOk returns a tuple with the MinimumTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTlsVersion

`func (o *ApplicationResults) SetMinimumTlsVersion(v string)`

SetMinimumTlsVersion sets MinimumTlsVersion field to given value.


### GetApplicationAcceleration

`func (o *ApplicationResults) GetApplicationAcceleration() bool`

GetApplicationAcceleration returns the ApplicationAcceleration field if non-nil, zero value otherwise.

### GetApplicationAccelerationOk

`func (o *ApplicationResults) GetApplicationAccelerationOk() (*bool, bool)`

GetApplicationAccelerationOk returns a tuple with the ApplicationAcceleration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAcceleration

`func (o *ApplicationResults) SetApplicationAcceleration(v bool)`

SetApplicationAcceleration sets ApplicationAcceleration field to given value.


### GetCaching

`func (o *ApplicationResults) GetCaching() bool`

GetCaching returns the Caching field if non-nil, zero value otherwise.

### GetCachingOk

`func (o *ApplicationResults) GetCachingOk() (*bool, bool)`

GetCachingOk returns a tuple with the Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaching

`func (o *ApplicationResults) SetCaching(v bool)`

SetCaching sets Caching field to given value.


### GetDeviceDetection

`func (o *ApplicationResults) GetDeviceDetection() bool`

GetDeviceDetection returns the DeviceDetection field if non-nil, zero value otherwise.

### GetDeviceDetectionOk

`func (o *ApplicationResults) GetDeviceDetectionOk() (*bool, bool)`

GetDeviceDetectionOk returns a tuple with the DeviceDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDetection

`func (o *ApplicationResults) SetDeviceDetection(v bool)`

SetDeviceDetection sets DeviceDetection field to given value.


### GetEdgeFirewall

`func (o *ApplicationResults) GetEdgeFirewall() bool`

GetEdgeFirewall returns the EdgeFirewall field if non-nil, zero value otherwise.

### GetEdgeFirewallOk

`func (o *ApplicationResults) GetEdgeFirewallOk() (*bool, bool)`

GetEdgeFirewallOk returns a tuple with the EdgeFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFirewall

`func (o *ApplicationResults) SetEdgeFirewall(v bool)`

SetEdgeFirewall sets EdgeFirewall field to given value.


### GetEdgeFunctions

`func (o *ApplicationResults) GetEdgeFunctions() bool`

GetEdgeFunctions returns the EdgeFunctions field if non-nil, zero value otherwise.

### GetEdgeFunctionsOk

`func (o *ApplicationResults) GetEdgeFunctionsOk() (*bool, bool)`

GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctions

`func (o *ApplicationResults) SetEdgeFunctions(v bool)`

SetEdgeFunctions sets EdgeFunctions field to given value.


### GetImageOptimization

`func (o *ApplicationResults) GetImageOptimization() bool`

GetImageOptimization returns the ImageOptimization field if non-nil, zero value otherwise.

### GetImageOptimizationOk

`func (o *ApplicationResults) GetImageOptimizationOk() (*bool, bool)`

GetImageOptimizationOk returns a tuple with the ImageOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOptimization

`func (o *ApplicationResults) SetImageOptimization(v bool)`

SetImageOptimization sets ImageOptimization field to given value.


### GetL2Caching

`func (o *ApplicationResults) GetL2Caching() bool`

GetL2Caching returns the L2Caching field if non-nil, zero value otherwise.

### GetL2CachingOk

`func (o *ApplicationResults) GetL2CachingOk() (*bool, bool)`

GetL2CachingOk returns a tuple with the L2Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2Caching

`func (o *ApplicationResults) SetL2Caching(v bool)`

SetL2Caching sets L2Caching field to given value.


### GetLoadBalancer

`func (o *ApplicationResults) GetLoadBalancer() bool`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *ApplicationResults) GetLoadBalancerOk() (*bool, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *ApplicationResults) SetLoadBalancer(v bool)`

SetLoadBalancer sets LoadBalancer field to given value.


### GetRawLogs

`func (o *ApplicationResults) GetRawLogs() bool`

GetRawLogs returns the RawLogs field if non-nil, zero value otherwise.

### GetRawLogsOk

`func (o *ApplicationResults) GetRawLogsOk() (*bool, bool)`

GetRawLogsOk returns a tuple with the RawLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawLogs

`func (o *ApplicationResults) SetRawLogs(v bool)`

SetRawLogs sets RawLogs field to given value.


### GetWebApplicationFirewall

`func (o *ApplicationResults) GetWebApplicationFirewall() bool`

GetWebApplicationFirewall returns the WebApplicationFirewall field if non-nil, zero value otherwise.

### GetWebApplicationFirewallOk

`func (o *ApplicationResults) GetWebApplicationFirewallOk() (*bool, bool)`

GetWebApplicationFirewallOk returns a tuple with the WebApplicationFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApplicationFirewall

`func (o *ApplicationResults) SetWebApplicationFirewall(v bool)`

SetWebApplicationFirewall sets WebApplicationFirewall field to given value.


### GetWebsocket

`func (o *ApplicationResults) GetWebsocket() bool`

GetWebsocket returns the Websocket field if non-nil, zero value otherwise.

### GetWebsocketOk

`func (o *ApplicationResults) GetWebsocketOk() (*bool, bool)`

GetWebsocketOk returns a tuple with the Websocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocket

`func (o *ApplicationResults) SetWebsocket(v bool)`

SetWebsocket sets Websocket field to given value.

### HasWebsocket

`func (o *ApplicationResults) HasWebsocket() bool`

HasWebsocket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


