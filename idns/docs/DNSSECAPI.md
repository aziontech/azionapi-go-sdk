# \DNSSECAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetZoneDnsSec**](DNSSECAPI.md#GetZoneDnsSec) | **Get** /intelligent_dns/{zone_id}/dnssec | Retrieve the DNSSEC zone status
[**PutZoneDnsSec**](DNSSECAPI.md#PutZoneDnsSec) | **Patch** /intelligent_dns/{zone_id}/dnssec | Update the DNSSEC zone



## GetZoneDnsSec

> GetOrPatchDnsSecResponse GetZoneDnsSec(ctx, zoneId).Execute()

Retrieve the DNSSEC zone status

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
    resp, r, err := apiClient.DNSSECAPI.GetZoneDnsSec(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSSECAPI.GetZoneDnsSec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZoneDnsSec`: GetOrPatchDnsSecResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSSECAPI.GetZoneDnsSec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int32** | The hosted zone id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneDnsSecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrPatchDnsSecResponse**](GetOrPatchDnsSecResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutZoneDnsSec

> GetOrPatchDnsSecResponse PutZoneDnsSec(ctx, zoneId).DnsSec(dnsSec).Execute()

Update the DNSSEC zone

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
    dnsSec := *openapiclient.NewDnsSec() // DnsSec |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DNSSECAPI.PutZoneDnsSec(context.Background(), zoneId).DnsSec(dnsSec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DNSSECAPI.PutZoneDnsSec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutZoneDnsSec`: GetOrPatchDnsSecResponse
    fmt.Fprintf(os.Stdout, "Response from `DNSSECAPI.PutZoneDnsSec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int32** | The hosted zone id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutZoneDnsSecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dnsSec** | [**DnsSec**](DnsSec.md) |  | 

### Return type

[**GetOrPatchDnsSecResponse**](GetOrPatchDnsSecResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

