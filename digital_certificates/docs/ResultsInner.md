# ResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SubjectName** | Pointer to **[]string** |  | [optional] 
**Validity** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CertificateType** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**Issuer** | Pointer to **NullableString** |  | [optional] 
**AzionInformation** | Pointer to **string** |  | [optional] 

## Methods

### NewResultsInner

`func NewResultsInner() *ResultsInner`

NewResultsInner instantiates a new ResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultsInnerWithDefaults

`func NewResultsInnerWithDefaults() *ResultsInner`

NewResultsInnerWithDefaults instantiates a new ResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResultsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResultsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResultsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResultsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResultsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubjectName

`func (o *ResultsInner) GetSubjectName() []string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *ResultsInner) GetSubjectNameOk() (*[]string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *ResultsInner) SetSubjectName(v []string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *ResultsInner) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetValidity

`func (o *ResultsInner) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *ResultsInner) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *ResultsInner) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *ResultsInner) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### SetValidityNil

`func (o *ResultsInner) SetValidityNil(b bool)`

 SetValidityNil sets the value for Validity to be an explicit nil

### UnsetValidity
`func (o *ResultsInner) UnsetValidity()`

UnsetValidity ensures that no value is present for Validity, not even an explicit nil
### GetStatus

`func (o *ResultsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResultsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResultsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResultsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCertificateType

`func (o *ResultsInner) GetCertificateType() string`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *ResultsInner) GetCertificateTypeOk() (*string, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *ResultsInner) SetCertificateType(v string)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *ResultsInner) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.

### GetManaged

`func (o *ResultsInner) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ResultsInner) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ResultsInner) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ResultsInner) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetIssuer

`func (o *ResultsInner) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ResultsInner) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ResultsInner) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ResultsInner) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *ResultsInner) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *ResultsInner) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetAzionInformation

`func (o *ResultsInner) GetAzionInformation() string`

GetAzionInformation returns the AzionInformation field if non-nil, zero value otherwise.

### GetAzionInformationOk

`func (o *ResultsInner) GetAzionInformationOk() (*string, bool)`

GetAzionInformationOk returns a tuple with the AzionInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzionInformation

`func (o *ResultsInner) SetAzionInformation(v string)`

SetAzionInformation sets AzionInformation field to given value.

### HasAzionInformation

`func (o *ResultsInner) HasAzionInformation() bool`

HasAzionInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


