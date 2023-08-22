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
)


// DataStreamingDomainApiService DataStreamingDomainApi service
type DataStreamingDomainApiService service

type ApiListDataStreamingRequest struct {
	ctx context.Context
	ApiService *DataStreamingDomainApiService
	name *string
	streamingId *int64
	selected *bool
}

// Domain&#39;s name in data streaming
func (r ApiListDataStreamingRequest) Name(name string) ApiListDataStreamingRequest {
	r.name = &name
	return r
}

func (r ApiListDataStreamingRequest) StreamingId(streamingId int64) ApiListDataStreamingRequest {
	r.streamingId = &streamingId
	return r
}

func (r ApiListDataStreamingRequest) Selected(selected bool) ApiListDataStreamingRequest {
	r.selected = &selected
	return r
}

func (r ApiListDataStreamingRequest) Execute() (*DataStreamingsDomainResponse, *http.Response, error) {
	return r.ApiService.ListDataStreamingExecute(r)
}

/*
ListDataStreaming List all domains used on data streaming

Use the GET method to list all available domains that can be used on Data Streaming operations.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDataStreamingRequest
*/
func (a *DataStreamingDomainApiService) ListDataStreaming(ctx context.Context) ApiListDataStreamingRequest {
	return ApiListDataStreamingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DataStreamingsDomainResponse
func (a *DataStreamingDomainApiService) ListDataStreamingExecute(r ApiListDataStreamingRequest) (*DataStreamingsDomainResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DataStreamingsDomainResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataStreamingDomainApiService.ListDataStreaming")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/data_streaming/domains"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.streamingId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "streaming_id", r.streamingId, "")
	}
	if r.selected != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "selected", r.selected, "")
	}
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
