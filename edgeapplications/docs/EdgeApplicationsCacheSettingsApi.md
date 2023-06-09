# \EdgeApplicationsCacheSettingsApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsDelete**](EdgeApplicationsCacheSettingsApi.md#EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsDelete) | **Delete** /edge_applications/{edge_application_id}/cache_settings/{cache_settings} | /edge_applications/:edge_application_id:/cache_settings/:cache_settings:
[**EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet**](EdgeApplicationsCacheSettingsApi.md#EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet) | **Get** /edge_applications/{edge_application_id}/cache_settings/{cache_settings_id} | /edge_applications/:edge_application_id:/cache_settings/:cache_settings_id:
[**EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut**](EdgeApplicationsCacheSettingsApi.md#EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut) | **Put** /edge_applications/{edge_application_id}/cache_settings/{cache_settings_id} | /edge_applications/:edge_application_id:/cache_settings/ca
[**EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsPatch**](EdgeApplicationsCacheSettingsApi.md#EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsPatch) | **Patch** /edge_applications/{edge_application_id}/cache_settings/{cache_settings} | /edge_applications/:edge_application_id:/cache_settings/:cache_settings_id:
[**EdgeApplicationsEdgeApplicationIdCacheSettingsGet**](EdgeApplicationsCacheSettingsApi.md#EdgeApplicationsEdgeApplicationIdCacheSettingsGet) | **Get** /edge_applications/{edge_application_id}/cache_settings | /edge_applications/{edge_application_id}/cache_settings
[**EdgeApplicationsEdgeApplicationIdCacheSettingsPost**](EdgeApplicationsCacheSettingsApi.md#EdgeApplicationsEdgeApplicationIdCacheSettingsPost) | **Post** /edge_applications/{edge_application_id}/cache_settings | /edge_applications/:edge_application_id:/cache_settings



## EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsDelete

> EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsDelete(ctx, edgeApplicationId, cacheSettings).Accept(accept).ContentType(contentType).Execute()

/edge_applications/:edge_application_id:/cache_settings/:cache_settings:

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
    edgeApplicationId := int64(789) // int64 | 
    cacheSettings := int64(789) // int64 | 
    accept := "application/json; version=3" // string |  (optional)
    contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsDelete(context.Background(), edgeApplicationId, cacheSettings).Accept(accept).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**cacheSettings** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 

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


## EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet

> ApplicationCacheGetOneResponse EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet(ctx, edgeApplicationId, cacheSettingsId).Accept(accept).Execute()

/edge_applications/:edge_application_id:/cache_settings/:cache_settings_id:

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
    edgeApplicationId := int64(789) // int64 | 
    cacheSettingsId := int64(789) // int64 | 
    accept := "application/json; version=3" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet(context.Background(), edgeApplicationId, cacheSettingsId).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet`: ApplicationCacheGetOneResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**cacheSettingsId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 

### Return type

[**ApplicationCacheGetOneResponse**](ApplicationCacheGetOneResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut

> ApplicationCachePutResponse EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut(ctx, edgeApplicationId, cacheSettingsId).Accept(accept).ContentType(contentType).ApplicationCachePutRequest(applicationCachePutRequest).Execute()

/edge_applications/:edge_application_id:/cache_settings/ca

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
    edgeApplicationId := int64(789) // int64 | 
    cacheSettingsId := int64(789) // int64 | 
    accept := "application/json; version=3" // string |  (optional)
    contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
    applicationCachePutRequest := *openapiclient.NewApplicationCachePutRequest("Name_example") // ApplicationCachePutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut(context.Background(), edgeApplicationId, cacheSettingsId).Accept(accept).ContentType(contentType).ApplicationCachePutRequest(applicationCachePutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut`: ApplicationCachePutResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**cacheSettingsId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **applicationCachePutRequest** | [**ApplicationCachePutRequest**](ApplicationCachePutRequest.md) |  | 

### Return type

[**ApplicationCachePutResponse**](ApplicationCachePutResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsPatch

> ApplicationCachePatchResponse EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsPatch(ctx, edgeApplicationId, cacheSettings).Accept(accept).ApplicationCachePatchRequest(applicationCachePatchRequest).Execute()

/edge_applications/:edge_application_id:/cache_settings/:cache_settings_id:

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
    edgeApplicationId := int64(789) // int64 | 
    cacheSettings := int64(789) // int64 | 
    accept := "application/json; version=3" // string |  (optional)
    applicationCachePatchRequest := *openapiclient.NewApplicationCachePatchRequest() // ApplicationCachePatchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsPatch(context.Background(), edgeApplicationId, cacheSettings).Accept(accept).ApplicationCachePatchRequest(applicationCachePatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsPatch`: ApplicationCachePatchResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**cacheSettings** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **applicationCachePatchRequest** | [**ApplicationCachePatchRequest**](ApplicationCachePatchRequest.md) |  | 

### Return type

[**ApplicationCachePatchResponse**](ApplicationCachePatchResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdCacheSettingsGet

> ApplicationCacheGetResponse EdgeApplicationsEdgeApplicationIdCacheSettingsGet(ctx, edgeApplicationId).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()

/edge_applications/{edge_application_id}/cache_settings

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
    edgeApplicationId := int64(789) // int64 | 
    page := int64(789) // int64 |  (optional)
    pageSize := int64(789) // int64 |  (optional)
    filter := "filter_example" // string |  (optional)
    orderBy := "orderBy_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    accept := "application/json; version=3" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsGet(context.Background(), edgeApplicationId).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdCacheSettingsGet`: ApplicationCacheGetResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdCacheSettingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | 
 **pageSize** | **int64** |  | 
 **filter** | **string** |  | 
 **orderBy** | **string** |  | 
 **sort** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**ApplicationCacheGetResponse**](ApplicationCacheGetResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdCacheSettingsPost

> ApplicationCacheCreateResponse EdgeApplicationsEdgeApplicationIdCacheSettingsPost(ctx, edgeApplicationId).Accept(accept).ContentType(contentType).ApplicationCacheCreateRequest(applicationCacheCreateRequest).Execute()

/edge_applications/:edge_application_id:/cache_settings

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
    edgeApplicationId := int64(789) // int64 | 
    accept := "application/json; version=3" // string |  (optional)
    contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
    applicationCacheCreateRequest := *openapiclient.NewApplicationCacheCreateRequest("Name_example") // ApplicationCacheCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsPost(context.Background(), edgeApplicationId).Accept(accept).ContentType(contentType).ApplicationCacheCreateRequest(applicationCacheCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdCacheSettingsPost`: ApplicationCacheCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsApi.EdgeApplicationsEdgeApplicationIdCacheSettingsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdCacheSettingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **applicationCacheCreateRequest** | [**ApplicationCacheCreateRequest**](ApplicationCacheCreateRequest.md) |  | 

### Return type

[**ApplicationCacheCreateResponse**](ApplicationCacheCreateResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

