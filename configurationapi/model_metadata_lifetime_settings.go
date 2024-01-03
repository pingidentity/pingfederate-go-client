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

// checks if the MetadataLifetimeSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataLifetimeSettings{}

// MetadataLifetimeSettings Metadata lifetime settings.
type MetadataLifetimeSettings struct {
	// This field adjusts the validity of your metadata in minutes. The default value is 1440 (1 day).
	CacheDuration *int64 `json:"cacheDuration,omitempty" tfsdk:"cache_duration"`
	// This field adjusts the frequency of automatic reloading of SAML metadata in minutes. The default value is 1440 (1 day).
	ReloadDelay *int64 `json:"reloadDelay,omitempty" tfsdk:"reload_delay"`
}

// NewMetadataLifetimeSettings instantiates a new MetadataLifetimeSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataLifetimeSettings() *MetadataLifetimeSettings {
	this := MetadataLifetimeSettings{}
	return &this
}

// NewMetadataLifetimeSettingsWithDefaults instantiates a new MetadataLifetimeSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataLifetimeSettingsWithDefaults() *MetadataLifetimeSettings {
	this := MetadataLifetimeSettings{}
	return &this
}

// GetCacheDuration returns the CacheDuration field value if set, zero value otherwise.
func (o *MetadataLifetimeSettings) GetCacheDuration() int64 {
	if o == nil || IsNil(o.CacheDuration) {
		var ret int64
		return ret
	}
	return *o.CacheDuration
}

// GetCacheDurationOk returns a tuple with the CacheDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataLifetimeSettings) GetCacheDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.CacheDuration) {
		return nil, false
	}
	return o.CacheDuration, true
}

// HasCacheDuration returns a boolean if a field has been set.
func (o *MetadataLifetimeSettings) HasCacheDuration() bool {
	if o != nil && !IsNil(o.CacheDuration) {
		return true
	}

	return false
}

// SetCacheDuration gets a reference to the given int64 and assigns it to the CacheDuration field.
func (o *MetadataLifetimeSettings) SetCacheDuration(v int64) {
	o.CacheDuration = &v
}

// GetReloadDelay returns the ReloadDelay field value if set, zero value otherwise.
func (o *MetadataLifetimeSettings) GetReloadDelay() int64 {
	if o == nil || IsNil(o.ReloadDelay) {
		var ret int64
		return ret
	}
	return *o.ReloadDelay
}

// GetReloadDelayOk returns a tuple with the ReloadDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataLifetimeSettings) GetReloadDelayOk() (*int64, bool) {
	if o == nil || IsNil(o.ReloadDelay) {
		return nil, false
	}
	return o.ReloadDelay, true
}

// HasReloadDelay returns a boolean if a field has been set.
func (o *MetadataLifetimeSettings) HasReloadDelay() bool {
	if o != nil && !IsNil(o.ReloadDelay) {
		return true
	}

	return false
}

// SetReloadDelay gets a reference to the given int64 and assigns it to the ReloadDelay field.
func (o *MetadataLifetimeSettings) SetReloadDelay(v int64) {
	o.ReloadDelay = &v
}

func (o MetadataLifetimeSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataLifetimeSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CacheDuration) {
		toSerialize["cacheDuration"] = o.CacheDuration
	}
	if !IsNil(o.ReloadDelay) {
		toSerialize["reloadDelay"] = o.ReloadDelay
	}
	return toSerialize, nil
}

type NullableMetadataLifetimeSettings struct {
	value *MetadataLifetimeSettings
	isSet bool
}

func (v NullableMetadataLifetimeSettings) Get() *MetadataLifetimeSettings {
	return v.value
}

func (v *NullableMetadataLifetimeSettings) Set(val *MetadataLifetimeSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataLifetimeSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataLifetimeSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataLifetimeSettings(val *MetadataLifetimeSettings) *NullableMetadataLifetimeSettings {
	return &NullableMetadataLifetimeSettings{value: val, isSet: true}
}

func (v NullableMetadataLifetimeSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataLifetimeSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
