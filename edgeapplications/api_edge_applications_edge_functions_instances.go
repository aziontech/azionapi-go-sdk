/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
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


// EdgeApplicationsEdgeFunctionsInstancesAPIService EdgeApplicationsEdgeFunctionsInstancesAPI service
type EdgeApplicationsEdgeFunctionsInstancesAPIService service

type ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsEdgeFunctionsInstancesAPIService
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
func (a *EdgeApplicationsEdgeFunctionsInstancesAPIService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDelete(ctx context.Context, edgeApplicationId string, functionsInstancesId string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest {
	return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		functionsInstancesId: functionsInstancesId,
	}
}

// Execute executes the request
func (a *EdgeApplicationsEdgeFunctionsInstancesAPIService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteExecute(r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsEdgeFunctionsInstancesAPIService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDelete")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
	}
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "simple", "")
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

type ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsEdgeFunctionsInstancesAPIService
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
func (a *EdgeApplicationsEdgeFunctionsInstancesAPIService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet(ctx context.Context, edgeApplicationId int64, functionsInstancesId int64) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetRequest {
	return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		functionsInstancesId: functionsInstancesId,
	}
}

// Execute executes the request
//  @return ApplicationInstancesGetOneResponse
func (a *EdgeApplicationsEdgeFunctionsInstancesAPIService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetExecute(r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetRequest) (*ApplicationInstancesGetOneResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationInstancesGetOneResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsEdgeFunctionsInstancesAPIService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
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

type ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsEdgeFunctionsInstancesAPIService
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
func (a *EdgeApplicationsEdgeFunctionsInstancesAPIService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch(ctx context.Context, edgeApplicationId string, functionsInstancesId string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest {
	return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		functionsInstancesId: functionsInstancesId,
	}
}

// Execute executes the request
//  @return ApplicationInstanceResults
func (a *EdgeApplicationsEdgeFunctionsInstancesAPIService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchExecute(r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest) (*ApplicationInstanceResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationInstanceResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsEdgeFunctionsInstancesAPIService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
	}
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "simple", "")
	}
	// body params
	localVarPostBody = r.applicationUpdateInstanceRequest
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

type ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsEdgeFunctionsInstancesAPIService
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
func (a *EdgeApplicationsEdgeFunctionsInstancesAPIService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut(ctx context.Context, edgeApplicationId string, functionsInstancesId string) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest {
	return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		functionsInstancesId: functionsInstancesId,
	}
}

// Execute executes the request
//  @return ApplicationInstanceResults
func (a *EdgeApplicationsEdgeFunctionsInstancesAPIService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutExecute(r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest) (*ApplicationInstanceResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationInstanceResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsEdgeFunctionsInstancesAPIService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
	}
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "simple", "")
	}
	// body params
	localVarPostBody = r.applicationPutInstanceRequest
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

type ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsEdgeFunctionsInstancesAPIService
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
func (a *EdgeApplicationsEdgeFunctionsInstancesAPIService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet(ctx context.Context, edgeApplicationId int64) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest {
	return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
	}
}

// Execute executes the request
//  @return ApplicationInstancesGetResponse
func (a *EdgeApplicationsEdgeFunctionsInstancesAPIService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesGetExecute(r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest) (*ApplicationInstancesGetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationInstancesGetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsEdgeFunctionsInstancesAPIService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/functions_instances"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
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

type ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsEdgeFunctionsInstancesAPIService
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
func (a *EdgeApplicationsEdgeFunctionsInstancesAPIService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost(ctx context.Context, edgeApplicationId int64) ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest {
	return ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
	}
}

// Execute executes the request
//  @return ApplicationInstanceResults
func (a *EdgeApplicationsEdgeFunctionsInstancesAPIService) EdgeApplicationsEdgeApplicationIdFunctionsInstancesPostExecute(r ApiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest) (*ApplicationInstanceResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationInstanceResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsEdgeFunctionsInstancesAPIService.EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost")
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
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
	}
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "simple", "")
	}
	// body params
	localVarPostBody = r.applicationCreateInstanceRequest
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
