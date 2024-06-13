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

// checks if the ApplicationResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationResults{}

// ApplicationResults struct for ApplicationResults
type ApplicationResults struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Active bool `json:"active"`
	DebugRules bool `json:"debug_rules"`
	Http3 bool `json:"http3"`
	SupportedCiphers string `json:"supported_ciphers"`
	DeliveryProtocol string `json:"delivery_protocol"`
	HttpPort interface{} `json:"http_port"`
	HttpsPort interface{} `json:"https_port"`
	MinimumTlsVersion string `json:"minimum_tls_version"`
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
}

type _ApplicationResults ApplicationResults

// NewApplicationResults instantiates a new ApplicationResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResults(id int64, name string, active bool, debugRules bool, http3 bool, supportedCiphers string, deliveryProtocol string, httpPort interface{}, httpsPort interface{}, minimumTlsVersion string, applicationAcceleration bool, caching bool, deviceDetection bool, edgeFirewall bool, edgeFunctions bool, imageOptimization bool, l2Caching bool, loadBalancer bool, rawLogs bool, webApplicationFirewall bool) *ApplicationResults {
	this := ApplicationResults{}
	this.Id = id
	this.Name = name
	this.Active = active
	this.DebugRules = debugRules
	this.Http3 = http3
	this.SupportedCiphers = supportedCiphers
	this.DeliveryProtocol = deliveryProtocol
	this.HttpPort = httpPort
	this.HttpsPort = httpsPort
	this.MinimumTlsVersion = minimumTlsVersion
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

// NewApplicationResultsWithDefaults instantiates a new ApplicationResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResultsWithDefaults() *ApplicationResults {
	this := ApplicationResults{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationResults) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationResults) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ApplicationResults) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationResults) SetName(v string) {
	o.Name = v
}

// GetActive returns the Active field value
func (o *ApplicationResults) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *ApplicationResults) SetActive(v bool) {
	o.Active = v
}

// GetDebugRules returns the DebugRules field value
func (o *ApplicationResults) GetDebugRules() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DebugRules
}

// GetDebugRulesOk returns a tuple with the DebugRules field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetDebugRulesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DebugRules, true
}

// SetDebugRules sets field value
func (o *ApplicationResults) SetDebugRules(v bool) {
	o.DebugRules = v
}

// GetHttp3 returns the Http3 field value
func (o *ApplicationResults) GetHttp3() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Http3
}

// GetHttp3Ok returns a tuple with the Http3 field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetHttp3Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Http3, true
}

// SetHttp3 sets field value
func (o *ApplicationResults) SetHttp3(v bool) {
	o.Http3 = v
}

// GetSupportedCiphers returns the SupportedCiphers field value
func (o *ApplicationResults) GetSupportedCiphers() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportedCiphers
}

// GetSupportedCiphersOk returns a tuple with the SupportedCiphers field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetSupportedCiphersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedCiphers, true
}

// SetSupportedCiphers sets field value
func (o *ApplicationResults) SetSupportedCiphers(v string) {
	o.SupportedCiphers = v
}

// GetDeliveryProtocol returns the DeliveryProtocol field value
func (o *ApplicationResults) GetDeliveryProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryProtocol
}

// GetDeliveryProtocolOk returns a tuple with the DeliveryProtocol field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetDeliveryProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryProtocol, true
}

// SetDeliveryProtocol sets field value
func (o *ApplicationResults) SetDeliveryProtocol(v string) {
	o.DeliveryProtocol = v
}

// GetHttpPort returns the HttpPort field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApplicationResults) GetHttpPort() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.HttpPort
}

// GetHttpPortOk returns a tuple with the HttpPort field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationResults) GetHttpPortOk() (*interface{}, bool) {
	if o == nil || IsNil(o.HttpPort) {
		return nil, false
	}
	return &o.HttpPort, true
}

// SetHttpPort sets field value
func (o *ApplicationResults) SetHttpPort(v interface{}) {
	o.HttpPort = v
}

// GetHttpsPort returns the HttpsPort field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApplicationResults) GetHttpsPort() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationResults) GetHttpsPortOk() (*interface{}, bool) {
	if o == nil || IsNil(o.HttpsPort) {
		return nil, false
	}
	return &o.HttpsPort, true
}

// SetHttpsPort sets field value
func (o *ApplicationResults) SetHttpsPort(v interface{}) {
	o.HttpsPort = v
}

// GetMinimumTlsVersion returns the MinimumTlsVersion field value
func (o *ApplicationResults) GetMinimumTlsVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinimumTlsVersion
}

// GetMinimumTlsVersionOk returns a tuple with the MinimumTlsVersion field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetMinimumTlsVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimumTlsVersion, true
}

// SetMinimumTlsVersion sets field value
func (o *ApplicationResults) SetMinimumTlsVersion(v string) {
	o.MinimumTlsVersion = v
}

// GetApplicationAcceleration returns the ApplicationAcceleration field value
func (o *ApplicationResults) GetApplicationAcceleration() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ApplicationAcceleration
}

// GetApplicationAccelerationOk returns a tuple with the ApplicationAcceleration field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetApplicationAccelerationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationAcceleration, true
}

