# \EdgeApplicationsOriginsApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeApplicationsEdgeApplicationIdOriginsGet**](EdgeApplicationsOriginsApi.md#EdgeApplicationsEdgeApplicationIdOriginsGet) | **Get** /edge_applications/{edge_application_id}/origins | /edge_applications/{edge_application_id}/origins
[**EdgeApplicationsEdgeApplicationIdOriginsOriginKeyDelete**](EdgeApplicationsOriginsApi.md#EdgeApplicationsEdgeApplicationIdOriginsOriginKeyDelete) | **Delete** /edge_applications/{edge_application_id}/origins/{origin_key} | /edge_applications/{edge_application_id}/origins/{origin_id}
[**EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGet**](EdgeApplicationsOriginsApi.md#EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGet) | **Get** /edge_applications/{edge_application_id}/origins/{origin_key} | /edge_applications/{edge_application_id}/origins/{origin_key}
[**EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatch**](EdgeApplicationsOriginsApi.md#EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatch) | **Patch** /edge_applications/{edge_application_id}/origins/{origin_key} | /edge_applications/:edge_application_id:/origins/:origin_id:
[**EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPut**](EdgeApplicationsOriginsApi.md#EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPut) | **Put** /edge_applications/{edge_application_id}/origins/{origin_key} | /edge_applications/{edge_application_id}/origins/{origin_id}
[**EdgeApplicationsEdgeApplicationIdOriginsPost**](EdgeApplicationsOriginsApi.md#EdgeApplicationsEdgeApplicationIdOriginsPost) | **Post** /edge_applications/{edge_application_id}/origins | /edge_applications/{edge_application_id}/origins



## EdgeApplicationsEdgeApplicationIdOriginsGet

> OriginsResponse EdgeApplicationsEdgeApplicationIdOriginsGet(ctx, edgeApplicationId).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()

/edge_applications/{edge_application_id}/origins

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
    resp, r, err := apiClient.EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsGet(context.Background(), edgeApplicationId).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdOriginsGet`: OriginsResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdOriginsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | 
 **pageSize** | **int64** |  | 
 **filter** | **string** |  | 
 **orderBy** | **string** |  | 
 **sort** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**OriginsResponse**](OriginsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdOriginsOriginKeyDelete

> EdgeApplicationsEdgeApplicationIdOriginsOriginKeyDelete(ctx, edgeApplicationId, originKey).Accept(accept).Execute()

/edge_applications/{edge_application_id}/origins/{origin_id}

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
    originKey := "originKey_example" // string | 
    accept := "application/json; version=3" // string | The id of the Origin that you plan to delete. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyDelete(context.Background(), edgeApplicationId, originKey).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**originKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** | The id of the Origin that you plan to delete. | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGet

> OriginsIdResponse EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGet(ctx, edgeApplicationId, originKey).Accept(accept).Execute()

/edge_applications/{edge_application_id}/origins/{origin_key}

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
    originKey := "originKey_example" // string | 
    accept := "application/json; version=3" // string | The id of the Origin that you plan to query. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGet(context.Background(), edgeApplicationId, originKey).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGet`: OriginsIdResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**originKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** | The id of the Origin that you plan to query. | 

### Return type

[**OriginsIdResponse**](OriginsIdResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatch

> OriginsIdResponse EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatch(ctx, edgeApplicationId, originKey).Accept(accept).ContentType(contentType).PatchOriginsRequest(patchOriginsRequest).Execute()

/edge_applications/:edge_application_id:/origins/:origin_id:

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
    originKey := "originKey_example" // string | 
    accept := "application/json; version=3" // string |  (optional)
    contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
    patchOriginsRequest := *openapiclient.NewPatchOriginsRequest() // PatchOriginsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatch(context.Background(), edgeApplicationId, originKey).Accept(accept).ContentType(contentType).PatchOriginsRequest(patchOriginsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatch`: OriginsIdResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**originKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **patchOriginsRequest** | [**PatchOriginsRequest**](PatchOriginsRequest.md) |  | 

### Return type

[**OriginsIdResponse**](OriginsIdResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPut

> OriginsIdResponse EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPut(ctx, edgeApplicationId, originKey).Accept(accept).ContentType(contentType).UpdateOriginsRequest(updateOriginsRequest).Execute()

/edge_applications/{edge_application_id}/origins/{origin_id}

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
    originKey := "originKey_example" // string | 
    accept := "application/json; version=3" // string |  (optional)
    contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
    updateOriginsRequest := *openapiclient.NewUpdateOriginsRequest("Name_example", []openapiclient.CreateOriginsRequestAddresses{*openapiclient.NewCreateOriginsRequestAddresses("Address_example")}, "HostHeader_example") // UpdateOriginsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPut(context.Background(), edgeApplicationId, originKey).Accept(accept).ContentType(contentType).UpdateOriginsRequest(updateOriginsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPut`: OriginsIdResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsOriginKeyPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**originKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdOriginsOriginKeyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **updateOriginsRequest** | [**UpdateOriginsRequest**](UpdateOriginsRequest.md) |  | 

### Return type

[**OriginsIdResponse**](OriginsIdResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdOriginsPost

> OriginsIdResponse EdgeApplicationsEdgeApplicationIdOriginsPost(ctx, edgeApplicationId).Accept(accept).ContentType(contentType).CreateOriginsRequest(createOriginsRequest).Execute()

/edge_applications/{edge_application_id}/origins

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
    createOriginsRequest := *openapiclient.NewCreateOriginsRequest("Name_example", []openapiclient.CreateOriginsRequestAddresses{*openapiclient.NewCreateOriginsRequestAddresses("Address_example")}, "HostHeader_example") // CreateOriginsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsPost(context.Background(), edgeApplicationId).Accept(accept).ContentType(contentType).CreateOriginsRequest(createOriginsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdOriginsPost`: OriginsIdResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsOriginsApi.EdgeApplicationsEdgeApplicationIdOriginsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdOriginsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **createOriginsRequest** | [**CreateOriginsRequest**](CreateOriginsRequest.md) |  | 

### Return type

[**OriginsIdResponse**](OriginsIdResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

