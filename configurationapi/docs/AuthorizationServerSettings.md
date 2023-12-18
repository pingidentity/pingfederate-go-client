# AuthorizationServerSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultScopeDescription** | **string** | The default scope description. | 
**Scopes** | Pointer to [**[]ScopeEntry**](ScopeEntry.md) | The list of common scopes. | [optional] 
**ScopeGroups** | Pointer to [**[]ScopeGroupEntry**](ScopeGroupEntry.md) | The list of common scope groups. | [optional] 
**ExclusiveScopes** | Pointer to [**[]ScopeEntry**](ScopeEntry.md) | The list of exclusive scopes. | [optional] 
**ExclusiveScopeGroups** | Pointer to [**[]ScopeGroupEntry**](ScopeGroupEntry.md) | The list of exclusive scope groups. | [optional] 
**AuthorizationCodeTimeout** | **int64** | The authorization code timeout, in seconds. | 
**AuthorizationCodeEntropy** | **int64** | The authorization code entropy, in bytes. | 
**DisallowPlainPKCE** | Pointer to **bool** | Determines whether PKCE&#39;s &#39;plain&#39; code challenge method will be disallowed. The default value is false. | [optional] 
**IncludeIssuerInAuthorizationResponse** | Pointer to **bool** | Determines whether the authorization server&#39;s issuer value is added to the authorization response or not. The default value is false. | [optional] 
**TrackUserSessionsForLogout** | Pointer to **bool** | Determines whether user sessions are tracked for logout. If this property is not provided on a PUT, the setting is left unchanged. | [optional] 
**TokenEndpointBaseUrl** | Pointer to **string** | The token endpoint base URL used to validate the &#39;aud&#39; claim during Private Key JWT Client Authentication. | [optional] 
**PersistentGrantLifetime** | Pointer to **int64** | The persistent grant lifetime. The default value is indefinite. -1 indicates an indefinite amount of time. | [optional] 
**PersistentGrantLifetimeUnit** | Pointer to **string** | The persistent grant lifetime unit. | [optional] 
**PersistentGrantIdleTimeout** | Pointer to **int64** | The persistent grant idle timeout. The default value is 30 (days). -1 indicates an indefinite amount of time. | [optional] 
**PersistentGrantIdleTimeoutTimeUnit** | Pointer to **string** | The persistent grant idle timeout time unit. | [optional] 
**RefreshTokenLength** | **int64** | The refresh token length in number of characters. | 
**RollRefreshTokenValues** | Pointer to **bool** | The roll refresh token values default policy. The default value is true. | [optional] 
**RefreshTokenRollingGracePeriod** | Pointer to **int64** | The grace period that a rolled refresh token remains valid in seconds. The default value is 60. | [optional] 
**RefreshRollingInterval** | **int64** | The minimum interval to roll refresh tokens, in hours. | 
**PersistentGrantReuseGrantTypes** | Pointer to **[]string** | The grant types that the OAuth AS can reuse rather than creating a new grant for each request. Only &#39;IMPLICIT&#39; or &#39;AUTHORIZATION_CODE&#39; or &#39;RESOURCE_OWNER_CREDENTIALS&#39; are valid grant types. | [optional] 
**PersistentGrantContract** | Pointer to [**PersistentGrantContract**](PersistentGrantContract.md) |  | [optional] 
**BypassAuthorizationForApprovedGrants** | Pointer to **bool** | Bypass authorization for previously approved persistent grants. The default value is false. | [optional] 
**AllowUnidentifiedClientROCreds** | Pointer to **bool** | Allow unidentified clients to request resource owner password credentials grants. The default value is false. | [optional] 
**AllowUnidentifiedClientExtensionGrants** | Pointer to **bool** | Allow unidentified clients to request extension grants. The default value is false. | [optional] 
**AdminWebServicePcvRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**AtmIdForOAuthGrantManagement** | Pointer to **string** | The ID of the Access Token Manager used for OAuth enabled grant management. | [optional] 
**ScopeForOAuthGrantManagement** | Pointer to **string** | The OAuth scope to validate when accessing grant management service. | [optional] 
**AllowedOrigins** | Pointer to **[]string** | The list of allowed origins. | [optional] 
**UserAuthorizationUrl** | Pointer to **string** | The URL used to generate &#39;verification_url&#39; and &#39;verification_url_complete&#39; values in a Device Authorization request | [optional] 
**RegisteredAuthorizationPath** | **string** | The Registered Authorization Path is concatenated to PingFederate base URL to generate &#39;verification_url&#39; and &#39;verification_url_complete&#39; values in a Device Authorization request. PingFederate listens to this path if specified | 
**PendingAuthorizationTimeout** | **int64** | The &#39;device_code&#39; and &#39;user_code&#39; timeout, in seconds. | 
**DevicePollingInterval** | **int64** | The amount of time client should wait between polling requests, in seconds. | 
**ActivationCodeCheckMode** | Pointer to **string** | Determines whether the user is prompted to enter or confirm the activation code after authenticating or before. The default is AFTER_AUTHENTICATION. | [optional] 
**BypassActivationCodeConfirmation** | **bool** | Indicates if the Activation Code Confirmation page should be bypassed if &#39;verification_url_complete&#39; is used by the end user to authorize a device. | 
**UserAuthorizationConsentPageSetting** | Pointer to **string** | User Authorization Consent Page setting to use PingFederate&#39;s internal consent page or an external system | [optional] 
**UserAuthorizationConsentAdapter** | Pointer to **string** | Adapter ID of the external consent adapter to be used for the consent page user interface. | [optional] 
**ApprovedScopesAttribute** | Pointer to **string** | Attribute from the external consent adapter&#39;s contract, intended for storing approved scopes returned by the external consent page. | [optional] 
**ApprovedAuthorizationDetailAttribute** | Pointer to **string** | Attribute from the external consent adapter&#39;s contract, intended for storing approved authorization details returned by the external consent page. | [optional] 
**ParReferenceTimeout** | Pointer to **int64** | The timeout, in seconds, of the pushed authorization request reference. The default value is 60. | [optional] 
**ParReferenceLength** | Pointer to **int64** | The entropy of pushed authorization request references, in bytes. The default value is 24. | [optional] 
**ParStatus** | Pointer to **string** | The status of pushed authorization request support. The default value is ENABLED. | [optional] 
**ClientSecretRetentionPeriod** | Pointer to **int64** | The length of time in minutes that client secrets will be retained as secondary secrets after secret change. The default value is 0, which will disable secondary client secret retention. | [optional] 
**JwtSecuredAuthorizationResponseModeLifetime** | Pointer to **int64** | The lifetime, in seconds, of the JWT Secured authorization response. The default value is 600. | [optional] 
**DpopProofRequireNonce** | Pointer to **bool** | Determines whether nonce is required in the Demonstrating Proof-of-Possession (DPoP) proof JWT. The default value is false. | [optional] 
**DpopProofLifetimeSeconds** | Pointer to **int64** | The lifetime, in seconds, of the Demonstrating Proof-of-Possession (DPoP) proof JWT. The default value is 120. | [optional] 
**DpopProofEnforceReplayPrevention** | Pointer to **bool** | Determines whether Demonstrating Proof-of-Possession (DPoP) proof JWT replay prevention is enforced. The default value is false. | [optional] 

