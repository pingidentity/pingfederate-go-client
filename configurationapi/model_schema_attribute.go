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

// checks if the SchemaAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaAttribute{}

// SchemaAttribute A custom SCIM attribute.
type SchemaAttribute struct {
	// Name of the attribute.
	Name *string `json:"name,omitempty" tfsdk:"name"`
	// Indicates whether the attribute is multi-valued.
	MultiValued *bool `json:"multiValued,omitempty" tfsdk:"multi_valued"`
	// Represents the name of each attribute type in case of multi-valued attribute.
	Types []string `json:"types,omitempty" tfsdk:"types"`
	// List of sub-attributes for an attribute.
	SubAttributes []string `json:"subAttributes,omitempty" tfsdk:"sub_attributes"`
}

// NewSchemaAttribute instantiates a new SchemaAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaAttribute() *SchemaAttribute {
	this := SchemaAttribute{}
	return &this
}

// NewSchemaAttributeWithDefaults instantiates a new SchemaAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaAttributeWithDefaults() *SchemaAttribute {
	this := SchemaAttribute{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SchemaAttribute) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SchemaAttribute) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SchemaAttribute) SetName(v string) {
	o.Name = &v
}

// GetMultiValued returns the MultiValued field value if set, zero value otherwise.
func (o *SchemaAttribute) GetMultiValued() bool {
	if o == nil || IsNil(o.MultiValued) {
		var ret bool
		return ret
	}
	return *o.MultiValued
}

// GetMultiValuedOk returns a tuple with the MultiValued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetMultiValuedOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiValued) {
		return nil, false
	}
	return o.MultiValued, true
}

// HasMultiValued returns a boolean if a field has been set.
func (o *SchemaAttribute) HasMultiValued() bool {
	if o != nil && !IsNil(o.MultiValued) {
		return true
	}

	return false
}

// SetMultiValued gets a reference to the given bool and assigns it to the MultiValued field.
func (o *SchemaAttribute) SetMultiValued(v bool) {
	o.MultiValued = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *SchemaAttribute) GetTypes() []string {
	if o == nil || IsNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *SchemaAttribute) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *SchemaAttribute) SetTypes(v []string) {
	o.Types = v
}

// GetSubAttributes returns the SubAttributes field value if set, zero value otherwise.
func (o *SchemaAttribute) GetSubAttributes() []string {
	if o == nil || IsNil(o.SubAttributes) {
		var ret []string
		return ret
	}
	return o.SubAttributes
}

// GetSubAttributesOk returns a tuple with the SubAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetSubAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.SubAttributes) {
		return nil, false
	}
	return o.SubAttributes, true
}

// HasSubAttributes returns a boolean if a field has been set.
func (o *SchemaAttribute) HasSubAttributes() bool {
	if o != nil && !IsNil(o.SubAttributes) {
		return true
	}

	return false
}

// SetSubAttributes gets a reference to the given []string and assigns it to the SubAttributes field.
func (o *SchemaAttribute) SetSubAttributes(v []string) {
	o.SubAttributes = v
}

func (o SchemaAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.MultiValued) {
		toSerialize["multiValued"] = o.MultiValued
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.SubAttributes) {
		toSerialize["subAttributes"] = o.SubAttributes
	}
	return toSerialize, nil
}

type NullableSchemaAttribute struct {
	value *SchemaAttribute
	isSet bool
}

func (v NullableSchemaAttribute) Get() *SchemaAttribute {
	return v.value
}

func (v *NullableSchemaAttribute) Set(val *SchemaAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaAttribute(val *SchemaAttribute) *NullableSchemaAttribute {
	return &NullableSchemaAttribute{value: val, isSet: true}
}

func (v NullableSchemaAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
