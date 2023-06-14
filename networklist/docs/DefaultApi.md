# \DefaultApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkListsGet**](DefaultApi.md#NetworkListsGet) | **Get** /network_lists | List all user Network Lists
[**NetworkListsPost**](DefaultApi.md#NetworkListsPost) | **Post** /network_lists | Create a Network Lists
[**NetworkListsUuidGet**](DefaultApi.md#NetworkListsUuidGet) | **Get** /network_lists/{uuid} | Retrieve a Network Lists set by uuid
[**NetworkListsUuidPut**](DefaultApi.md#NetworkListsUuidPut) | **Put** /network_lists/{uuid} | Overwrite some Network Lists attributes



## NetworkListsGet

> ListNetworkListsResponse NetworkListsGet(ctx).Page(page).Execute()

List all user Network Lists

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
    page := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.NetworkListsGet(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.NetworkListsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkListsGet`: ListNetworkListsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.NetworkListsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkListsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 

### Return type

[**ListNetworkListsResponse**](ListNetworkListsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkListsPost

> NetworkListsPost(ctx).CreateNetworkListsRequest(createNetworkListsRequest).Execute()

Create a Network Lists

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
    createNetworkListsRequest := *openapiclient.NewCreateNetworkListsRequest() // CreateNetworkListsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.NetworkListsPost(context.Background()).CreateNetworkListsRequest(createNetworkListsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.NetworkListsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkListsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkListsRequest** | [**CreateNetworkListsRequest**](CreateNetworkListsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkListsUuidGet

> NetworkListsResponse NetworkListsUuidGet(ctx, uuid).Execute()

Retrieve a Network Lists set by uuid

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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.NetworkListsUuidGet(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.NetworkListsUuidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkListsUuidGet`: NetworkListsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.NetworkListsUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkListsUuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkListsResponse**](NetworkListsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkListsUuidPut

> ListNetworkListsResponse NetworkListsUuidPut(ctx, uuid).UpdateNetworkListsRequest(updateNetworkListsRequest).Execute()

Overwrite some Network Lists attributes

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
    uuid := "uuid_example" // string | 
    updateNetworkListsRequest := *openapiclient.NewUpdateNetworkListsRequest() // UpdateNetworkListsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.NetworkListsUuidPut(context.Background(), uuid).UpdateNetworkListsRequest(updateNetworkListsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.NetworkListsUuidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkListsUuidPut`: ListNetworkListsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.NetworkListsUuidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkListsUuidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkListsRequest** | [**UpdateNetworkListsRequest**](UpdateNetworkListsRequest.md) |  | 

### Return type

[**ListNetworkListsResponse**](ListNetworkListsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

