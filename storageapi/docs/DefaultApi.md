# \DefaultApi

All URIs are relative to *https://storage-api.azion.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVersion**](DefaultApi.md#DeleteVersion) | **Delete** /storage/{version_id}/delete | /storage/:version_id/delete
[**StorageVersionIdPost**](DefaultApi.md#StorageVersionIdPost) | **Post** /storage/{version_id} | /storage/:version_id



## DeleteVersion

> DeleteVersion(ctx, versionId).Execute()

/storage/:version_id/delete



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
    versionId := "versionId_example" // string | The version identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteVersion(context.Background(), versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | The version identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageVersionIdPost

> interface{} StorageVersionIdPost(ctx, versionId).XAzionStaticPath(xAzionStaticPath).ContentType(contentType).Body(body).Execute()

/storage/:version_id



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
    xAzionStaticPath := "xAzionStaticPath_example" // string | Required in order to get the path and file name. i.e.: assets/css/main.css
    versionId := "versionId_example" // string | 
    contentType := "contentType_example" // string | The content type of the file (Example: text/plain). (optional) (default to "b2/x-auto")
    body := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.StorageVersionIdPost(context.Background(), versionId).XAzionStaticPath(xAzionStaticPath).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StorageVersionIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StorageVersionIdPost`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.StorageVersionIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageVersionIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAzionStaticPath** | **string** | Required in order to get the path and file name. i.e.: assets/css/main.css | 

 **contentType** | **string** | The content type of the file (Example: text/plain). | [default to &quot;b2/x-auto&quot;]
 **body** | ***os.File** |  | 

### Return type

**interface{}**

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

