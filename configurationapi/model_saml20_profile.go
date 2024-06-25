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

// checks if the SAML20Profile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAML20Profile{}

// SAML20Profile SAML 2.0 Profile.
type SAML20Profile struct {
	// Enable SAML2.0 profile.
	Enable *bool `json:"enable,omitempty" tfsdk:"enable"`
	// This property has been deprecated and no longer used
	EnableAutoConnect *bool `json:"enableAutoConnect,omitempty" tfsdk:"enable_auto_connect"`
}

// NewSAML20Profile instantiates a new SAML20Profile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAML20Profile() *SAML20Profile {
	this := SAML20Profile{}
	return &this
}

// NewSAML20ProfileWithDefaults instantiates a new SAML20Profile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAML20ProfileWithDefaults() *SAML20Profile {
	this := SAML20Profile{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *SAML20Profile) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAML20Profile) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *SAML20Profile) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *SAML20Profile) SetEnable(v bool) {
	o.Enable = &v
}

// GetEnableAutoConnect returns the EnableAutoConnect field value if set, zero value otherwise.
func (o *SAML20Profile) GetEnableAutoConnect() bool {
	if o == nil || IsNil(o.EnableAutoConnect) {
		var ret bool
		return ret
	}
	return *o.EnableAutoConnect
}

// GetEnableAutoConnectOk returns a tuple with the EnableAutoConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAML20Profile) GetEnableAutoConnectOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAutoConnect) {
		return nil, false
	}
	return o.EnableAutoConnect, true
}

// HasEnableAutoConnect returns a boolean if a field has been set.
func (o *SAML20Profile) HasEnableAutoConnect() bool {
	if o != nil && !IsNil(o.EnableAutoConnect) {
		return true
	}

	return false
}

// SetEnableAutoConnect gets a reference to the given bool and assigns it to the EnableAutoConnect field.
func (o *SAML20Profile) SetEnableAutoConnect(v bool) {
	o.EnableAutoConnect = &v
}

func (o SAML20Profile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAML20Profile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.EnableAutoConnect) {
		toSerialize["enableAutoConnect"] = o.EnableAutoConnect
	}
	return toSerialize, nil
}

type NullableSAML20Profile struct {
	value *SAML20Profile
	isSet bool
}

func (v NullableSAML20Profile) Get() *SAML20Profile {
	return v.value
}

func (v *NullableSAML20Profile) Set(val *SAML20Profile) {
	v.value = val
	v.isSet = true
}

func (v NullableSAML20Profile) IsSet() bool {
	return v.isSet
}

func (v *NullableSAML20Profile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAML20Profile(val *SAML20Profile) *NullableSAML20Profile {
	return &NullableSAML20Profile{value: val, isSet: true}
}

func (v NullableSAML20Profile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAML20Profile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
