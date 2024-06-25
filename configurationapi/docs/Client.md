# Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | A unique identifier the client provides to the Resource Server to identify itself. This identifier is included with every request the client makes. For PUT requests, this field is optional and it will be overridden by the &#39;id&#39; parameter of the PUT request. | 
**Enabled** | Pointer to **bool** | Specifies whether the client is enabled. The default value is true. | [optional] 
**RedirectUris** | Pointer to **[]string** | URIs to which the OAuth AS may redirect the resource owner&#39;s user agent after authorization is obtained. A redirection URI is used with the Authorization Code and Implicit grant types. Wildcards are allowed. However, for security reasons, make the URL as restrictive as possible.For example: https://_*.company.com/_* Important: If more than one URI is added or if a single URI uses wildcards, then Authorization Code grant and token requests must contain a specific matching redirect uri parameter. | [optional] 
**GrantTypes** | **[]string** | The grant types allowed for this client. The EXTENSION grant type applies to SAML/JWT assertion grants. | 
**Name** | **string** | A descriptive name for the client instance. This name appears when the user is prompted for authorization. | 
**Description** | Pointer to **string** | A description of what the client application does. This description appears when the user is prompted for authorization. | [optional] 
**ModificationDate** | Pointer to **time.Time** | The time at which the client was last changed. This property is read only and is ignored on PUT and POST requests. | [optional] 
**CreationDate** | Pointer to **time.Time** | The time at which the client was created. This property is read only and is ignored on PUT and POST requests. | [optional] 
**ReplicationStatus** | Pointer to **string** | This status indicates whether the client has been replicated to the cluster. This property only applies when using XML client storage and automatic replication of clients is enabled. It is read only and is ignored on PUT and POST requests. | [optional] 
**LogoUrl** | Pointer to **string** | The location of the logo used on user-facing OAuth grant authorization and revocation pages. | [optional] 
**DefaultAccessTokenManagerRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**RestrictToDefaultAccessTokenManager** | Pointer to **bool** | Determines whether the client is restricted to using only its default access token manager. The default is false. | [optional] 
**ValidateUsingAllEligibleAtms** | Pointer to **bool** | Validates token using all eligible access token managers for the client. This setting is ignored if &#39;restrictToDefaultAccessTokenManager&#39; is set to true. | [optional] 
**RefreshRolling** | Pointer to **string** | Use ROLL or DONT_ROLL to override the Roll Refresh Token Values setting on the Authorization Server Settings. SERVER_DEFAULT will default to the Roll Refresh Token Values setting on the Authorization Server Setting screen. Defaults to SERVER_DEFAULT. | [optional] 
**RefreshTokenRollingIntervalType** | Pointer to **string** | Use OVERRIDE_SERVER_DEFAULT to override the Refresh Token Rolling Interval value on the Authorization Server Settings. SERVER_DEFAULT will default to the Refresh Token Rolling Interval value on the Authorization Server Setting. Defaults to SERVER_DEFAULT. | [optional] 
**RefreshTokenRollingInterval** | Pointer to **int64** | The minimum interval to roll refresh tokens. This value will override the Refresh Token Rolling Interval Value on the Authorization Server Settings. | [optional] 
**RefreshTokenRollingIntervalTimeUnit** | Pointer to **string** | The refresh token rolling interval time unit. Defaults to HOURS. | [optional] 
**PersistentGrantExpirationType** | Pointer to **string** | Allows an administrator to override the Persistent Grant Lifetime set globally for the OAuth AS. Defaults to SERVER_DEFAULT. | [optional] 
**PersistentGrantExpirationTime** | Pointer to **int64** | The persistent grant expiration time. -1 indicates an indefinite amount of time. | [optional] 
**PersistentGrantExpirationTimeUnit** | Pointer to **string** | The persistent grant expiration time unit. | [optional] 
**PersistentGrantIdleTimeoutType** | Pointer to **string** | Allows an administrator to override the Persistent Grant Idle Timeout set globally for the OAuth AS. Defaults to SERVER_DEFAULT. | [optional] 
**PersistentGrantIdleTimeout** | Pointer to **int64** | The persistent grant idle timeout. | [optional] 
**PersistentGrantIdleTimeoutTimeUnit** | Pointer to **string** | The persistent grant idle timeout time unit. | [optional] 
**PersistentGrantReuseType** | Pointer to **string** | Allows and administrator to override the Reuse Existing Persistent Access Grants for Grant Types set globally for OAuth AS. Defaults to SERVER_DEFAULT. | [optional] 
**PersistentGrantReuseGrantTypes** | Pointer to **[]string** | The grant types that the OAuth AS can reuse rather than creating a new grant for each request. This value will override the Reuse Existing Persistent Access Grants for Grant Types on the Authorization Server Settings. Only &#39;IMPLICIT&#39; or &#39;AUTHORIZATION_CODE&#39; or &#39;RESOURCE_OWNER_CREDENTIALS&#39; are valid grant types. | [optional] 
**AllowAuthenticationApiInit** | Pointer to **bool** | Set to true to allow this client to initiate the authentication API redirectless flow. | [optional] 
**EnableCookielessAuthenticationApi** | Pointer to **bool** | Set to true to allow the authentication API redirectless flow to function without requiring any cookies. | [optional] 
**BypassApprovalPage** | Pointer to **bool** | Use this setting, for example, when you want to deploy a trusted application and authenticate end users via an IdP adapter or IdP connection. | [optional] 
**RestrictScopes** | Pointer to **bool** | Restricts this client&#39;s access to specific scopes. | [optional] 
**RestrictedScopes** | Pointer to **[]string** | The scopes available for this client. | [optional] 
**ExclusiveScopes** | Pointer to **[]string** | The exclusive scopes available for this client. | [optional] 
**AuthorizationDetailTypes** | Pointer to **[]string** | The authorization detail types available for this client. | [optional] 
**RestrictedResponseTypes** | Pointer to **[]string** | The response types allowed for this client. If omitted all response types are available to the client. | [optional] 
**RequirePushedAuthorizationRequests** | Pointer to **bool** | Determines whether pushed authorization requests are required when initiating an authorization request. The default is false. | [optional] 
**RequireJwtSecuredAuthorizationResponseMode** | Pointer to **bool** | Determines whether JWT Secured authorization response mode is required when initiating an authorization request. The default is false. | [optional] 
**RequireSignedRequests** | Pointer to **bool** | Determines whether signed requests are required for this client | [optional] 
**RequestObjectSigningAlgorithm** | Pointer to **string** | The JSON Web Signature [JWS] algorithm that must be used to sign the Request Object. All signing algorithms are allowed if value is not present &lt;br&gt;RS256 - RSA using SHA-256&lt;br&gt;RS384 - RSA using SHA-384&lt;br&gt;RS512 - RSA using SHA-512&lt;br&gt;ES256 - ECDSA using P256 Curve and SHA-256&lt;br&gt;ES384 - ECDSA using P384 Curve and SHA-384&lt;br&gt;ES512 - ECDSA using P521 Curve and SHA-512&lt;br&gt;PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256&lt;br&gt;PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384&lt;br&gt;PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512&lt;br&gt;RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11. | [optional] 
**OidcPolicy** | Pointer to [**ClientOIDCPolicy**](ClientOIDCPolicy.md) |  | [optional] 
**ClientAuth** | Pointer to [**ClientAuth**](ClientAuth.md) |  | [optional] 
**JwksSettings** | Pointer to [**JwksSettings**](JwksSettings.md) |  | [optional] 
**ExtendedParameters** | Pointer to [**map[string]ParameterValues**](ParameterValues.md) | OAuth Client Metadata can be extended to use custom Client Metadata Parameters. The names of these custom parameters should be defined in /extendedProperties. | [optional] 
**DeviceFlowSettingType** | Pointer to **string** | Allows an administrator to override the Device Authorization Settings set globally for the OAuth AS. Defaults to SERVER_DEFAULT. | [optional] 
**UserAuthorizationUrlOverride** | Pointer to **string** | The URL used as &#39;verification_url&#39; and &#39;verification_url_complete&#39; values in a Device Authorization request. This property overrides the &#39;userAuthorizationUrl&#39; value present in Authorization Server Settings. | [optional] 
**PendingAuthorizationTimeoutOverride** | Pointer to **int64** | The &#39;device_code&#39; and &#39;user_code&#39; timeout, in seconds. This overrides the &#39;pendingAuthorizationTimeout&#39; value present in Authorization Server Settings. | [optional] 
**DevicePollingIntervalOverride** | Pointer to **int64** | The amount of time client should wait between polling requests, in seconds. This overrides the &#39;devicePollingInterval&#39; value present in Authorization Server Settings. | [optional] 
**BypassActivationCodeConfirmationOverride** | Pointer to **bool** | Indicates if the Activation Code Confirmation page should be bypassed if &#39;verification_url_complete&#39; is used by the end user to authorize a device. This overrides the &#39;bypassUseCodeConfirmation&#39; value present in Authorization Server Settings. | [optional] 
**RequireProofKeyForCodeExchange** | Pointer to **bool** | Determines whether Proof Key for Code Exchange (PKCE) is required for this client. | [optional] 
**CibaDeliveryMode** | Pointer to **string** | The token delivery mode for the client.  The default value is &#39;POLL&#39;. | [optional] 
**CibaNotificationEndpoint** | Pointer to **string** | The endpoint the OP will call after a successful or failed end-user authentication. | [optional] 
**CibaPollingInterval** | Pointer to **int64** | The minimum amount of time in seconds that the Client must wait between polling requests to the token endpoint. The default is 3 seconds. | [optional] 
**CibaRequireSignedRequests** | Pointer to **bool** | Determines whether CIBA signed requests are required for this client. | [optional] 
**CibaRequestObjectSigningAlgorithm** | Pointer to **string** | The JSON Web Signature [JWS] algorithm that must be used to sign the CIBA Request Object. All signing algorithms are allowed if value is not present &lt;br&gt;RS256 - RSA using SHA-256&lt;br&gt;RS384 - RSA using SHA-384&lt;br&gt;RS512 - RSA using SHA-512&lt;br&gt;ES256 - ECDSA using P256 Curve and SHA-256&lt;br&gt;ES384 - ECDSA using P384 Curve and SHA-384&lt;br&gt;ES512 - ECDSA using P521 Curve and SHA-512&lt;br&gt;PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256&lt;br&gt;PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384&lt;br&gt;PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512&lt;br&gt;RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11. | [optional] 
**CibaUserCodeSupported** | Pointer to **bool** | Determines whether CIBA user code is supported for this client. | [optional] 
**RequestPolicyRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**TokenExchangeProcessorPolicyRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**RefreshTokenRollingGracePeriodType** | Pointer to **string** | When specified, it overrides the global Refresh Token Grace Period defined in the Authorization Server Settings. The default value is SERVER_DEFAULT | [optional] 
**RefreshTokenRollingGracePeriod** | Pointer to **int64** | The grace period that a rolled refresh token remains valid in seconds. | [optional] 
**ClientSecretRetentionPeriodType** | Pointer to **string** | Use OVERRIDE_SERVER_DEFAULT to override the Client Secret Retention Period value on the Authorization Server Settings. SERVER_DEFAULT will default to the Client Secret Retention Period value on the Authorization Server Setting. Defaults to SERVER_DEFAULT. | [optional] 
**ClientSecretRetentionPeriod** | Pointer to **int64** | The length of time in minutes that client secrets will be retained as secondary secrets after secret change. The default value is 0, which will disable secondary client secret retention. This value will override the Client Secret Retention Period value on the Authorization Server Settings. | [optional] 
**ClientSecretChangedTime** | Pointer to **time.Time** | The time at which the client secret was last changed. This property is read only and is ignored on PUT and POST requests. | [optional] 
**TokenIntrospectionSigningAlgorithm** | Pointer to **string** | The JSON Web Signature [JWS] algorithm required to sign the Token Introspection Response.&lt;br&gt;HS256 - HMAC using SHA-256&lt;br&gt;HS384 - HMAC using SHA-384&lt;br&gt;HS512 - HMAC using SHA-512&lt;br&gt;RS256 - RSA using SHA-256&lt;br&gt;RS384 - RSA using SHA-384&lt;br&gt;RS512 - RSA using SHA-512&lt;br&gt;ES256 - ECDSA using P256 Curve and SHA-256&lt;br&gt;ES384 - ECDSA using P384 Curve and SHA-384&lt;br&gt;ES512 - ECDSA using P521 Curve and SHA-512&lt;br&gt;PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256&lt;br&gt;PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384&lt;br&gt;PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512&lt;br&gt;A null value will represent the default algorithm which is RS256.&lt;br&gt;RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11 | [optional] 
**TokenIntrospectionEncryptionAlgorithm** | Pointer to **string** | The JSON Web Encryption [JWE] encryption algorithm used to encrypt the content-encryption key of the Token Introspection Response.&lt;br&gt;DIR - Direct Encryption with symmetric key&lt;br&gt;A128KW - AES-128 Key Wrap&lt;br&gt;A192KW - AES-192 Key Wrap&lt;br&gt;A256KW - AES-256 Key Wrap&lt;br&gt;A128GCMKW - AES-GCM-128 key encryption&lt;br&gt;A192GCMKW - AES-GCM-192 key encryption&lt;br&gt;A256GCMKW - AES-GCM-256 key encryption&lt;br&gt;ECDH_ES - ECDH-ES&lt;br&gt;ECDH_ES_A128KW - ECDH-ES with AES-128 Key Wrap&lt;br&gt;ECDH_ES_A192KW - ECDH-ES with AES-192 Key Wrap&lt;br&gt;ECDH_ES_A256KW - ECDH-ES with AES-256 Key Wrap&lt;br&gt;RSA_OAEP - RSAES OAEP&lt;br&gt;RSA_OAEP_256 - RSAES OAEP using SHA-256 and MGF1 with SHA-256 | [optional] 
**TokenIntrospectionContentEncryptionAlgorithm** | Pointer to **string** | The JSON Web Encryption [JWE] content-encryption algorithm for the Token Introspection Response.&lt;br&gt;AES_128_CBC_HMAC_SHA_256 - Composite AES-CBC-128 HMAC-SHA-256&lt;br&gt;AES_192_CBC_HMAC_SHA_384 - Composite AES-CBC-192 HMAC-SHA-384&lt;br&gt;AES_256_CBC_HMAC_SHA_512 - Composite AES-CBC-256 HMAC-SHA-512&lt;br&gt;AES_128_GCM - AES-GCM-128&lt;br&gt;AES_192_GCM - AES-GCM-192&lt;br&gt;AES_256_GCM - AES-GCM-256 | [optional] 
**JwtSecuredAuthorizationResponseModeSigningAlgorithm** | Pointer to **string** | The JSON Web Signature [JWS] algorithm required to sign the JWT Secured Authorization Response.&lt;br&gt;HS256 - HMAC using SHA-256&lt;br&gt;HS384 - HMAC using SHA-384&lt;br&gt;HS512 - HMAC using SHA-512&lt;br&gt;RS256 - RSA using SHA-256&lt;br&gt;RS384 - RSA using SHA-384&lt;br&gt;RS512 - RSA using SHA-512&lt;br&gt;ES256 - ECDSA using P256 Curve and SHA-256&lt;br&gt;ES384 - ECDSA using P384 Curve and SHA-384&lt;br&gt;ES512 - ECDSA using P521 Curve and SHA-512&lt;br&gt;PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256&lt;br&gt;PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384&lt;br&gt;PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512&lt;br&gt;A null value will represent the default algorithm which is RS256.&lt;br&gt;RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11 | [optional] 
**JwtSecuredAuthorizationResponseModeEncryptionAlgorithm** | Pointer to **string** | The JSON Web Encryption [JWE] encryption algorithm used to encrypt the content-encryption key of the JWT Secured Authorization Response.&lt;br&gt;DIR - Direct Encryption with symmetric key&lt;br&gt;A128KW - AES-128 Key Wrap&lt;br&gt;A192KW - AES-192 Key Wrap&lt;br&gt;A256KW - AES-256 Key Wrap&lt;br&gt;A128GCMKW - AES-GCM-128 key encryption&lt;br&gt;A192GCMKW - AES-GCM-192 key encryption&lt;br&gt;A256GCMKW - AES-GCM-256 key encryption&lt;br&gt;ECDH_ES - ECDH-ES&lt;br&gt;ECDH_ES_A128KW - ECDH-ES with AES-128 Key Wrap&lt;br&gt;ECDH_ES_A192KW - ECDH-ES with AES-192 Key Wrap&lt;br&gt;ECDH_ES_A256KW - ECDH-ES with AES-256 Key Wrap&lt;br&gt;RSA_OAEP - RSAES OAEP&lt;br&gt;RSA_OAEP_256 - RSAES OAEP using SHA-256 and MGF1 with SHA-256 | [optional] 
**JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm** | Pointer to **string** | The JSON Web Encryption [JWE] content-encryption algorithm for the JWT Secured Authorization Response.&lt;br&gt;AES_128_CBC_HMAC_SHA_256 - Composite AES-CBC-128 HMAC-SHA-256&lt;br&gt;AES_192_CBC_HMAC_SHA_384 - Composite AES-CBC-192 HMAC-SHA-384&lt;br&gt;AES_256_CBC_HMAC_SHA_512 - Composite AES-CBC-256 HMAC-SHA-512&lt;br&gt;AES_128_GCM - AES-GCM-128&lt;br&gt;AES_192_GCM - AES-GCM-192&lt;br&gt;AES_256_GCM - AES-GCM-256 | [optional] 
**RequireDpop** | Pointer to **bool** | Determines whether Demonstrating Proof-of-Possession (DPoP) is required for this client. | [optional] 
**RequireOfflineAccessScopeToIssueRefreshTokens** | Pointer to **string** | Determines whether offline_access scope is required to issue refresh tokens by this client or not. &#39;SERVER_DEFAULT&#39; is the default value.  | [optional] 
**OfflineAccessRequireConsentPrompt** | Pointer to **string** | Determines whether offline_access requires the prompt parameter value to be set to &#39;consent&#39; by this client or not. The value will be reset to default if the &#39;requireOfflineAccessScopeToIssueRefreshTokens&#39; attribute is set to &#39;SERVER_DEFAULT&#39; or &#39;false&#39;. &#39;SERVER_DEFAULT&#39; is the default value. | [optional] 

