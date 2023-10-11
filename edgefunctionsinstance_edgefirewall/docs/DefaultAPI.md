# \DefaultAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdDelete**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdDelete) | **Delete** /edge_firewall/{edge_firewall_id}/functions_instances/{edge_function_instance_id} | Delete an Edge Functions Instance by uuid
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdGet**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdGet) | **Get** /edge_firewall/{edge_firewall_id}/functions_instances/{edge_function_instance_id} | Retrieve an Edge Functions Instance set by uuid
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPatch**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPatch) | **Patch** /edge_firewall/{edge_firewall_id}/functions_instances/{edge_function_instance_id} | Update some Edge Functions Instance attributes
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPut**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPut) | **Put** /edge_firewall/{edge_firewall_id}/functions_instances/{edge_function_instance_id} | Overwrite some Edge Functions Instance attributes
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesGet**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesGet) | **Get** /edge_firewall/{edge_firewall_id}/functions_instances | List all user Edge Functions Instances
[**EdgeFirewallEdgeFirewallIdFunctionsInstancesPost**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdFunctionsInstancesPost) | **Post** /edge_firewall/{edge_firewall_id}/functions_instances | Create an Edge Functions Instance



## EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdDelete

> EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdDelete(ctx, edgeFirewallId, edgeFunctionInstanceId).Execute()

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
    edgeFirewallId := int64(789) // int64 | 
    edgeFunctionInstanceId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdDelete(context.Background(), edgeFirewallId, edgeFunctionInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **int64** |  | 
**edgeFunctionInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdDeleteRequest struct via the builder pattern


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


## EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdGet

> EdgeFunctionsInstanceResponse EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdGet(ctx, edgeFirewallId, edgeFunctionInstanceId).Execute()

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
    edgeFirewallId := int64(789) // int64 | 
    edgeFunctionInstanceId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdGet(context.Background(), edgeFirewallId, edgeFunctionInstanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdGet`: EdgeFunctionsInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **int64** |  | 
**edgeFunctionInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdGetRequest struct via the builder pattern


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


## EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPatch

> EdgeFunctionsInstanceResponse EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPatch(ctx, edgeFirewallId, edgeFunctionInstanceId).Body(body).Execute()

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
    edgeFirewallId := int64(789) // int64 | 
    edgeFunctionInstanceId := int64(789) // int64 | 
    body := CreateEdgeFunctionsInstancesRequest(987) // CreateEdgeFunctionsInstancesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPatch(context.Background(), edgeFirewallId, edgeFunctionInstanceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPatch`: EdgeFunctionsInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **int64** |  | 
**edgeFunctionInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPatchRequest struct via the builder pattern


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


## EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPut

> EdgeFunctionsInstanceResponse EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPut(ctx, edgeFirewallId, edgeFunctionInstanceId).Body(body).Execute()

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
    edgeFirewallId := int64(789) // int64 | 
    edgeFunctionInstanceId := int64(789) // int64 | 
    body := CreateEdgeFunctionsInstancesRequest(987) // CreateEdgeFunctionsInstancesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPut(context.Background(), edgeFirewallId, edgeFunctionInstanceId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPut`: EdgeFunctionsInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **int64** |  | 
**edgeFunctionInstanceId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdFunctionsInstancesEdgeFunctionInstanceIdPutRequest struct via the builder pattern


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


## EdgeFirewallEdgeFirewallIdFunctionsInstancesGet

> ListEdgeFunctionsInstancesResponse EdgeFirewallEdgeFirewallIdFunctionsInstancesGet(ctx, edgeFirewallId).Page(page).PageSize(pageSize).Sort(sort).OrderBy(orderBy).Execute()

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
    edgeFirewallId := int64(789) // int64 | 
    page := int64(789) // int64 |  (optional)
    pageSize := int64(789) // int64 |  (optional)
    sort := "sort_example" // string |  (optional)
    orderBy := "orderBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesGet(context.Background(), edgeFirewallId).Page(page).PageSize(pageSize).Sort(sort).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesGet`: ListEdgeFunctionsInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **int64** |  | 

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

> EdgeFunctionsInstanceResponse EdgeFirewallEdgeFirewallIdFunctionsInstancesPost(ctx, edgeFirewallId).CreateEdgeFunctionsInstancesRequest(createEdgeFunctionsInstancesRequest).Execute()

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
    edgeFirewallId := int64(789) // int64 | 
    createEdgeFunctionsInstancesRequest := *openapiclient.NewCreateEdgeFunctionsInstancesRequest() // CreateEdgeFunctionsInstancesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesPost(context.Background(), edgeFirewallId).CreateEdgeFunctionsInstancesRequest(createEdgeFunctionsInstancesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdFunctionsInstancesPost`: EdgeFunctionsInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdFunctionsInstancesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **int64** |  | 

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

