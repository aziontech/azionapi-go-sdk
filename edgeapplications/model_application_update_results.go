/*
Edge Application

## Welcome to the Azion API!  With the following APIs you can check, remove or update existing settings, besides creating new ones.  * * *  ## Overview  The Azion API is a RESTful API, based on HTTPS requests, that allows you to integrate your systems with our platform, simply, quickly, and securely.  Here you will find instructions on how our API works and what functionality is available.  This documentation is being constantly updated. Make sure you verify if there are any updates or changes before you perform any development in your application, even if you are familiar with our API.  * * *  Both HTTPS requests and responses must be in JavaScript Object Notation (JSON) format. All HTTPS requests and responses must follow these conventions.  **Base URL**  The base URL of the API, which must be used, is:  [**https://api.azionapi.net/_**](https://api.azionapi.net/)  **HTTP Methods**  Each HTTP method defines the type of operation that will be run.  | HTTP Method | CRUD | Whole Collection (e.g. /items) | Specific Item (e.g. /items/:item_id) | | --- | --- | --- | --- | | GET | Read | It retrieves the list of items in a Collection. | It retrieves a specific item in a Collection. | | POST | Create | It creates a new item in the Collection. | \\- | | PUT | Update/Replace | It replaces a whole Collection with a new one. | It replaces an item in a Collection with a new one. | | PATCH | Update/Modify | It partially updates a Collection. | It partially updates an item in a Collection | | DELETE | Delete | It deletes a whole Collection. | It deletes an item in a Collection. |  * * *  **Status Codes**  The HTTP return code defines the status – successful or not – after the requested operation is run.  | Status Code | Meaning | | --- | --- | | 200 OK | General Status for a successful operation. | | 201 CREATED | Successfully created a collection or item, by means of POST or PUT. | | 204 NO CONTENT | Successful operation, but without any content to be return to the Body. Generally used for DELETE or PUT operations. | | 207 MULTI-STATUS | A batch of operations were triggered by a single request, which resulted in different return codes when it was run, for some of the operations. The codes are displayed in the “status” field, in the body of the response, for each sub-batch of operations for whichever are applicable. | | 400 BAD REQUEST | Request error, such as invalid parameters, missing mandatory parameters or others. | | 401 UNAUTHORIZED | Error caused by a missing HTTP Authentication header. | | 403 FORBIDDEN | User does not have the permissions to run the requested operation. | | 404 NOT FOUND | The requested resource does not exist. | | 405 METHOD NOT ALLOWED | The requested method is not applicable with the URL. | | 406 NOT ACCEPTABLE | Accept header missing from the HTTP request or the header contains formatting or the version is not supported by the API. | | 409 CONFLICT | Conflict generated in running the request, such as a request to create an already existing record. | | 429 TOO MANY REQUESTS | The request was temporarily rejected, due to exceeding the limit on operations. Wait for the time defined in the X-Ratelimit-Reset header and try again. | | 500 INTERNAL SERVER ERROR | Unintentional error, due to an unidentified failure in the request process. | | \\--- |  | | **HTTP Headers** |  |  All requests must be generated with the HTTP header specifying the desired code format for responses and the API version used. The current version of the API is 3 and the format is application/json.  ``` Accept: application/json; version=3  ```  * * *  **Rate Limit**  The limit of operations that can be run via the API is defined by 3 HTTP response headers:  *   **X-ratelimit-limit:** number of operations allowed in one time-frame; *   **X-ratelimit-remaining:** number of operations still available in one time-frame; *   **X-ratelimit-reset:** is the date and time, in IOS 8601 format, which represents the point in the future when the time-frame will be closed and when the limits will, therefore, be reset.       Example of HTTP response headers and a request:  ``` Date: Thu, 02 Nov 2017 12:30:28 GMT X-ratelimit-remaining: 199 X-ratelimit-limit: 200 X-ratelimit-reset: 2017-11-02T12:31:28.675446  ```  In the example provided, 200 operations are allowed within a 1-minute time frame. Of those, there are 199 still available, because one has already been run. The total limit will be reset after 1 minute.  When the X-ratelimit-limit is reached, or in other words, when the X-ratelimit-remaining reaches zero, the API will give the error, HTTP 429 TOO MANY REQUESTS. If your application receives this error, you will need to wait until the time defined in X-ratelimit-reset has passed, to make another request.  *   **X-ratelimit-limit by product**       The *X-ratelimit limit* variations by product are the following:  *   **Real-Time Metrics:** 20 requests per minute. *   **Real-Time Purge:** 200 requests per minute; except for the Wildcard, which is 2000 a day.       > The rate-limit values are based on the client_id.  * * *  **Optional Parameters**  In this version, it is possible to include some optional parameters as part of the GET method, which can help and modify the form in which your data will be returned.  You can combine these parameters to get the result you want.  They are:  | Parameter | Description | Accepted values: | | --- | --- | --- | | order_by | Identifies which field the return should be sorted by. The default ordering is ascending. | Depends on the fields available from the endpoint structure | | sort | Defines which ordering to be used: ascending or descending. | asc  <br>  <br>desc | | page | Identifies which page should be returned, if the return is paginated. The default value is 1. | Page number. | | page_size | Identifies how many items should be returned per page. The default value is 10. | Desired number of items. |  * * *  **Request Exemple**  You can use one parameter, or a combination. See the example below, which uses all of them in the same request.  ``` GET /domains?order_by=name&page_size=20&sort=desc&page=3 Accept: application/json; version=3 Authorization: token 2909f3932069047f4736dc87e72baaddd19c9f75  ```  * * *  # Authentication  Authentication and authorization of operations via Azion API is done through Tokens.  The first step is to create the Token through authenticating a user registered in [Real-Time Manager](https://sso.azion.com/login).  * * *  ## Encoding Username and Password in Base64  Only token creation and cancelling operations are performed through Basic Authentication, that is, with a username and password. You can create and cancel the token through the API itself, but for that you need to encode your username and password in base64.  Base64 encoding can be done from the command line on a Unix terminal:  ``` $ echo 'user@domain:password'|base64 dXNlckBkb21haW46cGFzc3dvcmQK  ```  If you do not have a Unix terminal available, you can use the free online service at [https://www.base64encode.org/](https://www.base64encode.org/)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
)

// checks if the ApplicationUpdateResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationUpdateResults{}

// ApplicationUpdateResults struct for ApplicationUpdateResults
type ApplicationUpdateResults struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	DeliveryProtocol string `json:"delivery_protocol"`
	HttpPort interface{} `json:"http_port"`
	HttpsPort []int64 `json:"https_port"`
	MinimumTlsVersion string `json:"minimum_tls_version"`
	Active bool `json:"active"`
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

// NewApplicationUpdateResults instantiates a new ApplicationUpdateResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationUpdateResults(id int64, name string, deliveryProtocol string, httpPort interface{}, httpsPort []int64, minimumTlsVersion string, active bool, applicationAcceleration bool, caching bool, deviceDetection bool, edgeFirewall bool, edgeFunctions bool, imageOptimization bool, l2Caching bool, loadBalancer bool, rawLogs bool, webApplicationFirewall bool) *ApplicationUpdateResults {
	this := ApplicationUpdateResults{}
	this.Id = id
	this.Name = name
	this.DeliveryProtocol = deliveryProtocol
	this.HttpPort = httpPort
	this.HttpsPort = httpsPort
	this.MinimumTlsVersion = minimumTlsVersion
	this.Active = active
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
func (o *ApplicationUpdateResults) GetHttpPort() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.HttpPort
}

// GetHttpPortOk returns a tuple with the HttpPort field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetHttpPortOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpPort, true
}

// SetHttpPort sets field value
func (o *ApplicationUpdateResults) SetHttpPort(v interface{}) {
	o.HttpPort = v
}

// GetHttpsPort returns the HttpsPort field value
func (o *ApplicationUpdateResults) GetHttpsPort() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateResults) GetHttpsPortOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HttpsPort, true
}

// SetHttpsPort sets field value
func (o *ApplicationUpdateResults) SetHttpsPort(v []int64) {
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
	toSerialize["http_port"] = o.HttpPort
	toSerialize["https_port"] = o.HttpsPort
	toSerialize["minimum_tls_version"] = o.MinimumTlsVersion
	toSerialize["active"] = o.Active
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


