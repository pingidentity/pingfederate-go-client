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

// checks if the SaasPluginFieldOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaasPluginFieldOption{}

// SaasPluginFieldOption A plugin configuration field value.
type SaasPluginFieldOption struct {
	// The code that represents the field.
	Code string `json:"code" tfsdk:"code"`
	// The label for the field.
	Label string `json:"label" tfsdk:"label"`
}

// NewSaasPluginFieldOption instantiates a new SaasPluginFieldOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaasPluginFieldOption(code string, label string) *SaasPluginFieldOption {
	this := SaasPluginFieldOption{}
	this.Code = code
	this.Label = label
	return &this
}

// NewSaasPluginFieldOptionWithDefaults instantiates a new SaasPluginFieldOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaasPluginFieldOptionWithDefaults() *SaasPluginFieldOption {
	this := SaasPluginFieldOption{}
	return &this
}

// GetCode returns the Code field value
func (o *SaasPluginFieldOption) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SaasPluginFieldOption) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SaasPluginFieldOption) SetCode(v string) {
	o.Code = v
}

// GetLabel returns the Label field value
func (o *SaasPluginFieldOption) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *SaasPluginFieldOption) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *SaasPluginFieldOption) SetLabel(v string) {
	o.Label = v
}

func (o SaasPluginFieldOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaasPluginFieldOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

type NullableSaasPluginFieldOption struct {
	value *SaasPluginFieldOption
	isSet bool
}

func (v NullableSaasPluginFieldOption) Get() *SaasPluginFieldOption {
	return v.value
}

func (v *NullableSaasPluginFieldOption) Set(val *SaasPluginFieldOption) {
	v.value = val
	v.isSet = true
}

func (v NullableSaasPluginFieldOption) IsSet() bool {
	return v.isSet
}

func (v *NullableSaasPluginFieldOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaasPluginFieldOption(val *SaasPluginFieldOption) *NullableSaasPluginFieldOption {
	return &NullableSaasPluginFieldOption{value: val, isSet: true}
}

func (v NullableSaasPluginFieldOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaasPluginFieldOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
