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

// checks if the FieldConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldConfig{}

// FieldConfig A local identity profile field configuration.
type FieldConfig struct {
	// The field configuration for the local identity profile.
	Fields []LocalIdentityField `json:"fields,omitempty" tfsdk:"fields"`
	// Strip leading/trailing spaces from unique ID field. Default is true.
	StripSpaceFromUniqueField *bool `json:"stripSpaceFromUniqueField,omitempty" tfsdk:"strip_space_from_unique_field"`
}

// NewFieldConfig instantiates a new FieldConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldConfig() *FieldConfig {
	this := FieldConfig{}
	return &this
}

// NewFieldConfigWithDefaults instantiates a new FieldConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldConfigWithDefaults() *FieldConfig {
	this := FieldConfig{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *FieldConfig) GetFields() []LocalIdentityField {
	if o == nil || IsNil(o.Fields) {
		var ret []LocalIdentityField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldConfig) GetFieldsOk() ([]LocalIdentityField, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *FieldConfig) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []LocalIdentityField and assigns it to the Fields field.
func (o *FieldConfig) SetFields(v []LocalIdentityField) {
	o.Fields = v
}

// GetStripSpaceFromUniqueField returns the StripSpaceFromUniqueField field value if set, zero value otherwise.
func (o *FieldConfig) GetStripSpaceFromUniqueField() bool {
	if o == nil || IsNil(o.StripSpaceFromUniqueField) {
		var ret bool
		return ret
	}
	return *o.StripSpaceFromUniqueField
}

// GetStripSpaceFromUniqueFieldOk returns a tuple with the StripSpaceFromUniqueField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldConfig) GetStripSpaceFromUniqueFieldOk() (*bool, bool) {
	if o == nil || IsNil(o.StripSpaceFromUniqueField) {
		return nil, false
	}
	return o.StripSpaceFromUniqueField, true
}

// HasStripSpaceFromUniqueField returns a boolean if a field has been set.
func (o *FieldConfig) HasStripSpaceFromUniqueField() bool {
	if o != nil && !IsNil(o.StripSpaceFromUniqueField) {
		return true
	}

	return false
}

// SetStripSpaceFromUniqueField gets a reference to the given bool and assigns it to the StripSpaceFromUniqueField field.
func (o *FieldConfig) SetStripSpaceFromUniqueField(v bool) {
	o.StripSpaceFromUniqueField = &v
}

func (o FieldConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.StripSpaceFromUniqueField) {
		toSerialize["stripSpaceFromUniqueField"] = o.StripSpaceFromUniqueField
	}
	return toSerialize, nil
}

type NullableFieldConfig struct {
	value *FieldConfig
	isSet bool
}

func (v NullableFieldConfig) Get() *FieldConfig {
	return v.value
}

func (v *NullableFieldConfig) Set(val *FieldConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldConfig(val *FieldConfig) *NullableFieldConfig {
	return &NullableFieldConfig{value: val, isSet: true}
}

func (v NullableFieldConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}