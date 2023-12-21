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

// checks if the DateLocalIdentityField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateLocalIdentityField{}

// DateLocalIdentityField struct for DateLocalIdentityField
type DateLocalIdentityField struct {
	BaseDefaultValueLocalIdentityField
}

// NewDateLocalIdentityField instantiates a new DateLocalIdentityField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateLocalIdentityField(type_ string, id string, label string) *DateLocalIdentityField {
	this := DateLocalIdentityField{}
	this.Type = type_
	this.Id = id
	this.Label = label
	return &this
}

// NewDateLocalIdentityFieldWithDefaults instantiates a new DateLocalIdentityField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateLocalIdentityFieldWithDefaults() *DateLocalIdentityField {
	this := DateLocalIdentityField{}
	return &this
}

func (o DateLocalIdentityField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DateLocalIdentityField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseDefaultValueLocalIdentityField, errBaseDefaultValueLocalIdentityField := json.Marshal(o.BaseDefaultValueLocalIdentityField)
	if errBaseDefaultValueLocalIdentityField != nil {
		return map[string]interface{}{}, errBaseDefaultValueLocalIdentityField
	}
	errBaseDefaultValueLocalIdentityField = json.Unmarshal([]byte(serializedBaseDefaultValueLocalIdentityField), &toSerialize)
	if errBaseDefaultValueLocalIdentityField != nil {
		return map[string]interface{}{}, errBaseDefaultValueLocalIdentityField
	}
	return toSerialize, nil
}

type NullableDateLocalIdentityField struct {
	value *DateLocalIdentityField
	isSet bool
}

func (v NullableDateLocalIdentityField) Get() *DateLocalIdentityField {
	return v.value
}

func (v *NullableDateLocalIdentityField) Set(val *DateLocalIdentityField) {
	v.value = val
	v.isSet = true
}

func (v NullableDateLocalIdentityField) IsSet() bool {
	return v.isSet
}

func (v *NullableDateLocalIdentityField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateLocalIdentityField(val *DateLocalIdentityField) *NullableDateLocalIdentityField {
	return &NullableDateLocalIdentityField{value: val, isSet: true}
}

func (v NullableDateLocalIdentityField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateLocalIdentityField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
