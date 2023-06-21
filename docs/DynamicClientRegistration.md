# DynamicClientRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitialAccessTokenScope** | Pointer to **string** | The initial access token to prevent unwanted client registrations. | [optional] 
**RestrictCommonScopes** | Pointer to **bool** | Restrict common scopes. | [optional] 
**RestrictedCommonScopes** | Pointer to **[]string** | The common scopes to restrict. | [optional] 
**AllowedExclusiveScopes** | Pointer to **[]string** | The exclusive scopes to allow. | [optional] 
**AllowedAuthorizationDetailTypes** | Pointer to **[]string** | The authorization detail types to allow. | [optional] 
**EnforceReplayPrevention** | Pointer to **bool** | Enforce replay prevention. | [optional] 
**RequireSignedRequests** | Pointer to **bool** | Require signed requests. | [optional] 
**DefaultAccessTokenManagerRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**RestrictToDefaultAccessTokenManager** | Pointer to **bool** | Determines whether the client is restricted to using only its default access token manager. The default is false. | [optional] 
**PersistentGrantExpirationType** | Pointer to **string** | Allows an administrator to override the Persistent Grant Lifetime set globally for the OAuth AS. Defaults to SERVER_DEFAULT. | [optional] 
**PersistentGrantExpirationTime** | Pointer to **int64** | The persistent grant expiration time. | [optional] 
**PersistentGrantExpirationTimeUnit** | Pointer to **string** | The persistent grant expiration time unit. | [optional] 
**PersistentGrantIdleTimeoutType** | Pointer to **string** | Allows an administrator to override the Persistent Grant Idle Timeout set globally for the OAuth AS. Defaults to SERVER_DEFAULT. | [optional] 
**PersistentGrantIdleTimeout** | Pointer to **int64** | The persistent grant idle timeout. | [optional] 
**PersistentGrantIdleTimeoutTimeUnit** | Pointer to **string** | The persistent grant idle timeout time unit. | [optional] 
**ClientCertIssuerType** | Pointer to **string** | Client TLS Certificate Issuer Type. | [optional] 
**ClientCertIssuerRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**RefreshRolling** | Pointer to **string** | Use ROLL or DONT_ROLL to override the Roll Refresh Token Values setting on the Authorization Server Settings. SERVER_DEFAULT will default to the Roll Refresh Token Values setting on the Authorization Server Setting screen. Defaults to SERVER_DEFAULT. | [optional] 
**RefreshTokenRollingIntervalType** | Pointer to **string** | Use OVERRIDE_SERVER_DEFAULT to override the Refresh Token Rolling Interval value on the Authorization Server Settings. SERVER_DEFAULT will default to the Refresh Token Rolling Interval value on the Authorization Server Setting. Defaults to SERVER_DEFAULT. | [optional] 
**RefreshTokenRollingInterval** | Pointer to **int64** | The minimum interval to roll refresh tokens, in hours. This value will override the Refresh Token Rolling Interval Value on the Authorization Server Settings. | [optional] 
**OidcPolicy** | Pointer to [**ClientRegistrationOIDCPolicy**](ClientRegistrationOIDCPolicy.md) |  | [optional] 
**PolicyRefs** | Pointer to [**[]ResourceLink**](ResourceLink.md) | The client registration policies. | [optional] 
**DeviceFlowSettingType** | Pointer to **string** | Allows an administrator to override the Device Authorization Settings set globally for the OAuth AS. Defaults to SERVER_DEFAULT. | [optional] 
**UserAuthorizationUrlOverride** | Pointer to **string** | The URL is used as &#39;verification_url&#39; and &#39;verification_url_complete&#39; values in a Device Authorization request. | [optional] 
**PendingAuthorizationTimeoutOverride** | Pointer to **int64** | The &#39;device_code&#39; and &#39;user_code&#39; timeout, in seconds. | [optional] 
**DevicePollingIntervalOverride** | Pointer to **int64** | The amount of time client should wait between polling requests, in seconds. | [optional] 
**BypassActivationCodeConfirmationOverride** | Pointer to **bool** | Indicates if the Activation Code Confirmation page should be bypassed if &#39;verification_url_complete&#39; is used by the end user to authorize a device. | [optional] 
**RequireProofKeyForCodeExchange** | Pointer to **bool** | Determines whether Proof Key for Code Exchange (PKCE) is required for the dynamically created client. | [optional] 
**CibaPollingInterval** | Pointer to **int64** | The minimum amount of time in seconds that the Client must wait between polling requests to the token endpoint. The default is 3 seconds. | [optional] 
**CibaRequireSignedRequests** | Pointer to **bool** | Determines whether CIBA signed requests are required for this client. | [optional] 
**RequestPolicyRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**TokenExchangeProcessorPolicyRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**RotateClientSecret** | Pointer to **bool** | Rotate registration access token on dynamic client management requests. | [optional] 
**RotateRegistrationAccessToken** | Pointer to **bool** | Rotate client secret on dynamic client management requests. | [optional] 
**AllowClientDelete** | Pointer to **bool** | Allow client deletion from dynamic client management. | [optional] 
**DisableRegistrationAccessTokens** | Pointer to **bool** | Disable registration access tokens. Local standards may mandate different registration access token requirements. If applicable, implement custom validation and enforcement rules using the DynamicClientRegistrationPlugin interface from the PingFederate SDK, configure the client registration policies (policyRefs), and set this property (disableRegistrationAccessTokens) to true. CAUTION: When the disableRegistrationAccessTokens property is set to true, all clients, not just the ones created using the Dynamic Client Registration protocol, are vulnerable to unrestricted retrievals, updates (including modifications to the client authentication scheme and redirect URIs), and deletes at the /as/clients.oauth2 endpoint unless one or more client registration policies are in place to protect against unauthorized attempts. | [optional] 
**RefreshTokenRollingGracePeriodType** | Pointer to **string** | When specified, it overrides the global Refresh Token Grace Period defined in the Authorization Server Settings. The default value is SERVER_DEFAULT | [optional] 
**RefreshTokenRollingGracePeriod** | Pointer to **int64** | The grace period that a rolled refresh token remains valid in seconds. | [optional] 
**RetainClientSecret** | Pointer to **bool** | Temporarily retain the old client secret on client secret change. | [optional] 
**ClientSecretRetentionPeriodType** | Pointer to **string** | Use OVERRIDE_SERVER_DEFAULT to override the Client Secret Retention Period value on the Authorization Server Settings. SERVER_DEFAULT will default to the Client Secret Retention Period value on the Authorization Server Setting. Defaults to SERVER_DEFAULT. | [optional] 
**ClientSecretRetentionPeriodOverride** | Pointer to **int64** | The length of time in minutes that client secrets will be retained as secondary secrets after secret change. The default value is 0, which will disable secondary client secret retention. This value will override the Client Secret Retention Period value on the Authorization Server Settings. | [optional] 
**RequireJwtSecuredAuthorizationResponseMode** | Pointer to **bool** | Determines whether JWT Secured authorization response mode is required when initiating an authorization request. The default is false. | [optional] 