## Methods

### NewClient

`func NewClient(clientId string, grantTypes []string, name string, ) *Client`

NewClient instantiates a new Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWithDefaults

`func NewClientWithDefaults() *Client`

NewClientWithDefaults instantiates a new Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *Client) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Client) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Client) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetEnabled

`func (o *Client) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Client) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Client) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Client) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRedirectUris

`func (o *Client) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *Client) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *Client) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *Client) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetGrantTypes

`func (o *Client) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *Client) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *Client) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.


### GetName

`func (o *Client) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Client) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Client) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Client) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Client) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Client) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Client) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModificationDate

`func (o *Client) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *Client) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *Client) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *Client) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetCreationDate

`func (o *Client) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Client) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Client) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Client) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetReplicationStatus

`func (o *Client) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *Client) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *Client) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *Client) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetLogoUrl

`func (o *Client) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *Client) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *Client) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *Client) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetDefaultAccessTokenManagerRef

`func (o *Client) GetDefaultAccessTokenManagerRef() ResourceLink`

GetDefaultAccessTokenManagerRef returns the DefaultAccessTokenManagerRef field if non-nil, zero value otherwise.

### GetDefaultAccessTokenManagerRefOk

`func (o *Client) GetDefaultAccessTokenManagerRefOk() (*ResourceLink, bool)`