## Methods

### NewAuthorizationServerSettings

`func NewAuthorizationServerSettings(defaultScopeDescription string, authorizationCodeTimeout int64, authorizationCodeEntropy int64, refreshTokenLength int64, refreshRollingInterval int64, registeredAuthorizationPath string, pendingAuthorizationTimeout int64, devicePollingInterval int64, bypassActivationCodeConfirmation bool, ) *AuthorizationServerSettings`

NewAuthorizationServerSettings instantiates a new AuthorizationServerSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerSettingsWithDefaults

`func NewAuthorizationServerSettingsWithDefaults() *AuthorizationServerSettings`

NewAuthorizationServerSettingsWithDefaults instantiates a new AuthorizationServerSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultScopeDescription

`func (o *AuthorizationServerSettings) GetDefaultScopeDescription() string`

GetDefaultScopeDescription returns the DefaultScopeDescription field if non-nil, zero value otherwise.

### GetDefaultScopeDescriptionOk

`func (o *AuthorizationServerSettings) GetDefaultScopeDescriptionOk() (*string, bool)`

GetDefaultScopeDescriptionOk returns a tuple with the DefaultScopeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultScopeDescription

`func (o *AuthorizationServerSettings) SetDefaultScopeDescription(v string)`

SetDefaultScopeDescription sets DefaultScopeDescription field to given value.


### GetScopes

`func (o *AuthorizationServerSettings) GetScopes() []ScopeEntry`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthorizationServerSettings) GetScopesOk() (*[]ScopeEntry, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthorizationServerSettings) SetScopes(v []ScopeEntry)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthorizationServerSettings) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetScopeGroups

`func (o *AuthorizationServerSettings) GetScopeGroups() []ScopeGroupEntry`

GetScopeGroups returns the ScopeGroups field if non-nil, zero value otherwise.

### GetScopeGroupsOk

`func (o *AuthorizationServerSettings) GetScopeGroupsOk() (*[]ScopeGroupEntry, bool)`

GetScopeGroupsOk returns a tuple with the ScopeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeGroups

`func (o *AuthorizationServerSettings) SetScopeGroups(v []ScopeGroupEntry)`

SetScopeGroups sets ScopeGroups field to given value.

### HasScopeGroups

`func (o *AuthorizationServerSettings) HasScopeGroups() bool`

HasScopeGroups returns a boolean if a field has been set.

### GetExclusiveScopes

`func (o *AuthorizationServerSettings) GetExclusiveScopes() []ScopeEntry`

GetExclusiveScopes returns the ExclusiveScopes field if non-nil, zero value otherwise.

### GetExclusiveScopesOk

`func (o *AuthorizationServerSettings) GetExclusiveScopesOk() (*[]ScopeEntry, bool)`

GetExclusiveScopesOk returns a tuple with the ExclusiveScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveScopes

`func (o *AuthorizationServerSettings) SetExclusiveScopes(v []ScopeEntry)`

SetExclusiveScopes sets ExclusiveScopes field to given value.

### HasExclusiveScopes

`func (o *AuthorizationServerSettings) HasExclusiveScopes() bool`

HasExclusiveScopes returns a boolean if a field has been set.

### GetExclusiveScopeGroups

`func (o *AuthorizationServerSettings) GetExclusiveScopeGroups() []ScopeGroupEntry`

GetExclusiveScopeGroups returns the ExclusiveScopeGroups field if non-nil, zero value otherwise.

### GetExclusiveScopeGroupsOk

`func (o *AuthorizationServerSettings) GetExclusiveScopeGroupsOk() (*[]ScopeGroupEntry, bool)`

GetExclusiveScopeGroupsOk returns a tuple with the ExclusiveScopeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveScopeGroups

`func (o *AuthorizationServerSettings) SetExclusiveScopeGroups(v []ScopeGroupEntry)`

SetExclusiveScopeGroups sets ExclusiveScopeGroups field to given value.

### HasExclusiveScopeGroups

`func (o *AuthorizationServerSettings) HasExclusiveScopeGroups() bool`

HasExclusiveScopeGroups returns a boolean if a field has been set.

### GetAuthorizationCodeTimeout

`func (o *AuthorizationServerSettings) GetAuthorizationCodeTimeout() int64`

GetAuthorizationCodeTimeout returns the AuthorizationCodeTimeout field if non-nil, zero value otherwise.

### GetAuthorizationCodeTimeoutOk

`func (o *AuthorizationServerSettings) GetAuthorizationCodeTimeoutOk() (*int64, bool)`

GetAuthorizationCodeTimeoutOk returns a tuple with the AuthorizationCodeTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCodeTimeout

