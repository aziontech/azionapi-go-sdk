# SingleWAF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | Identification name for WAF Rule Set. | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**SqlInjection** | Pointer to **bool** |  | [optional] 
**SqlInjectionSensitivity** | Pointer to [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | [optional] 
**RemoteFileInclusion** | Pointer to **bool** |  | [optional] 
**RemoteFileInclusionSensitivity** | Pointer to [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | [optional] 
**DirectoryTraversal** | Pointer to **bool** |  | [optional] 
**DirectoryTraversalSensitivity** | Pointer to [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | [optional] 
**CrossSiteScripting** | Pointer to **bool** |  | [optional] 
**CrossSiteScriptingSensitivity** | Pointer to [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | [optional] 
**EvadingTricks** | Pointer to **bool** |  | [optional] 
**EvadingTricksSensitivity** | Pointer to [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | [optional] 
**FileUpload** | Pointer to **bool** |  | [optional] 
**FileUploadSensitivity** | Pointer to [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | [optional] 
**UnwantedAccess** | Pointer to **bool** |  | [optional] 
**UnwantedAccessSensitivity** | Pointer to [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | [optional] 
**IdentifiedAttack** | Pointer to **bool** |  | [optional] 
**IdentifiedAttackSensitivity** | Pointer to [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | [optional] 
**BypassAddresses** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSingleWAF

`func NewSingleWAF() *SingleWAF`

NewSingleWAF instantiates a new SingleWAF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleWAFWithDefaults

`func NewSingleWAFWithDefaults() *SingleWAF`

NewSingleWAFWithDefaults instantiates a new SingleWAF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SingleWAF) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SingleWAF) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SingleWAF) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SingleWAF) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SingleWAF) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SingleWAF) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SingleWAF) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SingleWAF) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMode

`func (o *SingleWAF) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SingleWAF) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SingleWAF) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SingleWAF) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetActive

`func (o *SingleWAF) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SingleWAF) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SingleWAF) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SingleWAF) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSqlInjection

`func (o *SingleWAF) GetSqlInjection() bool`

GetSqlInjection returns the SqlInjection field if non-nil, zero value otherwise.

### GetSqlInjectionOk

`func (o *SingleWAF) GetSqlInjectionOk() (*bool, bool)`

GetSqlInjectionOk returns a tuple with the SqlInjection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlInjection

`func (o *SingleWAF) SetSqlInjection(v bool)`

SetSqlInjection sets SqlInjection field to given value.

### HasSqlInjection

`func (o *SingleWAF) HasSqlInjection() bool`

HasSqlInjection returns a boolean if a field has been set.

### GetSqlInjectionSensitivity

`func (o *SingleWAF) GetSqlInjectionSensitivity() WAFSensitivityChoices`

GetSqlInjectionSensitivity returns the SqlInjectionSensitivity field if non-nil, zero value otherwise.

### GetSqlInjectionSensitivityOk

`func (o *SingleWAF) GetSqlInjectionSensitivityOk() (*WAFSensitivityChoices, bool)`

GetSqlInjectionSensitivityOk returns a tuple with the SqlInjectionSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlInjectionSensitivity

`func (o *SingleWAF) SetSqlInjectionSensitivity(v WAFSensitivityChoices)`

SetSqlInjectionSensitivity sets SqlInjectionSensitivity field to given value.

### HasSqlInjectionSensitivity

`func (o *SingleWAF) HasSqlInjectionSensitivity() bool`

HasSqlInjectionSensitivity returns a boolean if a field has been set.

### GetRemoteFileInclusion

`func (o *SingleWAF) GetRemoteFileInclusion() bool`

GetRemoteFileInclusion returns the RemoteFileInclusion field if non-nil, zero value otherwise.

### GetRemoteFileInclusionOk

`func (o *SingleWAF) GetRemoteFileInclusionOk() (*bool, bool)`

GetRemoteFileInclusionOk returns a tuple with the RemoteFileInclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFileInclusion

