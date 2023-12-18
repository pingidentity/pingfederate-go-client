# ClientOIDCPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdTokenSigningAlgorithm** | Pointer to **string** | The JSON Web Signature [JWS] algorithm required for the ID Token.&lt;br&gt;NONE - No signing algorithm&lt;br&gt;HS256 - HMAC using SHA-256&lt;br&gt;HS384 - HMAC using SHA-384&lt;br&gt;HS512 - HMAC using SHA-512&lt;br&gt;RS256 - RSA using SHA-256&lt;br&gt;RS384 - RSA using SHA-384&lt;br&gt;RS512 - RSA using SHA-512&lt;br&gt;ES256 - ECDSA using P256 Curve and SHA-256&lt;br&gt;ES384 - ECDSA using P384 Curve and SHA-384&lt;br&gt;ES512 - ECDSA using P521 Curve and SHA-512&lt;br&gt;PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256&lt;br&gt;PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384&lt;br&gt;PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512&lt;br&gt;A null value will represent the default algorithm which is RS256.&lt;br&gt;RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11 | [optional] 
**IdTokenEncryptionAlgorithm** | Pointer to **string** | The JSON Web Encryption [JWE] encryption algorithm used to encrypt the content encryption key for the ID Token.&lt;br&gt;DIR - Direct Encryption with symmetric key&lt;br&gt;A128KW - AES-128 Key Wrap&lt;br&gt;A192KW - AES-192 Key Wrap&lt;br&gt;A256KW - AES-256 Key Wrap&lt;br&gt;A128GCMKW - AES-GCM-128 key encryption&lt;br&gt;A192GCMKW - AES-GCM-192 key encryption&lt;br&gt;A256GCMKW - AES-GCM-256 key encryption&lt;br&gt;ECDH_ES - ECDH-ES&lt;br&gt;ECDH_ES_A128KW - ECDH-ES with AES-128 Key Wrap&lt;br&gt;ECDH_ES_A192KW - ECDH-ES with AES-192 Key Wrap&lt;br&gt;ECDH_ES_A256KW - ECDH-ES with AES-256 Key Wrap&lt;br&gt;RSA_OAEP - RSAES OAEP&lt;br&gt;RSA_OAEP_256 - RSAES OAEP using SHA-256 and MGF1 with SHA-256 | [optional] 
**IdTokenContentEncryptionAlgorithm** | Pointer to **string** | The JSON Web Encryption [JWE] content encryption algorithm for the ID Token.&lt;br&gt;AES_128_CBC_HMAC_SHA_256 - Composite AES-CBC-128 HMAC-SHA-256&lt;br&gt;AES_192_CBC_HMAC_SHA_384 - Composite AES-CBC-192 HMAC-SHA-384&lt;br&gt;AES_256_CBC_HMAC_SHA_512 - Composite AES-CBC-256 HMAC-SHA-512&lt;br&gt;AES_128_GCM - AES-GCM-128&lt;br&gt;AES_192_GCM - AES-GCM-192&lt;br&gt;AES_256_GCM - AES-GCM-256 | [optional] 
**PolicyGroup** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**GrantAccessSessionRevocationApi** | Pointer to **bool** | Determines whether this client is allowed to access the Session Revocation API. | [optional] 
**GrantAccessSessionSessionManagementApi** | Pointer to **bool** | Determines whether this client is allowed to access the Session Management API. | [optional] 
**LogoutMode** | Pointer to **string** | The logout mode for this client. The default is &#39;NONE&#39;. | [optional] 
**PingAccessLogoutCapable** | Pointer to **bool** | Set this value to true if you wish to enable client application logout, and the client is PingAccess, or its logout endpoints follow the PingAccess path convention. | [optional] 
**LogoutUris** | Pointer to **[]string** | A list of front-channel logout URIs for this client. | [optional] 
**BackChannelLogoutUri** | Pointer to **string** | The back-channel logout URI for this client. | [optional] 
**PairwiseIdentifierUserType** | Pointer to **bool** | Determines whether the subject identifier type is pairwise. | [optional] 
**SectorIdentifierUri** | Pointer to **string** | The URI references a file with a single JSON array of Redirect URI and JWKS URL values. | [optional] 

## Methods

### NewClientOIDCPolicy

`func NewClientOIDCPolicy() *ClientOIDCPolicy`

NewClientOIDCPolicy instantiates a new ClientOIDCPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientOIDCPolicyWithDefaults

`func NewClientOIDCPolicyWithDefaults() *ClientOIDCPolicy`

NewClientOIDCPolicyWithDefaults instantiates a new ClientOIDCPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdTokenSigningAlgorithm

`func (o *ClientOIDCPolicy) GetIdTokenSigningAlgorithm() string`

GetIdTokenSigningAlgorithm returns the IdTokenSigningAlgorithm field if non-nil, zero value otherwise.

### GetIdTokenSigningAlgorithmOk

`func (o *ClientOIDCPolicy) GetIdTokenSigningAlgorithmOk() (*string, bool)`

GetIdTokenSigningAlgorithmOk returns a tuple with the IdTokenSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenSigningAlgorithm

`func (o *ClientOIDCPolicy) SetIdTokenSigningAlgorithm(v string)`

SetIdTokenSigningAlgorithm sets IdTokenSigningAlgorithm field to given value.

### HasIdTokenSigningAlgorithm

`func (o *ClientOIDCPolicy) HasIdTokenSigningAlgorithm() bool`

HasIdTokenSigningAlgorithm returns a boolean if a field has been set.

### GetIdTokenEncryptionAlgorithm

`func (o *ClientOIDCPolicy) GetIdTokenEncryptionAlgorithm() string`

GetIdTokenEncryptionAlgorithm returns the IdTokenEncryptionAlgorithm field if non-nil, zero value otherwise.

### GetIdTokenEncryptionAlgorithmOk

`func (o *ClientOIDCPolicy) GetIdTokenEncryptionAlgorithmOk() (*string, bool)`

GetIdTokenEncryptionAlgorithmOk returns a tuple with the IdTokenEncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenEncryptionAlgorithm

`func (o *ClientOIDCPolicy) SetIdTokenEncryptionAlgorithm(v string)`

SetIdTokenEncryptionAlgorithm sets IdTokenEncryptionAlgorithm field to given value.

### HasIdTokenEncryptionAlgorithm

`func (o *ClientOIDCPolicy) HasIdTokenEncryptionAlgorithm() bool`

HasIdTokenEncryptionAlgorithm returns a boolean if a field has been set.

### GetIdTokenContentEncryptionAlgorithm

`func (o *ClientOIDCPolicy) GetIdTokenContentEncryptionAlgorithm() string`

GetIdTokenContentEncryptionAlgorithm returns the IdTokenContentEncryptionAlgorithm field if non-nil, zero value otherwise.

### GetIdTokenContentEncryptionAlgorithmOk

`func (o *ClientOIDCPolicy) GetIdTokenContentEncryptionAlgorithmOk() (*string, bool)`

GetIdTokenContentEncryptionAlgorithmOk returns a tuple with the IdTokenContentEncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenContentEncryptionAlgorithm

`func (o *ClientOIDCPolicy) SetIdTokenContentEncryptionAlgorithm(v string)`

SetIdTokenContentEncryptionAlgorithm sets IdTokenContentEncryptionAlgorithm field to given value.

### HasIdTokenContentEncryptionAlgorithm

`func (o *ClientOIDCPolicy) HasIdTokenContentEncryptionAlgorithm() bool`

HasIdTokenContentEncryptionAlgorithm returns a boolean if a field has been set.

### GetPolicyGroup

`func (o *ClientOIDCPolicy) GetPolicyGroup() ResourceLink`

GetPolicyGroup returns the PolicyGroup field if non-nil, zero value otherwise.

### GetPolicyGroupOk

`func (o *ClientOIDCPolicy) GetPolicyGroupOk() (*ResourceLink, bool)`

GetPolicyGroupOk returns a tuple with the PolicyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyGroup

`func (o *ClientOIDCPolicy) SetPolicyGroup(v ResourceLink)`

SetPolicyGroup sets PolicyGroup field to given value.

### HasPolicyGroup

`func (o *ClientOIDCPolicy) HasPolicyGroup() bool`

HasPolicyGroup returns a boolean if a field has been set.

### GetGrantAccessSessionRevocationApi

`func (o *ClientOIDCPolicy) GetGrantAccessSessionRevocationApi() bool`

GetGrantAccessSessionRevocationApi returns the GrantAccessSessionRevocationApi field if non-nil, zero value otherwise.

### GetGrantAccessSessionRevocationApiOk

`func (o *ClientOIDCPolicy) GetGrantAccessSessionRevocationApiOk() (*bool, bool)`

GetGrantAccessSessionRevocationApiOk returns a tuple with the GrantAccessSessionRevocationApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantAccessSessionRevocationApi

`func (o *ClientOIDCPolicy) SetGrantAccessSessionRevocationApi(v bool)`

SetGrantAccessSessionRevocationApi sets GrantAccessSessionRevocationApi field to given value.

### HasGrantAccessSessionRevocationApi

`func (o *ClientOIDCPolicy) HasGrantAccessSessionRevocationApi() bool`

HasGrantAccessSessionRevocationApi returns a boolean if a field has been set.

### GetGrantAccessSessionSessionManagementApi

`func (o *ClientOIDCPolicy) GetGrantAccessSessionSessionManagementApi() bool`

GetGrantAccessSessionSessionManagementApi returns the GrantAccessSessionSessionManagementApi field if non-nil, zero value otherwise.

### GetGrantAccessSessionSessionManagementApiOk

`func (o *ClientOIDCPolicy) GetGrantAccessSessionSessionManagementApiOk() (*bool, bool)`

GetGrantAccessSessionSessionManagementApiOk returns a tuple with the GrantAccessSessionSessionManagementApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantAccessSessionSessionManagementApi

`func (o *ClientOIDCPolicy) SetGrantAccessSessionSessionManagementApi(v bool)`

SetGrantAccessSessionSessionManagementApi sets GrantAccessSessionSessionManagementApi field to given value.

