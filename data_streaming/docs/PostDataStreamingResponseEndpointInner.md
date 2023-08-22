# PostDataStreamingResponseEndpointInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointType** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**LogLineSeparator** | Pointer to **string** |  | [optional] 
**PayloadFormat** | Pointer to **string** |  | [optional] 
**MaxSize** | Pointer to **int32** |  | [optional] 
**Headers** | Pointer to **[]map[string]string** |  | [optional] 
**KafkaTopic** | Pointer to **string** |  | [optional] 
**BootstrapServers** | Pointer to **string** |  | [optional] 
**UseTls** | Pointer to **bool** |  | [optional] 
**AccessKey** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ObjectKeyPrefix** | Pointer to **string** |  | [optional] 
**BucketName** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**HostUrl** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**DatasetId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**TableId** | Pointer to **string** |  | [optional] 
**ServiceAccountKey** | Pointer to [**EndpointGoogleBigQueryServiceAccountKey**](EndpointGoogleBigQueryServiceAccountKey.md) |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**StreamName** | Pointer to **string** |  | [optional] 
**LogType** | Pointer to **string** |  | [optional] 
**SharedKey** | Pointer to **string** |  | [optional] 
**TimeGeneratedField** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**StorageAccount** | Pointer to **string** |  | [optional] 
**ContainerName** | Pointer to **string** |  | [optional] 
**BlobSasToken** | Pointer to **string** |  | [optional] 

## Methods

### NewPostDataStreamingResponseEndpointInner

`func NewPostDataStreamingResponseEndpointInner() *PostDataStreamingResponseEndpointInner`

NewPostDataStreamingResponseEndpointInner instantiates a new PostDataStreamingResponseEndpointInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDataStreamingResponseEndpointInnerWithDefaults

`func NewPostDataStreamingResponseEndpointInnerWithDefaults() *PostDataStreamingResponseEndpointInner`

NewPostDataStreamingResponseEndpointInnerWithDefaults instantiates a new PostDataStreamingResponseEndpointInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointType

`func (o *PostDataStreamingResponseEndpointInner) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *PostDataStreamingResponseEndpointInner) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *PostDataStreamingResponseEndpointInner) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *PostDataStreamingResponseEndpointInner) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### GetUrl

`func (o *PostDataStreamingResponseEndpointInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PostDataStreamingResponseEndpointInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PostDataStreamingResponseEndpointInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PostDataStreamingResponseEndpointInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLogLineSeparator

`func (o *PostDataStreamingResponseEndpointInner) GetLogLineSeparator() string`

GetLogLineSeparator returns the LogLineSeparator field if non-nil, zero value otherwise.

### GetLogLineSeparatorOk

`func (o *PostDataStreamingResponseEndpointInner) GetLogLineSeparatorOk() (*string, bool)`

GetLogLineSeparatorOk returns a tuple with the LogLineSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLineSeparator

`func (o *PostDataStreamingResponseEndpointInner) SetLogLineSeparator(v string)`

SetLogLineSeparator sets LogLineSeparator field to given value.

### HasLogLineSeparator

`func (o *PostDataStreamingResponseEndpointInner) HasLogLineSeparator() bool`

HasLogLineSeparator returns a boolean if a field has been set.

### GetPayloadFormat

`func (o *PostDataStreamingResponseEndpointInner) GetPayloadFormat() string`

GetPayloadFormat returns the PayloadFormat field if non-nil, zero value otherwise.

### GetPayloadFormatOk

`func (o *PostDataStreamingResponseEndpointInner) GetPayloadFormatOk() (*string, bool)`

GetPayloadFormatOk returns a tuple with the PayloadFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadFormat

`func (o *PostDataStreamingResponseEndpointInner) SetPayloadFormat(v string)`

SetPayloadFormat sets PayloadFormat field to given value.

### HasPayloadFormat

`func (o *PostDataStreamingResponseEndpointInner) HasPayloadFormat() bool`

HasPayloadFormat returns a boolean if a field has been set.

### GetMaxSize

`func (o *PostDataStreamingResponseEndpointInner) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *PostDataStreamingResponseEndpointInner) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *PostDataStreamingResponseEndpointInner) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *PostDataStreamingResponseEndpointInner) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetHeaders

`func (o *PostDataStreamingResponseEndpointInner) GetHeaders() []map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PostDataStreamingResponseEndpointInner) GetHeadersOk() (*[]map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PostDataStreamingResponseEndpointInner) SetHeaders(v []map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PostDataStreamingResponseEndpointInner) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetKafkaTopic

`func (o *PostDataStreamingResponseEndpointInner) GetKafkaTopic() string`

GetKafkaTopic returns the KafkaTopic field if non-nil, zero value otherwise.

### GetKafkaTopicOk

`func (o *PostDataStreamingResponseEndpointInner) GetKafkaTopicOk() (*string, bool)`

