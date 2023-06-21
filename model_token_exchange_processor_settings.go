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

// checks if the TokenExchangeProcessorSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenExchangeProcessorSettings{}

// TokenExchangeProcessorSettings Settings for the OAuth Token Exchange Processor Policy configuration.
type TokenExchangeProcessorSettings struct {
	DefaultProcessorPolicyRef *ResourceLink `json:"defaultProcessorPolicyRef,omitempty"`
}

// NewTokenExchangeProcessorSettings instantiates a new TokenExchangeProcessorSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenExchangeProcessorSettings() *TokenExchangeProcessorSettings {
	this := TokenExchangeProcessorSettings{}
	return &this
}

// NewTokenExchangeProcessorSettingsWithDefaults instantiates a new TokenExchangeProcessorSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenExchangeProcessorSettingsWithDefaults() *TokenExchangeProcessorSettings {
	this := TokenExchangeProcessorSettings{}
	return &this
}

// GetDefaultProcessorPolicyRef returns the DefaultProcessorPolicyRef field value if set, zero value otherwise.
func (o *TokenExchangeProcessorSettings) GetDefaultProcessorPolicyRef() ResourceLink {
	if o == nil || IsNil(o.DefaultProcessorPolicyRef) {
		var ret ResourceLink
		return ret
	}
	return *o.DefaultProcessorPolicyRef
}

// GetDefaultProcessorPolicyRefOk returns a tuple with the DefaultProcessorPolicyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenExchangeProcessorSettings) GetDefaultProcessorPolicyRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.DefaultProcessorPolicyRef) {
		return nil, false
	}
	return o.DefaultProcessorPolicyRef, true
}

// HasDefaultProcessorPolicyRef returns a boolean if a field has been set.
func (o *TokenExchangeProcessorSettings) HasDefaultProcessorPolicyRef() bool {
	if o != nil && !IsNil(o.DefaultProcessorPolicyRef) {
		return true
	}

	return false
}

// SetDefaultProcessorPolicyRef gets a reference to the given ResourceLink and assigns it to the DefaultProcessorPolicyRef field.
func (o *TokenExchangeProcessorSettings) SetDefaultProcessorPolicyRef(v ResourceLink) {
	o.DefaultProcessorPolicyRef = &v
}

func (o TokenExchangeProcessorSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenExchangeProcessorSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultProcessorPolicyRef) {
		toSerialize["defaultProcessorPolicyRef"] = o.DefaultProcessorPolicyRef
	}
	return toSerialize, nil
}

type NullableTokenExchangeProcessorSettings struct {
	value *TokenExchangeProcessorSettings
	isSet bool
}

func (v NullableTokenExchangeProcessorSettings) Get() *TokenExchangeProcessorSettings {
	return v.value
}

func (v *NullableTokenExchangeProcessorSettings) Set(val *TokenExchangeProcessorSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenExchangeProcessorSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenExchangeProcessorSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenExchangeProcessorSettings(val *TokenExchangeProcessorSettings) *NullableTokenExchangeProcessorSettings {
	return &NullableTokenExchangeProcessorSettings{value: val, isSet: true}
}

func (v NullableTokenExchangeProcessorSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenExchangeProcessorSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
