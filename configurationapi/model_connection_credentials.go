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

// checks if the ConnectionCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionCredentials{}

// ConnectionCredentials The certificates and settings for encryption, signing, and signature verification.
type ConnectionCredentials struct {
	// If this property is set, the verification trust model is Anchored. The verification certificate must be signed by a trusted CA and included in the incoming message, and the subject DN of the expected certificate is specified in this property. If this property is not set, then a primary verification certificate must be specified in the certs array.
	VerificationSubjectDN *string `json:"verificationSubjectDN,omitempty" tfsdk:"verification_subject_dn"`
	// If a verification Subject DN is provided, you can optionally restrict the issuer to a specific trusted CA by specifying its DN in this field.
	VerificationIssuerDN *string `json:"verificationIssuerDN,omitempty" tfsdk:"verification_issuer_dn"`
	// The certificates used for signature verification and XML encryption.
	Certs []ConnectionCert `json:"certs,omitempty" tfsdk:"certs"`
	// The algorithm used to encrypt assertions sent to this partner. AES_128, AES_256, AES_128_GCM, AES_192_GCM, AES_256_GCM and Triple_DES are supported.
	BlockEncryptionAlgorithm *string `json:"blockEncryptionAlgorithm,omitempty" tfsdk:"block_encryption_algorithm"`
	// The algorithm used to transport keys to this partner. RSA_OAEP, RSA_OAEP_256 and RSA_v15 are supported.
	KeyTransportAlgorithm         *string                  `json:"keyTransportAlgorithm,omitempty" tfsdk:"key_transport_algorithm"`
	SigningSettings               *SigningSettings         `json:"signingSettings,omitempty" tfsdk:"signing_settings"`
	DecryptionKeyPairRef          *ResourceLink            `json:"decryptionKeyPairRef,omitempty" tfsdk:"decryption_key_pair_ref"`
	SecondaryDecryptionKeyPairRef *ResourceLink            `json:"secondaryDecryptionKeyPairRef,omitempty" tfsdk:"secondary_decryption_key_pair_ref"`
	OutboundBackChannelAuth       *OutboundBackChannelAuth `json:"outboundBackChannelAuth,omitempty" tfsdk:"outbound_back_channel_auth"`
	InboundBackChannelAuth        *InboundBackChannelAuth  `json:"inboundBackChannelAuth,omitempty" tfsdk:"inbound_back_channel_auth"`
}

// NewConnectionCredentials instantiates a new ConnectionCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionCredentials() *ConnectionCredentials {
	this := ConnectionCredentials{}
	return &this
}

// NewConnectionCredentialsWithDefaults instantiates a new ConnectionCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionCredentialsWithDefaults() *ConnectionCredentials {
	this := ConnectionCredentials{}
	return &this
}

// GetVerificationSubjectDN returns the VerificationSubjectDN field value if set, zero value otherwise.
func (o *ConnectionCredentials) GetVerificationSubjectDN() string {
	if o == nil || IsNil(o.VerificationSubjectDN) {
		var ret string
		return ret
	}
	return *o.VerificationSubjectDN
}

// GetVerificationSubjectDNOk returns a tuple with the VerificationSubjectDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCredentials) GetVerificationSubjectDNOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationSubjectDN) {
		return nil, false
	}
	return o.VerificationSubjectDN, true
}

// HasVerificationSubjectDN returns a boolean if a field has been set.
func (o *ConnectionCredentials) HasVerificationSubjectDN() bool {
	if o != nil && !IsNil(o.VerificationSubjectDN) {
		return true
	}

	return false
}

// SetVerificationSubjectDN gets a reference to the given string and assigns it to the VerificationSubjectDN field.
func (o *ConnectionCredentials) SetVerificationSubjectDN(v string) {
	o.VerificationSubjectDN = &v
}

// GetVerificationIssuerDN returns the VerificationIssuerDN field value if set, zero value otherwise.
func (o *ConnectionCredentials) GetVerificationIssuerDN() string {
	if o == nil || IsNil(o.VerificationIssuerDN) {
		var ret string
		return ret
	}
	return *o.VerificationIssuerDN
}

