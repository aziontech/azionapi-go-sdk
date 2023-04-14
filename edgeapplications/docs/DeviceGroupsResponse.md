# DeviceGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** |  | 
**TotalPages** | **int64** |  | 
**SchemaVersion** | **int64** |  | 
**Links** | [**DeviceGroupsResponseLinks**](DeviceGroupsResponseLinks.md) |  | 
**Results** | [**[]DeviceGroupsResultResponse**](DeviceGroupsResultResponse.md) |  | 

## Methods

### NewDeviceGroupsResponse

`func NewDeviceGroupsResponse(count int64, totalPages int64, schemaVersion int64, links DeviceGroupsResponseLinks, results []DeviceGroupsResultResponse, ) *DeviceGroupsResponse`

NewDeviceGroupsResponse instantiates a new DeviceGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceGroupsResponseWithDefaults

`func NewDeviceGroupsResponseWithDefaults() *DeviceGroupsResponse`

NewDeviceGroupsResponseWithDefaults instantiates a new DeviceGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DeviceGroupsResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DeviceGroupsResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DeviceGroupsResponse) SetCount(v int64)`

SetCount sets Count field to given value.


### GetTotalPages

`func (o *DeviceGroupsResponse) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *DeviceGroupsResponse) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *DeviceGroupsResponse) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.


### GetSchemaVersion

`func (o *DeviceGroupsResponse) GetSchemaVersion() int64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *DeviceGroupsResponse) GetSchemaVersionOk() (*int64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *DeviceGroupsResponse) SetSchemaVersion(v int64)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetLinks

`func (o *DeviceGroupsResponse) GetLinks() DeviceGroupsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceGroupsResponse) GetLinksOk() (*DeviceGroupsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceGroupsResponse) SetLinks(v DeviceGroupsResponseLinks)`

SetLinks sets Links field to given value.


### GetResults

`func (o *DeviceGroupsResponse) GetResults() []DeviceGroupsResultResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *DeviceGroupsResponse) GetResultsOk() (*[]DeviceGroupsResultResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *DeviceGroupsResponse) SetResults(v []DeviceGroupsResultResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


