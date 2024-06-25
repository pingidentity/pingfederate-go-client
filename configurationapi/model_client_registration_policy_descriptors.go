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

// checks if the ClientRegistrationPolicyDescriptors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientRegistrationPolicyDescriptors{}

// ClientRegistrationPolicyDescriptors A collection of client registration policy plugin descriptors.
type ClientRegistrationPolicyDescriptors struct {
	// The list of client registration policy plugin descriptors.
	Items []ClientRegistrationPolicyDescriptor `json:"items,omitempty" tfsdk:"items"`
}

// NewClientRegistrationPolicyDescriptors instantiates a new ClientRegistrationPolicyDescriptors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientRegistrationPolicyDescriptors() *ClientRegistrationPolicyDescriptors {
	this := ClientRegistrationPolicyDescriptors{}
	return &this
}

// NewClientRegistrationPolicyDescriptorsWithDefaults instantiates a new ClientRegistrationPolicyDescriptors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientRegistrationPolicyDescriptorsWithDefaults() *ClientRegistrationPolicyDescriptors {
	this := ClientRegistrationPolicyDescriptors{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ClientRegistrationPolicyDescriptors) GetItems() []ClientRegistrationPolicyDescriptor {
	if o == nil || IsNil(o.Items) {
		var ret []ClientRegistrationPolicyDescriptor
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationPolicyDescriptors) GetItemsOk() ([]ClientRegistrationPolicyDescriptor, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ClientRegistrationPolicyDescriptors) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ClientRegistrationPolicyDescriptor and assigns it to the Items field.
func (o *ClientRegistrationPolicyDescriptors) SetItems(v []ClientRegistrationPolicyDescriptor) {
	o.Items = v
}

func (o ClientRegistrationPolicyDescriptors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientRegistrationPolicyDescriptors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableClientRegistrationPolicyDescriptors struct {
	value *ClientRegistrationPolicyDescriptors
	isSet bool
}

func (v NullableClientRegistrationPolicyDescriptors) Get() *ClientRegistrationPolicyDescriptors {
	return v.value
}

func (v *NullableClientRegistrationPolicyDescriptors) Set(val *ClientRegistrationPolicyDescriptors) {
	v.value = val
	v.isSet = true
}

func (v NullableClientRegistrationPolicyDescriptors) IsSet() bool {
	return v.isSet
}

func (v *NullableClientRegistrationPolicyDescriptors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientRegistrationPolicyDescriptors(val *ClientRegistrationPolicyDescriptors) *NullableClientRegistrationPolicyDescriptors {
	return &NullableClientRegistrationPolicyDescriptors{value: val, isSet: true}
}

func (v NullableClientRegistrationPolicyDescriptors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientRegistrationPolicyDescriptors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
