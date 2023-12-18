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

// checks if the ClientOIDCPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientOIDCPolicy{}

// ClientOIDCPolicy OAuth Client Open ID Connect Policy.
type ClientOIDCPolicy struct {
	// The JSON Web Signature [JWS] algorithm required for the ID Token.<br>NONE - No signing algorithm<br>HS256 - HMAC using SHA-256<br>HS384 - HMAC using SHA-384<br>HS512 - HMAC using SHA-512<br>RS256 - RSA using SHA-256<br>RS384 - RSA using SHA-384<br>RS512 - RSA using SHA-512<br>ES256 - ECDSA using P256 Curve and SHA-256<br>ES384 - ECDSA using P384 Curve and SHA-384<br>ES512 - ECDSA using P521 Curve and SHA-512<br>PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256<br>PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384<br>PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512<br>A null value will represent the default algorithm which is RS256.<br>RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11
	IdTokenSigningAlgorithm *string `json:"idTokenSigningAlgorithm,omitempty" tfsdk:"id_token_signing_algorithm"`
	// The JSON Web Encryption [JWE] encryption algorithm used to encrypt the content encryption key for the ID Token.<br>DIR - Direct Encryption with symmetric key<br>A128KW - AES-128 Key Wrap<br>A192KW - AES-192 Key Wrap<br>A256KW - AES-256 Key Wrap<br>A128GCMKW - AES-GCM-128 key encryption<br>A192GCMKW - AES-GCM-192 key encryption<br>A256GCMKW - AES-GCM-256 key encryption<br>ECDH_ES - ECDH-ES<br>ECDH_ES_A128KW - ECDH-ES with AES-128 Key Wrap<br>ECDH_ES_A192KW - ECDH-ES with AES-192 Key Wrap<br>ECDH_ES_A256KW - ECDH-ES with AES-256 Key Wrap<br>RSA_OAEP - RSAES OAEP<br>RSA_OAEP_256 - RSAES OAEP using SHA-256 and MGF1 with SHA-256
	IdTokenEncryptionAlgorithm *string `json:"idTokenEncryptionAlgorithm,omitempty" tfsdk:"id_token_encryption_algorithm"`
	// The JSON Web Encryption [JWE] content encryption algorithm for the ID Token.<br>AES_128_CBC_HMAC_SHA_256 - Composite AES-CBC-128 HMAC-SHA-256<br>AES_192_CBC_HMAC_SHA_384 - Composite AES-CBC-192 HMAC-SHA-384<br>AES_256_CBC_HMAC_SHA_512 - Composite AES-CBC-256 HMAC-SHA-512<br>AES_128_GCM - AES-GCM-128<br>AES_192_GCM - AES-GCM-192<br>AES_256_GCM - AES-GCM-256
	IdTokenContentEncryptionAlgorithm *string       `json:"idTokenContentEncryptionAlgorithm,omitempty" tfsdk:"id_token_content_encryption_algorithm"`
	PolicyGroup                       *ResourceLink `json:"policyGroup,omitempty" tfsdk:"policy_group"`
	// Determines whether this client is allowed to access the Session Revocation API.
	GrantAccessSessionRevocationApi *bool `json:"grantAccessSessionRevocationApi,omitempty" tfsdk:"grant_access_session_revocation_api"`
	// Determines whether this client is allowed to access the Session Management API.
	GrantAccessSessionSessionManagementApi *bool `json:"grantAccessSessionSessionManagementApi,omitempty" tfsdk:"grant_access_session_session_management_api"`
	// The logout mode for this client. The default is 'NONE'.
	LogoutMode *string `json:"logoutMode,omitempty" tfsdk:"logout_mode"`
	// Set this value to true if you wish to enable client application logout, and the client is PingAccess, or its logout endpoints follow the PingAccess path convention.
	PingAccessLogoutCapable *bool `json:"pingAccessLogoutCapable,omitempty" tfsdk:"ping_access_logout_capable"`
	// A list of front-channel logout URIs for this client.
	LogoutUris []string `json:"logoutUris,omitempty" tfsdk:"logout_uris"`
	// The back-channel logout URI for this client.
	BackChannelLogoutUri *string `json:"backChannelLogoutUri,omitempty" tfsdk:"back_channel_logout_uri"`
	// Determines whether the subject identifier type is pairwise.
	PairwiseIdentifierUserType *bool `json:"pairwiseIdentifierUserType,omitempty" tfsdk:"pairwise_identifier_user_type"`
	// The URI references a file with a single JSON array of Redirect URI and JWKS URL values.
	SectorIdentifierUri *string `json:"sectorIdentifierUri,omitempty" tfsdk:"sector_identifier_uri"`
}

