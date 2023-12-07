# \BucketsAPI

All URIs are relative to *https://api.azion.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1StorageBucketsCreate**](BucketsAPI.md#ApiV1StorageBucketsCreate) | **Post** /v4/storage/buckets | /v4/storage/buckets
[**ApiV1StorageBucketsDestroy**](BucketsAPI.md#ApiV1StorageBucketsDestroy) | **Delete** /v4/storage/buckets/{name} | /v4/storage/buckets/:name
[**ApiV1StorageBucketsList**](BucketsAPI.md#ApiV1StorageBucketsList) | **Get** /v4/storage/buckets | /v4/storage/buckets
[**ApiV1StorageBucketsPartialUpdate**](BucketsAPI.md#ApiV1StorageBucketsPartialUpdate) | **Patch** /v4/storage/buckets/{name} | /v4/storage/buckets/:name



## ApiV1StorageBucketsCreate

> ResponseBucket ApiV1StorageBucketsCreate(ctx).BucketCreate(bucketCreate).Execute()

/v4/storage/buckets



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
    bucketCreate := *openapiclient.NewBucketCreate("Name_example", openapiclient.EdgeAccessEnum("read_only")) // BucketCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketsAPI.ApiV1StorageBucketsCreate(context.Background()).BucketCreate(bucketCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsAPI.ApiV1StorageBucketsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1StorageBucketsCreate`: ResponseBucket
    fmt.Fprintf(os.Stdout, "Response from `BucketsAPI.ApiV1StorageBucketsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1StorageBucketsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketCreate** | [**BucketCreate**](BucketCreate.md) |  | 

### Return type

[**ResponseBucket**](ResponseBucket.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1StorageBucketsDestroy

> ResponseDeleteBucket ApiV1StorageBucketsDestroy(ctx, name).Execute()

/v4/storage/buckets/:name



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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketsAPI.ApiV1StorageBucketsDestroy(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsAPI.ApiV1StorageBucketsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1StorageBucketsDestroy`: ResponseDeleteBucket
    fmt.Fprintf(os.Stdout, "Response from `BucketsAPI.ApiV1StorageBucketsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1StorageBucketsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeleteBucket**](ResponseDeleteBucket.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1StorageBucketsList

> PaginatedBucketList ApiV1StorageBucketsList(ctx).Page(page).PageSize(pageSize).Execute()

/v4/storage/buckets



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
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketsAPI.ApiV1StorageBucketsList(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsAPI.ApiV1StorageBucketsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1StorageBucketsList`: PaginatedBucketList
    fmt.Fprintf(os.Stdout, "Response from `BucketsAPI.ApiV1StorageBucketsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1StorageBucketsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedBucketList**](PaginatedBucketList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1StorageBucketsPartialUpdate

> ResponseBucket ApiV1StorageBucketsPartialUpdate(ctx, name).PatchedBucket(patchedBucket).Execute()

/v4/storage/buckets/:name



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
    name := "name_example" // string | 
    patchedBucket := *openapiclient.NewPatchedBucket() // PatchedBucket |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketsAPI.ApiV1StorageBucketsPartialUpdate(context.Background(), name).PatchedBucket(patchedBucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsAPI.ApiV1StorageBucketsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1StorageBucketsPartialUpdate`: ResponseBucket
    fmt.Fprintf(os.Stdout, "Response from `BucketsAPI.ApiV1StorageBucketsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1StorageBucketsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedBucket** | [**PatchedBucket**](PatchedBucket.md) |  | 

### Return type

[**ResponseBucket**](ResponseBucket.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

