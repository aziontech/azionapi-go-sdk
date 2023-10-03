# \DefaultAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeFirewallEdgeFirewallIdRulesEngineGet**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdRulesEngineGet) | **Get** /edge_firewall/{edge_firewall_id}/rules_engine | List all rule sets.
[**EdgeFirewallEdgeFirewallIdRulesEnginePost**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdRulesEnginePost) | **Post** /edge_firewall/{edge_firewall_id}/rules_engine | Create rule set.
[**EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdDelete**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdDelete) | **Delete** /edge_firewall/{edge_firewall_id}/rules_engine/{rule_set_id} | Delete rule set.
[**EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdGet**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdGet) | **Get** /edge_firewall/{edge_firewall_id}/rules_engine/{rule_set_id} | Retrieve rule set by ID.
[**EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPatch**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPatch) | **Patch** /edge_firewall/{edge_firewall_id}/rules_engine/{rule_set_id} | Edit rule set.
[**EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPut**](DefaultAPI.md#EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPut) | **Put** /edge_firewall/{edge_firewall_id}/rules_engine/{rule_set_id} | Overwrite rule set
[**EdgeFirewallGet**](DefaultAPI.md#EdgeFirewallGet) | **Get** /edge_firewall | List all user edge firewall
[**EdgeFirewallPost**](DefaultAPI.md#EdgeFirewallPost) | **Post** /edge_firewall | Create a edge firewall
[**EdgeFirewallUuidDelete**](DefaultAPI.md#EdgeFirewallUuidDelete) | **Delete** /edge_firewall/{uuid} | Delete an edge firewall by uuid
[**EdgeFirewallUuidGet**](DefaultAPI.md#EdgeFirewallUuidGet) | **Get** /edge_firewall/{uuid} | Retrieve an edge firewall set by uuid
[**EdgeFirewallUuidPatch**](DefaultAPI.md#EdgeFirewallUuidPatch) | **Patch** /edge_firewall/{uuid} | Update some edge firewall attributes, like \&quot;active\&quot;
[**EdgeFirewallUuidPut**](DefaultAPI.md#EdgeFirewallUuidPut) | **Put** /edge_firewall/{uuid} | Overwrite some edge firewall attributes, like \&quot;active\&quot;



## EdgeFirewallEdgeFirewallIdRulesEngineGet

> RuleSetResponseAll EdgeFirewallEdgeFirewallIdRulesEngineGet(ctx, edgeFirewallId).Execute()

List all rule sets.

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
    edgeFirewallId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineGet(context.Background(), edgeFirewallId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdRulesEngineGet`: RuleSetResponseAll
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdRulesEngineGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RuleSetResponseAll**](RuleSetResponseAll.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallEdgeFirewallIdRulesEnginePost

> RuleSetResponse EdgeFirewallEdgeFirewallIdRulesEnginePost(ctx, edgeFirewallId).CreateRuleSetRequest(createRuleSetRequest).Execute()

Create rule set.

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
    edgeFirewallId := int64(789) // int64 | 
    createRuleSetRequest := *openapiclient.NewCreateRuleSetRequest() // CreateRuleSetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEnginePost(context.Background(), edgeFirewallId).CreateRuleSetRequest(createRuleSetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEnginePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdRulesEnginePost`: RuleSetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEnginePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdRulesEnginePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRuleSetRequest** | [**CreateRuleSetRequest**](CreateRuleSetRequest.md) |  | 

### Return type

[**RuleSetResponse**](RuleSetResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdDelete

> EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdDelete(ctx, edgeFirewallId, ruleSetId).Execute()

Delete rule set.

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
    edgeFirewallId := int64(789) // int64 | 
    ruleSetId := int64(789) // int64 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdDelete(context.Background(), edgeFirewallId, ruleSetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **int64** |  | 
**ruleSetId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdGet

> RuleSetResult EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdGet(ctx, edgeFirewallId, ruleSetId).OrderBy(orderBy).Sort(sort).Page(page).PageSize(pageSize).Execute()

Retrieve rule set by ID.

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
    edgeFirewallId := int64(789) // int64 | 
    ruleSetId := int64(789) // int64 | 
    orderBy := "orderBy_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    page := int64(789) // int64 |  (optional) (default to 1)
    pageSize := int64(789) // int64 |  (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdGet(context.Background(), edgeFirewallId, ruleSetId).OrderBy(orderBy).Sort(sort).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdGet`: RuleSetResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **int64** |  | 
**ruleSetId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orderBy** | **string** |  | 
 **sort** | **string** |  | 
 **page** | **int64** |  | [default to 1]
 **pageSize** | **int64** |  | [default to 10]

### Return type

[**RuleSetResult**](RuleSetResult.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPatch

> RuleSetResult EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPatch(ctx, edgeFirewallId, ruleSetId).CreateRuleSetRequest(createRuleSetRequest).Execute()

Edit rule set.

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
    edgeFirewallId := int64(789) // int64 | 
    ruleSetId := int64(789) // int64 | 
    createRuleSetRequest := *openapiclient.NewCreateRuleSetRequest() // CreateRuleSetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPatch(context.Background(), edgeFirewallId, ruleSetId).CreateRuleSetRequest(createRuleSetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPatch`: RuleSetResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **int64** |  | 
**ruleSetId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createRuleSetRequest** | [**CreateRuleSetRequest**](CreateRuleSetRequest.md) |  | 

### Return type

[**RuleSetResult**](RuleSetResult.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPut

> RuleSetResult EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPut(ctx, edgeFirewallId, ruleSetId).CreateRuleSetRequest(createRuleSetRequest).Execute()

Overwrite rule set

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
    edgeFirewallId := int64(789) // int64 | 
    ruleSetId := int64(789) // int64 | 
    createRuleSetRequest := *openapiclient.NewCreateRuleSetRequest() // CreateRuleSetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPut(context.Background(), edgeFirewallId, ruleSetId).CreateRuleSetRequest(createRuleSetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPut`: RuleSetResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**edgeFirewallId** | **int64** |  | 
**ruleSetId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallEdgeFirewallIdRulesEngineRuleSetIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createRuleSetRequest** | [**CreateRuleSetRequest**](CreateRuleSetRequest.md) |  | 

### Return type

[**RuleSetResult**](RuleSetResult.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallGet

> ListEdgeFirewallResponse EdgeFirewallGet(ctx).Page(page).PageSize(pageSize).Sort(sort).OrderBy(orderBy).Execute()

List all user edge firewall

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
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallGet(context.Background()).Page(page).PageSize(pageSize).Sort(sort).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallGet`: ListEdgeFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** |  | 
 **pageSize** | **int64** |  | 
 **sort** | **string** |  | 
 **orderBy** | **string** |  | 

### Return type

[**ListEdgeFirewallResponse**](ListEdgeFirewallResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallPost

> EdgeFirewallResponse EdgeFirewallPost(ctx).CreateEdgeFirewallRequest(createEdgeFirewallRequest).Execute()

Create a edge firewall

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
    createEdgeFirewallRequest := *openapiclient.NewCreateEdgeFirewallRequest() // CreateEdgeFirewallRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallPost(context.Background()).CreateEdgeFirewallRequest(createEdgeFirewallRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallPost`: EdgeFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEdgeFirewallRequest** | [**CreateEdgeFirewallRequest**](CreateEdgeFirewallRequest.md) |  | 

### Return type

[**EdgeFirewallResponse**](EdgeFirewallResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallUuidDelete

> EdgeFirewallUuidDelete(ctx, uuid).Execute()

Delete an edge firewall by uuid

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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.EdgeFirewallUuidDelete(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallUuidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallUuidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## EdgeFirewallUuidGet

> EdgeFirewallResponse EdgeFirewallUuidGet(ctx, uuid).Execute()

Retrieve an edge firewall set by uuid

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
    uuid := "uuid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallUuidGet(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallUuidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallUuidGet`: EdgeFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallUuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EdgeFirewallResponse**](EdgeFirewallResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallUuidPatch

> EdgeFirewallResponse EdgeFirewallUuidPatch(ctx, uuid).UpdateEdgeFirewallRequest(updateEdgeFirewallRequest).Execute()

Update some edge firewall attributes, like \"active\"

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
    uuid := "uuid_example" // string | 
    updateEdgeFirewallRequest := *openapiclient.NewUpdateEdgeFirewallRequest() // UpdateEdgeFirewallRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallUuidPatch(context.Background(), uuid).UpdateEdgeFirewallRequest(updateEdgeFirewallRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallUuidPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallUuidPatch`: EdgeFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallUuidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallUuidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEdgeFirewallRequest** | [**UpdateEdgeFirewallRequest**](UpdateEdgeFirewallRequest.md) |  | 

### Return type

[**EdgeFirewallResponse**](EdgeFirewallResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeFirewallUuidPut

> EdgeFirewallResponse EdgeFirewallUuidPut(ctx, uuid).UpdateEdgeFirewallRequest(updateEdgeFirewallRequest).Execute()

Overwrite some edge firewall attributes, like \"active\"

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
    uuid := "uuid_example" // string | 
    updateEdgeFirewallRequest := *openapiclient.NewUpdateEdgeFirewallRequest() // UpdateEdgeFirewallRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.EdgeFirewallUuidPut(context.Background(), uuid).UpdateEdgeFirewallRequest(updateEdgeFirewallRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EdgeFirewallUuidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EdgeFirewallUuidPut`: EdgeFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EdgeFirewallUuidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeFirewallUuidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEdgeFirewallRequest** | [**UpdateEdgeFirewallRequest**](UpdateEdgeFirewallRequest.md) |  | 

### Return type

[**EdgeFirewallResponse**](EdgeFirewallResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

