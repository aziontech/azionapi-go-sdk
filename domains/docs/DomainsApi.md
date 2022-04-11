# \DomainsApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomain**](DomainsApi.md#CreateDomain) | **Post** /domains | /domains
[**DelDomain**](DomainsApi.md#DelDomain) | **Delete** /domains/{id} | /domains/:id
[**GetDomain**](DomainsApi.md#GetDomain) | **Get** /domains/{id} | /domains/:id
[**GetDomains**](DomainsApi.md#GetDomains) | **Get** /domains | /domains
[**PutDomain**](DomainsApi.md#PutDomain) | **Put** /domains/{domain_id} | /domains:/:domain_id
[**UpdateDomain**](DomainsApi.md#UpdateDomain) | **Patch** /domains/{domain_id} | /domains/:domain_id



## CreateDomain

> DomainResponseWithResult CreateDomain(ctx).Accept(accept).Authorization(authorization).ContentType(contentType).CreateDomainRequest(createDomainRequest).Execute()

/domains



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accept := "application/json; version=3" // string |  (optional)
    authorization := "Token [TOKEN VALUE]" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    createDomainRequest := *openapiclient.NewCreateDomainRequest([]string{"Cnames_example"}, false, "Name_example", false, int64(123), NullableInt64(123)) // CreateDomainRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.CreateDomain(context.Background()).Accept(accept).Authorization(authorization).ContentType(contentType).CreateDomainRequest(createDomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.CreateDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDomain`: DomainResponseWithResult
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.CreateDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** |  | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **createDomainRequest** | [**CreateDomainRequest**](CreateDomainRequest.md) |  | 

### Return type

[**DomainResponseWithResult**](DomainResponseWithResult.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DelDomain

> DelDomain(ctx, id).Accept(accept).Authorization(authorization).Execute()

/domains/:id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of the domain to be deleted. 
    accept := "application/json; version=3" // string |  (optional)
    authorization := "Token [TOKEN VALUE]" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.DelDomain(context.Background(), id).Accept(accept).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.DelDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the domain to be deleted.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomain

> DomainResponseWithResult GetDomain(ctx, id).Accept(accept).Authorization(authorization).Execute()

/domains/:id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of the domain to be consulted. 
    accept := "application/json; version=3" // string |  (optional)
    authorization := "Token [TOKEN VALUE]" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.GetDomain(context.Background(), id).Accept(accept).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.GetDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomain`: DomainResponseWithResult
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.GetDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the domain to be consulted.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**DomainResponseWithResult**](DomainResponseWithResult.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomains

> DomainResponseWithResults GetDomains(ctx).Accept(accept).Authorization(authorization).Execute()

/domains



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accept := "application/json; version=3" // string |  (optional)
    authorization := "Token [TOKEN VALUE]" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.GetDomains(context.Background()).Accept(accept).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.GetDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomains`: DomainResponseWithResults
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.GetDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** |  | 
 **authorization** | **string** |  | 

### Return type

[**DomainResponseWithResults**](DomainResponseWithResults.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDomain

> DomainResponseWithResult PutDomain(ctx, domainId).Accept(accept).Authorization(authorization).ContentType(contentType).PutDomainRequest(putDomainRequest).Execute()

/domains:/:domain_id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domainId := "domainId_example" // string | 
    accept := "application/json; version=3" // string |  (optional)
    authorization := "Token [TOKEN VALUE]" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    putDomainRequest := *openapiclient.NewPutDomainRequest([]string{"Cnames_example"}, false, "Name_example", false, int64(123), NullableInt64(123)) // PutDomainRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.PutDomain(context.Background(), domainId).Accept(accept).Authorization(authorization).ContentType(contentType).PutDomainRequest(putDomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.PutDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutDomain`: DomainResponseWithResult
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.PutDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **putDomainRequest** | [**PutDomainRequest**](PutDomainRequest.md) |  | 

### Return type

[**DomainResponseWithResult**](DomainResponseWithResult.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDomain

> DomainResponseWithResult UpdateDomain(ctx, domainId).Accept(accept).Authorization(authorization).ContentType(contentType).UpdateDomainRequest(updateDomainRequest).Execute()

/domains/:domain_id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domainId := "domainId_example" // string | 
    accept := "application/json; version=3" // string |  (optional)
    authorization := "Token [TOKEN VALUE]" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    updateDomainRequest := *openapiclient.NewUpdateDomainRequest() // UpdateDomainRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.UpdateDomain(context.Background(), domainId).Accept(accept).Authorization(authorization).ContentType(contentType).UpdateDomainRequest(updateDomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.UpdateDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDomain`: DomainResponseWithResult
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.UpdateDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **updateDomainRequest** | [**UpdateDomainRequest**](UpdateDomainRequest.md) |  | 

### Return type

[**DomainResponseWithResult**](DomainResponseWithResult.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

