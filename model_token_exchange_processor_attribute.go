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

// checks if the TokenExchangeProcessorAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenExchangeProcessorAttribute{}

// TokenExchangeProcessorAttribute An attribute for the OAuth 2.0 Token Exchange Processor policy attribute contract.
type TokenExchangeProcessorAttribute struct {
	// The name of this attribute.
	Name string `json:"name" tfsdk:"name"`
}

// NewTokenExchangeProcessorAttribute instantiates a new TokenExchangeProcessorAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenExchangeProcessorAttribute(name string) *TokenExchangeProcessorAttribute {
	this := TokenExchangeProcessorAttribute{}
	this.Name = name
	return &this
}

// NewTokenExchangeProcessorAttributeWithDefaults instantiates a new TokenExchangeProcessorAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenExchangeProcessorAttributeWithDefaults() *TokenExchangeProcessorAttribute {
	this := TokenExchangeProcessorAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *TokenExchangeProcessorAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TokenExchangeProcessorAttribute) SetName(v string) {
	o.Name = v
}

func (o TokenExchangeProcessorAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenExchangeProcessorAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableTokenExchangeProcessorAttribute struct {
	value *TokenExchangeProcessorAttribute
	isSet bool
}

func (v NullableTokenExchangeProcessorAttribute) Get() *TokenExchangeProcessorAttribute {
	return v.value
}

func (v *NullableTokenExchangeProcessorAttribute) Set(val *TokenExchangeProcessorAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenExchangeProcessorAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenExchangeProcessorAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenExchangeProcessorAttribute(val *TokenExchangeProcessorAttribute) *NullableTokenExchangeProcessorAttribute {
	return &NullableTokenExchangeProcessorAttribute{value: val, isSet: true}
}

func (v NullableTokenExchangeProcessorAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenExchangeProcessorAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
