# ApplicationPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**DeliveryProtocol** | Pointer to **string** |  | [optional] 
**HttpPort** | Pointer to **int64** |  | [optional] 
**HttpsPort** | Pointer to **int64** |  | [optional] 
**MinimumTlsVersion** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
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

## Methods

### NewApplicationPutRequest

`func NewApplicationPutRequest(name string, ) *ApplicationPutRequest`

NewApplicationPutRequest instantiates a new ApplicationPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPutRequestWithDefaults

`func NewApplicationPutRequestWithDefaults() *ApplicationPutRequest`

NewApplicationPutRequestWithDefaults instantiates a new ApplicationPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationPutRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationPutRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationPutRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationPutRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApplicationPutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationPutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationPutRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDeliveryProtocol

`func (o *ApplicationPutRequest) GetDeliveryProtocol() string`

GetDeliveryProtocol returns the DeliveryProtocol field if non-nil, zero value otherwise.

### GetDeliveryProtocolOk

`func (o *ApplicationPutRequest) GetDeliveryProtocolOk() (*string, bool)`

GetDeliveryProtocolOk returns a tuple with the DeliveryProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryProtocol

`func (o *ApplicationPutRequest) SetDeliveryProtocol(v string)`

SetDeliveryProtocol sets DeliveryProtocol field to given value.

### HasDeliveryProtocol

`func (o *ApplicationPutRequest) HasDeliveryProtocol() bool`

HasDeliveryProtocol returns a boolean if a field has been set.

### GetHttpPort

`func (o *ApplicationPutRequest) GetHttpPort() int64`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *ApplicationPutRequest) GetHttpPortOk() (*int64, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *ApplicationPutRequest) SetHttpPort(v int64)`

SetHttpPort sets HttpPort field to given value.

### HasHttpPort

`func (o *ApplicationPutRequest) HasHttpPort() bool`

HasHttpPort returns a boolean if a field has been set.

### GetHttpsPort

`func (o *ApplicationPutRequest) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *ApplicationPutRequest) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *ApplicationPutRequest) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *ApplicationPutRequest) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetMinimumTlsVersion

`func (o *ApplicationPutRequest) GetMinimumTlsVersion() string`

GetMinimumTlsVersion returns the MinimumTlsVersion field if non-nil, zero value otherwise.

### GetMinimumTlsVersionOk

`func (o *ApplicationPutRequest) GetMinimumTlsVersionOk() (*string, bool)`

GetMinimumTlsVersionOk returns a tuple with the MinimumTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTlsVersion

`func (o *ApplicationPutRequest) SetMinimumTlsVersion(v string)`

SetMinimumTlsVersion sets MinimumTlsVersion field to given value.

### HasMinimumTlsVersion

`func (o *ApplicationPutRequest) HasMinimumTlsVersion() bool`

HasMinimumTlsVersion returns a boolean if a field has been set.

### GetActive

