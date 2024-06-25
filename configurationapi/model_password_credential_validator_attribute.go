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

// checks if the PasswordCredentialValidatorAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordCredentialValidatorAttribute{}

// PasswordCredentialValidatorAttribute An attribute for the password credential validator attribute contract.
type PasswordCredentialValidatorAttribute struct {
	// The name of this attribute.
	Name string `json:"name" tfsdk:"name"`
}

// NewPasswordCredentialValidatorAttribute instantiates a new PasswordCredentialValidatorAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordCredentialValidatorAttribute(name string) *PasswordCredentialValidatorAttribute {
	this := PasswordCredentialValidatorAttribute{}
	this.Name = name
	return &this
}

// NewPasswordCredentialValidatorAttributeWithDefaults instantiates a new PasswordCredentialValidatorAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordCredentialValidatorAttributeWithDefaults() *PasswordCredentialValidatorAttribute {
	this := PasswordCredentialValidatorAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *PasswordCredentialValidatorAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PasswordCredentialValidatorAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PasswordCredentialValidatorAttribute) SetName(v string) {
	o.Name = v
}

func (o PasswordCredentialValidatorAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordCredentialValidatorAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullablePasswordCredentialValidatorAttribute struct {
	value *PasswordCredentialValidatorAttribute
	isSet bool
}

func (v NullablePasswordCredentialValidatorAttribute) Get() *PasswordCredentialValidatorAttribute {
	return v.value
}

func (v *NullablePasswordCredentialValidatorAttribute) Set(val *PasswordCredentialValidatorAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordCredentialValidatorAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordCredentialValidatorAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordCredentialValidatorAttribute(val *PasswordCredentialValidatorAttribute) *NullablePasswordCredentialValidatorAttribute {
	return &NullablePasswordCredentialValidatorAttribute{value: val, isSet: true}
}

func (v NullablePasswordCredentialValidatorAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordCredentialValidatorAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
