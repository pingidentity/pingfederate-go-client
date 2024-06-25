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

// checks if the PluginConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginConfiguration{}

// PluginConfiguration Configuration settings for a plugin instance.
type PluginConfiguration struct {
	// List of configuration tables.
	Tables []ConfigTable `json:"tables,omitempty" tfsdk:"tables"`
	// List of configuration fields.
	Fields []ConfigField `json:"fields,omitempty" tfsdk:"fields"`
}

// NewPluginConfiguration instantiates a new PluginConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginConfiguration() *PluginConfiguration {
	this := PluginConfiguration{}
	return &this
}

// NewPluginConfigurationWithDefaults instantiates a new PluginConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginConfigurationWithDefaults() *PluginConfiguration {
	this := PluginConfiguration{}
	return &this
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *PluginConfiguration) GetTables() []ConfigTable {
	if o == nil || IsNil(o.Tables) {
		var ret []ConfigTable
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfiguration) GetTablesOk() ([]ConfigTable, bool) {
	if o == nil || IsNil(o.Tables) {
		return nil, false
	}
	return o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *PluginConfiguration) HasTables() bool {
	if o != nil && !IsNil(o.Tables) {
		return true
	}

	return false
}

// SetTables gets a reference to the given []ConfigTable and assigns it to the Tables field.
func (o *PluginConfiguration) SetTables(v []ConfigTable) {
	o.Tables = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *PluginConfiguration) GetFields() []ConfigField {
	if o == nil || IsNil(o.Fields) {
		var ret []ConfigField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfiguration) GetFieldsOk() ([]ConfigField, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *PluginConfiguration) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []ConfigField and assigns it to the Fields field.
func (o *PluginConfiguration) SetFields(v []ConfigField) {
	o.Fields = v
}

func (o PluginConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tables) {
		toSerialize["tables"] = o.Tables
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}

type NullablePluginConfiguration struct {
	value *PluginConfiguration
	isSet bool
}

func (v NullablePluginConfiguration) Get() *PluginConfiguration {
	return v.value
}

func (v *NullablePluginConfiguration) Set(val *PluginConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginConfiguration(val *PluginConfiguration) *NullablePluginConfiguration {
	return &NullablePluginConfiguration{value: val, isSet: true}
}

func (v NullablePluginConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
