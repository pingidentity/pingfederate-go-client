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

// checks if the BaseSigningSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseSigningSettings{}

// BaseSigningSettings Settings related to signing messages.
type BaseSigningSettings struct {
	SigningKeyPairRef ResourceLink `json:"signingKeyPairRef" tfsdk:"signing_key_pair_ref"`
	// The algorithm used to sign messages sent to this partner. The default is SHA1withDSA for DSA certs, SHA256withRSA for RSA certs, and SHA256withECDSA for EC certs. For RSA certs, SHA1withRSA, SHA384withRSA, SHA512withRSA, SHA256withRSAandMGF1, SHA384withRSAandMGF1 and SHA512withRSAandMGF1 are also supported. For EC certs, SHA384withECDSA and SHA512withECDSA are also supported. If the connection is WS-Federation with JWT token type, then the possible values are RSA SHA256, RSA SHA384, RSA SHA512, RSASSA-PSS SHA256, RSASSA-PSS SHA384, RSASSA-PSS SHA512, ECDSA SHA256, ECDSA SHA384, ECDSA SHA512
	Algorithm *string `json:"algorithm,omitempty" tfsdk:"algorithm"`
	// Determines whether the signing certificate is included in the signature <KeyInfo> element.
	IncludeCertInSignature *bool `json:"includeCertInSignature,omitempty" tfsdk:"include_cert_in_signature"`
	// Determines whether the <KeyValue> element with the raw public key is included in the signature <KeyInfo> element.
	IncludeRawKeyInSignature *bool `json:"includeRawKeyInSignature,omitempty" tfsdk:"include_raw_key_in_signature"`
}

// NewBaseSigningSettings instantiates a new BaseSigningSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseSigningSettings(signingKeyPairRef ResourceLink) *BaseSigningSettings {
	this := BaseSigningSettings{}
	this.SigningKeyPairRef = signingKeyPairRef
	return &this
}

// NewBaseSigningSettingsWithDefaults instantiates a new BaseSigningSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseSigningSettingsWithDefaults() *BaseSigningSettings {
	this := BaseSigningSettings{}
	return &this
}

// GetSigningKeyPairRef returns the SigningKeyPairRef field value
func (o *BaseSigningSettings) GetSigningKeyPairRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.SigningKeyPairRef
}

// GetSigningKeyPairRefOk returns a tuple with the SigningKeyPairRef field value
// and a boolean to check if the value has been set.
func (o *BaseSigningSettings) GetSigningKeyPairRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningKeyPairRef, true
}

// SetSigningKeyPairRef sets field value
func (o *BaseSigningSettings) SetSigningKeyPairRef(v ResourceLink) {
	o.SigningKeyPairRef = v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *BaseSigningSettings) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSigningSettings) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *BaseSigningSettings) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *BaseSigningSettings) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetIncludeCertInSignature returns the IncludeCertInSignature field value if set, zero value otherwise.
func (o *BaseSigningSettings) GetIncludeCertInSignature() bool {
	if o == nil || IsNil(o.IncludeCertInSignature) {
		var ret bool
		return ret
	}
	return *o.IncludeCertInSignature
}

// GetIncludeCertInSignatureOk returns a tuple with the IncludeCertInSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSigningSettings) GetIncludeCertInSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeCertInSignature) {
		return nil, false
	}
	return o.IncludeCertInSignature, true
}

// HasIncludeCertInSignature returns a boolean if a field has been set.
func (o *BaseSigningSettings) HasIncludeCertInSignature() bool {
	if o != nil && !IsNil(o.IncludeCertInSignature) {
		return true
	}

	return false
}

// SetIncludeCertInSignature gets a reference to the given bool and assigns it to the IncludeCertInSignature field.
func (o *BaseSigningSettings) SetIncludeCertInSignature(v bool) {
	o.IncludeCertInSignature = &v
}

// GetIncludeRawKeyInSignature returns the IncludeRawKeyInSignature field value if set, zero value otherwise.
func (o *BaseSigningSettings) GetIncludeRawKeyInSignature() bool {
	if o == nil || IsNil(o.IncludeRawKeyInSignature) {
		var ret bool
		return ret
	}
	return *o.IncludeRawKeyInSignature
}

// GetIncludeRawKeyInSignatureOk returns a tuple with the IncludeRawKeyInSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSigningSettings) GetIncludeRawKeyInSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeRawKeyInSignature) {
		return nil, false
	}
	return o.IncludeRawKeyInSignature, true
}

// HasIncludeRawKeyInSignature returns a boolean if a field has been set.
func (o *BaseSigningSettings) HasIncludeRawKeyInSignature() bool {
	if o != nil && !IsNil(o.IncludeRawKeyInSignature) {
		return true
	}

	return false
}

// SetIncludeRawKeyInSignature gets a reference to the given bool and assigns it to the IncludeRawKeyInSignature field.
func (o *BaseSigningSettings) SetIncludeRawKeyInSignature(v bool) {
	o.IncludeRawKeyInSignature = &v
}

func (o BaseSigningSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseSigningSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signingKeyPairRef"] = o.SigningKeyPairRef
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.IncludeCertInSignature) {
		toSerialize["includeCertInSignature"] = o.IncludeCertInSignature
	}
	if !IsNil(o.IncludeRawKeyInSignature) {
		toSerialize["includeRawKeyInSignature"] = o.IncludeRawKeyInSignature
	}
	return toSerialize, nil
}

type NullableBaseSigningSettings struct {
	value *BaseSigningSettings
	isSet bool
}

func (v NullableBaseSigningSettings) Get() *BaseSigningSettings {
	return v.value
}

func (v *NullableBaseSigningSettings) Set(val *BaseSigningSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseSigningSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseSigningSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseSigningSettings(val *BaseSigningSettings) *NullableBaseSigningSettings {
	return &NullableBaseSigningSettings{value: val, isSet: true}
}

func (v NullableBaseSigningSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseSigningSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
