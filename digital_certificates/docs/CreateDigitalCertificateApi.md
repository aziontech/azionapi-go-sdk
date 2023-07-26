# \CreateDigitalCertificateApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificate**](CreateDigitalCertificateApi.md#CreateCertificate) | **Post** /digital_certificates | Create a new digital certificate



## CreateCertificate

> DC200 CreateCertificate(ctx).CreateCertificateRequest(createCertificateRequest).Execute()

Create a new digital certificate

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
    createCertificateRequest := *openapiclient.NewCreateCertificateRequest("Name_example", "Certificate_example", "PrivateKey_example") // CreateCertificateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreateDigitalCertificateApi.CreateCertificate(context.Background()).CreateCertificateRequest(createCertificateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateDigitalCertificateApi.CreateCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificate`: DC200
    fmt.Fprintf(os.Stdout, "Response from `CreateDigitalCertificateApi.CreateCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateRequest struct via the builder pattern


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

