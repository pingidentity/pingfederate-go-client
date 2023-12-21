/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the RadioGroupFieldDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RadioGroupFieldDescriptor{}

// RadioGroupFieldDescriptor struct for RadioGroupFieldDescriptor
type RadioGroupFieldDescriptor struct {
	BaseSelectionFieldDescriptor
}

// NewRadioGroupFieldDescriptor instantiates a new RadioGroupFieldDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRadioGroupFieldDescriptor() *RadioGroupFieldDescriptor {
	this := RadioGroupFieldDescriptor{}
	return &this
}

// NewRadioGroupFieldDescriptorWithDefaults instantiates a new RadioGroupFieldDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRadioGroupFieldDescriptorWithDefaults() *RadioGroupFieldDescriptor {
	this := RadioGroupFieldDescriptor{}
	return &this
}

func (o RadioGroupFieldDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RadioGroupFieldDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseSelectionFieldDescriptor, errBaseSelectionFieldDescriptor := json.Marshal(o.BaseSelectionFieldDescriptor)
	if errBaseSelectionFieldDescriptor != nil {
		return map[string]interface{}{}, errBaseSelectionFieldDescriptor
	}
	errBaseSelectionFieldDescriptor = json.Unmarshal([]byte(serializedBaseSelectionFieldDescriptor), &toSerialize)
	if errBaseSelectionFieldDescriptor != nil {
		return map[string]interface{}{}, errBaseSelectionFieldDescriptor
	}
	return toSerialize, nil
}

type NullableRadioGroupFieldDescriptor struct {
	value *RadioGroupFieldDescriptor
	isSet bool
}

func (v NullableRadioGroupFieldDescriptor) Get() *RadioGroupFieldDescriptor {
	return v.value
}

func (v *NullableRadioGroupFieldDescriptor) Set(val *RadioGroupFieldDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableRadioGroupFieldDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableRadioGroupFieldDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRadioGroupFieldDescriptor(val *RadioGroupFieldDescriptor) *NullableRadioGroupFieldDescriptor {
	return &NullableRadioGroupFieldDescriptor{value: val, isSet: true}
}

func (v NullableRadioGroupFieldDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRadioGroupFieldDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
