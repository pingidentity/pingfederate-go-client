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

// checks if the SaasPluginDescriptors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaasPluginDescriptors{}

// SaasPluginDescriptors A collection of SaaS plugins.
type SaasPluginDescriptors struct {
	// The actual list of SaaS plugins.
	Items []SaasPluginDescriptor `json:"items,omitempty" tfsdk:"items"`
}

// NewSaasPluginDescriptors instantiates a new SaasPluginDescriptors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaasPluginDescriptors() *SaasPluginDescriptors {
	this := SaasPluginDescriptors{}
	return &this
}

// NewSaasPluginDescriptorsWithDefaults instantiates a new SaasPluginDescriptors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaasPluginDescriptorsWithDefaults() *SaasPluginDescriptors {
	this := SaasPluginDescriptors{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SaasPluginDescriptors) GetItems() []SaasPluginDescriptor {
	if o == nil || IsNil(o.Items) {
		var ret []SaasPluginDescriptor
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasPluginDescriptors) GetItemsOk() ([]SaasPluginDescriptor, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SaasPluginDescriptors) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SaasPluginDescriptor and assigns it to the Items field.
func (o *SaasPluginDescriptors) SetItems(v []SaasPluginDescriptor) {
	o.Items = v
}

func (o SaasPluginDescriptors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaasPluginDescriptors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableSaasPluginDescriptors struct {
	value *SaasPluginDescriptors
	isSet bool
}

func (v NullableSaasPluginDescriptors) Get() *SaasPluginDescriptors {
	return v.value
}

func (v *NullableSaasPluginDescriptors) Set(val *SaasPluginDescriptors) {
	v.value = val
	v.isSet = true
}

func (v NullableSaasPluginDescriptors) IsSet() bool {
	return v.isSet
}

func (v *NullableSaasPluginDescriptors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaasPluginDescriptors(val *SaasPluginDescriptors) *NullableSaasPluginDescriptors {
	return &NullableSaasPluginDescriptors{value: val, isSet: true}
}

func (v NullableSaasPluginDescriptors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaasPluginDescriptors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