`func (o *AuthorizationServerSettings) SetAuthorizationCodeTimeout(v int64)`

SetAuthorizationCodeTimeout sets AuthorizationCodeTimeout field to given value.


### GetAuthorizationCodeEntropy

`func (o *AuthorizationServerSettings) GetAuthorizationCodeEntropy() int64`

GetAuthorizationCodeEntropy returns the AuthorizationCodeEntropy field if non-nil, zero value otherwise.

### GetAuthorizationCodeEntropyOk

`func (o *AuthorizationServerSettings) GetAuthorizationCodeEntropyOk() (*int64, bool)`

GetAuthorizationCodeEntropyOk returns a tuple with the AuthorizationCodeEntropy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCodeEntropy

`func (o *AuthorizationServerSettings) SetAuthorizationCodeEntropy(v int64)`

SetAuthorizationCodeEntropy sets AuthorizationCodeEntropy field to given value.


### GetDisallowPlainPKCE

`func (o *AuthorizationServerSettings) GetDisallowPlainPKCE() bool`

GetDisallowPlainPKCE returns the DisallowPlainPKCE field if non-nil, zero value otherwise.

### GetDisallowPlainPKCEOk

`func (o *AuthorizationServerSettings) GetDisallowPlainPKCEOk() (*bool, bool)`

GetDisallowPlainPKCEOk returns a tuple with the DisallowPlainPKCE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowPlainPKCE

`func (o *AuthorizationServerSettings) SetDisallowPlainPKCE(v bool)`

SetDisallowPlainPKCE sets DisallowPlainPKCE field to given value.

### HasDisallowPlainPKCE

`func (o *AuthorizationServerSettings) HasDisallowPlainPKCE() bool`

HasDisallowPlainPKCE returns a boolean if a field has been set.

### GetIncludeIssuerInAuthorizationResponse

`func (o *AuthorizationServerSettings) GetIncludeIssuerInAuthorizationResponse() bool`

GetIncludeIssuerInAuthorizationResponse returns the IncludeIssuerInAuthorizationResponse field if non-nil, zero value otherwise.

### GetIncludeIssuerInAuthorizationResponseOk

`func (o *AuthorizationServerSettings) GetIncludeIssuerInAuthorizationResponseOk() (*bool, bool)`

GetIncludeIssuerInAuthorizationResponseOk returns a tuple with the IncludeIssuerInAuthorizationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeIssuerInAuthorizationResponse

`func (o *AuthorizationServerSettings) SetIncludeIssuerInAuthorizationResponse(v bool)`

SetIncludeIssuerInAuthorizationResponse sets IncludeIssuerInAuthorizationResponse field to given value.

### HasIncludeIssuerInAuthorizationResponse

`func (o *AuthorizationServerSettings) HasIncludeIssuerInAuthorizationResponse() bool`

HasIncludeIssuerInAuthorizationResponse returns a boolean if a field has been set.

### GetTrackUserSessionsForLogout

`func (o *AuthorizationServerSettings) GetTrackUserSessionsForLogout() bool`

GetTrackUserSessionsForLogout returns the TrackUserSessionsForLogout field if non-nil, zero value otherwise.

### GetTrackUserSessionsForLogoutOk

`func (o *AuthorizationServerSettings) GetTrackUserSessionsForLogoutOk() (*bool, bool)`

GetTrackUserSessionsForLogoutOk returns a tuple with the TrackUserSessionsForLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackUserSessionsForLogout

`func (o *AuthorizationServerSettings) SetTrackUserSessionsForLogout(v bool)`

SetTrackUserSessionsForLogout sets TrackUserSessionsForLogout field to given value.

### HasTrackUserSessionsForLogout

`func (o *AuthorizationServerSettings) HasTrackUserSessionsForLogout() bool`

HasTrackUserSessionsForLogout returns a boolean if a field has been set.

### GetTokenEndpointBaseUrl

`func (o *AuthorizationServerSettings) GetTokenEndpointBaseUrl() string`

GetTokenEndpointBaseUrl returns the TokenEndpointBaseUrl field if non-nil, zero value otherwise.

### GetTokenEndpointBaseUrlOk

`func (o *AuthorizationServerSettings) GetTokenEndpointBaseUrlOk() (*string, bool)`

GetTokenEndpointBaseUrlOk returns a tuple with the TokenEndpointBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointBaseUrl

`func (o *AuthorizationServerSettings) SetTokenEndpointBaseUrl(v string)`

SetTokenEndpointBaseUrl sets TokenEndpointBaseUrl field to given value.

### HasTokenEndpointBaseUrl

`func (o *AuthorizationServerSettings) HasTokenEndpointBaseUrl() bool`

HasTokenEndpointBaseUrl returns a boolean if a field has been set.

### GetPersistentGrantLifetime

`func (o *AuthorizationServerSettings) GetPersistentGrantLifetime() int64`

GetPersistentGrantLifetime returns the PersistentGrantLifetime field if non-nil, zero value otherwise.

### GetPersistentGrantLifetimeOk

`func (o *AuthorizationServerSettings) GetPersistentGrantLifetimeOk() (*int64, bool)`

GetPersistentGrantLifetimeOk returns a tuple with the PersistentGrantLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantLifetime

`func (o *AuthorizationServerSettings) SetPersistentGrantLifetime(v int64)`

SetPersistentGrantLifetime sets PersistentGrantLifetime field to given value.

### HasPersistentGrantLifetime

`func (o *AuthorizationServerSettings) HasPersistentGrantLifetime() bool`

HasPersistentGrantLifetime returns a boolean if a field has been set.

### GetPersistentGrantLifetimeUnit

`func (o *AuthorizationServerSettings) GetPersistentGrantLifetimeUnit() string`

GetPersistentGrantLifetimeUnit returns the PersistentGrantLifetimeUnit field if non-nil, zero value otherwise.

### GetPersistentGrantLifetimeUnitOk

