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

// checks if the VirtualHostNameSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualHostNameSettings{}

// VirtualHostNameSettings Settings for virtual host names.
type VirtualHostNameSettings struct {
	// List of virtual host names.
	VirtualHostNames []string `json:"virtualHostNames,omitempty" tfsdk:"virtual_host_names"`
}

// NewVirtualHostNameSettings instantiates a new VirtualHostNameSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualHostNameSettings() *VirtualHostNameSettings {
	this := VirtualHostNameSettings{}
	return &this
}

// NewVirtualHostNameSettingsWithDefaults instantiates a new VirtualHostNameSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualHostNameSettingsWithDefaults() *VirtualHostNameSettings {
	this := VirtualHostNameSettings{}
	return &this
}

// GetVirtualHostNames returns the VirtualHostNames field value if set, zero value otherwise.
func (o *VirtualHostNameSettings) GetVirtualHostNames() []string {
	if o == nil || IsNil(o.VirtualHostNames) {
		var ret []string
		return ret
	}
	return o.VirtualHostNames
}

// GetVirtualHostNamesOk returns a tuple with the VirtualHostNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualHostNameSettings) GetVirtualHostNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.VirtualHostNames) {
		return nil, false
	}
	return o.VirtualHostNames, true
}

// HasVirtualHostNames returns a boolean if a field has been set.
func (o *VirtualHostNameSettings) HasVirtualHostNames() bool {
	if o != nil && !IsNil(o.VirtualHostNames) {
		return true
	}

	return false
}

// SetVirtualHostNames gets a reference to the given []string and assigns it to the VirtualHostNames field.
func (o *VirtualHostNameSettings) SetVirtualHostNames(v []string) {
	o.VirtualHostNames = v
}

func (o VirtualHostNameSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualHostNameSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VirtualHostNames) {
		toSerialize["virtualHostNames"] = o.VirtualHostNames
	}
	return toSerialize, nil
}

type NullableVirtualHostNameSettings struct {
	value *VirtualHostNameSettings
	isSet bool
}

func (v NullableVirtualHostNameSettings) Get() *VirtualHostNameSettings {
	return v.value
}

func (v *NullableVirtualHostNameSettings) Set(val *VirtualHostNameSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualHostNameSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualHostNameSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualHostNameSettings(val *VirtualHostNameSettings) *NullableVirtualHostNameSettings {
	return &NullableVirtualHostNameSettings{value: val, isSet: true}
}

func (v NullableVirtualHostNameSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualHostNameSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
