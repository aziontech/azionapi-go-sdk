/*
Storage API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storageapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiDeleteVersionRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	versionId string
}

func (r ApiDeleteVersionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteVersionExecute(r)
}

/*
DeleteVersion /domains/:version_id

Delete a version. A version is just um path prefix/sub-namespace for a set of files.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param versionId The version identifier
 @return ApiDeleteVersionRequest
*/
func (a *DefaultApiService) DeleteVersion(ctx context.Context, versionId string) ApiDeleteVersionRequest {
	return ApiDeleteVersionRequest{
		ApiService: a,
		ctx: ctx,
		versionId: versionId,
	}
}

// Execute executes the request
func (a *DefaultApiService) DeleteVersionExecute(r ApiDeleteVersionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeleteVersion")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/storage/{version_id}/delete"
	localVarPath = strings.Replace(localVarPath, "{"+"version_id"+"}", url.PathEscape(parameterValueToString(r.versionId, "versionId")), -1)

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

type ApiStorageVersionIdPostRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	xAzionStaticPath *string
	versionId string
	body *os.File
}

// Required in order to get the path and file name. i.e.: assets/css/main.css
func (r ApiStorageVersionIdPostRequest) XAzionStaticPath(xAzionStaticPath string) ApiStorageVersionIdPostRequest {
	r.xAzionStaticPath = &xAzionStaticPath
	return r
}

func (r ApiStorageVersionIdPostRequest) Body(body *os.File) ApiStorageVersionIdPostRequest {
	r.body = body
	return r
}

func (r ApiStorageVersionIdPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.StorageVersionIdPostExecute(r)
}

/*
StorageVersionIdPost /domains/:version_id

Upload file and transfer to remote storage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param versionId 
 @return ApiStorageVersionIdPostRequest
*/
func (a *DefaultApiService) StorageVersionIdPost(ctx context.Context, versionId string) ApiStorageVersionIdPostRequest {
	return ApiStorageVersionIdPostRequest{
		ApiService: a,
		ctx: ctx,
		versionId: versionId,
	}
}

// Execute executes the request
//  @return interface{}
func (a *DefaultApiService) StorageVersionIdPostExecute(r ApiStorageVersionIdPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.StorageVersionIdPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/storage/{version_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"version_id"+"}", url.PathEscape(parameterValueToString(r.versionId, "versionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAzionStaticPath == nil {
		return localVarReturnValue, nil, reportError("xAzionStaticPath is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"b2/x-auto"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Azion-Static-Path", r.xAzionStaticPath, "")
	// body params
	localVarPostBody = r.body
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