## Methods

### NewDynamicClientRegistration

`func NewDynamicClientRegistration() *DynamicClientRegistration`

NewDynamicClientRegistration instantiates a new DynamicClientRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicClientRegistrationWithDefaults

`func NewDynamicClientRegistrationWithDefaults() *DynamicClientRegistration`

NewDynamicClientRegistrationWithDefaults instantiates a new DynamicClientRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitialAccessTokenScope

`func (o *DynamicClientRegistration) GetInitialAccessTokenScope() string`

GetInitialAccessTokenScope returns the InitialAccessTokenScope field if non-nil, zero value otherwise.

### GetInitialAccessTokenScopeOk

`func (o *DynamicClientRegistration) GetInitialAccessTokenScopeOk() (*string, bool)`

GetInitialAccessTokenScopeOk returns a tuple with the InitialAccessTokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAccessTokenScope

`func (o *DynamicClientRegistration) SetInitialAccessTokenScope(v string)`

SetInitialAccessTokenScope sets InitialAccessTokenScope field to given value.

### HasInitialAccessTokenScope

`func (o *DynamicClientRegistration) HasInitialAccessTokenScope() bool`

HasInitialAccessTokenScope returns a boolean if a field has been set.

### GetRestrictCommonScopes

`func (o *DynamicClientRegistration) GetRestrictCommonScopes() bool`

GetRestrictCommonScopes returns the RestrictCommonScopes field if non-nil, zero value otherwise.

### GetRestrictCommonScopesOk

`func (o *DynamicClientRegistration) GetRestrictCommonScopesOk() (*bool, bool)`

GetRestrictCommonScopesOk returns a tuple with the RestrictCommonScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictCommonScopes

`func (o *DynamicClientRegistration) SetRestrictCommonScopes(v bool)`

SetRestrictCommonScopes sets RestrictCommonScopes field to given value.

### HasRestrictCommonScopes

`func (o *DynamicClientRegistration) HasRestrictCommonScopes() bool`

HasRestrictCommonScopes returns a boolean if a field has been set.

### GetRestrictedCommonScopes

`func (o *DynamicClientRegistration) GetRestrictedCommonScopes() []string`

GetRestrictedCommonScopes returns the RestrictedCommonScopes field if non-nil, zero value otherwise.

### GetRestrictedCommonScopesOk

`func (o *DynamicClientRegistration) GetRestrictedCommonScopesOk() (*[]string, bool)`

GetRestrictedCommonScopesOk returns a tuple with the RestrictedCommonScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedCommonScopes

`func (o *DynamicClientRegistration) SetRestrictedCommonScopes(v []string)`

SetRestrictedCommonScopes sets RestrictedCommonScopes field to given value.

### HasRestrictedCommonScopes

`func (o *DynamicClientRegistration) HasRestrictedCommonScopes() bool`

HasRestrictedCommonScopes returns a boolean if a field has been set.

### GetAllowedExclusiveScopes

`func (o *DynamicClientRegistration) GetAllowedExclusiveScopes() []string`

GetAllowedExclusiveScopes returns the AllowedExclusiveScopes field if non-nil, zero value otherwise.

### GetAllowedExclusiveScopesOk

`func (o *DynamicClientRegistration) GetAllowedExclusiveScopesOk() (*[]string, bool)`

GetAllowedExclusiveScopesOk returns a tuple with the AllowedExclusiveScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedExclusiveScopes

`func (o *DynamicClientRegistration) SetAllowedExclusiveScopes(v []string)`

SetAllowedExclusiveScopes sets AllowedExclusiveScopes field to given value.

### HasAllowedExclusiveScopes

`func (o *DynamicClientRegistration) HasAllowedExclusiveScopes() bool`

HasAllowedExclusiveScopes returns a boolean if a field has been set.

### GetAllowedAuthorizationDetailTypes

`func (o *DynamicClientRegistration) GetAllowedAuthorizationDetailTypes() []string`

GetAllowedAuthorizationDetailTypes returns the AllowedAuthorizationDetailTypes field if non-nil, zero value otherwise.

### GetAllowedAuthorizationDetailTypesOk

`func (o *DynamicClientRegistration) GetAllowedAuthorizationDetailTypesOk() (*[]string, bool)`

GetAllowedAuthorizationDetailTypesOk returns a tuple with the AllowedAuthorizationDetailTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthorizationDetailTypes

`func (o *DynamicClientRegistration) SetAllowedAuthorizationDetailTypes(v []string)`

