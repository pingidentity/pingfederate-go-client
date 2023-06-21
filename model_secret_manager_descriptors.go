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

// checks if the SecretManagerDescriptors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretManagerDescriptors{}

// SecretManagerDescriptors A collection of secret manager plugin descriptors.
type SecretManagerDescriptors struct {
	// The list of secret manager plugin descriptors.
	Items []SecretManagerDescriptor `json:"items,omitempty"`
}

// NewSecretManagerDescriptors instantiates a new SecretManagerDescriptors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretManagerDescriptors() *SecretManagerDescriptors {
	this := SecretManagerDescriptors{}
	return &this
}

// NewSecretManagerDescriptorsWithDefaults instantiates a new SecretManagerDescriptors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretManagerDescriptorsWithDefaults() *SecretManagerDescriptors {
	this := SecretManagerDescriptors{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SecretManagerDescriptors) GetItems() []SecretManagerDescriptor {
	if o == nil || IsNil(o.Items) {
		var ret []SecretManagerDescriptor
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretManagerDescriptors) GetItemsOk() ([]SecretManagerDescriptor, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SecretManagerDescriptors) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SecretManagerDescriptor and assigns it to the Items field.
func (o *SecretManagerDescriptors) SetItems(v []SecretManagerDescriptor) {
	o.Items = v
}

func (o SecretManagerDescriptors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretManagerDescriptors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableSecretManagerDescriptors struct {
	value *SecretManagerDescriptors
	isSet bool
}

func (v NullableSecretManagerDescriptors) Get() *SecretManagerDescriptors {
	return v.value
}

func (v *NullableSecretManagerDescriptors) Set(val *SecretManagerDescriptors) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretManagerDescriptors) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretManagerDescriptors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretManagerDescriptors(val *SecretManagerDescriptors) *NullableSecretManagerDescriptors {
	return &NullableSecretManagerDescriptors{value: val, isSet: true}
}

func (v NullableSecretManagerDescriptors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretManagerDescriptors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
