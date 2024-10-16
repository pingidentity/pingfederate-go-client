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

// checks if the BaseSelectionFieldDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseSelectionFieldDescriptor{}

// BaseSelectionFieldDescriptor struct for BaseSelectionFieldDescriptor
type BaseSelectionFieldDescriptor struct {
	FieldDescriptor
	// The list of option values for this selection field.
	OptionValues []OptionValue `json:"optionValues,omitempty" tfsdk:"option_values"`
}

// NewBaseSelectionFieldDescriptor instantiates a new BaseSelectionFieldDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseSelectionFieldDescriptor() *BaseSelectionFieldDescriptor {
	this := BaseSelectionFieldDescriptor{}
	return &this
}

// NewBaseSelectionFieldDescriptorWithDefaults instantiates a new BaseSelectionFieldDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseSelectionFieldDescriptorWithDefaults() *BaseSelectionFieldDescriptor {
	this := BaseSelectionFieldDescriptor{}
	return &this
}

// GetOptionValues returns the OptionValues field value if set, zero value otherwise.
func (o *BaseSelectionFieldDescriptor) GetOptionValues() []OptionValue {
	if o == nil || IsNil(o.OptionValues) {
		var ret []OptionValue
		return ret
	}
	return o.OptionValues
}

// GetOptionValuesOk returns a tuple with the OptionValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSelectionFieldDescriptor) GetOptionValuesOk() ([]OptionValue, bool) {
	if o == nil || IsNil(o.OptionValues) {
		return nil, false
	}
	return o.OptionValues, true
}

// HasOptionValues returns a boolean if a field has been set.
func (o *BaseSelectionFieldDescriptor) HasOptionValues() bool {
	if o != nil && !IsNil(o.OptionValues) {
		return true
	}

	return false
}

// SetOptionValues gets a reference to the given []OptionValue and assigns it to the OptionValues field.
func (o *BaseSelectionFieldDescriptor) SetOptionValues(v []OptionValue) {
	o.OptionValues = v
}

func (o BaseSelectionFieldDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseSelectionFieldDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFieldDescriptor, errFieldDescriptor := json.Marshal(o.FieldDescriptor)
	if errFieldDescriptor != nil {
		return map[string]interface{}{}, errFieldDescriptor
	}
	errFieldDescriptor = json.Unmarshal([]byte(serializedFieldDescriptor), &toSerialize)
	if errFieldDescriptor != nil {
		return map[string]interface{}{}, errFieldDescriptor
	}
	if !IsNil(o.OptionValues) {
		toSerialize["optionValues"] = o.OptionValues
	}
	return toSerialize, nil
}

type NullableBaseSelectionFieldDescriptor struct {
	value *BaseSelectionFieldDescriptor
	isSet bool
}

func (v NullableBaseSelectionFieldDescriptor) Get() *BaseSelectionFieldDescriptor {
	return v.value
}

func (v *NullableBaseSelectionFieldDescriptor) Set(val *BaseSelectionFieldDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseSelectionFieldDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseSelectionFieldDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseSelectionFieldDescriptor(val *BaseSelectionFieldDescriptor) *NullableBaseSelectionFieldDescriptor {
	return &NullableBaseSelectionFieldDescriptor{value: val, isSet: true}
}

func (v NullableBaseSelectionFieldDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseSelectionFieldDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
