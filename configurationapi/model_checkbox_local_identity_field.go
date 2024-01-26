/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"fmt"
)

// checks if the CheckboxLocalIdentityField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckboxLocalIdentityField{}

// CheckboxLocalIdentityField struct for CheckboxLocalIdentityField
type CheckboxLocalIdentityField struct {
	BaseDefaultValueLocalIdentityField
}

type _CheckboxLocalIdentityField CheckboxLocalIdentityField

// NewCheckboxLocalIdentityField instantiates a new CheckboxLocalIdentityField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckboxLocalIdentityField(type_ string, id string, label string) *CheckboxLocalIdentityField {
	this := CheckboxLocalIdentityField{}
	this.Type = type_
	this.Id = id
	this.Label = label
	return &this
}

// NewCheckboxLocalIdentityFieldWithDefaults instantiates a new CheckboxLocalIdentityField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckboxLocalIdentityFieldWithDefaults() *CheckboxLocalIdentityField {
	this := CheckboxLocalIdentityField{}
	return &this
}

func (o CheckboxLocalIdentityField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckboxLocalIdentityField) ToMap() (map[string]interface{}, error) {
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

func (o *CheckboxLocalIdentityField) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"label",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCheckboxLocalIdentityField := _CheckboxLocalIdentityField{}

	err = json.Unmarshal(bytes, &varCheckboxLocalIdentityField)

	if err != nil {
		return err
	}

	*o = CheckboxLocalIdentityField(varCheckboxLocalIdentityField)

	return err
}

type NullableCheckboxLocalIdentityField struct {
	value *CheckboxLocalIdentityField
	isSet bool
}

func (v NullableCheckboxLocalIdentityField) Get() *CheckboxLocalIdentityField {
	return v.value
}

func (v *NullableCheckboxLocalIdentityField) Set(val *CheckboxLocalIdentityField) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckboxLocalIdentityField) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckboxLocalIdentityField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckboxLocalIdentityField(val *CheckboxLocalIdentityField) *NullableCheckboxLocalIdentityField {
	return &NullableCheckboxLocalIdentityField{value: val, isSet: true}
}

func (v NullableCheckboxLocalIdentityField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckboxLocalIdentityField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
