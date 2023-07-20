# \UpdateDigitalCertificateApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateDigitalCertificate**](UpdateDigitalCertificateApi.md#UpdateDigitalCertificate) | **Patch** /digital_certificates/{digital_certificate_id} | Change only select settings of your digital certificate or CSR.



## UpdateDigitalCertificate

> DC200 UpdateDigitalCertificate(ctx, digitalCertificateId).UpdateDigitalCertificateRequest(updateDigitalCertificateRequest).Execute()

Change only select settings of your digital certificate or CSR.

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
    digitalCertificateId := int32(56) // int32 | ID of certificate to update
    updateDigitalCertificateRequest := *openapiclient.NewUpdateDigitalCertificateRequest() // UpdateDigitalCertificateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdateDigitalCertificateApi.UpdateDigitalCertificate(context.Background(), digitalCertificateId).UpdateDigitalCertificateRequest(updateDigitalCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateDigitalCertificateApi.UpdateDigitalCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDigitalCertificate`: DC200
    fmt.Fprintf(os.Stdout, "Response from `UpdateDigitalCertificateApi.UpdateDigitalCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalCertificateId** | **int32** | ID of certificate to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDigitalCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDigitalCertificateRequest** | [**UpdateDigitalCertificateRequest**](UpdateDigitalCertificateRequest.md) |  | 

### Return type

[**DC200**](DC200.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

