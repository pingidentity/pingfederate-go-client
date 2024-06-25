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

// checks if the BaseProviderRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseProviderRole{}

// BaseProviderRole Base Provider Role.
type BaseProviderRole struct {
	Enable *bool `json:"enable,omitempty" tfsdk:"enable"`
	// Enable SAML 1.1.
	EnableSaml11 *bool `json:"enableSaml11,omitempty" tfsdk:"enable_saml11"`
	// Enable SAML 1.0.
	EnableSaml10 *bool `json:"enableSaml10,omitempty" tfsdk:"enable_saml10"`
	// Enable WS Federation.
	EnableWsFed *bool `json:"enableWsFed,omitempty" tfsdk:"enable_ws_fed"`
	// Enable WS Trust.
	EnableWsTrust *bool `json:"enableWsTrust,omitempty" tfsdk:"enable_ws_trust"`
}

// NewBaseProviderRole instantiates a new BaseProviderRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseProviderRole() *BaseProviderRole {
	this := BaseProviderRole{}
	return &this
}

// NewBaseProviderRoleWithDefaults instantiates a new BaseProviderRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseProviderRoleWithDefaults() *BaseProviderRole {
	this := BaseProviderRole{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *BaseProviderRole) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseProviderRole) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *BaseProviderRole) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *BaseProviderRole) SetEnable(v bool) {
	o.Enable = &v
}

// GetEnableSaml11 returns the EnableSaml11 field value if set, zero value otherwise.
func (o *BaseProviderRole) GetEnableSaml11() bool {
	if o == nil || IsNil(o.EnableSaml11) {
		var ret bool
		return ret
	}
	return *o.EnableSaml11
}

// GetEnableSaml11Ok returns a tuple with the EnableSaml11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseProviderRole) GetEnableSaml11Ok() (*bool, bool) {
	if o == nil || IsNil(o.EnableSaml11) {
		return nil, false
	}
	return o.EnableSaml11, true
}

// HasEnableSaml11 returns a boolean if a field has been set.
func (o *BaseProviderRole) HasEnableSaml11() bool {
	if o != nil && !IsNil(o.EnableSaml11) {
		return true
	}

	return false
}

// SetEnableSaml11 gets a reference to the given bool and assigns it to the EnableSaml11 field.
func (o *BaseProviderRole) SetEnableSaml11(v bool) {
	o.EnableSaml11 = &v
}

// GetEnableSaml10 returns the EnableSaml10 field value if set, zero value otherwise.
func (o *BaseProviderRole) GetEnableSaml10() bool {
	if o == nil || IsNil(o.EnableSaml10) {
		var ret bool
		return ret
	}
	return *o.EnableSaml10
}

// GetEnableSaml10Ok returns a tuple with the EnableSaml10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseProviderRole) GetEnableSaml10Ok() (*bool, bool) {
	if o == nil || IsNil(o.EnableSaml10) {
		return nil, false
	}
	return o.EnableSaml10, true
}

// HasEnableSaml10 returns a boolean if a field has been set.
func (o *BaseProviderRole) HasEnableSaml10() bool {
	if o != nil && !IsNil(o.EnableSaml10) {
		return true
	}

	return false
}

// SetEnableSaml10 gets a reference to the given bool and assigns it to the EnableSaml10 field.
func (o *BaseProviderRole) SetEnableSaml10(v bool) {
	o.EnableSaml10 = &v
}

// GetEnableWsFed returns the EnableWsFed field value if set, zero value otherwise.
func (o *BaseProviderRole) GetEnableWsFed() bool {
	if o == nil || IsNil(o.EnableWsFed) {
		var ret bool
		return ret
	}
	return *o.EnableWsFed
}

// GetEnableWsFedOk returns a tuple with the EnableWsFed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseProviderRole) GetEnableWsFedOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableWsFed) {
		return nil, false
	}
	return o.EnableWsFed, true
}

// HasEnableWsFed returns a boolean if a field has been set.
func (o *BaseProviderRole) HasEnableWsFed() bool {
	if o != nil && !IsNil(o.EnableWsFed) {
		return true
	}

	return false
}

// SetEnableWsFed gets a reference to the given bool and assigns it to the EnableWsFed field.
func (o *BaseProviderRole) SetEnableWsFed(v bool) {
	o.EnableWsFed = &v
}

// GetEnableWsTrust returns the EnableWsTrust field value if set, zero value otherwise.
func (o *BaseProviderRole) GetEnableWsTrust() bool {
	if o == nil || IsNil(o.EnableWsTrust) {
		var ret bool
		return ret
	}
	return *o.EnableWsTrust
}

// GetEnableWsTrustOk returns a tuple with the EnableWsTrust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseProviderRole) GetEnableWsTrustOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableWsTrust) {
		return nil, false
	}
	return o.EnableWsTrust, true
}

// HasEnableWsTrust returns a boolean if a field has been set.
func (o *BaseProviderRole) HasEnableWsTrust() bool {
	if o != nil && !IsNil(o.EnableWsTrust) {
		return true
	}

	return false
}

// SetEnableWsTrust gets a reference to the given bool and assigns it to the EnableWsTrust field.
func (o *BaseProviderRole) SetEnableWsTrust(v bool) {
	o.EnableWsTrust = &v
}

func (o BaseProviderRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseProviderRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.EnableSaml11) {
		toSerialize["enableSaml11"] = o.EnableSaml11
	}
	if !IsNil(o.EnableSaml10) {
		toSerialize["enableSaml10"] = o.EnableSaml10
	}
	if !IsNil(o.EnableWsFed) {
		toSerialize["enableWsFed"] = o.EnableWsFed
	}
	if !IsNil(o.EnableWsTrust) {
		toSerialize["enableWsTrust"] = o.EnableWsTrust
	}
	return toSerialize, nil
}

type NullableBaseProviderRole struct {
	value *BaseProviderRole
	isSet bool
}

func (v NullableBaseProviderRole) Get() *BaseProviderRole {
	return v.value
}

func (v *NullableBaseProviderRole) Set(val *BaseProviderRole) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseProviderRole) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseProviderRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseProviderRole(val *BaseProviderRole) *NullableBaseProviderRole {
	return &NullableBaseProviderRole{value: val, isSet: true}
}

func (v NullableBaseProviderRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseProviderRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
