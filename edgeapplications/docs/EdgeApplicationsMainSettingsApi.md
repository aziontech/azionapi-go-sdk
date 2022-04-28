# \EdgeApplicationsMainSettingsApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeApplicationsGet**](EdgeApplicationsMainSettingsApi.md#EdgeApplicationsGet) | **Get** /edge_applications | /edge_applications
[**EdgeApplicationsIdDelete**](EdgeApplicationsMainSettingsApi.md#EdgeApplicationsIdDelete) | **Delete** /edge_applications/{id} | /edge_applications/:id
[**EdgeApplicationsIdGet**](EdgeApplicationsMainSettingsApi.md#EdgeApplicationsIdGet) | **Get** /edge_applications/{id} | /edge_applications/:id
[**EdgeApplicationsIdPatch**](EdgeApplicationsMainSettingsApi.md#EdgeApplicationsIdPatch) | **Patch** /edge_applications/{id} | /edge_applications/:id
[**EdgeApplicationsIdPut**](EdgeApplicationsMainSettingsApi.md#EdgeApplicationsIdPut) | **Put** /edge_applications/{id} | /edge_applications/:id
[**EdgeApplicationsPost**](EdgeApplicationsMainSettingsApi.md#EdgeApplicationsPost) | **Post** /edge_applications | /edge_applications



## EdgeApplicationsGet

> GetApplicationsResponse EdgeApplicationsGet(ctx).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()

/edge_applications

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
    page := int64(789) // int64 |  (optional)
    pageSize := int64(789) // int64 |  (optional)
    filter := "filter_example" // string |  (optional)
    orderBy := "orderBy_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    accept := "application/json; version=3" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EdgeApplicationsMainSettingsApi.EdgeApplicationsGet(context.Background()).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsMainSettingsApi.EdgeApplicationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsGet`: GetApplicationsResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsMainSettingsApi.EdgeApplicationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | 
 **pageSize** | **int64** |  | 
 **filter** | **string** |  | 
 **orderBy** | **string** |  | 
 **sort** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**GetApplicationsResponse**](GetApplicationsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsIdDelete

> EdgeApplicationsIdDelete(ctx, id).Accept(accept).Execute()

/edge_applications/:id

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
    id := "id_example" // string | The id of the edge application that you plan to delete.
    accept := "application/json; version=3" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EdgeApplicationsMainSettingsApi.EdgeApplicationsIdDelete(context.Background(), id).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsMainSettingsApi.EdgeApplicationsIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the edge application that you plan to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 

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


## EdgeApplicationsIdGet

> GetApplicationResponse EdgeApplicationsIdGet(ctx, id).Accept(accept).Execute()

/edge_applications/:id

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
    id := "id_example" // string | 
    accept := "application/json; version=3" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EdgeApplicationsMainSettingsApi.EdgeApplicationsIdGet(context.Background(), id).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsMainSettingsApi.EdgeApplicationsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsIdGet`: GetApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsMainSettingsApi.EdgeApplicationsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 

### Return type

[**GetApplicationResponse**](GetApplicationResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsIdPatch

> ApplicationUpdateResponse EdgeApplicationsIdPatch(ctx, id).Accept(accept).ContentType(contentType).ApplicationUpdateRequest(applicationUpdateRequest).Execute()

/edge_applications/:id

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
    id := "id_example" // string | 
    accept := "application/json; version=3" // string |  (optional)
    contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
    applicationUpdateRequest := *openapiclient.NewApplicationUpdateRequest() // ApplicationUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EdgeApplicationsMainSettingsApi.EdgeApplicationsIdPatch(context.Background(), id).Accept(accept).ContentType(contentType).ApplicationUpdateRequest(applicationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsMainSettingsApi.EdgeApplicationsIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsIdPatch`: ApplicationUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsMainSettingsApi.EdgeApplicationsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **applicationUpdateRequest** | [**ApplicationUpdateRequest**](ApplicationUpdateRequest.md) |  | 

### Return type

[**ApplicationUpdateResponse**](ApplicationUpdateResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsIdPut

> ApplicationPutResult EdgeApplicationsIdPut(ctx, id).Accept(accept).ContentType(contentType).ApplicationPutRequest(applicationPutRequest).Execute()

/edge_applications/:id

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
    id := "id_example" // string | The Id of the edge application to be overwritten. 
    accept := "application/json; version=3" // string |  (optional)
    contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
    applicationPutRequest := *openapiclient.NewApplicationPutRequest("Name_example") // ApplicationPutRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EdgeApplicationsMainSettingsApi.EdgeApplicationsIdPut(context.Background(), id).Accept(accept).ContentType(contentType).ApplicationPutRequest(applicationPutRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsMainSettingsApi.EdgeApplicationsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsIdPut`: ApplicationPutResult
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsMainSettingsApi.EdgeApplicationsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Id of the edge application to be overwritten.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **applicationPutRequest** | [**ApplicationPutRequest**](ApplicationPutRequest.md) |  | 

### Return type

[**ApplicationPutResult**](ApplicationPutResult.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsPost

> CreateApplicationResult EdgeApplicationsPost(ctx).Accept(accept).ContentType(contentType).ERRORUNKNOWN(eRRORUNKNOWN).CreateApplicationRequest(createApplicationRequest).Execute()

/edge_applications

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
    accept := "application/json; version=3" // string |  (optional)
    contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
    eRRORUNKNOWN := "eRRORUNKNOWN_example" // string |  (optional)
    createApplicationRequest := *openapiclient.NewCreateApplicationRequest("Name_example") // CreateApplicationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EdgeApplicationsMainSettingsApi.EdgeApplicationsPost(context.Background()).Accept(accept).ContentType(contentType).ERRORUNKNOWN(eRRORUNKNOWN).CreateApplicationRequest(createApplicationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsMainSettingsApi.EdgeApplicationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsPost`: CreateApplicationResult
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsMainSettingsApi.EdgeApplicationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **eRRORUNKNOWN** | **string** |  | 
 **createApplicationRequest** | [**CreateApplicationRequest**](CreateApplicationRequest.md) |  | 

### Return type

[**CreateApplicationResult**](CreateApplicationResult.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

