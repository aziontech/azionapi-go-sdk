/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApplicationUpdateResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationUpdateResults{}

// ApplicationUpdateResults struct for ApplicationUpdateResults
type ApplicationUpdateResults struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	DeliveryProtocol string `json:"delivery_protocol"`
	HttpPort interface{} `json:"http_port"`
	HttpsPort interface{} `json:"https_port"`
	MinimumTlsVersion string `json:"minimum_tls_version"`
	Active bool `json:"active"`
	DebugRules bool `json:"debug_rules"`
	Http3 bool `json:"http3"`
	SupportedCiphers string `json:"supported_ciphers"`
	ApplicationAcceleration bool `json:"application_acceleration"`
	Caching bool `json:"caching"`
	DeviceDetection bool `json:"device_detection"`
	EdgeFirewall bool `json:"edge_firewall"`
	EdgeFunctions bool `json:"edge_functions"`
	ImageOptimization bool `json:"image_optimization"`
	L2Caching bool `json:"l2_caching"`
	LoadBalancer bool `json:"load_balancer"`
	RawLogs bool `json:"raw_logs"`
	WebApplicationFirewall bool `json:"web_application_firewall"`
	Websocket *bool `json:"websocket,omitempty"`
}

type _ApplicationUpdateResults ApplicationUpdateResults

// NewApplicationUpdateResults instantiates a new ApplicationUpdateResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationUpdateResults(id int64, name string, deliveryProtocol string, httpPort interface{}, httpsPort interface{}, minimumTlsVersion string, active bool, debugRules bool, http3 bool, supportedCiphers string, applicationAcceleration bool, caching bool, deviceDetection bool, edgeFirewall bool, edgeFunctions bool, imageOptimization bool, l2Caching bool, loadBalancer bool, rawLogs bool, webApplicationFirewall bool) *ApplicationUpdateResults {
	this := ApplicationUpdateResults{}
	this.Id = id
	this.Name = name
	this.DeliveryProtocol = deliveryProtocol
	this.HttpPort = httpPort
	this.HttpsPort = httpsPort
	this.MinimumTlsVersion = minimumTlsVersion
	this.Active = active
	this.DebugRules = debugRules
	this.Http3 = http3
	this.SupportedCiphers = supportedCiphers
	this.ApplicationAcceleration = applicationAcceleration
	this.Caching = caching
	this.DeviceDetection = deviceDetection
	this.EdgeFirewall = edgeFirewall
	this.EdgeFunctions = edgeFunctions
	this.ImageOptimization = imageOptimization
	this.L2Caching = l2Caching
	this.LoadBalancer = loadBalancer
	this.RawLogs = rawLogs
	this.WebApplicationFirewall = webApplicationFirewall
	return &this
}

// NewApplicationUpdateResultsWithDefaults instantiates a new ApplicationUpdateResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationUpdateResultsWithDefaults() *ApplicationUpdateResults {
	this := ApplicationUpdateResults{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationUpdateResults) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationUpdateResults) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ApplicationUpdateResults) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationUpdateResults) SetName(v string) {
	o.Name = v
}

// GetDeliveryProtocol returns the DeliveryProtocol field value
func (o *ApplicationUpdateResults) GetDeliveryProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryProtocol
}

// GetDeliveryProtocolOk returns a tuple with the DeliveryProtocol field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetDeliveryProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryProtocol, true
}

// SetDeliveryProtocol sets field value
func (o *ApplicationUpdateResults) SetDeliveryProtocol(v string) {
	o.DeliveryProtocol = v
}

// GetHttpPort returns the HttpPort field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApplicationUpdateResults) GetHttpPort() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.HttpPort
}

// GetHttpPortOk returns a tuple with the HttpPort field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationUpdateResults) GetHttpPortOk() (*interface{}, bool) {
	if o == nil || IsNil(o.HttpPort) {
		return nil, false
	}
	return &o.HttpPort, true
}

// SetHttpPort sets field value
func (o *ApplicationUpdateResults) SetHttpPort(v interface{}) {
	o.HttpPort = v
}

// GetHttpsPort returns the HttpsPort field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApplicationUpdateResults) GetHttpsPort() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationUpdateResults) GetHttpsPortOk() (*interface{}, bool) {
	if o == nil || IsNil(o.HttpsPort) {
		return nil, false
	}
	return &o.HttpsPort, true
}

// SetHttpsPort sets field value
func (o *ApplicationUpdateResults) SetHttpsPort(v interface{}) {
	o.HttpsPort = v
}

// GetMinimumTlsVersion returns the MinimumTlsVersion field value
func (o *ApplicationUpdateResults) GetMinimumTlsVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinimumTlsVersion
}

