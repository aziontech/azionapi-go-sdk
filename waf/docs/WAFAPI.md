# \WAFAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewWAFRuleset**](WAFAPI.md#CreateNewWAFRuleset) | **Post** /waf/rulesets | Create a new WAF Rule Set in an account.
[**DeleteWAFRuleset**](WAFAPI.md#DeleteWAFRuleset) | **Delete** /waf/rulesets/{waf_rule_set_id} | Remove an WAF Rule Set from an account. Warning: this action cannot be undone.
[**GetWAFDomains**](WAFAPI.md#GetWAFDomains) | **Get** /waf/{wafId}/domains | List all domains attached to a Web Application Firewall (WAF) in an account.
[**GetWAFEvents**](WAFAPI.md#GetWAFEvents) | **Get** /waf/{wafId}/waf_events | Find WAF log events
[**GetWAFRuleset**](WAFAPI.md#GetWAFRuleset) | **Get** /waf/rulesets/{waf_rule_set_id} | List a specific Rule Set associated to a Web Application Firewall (WAF) in an account.
[**ListAllWAF**](WAFAPI.md#ListAllWAF) | **Get** /waf | List all Web Application Firewalls (WAFs) created in an account
[**ListAllWAFRulesets**](WAFAPI.md#ListAllWAFRulesets) | **Get** /waf/rulesets | list all Rule Sets associated to a Web Application Firewall (WAF) in an account.
[**UpdateWAFRuleset**](WAFAPI.md#UpdateWAFRuleset) | **Patch** /waf/rulesets/{waf_rule_set_id} | Change only select settings of a WAF Rule Set



## CreateNewWAFRuleset

> SingleWAF CreateNewWAFRuleset(ctx).CreateNewWAFRulesetRequest(createNewWAFRulesetRequest).Execute()

Create a new WAF Rule Set in an account.

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
    createNewWAFRulesetRequest := *openapiclient.NewCreateNewWAFRulesetRequest("Name_example", "Mode_example", false, false, openapiclient.WAFSensitivityChoices("lowest"), false, openapiclient.WAFSensitivityChoices("lowest"), false, openapiclient.WAFSensitivityChoices("lowest"), false, openapiclient.WAFSensitivityChoices("lowest"), false, openapiclient.WAFSensitivityChoices("lowest"), false, openapiclient.WAFSensitivityChoices("lowest"), false, openapiclient.WAFSensitivityChoices("lowest"), false, openapiclient.WAFSensitivityChoices("lowest"), []string{"BypassAddresses_example"}) // CreateNewWAFRulesetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFAPI.CreateNewWAFRuleset(context.Background()).CreateNewWAFRulesetRequest(createNewWAFRulesetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFAPI.CreateNewWAFRuleset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewWAFRuleset`: SingleWAF
    fmt.Fprintf(os.Stdout, "Response from `WAFAPI.CreateNewWAFRuleset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewWAFRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNewWAFRulesetRequest** | [**CreateNewWAFRulesetRequest**](CreateNewWAFRulesetRequest.md) |  | 

### Return type

[**SingleWAF**](SingleWAF.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWAFRuleset

> DeleteWAFRuleset(ctx, wafRuleSetId).Execute()

Remove an WAF Rule Set from an account. Warning: this action cannot be undone.

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
    wafRuleSetId := "wafRuleSetId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WAFAPI.DeleteWAFRuleset(context.Background(), wafRuleSetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFAPI.DeleteWAFRuleset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafRuleSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWAFRulesetRequest struct via the builder pattern


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


## GetWAFDomains

> WAFDomains200 GetWAFDomains(ctx, wafId).Name(name).Execute()

List all domains attached to a Web Application Firewall (WAF) in an account.

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
    wafId := int64(789) // int64 | ID of WAF to return
    name := "name_example" // string | searches WAF for name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFAPI.GetWAFDomains(context.Background(), wafId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFAPI.GetWAFDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWAFDomains`: WAFDomains200
    fmt.Fprintf(os.Stdout, "Response from `WAFAPI.GetWAFDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **int64** | ID of WAF to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWAFDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | searches WAF for name | 

### Return type

[**WAFDomains200**](WAFDomains200.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWAFEvents

> WAFEvents200 GetWAFEvents(ctx, wafId).HourRange(hourRange).DomainsIds(domainsIds).NetworkListId(networkListId).Execute()

Find WAF log events

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
    wafId := int64(789) // int64 | ID of WAF to return
    hourRange := int64(789) // int64 | Last log hours since now (it must be a integer number ranging between 1 and 72)
    domainsIds := "domainsIds_example" // string | Multiple domain's id (they must be separated by comma like 1233,1234)
    networkListId := int64(789) // int64 | Id of a network list (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFAPI.GetWAFEvents(context.Background(), wafId).HourRange(hourRange).DomainsIds(domainsIds).NetworkListId(networkListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFAPI.GetWAFEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWAFEvents`: WAFEvents200
    fmt.Fprintf(os.Stdout, "Response from `WAFAPI.GetWAFEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafId** | **int64** | ID of WAF to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWAFEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hourRange** | **int64** | Last log hours since now (it must be a integer number ranging between 1 and 72) | 
 **domainsIds** | **string** | Multiple domain&#39;s id (they must be separated by comma like 1233,1234) | 
 **networkListId** | **int64** | Id of a network list | 

### Return type

[**WAFEvents200**](WAFEvents200.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWAFRuleset

> WAFSingle200 GetWAFRuleset(ctx, wafRuleSetId).Execute()

List a specific Rule Set associated to a Web Application Firewall (WAF) in an account.

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
    wafRuleSetId := int64(789) // int64 | ID of WAF Ruleset to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFAPI.GetWAFRuleset(context.Background(), wafRuleSetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFAPI.GetWAFRuleset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWAFRuleset`: WAFSingle200
    fmt.Fprintf(os.Stdout, "Response from `WAFAPI.GetWAFRuleset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafRuleSetId** | **int64** | ID of WAF Ruleset to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWAFRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WAFSingle200**](WAFSingle200.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllWAF

> WAFList200 ListAllWAF(ctx).Page(page).PageSize(pageSize).Execute()

List all Web Application Firewalls (WAFs) created in an account

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
    page := int64(789) // int64 | Identifies which page should be returned, if the return is paginated. (optional) (default to 1)
    pageSize := int64(789) // int64 | Identifies how many items should be returned per page. (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFAPI.ListAllWAF(context.Background()).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFAPI.ListAllWAF``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllWAF`: WAFList200
    fmt.Fprintf(os.Stdout, "Response from `WAFAPI.ListAllWAF`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllWAFRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Identifies which page should be returned, if the return is paginated. | [default to 1]
 **pageSize** | **int64** | Identifies how many items should be returned per page. | [default to 10]

### Return type

[**WAFList200**](WAFList200.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllWAFRulesets

> WAFList200 ListAllWAFRulesets(ctx).OrderBy(orderBy).Sort(sort).Page(page).PageSize(pageSize).Execute()

list all Rule Sets associated to a Web Application Firewall (WAF) in an account.

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
    orderBy := "orderBy_example" // string | Identifies which property the return should be sorted by. (optional) (default to "name")
    sort := "sort_example" // string | Defines whether objects are shown in ascending or descending order depending on the value set in order_by. (optional) (default to "asc")
    page := int64(789) // int64 | Identifies which page should be returned, if the return is paginated. (optional) (default to 1)
    pageSize := int64(789) // int64 | Identifies how many items should be returned per page. (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFAPI.ListAllWAFRulesets(context.Background()).OrderBy(orderBy).Sort(sort).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFAPI.ListAllWAFRulesets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllWAFRulesets`: WAFList200
    fmt.Fprintf(os.Stdout, "Response from `WAFAPI.ListAllWAFRulesets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllWAFRulesetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **string** | Identifies which property the return should be sorted by. | [default to &quot;name&quot;]
 **sort** | **string** | Defines whether objects are shown in ascending or descending order depending on the value set in order_by. | [default to &quot;asc&quot;]
 **page** | **int64** | Identifies which page should be returned, if the return is paginated. | [default to 1]
 **pageSize** | **int64** | Identifies how many items should be returned per page. | [default to 10]

### Return type

[**WAFList200**](WAFList200.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWAFRuleset

> SingleWAF UpdateWAFRuleset(ctx, wafRuleSetId).SingleWAF(singleWAF).Execute()

Change only select settings of a WAF Rule Set

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
    wafRuleSetId := "wafRuleSetId_example" // string | 
    singleWAF := *openapiclient.NewSingleWAF() // SingleWAF |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WAFAPI.UpdateWAFRuleset(context.Background(), wafRuleSetId).SingleWAF(singleWAF).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WAFAPI.UpdateWAFRuleset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWAFRuleset`: SingleWAF
    fmt.Fprintf(os.Stdout, "Response from `WAFAPI.UpdateWAFRuleset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wafRuleSetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWAFRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **singleWAF** | [**SingleWAF**](SingleWAF.md) |  | 

### Return type

[**SingleWAF**](SingleWAF.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json; version=3
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

