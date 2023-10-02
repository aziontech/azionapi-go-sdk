# CreateNewWAFRulesetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | **string** | Identification name for WAF Rule Set. | 
**Mode** | **string** |  | 
**Active** | **bool** |  | 
**SqlInjection** | **bool** |  | 
**SqlInjectionSensitivity** | [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | 
**RemoteFileInclusion** | **bool** |  | 
**RemoteFileInclusionSensitivity** | [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | 
**DirectoryTraversal** | **bool** |  | 
**DirectoryTraversalSensitivity** | [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | 
**CrossSiteScripting** | **bool** |  | 
**CrossSiteScriptingSensitivity** | [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | 
**EvadingTricks** | **bool** |  | 
**EvadingTricksSensitivity** | [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | 
**FileUpload** | **bool** |  | 
**FileUploadSensitivity** | [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | 
**UnwantedAccess** | **bool** |  | 
**UnwantedAccessSensitivity** | [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | 
**IdentifiedAttack** | **bool** |  | 
**IdentifiedAttackSensitivity** | [**WAFSensitivityChoices**](WAFSensitivityChoices.md) |  | 
**BypassAddresses** | **[]string** |  | 

## Methods

### NewCreateNewWAFRulesetRequest

`func NewCreateNewWAFRulesetRequest(name string, mode string, active bool, sqlInjection bool, sqlInjectionSensitivity WAFSensitivityChoices, remoteFileInclusion bool, remoteFileInclusionSensitivity WAFSensitivityChoices, directoryTraversal bool, directoryTraversalSensitivity WAFSensitivityChoices, crossSiteScripting bool, crossSiteScriptingSensitivity WAFSensitivityChoices, evadingTricks bool, evadingTricksSensitivity WAFSensitivityChoices, fileUpload bool, fileUploadSensitivity WAFSensitivityChoices, unwantedAccess bool, unwantedAccessSensitivity WAFSensitivityChoices, identifiedAttack bool, identifiedAttackSensitivity WAFSensitivityChoices, bypassAddresses []string, ) *CreateNewWAFRulesetRequest`

NewCreateNewWAFRulesetRequest instantiates a new CreateNewWAFRulesetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNewWAFRulesetRequestWithDefaults

`func NewCreateNewWAFRulesetRequestWithDefaults() *CreateNewWAFRulesetRequest`

NewCreateNewWAFRulesetRequestWithDefaults instantiates a new CreateNewWAFRulesetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateNewWAFRulesetRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNewWAFRulesetRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNewWAFRulesetRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNewWAFRulesetRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateNewWAFRulesetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNewWAFRulesetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNewWAFRulesetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMode

`func (o *CreateNewWAFRulesetRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateNewWAFRulesetRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateNewWAFRulesetRequest) SetMode(v string)`

SetMode sets Mode field to given value.


### GetActive

`func (o *CreateNewWAFRulesetRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateNewWAFRulesetRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateNewWAFRulesetRequest) SetActive(v bool)`

SetActive sets Active field to given value.


### GetSqlInjection

`func (o *CreateNewWAFRulesetRequest) GetSqlInjection() bool`

GetSqlInjection returns the SqlInjection field if non-nil, zero value otherwise.

### GetSqlInjectionOk

`func (o *CreateNewWAFRulesetRequest) GetSqlInjectionOk() (*bool, bool)`

GetSqlInjectionOk returns a tuple with the SqlInjection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlInjection

`func (o *CreateNewWAFRulesetRequest) SetSqlInjection(v bool)`

SetSqlInjection sets SqlInjection field to given value.


### GetSqlInjectionSensitivity

`func (o *CreateNewWAFRulesetRequest) GetSqlInjectionSensitivity() WAFSensitivityChoices`

GetSqlInjectionSensitivity returns the SqlInjectionSensitivity field if non-nil, zero value otherwise.

### GetSqlInjectionSensitivityOk

`func (o *CreateNewWAFRulesetRequest) GetSqlInjectionSensitivityOk() (*WAFSensitivityChoices, bool)`

GetSqlInjectionSensitivityOk returns a tuple with the SqlInjectionSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlInjectionSensitivity

`func (o *CreateNewWAFRulesetRequest) SetSqlInjectionSensitivity(v WAFSensitivityChoices)`

SetSqlInjectionSensitivity sets SqlInjectionSensitivity field to given value.


### GetRemoteFileInclusion

`func (o *CreateNewWAFRulesetRequest) GetRemoteFileInclusion() bool`

GetRemoteFileInclusion returns the RemoteFileInclusion field if non-nil, zero value otherwise.

### GetRemoteFileInclusionOk

`func (o *CreateNewWAFRulesetRequest) GetRemoteFileInclusionOk() (*bool, bool)`

GetRemoteFileInclusionOk returns a tuple with the RemoteFileInclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFileInclusion

`func (o *CreateNewWAFRulesetRequest) SetRemoteFileInclusion(v bool)`

SetRemoteFileInclusion sets RemoteFileInclusion field to given value.


### GetRemoteFileInclusionSensitivity

`func (o *CreateNewWAFRulesetRequest) GetRemoteFileInclusionSensitivity() WAFSensitivityChoices`

GetRemoteFileInclusionSensitivity returns the RemoteFileInclusionSensitivity field if non-nil, zero value otherwise.

### GetRemoteFileInclusionSensitivityOk

`func (o *CreateNewWAFRulesetRequest) GetRemoteFileInclusionSensitivityOk() (*WAFSensitivityChoices, bool)`

GetRemoteFileInclusionSensitivityOk returns a tuple with the RemoteFileInclusionSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteFileInclusionSensitivity

`func (o *CreateNewWAFRulesetRequest) SetRemoteFileInclusionSensitivity(v WAFSensitivityChoices)`

SetRemoteFileInclusionSensitivity sets RemoteFileInclusionSensitivity field to given value.


### GetDirectoryTraversal

`func (o *CreateNewWAFRulesetRequest) GetDirectoryTraversal() bool`

GetDirectoryTraversal returns the DirectoryTraversal field if non-nil, zero value otherwise.

### GetDirectoryTraversalOk

`func (o *CreateNewWAFRulesetRequest) GetDirectoryTraversalOk() (*bool, bool)`

GetDirectoryTraversalOk returns a tuple with the DirectoryTraversal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTraversal

`func (o *CreateNewWAFRulesetRequest) SetDirectoryTraversal(v bool)`

SetDirectoryTraversal sets DirectoryTraversal field to given value.


### GetDirectoryTraversalSensitivity

`func (o *CreateNewWAFRulesetRequest) GetDirectoryTraversalSensitivity() WAFSensitivityChoices`

GetDirectoryTraversalSensitivity returns the DirectoryTraversalSensitivity field if non-nil, zero value otherwise.

### GetDirectoryTraversalSensitivityOk

`func (o *CreateNewWAFRulesetRequest) GetDirectoryTraversalSensitivityOk() (*WAFSensitivityChoices, bool)`

GetDirectoryTraversalSensitivityOk returns a tuple with the DirectoryTraversalSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryTraversalSensitivity

`func (o *CreateNewWAFRulesetRequest) SetDirectoryTraversalSensitivity(v WAFSensitivityChoices)`

SetDirectoryTraversalSensitivity sets DirectoryTraversalSensitivity field to given value.


### GetCrossSiteScripting

`func (o *CreateNewWAFRulesetRequest) GetCrossSiteScripting() bool`

GetCrossSiteScripting returns the CrossSiteScripting field if non-nil, zero value otherwise.

### GetCrossSiteScriptingOk

`func (o *CreateNewWAFRulesetRequest) GetCrossSiteScriptingOk() (*bool, bool)`

GetCrossSiteScriptingOk returns a tuple with the CrossSiteScripting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSiteScripting

`func (o *CreateNewWAFRulesetRequest) SetCrossSiteScripting(v bool)`

SetCrossSiteScripting sets CrossSiteScripting field to given value.


### GetCrossSiteScriptingSensitivity

`func (o *CreateNewWAFRulesetRequest) GetCrossSiteScriptingSensitivity() WAFSensitivityChoices`

GetCrossSiteScriptingSensitivity returns the CrossSiteScriptingSensitivity field if non-nil, zero value otherwise.

### GetCrossSiteScriptingSensitivityOk

`func (o *CreateNewWAFRulesetRequest) GetCrossSiteScriptingSensitivityOk() (*WAFSensitivityChoices, bool)`

GetCrossSiteScriptingSensitivityOk returns a tuple with the CrossSiteScriptingSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossSiteScriptingSensitivity

`func (o *CreateNewWAFRulesetRequest) SetCrossSiteScriptingSensitivity(v WAFSensitivityChoices)`

SetCrossSiteScriptingSensitivity sets CrossSiteScriptingSensitivity field to given value.


### GetEvadingTricks

`func (o *CreateNewWAFRulesetRequest) GetEvadingTricks() bool`

GetEvadingTricks returns the EvadingTricks field if non-nil, zero value otherwise.

### GetEvadingTricksOk

`func (o *CreateNewWAFRulesetRequest) GetEvadingTricksOk() (*bool, bool)`

GetEvadingTricksOk returns a tuple with the EvadingTricks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvadingTricks

`func (o *CreateNewWAFRulesetRequest) SetEvadingTricks(v bool)`

SetEvadingTricks sets EvadingTricks field to given value.


### GetEvadingTricksSensitivity

`func (o *CreateNewWAFRulesetRequest) GetEvadingTricksSensitivity() WAFSensitivityChoices`

GetEvadingTricksSensitivity returns the EvadingTricksSensitivity field if non-nil, zero value otherwise.

### GetEvadingTricksSensitivityOk

`func (o *CreateNewWAFRulesetRequest) GetEvadingTricksSensitivityOk() (*WAFSensitivityChoices, bool)`

GetEvadingTricksSensitivityOk returns a tuple with the EvadingTricksSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvadingTricksSensitivity

`func (o *CreateNewWAFRulesetRequest) SetEvadingTricksSensitivity(v WAFSensitivityChoices)`

SetEvadingTricksSensitivity sets EvadingTricksSensitivity field to given value.


### GetFileUpload

`func (o *CreateNewWAFRulesetRequest) GetFileUpload() bool`

GetFileUpload returns the FileUpload field if non-nil, zero value otherwise.

### GetFileUploadOk

`func (o *CreateNewWAFRulesetRequest) GetFileUploadOk() (*bool, bool)`

GetFileUploadOk returns a tuple with the FileUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUpload

`func (o *CreateNewWAFRulesetRequest) SetFileUpload(v bool)`

SetFileUpload sets FileUpload field to given value.


### GetFileUploadSensitivity

`func (o *CreateNewWAFRulesetRequest) GetFileUploadSensitivity() WAFSensitivityChoices`

GetFileUploadSensitivity returns the FileUploadSensitivity field if non-nil, zero value otherwise.

### GetFileUploadSensitivityOk

`func (o *CreateNewWAFRulesetRequest) GetFileUploadSensitivityOk() (*WAFSensitivityChoices, bool)`

GetFileUploadSensitivityOk returns a tuple with the FileUploadSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadSensitivity

`func (o *CreateNewWAFRulesetRequest) SetFileUploadSensitivity(v WAFSensitivityChoices)`

SetFileUploadSensitivity sets FileUploadSensitivity field to given value.


### GetUnwantedAccess

`func (o *CreateNewWAFRulesetRequest) GetUnwantedAccess() bool`

GetUnwantedAccess returns the UnwantedAccess field if non-nil, zero value otherwise.

### GetUnwantedAccessOk

`func (o *CreateNewWAFRulesetRequest) GetUnwantedAccessOk() (*bool, bool)`

GetUnwantedAccessOk returns a tuple with the UnwantedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwantedAccess

`func (o *CreateNewWAFRulesetRequest) SetUnwantedAccess(v bool)`

SetUnwantedAccess sets UnwantedAccess field to given value.


### GetUnwantedAccessSensitivity

`func (o *CreateNewWAFRulesetRequest) GetUnwantedAccessSensitivity() WAFSensitivityChoices`

GetUnwantedAccessSensitivity returns the UnwantedAccessSensitivity field if non-nil, zero value otherwise.

### GetUnwantedAccessSensitivityOk

`func (o *CreateNewWAFRulesetRequest) GetUnwantedAccessSensitivityOk() (*WAFSensitivityChoices, bool)`

GetUnwantedAccessSensitivityOk returns a tuple with the UnwantedAccessSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwantedAccessSensitivity

`func (o *CreateNewWAFRulesetRequest) SetUnwantedAccessSensitivity(v WAFSensitivityChoices)`

SetUnwantedAccessSensitivity sets UnwantedAccessSensitivity field to given value.


### GetIdentifiedAttack

`func (o *CreateNewWAFRulesetRequest) GetIdentifiedAttack() bool`

GetIdentifiedAttack returns the IdentifiedAttack field if non-nil, zero value otherwise.

### GetIdentifiedAttackOk

`func (o *CreateNewWAFRulesetRequest) GetIdentifiedAttackOk() (*bool, bool)`

GetIdentifiedAttackOk returns a tuple with the IdentifiedAttack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedAttack

`func (o *CreateNewWAFRulesetRequest) SetIdentifiedAttack(v bool)`

SetIdentifiedAttack sets IdentifiedAttack field to given value.


### GetIdentifiedAttackSensitivity

`func (o *CreateNewWAFRulesetRequest) GetIdentifiedAttackSensitivity() WAFSensitivityChoices`

GetIdentifiedAttackSensitivity returns the IdentifiedAttackSensitivity field if non-nil, zero value otherwise.

### GetIdentifiedAttackSensitivityOk

`func (o *CreateNewWAFRulesetRequest) GetIdentifiedAttackSensitivityOk() (*WAFSensitivityChoices, bool)`

GetIdentifiedAttackSensitivityOk returns a tuple with the IdentifiedAttackSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedAttackSensitivity

`func (o *CreateNewWAFRulesetRequest) SetIdentifiedAttackSensitivity(v WAFSensitivityChoices)`

SetIdentifiedAttackSensitivity sets IdentifiedAttackSensitivity field to given value.


### GetBypassAddresses

`func (o *CreateNewWAFRulesetRequest) GetBypassAddresses() []string`

GetBypassAddresses returns the BypassAddresses field if non-nil, zero value otherwise.

### GetBypassAddressesOk

`func (o *CreateNewWAFRulesetRequest) GetBypassAddressesOk() (*[]string, bool)`

GetBypassAddressesOk returns a tuple with the BypassAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassAddresses

`func (o *CreateNewWAFRulesetRequest) SetBypassAddresses(v []string)`

SetBypassAddresses sets BypassAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


