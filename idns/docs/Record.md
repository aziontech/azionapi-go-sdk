# Record

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordId** | Pointer to **int32** |  | [optional] 
**Entry** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AnswersList** | Pointer to **[]string** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 

## Methods

### NewRecord

`func NewRecord() *Record`

NewRecord instantiates a new Record object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordWithDefaults

`func NewRecordWithDefaults() *Record`

NewRecordWithDefaults instantiates a new Record object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordId

`func (o *Record) GetRecordId() int32`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *Record) GetRecordIdOk() (*int32, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *Record) SetRecordId(v int32)`

SetRecordId sets RecordId field to given value.

### HasRecordId

`func (o *Record) HasRecordId() bool`

HasRecordId returns a boolean if a field has been set.

### GetEntry

`func (o *Record) GetEntry() string`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *Record) GetEntryOk() (*string, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *Record) SetEntry(v string)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *Record) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetDescription

`func (o *Record) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Record) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Record) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Record) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAnswersList

`func (o *Record) GetAnswersList() []string`

GetAnswersList returns the AnswersList field if non-nil, zero value otherwise.

### GetAnswersListOk

`func (o *Record) GetAnswersListOk() (*[]string, bool)`

GetAnswersListOk returns a tuple with the AnswersList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswersList

`func (o *Record) SetAnswersList(v []string)`

SetAnswersList sets AnswersList field to given value.

### HasAnswersList

`func (o *Record) HasAnswersList() bool`

HasAnswersList returns a boolean if a field has been set.

### GetPolicy

`func (o *Record) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *Record) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *Record) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *Record) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRecordType

`func (o *Record) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *Record) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *Record) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *Record) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetTtl

`func (o *Record) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Record) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Record) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Record) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


