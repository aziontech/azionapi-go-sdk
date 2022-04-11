# UpdateDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cnames** | Pointer to **[]string** |  | [optional] 
**CnameAccessOnly** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**EdgeApplicationId** | Pointer to **int64** |  | [optional] 
**DigitalCertificateId** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewUpdateDomainRequest

`func NewUpdateDomainRequest() *UpdateDomainRequest`

NewUpdateDomainRequest instantiates a new UpdateDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomainRequestWithDefaults

`func NewUpdateDomainRequestWithDefaults() *UpdateDomainRequest`

NewUpdateDomainRequestWithDefaults instantiates a new UpdateDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCnames

`func (o *UpdateDomainRequest) GetCnames() []string`

GetCnames returns the Cnames field if non-nil, zero value otherwise.

### GetCnamesOk

`func (o *UpdateDomainRequest) GetCnamesOk() (*[]string, bool)`

GetCnamesOk returns a tuple with the Cnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnames

`func (o *UpdateDomainRequest) SetCnames(v []string)`

SetCnames sets Cnames field to given value.

### HasCnames

`func (o *UpdateDomainRequest) HasCnames() bool`

HasCnames returns a boolean if a field has been set.

### GetCnameAccessOnly

`func (o *UpdateDomainRequest) GetCnameAccessOnly() bool`

GetCnameAccessOnly returns the CnameAccessOnly field if non-nil, zero value otherwise.

### GetCnameAccessOnlyOk

`func (o *UpdateDomainRequest) GetCnameAccessOnlyOk() (*bool, bool)`

GetCnameAccessOnlyOk returns a tuple with the CnameAccessOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameAccessOnly

`func (o *UpdateDomainRequest) SetCnameAccessOnly(v bool)`

SetCnameAccessOnly sets CnameAccessOnly field to given value.

### HasCnameAccessOnly

`func (o *UpdateDomainRequest) HasCnameAccessOnly() bool`

HasCnameAccessOnly returns a boolean if a field has been set.

### GetName

`func (o *UpdateDomainRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDomainRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDomainRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDomainRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateDomainRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateDomainRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateDomainRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateDomainRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEdgeApplicationId

`func (o *UpdateDomainRequest) GetEdgeApplicationId() int64`

GetEdgeApplicationId returns the EdgeApplicationId field if non-nil, zero value otherwise.

### GetEdgeApplicationIdOk

`func (o *UpdateDomainRequest) GetEdgeApplicationIdOk() (*int64, bool)`

GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplicationId

`func (o *UpdateDomainRequest) SetEdgeApplicationId(v int64)`

SetEdgeApplicationId sets EdgeApplicationId field to given value.

### HasEdgeApplicationId

`func (o *UpdateDomainRequest) HasEdgeApplicationId() bool`

HasEdgeApplicationId returns a boolean if a field has been set.

### GetDigitalCertificateId

`func (o *UpdateDomainRequest) GetDigitalCertificateId() int64`

GetDigitalCertificateId returns the DigitalCertificateId field if non-nil, zero value otherwise.

### GetDigitalCertificateIdOk

`func (o *UpdateDomainRequest) GetDigitalCertificateIdOk() (*int64, bool)`

GetDigitalCertificateIdOk returns a tuple with the DigitalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalCertificateId

`func (o *UpdateDomainRequest) SetDigitalCertificateId(v int64)`

SetDigitalCertificateId sets DigitalCertificateId field to given value.

### HasDigitalCertificateId

`func (o *UpdateDomainRequest) HasDigitalCertificateId() bool`

HasDigitalCertificateId returns a boolean if a field has been set.

### SetDigitalCertificateIdNil

`func (o *UpdateDomainRequest) SetDigitalCertificateIdNil(b bool)`

 SetDigitalCertificateIdNil sets the value for DigitalCertificateId to be an explicit nil

### UnsetDigitalCertificateId
`func (o *UpdateDomainRequest) UnsetDigitalCertificateId()`

UnsetDigitalCertificateId ensures that no value is present for DigitalCertificateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