// GetMinimumTlsVersionOk returns a tuple with the MinimumTlsVersion field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetMinimumTlsVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumTlsVersion, true
}

// SetMinimumTlsVersion sets field value
func (o *ApplicationUpdateResults) SetMinimumTlsVersion(v string) {
	o.MinimumTlsVersion = v
}

// GetActive returns the Active field value
func (o *ApplicationUpdateResults) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *ApplicationUpdateResults) SetActive(v bool) {
	o.Active = v
}

// GetDebugRules returns the DebugRules field value
func (o *ApplicationUpdateResults) GetDebugRules() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DebugRules
}

// GetDebugRulesOk returns a tuple with the DebugRules field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetDebugRulesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DebugRules, true
}

// SetDebugRules sets field value
func (o *ApplicationUpdateResults) SetDebugRules(v bool) {
	o.DebugRules = v
}

// GetHttp3 returns the Http3 field value
func (o *ApplicationUpdateResults) GetHttp3() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Http3
}

// GetHttp3Ok returns a tuple with the Http3 field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetHttp3Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Http3, true
}

// SetHttp3 sets field value
func (o *ApplicationUpdateResults) SetHttp3(v bool) {
	o.Http3 = v
}

// GetSupportedCiphers returns the SupportedCiphers field value
func (o *ApplicationUpdateResults) GetSupportedCiphers() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportedCiphers
}

// GetSupportedCiphersOk returns a tuple with the SupportedCiphers field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetSupportedCiphersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedCiphers, true
}

// SetSupportedCiphers sets field value
func (o *ApplicationUpdateResults) SetSupportedCiphers(v string) {
	o.SupportedCiphers = v
}

// GetApplicationAcceleration returns the ApplicationAcceleration field value
func (o *ApplicationUpdateResults) GetApplicationAcceleration() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ApplicationAcceleration
}

// GetApplicationAccelerationOk returns a tuple with the ApplicationAcceleration field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetApplicationAccelerationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationAcceleration, true
}

// SetApplicationAcceleration sets field value
func (o *ApplicationUpdateResults) SetApplicationAcceleration(v bool) {
	o.ApplicationAcceleration = v
}

// GetCaching returns the Caching field value
func (o *ApplicationUpdateResults) GetCaching() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Caching
}

// GetCachingOk returns a tuple with the Caching field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetCachingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Caching, true
}

// SetCaching sets field value
func (o *ApplicationUpdateResults) SetCaching(v bool) {
	o.Caching = v
}

// GetDeviceDetection returns the DeviceDetection field value
func (o *ApplicationUpdateResults) GetDeviceDetection() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DeviceDetection
}

// GetDeviceDetectionOk returns a tuple with the DeviceDetection field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetDeviceDetectionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceDetection, true
}

// SetDeviceDetection sets field value
func (o *ApplicationUpdateResults) SetDeviceDetection(v bool) {
	o.DeviceDetection = v
}

// GetEdgeFirewall returns the EdgeFirewall field value
func (o *ApplicationUpdateResults) GetEdgeFirewall() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EdgeFirewall
}

// GetEdgeFirewallOk returns a tuple with the EdgeFirewall field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetEdgeFirewallOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFirewall, true
}

// SetEdgeFirewall sets field value
func (o *ApplicationUpdateResults) SetEdgeFirewall(v bool) {
	o.EdgeFirewall = v
}

// GetEdgeFunctions returns the EdgeFunctions field value
func (o *ApplicationUpdateResults) GetEdgeFunctions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EdgeFunctions
}

// GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetEdgeFunctionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFunctions, true
}

// SetEdgeFunctions sets field value
func (o *ApplicationUpdateResults) SetEdgeFunctions(v bool) {
	o.EdgeFunctions = v
}

// GetImageOptimization returns the ImageOptimization field value
func (o *ApplicationUpdateResults) GetImageOptimization() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ImageOptimization
}

// GetImageOptimizationOk returns a tuple with the ImageOptimization field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetImageOptimizationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageOptimization, true
}

// SetImageOptimization sets field value
func (o *ApplicationUpdateResults) SetImageOptimization(v bool) {
	o.ImageOptimization = v
}

// GetL2Caching returns the L2Caching field value
func (o *ApplicationUpdateResults) GetL2Caching() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.L2Caching
}

// GetL2CachingOk returns a tuple with the L2Caching field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetL2CachingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.L2Caching, true
}

// SetL2Caching sets field value
func (o *ApplicationUpdateResults) SetL2Caching(v bool) {
	o.L2Caching = v
}

// GetLoadBalancer returns the LoadBalancer field value
func (o *ApplicationUpdateResults) GetLoadBalancer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LoadBalancer
}

