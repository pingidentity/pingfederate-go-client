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

// checks if the AuthenticationSelectorDescriptors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationSelectorDescriptors{}

// AuthenticationSelectorDescriptors A collection of Authentication Selector descriptors.
type AuthenticationSelectorDescriptors struct {
	// The list of Authentication Selector descriptors.
	Items []AuthenticationSelectorDescriptor `json:"items,omitempty" tfsdk:"items"`
}

// NewAuthenticationSelectorDescriptors instantiates a new AuthenticationSelectorDescriptors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationSelectorDescriptors() *AuthenticationSelectorDescriptors {
	this := AuthenticationSelectorDescriptors{}
	return &this
}

// NewAuthenticationSelectorDescriptorsWithDefaults instantiates a new AuthenticationSelectorDescriptors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationSelectorDescriptorsWithDefaults() *AuthenticationSelectorDescriptors {
	this := AuthenticationSelectorDescriptors{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AuthenticationSelectorDescriptors) GetItems() []AuthenticationSelectorDescriptor {
	if o == nil || IsNil(o.Items) {
		var ret []AuthenticationSelectorDescriptor
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationSelectorDescriptors) GetItemsOk() ([]AuthenticationSelectorDescriptor, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AuthenticationSelectorDescriptors) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AuthenticationSelectorDescriptor and assigns it to the Items field.
func (o *AuthenticationSelectorDescriptors) SetItems(v []AuthenticationSelectorDescriptor) {
	o.Items = v
}

func (o AuthenticationSelectorDescriptors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationSelectorDescriptors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableAuthenticationSelectorDescriptors struct {
	value *AuthenticationSelectorDescriptors
	isSet bool
}

func (v NullableAuthenticationSelectorDescriptors) Get() *AuthenticationSelectorDescriptors {
	return v.value
}

func (v *NullableAuthenticationSelectorDescriptors) Set(val *AuthenticationSelectorDescriptors) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationSelectorDescriptors) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationSelectorDescriptors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationSelectorDescriptors(val *AuthenticationSelectorDescriptors) *NullableAuthenticationSelectorDescriptors {
	return &NullableAuthenticationSelectorDescriptors{value: val, isSet: true}
}

func (v NullableAuthenticationSelectorDescriptors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationSelectorDescriptors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
