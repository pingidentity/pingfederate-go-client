/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the ProxySettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxySettings{}

// ProxySettings Proxy settings.
type ProxySettings struct {
	// Host name.
	Host *string `json:"host,omitempty" tfsdk:"host"`
	// Port number.
	Port *int64 `json:"port,omitempty" tfsdk:"port"`
}

// NewProxySettings instantiates a new ProxySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxySettings() *ProxySettings {
	this := ProxySettings{}
	return &this
}

// NewProxySettingsWithDefaults instantiates a new ProxySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxySettingsWithDefaults() *ProxySettings {
	this := ProxySettings{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ProxySettings) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySettings) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ProxySettings) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ProxySettings) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ProxySettings) GetPort() int64 {
	if o == nil || IsNil(o.Port) {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxySettings) GetPortOk() (*int64, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ProxySettings) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *ProxySettings) SetPort(v int64) {
	o.Port = &v
}

func (o ProxySettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxySettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableProxySettings struct {
	value *ProxySettings
	isSet bool
}

func (v NullableProxySettings) Get() *ProxySettings {
	return v.value
}

func (v *NullableProxySettings) Set(val *ProxySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableProxySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableProxySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxySettings(val *ProxySettings) *NullableProxySettings {
	return &NullableProxySettings{value: val, isSet: true}
}

func (v NullableProxySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
