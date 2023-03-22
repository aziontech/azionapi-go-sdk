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

// checks if the RulesEngineResultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RulesEngineResultResponse{}

// RulesEngineResultResponse struct for RulesEngineResultResponse
type RulesEngineResultResponse struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Phase string `json:"phase"`
	Behaviors []RulesEngineResultResponseBehaviors `json:"behaviors,omitempty"`
	Criteria [][]RulesEngineCriteria `json:"criteria"`
	IsActive bool `json:"is_active"`
	Order int64 `json:"order"`
}

// NewRulesEngineResultResponse instantiates a new RulesEngineResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRulesEngineResultResponse(id int64, name string, phase string, criteria [][]RulesEngineCriteria, isActive bool, order int64) *RulesEngineResultResponse {
	this := RulesEngineResultResponse{}
	this.Id = id
	this.Name = name
	this.Phase = phase
	this.Criteria = criteria
	this.IsActive = isActive
	this.Order = order
	return &this
}

// NewRulesEngineResultResponseWithDefaults instantiates a new RulesEngineResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRulesEngineResultResponseWithDefaults() *RulesEngineResultResponse {
	this := RulesEngineResultResponse{}
	return &this
}

// GetId returns the Id field value
func (o *RulesEngineResultResponse) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RulesEngineResultResponse) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *RulesEngineResultResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RulesEngineResultResponse) SetName(v string) {
	o.Name = v
}

// GetPhase returns the Phase field value
func (o *RulesEngineResultResponse) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *RulesEngineResultResponse) SetPhase(v string) {
	o.Phase = v
}

// GetBehaviors returns the Behaviors field value if set, zero value otherwise.
func (o *RulesEngineResultResponse) GetBehaviors() []RulesEngineResultResponseBehaviors {
	if o == nil || IsNil(o.Behaviors) {
		var ret []RulesEngineResultResponseBehaviors
		return ret
	}
	return o.Behaviors
}

// GetBehaviorsOk returns a tuple with the Behaviors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetBehaviorsOk() ([]RulesEngineResultResponseBehaviors, bool) {
	if o == nil || IsNil(o.Behaviors) {
		return nil, false
	}
	return o.Behaviors, true
}

// HasBehaviors returns a boolean if a field has been set.
func (o *RulesEngineResultResponse) HasBehaviors() bool {
	if o != nil && !IsNil(o.Behaviors) {
		return true
	}

	return false
}

// SetBehaviors gets a reference to the given []RulesEngineResultResponseBehaviors and assigns it to the Behaviors field.
func (o *RulesEngineResultResponse) SetBehaviors(v []RulesEngineResultResponseBehaviors) {
	o.Behaviors = v
}

// GetCriteria returns the Criteria field value
func (o *RulesEngineResultResponse) GetCriteria() [][]RulesEngineCriteria {
	if o == nil {
		var ret [][]RulesEngineCriteria
		return ret
	}

	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetCriteriaOk() ([][]RulesEngineCriteria, bool) {
	if o == nil {
		return nil, false
	}
	return o.Criteria, true
}

// SetCriteria sets field value
func (o *RulesEngineResultResponse) SetCriteria(v [][]RulesEngineCriteria) {
	o.Criteria = v
}

// GetIsActive returns the IsActive field value
func (o *RulesEngineResultResponse) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *RulesEngineResultResponse) SetIsActive(v bool) {
	o.IsActive = v
}

// GetOrder returns the Order field value
func (o *RulesEngineResultResponse) GetOrder() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *RulesEngineResultResponse) GetOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *RulesEngineResultResponse) SetOrder(v int64) {
	o.Order = v
}

func (o RulesEngineResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RulesEngineResultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["phase"] = o.Phase
	if !IsNil(o.Behaviors) {
		toSerialize["behaviors"] = o.Behaviors
	}
	toSerialize["criteria"] = o.Criteria
	toSerialize["is_active"] = o.IsActive
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

type NullableRulesEngineResultResponse struct {
	value *RulesEngineResultResponse
	isSet bool
}

func (v NullableRulesEngineResultResponse) Get() *RulesEngineResultResponse {
	return v.value
}

func (v *NullableRulesEngineResultResponse) Set(val *RulesEngineResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRulesEngineResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRulesEngineResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRulesEngineResultResponse(val *RulesEngineResultResponse) *NullableRulesEngineResultResponse {
	return &NullableRulesEngineResultResponse{value: val, isSet: true}
}

func (v NullableRulesEngineResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRulesEngineResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


