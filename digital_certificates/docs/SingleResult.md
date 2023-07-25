# SingleResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SubjectName** | Pointer to **[]string** |  | [optional] 
**Issuer** | Pointer to **NullableString** |  | [optional] 
**Validity** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CertificateType** | Pointer to **string** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**Csr** | Pointer to **NullableString** |  | [optional] 
**CertificateContent** | Pointer to **NullableString** |  | [optional] 
**AzionInformation** | Pointer to **string** |  | [optional] 

## Methods

### NewSingleResult

`func NewSingleResult() *SingleResult`

NewSingleResult instantiates a new SingleResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleResultWithDefaults

`func NewSingleResultWithDefaults() *SingleResult`

NewSingleResultWithDefaults instantiates a new SingleResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SingleResult) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SingleResult) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SingleResult) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SingleResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SingleResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SingleResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SingleResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SingleResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubjectName

`func (o *SingleResult) GetSubjectName() []string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *SingleResult) GetSubjectNameOk() (*[]string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *SingleResult) SetSubjectName(v []string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *SingleResult) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetIssuer

`func (o *SingleResult) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SingleResult) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SingleResult) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SingleResult) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *SingleResult) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *SingleResult) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetValidity

`func (o *SingleResult) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *SingleResult) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *SingleResult) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *SingleResult) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### SetValidityNil

`func (o *SingleResult) SetValidityNil(b bool)`

 SetValidityNil sets the value for Validity to be an explicit nil

### UnsetValidity
`func (o *SingleResult) UnsetValidity()`

UnsetValidity ensures that no value is present for Validity, not even an explicit nil
### GetStatus

`func (o *SingleResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SingleResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SingleResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SingleResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCertificateType

`func (o *SingleResult) GetCertificateType() string`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *SingleResult) GetCertificateTypeOk() (*string, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *SingleResult) SetCertificateType(v string)`

SetCertificateType sets CertificateType field to given value.

### HasCertificateType

`func (o *SingleResult) HasCertificateType() bool`

HasCertificateType returns a boolean if a field has been set.

### GetManaged

`func (o *SingleResult) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *SingleResult) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *SingleResult) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *SingleResult) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetCsr

`func (o *SingleResult) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *SingleResult) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *SingleResult) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *SingleResult) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### SetCsrNil

`func (o *SingleResult) SetCsrNil(b bool)`

 SetCsrNil sets the value for Csr to be an explicit nil

### UnsetCsr
`func (o *SingleResult) UnsetCsr()`

UnsetCsr ensures that no value is present for Csr, not even an explicit nil
### GetCertificateContent

`func (o *SingleResult) GetCertificateContent() string`

GetCertificateContent returns the CertificateContent field if non-nil, zero value otherwise.

### GetCertificateContentOk

`func (o *SingleResult) GetCertificateContentOk() (*string, bool)`

GetCertificateContentOk returns a tuple with the CertificateContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateContent

`func (o *SingleResult) SetCertificateContent(v string)`

SetCertificateContent sets CertificateContent field to given value.

### HasCertificateContent

`func (o *SingleResult) HasCertificateContent() bool`

HasCertificateContent returns a boolean if a field has been set.

### SetCertificateContentNil

`func (o *SingleResult) SetCertificateContentNil(b bool)`

 SetCertificateContentNil sets the value for CertificateContent to be an explicit nil

### UnsetCertificateContent
`func (o *SingleResult) UnsetCertificateContent()`

UnsetCertificateContent ensures that no value is present for CertificateContent, not even an explicit nil
### GetAzionInformation

`func (o *SingleResult) GetAzionInformation() string`

GetAzionInformation returns the AzionInformation field if non-nil, zero value otherwise.

### GetAzionInformationOk

`func (o *SingleResult) GetAzionInformationOk() (*string, bool)`

GetAzionInformationOk returns a tuple with the AzionInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzionInformation

`func (o *SingleResult) SetAzionInformation(v string)`

SetAzionInformation sets AzionInformation field to given value.

### HasAzionInformation

`func (o *SingleResult) HasAzionInformation() bool`

HasAzionInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