// GetVerificationIssuerDNOk returns a tuple with the VerificationIssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCredentials) GetVerificationIssuerDNOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationIssuerDN) {
		return nil, false
	}
	return o.VerificationIssuerDN, true
}

// HasVerificationIssuerDN returns a boolean if a field has been set.
func (o *ConnectionCredentials) HasVerificationIssuerDN() bool {
	if o != nil && !IsNil(o.VerificationIssuerDN) {
		return true
	}

	return false
}

// SetVerificationIssuerDN gets a reference to the given string and assigns it to the VerificationIssuerDN field.
func (o *ConnectionCredentials) SetVerificationIssuerDN(v string) {
	o.VerificationIssuerDN = &v
}

// GetCerts returns the Certs field value if set, zero value otherwise.
func (o *ConnectionCredentials) GetCerts() []ConnectionCert {
	if o == nil || IsNil(o.Certs) {
		var ret []ConnectionCert
		return ret
	}
	return o.Certs
}

// GetCertsOk returns a tuple with the Certs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCredentials) GetCertsOk() ([]ConnectionCert, bool) {
	if o == nil || IsNil(o.Certs) {
		return nil, false
	}
	return o.Certs, true
}

// HasCerts returns a boolean if a field has been set.
func (o *ConnectionCredentials) HasCerts() bool {
	if o != nil && !IsNil(o.Certs) {
		return true
	}

	return false
}

// SetCerts gets a reference to the given []ConnectionCert and assigns it to the Certs field.
func (o *ConnectionCredentials) SetCerts(v []ConnectionCert) {
	o.Certs = v
}

// GetBlockEncryptionAlgorithm returns the BlockEncryptionAlgorithm field value if set, zero value otherwise.
func (o *ConnectionCredentials) GetBlockEncryptionAlgorithm() string {
	if o == nil || IsNil(o.BlockEncryptionAlgorithm) {
		var ret string
		return ret
	}
	return *o.BlockEncryptionAlgorithm
}

// GetBlockEncryptionAlgorithmOk returns a tuple with the BlockEncryptionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCredentials) GetBlockEncryptionAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.BlockEncryptionAlgorithm) {
		return nil, false
	}
	return o.BlockEncryptionAlgorithm, true
}

// HasBlockEncryptionAlgorithm returns a boolean if a field has been set.
func (o *ConnectionCredentials) HasBlockEncryptionAlgorithm() bool {
	if o != nil && !IsNil(o.BlockEncryptionAlgorithm) {
		return true
	}

	return false
}

// SetBlockEncryptionAlgorithm gets a reference to the given string and assigns it to the BlockEncryptionAlgorithm field.
func (o *ConnectionCredentials) SetBlockEncryptionAlgorithm(v string) {
	o.BlockEncryptionAlgorithm = &v
}

// GetKeyTransportAlgorithm returns the KeyTransportAlgorithm field value if set, zero value otherwise.
func (o *ConnectionCredentials) GetKeyTransportAlgorithm() string {
	if o == nil || IsNil(o.KeyTransportAlgorithm) {
		var ret string
		return ret
	}
	return *o.KeyTransportAlgorithm
}

// GetKeyTransportAlgorithmOk returns a tuple with the KeyTransportAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCredentials) GetKeyTransportAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.KeyTransportAlgorithm) {
		return nil, false
	}
	return o.KeyTransportAlgorithm, true
}

// HasKeyTransportAlgorithm returns a boolean if a field has been set.
func (o *ConnectionCredentials) HasKeyTransportAlgorithm() bool {
	if o != nil && !IsNil(o.KeyTransportAlgorithm) {
		return true
	}

	return false
}

// SetKeyTransportAlgorithm gets a reference to the given string and assigns it to the KeyTransportAlgorithm field.
func (o *ConnectionCredentials) SetKeyTransportAlgorithm(v string) {
	o.KeyTransportAlgorithm = &v
}

