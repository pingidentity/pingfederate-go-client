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

// checks if the RsaAlgKeyId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RsaAlgKeyId{}

// RsaAlgKeyId This class represents a pair of RSA algorithm type and key identifier.
type RsaAlgKeyId struct {
	// The RSA signing algorithm type. The supported RSA signing algorithm types are RS256, RS384, RS512, PS256, PS384 and PS512.
	RsaAlgType string `json:"rsaAlgType" tfsdk:"rsa_alg_type"`
	// Unique key identifier.
	KeyId string `json:"keyId" tfsdk:"key_id"`
}

// NewRsaAlgKeyId instantiates a new RsaAlgKeyId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRsaAlgKeyId(rsaAlgType string, keyId string) *RsaAlgKeyId {
	this := RsaAlgKeyId{}
	this.RsaAlgType = rsaAlgType
	this.KeyId = keyId
	return &this
}

// NewRsaAlgKeyIdWithDefaults instantiates a new RsaAlgKeyId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRsaAlgKeyIdWithDefaults() *RsaAlgKeyId {
	this := RsaAlgKeyId{}
	return &this
}

// GetRsaAlgType returns the RsaAlgType field value
func (o *RsaAlgKeyId) GetRsaAlgType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RsaAlgType
}

// GetRsaAlgTypeOk returns a tuple with the RsaAlgType field value
// and a boolean to check if the value has been set.
func (o *RsaAlgKeyId) GetRsaAlgTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RsaAlgType, true
}

// SetRsaAlgType sets field value
func (o *RsaAlgKeyId) SetRsaAlgType(v string) {
	o.RsaAlgType = v
}

// GetKeyId returns the KeyId field value
func (o *RsaAlgKeyId) GetKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *RsaAlgKeyId) GetKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyId, true
}

// SetKeyId sets field value
func (o *RsaAlgKeyId) SetKeyId(v string) {
	o.KeyId = v
}

func (o RsaAlgKeyId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RsaAlgKeyId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rsaAlgType"] = o.RsaAlgType
	toSerialize["keyId"] = o.KeyId
	return toSerialize, nil
}

type NullableRsaAlgKeyId struct {
	value *RsaAlgKeyId
	isSet bool
}

func (v NullableRsaAlgKeyId) Get() *RsaAlgKeyId {
	return v.value
}

func (v *NullableRsaAlgKeyId) Set(val *RsaAlgKeyId) {
	v.value = val
	v.isSet = true
}

func (v NullableRsaAlgKeyId) IsSet() bool {
	return v.isSet
}

func (v *NullableRsaAlgKeyId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRsaAlgKeyId(val *RsaAlgKeyId) *NullableRsaAlgKeyId {
	return &NullableRsaAlgKeyId{value: val, isSet: true}
}

func (v NullableRsaAlgKeyId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRsaAlgKeyId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
