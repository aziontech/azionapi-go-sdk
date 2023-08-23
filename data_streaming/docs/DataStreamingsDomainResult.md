# DataStreamingsDomainResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Selected** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewDataStreamingsDomainResult

`func NewDataStreamingsDomainResult() *DataStreamingsDomainResult`

NewDataStreamingsDomainResult instantiates a new DataStreamingsDomainResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStreamingsDomainResultWithDefaults

`func NewDataStreamingsDomainResultWithDefaults() *DataStreamingsDomainResult`

NewDataStreamingsDomainResultWithDefaults instantiates a new DataStreamingsDomainResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *DataStreamingsDomainResult) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DataStreamingsDomainResult) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DataStreamingsDomainResult) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *DataStreamingsDomainResult) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetName

`func (o *DataStreamingsDomainResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataStreamingsDomainResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataStreamingsDomainResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataStreamingsDomainResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelected

`func (o *DataStreamingsDomainResult) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *DataStreamingsDomainResult) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *DataStreamingsDomainResult) SetSelected(v bool)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *DataStreamingsDomainResult) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### SetSelectedNil

`func (o *DataStreamingsDomainResult) SetSelectedNil(b bool)`

 SetSelectedNil sets the value for Selected to be an explicit nil

### UnsetSelected
`func (o *DataStreamingsDomainResult) UnsetSelected()`

UnsetSelected ensures that no value is present for Selected, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


