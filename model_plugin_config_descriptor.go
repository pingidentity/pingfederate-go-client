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

// checks if the PluginConfigDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginConfigDescriptor{}

// PluginConfigDescriptor Defines the configuration fields available for a plugin.
type PluginConfigDescriptor struct {
	// The description of this plugin.
	Description *string `json:"description,omitempty"`
	// The configuration fields available for this plugin.
	Fields []FieldDescriptor `json:"fields,omitempty"`
	// Configuration tables available for this plugin.
	Tables []TableDescriptor `json:"tables,omitempty"`
	// The available actions for this plugin.
	ActionDescriptors []ActionDescriptor `json:"actionDescriptors,omitempty"`
}

// NewPluginConfigDescriptor instantiates a new PluginConfigDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginConfigDescriptor() *PluginConfigDescriptor {
	this := PluginConfigDescriptor{}
	return &this
}

// NewPluginConfigDescriptorWithDefaults instantiates a new PluginConfigDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginConfigDescriptorWithDefaults() *PluginConfigDescriptor {
	this := PluginConfigDescriptor{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PluginConfigDescriptor) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigDescriptor) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PluginConfigDescriptor) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PluginConfigDescriptor) SetDescription(v string) {
	o.Description = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *PluginConfigDescriptor) GetFields() []FieldDescriptor {
	if o == nil || IsNil(o.Fields) {
		var ret []FieldDescriptor
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigDescriptor) GetFieldsOk() ([]FieldDescriptor, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *PluginConfigDescriptor) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []FieldDescriptor and assigns it to the Fields field.
func (o *PluginConfigDescriptor) SetFields(v []FieldDescriptor) {
	o.Fields = v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *PluginConfigDescriptor) GetTables() []TableDescriptor {
	if o == nil || IsNil(o.Tables) {
		var ret []TableDescriptor
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigDescriptor) GetTablesOk() ([]TableDescriptor, bool) {
	if o == nil || IsNil(o.Tables) {
		return nil, false
	}
	return o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *PluginConfigDescriptor) HasTables() bool {
	if o != nil && !IsNil(o.Tables) {
		return true
	}

	return false
}

// SetTables gets a reference to the given []TableDescriptor and assigns it to the Tables field.
func (o *PluginConfigDescriptor) SetTables(v []TableDescriptor) {
	o.Tables = v
}

// GetActionDescriptors returns the ActionDescriptors field value if set, zero value otherwise.
func (o *PluginConfigDescriptor) GetActionDescriptors() []ActionDescriptor {
	if o == nil || IsNil(o.ActionDescriptors) {
		var ret []ActionDescriptor
		return ret
	}
	return o.ActionDescriptors
}

// GetActionDescriptorsOk returns a tuple with the ActionDescriptors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginConfigDescriptor) GetActionDescriptorsOk() ([]ActionDescriptor, bool) {
	if o == nil || IsNil(o.ActionDescriptors) {
		return nil, false
	}
	return o.ActionDescriptors, true
}

// HasActionDescriptors returns a boolean if a field has been set.
func (o *PluginConfigDescriptor) HasActionDescriptors() bool {
	if o != nil && !IsNil(o.ActionDescriptors) {
		return true
	}

	return false
}

// SetActionDescriptors gets a reference to the given []ActionDescriptor and assigns it to the ActionDescriptors field.
func (o *PluginConfigDescriptor) SetActionDescriptors(v []ActionDescriptor) {
	o.ActionDescriptors = v
}

func (o PluginConfigDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginConfigDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Tables) {
		toSerialize["tables"] = o.Tables
	}
	if !IsNil(o.ActionDescriptors) {
		toSerialize["actionDescriptors"] = o.ActionDescriptors
	}
	return toSerialize, nil
}

type NullablePluginConfigDescriptor struct {
	value *PluginConfigDescriptor
	isSet bool
}

func (v NullablePluginConfigDescriptor) Get() *PluginConfigDescriptor {
	return v.value
}

func (v *NullablePluginConfigDescriptor) Set(val *PluginConfigDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginConfigDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginConfigDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginConfigDescriptor(val *PluginConfigDescriptor) *NullablePluginConfigDescriptor {
	return &NullablePluginConfigDescriptor{value: val, isSet: true}
}

func (v NullablePluginConfigDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginConfigDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