// SetApplicationAcceleration sets field value
func (o *ApplicationResults) SetApplicationAcceleration(v bool) {
	o.ApplicationAcceleration = v
}

// GetCaching returns the Caching field value
func (o *ApplicationResults) GetCaching() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Caching
}

// GetCachingOk returns a tuple with the Caching field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetCachingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Caching, true
}

// SetCaching sets field value
func (o *ApplicationResults) SetCaching(v bool) {
	o.Caching = v
}

// GetDeviceDetection returns the DeviceDetection field value
func (o *ApplicationResults) GetDeviceDetection() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DeviceDetection
}

// GetDeviceDetectionOk returns a tuple with the DeviceDetection field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetDeviceDetectionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceDetection, true
}

// SetDeviceDetection sets field value
func (o *ApplicationResults) SetDeviceDetection(v bool) {
	o.DeviceDetection = v
}

// GetEdgeFirewall returns the EdgeFirewall field value
func (o *ApplicationResults) GetEdgeFirewall() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EdgeFirewall
}

// GetEdgeFirewallOk returns a tuple with the EdgeFirewall field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetEdgeFirewallOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFirewall, true
}

// SetEdgeFirewall sets field value
func (o *ApplicationResults) SetEdgeFirewall(v bool) {
	o.EdgeFirewall = v
}

// GetEdgeFunctions returns the EdgeFunctions field value
func (o *ApplicationResults) GetEdgeFunctions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EdgeFunctions
}

// GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetEdgeFunctionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFunctions, true
}

// SetEdgeFunctions sets field value
func (o *ApplicationResults) SetEdgeFunctions(v bool) {
	o.EdgeFunctions = v
}

// GetImageOptimization returns the ImageOptimization field value
func (o *ApplicationResults) GetImageOptimization() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ImageOptimization
}

// GetImageOptimizationOk returns a tuple with the ImageOptimization field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetImageOptimizationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageOptimization, true
}

// SetImageOptimization sets field value
func (o *ApplicationResults) SetImageOptimization(v bool) {
	o.ImageOptimization = v
}

// GetL2Caching returns the L2Caching field value
func (o *ApplicationResults) GetL2Caching() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.L2Caching
}

// GetL2CachingOk returns a tuple with the L2Caching field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetL2CachingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.L2Caching, true
}

// SetL2Caching sets field value
func (o *ApplicationResults) SetL2Caching(v bool) {
	o.L2Caching = v
}

// GetLoadBalancer returns the LoadBalancer field value
func (o *ApplicationResults) GetLoadBalancer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LoadBalancer
}

// GetLoadBalancerOk returns a tuple with the LoadBalancer field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetLoadBalancerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancer, true
}

// SetLoadBalancer sets field value
func (o *ApplicationResults) SetLoadBalancer(v bool) {
	o.LoadBalancer = v
}

// GetRawLogs returns the RawLogs field value
func (o *ApplicationResults) GetRawLogs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RawLogs
}

// GetRawLogsOk returns a tuple with the RawLogs field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetRawLogsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawLogs, true
}

// SetRawLogs sets field value
func (o *ApplicationResults) SetRawLogs(v bool) {
	o.RawLogs = v
}

// GetWebApplicationFirewall returns the WebApplicationFirewall field value
func (o *ApplicationResults) GetWebApplicationFirewall() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WebApplicationFirewall
}

// GetWebApplicationFirewallOk returns a tuple with the WebApplicationFirewall field value
// and a boolean to check if the value has been set.
func (o *ApplicationResults) GetWebApplicationFirewallOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebApplicationFirewall, true
}

// SetWebApplicationFirewall sets field value
func (o *ApplicationResults) SetWebApplicationFirewall(v bool) {
	o.WebApplicationFirewall = v
}

func (o ApplicationResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["active"] = o.Active
	toSerialize["debug_rules"] = o.DebugRules
	toSerialize["http3"] = o.Http3
	toSerialize["supported_ciphers"] = o.SupportedCiphers
	toSerialize["delivery_protocol"] = o.DeliveryProtocol
	if o.HttpPort != nil {
		toSerialize["http_port"] = o.HttpPort
	}
	if o.HttpsPort != nil {
		toSerialize["https_port"] = o.HttpsPort
	}
	toSerialize["minimum_tls_version"] = o.MinimumTlsVersion
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
	return toSerialize, nil
}

func (o *ApplicationResults) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"active",
		"debug_rules",
		"http3",
		"supported_ciphers",
		"delivery_protocol",
		"http_port",
		"https_port",
		"minimum_tls_version",
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

	varApplicationResults := _ApplicationResults{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationResults)

	if err != nil {
		return err
	}

	*o = ApplicationResults(varApplicationResults)

	return err
}

type NullableApplicationResults struct {
	value *ApplicationResults
	isSet bool
}

func (v NullableApplicationResults) Get() *ApplicationResults {
	return v.value
}

func (v *NullableApplicationResults) Set(val *ApplicationResults) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResults) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResults(val *ApplicationResults) *NullableApplicationResults {
	return &NullableApplicationResults{value: val, isSet: true}
}

func (v NullableApplicationResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


