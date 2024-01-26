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

// checks if the DropDownLocalIdentityField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DropDownLocalIdentityField{}

// DropDownLocalIdentityField struct for DropDownLocalIdentityField
type DropDownLocalIdentityField struct {
	BaseSelectionLocalIdentityField
	// The default value for this field.
	DefaultValue *string `json:"defaultValue,omitempty" tfsdk:"default_value"`
	// The list of options for this selection field.
	Options []string `json:"options" tfsdk:"options"`
}

type _DropDownLocalIdentityField DropDownLocalIdentityField

// NewDropDownLocalIdentityField instantiates a new DropDownLocalIdentityField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDropDownLocalIdentityField(options []string, type_ string, id string, label string) *DropDownLocalIdentityField {
	this := DropDownLocalIdentityField{}
	this.Type = type_
	this.Id = id
	this.Label = label
	this.Options = options
	return &this
}

// NewDropDownLocalIdentityFieldWithDefaults instantiates a new DropDownLocalIdentityField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDropDownLocalIdentityFieldWithDefaults() *DropDownLocalIdentityField {
	this := DropDownLocalIdentityField{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *DropDownLocalIdentityField) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DropDownLocalIdentityField) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *DropDownLocalIdentityField) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *DropDownLocalIdentityField) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetOptions returns the Options field value
func (o *DropDownLocalIdentityField) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *DropDownLocalIdentityField) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *DropDownLocalIdentityField) SetOptions(v []string) {
	o.Options = v
}

func (o DropDownLocalIdentityField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DropDownLocalIdentityField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseSelectionLocalIdentityField, errBaseSelectionLocalIdentityField := json.Marshal(o.BaseSelectionLocalIdentityField)
	if errBaseSelectionLocalIdentityField != nil {
		return map[string]interface{}{}, errBaseSelectionLocalIdentityField
	}
	errBaseSelectionLocalIdentityField = json.Unmarshal([]byte(serializedBaseSelectionLocalIdentityField), &toSerialize)
	if errBaseSelectionLocalIdentityField != nil {
		return map[string]interface{}{}, errBaseSelectionLocalIdentityField
	}
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	toSerialize["options"] = o.Options
	return toSerialize, nil
}

func (o *DropDownLocalIdentityField) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"options",
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

	varDropDownLocalIdentityField := _DropDownLocalIdentityField{}

	err = json.Unmarshal(bytes, &varDropDownLocalIdentityField)

	if err != nil {
		return err
	}

	*o = DropDownLocalIdentityField(varDropDownLocalIdentityField)

	return err
}

type NullableDropDownLocalIdentityField struct {
	value *DropDownLocalIdentityField
	isSet bool
}

func (v NullableDropDownLocalIdentityField) Get() *DropDownLocalIdentityField {
	return v.value
}

func (v *NullableDropDownLocalIdentityField) Set(val *DropDownLocalIdentityField) {
	v.value = val
	v.isSet = true
}

func (v NullableDropDownLocalIdentityField) IsSet() bool {
	return v.isSet
}

func (v *NullableDropDownLocalIdentityField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDropDownLocalIdentityField(val *DropDownLocalIdentityField) *NullableDropDownLocalIdentityField {
	return &NullableDropDownLocalIdentityField{value: val, isSet: true}
}

func (v NullableDropDownLocalIdentityField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDropDownLocalIdentityField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
