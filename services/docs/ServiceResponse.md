# ServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**BoundNodes** | **int64** |  | 
**Id** | **int64** |  | 
**LastEditor** | **string** |  | 
**Name** | **string** |  | 
**Permissions** | **[]string** |  | 
**UpdatedAt** | **string** |  | 
**Variables** | Pointer to [**[]Variable**](Variable.md) |  | [optional] 

## Methods

### NewServiceResponse

`func NewServiceResponse(active bool, boundNodes int64, id int64, lastEditor string, name string, permissions []string, updatedAt string, ) *ServiceResponse`

NewServiceResponse instantiates a new ServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceResponseWithDefaults

`func NewServiceResponseWithDefaults() *ServiceResponse`

NewServiceResponseWithDefaults instantiates a new ServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ServiceResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ServiceResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ServiceResponse) SetActive(v bool)`

SetActive sets Active field to given value.


### GetBoundNodes

`func (o *ServiceResponse) GetBoundNodes() int64`

GetBoundNodes returns the BoundNodes field if non-nil, zero value otherwise.

### GetBoundNodesOk

`func (o *ServiceResponse) GetBoundNodesOk() (*int64, bool)`

GetBoundNodesOk returns a tuple with the BoundNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundNodes

`func (o *ServiceResponse) SetBoundNodes(v int64)`

SetBoundNodes sets BoundNodes field to given value.


### GetId

`func (o *ServiceResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceResponse) SetId(v int64)`

SetId sets Id field to given value.


### GetLastEditor

`func (o *ServiceResponse) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ServiceResponse) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ServiceResponse) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetName

`func (o *ServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *ServiceResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ServiceResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ServiceResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetUpdatedAt

`func (o *ServiceResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVariables

`func (o *ServiceResponse) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ServiceResponse) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ServiceResponse) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ServiceResponse) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


