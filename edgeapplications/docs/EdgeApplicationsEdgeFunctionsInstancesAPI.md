# \EdgeApplicationsEdgeFunctionsInstancesAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDelete**](EdgeApplicationsEdgeFunctionsInstancesAPI.md#EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDelete) | **Delete** /edge_applications/{edge_application_id}/functions_instances/{functions_instances_id} | /edge_applications/:edge_application_id:/functions_instances/:functions_instances_id:
[**EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet**](EdgeApplicationsEdgeFunctionsInstancesAPI.md#EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet) | **Get** /edge_applications/{edge_application_id}/functions_instances/{functions_instances_id} | /edge_applications/:edge_application_id:/functions_instances/:functions_instances_id:
[**EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch**](EdgeApplicationsEdgeFunctionsInstancesAPI.md#EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch) | **Patch** /edge_applications/{edge_application_id}/functions_instances/{functions_instances_id} | /edge_applications/:edge_application_id:/functions_instances/:functions_instances_id:
[**EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut**](EdgeApplicationsEdgeFunctionsInstancesAPI.md#EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut) | **Put** /edge_applications/{edge_application_id}/functions_instances/{functions_instances_id} | /edge_applications/:edge_application_id:/functions_instances/:functions_instances_id:
[**EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet**](EdgeApplicationsEdgeFunctionsInstancesAPI.md#EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet) | **Get** /edge_applications/{edge_application_id}/functions_instances | /edge_applications/:edge_application_id:/functions_instances
[**EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost**](EdgeApplicationsEdgeFunctionsInstancesAPI.md#EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost) | **Post** /edge_applications/{edge_application_id}/functions_instances | edge_application/:edge_application_id:/functions_instances



## EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDelete

> EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDelete(ctx, edgeApplicationId, functionsInstancesId).Accept(accept).ContentType(contentType).Execute()

/edge_applications/:edge_application_id:/functions_instances/:functions_instances_id:

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
	edgeApplicationId := "edgeApplicationId_example" // string | 
	functionsInstancesId := "functionsInstancesId_example" // string | 
	accept := "application/json; version=3" // string |  (optional)
	contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDelete(context.Background(), edgeApplicationId, functionsInstancesId).Accept(accept).ContentType(contentType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** |  | 
**functionsInstancesId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdDeleteRequest struct via the builder pattern


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


## EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet

> ApplicationInstancesGetOneResponse EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet(ctx, edgeApplicationId, functionsInstancesId).Accept(accept).Execute()

/edge_applications/:edge_application_id:/functions_instances/:functions_instances_id:

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
	functionsInstancesId := int64(789) // int64 | 
	accept := "application/json; version=3" // string | The id of the edge function instance you plan to query.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet(context.Background(), edgeApplicationId, functionsInstancesId).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet`: ApplicationInstancesGetOneResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**functionsInstancesId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** | The id of the edge function instance you plan to query.  | 

### Return type

[**ApplicationInstancesGetOneResponse**](ApplicationInstancesGetOneResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch

> ApplicationInstanceResults EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch(ctx, edgeApplicationId, functionsInstancesId).Accept(accept).ContentType(contentType).ApplicationUpdateInstanceRequest(applicationUpdateInstanceRequest).Execute()

/edge_applications/:edge_application_id:/functions_instances/:functions_instances_id:

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
	edgeApplicationId := "edgeApplicationId_example" // string | The id of the edge application you plan to overwrite 
	functionsInstancesId := "functionsInstancesId_example" // string | The id of the edge function instance you plan to overwrite.
	accept := "application/json; version=3" // string |  (optional)
	contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
	applicationUpdateInstanceRequest := *openapiclient.NewApplicationUpdateInstanceRequest("Name_example", NullableInt64(123), openapiclient.ApplicationInstancesResults_args{ArrayOfMapmapOfStringAny: new([]map[string]interface{})}) // ApplicationUpdateInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch(context.Background(), edgeApplicationId, functionsInstancesId).Accept(accept).ContentType(contentType).ApplicationUpdateInstanceRequest(applicationUpdateInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch`: ApplicationInstanceResults
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** | The id of the edge application you plan to overwrite  | 
**functionsInstancesId** | **string** | The id of the edge function instance you plan to overwrite. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **applicationUpdateInstanceRequest** | [**ApplicationUpdateInstanceRequest**](ApplicationUpdateInstanceRequest.md) |  | 

### Return type

[**ApplicationInstanceResults**](ApplicationInstanceResults.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut

> ApplicationInstanceResults EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut(ctx, edgeApplicationId, functionsInstancesId).Accept(accept).ContentType(contentType).ApplicationPutInstanceRequest(applicationPutInstanceRequest).Execute()

/edge_applications/:edge_application_id:/functions_instances/:functions_instances_id:

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
	edgeApplicationId := "edgeApplicationId_example" // string | The id of the edge application you plan to overwrite 
	functionsInstancesId := "functionsInstancesId_example" // string | The id of the edge function instance you plan to overwrite.
	accept := "application/json; version=3" // string |  (optional)
	contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
	applicationPutInstanceRequest := *openapiclient.NewApplicationPutInstanceRequest("Name_example", int64(123), openapiclient.ApplicationInstancesResults_args{ArrayOfMapmapOfStringAny: new([]map[string]interface{})}) // ApplicationPutInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut(context.Background(), edgeApplicationId, functionsInstancesId).Accept(accept).ContentType(contentType).ApplicationPutInstanceRequest(applicationPutInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut`: ApplicationInstanceResults
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **string** | The id of the edge application you plan to overwrite  | 
**functionsInstancesId** | **string** | The id of the edge function instance you plan to overwrite. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdFunctionsInstancesFunctionsInstancesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **applicationPutInstanceRequest** | [**ApplicationPutInstanceRequest**](ApplicationPutInstanceRequest.md) |  | 

### Return type

[**ApplicationInstanceResults**](ApplicationInstanceResults.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet

> ApplicationInstancesGetResponse EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet(ctx, edgeApplicationId).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()

/edge_applications/:edge_application_id:/functions_instances

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
	resp, r, err := apiClient.EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet(context.Background(), edgeApplicationId).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet`: ApplicationInstancesGetResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdFunctionsInstancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | 
 **pageSize** | **int64** |  | 
 **filter** | **string** |  | 
 **orderBy** | **string** |  | 
 **sort** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**ApplicationInstancesGetResponse**](ApplicationInstancesGetResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost

> ApplicationInstanceResults EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost(ctx, edgeApplicationId).Accept(accept).ContentType(contentType).ApplicationCreateInstanceRequest(applicationCreateInstanceRequest).Execute()

edge_application/:edge_application_id:/functions_instances

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
	applicationCreateInstanceRequest := *openapiclient.NewApplicationCreateInstanceRequest("Name_example", int64(123), openapiclient.ApplicationInstancesResults_args{ArrayOfMapmapOfStringAny: new([]map[string]interface{})}) // ApplicationCreateInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost(context.Background(), edgeApplicationId).Accept(accept).ContentType(contentType).ApplicationCreateInstanceRequest(applicationCreateInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost`: ApplicationInstanceResults
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsEdgeFunctionsInstancesAPI.EdgeApplicationsEdgeApplicationIdFunctionsInstancesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdFunctionsInstancesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **applicationCreateInstanceRequest** | [**ApplicationCreateInstanceRequest**](ApplicationCreateInstanceRequest.md) |  | 

### Return type

[**ApplicationInstanceResults**](ApplicationInstanceResults.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