SetAllowedAuthorizationDetailTypes sets AllowedAuthorizationDetailTypes field to given value.

### HasAllowedAuthorizationDetailTypes

`func (o *DynamicClientRegistration) HasAllowedAuthorizationDetailTypes() bool`

HasAllowedAuthorizationDetailTypes returns a boolean if a field has been set.

### GetEnforceReplayPrevention

`func (o *DynamicClientRegistration) GetEnforceReplayPrevention() bool`

GetEnforceReplayPrevention returns the EnforceReplayPrevention field if non-nil, zero value otherwise.

### GetEnforceReplayPreventionOk

`func (o *DynamicClientRegistration) GetEnforceReplayPreventionOk() (*bool, bool)`

GetEnforceReplayPreventionOk returns a tuple with the EnforceReplayPrevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceReplayPrevention

`func (o *DynamicClientRegistration) SetEnforceReplayPrevention(v bool)`

SetEnforceReplayPrevention sets EnforceReplayPrevention field to given value.

### HasEnforceReplayPrevention

`func (o *DynamicClientRegistration) HasEnforceReplayPrevention() bool`

HasEnforceReplayPrevention returns a boolean if a field has been set.

### GetRequireSignedRequests

`func (o *DynamicClientRegistration) GetRequireSignedRequests() bool`

GetRequireSignedRequests returns the RequireSignedRequests field if non-nil, zero value otherwise.

### GetRequireSignedRequestsOk

`func (o *DynamicClientRegistration) GetRequireSignedRequestsOk() (*bool, bool)`

GetRequireSignedRequestsOk returns a tuple with the RequireSignedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedRequests

`func (o *DynamicClientRegistration) SetRequireSignedRequests(v bool)`

SetRequireSignedRequests sets RequireSignedRequests field to given value.

### HasRequireSignedRequests

`func (o *DynamicClientRegistration) HasRequireSignedRequests() bool`

HasRequireSignedRequests returns a boolean if a field has been set.

### GetDefaultAccessTokenManagerRef

`func (o *DynamicClientRegistration) GetDefaultAccessTokenManagerRef() ResourceLink`

GetDefaultAccessTokenManagerRef returns the DefaultAccessTokenManagerRef field if non-nil, zero value otherwise.

### GetDefaultAccessTokenManagerRefOk

`func (o *DynamicClientRegistration) GetDefaultAccessTokenManagerRefOk() (*ResourceLink, bool)`

GetDefaultAccessTokenManagerRefOk returns a tuple with the DefaultAccessTokenManagerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccessTokenManagerRef

`func (o *DynamicClientRegistration) SetDefaultAccessTokenManagerRef(v ResourceLink)`

SetDefaultAccessTokenManagerRef sets DefaultAccessTokenManagerRef field to given value.

### HasDefaultAccessTokenManagerRef

`func (o *DynamicClientRegistration) HasDefaultAccessTokenManagerRef() bool`

HasDefaultAccessTokenManagerRef returns a boolean if a field has been set.

### GetRestrictToDefaultAccessTokenManager

`func (o *DynamicClientRegistration) GetRestrictToDefaultAccessTokenManager() bool`

GetRestrictToDefaultAccessTokenManager returns the RestrictToDefaultAccessTokenManager field if non-nil, zero value otherwise.

### GetRestrictToDefaultAccessTokenManagerOk

`func (o *DynamicClientRegistration) GetRestrictToDefaultAccessTokenManagerOk() (*bool, bool)`

GetRestrictToDefaultAccessTokenManagerOk returns a tuple with the RestrictToDefaultAccessTokenManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictToDefaultAccessTokenManager

`func (o *DynamicClientRegistration) SetRestrictToDefaultAccessTokenManager(v bool)`

SetRestrictToDefaultAccessTokenManager sets RestrictToDefaultAccessTokenManager field to given value.

### HasRestrictToDefaultAccessTokenManager

`func (o *DynamicClientRegistration) HasRestrictToDefaultAccessTokenManager() bool`

HasRestrictToDefaultAccessTokenManager returns a boolean if a field has been set.

### GetPersistentGrantExpirationType

`func (o *DynamicClientRegistration) GetPersistentGrantExpirationType() string`

GetPersistentGrantExpirationType returns the PersistentGrantExpirationType field if non-nil, zero value otherwise.

### GetPersistentGrantExpirationTypeOk

`func (o *DynamicClientRegistration) GetPersistentGrantExpirationTypeOk() (*string, bool)`

GetPersistentGrantExpirationTypeOk returns a tuple with the PersistentGrantExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantExpirationType

`func (o *DynamicClientRegistration) SetPersistentGrantExpirationType(v string)`

SetPersistentGrantExpirationType sets PersistentGrantExpirationType field to given value.

### HasPersistentGrantExpirationType

`func (o *DynamicClientRegistration) HasPersistentGrantExpirationType() bool`

HasPersistentGrantExpirationType returns a boolean if a field has been set.

### GetPersistentGrantExpirationTime

`func (o *DynamicClientRegistration) GetPersistentGrantExpirationTime() int64`

GetPersistentGrantExpirationTime returns the PersistentGrantExpirationTime field if non-nil, zero value otherwise.

### GetPersistentGrantExpirationTimeOk

`func (o *DynamicClientRegistration) GetPersistentGrantExpirationTimeOk() (*int64, bool)`

GetPersistentGrantExpirationTimeOk returns a tuple with the PersistentGrantExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantExpirationTime

`func (o *DynamicClientRegistration) SetPersistentGrantExpirationTime(v int64)`

SetPersistentGrantExpirationTime sets PersistentGrantExpirationTime field to given value.

### HasPersistentGrantExpirationTime

`func (o *DynamicClientRegistration) HasPersistentGrantExpirationTime() bool`

