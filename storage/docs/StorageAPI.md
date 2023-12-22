# \StorageAPI

All URIs are relative to *https://api.azion.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageApiBucketsCreate**](StorageAPI.md#StorageApiBucketsCreate) | **Post** /v4/storage/buckets | Create a new bucket
[**StorageApiBucketsDestroy**](StorageAPI.md#StorageApiBucketsDestroy) | **Delete** /v4/storage/buckets/{name} | Delete a bucket
[**StorageApiBucketsList**](StorageAPI.md#StorageApiBucketsList) | **Get** /v4/storage/buckets | List buckets
[**StorageApiBucketsObjectsCreate**](StorageAPI.md#StorageApiBucketsObjectsCreate) | **Post** /v4/storage/buckets/{bucket_name}/objects/{object_key} | Create new object key
[**StorageApiBucketsObjectsDestroy**](StorageAPI.md#StorageApiBucketsObjectsDestroy) | **Delete** /v4/storage/buckets/{bucket_name}/objects/{object_key} | Delete object key
[**StorageApiBucketsObjectsList**](StorageAPI.md#StorageApiBucketsObjectsList) | **Get** /v4/storage/buckets/{bucket_name}/objects | List buckets objects
[**StorageApiBucketsObjectsRetrieve**](StorageAPI.md#StorageApiBucketsObjectsRetrieve) | **Get** /v4/storage/buckets/{bucket_name}/objects/{object_key} | Download object
[**StorageApiBucketsObjectsUpdate**](StorageAPI.md#StorageApiBucketsObjectsUpdate) | **Put** /v4/storage/buckets/{bucket_name}/objects/{object_key} | Update the object key
[**StorageApiBucketsPartialUpdate**](StorageAPI.md#StorageApiBucketsPartialUpdate) | **Patch** /v4/storage/buckets/{name} | Update bucket info



## StorageApiBucketsCreate

> ResponseBucket StorageApiBucketsCreate(ctx).BucketCreate(bucketCreate).Execute()

Create a new bucket



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
	resp, r, err := apiClient.StorageAPI.StorageApiBucketsCreate(context.Background()).BucketCreate(bucketCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiBucketsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageApiBucketsCreate`: ResponseBucket
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageApiBucketsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageApiBucketsCreateRequest struct via the builder pattern


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


## StorageApiBucketsDestroy

> SuccessBucketOperation StorageApiBucketsDestroy(ctx, name).Execute()

Delete a bucket



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
	resp, r, err := apiClient.StorageAPI.StorageApiBucketsDestroy(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiBucketsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageApiBucketsDestroy`: SuccessBucketOperation
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageApiBucketsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageApiBucketsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessBucketOperation**](SuccessBucketOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageApiBucketsList

> PaginatedBucketList StorageApiBucketsList(ctx).Page(page).PageSize(pageSize).Execute()

List buckets



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
	resp, r, err := apiClient.StorageAPI.StorageApiBucketsList(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiBucketsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageApiBucketsList`: PaginatedBucketList
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageApiBucketsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageApiBucketsListRequest struct via the builder pattern


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


## StorageApiBucketsObjectsCreate

> SuccessObjectOperation StorageApiBucketsObjectsCreate(ctx, bucketName, objectKey).Body(body).Execute()

Create new object key



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
	bucketName := "bucketName_example" // string | 
	objectKey := "objectKey_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageApiBucketsObjectsCreate(context.Background(), bucketName, objectKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiBucketsObjectsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageApiBucketsObjectsCreate`: SuccessObjectOperation
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageApiBucketsObjectsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 
**objectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageApiBucketsObjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 

### Return type

[**SuccessObjectOperation**](SuccessObjectOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageApiBucketsObjectsDestroy

> SuccessObjectOperation StorageApiBucketsObjectsDestroy(ctx, bucketName, objectKey).Execute()

Delete object key



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
	bucketName := "bucketName_example" // string | 
	objectKey := "objectKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageApiBucketsObjectsDestroy(context.Background(), bucketName, objectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiBucketsObjectsDestroy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageApiBucketsObjectsDestroy`: SuccessObjectOperation
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageApiBucketsObjectsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 
**objectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageApiBucketsObjectsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SuccessObjectOperation**](SuccessObjectOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageApiBucketsObjectsList

> PaginatedBucketObjectList StorageApiBucketsObjectsList(ctx, bucketName).Page(page).PageSize(pageSize).Execute()

List buckets objects



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
	bucketName := "bucketName_example" // string | 
	page := int32(56) // int32 | A page number within the paginated result set. (optional)
	pageSize := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageApiBucketsObjectsList(context.Background(), bucketName).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiBucketsObjectsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageApiBucketsObjectsList`: PaginatedBucketObjectList
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageApiBucketsObjectsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageApiBucketsObjectsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 

### Return type

[**PaginatedBucketObjectList**](PaginatedBucketObjectList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageApiBucketsObjectsRetrieve

> *os.File StorageApiBucketsObjectsRetrieve(ctx, bucketName, objectKey).Execute()

Download object



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
	bucketName := "bucketName_example" // string | 
	objectKey := "objectKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageApiBucketsObjectsRetrieve(context.Background(), bucketName, objectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiBucketsObjectsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageApiBucketsObjectsRetrieve`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageApiBucketsObjectsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 
**objectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageApiBucketsObjectsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageApiBucketsObjectsUpdate

> SuccessObjectOperation StorageApiBucketsObjectsUpdate(ctx, bucketName, objectKey).Body(body).Execute()

Update the object key



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
	bucketName := "bucketName_example" // string | 
	objectKey := "objectKey_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageApiBucketsObjectsUpdate(context.Background(), bucketName, objectKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiBucketsObjectsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageApiBucketsObjectsUpdate`: SuccessObjectOperation
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageApiBucketsObjectsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucketName** | **string** |  | 
**objectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageApiBucketsObjectsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 

### Return type

[**SuccessObjectOperation**](SuccessObjectOperation.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageApiBucketsPartialUpdate

> ResponseBucket StorageApiBucketsPartialUpdate(ctx, name).Execute()

Update bucket info



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
	resp, r, err := apiClient.StorageAPI.StorageApiBucketsPartialUpdate(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiBucketsPartialUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageApiBucketsPartialUpdate`: ResponseBucket
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageApiBucketsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageApiBucketsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseBucket**](ResponseBucket.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