// GetSigningSettings returns the SigningSettings field value if set, zero value otherwise.
func (o *ConnectionCredentials) GetSigningSettings() SigningSettings {
	if o == nil || IsNil(o.SigningSettings) {
		var ret SigningSettings
		return ret
	}
	return *o.SigningSettings
}

// GetSigningSettingsOk returns a tuple with the SigningSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCredentials) GetSigningSettingsOk() (*SigningSettings, bool) {
	if o == nil || IsNil(o.SigningSettings) {
		return nil, false
	}
	return o.SigningSettings, true
}

// HasSigningSettings returns a boolean if a field has been set.
func (o *ConnectionCredentials) HasSigningSettings() bool {
	if o != nil && !IsNil(o.SigningSettings) {
		return true
	}

	return false
}

// SetSigningSettings gets a reference to the given SigningSettings and assigns it to the SigningSettings field.
func (o *ConnectionCredentials) SetSigningSettings(v SigningSettings) {
	o.SigningSettings = &v
}

// GetDecryptionKeyPairRef returns the DecryptionKeyPairRef field value if set, zero value otherwise.
func (o *ConnectionCredentials) GetDecryptionKeyPairRef() ResourceLink {
	if o == nil || IsNil(o.DecryptionKeyPairRef) {
		var ret ResourceLink
		return ret
	}
	return *o.DecryptionKeyPairRef
}

// GetDecryptionKeyPairRefOk returns a tuple with the DecryptionKeyPairRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCredentials) GetDecryptionKeyPairRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.DecryptionKeyPairRef) {
		return nil, false
	}
	return o.DecryptionKeyPairRef, true
}

// HasDecryptionKeyPairRef returns a boolean if a field has been set.
func (o *ConnectionCredentials) HasDecryptionKeyPairRef() bool {
	if o != nil && !IsNil(o.DecryptionKeyPairRef) {
		return true
	}

	return false
}

// SetDecryptionKeyPairRef gets a reference to the given ResourceLink and assigns it to the DecryptionKeyPairRef field.
func (o *ConnectionCredentials) SetDecryptionKeyPairRef(v ResourceLink) {
	o.DecryptionKeyPairRef = &v
}

// GetSecondaryDecryptionKeyPairRef returns the SecondaryDecryptionKeyPairRef field value if set, zero value otherwise.
func (o *ConnectionCredentials) GetSecondaryDecryptionKeyPairRef() ResourceLink {
	if o == nil || IsNil(o.SecondaryDecryptionKeyPairRef) {
		var ret ResourceLink
		return ret
	}
	return *o.SecondaryDecryptionKeyPairRef
}

// GetSecondaryDecryptionKeyPairRefOk returns a tuple with the SecondaryDecryptionKeyPairRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCredentials) GetSecondaryDecryptionKeyPairRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.SecondaryDecryptionKeyPairRef) {
		return nil, false
	}
	return o.SecondaryDecryptionKeyPairRef, true
}

// HasSecondaryDecryptionKeyPairRef returns a boolean if a field has been set.
func (o *ConnectionCredentials) HasSecondaryDecryptionKeyPairRef() bool {
	if o != nil && !IsNil(o.SecondaryDecryptionKeyPairRef) {
		return true
	}

	return false
}

// SetSecondaryDecryptionKeyPairRef gets a reference to the given ResourceLink and assigns it to the SecondaryDecryptionKeyPairRef field.
func (o *ConnectionCredentials) SetSecondaryDecryptionKeyPairRef(v ResourceLink) {
	o.SecondaryDecryptionKeyPairRef = &v
}

// GetOutboundBackChannelAuth returns the OutboundBackChannelAuth field value if set, zero value otherwise.
func (o *ConnectionCredentials) GetOutboundBackChannelAuth() OutboundBackChannelAuth {
	if o == nil || IsNil(o.OutboundBackChannelAuth) {
		var ret OutboundBackChannelAuth
		return ret
	}
	return *o.OutboundBackChannelAuth
}

