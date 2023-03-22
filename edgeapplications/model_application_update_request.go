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

// checks if the ApplicationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationUpdateRequest{}

// ApplicationUpdateRequest struct for ApplicationUpdateRequest
type ApplicationUpdateRequest struct {
	Name *string `json:"name,omitempty"`
	DeliveryProtocol *string `json:"delivery_protocol,omitempty"`
	HttpPort *interface{} `json:"http_port,omitempty"`
	HttpsPort *interface{} `json:"https_port,omitempty"`
	MinimumTlsVersion *string `json:"minimum_tls_version,omitempty"`
	Active *bool `json:"active,omitempty"`
	ApplicationAcceleration *bool `json:"application_acceleration,omitempty"`
	Caching *bool `json:"caching,omitempty"`
	DeviceDetection *bool `json:"device_detection,omitempty"`
	EdgeFirewall *bool `json:"edge_firewall,omitempty"`
	EdgeFunctions *bool `json:"edge_functions,omitempty"`
	ImageOptimization *bool `json:"image_optimization,omitempty"`
	L2Caching *bool `json:"l2_caching,omitempty"`
	LoadBalancer *bool `json:"load_balancer,omitempty"`
	RawLogs *bool `json:"raw_logs,omitempty"`
	WebApplicationFirewall *bool `json:"web_application_firewall,omitempty"`
}

// NewApplicationUpdateRequest instantiates a new ApplicationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationUpdateRequest() *ApplicationUpdateRequest {
	this := ApplicationUpdateRequest{}
	return &this
}

// NewApplicationUpdateRequestWithDefaults instantiates a new ApplicationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationUpdateRequestWithDefaults() *ApplicationUpdateRequest {
	this := ApplicationUpdateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetDeliveryProtocol returns the DeliveryProtocol field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetDeliveryProtocol() string {
	if o == nil || IsNil(o.DeliveryProtocol) {
		var ret string
		return ret
	}
	return *o.DeliveryProtocol
}

// GetDeliveryProtocolOk returns a tuple with the DeliveryProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetDeliveryProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryProtocol) {
		return nil, false
	}
	return o.DeliveryProtocol, true
}

// HasDeliveryProtocol returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasDeliveryProtocol() bool {
	if o != nil && !IsNil(o.DeliveryProtocol) {
		return true
	}

	return false
}

// SetDeliveryProtocol gets a reference to the given string and assigns it to the DeliveryProtocol field.
func (o *ApplicationUpdateRequest) SetDeliveryProtocol(v string) {
	o.DeliveryProtocol = &v
}

// GetHttpPort returns the HttpPort field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetHttpPort() interface{} {
	if o == nil || IsNil(o.HttpPort) {
		var ret interface{}
		return ret
	}
	return *o.HttpPort
}

// GetHttpPortOk returns a tuple with the HttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetHttpPortOk() (*interface{}, bool) {
	if o == nil || IsNil(o.HttpPort) {
		return nil, false
	}
	return o.HttpPort, true
}

// HasHttpPort returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasHttpPort() bool {
	if o != nil && !IsNil(o.HttpPort) {
		return true
	}

	return false
}

// SetHttpPort gets a reference to the given interface{} and assigns it to the HttpPort field.
func (o *ApplicationUpdateRequest) SetHttpPort(v interface{}) {
	o.HttpPort = &v
}

// GetHttpsPort returns the HttpsPort field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetHttpsPort() interface{} {
	if o == nil || IsNil(o.HttpsPort) {
		var ret interface{}
		return ret
	}
	return *o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetHttpsPortOk() (*interface{}, bool) {
	if o == nil || IsNil(o.HttpsPort) {
		return nil, false
	}
	return o.HttpsPort, true
}

// HasHttpsPort returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasHttpsPort() bool {
	if o != nil && !IsNil(o.HttpsPort) {
		return true
	}

	return false
}

