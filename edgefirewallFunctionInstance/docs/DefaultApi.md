# {{classname}}

All URIs are relative to *https://api.azionapi.net/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeFirewalledgeFirewallIdFunctionsInstancesGet**](DefaultApi.md#EdgeFirewalledgeFirewallIdFunctionsInstancesGet) | **Get** /edge_firewall/:edge_firewall_id:/functions_instances | List all user Edge Functions Instances
[**EdgeFirewalledgeFirewallIdFunctionsInstancesPost**](DefaultApi.md#EdgeFirewalledgeFirewallIdFunctionsInstancesPost) | **Post** /edge_firewall/:edge_firewall_id:/functions_instances | Create an Edge Functions Instance
[**EdgeFirewalledgeFirewallIdFunctionsInstancesUuidDelete**](DefaultApi.md#EdgeFirewalledgeFirewallIdFunctionsInstancesUuidDelete) | **Delete** /edge_firewall/:edge_firewall_id:/functions_instances/{uuid} | Delete an Edge Functions Instance by uuid
[**EdgeFirewalledgeFirewallIdFunctionsInstancesUuidGet**](DefaultApi.md#EdgeFirewalledgeFirewallIdFunctionsInstancesUuidGet) | **Get** /edge_firewall/:edge_firewall_id:/functions_instances/{uuid} | Retrieve an Edge Functions Instance set by uuid
[**EdgeFirewalledgeFirewallIdFunctionsInstancesUuidPatch**](DefaultApi.md#EdgeFirewalledgeFirewallIdFunctionsInstancesUuidPatch) | **Patch** /edge_firewall/:edge_firewall_id:/functions_instances/{uuid} | Update some Edge Functions Instance attributes
[**EdgeFirewalledgeFirewallIdFunctionsInstancesUuidPut**](DefaultApi.md#EdgeFirewalledgeFirewallIdFunctionsInstancesUuidPut) | **Put** /edge_firewall/:edge_firewall_id:/functions_instances/{uuid} | Overwrite some Edge Functions Instance attributes

# **EdgeFirewalledgeFirewallIdFunctionsInstancesGet**
> ListEdgeFunctionsInstancesResponse EdgeFirewalledgeFirewallIdFunctionsInstancesGet(ctx, optional)
List all user Edge Functions Instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiEdgeFirewalledgeFirewallIdFunctionsInstancesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiEdgeFirewalledgeFirewallIdFunctionsInstancesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **orderBy** | **optional.String**|  | 

### Return type

[**ListEdgeFunctionsInstancesResponse**](ListEdgeFunctionsInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeFirewalledgeFirewallIdFunctionsInstancesPost**
> EdgeFirewalledgeFirewallIdFunctionsInstancesPost(ctx, body)
Create an Edge Functions Instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEdgeFunctionsInstancesRequest**](CreateEdgeFunctionsInstancesRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeFirewalledgeFirewallIdFunctionsInstancesUuidDelete**
> EdgeFirewalledgeFirewallIdFunctionsInstancesUuidDelete(ctx, uuid)
Delete an Edge Functions Instance by uuid

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

# **EdgeFirewalledgeFirewallIdFunctionsInstancesUuidGet**
> EdgeFunctionsInstanceResponse EdgeFirewalledgeFirewallIdFunctionsInstancesUuidGet(ctx, uuid)
Retrieve an Edge Functions Instance set by uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**|  | 

### Return type

[**EdgeFunctionsInstanceResponse**](EdgeFunctionsInstanceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeFirewalledgeFirewallIdFunctionsInstancesUuidPatch**
> ListEdgeFunctionsInstancesResponse EdgeFirewalledgeFirewallIdFunctionsInstancesUuidPatch(ctx, body, uuid)
Update some Edge Functions Instance attributes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEdgeFunctionsInstancesRequest**](CreateEdgeFunctionsInstancesRequest.md)|  | 
  **uuid** | **string**|  | 

### Return type

[**ListEdgeFunctionsInstancesResponse**](ListEdgeFunctionsInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeFirewalledgeFirewallIdFunctionsInstancesUuidPut**
> ListEdgeFunctionsInstancesResponse EdgeFirewalledgeFirewallIdFunctionsInstancesUuidPut(ctx, body, uuid)
Overwrite some Edge Functions Instance attributes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateEdgeFunctionsInstancesRequest**](CreateEdgeFunctionsInstancesRequest.md)|  | 
  **uuid** | **string**|  | 

### Return type

[**ListEdgeFunctionsInstancesResponse**](ListEdgeFunctionsInstancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