### HasGrantAccessSessionSessionManagementApi

`func (o *ClientOIDCPolicy) HasGrantAccessSessionSessionManagementApi() bool`

HasGrantAccessSessionSessionManagementApi returns a boolean if a field has been set.

### GetLogoutMode

`func (o *ClientOIDCPolicy) GetLogoutMode() string`

GetLogoutMode returns the LogoutMode field if non-nil, zero value otherwise.

### GetLogoutModeOk

`func (o *ClientOIDCPolicy) GetLogoutModeOk() (*string, bool)`

GetLogoutModeOk returns a tuple with the LogoutMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutMode

`func (o *ClientOIDCPolicy) SetLogoutMode(v string)`

SetLogoutMode sets LogoutMode field to given value.

### HasLogoutMode

`func (o *ClientOIDCPolicy) HasLogoutMode() bool`

HasLogoutMode returns a boolean if a field has been set.

### GetPingAccessLogoutCapable

`func (o *ClientOIDCPolicy) GetPingAccessLogoutCapable() bool`

GetPingAccessLogoutCapable returns the PingAccessLogoutCapable field if non-nil, zero value otherwise.

### GetPingAccessLogoutCapableOk

`func (o *ClientOIDCPolicy) GetPingAccessLogoutCapableOk() (*bool, bool)`

GetPingAccessLogoutCapableOk returns a tuple with the PingAccessLogoutCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingAccessLogoutCapable

`func (o *ClientOIDCPolicy) SetPingAccessLogoutCapable(v bool)`

SetPingAccessLogoutCapable sets PingAccessLogoutCapable field to given value.

### HasPingAccessLogoutCapable

`func (o *ClientOIDCPolicy) HasPingAccessLogoutCapable() bool`

HasPingAccessLogoutCapable returns a boolean if a field has been set.

### GetLogoutUris

`func (o *ClientOIDCPolicy) GetLogoutUris() []string`

GetLogoutUris returns the LogoutUris field if non-nil, zero value otherwise.

### GetLogoutUrisOk

`func (o *ClientOIDCPolicy) GetLogoutUrisOk() (*[]string, bool)`

GetLogoutUrisOk returns a tuple with the LogoutUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUris

`func (o *ClientOIDCPolicy) SetLogoutUris(v []string)`

SetLogoutUris sets LogoutUris field to given value.

### HasLogoutUris

`func (o *ClientOIDCPolicy) HasLogoutUris() bool`

HasLogoutUris returns a boolean if a field has been set.

### GetBackChannelLogoutUri

`func (o *ClientOIDCPolicy) GetBackChannelLogoutUri() string`

GetBackChannelLogoutUri returns the BackChannelLogoutUri field if non-nil, zero value otherwise.

### GetBackChannelLogoutUriOk

`func (o *ClientOIDCPolicy) GetBackChannelLogoutUriOk() (*string, bool)`

GetBackChannelLogoutUriOk returns a tuple with the BackChannelLogoutUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackChannelLogoutUri

`func (o *ClientOIDCPolicy) SetBackChannelLogoutUri(v string)`

SetBackChannelLogoutUri sets BackChannelLogoutUri field to given value.

### HasBackChannelLogoutUri

`func (o *ClientOIDCPolicy) HasBackChannelLogoutUri() bool`

HasBackChannelLogoutUri returns a boolean if a field has been set.

### GetPairwiseIdentifierUserType

`func (o *ClientOIDCPolicy) GetPairwiseIdentifierUserType() bool`

GetPairwiseIdentifierUserType returns the PairwiseIdentifierUserType field if non-nil, zero value otherwise.

### GetPairwiseIdentifierUserTypeOk

`func (o *ClientOIDCPolicy) GetPairwiseIdentifierUserTypeOk() (*bool, bool)`

GetPairwiseIdentifierUserTypeOk returns a tuple with the PairwiseIdentifierUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairwiseIdentifierUserType

`func (o *ClientOIDCPolicy) SetPairwiseIdentifierUserType(v bool)`

SetPairwiseIdentifierUserType sets PairwiseIdentifierUserType field to given value.

### HasPairwiseIdentifierUserType

`func (o *ClientOIDCPolicy) HasPairwiseIdentifierUserType() bool`

HasPairwiseIdentifierUserType returns a boolean if a field has been set.

### GetSectorIdentifierUri

`func (o *ClientOIDCPolicy) GetSectorIdentifierUri() string`

GetSectorIdentifierUri returns the SectorIdentifierUri field if non-nil, zero value otherwise.

### GetSectorIdentifierUriOk

`func (o *ClientOIDCPolicy) GetSectorIdentifierUriOk() (*string, bool)`

GetSectorIdentifierUriOk returns a tuple with the SectorIdentifierUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorIdentifierUri

`func (o *ClientOIDCPolicy) SetSectorIdentifierUri(v string)`

SetSectorIdentifierUri sets SectorIdentifierUri field to given value.

### HasSectorIdentifierUri

`func (o *ClientOIDCPolicy) HasSectorIdentifierUri() bool`

HasSectorIdentifierUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


