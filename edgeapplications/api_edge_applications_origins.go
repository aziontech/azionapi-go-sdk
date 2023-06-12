/*
Edge Application

## Welcome to the Azion API!  With the following APIs you can check, remove or update existing settings, besides creating new ones.  * * *  ## Overview  The Azion API is a RESTful API, based on HTTPS requests, that allows you to integrate your systems with our platform, simply, quickly, and securely.  Here you will find instructions on how our API works and what functionality is available.  This documentation is being constantly updated. Make sure you verify if there are any updates or changes before you perform any development in your application, even if you are familiar with our API.  * * *  Both HTTPS requests and responses must be in JavaScript Object Notation (JSON) format. All HTTPS requests and responses must follow these conventions.  **Base URL**  The base URL of the API, which must be used, is:  [**https://api.azionapi.net/_**](https://api.azionapi.net/)  **HTTP Methods**  Each HTTP method defines the type of operation that will be run.  | HTTP Method | CRUD | Whole Collection (e.g. /items) | Specific Item (e.g. /items/:item_id) | | --- | --- | --- | --- | | GET | Read | It retrieves the list of items in a Collection. | It retrieves a specific item in a Collection. | | POST | Create | It creates a new item in the Collection. | \\- | | PUT | Update/Replace | It replaces a whole Collection with a new one. | It replaces an item in a Collection with a new one. | | PATCH | Update/Modify | It partially updates a Collection. | It partially updates an item in a Collection | | DELETE | Delete | It deletes a whole Collection. | It deletes an item in a Collection. |  * * *  **Status Codes**  The HTTP return code defines the status – successful or not – after the requested operation is run.  | Status Code | Meaning | | --- | --- | | 200 OK | General Status for a successful operation. | | 201 CREATED | Successfully created a collection or item, by means of POST or PUT. | | 204 NO CONTENT | Successful operation, but without any content to be return to the Body. Generally used for DELETE or PUT operations. | | 207 MULTI-STATUS | A batch of operations were triggered by a single request, which resulted in different return codes when it was run, for some of the operations. The codes are displayed in the “status” field, in the body of the response, for each sub-batch of operations for whichever are applicable. | | 400 BAD REQUEST | Request error, such as invalid parameters, missing mandatory parameters or others. | | 401 UNAUTHORIZED | Error caused by a missing HTTP Authentication header. | | 403 FORBIDDEN | User does not have the permissions to run the requested operation. | | 404 NOT FOUND | The requested resource does not exist. | | 405 METHOD NOT ALLOWED | The requested method is not applicable with the URL. | | 406 NOT ACCEPTABLE | Accept header missing from the HTTP request or the header contains formatting or the version is not supported by the API. | | 409 CONFLICT | Conflict generated in running the request, such as a request to create an already existing record. | | 429 TOO MANY REQUESTS | The request was temporarily rejected, due to exceeding the limit on operations. Wait for the time defined in the X-Ratelimit-Reset header and try again. | | 500 INTERNAL SERVER ERROR | Unintentional error, due to an unidentified failure in the request process. | | \\--- |  | | **HTTP Headers** |  |  All requests must be generated with the HTTP header specifying the desired code format for responses and the API version used. The current version of the API is 3 and the format is application/json.  ``` Accept: application/json; version=3  ```  * * *  **Rate Limit**  The limit of operations that can be run via the API is defined by 3 HTTP response headers:  *   **X-ratelimit-limit:** number of operations allowed in one time-frame; *   **X-ratelimit-remaining:** number of operations still available in one time-frame; *   **X-ratelimit-reset:** is the date and time, in IOS 8601 format, which represents the point in the future when the time-frame will be closed and when the limits will, therefore, be reset.       Example of HTTP response headers and a request:  ``` Date: Thu, 02 Nov 2017 12:30:28 GMT X-ratelimit-remaining: 199 X-ratelimit-limit: 200 X-ratelimit-reset: 2017-11-02T12:31:28.675446  ```  In the example provided, 200 operations are allowed within a 1-minute time frame. Of those, there are 199 still available, because one has already been run. The total limit will be reset after 1 minute.  When the X-ratelimit-limit is reached, or in other words, when the X-ratelimit-remaining reaches zero, the API will give the error, HTTP 429 TOO MANY REQUESTS. If your application receives this error, you will need to wait until the time defined in X-ratelimit-reset has passed, to make another request.  *   **X-ratelimit-limit by product**       The *X-ratelimit limit* variations by product are the following:  *   **Real-Time Metrics:** 20 requests per minute. *   **Real-Time Purge:** 200 requests per minute; except for the Wildcard, which is 2000 a day.       > The rate-limit values are based on the client_id.  * * *  **Optional Parameters**  In this version, it is possible to include some optional parameters as part of the GET method, which can help and modify the form in which your data will be returned.  You can combine these parameters to get the result you want.  They are:  | Parameter | Description | Accepted values: | | --- | --- | --- | | order_by | Identifies which field the return should be sorted by. The default ordering is ascending. | Depends on the fields available from the endpoint structure | | sort | Defines which ordering to be used: ascending or descending. | asc  <br>  <br>desc | | page | Identifies which page should be returned, if the return is paginated. The default value is 1. | Page number. | | page_size | Identifies how many items should be returned per page. The default value is 10. | Desired number of items. |  * * *  **Request Exemple**  You can use one parameter, or a combination. See the example below, which uses all of them in the same request.  ``` GET /domains?order_by=name&page_size=20&sort=desc&page=3 Accept: application/json; version=3 Authorization: token 2909f3932069047f4736dc87e72baaddd19c9f75  ```  * * *  # Authentication  Authentication and authorization of operations via Azion API is done through Tokens.  The first step is to create the Token through authenticating a user registered in [Real-Time Manager](https://sso.azion.com/login).  * * *  ## Encoding Username and Password in Base64  Only token creation and cancelling operations are performed through Basic Authentication, that is, with a username and password. You can create and cancel the token through the API itself, but for that you need to encode your username and password in base64.  Base64 encoding can be done from the command line on a Unix terminal:  ``` $ echo 'user@domain:password'|base64 dXNlckBkb21haW46cGFzc3dvcmQK  ```  If you do not have a Unix terminal available, you can use the free online service at [https://www.base64encode.org/](https://www.base64encode.org/)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// EdgeApplicationsOriginsApiService EdgeApplicationsOriginsApi service
type EdgeApplicationsOriginsApiService service

type ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsOriginsApiService
	edgeApplicationId int64
	page *int64
	pageSize *int64
	filter *string
	orderBy *string
	sort *string
	accept *string
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest) Page(page int64) ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest {
	r.page = &page
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest) PageSize(pageSize int64) ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest) Filter(filter string) ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest {
	r.filter = &filter
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest) OrderBy(orderBy string) ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest) Sort(sort string) ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest {
	r.sort = &sort
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest {
	r.accept = &accept
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest) Execute() (*OriginsResponse, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdOriginsGetExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdOriginsGet /edge_applications/{edge_application_id}/origins

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @return ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest
*/
func (a *EdgeApplicationsOriginsApiService) EdgeApplicationsEdgeApplicationIdOriginsGet(ctx context.Context, edgeApplicationId int64) ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest {
	return ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
	}
}

