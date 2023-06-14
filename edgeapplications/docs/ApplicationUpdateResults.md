# ApplicationUpdateResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**DeliveryProtocol** | **string** |  | 
**HttpPort** | **interface{}** |  | 
**HttpsPort** | **interface{}** |  | 
**MinimumTlsVersion** | **string** |  | 
**Active** | **bool** |  | 
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

## Methods

### NewApplicationUpdateResults

`func NewApplicationUpdateResults(id int64, name string, deliveryProtocol string, httpPort interface{}, httpsPort interface{}, minimumTlsVersion string, active bool, applicationAcceleration bool, caching bool, deviceDetection bool, edgeFirewall bool, edgeFunctions bool, imageOptimization bool, l2Caching bool, loadBalancer bool, rawLogs bool, webApplicationFirewall bool, ) *ApplicationUpdateResults`

NewApplicationUpdateResults instantiates a new ApplicationUpdateResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationUpdateResultsWithDefaults

`func NewApplicationUpdateResultsWithDefaults() *ApplicationUpdateResults`

NewApplicationUpdateResultsWithDefaults instantiates a new ApplicationUpdateResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationUpdateResults) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationUpdateResults) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationUpdateResults) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationUpdateResults) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationUpdateResults) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationUpdateResults) SetName(v string)`

SetName sets Name field to given value.


### GetDeliveryProtocol

`func (o *ApplicationUpdateResults) GetDeliveryProtocol() string`

GetDeliveryProtocol returns the DeliveryProtocol field if non-nil, zero value otherwise.

### GetDeliveryProtocolOk

`func (o *ApplicationUpdateResults) GetDeliveryProtocolOk() (*string, bool)`

GetDeliveryProtocolOk returns a tuple with the DeliveryProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryProtocol

`func (o *ApplicationUpdateResults) SetDeliveryProtocol(v string)`

SetDeliveryProtocol sets DeliveryProtocol field to given value.


### GetHttpPort

`func (o *ApplicationUpdateResults) GetHttpPort() interface{}`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *ApplicationUpdateResults) GetHttpPortOk() (*interface{}, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *ApplicationUpdateResults) SetHttpPort(v interface{})`

SetHttpPort sets HttpPort field to given value.


### SetHttpPortNil

`func (o *ApplicationUpdateResults) SetHttpPortNil(b bool)`

 SetHttpPortNil sets the value for HttpPort to be an explicit nil

### UnsetHttpPort
`func (o *ApplicationUpdateResults) UnsetHttpPort()`

UnsetHttpPort ensures that no value is present for HttpPort, not even an explicit nil
### GetHttpsPort