`func (o *AuthorizationServerSettings) GetPersistentGrantLifetimeUnitOk() (*string, bool)`

GetPersistentGrantLifetimeUnitOk returns a tuple with the PersistentGrantLifetimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantLifetimeUnit

`func (o *AuthorizationServerSettings) SetPersistentGrantLifetimeUnit(v string)`

SetPersistentGrantLifetimeUnit sets PersistentGrantLifetimeUnit field to given value.

### HasPersistentGrantLifetimeUnit

`func (o *AuthorizationServerSettings) HasPersistentGrantLifetimeUnit() bool`

HasPersistentGrantLifetimeUnit returns a boolean if a field has been set.

### GetPersistentGrantIdleTimeout

`func (o *AuthorizationServerSettings) GetPersistentGrantIdleTimeout() int64`

GetPersistentGrantIdleTimeout returns the PersistentGrantIdleTimeout field if non-nil, zero value otherwise.

### GetPersistentGrantIdleTimeoutOk

`func (o *AuthorizationServerSettings) GetPersistentGrantIdleTimeoutOk() (*int64, bool)`

GetPersistentGrantIdleTimeoutOk returns a tuple with the PersistentGrantIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantIdleTimeout

`func (o *AuthorizationServerSettings) SetPersistentGrantIdleTimeout(v int64)`

SetPersistentGrantIdleTimeout sets PersistentGrantIdleTimeout field to given value.

### HasPersistentGrantIdleTimeout

`func (o *AuthorizationServerSettings) HasPersistentGrantIdleTimeout() bool`

HasPersistentGrantIdleTimeout returns a boolean if a field has been set.

### GetPersistentGrantIdleTimeoutTimeUnit

`func (o *AuthorizationServerSettings) GetPersistentGrantIdleTimeoutTimeUnit() string`

GetPersistentGrantIdleTimeoutTimeUnit returns the PersistentGrantIdleTimeoutTimeUnit field if non-nil, zero value otherwise.

### GetPersistentGrantIdleTimeoutTimeUnitOk

`func (o *AuthorizationServerSettings) GetPersistentGrantIdleTimeoutTimeUnitOk() (*string, bool)`

GetPersistentGrantIdleTimeoutTimeUnitOk returns a tuple with the PersistentGrantIdleTimeoutTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantIdleTimeoutTimeUnit

`func (o *AuthorizationServerSettings) SetPersistentGrantIdleTimeoutTimeUnit(v string)`

SetPersistentGrantIdleTimeoutTimeUnit sets PersistentGrantIdleTimeoutTimeUnit field to given value.

### HasPersistentGrantIdleTimeoutTimeUnit

`func (o *AuthorizationServerSettings) HasPersistentGrantIdleTimeoutTimeUnit() bool`

HasPersistentGrantIdleTimeoutTimeUnit returns a boolean if a field has been set.

### GetRefreshTokenLength

`func (o *AuthorizationServerSettings) GetRefreshTokenLength() int64`

GetRefreshTokenLength returns the RefreshTokenLength field if non-nil, zero value otherwise.

### GetRefreshTokenLengthOk

`func (o *AuthorizationServerSettings) GetRefreshTokenLengthOk() (*int64, bool)`

GetRefreshTokenLengthOk returns a tuple with the RefreshTokenLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenLength

`func (o *AuthorizationServerSettings) SetRefreshTokenLength(v int64)`

SetRefreshTokenLength sets RefreshTokenLength field to given value.


### GetRollRefreshTokenValues

`func (o *AuthorizationServerSettings) GetRollRefreshTokenValues() bool`

GetRollRefreshTokenValues returns the RollRefreshTokenValues field if non-nil, zero value otherwise.

### GetRollRefreshTokenValuesOk

`func (o *AuthorizationServerSettings) GetRollRefreshTokenValuesOk() (*bool, bool)`

GetRollRefreshTokenValuesOk returns a tuple with the RollRefreshTokenValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollRefreshTokenValues

`func (o *AuthorizationServerSettings) SetRollRefreshTokenValues(v bool)`

SetRollRefreshTokenValues sets RollRefreshTokenValues field to given value.

### HasRollRefreshTokenValues

`func (o *AuthorizationServerSettings) HasRollRefreshTokenValues() bool`

HasRollRefreshTokenValues returns a boolean if a field has been set.

### GetRefreshTokenRollingGracePeriod

`func (o *AuthorizationServerSettings) GetRefreshTokenRollingGracePeriod() int64`

GetRefreshTokenRollingGracePeriod returns the RefreshTokenRollingGracePeriod field if non-nil, zero value otherwise.

### GetRefreshTokenRollingGracePeriodOk

`func (o *AuthorizationServerSettings) GetRefreshTokenRollingGracePeriodOk() (*int64, bool)`

GetRefreshTokenRollingGracePeriodOk returns a tuple with the RefreshTokenRollingGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingGracePeriod

`func (o *AuthorizationServerSettings) SetRefreshTokenRollingGracePeriod(v int64)`

SetRefreshTokenRollingGracePeriod sets RefreshTokenRollingGracePeriod field to given value.

### HasRefreshTokenRollingGracePeriod

`func (o *AuthorizationServerSettings) HasRefreshTokenRollingGracePeriod() bool`

HasRefreshTokenRollingGracePeriod returns a boolean if a field has been set.

### GetRefreshRollingInterval

`func (o *AuthorizationServerSettings) GetRefreshRollingInterval() int64`

GetRefreshRollingInterval returns the RefreshRollingInterval field if non-nil, zero value otherwise.

### GetRefreshRollingIntervalOk

`func (o *AuthorizationServerSettings) GetRefreshRollingIntervalOk() (*int64, bool)`

GetRefreshRollingIntervalOk returns a tuple with the RefreshRollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshRollingInterval

`func (o *AuthorizationServerSettings) SetRefreshRollingInterval(v int64)`

SetRefreshRollingInterval sets RefreshRollingInterval field to given value.


### GetPersistentGrantReuseGrantTypes

