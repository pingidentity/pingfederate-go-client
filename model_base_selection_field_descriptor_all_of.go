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

// checks if the BaseSelectionFieldDescriptorAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseSelectionFieldDescriptorAllOf{}

// BaseSelectionFieldDescriptorAllOf Holds fields that are shared by all selection-type field descriptors.
type BaseSelectionFieldDescriptorAllOf struct {
	// The list of option values for this selection field.
	OptionValues []OptionValue `json:"optionValues,omitempty" tfsdk:"option_values"`
}

// NewBaseSelectionFieldDescriptorAllOf instantiates a new BaseSelectionFieldDescriptorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseSelectionFieldDescriptorAllOf() *BaseSelectionFieldDescriptorAllOf {
	this := BaseSelectionFieldDescriptorAllOf{}
	return &this
}

// NewBaseSelectionFieldDescriptorAllOfWithDefaults instantiates a new BaseSelectionFieldDescriptorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseSelectionFieldDescriptorAllOfWithDefaults() *BaseSelectionFieldDescriptorAllOf {
	this := BaseSelectionFieldDescriptorAllOf{}
	return &this
}

// GetOptionValues returns the OptionValues field value if set, zero value otherwise.
func (o *BaseSelectionFieldDescriptorAllOf) GetOptionValues() []OptionValue {
	if o == nil || IsNil(o.OptionValues) {
		var ret []OptionValue
		return ret
	}
	return o.OptionValues
}

// GetOptionValuesOk returns a tuple with the OptionValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSelectionFieldDescriptorAllOf) GetOptionValuesOk() ([]OptionValue, bool) {
	if o == nil || IsNil(o.OptionValues) {
		return nil, false
	}
	return o.OptionValues, true
}

// HasOptionValues returns a boolean if a field has been set.
func (o *BaseSelectionFieldDescriptorAllOf) HasOptionValues() bool {
	if o != nil && !IsNil(o.OptionValues) {
		return true
	}

	return false
}

// SetOptionValues gets a reference to the given []OptionValue and assigns it to the OptionValues field.
func (o *BaseSelectionFieldDescriptorAllOf) SetOptionValues(v []OptionValue) {
	o.OptionValues = v
}

func (o BaseSelectionFieldDescriptorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseSelectionFieldDescriptorAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptionValues) {
		toSerialize["optionValues"] = o.OptionValues
	}
	return toSerialize, nil
}

type NullableBaseSelectionFieldDescriptorAllOf struct {
	value *BaseSelectionFieldDescriptorAllOf
	isSet bool
}

func (v NullableBaseSelectionFieldDescriptorAllOf) Get() *BaseSelectionFieldDescriptorAllOf {
	return v.value
}

func (v *NullableBaseSelectionFieldDescriptorAllOf) Set(val *BaseSelectionFieldDescriptorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseSelectionFieldDescriptorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseSelectionFieldDescriptorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseSelectionFieldDescriptorAllOf(val *BaseSelectionFieldDescriptorAllOf) *NullableBaseSelectionFieldDescriptorAllOf {
	return &NullableBaseSelectionFieldDescriptorAllOf{value: val, isSet: true}
}

func (v NullableBaseSelectionFieldDescriptorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseSelectionFieldDescriptorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
