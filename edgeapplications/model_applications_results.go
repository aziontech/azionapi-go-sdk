/*
Edge Application API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgeapplications

import (
	"encoding/json"
)

// checks if the ApplicationsResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationsResults{}

// ApplicationsResults struct for ApplicationsResults
type ApplicationsResults struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	DebugRules *bool `json:"debug_rules,omitempty"`
	LastEditor *string `json:"last_editor,omitempty"`
	LastModified *string `json:"last_modified,omitempty"`
	Active *bool `json:"active,omitempty"`
	Origins []ApplicationOrigins `json:"origins,omitempty"`
}

// NewApplicationsResults instantiates a new ApplicationsResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationsResults() *ApplicationsResults {
	this := ApplicationsResults{}
	return &this
}

// NewApplicationsResultsWithDefaults instantiates a new ApplicationsResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationsResultsWithDefaults() *ApplicationsResults {
	this := ApplicationsResults{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationsResults) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResults) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationsResults) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ApplicationsResults) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationsResults) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResults) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationsResults) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationsResults) SetName(v string) {
	o.Name = &v
}

// GetDebugRules returns the DebugRules field value if set, zero value otherwise.
func (o *ApplicationsResults) GetDebugRules() bool {
	if o == nil || IsNil(o.DebugRules) {
		var ret bool
		return ret
	}
	return *o.DebugRules
}

// GetDebugRulesOk returns a tuple with the DebugRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResults) GetDebugRulesOk() (*bool, bool) {
	if o == nil || IsNil(o.DebugRules) {
		return nil, false
	}
	return o.DebugRules, true
}

// HasDebugRules returns a boolean if a field has been set.
func (o *ApplicationsResults) HasDebugRules() bool {
	if o != nil && !IsNil(o.DebugRules) {
		return true
	}

	return false
}

// SetDebugRules gets a reference to the given bool and assigns it to the DebugRules field.
func (o *ApplicationsResults) SetDebugRules(v bool) {
	o.DebugRules = &v
}

// GetLastEditor returns the LastEditor field value if set, zero value otherwise.
func (o *ApplicationsResults) GetLastEditor() string {
	if o == nil || IsNil(o.LastEditor) {
		var ret string
		return ret
	}
	return *o.LastEditor
}

// GetLastEditorOk returns a tuple with the LastEditor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResults) GetLastEditorOk() (*string, bool) {
	if o == nil || IsNil(o.LastEditor) {
		return nil, false
	}
	return o.LastEditor, true
}

// HasLastEditor returns a boolean if a field has been set.
func (o *ApplicationsResults) HasLastEditor() bool {
	if o != nil && !IsNil(o.LastEditor) {
		return true
	}

	return false
}

// SetLastEditor gets a reference to the given string and assigns it to the LastEditor field.
func (o *ApplicationsResults) SetLastEditor(v string) {
	o.LastEditor = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *ApplicationsResults) GetLastModified() string {
	if o == nil || IsNil(o.LastModified) {
		var ret string
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResults) GetLastModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *ApplicationsResults) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given string and assigns it to the LastModified field.
func (o *ApplicationsResults) SetLastModified(v string) {
	o.LastModified = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApplicationsResults) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResults) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApplicationsResults) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApplicationsResults) SetActive(v bool) {
	o.Active = &v
}

// GetOrigins returns the Origins field value if set, zero value otherwise.
func (o *ApplicationsResults) GetOrigins() []ApplicationOrigins {
	if o == nil || IsNil(o.Origins) {
		var ret []ApplicationOrigins
		return ret
	}
	return o.Origins
}

// GetOriginsOk returns a tuple with the Origins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationsResults) GetOriginsOk() ([]ApplicationOrigins, bool) {
	if o == nil || IsNil(o.Origins) {
		return nil, false
	}
	return o.Origins, true
}

// HasOrigins returns a boolean if a field has been set.
func (o *ApplicationsResults) HasOrigins() bool {
	if o != nil && !IsNil(o.Origins) {
		return true
	}

	return false
}

// SetOrigins gets a reference to the given []ApplicationOrigins and assigns it to the Origins field.
func (o *ApplicationsResults) SetOrigins(v []ApplicationOrigins) {
	o.Origins = v
}

func (o ApplicationsResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationsResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DebugRules) {
		toSerialize["debug_rules"] = o.DebugRules
	}
	if !IsNil(o.LastEditor) {
		toSerialize["last_editor"] = o.LastEditor
	}
	if !IsNil(o.LastModified) {
		toSerialize["last_modified"] = o.LastModified
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Origins) {
		toSerialize["origins"] = o.Origins
	}
	return toSerialize, nil
}

type NullableApplicationsResults struct {
	value *ApplicationsResults
	isSet bool
}

func (v NullableApplicationsResults) Get() *ApplicationsResults {
	return v.value
}

func (v *NullableApplicationsResults) Set(val *ApplicationsResults) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationsResults) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationsResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationsResults(val *ApplicationsResults) *NullableApplicationsResults {
	return &NullableApplicationsResults{value: val, isSet: true}
}

func (v NullableApplicationsResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationsResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


