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

// checks if the AtmAccessControlSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AtmAccessControlSettings{}

// AtmAccessControlSettings Access control settings for an access token management plugin instance.
type AtmAccessControlSettings struct {
	// Determines whether access to this token manager is restricted to specific OAuth clients. If false, the 'allowedClients' field is ignored. The default value is false.
	RestrictClients *bool `json:"restrictClients,omitempty" tfsdk:"restrict_clients"`
	// If 'restrictClients' is true, this field defines the list of OAuth clients that are allowed to access the token manager.
	AllowedClients []ResourceLink `json:"allowedClients,omitempty" tfsdk:"allowed_clients"`
}

// NewAtmAccessControlSettings instantiates a new AtmAccessControlSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtmAccessControlSettings() *AtmAccessControlSettings {
	this := AtmAccessControlSettings{}
	return &this
}

// NewAtmAccessControlSettingsWithDefaults instantiates a new AtmAccessControlSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtmAccessControlSettingsWithDefaults() *AtmAccessControlSettings {
	this := AtmAccessControlSettings{}
	return &this
}

// GetRestrictClients returns the RestrictClients field value if set, zero value otherwise.
func (o *AtmAccessControlSettings) GetRestrictClients() bool {
	if o == nil || IsNil(o.RestrictClients) {
		var ret bool
		return ret
	}
	return *o.RestrictClients
}

// GetRestrictClientsOk returns a tuple with the RestrictClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtmAccessControlSettings) GetRestrictClientsOk() (*bool, bool) {
	if o == nil || IsNil(o.RestrictClients) {
		return nil, false
	}
	return o.RestrictClients, true
}

// HasRestrictClients returns a boolean if a field has been set.
func (o *AtmAccessControlSettings) HasRestrictClients() bool {
	if o != nil && !IsNil(o.RestrictClients) {
		return true
	}

	return false
}

// SetRestrictClients gets a reference to the given bool and assigns it to the RestrictClients field.
func (o *AtmAccessControlSettings) SetRestrictClients(v bool) {
	o.RestrictClients = &v
}

// GetAllowedClients returns the AllowedClients field value if set, zero value otherwise.
func (o *AtmAccessControlSettings) GetAllowedClients() []ResourceLink {
	if o == nil || IsNil(o.AllowedClients) {
		var ret []ResourceLink
		return ret
	}
	return o.AllowedClients
}

// GetAllowedClientsOk returns a tuple with the AllowedClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtmAccessControlSettings) GetAllowedClientsOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.AllowedClients) {
		return nil, false
	}
	return o.AllowedClients, true
}

// HasAllowedClients returns a boolean if a field has been set.
func (o *AtmAccessControlSettings) HasAllowedClients() bool {
	if o != nil && !IsNil(o.AllowedClients) {
		return true
	}

	return false
}

// SetAllowedClients gets a reference to the given []ResourceLink and assigns it to the AllowedClients field.
func (o *AtmAccessControlSettings) SetAllowedClients(v []ResourceLink) {
	o.AllowedClients = v
}

func (o AtmAccessControlSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AtmAccessControlSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RestrictClients) {
		toSerialize["restrictClients"] = o.RestrictClients
	}
	if !IsNil(o.AllowedClients) {
		toSerialize["allowedClients"] = o.AllowedClients
	}
	return toSerialize, nil
}

type NullableAtmAccessControlSettings struct {
	value *AtmAccessControlSettings
	isSet bool
}

func (v NullableAtmAccessControlSettings) Get() *AtmAccessControlSettings {
	return v.value
}

func (v *NullableAtmAccessControlSettings) Set(val *AtmAccessControlSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAtmAccessControlSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAtmAccessControlSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtmAccessControlSettings(val *AtmAccessControlSettings) *NullableAtmAccessControlSettings {
	return &NullableAtmAccessControlSettings{value: val, isSet: true}
}

func (v NullableAtmAccessControlSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtmAccessControlSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
