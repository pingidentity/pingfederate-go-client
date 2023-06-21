/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the OutboundBackChannelAuthAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutboundBackChannelAuthAllOf{}

// OutboundBackChannelAuthAllOf struct for OutboundBackChannelAuthAllOf
type OutboundBackChannelAuthAllOf struct {
	SslAuthKeyPairRef *ResourceLink `json:"sslAuthKeyPairRef,omitempty"`
	// Validate the partner server certificate. Default is true.
	ValidatePartnerCert *bool `json:"validatePartnerCert,omitempty"`
}

// NewOutboundBackChannelAuthAllOf instantiates a new OutboundBackChannelAuthAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutboundBackChannelAuthAllOf() *OutboundBackChannelAuthAllOf {
	this := OutboundBackChannelAuthAllOf{}
	return &this
}

// NewOutboundBackChannelAuthAllOfWithDefaults instantiates a new OutboundBackChannelAuthAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutboundBackChannelAuthAllOfWithDefaults() *OutboundBackChannelAuthAllOf {
	this := OutboundBackChannelAuthAllOf{}
	return &this
}

// GetSslAuthKeyPairRef returns the SslAuthKeyPairRef field value if set, zero value otherwise.
func (o *OutboundBackChannelAuthAllOf) GetSslAuthKeyPairRef() ResourceLink {
	if o == nil || IsNil(o.SslAuthKeyPairRef) {
		var ret ResourceLink
		return ret
	}
	return *o.SslAuthKeyPairRef
}

// GetSslAuthKeyPairRefOk returns a tuple with the SslAuthKeyPairRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundBackChannelAuthAllOf) GetSslAuthKeyPairRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.SslAuthKeyPairRef) {
		return nil, false
	}
	return o.SslAuthKeyPairRef, true
}

// HasSslAuthKeyPairRef returns a boolean if a field has been set.
func (o *OutboundBackChannelAuthAllOf) HasSslAuthKeyPairRef() bool {
	if o != nil && !IsNil(o.SslAuthKeyPairRef) {
		return true
	}

	return false
}

// SetSslAuthKeyPairRef gets a reference to the given ResourceLink and assigns it to the SslAuthKeyPairRef field.
func (o *OutboundBackChannelAuthAllOf) SetSslAuthKeyPairRef(v ResourceLink) {
	o.SslAuthKeyPairRef = &v
}

// GetValidatePartnerCert returns the ValidatePartnerCert field value if set, zero value otherwise.
func (o *OutboundBackChannelAuthAllOf) GetValidatePartnerCert() bool {
	if o == nil || IsNil(o.ValidatePartnerCert) {
		var ret bool
		return ret
	}
	return *o.ValidatePartnerCert
}

// GetValidatePartnerCertOk returns a tuple with the ValidatePartnerCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutboundBackChannelAuthAllOf) GetValidatePartnerCertOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidatePartnerCert) {
		return nil, false
	}
	return o.ValidatePartnerCert, true
}

// HasValidatePartnerCert returns a boolean if a field has been set.
func (o *OutboundBackChannelAuthAllOf) HasValidatePartnerCert() bool {
	if o != nil && !IsNil(o.ValidatePartnerCert) {
		return true
	}

	return false
}

// SetValidatePartnerCert gets a reference to the given bool and assigns it to the ValidatePartnerCert field.
func (o *OutboundBackChannelAuthAllOf) SetValidatePartnerCert(v bool) {
	o.ValidatePartnerCert = &v
}

func (o OutboundBackChannelAuthAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutboundBackChannelAuthAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SslAuthKeyPairRef) {
		toSerialize["sslAuthKeyPairRef"] = o.SslAuthKeyPairRef
	}
	if !IsNil(o.ValidatePartnerCert) {
		toSerialize["validatePartnerCert"] = o.ValidatePartnerCert
	}
	return toSerialize, nil
}

type NullableOutboundBackChannelAuthAllOf struct {
	value *OutboundBackChannelAuthAllOf
	isSet bool
}

func (v NullableOutboundBackChannelAuthAllOf) Get() *OutboundBackChannelAuthAllOf {
	return v.value
}

func (v *NullableOutboundBackChannelAuthAllOf) Set(val *OutboundBackChannelAuthAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOutboundBackChannelAuthAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOutboundBackChannelAuthAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutboundBackChannelAuthAllOf(val *OutboundBackChannelAuthAllOf) *NullableOutboundBackChannelAuthAllOf {
	return &NullableOutboundBackChannelAuthAllOf{value: val, isSet: true}
}

func (v NullableOutboundBackChannelAuthAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutboundBackChannelAuthAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
