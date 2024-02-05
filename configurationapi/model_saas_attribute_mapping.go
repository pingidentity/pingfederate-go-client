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

// checks if the SaasAttributeMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaasAttributeMapping{}

// SaasAttributeMapping Settings to map the source record attributes to target attributes.
type SaasAttributeMapping struct {
	// The name of target field.
	FieldName     string                 `json:"fieldName" tfsdk:"field_name"`
	SaasFieldInfo SaasFieldConfiguration `json:"saasFieldInfo" tfsdk:"saas_field_info"`
}

// NewSaasAttributeMapping instantiates a new SaasAttributeMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaasAttributeMapping(fieldName string, saasFieldInfo SaasFieldConfiguration) *SaasAttributeMapping {
	this := SaasAttributeMapping{}
	this.FieldName = fieldName
	this.SaasFieldInfo = saasFieldInfo
	return &this
}

// NewSaasAttributeMappingWithDefaults instantiates a new SaasAttributeMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaasAttributeMappingWithDefaults() *SaasAttributeMapping {
	this := SaasAttributeMapping{}
	return &this
}

// GetFieldName returns the FieldName field value
func (o *SaasAttributeMapping) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *SaasAttributeMapping) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *SaasAttributeMapping) SetFieldName(v string) {
	o.FieldName = v
}

// GetSaasFieldInfo returns the SaasFieldInfo field value
func (o *SaasAttributeMapping) GetSaasFieldInfo() SaasFieldConfiguration {
	if o == nil {
		var ret SaasFieldConfiguration
		return ret
	}

	return o.SaasFieldInfo
}

// GetSaasFieldInfoOk returns a tuple with the SaasFieldInfo field value
// and a boolean to check if the value has been set.
func (o *SaasAttributeMapping) GetSaasFieldInfoOk() (*SaasFieldConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SaasFieldInfo, true
}

// SetSaasFieldInfo sets field value
func (o *SaasAttributeMapping) SetSaasFieldInfo(v SaasFieldConfiguration) {
	o.SaasFieldInfo = v
}

func (o SaasAttributeMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaasAttributeMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fieldName"] = o.FieldName
	toSerialize["saasFieldInfo"] = o.SaasFieldInfo
	return toSerialize, nil
}

type NullableSaasAttributeMapping struct {
	value *SaasAttributeMapping
	isSet bool
}

func (v NullableSaasAttributeMapping) Get() *SaasAttributeMapping {
	return v.value
}

func (v *NullableSaasAttributeMapping) Set(val *SaasAttributeMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableSaasAttributeMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableSaasAttributeMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaasAttributeMapping(val *SaasAttributeMapping) *NullableSaasAttributeMapping {
	return &NullableSaasAttributeMapping{value: val, isSet: true}
}

func (v NullableSaasAttributeMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaasAttributeMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