`func (o *SingleWAF) SetRemoteFileInclusion(v bool)`

SetRemoteFileInclusion sets RemoteFileInclusion field to given value.

### HasRemoteFileInclusion

`func (o *SingleWAF) HasRemoteFileInclusion() bool`

HasRemoteFileInclusion returns a boolean if a field has been set.

### GetRemoteFileInclusionSensitivity

`func (o *SingleWAF) GetRemoteFileInclusionSensitivity() WAFSensitivityChoices`

GetRemoteFileInclusionSensitivity returns the RemoteFileInclusionSensitivity field if non-nil, zero value otherwise.

### GetRemoteFileInclusionSensitivityOk

`func (o *SingleWAF) GetRemoteFileInclusionSensitivityOk() (*WAFSensitivityChoices, bool)`

GetRemoteFileInclusionSensitivityOk returns a tuple with the RemoteFileInclusionSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFileInclusionSensitivity

`func (o *SingleWAF) SetRemoteFileInclusionSensitivity(v WAFSensitivityChoices)`

SetRemoteFileInclusionSensitivity sets RemoteFileInclusionSensitivity field to given value.

### HasRemoteFileInclusionSensitivity

`func (o *SingleWAF) HasRemoteFileInclusionSensitivity() bool`

HasRemoteFileInclusionSensitivity returns a boolean if a field has been set.

### GetDirectoryTraversal

`func (o *SingleWAF) GetDirectoryTraversal() bool`

GetDirectoryTraversal returns the DirectoryTraversal field if non-nil, zero value otherwise.

### GetDirectoryTraversalOk

`func (o *SingleWAF) GetDirectoryTraversalOk() (*bool, bool)`

GetDirectoryTraversalOk returns a tuple with the DirectoryTraversal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTraversal

`func (o *SingleWAF) SetDirectoryTraversal(v bool)`

SetDirectoryTraversal sets DirectoryTraversal field to given value.

### HasDirectoryTraversal

`func (o *SingleWAF) HasDirectoryTraversal() bool`

HasDirectoryTraversal returns a boolean if a field has been set.

### GetDirectoryTraversalSensitivity

`func (o *SingleWAF) GetDirectoryTraversalSensitivity() WAFSensitivityChoices`

GetDirectoryTraversalSensitivity returns the DirectoryTraversalSensitivity field if non-nil, zero value otherwise.

### GetDirectoryTraversalSensitivityOk

`func (o *SingleWAF) GetDirectoryTraversalSensitivityOk() (*WAFSensitivityChoices, bool)`

GetDirectoryTraversalSensitivityOk returns a tuple with the DirectoryTraversalSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTraversalSensitivity

`func (o *SingleWAF) SetDirectoryTraversalSensitivity(v WAFSensitivityChoices)`

SetDirectoryTraversalSensitivity sets DirectoryTraversalSensitivity field to given value.

### HasDirectoryTraversalSensitivity

`func (o *SingleWAF) HasDirectoryTraversalSensitivity() bool`

HasDirectoryTraversalSensitivity returns a boolean if a field has been set.

### GetCrossSiteScripting

`func (o *SingleWAF) GetCrossSiteScripting() bool`

GetCrossSiteScripting returns the CrossSiteScripting field if non-nil, zero value otherwise.

### GetCrossSiteScriptingOk

`func (o *SingleWAF) GetCrossSiteScriptingOk() (*bool, bool)`

GetCrossSiteScriptingOk returns a tuple with the CrossSiteScripting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSiteScripting

`func (o *SingleWAF) SetCrossSiteScripting(v bool)`

SetCrossSiteScripting sets CrossSiteScripting field to given value.

### HasCrossSiteScripting

`func (o *SingleWAF) HasCrossSiteScripting() bool`

HasCrossSiteScripting returns a boolean if a field has been set.

### GetCrossSiteScriptingSensitivity

`func (o *SingleWAF) GetCrossSiteScriptingSensitivity() WAFSensitivityChoices`