GetDefaultAccessTokenManagerRefOk returns a tuple with the DefaultAccessTokenManagerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccessTokenManagerRef

`func (o *Client) SetDefaultAccessTokenManagerRef(v ResourceLink)`

SetDefaultAccessTokenManagerRef sets DefaultAccessTokenManagerRef field to given value.

### HasDefaultAccessTokenManagerRef

`func (o *Client) HasDefaultAccessTokenManagerRef() bool`

HasDefaultAccessTokenManagerRef returns a boolean if a field has been set.

### GetRestrictToDefaultAccessTokenManager

`func (o *Client) GetRestrictToDefaultAccessTokenManager() bool`

GetRestrictToDefaultAccessTokenManager returns the RestrictToDefaultAccessTokenManager field if non-nil, zero value otherwise.

### GetRestrictToDefaultAccessTokenManagerOk

`func (o *Client) GetRestrictToDefaultAccessTokenManagerOk() (*bool, bool)`

GetRestrictToDefaultAccessTokenManagerOk returns a tuple with the RestrictToDefaultAccessTokenManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictToDefaultAccessTokenManager

`func (o *Client) SetRestrictToDefaultAccessTokenManager(v bool)`

SetRestrictToDefaultAccessTokenManager sets RestrictToDefaultAccessTokenManager field to given value.

### HasRestrictToDefaultAccessTokenManager

`func (o *Client) HasRestrictToDefaultAccessTokenManager() bool`

HasRestrictToDefaultAccessTokenManager returns a boolean if a field has been set.

### GetValidateUsingAllEligibleAtms

`func (o *Client) GetValidateUsingAllEligibleAtms() bool`

GetValidateUsingAllEligibleAtms returns the ValidateUsingAllEligibleAtms field if non-nil, zero value otherwise.

### GetValidateUsingAllEligibleAtmsOk

`func (o *Client) GetValidateUsingAllEligibleAtmsOk() (*bool, bool)`

GetValidateUsingAllEligibleAtmsOk returns a tuple with the ValidateUsingAllEligibleAtms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateUsingAllEligibleAtms

`func (o *Client) SetValidateUsingAllEligibleAtms(v bool)`

SetValidateUsingAllEligibleAtms sets ValidateUsingAllEligibleAtms field to given value.

### HasValidateUsingAllEligibleAtms

`func (o *Client) HasValidateUsingAllEligibleAtms() bool`

HasValidateUsingAllEligibleAtms returns a boolean if a field has been set.

### GetRefreshRolling

`func (o *Client) GetRefreshRolling() string`

GetRefreshRolling returns the RefreshRolling field if non-nil, zero value otherwise.

### GetRefreshRollingOk

`func (o *Client) GetRefreshRollingOk() (*string, bool)`

GetRefreshRollingOk returns a tuple with the RefreshRolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshRolling

`func (o *Client) SetRefreshRolling(v string)`

SetRefreshRolling sets RefreshRolling field to given value.

### HasRefreshRolling

`func (o *Client) HasRefreshRolling() bool`

HasRefreshRolling returns a boolean if a field has been set.

### GetRefreshTokenRollingIntervalType

`func (o *Client) GetRefreshTokenRollingIntervalType() string`

GetRefreshTokenRollingIntervalType returns the RefreshTokenRollingIntervalType field if non-nil, zero value otherwise.

### GetRefreshTokenRollingIntervalTypeOk

`func (o *Client) GetRefreshTokenRollingIntervalTypeOk() (*string, bool)`

GetRefreshTokenRollingIntervalTypeOk returns a tuple with the RefreshTokenRollingIntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingIntervalType

`func (o *Client) SetRefreshTokenRollingIntervalType(v string)`

SetRefreshTokenRollingIntervalType sets RefreshTokenRollingIntervalType field to given value.

### HasRefreshTokenRollingIntervalType

`func (o *Client) HasRefreshTokenRollingIntervalType() bool`

HasRefreshTokenRollingIntervalType returns a boolean if a field has been set.

### GetRefreshTokenRollingInterval

`func (o *Client) GetRefreshTokenRollingInterval() int64`

GetRefreshTokenRollingInterval returns the RefreshTokenRollingInterval field if non-nil, zero value otherwise.

### GetRefreshTokenRollingIntervalOk

`func (o *Client) GetRefreshTokenRollingIntervalOk() (*int64, bool)`

GetRefreshTokenRollingIntervalOk returns a tuple with the RefreshTokenRollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingInterval

`func (o *Client) SetRefreshTokenRollingInterval(v int64)`

SetRefreshTokenRollingInterval sets RefreshTokenRollingInterval field to given value.

### HasRefreshTokenRollingInterval

`func (o *Client) HasRefreshTokenRollingInterval() bool`

HasRefreshTokenRollingInterval returns a boolean if a field has been set.

### GetRefreshTokenRollingIntervalTimeUnit

`func (o *Client) GetRefreshTokenRollingIntervalTimeUnit() string`

GetRefreshTokenRollingIntervalTimeUnit returns the RefreshTokenRollingIntervalTimeUnit field if non-nil, zero value otherwise.

### GetRefreshTokenRollingIntervalTimeUnitOk

`func (o *Client) GetRefreshTokenRollingIntervalTimeUnitOk() (*string, bool)`

GetRefreshTokenRollingIntervalTimeUnitOk returns a tuple with the RefreshTokenRollingIntervalTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingIntervalTimeUnit

`func (o *Client) SetRefreshTokenRollingIntervalTimeUnit(v string)`

SetRefreshTokenRollingIntervalTimeUnit sets RefreshTokenRollingIntervalTimeUnit field to given value.

### HasRefreshTokenRollingIntervalTimeUnit

`func (o *Client) HasRefreshTokenRollingIntervalTimeUnit() bool`

HasRefreshTokenRollingIntervalTimeUnit returns a boolean if a field has been set.

### GetPersistentGrantExpirationType

`func (o *Client) GetPersistentGrantExpirationType() string`

GetPersistentGrantExpirationType returns the PersistentGrantExpirationType field if non-nil, zero value otherwise.

### GetPersistentGrantExpirationTypeOk

`func (o *Client) GetPersistentGrantExpirationTypeOk() (*string, bool)`

GetPersistentGrantExpirationTypeOk returns a tuple with the PersistentGrantExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantExpirationType

`func (o *Client) SetPersistentGrantExpirationType(v string)`

SetPersistentGrantExpirationType sets PersistentGrantExpirationType field to given value.

### HasPersistentGrantExpirationType

`func (o *Client) HasPersistentGrantExpirationType() bool`

