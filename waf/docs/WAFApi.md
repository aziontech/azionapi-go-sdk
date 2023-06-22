# \WAFApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWAFDomains**](WAFApi.md#GetWAFDomains) | **Get** /waf/{wafId}/domains | Find domains attached to a WAF
[**GetWAFEvents**](WAFApi.md#GetWAFEvents) | **Get** /waf/{wafId}/waf_events | Find WAF log events



## GetWAFDomains

> WAFDomains200 GetWAFDomains(ctx, wafId).Name(name).Execute()

Find domains attached to a WAF

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
    wafId := int64(789) // int64 | ID of WAF to return
    name := "name_example" // string | searches WAF for name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.GetWAFDomains(context.Background(), wafId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.GetWAFDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWAFDomains`: WAFDomains200
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.GetWAFDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **int64** | ID of WAF to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWAFDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | searches WAF for name | 

### Return type

[**WAFDomains200**](WAFDomains200.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWAFEvents

> WAFEvents200 GetWAFEvents(ctx, wafId).HourRange(hourRange).DomainsIds(domainsIds).NetworkListId(networkListId).Execute()

Find WAF log events

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
    wafId := int64(789) // int64 | ID of WAF to return
    hourRange := int64(789) // int64 | Last log hours since now (it must be a integer number ranging between 1 and 72)
    domainsIds := "domainsIds_example" // string | Multiple domain's id (they must be separated by comma like 1233,1234)
    networkListId := int64(789) // int64 | Id of a network list (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFApi.GetWAFEvents(context.Background(), wafId).HourRange(hourRange).DomainsIds(domainsIds).NetworkListId(networkListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFApi.GetWAFEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWAFEvents`: WAFEvents200
    fmt.Fprintf(os.Stdout, "Response from `WAFApi.GetWAFEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **int64** | ID of WAF to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWAFEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hourRange** | **int64** | Last log hours since now (it must be a integer number ranging between 1 and 72) | 
 **domainsIds** | **string** | Multiple domain&#39;s id (they must be separated by comma like 1233,1234) | 
 **networkListId** | **int64** | Id of a network list | 

### Return type

[**WAFEvents200**](WAFEvents200.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

