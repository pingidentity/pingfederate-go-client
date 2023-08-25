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

// checks if the ClientRegistrationOIDCPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientRegistrationOIDCPolicy{}

// ClientRegistrationOIDCPolicy Client Registration Open ID Connect Policy settings.
type ClientRegistrationOIDCPolicy struct {
	// The JSON Web Signature [JWS] algorithm required for the ID Token.<br>NONE - No signing algorithm<br>HS256 - HMAC using SHA-256<br>HS384 - HMAC using SHA-384<br>HS512 - HMAC using SHA-512<br>RS256 - RSA using SHA-256<br>RS384 - RSA using SHA-384<br>RS512 - RSA using SHA-512<br>ES256 - ECDSA using P256 Curve and SHA-256<br>ES384 - ECDSA using P384 Curve and SHA-384<br>ES512 - ECDSA using P521 Curve and SHA-512<br>PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256<br>PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384<br>PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512<br>A null value will represent the default algorithm which is RS256.<br>RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11
	IdTokenSigningAlgorithm *string `json:"idTokenSigningAlgorithm,omitempty" tfsdk:"id_token_signing_algorithm"`
	// The JSON Web Encryption [JWE] encryption algorithm used to encrypt the content encryption key for the ID Token.<br>DIR - Direct Encryption with symmetric key<br>A128KW - AES-128 Key Wrap<br>A192KW - AES-192 Key Wrap<br>A256KW - AES-256 Key Wrap<br>A128GCMKW - AES-GCM-128 key encryption<br>A192GCMKW - AES-GCM-192 key encryption<br>A256GCMKW - AES-GCM-256 key encryption<br>ECDH_ES - ECDH-ES<br>ECDH_ES_A128KW - ECDH-ES with AES-128 Key Wrap<br>ECDH_ES_A192KW - ECDH-ES with AES-192 Key Wrap<br>ECDH_ES_A256KW - ECDH-ES with AES-256 Key Wrap<br>RSA_OAEP - RSAES OAEP<br>RSA_OAEP_256 - RSAES OAEP using SHA-256 and MGF1 with SHA-256
	IdTokenEncryptionAlgorithm *string `json:"idTokenEncryptionAlgorithm,omitempty" tfsdk:"id_token_encryption_algorithm"`
	// The JSON Web Encryption [JWE] content encryption algorithm for the ID Token.<br>AES_128_CBC_HMAC_SHA_256 - Composite AES-CBC-128 HMAC-SHA-256<br>AES_192_CBC_HMAC_SHA_384 - Composite AES-CBC-192 HMAC-SHA-384<br>AES_256_CBC_HMAC_SHA_512 - Composite AES-CBC-256 HMAC-SHA-512<br>AES_128_GCM - AES-GCM-128<br>AES_192_GCM - AES-GCM-192<br>AES_256_GCM - AES-GCM-256
	IdTokenContentEncryptionAlgorithm *string       `json:"idTokenContentEncryptionAlgorithm,omitempty" tfsdk:"id_token_content_encryption_algorithm"`
	PolicyGroup                       *ResourceLink `json:"policyGroup,omitempty" tfsdk:"policy_group"`
}

// NewClientRegistrationOIDCPolicy instantiates a new ClientRegistrationOIDCPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientRegistrationOIDCPolicy() *ClientRegistrationOIDCPolicy {
	this := ClientRegistrationOIDCPolicy{}
	return &this
}

// NewClientRegistrationOIDCPolicyWithDefaults instantiates a new ClientRegistrationOIDCPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientRegistrationOIDCPolicyWithDefaults() *ClientRegistrationOIDCPolicy {
	this := ClientRegistrationOIDCPolicy{}
	return &this
}

// GetIdTokenSigningAlgorithm returns the IdTokenSigningAlgorithm field value if set, zero value otherwise.
func (o *ClientRegistrationOIDCPolicy) GetIdTokenSigningAlgorithm() string {
	if o == nil || IsNil(o.IdTokenSigningAlgorithm) {
		var ret string
		return ret
	}
	return *o.IdTokenSigningAlgorithm
}

// GetIdTokenSigningAlgorithmOk returns a tuple with the IdTokenSigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationOIDCPolicy) GetIdTokenSigningAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.IdTokenSigningAlgorithm) {
		return nil, false
	}
	return o.IdTokenSigningAlgorithm, true
}

// HasIdTokenSigningAlgorithm returns a boolean if a field has been set.
func (o *ClientRegistrationOIDCPolicy) HasIdTokenSigningAlgorithm() bool {
	if o != nil && !IsNil(o.IdTokenSigningAlgorithm) {
		return true
	}

	return false
}

// SetIdTokenSigningAlgorithm gets a reference to the given string and assigns it to the IdTokenSigningAlgorithm field.
func (o *ClientRegistrationOIDCPolicy) SetIdTokenSigningAlgorithm(v string) {
	o.IdTokenSigningAlgorithm = &v
}

