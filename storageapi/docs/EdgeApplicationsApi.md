# \EdgeApplicationsApi

All URIs are relative to *https://edge-application-statics-storage.azion.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVersion**](EdgeApplicationsApi.md#CreateVersion) | **Post** /edge-applications/{edge_application_id}/statics | 
[**EdgeApplicationsEdgeApplicationIdStaticsVersionIdFilesPost**](EdgeApplicationsApi.md#EdgeApplicationsEdgeApplicationIdStaticsVersionIdFilesPost) | **Post** /edge-applications/{edge_application_id}/statics/{version_id}/files | 



## CreateVersion

> CreateVersion201Response CreateVersion(ctx, edgeApplicationId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    edgeApplicationId := "edgeApplicationId_example" // string | 
    body := interface{}(987) // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsApi.CreateVersion(context.Background(), edgeApplicationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsApi.CreateVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVersion`: CreateVersion201Response
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsApi.CreateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

### Return type

[**CreateVersion201Response**](CreateVersion201Response.md)

### Authorization

[authToken](../README.md#authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdStaticsVersionIdFilesPost

> interface{} EdgeApplicationsEdgeApplicationIdStaticsVersionIdFilesPost(ctx, edgeApplicationId, versionId).ContentDisposition(contentDisposition).XAzionStaticPath(xAzionStaticPath).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    contentDisposition := "contentDisposition_example" // string | Required in order to get the file name. Example: Content-Disposition: attachment; filename=index.js
    edgeApplicationId := "edgeApplicationId_example" // string | Edge Application id (global_id)
    versionId := "versionId_example" // string | 
    xAzionStaticPath := "xAzionStaticPath_example" // string | Original path file being uploaded. Given an original file path like 'assets/css/main.css' the value would be: assets/css (optional)
    body := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsApi.EdgeApplicationsEdgeApplicationIdStaticsVersionIdFilesPost(context.Background(), edgeApplicationId, versionId).ContentDisposition(contentDisposition).XAzionStaticPath(xAzionStaticPath).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsApi.EdgeApplicationsEdgeApplicationIdStaticsVersionIdFilesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdStaticsVersionIdFilesPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsApi.EdgeApplicationsEdgeApplicationIdStaticsVersionIdFilesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** | Edge Application id (global_id) | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdStaticsVersionIdFilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentDisposition** | **string** | Required in order to get the file name. Example: Content-Disposition: attachment; filename&#x3D;index.js | 


 **xAzionStaticPath** | **string** | Original path file being uploaded. Given an original file path like &#39;assets/css/main.css&#39; the value would be: assets/css | 
 **body** | ***os.File** |  | 

### Return type

**interface{}**

### Authorization

[authToken](../README.md#authToken)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

