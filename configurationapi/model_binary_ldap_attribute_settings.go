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

// checks if the BinaryLdapAttributeSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BinaryLdapAttributeSettings{}

// BinaryLdapAttributeSettings Binary settings for a LDAP attribute.
type BinaryLdapAttributeSettings struct {
	// Get the encoding type for this attribute. If not specified, the default is BASE64.
	BinaryEncoding *string `json:"binaryEncoding,omitempty" tfsdk:"binary_encoding"`
}

// NewBinaryLdapAttributeSettings instantiates a new BinaryLdapAttributeSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBinaryLdapAttributeSettings() *BinaryLdapAttributeSettings {
	this := BinaryLdapAttributeSettings{}
	return &this
}

// NewBinaryLdapAttributeSettingsWithDefaults instantiates a new BinaryLdapAttributeSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBinaryLdapAttributeSettingsWithDefaults() *BinaryLdapAttributeSettings {
	this := BinaryLdapAttributeSettings{}
	return &this
}

// GetBinaryEncoding returns the BinaryEncoding field value if set, zero value otherwise.
func (o *BinaryLdapAttributeSettings) GetBinaryEncoding() string {
	if o == nil || IsNil(o.BinaryEncoding) {
		var ret string
		return ret
	}
	return *o.BinaryEncoding
}

// GetBinaryEncodingOk returns a tuple with the BinaryEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BinaryLdapAttributeSettings) GetBinaryEncodingOk() (*string, bool) {
	if o == nil || IsNil(o.BinaryEncoding) {
		return nil, false
	}
	return o.BinaryEncoding, true
}

// HasBinaryEncoding returns a boolean if a field has been set.
func (o *BinaryLdapAttributeSettings) HasBinaryEncoding() bool {
	if o != nil && !IsNil(o.BinaryEncoding) {
		return true
	}

	return false
}

// SetBinaryEncoding gets a reference to the given string and assigns it to the BinaryEncoding field.
func (o *BinaryLdapAttributeSettings) SetBinaryEncoding(v string) {
	o.BinaryEncoding = &v
}

func (o BinaryLdapAttributeSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BinaryLdapAttributeSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BinaryEncoding) {
		toSerialize["binaryEncoding"] = o.BinaryEncoding
	}
	return toSerialize, nil
}

type NullableBinaryLdapAttributeSettings struct {
	value *BinaryLdapAttributeSettings
	isSet bool
}

func (v NullableBinaryLdapAttributeSettings) Get() *BinaryLdapAttributeSettings {
	return v.value
}

func (v *NullableBinaryLdapAttributeSettings) Set(val *BinaryLdapAttributeSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableBinaryLdapAttributeSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableBinaryLdapAttributeSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBinaryLdapAttributeSettings(val *BinaryLdapAttributeSettings) *NullableBinaryLdapAttributeSettings {
	return &NullableBinaryLdapAttributeSettings{value: val, isSet: true}
}

func (v NullableBinaryLdapAttributeSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBinaryLdapAttributeSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}