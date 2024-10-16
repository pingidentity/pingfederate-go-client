/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"time"
)

// checks if the ConfigurationEncryptionKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationEncryptionKey{}

// ConfigurationEncryptionKey Configuration Encryption Key.
type ConfigurationEncryptionKey struct {
	// The id of the key.
	KeyId *string `json:"keyId,omitempty" tfsdk:"key_id"`
	// The creation date of the key.
	CreationDate *time.Time `json:"creationDate,omitempty" tfsdk:"creation_date"`
}

// NewConfigurationEncryptionKey instantiates a new ConfigurationEncryptionKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationEncryptionKey() *ConfigurationEncryptionKey {
	this := ConfigurationEncryptionKey{}
	return &this
}

// NewConfigurationEncryptionKeyWithDefaults instantiates a new ConfigurationEncryptionKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationEncryptionKeyWithDefaults() *ConfigurationEncryptionKey {
	this := ConfigurationEncryptionKey{}
	return &this
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *ConfigurationEncryptionKey) GetKeyId() string {
	if o == nil || IsNil(o.KeyId) {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationEncryptionKey) GetKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *ConfigurationEncryptionKey) HasKeyId() bool {
	if o != nil && !IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *ConfigurationEncryptionKey) SetKeyId(v string) {
	o.KeyId = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *ConfigurationEncryptionKey) GetCreationDate() time.Time {
	if o == nil || IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationEncryptionKey) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *ConfigurationEncryptionKey) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *ConfigurationEncryptionKey) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

func (o ConfigurationEncryptionKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationEncryptionKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeyId) {
		toSerialize["keyId"] = o.KeyId
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	return toSerialize, nil
}

type NullableConfigurationEncryptionKey struct {
	value *ConfigurationEncryptionKey
	isSet bool
}

func (v NullableConfigurationEncryptionKey) Get() *ConfigurationEncryptionKey {
	return v.value
}

func (v *NullableConfigurationEncryptionKey) Set(val *ConfigurationEncryptionKey) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationEncryptionKey) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationEncryptionKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationEncryptionKey(val *ConfigurationEncryptionKey) *NullableConfigurationEncryptionKey {
	return &NullableConfigurationEncryptionKey{value: val, isSet: true}
}

func (v NullableConfigurationEncryptionKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationEncryptionKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