// SetHttpsPort gets a reference to the given interface{} and assigns it to the HttpsPort field.
func (o *ApplicationUpdateRequest) SetHttpsPort(v interface{}) {
	o.HttpsPort = &v
}

// GetMinimumTlsVersion returns the MinimumTlsVersion field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetMinimumTlsVersion() string {
	if o == nil || IsNil(o.MinimumTlsVersion) {
		var ret string
		return ret
	}
	return *o.MinimumTlsVersion
}

// GetMinimumTlsVersionOk returns a tuple with the MinimumTlsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetMinimumTlsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumTlsVersion) {
		return nil, false
	}
	return o.MinimumTlsVersion, true
}

// HasMinimumTlsVersion returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasMinimumTlsVersion() bool {
	if o != nil && !IsNil(o.MinimumTlsVersion) {
		return true
	}

	return false
}

// SetMinimumTlsVersion gets a reference to the given string and assigns it to the MinimumTlsVersion field.
func (o *ApplicationUpdateRequest) SetMinimumTlsVersion(v string) {
	o.MinimumTlsVersion = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApplicationUpdateRequest) SetActive(v bool) {
	o.Active = &v
}

// GetApplicationAcceleration returns the ApplicationAcceleration field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetApplicationAcceleration() bool {
	if o == nil || IsNil(o.ApplicationAcceleration) {
		var ret bool
		return ret
	}
	return *o.ApplicationAcceleration
}

// GetApplicationAccelerationOk returns a tuple with the ApplicationAcceleration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetApplicationAccelerationOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplicationAcceleration) {
		return nil, false
	}
	return o.ApplicationAcceleration, true
}

// HasApplicationAcceleration returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasApplicationAcceleration() bool {
	if o != nil && !IsNil(o.ApplicationAcceleration) {
		return true
	}

	return false
}

// SetApplicationAcceleration gets a reference to the given bool and assigns it to the ApplicationAcceleration field.
func (o *ApplicationUpdateRequest) SetApplicationAcceleration(v bool) {
	o.ApplicationAcceleration = &v
}

// GetCaching returns the Caching field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetCaching() bool {
	if o == nil || IsNil(o.Caching) {
		var ret bool
		return ret
	}
	return *o.Caching
}

// GetCachingOk returns a tuple with the Caching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetCachingOk() (*bool, bool) {
	if o == nil || IsNil(o.Caching) {
		return nil, false
	}
	return o.Caching, true
}

// HasCaching returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasCaching() bool {
	if o != nil && !IsNil(o.Caching) {
		return true
	}

	return false
}

// SetCaching gets a reference to the given bool and assigns it to the Caching field.
func (o *ApplicationUpdateRequest) SetCaching(v bool) {
	o.Caching = &v
}

// GetDeviceDetection returns the DeviceDetection field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetDeviceDetection() bool {
	if o == nil || IsNil(o.DeviceDetection) {
		var ret bool
		return ret
	}
	return *o.DeviceDetection
}

// GetDeviceDetectionOk returns a tuple with the DeviceDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetDeviceDetectionOk() (*bool, bool) {
	if o == nil || IsNil(o.DeviceDetection) {
		return nil, false
	}
	return o.DeviceDetection, true
}

// HasDeviceDetection returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasDeviceDetection() bool {
	if o != nil && !IsNil(o.DeviceDetection) {
		return true
	}

	return false
}

// SetDeviceDetection gets a reference to the given bool and assigns it to the DeviceDetection field.
func (o *ApplicationUpdateRequest) SetDeviceDetection(v bool) {
	o.DeviceDetection = &v
}

// GetEdgeFirewall returns the EdgeFirewall field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetEdgeFirewall() bool {
	if o == nil || IsNil(o.EdgeFirewall) {
		var ret bool
		return ret
	}
	return *o.EdgeFirewall
}

// GetEdgeFirewallOk returns a tuple with the EdgeFirewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetEdgeFirewallOk() (*bool, bool) {
	if o == nil || IsNil(o.EdgeFirewall) {
		return nil, false
	}
	return o.EdgeFirewall, true
}