// GetLoadBalancerOk returns a tuple with the LoadBalancer field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetLoadBalancerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancer, true
}

// SetLoadBalancer sets field value
func (o *ApplicationUpdateResults) SetLoadBalancer(v bool) {
	o.LoadBalancer = v
}

// GetRawLogs returns the RawLogs field value
func (o *ApplicationUpdateResults) GetRawLogs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RawLogs
}

// GetRawLogsOk returns a tuple with the RawLogs field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetRawLogsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawLogs, true
}

// SetRawLogs sets field value
func (o *ApplicationUpdateResults) SetRawLogs(v bool) {
	o.RawLogs = v
}

// GetWebApplicationFirewall returns the WebApplicationFirewall field value
func (o *ApplicationUpdateResults) GetWebApplicationFirewall() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WebApplicationFirewall
}

// GetWebApplicationFirewallOk returns a tuple with the WebApplicationFirewall field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetWebApplicationFirewallOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebApplicationFirewall, true
}

// SetWebApplicationFirewall sets field value
func (o *ApplicationUpdateResults) SetWebApplicationFirewall(v bool) {
	o.WebApplicationFirewall = v
}

// GetWebsocket returns the Websocket field value if set, zero value otherwise.
func (o *ApplicationUpdateResults) GetWebsocket() bool {
	if o == nil || IsNil(o.Websocket) {
		var ret bool
		return ret
	}
	return *o.Websocket
}

// GetWebsocketOk returns a tuple with the Websocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetWebsocketOk() (*bool, bool) {
	if o == nil || IsNil(o.Websocket) {
		return nil, false
	}
	return o.Websocket, true
}

// HasWebsocket returns a boolean if a field has been set.
func (o *ApplicationUpdateResults) HasWebsocket() bool {
	if o != nil && !IsNil(o.Websocket) {
		return true
	}

	return false
}

// SetWebsocket gets a reference to the given bool and assigns it to the Websocket field.
func (o *ApplicationUpdateResults) SetWebsocket(v bool) {
	o.Websocket = &v
}

func (o ApplicationUpdateResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationUpdateResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["delivery_protocol"] = o.DeliveryProtocol
	if o.HttpPort != nil {
		toSerialize["http_port"] = o.HttpPort
	}
	if o.HttpsPort != nil {
		toSerialize["https_port"] = o.HttpsPort
	}
	toSerialize["minimum_tls_version"] = o.MinimumTlsVersion
	toSerialize["active"] = o.Active
	toSerialize["debug_rules"] = o.DebugRules
	toSerialize["http3"] = o.Http3
	toSerialize["supported_ciphers"] = o.SupportedCiphers
	toSerialize["application_acceleration"] = o.ApplicationAcceleration
	toSerialize["caching"] = o.Caching
	toSerialize["device_detection"] = o.DeviceDetection
	toSerialize["edge_firewall"] = o.EdgeFirewall
	toSerialize["edge_functions"] = o.EdgeFunctions
	toSerialize["image_optimization"] = o.ImageOptimization
	toSerialize["l2_caching"] = o.L2Caching
	toSerialize["load_balancer"] = o.LoadBalancer
	toSerialize["raw_logs"] = o.RawLogs
	toSerialize["web_application_firewall"] = o.WebApplicationFirewall
	if !IsNil(o.Websocket) {
		toSerialize["websocket"] = o.Websocket
	}
	return toSerialize, nil
}

func (o *ApplicationUpdateResults) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"delivery_protocol",
		"http_port",
		"https_port",
		"minimum_tls_version",
		"active",
		"debug_rules",
		"http3",
		"supported_ciphers",
		"application_acceleration",
		"caching",
		"device_detection",
		"edge_firewall",
		"edge_functions",
		"image_optimization",
		"l2_caching",
		"load_balancer",
		"raw_logs",
		"web_application_firewall",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApplicationUpdateResults := _ApplicationUpdateResults{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationUpdateResults)

	if err != nil {
		return err
	}

	*o = ApplicationUpdateResults(varApplicationUpdateResults)

	return err
}

type NullableApplicationUpdateResults struct {
	value *ApplicationUpdateResults
	isSet bool
}

func (v NullableApplicationUpdateResults) Get() *ApplicationUpdateResults {
	return v.value
}

func (v *NullableApplicationUpdateResults) Set(val *ApplicationUpdateResults) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationUpdateResults) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationUpdateResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationUpdateResults(val *ApplicationUpdateResults) *NullableApplicationUpdateResults {
	return &NullableApplicationUpdateResults{value: val, isSet: true}
}

func (v NullableApplicationUpdateResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationUpdateResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