HasPersistentGrantExpirationType returns a boolean if a field has been set.

### GetPersistentGrantExpirationTime

`func (o *Client) GetPersistentGrantExpirationTime() int64`

GetPersistentGrantExpirationTime returns the PersistentGrantExpirationTime field if non-nil, zero value otherwise.

### GetPersistentGrantExpirationTimeOk

`func (o *Client) GetPersistentGrantExpirationTimeOk() (*int64, bool)`

GetPersistentGrantExpirationTimeOk returns a tuple with the PersistentGrantExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantExpirationTime

`func (o *Client) SetPersistentGrantExpirationTime(v int64)`

SetPersistentGrantExpirationTime sets PersistentGrantExpirationTime field to given value.

### HasPersistentGrantExpirationTime

`func (o *Client) HasPersistentGrantExpirationTime() bool`

HasPersistentGrantExpirationTime returns a boolean if a field has been set.

### GetPersistentGrantExpirationTimeUnit

`func (o *Client) GetPersistentGrantExpirationTimeUnit() string`

GetPersistentGrantExpirationTimeUnit returns the PersistentGrantExpirationTimeUnit field if non-nil, zero value otherwise.

### GetPersistentGrantExpirationTimeUnitOk

`func (o *Client) GetPersistentGrantExpirationTimeUnitOk() (*string, bool)`

GetPersistentGrantExpirationTimeUnitOk returns a tuple with the PersistentGrantExpirationTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantExpirationTimeUnit

`func (o *Client) SetPersistentGrantExpirationTimeUnit(v string)`

SetPersistentGrantExpirationTimeUnit sets PersistentGrantExpirationTimeUnit field to given value.

### HasPersistentGrantExpirationTimeUnit

`func (o *Client) HasPersistentGrantExpirationTimeUnit() bool`

HasPersistentGrantExpirationTimeUnit returns a boolean if a field has been set.

### GetPersistentGrantIdleTimeoutType

`func (o *Client) GetPersistentGrantIdleTimeoutType() string`

GetPersistentGrantIdleTimeoutType returns the PersistentGrantIdleTimeoutType field if non-nil, zero value otherwise.

### GetPersistentGrantIdleTimeoutTypeOk

`func (o *Client) GetPersistentGrantIdleTimeoutTypeOk() (*string, bool)`

GetPersistentGrantIdleTimeoutTypeOk returns a tuple with the PersistentGrantIdleTimeoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantIdleTimeoutType

`func (o *Client) SetPersistentGrantIdleTimeoutType(v string)`

SetPersistentGrantIdleTimeoutType sets PersistentGrantIdleTimeoutType field to given value.

### HasPersistentGrantIdleTimeoutType

`func (o *Client) HasPersistentGrantIdleTimeoutType() bool`

HasPersistentGrantIdleTimeoutType returns a boolean if a field has been set.

### GetPersistentGrantIdleTimeout

`func (o *Client) GetPersistentGrantIdleTimeout() int64`

GetPersistentGrantIdleTimeout returns the PersistentGrantIdleTimeout field if non-nil, zero value otherwise.

### GetPersistentGrantIdleTimeoutOk

`func (o *Client) GetPersistentGrantIdleTimeoutOk() (*int64, bool)`

GetPersistentGrantIdleTimeoutOk returns a tuple with the PersistentGrantIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantIdleTimeout

`func (o *Client) SetPersistentGrantIdleTimeout(v int64)`

SetPersistentGrantIdleTimeout sets PersistentGrantIdleTimeout field to given value.

### HasPersistentGrantIdleTimeout

`func (o *Client) HasPersistentGrantIdleTimeout() bool`

HasPersistentGrantIdleTimeout returns a boolean if a field has been set.

### GetPersistentGrantIdleTimeoutTimeUnit

`func (o *Client) GetPersistentGrantIdleTimeoutTimeUnit() string`

GetPersistentGrantIdleTimeoutTimeUnit returns the PersistentGrantIdleTimeoutTimeUnit field if non-nil, zero value otherwise.

### GetPersistentGrantIdleTimeoutTimeUnitOk

`func (o *Client) GetPersistentGrantIdleTimeoutTimeUnitOk() (*string, bool)`

GetPersistentGrantIdleTimeoutTimeUnitOk returns a tuple with the PersistentGrantIdleTimeoutTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantIdleTimeoutTimeUnit

`func (o *Client) SetPersistentGrantIdleTimeoutTimeUnit(v string)`

SetPersistentGrantIdleTimeoutTimeUnit sets PersistentGrantIdleTimeoutTimeUnit field to given value.

### HasPersistentGrantIdleTimeoutTimeUnit

`func (o *Client) HasPersistentGrantIdleTimeoutTimeUnit() bool`

HasPersistentGrantIdleTimeoutTimeUnit returns a boolean if a field has been set.

### GetPersistentGrantReuseType

`func (o *Client) GetPersistentGrantReuseType() string`

GetPersistentGrantReuseType returns the PersistentGrantReuseType field if non-nil, zero value otherwise.

### GetPersistentGrantReuseTypeOk

`func (o *Client) GetPersistentGrantReuseTypeOk() (*string, bool)`

GetPersistentGrantReuseTypeOk returns a tuple with the PersistentGrantReuseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantReuseType

`func (o *Client) SetPersistentGrantReuseType(v string)`

SetPersistentGrantReuseType sets PersistentGrantReuseType field to given value.

### HasPersistentGrantReuseType

`func (o *Client) HasPersistentGrantReuseType() bool`

HasPersistentGrantReuseType returns a boolean if a field has been set.

### GetPersistentGrantReuseGrantTypes

`func (o *Client) GetPersistentGrantReuseGrantTypes() []string`

GetPersistentGrantReuseGrantTypes returns the PersistentGrantReuseGrantTypes field if non-nil, zero value otherwise.

### GetPersistentGrantReuseGrantTypesOk

`func (o *Client) GetPersistentGrantReuseGrantTypesOk() (*[]string, bool)`

GetPersistentGrantReuseGrantTypesOk returns a tuple with the PersistentGrantReuseGrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantReuseGrantTypes

`func (o *Client) SetPersistentGrantReuseGrantTypes(v []string)`

SetPersistentGrantReuseGrantTypes sets PersistentGrantReuseGrantTypes field to given value.

### HasPersistentGrantReuseGrantTypes

`func (o *Client) HasPersistentGrantReuseGrantTypes() bool`

HasPersistentGrantReuseGrantTypes returns a boolean if a field has been set.

### GetAllowAuthenticationApiInit

`func (o *Client) GetAllowAuthenticationApiInit() bool`

GetAllowAuthenticationApiInit returns the AllowAuthenticationApiInit field if non-nil, zero value otherwise.

### GetAllowAuthenticationApiInitOk

`func (o *Client) GetAllowAuthenticationApiInitOk() (*bool, bool)`

GetAllowAuthenticationApiInitOk returns a tuple with the AllowAuthenticationApiInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAuthenticationApiInit

`func (o *Client) SetAllowAuthenticationApiInit(v bool)`

SetAllowAuthenticationApiInit sets AllowAuthenticationApiInit field to given value.

### HasAllowAuthenticationApiInit

`func (o *Client) HasAllowAuthenticationApiInit() bool`

HasAllowAuthenticationApiInit returns a boolean if a field has been set.

### GetEnableCookielessAuthenticationApi

`func (o *Client) GetEnableCookielessAuthenticationApi() bool`

GetEnableCookielessAuthenticationApi returns the EnableCookielessAuthenticationApi field if non-nil, zero value otherwise.

### GetEnableCookielessAuthenticationApiOk

`func (o *Client) GetEnableCookielessAuthenticationApiOk() (*bool, bool)`

GetEnableCookielessAuthenticationApiOk returns a tuple with the EnableCookielessAuthenticationApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCookielessAuthenticationApi

`func (o *Client) SetEnableCookielessAuthenticationApi(v bool)`

SetEnableCookielessAuthenticationApi sets EnableCookielessAuthenticationApi field to given value.

### HasEnableCookielessAuthenticationApi

`func (o *Client) HasEnableCookielessAuthenticationApi() bool`

HasEnableCookielessAuthenticationApi returns a boolean if a field has been set.

### GetBypassApprovalPage

`func (o *Client) GetBypassApprovalPage() bool`

GetBypassApprovalPage returns the BypassApprovalPage field if non-nil, zero value otherwise.

### GetBypassApprovalPageOk

`func (o *Client) GetBypassApprovalPageOk() (*bool, bool)`

GetBypassApprovalPageOk returns a tuple with the BypassApprovalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassApprovalPage

`func (o *Client) SetBypassApprovalPage(v bool)`

SetBypassApprovalPage sets BypassApprovalPage field to given value.

### HasBypassApprovalPage

`func (o *Client) HasBypassApprovalPage() bool`