// HasEdgeFirewall returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasEdgeFirewall() bool {
	if o != nil && !IsNil(o.EdgeFirewall) {
		return true
	}

	return false
}

// SetEdgeFirewall gets a reference to the given bool and assigns it to the EdgeFirewall field.
func (o *ApplicationUpdateRequest) SetEdgeFirewall(v bool) {
	o.EdgeFirewall = &v
}

// GetEdgeFunctions returns the EdgeFunctions field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetEdgeFunctions() bool {
	if o == nil || IsNil(o.EdgeFunctions) {
		var ret bool
		return ret
	}
	return *o.EdgeFunctions
}

// GetEdgeFunctionsOk returns a tuple with the EdgeFunctions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetEdgeFunctionsOk() (*bool, bool) {
	if o == nil || IsNil(o.EdgeFunctions) {
		return nil, false
	}
	return o.EdgeFunctions, true
}

// HasEdgeFunctions returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasEdgeFunctions() bool {
	if o != nil && !IsNil(o.EdgeFunctions) {
		return true
	}

	return false
}

// SetEdgeFunctions gets a reference to the given bool and assigns it to the EdgeFunctions field.
func (o *ApplicationUpdateRequest) SetEdgeFunctions(v bool) {
	o.EdgeFunctions = &v
}

// GetImageOptimization returns the ImageOptimization field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetImageOptimization() bool {
	if o == nil || IsNil(o.ImageOptimization) {
		var ret bool
		return ret
	}
	return *o.ImageOptimization
}

// GetImageOptimizationOk returns a tuple with the ImageOptimization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetImageOptimizationOk() (*bool, bool) {
	if o == nil || IsNil(o.ImageOptimization) {
		return nil, false
	}
	return o.ImageOptimization, true
}

// HasImageOptimization returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasImageOptimization() bool {
	if o != nil && !IsNil(o.ImageOptimization) {
		return true
	}

	return false
}

// SetImageOptimization gets a reference to the given bool and assigns it to the ImageOptimization field.
func (o *ApplicationUpdateRequest) SetImageOptimization(v bool) {
	o.ImageOptimization = &v
}

// GetL2Caching returns the L2Caching field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetL2Caching() bool {
	if o == nil || IsNil(o.L2Caching) {
		var ret bool
		return ret
	}
	return *o.L2Caching
}

// GetL2CachingOk returns a tuple with the L2Caching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetL2CachingOk() (*bool, bool) {
	if o == nil || IsNil(o.L2Caching) {
		return nil, false
	}
	return o.L2Caching, true
}

// HasL2Caching returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasL2Caching() bool {
	if o != nil && !IsNil(o.L2Caching) {
		return true
	}

	return false
}

// SetL2Caching gets a reference to the given bool and assigns it to the L2Caching field.
func (o *ApplicationUpdateRequest) SetL2Caching(v bool) {
	o.L2Caching = &v
}

// GetLoadBalancer returns the LoadBalancer field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetLoadBalancer() bool {
	if o == nil || IsNil(o.LoadBalancer) {
		var ret bool
		return ret
	}
	return *o.LoadBalancer
}

// GetLoadBalancerOk returns a tuple with the LoadBalancer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetLoadBalancerOk() (*bool, bool) {
	if o == nil || IsNil(o.LoadBalancer) {
		return nil, false
	}
	return o.LoadBalancer, true
}

// HasLoadBalancer returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasLoadBalancer() bool {
	if o != nil && !IsNil(o.LoadBalancer) {
		return true
	}

	return false
}

// SetLoadBalancer gets a reference to the given bool and assigns it to the LoadBalancer field.
func (o *ApplicationUpdateRequest) SetLoadBalancer(v bool) {
	o.LoadBalancer = &v
}

// GetRawLogs returns the RawLogs field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetRawLogs() bool {
	if o == nil || IsNil(o.RawLogs) {
		var ret bool
		return ret
	}
	return *o.RawLogs
}

