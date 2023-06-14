# ApplicationResultsCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Active** | **bool** |  | 
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
**LoadBalancer** | **bool** |  | 
**RawLogs** | **bool** |  | 
**WebApplicationFirewall** | **bool** |  | 

## Methods

### NewApplicationResultsCreate

`func NewApplicationResultsCreate(id int64, name string, active bool, deliveryProtocol string, httpPort interface{}, httpsPort interface{}, minimumTlsVersion string, applicationAcceleration bool, caching bool, deviceDetection bool, edgeFirewall bool, edgeFunctions bool, imageOptimization bool, loadBalancer bool, rawLogs bool, webApplicationFirewall bool, ) *ApplicationResultsCreate`

NewApplicationResultsCreate instantiates a new ApplicationResultsCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResultsCreateWithDefaults

`func NewApplicationResultsCreateWithDefaults() *ApplicationResultsCreate`

NewApplicationResultsCreateWithDefaults instantiates a new ApplicationResultsCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationResultsCreate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResultsCreate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResultsCreate) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationResultsCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationResultsCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationResultsCreate) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *ApplicationResultsCreate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationResultsCreate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationResultsCreate) SetActive(v bool)`

SetActive sets Active field to given value.


### GetDeliveryProtocol

`func (o *ApplicationResultsCreate) GetDeliveryProtocol() string`

GetDeliveryProtocol returns the DeliveryProtocol field if non-nil, zero value otherwise.

### GetDeliveryProtocolOk

`func (o *ApplicationResultsCreate) GetDeliveryProtocolOk() (*string, bool)`

GetDeliveryProtocolOk returns a tuple with the DeliveryProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryProtocol

`func (o *ApplicationResultsCreate) SetDeliveryProtocol(v string)`

SetDeliveryProtocol sets DeliveryProtocol field to given value.


### GetHttpPort

`func (o *ApplicationResultsCreate) GetHttpPort() interface{}`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *ApplicationResultsCreate) GetHttpPortOk() (*interface{}, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *ApplicationResultsCreate) SetHttpPort(v interface{})`

SetHttpPort sets HttpPort field to given value.


### SetHttpPortNil

`func (o *ApplicationResultsCreate) SetHttpPortNil(b bool)`

 SetHttpPortNil sets the value for HttpPort to be an explicit nil

### UnsetHttpPort
`func (o *ApplicationResultsCreate) UnsetHttpPort()`

UnsetHttpPort ensures that no value is present for HttpPort, not even an explicit nil
### GetHttpsPort

