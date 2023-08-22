/*
Data Streaming - OpenAPI

The Data Streaming API allows you to manage your existing data streamings and templates. Data Streaming allows you to feed your stream processing, SIEM, and big data platforms with the event logs from your applications on Azion in real time. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data_streaming

import (
	"encoding/json"
)

// checks if the EndpointGoogleBigQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointGoogleBigQuery{}

// EndpointGoogleBigQuery struct for EndpointGoogleBigQuery
type EndpointGoogleBigQuery struct {
	EndpointType *string `json:"endpoint_type,omitempty"`
	DatasetId *string `json:"dataset_id,omitempty"`
	ProjectId *string `json:"project_id,omitempty"`
	TableId *string `json:"table_id,omitempty"`
	ServiceAccountKey *EndpointGoogleBigQueryServiceAccountKey `json:"service_account_key,omitempty"`
}

// NewEndpointGoogleBigQuery instantiates a new EndpointGoogleBigQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointGoogleBigQuery() *EndpointGoogleBigQuery {
	this := EndpointGoogleBigQuery{}
	return &this
}

// NewEndpointGoogleBigQueryWithDefaults instantiates a new EndpointGoogleBigQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointGoogleBigQueryWithDefaults() *EndpointGoogleBigQuery {
	this := EndpointGoogleBigQuery{}
	return &this
}

// GetEndpointType returns the EndpointType field value if set, zero value otherwise.
func (o *EndpointGoogleBigQuery) GetEndpointType() string {
	if o == nil || IsNil(o.EndpointType) {
		var ret string
		return ret
	}
	return *o.EndpointType
}

// GetEndpointTypeOk returns a tuple with the EndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointGoogleBigQuery) GetEndpointTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointType) {
		return nil, false
	}
	return o.EndpointType, true
}

// HasEndpointType returns a boolean if a field has been set.
func (o *EndpointGoogleBigQuery) HasEndpointType() bool {
	if o != nil && !IsNil(o.EndpointType) {
		return true
	}

	return false
}

// SetEndpointType gets a reference to the given string and assigns it to the EndpointType field.
func (o *EndpointGoogleBigQuery) SetEndpointType(v string) {
	o.EndpointType = &v
}

// GetDatasetId returns the DatasetId field value if set, zero value otherwise.
func (o *EndpointGoogleBigQuery) GetDatasetId() string {
	if o == nil || IsNil(o.DatasetId) {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointGoogleBigQuery) GetDatasetIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetId) {
		return nil, false
	}
	return o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *EndpointGoogleBigQuery) HasDatasetId() bool {
	if o != nil && !IsNil(o.DatasetId) {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *EndpointGoogleBigQuery) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *EndpointGoogleBigQuery) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointGoogleBigQuery) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *EndpointGoogleBigQuery) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *EndpointGoogleBigQuery) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetTableId returns the TableId field value if set, zero value otherwise.
func (o *EndpointGoogleBigQuery) GetTableId() string {
	if o == nil || IsNil(o.TableId) {
		var ret string
		return ret
	}
	return *o.TableId
}

// GetTableIdOk returns a tuple with the TableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointGoogleBigQuery) GetTableIdOk() (*string, bool) {
	if o == nil || IsNil(o.TableId) {
		return nil, false
	}
	return o.TableId, true
}

// HasTableId returns a boolean if a field has been set.
func (o *EndpointGoogleBigQuery) HasTableId() bool {
	if o != nil && !IsNil(o.TableId) {
		return true
	}

	return false
}

// SetTableId gets a reference to the given string and assigns it to the TableId field.
func (o *EndpointGoogleBigQuery) SetTableId(v string) {
	o.TableId = &v
}

// GetServiceAccountKey returns the ServiceAccountKey field value if set, zero value otherwise.
func (o *EndpointGoogleBigQuery) GetServiceAccountKey() EndpointGoogleBigQueryServiceAccountKey {
	if o == nil || IsNil(o.ServiceAccountKey) {
		var ret EndpointGoogleBigQueryServiceAccountKey
		return ret
	}
	return *o.ServiceAccountKey
}

// GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointGoogleBigQuery) GetServiceAccountKeyOk() (*EndpointGoogleBigQueryServiceAccountKey, bool) {
	if o == nil || IsNil(o.ServiceAccountKey) {
		return nil, false
	}
	return o.ServiceAccountKey, true
}

// HasServiceAccountKey returns a boolean if a field has been set.
func (o *EndpointGoogleBigQuery) HasServiceAccountKey() bool {
	if o != nil && !IsNil(o.ServiceAccountKey) {
		return true
	}

	return false
}

// SetServiceAccountKey gets a reference to the given EndpointGoogleBigQueryServiceAccountKey and assigns it to the ServiceAccountKey field.
func (o *EndpointGoogleBigQuery) SetServiceAccountKey(v EndpointGoogleBigQueryServiceAccountKey) {
	o.ServiceAccountKey = &v
}

func (o EndpointGoogleBigQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointGoogleBigQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointType) {
		toSerialize["endpoint_type"] = o.EndpointType
	}
	if !IsNil(o.DatasetId) {
		toSerialize["dataset_id"] = o.DatasetId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.TableId) {
		toSerialize["table_id"] = o.TableId
	}
	if !IsNil(o.ServiceAccountKey) {
		toSerialize["service_account_key"] = o.ServiceAccountKey
	}
	return toSerialize, nil
}

type NullableEndpointGoogleBigQuery struct {
	value *EndpointGoogleBigQuery
	isSet bool
}

func (v NullableEndpointGoogleBigQuery) Get() *EndpointGoogleBigQuery {
	return v.value
}

func (v *NullableEndpointGoogleBigQuery) Set(val *EndpointGoogleBigQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointGoogleBigQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointGoogleBigQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointGoogleBigQuery(val *EndpointGoogleBigQuery) *NullableEndpointGoogleBigQuery {
	return &NullableEndpointGoogleBigQuery{value: val, isSet: true}
}

func (v NullableEndpointGoogleBigQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointGoogleBigQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


