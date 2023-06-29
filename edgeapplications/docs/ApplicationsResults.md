# ApplicationsResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**DebugRules** | **bool** |  | 
**LastEditor** | **string** |  | 
**LastModified** | **string** |  | 
**Active** | **bool** |  | 
**Origins** | [**[]ApplicationOrigins**](ApplicationOrigins.md) |  | 

## Methods

### NewApplicationsResults

`func NewApplicationsResults(id int64, name string, debugRules bool, lastEditor string, lastModified string, active bool, origins []ApplicationOrigins, ) *ApplicationsResults`

NewApplicationsResults instantiates a new ApplicationsResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationsResultsWithDefaults

`func NewApplicationsResultsWithDefaults() *ApplicationsResults`

NewApplicationsResultsWithDefaults instantiates a new ApplicationsResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationsResults) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationsResults) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationsResults) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ApplicationsResults) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationsResults) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationsResults) SetName(v string)`

SetName sets Name field to given value.


### GetDebugRules

`func (o *ApplicationsResults) GetDebugRules() bool`

GetDebugRules returns the DebugRules field if non-nil, zero value otherwise.

### GetDebugRulesOk

`func (o *ApplicationsResults) GetDebugRulesOk() (*bool, bool)`

GetDebugRulesOk returns a tuple with the DebugRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugRules

`func (o *ApplicationsResults) SetDebugRules(v bool)`

SetDebugRules sets DebugRules field to given value.


### GetLastEditor

`func (o *ApplicationsResults) GetLastEditor() string`

GetLastEditor returns the LastEditor field if non-nil, zero value otherwise.

### GetLastEditorOk

`func (o *ApplicationsResults) GetLastEditorOk() (*string, bool)`

GetLastEditorOk returns a tuple with the LastEditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditor

`func (o *ApplicationsResults) SetLastEditor(v string)`

SetLastEditor sets LastEditor field to given value.


### GetLastModified

`func (o *ApplicationsResults) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ApplicationsResults) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ApplicationsResults) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### GetActive

`func (o *ApplicationsResults) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationsResults) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationsResults) SetActive(v bool)`

SetActive sets Active field to given value.


### GetOrigins

`func (o *ApplicationsResults) GetOrigins() []ApplicationOrigins`

GetOrigins returns the Origins field if non-nil, zero value otherwise.

### GetOriginsOk

`func (o *ApplicationsResults) GetOriginsOk() (*[]ApplicationOrigins, bool)`

GetOriginsOk returns a tuple with the Origins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigins

`func (o *ApplicationsResults) SetOrigins(v []ApplicationOrigins)`

SetOrigins sets Origins field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