`func (o *ApplicationResultsCreate) GetHttpsPort() interface{}`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *ApplicationResultsCreate) GetHttpsPortOk() (*interface{}, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *ApplicationResultsCreate) SetHttpsPort(v interface{})`

SetHttpsPort sets HttpsPort field to given value.


### SetHttpsPortNil

`func (o *ApplicationResultsCreate) SetHttpsPortNil(b bool)`

 SetHttpsPortNil sets the value for HttpsPort to be an explicit nil

### UnsetHttpsPort
`func (o *ApplicationResultsCreate) UnsetHttpsPort()`

UnsetHttpsPort ensures that no value is present for HttpsPort, not even an explicit nil
### GetMinimumTlsVersion

`func (o *ApplicationResultsCreate) GetMinimumTlsVersion() string`

GetMinimumTlsVersion returns the MinimumTlsVersion field if non-nil, zero value otherwise.

### GetMinimumTlsVersionOk

`func (o *ApplicationResultsCreate) GetMinimumTlsVersionOk() (*string, bool)`

GetMinimumTlsVersionOk returns a tuple with the MinimumTlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumTlsVersion

`func (o *ApplicationResultsCreate) SetMinimumTlsVersion(v string)`

SetMinimumTlsVersion sets MinimumTlsVersion field to given value.


### GetApplicationAcceleration

`func (o *ApplicationResultsCreate) GetApplicationAcceleration() bool`

GetApplicationAcceleration returns the ApplicationAcceleration field if non-nil, zero value otherwise.

### GetApplicationAccelerationOk

`func (o *ApplicationResultsCreate) GetApplicationAccelerationOk() (*bool, bool)`

GetApplicationAccelerationOk returns a tuple with the ApplicationAcceleration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationAcceleration

`func (o *ApplicationResultsCreate) SetApplicationAcceleration(v bool)`

SetApplicationAcceleration sets ApplicationAcceleration field to given value.


### GetCaching

`func (o *ApplicationResultsCreate) GetCaching() bool`

GetCaching returns the Caching field if non-nil, zero value otherwise.

### GetCachingOk

`func (o *ApplicationResultsCreate) GetCachingOk() (*bool, bool)`

GetCachingOk returns a tuple with the Caching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaching

`func (o *ApplicationResultsCreate) SetCaching(v bool)`

SetCaching sets Caching field to given value.


### GetDeviceDetection

`func (o *ApplicationResultsCreate) GetDeviceDetection() bool`

GetDeviceDetection returns the DeviceDetection field if non-nil, zero value otherwise.

### GetDeviceDetectionOk

`func (o *ApplicationResultsCreate) GetDeviceDetectionOk() (*bool, bool)`

GetDeviceDetectionOk returns a tuple with the DeviceDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDetection

`func (o *ApplicationResultsCreate) SetDeviceDetection(v bool)`

SetDeviceDetection sets DeviceDetection field to given value.


### GetEdgeFirewall

`func (o *ApplicationResultsCreate) GetEdgeFirewall() bool`

GetEdgeFirewall returns the EdgeFirewall field if non-nil, zero value otherwise.

### GetEdgeFirewallOk

`func (o *ApplicationResultsCreate) GetEdgeFirewallOk() (*bool, bool)`

GetEdgeFirewallOk returns a tuple with the EdgeFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFirewall

`func (o *ApplicationResultsCreate) SetEdgeFirewall(v bool)`

SetEdgeFirewall sets EdgeFirewall field to given value.


### GetEdgeFunctions

`func (o *ApplicationResultsCreate) GetEdgeFunctions() bool`

GetEdgeFunctions returns the EdgeFunctions field if non-nil, zero value otherwise.

### GetEdgeFunctionsOk

`func (o *ApplicationResultsCreate) GetEdgeFunctionsOk() (*bool, bool)`

GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeFunctions

`func (o *ApplicationResultsCreate) SetEdgeFunctions(v bool)`

SetEdgeFunctions sets EdgeFunctions field to given value.


### GetImageOptimization

`func (o *ApplicationResultsCreate) GetImageOptimization() bool`

GetImageOptimization returns the ImageOptimization field if non-nil, zero value otherwise.

### GetImageOptimizationOk

`func (o *ApplicationResultsCreate) GetImageOptimizationOk() (*bool, bool)`

GetImageOptimizationOk returns a tuple with the ImageOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOptimization

`func (o *ApplicationResultsCreate) SetImageOptimization(v bool)`

SetImageOptimization sets ImageOptimization field to given value.


### GetLoadBalancer

`func (o *ApplicationResultsCreate) GetLoadBalancer() bool`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *ApplicationResultsCreate) GetLoadBalancerOk() (*bool, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *ApplicationResultsCreate) SetLoadBalancer(v bool)`

SetLoadBalancer sets LoadBalancer field to given value.


### GetRawLogs

`func (o *ApplicationResultsCreate) GetRawLogs() bool`

GetRawLogs returns the RawLogs field if non-nil, zero value otherwise.

### GetRawLogsOk

`func (o *ApplicationResultsCreate) GetRawLogsOk() (*bool, bool)`

GetRawLogsOk returns a tuple with the RawLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawLogs

`func (o *ApplicationResultsCreate) SetRawLogs(v bool)`

SetRawLogs sets RawLogs field to given value.


### GetWebApplicationFirewall

`func (o *ApplicationResultsCreate) GetWebApplicationFirewall() bool`

GetWebApplicationFirewall returns the WebApplicationFirewall field if non-nil, zero value otherwise.

### GetWebApplicationFirewallOk

`func (o *ApplicationResultsCreate) GetWebApplicationFirewallOk() (*bool, bool)`

GetWebApplicationFirewallOk returns a tuple with the WebApplicationFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApplicationFirewall

`func (o *ApplicationResultsCreate) SetWebApplicationFirewall(v bool)`

SetWebApplicationFirewall sets WebApplicationFirewall field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


