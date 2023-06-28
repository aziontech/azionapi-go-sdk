/*
Web Application Firewall API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package waf

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// WAFApiService WAFApi service
type WAFApiService service

type ApiGetWAFDomainsRequest struct {
	ctx context.Context
	ApiService *WAFApiService
	wafId int64
	name *string
}

// searches WAF for name
func (r ApiGetWAFDomainsRequest) Name(name string) ApiGetWAFDomainsRequest {
	r.name = &name
	return r
}

func (r ApiGetWAFDomainsRequest) Execute() (*WAFDomains200, *http.Response, error) {
	return r.ApiService.GetWAFDomainsExecute(r)
}

/*
GetWAFDomains Find domains attached to a WAF

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wafId ID of WAF to return
 @return ApiGetWAFDomainsRequest
*/
func (a *WAFApiService) GetWAFDomains(ctx context.Context, wafId int64) ApiGetWAFDomainsRequest {
	return ApiGetWAFDomainsRequest{
		ApiService: a,
		ctx: ctx,
		wafId: wafId,
	}
}

// Execute executes the request
//  @return WAFDomains200
func (a *WAFApiService) GetWAFDomainsExecute(r ApiGetWAFDomainsRequest) (*WAFDomains200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WAFDomains200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WAFApiService.GetWAFDomains")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/{wafId}/domains"
	localVarPath = strings.Replace(localVarPath, "{"+"wafId"+"}", url.PathEscape(parameterValueToString(r.wafId, "wafId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
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

type ApiGetWAFEventsRequest struct {
	ctx context.Context
	ApiService *WAFApiService
	wafId int64
	hourRange *int64
	domainsIds *string
	networkListId *int64
}

// Last log hours since now (it must be a integer number ranging between 1 and 72)
func (r ApiGetWAFEventsRequest) HourRange(hourRange int64) ApiGetWAFEventsRequest {
	r.hourRange = &hourRange
	return r
}

// Multiple domain&#39;s id (they must be separated by comma like 1233,1234)
func (r ApiGetWAFEventsRequest) DomainsIds(domainsIds string) ApiGetWAFEventsRequest {
	r.domainsIds = &domainsIds
	return r
}

// Id of a network list
func (r ApiGetWAFEventsRequest) NetworkListId(networkListId int64) ApiGetWAFEventsRequest {
	r.networkListId = &networkListId
	return r
}

func (r ApiGetWAFEventsRequest) Execute() (*WAFEvents200, *http.Response, error) {
	return r.ApiService.GetWAFEventsExecute(r)
}

/*
GetWAFEvents Find WAF log events

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wafId ID of WAF to return
 @return ApiGetWAFEventsRequest
*/
func (a *WAFApiService) GetWAFEvents(ctx context.Context, wafId int64) ApiGetWAFEventsRequest {
	return ApiGetWAFEventsRequest{
		ApiService: a,
		ctx: ctx,
		wafId: wafId,
	}
}

// Execute executes the request
//  @return WAFEvents200
func (a *WAFApiService) GetWAFEventsExecute(r ApiGetWAFEventsRequest) (*WAFEvents200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WAFEvents200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WAFApiService.GetWAFEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/waf/{wafId}/waf_events"
	localVarPath = strings.Replace(localVarPath, "{"+"wafId"+"}", url.PathEscape(parameterValueToString(r.wafId, "wafId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.hourRange == nil {
		return localVarReturnValue, nil, reportError("hourRange is required and must be specified")
	}
	if r.domainsIds == nil {
		return localVarReturnValue, nil, reportError("domainsIds is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "hour_range", r.hourRange, "")
	if r.networkListId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "network_list_id", r.networkListId, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "domains_ids", r.domainsIds, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v WAFEvents400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v WAFEvents404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v WAFEvents401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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