HasPersistentGrantExpirationTime returns a boolean if a field has been set.

### GetPersistentGrantExpirationTimeUnit

`func (o *DynamicClientRegistration) GetPersistentGrantExpirationTimeUnit() string`

GetPersistentGrantExpirationTimeUnit returns the PersistentGrantExpirationTimeUnit field if non-nil, zero value otherwise.

### GetPersistentGrantExpirationTimeUnitOk

`func (o *DynamicClientRegistration) GetPersistentGrantExpirationTimeUnitOk() (*string, bool)`

GetPersistentGrantExpirationTimeUnitOk returns a tuple with the PersistentGrantExpirationTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantExpirationTimeUnit

`func (o *DynamicClientRegistration) SetPersistentGrantExpirationTimeUnit(v string)`

SetPersistentGrantExpirationTimeUnit sets PersistentGrantExpirationTimeUnit field to given value.

### HasPersistentGrantExpirationTimeUnit

`func (o *DynamicClientRegistration) HasPersistentGrantExpirationTimeUnit() bool`

HasPersistentGrantExpirationTimeUnit returns a boolean if a field has been set.

### GetPersistentGrantIdleTimeoutType

`func (o *DynamicClientRegistration) GetPersistentGrantIdleTimeoutType() string`

GetPersistentGrantIdleTimeoutType returns the PersistentGrantIdleTimeoutType field if non-nil, zero value otherwise.

### GetPersistentGrantIdleTimeoutTypeOk

`func (o *DynamicClientRegistration) GetPersistentGrantIdleTimeoutTypeOk() (*string, bool)`

GetPersistentGrantIdleTimeoutTypeOk returns a tuple with the PersistentGrantIdleTimeoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantIdleTimeoutType

`func (o *DynamicClientRegistration) SetPersistentGrantIdleTimeoutType(v string)`

SetPersistentGrantIdleTimeoutType sets PersistentGrantIdleTimeoutType field to given value.

### HasPersistentGrantIdleTimeoutType

`func (o *DynamicClientRegistration) HasPersistentGrantIdleTimeoutType() bool`

HasPersistentGrantIdleTimeoutType returns a boolean if a field has been set.

### GetPersistentGrantIdleTimeout

`func (o *DynamicClientRegistration) GetPersistentGrantIdleTimeout() int64`

GetPersistentGrantIdleTimeout returns the PersistentGrantIdleTimeout field if non-nil, zero value otherwise.

### GetPersistentGrantIdleTimeoutOk

`func (o *DynamicClientRegistration) GetPersistentGrantIdleTimeoutOk() (*int64, bool)`

GetPersistentGrantIdleTimeoutOk returns a tuple with the PersistentGrantIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantIdleTimeout

`func (o *DynamicClientRegistration) SetPersistentGrantIdleTimeout(v int64)`

SetPersistentGrantIdleTimeout sets PersistentGrantIdleTimeout field to given value.

### HasPersistentGrantIdleTimeout

`func (o *DynamicClientRegistration) HasPersistentGrantIdleTimeout() bool`

HasPersistentGrantIdleTimeout returns a boolean if a field has been set.

### GetPersistentGrantIdleTimeoutTimeUnit

`func (o *DynamicClientRegistration) GetPersistentGrantIdleTimeoutTimeUnit() string`

GetPersistentGrantIdleTimeoutTimeUnit returns the PersistentGrantIdleTimeoutTimeUnit field if non-nil, zero value otherwise.

### GetPersistentGrantIdleTimeoutTimeUnitOk

`func (o *DynamicClientRegistration) GetPersistentGrantIdleTimeoutTimeUnitOk() (*string, bool)`

GetPersistentGrantIdleTimeoutTimeUnitOk returns a tuple with the PersistentGrantIdleTimeoutTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantIdleTimeoutTimeUnit

`func (o *DynamicClientRegistration) SetPersistentGrantIdleTimeoutTimeUnit(v string)`

SetPersistentGrantIdleTimeoutTimeUnit sets PersistentGrantIdleTimeoutTimeUnit field to given value.

### HasPersistentGrantIdleTimeoutTimeUnit

`func (o *DynamicClientRegistration) HasPersistentGrantIdleTimeoutTimeUnit() bool`

HasPersistentGrantIdleTimeoutTimeUnit returns a boolean if a field has been set.

### GetClientCertIssuerType

`func (o *DynamicClientRegistration) GetClientCertIssuerType() string`

GetClientCertIssuerType returns the ClientCertIssuerType field if non-nil, zero value otherwise.

### GetClientCertIssuerTypeOk

`func (o *DynamicClientRegistration) GetClientCertIssuerTypeOk() (*string, bool)`

GetClientCertIssuerTypeOk returns a tuple with the ClientCertIssuerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertIssuerType

`func (o *DynamicClientRegistration) SetClientCertIssuerType(v string)`

SetClientCertIssuerType sets ClientCertIssuerType field to given value.

### HasClientCertIssuerType

`func (o *DynamicClientRegistration) HasClientCertIssuerType() bool`

HasClientCertIssuerType returns a boolean if a field has been set.

### GetClientCertIssuerRef

`func (o *DynamicClientRegistration) GetClientCertIssuerRef() ResourceLink`

GetClientCertIssuerRef returns the ClientCertIssuerRef field if non-nil, zero value otherwise.

### GetClientCertIssuerRefOk

`func (o *DynamicClientRegistration) GetClientCertIssuerRefOk() (*ResourceLink, bool)`

GetClientCertIssuerRefOk returns a tuple with the ClientCertIssuerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertIssuerRef

`func (o *DynamicClientRegistration) SetClientCertIssuerRef(v ResourceLink)`

SetClientCertIssuerRef sets ClientCertIssuerRef field to given value.

### HasClientCertIssuerRef