// NewClientOIDCPolicy instantiates a new ClientOIDCPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientOIDCPolicy() *ClientOIDCPolicy {
	this := ClientOIDCPolicy{}
	return &this
}

// NewClientOIDCPolicyWithDefaults instantiates a new ClientOIDCPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientOIDCPolicyWithDefaults() *ClientOIDCPolicy {
	this := ClientOIDCPolicy{}
	return &this
}

// GetIdTokenSigningAlgorithm returns the IdTokenSigningAlgorithm field value if set, zero value otherwise.
func (o *ClientOIDCPolicy) GetIdTokenSigningAlgorithm() string {
	if o == nil || IsNil(o.IdTokenSigningAlgorithm) {
		var ret string
		return ret
	}
	return *o.IdTokenSigningAlgorithm
}

// GetIdTokenSigningAlgorithmOk returns a tuple with the IdTokenSigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOIDCPolicy) GetIdTokenSigningAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.IdTokenSigningAlgorithm) {
		return nil, false
	}
	return o.IdTokenSigningAlgorithm, true
}

// HasIdTokenSigningAlgorithm returns a boolean if a field has been set.
func (o *ClientOIDCPolicy) HasIdTokenSigningAlgorithm() bool {
	if o != nil && !IsNil(o.IdTokenSigningAlgorithm) {
		return true
	}

	return false
}

// SetIdTokenSigningAlgorithm gets a reference to the given string and assigns it to the IdTokenSigningAlgorithm field.
func (o *ClientOIDCPolicy) SetIdTokenSigningAlgorithm(v string) {
	o.IdTokenSigningAlgorithm = &v
}

// GetIdTokenEncryptionAlgorithm returns the IdTokenEncryptionAlgorithm field value if set, zero value otherwise.
func (o *ClientOIDCPolicy) GetIdTokenEncryptionAlgorithm() string {
	if o == nil || IsNil(o.IdTokenEncryptionAlgorithm) {
		var ret string
		return ret
	}
	return *o.IdTokenEncryptionAlgorithm
}

// GetIdTokenEncryptionAlgorithmOk returns a tuple with the IdTokenEncryptionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOIDCPolicy) GetIdTokenEncryptionAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.IdTokenEncryptionAlgorithm) {
		return nil, false
	}
	return o.IdTokenEncryptionAlgorithm, true
}

// HasIdTokenEncryptionAlgorithm returns a boolean if a field has been set.
func (o *ClientOIDCPolicy) HasIdTokenEncryptionAlgorithm() bool {
	if o != nil && !IsNil(o.IdTokenEncryptionAlgorithm) {
		return true
	}

	return false
}

// SetIdTokenEncryptionAlgorithm gets a reference to the given string and assigns it to the IdTokenEncryptionAlgorithm field.
func (o *ClientOIDCPolicy) SetIdTokenEncryptionAlgorithm(v string) {
	o.IdTokenEncryptionAlgorithm = &v
}

// GetIdTokenContentEncryptionAlgorithm returns the IdTokenContentEncryptionAlgorithm field value if set, zero value otherwise.
func (o *ClientOIDCPolicy) GetIdTokenContentEncryptionAlgorithm() string {
	if o == nil || IsNil(o.IdTokenContentEncryptionAlgorithm) {
		var ret string
		return ret
	}
	return *o.IdTokenContentEncryptionAlgorithm
}

// GetIdTokenContentEncryptionAlgorithmOk returns a tuple with the IdTokenContentEncryptionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOIDCPolicy) GetIdTokenContentEncryptionAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.IdTokenContentEncryptionAlgorithm) {
		return nil, false
	}
	return o.IdTokenContentEncryptionAlgorithm, true
}

// HasIdTokenContentEncryptionAlgorithm returns a boolean if a field has been set.
func (o *ClientOIDCPolicy) HasIdTokenContentEncryptionAlgorithm() bool {
	if o != nil && !IsNil(o.IdTokenContentEncryptionAlgorithm) {
		return true
	}

	return false
}

