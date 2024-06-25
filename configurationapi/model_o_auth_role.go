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

// checks if the OAuthRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthRole{}

// OAuthRole This property has been deprecated and is no longer used. OAuth and OpenID Connect are always enabled.
type OAuthRole struct {
	// Enable OAuth 2.0 Authorization Server (AS) Role.
	EnableOauth *bool `json:"enableOauth,omitempty" tfsdk:"enable_oauth"`
	// Enable Open ID Connect.
	EnableOpenIdConnect *bool `json:"enableOpenIdConnect,omitempty" tfsdk:"enable_open_id_connect"`
}

// NewOAuthRole instantiates a new OAuthRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthRole() *OAuthRole {
	this := OAuthRole{}
	return &this
}

// NewOAuthRoleWithDefaults instantiates a new OAuthRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthRoleWithDefaults() *OAuthRole {
	this := OAuthRole{}
	return &this
}

// GetEnableOauth returns the EnableOauth field value if set, zero value otherwise.
func (o *OAuthRole) GetEnableOauth() bool {
	if o == nil || IsNil(o.EnableOauth) {
		var ret bool
		return ret
	}
	return *o.EnableOauth
}

// GetEnableOauthOk returns a tuple with the EnableOauth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthRole) GetEnableOauthOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableOauth) {
		return nil, false
	}
	return o.EnableOauth, true
}

// HasEnableOauth returns a boolean if a field has been set.
func (o *OAuthRole) HasEnableOauth() bool {
	if o != nil && !IsNil(o.EnableOauth) {
		return true
	}

	return false
}

// SetEnableOauth gets a reference to the given bool and assigns it to the EnableOauth field.
func (o *OAuthRole) SetEnableOauth(v bool) {
	o.EnableOauth = &v
}

// GetEnableOpenIdConnect returns the EnableOpenIdConnect field value if set, zero value otherwise.
func (o *OAuthRole) GetEnableOpenIdConnect() bool {
	if o == nil || IsNil(o.EnableOpenIdConnect) {
		var ret bool
		return ret
	}
	return *o.EnableOpenIdConnect
}

// GetEnableOpenIdConnectOk returns a tuple with the EnableOpenIdConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthRole) GetEnableOpenIdConnectOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableOpenIdConnect) {
		return nil, false
	}
	return o.EnableOpenIdConnect, true
}

// HasEnableOpenIdConnect returns a boolean if a field has been set.
func (o *OAuthRole) HasEnableOpenIdConnect() bool {
	if o != nil && !IsNil(o.EnableOpenIdConnect) {
		return true
	}

	return false
}

// SetEnableOpenIdConnect gets a reference to the given bool and assigns it to the EnableOpenIdConnect field.
func (o *OAuthRole) SetEnableOpenIdConnect(v bool) {
	o.EnableOpenIdConnect = &v
}

func (o OAuthRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableOauth) {
		toSerialize["enableOauth"] = o.EnableOauth
	}
	if !IsNil(o.EnableOpenIdConnect) {
		toSerialize["enableOpenIdConnect"] = o.EnableOpenIdConnect
	}
	return toSerialize, nil
}

type NullableOAuthRole struct {
	value *OAuthRole
	isSet bool
}

func (v NullableOAuthRole) Get() *OAuthRole {
	return v.value
}

func (v *NullableOAuthRole) Set(val *OAuthRole) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthRole) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthRole(val *OAuthRole) *NullableOAuthRole {
	return &NullableOAuthRole{value: val, isSet: true}
}

func (v NullableOAuthRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
