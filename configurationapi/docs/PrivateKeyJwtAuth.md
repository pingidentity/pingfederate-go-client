# PrivateKeyJwtAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies whether private key JWT authentication is enabled. | [optional] 
**EnforceReplayPrevention** | Pointer to **bool** | Enforce replay prevention on JSON Web Tokens. | [optional] 
**TokenEndpointAuthSigningAlgorithm** | Pointer to **string** | The JSON Web Signature [JWS] algorithm that must be used to sign the JSON Web Tokens. All asymmetric signing algorithms are allowed if value is not present.&lt;br&gt;RS256 - RSA using SHA-256&lt;br&gt;RS384 - RSA using SHA-384&lt;br&gt;RS512 - RSA using SHA-512&lt;br&gt;ES256 - ECDSA using P256 Curve and SHA-256&lt;br&gt;ES384 - ECDSA using P384 Curve and SHA-384&lt;br&gt;ES512 - ECDSA using P521 Curve and SHA-512&lt;br&gt;PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256&lt;br&gt;PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384&lt;br&gt;PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512&lt;br&gt;RSASSA-PSS is only supported with Thales Luna, Entrust nShield Connect or Java 11. | [optional] 

## Methods

### NewPrivateKeyJwtAuth

`func NewPrivateKeyJwtAuth() *PrivateKeyJwtAuth`

NewPrivateKeyJwtAuth instantiates a new PrivateKeyJwtAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateKeyJwtAuthWithDefaults

`func NewPrivateKeyJwtAuthWithDefaults() *PrivateKeyJwtAuth`

NewPrivateKeyJwtAuthWithDefaults instantiates a new PrivateKeyJwtAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PrivateKeyJwtAuth) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PrivateKeyJwtAuth) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PrivateKeyJwtAuth) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PrivateKeyJwtAuth) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnforceReplayPrevention

`func (o *PrivateKeyJwtAuth) GetEnforceReplayPrevention() bool`

GetEnforceReplayPrevention returns the EnforceReplayPrevention field if non-nil, zero value otherwise.

### GetEnforceReplayPreventionOk

`func (o *PrivateKeyJwtAuth) GetEnforceReplayPreventionOk() (*bool, bool)`

GetEnforceReplayPreventionOk returns a tuple with the EnforceReplayPrevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceReplayPrevention

`func (o *PrivateKeyJwtAuth) SetEnforceReplayPrevention(v bool)`

SetEnforceReplayPrevention sets EnforceReplayPrevention field to given value.

### HasEnforceReplayPrevention

`func (o *PrivateKeyJwtAuth) HasEnforceReplayPrevention() bool`

HasEnforceReplayPrevention returns a boolean if a field has been set.

### GetTokenEndpointAuthSigningAlgorithm

`func (o *PrivateKeyJwtAuth) GetTokenEndpointAuthSigningAlgorithm() string`

GetTokenEndpointAuthSigningAlgorithm returns the TokenEndpointAuthSigningAlgorithm field if non-nil, zero value otherwise.

### GetTokenEndpointAuthSigningAlgorithmOk

`func (o *PrivateKeyJwtAuth) GetTokenEndpointAuthSigningAlgorithmOk() (*string, bool)`

GetTokenEndpointAuthSigningAlgorithmOk returns a tuple with the TokenEndpointAuthSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthSigningAlgorithm

`func (o *PrivateKeyJwtAuth) SetTokenEndpointAuthSigningAlgorithm(v string)`

SetTokenEndpointAuthSigningAlgorithm sets TokenEndpointAuthSigningAlgorithm field to given value.

### HasTokenEndpointAuthSigningAlgorithm

`func (o *PrivateKeyJwtAuth) HasTokenEndpointAuthSigningAlgorithm() bool`

HasTokenEndpointAuthSigningAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