`func (o *ApplicationUpdateResults) GetHttpsPort() interface{}`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *ApplicationUpdateResults) GetHttpsPortOk() (*interface{}, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *ApplicationUpdateResults) SetHttpsPort(v interface{})`

SetHttpsPort sets HttpsPort field to given value.


### SetHttpsPortNil

`func (o *ApplicationUpdateResults) SetHttpsPortNil(b bool)`

 SetHttpsPortNil sets the value for HttpsPort to be an explicit nil

### UnsetHttpsPort
`func (o *ApplicationUpdateResults) UnsetHttpsPort()`

UnsetHttpsPort ensures that no value is present for HttpsPort, not even an explicit nil
### GetMinimumTlsVersion

`func (o *ApplicationUpdateResults) GetMinimumTlsVersion() string`

GetMinimumTlsVersion returns the MinimumTlsVersion field if non-nil, zero value otherwise.

### GetMinimumTlsVersionOk

`func (o *ApplicationUpdateResults) GetMinimumTlsVersionOk() (*string, bool)`

GetMinimumTlsVersionOk returns a tuple with the MinimumTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTlsVersion

`func (o *ApplicationUpdateResults) SetMinimumTlsVersion(v string)`

SetMinimumTlsVersion sets MinimumTlsVersion field to given value.


### GetActive

`func (o *ApplicationUpdateResults) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationUpdateResults) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationUpdateResults) SetActive(v bool)`

SetActive sets Active field to given value.


### GetApplicationAcceleration

`func (o *ApplicationUpdateResults) GetApplicationAcceleration() bool`

GetApplicationAcceleration returns the ApplicationAcceleration field if non-nil, zero value otherwise.

### GetApplicationAccelerationOk

`func (o *ApplicationUpdateResults) GetApplicationAccelerationOk() (*bool, bool)`

GetApplicationAccelerationOk returns a tuple with the ApplicationAcceleration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAcceleration

`func (o *ApplicationUpdateResults) SetApplicationAcceleration(v bool)`

SetApplicationAcceleration sets ApplicationAcceleration field to given value.


### GetCaching

`func (o *ApplicationUpdateResults) GetCaching() bool`

GetCaching returns the Caching field if non-nil, zero value otherwise.

### GetCachingOk

`func (o *ApplicationUpdateResults) GetCachingOk() (*bool, bool)`

GetCachingOk returns a tuple with the Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaching

`func (o *ApplicationUpdateResults) SetCaching(v bool)`

SetCaching sets Caching field to given value.


### GetDeviceDetection

`func (o *ApplicationUpdateResults) GetDeviceDetection() bool`

GetDeviceDetection returns the DeviceDetection field if non-nil, zero value otherwise.

### GetDeviceDetectionOk

`func (o *ApplicationUpdateResults) GetDeviceDetectionOk() (*bool, bool)`

GetDeviceDetectionOk returns a tuple with the DeviceDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDetection

`func (o *ApplicationUpdateResults) SetDeviceDetection(v bool)`

SetDeviceDetection sets DeviceDetection field to given value.


### GetEdgeFirewall

`func (o *ApplicationUpdateResults) GetEdgeFirewall() bool`

GetEdgeFirewall returns the EdgeFirewall field if non-nil, zero value otherwise.

### GetEdgeFirewallOk

`func (o *ApplicationUpdateResults) GetEdgeFirewallOk() (*bool, bool)`

GetEdgeFirewallOk returns a tuple with the EdgeFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFirewall

`func (o *ApplicationUpdateResults) SetEdgeFirewall(v bool)`

SetEdgeFirewall sets EdgeFirewall field to given value.


### GetEdgeFunctions

`func (o *ApplicationUpdateResults) GetEdgeFunctions() bool`

GetEdgeFunctions returns the EdgeFunctions field if non-nil, zero value otherwise.

### GetEdgeFunctionsOk

`func (o *ApplicationUpdateResults) GetEdgeFunctionsOk() (*bool, bool)`

GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctions

`func (o *ApplicationUpdateResults) SetEdgeFunctions(v bool)`

SetEdgeFunctions sets EdgeFunctions field to given value.


### GetImageOptimization

`func (o *ApplicationUpdateResults) GetImageOptimization() bool`

GetImageOptimization returns the ImageOptimization field if non-nil, zero value otherwise.

### GetImageOptimizationOk

`func (o *ApplicationUpdateResults) GetImageOptimizationOk() (*bool, bool)`

GetImageOptimizationOk returns a tuple with the ImageOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOptimization

`func (o *ApplicationUpdateResults) SetImageOptimization(v bool)`

SetImageOptimization sets ImageOptimization field to given value.


### GetL2Caching

`func (o *ApplicationUpdateResults) GetL2Caching() bool`

GetL2Caching returns the L2Caching field if non-nil, zero value otherwise.

### GetL2CachingOk

`func (o *ApplicationUpdateResults) GetL2CachingOk() (*bool, bool)`

GetL2CachingOk returns a tuple with the L2Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2Caching

`func (o *ApplicationUpdateResults) SetL2Caching(v bool)`

SetL2Caching sets L2Caching field to given value.


### GetLoadBalancer

`func (o *ApplicationUpdateResults) GetLoadBalancer() bool`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *ApplicationUpdateResults) GetLoadBalancerOk() (*bool, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *ApplicationUpdateResults) SetLoadBalancer(v bool)`

SetLoadBalancer sets LoadBalancer field to given value.


### GetRawLogs

`func (o *ApplicationUpdateResults) GetRawLogs() bool`

GetRawLogs returns the RawLogs field if non-nil, zero value otherwise.

### GetRawLogsOk

`func (o *ApplicationUpdateResults) GetRawLogsOk() (*bool, bool)`

GetRawLogsOk returns a tuple with the RawLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawLogs

`func (o *ApplicationUpdateResults) SetRawLogs(v bool)`

SetRawLogs sets RawLogs field to given value.


### GetWebApplicationFirewall

`func (o *ApplicationUpdateResults) GetWebApplicationFirewall() bool`

GetWebApplicationFirewall returns the WebApplicationFirewall field if non-nil, zero value otherwise.

### GetWebApplicationFirewallOk

`func (o *ApplicationUpdateResults) GetWebApplicationFirewallOk() (*bool, bool)`

GetWebApplicationFirewallOk returns a tuple with the WebApplicationFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApplicationFirewall

`func (o *ApplicationUpdateResults) SetWebApplicationFirewall(v bool)`

SetWebApplicationFirewall sets WebApplicationFirewall field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