GetCrossSiteScriptingSensitivity returns the CrossSiteScriptingSensitivity field if non-nil, zero value otherwise.

### GetCrossSiteScriptingSensitivityOk

`func (o *SingleWAF) GetCrossSiteScriptingSensitivityOk() (*WAFSensitivityChoices, bool)`

GetCrossSiteScriptingSensitivityOk returns a tuple with the CrossSiteScriptingSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSiteScriptingSensitivity

`func (o *SingleWAF) SetCrossSiteScriptingSensitivity(v WAFSensitivityChoices)`

SetCrossSiteScriptingSensitivity sets CrossSiteScriptingSensitivity field to given value.

### HasCrossSiteScriptingSensitivity

`func (o *SingleWAF) HasCrossSiteScriptingSensitivity() bool`

HasCrossSiteScriptingSensitivity returns a boolean if a field has been set.

### GetEvadingTricks

`func (o *SingleWAF) GetEvadingTricks() bool`

GetEvadingTricks returns the EvadingTricks field if non-nil, zero value otherwise.

### GetEvadingTricksOk

`func (o *SingleWAF) GetEvadingTricksOk() (*bool, bool)`

GetEvadingTricksOk returns a tuple with the EvadingTricks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvadingTricks

`func (o *SingleWAF) SetEvadingTricks(v bool)`

SetEvadingTricks sets EvadingTricks field to given value.

### HasEvadingTricks

`func (o *SingleWAF) HasEvadingTricks() bool`

HasEvadingTricks returns a boolean if a field has been set.

### GetEvadingTricksSensitivity

`func (o *SingleWAF) GetEvadingTricksSensitivity() WAFSensitivityChoices`

GetEvadingTricksSensitivity returns the EvadingTricksSensitivity field if non-nil, zero value otherwise.

### GetEvadingTricksSensitivityOk

`func (o *SingleWAF) GetEvadingTricksSensitivityOk() (*WAFSensitivityChoices, bool)`

GetEvadingTricksSensitivityOk returns a tuple with the EvadingTricksSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvadingTricksSensitivity

`func (o *SingleWAF) SetEvadingTricksSensitivity(v WAFSensitivityChoices)`

SetEvadingTricksSensitivity sets EvadingTricksSensitivity field to given value.

### HasEvadingTricksSensitivity

`func (o *SingleWAF) HasEvadingTricksSensitivity() bool`

HasEvadingTricksSensitivity returns a boolean if a field has been set.

### GetFileUpload

`func (o *SingleWAF) GetFileUpload() bool`

GetFileUpload returns the FileUpload field if non-nil, zero value otherwise.

### GetFileUploadOk

`func (o *SingleWAF) GetFileUploadOk() (*bool, bool)`

GetFileUploadOk returns a tuple with the FileUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUpload

`func (o *SingleWAF) SetFileUpload(v bool)`

SetFileUpload sets FileUpload field to given value.

### HasFileUpload

`func (o *SingleWAF) HasFileUpload() bool`

HasFileUpload returns a boolean if a field has been set.

### GetFileUploadSensitivity

`func (o *SingleWAF) GetFileUploadSensitivity() WAFSensitivityChoices`

GetFileUploadSensitivity returns the FileUploadSensitivity field if non-nil, zero value otherwise.

### GetFileUploadSensitivityOk

`func (o *SingleWAF) GetFileUploadSensitivityOk() (*WAFSensitivityChoices, bool)`

GetFileUploadSensitivityOk returns a tuple with the FileUploadSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadSensitivity

`func (o *SingleWAF) SetFileUploadSensitivity(v WAFSensitivityChoices)`

SetFileUploadSensitivity sets FileUploadSensitivity field to given value.

### HasFileUploadSensitivity

`func (o *SingleWAF) HasFileUploadSensitivity() bool`

HasFileUploadSensitivity returns a boolean if a field has been set.

### GetUnwantedAccess

