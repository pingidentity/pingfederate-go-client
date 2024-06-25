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

// checks if the LdapDataStoreAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapDataStoreAttribute{}

// LdapDataStoreAttribute struct for LdapDataStoreAttribute
type LdapDataStoreAttribute struct {
	DataStoreAttribute
}

// NewLdapDataStoreAttribute instantiates a new LdapDataStoreAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapDataStoreAttribute(type_ string, name string) *LdapDataStoreAttribute {
	this := LdapDataStoreAttribute{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewLdapDataStoreAttributeWithDefaults instantiates a new LdapDataStoreAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapDataStoreAttributeWithDefaults() *LdapDataStoreAttribute {
	this := LdapDataStoreAttribute{}
	return &this
}

func (o LdapDataStoreAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapDataStoreAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDataStoreAttribute, errDataStoreAttribute := json.Marshal(o.DataStoreAttribute)
	if errDataStoreAttribute != nil {
		return map[string]interface{}{}, errDataStoreAttribute
	}
	errDataStoreAttribute = json.Unmarshal([]byte(serializedDataStoreAttribute), &toSerialize)
	if errDataStoreAttribute != nil {
		return map[string]interface{}{}, errDataStoreAttribute
	}
	return toSerialize, nil
}

type NullableLdapDataStoreAttribute struct {
	value *LdapDataStoreAttribute
	isSet bool
}

func (v NullableLdapDataStoreAttribute) Get() *LdapDataStoreAttribute {
	return v.value
}

func (v *NullableLdapDataStoreAttribute) Set(val *LdapDataStoreAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapDataStoreAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapDataStoreAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapDataStoreAttribute(val *LdapDataStoreAttribute) *NullableLdapDataStoreAttribute {
	return &NullableLdapDataStoreAttribute{value: val, isSet: true}
}

func (v NullableLdapDataStoreAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapDataStoreAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
