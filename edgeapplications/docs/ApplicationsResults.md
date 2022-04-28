# ApplicationsResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Next** | **string** |  | 
**Active** | Pointer to **bool** |  | [optional] 
**Links** | Pointer to [**ApplicationOrigins**](ApplicationOrigins.md) |  | [optional] 

## Methods

### NewApplicationsResults

`func NewApplicationsResults(next string, ) *ApplicationsResults`

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

### HasId

`func (o *ApplicationsResults) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNext

`func (o *ApplicationsResults) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ApplicationsResults) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ApplicationsResults) SetNext(v string)`

SetNext sets Next field to given value.


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

### HasActive

`func (o *ApplicationsResults) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLinks

`func (o *ApplicationsResults) GetLinks() ApplicationOrigins`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationsResults) GetLinksOk() (*ApplicationOrigins, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationsResults) SetLinks(v ApplicationOrigins)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationsResults) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


