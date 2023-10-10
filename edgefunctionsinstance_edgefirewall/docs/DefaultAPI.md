# \DefaultAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesGet**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesGet) | **Get** /edge_firewall/:edge_firewall_id:/functions_instances | List all user Edge Functions Instances
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesPost**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesPost) | **Post** /edge_firewall/:edge_firewall_id:/functions_instances | Create an Edge Functions Instance
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidDelete**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidDelete) | **Delete** /edge_firewall/:edge_firewall_id:/functions_instances/{uuid} | Delete an Edge Functions Instance by uuid
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet) | **Get** /edge_firewall/:edge_firewall_id:/functions_instances/{uuid} | Retrieve an Edge Functions Instance set by uuid
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch) | **Patch** /edge_firewall/:edge_firewall_id:/functions_instances/{uuid} | Update some Edge Functions Instance attributes
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut) | **Put** /edge_firewall/:edge_firewall_id:/functions_instances/{uuid} | Overwrite some Edge Functions Instance attributes



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
    page := int64(789) // int64 |  (optional)
    pageSize := int64(789) // int64 |  (optional)
    sort := "sort_example" // string |  (optional)
    orderBy := "orderBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesGet`: ListEdgeFunctionsInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdFunctionsInstancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | 
 **pageSize** | **int64** |  | 
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

> EdgeFunctionsInstanceResponse EdgeFirewallEdgeFirewallIdFunctionsInstancesPost(ctx).CreateEdgeFunctionsInstancesRequest(createEdgeFunctionsInstancesRequest).Execute()

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
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesPost(context.Background()).CreateEdgeFunctionsInstancesRequest(createEdgeFunctionsInstancesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesPost`: EdgeFunctionsInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdFunctionsInstancesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEdgeFunctionsInstancesRequest** | [**CreateEdgeFunctionsInstancesRequest**](CreateEdgeFunctionsInstancesRequest.md) |  | 

### Return type

[**EdgeFunctionsInstanceResponse**](EdgeFunctionsInstanceResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
    r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidDelete(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidDelete``: %v\n", err)
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
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet`: EdgeFunctionsInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidGet`: %v\n", resp)
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

> EdgeFunctionsInstanceResponse EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch(ctx, uuid).Body(body).Execute()

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
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch(context.Background(), uuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch`: EdgeFunctionsInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPatch`: %v\n", resp)
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

[**EdgeFunctionsInstanceResponse**](EdgeFunctionsInstanceResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut

> EdgeFunctionsInstanceResponse EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut(ctx, uuid).Body(body).Execute()

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
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut(context.Background(), uuid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut`: EdgeFunctionsInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesUuidPut`: %v\n", resp)
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

[**EdgeFunctionsInstanceResponse**](EdgeFunctionsInstanceResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

