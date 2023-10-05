# \ZonesAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteZone**](ZonesAPI.md#DeleteZone) | **Delete** /intelligent_dns/{zone_id} | Remove an Intelligent DNS hosted zone
[**GetZone**](ZonesAPI.md#GetZone) | **Get** /intelligent_dns/{zone_id} | Get an Intelligent DNS hosted zone
[**GetZones**](ZonesAPI.md#GetZones) | **Get** /intelligent_dns | Get a collection of Intelligent DNS zones
[**PostZone**](ZonesAPI.md#PostZone) | **Post** /intelligent_dns | Add a new Intelligent DNS zone
[**PutZone**](ZonesAPI.md#PutZone) | **Put** /intelligent_dns/{zone_id} | Update an Intelligent DNS hosted zone



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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    zoneId := int32(56) // int32 | The hosted zone id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesAPI.DeleteZone(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.DeleteZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteZone`: string
    fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.DeleteZone`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    zoneId := int32(56) // int32 | The hosted zone id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesAPI.GetZone(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.GetZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZone`: GetZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.GetZone`: %v\n", resp)
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

> GetZonesResponse GetZones(ctx).OrderBy(orderBy).Sort(sort).Page(page).PageSize(pageSize).Execute()

Get a collection of Intelligent DNS zones

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
    orderBy := "orderBy_example" // string | Identifies which property the return should be sorted by. (optional) (default to "name")
    sort := "sort_example" // string | Defines whether objects are shown in ascending or descending order depending on the value set in order_by. (optional) (default to "asc")
    page := int64(789) // int64 | Identifies which page should be returned, if the return is paginated. (optional) (default to 1)
    pageSize := int64(789) // int64 | Identifies how many items should be returned per page. (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesAPI.GetZones(context.Background()).OrderBy(orderBy).Sort(sort).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.GetZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZones`: GetZonesResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.GetZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **string** | Identifies which property the return should be sorted by. | [default to &quot;name&quot;]
 **sort** | **string** | Defines whether objects are shown in ascending or descending order depending on the value set in order_by. | [default to &quot;asc&quot;]
 **page** | **int64** | Identifies which page should be returned, if the return is paginated. | [default to 1]
 **pageSize** | **int64** | Identifies how many items should be returned per page. | [default to 10]

### Return type

[**GetZonesResponse**](GetZonesResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; version=3

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    zone := *openapiclient.NewZone() // Zone |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesAPI.PostZone(context.Background()).Zone(zone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.PostZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostZone`: PostOrPutZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.PostZone`: %v\n", resp)
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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    zoneId := int32(56) // int32 | The hosted zone id
    zone := *openapiclient.NewZone() // Zone |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ZonesAPI.PutZone(context.Background(), zoneId).Zone(zone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.PutZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutZone`: PostOrPutZoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.PutZone`: %v\n", resp)
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

