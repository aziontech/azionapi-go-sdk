# \RetrieveDigitalCertificateListApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDigitalCertificates**](RetrieveDigitalCertificateListApi.md#ListDigitalCertificates) | **Get** /digital_certificates | List all existing digital certificates and CSRs registered to your account



## ListDigitalCertificates

> DC200List ListDigitalCertificates(ctx).OrderBy(orderBy).Sort(sort).Name(name).Search(search).Execute()

List all existing digital certificates and CSRs registered to your account

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
    orderBy := "orderBy_example" // string | Response field to order by, current options are \"name\" or \"id\" (optional) (default to "id")
    sort := "sort_example" // string | Sorting direction (optional) (default to "asc")
    name := "name_example" // string | Searches certificate for name (optional)
    search := "search_example" // string | Searches for string in certificate name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetrieveDigitalCertificateListApi.ListDigitalCertificates(context.Background()).OrderBy(orderBy).Sort(sort).Name(name).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetrieveDigitalCertificateListApi.ListDigitalCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDigitalCertificates`: DC200List
    fmt.Fprintf(os.Stdout, "Response from `RetrieveDigitalCertificateListApi.ListDigitalCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDigitalCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **string** | Response field to order by, current options are \&quot;name\&quot; or \&quot;id\&quot; | [default to &quot;id&quot;]
 **sort** | **string** | Sorting direction | [default to &quot;asc&quot;]
 **name** | **string** | Searches certificate for name | 
 **search** | **string** | Searches for string in certificate name | 

### Return type

[**DC200List**](DC200List.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

