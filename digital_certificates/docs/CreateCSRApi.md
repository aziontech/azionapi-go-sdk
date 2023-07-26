# \CreateCSRApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCSR**](CreateCSRApi.md#CreateCSR) | **Post** /digital_certificates/csr | Create an encrypted Certificate Request with Azion, which can then be sent for signing to a CA



## CreateCSR

> DC200 CreateCSR(ctx).CreateCSRRequest(createCSRRequest).Execute()

Create an encrypted Certificate Request with Azion, which can then be sent for signing to a CA

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
    createCSRRequest := *openapiclient.NewCreateCSRRequest() // CreateCSRRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreateCSRApi.CreateCSR(context.Background()).CreateCSRRequest(createCSRRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreateCSRApi.CreateCSR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCSR`: DC200
    fmt.Fprintf(os.Stdout, "Response from `CreateCSRApi.CreateCSR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCSRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCSRRequest** | [**CreateCSRRequest**](CreateCSRRequest.md) |  | 

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

