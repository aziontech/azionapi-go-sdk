/*
Digital Certificates API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package digital_certificates

import (
	"encoding/json"
)

// checks if the DC200List type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DC200List{}

// DC200List struct for DC200List
type DC200List struct {
	Count *int32 `json:"count,omitempty"`
	TotalPages *int32 `json:"total_pages,omitempty"`
	SchemaVersion *int32 `json:"schema_version,omitempty"`
	Links *DC200ListLinks `json:"links,omitempty"`
	Results [][]ResultsInner `json:"results,omitempty"`
}

// NewDC200List instantiates a new DC200List object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDC200List() *DC200List {
	this := DC200List{}
	return &this
}

// NewDC200ListWithDefaults instantiates a new DC200List object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDC200ListWithDefaults() *DC200List {
	this := DC200List{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *DC200List) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DC200List) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *DC200List) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *DC200List) SetCount(v int32) {
	o.Count = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *DC200List) GetTotalPages() int32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DC200List) GetTotalPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *DC200List) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *DC200List) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *DC200List) GetSchemaVersion() int32 {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret int32
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DC200List) GetSchemaVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *DC200List) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given int32 and assigns it to the SchemaVersion field.
func (o *DC200List) SetSchemaVersion(v int32) {
	o.SchemaVersion = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DC200List) GetLinks() DC200ListLinks {
	if o == nil || IsNil(o.Links) {
		var ret DC200ListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DC200List) GetLinksOk() (*DC200ListLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DC200List) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given DC200ListLinks and assigns it to the Links field.
func (o *DC200List) SetLinks(v DC200ListLinks) {
	o.Links = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DC200List) GetResults() [][]ResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret [][]ResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DC200List) GetResultsOk() ([][]ResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DC200List) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given [][]ResultsInner and assigns it to the Results field.
func (o *DC200List) SetResults(v [][]ResultsInner) {
	o.Results = v
}

func (o DC200List) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DC200List) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.TotalPages) {
		toSerialize["total_pages"] = o.TotalPages
	}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableDC200List struct {
	value *DC200List
	isSet bool
}

func (v NullableDC200List) Get() *DC200List {
	return v.value
}

func (v *NullableDC200List) Set(val *DC200List) {
	v.value = val
	v.isSet = true
}

func (v NullableDC200List) IsSet() bool {
	return v.isSet
}

func (v *NullableDC200List) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDC200List(val *DC200List) *NullableDC200List {
	return &NullableDC200List{value: val, isSet: true}
}

func (v NullableDC200List) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDC200List) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

