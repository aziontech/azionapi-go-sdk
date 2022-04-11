# \RealTimePurgeApi

All URIs are relative to *http://localhost:3002*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PurgeCacheKey**](RealTimePurgeApi.md#PurgeCacheKey) | **Post** /purge/cachekey | /purge/cachekey
[**PurgeUrl**](RealTimePurgeApi.md#PurgeUrl) | **Post** /purge/url | /purge/url
[**PurgeWildcard**](RealTimePurgeApi.md#PurgeWildcard) | **Post** /purge/wildcard | /purge/wildcard



## PurgeCacheKey

> PurgeCacheKey(ctx).Accept(accept).ContentType(contentType).PurgeCacheKeyRequest(purgeCacheKeyRequest).Execute()

/purge/cachekey



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
    contentType := "application/json" // string |  (optional)
    purgeCacheKeyRequest := *openapiclient.NewPurgeCacheKeyRequest([]string{"Urls_example"}, "Method_example", "Layer_example") // PurgeCacheKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RealTimePurgeApi.PurgeCacheKey(context.Background()).Accept(accept).ContentType(contentType).PurgeCacheKeyRequest(purgeCacheKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimePurgeApi.PurgeCacheKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurgeCacheKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** |  | 
 **contentType** | **string** |  | 
 **purgeCacheKeyRequest** | [**PurgeCacheKeyRequest**](PurgeCacheKeyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurgeUrl

> PurgeUrl(ctx).Accept(accept).ContentType(contentType).PurgeUrlRequest(purgeUrlRequest).Execute()

/purge/url



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
    contentType := "application/json" // string |  (optional)
    purgeUrlRequest := *openapiclient.NewPurgeUrlRequest([]string{"Urls_example"}, "Method_example") // PurgeUrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RealTimePurgeApi.PurgeUrl(context.Background()).Accept(accept).ContentType(contentType).PurgeUrlRequest(purgeUrlRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimePurgeApi.PurgeUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurgeUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** |  | 
 **contentType** | **string** |  | 
 **purgeUrlRequest** | [**PurgeUrlRequest**](PurgeUrlRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurgeWildcard

> PurgeWildcard(ctx).Accept(accept).ContentType(contentType).PurgeWildcardRequest(purgeWildcardRequest).Execute()

/purge/wildcard



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
    contentType := "application/json" // string |  (optional)
    purgeWildcardRequest := *openapiclient.NewPurgeWildcardRequest([]string{"Urls_example"}, "Method_example") // PurgeWildcardRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RealTimePurgeApi.PurgeWildcard(context.Background()).Accept(accept).ContentType(contentType).PurgeWildcardRequest(purgeWildcardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RealTimePurgeApi.PurgeWildcard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurgeWildcardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** |  | 
 **contentType** | **string** |  | 
 **purgeWildcardRequest** | [**PurgeWildcardRequest**](PurgeWildcardRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

