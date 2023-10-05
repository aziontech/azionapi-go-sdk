# \RecordsAPI

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteZoneRecord**](RecordsAPI.md#DeleteZoneRecord) | **Delete** /intelligent_dns/{zone_id}/records/{record_id} | Remove an Intelligent DNS zone record
[**GetZoneRecords**](RecordsAPI.md#GetZoneRecords) | **Get** /intelligent_dns/{zone_id}/records | Get a collection of Intelligent DNS zone records
[**PostZoneRecord**](RecordsAPI.md#PostZoneRecord) | **Post** /intelligent_dns/{zone_id}/records | Create a new Intelligent DNS zone record
[**PutZoneRecord**](RecordsAPI.md#PutZoneRecord) | **Put** /intelligent_dns/{zone_id}/records/{record_id} | Update an Intelligent DNS zone record



## DeleteZoneRecord

> string DeleteZoneRecord(ctx, zoneId, recordId).Execute()

Remove an Intelligent DNS zone record

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
    zoneId := int32(56) // int32 | The hosted zone id
    recordId := int32(56) // int32 | The zone record id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordsAPI.DeleteZoneRecord(context.Background(), zoneId, recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.DeleteZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteZoneRecord`: string
    fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.DeleteZoneRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int32** | The hosted zone id | 
**recordId** | **int32** | The zone record id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZoneRecords

> GetRecordsResponse GetZoneRecords(ctx, zoneId).Page(page).PageSize(pageSize).Execute()

Get a collection of Intelligent DNS zone records

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
    zoneId := int32(56) // int32 | The hosted zone id
    page := int64(789) // int64 | Identifies which page should be returned, if the return is paginated. (optional) (default to 1)
    pageSize := int64(789) // int64 | Identifies how many items should be returned per page. (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordsAPI.GetZoneRecords(context.Background(), zoneId).Page(page).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.GetZoneRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZoneRecords`: GetRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.GetZoneRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int32** | The hosted zone id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Identifies which page should be returned, if the return is paginated. | [default to 1]
 **pageSize** | **int64** | Identifies how many items should be returned per page. | [default to 10]

### Return type

[**GetRecordsResponse**](GetRecordsResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostZoneRecord

> PostOrPutRecordResponse PostZoneRecord(ctx, zoneId).RecordPostOrPut(recordPostOrPut).Execute()

Create a new Intelligent DNS zone record

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
    zoneId := int32(56) // int32 | The hosted zone id
    recordPostOrPut := *openapiclient.NewRecordPostOrPut() // RecordPostOrPut |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordsAPI.PostZoneRecord(context.Background(), zoneId).RecordPostOrPut(recordPostOrPut).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.PostZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostZoneRecord`: PostOrPutRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.PostZoneRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int32** | The hosted zone id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordPostOrPut** | [**RecordPostOrPut**](RecordPostOrPut.md) |  | 

### Return type

[**PostOrPutRecordResponse**](PostOrPutRecordResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutZoneRecord

> PostOrPutRecordResponse PutZoneRecord(ctx, zoneId, recordId).RecordPostOrPut(recordPostOrPut).Execute()

Update an Intelligent DNS zone record

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
    zoneId := int32(56) // int32 | The hosted zone id
    recordId := int32(56) // int32 | The zone record id
    recordPostOrPut := *openapiclient.NewRecordPostOrPut() // RecordPostOrPut |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordsAPI.PutZoneRecord(context.Background(), zoneId, recordId).RecordPostOrPut(recordPostOrPut).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsAPI.PutZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutZoneRecord`: PostOrPutRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordsAPI.PutZoneRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **int32** | The hosted zone id | 
**recordId** | **int32** | The zone record id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutZoneRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **recordPostOrPut** | [**RecordPostOrPut**](RecordPostOrPut.md) |  | 

### Return type

[**PostOrPutRecordResponse**](PostOrPutRecordResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; version=3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

