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

// checks if the CheckboxGroupLocalIdentityField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckboxGroupLocalIdentityField{}

// CheckboxGroupLocalIdentityField struct for CheckboxGroupLocalIdentityField
type CheckboxGroupLocalIdentityField struct {
	BaseSelectionLocalIdentityField
	// The list of options for this selection field.
	Options []string `json:"options" tfsdk:"options"`
}

// NewCheckboxGroupLocalIdentityField instantiates a new CheckboxGroupLocalIdentityField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckboxGroupLocalIdentityField(options []string, type_ string, id string, label string) *CheckboxGroupLocalIdentityField {
	this := CheckboxGroupLocalIdentityField{}
	this.Type = type_
	this.Id = id
	this.Label = label
	this.Options = options
	return &this
}

// NewCheckboxGroupLocalIdentityFieldWithDefaults instantiates a new CheckboxGroupLocalIdentityField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckboxGroupLocalIdentityFieldWithDefaults() *CheckboxGroupLocalIdentityField {
	this := CheckboxGroupLocalIdentityField{}
	return &this
}

// GetOptions returns the Options field value
func (o *CheckboxGroupLocalIdentityField) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *CheckboxGroupLocalIdentityField) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *CheckboxGroupLocalIdentityField) SetOptions(v []string) {
	o.Options = v
}

func (o CheckboxGroupLocalIdentityField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckboxGroupLocalIdentityField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseSelectionLocalIdentityField, errBaseSelectionLocalIdentityField := json.Marshal(o.BaseSelectionLocalIdentityField)
	if errBaseSelectionLocalIdentityField != nil {
		return map[string]interface{}{}, errBaseSelectionLocalIdentityField
	}
	errBaseSelectionLocalIdentityField = json.Unmarshal([]byte(serializedBaseSelectionLocalIdentityField), &toSerialize)
	if errBaseSelectionLocalIdentityField != nil {
		return map[string]interface{}{}, errBaseSelectionLocalIdentityField
	}
	toSerialize["options"] = o.Options
	return toSerialize, nil
}

type NullableCheckboxGroupLocalIdentityField struct {
	value *CheckboxGroupLocalIdentityField
	isSet bool
}

func (v NullableCheckboxGroupLocalIdentityField) Get() *CheckboxGroupLocalIdentityField {
	return v.value
}

func (v *NullableCheckboxGroupLocalIdentityField) Set(val *CheckboxGroupLocalIdentityField) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckboxGroupLocalIdentityField) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckboxGroupLocalIdentityField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckboxGroupLocalIdentityField(val *CheckboxGroupLocalIdentityField) *NullableCheckboxGroupLocalIdentityField {
	return &NullableCheckboxGroupLocalIdentityField{value: val, isSet: true}
}

func (v NullableCheckboxGroupLocalIdentityField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckboxGroupLocalIdentityField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
