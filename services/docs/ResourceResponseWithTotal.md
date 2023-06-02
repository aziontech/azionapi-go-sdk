# ResourceResponseWithTotal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | [**[]ResourceResponse**](ResourceResponse.md) |  | 
**Total** | **int64** |  | 

## Methods

### NewResourceResponseWithTotal

`func NewResourceResponseWithTotal(resources []ResourceResponse, total int64, ) *ResourceResponseWithTotal`

NewResourceResponseWithTotal instantiates a new ResourceResponseWithTotal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceResponseWithTotalWithDefaults

`func NewResourceResponseWithTotalWithDefaults() *ResourceResponseWithTotal`

NewResourceResponseWithTotalWithDefaults instantiates a new ResourceResponseWithTotal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *ResourceResponseWithTotal) GetResources() []ResourceResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ResourceResponseWithTotal) GetResourcesOk() (*[]ResourceResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ResourceResponseWithTotal) SetResources(v []ResourceResponse)`

SetResources sets Resources field to given value.


### GetTotal

`func (o *ResourceResponseWithTotal) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResourceResponseWithTotal) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResourceResponseWithTotal) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


