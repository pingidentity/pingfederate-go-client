/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the ConfigField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigField{}

// ConfigField A plugin configuration field value.
type ConfigField struct {
	// The name of the configuration field.
	Name string `json:"name" tfsdk:"name"`
	// The value for the configuration field. For encrypted or hashed fields, GETs will not return this attribute. To update an encrypted or hashed field, specify the new value in this attribute.
	Value *string `json:"value,omitempty" tfsdk:"value"`
	// For encrypted or hashed fields, this attribute contains the encrypted representation of the field's value, if a value is defined. If you do not want to update the stored value, this attribute should be passed back unchanged.
	EncryptedValue *string `json:"encryptedValue,omitempty" tfsdk:"encrypted_value"`
}

// NewConfigField instantiates a new ConfigField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigField(name string) *ConfigField {
	this := ConfigField{}
	this.Name = name
	return &this
}

// NewConfigFieldWithDefaults instantiates a new ConfigField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigFieldWithDefaults() *ConfigField {
	this := ConfigField{}
	return &this
}

// GetName returns the Name field value
func (o *ConfigField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigField) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConfigField) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigField) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConfigField) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ConfigField) SetValue(v string) {
	o.Value = &v
}

// GetEncryptedValue returns the EncryptedValue field value if set, zero value otherwise.
func (o *ConfigField) GetEncryptedValue() string {
	if o == nil || IsNil(o.EncryptedValue) {
		var ret string
		return ret
	}
	return *o.EncryptedValue
}

// GetEncryptedValueOk returns a tuple with the EncryptedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigField) GetEncryptedValueOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedValue) {
		return nil, false
	}
	return o.EncryptedValue, true
}

// HasEncryptedValue returns a boolean if a field has been set.
func (o *ConfigField) HasEncryptedValue() bool {
	if o != nil && !IsNil(o.EncryptedValue) {
		return true
	}

	return false
}

// SetEncryptedValue gets a reference to the given string and assigns it to the EncryptedValue field.
func (o *ConfigField) SetEncryptedValue(v string) {
	o.EncryptedValue = &v
}

func (o ConfigField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.EncryptedValue) {
		toSerialize["encryptedValue"] = o.EncryptedValue
	}
	return toSerialize, nil
}

type NullableConfigField struct {
	value *ConfigField
	isSet bool
}

func (v NullableConfigField) Get() *ConfigField {
	return v.value
}

func (v *NullableConfigField) Set(val *ConfigField) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigField) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigField(val *ConfigField) *NullableConfigField {
	return &NullableConfigField{value: val, isSet: true}
}

func (v NullableConfigField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