`func (o *DynamicClientRegistration) HasClientCertIssuerRef() bool`

HasClientCertIssuerRef returns a boolean if a field has been set.

### GetRefreshRolling

`func (o *DynamicClientRegistration) GetRefreshRolling() string`

GetRefreshRolling returns the RefreshRolling field if non-nil, zero value otherwise.

### GetRefreshRollingOk

`func (o *DynamicClientRegistration) GetRefreshRollingOk() (*string, bool)`

GetRefreshRollingOk returns a tuple with the RefreshRolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshRolling

`func (o *DynamicClientRegistration) SetRefreshRolling(v string)`

SetRefreshRolling sets RefreshRolling field to given value.

### HasRefreshRolling

`func (o *DynamicClientRegistration) HasRefreshRolling() bool`

HasRefreshRolling returns a boolean if a field has been set.

### GetRefreshTokenRollingIntervalType

`func (o *DynamicClientRegistration) GetRefreshTokenRollingIntervalType() string`

GetRefreshTokenRollingIntervalType returns the RefreshTokenRollingIntervalType field if non-nil, zero value otherwise.

### GetRefreshTokenRollingIntervalTypeOk

`func (o *DynamicClientRegistration) GetRefreshTokenRollingIntervalTypeOk() (*string, bool)`

GetRefreshTokenRollingIntervalTypeOk returns a tuple with the RefreshTokenRollingIntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingIntervalType

`func (o *DynamicClientRegistration) SetRefreshTokenRollingIntervalType(v string)`

SetRefreshTokenRollingIntervalType sets RefreshTokenRollingIntervalType field to given value.

### HasRefreshTokenRollingIntervalType

`func (o *DynamicClientRegistration) HasRefreshTokenRollingIntervalType() bool`

HasRefreshTokenRollingIntervalType returns a boolean if a field has been set.

### GetRefreshTokenRollingInterval

`func (o *DynamicClientRegistration) GetRefreshTokenRollingInterval() int64`

GetRefreshTokenRollingInterval returns the RefreshTokenRollingInterval field if non-nil, zero value otherwise.

### GetRefreshTokenRollingIntervalOk

`func (o *DynamicClientRegistration) GetRefreshTokenRollingIntervalOk() (*int64, bool)`

GetRefreshTokenRollingIntervalOk returns a tuple with the RefreshTokenRollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingInterval

`func (o *DynamicClientRegistration) SetRefreshTokenRollingInterval(v int64)`

SetRefreshTokenRollingInterval sets RefreshTokenRollingInterval field to given value.

### HasRefreshTokenRollingInterval

`func (o *DynamicClientRegistration) HasRefreshTokenRollingInterval() bool`

HasRefreshTokenRollingInterval returns a boolean if a field has been set.

### GetOidcPolicy

`func (o *DynamicClientRegistration) GetOidcPolicy() ClientRegistrationOIDCPolicy`

GetOidcPolicy returns the OidcPolicy field if non-nil, zero value otherwise.

### GetOidcPolicyOk

`func (o *DynamicClientRegistration) GetOidcPolicyOk() (*ClientRegistrationOIDCPolicy, bool)`

GetOidcPolicyOk returns a tuple with the OidcPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcPolicy

`func (o *DynamicClientRegistration) SetOidcPolicy(v ClientRegistrationOIDCPolicy)`

SetOidcPolicy sets OidcPolicy field to given value.

### HasOidcPolicy

`func (o *DynamicClientRegistration) HasOidcPolicy() bool`

HasOidcPolicy returns a boolean if a field has been set.

### GetPolicyRefs

`func (o *DynamicClientRegistration) GetPolicyRefs() []ResourceLink`

GetPolicyRefs returns the PolicyRefs field if non-nil, zero value otherwise.

### GetPolicyRefsOk

`func (o *DynamicClientRegistration) GetPolicyRefsOk() (*[]ResourceLink, bool)`

GetPolicyRefsOk returns a tuple with the PolicyRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRefs

`func (o *DynamicClientRegistration) SetPolicyRefs(v []ResourceLink)`

SetPolicyRefs sets PolicyRefs field to given value.

### HasPolicyRefs

`func (o *DynamicClientRegistration) HasPolicyRefs() bool`

HasPolicyRefs returns a boolean if a field has been set.

### GetDeviceFlowSettingType

`func (o *DynamicClientRegistration) GetDeviceFlowSettingType() string`

GetDeviceFlowSettingType returns the DeviceFlowSettingType field if non-nil, zero value otherwise.

### GetDeviceFlowSettingTypeOk

`func (o *DynamicClientRegistration) GetDeviceFlowSettingTypeOk() (*string, bool)`

GetDeviceFlowSettingTypeOk returns a tuple with the DeviceFlowSettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFlowSettingType

`func (o *DynamicClientRegistration) SetDeviceFlowSettingType(v string)`

SetDeviceFlowSettingType sets DeviceFlowSettingType field to given value.

### HasDeviceFlowSettingType

`func (o *DynamicClientRegistration) HasDeviceFlowSettingType() bool`

HasDeviceFlowSettingType returns a boolean if a field has been set.

### GetUserAuthorizationUrlOverride

`func (o *DynamicClientRegistration) GetUserAuthorizationUrlOverride() string`

GetUserAuthorizationUrlOverride returns the UserAuthorizationUrlOverride field if non-nil, zero value otherwise.

### GetUserAuthorizationUrlOverrideOk

`func (o *DynamicClientRegistration) GetUserAuthorizationUrlOverrideOk() (*string, bool)`

GetUserAuthorizationUrlOverrideOk returns a tuple with the UserAuthorizationUrlOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthorizationUrlOverride

`func (o *DynamicClientRegistration) SetUserAuthorizationUrlOverride(v string)`

