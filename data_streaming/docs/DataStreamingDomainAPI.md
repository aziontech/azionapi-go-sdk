# \DataStreamingDomainAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDataStreaming**](DataStreamingDomainAPI.md#ListDataStreaming) | **Get** /data_streaming/domains | List all domains used on data streaming



## ListDataStreaming

> DataStreamingsDomainResponse ListDataStreaming(ctx).Name(name).StreamingId(streamingId).Selected(selected).Execute()

List all domains used on data streaming



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
    name := "name_example" // string | Domain's name in data streaming (optional)
    streamingId := int64(789) // int64 |  (optional)
    selected := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataStreamingDomainAPI.ListDataStreaming(context.Background()).Name(name).StreamingId(streamingId).Selected(selected).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataStreamingDomainAPI.ListDataStreaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDataStreaming`: DataStreamingsDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DataStreamingDomainAPI.ListDataStreaming`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataStreamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Domain&#39;s name in data streaming | 
 **streamingId** | **int64** |  | 
 **selected** | **bool** |  | 

### Return type

[**DataStreamingsDomainResponse**](DataStreamingsDomainResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

