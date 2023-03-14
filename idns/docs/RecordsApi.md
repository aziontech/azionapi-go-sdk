# \RecordsApi

All URIs are relative to *https://api.azionapi.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteZoneRecord**](RecordsApi.md#DeleteZoneRecord) | **Delete** /intelligent_dns/{zone_id}/records/{record_id} | Remove an Intelligent DNS zone record
[**GetZoneRecords**](RecordsApi.md#GetZoneRecords) | **Get** /intelligent_dns/{zone_id}/records | Get a collection of Intelligent DNS zone records
[**PostZoneRecord**](RecordsApi.md#PostZoneRecord) | **Post** /intelligent_dns/{zone_id}/records | Create a new Intelligent DNS zone record
[**PutZoneRecord**](RecordsApi.md#PutZoneRecord) | **Put** /intelligent_dns/{zone_id}/records/{record_id} | Update an Intelligent DNS zone record



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
    resp, r, err := apiClient.RecordsApi.DeleteZoneRecord(context.Background(), zoneId, recordId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.DeleteZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteZoneRecord`: string
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.DeleteZoneRecord`: %v\n", resp)
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

> GetRecordsResponse GetZoneRecords(ctx, zoneId).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordsApi.GetZoneRecords(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.GetZoneRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetZoneRecords`: GetRecordsResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.GetZoneRecords`: %v\n", resp)
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

> PostOrPutRecordResponse PostZoneRecord(ctx, zoneId).Record(record).Execute()

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
    record := *openapiclient.NewRecord() // Record |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordsApi.PostZoneRecord(context.Background(), zoneId).Record(record).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.PostZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostZoneRecord`: PostOrPutRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.PostZoneRecord`: %v\n", resp)
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

 **record** | [**Record**](Record.md) |  | 

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

> PostOrPutRecordResponse PutZoneRecord(ctx, zoneId, recordId).Record(record).Execute()

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
    record := *openapiclient.NewRecord() // Record |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordsApi.PutZoneRecord(context.Background(), zoneId, recordId).Record(record).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordsApi.PutZoneRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutZoneRecord`: PostOrPutRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `RecordsApi.PutZoneRecord`: %v\n", resp)
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


 **record** | [**Record**](Record.md) |  | 

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

