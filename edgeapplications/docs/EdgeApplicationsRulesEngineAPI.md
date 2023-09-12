# \EdgeApplicationsRulesEngineAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGet**](EdgeApplicationsRulesEngineAPI.md#EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGet) | **Get** /edge_applications/{edge_application_id}/rules_engine/{phase}/rules | /edge_applications/{edge_application_id}/rules_engine/{phase}/rules
[**EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPost**](EdgeApplicationsRulesEngineAPI.md#EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPost) | **Post** /edge_applications/{edge_application_id}/rules_engine/{phase}/rules | /edge_applications/{edge_application_id}/rules_engine/{phase}/rules
[**EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDelete**](EdgeApplicationsRulesEngineAPI.md#EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDelete) | **Delete** /edge_applications/{edge_application_id}/rules_engine/{phase}/rules/{rule_id} | /edge_applications/{edge_application_id}/rules_engine/{phase}/rules
[**EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGet**](EdgeApplicationsRulesEngineAPI.md#EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGet) | **Get** /edge_applications/{edge_application_id}/rules_engine/{phase}/rules/{rule_id} | /edge_applications/{edge_application_id}/rules_engine/{phase}/rules
[**EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatch**](EdgeApplicationsRulesEngineAPI.md#EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatch) | **Patch** /edge_applications/{edge_application_id}/rules_engine/{phase}/rules/{rule_id} | /edge_applications/:edge_application_id:/rules_engine/:phase:/rules/:rule_id:
[**EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPut**](EdgeApplicationsRulesEngineAPI.md#EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPut) | **Put** /edge_applications/{edge_application_id}/rules_engine/{phase}/rules/{rule_id} | /edge_applications/:edge_application_id:/rules_engine/:phase:/rules/:rule_id:



## EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGet

> RulesEngineResponse EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGet(ctx, edgeApplicationId, phase).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()

/edge_applications/{edge_application_id}/rules_engine/{phase}/rules

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
    phase := "phase_example" // string | 
    page := int64(789) // int64 |  (optional)
    pageSize := int64(789) // int64 |  (optional)
    filter := "filter_example" // string |  (optional)
    orderBy := "orderBy_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    accept := "application/json; version=3" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGet(context.Background(), edgeApplicationId, phase).Page(page).PageSize(pageSize).Filter(filter).OrderBy(orderBy).Sort(sort).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGet`: RulesEngineResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**phase** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** |  | 
 **pageSize** | **int64** |  | 
 **filter** | **string** |  | 
 **orderBy** | **string** |  | 
 **sort** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**RulesEngineResponse**](RulesEngineResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPost

> RulesEngineIdResponse EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPost(ctx, edgeApplicationId, phase).Accept(accept).ContentType(contentType).CreateRulesEngineRequest(createRulesEngineRequest).Execute()

/edge_applications/{edge_application_id}/rules_engine/{phase}/rules



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
    phase := "phase_example" // string | 
    accept := "application/json; version=3" // string |  (optional)
    contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
    createRulesEngineRequest := *openapiclient.NewCreateRulesEngineRequest("Name_example", [][]RulesEngineCriteria{[]openapiclient.RulesEngineCriteria{*openapiclient.NewRulesEngineCriteria("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.RulesEngineBehavior{*openapiclient.NewRulesEngineBehavior("Name_example")}) // CreateRulesEngineRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPost(context.Background(), edgeApplicationId, phase).Accept(accept).ContentType(contentType).CreateRulesEngineRequest(createRulesEngineRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPost`: RulesEngineIdResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**phase** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **createRulesEngineRequest** | [**CreateRulesEngineRequest**](CreateRulesEngineRequest.md) |  | 

### Return type

[**RulesEngineIdResponse**](RulesEngineIdResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDelete

> EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDelete(ctx, edgeApplicationId, phase, ruleId).Accept(accept).Execute()

/edge_applications/{edge_application_id}/rules_engine/{phase}/rules

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
    edgeApplicationId := int64(789) // int64 | The id of the edge application you plan to delete. 
    phase := "phase_example" // string | 
    ruleId := int64(789) // int64 | The id of the rule you plan to delete. 
    accept := "application/json; version=3" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDelete(context.Background(), edgeApplicationId, phase, ruleId).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** | The id of the edge application you plan to delete.  | 
**phase** | **string** |  | 
**ruleId** | **int64** | The id of the rule you plan to delete.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accept** | **string** |  | 

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


## EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGet

> RulesEngineIdResponse EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGet(ctx, edgeApplicationId, phase, ruleId).Accept(accept).Execute()

/edge_applications/{edge_application_id}/rules_engine/{phase}/rules

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
    edgeApplicationId := int64(789) // int64 | The id of the edge application you want to get. 
    phase := "phase_example" // string | 
    ruleId := int64(789) // int64 | The id of the rule you plan to delete. 
    accept := "application/json; version=3" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGet(context.Background(), edgeApplicationId, phase, ruleId).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGet`: RulesEngineIdResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** | The id of the edge application you want to get.  | 
**phase** | **string** |  | 
**ruleId** | **int64** | The id of the rule you plan to delete.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accept** | **string** |  | 

### Return type

[**RulesEngineIdResponse**](RulesEngineIdResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatch

> RulesEngineIdResponse EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatch(ctx, edgeApplicationId, phase, ruleId).Accept(accept).ContentType(contentType).PatchRulesEngineRequest(patchRulesEngineRequest).Execute()

/edge_applications/:edge_application_id:/rules_engine/:phase:/rules/:rule_id:

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
    phase := "phase_example" // string | 
    ruleId := int64(789) // int64 | 
    accept := "application/json; version=3" // string |  (optional)
    contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
    patchRulesEngineRequest := *openapiclient.NewPatchRulesEngineRequest() // PatchRulesEngineRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatch(context.Background(), edgeApplicationId, phase, ruleId).Accept(accept).ContentType(contentType).PatchRulesEngineRequest(patchRulesEngineRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatch`: RulesEngineIdResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**phase** | **string** |  | 
**ruleId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **patchRulesEngineRequest** | [**PatchRulesEngineRequest**](PatchRulesEngineRequest.md) |  | 

### Return type

[**RulesEngineIdResponse**](RulesEngineIdResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPut

> RulesEngineIdResponse EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPut(ctx, edgeApplicationId, phase, ruleId).Accept(accept).ContentType(contentType).UpdateRulesEngineRequest(updateRulesEngineRequest).Execute()

/edge_applications/:edge_application_id:/rules_engine/:phase:/rules/:rule_id:

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
    phase := "phase_example" // string | 
    ruleId := int64(789) // int64 | 
    accept := "application/json; version=3" // string |  (optional)
    contentType := "application/json" // string | The type of coding used in the Body (application/json). <br>  Example: Content-Type: application/json (optional)
    updateRulesEngineRequest := *openapiclient.NewUpdateRulesEngineRequest("Name_example", [][]RulesEngineCriteria{[]openapiclient.RulesEngineCriteria{*openapiclient.NewRulesEngineCriteria("Conditional_example", "Variable_example", "Operator_example")}}, []openapiclient.RulesEngineBehavior{*openapiclient.NewRulesEngineBehavior("Name_example")}) // UpdateRulesEngineRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPut(context.Background(), edgeApplicationId, phase, ruleId).Accept(accept).ContentType(contentType).UpdateRulesEngineRequest(updateRulesEngineRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPut`: RulesEngineIdResponse
    fmt.Fprintf(os.Stdout, "Response from `EdgeApplicationsRulesEngineAPI.EdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeApplicationId** | **int64** |  | 
**phase** | **string** |  | 
**ruleId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeApplicationsEdgeApplicationIdRulesEnginePhaseRulesRuleIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accept** | **string** |  | 
 **contentType** | **string** | The type of coding used in the Body (application/json). &lt;br&gt;  Example: Content-Type: application/json | 
 **updateRulesEngineRequest** | [**UpdateRulesEngineRequest**](UpdateRulesEngineRequest.md) |  | 

### Return type

[**RulesEngineIdResponse**](RulesEngineIdResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