GetKafkaTopicOk returns a tuple with the KafkaTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaTopic

`func (o *PostDataStreamingResponseEndpointInner) SetKafkaTopic(v string)`

SetKafkaTopic sets KafkaTopic field to given value.

### HasKafkaTopic

`func (o *PostDataStreamingResponseEndpointInner) HasKafkaTopic() bool`

HasKafkaTopic returns a boolean if a field has been set.

### GetBootstrapServers

`func (o *PostDataStreamingResponseEndpointInner) GetBootstrapServers() string`

GetBootstrapServers returns the BootstrapServers field if non-nil, zero value otherwise.

### GetBootstrapServersOk

`func (o *PostDataStreamingResponseEndpointInner) GetBootstrapServersOk() (*string, bool)`

GetBootstrapServersOk returns a tuple with the BootstrapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServers

`func (o *PostDataStreamingResponseEndpointInner) SetBootstrapServers(v string)`

SetBootstrapServers sets BootstrapServers field to given value.

### HasBootstrapServers

`func (o *PostDataStreamingResponseEndpointInner) HasBootstrapServers() bool`

HasBootstrapServers returns a boolean if a field has been set.

### GetUseTls

`func (o *PostDataStreamingResponseEndpointInner) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *PostDataStreamingResponseEndpointInner) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *PostDataStreamingResponseEndpointInner) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *PostDataStreamingResponseEndpointInner) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetAccessKey

`func (o *PostDataStreamingResponseEndpointInner) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *PostDataStreamingResponseEndpointInner) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *PostDataStreamingResponseEndpointInner) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *PostDataStreamingResponseEndpointInner) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetRegion

`func (o *PostDataStreamingResponseEndpointInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostDataStreamingResponseEndpointInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostDataStreamingResponseEndpointInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PostDataStreamingResponseEndpointInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetObjectKeyPrefix

`func (o *PostDataStreamingResponseEndpointInner) GetObjectKeyPrefix() string`

GetObjectKeyPrefix returns the ObjectKeyPrefix field if non-nil, zero value otherwise.

### GetObjectKeyPrefixOk

`func (o *PostDataStreamingResponseEndpointInner) GetObjectKeyPrefixOk() (*string, bool)`

GetObjectKeyPrefixOk returns a tuple with the ObjectKeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectKeyPrefix

`func (o *PostDataStreamingResponseEndpointInner) SetObjectKeyPrefix(v string)`

SetObjectKeyPrefix sets ObjectKeyPrefix field to given value.

### HasObjectKeyPrefix

`func (o *PostDataStreamingResponseEndpointInner) HasObjectKeyPrefix() bool`

HasObjectKeyPrefix returns a boolean if a field has been set.

### GetBucketName

`func (o *PostDataStreamingResponseEndpointInner) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *PostDataStreamingResponseEndpointInner) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *PostDataStreamingResponseEndpointInner) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *PostDataStreamingResponseEndpointInner) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetContentType

`func (o *PostDataStreamingResponseEndpointInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PostDataStreamingResponseEndpointInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PostDataStreamingResponseEndpointInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PostDataStreamingResponseEndpointInner) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetHostUrl

`func (o *PostDataStreamingResponseEndpointInner) GetHostUrl() string`

GetHostUrl returns the HostUrl field if non-nil, zero value otherwise.

### GetHostUrlOk

`func (o *PostDataStreamingResponseEndpointInner) GetHostUrlOk() (*string, bool)`

GetHostUrlOk returns a tuple with the HostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUrl

`func (o *PostDataStreamingResponseEndpointInner) SetHostUrl(v string)`

SetHostUrl sets HostUrl field to given value.

### HasHostUrl

`func (o *PostDataStreamingResponseEndpointInner) HasHostUrl() bool`

HasHostUrl returns a boolean if a field has been set.

### GetSecretKey

`func (o *PostDataStreamingResponseEndpointInner) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *PostDataStreamingResponseEndpointInner) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *PostDataStreamingResponseEndpointInner) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *PostDataStreamingResponseEndpointInner) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetDatasetId

