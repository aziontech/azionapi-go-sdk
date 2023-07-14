# PersonalTokenResponseGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPersonalTokenResponseGet

`func NewPersonalTokenResponseGet() *PersonalTokenResponseGet`

NewPersonalTokenResponseGet instantiates a new PersonalTokenResponseGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalTokenResponseGetWithDefaults

`func NewPersonalTokenResponseGetWithDefaults() *PersonalTokenResponseGet`

NewPersonalTokenResponseGetWithDefaults instantiates a new PersonalTokenResponseGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PersonalTokenResponseGet) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PersonalTokenResponseGet) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PersonalTokenResponseGet) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PersonalTokenResponseGet) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *PersonalTokenResponseGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PersonalTokenResponseGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PersonalTokenResponseGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PersonalTokenResponseGet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreated

`func (o *PersonalTokenResponseGet) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PersonalTokenResponseGet) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PersonalTokenResponseGet) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PersonalTokenResponseGet) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PersonalTokenResponseGet) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PersonalTokenResponseGet) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PersonalTokenResponseGet) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PersonalTokenResponseGet) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetDescription

`func (o *PersonalTokenResponseGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PersonalTokenResponseGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PersonalTokenResponseGet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PersonalTokenResponseGet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PersonalTokenResponseGet) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PersonalTokenResponseGet) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