// SetIdTokenContentEncryptionAlgorithm gets a reference to the given string and assigns it to the IdTokenContentEncryptionAlgorithm field.
func (o *ClientOIDCPolicy) SetIdTokenContentEncryptionAlgorithm(v string) {
	o.IdTokenContentEncryptionAlgorithm = &v
}

// GetPolicyGroup returns the PolicyGroup field value if set, zero value otherwise.
func (o *ClientOIDCPolicy) GetPolicyGroup() ResourceLink {
	if o == nil || IsNil(o.PolicyGroup) {
		var ret ResourceLink
		return ret
	}
	return *o.PolicyGroup
}

// GetPolicyGroupOk returns a tuple with the PolicyGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOIDCPolicy) GetPolicyGroupOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.PolicyGroup) {
		return nil, false
	}
	return o.PolicyGroup, true
}

// HasPolicyGroup returns a boolean if a field has been set.
func (o *ClientOIDCPolicy) HasPolicyGroup() bool {
	if o != nil && !IsNil(o.PolicyGroup) {
		return true
	}

	return false
}

// SetPolicyGroup gets a reference to the given ResourceLink and assigns it to the PolicyGroup field.
func (o *ClientOIDCPolicy) SetPolicyGroup(v ResourceLink) {
	o.PolicyGroup = &v
}

// GetGrantAccessSessionRevocationApi returns the GrantAccessSessionRevocationApi field value if set, zero value otherwise.
func (o *ClientOIDCPolicy) GetGrantAccessSessionRevocationApi() bool {
	if o == nil || IsNil(o.GrantAccessSessionRevocationApi) {
		var ret bool
		return ret
	}
	return *o.GrantAccessSessionRevocationApi
}

// GetGrantAccessSessionRevocationApiOk returns a tuple with the GrantAccessSessionRevocationApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOIDCPolicy) GetGrantAccessSessionRevocationApiOk() (*bool, bool) {
	if o == nil || IsNil(o.GrantAccessSessionRevocationApi) {
		return nil, false
	}
	return o.GrantAccessSessionRevocationApi, true
}

// HasGrantAccessSessionRevocationApi returns a boolean if a field has been set.
func (o *ClientOIDCPolicy) HasGrantAccessSessionRevocationApi() bool {
	if o != nil && !IsNil(o.GrantAccessSessionRevocationApi) {
		return true
	}

	return false
}

// SetGrantAccessSessionRevocationApi gets a reference to the given bool and assigns it to the GrantAccessSessionRevocationApi field.
func (o *ClientOIDCPolicy) SetGrantAccessSessionRevocationApi(v bool) {
	o.GrantAccessSessionRevocationApi = &v
}

// GetGrantAccessSessionSessionManagementApi returns the GrantAccessSessionSessionManagementApi field value if set, zero value otherwise.
func (o *ClientOIDCPolicy) GetGrantAccessSessionSessionManagementApi() bool {
	if o == nil || IsNil(o.GrantAccessSessionSessionManagementApi) {
		var ret bool
		return ret
	}
	return *o.GrantAccessSessionSessionManagementApi
}

// GetGrantAccessSessionSessionManagementApiOk returns a tuple with the GrantAccessSessionSessionManagementApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOIDCPolicy) GetGrantAccessSessionSessionManagementApiOk() (*bool, bool) {
	if o == nil || IsNil(o.GrantAccessSessionSessionManagementApi) {
		return nil, false
	}
	return o.GrantAccessSessionSessionManagementApi, true
}

// HasGrantAccessSessionSessionManagementApi returns a boolean if a field has been set.
func (o *ClientOIDCPolicy) HasGrantAccessSessionSessionManagementApi() bool {
	if o != nil && !IsNil(o.GrantAccessSessionSessionManagementApi) {
		return true
	}

	return false
}

// SetGrantAccessSessionSessionManagementApi gets a reference to the given bool and assigns it to the GrantAccessSessionSessionManagementApi field.
func (o *ClientOIDCPolicy) SetGrantAccessSessionSessionManagementApi(v bool) {
	o.GrantAccessSessionSessionManagementApi = &v
}