// Execute executes the request
//  @return OriginsResponse
func (a *EdgeApplicationsOriginsApiService) EdgeApplicationsEdgeApplicationIdOriginsGetExecute(r ApiEdgeApplicationsEdgeApplicationIdOriginsGetRequest) (*OriginsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OriginsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsOriginsApiService.EdgeApplicationsEdgeApplicationIdOriginsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/origins"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; version=3"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyDeleteRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsOriginsApiService
	edgeApplicationId int64
	originKey string
	accept *string
}

// The id of the Origin that you plan to delete.
func (r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyDeleteRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyDeleteRequest {
	r.accept = &accept
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyDeleteExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdOriginsOriginKeyDelete /edge_applications/{edge_application_id}/origins/{origin_id}

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @param originKey
 @return ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyDeleteRequest
*/
func (a *EdgeApplicationsOriginsApiService) EdgeApplicationsEdgeApplicationIdOriginsOriginKeyDelete(ctx context.Context, edgeApplicationId int64, originKey string) ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyDeleteRequest {
	return ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyDeleteRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		originKey: originKey,
	}
}

// Execute executes the request
func (a *EdgeApplicationsOriginsApiService) EdgeApplicationsEdgeApplicationIdOriginsOriginKeyDeleteExecute(r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsOriginsApiService.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/origins/{origin_key}"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"origin_key"+"}", url.PathEscape(parameterValueToString(r.originKey, "originKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyGetRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsOriginsApiService
	edgeApplicationId int64
	originKey string
	accept *string
}