SetUserAuthorizationUrlOverride sets UserAuthorizationUrlOverride field to given value.

### HasUserAuthorizationUrlOverride

`func (o *DynamicClientRegistration) HasUserAuthorizationUrlOverride() bool`

HasUserAuthorizationUrlOverride returns a boolean if a field has been set.

### GetPendingAuthorizationTimeoutOverride

`func (o *DynamicClientRegistration) GetPendingAuthorizationTimeoutOverride() int64`

GetPendingAuthorizationTimeoutOverride returns the PendingAuthorizationTimeoutOverride field if non-nil, zero value otherwise.

### GetPendingAuthorizationTimeoutOverrideOk

`func (o *DynamicClientRegistration) GetPendingAuthorizationTimeoutOverrideOk() (*int64, bool)`

GetPendingAuthorizationTimeoutOverrideOk returns a tuple with the PendingAuthorizationTimeoutOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAuthorizationTimeoutOverride

`func (o *DynamicClientRegistration) SetPendingAuthorizationTimeoutOverride(v int64)`

SetPendingAuthorizationTimeoutOverride sets PendingAuthorizationTimeoutOverride field to given value.

### HasPendingAuthorizationTimeoutOverride

`func (o *DynamicClientRegistration) HasPendingAuthorizationTimeoutOverride() bool`

HasPendingAuthorizationTimeoutOverride returns a boolean if a field has been set.

### GetDevicePollingIntervalOverride

`func (o *DynamicClientRegistration) GetDevicePollingIntervalOverride() int64`

GetDevicePollingIntervalOverride returns the DevicePollingIntervalOverride field if non-nil, zero value otherwise.

### GetDevicePollingIntervalOverrideOk

`func (o *DynamicClientRegistration) GetDevicePollingIntervalOverrideOk() (*int64, bool)`

GetDevicePollingIntervalOverrideOk returns a tuple with the DevicePollingIntervalOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePollingIntervalOverride

`func (o *DynamicClientRegistration) SetDevicePollingIntervalOverride(v int64)`

SetDevicePollingIntervalOverride sets DevicePollingIntervalOverride field to given value.

### HasDevicePollingIntervalOverride

`func (o *DynamicClientRegistration) HasDevicePollingIntervalOverride() bool`

HasDevicePollingIntervalOverride returns a boolean if a field has been set.

### GetBypassActivationCodeConfirmationOverride

`func (o *DynamicClientRegistration) GetBypassActivationCodeConfirmationOverride() bool`

GetBypassActivationCodeConfirmationOverride returns the BypassActivationCodeConfirmationOverride field if non-nil, zero value otherwise.

### GetBypassActivationCodeConfirmationOverrideOk

`func (o *DynamicClientRegistration) GetBypassActivationCodeConfirmationOverrideOk() (*bool, bool)`

GetBypassActivationCodeConfirmationOverrideOk returns a tuple with the BypassActivationCodeConfirmationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassActivationCodeConfirmationOverride

`func (o *DynamicClientRegistration) SetBypassActivationCodeConfirmationOverride(v bool)`

SetBypassActivationCodeConfirmationOverride sets BypassActivationCodeConfirmationOverride field to given value.

### HasBypassActivationCodeConfirmationOverride

`func (o *DynamicClientRegistration) HasBypassActivationCodeConfirmationOverride() bool`

HasBypassActivationCodeConfirmationOverride returns a boolean if a field has been set.

### GetRequireProofKeyForCodeExchange

`func (o *DynamicClientRegistration) GetRequireProofKeyForCodeExchange() bool`

GetRequireProofKeyForCodeExchange returns the RequireProofKeyForCodeExchange field if non-nil, zero value otherwise.

### GetRequireProofKeyForCodeExchangeOk

`func (o *DynamicClientRegistration) GetRequireProofKeyForCodeExchangeOk() (*bool, bool)`

GetRequireProofKeyForCodeExchangeOk returns a tuple with the RequireProofKeyForCodeExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireProofKeyForCodeExchange

`func (o *DynamicClientRegistration) SetRequireProofKeyForCodeExchange(v bool)`

SetRequireProofKeyForCodeExchange sets RequireProofKeyForCodeExchange field to given value.

### HasRequireProofKeyForCodeExchange

`func (o *DynamicClientRegistration) HasRequireProofKeyForCodeExchange() bool`

HasRequireProofKeyForCodeExchange returns a boolean if a field has been set.

### GetCibaPollingInterval

`func (o *DynamicClientRegistration) GetCibaPollingInterval() int64`

GetCibaPollingInterval returns the CibaPollingInterval field if non-nil, zero value otherwise.

### GetCibaPollingIntervalOk

`func (o *DynamicClientRegistration) GetCibaPollingIntervalOk() (*int64, bool)`

GetCibaPollingIntervalOk returns a tuple with the CibaPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCibaPollingInterval

`func (o *DynamicClientRegistration) SetCibaPollingInterval(v int64)`

SetCibaPollingInterval sets CibaPollingInterval field to given value.

### HasCibaPollingInterval

`func (o *DynamicClientRegistration) HasCibaPollingInterval() bool`

HasCibaPollingInterval returns a boolean if a field has been set.

### GetCibaRequireSignedRequests

`func (o *DynamicClientRegistration) GetCibaRequireSignedRequests() bool`

GetCibaRequireSignedRequests returns the CibaRequireSignedRequests field if non-nil, zero value otherwise.

### GetCibaRequireSignedRequestsOk

`func (o *DynamicClientRegistration) GetCibaRequireSignedRequestsOk() (*bool, bool)`

GetCibaRequireSignedRequestsOk returns a tuple with the CibaRequireSignedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCibaRequireSignedRequests

`func (o *DynamicClientRegistration) SetCibaRequireSignedRequests(v bool)`

