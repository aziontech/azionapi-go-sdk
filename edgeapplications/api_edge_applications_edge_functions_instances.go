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


// EdgeApplicationsEdgeFunctionsInstancesApiService EdgeApplicationsEdgeFunctionsInstancesApi service
type EdgeApplicationsEdgeFunctionsInstancesApiService service

type ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsEdgeFunctionsInstancesApiService
	edgeApplicationId string
	functionsInstancesId string
	accept *string
	contentType *string
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest {
	r.accept = &accept
	return r
}

// The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json
func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest) ContentType(contentType string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest {
	r.contentType = &contentType
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDelete /edge_applications/:edge_application_id:/functions_instances/:functions_instances_id:

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @param functionsInstancesId
 @return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest
*/
func (a *EdgeApplicationsEdgeFunctionsInstancesApiService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDelete(ctx context.Context, edgeApplicationId string, functionsInstancesId string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest {
	return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		functionsInstancesId: functionsInstancesId,
	}
}

// Execute executes the request
func (a *EdgeApplicationsEdgeFunctionsInstancesApiService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteExecute(r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsEdgeFunctionsInstancesApiService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/functions_instances/{functions_instances_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functions_instances_id"+"}", url.PathEscape(parameterValueToString(r.functionsInstancesId, "functionsInstancesId")), -1)

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
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWT"]; ok {
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

type ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsEdgeFunctionsInstancesApiService
	edgeApplicationId int64
	functionsInstancesId int64
	accept *string
}

// The id of the edge function instance you plan to query. 
func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetRequest {
	r.accept = &accept
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetRequest) Execute() (*ApplicationInstancesGetOneResponse, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet /edge_applications/:edge_application_id:/functions_instances/:functions_instances_id:

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @param functionsInstancesId
 @return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetRequest
*/
func (a *EdgeApplicationsEdgeFunctionsInstancesApiService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet(ctx context.Context, edgeApplicationId int64, functionsInstancesId int64) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetRequest {
	return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		functionsInstancesId: functionsInstancesId,
	}
}

// Execute executes the request
//  @return ApplicationInstancesGetOneResponse
func (a *EdgeApplicationsEdgeFunctionsInstancesApiService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetExecute(r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetRequest) (*ApplicationInstancesGetOneResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationInstancesGetOneResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsEdgeFunctionsInstancesApiService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/functions_instances/{functions_instances_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functions_instances_id"+"}", url.PathEscape(parameterValueToString(r.functionsInstancesId, "functionsInstancesId")), -1)

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
			if apiKey, ok := auth["JWT"]; ok {
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

type ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsEdgeFunctionsInstancesApiService
	edgeApplicationId string
	functionsInstancesId string
	accept *string
	contentType *string
	applicationUpdateInstanceRequest *ApplicationUpdateInstanceRequest
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest {
	r.accept = &accept
	return r
}

// The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json
func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest) ContentType(contentType string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest {
	r.contentType = &contentType
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest) ApplicationUpdateInstanceRequest(applicationUpdateInstanceRequest ApplicationUpdateInstanceRequest) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest {
	r.applicationUpdateInstanceRequest = &applicationUpdateInstanceRequest
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest) Execute() (*ApplicationInstanceResults, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch /edge_applications/:edge_application_id:/functions_instances/:functions_instances_id:

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId The id of the edge application you plan to overwrite 
 @param functionsInstancesId The id of the edge function instance you plan to overwrite.
 @return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest
*/
func (a *EdgeApplicationsEdgeFunctionsInstancesApiService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch(ctx context.Context, edgeApplicationId string, functionsInstancesId string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest {
	return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		functionsInstancesId: functionsInstancesId,
	}
}

// Execute executes the request
//  @return ApplicationInstanceResults
func (a *EdgeApplicationsEdgeFunctionsInstancesApiService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchExecute(r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest) (*ApplicationInstanceResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationInstanceResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsEdgeFunctionsInstancesApiService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/functions_instances/{functions_instances_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functions_instances_id"+"}", url.PathEscape(parameterValueToString(r.functionsInstancesId, "functionsInstancesId")), -1)

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
	localVarPostBody = r.applicationUpdateInstanceRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWT"]; ok {
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

type ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsEdgeFunctionsInstancesApiService
	edgeApplicationId string
	functionsInstancesId string
	accept *string
	contentType *string
	applicationPutInstanceRequest *ApplicationPutInstanceRequest
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest {
	r.accept = &accept
	return r
}

// The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json
func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest) ContentType(contentType string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest {
	r.contentType = &contentType
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest) ApplicationPutInstanceRequest(applicationPutInstanceRequest ApplicationPutInstanceRequest) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest {
	r.applicationPutInstanceRequest = &applicationPutInstanceRequest
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest) Execute() (*ApplicationInstanceResults, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut /edge_applications/:edge_application_id:/functions_instances/:functions_instances_id:

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId The id of the edge application you plan to overwrite 
 @param functionsInstancesId The id of the edge function instance you plan to overwrite.
 @return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest
*/
func (a *EdgeApplicationsEdgeFunctionsInstancesApiService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut(ctx context.Context, edgeApplicationId string, functionsInstancesId string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest {
	return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		functionsInstancesId: functionsInstancesId,
	}
}

// Execute executes the request
//  @return ApplicationInstanceResults
func (a *EdgeApplicationsEdgeFunctionsInstancesApiService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutExecute(r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest) (*ApplicationInstanceResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationInstanceResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsEdgeFunctionsInstancesApiService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/functions_instances/{functions_instances_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"functions_instances_id"+"}", url.PathEscape(parameterValueToString(r.functionsInstancesId, "functionsInstancesId")), -1)

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
	localVarPostBody = r.applicationPutInstanceRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWT"]; ok {
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

type ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsEdgeFunctionsInstancesApiService
	edgeApplicationId int64
	page *int64
	pageSize *int64
	filter *string
	orderBy *string
	sort *string
	accept *string
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest) Page(page int64) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest {
	r.page = &page
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest) PageSize(pageSize int64) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest) Filter(filter string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest {
	r.filter = &filter
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest) OrderBy(orderBy string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest) Sort(sort string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest {
	r.sort = &sort
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest {
	r.accept = &accept
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest) Execute() (*ApplicationInstancesGetResponse, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesGetExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet /edge_applications/:edge_application_id:/functions_instances

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest
*/
func (a *EdgeApplicationsEdgeFunctionsInstancesApiService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet(ctx context.Context, edgeApplicationId int64) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest {
	return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
	}
}

// Execute executes the request
//  @return ApplicationInstancesGetResponse
func (a *EdgeApplicationsEdgeFunctionsInstancesApiService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesGetExecute(r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest) (*ApplicationInstancesGetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationInstancesGetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsEdgeFunctionsInstancesApiService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/functions_instances"
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
			if apiKey, ok := auth["JWT"]; ok {
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

type ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsEdgeFunctionsInstancesApiService
	edgeApplicationId int64
	accept *string
	contentType *string
	applicationCreateInstanceRequest *ApplicationCreateInstanceRequest
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest {
	r.accept = &accept
	return r
}

// The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json
func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest) ContentType(contentType string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest {
	r.contentType = &contentType
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest) ApplicationCreateInstanceRequest(applicationCreateInstanceRequest ApplicationCreateInstanceRequest) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest {
	r.applicationCreateInstanceRequest = &applicationCreateInstanceRequest
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest) Execute() (*ApplicationInstanceResults, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesPostExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost edge_application/:edge_application_id:/functions_instances

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest
*/
func (a *EdgeApplicationsEdgeFunctionsInstancesApiService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost(ctx context.Context, edgeApplicationId int64) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest {
	return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
	}
}

// Execute executes the request
//  @return ApplicationInstanceResults
func (a *EdgeApplicationsEdgeFunctionsInstancesApiService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesPostExecute(r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest) (*ApplicationInstanceResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationInstanceResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsEdgeFunctionsInstancesApiService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/functions_instances"
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
	localVarPostBody = r.applicationCreateInstanceRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["JWT"]; ok {
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
