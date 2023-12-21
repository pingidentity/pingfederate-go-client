/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"time"
)

// checks if the SystemKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemKey{}

// SystemKey A system key.
type SystemKey struct {
	// Creation time of the key.
	CreationDate *time.Time `json:"creationDate,omitempty" tfsdk:"creation_date"`
	// The system key encrypted.
	EncryptedKeyData *string `json:"encryptedKeyData,omitempty" tfsdk:"encrypted_key_data"`
	// The clear text system key base 64 encoded. The system key must be 32 bytes before base 64 encoding.
	KeyData *string `json:"keyData,omitempty" tfsdk:"key_data"`
}

// NewSystemKey instantiates a new SystemKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemKey() *SystemKey {
	this := SystemKey{}
	return &this
}

// NewSystemKeyWithDefaults instantiates a new SystemKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemKeyWithDefaults() *SystemKey {
	this := SystemKey{}
	return &this
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *SystemKey) GetCreationDate() time.Time {
	if o == nil || IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemKey) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *SystemKey) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *SystemKey) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetEncryptedKeyData returns the EncryptedKeyData field value if set, zero value otherwise.
func (o *SystemKey) GetEncryptedKeyData() string {
	if o == nil || IsNil(o.EncryptedKeyData) {
		var ret string
		return ret
	}
	return *o.EncryptedKeyData
}

// GetEncryptedKeyDataOk returns a tuple with the EncryptedKeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemKey) GetEncryptedKeyDataOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedKeyData) {
		return nil, false
	}
	return o.EncryptedKeyData, true
}

// HasEncryptedKeyData returns a boolean if a field has been set.
func (o *SystemKey) HasEncryptedKeyData() bool {
	if o != nil && !IsNil(o.EncryptedKeyData) {
		return true
	}

	return false
}

// SetEncryptedKeyData gets a reference to the given string and assigns it to the EncryptedKeyData field.
func (o *SystemKey) SetEncryptedKeyData(v string) {
	o.EncryptedKeyData = &v
}

// GetKeyData returns the KeyData field value if set, zero value otherwise.
func (o *SystemKey) GetKeyData() string {
	if o == nil || IsNil(o.KeyData) {
		var ret string
		return ret
	}
	return *o.KeyData
}

// GetKeyDataOk returns a tuple with the KeyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemKey) GetKeyDataOk() (*string, bool) {
	if o == nil || IsNil(o.KeyData) {
		return nil, false
	}
	return o.KeyData, true
}

// HasKeyData returns a boolean if a field has been set.
func (o *SystemKey) HasKeyData() bool {
	if o != nil && !IsNil(o.KeyData) {
		return true
	}

	return false
}

// SetKeyData gets a reference to the given string and assigns it to the KeyData field.
func (o *SystemKey) SetKeyData(v string) {
	o.KeyData = &v
}

func (o SystemKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !IsNil(o.EncryptedKeyData) {
		toSerialize["encryptedKeyData"] = o.EncryptedKeyData
	}
	if !IsNil(o.KeyData) {
		toSerialize["keyData"] = o.KeyData
	}
	return toSerialize, nil
}

type NullableSystemKey struct {
	value *SystemKey
	isSet bool
}

func (v NullableSystemKey) Get() *SystemKey {
	return v.value
}

func (v *NullableSystemKey) Set(val *SystemKey) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemKey) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemKey(val *SystemKey) *NullableSystemKey {
	return &NullableSystemKey{value: val, isSet: true}
}

func (v NullableSystemKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
