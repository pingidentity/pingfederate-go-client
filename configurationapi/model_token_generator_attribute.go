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

// checks if the TokenGeneratorAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenGeneratorAttribute{}

// TokenGeneratorAttribute An attribute for the token generator attribute contract.
type TokenGeneratorAttribute struct {
	// The name of this attribute.
	Name string `json:"name" tfsdk:"name"`
}

// NewTokenGeneratorAttribute instantiates a new TokenGeneratorAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenGeneratorAttribute(name string) *TokenGeneratorAttribute {
	this := TokenGeneratorAttribute{}
	this.Name = name
	return &this
}

// NewTokenGeneratorAttributeWithDefaults instantiates a new TokenGeneratorAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenGeneratorAttributeWithDefaults() *TokenGeneratorAttribute {
	this := TokenGeneratorAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *TokenGeneratorAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TokenGeneratorAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TokenGeneratorAttribute) SetName(v string) {
	o.Name = v
}

func (o TokenGeneratorAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenGeneratorAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableTokenGeneratorAttribute struct {
	value *TokenGeneratorAttribute
	isSet bool
}

func (v NullableTokenGeneratorAttribute) Get() *TokenGeneratorAttribute {
	return v.value
}

func (v *NullableTokenGeneratorAttribute) Set(val *TokenGeneratorAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenGeneratorAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenGeneratorAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenGeneratorAttribute(val *TokenGeneratorAttribute) *NullableTokenGeneratorAttribute {
	return &NullableTokenGeneratorAttribute{value: val, isSet: true}
}

func (v NullableTokenGeneratorAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenGeneratorAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
