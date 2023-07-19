# \DeleteDigitalCertificateApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemoveDigitalCertificates**](DeleteDigitalCertificateApi.md#RemoveDigitalCertificates) | **Delete** /digital_certificates/{digital_certificate_id} | Remove a digital certificate or CSR from your account



## RemoveDigitalCertificates

> RemoveDigitalCertificates(ctx, digitalCertificateId).Execute()

Remove a digital certificate or CSR from your account

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
    digitalCertificateId := int32(56) // int32 | ID of certificate to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeleteDigitalCertificateApi.RemoveDigitalCertificates(context.Background(), digitalCertificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeleteDigitalCertificateApi.RemoveDigitalCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalCertificateId** | **int32** | ID of certificate to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDigitalCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