`func (o *SingleWAF) GetUnwantedAccess() bool`

GetUnwantedAccess returns the UnwantedAccess field if non-nil, zero value otherwise.

### GetUnwantedAccessOk

`func (o *SingleWAF) GetUnwantedAccessOk() (*bool, bool)`

GetUnwantedAccessOk returns a tuple with the UnwantedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwantedAccess

`func (o *SingleWAF) SetUnwantedAccess(v bool)`

SetUnwantedAccess sets UnwantedAccess field to given value.

### HasUnwantedAccess

`func (o *SingleWAF) HasUnwantedAccess() bool`

HasUnwantedAccess returns a boolean if a field has been set.

### GetUnwantedAccessSensitivity

`func (o *SingleWAF) GetUnwantedAccessSensitivity() WAFSensitivityChoices`

GetUnwantedAccessSensitivity returns the UnwantedAccessSensitivity field if non-nil, zero value otherwise.

### GetUnwantedAccessSensitivityOk

`func (o *SingleWAF) GetUnwantedAccessSensitivityOk() (*WAFSensitivityChoices, bool)`

GetUnwantedAccessSensitivityOk returns a tuple with the UnwantedAccessSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwantedAccessSensitivity

`func (o *SingleWAF) SetUnwantedAccessSensitivity(v WAFSensitivityChoices)`

SetUnwantedAccessSensitivity sets UnwantedAccessSensitivity field to given value.

### HasUnwantedAccessSensitivity

`func (o *SingleWAF) HasUnwantedAccessSensitivity() bool`

HasUnwantedAccessSensitivity returns a boolean if a field has been set.

### GetIdentifiedAttack

`func (o *SingleWAF) GetIdentifiedAttack() bool`

GetIdentifiedAttack returns the IdentifiedAttack field if non-nil, zero value otherwise.

### GetIdentifiedAttackOk

`func (o *SingleWAF) GetIdentifiedAttackOk() (*bool, bool)`

GetIdentifiedAttackOk returns a tuple with the IdentifiedAttack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedAttack

`func (o *SingleWAF) SetIdentifiedAttack(v bool)`

SetIdentifiedAttack sets IdentifiedAttack field to given value.

### HasIdentifiedAttack

`func (o *SingleWAF) HasIdentifiedAttack() bool`

HasIdentifiedAttack returns a boolean if a field has been set.

### GetIdentifiedAttackSensitivity

`func (o *SingleWAF) GetIdentifiedAttackSensitivity() WAFSensitivityChoices`

GetIdentifiedAttackSensitivity returns the IdentifiedAttackSensitivity field if non-nil, zero value otherwise.

### GetIdentifiedAttackSensitivityOk

`func (o *SingleWAF) GetIdentifiedAttackSensitivityOk() (*WAFSensitivityChoices, bool)`

GetIdentifiedAttackSensitivityOk returns a tuple with the IdentifiedAttackSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedAttackSensitivity

`func (o *SingleWAF) SetIdentifiedAttackSensitivity(v WAFSensitivityChoices)`

SetIdentifiedAttackSensitivity sets IdentifiedAttackSensitivity field to given value.

### HasIdentifiedAttackSensitivity

`func (o *SingleWAF) HasIdentifiedAttackSensitivity() bool`

HasIdentifiedAttackSensitivity returns a boolean if a field has been set.

### GetBypassAddresses

`func (o *SingleWAF) GetBypassAddresses() []string`

GetBypassAddresses returns the BypassAddresses field if non-nil, zero value otherwise.

### GetBypassAddressesOk

`func (o *SingleWAF) GetBypassAddressesOk() (*[]string, bool)`

GetBypassAddressesOk returns a tuple with the BypassAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassAddresses

`func (o *SingleWAF) SetBypassAddresses(v []string)`

SetBypassAddresses sets BypassAddresses field to given value.

### HasBypassAddresses

`func (o *SingleWAF) HasBypassAddresses() bool`

HasBypassAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