// GetLogoutMode returns the LogoutMode field value if set, zero value otherwise.
func (o *ClientOIDCPolicy) GetLogoutMode() string {
	if o == nil || IsNil(o.LogoutMode) {
		var ret string
		return ret
	}
	return *o.LogoutMode
}

// GetLogoutModeOk returns a tuple with the LogoutMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOIDCPolicy) GetLogoutModeOk() (*string, bool) {
	if o == nil || IsNil(o.LogoutMode) {
		return nil, false
	}
	return o.LogoutMode, true
}

// HasLogoutMode returns a boolean if a field has been set.
func (o *ClientOIDCPolicy) HasLogoutMode() bool {
	if o != nil && !IsNil(o.LogoutMode) {
		return true
	}

	return false
}

// SetLogoutMode gets a reference to the given string and assigns it to the LogoutMode field.
func (o *ClientOIDCPolicy) SetLogoutMode(v string) {
	o.LogoutMode = &v
}

// GetPingAccessLogoutCapable returns the PingAccessLogoutCapable field value if set, zero value otherwise.
func (o *ClientOIDCPolicy) GetPingAccessLogoutCapable() bool {
	if o == nil || IsNil(o.PingAccessLogoutCapable) {
		var ret bool
		return ret
	}
	return *o.PingAccessLogoutCapable
}

// GetPingAccessLogoutCapableOk returns a tuple with the PingAccessLogoutCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOIDCPolicy) GetPingAccessLogoutCapableOk() (*bool, bool) {
	if o == nil || IsNil(o.PingAccessLogoutCapable) {
		return nil, false
	}
	return o.PingAccessLogoutCapable, true
}

// HasPingAccessLogoutCapable returns a boolean if a field has been set.
func (o *ClientOIDCPolicy) HasPingAccessLogoutCapable() bool {
	if o != nil && !IsNil(o.PingAccessLogoutCapable) {
		return true
	}

	return false
}

// SetPingAccessLogoutCapable gets a reference to the given bool and assigns it to the PingAccessLogoutCapable field.
func (o *ClientOIDCPolicy) SetPingAccessLogoutCapable(v bool) {
	o.PingAccessLogoutCapable = &v
}

// GetLogoutUris returns the LogoutUris field value if set, zero value otherwise.
func (o *ClientOIDCPolicy) GetLogoutUris() []string {
	if o == nil || IsNil(o.LogoutUris) {
		var ret []string
		return ret
	}
	return o.LogoutUris
}

// GetLogoutUrisOk returns a tuple with the LogoutUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOIDCPolicy) GetLogoutUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.LogoutUris) {
		return nil, false
	}
	return o.LogoutUris, true
}

// HasLogoutUris returns a boolean if a field has been set.
func (o *ClientOIDCPolicy) HasLogoutUris() bool {
	if o != nil && !IsNil(o.LogoutUris) {
		return true
	}

	return false
}

// SetLogoutUris gets a reference to the given []string and assigns it to the LogoutUris field.
func (o *ClientOIDCPolicy) SetLogoutUris(v []string) {
	o.LogoutUris = v
}

// GetBackChannelLogoutUri returns the BackChannelLogoutUri field value if set, zero value otherwise.
func (o *ClientOIDCPolicy) GetBackChannelLogoutUri() string {
	if o == nil || IsNil(o.BackChannelLogoutUri) {
		var ret string
		return ret
	}
	return *o.BackChannelLogoutUri
}

// GetBackChannelLogoutUriOk returns a tuple with the BackChannelLogoutUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOIDCPolicy) GetBackChannelLogoutUriOk() (*string, bool) {
	if o == nil || IsNil(o.BackChannelLogoutUri) {
		return nil, false
	}
	return o.BackChannelLogoutUri, true
}

// HasBackChannelLogoutUri returns a boolean if a field has been set.
func (o *ClientOIDCPolicy) HasBackChannelLogoutUri() bool {
	if o != nil && !IsNil(o.BackChannelLogoutUri) {
		return true
	}

	return false
}

// SetBackChannelLogoutUri gets a reference to the given string and assigns it to the BackChannelLogoutUri field.
func (o *ClientOIDCPolicy) SetBackChannelLogoutUri(v string) {
	o.BackChannelLogoutUri = &v
}

