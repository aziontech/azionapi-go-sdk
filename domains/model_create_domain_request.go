/*
Domain

## Welcome to the Azion API!  With the following APIs you can check, remove or update existing settings, besides creating new ones.  * * *  ## Overview  The Azion API is a RESTful API, based on HTTPS requests, that allows you to integrate your systems with our platform, simply, quickly, and securely.  Here you will find instructions on how our API works and what functionality is available.  This documentation is being constantly updated. Make sure you verify if there are any updates or changes before you perform any development in your application, even if you are familiar with our API.  * * *  Both HTTPS requests and responses must be in JavaScript Object Notation (JSON) format. All HTTPS requests and responses must follow these conventions.  **Base URL**  The base URL of the API, which must be used, is:  [**https://api.azionapi.net/_**](https://api.azionapi.net/)  **HTTP Methods**  Each HTTP method defines the type of operation that will be run.  | HTTP Method | CRUD | Whole Collection (e.g. /items) | Specific Item (e.g. /items/:item_id) | | --- | --- | --- | --- | | GET | Read | It retrieves the list of items in a Collection. | It retrieves a specific item in a Collection. | | POST | Create | It creates a new item in the Collection. | \\- | | PUT | Update/Replace | It replaces a whole Collection with a new one. | It replaces an item in a Collection with a new one. | | PATCH | Update/Modify | It partially updates a Collection. | It partially updates an item in a Collection | | DELETE | Delete | It deletes a whole Collection. | It deletes an item in a Collection. |  * * *  **Status Codes**  The HTTP return code defines the status – successful or not – after the requested operation is run.  | Status Code | Meaning | | --- | --- | | 200 OK | General Status for a successful operation. | | 201 CREATED | Successfully created a collection or item, by means of POST or PUT. | | 204 NO CONTENT | Successful operation, but without any content to be return to the Body. Generally used for DELETE or PUT operations. | | 207 MULTI-STATUS | A batch of operations were triggered by a single request, which resulted in different return codes when it was run, for some of the operations. The codes are displayed in the “status” field, in the body of the response, for each sub-batch of operations for whichever are applicable. | | 400 BAD REQUEST | Request error, such as invalid parameters, missing mandatory parameters or others. | | 401 UNAUTHORIZED | Error caused by a missing HTTP Authentication header. | | 403 FORBIDDEN | User does not have the permissions to run the requested operation. | | 404 NOT FOUND | The requested resource does not exist. | | 405 METHOD NOT ALLOWED | The requested method is not applicable with the URL. | | 406 NOT ACCEPTABLE | Accept header missing from the HTTP request or the header contains formatting or the version is not supported by the API. | | 409 CONFLICT | Conflict generated in running the request, such as a request to create an already existing record. | | 429 TOO MANY REQUESTS | The request was temporarily rejected, due to exceeding the limit on operations. Wait for the time defined in the X-Ratelimit-Reset header and try again. | | 500 INTERNAL SERVER ERROR | Unintentional error, due to an unidentified failure in the request process. | | \\--- |  | | **HTTP Headers** |  |  All requests must be generated with the HTTP header specifying the desired code format for responses and the API version used. The current version of the API is 3 and the format is application/json.  ``` Accept: application/json; version=3  ```  * * *  **Rate Limit**  The limit of operations that can be run via the API is defined by 3 HTTP response headers:  *   **X-ratelimit-limit:** number of operations allowed in one time-frame; *   **X-ratelimit-remaining:** number of operations still available in one time-frame; *   **X-ratelimit-reset:** is the date and time, in IOS 8601 format, which represents the point in the future when the time-frame will be closed and when the limits will, therefore, be reset.       Example of HTTP response headers and a request:  ``` Date: Thu, 02 Nov 2017 12:30:28 GMT X-ratelimit-remaining: 199 X-ratelimit-limit: 200 X-ratelimit-reset: 2017-11-02T12:31:28.675446  ```  In the example provided, 200 operations are allowed within a 1-minute time frame. Of those, there are 199 still available, because one has already been run. The total limit will be reset after 1 minute.  When the X-ratelimit-limit is reached, or in other words, when the X-ratelimit-remaining reaches zero, the API will give the error, HTTP 429 TOO MANY REQUESTS. If your application receives this error, you will need to wait until the time defined in X-ratelimit-reset has passed, to make another request.  *   **X-ratelimit-limit by product**       The *X-ratelimit limit* variations by product are the following:  *   **Real-Time Metrics:** 20 requests per minute. *   **Real-Time Purge:** 200 requests per minute; except for the Wildcard, which is 2000 a day.       > The rate-limit values are based on the client_id.  * * *  **Optional Parameters**  In this version, it is possible to include some optional parameters as part of the GET method, which can help and modify the form in which your data will be returned.  You can combine these parameters to get the result you want.  They are:  | Parameter | Description | Accepted values: | | --- | --- | --- | | order_by | Identifies which field the return should be sorted by. The default ordering is ascending. | Depends on the fields available from the endpoint structure | | sort | Defines which ordering to be used: ascending or descending. | asc  <br>  <br>desc | | page | Identifies which page should be returned, if the return is paginated. The default value is 1. | Page number. | | page_size | Identifies how many items should be returned per page. The default value is 10. | Desired number of items. |  * * *  **Request Exemple**  You can use one parameter, or a combination. See the example below, which uses all of them in the same request.  ``` GET /domains?order_by=name&page_size=20&sort=desc&page=3 Accept: application/json; version=3 Authorization: token 2909f3932069047f4736dc87e72baaddd19c9f75  ```  * * *  # Authentication  Authentication and authorization of operations via Azion API is done through Tokens.  The first step is to create the Token through authenticating a user registered in [Real-Time Manager](https://sso.azion.com/login).  * * *  ## Encoding Username and Password in Base64  Only token creation and cancelling operations are performed through Basic Authentication, that is, with a username and password. You can create and cancel the token through the API itself, but for that you need to encode your username and password in base64.  Base64 encoding can be done from the command line on a Unix terminal:  ``` $ echo 'user@domain:password'|base64 dXNlckBkb21haW46cGFzc3dvcmQK  ```  If you do not have a Unix terminal available, you can use the free online service at [https://www.base64encode.org/](https://www.base64encode.org/)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package domains

import (
	"encoding/json"
)

// CreateDomainRequest struct for CreateDomainRequest
type CreateDomainRequest struct {
	Cnames []string `json:"cnames"`
	CnameAccessOnly bool `json:"cname_access_only"`
	Name string `json:"name"`
	IsActive bool `json:"is_active"`
	EdgeApplicationId int64 `json:"edge_application_id"`
	DigitalCertificateId NullableInt64 `json:"digital_certificate_id"`
}

// NewCreateDomainRequest instantiates a new CreateDomainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDomainRequest(cnames []string, cnameAccessOnly bool, name string, isActive bool, edgeApplicationId int64, digitalCertificateId NullableInt64) *CreateDomainRequest {
	this := CreateDomainRequest{}
	this.Cnames = cnames
	this.CnameAccessOnly = cnameAccessOnly
	this.Name = name
	this.IsActive = isActive
	this.EdgeApplicationId = edgeApplicationId
	this.DigitalCertificateId = digitalCertificateId
	return &this
}

// NewCreateDomainRequestWithDefaults instantiates a new CreateDomainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDomainRequestWithDefaults() *CreateDomainRequest {
	this := CreateDomainRequest{}
	return &this
}

// GetCnames returns the Cnames field value
func (o *CreateDomainRequest) GetCnames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Cnames
}

// GetCnamesOk returns a tuple with the Cnames field value
// and a boolean to check if the value has been set.
func (o *CreateDomainRequest) GetCnamesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cnames, true
}

// SetCnames sets field value
func (o *CreateDomainRequest) SetCnames(v []string) {
	o.Cnames = v
}

// GetCnameAccessOnly returns the CnameAccessOnly field value
func (o *CreateDomainRequest) GetCnameAccessOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CnameAccessOnly
}

// GetCnameAccessOnlyOk returns a tuple with the CnameAccessOnly field value
// and a boolean to check if the value has been set.
func (o *CreateDomainRequest) GetCnameAccessOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CnameAccessOnly, true
}

// SetCnameAccessOnly sets field value
func (o *CreateDomainRequest) SetCnameAccessOnly(v bool) {
	o.CnameAccessOnly = v
}

// GetName returns the Name field value
func (o *CreateDomainRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDomainRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateDomainRequest) SetName(v string) {
	o.Name = v
}

// GetIsActive returns the IsActive field value
func (o *CreateDomainRequest) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *CreateDomainRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *CreateDomainRequest) SetIsActive(v bool) {
	o.IsActive = v
}

// GetEdgeApplicationId returns the EdgeApplicationId field value
func (o *CreateDomainRequest) GetEdgeApplicationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EdgeApplicationId
}

// GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field value
// and a boolean to check if the value has been set.
func (o *CreateDomainRequest) GetEdgeApplicationIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EdgeApplicationId, true
}

// SetEdgeApplicationId sets field value
func (o *CreateDomainRequest) SetEdgeApplicationId(v int64) {
	o.EdgeApplicationId = v
}

// GetDigitalCertificateId returns the DigitalCertificateId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *CreateDomainRequest) GetDigitalCertificateId() int64 {
	if o == nil || o.DigitalCertificateId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.DigitalCertificateId.Get()
}

// GetDigitalCertificateIdOk returns a tuple with the DigitalCertificateId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateDomainRequest) GetDigitalCertificateIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DigitalCertificateId.Get(), o.DigitalCertificateId.IsSet()
}

// SetDigitalCertificateId sets field value
func (o *CreateDomainRequest) SetDigitalCertificateId(v int64) {
	o.DigitalCertificateId.Set(&v)
}

func (o CreateDomainRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cnames"] = o.Cnames
	}
	if true {
		toSerialize["cname_access_only"] = o.CnameAccessOnly
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["is_active"] = o.IsActive
	}
	if true {
		toSerialize["edge_application_id"] = o.EdgeApplicationId
	}
	if true {
		toSerialize["digital_certificate_id"] = o.DigitalCertificateId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDomainRequest struct {
	value *CreateDomainRequest
	isSet bool
}

func (v NullableCreateDomainRequest) Get() *CreateDomainRequest {
	return v.value
}

func (v *NullableCreateDomainRequest) Set(val *CreateDomainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDomainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDomainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDomainRequest(val *CreateDomainRequest) *NullableCreateDomainRequest {
	return &NullableCreateDomainRequest{value: val, isSet: true}
}

func (v NullableCreateDomainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDomainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


