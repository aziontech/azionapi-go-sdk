# \ZonesApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteZone**](ZonesApi.md#DeleteZone) | **Delete** /intelligent_dns/{zone_id} | Remove an Intelligent DNS hosted zone
[**GetZone**](ZonesApi.md#GetZone) | **Get** /intelligent_dns/{zone_id} | Get an Intelligent DNS hosted zone
[**GetZones**](ZonesApi.md#GetZones) | **Get** /intelligent_dns | Get a collection of Intelligent DNS zones
[**PostZone**](ZonesApi.md#PostZone) | **Post** /intelligent_dns | Add a new Intelligent DNS zone
[**PutZone**](ZonesApi.md#PutZone) | **Put** /intelligent_dns/{zone_id} | Update an Intelligent DNS hosted zone



## DeleteZone

> string DeleteZone(ctx, zoneId).Execute()

Remove an Intelligent DNS hosted zone

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
    zoneId := int32(56) // int32 | The hosted zone id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesApi.DeleteZone(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.DeleteZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteZone`: string
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.DeleteZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int32** | The hosted zone id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZone

> GetZoneResponse GetZone(ctx, zoneId).Execute()

Get an Intelligent DNS hosted zone

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
    zoneId := int32(56) // int32 | The hosted zone id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesApi.GetZone(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.GetZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZone`: GetZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.GetZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int32** | The hosted zone id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetZoneResponse**](GetZoneResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZones

> GetZonesResponse GetZones(ctx).Execute()

Get a collection of Intelligent DNS zones

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesApi.GetZones(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.GetZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZones`: GetZonesResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.GetZones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetZonesRequest struct via the builder pattern


### Return type

[**GetZonesResponse**](GetZonesResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostZone

> PostOrPutZoneResponse PostZone(ctx).Zone(zone).Execute()

Add a new Intelligent DNS zone

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
    zone := *openapiclient.NewZone() // Zone |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesApi.PostZone(context.Background()).Zone(zone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.PostZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostZone`: PostOrPutZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.PostZone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zone** | [**Zone**](Zone.md) |  | 

### Return type

[**PostOrPutZoneResponse**](PostOrPutZoneResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutZone

> PostOrPutZoneResponse PutZone(ctx, zoneId).Zone(zone).Execute()

Update an Intelligent DNS hosted zone

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
    zoneId := int32(56) // int32 | The hosted zone id
    zone := *openapiclient.NewZone() // Zone |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesApi.PutZone(context.Background(), zoneId).Zone(zone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesApi.PutZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutZone`: PostOrPutZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesApi.PutZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int32** | The hosted zone id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zone** | [**Zone**](Zone.md) |  | 

### Return type

[**PostOrPutZoneResponse**](PostOrPutZoneResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