HasBypassApprovalPage returns a boolean if a field has been set.

### GetRestrictScopes

`func (o *Client) GetRestrictScopes() bool`

GetRestrictScopes returns the RestrictScopes field if non-nil, zero value otherwise.

### GetRestrictScopesOk

`func (o *Client) GetRestrictScopesOk() (*bool, bool)`

GetRestrictScopesOk returns a tuple with the RestrictScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictScopes

`func (o *Client) SetRestrictScopes(v bool)`

SetRestrictScopes sets RestrictScopes field to given value.

### HasRestrictScopes

`func (o *Client) HasRestrictScopes() bool`

HasRestrictScopes returns a boolean if a field has been set.

### GetRestrictedScopes

`func (o *Client) GetRestrictedScopes() []string`

GetRestrictedScopes returns the RestrictedScopes field if non-nil, zero value otherwise.

### GetRestrictedScopesOk

`func (o *Client) GetRestrictedScopesOk() (*[]string, bool)`

GetRestrictedScopesOk returns a tuple with the RestrictedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedScopes

`func (o *Client) SetRestrictedScopes(v []string)`

SetRestrictedScopes sets RestrictedScopes field to given value.

### HasRestrictedScopes

`func (o *Client) HasRestrictedScopes() bool`

HasRestrictedScopes returns a boolean if a field has been set.

### GetExclusiveScopes

`func (o *Client) GetExclusiveScopes() []string`

GetExclusiveScopes returns the ExclusiveScopes field if non-nil, zero value otherwise.

### GetExclusiveScopesOk

`func (o *Client) GetExclusiveScopesOk() (*[]string, bool)`

GetExclusiveScopesOk returns a tuple with the ExclusiveScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveScopes

`func (o *Client) SetExclusiveScopes(v []string)`

SetExclusiveScopes sets ExclusiveScopes field to given value.

### HasExclusiveScopes

`func (o *Client) HasExclusiveScopes() bool`

HasExclusiveScopes returns a boolean if a field has been set.

### GetAuthorizationDetailTypes

`func (o *Client) GetAuthorizationDetailTypes() []string`

GetAuthorizationDetailTypes returns the AuthorizationDetailTypes field if non-nil, zero value otherwise.

### GetAuthorizationDetailTypesOk

`func (o *Client) GetAuthorizationDetailTypesOk() (*[]string, bool)`

GetAuthorizationDetailTypesOk returns a tuple with the AuthorizationDetailTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationDetailTypes

`func (o *Client) SetAuthorizationDetailTypes(v []string)`

SetAuthorizationDetailTypes sets AuthorizationDetailTypes field to given value.

### HasAuthorizationDetailTypes

`func (o *Client) HasAuthorizationDetailTypes() bool`

HasAuthorizationDetailTypes returns a boolean if a field has been set.

### GetRestrictedResponseTypes

`func (o *Client) GetRestrictedResponseTypes() []string`

GetRestrictedResponseTypes returns the RestrictedResponseTypes field if non-nil, zero value otherwise.

### GetRestrictedResponseTypesOk

`func (o *Client) GetRestrictedResponseTypesOk() (*[]string, bool)`

GetRestrictedResponseTypesOk returns a tuple with the RestrictedResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedResponseTypes

`func (o *Client) SetRestrictedResponseTypes(v []string)`

SetRestrictedResponseTypes sets RestrictedResponseTypes field to given value.

### HasRestrictedResponseTypes

`func (o *Client) HasRestrictedResponseTypes() bool`

HasRestrictedResponseTypes returns a boolean if a field has been set.

### GetRequirePushedAuthorizationRequests

`func (o *Client) GetRequirePushedAuthorizationRequests() bool`

GetRequirePushedAuthorizationRequests returns the RequirePushedAuthorizationRequests field if non-nil, zero value otherwise.

### GetRequirePushedAuthorizationRequestsOk

`func (o *Client) GetRequirePushedAuthorizationRequestsOk() (*bool, bool)`

GetRequirePushedAuthorizationRequestsOk returns a tuple with the RequirePushedAuthorizationRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePushedAuthorizationRequests

`func (o *Client) SetRequirePushedAuthorizationRequests(v bool)`

SetRequirePushedAuthorizationRequests sets RequirePushedAuthorizationRequests field to given value.

### HasRequirePushedAuthorizationRequests

`func (o *Client) HasRequirePushedAuthorizationRequests() bool`

HasRequirePushedAuthorizationRequests returns a boolean if a field has been set.

### GetRequireJwtSecuredAuthorizationResponseMode

`func (o *Client) GetRequireJwtSecuredAuthorizationResponseMode() bool`

GetRequireJwtSecuredAuthorizationResponseMode returns the RequireJwtSecuredAuthorizationResponseMode field if non-nil, zero value otherwise.

### GetRequireJwtSecuredAuthorizationResponseModeOk

`func (o *Client) GetRequireJwtSecuredAuthorizationResponseModeOk() (*bool, bool)`

GetRequireJwtSecuredAuthorizationResponseModeOk returns a tuple with the RequireJwtSecuredAuthorizationResponseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireJwtSecuredAuthorizationResponseMode

`func (o *Client) SetRequireJwtSecuredAuthorizationResponseMode(v bool)`

SetRequireJwtSecuredAuthorizationResponseMode sets RequireJwtSecuredAuthorizationResponseMode field to given value.

### HasRequireJwtSecuredAuthorizationResponseMode

`func (o *Client) HasRequireJwtSecuredAuthorizationResponseMode() bool`

HasRequireJwtSecuredAuthorizationResponseMode returns a boolean if a field has been set.

### GetRequireSignedRequests

`func (o *Client) GetRequireSignedRequests() bool`

GetRequireSignedRequests returns the RequireSignedRequests field if non-nil, zero value otherwise.

### GetRequireSignedRequestsOk

`func (o *Client) GetRequireSignedRequestsOk() (*bool, bool)`

GetRequireSignedRequestsOk returns a tuple with the RequireSignedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedRequests

`func (o *Client) SetRequireSignedRequests(v bool)`

SetRequireSignedRequests sets RequireSignedRequests field to given value.

### HasRequireSignedRequests

`func (o *Client) HasRequireSignedRequests() bool`

HasRequireSignedRequests returns a boolean if a field has been set.

### GetRequestObjectSigningAlgorithm

`func (o *Client) GetRequestObjectSigningAlgorithm() string`

GetRequestObjectSigningAlgorithm returns the RequestObjectSigningAlgorithm field if non-nil, zero value otherwise.

### GetRequestObjectSigningAlgorithmOk

`func (o *Client) GetRequestObjectSigningAlgorithmOk() (*string, bool)`

GetRequestObjectSigningAlgorithmOk returns a tuple with the RequestObjectSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectSigningAlgorithm

`func (o *Client) SetRequestObjectSigningAlgorithm(v string)`

SetRequestObjectSigningAlgorithm sets RequestObjectSigningAlgorithm field to given value.

### HasRequestObjectSigningAlgorithm

`func (o *Client) HasRequestObjectSigningAlgorithm() bool`

HasRequestObjectSigningAlgorithm returns a boolean if a field has been set.

### GetOidcPolicy

`func (o *Client) GetOidcPolicy() ClientOIDCPolicy`

GetOidcPolicy returns the OidcPolicy field if non-nil, zero value otherwise.

### GetOidcPolicyOk

`func (o *Client) GetOidcPolicyOk() (*ClientOIDCPolicy, bool)`

GetOidcPolicyOk returns a tuple with the OidcPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcPolicy

`func (o *Client) SetOidcPolicy(v ClientOIDCPolicy)`

SetOidcPolicy sets OidcPolicy field to given value.

### HasOidcPolicy

`func (o *Client) HasOidcPolicy() bool`

HasOidcPolicy returns a boolean if a field has been set.

### GetClientAuth

`func (o *Client) GetClientAuth() ClientAuth`

GetClientAuth returns the ClientAuth field if non-nil, zero value otherwise.

### GetClientAuthOk

`func (o *Client) GetClientAuthOk() (*ClientAuth, bool)`

GetClientAuthOk returns a tuple with the ClientAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuth

`func (o *Client) SetClientAuth(v ClientAuth)`

SetClientAuth sets ClientAuth field to given value.

### HasClientAuth

`func (o *Client) HasClientAuth() bool`

HasClientAuth returns a boolean if a field has been set.

### GetJwksSettings

`func (o *Client) GetJwksSettings() JwksSettings`

GetJwksSettings returns the JwksSettings field if non-nil, zero value otherwise.

### GetJwksSettingsOk

`func (o *Client) GetJwksSettingsOk() (*JwksSettings, bool)`

GetJwksSettingsOk returns a tuple with the JwksSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksSettings

`func (o *Client) SetJwksSettings(v JwksSettings)`

