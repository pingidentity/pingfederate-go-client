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

// checks if the BaseDefaultValueLocalIdentityField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseDefaultValueLocalIdentityField{}

// BaseDefaultValueLocalIdentityField struct for BaseDefaultValueLocalIdentityField
type BaseDefaultValueLocalIdentityField struct {
	LocalIdentityField
	// The default value for this field.
	DefaultValue *string `json:"defaultValue,omitempty" tfsdk:"default_value"`
}

// NewBaseDefaultValueLocalIdentityField instantiates a new BaseDefaultValueLocalIdentityField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseDefaultValueLocalIdentityField(type_ string, id string, label string) *BaseDefaultValueLocalIdentityField {
	this := BaseDefaultValueLocalIdentityField{}
	this.Type = type_
	this.Id = id
	this.Label = label
	return &this
}

// NewBaseDefaultValueLocalIdentityFieldWithDefaults instantiates a new BaseDefaultValueLocalIdentityField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseDefaultValueLocalIdentityFieldWithDefaults() *BaseDefaultValueLocalIdentityField {
	this := BaseDefaultValueLocalIdentityField{}
	return &this
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *BaseDefaultValueLocalIdentityField) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseDefaultValueLocalIdentityField) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *BaseDefaultValueLocalIdentityField) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *BaseDefaultValueLocalIdentityField) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

func (o BaseDefaultValueLocalIdentityField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseDefaultValueLocalIdentityField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedLocalIdentityField, errLocalIdentityField := json.Marshal(o.LocalIdentityField)
	if errLocalIdentityField != nil {
		return map[string]interface{}{}, errLocalIdentityField
	}
	errLocalIdentityField = json.Unmarshal([]byte(serializedLocalIdentityField), &toSerialize)
	if errLocalIdentityField != nil {
		return map[string]interface{}{}, errLocalIdentityField
	}
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	return toSerialize, nil
}

type NullableBaseDefaultValueLocalIdentityField struct {
	value *BaseDefaultValueLocalIdentityField
	isSet bool
}

func (v NullableBaseDefaultValueLocalIdentityField) Get() *BaseDefaultValueLocalIdentityField {
	return v.value
}

func (v *NullableBaseDefaultValueLocalIdentityField) Set(val *BaseDefaultValueLocalIdentityField) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseDefaultValueLocalIdentityField) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseDefaultValueLocalIdentityField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseDefaultValueLocalIdentityField(val *BaseDefaultValueLocalIdentityField) *NullableBaseDefaultValueLocalIdentityField {
	return &NullableBaseDefaultValueLocalIdentityField{value: val, isSet: true}
}

func (v NullableBaseDefaultValueLocalIdentityField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseDefaultValueLocalIdentityField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
