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

// checks if the OriginsResultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OriginsResultResponse{}

// OriginsResultResponse struct for OriginsResultResponse
type OriginsResultResponse struct {
	OriginId int64 `json:"origin_id"`
	OriginKey string `json:"origin_key"`
	Name string `json:"name"`
	OriginType string `json:"origin_type"`
	Addresses []OriginsResultResponseAddresses `json:"addresses"`
	OriginProtocolPolicy string `json:"origin_protocol_policy"`
	IsOriginRedirectionEnabled bool `json:"is_origin_redirection_enabled"`
	HostHeader string `json:"host_header"`
	Method string `json:"method"`
	OriginPath string `json:"origin_path"`
	ConnectionTimeout int64 `json:"connection_timeout"`
	TimeoutBetweenBytes int64 `json:"timeout_between_bytes"`
	HmacAuthentication bool `json:"hmac_authentication"`
	HmacRegionName string `json:"hmac_region_name"`
	HmacAccessKey string `json:"hmac_access_key"`
	HmacSecretKey string `json:"hmac_secret_key"`
}

// NewOriginsResultResponse instantiates a new OriginsResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginsResultResponse(originId int64, originKey string, name string, originType string, addresses []OriginsResultResponseAddresses, originProtocolPolicy string, isOriginRedirectionEnabled bool, hostHeader string, method string, originPath string, connectionTimeout int64, timeoutBetweenBytes int64, hmacAuthentication bool, hmacRegionName string, hmacAccessKey string, hmacSecretKey string) *OriginsResultResponse {
	this := OriginsResultResponse{}
	this.OriginId = originId
	this.OriginKey = originKey
	this.Name = name
	this.OriginType = originType
	this.Addresses = addresses
	this.OriginProtocolPolicy = originProtocolPolicy
	this.IsOriginRedirectionEnabled = isOriginRedirectionEnabled
	this.HostHeader = hostHeader
	this.Method = method
	this.OriginPath = originPath
	this.ConnectionTimeout = connectionTimeout
	this.TimeoutBetweenBytes = timeoutBetweenBytes
	this.HmacAuthentication = hmacAuthentication
	this.HmacRegionName = hmacRegionName
	this.HmacAccessKey = hmacAccessKey
	this.HmacSecretKey = hmacSecretKey
	return &this
}

// NewOriginsResultResponseWithDefaults instantiates a new OriginsResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginsResultResponseWithDefaults() *OriginsResultResponse {
	this := OriginsResultResponse{}
	return &this
}

// GetOriginId returns the OriginId field value
func (o *OriginsResultResponse) GetOriginId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OriginId
}

// GetOriginIdOk returns a tuple with the OriginId field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetOriginIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginId, true
}

// SetOriginId sets field value
func (o *OriginsResultResponse) SetOriginId(v int64) {
	o.OriginId = v
}

// GetOriginKey returns the OriginKey field value
func (o *OriginsResultResponse) GetOriginKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginKey
}

// GetOriginKeyOk returns a tuple with the OriginKey field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetOriginKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginKey, true
}

// SetOriginKey sets field value
func (o *OriginsResultResponse) SetOriginKey(v string) {
	o.OriginKey = v
}

// GetName returns the Name field value
func (o *OriginsResultResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OriginsResultResponse) SetName(v string) {
	o.Name = v
}

// GetOriginType returns the OriginType field value
func (o *OriginsResultResponse) GetOriginType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginType
}

// GetOriginTypeOk returns a tuple with the OriginType field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetOriginTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginType, true
}

// SetOriginType sets field value
func (o *OriginsResultResponse) SetOriginType(v string) {
	o.OriginType = v
}

// GetAddresses returns the Addresses field value
func (o *OriginsResultResponse) GetAddresses() []OriginsResultResponseAddresses {
	if o == nil {
		var ret []OriginsResultResponseAddresses
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetAddressesOk() ([]OriginsResultResponseAddresses, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *OriginsResultResponse) SetAddresses(v []OriginsResultResponseAddresses) {
	o.Addresses = v
}

// GetOriginProtocolPolicy returns the OriginProtocolPolicy field value
func (o *OriginsResultResponse) GetOriginProtocolPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginProtocolPolicy
}

// GetOriginProtocolPolicyOk returns a tuple with the OriginProtocolPolicy field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetOriginProtocolPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginProtocolPolicy, true
}

// SetOriginProtocolPolicy sets field value
func (o *OriginsResultResponse) SetOriginProtocolPolicy(v string) {
	o.OriginProtocolPolicy = v
}

// GetIsOriginRedirectionEnabled returns the IsOriginRedirectionEnabled field value
func (o *OriginsResultResponse) GetIsOriginRedirectionEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsOriginRedirectionEnabled
}

// GetIsOriginRedirectionEnabledOk returns a tuple with the IsOriginRedirectionEnabled field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetIsOriginRedirectionEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsOriginRedirectionEnabled, true
}

// SetIsOriginRedirectionEnabled sets field value
func (o *OriginsResultResponse) SetIsOriginRedirectionEnabled(v bool) {
	o.IsOriginRedirectionEnabled = v
}

