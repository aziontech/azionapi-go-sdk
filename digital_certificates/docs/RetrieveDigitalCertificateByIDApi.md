# \RetrieveDigitalCertificateByIDApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCertificate**](RetrieveDigitalCertificateByIDApi.md#GetCertificate) | **Get** /digital_certificates/{digital_certificate_id} | Get more data on a specific digital certificate or CSR.



## GetCertificate

> DC200 GetCertificate(ctx, digitalCertificateId).Execute()

Get more data on a specific digital certificate or CSR.

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
    digitalCertificateId := int64(789) // int64 | ID of certificate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetrieveDigitalCertificateByIDApi.GetCertificate(context.Background(), digitalCertificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetrieveDigitalCertificateByIDApi.GetCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificate`: DC200
    fmt.Fprintf(os.Stdout, "Response from `RetrieveDigitalCertificateByIDApi.GetCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalCertificateId** | **int64** | ID of certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DC200**](DC200.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

