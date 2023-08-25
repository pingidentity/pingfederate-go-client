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

// checks if the ClientAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientAuth{}

// ClientAuth Client Authentication.
type ClientAuth struct {
	// Client authentication type.<br>The required field for type SECRET is secret.<br>The required fields for type CERTIFICATE are clientCertIssuerDn and clientCertSubjectDn.<br>The required field for type PRIVATE_KEY_JWT is: either jwks or jwksUrl.
	Type *string `json:"type,omitempty" tfsdk:"type"`
	// Client secret for Basic Authentication.  To update the client secret, specify the plaintext value in this field.  This field will not be populated for GET requests.
	Secret *string `json:"secret,omitempty" tfsdk:"secret"`
	// For GET requests, this field contains the encrypted client secret, if one exists.  For POST and PUT requests, if you wish to reuse the existing secret, this field should be passed back unchanged.
	EncryptedSecret *string `json:"encryptedSecret,omitempty" tfsdk:"encrypted_secret"`
	// The list of secondary client secrets that are temporarily retained.
	SecondarySecrets []SecondarySecret `json:"secondarySecrets,omitempty" tfsdk:"secondary_secrets"`
	// Client TLS Certificate Issuer DN.
	ClientCertIssuerDn *string `json:"clientCertIssuerDn,omitempty" tfsdk:"client_cert_issuer_dn"`
	// Client TLS Certificate Subject DN.
	ClientCertSubjectDn *string `json:"clientCertSubjectDn,omitempty" tfsdk:"client_cert_subject_dn"`
	// Enforce replay prevention on JSON Web Tokens. This field is applicable only for Private Key JWT Client Authentication.
	EnforceReplayPrevention *bool `json:"enforceReplayPrevention,omitempty" tfsdk:"enforce_replay_prevention"`
	// The JSON Web Signature [JWS] algorithm that must be used to sign the JSON Web Tokens. This field is applicable only for Private Key JWT Client Authentication. All signing algorithms are allowed if value is not present <br>RS256 - RSA using SHA-256<br>RS384 - RSA using SHA-384<br>RS512 - RSA using SHA-512<br>ES256 - ECDSA using P256 Curve and SHA-256<br>ES384 - ECDSA using P384 Curve and SHA-384<br>ES512 - ECDSA using P521 Curve and SHA-512<br>PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256<br>PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384<br>PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512<br>RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11.
	TokenEndpointAuthSigningAlgorithm *string `json:"tokenEndpointAuthSigningAlgorithm,omitempty" tfsdk:"token_endpoint_auth_signing_algorithm"`
}

// NewClientAuth instantiates a new ClientAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientAuth() *ClientAuth {
	this := ClientAuth{}
	return &this
}

// NewClientAuthWithDefaults instantiates a new ClientAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientAuthWithDefaults() *ClientAuth {
	this := ClientAuth{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ClientAuth) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuth) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClientAuth) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClientAuth) SetType(v string) {
	o.Type = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ClientAuth) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuth) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ClientAuth) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ClientAuth) SetSecret(v string) {
	o.Secret = &v
}

// GetEncryptedSecret returns the EncryptedSecret field value if set, zero value otherwise.
func (o *ClientAuth) GetEncryptedSecret() string {
	if o == nil || IsNil(o.EncryptedSecret) {
		var ret string
		return ret
	}
	return *o.EncryptedSecret
}

// GetEncryptedSecretOk returns a tuple with the EncryptedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuth) GetEncryptedSecretOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedSecret) {
		return nil, false
	}
	return o.EncryptedSecret, true
}

// HasEncryptedSecret returns a boolean if a field has been set.
func (o *ClientAuth) HasEncryptedSecret() bool {
	if o != nil && !IsNil(o.EncryptedSecret) {
		return true
	}

	return false
}

// SetEncryptedSecret gets a reference to the given string and assigns it to the EncryptedSecret field.
func (o *ClientAuth) SetEncryptedSecret(v string) {
	o.EncryptedSecret = &v
}

// GetSecondarySecrets returns the SecondarySecrets field value if set, zero value otherwise.
func (o *ClientAuth) GetSecondarySecrets() []SecondarySecret {
	if o == nil || IsNil(o.SecondarySecrets) {
		var ret []SecondarySecret
		return ret
	}
	return o.SecondarySecrets
}

// GetSecondarySecretsOk returns a tuple with the SecondarySecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuth) GetSecondarySecretsOk() ([]SecondarySecret, bool) {
	if o == nil || IsNil(o.SecondarySecrets) {
		return nil, false
	}
	return o.SecondarySecrets, true
}

// HasSecondarySecrets returns a boolean if a field has been set.
func (o *ClientAuth) HasSecondarySecrets() bool {
	if o != nil && !IsNil(o.SecondarySecrets) {
		return true
	}

	return false
}

// SetSecondarySecrets gets a reference to the given []SecondarySecret and assigns it to the SecondarySecrets field.
func (o *ClientAuth) SetSecondarySecrets(v []SecondarySecret) {
	o.SecondarySecrets = v
}

// GetClientCertIssuerDn returns the ClientCertIssuerDn field value if set, zero value otherwise.
func (o *ClientAuth) GetClientCertIssuerDn() string {
	if o == nil || IsNil(o.ClientCertIssuerDn) {
		var ret string
		return ret
	}
	return *o.ClientCertIssuerDn
}

