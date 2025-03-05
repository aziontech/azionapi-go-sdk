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


// EdgeApplicationsRulesEngineAPIService EdgeApplicationsRulesEngineAPI service
type EdgeApplicationsRulesEngineAPIService service

type ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsRulesEngineAPIService
	edgeApplicationId int64
	phase string
	page *int64
	pageSize *int64
	filter *string
	orderBy *string
	sort *string
	accept *string
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest) Page(page int64) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest {
	r.page = &page
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest) PageSize(pageSize int64) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest) Filter(filter string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest {
	r.filter = &filter
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest) OrderBy(orderBy string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest) Sort(sort string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest {
	r.sort = &sort
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest {
	r.accept = &accept
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest) Execute() (*RulesEngineResponse, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGet /edge_applications/{edge_application_id}/rules_engine/{phase}/rules

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @param phase
 @return ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest
*/
func (a *EdgeApplicationsRulesEngineAPIService) EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGet(ctx context.Context, edgeApplicationId int64, phase string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest {
	return ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		phase: phase,
	}
}

// Execute executes the request
//  @return RulesEngineResponse
func (a *EdgeApplicationsRulesEngineAPIService) EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetExecute(r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest) (*RulesEngineResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RulesEngineResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsRulesEngineAPIService.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/rules_engine/{phase}/rules"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phase"+"}", url.PathEscape(parameterValueToString(r.phase, "phase")), -1)

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

type ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsRulesEngineAPIService
	edgeApplicationId int64
	phase string
	accept *string
	contentType *string
	createRulesEngineRequest *CreateRulesEngineRequest
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostRequest {
	r.accept = &accept
	return r
}

// The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json
func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostRequest) ContentType(contentType string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostRequest {
	r.contentType = &contentType
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostRequest) CreateRulesEngineRequest(createRulesEngineRequest CreateRulesEngineRequest) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostRequest {
	r.createRulesEngineRequest = &createRulesEngineRequest
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostRequest) Execute() (*RulesEngineIdResponse, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPost /edge_applications/{edge_application_id}/rules_engine/{phase}/rules

Check below the list of behaviors that can be applied:

| Name                                | Behavior               |
| ----------------------------------- | ---------------------- |
| Add Request Cookie                  | add_request_cookie     |
| Add Request Header                  | add_request_header     |
| Add Response Cookie                 | set_cookie             |
| Add Response Header                 | add_response_header    |
| Bypass Cache                        | bypass_cache_phase     |
| Capture Match Groups                | capture_match_groups   |
| Deliver                             | deliver                |
| Deny (403 Forbidden)                | deny                   |
| Enable Gzip                         | enable_gzip            |
| Filter Request Cookie               | filter_request_cookie  |
| Filter Request Header               | filter_request_header  |
| Filter Response Cookie              | filter_response_cookie |
| Filter Response Header              | filter_response_header |
| Finish Request Phase                | finish_request_phase   |
| Forward Cookies                     | forward_cookies        |
| Optimize Images                     | optimize_images        |
| Redirect HTTP to HTTPS              | redirect_http_to_https |
| Redirect To (301 Moved Permanently) | redirect_to_301        |
| Redirect To (302 Found)             | redirect_to_302        |
| Rewrite Request                     | rewrite_request        |
| Run Function                        | run_function           |
| Set Cache Policy                    | set_cache_policy       |
| Set Origin                          | set_origin             |

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @param phase
 @return ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostRequest
*/
func (a *EdgeApplicationsRulesEngineAPIService) EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPost(ctx context.Context, edgeApplicationId int64, phase string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostRequest {
	return ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		phase: phase,
	}
}

// Execute executes the request
//  @return RulesEngineIdResponse
func (a *EdgeApplicationsRulesEngineAPIService) EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostExecute(r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostRequest) (*RulesEngineIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RulesEngineIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsRulesEngineAPIService.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/rules_engine/{phase}/rules"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phase"+"}", url.PathEscape(parameterValueToString(r.phase, "phase")), -1)

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
	localVarPostBody = r.createRulesEngineRequest
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

type ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDeleteRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsRulesEngineAPIService
	edgeApplicationId int64
	phase string
	ruleId int64
	accept *string
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDeleteRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDeleteRequest {
	r.accept = &accept
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDeleteExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDelete /edge_applications/{edge_application_id}/rules_engine/{phase}/rules

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId The id of the edge application you plan to delete. 
 @param phase
 @param ruleId The id of the rule you plan to delete. 
 @return ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDeleteRequest
*/
func (a *EdgeApplicationsRulesEngineAPIService) EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDelete(ctx context.Context, edgeApplicationId int64, phase string, ruleId int64) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDeleteRequest {
	return ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		phase: phase,
		ruleId: ruleId,
	}
}

// Execute executes the request
func (a *EdgeApplicationsRulesEngineAPIService) EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDeleteExecute(r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsRulesEngineAPIService.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/rules_engine/{phase}/rules/{rule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phase"+"}", url.PathEscape(parameterValueToString(r.phase, "phase")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rule_id"+"}", url.PathEscape(parameterValueToString(r.ruleId, "ruleId")), -1)

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

type ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGetRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsRulesEngineAPIService
	edgeApplicationId int64
	phase string
	ruleId int64
	accept *string
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGetRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGetRequest {
	r.accept = &accept
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGetRequest) Execute() (*RulesEngineIdResponse, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGetExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGet /edge_applications/{edge_application_id}/rules_engine/{phase}/rules

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId The id of the edge application you want to get. 
 @param phase
 @param ruleId The id of the rule you plan to delete. 
 @return ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGetRequest
*/
func (a *EdgeApplicationsRulesEngineAPIService) EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGet(ctx context.Context, edgeApplicationId int64, phase string, ruleId int64) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGetRequest {
	return ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGetRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		phase: phase,
		ruleId: ruleId,
	}
}

// Execute executes the request
//  @return RulesEngineIdResponse
func (a *EdgeApplicationsRulesEngineAPIService) EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGetExecute(r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGetRequest) (*RulesEngineIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RulesEngineIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsRulesEngineAPIService.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/rules_engine/{phase}/rules/{rule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phase"+"}", url.PathEscape(parameterValueToString(r.phase, "phase")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rule_id"+"}", url.PathEscape(parameterValueToString(r.ruleId, "ruleId")), -1)

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

type ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsRulesEngineAPIService
	edgeApplicationId int64
	phase string
	ruleId int64
	accept *string
	contentType *string
	patchRulesEngineRequest *PatchRulesEngineRequest
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchRequest {
	r.accept = &accept
	return r
}

// The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json
func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchRequest) ContentType(contentType string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchRequest {
	r.contentType = &contentType
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchRequest) PatchRulesEngineRequest(patchRulesEngineRequest PatchRulesEngineRequest) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchRequest {
	r.patchRulesEngineRequest = &patchRulesEngineRequest
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchRequest) Execute() (*RulesEngineIdResponse, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatch /edge_applications/:edge_application_id:/rules_engine/:phase:/rules/:rule_id:

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @param phase
 @param ruleId
 @return ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchRequest
*/
func (a *EdgeApplicationsRulesEngineAPIService) EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatch(ctx context.Context, edgeApplicationId int64, phase string, ruleId int64) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchRequest {
	return ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		phase: phase,
		ruleId: ruleId,
	}
}

// Execute executes the request
//  @return RulesEngineIdResponse
func (a *EdgeApplicationsRulesEngineAPIService) EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchExecute(r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchRequest) (*RulesEngineIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RulesEngineIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsRulesEngineAPIService.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/rules_engine/{phase}/rules/{rule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phase"+"}", url.PathEscape(parameterValueToString(r.phase, "phase")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rule_id"+"}", url.PathEscape(parameterValueToString(r.ruleId, "ruleId")), -1)

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
	localVarPostBody = r.patchRulesEngineRequest
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

type ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutRequest struct {
	ctx context.Context
	ApiService *EdgeApplicationsRulesEngineAPIService
	edgeApplicationId int64
	phase string
	ruleId int64
	accept *string
	contentType *string
	updateRulesEngineRequest *UpdateRulesEngineRequest
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutRequest) Accept(accept string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutRequest {
	r.accept = &accept
	return r
}

// The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json
func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutRequest) ContentType(contentType string) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutRequest {
	r.contentType = &contentType
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutRequest) UpdateRulesEngineRequest(updateRulesEngineRequest UpdateRulesEngineRequest) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutRequest {
	r.updateRulesEngineRequest = &updateRulesEngineRequest
	return r
}

func (r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutRequest) Execute() (*RulesEngineIdResponse, *http.Response, error) {
	return r.ApiService.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutExecute(r)
}

/*
EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPut /edge_applications/:edge_application_id:/rules_engine/:phase:/rules/:rule_id:

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param edgeApplicationId
 @param phase
 @param ruleId
 @return ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutRequest
*/
func (a *EdgeApplicationsRulesEngineAPIService) EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPut(ctx context.Context, edgeApplicationId int64, phase string, ruleId int64) ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutRequest {
	return ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutRequest{
		ApiService: a,
		ctx: ctx,
		edgeApplicationId: edgeApplicationId,
		phase: phase,
		ruleId: ruleId,
	}
}

// Execute executes the request
//  @return RulesEngineIdResponse
func (a *EdgeApplicationsRulesEngineAPIService) EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutExecute(r ApiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutRequest) (*RulesEngineIdResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RulesEngineIdResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EdgeApplicationsRulesEngineAPIService.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/edge_applications/{edge_application_id}/rules_engine/{phase}/rules/{rule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"edge_application_id"+"}", url.PathEscape(parameterValueToString(r.edgeApplicationId, "edgeApplicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"phase"+"}", url.PathEscape(parameterValueToString(r.phase, "phase")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"rule_id"+"}", url.PathEscape(parameterValueToString(r.ruleId, "ruleId")), -1)

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
	localVarPostBody = r.updateRulesEngineRequest
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
