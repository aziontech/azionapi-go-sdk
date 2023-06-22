# \DefaultApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeFirewallGet**](DefaultApi.md#EdgeFirewallGet) | **Get** /edge_firewall | List all user edge firewall
[**EdgeFirewallPost**](DefaultApi.md#EdgeFirewallPost) | **Post** /edge_firewall | Create a edge firewall
[**EdgeFirewallUuidDelete**](DefaultApi.md#EdgeFirewallUuidDelete) | **Delete** /edge_firewall/{uuid} | Delete an edge firewall by uuid
[**EdgeFirewallUuidGet**](DefaultApi.md#EdgeFirewallUuidGet) | **Get** /edge_firewall/{uuid} | Retrieve an edge firewall set by uuid
[**EdgeFirewallUuidPatch**](DefaultApi.md#EdgeFirewallUuidPatch) | **Patch** /edge_firewall/{uuid} | Update some edge firewall attributes, like \&quot;active\&quot;
[**EdgeFirewallUuidPut**](DefaultApi.md#EdgeFirewallUuidPut) | **Put** /edge_firewall/{uuid} | Overwrite some edge firewall attributes, like \&quot;active\&quot;



## EdgeFirewallGet

> ListEdgeFirewallResponse EdgeFirewallGet(ctx).Page(page).PageSize(pageSize).Sort(sort).OrderBy(orderBy).Execute()

List all user edge firewall

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
    pageSize := int32(56) // int32 |  (optional)
    sort := "sort_example" // string |  (optional)
    orderBy := "orderBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EdgeFirewallGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EdgeFirewallGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallGet`: ListEdgeFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EdgeFirewallGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **sort** | **string** |  | 
 **orderBy** | **string** |  | 

### Return type

[**ListEdgeFirewallResponse**](ListEdgeFirewallResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallPost

> EdgeFirewallPost(ctx).CreateEdgeFirewallRequest(createEdgeFirewallRequest).Execute()

Create a edge firewall

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
    createEdgeFirewallRequest := *openapiclient.NewCreateEdgeFirewallRequest() // CreateEdgeFirewallRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.EdgeFirewallPost(context.Background()).CreateEdgeFirewallRequest(createEdgeFirewallRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EdgeFirewallPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEdgeFirewallRequest** | [**CreateEdgeFirewallRequest**](CreateEdgeFirewallRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallUuidDelete

> EdgeFirewallUuidDelete(ctx, uuid).Execute()

Delete an edge firewall by uuid

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
    r, err := apiClient.DefaultApi.EdgeFirewallUuidDelete(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EdgeFirewallUuidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallUuidDeleteRequest struct via the builder pattern


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


## EdgeFirewallUuidGet

> EdgeFirewallResponse EdgeFirewallUuidGet(ctx, uuid).Execute()

Retrieve an edge firewall set by uuid

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
    resp, r, err := apiClient.DefaultApi.EdgeFirewallUuidGet(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EdgeFirewallUuidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallUuidGet`: EdgeFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EdgeFirewallUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallUuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EdgeFirewallResponse**](EdgeFirewallResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallUuidPatch

> ListEdgeFirewallResponse EdgeFirewallUuidPatch(ctx, uuid).Body(body).Execute()

Update some edge firewall attributes, like \"active\"

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
    body := ListEdgeFirewallResponse(987) // ListEdgeFirewallResponse | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EdgeFirewallUuidPatch(context.Background(), uuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EdgeFirewallUuidPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallUuidPatch`: ListEdgeFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EdgeFirewallUuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallUuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **ListEdgeFirewallResponse** |  | 

### Return type

[**ListEdgeFirewallResponse**](ListEdgeFirewallResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallUuidPut

> ListEdgeFirewallResponse EdgeFirewallUuidPut(ctx, uuid).Body(body).Execute()

Overwrite some edge firewall attributes, like \"active\"

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
    body := ListEdgeFirewallResponse(987) // ListEdgeFirewallResponse | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EdgeFirewallUuidPut(context.Background(), uuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EdgeFirewallUuidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallUuidPut`: ListEdgeFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EdgeFirewallUuidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallUuidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **ListEdgeFirewallResponse** |  | 

### Return type

[**ListEdgeFirewallResponse**](ListEdgeFirewallResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

