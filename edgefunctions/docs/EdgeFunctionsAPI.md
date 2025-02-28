# \EdgeFunctionsAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeFunctionsGet**](EdgeFunctionsAPI.md#EdgeFunctionsGet) | **Get** /edge_functions | edge_functions
[**EdgeFunctionsIdDelete**](EdgeFunctionsAPI.md#EdgeFunctionsIdDelete) | **Delete** /edge_functions/{id} | edge_functions
[**EdgeFunctionsIdGet**](EdgeFunctionsAPI.md#EdgeFunctionsIdGet) | **Get** /edge_functions/{id} | edge_functions
[**EdgeFunctionsIdPatch**](EdgeFunctionsAPI.md#EdgeFunctionsIdPatch) | **Patch** /edge_functions/{id} | edge_functions
[**EdgeFunctionsIdPut**](EdgeFunctionsAPI.md#EdgeFunctionsIdPut) | **Put** /edge_functions/{id} | edge_functions
[**EdgeFunctionsPost**](EdgeFunctionsAPI.md#EdgeFunctionsPost) | **Post** /edge_functions | edge_functions



## EdgeFunctionsGet

> ListEdgeFunctionResponse EdgeFunctionsGet(ctx).Page(page).PageSize(pageSize).Sort(sort).OrderBy(orderBy).Execute()

edge_functions

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
	page := int64(789) // int64 |  (optional)
	pageSize := int64(789) // int64 |  (optional)
	sort := "sort_example" // string |  (optional)
	orderBy := "orderBy_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.EdgeFunctionsGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeFunctionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeFunctionsGet`: ListEdgeFunctionResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.EdgeFunctionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFunctionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | 
 **pageSize** | **int64** |  | 
 **sort** | **string** |  | 
 **orderBy** | **string** |  | 

### Return type

[**ListEdgeFunctionResponse**](ListEdgeFunctionResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFunctionsIdDelete

> EdgeFunctionsIdDelete(ctx, id).Execute()

edge_functions

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
	id := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EdgeFunctionsAPI.EdgeFunctionsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeFunctionsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFunctionsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFunctionsIdGet

> EdgeFunctionResponse EdgeFunctionsIdGet(ctx, id).Execute()

edge_functions

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
	id := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.EdgeFunctionsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeFunctionsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeFunctionsIdGet`: EdgeFunctionResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.EdgeFunctionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFunctionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EdgeFunctionResponse**](EdgeFunctionResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFunctionsIdPatch

> EdgeFunctionResponse EdgeFunctionsIdPatch(ctx, id).PatchEdgeFunctionRequest(patchEdgeFunctionRequest).Execute()

edge_functions

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
	id := int64(789) // int64 | 
	patchEdgeFunctionRequest := *openapiclient.NewPatchEdgeFunctionRequest() // PatchEdgeFunctionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.EdgeFunctionsIdPatch(context.Background(), id).PatchEdgeFunctionRequest(patchEdgeFunctionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeFunctionsIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeFunctionsIdPatch`: EdgeFunctionResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.EdgeFunctionsIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFunctionsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchEdgeFunctionRequest** | [**PatchEdgeFunctionRequest**](PatchEdgeFunctionRequest.md) |  | 

### Return type

[**EdgeFunctionResponse**](EdgeFunctionResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFunctionsIdPut

> EdgeFunctionResponse EdgeFunctionsIdPut(ctx, id).PutEdgeFunctionRequest(putEdgeFunctionRequest).Execute()

edge_functions

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
	id := int64(789) // int64 | 
	putEdgeFunctionRequest := *openapiclient.NewPutEdgeFunctionRequest() // PutEdgeFunctionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.EdgeFunctionsIdPut(context.Background(), id).PutEdgeFunctionRequest(putEdgeFunctionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeFunctionsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeFunctionsIdPut`: EdgeFunctionResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.EdgeFunctionsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFunctionsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putEdgeFunctionRequest** | [**PutEdgeFunctionRequest**](PutEdgeFunctionRequest.md) |  | 

### Return type

[**EdgeFunctionResponse**](EdgeFunctionResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFunctionsPost

> EdgeFunctionResponse EdgeFunctionsPost(ctx).CreateEdgeFunctionRequest(createEdgeFunctionRequest).Execute()

edge_functions

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
	createEdgeFunctionRequest := *openapiclient.NewCreateEdgeFunctionRequest() // CreateEdgeFunctionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.EdgeFunctionsPost(context.Background()).CreateEdgeFunctionRequest(createEdgeFunctionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeFunctionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeFunctionsPost`: EdgeFunctionResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.EdgeFunctionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFunctionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEdgeFunctionRequest** | [**CreateEdgeFunctionRequest**](CreateEdgeFunctionRequest.md) |  | 

### Return type

[**EdgeFunctionResponse**](EdgeFunctionResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