// GetIdTokenEncryptionAlgorithm returns the IdTokenEncryptionAlgorithm field value if set, zero value otherwise.
func (o *ClientRegistrationOIDCPolicy) GetIdTokenEncryptionAlgorithm() string {
	if o == nil || IsNil(o.IdTokenEncryptionAlgorithm) {
		var ret string
		return ret
	}
	return *o.IdTokenEncryptionAlgorithm
}

// GetIdTokenEncryptionAlgorithmOk returns a tuple with the IdTokenEncryptionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationOIDCPolicy) GetIdTokenEncryptionAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.IdTokenEncryptionAlgorithm) {
		return nil, false
	}
	return o.IdTokenEncryptionAlgorithm, true
}

// HasIdTokenEncryptionAlgorithm returns a boolean if a field has been set.
func (o *ClientRegistrationOIDCPolicy) HasIdTokenEncryptionAlgorithm() bool {
	if o != nil && !IsNil(o.IdTokenEncryptionAlgorithm) {
		return true
	}

	return false
}

// SetIdTokenEncryptionAlgorithm gets a reference to the given string and assigns it to the IdTokenEncryptionAlgorithm field.
func (o *ClientRegistrationOIDCPolicy) SetIdTokenEncryptionAlgorithm(v string) {
	o.IdTokenEncryptionAlgorithm = &v
}

// GetIdTokenContentEncryptionAlgorithm returns the IdTokenContentEncryptionAlgorithm field value if set, zero value otherwise.
func (o *ClientRegistrationOIDCPolicy) GetIdTokenContentEncryptionAlgorithm() string {
	if o == nil || IsNil(o.IdTokenContentEncryptionAlgorithm) {
		var ret string
		return ret
	}
	return *o.IdTokenContentEncryptionAlgorithm
}

// GetIdTokenContentEncryptionAlgorithmOk returns a tuple with the IdTokenContentEncryptionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationOIDCPolicy) GetIdTokenContentEncryptionAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.IdTokenContentEncryptionAlgorithm) {
		return nil, false
	}
	return o.IdTokenContentEncryptionAlgorithm, true
}

// HasIdTokenContentEncryptionAlgorithm returns a boolean if a field has been set.
func (o *ClientRegistrationOIDCPolicy) HasIdTokenContentEncryptionAlgorithm() bool {
	if o != nil && !IsNil(o.IdTokenContentEncryptionAlgorithm) {
		return true
	}

	return false
}

// SetIdTokenContentEncryptionAlgorithm gets a reference to the given string and assigns it to the IdTokenContentEncryptionAlgorithm field.
func (o *ClientRegistrationOIDCPolicy) SetIdTokenContentEncryptionAlgorithm(v string) {
	o.IdTokenContentEncryptionAlgorithm = &v
}

// GetPolicyGroup returns the PolicyGroup field value if set, zero value otherwise.
func (o *ClientRegistrationOIDCPolicy) GetPolicyGroup() ResourceLink {
	if o == nil || IsNil(o.PolicyGroup) {
		var ret ResourceLink
		return ret
	}
	return *o.PolicyGroup
}

// GetPolicyGroupOk returns a tuple with the PolicyGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientRegistrationOIDCPolicy) GetPolicyGroupOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.PolicyGroup) {
		return nil, false
	}
	return o.PolicyGroup, true
}

// HasPolicyGroup returns a boolean if a field has been set.
func (o *ClientRegistrationOIDCPolicy) HasPolicyGroup() bool {
	if o != nil && !IsNil(o.PolicyGroup) {
		return true
	}

	return false
}

// SetPolicyGroup gets a reference to the given ResourceLink and assigns it to the PolicyGroup field.
func (o *ClientRegistrationOIDCPolicy) SetPolicyGroup(v ResourceLink) {
	o.PolicyGroup = &v
}

func (o ClientRegistrationOIDCPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientRegistrationOIDCPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdTokenSigningAlgorithm) {
		toSerialize["idTokenSigningAlgorithm"] = o.IdTokenSigningAlgorithm
	}
	if !IsNil(o.IdTokenEncryptionAlgorithm) {
		toSerialize["idTokenEncryptionAlgorithm"] = o.IdTokenEncryptionAlgorithm
	}
	if !IsNil(o.IdTokenContentEncryptionAlgorithm) {
		toSerialize["idTokenContentEncryptionAlgorithm"] = o.IdTokenContentEncryptionAlgorithm
	}
	if !IsNil(o.PolicyGroup) {
		toSerialize["policyGroup"] = o.PolicyGroup
	}
	return toSerialize, nil
}

type NullableClientRegistrationOIDCPolicy struct {
	value *ClientRegistrationOIDCPolicy
	isSet bool
}

func (v NullableClientRegistrationOIDCPolicy) Get() *ClientRegistrationOIDCPolicy {
	return v.value
}

func (v *NullableClientRegistrationOIDCPolicy) Set(val *ClientRegistrationOIDCPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableClientRegistrationOIDCPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableClientRegistrationOIDCPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientRegistrationOIDCPolicy(val *ClientRegistrationOIDCPolicy) *NullableClientRegistrationOIDCPolicy {
	return &NullableClientRegistrationOIDCPolicy{value: val, isSet: true}
}

func (v NullableClientRegistrationOIDCPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientRegistrationOIDCPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