SetJwksSettings sets JwksSettings field to given value.

### HasJwksSettings

`func (o *Client) HasJwksSettings() bool`

HasJwksSettings returns a boolean if a field has been set.

### GetExtendedParameters

`func (o *Client) GetExtendedParameters() map[string]ParameterValues`

GetExtendedParameters returns the ExtendedParameters field if non-nil, zero value otherwise.

### GetExtendedParametersOk

`func (o *Client) GetExtendedParametersOk() (*map[string]ParameterValues, bool)`

GetExtendedParametersOk returns a tuple with the ExtendedParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedParameters

`func (o *Client) SetExtendedParameters(v map[string]ParameterValues)`

SetExtendedParameters sets ExtendedParameters field to given value.

### HasExtendedParameters

`func (o *Client) HasExtendedParameters() bool`

HasExtendedParameters returns a boolean if a field has been set.

### GetDeviceFlowSettingType

`func (o *Client) GetDeviceFlowSettingType() string`

GetDeviceFlowSettingType returns the DeviceFlowSettingType field if non-nil, zero value otherwise.

### GetDeviceFlowSettingTypeOk

`func (o *Client) GetDeviceFlowSettingTypeOk() (*string, bool)`

GetDeviceFlowSettingTypeOk returns a tuple with the DeviceFlowSettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFlowSettingType

`func (o *Client) SetDeviceFlowSettingType(v string)`

SetDeviceFlowSettingType sets DeviceFlowSettingType field to given value.

### HasDeviceFlowSettingType

`func (o *Client) HasDeviceFlowSettingType() bool`

HasDeviceFlowSettingType returns a boolean if a field has been set.

### GetUserAuthorizationUrlOverride

`func (o *Client) GetUserAuthorizationUrlOverride() string`

GetUserAuthorizationUrlOverride returns the UserAuthorizationUrlOverride field if non-nil, zero value otherwise.

### GetUserAuthorizationUrlOverrideOk

`func (o *Client) GetUserAuthorizationUrlOverrideOk() (*string, bool)`

GetUserAuthorizationUrlOverrideOk returns a tuple with the UserAuthorizationUrlOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthorizationUrlOverride

`func (o *Client) SetUserAuthorizationUrlOverride(v string)`

SetUserAuthorizationUrlOverride sets UserAuthorizationUrlOverride field to given value.

### HasUserAuthorizationUrlOverride

`func (o *Client) HasUserAuthorizationUrlOverride() bool`

HasUserAuthorizationUrlOverride returns a boolean if a field has been set.

### GetPendingAuthorizationTimeoutOverride

`func (o *Client) GetPendingAuthorizationTimeoutOverride() int64`

GetPendingAuthorizationTimeoutOverride returns the PendingAuthorizationTimeoutOverride field if non-nil, zero value otherwise.

### GetPendingAuthorizationTimeoutOverrideOk

`func (o *Client) GetPendingAuthorizationTimeoutOverrideOk() (*int64, bool)`

GetPendingAuthorizationTimeoutOverrideOk returns a tuple with the PendingAuthorizationTimeoutOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAuthorizationTimeoutOverride

`func (o *Client) SetPendingAuthorizationTimeoutOverride(v int64)`

SetPendingAuthorizationTimeoutOverride sets PendingAuthorizationTimeoutOverride field to given value.

### HasPendingAuthorizationTimeoutOverride

`func (o *Client) HasPendingAuthorizationTimeoutOverride() bool`

HasPendingAuthorizationTimeoutOverride returns a boolean if a field has been set.

### GetDevicePollingIntervalOverride

`func (o *Client) GetDevicePollingIntervalOverride() int64`

GetDevicePollingIntervalOverride returns the DevicePollingIntervalOverride field if non-nil, zero value otherwise.

### GetDevicePollingIntervalOverrideOk

`func (o *Client) GetDevicePollingIntervalOverrideOk() (*int64, bool)`

GetDevicePollingIntervalOverrideOk returns a tuple with the DevicePollingIntervalOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePollingIntervalOverride

`func (o *Client) SetDevicePollingIntervalOverride(v int64)`

SetDevicePollingIntervalOverride sets DevicePollingIntervalOverride field to given value.

### HasDevicePollingIntervalOverride

`func (o *Client) HasDevicePollingIntervalOverride() bool`

HasDevicePollingIntervalOverride returns a boolean if a field has been set.

### GetBypassActivationCodeConfirmationOverride

`func (o *Client) GetBypassActivationCodeConfirmationOverride() bool`

GetBypassActivationCodeConfirmationOverride returns the BypassActivationCodeConfirmationOverride field if non-nil, zero value otherwise.

### GetBypassActivationCodeConfirmationOverrideOk

`func (o *Client) GetBypassActivationCodeConfirmationOverrideOk() (*bool, bool)`

GetBypassActivationCodeConfirmationOverrideOk returns a tuple with the BypassActivationCodeConfirmationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassActivationCodeConfirmationOverride

`func (o *Client) SetBypassActivationCodeConfirmationOverride(v bool)`

SetBypassActivationCodeConfirmationOverride sets BypassActivationCodeConfirmationOverride field to given value.

### HasBypassActivationCodeConfirmationOverride

`func (o *Client) HasBypassActivationCodeConfirmationOverride() bool`

HasBypassActivationCodeConfirmationOverride returns a boolean if a field has been set.

### GetRequireProofKeyForCodeExchange

`func (o *Client) GetRequireProofKeyForCodeExchange() bool`

GetRequireProofKeyForCodeExchange returns the RequireProofKeyForCodeExchange field if non-nil, zero value otherwise.

### GetRequireProofKeyForCodeExchangeOk

`func (o *Client) GetRequireProofKeyForCodeExchangeOk() (*bool, bool)`

GetRequireProofKeyForCodeExchangeOk returns a tuple with the RequireProofKeyForCodeExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireProofKeyForCodeExchange

`func (o *Client) SetRequireProofKeyForCodeExchange(v bool)`

SetRequireProofKeyForCodeExchange sets RequireProofKeyForCodeExchange field to given value.

### HasRequireProofKeyForCodeExchange

`func (o *Client) HasRequireProofKeyForCodeExchange() bool`

HasRequireProofKeyForCodeExchange returns a boolean if a field has been set.

### GetCibaDeliveryMode

`func (o *Client) GetCibaDeliveryMode() string`

GetCibaDeliveryMode returns the CibaDeliveryMode field if non-nil, zero value otherwise.

### GetCibaDeliveryModeOk

`func (o *Client) GetCibaDeliveryModeOk() (*string, bool)`

GetCibaDeliveryModeOk returns a tuple with the CibaDeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCibaDeliveryMode

`func (o *Client) SetCibaDeliveryMode(v string)`

SetCibaDeliveryMode sets CibaDeliveryMode field to given value.

### HasCibaDeliveryMode

`func (o *Client) HasCibaDeliveryMode() bool`

HasCibaDeliveryMode returns a boolean if a field has been set.

### GetCibaNotificationEndpoint

`func (o *Client) GetCibaNotificationEndpoint() string`

GetCibaNotificationEndpoint returns the CibaNotificationEndpoint field if non-nil, zero value otherwise.

### GetCibaNotificationEndpointOk

`func (o *Client) GetCibaNotificationEndpointOk() (*string, bool)`

GetCibaNotificationEndpointOk returns a tuple with the CibaNotificationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCibaNotificationEndpoint

`func (o *Client) SetCibaNotificationEndpoint(v string)`

SetCibaNotificationEndpoint sets CibaNotificationEndpoint field to given value.

### HasCibaNotificationEndpoint

`func (o *Client) HasCibaNotificationEndpoint() bool`

HasCibaNotificationEndpoint returns a boolean if a field has been set.

### GetCibaPollingInterval

`func (o *Client) GetCibaPollingInterval() int64`

GetCibaPollingInterval returns the CibaPollingInterval field if non-nil, zero value otherwise.

### GetCibaPollingIntervalOk

`func (o *Client) GetCibaPollingIntervalOk() (*int64, bool)`

GetCibaPollingIntervalOk returns a tuple with the CibaPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCibaPollingInterval

`func (o *Client) SetCibaPollingInterval(v int64)`

SetCibaPollingInterval sets CibaPollingInterval field to given value.

### HasCibaPollingInterval

`func (o *Client) HasCibaPollingInterval() bool`

HasCibaPollingInterval returns a boolean if a field has been set.

### GetCibaRequireSignedRequests

`func (o *Client) GetCibaRequireSignedRequests() bool`

GetCibaRequireSignedRequests returns the CibaRequireSignedRequests field if non-nil, zero value otherwise.

### GetCibaRequireSignedRequestsOk

`func (o *Client) GetCibaRequireSignedRequestsOk() (*bool, bool)`

