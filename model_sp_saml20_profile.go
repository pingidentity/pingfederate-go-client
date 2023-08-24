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

// checks if the SpSAML20Profile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpSAML20Profile{}

// SpSAML20Profile struct for SpSAML20Profile
type SpSAML20Profile struct {
	// Enable SAML2.0 profile.
	Enable *bool `json:"enable,omitempty" tfsdk:"enable"`
	// This property has been deprecated and no longer used
	EnableAutoConnect *bool `json:"enableAutoConnect,omitempty" tfsdk:"enable_auto_connect"`
	// Enable Attribute Requester Mapping for X.509 Attribute Sharing Profile (XASP)
	EnableXASP *bool `json:"enableXASP,omitempty" tfsdk:"enable_xasp"`
}

// NewSpSAML20Profile instantiates a new SpSAML20Profile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpSAML20Profile() *SpSAML20Profile {
	this := SpSAML20Profile{}
	return &this
}

// NewSpSAML20ProfileWithDefaults instantiates a new SpSAML20Profile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpSAML20ProfileWithDefaults() *SpSAML20Profile {
	this := SpSAML20Profile{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *SpSAML20Profile) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpSAML20Profile) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *SpSAML20Profile) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *SpSAML20Profile) SetEnable(v bool) {
	o.Enable = &v
}

// GetEnableAutoConnect returns the EnableAutoConnect field value if set, zero value otherwise.
func (o *SpSAML20Profile) GetEnableAutoConnect() bool {
	if o == nil || IsNil(o.EnableAutoConnect) {
		var ret bool
		return ret
	}
	return *o.EnableAutoConnect
}

// GetEnableAutoConnectOk returns a tuple with the EnableAutoConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpSAML20Profile) GetEnableAutoConnectOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAutoConnect) {
		return nil, false
	}
	return o.EnableAutoConnect, true
}

// HasEnableAutoConnect returns a boolean if a field has been set.
func (o *SpSAML20Profile) HasEnableAutoConnect() bool {
	if o != nil && !IsNil(o.EnableAutoConnect) {
		return true
	}

	return false
}

// SetEnableAutoConnect gets a reference to the given bool and assigns it to the EnableAutoConnect field.
func (o *SpSAML20Profile) SetEnableAutoConnect(v bool) {
	o.EnableAutoConnect = &v
}

// GetEnableXASP returns the EnableXASP field value if set, zero value otherwise.
func (o *SpSAML20Profile) GetEnableXASP() bool {
	if o == nil || IsNil(o.EnableXASP) {
		var ret bool
		return ret
	}
	return *o.EnableXASP
}

// GetEnableXASPOk returns a tuple with the EnableXASP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpSAML20Profile) GetEnableXASPOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableXASP) {
		return nil, false
	}
	return o.EnableXASP, true
}

// HasEnableXASP returns a boolean if a field has been set.
func (o *SpSAML20Profile) HasEnableXASP() bool {
	if o != nil && !IsNil(o.EnableXASP) {
		return true
	}

	return false
}

// SetEnableXASP gets a reference to the given bool and assigns it to the EnableXASP field.
func (o *SpSAML20Profile) SetEnableXASP(v bool) {
	o.EnableXASP = &v
}

func (o SpSAML20Profile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpSAML20Profile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.EnableAutoConnect) {
		toSerialize["enableAutoConnect"] = o.EnableAutoConnect
	}
	if !IsNil(o.EnableXASP) {
		toSerialize["enableXASP"] = o.EnableXASP
	}
	return toSerialize, nil
}

type NullableSpSAML20Profile struct {
	value *SpSAML20Profile
	isSet bool
}

func (v NullableSpSAML20Profile) Get() *SpSAML20Profile {
	return v.value
}

func (v *NullableSpSAML20Profile) Set(val *SpSAML20Profile) {
	v.value = val
	v.isSet = true
}

func (v NullableSpSAML20Profile) IsSet() bool {
	return v.isSet
}

func (v *NullableSpSAML20Profile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpSAML20Profile(val *SpSAML20Profile) *NullableSpSAML20Profile {
	return &NullableSpSAML20Profile{value: val, isSet: true}
}

func (v NullableSpSAML20Profile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpSAML20Profile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