// GetHostHeader returns the HostHeader field value
func (o *OriginsResultResponse) GetHostHeader() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostHeader
}

// GetHostHeaderOk returns a tuple with the HostHeader field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetHostHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostHeader, true
}

// SetHostHeader sets field value
func (o *OriginsResultResponse) SetHostHeader(v string) {
	o.HostHeader = v
}

// GetMethod returns the Method field value
func (o *OriginsResultResponse) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *OriginsResultResponse) SetMethod(v string) {
	o.Method = v
}

// GetOriginPath returns the OriginPath field value
func (o *OriginsResultResponse) GetOriginPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginPath
}

// GetOriginPathOk returns a tuple with the OriginPath field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetOriginPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginPath, true
}

// SetOriginPath sets field value
func (o *OriginsResultResponse) SetOriginPath(v string) {
	o.OriginPath = v
}

// GetConnectionTimeout returns the ConnectionTimeout field value
func (o *OriginsResultResponse) GetConnectionTimeout() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ConnectionTimeout
}

// GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetConnectionTimeoutOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionTimeout, true
}

// SetConnectionTimeout sets field value
func (o *OriginsResultResponse) SetConnectionTimeout(v int64) {
	o.ConnectionTimeout = v
}

// GetTimeoutBetweenBytes returns the TimeoutBetweenBytes field value
func (o *OriginsResultResponse) GetTimeoutBetweenBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimeoutBetweenBytes
}

// GetTimeoutBetweenBytesOk returns a tuple with the TimeoutBetweenBytes field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetTimeoutBetweenBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeoutBetweenBytes, true
}

// SetTimeoutBetweenBytes sets field value
func (o *OriginsResultResponse) SetTimeoutBetweenBytes(v int64) {
	o.TimeoutBetweenBytes = v
}

// GetHmacAuthentication returns the HmacAuthentication field value
func (o *OriginsResultResponse) GetHmacAuthentication() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HmacAuthentication
}

// GetHmacAuthenticationOk returns a tuple with the HmacAuthentication field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetHmacAuthenticationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HmacAuthentication, true
}

// SetHmacAuthentication sets field value
func (o *OriginsResultResponse) SetHmacAuthentication(v bool) {
	o.HmacAuthentication = v
}

// GetHmacRegionName returns the HmacRegionName field value
func (o *OriginsResultResponse) GetHmacRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HmacRegionName
}

// GetHmacRegionNameOk returns a tuple with the HmacRegionName field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetHmacRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HmacRegionName, true
}

// SetHmacRegionName sets field value
func (o *OriginsResultResponse) SetHmacRegionName(v string) {
	o.HmacRegionName = v
}

// GetHmacAccessKey returns the HmacAccessKey field value
func (o *OriginsResultResponse) GetHmacAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HmacAccessKey
}

// GetHmacAccessKeyOk returns a tuple with the HmacAccessKey field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetHmacAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HmacAccessKey, true
}

// SetHmacAccessKey sets field value
func (o *OriginsResultResponse) SetHmacAccessKey(v string) {
	o.HmacAccessKey = v
}

// GetHmacSecretKey returns the HmacSecretKey field value
func (o *OriginsResultResponse) GetHmacSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HmacSecretKey
}

// GetHmacSecretKeyOk returns a tuple with the HmacSecretKey field value
// and a boolean to check if the value has been set.
func (o *OriginsResultResponse) GetHmacSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HmacSecretKey, true
}

// SetHmacSecretKey sets field value
func (o *OriginsResultResponse) SetHmacSecretKey(v string) {
	o.HmacSecretKey = v
}

func (o OriginsResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginsResultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["origin_id"] = o.OriginId
	toSerialize["origin_key"] = o.OriginKey
	toSerialize["name"] = o.Name
	toSerialize["origin_type"] = o.OriginType
	toSerialize["addresses"] = o.Addresses
	toSerialize["origin_protocol_policy"] = o.OriginProtocolPolicy
	toSerialize["is_origin_redirection_enabled"] = o.IsOriginRedirectionEnabled
	toSerialize["host_header"] = o.HostHeader
	toSerialize["method"] = o.Method
	toSerialize["origin_path"] = o.OriginPath
	toSerialize["connection_timeout"] = o.ConnectionTimeout
	toSerialize["timeout_between_bytes"] = o.TimeoutBetweenBytes
	toSerialize["hmac_authentication"] = o.HmacAuthentication
	toSerialize["hmac_region_name"] = o.HmacRegionName
	toSerialize["hmac_access_key"] = o.HmacAccessKey
	toSerialize["hmac_secret_key"] = o.HmacSecretKey
	return toSerialize, nil
}

type NullableOriginsResultResponse struct {
	value *OriginsResultResponse
	isSet bool
}

func (v NullableOriginsResultResponse) Get() *OriginsResultResponse {
	return v.value
}

func (v *NullableOriginsResultResponse) Set(val *OriginsResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginsResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginsResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginsResultResponse(val *OriginsResultResponse) *NullableOriginsResultResponse {
	return &NullableOriginsResultResponse{value: val, isSet: true}
}

func (v NullableOriginsResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginsResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


