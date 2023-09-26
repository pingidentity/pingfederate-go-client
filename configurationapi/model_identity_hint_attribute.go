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

// checks if the IdentityHintAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityHintAttribute{}

// IdentityHintAttribute An attribute for the ciba request policy's identity hint attribute contract.
type IdentityHintAttribute struct {
	// The name of this attribute.
	Name string `json:"name" tfsdk:"name"`
}

// NewIdentityHintAttribute instantiates a new IdentityHintAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityHintAttribute(name string) *IdentityHintAttribute {
	this := IdentityHintAttribute{}
	this.Name = name
	return &this
}

// NewIdentityHintAttributeWithDefaults instantiates a new IdentityHintAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityHintAttributeWithDefaults() *IdentityHintAttribute {
	this := IdentityHintAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *IdentityHintAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdentityHintAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdentityHintAttribute) SetName(v string) {
	o.Name = v
}

func (o IdentityHintAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityHintAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableIdentityHintAttribute struct {
	value *IdentityHintAttribute
	isSet bool
}

func (v NullableIdentityHintAttribute) Get() *IdentityHintAttribute {
	return v.value
}

func (v *NullableIdentityHintAttribute) Set(val *IdentityHintAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityHintAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityHintAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityHintAttribute(val *IdentityHintAttribute) *NullableIdentityHintAttribute {
	return &NullableIdentityHintAttribute{value: val, isSet: true}
}

func (v NullableIdentityHintAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityHintAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}