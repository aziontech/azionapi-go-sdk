# EndpointKafka

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointType** | Pointer to **string** |  | [optional] 
**KafkaTopic** | Pointer to **string** |  | [optional] 
**BootstrapServers** | Pointer to **string** |  | [optional] 
**UseTls** | Pointer to **bool** |  | [optional] 

## Methods

### NewEndpointKafka

`func NewEndpointKafka() *EndpointKafka`

NewEndpointKafka instantiates a new EndpointKafka object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointKafkaWithDefaults

`func NewEndpointKafkaWithDefaults() *EndpointKafka`

NewEndpointKafkaWithDefaults instantiates a new EndpointKafka object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointType

`func (o *EndpointKafka) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *EndpointKafka) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *EndpointKafka) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *EndpointKafka) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### GetKafkaTopic

`func (o *EndpointKafka) GetKafkaTopic() string`

GetKafkaTopic returns the KafkaTopic field if non-nil, zero value otherwise.

### GetKafkaTopicOk

`func (o *EndpointKafka) GetKafkaTopicOk() (*string, bool)`

GetKafkaTopicOk returns a tuple with the KafkaTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaTopic

`func (o *EndpointKafka) SetKafkaTopic(v string)`

SetKafkaTopic sets KafkaTopic field to given value.

### HasKafkaTopic

`func (o *EndpointKafka) HasKafkaTopic() bool`

HasKafkaTopic returns a boolean if a field has been set.

### GetBootstrapServers

`func (o *EndpointKafka) GetBootstrapServers() string`

GetBootstrapServers returns the BootstrapServers field if non-nil, zero value otherwise.

### GetBootstrapServersOk

`func (o *EndpointKafka) GetBootstrapServersOk() (*string, bool)`

GetBootstrapServersOk returns a tuple with the BootstrapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServers

`func (o *EndpointKafka) SetBootstrapServers(v string)`

SetBootstrapServers sets BootstrapServers field to given value.

### HasBootstrapServers

`func (o *EndpointKafka) HasBootstrapServers() bool`

HasBootstrapServers returns a boolean if a field has been set.

### GetUseTls

`func (o *EndpointKafka) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *EndpointKafka) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *EndpointKafka) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *EndpointKafka) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


