# \VariablesAPI

All URIs are relative to *https://api.azionapi.net/variables*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiVariablesCreate**](VariablesAPI.md#ApiVariablesCreate) | **Post** /variables | /variables
[**ApiVariablesDestroy**](VariablesAPI.md#ApiVariablesDestroy) | **Delete** /variables/{uuid} | /variables/:uuid
[**ApiVariablesList**](VariablesAPI.md#ApiVariablesList) | **Get** /variables | /variables
[**ApiVariablesRetrieve**](VariablesAPI.md#ApiVariablesRetrieve) | **Get** /variables/{uuid} | /variables/:uuid
[**ApiVariablesUpdate**](VariablesAPI.md#ApiVariablesUpdate) | **Put** /variables/{uuid} | /variables/:uuid



## ApiVariablesCreate

> VariableGet ApiVariablesCreate(ctx).VariableCreate(variableCreate).Execute()

/variables



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
    variableCreate := *openapiclient.NewVariableCreate("Key_example", "Value_example") // VariableCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesAPI.ApiVariablesCreate(context.Background()).VariableCreate(variableCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.ApiVariablesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiVariablesCreate`: VariableGet
    fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.ApiVariablesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiVariablesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variableCreate** | [**VariableCreate**](VariableCreate.md) |  | 

### Return type

[**VariableGet**](VariableGet.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiVariablesDestroy

> ApiVariablesDestroy(ctx, uuid).Execute()

/variables/:uuid



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
    r, err := apiClient.VariablesAPI.ApiVariablesDestroy(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.ApiVariablesDestroy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiVariablesDestroyRequest struct via the builder pattern


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


## ApiVariablesList

> []Variable ApiVariablesList(ctx).Execute()

/variables



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesAPI.ApiVariablesList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.ApiVariablesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiVariablesList`: []Variable
    fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.ApiVariablesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiVariablesListRequest struct via the builder pattern


### Return type

[**[]Variable**](Variable.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiVariablesRetrieve

> Variable ApiVariablesRetrieve(ctx, uuid).Execute()

/variables/:uuid



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
    resp, r, err := apiClient.VariablesAPI.ApiVariablesRetrieve(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.ApiVariablesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiVariablesRetrieve`: Variable
    fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.ApiVariablesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiVariablesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Variable**](Variable.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiVariablesUpdate

> VariableGet ApiVariablesUpdate(ctx, uuid).VariableCreate(variableCreate).Execute()

/variables/:uuid



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
    variableCreate := *openapiclient.NewVariableCreate("Key_example", "Value_example") // VariableCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesAPI.ApiVariablesUpdate(context.Background(), uuid).VariableCreate(variableCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesAPI.ApiVariablesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiVariablesUpdate`: VariableGet
    fmt.Fprintf(os.Stdout, "Response from `VariablesAPI.ApiVariablesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiVariablesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variableCreate** | [**VariableCreate**](VariableCreate.md) |  | 

### Return type

[**VariableGet**](VariableGet.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

