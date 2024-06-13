# \EdgeApplicationsDeviceGroupsAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdDelete**](EdgeApplicationsDeviceGroupsAPI.md#EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdDelete) | **Delete** /edge_applications/{edge_application_id}/device_groups/{device_group_id} | /edge_applications/{edge_application_id}/device_groups/{device_group_id}
[**EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdGet**](EdgeApplicationsDeviceGroupsAPI.md#EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdGet) | **Get** /edge_applications/{edge_application_id}/device_groups/{device_group_id} | /edge_applications/{edge_application_id}/device_groups/{device_group_id}
[**EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPatch**](EdgeApplicationsDeviceGroupsAPI.md#EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPatch) | **Patch** /edge_applications/{edge_application_id}/device_groups/{device_group_id} | /edge_applications/{edge_application_id}/device_groups/{device_group_id}
[**EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPut**](EdgeApplicationsDeviceGroupsAPI.md#EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPut) | **Put** /edge_applications/{edge_application_id}/device_groups/{device_group_id} | /edge_applications/{edge_application_id}/device_groups/{device_group_id}
[**EdgeApplicationsEdgeApplicationIdDeviceGroupsGet**](EdgeApplicationsDeviceGroupsAPI.md#EdgeApplicationsEdgeApplicationIdDeviceGroupsGet) | **Get** /edge_applications/{edge_application_id}/device_groups | /edge_applications/{edge_application_id}/device_groups
[**EdgeApplicationsEdgeApplicationIdDeviceGroupsPost**](EdgeApplicationsDeviceGroupsAPI.md#EdgeApplicationsEdgeApplicationIdDeviceGroupsPost) | **Post** /edge_applications/{edge_application_id}/device_groups | /edge_applications/{edge_application_id}/device_groups



## EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdDelete

> EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdDelete(ctx, edgeApplicationId, deviceGroupId).Accept(accept).Execute()

/edge_applications/{edge_application_id}/device_groups/{device_group_id}

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
	deviceGroupId := int64(789) // int64 | 
	accept := "application/json; version=3" // string | The id of the Device Groups that you plan to delete. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdDelete(context.Background(), edgeApplicationId, deviceGroupId).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**deviceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** | The id of the Device Groups that you plan to delete. | 

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


## EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdGet

> DeviceGroupsIdResponse EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdGet(ctx, edgeApplicationId, deviceGroupId).Accept(accept).Execute()

/edge_applications/{edge_application_id}/device_groups/{device_group_id}

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
	deviceGroupId := int64(789) // int64 | 
	accept := "application/json; version=3" // string | The id of the Device Groups that you plan to query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdGet(context.Background(), edgeApplicationId, deviceGroupId).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdGet`: DeviceGroupsIdResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**deviceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** | The id of the Device Groups that you plan to query. | 

### Return type

[**DeviceGroupsIdResponse**](DeviceGroupsIdResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPatch

> DeviceGroupsIdResponse EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPatch(ctx, edgeApplicationId, deviceGroupId).Accept(accept).ContentType(contentType).PatchDeviceGroupsRequest(patchDeviceGroupsRequest).Execute()

/edge_applications/{edge_application_id}/device_groups/{device_group_id}

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
	deviceGroupId := int64(789) // int64 | 
	accept := "application/json; version=3" // string |  (optional)
	contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
	patchDeviceGroupsRequest := *openapiclient.NewPatchDeviceGroupsRequest() // PatchDeviceGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPatch(context.Background(), edgeApplicationId, deviceGroupId).Accept(accept).ContentType(contentType).PatchDeviceGroupsRequest(patchDeviceGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPatch`: DeviceGroupsIdResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**deviceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **patchDeviceGroupsRequest** | [**PatchDeviceGroupsRequest**](PatchDeviceGroupsRequest.md) |  | 

### Return type

[**DeviceGroupsIdResponse**](DeviceGroupsIdResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPut

> DeviceGroupsIdResponse EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPut(ctx, edgeApplicationId, deviceGroupId).Accept(accept).ContentType(contentType).UpdateDeviceGroupsRequest(updateDeviceGroupsRequest).Execute()

/edge_applications/{edge_application_id}/device_groups/{device_group_id}

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
	deviceGroupId := int64(789) // int64 | 
	accept := "application/json; version=3" // string |  (optional)
	contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
	updateDeviceGroupsRequest := *openapiclient.NewUpdateDeviceGroupsRequest() // UpdateDeviceGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPut(context.Background(), edgeApplicationId, deviceGroupId).Accept(accept).ContentType(contentType).UpdateDeviceGroupsRequest(updateDeviceGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPut`: DeviceGroupsIdResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**deviceGroupId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdDeviceGroupsDeviceGroupIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **updateDeviceGroupsRequest** | [**UpdateDeviceGroupsRequest**](UpdateDeviceGroupsRequest.md) |  | 

### Return type

[**DeviceGroupsIdResponse**](DeviceGroupsIdResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdDeviceGroupsGet

> DeviceGroupsResponse EdgeApplicationsEdgeApplicationIdDeviceGroupsGet(ctx, edgeApplicationId).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()

/edge_applications/{edge_application_id}/device_groups

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
	resp, r, err := apiClient.EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsGet(context.Background(), edgeApplicationId).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdDeviceGroupsGet`: DeviceGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdDeviceGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** |  | 
 **pageSize** | **int64** |  | 
 **filter** | **string** |  | 
 **orderBy** | **string** |  | 
 **sort** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**DeviceGroupsResponse**](DeviceGroupsResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdDeviceGroupsPost

> DeviceGroupsIdResponse EdgeApplicationsEdgeApplicationIdDeviceGroupsPost(ctx, edgeApplicationId).Accept(accept).ContentType(contentType).CreateDeviceGroupsRequest(createDeviceGroupsRequest).Execute()

/edge_applications/{edge_application_id}/device_groups

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
	createDeviceGroupsRequest := *openapiclient.NewCreateDeviceGroupsRequest("UserAgent_example", "Addresses_example") // CreateDeviceGroupsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsPost(context.Background(), edgeApplicationId).Accept(accept).ContentType(contentType).CreateDeviceGroupsRequest(createDeviceGroupsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeApplicationsEdgeApplicationIdDeviceGroupsPost`: DeviceGroupsIdResponse
	fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsDeviceGroupsAPI.EdgeApplicationsEdgeApplicationIdDeviceGroupsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdDeviceGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **createDeviceGroupsRequest** | [**CreateDeviceGroupsRequest**](CreateDeviceGroupsRequest.md) |  | 

### Return type

[**DeviceGroupsIdResponse**](DeviceGroupsIdResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