// GetPairwiseIdentifierUserType returns the PairwiseIdentifierUserType field value if set, zero value otherwise.
func (o *ClientOIDCPolicy) GetPairwiseIdentifierUserType() bool {
	if o == nil || IsNil(o.PairwiseIdentifierUserType) {
		var ret bool
		return ret
	}
	return *o.PairwiseIdentifierUserType
}

// GetPairwiseIdentifierUserTypeOk returns a tuple with the PairwiseIdentifierUserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOIDCPolicy) GetPairwiseIdentifierUserTypeOk() (*bool, bool) {
	if o == nil || IsNil(o.PairwiseIdentifierUserType) {
		return nil, false
	}
	return o.PairwiseIdentifierUserType, true
}

// HasPairwiseIdentifierUserType returns a boolean if a field has been set.
func (o *ClientOIDCPolicy) HasPairwiseIdentifierUserType() bool {
	if o != nil && !IsNil(o.PairwiseIdentifierUserType) {
		return true
	}

	return false
}

// SetPairwiseIdentifierUserType gets a reference to the given bool and assigns it to the PairwiseIdentifierUserType field.
func (o *ClientOIDCPolicy) SetPairwiseIdentifierUserType(v bool) {
	o.PairwiseIdentifierUserType = &v
}

// GetSectorIdentifierUri returns the SectorIdentifierUri field value if set, zero value otherwise.
func (o *ClientOIDCPolicy) GetSectorIdentifierUri() string {
	if o == nil || IsNil(o.SectorIdentifierUri) {
		var ret string
		return ret
	}
	return *o.SectorIdentifierUri
}

// GetSectorIdentifierUriOk returns a tuple with the SectorIdentifierUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientOIDCPolicy) GetSectorIdentifierUriOk() (*string, bool) {
	if o == nil || IsNil(o.SectorIdentifierUri) {
		return nil, false
	}
	return o.SectorIdentifierUri, true
}

// HasSectorIdentifierUri returns a boolean if a field has been set.
func (o *ClientOIDCPolicy) HasSectorIdentifierUri() bool {
	if o != nil && !IsNil(o.SectorIdentifierUri) {
		return true
	}

	return false
}

// SetSectorIdentifierUri gets a reference to the given string and assigns it to the SectorIdentifierUri field.
func (o *ClientOIDCPolicy) SetSectorIdentifierUri(v string) {
	o.SectorIdentifierUri = &v
}

func (o ClientOIDCPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientOIDCPolicy) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.GrantAccessSessionRevocationApi) {
		toSerialize["grantAccessSessionRevocationApi"] = o.GrantAccessSessionRevocationApi
	}
	if !IsNil(o.GrantAccessSessionSessionManagementApi) {
		toSerialize["grantAccessSessionSessionManagementApi"] = o.GrantAccessSessionSessionManagementApi
	}
	if !IsNil(o.LogoutMode) {
		toSerialize["logoutMode"] = o.LogoutMode
	}
	if !IsNil(o.PingAccessLogoutCapable) {
		toSerialize["pingAccessLogoutCapable"] = o.PingAccessLogoutCapable
	}
	if !IsNil(o.LogoutUris) {
		toSerialize["logoutUris"] = o.LogoutUris
	}
	if !IsNil(o.BackChannelLogoutUri) {
		toSerialize["backChannelLogoutUri"] = o.BackChannelLogoutUri
	}
	if !IsNil(o.PairwiseIdentifierUserType) {
		toSerialize["pairwiseIdentifierUserType"] = o.PairwiseIdentifierUserType
	}
	if !IsNil(o.SectorIdentifierUri) {
		toSerialize["sectorIdentifierUri"] = o.SectorIdentifierUri
	}
	return toSerialize, nil
}

type NullableClientOIDCPolicy struct {
	value *ClientOIDCPolicy
	isSet bool
}

func (v NullableClientOIDCPolicy) Get() *ClientOIDCPolicy {
	return v.value
}

func (v *NullableClientOIDCPolicy) Set(val *ClientOIDCPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableClientOIDCPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableClientOIDCPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientOIDCPolicy(val *ClientOIDCPolicy) *NullableClientOIDCPolicy {
	return &NullableClientOIDCPolicy{value: val, isSet: true}
}

func (v NullableClientOIDCPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientOIDCPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