GetCibaRequireSignedRequestsOk returns a tuple with the CibaRequireSignedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCibaRequireSignedRequests

`func (o *Client) SetCibaRequireSignedRequests(v bool)`

SetCibaRequireSignedRequests sets CibaRequireSignedRequests field to given value.

### HasCibaRequireSignedRequests

`func (o *Client) HasCibaRequireSignedRequests() bool`

HasCibaRequireSignedRequests returns a boolean if a field has been set.

### GetCibaRequestObjectSigningAlgorithm

`func (o *Client) GetCibaRequestObjectSigningAlgorithm() string`

GetCibaRequestObjectSigningAlgorithm returns the CibaRequestObjectSigningAlgorithm field if non-nil, zero value otherwise.

### GetCibaRequestObjectSigningAlgorithmOk

`func (o *Client) GetCibaRequestObjectSigningAlgorithmOk() (*string, bool)`

GetCibaRequestObjectSigningAlgorithmOk returns a tuple with the CibaRequestObjectSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCibaRequestObjectSigningAlgorithm

`func (o *Client) SetCibaRequestObjectSigningAlgorithm(v string)`

SetCibaRequestObjectSigningAlgorithm sets CibaRequestObjectSigningAlgorithm field to given value.

### HasCibaRequestObjectSigningAlgorithm

`func (o *Client) HasCibaRequestObjectSigningAlgorithm() bool`

HasCibaRequestObjectSigningAlgorithm returns a boolean if a field has been set.

### GetCibaUserCodeSupported

`func (o *Client) GetCibaUserCodeSupported() bool`

GetCibaUserCodeSupported returns the CibaUserCodeSupported field if non-nil, zero value otherwise.

### GetCibaUserCodeSupportedOk

`func (o *Client) GetCibaUserCodeSupportedOk() (*bool, bool)`

GetCibaUserCodeSupportedOk returns a tuple with the CibaUserCodeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCibaUserCodeSupported

`func (o *Client) SetCibaUserCodeSupported(v bool)`

SetCibaUserCodeSupported sets CibaUserCodeSupported field to given value.

### HasCibaUserCodeSupported

`func (o *Client) HasCibaUserCodeSupported() bool`

HasCibaUserCodeSupported returns a boolean if a field has been set.

### GetRequestPolicyRef

`func (o *Client) GetRequestPolicyRef() ResourceLink`

GetRequestPolicyRef returns the RequestPolicyRef field if non-nil, zero value otherwise.

### GetRequestPolicyRefOk

`func (o *Client) GetRequestPolicyRefOk() (*ResourceLink, bool)`

GetRequestPolicyRefOk returns a tuple with the RequestPolicyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPolicyRef

`func (o *Client) SetRequestPolicyRef(v ResourceLink)`

SetRequestPolicyRef sets RequestPolicyRef field to given value.

### HasRequestPolicyRef

`func (o *Client) HasRequestPolicyRef() bool`

HasRequestPolicyRef returns a boolean if a field has been set.

### GetTokenExchangeProcessorPolicyRef

`func (o *Client) GetTokenExchangeProcessorPolicyRef() ResourceLink`

GetTokenExchangeProcessorPolicyRef returns the TokenExchangeProcessorPolicyRef field if non-nil, zero value otherwise.

### GetTokenExchangeProcessorPolicyRefOk

`func (o *Client) GetTokenExchangeProcessorPolicyRefOk() (*ResourceLink, bool)`

GetTokenExchangeProcessorPolicyRefOk returns a tuple with the TokenExchangeProcessorPolicyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangeProcessorPolicyRef

`func (o *Client) SetTokenExchangeProcessorPolicyRef(v ResourceLink)`

SetTokenExchangeProcessorPolicyRef sets TokenExchangeProcessorPolicyRef field to given value.

### HasTokenExchangeProcessorPolicyRef

`func (o *Client) HasTokenExchangeProcessorPolicyRef() bool`

HasTokenExchangeProcessorPolicyRef returns a boolean if a field has been set.

### GetRefreshTokenRollingGracePeriodType

`func (o *Client) GetRefreshTokenRollingGracePeriodType() string`

GetRefreshTokenRollingGracePeriodType returns the RefreshTokenRollingGracePeriodType field if non-nil, zero value otherwise.

### GetRefreshTokenRollingGracePeriodTypeOk

`func (o *Client) GetRefreshTokenRollingGracePeriodTypeOk() (*string, bool)`

GetRefreshTokenRollingGracePeriodTypeOk returns a tuple with the RefreshTokenRollingGracePeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingGracePeriodType

`func (o *Client) SetRefreshTokenRollingGracePeriodType(v string)`

SetRefreshTokenRollingGracePeriodType sets RefreshTokenRollingGracePeriodType field to given value.

### HasRefreshTokenRollingGracePeriodType

`func (o *Client) HasRefreshTokenRollingGracePeriodType() bool`

HasRefreshTokenRollingGracePeriodType returns a boolean if a field has been set.

### GetRefreshTokenRollingGracePeriod

`func (o *Client) GetRefreshTokenRollingGracePeriod() int64`

GetRefreshTokenRollingGracePeriod returns the RefreshTokenRollingGracePeriod field if non-nil, zero value otherwise.

### GetRefreshTokenRollingGracePeriodOk

`func (o *Client) GetRefreshTokenRollingGracePeriodOk() (*int64, bool)`

GetRefreshTokenRollingGracePeriodOk returns a tuple with the RefreshTokenRollingGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingGracePeriod

`func (o *Client) SetRefreshTokenRollingGracePeriod(v int64)`

SetRefreshTokenRollingGracePeriod sets RefreshTokenRollingGracePeriod field to given value.

### HasRefreshTokenRollingGracePeriod

`func (o *Client) HasRefreshTokenRollingGracePeriod() bool`

HasRefreshTokenRollingGracePeriod returns a boolean if a field has been set.

### GetClientSecretRetentionPeriodType

`func (o *Client) GetClientSecretRetentionPeriodType() string`

GetClientSecretRetentionPeriodType returns the ClientSecretRetentionPeriodType field if non-nil, zero value otherwise.

### GetClientSecretRetentionPeriodTypeOk

`func (o *Client) GetClientSecretRetentionPeriodTypeOk() (*string, bool)`

GetClientSecretRetentionPeriodTypeOk returns a tuple with the ClientSecretRetentionPeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretRetentionPeriodType

`func (o *Client) SetClientSecretRetentionPeriodType(v string)`

SetClientSecretRetentionPeriodType sets ClientSecretRetentionPeriodType field to given value.

### HasClientSecretRetentionPeriodType

`func (o *Client) HasClientSecretRetentionPeriodType() bool`

HasClientSecretRetentionPeriodType returns a boolean if a field has been set.

### GetClientSecretRetentionPeriod

`func (o *Client) GetClientSecretRetentionPeriod() int64`

GetClientSecretRetentionPeriod returns the ClientSecretRetentionPeriod field if non-nil, zero value otherwise.

### GetClientSecretRetentionPeriodOk

`func (o *Client) GetClientSecretRetentionPeriodOk() (*int64, bool)`

GetClientSecretRetentionPeriodOk returns a tuple with the ClientSecretRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretRetentionPeriod

`func (o *Client) SetClientSecretRetentionPeriod(v int64)`

SetClientSecretRetentionPeriod sets ClientSecretRetentionPeriod field to given value.

### HasClientSecretRetentionPeriod

`func (o *Client) HasClientSecretRetentionPeriod() bool`

HasClientSecretRetentionPeriod returns a boolean if a field has been set.

### GetClientSecretChangedTime

`func (o *Client) GetClientSecretChangedTime() time.Time`

GetClientSecretChangedTime returns the ClientSecretChangedTime field if non-nil, zero value otherwise.

### GetClientSecretChangedTimeOk

`func (o *Client) GetClientSecretChangedTimeOk() (*time.Time, bool)`

GetClientSecretChangedTimeOk returns a tuple with the ClientSecretChangedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretChangedTime

`func (o *Client) SetClientSecretChangedTime(v time.Time)`

SetClientSecretChangedTime sets ClientSecretChangedTime field to given value.

### HasClientSecretChangedTime

`func (o *Client) HasClientSecretChangedTime() bool`

HasClientSecretChangedTime returns a boolean if a field has been set.

### GetTokenIntrospectionSigningAlgorithm

`func (o *Client) GetTokenIntrospectionSigningAlgorithm() string`

GetTokenIntrospectionSigningAlgorithm returns the TokenIntrospectionSigningAlgorithm field if non-nil, zero value otherwise.

### GetTokenIntrospectionSigningAlgorithmOk

`func (o *Client) GetTokenIntrospectionSigningAlgorithmOk() (*string, bool)`

GetTokenIntrospectionSigningAlgorithmOk returns a tuple with the TokenIntrospectionSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIntrospectionSigningAlgorithm

`func (o *Client) SetTokenIntrospectionSigningAlgorithm(v string)`

SetTokenIntrospectionSigningAlgorithm sets TokenIntrospectionSigningAlgorithm field to given value.

### HasTokenIntrospectionSigningAlgorithm

`func (o *Client) HasTokenIntrospectionSigningAlgorithm() bool`

HasTokenIntrospectionSigningAlgorithm returns a boolean if a field has been set.

### GetTokenIntrospectionEncryptionAlgorithm

`func (o *Client) GetTokenIntrospectionEncryptionAlgorithm() string`

GetTokenIntrospectionEncryptionAlgorithm returns the TokenIntrospectionEncryptionAlgorithm field if non-nil, zero value otherwise.

### GetTokenIntrospectionEncryptionAlgorithmOk

`func (o *Client) GetTokenIntrospectionEncryptionAlgorithmOk() (*string, bool)`

GetTokenIntrospectionEncryptionAlgorithmOk returns a tuple with the TokenIntrospectionEncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIntrospectionEncryptionAlgorithm

`func (o *Client) SetTokenIntrospectionEncryptionAlgorithm(v string)`

SetTokenIntrospectionEncryptionAlgorithm sets TokenIntrospectionEncryptionAlgorithm field to given value.

### HasTokenIntrospectionEncryptionAlgorithm

`func (o *Client) HasTokenIntrospectionEncryptionAlgorithm() bool`

HasTokenIntrospectionEncryptionAlgorithm returns a boolean if a field has been set.

### GetTokenIntrospectionContentEncryptionAlgorithm

`func (o *Client) GetTokenIntrospectionContentEncryptionAlgorithm() string`

GetTokenIntrospectionContentEncryptionAlgorithm returns the TokenIntrospectionContentEncryptionAlgorithm field if non-nil, zero value otherwise.

### GetTokenIntrospectionContentEncryptionAlgorithmOk

`func (o *Client) GetTokenIntrospectionContentEncryptionAlgorithmOk() (*string, bool)`

GetTokenIntrospectionContentEncryptionAlgorithmOk returns a tuple with the TokenIntrospectionContentEncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIntrospectionContentEncryptionAlgorithm

`func (o *Client) SetTokenIntrospectionContentEncryptionAlgorithm(v string)`

SetTokenIntrospectionContentEncryptionAlgorithm sets TokenIntrospectionContentEncryptionAlgorithm field to given value.

### HasTokenIntrospectionContentEncryptionAlgorithm

`func (o *Client) HasTokenIntrospectionContentEncryptionAlgorithm() bool`

HasTokenIntrospectionContentEncryptionAlgorithm returns a boolean if a field has been set.

### GetJwtSecuredAuthorizationResponseModeSigningAlgorithm

`func (o *Client) GetJwtSecuredAuthorizationResponseModeSigningAlgorithm() string`

GetJwtSecuredAuthorizationResponseModeSigningAlgorithm returns the JwtSecuredAuthorizationResponseModeSigningAlgorithm field if non-nil, zero value otherwise.

### GetJwtSecuredAuthorizationResponseModeSigningAlgorithmOk

`func (o *Client) GetJwtSecuredAuthorizationResponseModeSigningAlgorithmOk() (*string, bool)`

GetJwtSecuredAuthorizationResponseModeSigningAlgorithmOk returns a tuple with the JwtSecuredAuthorizationResponseModeSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSecuredAuthorizationResponseModeSigningAlgorithm

`func (o *Client) SetJwtSecuredAuthorizationResponseModeSigningAlgorithm(v string)`

SetJwtSecuredAuthorizationResponseModeSigningAlgorithm sets JwtSecuredAuthorizationResponseModeSigningAlgorithm field to given value.

### HasJwtSecuredAuthorizationResponseModeSigningAlgorithm

`func (o *Client) HasJwtSecuredAuthorizationResponseModeSigningAlgorithm() bool`

HasJwtSecuredAuthorizationResponseModeSigningAlgorithm returns a boolean if a field has been set.

### GetJwtSecuredAuthorizationResponseModeEncryptionAlgorithm

`func (o *Client) GetJwtSecuredAuthorizationResponseModeEncryptionAlgorithm() string`

GetJwtSecuredAuthorizationResponseModeEncryptionAlgorithm returns the JwtSecuredAuthorizationResponseModeEncryptionAlgorithm field if non-nil, zero value otherwise.

### GetJwtSecuredAuthorizationResponseModeEncryptionAlgorithmOk

`func (o *Client) GetJwtSecuredAuthorizationResponseModeEncryptionAlgorithmOk() (*string, bool)`

GetJwtSecuredAuthorizationResponseModeEncryptionAlgorithmOk returns a tuple with the JwtSecuredAuthorizationResponseModeEncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSecuredAuthorizationResponseModeEncryptionAlgorithm

`func (o *Client) SetJwtSecuredAuthorizationResponseModeEncryptionAlgorithm(v string)`

SetJwtSecuredAuthorizationResponseModeEncryptionAlgorithm sets JwtSecuredAuthorizationResponseModeEncryptionAlgorithm field to given value.

### HasJwtSecuredAuthorizationResponseModeEncryptionAlgorithm

`func (o *Client) HasJwtSecuredAuthorizationResponseModeEncryptionAlgorithm() bool`

HasJwtSecuredAuthorizationResponseModeEncryptionAlgorithm returns a boolean if a field has been set.

### GetJwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm

`func (o *Client) GetJwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm() string`

GetJwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm returns the JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm field if non-nil, zero value otherwise.

### GetJwtSecuredAuthorizationResponseModeContentEncryptionAlgorithmOk

`func (o *Client) GetJwtSecuredAuthorizationResponseModeContentEncryptionAlgorithmOk() (*string, bool)`

GetJwtSecuredAuthorizationResponseModeContentEncryptionAlgorithmOk returns a tuple with the JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm

`func (o *Client) SetJwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm(v string)`

SetJwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm sets JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm field to given value.

### HasJwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm

`func (o *Client) HasJwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm() bool`

HasJwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm returns a boolean if a field has been set.

### GetRequireDpop

`func (o *Client) GetRequireDpop() bool`

GetRequireDpop returns the RequireDpop field if non-nil, zero value otherwise.

### GetRequireDpopOk

`func (o *Client) GetRequireDpopOk() (*bool, bool)`

GetRequireDpopOk returns a tuple with the RequireDpop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireDpop

`func (o *Client) SetRequireDpop(v bool)`

SetRequireDpop sets RequireDpop field to given value.

### HasRequireDpop

`func (o *Client) HasRequireDpop() bool`

HasRequireDpop returns a boolean if a field has been set.

### GetRequireOfflineAccessScopeToIssueRefreshTokens

`func (o *Client) GetRequireOfflineAccessScopeToIssueRefreshTokens() string`

GetRequireOfflineAccessScopeToIssueRefreshTokens returns the RequireOfflineAccessScopeToIssueRefreshTokens field if non-nil, zero value otherwise.

### GetRequireOfflineAccessScopeToIssueRefreshTokensOk

`func (o *Client) GetRequireOfflineAccessScopeToIssueRefreshTokensOk() (*string, bool)`

GetRequireOfflineAccessScopeToIssueRefreshTokensOk returns a tuple with the RequireOfflineAccessScopeToIssueRefreshTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOfflineAccessScopeToIssueRefreshTokens

`func (o *Client) SetRequireOfflineAccessScopeToIssueRefreshTokens(v string)`

SetRequireOfflineAccessScopeToIssueRefreshTokens sets RequireOfflineAccessScopeToIssueRefreshTokens field to given value.

### HasRequireOfflineAccessScopeToIssueRefreshTokens

`func (o *Client) HasRequireOfflineAccessScopeToIssueRefreshTokens() bool`

HasRequireOfflineAccessScopeToIssueRefreshTokens returns a boolean if a field has been set.

### GetOfflineAccessRequireConsentPrompt

`func (o *Client) GetOfflineAccessRequireConsentPrompt() string`

GetOfflineAccessRequireConsentPrompt returns the OfflineAccessRequireConsentPrompt field if non-nil, zero value otherwise.

### GetOfflineAccessRequireConsentPromptOk

`func (o *Client) GetOfflineAccessRequireConsentPromptOk() (*string, bool)`

GetOfflineAccessRequireConsentPromptOk returns a tuple with the OfflineAccessRequireConsentPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineAccessRequireConsentPrompt

`func (o *Client) SetOfflineAccessRequireConsentPrompt(v string)`

SetOfflineAccessRequireConsentPrompt sets OfflineAccessRequireConsentPrompt field to given value.

### HasOfflineAccessRequireConsentPrompt

`func (o *Client) HasOfflineAccessRequireConsentPrompt() bool`

HasOfflineAccessRequireConsentPrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