`func (o *PostDataStreamingResponseEndpointInner) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *PostDataStreamingResponseEndpointInner) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *PostDataStreamingResponseEndpointInner) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *PostDataStreamingResponseEndpointInner) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### GetProjectId

`func (o *PostDataStreamingResponseEndpointInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PostDataStreamingResponseEndpointInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PostDataStreamingResponseEndpointInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *PostDataStreamingResponseEndpointInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTableId

`func (o *PostDataStreamingResponseEndpointInner) GetTableId() string`

GetTableId returns the TableId field if non-nil, zero value otherwise.

### GetTableIdOk

`func (o *PostDataStreamingResponseEndpointInner) GetTableIdOk() (*string, bool)`

GetTableIdOk returns a tuple with the TableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableId

`func (o *PostDataStreamingResponseEndpointInner) SetTableId(v string)`

SetTableId sets TableId field to given value.

### HasTableId

`func (o *PostDataStreamingResponseEndpointInner) HasTableId() bool`

HasTableId returns a boolean if a field has been set.

### GetServiceAccountKey

`func (o *PostDataStreamingResponseEndpointInner) GetServiceAccountKey() EndpointGoogleBigQueryServiceAccountKey`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *PostDataStreamingResponseEndpointInner) GetServiceAccountKeyOk() (*EndpointGoogleBigQueryServiceAccountKey, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *PostDataStreamingResponseEndpointInner) SetServiceAccountKey(v EndpointGoogleBigQueryServiceAccountKey)`

SetServiceAccountKey sets ServiceAccountKey field to given value.

### HasServiceAccountKey

`func (o *PostDataStreamingResponseEndpointInner) HasServiceAccountKey() bool`

HasServiceAccountKey returns a boolean if a field has been set.

### GetApiKey

`func (o *PostDataStreamingResponseEndpointInner) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *PostDataStreamingResponseEndpointInner) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *PostDataStreamingResponseEndpointInner) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *PostDataStreamingResponseEndpointInner) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetStreamName

`func (o *PostDataStreamingResponseEndpointInner) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *PostDataStreamingResponseEndpointInner) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *PostDataStreamingResponseEndpointInner) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.

### HasStreamName

`func (o *PostDataStreamingResponseEndpointInner) HasStreamName() bool`

HasStreamName returns a boolean if a field has been set.

### GetLogType

`func (o *PostDataStreamingResponseEndpointInner) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *PostDataStreamingResponseEndpointInner) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *PostDataStreamingResponseEndpointInner) SetLogType(v string)`

SetLogType sets LogType field to given value.

### HasLogType

`func (o *PostDataStreamingResponseEndpointInner) HasLogType() bool`

HasLogType returns a boolean if a field has been set.

### GetSharedKey

`func (o *PostDataStreamingResponseEndpointInner) GetSharedKey() string`

GetSharedKey returns the SharedKey field if non-nil, zero value otherwise.

### GetSharedKeyOk

`func (o *PostDataStreamingResponseEndpointInner) GetSharedKeyOk() (*string, bool)`

GetSharedKeyOk returns a tuple with the SharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedKey

`func (o *PostDataStreamingResponseEndpointInner) SetSharedKey(v string)`

SetSharedKey sets SharedKey field to given value.

### HasSharedKey

`func (o *PostDataStreamingResponseEndpointInner) HasSharedKey() bool`

HasSharedKey returns a boolean if a field has been set.

### GetTimeGeneratedField

`func (o *PostDataStreamingResponseEndpointInner) GetTimeGeneratedField() string`

GetTimeGeneratedField returns the TimeGeneratedField field if non-nil, zero value otherwise.

### GetTimeGeneratedFieldOk

`func (o *PostDataStreamingResponseEndpointInner) GetTimeGeneratedFieldOk() (*string, bool)`

GetTimeGeneratedFieldOk returns a tuple with the TimeGeneratedField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeGeneratedField

`func (o *PostDataStreamingResponseEndpointInner) SetTimeGeneratedField(v string)`

SetTimeGeneratedField sets TimeGeneratedField field to given value.

### HasTimeGeneratedField

`func (o *PostDataStreamingResponseEndpointInner) HasTimeGeneratedField() bool`

HasTimeGeneratedField returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *PostDataStreamingResponseEndpointInner) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *PostDataStreamingResponseEndpointInner) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *PostDataStreamingResponseEndpointInner) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *PostDataStreamingResponseEndpointInner) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetStorageAccount

`func (o *PostDataStreamingResponseEndpointInner) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *PostDataStreamingResponseEndpointInner) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *PostDataStreamingResponseEndpointInner) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *PostDataStreamingResponseEndpointInner) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### GetContainerName

`func (o *PostDataStreamingResponseEndpointInner) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *PostDataStreamingResponseEndpointInner) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *PostDataStreamingResponseEndpointInner) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *PostDataStreamingResponseEndpointInner) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetBlobSasToken

`func (o *PostDataStreamingResponseEndpointInner) GetBlobSasToken() string`

GetBlobSasToken returns the BlobSasToken field if non-nil, zero value otherwise.

### GetBlobSasTokenOk

`func (o *PostDataStreamingResponseEndpointInner) GetBlobSasTokenOk() (*string, bool)`

GetBlobSasTokenOk returns a tuple with the BlobSasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobSasToken

`func (o *PostDataStreamingResponseEndpointInner) SetBlobSasToken(v string)`

SetBlobSasToken sets BlobSasToken field to given value.

### HasBlobSasToken

`func (o *PostDataStreamingResponseEndpointInner) HasBlobSasToken() bool`

HasBlobSasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


