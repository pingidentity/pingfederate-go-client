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

// checks if the DecryptionKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DecryptionKeys{}

// DecryptionKeys Decryption keys used to decrypt message content received from the partner.
type DecryptionKeys struct {
	PrimaryKeyRef       *ResourceLink `json:"primaryKeyRef,omitempty" tfsdk:"primary_key_ref"`
	SecondaryKeyPairRef *ResourceLink `json:"secondaryKeyPairRef,omitempty" tfsdk:"secondary_key_pair_ref"`
}

// NewDecryptionKeys instantiates a new DecryptionKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecryptionKeys() *DecryptionKeys {
	this := DecryptionKeys{}
	return &this
}

// NewDecryptionKeysWithDefaults instantiates a new DecryptionKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecryptionKeysWithDefaults() *DecryptionKeys {
	this := DecryptionKeys{}
	return &this
}

// GetPrimaryKeyRef returns the PrimaryKeyRef field value if set, zero value otherwise.
func (o *DecryptionKeys) GetPrimaryKeyRef() ResourceLink {
	if o == nil || IsNil(o.PrimaryKeyRef) {
		var ret ResourceLink
		return ret
	}
	return *o.PrimaryKeyRef
}

// GetPrimaryKeyRefOk returns a tuple with the PrimaryKeyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptionKeys) GetPrimaryKeyRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.PrimaryKeyRef) {
		return nil, false
	}
	return o.PrimaryKeyRef, true
}

// HasPrimaryKeyRef returns a boolean if a field has been set.
func (o *DecryptionKeys) HasPrimaryKeyRef() bool {
	if o != nil && !IsNil(o.PrimaryKeyRef) {
		return true
	}

	return false
}

// SetPrimaryKeyRef gets a reference to the given ResourceLink and assigns it to the PrimaryKeyRef field.
func (o *DecryptionKeys) SetPrimaryKeyRef(v ResourceLink) {
	o.PrimaryKeyRef = &v
}

// GetSecondaryKeyPairRef returns the SecondaryKeyPairRef field value if set, zero value otherwise.
func (o *DecryptionKeys) GetSecondaryKeyPairRef() ResourceLink {
	if o == nil || IsNil(o.SecondaryKeyPairRef) {
		var ret ResourceLink
		return ret
	}
	return *o.SecondaryKeyPairRef
}

// GetSecondaryKeyPairRefOk returns a tuple with the SecondaryKeyPairRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptionKeys) GetSecondaryKeyPairRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.SecondaryKeyPairRef) {
		return nil, false
	}
	return o.SecondaryKeyPairRef, true
}

// HasSecondaryKeyPairRef returns a boolean if a field has been set.
func (o *DecryptionKeys) HasSecondaryKeyPairRef() bool {
	if o != nil && !IsNil(o.SecondaryKeyPairRef) {
		return true
	}

	return false
}

// SetSecondaryKeyPairRef gets a reference to the given ResourceLink and assigns it to the SecondaryKeyPairRef field.
func (o *DecryptionKeys) SetSecondaryKeyPairRef(v ResourceLink) {
	o.SecondaryKeyPairRef = &v
}

func (o DecryptionKeys) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DecryptionKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrimaryKeyRef) {
		toSerialize["primaryKeyRef"] = o.PrimaryKeyRef
	}
	if !IsNil(o.SecondaryKeyPairRef) {
		toSerialize["secondaryKeyPairRef"] = o.SecondaryKeyPairRef
	}
	return toSerialize, nil
}

type NullableDecryptionKeys struct {
	value *DecryptionKeys
	isSet bool
}

func (v NullableDecryptionKeys) Get() *DecryptionKeys {
	return v.value
}

func (v *NullableDecryptionKeys) Set(val *DecryptionKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableDecryptionKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableDecryptionKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecryptionKeys(val *DecryptionKeys) *NullableDecryptionKeys {
	return &NullableDecryptionKeys{value: val, isSet: true}
}

func (v NullableDecryptionKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecryptionKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