// GetOutboundBackChannelAuthOk returns a tuple with the OutboundBackChannelAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCredentials) GetOutboundBackChannelAuthOk() (*OutboundBackChannelAuth, bool) {
	if o == nil || IsNil(o.OutboundBackChannelAuth) {
		return nil, false
	}
	return o.OutboundBackChannelAuth, true
}

// HasOutboundBackChannelAuth returns a boolean if a field has been set.
func (o *ConnectionCredentials) HasOutboundBackChannelAuth() bool {
	if o != nil && !IsNil(o.OutboundBackChannelAuth) {
		return true
	}

	return false
}

// SetOutboundBackChannelAuth gets a reference to the given OutboundBackChannelAuth and assigns it to the OutboundBackChannelAuth field.
func (o *ConnectionCredentials) SetOutboundBackChannelAuth(v OutboundBackChannelAuth) {
	o.OutboundBackChannelAuth = &v
}

// GetInboundBackChannelAuth returns the InboundBackChannelAuth field value if set, zero value otherwise.
func (o *ConnectionCredentials) GetInboundBackChannelAuth() InboundBackChannelAuth {
	if o == nil || IsNil(o.InboundBackChannelAuth) {
		var ret InboundBackChannelAuth
		return ret
	}
	return *o.InboundBackChannelAuth
}

// GetInboundBackChannelAuthOk returns a tuple with the InboundBackChannelAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionCredentials) GetInboundBackChannelAuthOk() (*InboundBackChannelAuth, bool) {
	if o == nil || IsNil(o.InboundBackChannelAuth) {
		return nil, false
	}
	return o.InboundBackChannelAuth, true
}

// HasInboundBackChannelAuth returns a boolean if a field has been set.
func (o *ConnectionCredentials) HasInboundBackChannelAuth() bool {
	if o != nil && !IsNil(o.InboundBackChannelAuth) {
		return true
	}

	return false
}

// SetInboundBackChannelAuth gets a reference to the given InboundBackChannelAuth and assigns it to the InboundBackChannelAuth field.
func (o *ConnectionCredentials) SetInboundBackChannelAuth(v InboundBackChannelAuth) {
	o.InboundBackChannelAuth = &v
}

func (o ConnectionCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VerificationSubjectDN) {
		toSerialize["verificationSubjectDN"] = o.VerificationSubjectDN
	}
	if !IsNil(o.VerificationIssuerDN) {
		toSerialize["verificationIssuerDN"] = o.VerificationIssuerDN
	}
	if !IsNil(o.Certs) {
		toSerialize["certs"] = o.Certs
	}
	if !IsNil(o.BlockEncryptionAlgorithm) {
		toSerialize["blockEncryptionAlgorithm"] = o.BlockEncryptionAlgorithm
	}
	if !IsNil(o.KeyTransportAlgorithm) {
		toSerialize["keyTransportAlgorithm"] = o.KeyTransportAlgorithm
	}
	if !IsNil(o.SigningSettings) {
		toSerialize["signingSettings"] = o.SigningSettings
	}
	if !IsNil(o.DecryptionKeyPairRef) {
		toSerialize["decryptionKeyPairRef"] = o.DecryptionKeyPairRef
	}
	if !IsNil(o.SecondaryDecryptionKeyPairRef) {
		toSerialize["secondaryDecryptionKeyPairRef"] = o.SecondaryDecryptionKeyPairRef
	}
	if !IsNil(o.OutboundBackChannelAuth) {
		toSerialize["outboundBackChannelAuth"] = o.OutboundBackChannelAuth
	}
	if !IsNil(o.InboundBackChannelAuth) {
		toSerialize["inboundBackChannelAuth"] = o.InboundBackChannelAuth
	}
	return toSerialize, nil
}

type NullableConnectionCredentials struct {
	value *ConnectionCredentials
	isSet bool
}

func (v NullableConnectionCredentials) Get() *ConnectionCredentials {
	return v.value
}

func (v *NullableConnectionCredentials) Set(val *ConnectionCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionCredentials(val *ConnectionCredentials) *NullableConnectionCredentials {
	return &NullableConnectionCredentials{value: val, isSet: true}
}

func (v NullableConnectionCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