`func (o *ApplicationPutRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationPutRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationPutRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApplicationPutRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetApplicationAcceleration

`func (o *ApplicationPutRequest) GetApplicationAcceleration() bool`

GetApplicationAcceleration returns the ApplicationAcceleration field if non-nil, zero value otherwise.

### GetApplicationAccelerationOk

`func (o *ApplicationPutRequest) GetApplicationAccelerationOk() (*bool, bool)`

GetApplicationAccelerationOk returns a tuple with the ApplicationAcceleration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAcceleration

`func (o *ApplicationPutRequest) SetApplicationAcceleration(v bool)`

SetApplicationAcceleration sets ApplicationAcceleration field to given value.

### HasApplicationAcceleration

`func (o *ApplicationPutRequest) HasApplicationAcceleration() bool`

HasApplicationAcceleration returns a boolean if a field has been set.

### GetCaching

`func (o *ApplicationPutRequest) GetCaching() bool`

GetCaching returns the Caching field if non-nil, zero value otherwise.

### GetCachingOk

`func (o *ApplicationPutRequest) GetCachingOk() (*bool, bool)`

GetCachingOk returns a tuple with the Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaching

`func (o *ApplicationPutRequest) SetCaching(v bool)`

SetCaching sets Caching field to given value.

### HasCaching

`func (o *ApplicationPutRequest) HasCaching() bool`

HasCaching returns a boolean if a field has been set.

### GetDeviceDetection

`func (o *ApplicationPutRequest) GetDeviceDetection() bool`

GetDeviceDetection returns the DeviceDetection field if non-nil, zero value otherwise.

### GetDeviceDetectionOk

`func (o *ApplicationPutRequest) GetDeviceDetectionOk() (*bool, bool)`

GetDeviceDetectionOk returns a tuple with the DeviceDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDetection

`func (o *ApplicationPutRequest) SetDeviceDetection(v bool)`

SetDeviceDetection sets DeviceDetection field to given value.

### HasDeviceDetection

`func (o *ApplicationPutRequest) HasDeviceDetection() bool`

HasDeviceDetection returns a boolean if a field has been set.

### GetEdgeFirewall

`func (o *ApplicationPutRequest) GetEdgeFirewall() bool`

GetEdgeFirewall returns the EdgeFirewall field if non-nil, zero value otherwise.

### GetEdgeFirewallOk

`func (o *ApplicationPutRequest) GetEdgeFirewallOk() (*bool, bool)`

GetEdgeFirewallOk returns a tuple with the EdgeFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFirewall

`func (o *ApplicationPutRequest) SetEdgeFirewall(v bool)`

SetEdgeFirewall sets EdgeFirewall field to given value.

### HasEdgeFirewall

`func (o *ApplicationPutRequest) HasEdgeFirewall() bool`

HasEdgeFirewall returns a boolean if a field has been set.

### GetEdgeFunctions

`func (o *ApplicationPutRequest) GetEdgeFunctions() bool`

GetEdgeFunctions returns the EdgeFunctions field if non-nil, zero value otherwise.

### GetEdgeFunctionsOk

`func (o *ApplicationPutRequest) GetEdgeFunctionsOk() (*bool, bool)`

GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctions

`func (o *ApplicationPutRequest) SetEdgeFunctions(v bool)`

SetEdgeFunctions sets EdgeFunctions field to given value.

### HasEdgeFunctions

`func (o *ApplicationPutRequest) HasEdgeFunctions() bool`

HasEdgeFunctions returns a boolean if a field has been set.

### GetImageOptimization

`func (o *ApplicationPutRequest) GetImageOptimization() bool`

GetImageOptimization returns the ImageOptimization field if non-nil, zero value otherwise.

### GetImageOptimizationOk

`func (o *ApplicationPutRequest) GetImageOptimizationOk() (*bool, bool)`

GetImageOptimizationOk returns a tuple with the ImageOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOptimization

`func (o *ApplicationPutRequest) SetImageOptimization(v bool)`

SetImageOptimization sets ImageOptimization field to given value.

### HasImageOptimization

`func (o *ApplicationPutRequest) HasImageOptimization() bool`

HasImageOptimization returns a boolean if a field has been set.

### GetL2Caching

`func (o *ApplicationPutRequest) GetL2Caching() bool`

GetL2Caching returns the L2Caching field if non-nil, zero value otherwise.

### GetL2CachingOk

`func (o *ApplicationPutRequest) GetL2CachingOk() (*bool, bool)`

GetL2CachingOk returns a tuple with the L2Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2Caching

`func (o *ApplicationPutRequest) SetL2Caching(v bool)`

SetL2Caching sets L2Caching field to given value.

### HasL2Caching

`func (o *ApplicationPutRequest) HasL2Caching() bool`

HasL2Caching returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *ApplicationPutRequest) GetLoadBalancer() bool`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *ApplicationPutRequest) GetLoadBalancerOk() (*bool, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *ApplicationPutRequest) SetLoadBalancer(v bool)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *ApplicationPutRequest) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetRawLogs

`func (o *ApplicationPutRequest) GetRawLogs() bool`

GetRawLogs returns the RawLogs field if non-nil, zero value otherwise.

### GetRawLogsOk

`func (o *ApplicationPutRequest) GetRawLogsOk() (*bool, bool)`

GetRawLogsOk returns a tuple with the RawLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawLogs

`func (o *ApplicationPutRequest) SetRawLogs(v bool)`

SetRawLogs sets RawLogs field to given value.

### HasRawLogs

`func (o *ApplicationPutRequest) HasRawLogs() bool`

HasRawLogs returns a boolean if a field has been set.

### GetWebApplicationFirewall

`func (o *ApplicationPutRequest) GetWebApplicationFirewall() bool`

GetWebApplicationFirewall returns the WebApplicationFirewall field if non-nil, zero value otherwise.

### GetWebApplicationFirewallOk

`func (o *ApplicationPutRequest) GetWebApplicationFirewallOk() (*bool, bool)`

GetWebApplicationFirewallOk returns a tuple with the WebApplicationFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApplicationFirewall

`func (o *ApplicationPutRequest) SetWebApplicationFirewall(v bool)`

SetWebApplicationFirewall sets WebApplicationFirewall field to given value.

### HasWebApplicationFirewall

`func (o *ApplicationPutRequest) HasWebApplicationFirewall() bool`

HasWebApplicationFirewall returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


