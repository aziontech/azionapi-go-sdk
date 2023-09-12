# \DataStreamingTemplatesAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataStramingTemplateById**](DataStreamingTemplatesAPI.md#GetDataStramingTemplateById) | **Get** /data_streaming/templates/{template_id} | Get an global Template info by template ID
[**ListDataStreamingTemplates**](DataStreamingTemplatesAPI.md#ListDataStreamingTemplates) | **Get** /data_streaming/templates | List all global Templates that can be used on Data Streaming operations



## GetDataStramingTemplateById

> TemplateResultById GetDataStramingTemplateById(ctx, templateId).Execute()

Get an global Template info by template ID



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
    templateId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStreamingTemplatesAPI.GetDataStramingTemplateById(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStreamingTemplatesAPI.GetDataStramingTemplateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataStramingTemplateById`: TemplateResultById
    fmt.Fprintf(os.Stdout, "Response from `DataStreamingTemplatesAPI.GetDataStramingTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataStramingTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateResultById**](TemplateResultById.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataStreamingTemplates

> TemplateResults ListDataStreamingTemplates(ctx).Execute()

List all global Templates that can be used on Data Streaming operations



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStreamingTemplatesAPI.ListDataStreamingTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStreamingTemplatesAPI.ListDataStreamingTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDataStreamingTemplates`: TemplateResults
    fmt.Fprintf(os.Stdout, "Response from `DataStreamingTemplatesAPI.ListDataStreamingTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDataStreamingTemplatesRequest struct via the builder pattern


### Return type

[**TemplateResults**](TemplateResults.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

