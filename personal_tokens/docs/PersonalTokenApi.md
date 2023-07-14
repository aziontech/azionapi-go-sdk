# \PersonalTokenApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePersonalToken**](PersonalTokenApi.md#CreatePersonalToken) | **Post** /iam/personal_tokens | Create a new personal token
[**DeletePersonalToken**](PersonalTokenApi.md#DeletePersonalToken) | **Delete** /iam/personal_tokens/{personalTokenId} | Delete a personal token by id
[**GetPersonalToken**](PersonalTokenApi.md#GetPersonalToken) | **Get** /iam/personal_tokens/{personalTokenId} | Get a personal token info
[**ListPersonalToken**](PersonalTokenApi.md#ListPersonalToken) | **Get** /iam/personal_tokens | List all existing personal token



## CreatePersonalToken

> CreatePersonalTokenResponse CreatePersonalToken(ctx).CreatePersonalTokenRequest(createPersonalTokenRequest).Execute()

Create a new personal token



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
    createPersonalTokenRequest := *openapiclient.NewCreatePersonalTokenRequest() // CreatePersonalTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonalTokenApi.CreatePersonalToken(context.Background()).CreatePersonalTokenRequest(createPersonalTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalTokenApi.CreatePersonalToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePersonalToken`: CreatePersonalTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `PersonalTokenApi.CreatePersonalToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePersonalTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPersonalTokenRequest** | [**CreatePersonalTokenRequest**](CreatePersonalTokenRequest.md) |  | 

### Return type

[**CreatePersonalTokenResponse**](CreatePersonalTokenResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePersonalToken

> DeletePersonalToken(ctx, personalTokenId).Execute()

Delete a personal token by id



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
    personalTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PersonalTokenApi.DeletePersonalToken(context.Background(), personalTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalTokenApi.DeletePersonalToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personalTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonalTokenRequest struct via the builder pattern


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


## GetPersonalToken

> PersonalTokenResponseGet GetPersonalToken(ctx, personalTokenId).Execute()

Get a personal token info



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
    personalTokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonalTokenApi.GetPersonalToken(context.Background(), personalTokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalTokenApi.GetPersonalToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersonalToken`: PersonalTokenResponseGet
    fmt.Fprintf(os.Stdout, "Response from `PersonalTokenApi.GetPersonalToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**personalTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonalTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PersonalTokenResponseGet**](PersonalTokenResponseGet.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPersonalToken

> PersonalTokenResponseWithResults ListPersonalToken(ctx).Execute()

List all existing personal token



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
    resp, r, err := apiClient.PersonalTokenApi.ListPersonalToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalTokenApi.ListPersonalToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPersonalToken`: PersonalTokenResponseWithResults
    fmt.Fprintf(os.Stdout, "Response from `PersonalTokenApi.ListPersonalToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPersonalTokenRequest struct via the builder pattern


### Return type

[**PersonalTokenResponseWithResults**](PersonalTokenResponseWithResults.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

