/*
Edge Functions Instances API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package edgefunctionsinstance_edgefirewall

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// CreateEdgeFunctionsInstancesRequestJsonArgs - struct for CreateEdgeFunctionsInstancesRequestJsonArgs
type CreateEdgeFunctionsInstancesRequestJsonArgs struct {
	ArrayOfMapmapOfStringAny *[]map[string]interface{}
	MapmapOfStringAny *map[string]interface{}
}

// []map[string]interface{}AsCreateEdgeFunctionsInstancesRequestJsonArgs is a convenience function that returns []map[string]interface{} wrapped in CreateEdgeFunctionsInstancesRequestJsonArgs
func ArrayOfMapmapOfStringAnyAsCreateEdgeFunctionsInstancesRequestJsonArgs(v *[]map[string]interface{}) CreateEdgeFunctionsInstancesRequestJsonArgs {
	return CreateEdgeFunctionsInstancesRequestJsonArgs{
		ArrayOfMapmapOfStringAny: v,
	}
}

// map[string]interface{}AsCreateEdgeFunctionsInstancesRequestJsonArgs is a convenience function that returns map[string]interface{} wrapped in CreateEdgeFunctionsInstancesRequestJsonArgs
func MapmapOfStringAnyAsCreateEdgeFunctionsInstancesRequestJsonArgs(v *map[string]interface{}) CreateEdgeFunctionsInstancesRequestJsonArgs {
	return CreateEdgeFunctionsInstancesRequestJsonArgs{
		MapmapOfStringAny: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateEdgeFunctionsInstancesRequestJsonArgs) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfMapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.ArrayOfMapmapOfStringAny)
	if err == nil {
		jsonArrayOfMapmapOfStringAny, _ := json.Marshal(dst.ArrayOfMapmapOfStringAny)
		if string(jsonArrayOfMapmapOfStringAny) == "{}" { // empty struct
			dst.ArrayOfMapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.ArrayOfMapmapOfStringAny); err != nil {
				dst.ArrayOfMapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfMapmapOfStringAny = nil
	}

	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringAny); err != nil {
				dst.MapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfMapmapOfStringAny = nil
		dst.MapmapOfStringAny = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateEdgeFunctionsInstancesRequestJsonArgs)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateEdgeFunctionsInstancesRequestJsonArgs)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateEdgeFunctionsInstancesRequestJsonArgs) MarshalJSON() ([]byte, error) {
	if src.ArrayOfMapmapOfStringAny != nil {
		return json.Marshal(&src.ArrayOfMapmapOfStringAny)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateEdgeFunctionsInstancesRequestJsonArgs) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfMapmapOfStringAny != nil {
		return obj.ArrayOfMapmapOfStringAny
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj CreateEdgeFunctionsInstancesRequestJsonArgs) GetActualInstanceValue() (interface{}) {
	if obj.ArrayOfMapmapOfStringAny != nil {
		return *obj.ArrayOfMapmapOfStringAny
	}

	if obj.MapmapOfStringAny != nil {
		return *obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

type NullableCreateEdgeFunctionsInstancesRequestJsonArgs struct {
	value *CreateEdgeFunctionsInstancesRequestJsonArgs
	isSet bool
}

func (v NullableCreateEdgeFunctionsInstancesRequestJsonArgs) Get() *CreateEdgeFunctionsInstancesRequestJsonArgs {
	return v.value
}

func (v *NullableCreateEdgeFunctionsInstancesRequestJsonArgs) Set(val *CreateEdgeFunctionsInstancesRequestJsonArgs) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEdgeFunctionsInstancesRequestJsonArgs) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEdgeFunctionsInstancesRequestJsonArgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEdgeFunctionsInstancesRequestJsonArgs(val *CreateEdgeFunctionsInstancesRequestJsonArgs) *NullableCreateEdgeFunctionsInstancesRequestJsonArgs {
	return &NullableCreateEdgeFunctionsInstancesRequestJsonArgs{value: val, isSet: true}
}

func (v NullableCreateEdgeFunctionsInstancesRequestJsonArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEdgeFunctionsInstancesRequestJsonArgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


