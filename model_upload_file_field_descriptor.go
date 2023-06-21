/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the UploadFileFieldDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadFileFieldDescriptor{}

// UploadFileFieldDescriptor struct for UploadFileFieldDescriptor
type UploadFileFieldDescriptor struct {
	FieldDescriptor
}

// NewUploadFileFieldDescriptor instantiates a new UploadFileFieldDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadFileFieldDescriptor() *UploadFileFieldDescriptor {
	this := UploadFileFieldDescriptor{}
	return &this
}

// NewUploadFileFieldDescriptorWithDefaults instantiates a new UploadFileFieldDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadFileFieldDescriptorWithDefaults() *UploadFileFieldDescriptor {
	this := UploadFileFieldDescriptor{}
	return &this
}

func (o UploadFileFieldDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadFileFieldDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFieldDescriptor, errFieldDescriptor := json.Marshal(o.FieldDescriptor)
	if errFieldDescriptor != nil {
		return map[string]interface{}{}, errFieldDescriptor
	}
	errFieldDescriptor = json.Unmarshal([]byte(serializedFieldDescriptor), &toSerialize)
	if errFieldDescriptor != nil {
		return map[string]interface{}{}, errFieldDescriptor
	}
	return toSerialize, nil
}

type NullableUploadFileFieldDescriptor struct {
	value *UploadFileFieldDescriptor
	isSet bool
}

func (v NullableUploadFileFieldDescriptor) Get() *UploadFileFieldDescriptor {
	return v.value
}

func (v *NullableUploadFileFieldDescriptor) Set(val *UploadFileFieldDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadFileFieldDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadFileFieldDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadFileFieldDescriptor(val *UploadFileFieldDescriptor) *NullableUploadFileFieldDescriptor {
	return &NullableUploadFileFieldDescriptor{value: val, isSet: true}
}

func (v NullableUploadFileFieldDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadFileFieldDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
