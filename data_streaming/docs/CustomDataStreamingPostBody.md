# CustomDataStreamingPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DataSource** | Pointer to **NullableString** | Options:  * &#x60;http&#x60; - Edge Applications (default)  * &#x60;waf&#x60; - WAF Events  * &#x60;cells_console&#x60; - Edge Functions  * &#x60;rtm_activity&#x60; - Activity History    | [optional] 
**TemplateModel** | Pointer to **string** | Note:  * Add all variables and values that will be used to stream according to the data source you choose to use.    * All data streaming [variables can be found on the reference documentation](https://www.azion.com/en/documentation/products/data-streaming/#selecting-data-sources).    | [optional] 
**Active** | Pointer to **NullableBool** |  | [optional] [default to true]

## Methods

### NewCustomDataStreamingPostBody

`func NewCustomDataStreamingPostBody() *CustomDataStreamingPostBody`

NewCustomDataStreamingPostBody instantiates a new CustomDataStreamingPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDataStreamingPostBodyWithDefaults

`func NewCustomDataStreamingPostBodyWithDefaults() *CustomDataStreamingPostBody`

NewCustomDataStreamingPostBodyWithDefaults instantiates a new CustomDataStreamingPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomDataStreamingPostBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomDataStreamingPostBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomDataStreamingPostBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomDataStreamingPostBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDataSource

`func (o *CustomDataStreamingPostBody) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *CustomDataStreamingPostBody) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *CustomDataStreamingPostBody) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *CustomDataStreamingPostBody) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### SetDataSourceNil

`func (o *CustomDataStreamingPostBody) SetDataSourceNil(b bool)`

 SetDataSourceNil sets the value for DataSource to be an explicit nil

### UnsetDataSource
`func (o *CustomDataStreamingPostBody) UnsetDataSource()`

UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
### GetTemplateModel

`func (o *CustomDataStreamingPostBody) GetTemplateModel() string`

GetTemplateModel returns the TemplateModel field if non-nil, zero value otherwise.

### GetTemplateModelOk

`func (o *CustomDataStreamingPostBody) GetTemplateModelOk() (*string, bool)`

GetTemplateModelOk returns a tuple with the TemplateModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateModel

`func (o *CustomDataStreamingPostBody) SetTemplateModel(v string)`

SetTemplateModel sets TemplateModel field to given value.

### HasTemplateModel

`func (o *CustomDataStreamingPostBody) HasTemplateModel() bool`

HasTemplateModel returns a boolean if a field has been set.

### GetActive

`func (o *CustomDataStreamingPostBody) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CustomDataStreamingPostBody) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CustomDataStreamingPostBody) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CustomDataStreamingPostBody) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *CustomDataStreamingPostBody) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *CustomDataStreamingPostBody) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


