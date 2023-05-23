# {{classname}}

All URIs are relative to *https://api.azionapi.net/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeFirewallGet**](DefaultApi.md#EdgeFirewallGet) | **Get** /edge_firewall | List all user edge firewall
[**EdgeFirewallPost**](DefaultApi.md#EdgeFirewallPost) | **Post** /edge_firewall | Create a edge firewall
[**EdgeFirewallUuidDelete**](DefaultApi.md#EdgeFirewallUuidDelete) | **Delete** /edge_firewall/{uuid} | Delete an edge firewall by uuid
[**EdgeFirewallUuidGet**](DefaultApi.md#EdgeFirewallUuidGet) | **Get** /edge_firewall/{uuid} | Retrieve an edge firewall set by uuid
[**EdgeFirewallUuidPatch**](DefaultApi.md#EdgeFirewallUuidPatch) | **Patch** /edge_firewall/{uuid} | Update some edge firewall attributes, like \&quot;active\&quot;
[**EdgeFirewallUuidPut**](DefaultApi.md#EdgeFirewallUuidPut) | **Put** /edge_firewall/{uuid} | Overwrite some edge firewall attributes, like \&quot;active\&quot;

# **EdgeFirewallGet**
> ListEdgeFirewallResponse EdgeFirewallGet(ctx, optional)
List all user edge firewall

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiEdgeFirewallGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiEdgeFirewallGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **orderBy** | **optional.String**|  | 

### Return type

[**ListEdgeFirewallResponse**](ListEdgeFirewallResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeFirewallPost**
> EdgeFirewallPost(ctx, body)
Create a edge firewall

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEdgeFirewallRequest**](CreateEdgeFirewallRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeFirewallUuidDelete**
> EdgeFirewallUuidDelete(ctx, uuid)
Delete an edge firewall by uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeFirewallUuidGet**
> EdgeFirewallResponse EdgeFirewallUuidGet(ctx, uuid)
Retrieve an edge firewall set by uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**|  | 

### Return type

[**EdgeFirewallResponse**](EdgeFirewallResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeFirewallUuidPatch**
> ListEdgeFirewallResponse EdgeFirewallUuidPatch(ctx, body, uuid)
Update some edge firewall attributes, like \"active\"

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ListEdgeFirewallResponse**](ListEdgeFirewallResponse.md)|  | 
  **uuid** | **string**|  | 

### Return type

[**ListEdgeFirewallResponse**](ListEdgeFirewallResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeFirewallUuidPut**
> ListEdgeFirewallResponse EdgeFirewallUuidPut(ctx, body, uuid)
Overwrite some edge firewall attributes, like \"active\"

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ListEdgeFirewallResponse**](ListEdgeFirewallResponse.md)|  | 
  **uuid** | **string**|  | 

### Return type

[**ListEdgeFirewallResponse**](ListEdgeFirewallResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