SetCibaRequireSignedRequests sets CibaRequireSignedRequests field to given value.

### HasCibaRequireSignedRequests

`func (o *DynamicClientRegistration) HasCibaRequireSignedRequests() bool`

HasCibaRequireSignedRequests returns a boolean if a field has been set.

### GetRequestPolicyRef

`func (o *DynamicClientRegistration) GetRequestPolicyRef() ResourceLink`

GetRequestPolicyRef returns the RequestPolicyRef field if non-nil, zero value otherwise.

### GetRequestPolicyRefOk

`func (o *DynamicClientRegistration) GetRequestPolicyRefOk() (*ResourceLink, bool)`

GetRequestPolicyRefOk returns a tuple with the RequestPolicyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPolicyRef

`func (o *DynamicClientRegistration) SetRequestPolicyRef(v ResourceLink)`

SetRequestPolicyRef sets RequestPolicyRef field to given value.

### HasRequestPolicyRef

`func (o *DynamicClientRegistration) HasRequestPolicyRef() bool`

HasRequestPolicyRef returns a boolean if a field has been set.

### GetTokenExchangeProcessorPolicyRef

`func (o *DynamicClientRegistration) GetTokenExchangeProcessorPolicyRef() ResourceLink`

GetTokenExchangeProcessorPolicyRef returns the TokenExchangeProcessorPolicyRef field if non-nil, zero value otherwise.

### GetTokenExchangeProcessorPolicyRefOk

`func (o *DynamicClientRegistration) GetTokenExchangeProcessorPolicyRefOk() (*ResourceLink, bool)`

GetTokenExchangeProcessorPolicyRefOk returns a tuple with the TokenExchangeProcessorPolicyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangeProcessorPolicyRef

`func (o *DynamicClientRegistration) SetTokenExchangeProcessorPolicyRef(v ResourceLink)`

SetTokenExchangeProcessorPolicyRef sets TokenExchangeProcessorPolicyRef field to given value.

### HasTokenExchangeProcessorPolicyRef

`func (o *DynamicClientRegistration) HasTokenExchangeProcessorPolicyRef() bool`

HasTokenExchangeProcessorPolicyRef returns a boolean if a field has been set.

### GetRotateClientSecret

`func (o *DynamicClientRegistration) GetRotateClientSecret() bool`

GetRotateClientSecret returns the RotateClientSecret field if non-nil, zero value otherwise.

### GetRotateClientSecretOk

`func (o *DynamicClientRegistration) GetRotateClientSecretOk() (*bool, bool)`

GetRotateClientSecretOk returns a tuple with the RotateClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateClientSecret

`func (o *DynamicClientRegistration) SetRotateClientSecret(v bool)`

SetRotateClientSecret sets RotateClientSecret field to given value.

### HasRotateClientSecret

`func (o *DynamicClientRegistration) HasRotateClientSecret() bool`

HasRotateClientSecret returns a boolean if a field has been set.

### GetRotateRegistrationAccessToken

`func (o *DynamicClientRegistration) GetRotateRegistrationAccessToken() bool`

GetRotateRegistrationAccessToken returns the RotateRegistrationAccessToken field if non-nil, zero value otherwise.

### GetRotateRegistrationAccessTokenOk

`func (o *DynamicClientRegistration) GetRotateRegistrationAccessTokenOk() (*bool, bool)`

GetRotateRegistrationAccessTokenOk returns a tuple with the RotateRegistrationAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateRegistrationAccessToken

`func (o *DynamicClientRegistration) SetRotateRegistrationAccessToken(v bool)`

SetRotateRegistrationAccessToken sets RotateRegistrationAccessToken field to given value.

### HasRotateRegistrationAccessToken

`func (o *DynamicClientRegistration) HasRotateRegistrationAccessToken() bool`

HasRotateRegistrationAccessToken returns a boolean if a field has been set.

### GetAllowClientDelete

`func (o *DynamicClientRegistration) GetAllowClientDelete() bool`

GetAllowClientDelete returns the AllowClientDelete field if non-nil, zero value otherwise.

### GetAllowClientDeleteOk

`func (o *DynamicClientRegistration) GetAllowClientDeleteOk() (*bool, bool)`

GetAllowClientDeleteOk returns a tuple with the AllowClientDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowClientDelete

`func (o *DynamicClientRegistration) SetAllowClientDelete(v bool)`

SetAllowClientDelete sets AllowClientDelete field to given value.

### HasAllowClientDelete

`func (o *DynamicClientRegistration) HasAllowClientDelete() bool`

HasAllowClientDelete returns a boolean if a field has been set.

### GetDisableRegistrationAccessTokens

`func (o *DynamicClientRegistration) GetDisableRegistrationAccessTokens() bool`

GetDisableRegistrationAccessTokens returns the DisableRegistrationAccessTokens field if non-nil, zero value otherwise.

### GetDisableRegistrationAccessTokensOk

`func (o *DynamicClientRegistration) GetDisableRegistrationAccessTokensOk() (*bool, bool)`

GetDisableRegistrationAccessTokensOk returns a tuple with the DisableRegistrationAccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRegistrationAccessTokens

`func (o *DynamicClientRegistration) SetDisableRegistrationAccessTokens(v bool)`

SetDisableRegistrationAccessTokens sets DisableRegistrationAccessTokens field to given value.

### HasDisableRegistrationAccessTokens

`func (o *DynamicClientRegistration) HasDisableRegistrationAccessTokens() bool`

HasDisableRegistrationAccessTokens returns a boolean if a field has been set.

### GetRefreshTokenRollingGracePeriodType

`func (o *DynamicClientRegistration) GetRefreshTokenRollingGracePeriodType() string`

GetRefreshTokenRollingGracePeriodType returns the RefreshTokenRollingGracePeriodType field if non-nil, zero value otherwise.

