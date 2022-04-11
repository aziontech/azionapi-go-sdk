# DomainResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Cnames** | Pointer to **[]string** |  | [optional] 
**CnameAccessOnly** | Pointer to **bool** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**EdgeApplicationId** | Pointer to **int64** |  | [optional] 
**DigitalCertificateId** | Pointer to **NullableInt64** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainResults

`func NewDomainResults(id int64, name string, ) *DomainResults`

NewDomainResults instantiates a new DomainResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainResultsWithDefaults

`func NewDomainResultsWithDefaults() *DomainResults`

NewDomainResultsWithDefaults instantiates a new DomainResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DomainResults) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DomainResults) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DomainResults) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *DomainResults) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainResults) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainResults) SetName(v string)`

SetName sets Name field to given value.


### GetCnames

`func (o *DomainResults) GetCnames() []string`

GetCnames returns the Cnames field if non-nil, zero value otherwise.

### GetCnamesOk

`func (o *DomainResults) GetCnamesOk() (*[]string, bool)`

GetCnamesOk returns a tuple with the Cnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnames

`func (o *DomainResults) SetCnames(v []string)`

SetCnames sets Cnames field to given value.

### HasCnames

`func (o *DomainResults) HasCnames() bool`

HasCnames returns a boolean if a field has been set.

### GetCnameAccessOnly

`func (o *DomainResults) GetCnameAccessOnly() bool`

GetCnameAccessOnly returns the CnameAccessOnly field if non-nil, zero value otherwise.

### GetCnameAccessOnlyOk

`func (o *DomainResults) GetCnameAccessOnlyOk() (*bool, bool)`

GetCnameAccessOnlyOk returns a tuple with the CnameAccessOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameAccessOnly

`func (o *DomainResults) SetCnameAccessOnly(v bool)`

SetCnameAccessOnly sets CnameAccessOnly field to given value.

### HasCnameAccessOnly

`func (o *DomainResults) HasCnameAccessOnly() bool`

HasCnameAccessOnly returns a boolean if a field has been set.

### GetIsActive

`func (o *DomainResults) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DomainResults) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DomainResults) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DomainResults) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetEdgeApplicationId

`func (o *DomainResults) GetEdgeApplicationId() int64`

GetEdgeApplicationId returns the EdgeApplicationId field if non-nil, zero value otherwise.

### GetEdgeApplicationIdOk

`func (o *DomainResults) GetEdgeApplicationIdOk() (*int64, bool)`

GetEdgeApplicationIdOk returns a tuple with the EdgeApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeApplicationId

`func (o *DomainResults) SetEdgeApplicationId(v int64)`

SetEdgeApplicationId sets EdgeApplicationId field to given value.

### HasEdgeApplicationId

`func (o *DomainResults) HasEdgeApplicationId() bool`

HasEdgeApplicationId returns a boolean if a field has been set.

### GetDigitalCertificateId

`func (o *DomainResults) GetDigitalCertificateId() int64`

GetDigitalCertificateId returns the DigitalCertificateId field if non-nil, zero value otherwise.

### GetDigitalCertificateIdOk

`func (o *DomainResults) GetDigitalCertificateIdOk() (*int64, bool)`

GetDigitalCertificateIdOk returns a tuple with the DigitalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalCertificateId

`func (o *DomainResults) SetDigitalCertificateId(v int64)`

SetDigitalCertificateId sets DigitalCertificateId field to given value.

### HasDigitalCertificateId

`func (o *DomainResults) HasDigitalCertificateId() bool`

HasDigitalCertificateId returns a boolean if a field has been set.

### SetDigitalCertificateIdNil

`func (o *DomainResults) SetDigitalCertificateIdNil(b bool)`

 SetDigitalCertificateIdNil sets the value for DigitalCertificateId to be an explicit nil

### UnsetDigitalCertificateId
`func (o *DomainResults) UnsetDigitalCertificateId()`

UnsetDigitalCertificateId ensures that no value is present for DigitalCertificateId, not even an explicit nil
### GetDomainName

`func (o *DomainResults) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DomainResults) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DomainResults) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DomainResults) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetEnvironment

`func (o *DomainResults) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DomainResults) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DomainResults) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DomainResults) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


