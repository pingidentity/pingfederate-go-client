/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the KerberosKeySet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KerberosKeySet{}

// KerberosKeySet Represents a set of Kerberos encryption keys.
type KerberosKeySet struct {
	// The encrypted key set.
	EncryptedKeySet string `json:"encryptedKeySet" tfsdk:"encrypted_key_set"`
	// Time at which the key set was deactivated due to password change. This field is not populated if the key set is active.
	DeactivatedAt *time.Time `json:"deactivatedAt,omitempty" tfsdk:"deactivated_at"`
}

type _KerberosKeySet KerberosKeySet

// NewKerberosKeySet instantiates a new KerberosKeySet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKerberosKeySet(encryptedKeySet string) *KerberosKeySet {
	this := KerberosKeySet{}
	this.EncryptedKeySet = encryptedKeySet
	return &this
}

// NewKerberosKeySetWithDefaults instantiates a new KerberosKeySet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKerberosKeySetWithDefaults() *KerberosKeySet {
	this := KerberosKeySet{}
	return &this
}

// GetEncryptedKeySet returns the EncryptedKeySet field value
func (o *KerberosKeySet) GetEncryptedKeySet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncryptedKeySet
}

// GetEncryptedKeySetOk returns a tuple with the EncryptedKeySet field value
// and a boolean to check if the value has been set.
func (o *KerberosKeySet) GetEncryptedKeySetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptedKeySet, true
}

// SetEncryptedKeySet sets field value
func (o *KerberosKeySet) SetEncryptedKeySet(v string) {
	o.EncryptedKeySet = v
}

// GetDeactivatedAt returns the DeactivatedAt field value if set, zero value otherwise.
func (o *KerberosKeySet) GetDeactivatedAt() time.Time {
	if o == nil || IsNil(o.DeactivatedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeactivatedAt
}

// GetDeactivatedAtOk returns a tuple with the DeactivatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KerberosKeySet) GetDeactivatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeactivatedAt) {
		return nil, false
	}
	return o.DeactivatedAt, true
}

// HasDeactivatedAt returns a boolean if a field has been set.
func (o *KerberosKeySet) HasDeactivatedAt() bool {
	if o != nil && !IsNil(o.DeactivatedAt) {
		return true
	}

	return false
}

// SetDeactivatedAt gets a reference to the given time.Time and assigns it to the DeactivatedAt field.
func (o *KerberosKeySet) SetDeactivatedAt(v time.Time) {
	o.DeactivatedAt = &v
}

func (o KerberosKeySet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KerberosKeySet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["encryptedKeySet"] = o.EncryptedKeySet
	if !IsNil(o.DeactivatedAt) {
		toSerialize["deactivatedAt"] = o.DeactivatedAt
	}
	return toSerialize, nil
}

func (o *KerberosKeySet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"encryptedKeySet",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKerberosKeySet := _KerberosKeySet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varKerberosKeySet)

	if err != nil {
		return err
	}

	*o = KerberosKeySet(varKerberosKeySet)

	return err
}

type NullableKerberosKeySet struct {
	value *KerberosKeySet
	isSet bool
}

func (v NullableKerberosKeySet) Get() *KerberosKeySet {
	return v.value
}

func (v *NullableKerberosKeySet) Set(val *KerberosKeySet) {
	v.value = val
	v.isSet = true
}

func (v NullableKerberosKeySet) IsSet() bool {
	return v.isSet
}

func (v *NullableKerberosKeySet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKerberosKeySet(val *KerberosKeySet) *NullableKerberosKeySet {
	return &NullableKerberosKeySet{value: val, isSet: true}
}

func (v NullableKerberosKeySet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKerberosKeySet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
