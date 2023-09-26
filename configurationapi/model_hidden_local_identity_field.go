/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the HiddenLocalIdentityField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HiddenLocalIdentityField{}

// HiddenLocalIdentityField struct for HiddenLocalIdentityField
type HiddenLocalIdentityField struct {
	LocalIdentityField
}

// NewHiddenLocalIdentityField instantiates a new HiddenLocalIdentityField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHiddenLocalIdentityField(type_ string, id string, label string) *HiddenLocalIdentityField {
	this := HiddenLocalIdentityField{}
	this.Type = type_
	this.Id = id
	this.Label = label
	return &this
}

// NewHiddenLocalIdentityFieldWithDefaults instantiates a new HiddenLocalIdentityField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHiddenLocalIdentityFieldWithDefaults() *HiddenLocalIdentityField {
	this := HiddenLocalIdentityField{}
	return &this
}

func (o HiddenLocalIdentityField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HiddenLocalIdentityField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedLocalIdentityField, errLocalIdentityField := json.Marshal(o.LocalIdentityField)
	if errLocalIdentityField != nil {
		return map[string]interface{}{}, errLocalIdentityField
	}
	errLocalIdentityField = json.Unmarshal([]byte(serializedLocalIdentityField), &toSerialize)
	if errLocalIdentityField != nil {
		return map[string]interface{}{}, errLocalIdentityField
	}
	return toSerialize, nil
}

type NullableHiddenLocalIdentityField struct {
	value *HiddenLocalIdentityField
	isSet bool
}

func (v NullableHiddenLocalIdentityField) Get() *HiddenLocalIdentityField {
	return v.value
}

func (v *NullableHiddenLocalIdentityField) Set(val *HiddenLocalIdentityField) {
	v.value = val
	v.isSet = true
}

func (v NullableHiddenLocalIdentityField) IsSet() bool {
	return v.isSet
}

func (v *NullableHiddenLocalIdentityField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHiddenLocalIdentityField(val *HiddenLocalIdentityField) *NullableHiddenLocalIdentityField {
	return &NullableHiddenLocalIdentityField{value: val, isSet: true}
}

func (v NullableHiddenLocalIdentityField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHiddenLocalIdentityField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}