`func (o *AuthorizationServerSettings) GetPersistentGrantReuseGrantTypes() []string`

GetPersistentGrantReuseGrantTypes returns the PersistentGrantReuseGrantTypes field if non-nil, zero value otherwise.

### GetPersistentGrantReuseGrantTypesOk

`func (o *AuthorizationServerSettings) GetPersistentGrantReuseGrantTypesOk() (*[]string, bool)`

GetPersistentGrantReuseGrantTypesOk returns a tuple with the PersistentGrantReuseGrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantReuseGrantTypes

`func (o *AuthorizationServerSettings) SetPersistentGrantReuseGrantTypes(v []string)`

SetPersistentGrantReuseGrantTypes sets PersistentGrantReuseGrantTypes field to given value.

### HasPersistentGrantReuseGrantTypes

`func (o *AuthorizationServerSettings) HasPersistentGrantReuseGrantTypes() bool`

HasPersistentGrantReuseGrantTypes returns a boolean if a field has been set.

### GetPersistentGrantContract

`func (o *AuthorizationServerSettings) GetPersistentGrantContract() PersistentGrantContract`

GetPersistentGrantContract returns the PersistentGrantContract field if non-nil, zero value otherwise.

### GetPersistentGrantContractOk

`func (o *AuthorizationServerSettings) GetPersistentGrantContractOk() (*PersistentGrantContract, bool)`

GetPersistentGrantContractOk returns a tuple with the PersistentGrantContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantContract

`func (o *AuthorizationServerSettings) SetPersistentGrantContract(v PersistentGrantContract)`

SetPersistentGrantContract sets PersistentGrantContract field to given value.

### HasPersistentGrantContract

`func (o *AuthorizationServerSettings) HasPersistentGrantContract() bool`

HasPersistentGrantContract returns a boolean if a field has been set.

### GetBypassAuthorizationForApprovedGrants

`func (o *AuthorizationServerSettings) GetBypassAuthorizationForApprovedGrants() bool`

GetBypassAuthorizationForApprovedGrants returns the BypassAuthorizationForApprovedGrants field if non-nil, zero value otherwise.

### GetBypassAuthorizationForApprovedGrantsOk

`func (o *AuthorizationServerSettings) GetBypassAuthorizationForApprovedGrantsOk() (*bool, bool)`

GetBypassAuthorizationForApprovedGrantsOk returns a tuple with the BypassAuthorizationForApprovedGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassAuthorizationForApprovedGrants

`func (o *AuthorizationServerSettings) SetBypassAuthorizationForApprovedGrants(v bool)`

SetBypassAuthorizationForApprovedGrants sets BypassAuthorizationForApprovedGrants field to given value.

### HasBypassAuthorizationForApprovedGrants

`func (o *AuthorizationServerSettings) HasBypassAuthorizationForApprovedGrants() bool`

HasBypassAuthorizationForApprovedGrants returns a boolean if a field has been set.

### GetAllowUnidentifiedClientROCreds

`func (o *AuthorizationServerSettings) GetAllowUnidentifiedClientROCreds() bool`

GetAllowUnidentifiedClientROCreds returns the AllowUnidentifiedClientROCreds field if non-nil, zero value otherwise.

### GetAllowUnidentifiedClientROCredsOk

`func (o *AuthorizationServerSettings) GetAllowUnidentifiedClientROCredsOk() (*bool, bool)`

GetAllowUnidentifiedClientROCredsOk returns a tuple with the AllowUnidentifiedClientROCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnidentifiedClientROCreds

`func (o *AuthorizationServerSettings) SetAllowUnidentifiedClientROCreds(v bool)`

SetAllowUnidentifiedClientROCreds sets AllowUnidentifiedClientROCreds field to given value.

### HasAllowUnidentifiedClientROCreds

`func (o *AuthorizationServerSettings) HasAllowUnidentifiedClientROCreds() bool`

HasAllowUnidentifiedClientROCreds returns a boolean if a field has been set.

### GetAllowUnidentifiedClientExtensionGrants

`func (o *AuthorizationServerSettings) GetAllowUnidentifiedClientExtensionGrants() bool`

GetAllowUnidentifiedClientExtensionGrants returns the AllowUnidentifiedClientExtensionGrants field if non-nil, zero value otherwise.

### GetAllowUnidentifiedClientExtensionGrantsOk

`func (o *AuthorizationServerSettings) GetAllowUnidentifiedClientExtensionGrantsOk() (*bool, bool)`

GetAllowUnidentifiedClientExtensionGrantsOk returns a tuple with the AllowUnidentifiedClientExtensionGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnidentifiedClientExtensionGrants

`func (o *AuthorizationServerSettings) SetAllowUnidentifiedClientExtensionGrants(v bool)`

SetAllowUnidentifiedClientExtensionGrants sets AllowUnidentifiedClientExtensionGrants field to given value.

### HasAllowUnidentifiedClientExtensionGrants

`func (o *AuthorizationServerSettings) HasAllowUnidentifiedClientExtensionGrants() bool`

HasAllowUnidentifiedClientExtensionGrants returns a boolean if a field has been set.

### GetAdminWebServicePcvRef

`func (o *AuthorizationServerSettings) GetAdminWebServicePcvRef() ResourceLink`

GetAdminWebServicePcvRef returns the AdminWebServicePcvRef field if non-nil, zero value otherwise.

### GetAdminWebServicePcvRefOk

`func (o *AuthorizationServerSettings) GetAdminWebServicePcvRefOk() (*ResourceLink, bool)`

GetAdminWebServicePcvRefOk returns a tuple with the AdminWebServicePcvRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminWebServicePcvRef

`func (o *AuthorizationServerSettings) SetAdminWebServicePcvRef(v ResourceLink)`

SetAdminWebServicePcvRef sets AdminWebServicePcvRef field to given value.

### HasAdminWebServicePcvRef

`func (o *AuthorizationServerSettings) HasAdminWebServicePcvRef() bool`

