/*
Data Streaming - OpenAPI

The Data Streaming API allows you to manage your existing data streamings and templates. Data Streaming allows you to feed your stream processing, SIEM, and big data platforms with the event logs from your applications on Azion in real time. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data_streaming

import (
	"encoding/json"
	"fmt"
)

// CreateNewDataStreaming201Response - struct for CreateNewDataStreaming201Response
type CreateNewDataStreaming201Response struct {
	CreateCustomDataStreamingResponse *CreateCustomDataStreamingResponse
	CreateDataStreamingResponse *CreateDataStreamingResponse
}

// CreateCustomDataStreamingResponseAsCreateNewDataStreaming201Response is a convenience function that returns CreateCustomDataStreamingResponse wrapped in CreateNewDataStreaming201Response
func CreateCustomDataStreamingResponseAsCreateNewDataStreaming201Response(v *CreateCustomDataStreamingResponse) CreateNewDataStreaming201Response {
	return CreateNewDataStreaming201Response{
		CreateCustomDataStreamingResponse: v,
	}
}

// CreateDataStreamingResponseAsCreateNewDataStreaming201Response is a convenience function that returns CreateDataStreamingResponse wrapped in CreateNewDataStreaming201Response
func CreateDataStreamingResponseAsCreateNewDataStreaming201Response(v *CreateDataStreamingResponse) CreateNewDataStreaming201Response {
	return CreateNewDataStreaming201Response{
		CreateDataStreamingResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateNewDataStreaming201Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateCustomDataStreamingResponse
	err = newStrictDecoder(data).Decode(&dst.CreateCustomDataStreamingResponse)
	if err == nil {
		jsonCreateCustomDataStreamingResponse, _ := json.Marshal(dst.CreateCustomDataStreamingResponse)
		if string(jsonCreateCustomDataStreamingResponse) == "{}" { // empty struct
			dst.CreateCustomDataStreamingResponse = nil
		} else {
			match++
		}
	} else {
		dst.CreateCustomDataStreamingResponse = nil
	}

	// try to unmarshal data into CreateDataStreamingResponse
	err = newStrictDecoder(data).Decode(&dst.CreateDataStreamingResponse)
	if err == nil {
		jsonCreateDataStreamingResponse, _ := json.Marshal(dst.CreateDataStreamingResponse)
		if string(jsonCreateDataStreamingResponse) == "{}" { // empty struct
			dst.CreateDataStreamingResponse = nil
		} else {
			match++
		}
	} else {
		dst.CreateDataStreamingResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateCustomDataStreamingResponse = nil
		dst.CreateDataStreamingResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateNewDataStreaming201Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateNewDataStreaming201Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateNewDataStreaming201Response) MarshalJSON() ([]byte, error) {
	if src.CreateCustomDataStreamingResponse != nil {
		return json.Marshal(&src.CreateCustomDataStreamingResponse)
	}

	if src.CreateDataStreamingResponse != nil {
		return json.Marshal(&src.CreateDataStreamingResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateNewDataStreaming201Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CreateCustomDataStreamingResponse != nil {
		return obj.CreateCustomDataStreamingResponse
	}

	if obj.CreateDataStreamingResponse != nil {
		return obj.CreateDataStreamingResponse
	}

	// all schemas are nil
	return nil
}

type NullableCreateNewDataStreaming201Response struct {
	value *CreateNewDataStreaming201Response
	isSet bool
}

func (v NullableCreateNewDataStreaming201Response) Get() *CreateNewDataStreaming201Response {
	return v.value
}

func (v *NullableCreateNewDataStreaming201Response) Set(val *CreateNewDataStreaming201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNewDataStreaming201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNewDataStreaming201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNewDataStreaming201Response(val *CreateNewDataStreaming201Response) *NullableCreateNewDataStreaming201Response {
	return &NullableCreateNewDataStreaming201Response{value: val, isSet: true}
}

func (v NullableCreateNewDataStreaming201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNewDataStreaming201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