// The id of the Origin that you plan to query.
func (r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyGetRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyGetRequest {
	r.accept = &accept
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyGetRequest) Execute() (*OriginsIdResponse, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGetExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGet /edge_applications/{edge_application_id}/origins/{origin_key}

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @param originKey
 @return ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyGetRequest
*/
func (a *EdgeApplicationsOriginsApiService) EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGet(ctx context.Context, edgeApplicationId int64, originKey string) ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyGetRequest {
	return ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyGetRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		originKey: originKey,
	}
}

// Execute executes the request
//  @return OriginsIdResponse
func (a *EdgeApplicationsOriginsApiService) EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGetExecute(r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyGetRequest) (*OriginsIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OriginsIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsOriginsApiService.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/origins/{origin_key}"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"origin_key"+"}", url.PathEscape(parameterValueToString(r.originKey, "originKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; version=3"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsOriginsApiService
	edgeApplicationId int64
	originKey string
	accept *string
	contentType *string
	patchOriginsRequest *PatchOriginsRequest
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchRequest {
	r.accept = &accept
	return r
}

// The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json
func (r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchRequest) ContentType(contentType string) ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchRequest {
	r.contentType = &contentType
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchRequest) PatchOriginsRequest(patchOriginsRequest PatchOriginsRequest) ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchRequest {
	r.patchOriginsRequest = &patchOriginsRequest
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchRequest) Execute() (*OriginsIdResponse, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatch /edge_applications/:edge_application_id:/origins/:origin_id:

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @param originKey
 @return ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchRequest
*/
func (a *EdgeApplicationsOriginsApiService) EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatch(ctx context.Context, edgeApplicationId int64, originKey string) ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchRequest {
	return ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		originKey: originKey,
	}
}

// Execute executes the request
//  @return OriginsIdResponse
func (a *EdgeApplicationsOriginsApiService) EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchExecute(r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchRequest) (*OriginsIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OriginsIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsOriginsApiService.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/origins/{origin_key}"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"origin_key"+"}", url.PathEscape(parameterValueToString(r.originKey, "originKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; version=3"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; version=3"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	}
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
	}
	// body params
	localVarPostBody = r.patchOriginsRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsOriginsApiService
	edgeApplicationId int64
	originKey string
	accept *string
	contentType *string
	updateOriginsRequest *UpdateOriginsRequest
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutRequest {
	r.accept = &accept
	return r
}

// The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json
func (r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutRequest) ContentType(contentType string) ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutRequest {
	r.contentType = &contentType
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutRequest) UpdateOriginsRequest(updateOriginsRequest UpdateOriginsRequest) ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutRequest {
	r.updateOriginsRequest = &updateOriginsRequest
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutRequest) Execute() (*OriginsIdResponse, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPut /edge_applications/{edge_application_id}/origins/{origin_id}

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @param originKey
 @return ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutRequest
*/
func (a *EdgeApplicationsOriginsApiService) EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPut(ctx context.Context, edgeApplicationId int64, originKey string) ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutRequest {
	return ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		originKey: originKey,
	}
}

// Execute executes the request
//  @return OriginsIdResponse
func (a *EdgeApplicationsOriginsApiService) EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutExecute(r ApiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutRequest) (*OriginsIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OriginsIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsOriginsApiService.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/origins/{origin_key}"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"origin_key"+"}", url.PathEscape(parameterValueToString(r.originKey, "originKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; version=3"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; version=3"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	}
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
	}
	// body params
	localVarPostBody = r.updateOriginsRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEdgeApplicationsEdgeApplicationIdOriginsPostRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsOriginsApiService
	edgeApplicationId int64
	accept *string
	contentType *string
	createOriginsRequest *CreateOriginsRequest
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsPostRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdOriginsPostRequest {
	r.accept = &accept
	return r
}

// The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json
func (r ApiEdgeApplicationsEdgeApplicationIdOriginsPostRequest) ContentType(contentType string) ApiEdgeApplicationsEdgeApplicationIdOriginsPostRequest {
	r.contentType = &contentType
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsPostRequest) CreateOriginsRequest(createOriginsRequest CreateOriginsRequest) ApiEdgeApplicationsEdgeApplicationIdOriginsPostRequest {
	r.createOriginsRequest = &createOriginsRequest
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdOriginsPostRequest) Execute() (*OriginsIdResponse, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdOriginsPostExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdOriginsPost /edge_applications/{edge_application_id}/origins

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @return ApiEdgeApplicationsEdgeApplicationIdOriginsPostRequest
*/
func (a *EdgeApplicationsOriginsApiService) EdgeApplicationsEdgeApplicationIdOriginsPost(ctx context.Context, edgeApplicationId int64) ApiEdgeApplicationsEdgeApplicationIdOriginsPostRequest {
	return ApiEdgeApplicationsEdgeApplicationIdOriginsPostRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
	}
}

// Execute executes the request
//  @return OriginsIdResponse
func (a *EdgeApplicationsOriginsApiService) EdgeApplicationsEdgeApplicationIdOriginsPostExecute(r ApiEdgeApplicationsEdgeApplicationIdOriginsPostRequest) (*OriginsIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OriginsIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsOriginsApiService.EdgeApplicationsEdgeApplicationIdOriginsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/origins"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; version=3"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; version=3"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	}
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
	}
	// body params
	localVarPostBody = r.createOriginsRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
