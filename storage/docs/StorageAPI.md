# \StorageAPI

All URIs are relative to *https://api.azion.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageApiBucketsCreate**](StorageAPI.md#StorageApiBucketsCreate) | **Post** /v4/edge_storage/buckets | Create a new bucket
[**StorageApiBucketsDestroy**](StorageAPI.md#StorageApiBucketsDestroy) | **Delete** /v4/edge_storage/buckets/{name} | Delete a bucket
[**StorageApiBucketsList**](StorageAPI.md#StorageApiBucketsList) | **Get** /v4/edge_storage/buckets | List buckets
[**StorageApiBucketsObjectsCreate**](StorageAPI.md#StorageApiBucketsObjectsCreate) | **Post** /v4/edge_storage/buckets/{bucket_name}/objects/{object_key} | Create new object key
[**StorageApiBucketsObjectsDestroy**](StorageAPI.md#StorageApiBucketsObjectsDestroy) | **Delete** /v4/edge_storage/buckets/{bucket_name}/objects/{object_key} | Delete object key
[**StorageApiBucketsObjectsList**](StorageAPI.md#StorageApiBucketsObjectsList) | **Get** /v4/edge_storage/buckets/{bucket_name}/objects | List buckets objects
[**StorageApiBucketsObjectsRetrieve**](StorageAPI.md#StorageApiBucketsObjectsRetrieve) | **Get** /v4/edge_storage/buckets/{bucket_name}/objects/{object_key} | Download object
[**StorageApiBucketsObjectsUpdate**](StorageAPI.md#StorageApiBucketsObjectsUpdate) | **Put** /v4/edge_storage/buckets/{bucket_name}/objects/{object_key} | Update the object key
[**StorageApiBucketsPartialUpdate**](StorageAPI.md#StorageApiBucketsPartialUpdate) | **Patch** /v4/edge_storage/buckets/{name} | Update bucket info
[**StorageApiS3CredentialsByAccessKey**](StorageAPI.md#StorageApiS3CredentialsByAccessKey) | **Get** /v4/edge_storage/s3-credentials/{s3_credential_access_key} | get by s3 credentials by access key
[**StorageApiS3CredentialsCreate**](StorageAPI.md#StorageApiS3CredentialsCreate) | **Post** /v4/edge_storage/s3-credentials | create s3 credentials
[**StorageApiS3CredentialsDelete**](StorageAPI.md#StorageApiS3CredentialsDelete) | **Delete** /v4/edge_storage/s3-credentials/{s3_credential_access_key} | delete by s3 credentials
[**StorageApiS3CredentialsList**](StorageAPI.md#StorageApiS3CredentialsList) | **Get** /v4/edge_storage/s3-credentials | List s3 credentials



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

> SuccessObjectOperation StorageApiBucketsObjectsCreate(ctx, bucketName, objectKey).ContentType(contentType).Body(body).Execute()

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
	contentType := "contentType_example" // string | The content type of the file (Example: text/plain). (optional)
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageApiBucketsObjectsCreate(context.Background(), bucketName, objectKey).ContentType(contentType).Body(body).Execute()
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


 **contentType** | **string** | The content type of the file (Example: text/plain). | 
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

> PaginatedBucketObjectList StorageApiBucketsObjectsList(ctx, bucketName).ContinuationToken(continuationToken).MaxObjectCount(maxObjectCount).Execute()

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
	continuationToken := "continuationToken_example" // string | Token for next page. (optional)
	maxObjectCount := int32(56) // int32 | Number of results to return per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageApiBucketsObjectsList(context.Background(), bucketName).ContinuationToken(continuationToken).MaxObjectCount(maxObjectCount).Execute()
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

 **continuationToken** | **string** | Token for next page. | 
 **maxObjectCount** | **int32** | Number of results to return per page. | 

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

> StorageApiBucketsObjectsRetrieve(ctx, bucketName, objectKey).Execute()

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
	r, err := apiClient.StorageAPI.StorageApiBucketsObjectsRetrieve(context.Background(), bucketName, objectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiBucketsObjectsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
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

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html, application/json, application/xml, text/plain, image/jpeg, image/png, image/gif, video/mp4, audio/mpeg, application/pdf, application/javascript, text/css, application/octet-stream, multipart/form-data, application/x-www-form-urlencoded

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageApiBucketsObjectsUpdate

> SuccessObjectOperation StorageApiBucketsObjectsUpdate(ctx, bucketName, objectKey).ContentType(contentType).Body(body).Execute()

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
	contentType := "contentType_example" // string | The content type of the file (Example: text/plain). (optional)
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageApiBucketsObjectsUpdate(context.Background(), bucketName, objectKey).ContentType(contentType).Body(body).Execute()
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


 **contentType** | **string** | The content type of the file (Example: text/plain). | 
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

> ResponseBucket StorageApiBucketsPartialUpdate(ctx, name).BucketUpdate(bucketUpdate).Execute()

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
	bucketUpdate := *openapiclient.NewBucketUpdate(openapiclient.EdgeAccessEnum("read_only")) // BucketUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageApiBucketsPartialUpdate(context.Background(), name).BucketUpdate(bucketUpdate).Execute()
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

 **bucketUpdate** | [**BucketUpdate**](BucketUpdate.md) |  | 

### Return type

[**ResponseBucket**](ResponseBucket.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageApiS3CredentialsByAccessKey

> ResponseS3Credential StorageApiS3CredentialsByAccessKey(ctx, s3CredentialAccessKey).Execute()

get by s3 credentials by access key



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
	s3CredentialAccessKey := "s3CredentialAccessKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageApiS3CredentialsByAccessKey(context.Background(), s3CredentialAccessKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiS3CredentialsByAccessKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageApiS3CredentialsByAccessKey`: ResponseS3Credential
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageApiS3CredentialsByAccessKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**s3CredentialAccessKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageApiS3CredentialsByAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseS3Credential**](ResponseS3Credential.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageApiS3CredentialsCreate

> ResponseS3Credential StorageApiS3CredentialsCreate(ctx).S3CredentialCreate(s3CredentialCreate).Execute()

create s3 credentials



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
	s3CredentialCreate := *openapiclient.NewS3CredentialCreate() // S3CredentialCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageApiS3CredentialsCreate(context.Background()).S3CredentialCreate(s3CredentialCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiS3CredentialsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageApiS3CredentialsCreate`: ResponseS3Credential
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageApiS3CredentialsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageApiS3CredentialsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s3CredentialCreate** | [**S3CredentialCreate**](S3CredentialCreate.md) |  | 

### Return type

[**ResponseS3Credential**](ResponseS3Credential.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageApiS3CredentialsDelete

> ResponseS3Credential StorageApiS3CredentialsDelete(ctx, s3CredentialAccessKey).Execute()

delete by s3 credentials



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
	s3CredentialAccessKey := "s3CredentialAccessKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageApiS3CredentialsDelete(context.Background(), s3CredentialAccessKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiS3CredentialsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageApiS3CredentialsDelete`: ResponseS3Credential
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageApiS3CredentialsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**s3CredentialAccessKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageApiS3CredentialsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseS3Credential**](ResponseS3Credential.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageApiS3CredentialsList

> PaginatedS3CredentialList StorageApiS3CredentialsList(ctx).Key(key).LastModified(lastModified).Size(size).ContinuationToken(continuationToken).Execute()

List s3 credentials



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
	key := "key_example" // string | Object key. Used to identify the object for requests. Sent in POST requests as a path variable. (optional)
	lastModified := "lastModified_example" // string | Timestamp of the last modification to the object. (optional)
	size := int32(56) // int32 | Size of file in bytes. (optional)
	continuationToken := "continuationToken_example" // string | Hash that can be added to the continuation_token query to skip list to the next page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageApiS3CredentialsList(context.Background()).Key(key).LastModified(lastModified).Size(size).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageApiS3CredentialsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageApiS3CredentialsList`: PaginatedS3CredentialList
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageApiS3CredentialsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageApiS3CredentialsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | Object key. Used to identify the object for requests. Sent in POST requests as a path variable. | 
 **lastModified** | **string** | Timestamp of the last modification to the object. | 
 **size** | **int32** | Size of file in bytes. | 
 **continuationToken** | **string** | Hash that can be added to the continuation_token query to skip list to the next page. | 

### Return type

[**PaginatedS3CredentialList**](PaginatedS3CredentialList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

