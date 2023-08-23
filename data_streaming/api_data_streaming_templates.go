/*
Data Streaming - OpenAPI

The Data Streaming API allows you to manage your existing data streamings and templates. Data Streaming allows you to feed your stream processing, SIEM, and big data platforms with the event logs from your applications on Azion in real time. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data_streaming

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// DataStreamingTemplatesApiService DataStreamingTemplatesApi service
type DataStreamingTemplatesApiService service

type ApiGetDataStramingTemplateByIdRequest struct {
	ctx context.Context
	ApiService *DataStreamingTemplatesApiService
	templateId int32
}

func (r ApiGetDataStramingTemplateByIdRequest) Execute() (*TemplateResultById, *http.Response, error) {
	return r.ApiService.GetDataStramingTemplateByIdExecute(r)
}

/*
GetDataStramingTemplateById Get an global Template info by template ID

Use the GET method and add the data streaming's ID to the URI of the request to get more data on a specific data streaming global template.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateId
 @return ApiGetDataStramingTemplateByIdRequest
*/
func (a *DataStreamingTemplatesApiService) GetDataStramingTemplateById(ctx context.Context, templateId int32) ApiGetDataStramingTemplateByIdRequest {
	return ApiGetDataStramingTemplateByIdRequest{
		ApiService: a,
		ctx: ctx,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return TemplateResultById
func (a *DataStreamingTemplatesApiService) GetDataStramingTemplateByIdExecute(r ApiGetDataStramingTemplateByIdRequest) (*TemplateResultById, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TemplateResultById
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataStreamingTemplatesApiService.GetDataStramingTemplateById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data_streaming/templates/{template_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"template_id"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiListDataStreamingTemplatesRequest struct {
	ctx context.Context
	ApiService *DataStreamingTemplatesApiService
}

func (r ApiListDataStreamingTemplatesRequest) Execute() (*TemplateResults, *http.Response, error) {
	return r.ApiService.ListDataStreamingTemplatesExecute(r)
}

/*
ListDataStreamingTemplates List all global Templates that can be used on Data Streaming operations

Use the GET method to list all global templates that can be used on Data Streaming operations.

**Note:** Customized templates won't be listed.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDataStreamingTemplatesRequest
*/
func (a *DataStreamingTemplatesApiService) ListDataStreamingTemplates(ctx context.Context) ApiListDataStreamingTemplatesRequest {
	return ApiListDataStreamingTemplatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TemplateResults
func (a *DataStreamingTemplatesApiService) ListDataStreamingTemplatesExecute(r ApiListDataStreamingTemplatesRequest) (*TemplateResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TemplateResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataStreamingTemplatesApiService.ListDataStreamingTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data_streaming/templates"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