HasAdminWebServicePcvRef returns a boolean if a field has been set.

### GetAtmIdForOAuthGrantManagement

`func (o *AuthorizationServerSettings) GetAtmIdForOAuthGrantManagement() string`

GetAtmIdForOAuthGrantManagement returns the AtmIdForOAuthGrantManagement field if non-nil, zero value otherwise.

### GetAtmIdForOAuthGrantManagementOk

`func (o *AuthorizationServerSettings) GetAtmIdForOAuthGrantManagementOk() (*string, bool)`

GetAtmIdForOAuthGrantManagementOk returns a tuple with the AtmIdForOAuthGrantManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtmIdForOAuthGrantManagement

`func (o *AuthorizationServerSettings) SetAtmIdForOAuthGrantManagement(v string)`

SetAtmIdForOAuthGrantManagement sets AtmIdForOAuthGrantManagement field to given value.

### HasAtmIdForOAuthGrantManagement

`func (o *AuthorizationServerSettings) HasAtmIdForOAuthGrantManagement() bool`

HasAtmIdForOAuthGrantManagement returns a boolean if a field has been set.

### GetScopeForOAuthGrantManagement

`func (o *AuthorizationServerSettings) GetScopeForOAuthGrantManagement() string`

GetScopeForOAuthGrantManagement returns the ScopeForOAuthGrantManagement field if non-nil, zero value otherwise.

### GetScopeForOAuthGrantManagementOk

`func (o *AuthorizationServerSettings) GetScopeForOAuthGrantManagementOk() (*string, bool)`

GetScopeForOAuthGrantManagementOk returns a tuple with the ScopeForOAuthGrantManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeForOAuthGrantManagement

`func (o *AuthorizationServerSettings) SetScopeForOAuthGrantManagement(v string)`

SetScopeForOAuthGrantManagement sets ScopeForOAuthGrantManagement field to given value.

### HasScopeForOAuthGrantManagement

`func (o *AuthorizationServerSettings) HasScopeForOAuthGrantManagement() bool`

HasScopeForOAuthGrantManagement returns a boolean if a field has been set.

### GetAllowedOrigins

