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

// checks if the ApplicationPutInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationPutInstanceRequest{}

// ApplicationPutInstanceRequest struct for ApplicationPutInstanceRequest
type ApplicationPutInstanceRequest struct {
	Name string `json:"name"`
	EdgeFunctionId int64 `json:"edge_function_id"`
	Args interface{} `json:"args"`
}

// NewApplicationPutInstanceRequest instantiates a new ApplicationPutInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationPutInstanceRequest(name string, edgeFunctionId int64, args interface{}) *ApplicationPutInstanceRequest {
	this := ApplicationPutInstanceRequest{}
	this.Name = name
	this.EdgeFunctionId = edgeFunctionId
	this.Args = args
	return &this
}

// NewApplicationPutInstanceRequestWithDefaults instantiates a new ApplicationPutInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationPutInstanceRequestWithDefaults() *ApplicationPutInstanceRequest {
	this := ApplicationPutInstanceRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationPutInstanceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationPutInstanceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationPutInstanceRequest) SetName(v string) {
	o.Name = v
}

// GetEdgeFunctionId returns the EdgeFunctionId field value
func (o *ApplicationPutInstanceRequest) GetEdgeFunctionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EdgeFunctionId
}

// GetEdgeFunctionIdOk returns a tuple with the EdgeFunctionId field value
// and a boolean to check if the value has been set.
func (o *ApplicationPutInstanceRequest) GetEdgeFunctionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EdgeFunctionId, true
}

// SetEdgeFunctionId sets field value
func (o *ApplicationPutInstanceRequest) SetEdgeFunctionId(v int64) {
	o.EdgeFunctionId = v
}

// GetArgs returns the Args field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApplicationPutInstanceRequest) GetArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Args
}

// GetArgsOk returns a tuple with the Args field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationPutInstanceRequest) GetArgsOk() (*interface{}, bool) {
	if o == nil || isNil(o.Args) {
		return nil, false
	}
	return &o.Args, true
}

// SetArgs sets field value
func (o *ApplicationPutInstanceRequest) SetArgs(v interface{}) {
	o.Args = v
}

func (o ApplicationPutInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationPutInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["edge_function_id"] = o.EdgeFunctionId
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	return toSerialize, nil
}

type NullableApplicationPutInstanceRequest struct {
	value *ApplicationPutInstanceRequest
	isSet bool
}

func (v NullableApplicationPutInstanceRequest) Get() *ApplicationPutInstanceRequest {
	return v.value
}

func (v *NullableApplicationPutInstanceRequest) Set(val *ApplicationPutInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationPutInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationPutInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationPutInstanceRequest(val *ApplicationPutInstanceRequest) *NullableApplicationPutInstanceRequest {
	return &NullableApplicationPutInstanceRequest{value: val, isSet: true}
}

func (v NullableApplicationPutInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationPutInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


