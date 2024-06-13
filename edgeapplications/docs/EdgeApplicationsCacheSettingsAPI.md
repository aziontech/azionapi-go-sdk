# \EdgeApplicationsCacheSettingsAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdDelete**](EdgeApplicationsCacheSettingsAPI.md#EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdDelete) | **Delete** /edge_applications/{edge_application_id}/cache_settings/{cache_settings_id} | /edge_applications/:edge_application_id:/cache_settings/:cache_settings_id:
[**EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet**](EdgeApplicationsCacheSettingsAPI.md#EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet) | **Get** /edge_applications/{edge_application_id}/cache_settings/{cache_settings_id} | /edge_applications/:edge_application_id:/cache_settings/:cache_settings_id:
[**EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPatch**](EdgeApplicationsCacheSettingsAPI.md#EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPatch) | **Patch** /edge_applications/{edge_application_id}/cache_settings/{cache_settings_id} | /edge_applications/:edge_application_id:/cache_settings/:cache_settings_id:
[**EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut**](EdgeApplicationsCacheSettingsAPI.md#EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut) | **Put** /edge_applications/{edge_application_id}/cache_settings/{cache_settings_id} | /edge_applications/:edge_application_id:/cache_settings/ca
[**EdgeApplicationsEdgeApplicationIdCacheSettingsGet**](EdgeApplicationsCacheSettingsAPI.md#EdgeApplicationsEdgeApplicationIdCacheSettingsGet) | **Get** /edge_applications/{edge_application_id}/cache_settings | /edge_applications/{edge_application_id}/cache_settings
[**EdgeApplicationsEdgeApplicationIdCacheSettingsPost**](EdgeApplicationsCacheSettingsAPI.md#EdgeApplicationsEdgeApplicationIdCacheSettingsPost) | **Post** /edge_applications/{edge_application_id}/cache_settings | /edge_applications/:edge_application_id:/cache_settings



## EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdDelete

> EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdDelete(ctx, edgeApplicationId, cacheSettingsId).Accept(accept).ContentType(contentType).Execute()

/edge_applications/:edge_application_id:/cache_settings/:cache_settings_id:

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
	edgeApplicationId := int64(789) // int64 | 
	cacheSettingsId := int64(789) // int64 | 
	accept := "application/json; version=3" // string |  (optional)
	contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdDelete(context.Background(), edgeApplicationId, cacheSettingsId).Accept(accept).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**cacheSettingsId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet

> ApplicationCacheGetOneResponse EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet(ctx, edgeApplicationId, cacheSettingsId).Accept(accept).Execute()

/edge_applications/:edge_application_id:/cache_settings/:cache_settings_id:

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
	edgeApplicationId := int64(789) // int64 | 
	cacheSettingsId := int64(789) // int64 | 
	accept := "application/json; version=3" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet(context.Background(), edgeApplicationId, cacheSettingsId).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet`: ApplicationCacheGetOneResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**cacheSettingsId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 

### Return type

[**ApplicationCacheGetOneResponse**](ApplicationCacheGetOneResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPatch

> ApplicationCachePatchResponse EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPatch(ctx, edgeApplicationId, cacheSettingsId).Accept(accept).ApplicationCachePatchRequest(applicationCachePatchRequest).Execute()

/edge_applications/:edge_application_id:/cache_settings/:cache_settings_id:

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
	edgeApplicationId := int64(789) // int64 | 
	cacheSettingsId := int64(789) // int64 | 
	accept := "application/json; version=3" // string |  (optional)
	applicationCachePatchRequest := *openapiclient.NewApplicationCachePatchRequest() // ApplicationCachePatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPatch(context.Background(), edgeApplicationId, cacheSettingsId).Accept(accept).ApplicationCachePatchRequest(applicationCachePatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPatch`: ApplicationCachePatchResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**cacheSettingsId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **applicationCachePatchRequest** | [**ApplicationCachePatchRequest**](ApplicationCachePatchRequest.md) |  | 

### Return type

[**ApplicationCachePatchResponse**](ApplicationCachePatchResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut

> ApplicationCachePutResponse EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut(ctx, edgeApplicationId, cacheSettingsId).Accept(accept).ContentType(contentType).ApplicationCachePutRequest(applicationCachePutRequest).Execute()

/edge_applications/:edge_application_id:/cache_settings/ca

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
	edgeApplicationId := int64(789) // int64 | 
	cacheSettingsId := int64(789) // int64 | 
	accept := "application/json; version=3" // string |  (optional)
	contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
	applicationCachePutRequest := *openapiclient.NewApplicationCachePutRequest("Name_example") // ApplicationCachePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut(context.Background(), edgeApplicationId, cacheSettingsId).Accept(accept).ContentType(contentType).ApplicationCachePutRequest(applicationCachePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut`: ApplicationCachePutResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**cacheSettingsId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdCacheSettingsCacheSettingsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **applicationCachePutRequest** | [**ApplicationCachePutRequest**](ApplicationCachePutRequest.md) |  | 

### Return type

[**ApplicationCachePutResponse**](ApplicationCachePutResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdCacheSettingsGet

> ApplicationCacheGetResponse EdgeApplicationsEdgeApplicationIdCacheSettingsGet(ctx, edgeApplicationId).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()

/edge_applications/{edge_application_id}/cache_settings

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
	edgeApplicationId := int64(789) // int64 | 
	page := int64(789) // int64 |  (optional)
	pageSize := int64(789) // int64 |  (optional)
	filter := "filter_example" // string |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	sort := "sort_example" // string |  (optional)
	accept := "application/json; version=3" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsGet(context.Background(), edgeApplicationId).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdCacheSettingsGet`: ApplicationCacheGetResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdCacheSettingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | 
 **pageSize** | **int64** |  | 
 **filter** | **string** |  | 
 **orderBy** | **string** |  | 
 **sort** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**ApplicationCacheGetResponse**](ApplicationCacheGetResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdCacheSettingsPost

> ApplicationCacheCreateResponse EdgeApplicationsEdgeApplicationIdCacheSettingsPost(ctx, edgeApplicationId).Accept(accept).ContentType(contentType).ApplicationCacheCreateRequest(applicationCacheCreateRequest).Execute()

/edge_applications/:edge_application_id:/cache_settings

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
	edgeApplicationId := int64(789) // int64 | 
	accept := "application/json; version=3" // string |  (optional)
	contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
	applicationCacheCreateRequest := *openapiclient.NewApplicationCacheCreateRequest("Name_example") // ApplicationCacheCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsPost(context.Background(), edgeApplicationId).Accept(accept).ContentType(contentType).ApplicationCacheCreateRequest(applicationCacheCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdCacheSettingsPost`: ApplicationCacheCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsCacheSettingsAPI.EdgeApplicationsEdgeApplicationIdCacheSettingsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdCacheSettingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **applicationCacheCreateRequest** | [**ApplicationCacheCreateRequest**](ApplicationCacheCreateRequest.md) |  | 

### Return type

[**ApplicationCacheCreateResponse**](ApplicationCacheCreateResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

