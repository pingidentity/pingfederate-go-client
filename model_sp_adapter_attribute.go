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

// checks if the SpAdapterAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpAdapterAttribute{}

// SpAdapterAttribute An attribute for the SP adapter attribute contract.
type SpAdapterAttribute struct {
	// The name of this attribute.
	Name string `json:"name"`
}

// NewSpAdapterAttribute instantiates a new SpAdapterAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpAdapterAttribute(name string) *SpAdapterAttribute {
	this := SpAdapterAttribute{}
	this.Name = name
	return &this
}

// NewSpAdapterAttributeWithDefaults instantiates a new SpAdapterAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpAdapterAttributeWithDefaults() *SpAdapterAttribute {
	this := SpAdapterAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *SpAdapterAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SpAdapterAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SpAdapterAttribute) SetName(v string) {
	o.Name = v
}

func (o SpAdapterAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpAdapterAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableSpAdapterAttribute struct {
	value *SpAdapterAttribute
	isSet bool
}

func (v NullableSpAdapterAttribute) Get() *SpAdapterAttribute {
	return v.value
}

func (v *NullableSpAdapterAttribute) Set(val *SpAdapterAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableSpAdapterAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableSpAdapterAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpAdapterAttribute(val *SpAdapterAttribute) *NullableSpAdapterAttribute {
	return &NullableSpAdapterAttribute{value: val, isSet: true}
}

func (v NullableSpAdapterAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpAdapterAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
