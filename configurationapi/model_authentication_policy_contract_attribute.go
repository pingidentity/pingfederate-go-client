/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AuthenticationPolicyContractAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationPolicyContractAttribute{}

// AuthenticationPolicyContractAttribute An attribute for the Authentication Policy Contract.
type AuthenticationPolicyContractAttribute struct {
	// The name of this attribute.
	Name string `json:"name" tfsdk:"name"`
}

// NewAuthenticationPolicyContractAttribute instantiates a new AuthenticationPolicyContractAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationPolicyContractAttribute(name string) *AuthenticationPolicyContractAttribute {
	this := AuthenticationPolicyContractAttribute{}
	this.Name = name
	return &this
}

// NewAuthenticationPolicyContractAttributeWithDefaults instantiates a new AuthenticationPolicyContractAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationPolicyContractAttributeWithDefaults() *AuthenticationPolicyContractAttribute {
	this := AuthenticationPolicyContractAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *AuthenticationPolicyContractAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthenticationPolicyContractAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthenticationPolicyContractAttribute) SetName(v string) {
	o.Name = v
}

func (o AuthenticationPolicyContractAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationPolicyContractAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableAuthenticationPolicyContractAttribute struct {
	value *AuthenticationPolicyContractAttribute
	isSet bool
}

func (v NullableAuthenticationPolicyContractAttribute) Get() *AuthenticationPolicyContractAttribute {
	return v.value
}

func (v *NullableAuthenticationPolicyContractAttribute) Set(val *AuthenticationPolicyContractAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationPolicyContractAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationPolicyContractAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationPolicyContractAttribute(val *AuthenticationPolicyContractAttribute) *NullableAuthenticationPolicyContractAttribute {
	return &NullableAuthenticationPolicyContractAttribute{value: val, isSet: true}
}

func (v NullableAuthenticationPolicyContractAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationPolicyContractAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
