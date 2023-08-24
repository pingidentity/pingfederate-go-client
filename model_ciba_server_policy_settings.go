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

// checks if the CibaServerPolicySettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CibaServerPolicySettings{}

// CibaServerPolicySettings Settings for the CIBA request policy configuration.
type CibaServerPolicySettings struct {
	DefaultRequestPolicyRef *ResourceLink `json:"defaultRequestPolicyRef,omitempty" tfsdk:"default_request_policy_ref"`
}

// NewCibaServerPolicySettings instantiates a new CibaServerPolicySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCibaServerPolicySettings() *CibaServerPolicySettings {
	this := CibaServerPolicySettings{}
	return &this
}

// NewCibaServerPolicySettingsWithDefaults instantiates a new CibaServerPolicySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCibaServerPolicySettingsWithDefaults() *CibaServerPolicySettings {
	this := CibaServerPolicySettings{}
	return &this
}

// GetDefaultRequestPolicyRef returns the DefaultRequestPolicyRef field value if set, zero value otherwise.
func (o *CibaServerPolicySettings) GetDefaultRequestPolicyRef() ResourceLink {
	if o == nil || IsNil(o.DefaultRequestPolicyRef) {
		var ret ResourceLink
		return ret
	}
	return *o.DefaultRequestPolicyRef
}

// GetDefaultRequestPolicyRefOk returns a tuple with the DefaultRequestPolicyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CibaServerPolicySettings) GetDefaultRequestPolicyRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.DefaultRequestPolicyRef) {
		return nil, false
	}
	return o.DefaultRequestPolicyRef, true
}

// HasDefaultRequestPolicyRef returns a boolean if a field has been set.
func (o *CibaServerPolicySettings) HasDefaultRequestPolicyRef() bool {
	if o != nil && !IsNil(o.DefaultRequestPolicyRef) {
		return true
	}

	return false
}

// SetDefaultRequestPolicyRef gets a reference to the given ResourceLink and assigns it to the DefaultRequestPolicyRef field.
func (o *CibaServerPolicySettings) SetDefaultRequestPolicyRef(v ResourceLink) {
	o.DefaultRequestPolicyRef = &v
}

func (o CibaServerPolicySettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CibaServerPolicySettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultRequestPolicyRef) {
		toSerialize["defaultRequestPolicyRef"] = o.DefaultRequestPolicyRef
	}
	return toSerialize, nil
}

type NullableCibaServerPolicySettings struct {
	value *CibaServerPolicySettings
	isSet bool
}

func (v NullableCibaServerPolicySettings) Get() *CibaServerPolicySettings {
	return v.value
}

func (v *NullableCibaServerPolicySettings) Set(val *CibaServerPolicySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCibaServerPolicySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCibaServerPolicySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCibaServerPolicySettings(val *CibaServerPolicySettings) *NullableCibaServerPolicySettings {
	return &NullableCibaServerPolicySettings{value: val, isSet: true}
}

func (v NullableCibaServerPolicySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCibaServerPolicySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