### GetRefreshTokenRollingGracePeriodTypeOk

`func (o *DynamicClientRegistration) GetRefreshTokenRollingGracePeriodTypeOk() (*string, bool)`

GetRefreshTokenRollingGracePeriodTypeOk returns a tuple with the RefreshTokenRollingGracePeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingGracePeriodType

`func (o *DynamicClientRegistration) SetRefreshTokenRollingGracePeriodType(v string)`

SetRefreshTokenRollingGracePeriodType sets RefreshTokenRollingGracePeriodType field to given value.

### HasRefreshTokenRollingGracePeriodType

`func (o *DynamicClientRegistration) HasRefreshTokenRollingGracePeriodType() bool`

HasRefreshTokenRollingGracePeriodType returns a boolean if a field has been set.

### GetRefreshTokenRollingGracePeriod

`func (o *DynamicClientRegistration) GetRefreshTokenRollingGracePeriod() int64`

GetRefreshTokenRollingGracePeriod returns the RefreshTokenRollingGracePeriod field if non-nil, zero value otherwise.

### GetRefreshTokenRollingGracePeriodOk

`func (o *DynamicClientRegistration) GetRefreshTokenRollingGracePeriodOk() (*int64, bool)`

GetRefreshTokenRollingGracePeriodOk returns a tuple with the RefreshTokenRollingGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingGracePeriod

`func (o *DynamicClientRegistration) SetRefreshTokenRollingGracePeriod(v int64)`

SetRefreshTokenRollingGracePeriod sets RefreshTokenRollingGracePeriod field to given value.

### HasRefreshTokenRollingGracePeriod

`func (o *DynamicClientRegistration) HasRefreshTokenRollingGracePeriod() bool`

HasRefreshTokenRollingGracePeriod returns a boolean if a field has been set.

### GetRetainClientSecret

`func (o *DynamicClientRegistration) GetRetainClientSecret() bool`

GetRetainClientSecret returns the RetainClientSecret field if non-nil, zero value otherwise.

### GetRetainClientSecretOk

`func (o *DynamicClientRegistration) GetRetainClientSecretOk() (*bool, bool)`

GetRetainClientSecretOk returns a tuple with the RetainClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainClientSecret

`func (o *DynamicClientRegistration) SetRetainClientSecret(v bool)`

SetRetainClientSecret sets RetainClientSecret field to given value.

### HasRetainClientSecret

`func (o *DynamicClientRegistration) HasRetainClientSecret() bool`

HasRetainClientSecret returns a boolean if a field has been set.

### GetClientSecretRetentionPeriodType

`func (o *DynamicClientRegistration) GetClientSecretRetentionPeriodType() string`

GetClientSecretRetentionPeriodType returns the ClientSecretRetentionPeriodType field if non-nil, zero value otherwise.

### GetClientSecretRetentionPeriodTypeOk

`func (o *DynamicClientRegistration) GetClientSecretRetentionPeriodTypeOk() (*string, bool)`

GetClientSecretRetentionPeriodTypeOk returns a tuple with the ClientSecretRetentionPeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretRetentionPeriodType

`func (o *DynamicClientRegistration) SetClientSecretRetentionPeriodType(v string)`

SetClientSecretRetentionPeriodType sets ClientSecretRetentionPeriodType field to given value.

### HasClientSecretRetentionPeriodType

`func (o *DynamicClientRegistration) HasClientSecretRetentionPeriodType() bool`

HasClientSecretRetentionPeriodType returns a boolean if a field has been set.

### GetClientSecretRetentionPeriodOverride

`func (o *DynamicClientRegistration) GetClientSecretRetentionPeriodOverride() int64`

GetClientSecretRetentionPeriodOverride returns the ClientSecretRetentionPeriodOverride field if non-nil, zero value otherwise.

### GetClientSecretRetentionPeriodOverrideOk

`func (o *DynamicClientRegistration) GetClientSecretRetentionPeriodOverrideOk() (*int64, bool)`

GetClientSecretRetentionPeriodOverrideOk returns a tuple with the ClientSecretRetentionPeriodOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretRetentionPeriodOverride

`func (o *DynamicClientRegistration) SetClientSecretRetentionPeriodOverride(v int64)`

SetClientSecretRetentionPeriodOverride sets ClientSecretRetentionPeriodOverride field to given value.

### HasClientSecretRetentionPeriodOverride

`func (o *DynamicClientRegistration) HasClientSecretRetentionPeriodOverride() bool`

HasClientSecretRetentionPeriodOverride returns a boolean if a field has been set.

### GetRequireJwtSecuredAuthorizationResponseMode

`func (o *DynamicClientRegistration) GetRequireJwtSecuredAuthorizationResponseMode() bool`

GetRequireJwtSecuredAuthorizationResponseMode returns the RequireJwtSecuredAuthorizationResponseMode field if non-nil, zero value otherwise.

### GetRequireJwtSecuredAuthorizationResponseModeOk

`func (o *DynamicClientRegistration) GetRequireJwtSecuredAuthorizationResponseModeOk() (*bool, bool)`

GetRequireJwtSecuredAuthorizationResponseModeOk returns a tuple with the RequireJwtSecuredAuthorizationResponseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireJwtSecuredAuthorizationResponseMode

`func (o *DynamicClientRegistration) SetRequireJwtSecuredAuthorizationResponseMode(v bool)`

SetRequireJwtSecuredAuthorizationResponseMode sets RequireJwtSecuredAuthorizationResponseMode field to given value.

### HasRequireJwtSecuredAuthorizationResponseMode

`func (o *DynamicClientRegistration) HasRequireJwtSecuredAuthorizationResponseMode() bool`

HasRequireJwtSecuredAuthorizationResponseMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


