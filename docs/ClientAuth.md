# ClientAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Client authentication type.&lt;br&gt;The required field for type SECRET is secret.&lt;br&gt;The required fields for type CERTIFICATE are clientCertIssuerDn and clientCertSubjectDn.&lt;br&gt;The required field for type PRIVATE_KEY_JWT is: either jwks or jwksUrl. | [optional] 
**Secret** | Pointer to **string** | Client secret for Basic Authentication.  To update the client secret, specify the plaintext value in this field.  This field will not be populated for GET requests. | [optional] 
**EncryptedSecret** | Pointer to **string** | For GET requests, this field contains the encrypted client secret, if one exists.  For POST and PUT requests, if you wish to reuse the existing secret, this field should be passed back unchanged. | [optional] 
**SecondarySecrets** | Pointer to [**[]SecondarySecret**](SecondarySecret.md) | The list of secondary client secrets that are temporarily retained. | [optional] 
**ClientCertIssuerDn** | Pointer to **string** | Client TLS Certificate Issuer DN. | [optional] 
**ClientCertSubjectDn** | Pointer to **string** | Client TLS Certificate Subject DN. | [optional] 
**EnforceReplayPrevention** | Pointer to **bool** | Enforce replay prevention on JSON Web Tokens. This field is applicable only for Private Key JWT Client Authentication. | [optional] 
**TokenEndpointAuthSigningAlgorithm** | Pointer to **string** | The JSON Web Signature [JWS] algorithm that must be used to sign the JSON Web Tokens. This field is applicable only for Private Key JWT Client Authentication. All signing algorithms are allowed if value is not present &lt;br&gt;RS256 - RSA using SHA-256&lt;br&gt;RS384 - RSA using SHA-384&lt;br&gt;RS512 - RSA using SHA-512&lt;br&gt;ES256 - ECDSA using P256 Curve and SHA-256&lt;br&gt;ES384 - ECDSA using P384 Curve and SHA-384&lt;br&gt;ES512 - ECDSA using P521 Curve and SHA-512&lt;br&gt;PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256&lt;br&gt;PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384&lt;br&gt;PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512&lt;br&gt;RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11. | [optional] 

## Methods

### NewClientAuth

`func NewClientAuth() *ClientAuth`

NewClientAuth instantiates a new ClientAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAuthWithDefaults

`func NewClientAuthWithDefaults() *ClientAuth`

NewClientAuthWithDefaults instantiates a new ClientAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClientAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientAuth) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientAuth) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSecret

`func (o *ClientAuth) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ClientAuth) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ClientAuth) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ClientAuth) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetEncryptedSecret

`func (o *ClientAuth) GetEncryptedSecret() string`

GetEncryptedSecret returns the EncryptedSecret field if non-nil, zero value otherwise.

### GetEncryptedSecretOk

`func (o *ClientAuth) GetEncryptedSecretOk() (*string, bool)`

GetEncryptedSecretOk returns a tuple with the EncryptedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedSecret

`func (o *ClientAuth) SetEncryptedSecret(v string)`

SetEncryptedSecret sets EncryptedSecret field to given value.

### HasEncryptedSecret

`func (o *ClientAuth) HasEncryptedSecret() bool`

HasEncryptedSecret returns a boolean if a field has been set.

### GetSecondarySecrets

`func (o *ClientAuth) GetSecondarySecrets() []SecondarySecret`

GetSecondarySecrets returns the SecondarySecrets field if non-nil, zero value otherwise.

### GetSecondarySecretsOk

`func (o *ClientAuth) GetSecondarySecretsOk() (*[]SecondarySecret, bool)`

GetSecondarySecretsOk returns a tuple with the SecondarySecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySecrets

`func (o *ClientAuth) SetSecondarySecrets(v []SecondarySecret)`

SetSecondarySecrets sets SecondarySecrets field to given value.

### HasSecondarySecrets

`func (o *ClientAuth) HasSecondarySecrets() bool`

HasSecondarySecrets returns a boolean if a field has been set.

### GetClientCertIssuerDn

`func (o *ClientAuth) GetClientCertIssuerDn() string`

GetClientCertIssuerDn returns the ClientCertIssuerDn field if non-nil, zero value otherwise.

### GetClientCertIssuerDnOk

`func (o *ClientAuth) GetClientCertIssuerDnOk() (*string, bool)`

GetClientCertIssuerDnOk returns a tuple with the ClientCertIssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertIssuerDn

`func (o *ClientAuth) SetClientCertIssuerDn(v string)`

SetClientCertIssuerDn sets ClientCertIssuerDn field to given value.

### HasClientCertIssuerDn

`func (o *ClientAuth) HasClientCertIssuerDn() bool`

HasClientCertIssuerDn returns a boolean if a field has been set.

### GetClientCertSubjectDn

`func (o *ClientAuth) GetClientCertSubjectDn() string`

GetClientCertSubjectDn returns the ClientCertSubjectDn field if non-nil, zero value otherwise.

### GetClientCertSubjectDnOk

`func (o *ClientAuth) GetClientCertSubjectDnOk() (*string, bool)`

GetClientCertSubjectDnOk returns a tuple with the ClientCertSubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertSubjectDn

`func (o *ClientAuth) SetClientCertSubjectDn(v string)`

SetClientCertSubjectDn sets ClientCertSubjectDn field to given value.

### HasClientCertSubjectDn

`func (o *ClientAuth) HasClientCertSubjectDn() bool`

HasClientCertSubjectDn returns a boolean if a field has been set.

### GetEnforceReplayPrevention

`func (o *ClientAuth) GetEnforceReplayPrevention() bool`

GetEnforceReplayPrevention returns the EnforceReplayPrevention field if non-nil, zero value otherwise.

### GetEnforceReplayPreventionOk

`func (o *ClientAuth) GetEnforceReplayPreventionOk() (*bool, bool)`

GetEnforceReplayPreventionOk returns a tuple with the EnforceReplayPrevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceReplayPrevention

`func (o *ClientAuth) SetEnforceReplayPrevention(v bool)`

SetEnforceReplayPrevention sets EnforceReplayPrevention field to given value.

### HasEnforceReplayPrevention

`func (o *ClientAuth) HasEnforceReplayPrevention() bool`

HasEnforceReplayPrevention returns a boolean if a field has been set.

### GetTokenEndpointAuthSigningAlgorithm

`func (o *ClientAuth) GetTokenEndpointAuthSigningAlgorithm() string`

GetTokenEndpointAuthSigningAlgorithm returns the TokenEndpointAuthSigningAlgorithm field if non-nil, zero value otherwise.

### GetTokenEndpointAuthSigningAlgorithmOk

`func (o *ClientAuth) GetTokenEndpointAuthSigningAlgorithmOk() (*string, bool)`

GetTokenEndpointAuthSigningAlgorithmOk returns a tuple with the TokenEndpointAuthSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthSigningAlgorithm

`func (o *ClientAuth) SetTokenEndpointAuthSigningAlgorithm(v string)`

SetTokenEndpointAuthSigningAlgorithm sets TokenEndpointAuthSigningAlgorithm field to given value.

### HasTokenEndpointAuthSigningAlgorithm

`func (o *ClientAuth) HasTokenEndpointAuthSigningAlgorithm() bool`

HasTokenEndpointAuthSigningAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