`func (o *AuthorizationServerSettings) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *AuthorizationServerSettings) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *AuthorizationServerSettings) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *AuthorizationServerSettings) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### GetUserAuthorizationUrl

`func (o *AuthorizationServerSettings) GetUserAuthorizationUrl() string`

GetUserAuthorizationUrl returns the UserAuthorizationUrl field if non-nil, zero value otherwise.

### GetUserAuthorizationUrlOk

`func (o *AuthorizationServerSettings) GetUserAuthorizationUrlOk() (*string, bool)`

GetUserAuthorizationUrlOk returns a tuple with the UserAuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthorizationUrl

`func (o *AuthorizationServerSettings) SetUserAuthorizationUrl(v string)`

SetUserAuthorizationUrl sets UserAuthorizationUrl field to given value.

### HasUserAuthorizationUrl

`func (o *AuthorizationServerSettings) HasUserAuthorizationUrl() bool`

HasUserAuthorizationUrl returns a boolean if a field has been set.

### GetRegisteredAuthorizationPath

`func (o *AuthorizationServerSettings) GetRegisteredAuthorizationPath() string`

GetRegisteredAuthorizationPath returns the RegisteredAuthorizationPath field if non-nil, zero value otherwise.

### GetRegisteredAuthorizationPathOk

`func (o *AuthorizationServerSettings) GetRegisteredAuthorizationPathOk() (*string, bool)`

GetRegisteredAuthorizationPathOk returns a tuple with the RegisteredAuthorizationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAuthorizationPath

`func (o *AuthorizationServerSettings) SetRegisteredAuthorizationPath(v string)`

SetRegisteredAuthorizationPath sets RegisteredAuthorizationPath field to given value.


### GetPendingAuthorizationTimeout

`func (o *AuthorizationServerSettings) GetPendingAuthorizationTimeout() int64`

GetPendingAuthorizationTimeout returns the PendingAuthorizationTimeout field if non-nil, zero value otherwise.

### GetPendingAuthorizationTimeoutOk

`func (o *AuthorizationServerSettings) GetPendingAuthorizationTimeoutOk() (*int64, bool)`

GetPendingAuthorizationTimeoutOk returns a tuple with the PendingAuthorizationTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAuthorizationTimeout

`func (o *AuthorizationServerSettings) SetPendingAuthorizationTimeout(v int64)`

SetPendingAuthorizationTimeout sets PendingAuthorizationTimeout field to given value.


### GetDevicePollingInterval

`func (o *AuthorizationServerSettings) GetDevicePollingInterval() int64`

GetDevicePollingInterval returns the DevicePollingInterval field if non-nil, zero value otherwise.

### GetDevicePollingIntervalOk

`func (o *AuthorizationServerSettings) GetDevicePollingIntervalOk() (*int64, bool)`

GetDevicePollingIntervalOk returns a tuple with the DevicePollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePollingInterval

`func (o *AuthorizationServerSettings) SetDevicePollingInterval(v int64)`

SetDevicePollingInterval sets DevicePollingInterval field to given value.


### GetActivationCodeCheckMode

`func (o *AuthorizationServerSettings) GetActivationCodeCheckMode() string`

GetActivationCodeCheckMode returns the ActivationCodeCheckMode field if non-nil, zero value otherwise.

### GetActivationCodeCheckModeOk

`func (o *AuthorizationServerSettings) GetActivationCodeCheckModeOk() (*string, bool)`

GetActivationCodeCheckModeOk returns a tuple with the ActivationCodeCheckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCodeCheckMode

`func (o *AuthorizationServerSettings) SetActivationCodeCheckMode(v string)`

SetActivationCodeCheckMode sets ActivationCodeCheckMode field to given value.

### HasActivationCodeCheckMode

`func (o *AuthorizationServerSettings) HasActivationCodeCheckMode() bool`

HasActivationCodeCheckMode returns a boolean if a field has been set.

### GetBypassActivationCodeConfirmation

`func (o *AuthorizationServerSettings) GetBypassActivationCodeConfirmation() bool`

GetBypassActivationCodeConfirmation returns the BypassActivationCodeConfirmation field if non-nil, zero value otherwise.

### GetBypassActivationCodeConfirmationOk

`func (o *AuthorizationServerSettings) GetBypassActivationCodeConfirmationOk() (*bool, bool)`

GetBypassActivationCodeConfirmationOk returns a tuple with the BypassActivationCodeConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassActivationCodeConfirmation

`func (o *AuthorizationServerSettings) SetBypassActivationCodeConfirmation(v bool)`

SetBypassActivationCodeConfirmation sets BypassActivationCodeConfirmation field to given value.


### GetUserAuthorizationConsentPageSetting

`func (o *AuthorizationServerSettings) GetUserAuthorizationConsentPageSetting() string`

GetUserAuthorizationConsentPageSetting returns the UserAuthorizationConsentPageSetting field if non-nil, zero value otherwise.

### GetUserAuthorizationConsentPageSettingOk

`func (o *AuthorizationServerSettings) GetUserAuthorizationConsentPageSettingOk() (*string, bool)`

GetUserAuthorizationConsentPageSettingOk returns a tuple with the UserAuthorizationConsentPageSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthorizationConsentPageSetting

`func (o *AuthorizationServerSettings) SetUserAuthorizationConsentPageSetting(v string)`

SetUserAuthorizationConsentPageSetting sets UserAuthorizationConsentPageSetting field to given value.

### HasUserAuthorizationConsentPageSetting

`func (o *AuthorizationServerSettings) HasUserAuthorizationConsentPageSetting() bool`

HasUserAuthorizationConsentPageSetting returns a boolean if a field has been set.

### GetUserAuthorizationConsentAdapter

`func (o *AuthorizationServerSettings) GetUserAuthorizationConsentAdapter() string`

GetUserAuthorizationConsentAdapter returns the UserAuthorizationConsentAdapter field if non-nil, zero value otherwise.

### GetUserAuthorizationConsentAdapterOk

`func (o *AuthorizationServerSettings) GetUserAuthorizationConsentAdapterOk() (*string, bool)`

GetUserAuthorizationConsentAdapterOk returns a tuple with the UserAuthorizationConsentAdapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthorizationConsentAdapter

`func (o *AuthorizationServerSettings) SetUserAuthorizationConsentAdapter(v string)`

SetUserAuthorizationConsentAdapter sets UserAuthorizationConsentAdapter field to given value.

### HasUserAuthorizationConsentAdapter

`func (o *AuthorizationServerSettings) HasUserAuthorizationConsentAdapter() bool`

HasUserAuthorizationConsentAdapter returns a boolean if a field has been set.

### GetApprovedScopesAttribute

`func (o *AuthorizationServerSettings) GetApprovedScopesAttribute() string`

GetApprovedScopesAttribute returns the ApprovedScopesAttribute field if non-nil, zero value otherwise.

### GetApprovedScopesAttributeOk

`func (o *AuthorizationServerSettings) GetApprovedScopesAttributeOk() (*string, bool)`

GetApprovedScopesAttributeOk returns a tuple with the ApprovedScopesAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedScopesAttribute

`func (o *AuthorizationServerSettings) SetApprovedScopesAttribute(v string)`

SetApprovedScopesAttribute sets ApprovedScopesAttribute field to given value.

### HasApprovedScopesAttribute

`func (o *AuthorizationServerSettings) HasApprovedScopesAttribute() bool`

HasApprovedScopesAttribute returns a boolean if a field has been set.

### GetApprovedAuthorizationDetailAttribute

`func (o *AuthorizationServerSettings) GetApprovedAuthorizationDetailAttribute() string`

GetApprovedAuthorizationDetailAttribute returns the ApprovedAuthorizationDetailAttribute field if non-nil, zero value otherwise.

### GetApprovedAuthorizationDetailAttributeOk

`func (o *AuthorizationServerSettings) GetApprovedAuthorizationDetailAttributeOk() (*string, bool)`

GetApprovedAuthorizationDetailAttributeOk returns a tuple with the ApprovedAuthorizationDetailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAuthorizationDetailAttribute

`func (o *AuthorizationServerSettings) SetApprovedAuthorizationDetailAttribute(v string)`

SetApprovedAuthorizationDetailAttribute sets ApprovedAuthorizationDetailAttribute field to given value.

### HasApprovedAuthorizationDetailAttribute

`func (o *AuthorizationServerSettings) HasApprovedAuthorizationDetailAttribute() bool`

HasApprovedAuthorizationDetailAttribute returns a boolean if a field has been set.

### GetParReferenceTimeout

`func (o *AuthorizationServerSettings) GetParReferenceTimeout() int64`

GetParReferenceTimeout returns the ParReferenceTimeout field if non-nil, zero value otherwise.

### GetParReferenceTimeoutOk

`func (o *AuthorizationServerSettings) GetParReferenceTimeoutOk() (*int64, bool)`

GetParReferenceTimeoutOk returns a tuple with the ParReferenceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParReferenceTimeout

`func (o *AuthorizationServerSettings) SetParReferenceTimeout(v int64)`

SetParReferenceTimeout sets ParReferenceTimeout field to given value.

### HasParReferenceTimeout

`func (o *AuthorizationServerSettings) HasParReferenceTimeout() bool`

HasParReferenceTimeout returns a boolean if a field has been set.

### GetParReferenceLength

`func (o *AuthorizationServerSettings) GetParReferenceLength() int64`

GetParReferenceLength returns the ParReferenceLength field if non-nil, zero value otherwise.

### GetParReferenceLengthOk

`func (o *AuthorizationServerSettings) GetParReferenceLengthOk() (*int64, bool)`

GetParReferenceLengthOk returns a tuple with the ParReferenceLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParReferenceLength

`func (o *AuthorizationServerSettings) SetParReferenceLength(v int64)`

SetParReferenceLength sets ParReferenceLength field to given value.

### HasParReferenceLength

`func (o *AuthorizationServerSettings) HasParReferenceLength() bool`

HasParReferenceLength returns a boolean if a field has been set.

### GetParStatus

`func (o *AuthorizationServerSettings) GetParStatus() string`

GetParStatus returns the ParStatus field if non-nil, zero value otherwise.

### GetParStatusOk

`func (o *AuthorizationServerSettings) GetParStatusOk() (*string, bool)`

GetParStatusOk returns a tuple with the ParStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParStatus

`func (o *AuthorizationServerSettings) SetParStatus(v string)`

SetParStatus sets ParStatus field to given value.

### HasParStatus

`func (o *AuthorizationServerSettings) HasParStatus() bool`

HasParStatus returns a boolean if a field has been set.

### GetClientSecretRetentionPeriod

`func (o *AuthorizationServerSettings) GetClientSecretRetentionPeriod() int64`

GetClientSecretRetentionPeriod returns the ClientSecretRetentionPeriod field if non-nil, zero value otherwise.

### GetClientSecretRetentionPeriodOk

`func (o *AuthorizationServerSettings) GetClientSecretRetentionPeriodOk() (*int64, bool)`

GetClientSecretRetentionPeriodOk returns a tuple with the ClientSecretRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretRetentionPeriod

`func (o *AuthorizationServerSettings) SetClientSecretRetentionPeriod(v int64)`

SetClientSecretRetentionPeriod sets ClientSecretRetentionPeriod field to given value.

### HasClientSecretRetentionPeriod

`func (o *AuthorizationServerSettings) HasClientSecretRetentionPeriod() bool`

HasClientSecretRetentionPeriod returns a boolean if a field has been set.

### GetJwtSecuredAuthorizationResponseModeLifetime

`func (o *AuthorizationServerSettings) GetJwtSecuredAuthorizationResponseModeLifetime() int64`

GetJwtSecuredAuthorizationResponseModeLifetime returns the JwtSecuredAuthorizationResponseModeLifetime field if non-nil, zero value otherwise.

### GetJwtSecuredAuthorizationResponseModeLifetimeOk

`func (o *AuthorizationServerSettings) GetJwtSecuredAuthorizationResponseModeLifetimeOk() (*int64, bool)`

GetJwtSecuredAuthorizationResponseModeLifetimeOk returns a tuple with the JwtSecuredAuthorizationResponseModeLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSecuredAuthorizationResponseModeLifetime

`func (o *AuthorizationServerSettings) SetJwtSecuredAuthorizationResponseModeLifetime(v int64)`

SetJwtSecuredAuthorizationResponseModeLifetime sets JwtSecuredAuthorizationResponseModeLifetime field to given value.

### HasJwtSecuredAuthorizationResponseModeLifetime

`func (o *AuthorizationServerSettings) HasJwtSecuredAuthorizationResponseModeLifetime() bool`

HasJwtSecuredAuthorizationResponseModeLifetime returns a boolean if a field has been set.

### GetDpopProofRequireNonce

`func (o *AuthorizationServerSettings) GetDpopProofRequireNonce() bool`

GetDpopProofRequireNonce returns the DpopProofRequireNonce field if non-nil, zero value otherwise.

### GetDpopProofRequireNonceOk

`func (o *AuthorizationServerSettings) GetDpopProofRequireNonceOk() (*bool, bool)`

GetDpopProofRequireNonceOk returns a tuple with the DpopProofRequireNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopProofRequireNonce

`func (o *AuthorizationServerSettings) SetDpopProofRequireNonce(v bool)`

SetDpopProofRequireNonce sets DpopProofRequireNonce field to given value.

### HasDpopProofRequireNonce

`func (o *AuthorizationServerSettings) HasDpopProofRequireNonce() bool`

HasDpopProofRequireNonce returns a boolean if a field has been set.

### GetDpopProofLifetimeSeconds

`func (o *AuthorizationServerSettings) GetDpopProofLifetimeSeconds() int64`

GetDpopProofLifetimeSeconds returns the DpopProofLifetimeSeconds field if non-nil, zero value otherwise.

### GetDpopProofLifetimeSecondsOk

`func (o *AuthorizationServerSettings) GetDpopProofLifetimeSecondsOk() (*int64, bool)`

GetDpopProofLifetimeSecondsOk returns a tuple with the DpopProofLifetimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopProofLifetimeSeconds

`func (o *AuthorizationServerSettings) SetDpopProofLifetimeSeconds(v int64)`

SetDpopProofLifetimeSeconds sets DpopProofLifetimeSeconds field to given value.

### HasDpopProofLifetimeSeconds

`func (o *AuthorizationServerSettings) HasDpopProofLifetimeSeconds() bool`

HasDpopProofLifetimeSeconds returns a boolean if a field has been set.

### GetDpopProofEnforceReplayPrevention

`func (o *AuthorizationServerSettings) GetDpopProofEnforceReplayPrevention() bool`

GetDpopProofEnforceReplayPrevention returns the DpopProofEnforceReplayPrevention field if non-nil, zero value otherwise.

### GetDpopProofEnforceReplayPreventionOk

`func (o *AuthorizationServerSettings) GetDpopProofEnforceReplayPreventionOk() (*bool, bool)`

GetDpopProofEnforceReplayPreventionOk returns a tuple with the DpopProofEnforceReplayPrevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopProofEnforceReplayPrevention

`func (o *AuthorizationServerSettings) SetDpopProofEnforceReplayPrevention(v bool)`

SetDpopProofEnforceReplayPrevention sets DpopProofEnforceReplayPrevention field to given value.

### HasDpopProofEnforceReplayPrevention

`func (o *AuthorizationServerSettings) HasDpopProofEnforceReplayPrevention() bool`

HasDpopProofEnforceReplayPrevention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