// GetClientCertIssuerDnOk returns a tuple with the ClientCertIssuerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuth) GetClientCertIssuerDnOk() (*string, bool) {
	if o == nil || IsNil(o.ClientCertIssuerDn) {
		return nil, false
	}
	return o.ClientCertIssuerDn, true
}

// HasClientCertIssuerDn returns a boolean if a field has been set.
func (o *ClientAuth) HasClientCertIssuerDn() bool {
	if o != nil && !IsNil(o.ClientCertIssuerDn) {
		return true
	}

	return false
}

// SetClientCertIssuerDn gets a reference to the given string and assigns it to the ClientCertIssuerDn field.
func (o *ClientAuth) SetClientCertIssuerDn(v string) {
	o.ClientCertIssuerDn = &v
}

// GetClientCertSubjectDn returns the ClientCertSubjectDn field value if set, zero value otherwise.
func (o *ClientAuth) GetClientCertSubjectDn() string {
	if o == nil || IsNil(o.ClientCertSubjectDn) {
		var ret string
		return ret
	}
	return *o.ClientCertSubjectDn
}

// GetClientCertSubjectDnOk returns a tuple with the ClientCertSubjectDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuth) GetClientCertSubjectDnOk() (*string, bool) {
	if o == nil || IsNil(o.ClientCertSubjectDn) {
		return nil, false
	}
	return o.ClientCertSubjectDn, true
}

// HasClientCertSubjectDn returns a boolean if a field has been set.
func (o *ClientAuth) HasClientCertSubjectDn() bool {
	if o != nil && !IsNil(o.ClientCertSubjectDn) {
		return true
	}

	return false
}

// SetClientCertSubjectDn gets a reference to the given string and assigns it to the ClientCertSubjectDn field.
func (o *ClientAuth) SetClientCertSubjectDn(v string) {
	o.ClientCertSubjectDn = &v
}

// GetEnforceReplayPrevention returns the EnforceReplayPrevention field value if set, zero value otherwise.
func (o *ClientAuth) GetEnforceReplayPrevention() bool {
	if o == nil || IsNil(o.EnforceReplayPrevention) {
		var ret bool
		return ret
	}
	return *o.EnforceReplayPrevention
}

// GetEnforceReplayPreventionOk returns a tuple with the EnforceReplayPrevention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuth) GetEnforceReplayPreventionOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceReplayPrevention) {
		return nil, false
	}
	return o.EnforceReplayPrevention, true
}

// HasEnforceReplayPrevention returns a boolean if a field has been set.
func (o *ClientAuth) HasEnforceReplayPrevention() bool {
	if o != nil && !IsNil(o.EnforceReplayPrevention) {
		return true
	}

	return false
}

// SetEnforceReplayPrevention gets a reference to the given bool and assigns it to the EnforceReplayPrevention field.
func (o *ClientAuth) SetEnforceReplayPrevention(v bool) {
	o.EnforceReplayPrevention = &v
}

// GetTokenEndpointAuthSigningAlgorithm returns the TokenEndpointAuthSigningAlgorithm field value if set, zero value otherwise.
func (o *ClientAuth) GetTokenEndpointAuthSigningAlgorithm() string {
	if o == nil || IsNil(o.TokenEndpointAuthSigningAlgorithm) {
		var ret string
		return ret
	}
	return *o.TokenEndpointAuthSigningAlgorithm
}

// GetTokenEndpointAuthSigningAlgorithmOk returns a tuple with the TokenEndpointAuthSigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAuth) GetTokenEndpointAuthSigningAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.TokenEndpointAuthSigningAlgorithm) {
		return nil, false
	}
	return o.TokenEndpointAuthSigningAlgorithm, true
}

// HasTokenEndpointAuthSigningAlgorithm returns a boolean if a field has been set.
func (o *ClientAuth) HasTokenEndpointAuthSigningAlgorithm() bool {
	if o != nil && !IsNil(o.TokenEndpointAuthSigningAlgorithm) {
		return true
	}

	return false
}

// SetTokenEndpointAuthSigningAlgorithm gets a reference to the given string and assigns it to the TokenEndpointAuthSigningAlgorithm field.
func (o *ClientAuth) SetTokenEndpointAuthSigningAlgorithm(v string) {
	o.TokenEndpointAuthSigningAlgorithm = &v
}

func (o ClientAuth) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.EncryptedSecret) {
		toSerialize["encryptedSecret"] = o.EncryptedSecret
	}
	if !IsNil(o.SecondarySecrets) {
		toSerialize["secondarySecrets"] = o.SecondarySecrets
	}
	if !IsNil(o.ClientCertIssuerDn) {
		toSerialize["clientCertIssuerDn"] = o.ClientCertIssuerDn
	}
	if !IsNil(o.ClientCertSubjectDn) {
		toSerialize["clientCertSubjectDn"] = o.ClientCertSubjectDn
	}
	if !IsNil(o.EnforceReplayPrevention) {
		toSerialize["enforceReplayPrevention"] = o.EnforceReplayPrevention
	}
	if !IsNil(o.TokenEndpointAuthSigningAlgorithm) {
		toSerialize["tokenEndpointAuthSigningAlgorithm"] = o.TokenEndpointAuthSigningAlgorithm
	}
	return toSerialize, nil
}

type NullableClientAuth struct {
	value *ClientAuth
	isSet bool
}

func (v NullableClientAuth) Get() *ClientAuth {
	return v.value
}

func (v *NullableClientAuth) Set(val *ClientAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAuth(val *ClientAuth) *NullableClientAuth {
	return &NullableClientAuth{value: val, isSet: true}
}

func (v NullableClientAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
