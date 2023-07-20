# CreateCSRRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Locality** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**OrganizationUnity** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**PrivateKeyType** | Pointer to **string** |  | [optional] 
**Sans** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateCSRRequest

`func NewCreateCSRRequest() *CreateCSRRequest`

NewCreateCSRRequest instantiates a new CreateCSRRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCSRRequestWithDefaults

`func NewCreateCSRRequestWithDefaults() *CreateCSRRequest`

NewCreateCSRRequestWithDefaults instantiates a new CreateCSRRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCSRRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCSRRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCSRRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateCSRRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCommonName

`func (o *CreateCSRRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CreateCSRRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CreateCSRRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CreateCSRRequest) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetCountry

`func (o *CreateCSRRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateCSRRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateCSRRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CreateCSRRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *CreateCSRRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateCSRRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateCSRRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateCSRRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLocality

`func (o *CreateCSRRequest) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *CreateCSRRequest) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *CreateCSRRequest) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *CreateCSRRequest) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetOrganization

`func (o *CreateCSRRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CreateCSRRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CreateCSRRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CreateCSRRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationUnity

`func (o *CreateCSRRequest) GetOrganizationUnity() string`

GetOrganizationUnity returns the OrganizationUnity field if non-nil, zero value otherwise.

### GetOrganizationUnityOk

`func (o *CreateCSRRequest) GetOrganizationUnityOk() (*string, bool)`

GetOrganizationUnityOk returns a tuple with the OrganizationUnity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnity

`func (o *CreateCSRRequest) SetOrganizationUnity(v string)`

SetOrganizationUnity sets OrganizationUnity field to given value.

### HasOrganizationUnity

`func (o *CreateCSRRequest) HasOrganizationUnity() bool`

HasOrganizationUnity returns a boolean if a field has been set.

### GetEmail

`func (o *CreateCSRRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateCSRRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateCSRRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateCSRRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPrivateKeyType

`func (o *CreateCSRRequest) GetPrivateKeyType() string`

GetPrivateKeyType returns the PrivateKeyType field if non-nil, zero value otherwise.

### GetPrivateKeyTypeOk

`func (o *CreateCSRRequest) GetPrivateKeyTypeOk() (*string, bool)`

GetPrivateKeyTypeOk returns a tuple with the PrivateKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyType

`func (o *CreateCSRRequest) SetPrivateKeyType(v string)`

SetPrivateKeyType sets PrivateKeyType field to given value.

### HasPrivateKeyType

`func (o *CreateCSRRequest) HasPrivateKeyType() bool`

HasPrivateKeyType returns a boolean if a field has been set.

### GetSans

`func (o *CreateCSRRequest) GetSans() []string`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *CreateCSRRequest) GetSansOk() (*[]string, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *CreateCSRRequest) SetSans(v []string)`

SetSans sets Sans field to given value.

### HasSans

`func (o *CreateCSRRequest) HasSans() bool`

HasSans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


