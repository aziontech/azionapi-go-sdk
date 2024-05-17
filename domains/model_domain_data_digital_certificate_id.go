/*
Domain API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package domains

import (
	"encoding/json"
	"fmt"
)

// DomainDataDigitalCertificateId struct for DomainDataDigitalCertificateId
type DomainDataDigitalCertificateId struct {
	int64 *int64
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DomainDataDigitalCertificateId) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into int64
	err = json.Unmarshal(data, &dst.int64);
	if err == nil {
		jsonint64, _ := json.Marshal(dst.int64)
		if string(jsonint64) == "{}" { // empty struct
			dst.int64 = nil
		} else {
			return nil // data stored in dst.int64, return on the first match
		}
	} else {
		dst.int64 = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(DomainDataDigitalCertificateId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DomainDataDigitalCertificateId) MarshalJSON() ([]byte, error) {
	if src.int64 != nil {
		return json.Marshal(&src.int64)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDomainDataDigitalCertificateId struct {
	value *DomainDataDigitalCertificateId
	isSet bool
}

func (v NullableDomainDataDigitalCertificateId) Get() *DomainDataDigitalCertificateId {
	return v.value
}

func (v *NullableDomainDataDigitalCertificateId) Set(val *DomainDataDigitalCertificateId) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainDataDigitalCertificateId) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainDataDigitalCertificateId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainDataDigitalCertificateId(val *DomainDataDigitalCertificateId) *NullableDomainDataDigitalCertificateId {
	return &NullableDomainDataDigitalCertificateId{value: val, isSet: true}
}

func (v NullableDomainDataDigitalCertificateId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainDataDigitalCertificateId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

