# \DataStreamingApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewDataStreaming**](DataStreamingApi.md#CreateNewDataStreaming) | **Post** /data_streaming/streamings | Create a new data streaming
[**DeleteDataStreamingById**](DataStreamingApi.md#DeleteDataStreamingById) | **Delete** /data_streaming/streamings/{data_streaming_id} | Delete data streaming
[**EditDataStreamingById**](DataStreamingApi.md#EditDataStreamingById) | **Patch** /data_streaming/streamings/{data_streaming_id} | Edit data streaming
[**ListDataStreamingById**](DataStreamingApi.md#ListDataStreamingById) | **Get** /data_streaming/streamings/{data_streaming_id} | Get expecific data streaming by Data Streaming ID
[**ListDataStreamings**](DataStreamingApi.md#ListDataStreamings) | **Get** /data_streaming/streamings | List all exist data streamings
[**OverwriteDataStreamingById**](DataStreamingApi.md#OverwriteDataStreamingById) | **Put** /data_streaming/streamings/{data_streaming_id} | Overwrite data streaming



## CreateNewDataStreaming

> CreateNewDataStreaming201Response CreateNewDataStreaming(ctx).CreateNewDataStreamingRequest(createNewDataStreamingRequest).Execute()

Create a new data streaming



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createNewDataStreamingRequest := openapiclient.CreateNewDataStreaming_request{CustomDataStreamingPostBody: openapiclient.NewCustomDataStreamingPostBody()} // CreateNewDataStreamingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStreamingApi.CreateNewDataStreaming(context.Background()).CreateNewDataStreamingRequest(createNewDataStreamingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStreamingApi.CreateNewDataStreaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewDataStreaming`: CreateNewDataStreaming201Response
    fmt.Fprintf(os.Stdout, "Response from `DataStreamingApi.CreateNewDataStreaming`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewDataStreamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNewDataStreamingRequest** | [**CreateNewDataStreamingRequest**](CreateNewDataStreamingRequest.md) |  | 

### Return type

[**CreateNewDataStreaming201Response**](CreateNewDataStreaming201Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDataStreamingById

> DeleteDataStreamingById(ctx, dataStreamingId).Execute()

Delete data streaming



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataStreamingId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DataStreamingApi.DeleteDataStreamingById(context.Background(), dataStreamingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStreamingApi.DeleteDataStreamingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataStreamingId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataStreamingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditDataStreamingById

> CreateNewDataStreaming201Response EditDataStreamingById(ctx, dataStreamingId).CreateNewDataStreamingRequest(createNewDataStreamingRequest).Execute()

Edit data streaming



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataStreamingId := int32(56) // int32 | 
    createNewDataStreamingRequest := openapiclient.CreateNewDataStreaming_request{CustomDataStreamingPostBody: openapiclient.NewCustomDataStreamingPostBody()} // CreateNewDataStreamingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStreamingApi.EditDataStreamingById(context.Background(), dataStreamingId).CreateNewDataStreamingRequest(createNewDataStreamingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStreamingApi.EditDataStreamingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditDataStreamingById`: CreateNewDataStreaming201Response
    fmt.Fprintf(os.Stdout, "Response from `DataStreamingApi.EditDataStreamingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataStreamingId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditDataStreamingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNewDataStreamingRequest** | [**CreateNewDataStreamingRequest**](CreateNewDataStreamingRequest.md) |  | 

### Return type

[**CreateNewDataStreaming201Response**](CreateNewDataStreaming201Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataStreamingById

> DataStreamingsById ListDataStreamingById(ctx, dataStreamingId).Execute()

Get expecific data streaming by Data Streaming ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataStreamingId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStreamingApi.ListDataStreamingById(context.Background(), dataStreamingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStreamingApi.ListDataStreamingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDataStreamingById`: DataStreamingsById
    fmt.Fprintf(os.Stdout, "Response from `DataStreamingApi.ListDataStreamingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataStreamingId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDataStreamingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataStreamingsById**](DataStreamingsById.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataStreamings

> DataStreamingResponseWithResults ListDataStreamings(ctx).Execute()

List all exist data streamings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStreamingApi.ListDataStreamings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStreamingApi.ListDataStreamings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDataStreamings`: DataStreamingResponseWithResults
    fmt.Fprintf(os.Stdout, "Response from `DataStreamingApi.ListDataStreamings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDataStreamingsRequest struct via the builder pattern


### Return type

[**DataStreamingResponseWithResults**](DataStreamingResponseWithResults.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverwriteDataStreamingById

> CreateNewDataStreaming201Response OverwriteDataStreamingById(ctx, dataStreamingId).CreateNewDataStreamingRequest(createNewDataStreamingRequest).Execute()

Overwrite data streaming



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    dataStreamingId := int32(56) // int32 | 
    createNewDataStreamingRequest := openapiclient.CreateNewDataStreaming_request{CustomDataStreamingPostBody: openapiclient.NewCustomDataStreamingPostBody()} // CreateNewDataStreamingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStreamingApi.OverwriteDataStreamingById(context.Background(), dataStreamingId).CreateNewDataStreamingRequest(createNewDataStreamingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStreamingApi.OverwriteDataStreamingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverwriteDataStreamingById`: CreateNewDataStreaming201Response
    fmt.Fprintf(os.Stdout, "Response from `DataStreamingApi.OverwriteDataStreamingById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataStreamingId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverwriteDataStreamingByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNewDataStreamingRequest** | [**CreateNewDataStreamingRequest**](CreateNewDataStreamingRequest.md) |  | 

### Return type

[**CreateNewDataStreaming201Response**](CreateNewDataStreaming201Response.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