// GetRawLogsOk returns a tuple with the RawLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetRawLogsOk() (*bool, bool) {
	if o == nil || IsNil(o.RawLogs) {
		return nil, false
	}
	return o.RawLogs, true
}

// HasRawLogs returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasRawLogs() bool {
	if o != nil && !IsNil(o.RawLogs) {
		return true
	}

	return false
}

// SetRawLogs gets a reference to the given bool and assigns it to the RawLogs field.
func (o *ApplicationUpdateRequest) SetRawLogs(v bool) {
	o.RawLogs = &v
}

// GetWebApplicationFirewall returns the WebApplicationFirewall field value if set, zero value otherwise.
func (o *ApplicationUpdateRequest) GetWebApplicationFirewall() bool {
	if o == nil || IsNil(o.WebApplicationFirewall) {
		var ret bool
		return ret
	}
	return *o.WebApplicationFirewall
}

// GetWebApplicationFirewallOk returns a tuple with the WebApplicationFirewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRequest) GetWebApplicationFirewallOk() (*bool, bool) {
	if o == nil || IsNil(o.WebApplicationFirewall) {
		return nil, false
	}
	return o.WebApplicationFirewall, true
}

// HasWebApplicationFirewall returns a boolean if a field has been set.
func (o *ApplicationUpdateRequest) HasWebApplicationFirewall() bool {
	if o != nil && !IsNil(o.WebApplicationFirewall) {
		return true
	}

	return false
}

// SetWebApplicationFirewall gets a reference to the given bool and assigns it to the WebApplicationFirewall field.
func (o *ApplicationUpdateRequest) SetWebApplicationFirewall(v bool) {
	o.WebApplicationFirewall = &v
}

func (o ApplicationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DeliveryProtocol) {
		toSerialize["delivery_protocol"] = o.DeliveryProtocol
	}
	if !IsNil(o.HttpPort) {
		toSerialize["http_port"] = o.HttpPort
	}
	if !IsNil(o.HttpsPort) {
		toSerialize["https_port"] = o.HttpsPort
	}
	if !IsNil(o.MinimumTlsVersion) {
		toSerialize["minimum_tls_version"] = o.MinimumTlsVersion
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.ApplicationAcceleration) {
		toSerialize["application_acceleration"] = o.ApplicationAcceleration
	}
	if !IsNil(o.Caching) {
		toSerialize["caching"] = o.Caching
	}
	if !IsNil(o.DeviceDetection) {
		toSerialize["device_detection"] = o.DeviceDetection
	}
	if !IsNil(o.EdgeFirewall) {
		toSerialize["edge_firewall"] = o.EdgeFirewall
	}
	if !IsNil(o.EdgeFunctions) {
		toSerialize["edge_functions"] = o.EdgeFunctions
	}
	if !IsNil(o.ImageOptimization) {
		toSerialize["image_optimization"] = o.ImageOptimization
	}
	if !IsNil(o.L2Caching) {
		toSerialize["l2_caching"] = o.L2Caching
	}
	if !IsNil(o.LoadBalancer) {
		toSerialize["load_balancer"] = o.LoadBalancer
	}
	if !IsNil(o.RawLogs) {
		toSerialize["raw_logs"] = o.RawLogs
	}
	if !IsNil(o.WebApplicationFirewall) {
		toSerialize["web_application_firewall"] = o.WebApplicationFirewall
	}
	return toSerialize, nil
}

type NullableApplicationUpdateRequest struct {
	value *ApplicationUpdateRequest
	isSet bool
}

func (v NullableApplicationUpdateRequest) Get() *ApplicationUpdateRequest {
	return v.value
}

func (v *NullableApplicationUpdateRequest) Set(val *ApplicationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationUpdateRequest(val *ApplicationUpdateRequest) *NullableApplicationUpdateRequest {
	return &NullableApplicationUpdateRequest{value: val, isSet: true}
}

func (v NullableApplicationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


