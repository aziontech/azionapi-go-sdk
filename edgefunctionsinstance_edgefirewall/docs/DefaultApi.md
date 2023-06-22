# \DefaultApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesGet**](DefaultApi.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesGet) | **Get** /edge_firewall/:edge_firewall_id:/functions_instances | List all user Edge Functions Instances
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesPost**](DefaultApi.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesPost) | **Post** /edge_firewall/:edge_firewall_id:/functions_instances | Create an Edge Functions Instance
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidDelete**](DefaultApi.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidDelete) | **Delete** /edge_firewall/:edge_firewall_id:/functions_instances/{uuid} | Delete an Edge Functions Instance by uuid
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet**](DefaultApi.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet) | **Get** /edge_firewall/:edge_firewall_id:/functions_instances/{uuid} | Retrieve an Edge Functions Instance set by uuid
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch**](DefaultApi.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch) | **Patch** /edge_firewall/:edge_firewall_id:/functions_instances/{uuid} | Update some Edge Functions Instance attributes
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut**](DefaultApi.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut) | **Put** /edge_firewall/:edge_firewall_id:/functions_instances/{uuid} | Overwrite some Edge Functions Instance attributes



## EdgeFirewallEdgeFirewallIdFunctionsInstancesGet

> ListEdgeFunctionsInstancesResponse EdgeFirewallEdgeFirewallIdFunctionsInstancesGet(ctx).Page(page).PageSize(pageSize).Sort(sort).OrderBy(orderBy).Execute()

List all user Edge Functions Instances

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
    resp, r, err := apiClient.DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesGet`: ListEdgeFunctionsInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdFunctionsInstancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **sort** | **string** |  | 
 **orderBy** | **string** |  | 

### Return type

[**ListEdgeFunctionsInstancesResponse**](ListEdgeFunctionsInstancesResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallEdgeFirewallIdFunctionsInstancesPost

> EdgeFirewallEdgeFirewallIdFunctionsInstancesPost(ctx).CreateEdgeFunctionsInstancesRequest(createEdgeFunctionsInstancesRequest).Execute()

Create an Edge Functions Instance

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
    createEdgeFunctionsInstancesRequest := *openapiclient.NewCreateEdgeFunctionsInstancesRequest() // CreateEdgeFunctionsInstancesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesPost(context.Background()).CreateEdgeFunctionsInstancesRequest(createEdgeFunctionsInstancesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdFunctionsInstancesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEdgeFunctionsInstancesRequest** | [**CreateEdgeFunctionsInstancesRequest**](CreateEdgeFunctionsInstancesRequest.md) |  | 

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


## EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidDelete

> EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidDelete(ctx, uuid).Execute()

Delete an Edge Functions Instance by uuid

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
    r, err := apiClient.DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidDelete(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdFunctionsInstancesUuidDeleteRequest struct via the builder pattern


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


## EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet

> EdgeFunctionsInstanceResponse EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet(ctx, uuid).Execute()

Retrieve an Edge Functions Instance set by uuid

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
    resp, r, err := apiClient.DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet`: EdgeFunctionsInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EdgeFunctionsInstanceResponse**](EdgeFunctionsInstanceResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch

> ListEdgeFunctionsInstancesResponse EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch(ctx, uuid).Body(body).Execute()

Update some Edge Functions Instance attributes

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
    body := CreateEdgeFunctionsInstancesRequest(987) // CreateEdgeFunctionsInstancesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch(context.Background(), uuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch`: ListEdgeFunctionsInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **CreateEdgeFunctionsInstancesRequest** |  | 

### Return type

[**ListEdgeFunctionsInstancesResponse**](ListEdgeFunctionsInstancesResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut

> ListEdgeFunctionsInstancesResponse EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut(ctx, uuid).Body(body).Execute()

Overwrite some Edge Functions Instance attributes

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
    body := CreateEdgeFunctionsInstancesRequest(987) // CreateEdgeFunctionsInstancesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut(context.Background(), uuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut`: ListEdgeFunctionsInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **CreateEdgeFunctionsInstancesRequest** |  | 

### Return type

[**ListEdgeFunctionsInstancesResponse**](ListEdgeFunctionsInstancesResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

