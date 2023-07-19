# \OverwriteDigitalCertificateApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OverwriteDigitalCertificate**](OverwriteDigitalCertificateApi.md#OverwriteDigitalCertificate) | **Put** /digital_certificates/{digital_certificate_id} | Overwrite a digital certificate with another complete digital certificate



## OverwriteDigitalCertificate

> DC200 OverwriteDigitalCertificate(ctx, digitalCertificateId).CreateCertificateRequest(createCertificateRequest).Execute()

Overwrite a digital certificate with another complete digital certificate

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
    digitalCertificateId := int32(56) // int32 | ID of certificate to overwrite
    createCertificateRequest := *openapiclient.NewCreateCertificateRequest("Name_example", "Certificate_example", "PrivateKey_example") // CreateCertificateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverwriteDigitalCertificateApi.OverwriteDigitalCertificate(context.Background(), digitalCertificateId).CreateCertificateRequest(createCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverwriteDigitalCertificateApi.OverwriteDigitalCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverwriteDigitalCertificate`: DC200
    fmt.Fprintf(os.Stdout, "Response from `OverwriteDigitalCertificateApi.OverwriteDigitalCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalCertificateId** | **int32** | ID of certificate to overwrite | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverwriteDigitalCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCertificateRequest** | [**CreateCertificateRequest**](CreateCertificateRequest.md) |  | 

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

