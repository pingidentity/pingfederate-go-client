# CimdPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The policy ID. Client-supplied; must be unique. | [optional] 
**Name** | **string** | The policy name. | 
**Description** | Pointer to **string** | An optional description of the policy. | [optional] 
**Enabled** | **bool** | Whether the policy is enabled. | 
**AllowLoopbackRedirectUris** | Pointer to **bool** | When true, loopback addresses (127.0.0.0/8, ::1, localhost) are permitted in the client&#39;s redirect_uris and an incoming redirect_uri with a loopback host can match any registered loopback redirect URIs regardless of port. Default: false. | [optional] 
**PersistentGrantExpirationType** | Pointer to **string** | Allows an administrator to override the Persistent Grant Lifetime set globally for the OAuth AS. | [optional] 
**PersistentGrantExpirationTime** | Pointer to **int64** | The persistent grant expiration time. | [optional] 
**PersistentGrantExpirationTimeUnit** | Pointer to **string** | The persistent grant expiration time unit. | [optional] 
**PersistentGrantIdleTimeoutType** | Pointer to **string** | Allows an administrator to override the Persistent Grant Idle Timeout set globally for the OAuth AS. | [optional] 
**PersistentGrantIdleTimeout** | Pointer to **int64** | The persistent grant idle timeout. | [optional] 
**PersistentGrantIdleTimeoutTimeUnit** | Pointer to **string** | The persistent grant idle timeout time unit. | [optional] 
**RefreshRolling** | Pointer to **string** | Use ROLL or DONT_ROLL to override the Roll Refresh Token Values setting on the Authorization Server Settings. SERVER_DEFAULT will default to the Roll Refresh Token Values setting on the Authorization Server Setting screen. | [optional] 
**RefreshTokenRollingIntervalType** | Pointer to **string** | Use OVERRIDE_SERVER_DEFAULT to override the Refresh Token Rolling Interval value on the Authorization Server Settings. SERVER_DEFAULT will default to the Refresh Token Rolling Interval value on the Authorization Server Setting. | [optional] 
**RefreshTokenRollingInterval** | Pointer to **int64** | The minimum interval to roll refresh tokens. This value will override the Refresh Token Rolling Interval Value on the Authorization Server Settings. | [optional] 
**RefreshTokenRollingIntervalTimeUnit** | Pointer to **string** | The refresh token rolling interval time unit. Defaults to HOURS. | [optional] 
**RefreshTokenRollingGracePeriodType** | Pointer to **string** | When specified, it overrides the global Refresh Token Grace Period defined in the Authorization Server Settings. | [optional] 
**RefreshTokenRollingGracePeriod** | Pointer to **int64** | The grace period that a rolled refresh token remains valid in seconds. | [optional] 
**RequireOfflineAccessScopeToIssueRefreshTokens** | Pointer to **string** | Determines whether offline_access scope is required to issue refresh tokens or not. | [optional] 
**OfflineAccessRequireConsentPrompt** | Pointer to **string** | Determines whether offline_access requires the prompt parameter value to be set to &#39;consent&#39; or not. | [optional] 
**DefaultAccessTokenManagerRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**RestrictToDefaultAccessTokenManager** | Pointer to **bool** | Determines whether the client is restricted to using only the default access token manager. | [optional] 
**DeviceFlowSettingType** | Pointer to **string** | Allows an administrator to override the device authorization flow setting. SERVER_DEFAULT will use the server default. | [optional] 
**UserAuthorizationUrlOverride** | Pointer to **string** | The URL to redirect the user to for device authorization. Overrides the server default. | [optional] 
**PendingAuthorizationTimeoutOverride** | Pointer to **int64** | The pending authorization timeout override in seconds. | [optional] 
**DevicePollingIntervalOverride** | Pointer to **int64** | The device authorization polling interval override in seconds. | [optional] 
**BypassActivationCodeConfirmationOverride** | Pointer to **bool** | Bypass activation code confirmation for device authorization. | [optional] 
**LockoutMaxMaliciousActionsType** | Pointer to **string** | Allows an administrator to override the lockout maximum malicious actions setting. SERVER_DEFAULT will use the server default. | [optional] 
**LockoutMaxMaliciousActions** | Pointer to **int64** | The maximum number of malicious actions before lockout. | [optional] 
**ClientCertIssuerType** | Pointer to **string** | The type of client certificate issuer validation. Valid values: NONE, TRUST_ANY, CERTIFICATE. | [optional] 
**ClientCertIssuerRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**DpopProofSettings** | Pointer to [**DpopProofSettings**](DpopProofSettings.md) |  | [optional] 
**OidcPolicy** | Pointer to [**ClientRegistrationOIDCPolicy**](ClientRegistrationOIDCPolicy.md) |  | [optional] 
**CibaRequireSignedRequests** | Pointer to **bool** | Whether CIBA requests must be signed. Overrides the server default when set. | [optional] 
**CibaPollingInterval** | Pointer to **int64** | The CIBA polling interval in seconds. Overrides the server default when set. | [optional] 
**CibaRequestPolicyRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**TokenExchangeProcessorPolicyRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**MetadataUrls** | **[]string** | A list of URL patterns that match client_id values subject to this policy. Only one item is allowed currently. Example URL patterns: \&quot;https://_*.example.com/cimd\&quot; or \&quot;https://www.example.com/_*\&quot; | 
**CreatedAt** | Pointer to **time.Time** | The time at which this policy was created. Server-set on POST; read-only. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time at which this policy was last updated. Server-set on POST/PUT; read-only. | [optional] [readonly] 
**Tags** | Pointer to [**[]ResourceLink**](ResourceLink.md) | The tags assigned to clients created by this policy. | [optional] 
**EnforceReplayPrevention** | Pointer to **bool** | When true, configures client to enforce replay prevention for Private Key JWT client authentication. | [optional] 
**RequireSignedRequests** | Pointer to **bool** | When true, configures clients to require signed requests (JAR/PAR). | [optional] 
**RequireJwtSecuredAuthorizationResponseMode** | Pointer to **bool** | When true, configures clients to require JWT Secured Authorization Response Mode (JARM). | [optional] 
**RequireProofKeyForCodeExchange** | Pointer to **bool** | When true, configures clients to require Proof Key for Code Exchange (PKCE). | [optional] 

## Methods

### NewCimdPolicy

`func NewCimdPolicy(name string, enabled bool, metadataUrls []string, ) *CimdPolicy`

NewCimdPolicy instantiates a new CimdPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCimdPolicyWithDefaults

`func NewCimdPolicyWithDefaults() *CimdPolicy`

NewCimdPolicyWithDefaults instantiates a new CimdPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CimdPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CimdPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CimdPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CimdPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CimdPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CimdPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CimdPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CimdPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CimdPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CimdPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CimdPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CimdPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CimdPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CimdPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowLoopbackRedirectUris

`func (o *CimdPolicy) GetAllowLoopbackRedirectUris() bool`

GetAllowLoopbackRedirectUris returns the AllowLoopbackRedirectUris field if non-nil, zero value otherwise.

### GetAllowLoopbackRedirectUrisOk

`func (o *CimdPolicy) GetAllowLoopbackRedirectUrisOk() (*bool, bool)`

GetAllowLoopbackRedirectUrisOk returns a tuple with the AllowLoopbackRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLoopbackRedirectUris

`func (o *CimdPolicy) SetAllowLoopbackRedirectUris(v bool)`

SetAllowLoopbackRedirectUris sets AllowLoopbackRedirectUris field to given value.

### HasAllowLoopbackRedirectUris

`func (o *CimdPolicy) HasAllowLoopbackRedirectUris() bool`

HasAllowLoopbackRedirectUris returns a boolean if a field has been set.

### GetPersistentGrantExpirationType

`func (o *CimdPolicy) GetPersistentGrantExpirationType() string`

GetPersistentGrantExpirationType returns the PersistentGrantExpirationType field if non-nil, zero value otherwise.

### GetPersistentGrantExpirationTypeOk

`func (o *CimdPolicy) GetPersistentGrantExpirationTypeOk() (*string, bool)`

GetPersistentGrantExpirationTypeOk returns a tuple with the PersistentGrantExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantExpirationType

`func (o *CimdPolicy) SetPersistentGrantExpirationType(v string)`

SetPersistentGrantExpirationType sets PersistentGrantExpirationType field to given value.

### HasPersistentGrantExpirationType

`func (o *CimdPolicy) HasPersistentGrantExpirationType() bool`

HasPersistentGrantExpirationType returns a boolean if a field has been set.

### GetPersistentGrantExpirationTime

`func (o *CimdPolicy) GetPersistentGrantExpirationTime() int64`

GetPersistentGrantExpirationTime returns the PersistentGrantExpirationTime field if non-nil, zero value otherwise.

### GetPersistentGrantExpirationTimeOk

`func (o *CimdPolicy) GetPersistentGrantExpirationTimeOk() (*int64, bool)`

GetPersistentGrantExpirationTimeOk returns a tuple with the PersistentGrantExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantExpirationTime

`func (o *CimdPolicy) SetPersistentGrantExpirationTime(v int64)`

SetPersistentGrantExpirationTime sets PersistentGrantExpirationTime field to given value.

### HasPersistentGrantExpirationTime

`func (o *CimdPolicy) HasPersistentGrantExpirationTime() bool`

HasPersistentGrantExpirationTime returns a boolean if a field has been set.

### GetPersistentGrantExpirationTimeUnit

`func (o *CimdPolicy) GetPersistentGrantExpirationTimeUnit() string`

GetPersistentGrantExpirationTimeUnit returns the PersistentGrantExpirationTimeUnit field if non-nil, zero value otherwise.

### GetPersistentGrantExpirationTimeUnitOk

`func (o *CimdPolicy) GetPersistentGrantExpirationTimeUnitOk() (*string, bool)`

GetPersistentGrantExpirationTimeUnitOk returns a tuple with the PersistentGrantExpirationTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantExpirationTimeUnit

`func (o *CimdPolicy) SetPersistentGrantExpirationTimeUnit(v string)`

SetPersistentGrantExpirationTimeUnit sets PersistentGrantExpirationTimeUnit field to given value.

### HasPersistentGrantExpirationTimeUnit

`func (o *CimdPolicy) HasPersistentGrantExpirationTimeUnit() bool`

HasPersistentGrantExpirationTimeUnit returns a boolean if a field has been set.

### GetPersistentGrantIdleTimeoutType

`func (o *CimdPolicy) GetPersistentGrantIdleTimeoutType() string`

GetPersistentGrantIdleTimeoutType returns the PersistentGrantIdleTimeoutType field if non-nil, zero value otherwise.

### GetPersistentGrantIdleTimeoutTypeOk

`func (o *CimdPolicy) GetPersistentGrantIdleTimeoutTypeOk() (*string, bool)`

GetPersistentGrantIdleTimeoutTypeOk returns a tuple with the PersistentGrantIdleTimeoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantIdleTimeoutType

`func (o *CimdPolicy) SetPersistentGrantIdleTimeoutType(v string)`

SetPersistentGrantIdleTimeoutType sets PersistentGrantIdleTimeoutType field to given value.

### HasPersistentGrantIdleTimeoutType

`func (o *CimdPolicy) HasPersistentGrantIdleTimeoutType() bool`

HasPersistentGrantIdleTimeoutType returns a boolean if a field has been set.

### GetPersistentGrantIdleTimeout

`func (o *CimdPolicy) GetPersistentGrantIdleTimeout() int64`

GetPersistentGrantIdleTimeout returns the PersistentGrantIdleTimeout field if non-nil, zero value otherwise.

### GetPersistentGrantIdleTimeoutOk

`func (o *CimdPolicy) GetPersistentGrantIdleTimeoutOk() (*int64, bool)`

GetPersistentGrantIdleTimeoutOk returns a tuple with the PersistentGrantIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantIdleTimeout

`func (o *CimdPolicy) SetPersistentGrantIdleTimeout(v int64)`

SetPersistentGrantIdleTimeout sets PersistentGrantIdleTimeout field to given value.

### HasPersistentGrantIdleTimeout

`func (o *CimdPolicy) HasPersistentGrantIdleTimeout() bool`

HasPersistentGrantIdleTimeout returns a boolean if a field has been set.

### GetPersistentGrantIdleTimeoutTimeUnit

`func (o *CimdPolicy) GetPersistentGrantIdleTimeoutTimeUnit() string`

GetPersistentGrantIdleTimeoutTimeUnit returns the PersistentGrantIdleTimeoutTimeUnit field if non-nil, zero value otherwise.

### GetPersistentGrantIdleTimeoutTimeUnitOk

`func (o *CimdPolicy) GetPersistentGrantIdleTimeoutTimeUnitOk() (*string, bool)`

GetPersistentGrantIdleTimeoutTimeUnitOk returns a tuple with the PersistentGrantIdleTimeoutTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentGrantIdleTimeoutTimeUnit

`func (o *CimdPolicy) SetPersistentGrantIdleTimeoutTimeUnit(v string)`

SetPersistentGrantIdleTimeoutTimeUnit sets PersistentGrantIdleTimeoutTimeUnit field to given value.

### HasPersistentGrantIdleTimeoutTimeUnit

`func (o *CimdPolicy) HasPersistentGrantIdleTimeoutTimeUnit() bool`

HasPersistentGrantIdleTimeoutTimeUnit returns a boolean if a field has been set.

### GetRefreshRolling

`func (o *CimdPolicy) GetRefreshRolling() string`

GetRefreshRolling returns the RefreshRolling field if non-nil, zero value otherwise.

### GetRefreshRollingOk

`func (o *CimdPolicy) GetRefreshRollingOk() (*string, bool)`

GetRefreshRollingOk returns a tuple with the RefreshRolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshRolling

`func (o *CimdPolicy) SetRefreshRolling(v string)`

SetRefreshRolling sets RefreshRolling field to given value.

### HasRefreshRolling

`func (o *CimdPolicy) HasRefreshRolling() bool`

HasRefreshRolling returns a boolean if a field has been set.

### GetRefreshTokenRollingIntervalType

`func (o *CimdPolicy) GetRefreshTokenRollingIntervalType() string`

GetRefreshTokenRollingIntervalType returns the RefreshTokenRollingIntervalType field if non-nil, zero value otherwise.

### GetRefreshTokenRollingIntervalTypeOk

`func (o *CimdPolicy) GetRefreshTokenRollingIntervalTypeOk() (*string, bool)`

GetRefreshTokenRollingIntervalTypeOk returns a tuple with the RefreshTokenRollingIntervalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingIntervalType

`func (o *CimdPolicy) SetRefreshTokenRollingIntervalType(v string)`

SetRefreshTokenRollingIntervalType sets RefreshTokenRollingIntervalType field to given value.

### HasRefreshTokenRollingIntervalType

`func (o *CimdPolicy) HasRefreshTokenRollingIntervalType() bool`

HasRefreshTokenRollingIntervalType returns a boolean if a field has been set.

### GetRefreshTokenRollingInterval

`func (o *CimdPolicy) GetRefreshTokenRollingInterval() int64`

GetRefreshTokenRollingInterval returns the RefreshTokenRollingInterval field if non-nil, zero value otherwise.

### GetRefreshTokenRollingIntervalOk

`func (o *CimdPolicy) GetRefreshTokenRollingIntervalOk() (*int64, bool)`

GetRefreshTokenRollingIntervalOk returns a tuple with the RefreshTokenRollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingInterval

`func (o *CimdPolicy) SetRefreshTokenRollingInterval(v int64)`

SetRefreshTokenRollingInterval sets RefreshTokenRollingInterval field to given value.

### HasRefreshTokenRollingInterval

`func (o *CimdPolicy) HasRefreshTokenRollingInterval() bool`

HasRefreshTokenRollingInterval returns a boolean if a field has been set.

### GetRefreshTokenRollingIntervalTimeUnit

`func (o *CimdPolicy) GetRefreshTokenRollingIntervalTimeUnit() string`

GetRefreshTokenRollingIntervalTimeUnit returns the RefreshTokenRollingIntervalTimeUnit field if non-nil, zero value otherwise.

### GetRefreshTokenRollingIntervalTimeUnitOk

`func (o *CimdPolicy) GetRefreshTokenRollingIntervalTimeUnitOk() (*string, bool)`

GetRefreshTokenRollingIntervalTimeUnitOk returns a tuple with the RefreshTokenRollingIntervalTimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingIntervalTimeUnit

`func (o *CimdPolicy) SetRefreshTokenRollingIntervalTimeUnit(v string)`

SetRefreshTokenRollingIntervalTimeUnit sets RefreshTokenRollingIntervalTimeUnit field to given value.

### HasRefreshTokenRollingIntervalTimeUnit

`func (o *CimdPolicy) HasRefreshTokenRollingIntervalTimeUnit() bool`

HasRefreshTokenRollingIntervalTimeUnit returns a boolean if a field has been set.

### GetRefreshTokenRollingGracePeriodType

`func (o *CimdPolicy) GetRefreshTokenRollingGracePeriodType() string`

GetRefreshTokenRollingGracePeriodType returns the RefreshTokenRollingGracePeriodType field if non-nil, zero value otherwise.

### GetRefreshTokenRollingGracePeriodTypeOk

`func (o *CimdPolicy) GetRefreshTokenRollingGracePeriodTypeOk() (*string, bool)`

GetRefreshTokenRollingGracePeriodTypeOk returns a tuple with the RefreshTokenRollingGracePeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingGracePeriodType

`func (o *CimdPolicy) SetRefreshTokenRollingGracePeriodType(v string)`

SetRefreshTokenRollingGracePeriodType sets RefreshTokenRollingGracePeriodType field to given value.

### HasRefreshTokenRollingGracePeriodType

`func (o *CimdPolicy) HasRefreshTokenRollingGracePeriodType() bool`

HasRefreshTokenRollingGracePeriodType returns a boolean if a field has been set.

### GetRefreshTokenRollingGracePeriod

`func (o *CimdPolicy) GetRefreshTokenRollingGracePeriod() int64`

GetRefreshTokenRollingGracePeriod returns the RefreshTokenRollingGracePeriod field if non-nil, zero value otherwise.

### GetRefreshTokenRollingGracePeriodOk

`func (o *CimdPolicy) GetRefreshTokenRollingGracePeriodOk() (*int64, bool)`

GetRefreshTokenRollingGracePeriodOk returns a tuple with the RefreshTokenRollingGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingGracePeriod

`func (o *CimdPolicy) SetRefreshTokenRollingGracePeriod(v int64)`

SetRefreshTokenRollingGracePeriod sets RefreshTokenRollingGracePeriod field to given value.

### HasRefreshTokenRollingGracePeriod

`func (o *CimdPolicy) HasRefreshTokenRollingGracePeriod() bool`

HasRefreshTokenRollingGracePeriod returns a boolean if a field has been set.

### GetRequireOfflineAccessScopeToIssueRefreshTokens

`func (o *CimdPolicy) GetRequireOfflineAccessScopeToIssueRefreshTokens() string`

GetRequireOfflineAccessScopeToIssueRefreshTokens returns the RequireOfflineAccessScopeToIssueRefreshTokens field if non-nil, zero value otherwise.

### GetRequireOfflineAccessScopeToIssueRefreshTokensOk

`func (o *CimdPolicy) GetRequireOfflineAccessScopeToIssueRefreshTokensOk() (*string, bool)`

GetRequireOfflineAccessScopeToIssueRefreshTokensOk returns a tuple with the RequireOfflineAccessScopeToIssueRefreshTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOfflineAccessScopeToIssueRefreshTokens

`func (o *CimdPolicy) SetRequireOfflineAccessScopeToIssueRefreshTokens(v string)`

SetRequireOfflineAccessScopeToIssueRefreshTokens sets RequireOfflineAccessScopeToIssueRefreshTokens field to given value.

### HasRequireOfflineAccessScopeToIssueRefreshTokens

`func (o *CimdPolicy) HasRequireOfflineAccessScopeToIssueRefreshTokens() bool`

HasRequireOfflineAccessScopeToIssueRefreshTokens returns a boolean if a field has been set.

### GetOfflineAccessRequireConsentPrompt

`func (o *CimdPolicy) GetOfflineAccessRequireConsentPrompt() string`

GetOfflineAccessRequireConsentPrompt returns the OfflineAccessRequireConsentPrompt field if non-nil, zero value otherwise.

### GetOfflineAccessRequireConsentPromptOk

`func (o *CimdPolicy) GetOfflineAccessRequireConsentPromptOk() (*string, bool)`

GetOfflineAccessRequireConsentPromptOk returns a tuple with the OfflineAccessRequireConsentPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineAccessRequireConsentPrompt

`func (o *CimdPolicy) SetOfflineAccessRequireConsentPrompt(v string)`

SetOfflineAccessRequireConsentPrompt sets OfflineAccessRequireConsentPrompt field to given value.

### HasOfflineAccessRequireConsentPrompt

`func (o *CimdPolicy) HasOfflineAccessRequireConsentPrompt() bool`

HasOfflineAccessRequireConsentPrompt returns a boolean if a field has been set.

### GetDefaultAccessTokenManagerRef

`func (o *CimdPolicy) GetDefaultAccessTokenManagerRef() ResourceLink`

GetDefaultAccessTokenManagerRef returns the DefaultAccessTokenManagerRef field if non-nil, zero value otherwise.

### GetDefaultAccessTokenManagerRefOk

`func (o *CimdPolicy) GetDefaultAccessTokenManagerRefOk() (*ResourceLink, bool)`

GetDefaultAccessTokenManagerRefOk returns a tuple with the DefaultAccessTokenManagerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccessTokenManagerRef

`func (o *CimdPolicy) SetDefaultAccessTokenManagerRef(v ResourceLink)`

SetDefaultAccessTokenManagerRef sets DefaultAccessTokenManagerRef field to given value.

### HasDefaultAccessTokenManagerRef

`func (o *CimdPolicy) HasDefaultAccessTokenManagerRef() bool`

HasDefaultAccessTokenManagerRef returns a boolean if a field has been set.

### GetRestrictToDefaultAccessTokenManager

`func (o *CimdPolicy) GetRestrictToDefaultAccessTokenManager() bool`

GetRestrictToDefaultAccessTokenManager returns the RestrictToDefaultAccessTokenManager field if non-nil, zero value otherwise.

### GetRestrictToDefaultAccessTokenManagerOk

`func (o *CimdPolicy) GetRestrictToDefaultAccessTokenManagerOk() (*bool, bool)`

GetRestrictToDefaultAccessTokenManagerOk returns a tuple with the RestrictToDefaultAccessTokenManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictToDefaultAccessTokenManager

`func (o *CimdPolicy) SetRestrictToDefaultAccessTokenManager(v bool)`

SetRestrictToDefaultAccessTokenManager sets RestrictToDefaultAccessTokenManager field to given value.

### HasRestrictToDefaultAccessTokenManager

`func (o *CimdPolicy) HasRestrictToDefaultAccessTokenManager() bool`

HasRestrictToDefaultAccessTokenManager returns a boolean if a field has been set.

### GetDeviceFlowSettingType

`func (o *CimdPolicy) GetDeviceFlowSettingType() string`

GetDeviceFlowSettingType returns the DeviceFlowSettingType field if non-nil, zero value otherwise.

### GetDeviceFlowSettingTypeOk

`func (o *CimdPolicy) GetDeviceFlowSettingTypeOk() (*string, bool)`

GetDeviceFlowSettingTypeOk returns a tuple with the DeviceFlowSettingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFlowSettingType

`func (o *CimdPolicy) SetDeviceFlowSettingType(v string)`

SetDeviceFlowSettingType sets DeviceFlowSettingType field to given value.

### HasDeviceFlowSettingType

`func (o *CimdPolicy) HasDeviceFlowSettingType() bool`

HasDeviceFlowSettingType returns a boolean if a field has been set.

### GetUserAuthorizationUrlOverride

`func (o *CimdPolicy) GetUserAuthorizationUrlOverride() string`

GetUserAuthorizationUrlOverride returns the UserAuthorizationUrlOverride field if non-nil, zero value otherwise.

### GetUserAuthorizationUrlOverrideOk

`func (o *CimdPolicy) GetUserAuthorizationUrlOverrideOk() (*string, bool)`

GetUserAuthorizationUrlOverrideOk returns a tuple with the UserAuthorizationUrlOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuthorizationUrlOverride

`func (o *CimdPolicy) SetUserAuthorizationUrlOverride(v string)`

SetUserAuthorizationUrlOverride sets UserAuthorizationUrlOverride field to given value.

### HasUserAuthorizationUrlOverride

`func (o *CimdPolicy) HasUserAuthorizationUrlOverride() bool`

HasUserAuthorizationUrlOverride returns a boolean if a field has been set.

### GetPendingAuthorizationTimeoutOverride

`func (o *CimdPolicy) GetPendingAuthorizationTimeoutOverride() int64`

GetPendingAuthorizationTimeoutOverride returns the PendingAuthorizationTimeoutOverride field if non-nil, zero value otherwise.

### GetPendingAuthorizationTimeoutOverrideOk

`func (o *CimdPolicy) GetPendingAuthorizationTimeoutOverrideOk() (*int64, bool)`

GetPendingAuthorizationTimeoutOverrideOk returns a tuple with the PendingAuthorizationTimeoutOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAuthorizationTimeoutOverride

`func (o *CimdPolicy) SetPendingAuthorizationTimeoutOverride(v int64)`

SetPendingAuthorizationTimeoutOverride sets PendingAuthorizationTimeoutOverride field to given value.

### HasPendingAuthorizationTimeoutOverride

`func (o *CimdPolicy) HasPendingAuthorizationTimeoutOverride() bool`

HasPendingAuthorizationTimeoutOverride returns a boolean if a field has been set.

### GetDevicePollingIntervalOverride

`func (o *CimdPolicy) GetDevicePollingIntervalOverride() int64`

GetDevicePollingIntervalOverride returns the DevicePollingIntervalOverride field if non-nil, zero value otherwise.

### GetDevicePollingIntervalOverrideOk

`func (o *CimdPolicy) GetDevicePollingIntervalOverrideOk() (*int64, bool)`

GetDevicePollingIntervalOverrideOk returns a tuple with the DevicePollingIntervalOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePollingIntervalOverride

`func (o *CimdPolicy) SetDevicePollingIntervalOverride(v int64)`

SetDevicePollingIntervalOverride sets DevicePollingIntervalOverride field to given value.

### HasDevicePollingIntervalOverride

`func (o *CimdPolicy) HasDevicePollingIntervalOverride() bool`

HasDevicePollingIntervalOverride returns a boolean if a field has been set.

### GetBypassActivationCodeConfirmationOverride

`func (o *CimdPolicy) GetBypassActivationCodeConfirmationOverride() bool`

GetBypassActivationCodeConfirmationOverride returns the BypassActivationCodeConfirmationOverride field if non-nil, zero value otherwise.

### GetBypassActivationCodeConfirmationOverrideOk

`func (o *CimdPolicy) GetBypassActivationCodeConfirmationOverrideOk() (*bool, bool)`

GetBypassActivationCodeConfirmationOverrideOk returns a tuple with the BypassActivationCodeConfirmationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassActivationCodeConfirmationOverride

`func (o *CimdPolicy) SetBypassActivationCodeConfirmationOverride(v bool)`

SetBypassActivationCodeConfirmationOverride sets BypassActivationCodeConfirmationOverride field to given value.

### HasBypassActivationCodeConfirmationOverride

`func (o *CimdPolicy) HasBypassActivationCodeConfirmationOverride() bool`

HasBypassActivationCodeConfirmationOverride returns a boolean if a field has been set.

### GetLockoutMaxMaliciousActionsType

`func (o *CimdPolicy) GetLockoutMaxMaliciousActionsType() string`

GetLockoutMaxMaliciousActionsType returns the LockoutMaxMaliciousActionsType field if non-nil, zero value otherwise.

### GetLockoutMaxMaliciousActionsTypeOk

`func (o *CimdPolicy) GetLockoutMaxMaliciousActionsTypeOk() (*string, bool)`

GetLockoutMaxMaliciousActionsTypeOk returns a tuple with the LockoutMaxMaliciousActionsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutMaxMaliciousActionsType

`func (o *CimdPolicy) SetLockoutMaxMaliciousActionsType(v string)`

SetLockoutMaxMaliciousActionsType sets LockoutMaxMaliciousActionsType field to given value.

### HasLockoutMaxMaliciousActionsType

`func (o *CimdPolicy) HasLockoutMaxMaliciousActionsType() bool`

HasLockoutMaxMaliciousActionsType returns a boolean if a field has been set.

### GetLockoutMaxMaliciousActions

`func (o *CimdPolicy) GetLockoutMaxMaliciousActions() int64`

GetLockoutMaxMaliciousActions returns the LockoutMaxMaliciousActions field if non-nil, zero value otherwise.

### GetLockoutMaxMaliciousActionsOk

`func (o *CimdPolicy) GetLockoutMaxMaliciousActionsOk() (*int64, bool)`

GetLockoutMaxMaliciousActionsOk returns a tuple with the LockoutMaxMaliciousActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutMaxMaliciousActions

`func (o *CimdPolicy) SetLockoutMaxMaliciousActions(v int64)`

SetLockoutMaxMaliciousActions sets LockoutMaxMaliciousActions field to given value.

### HasLockoutMaxMaliciousActions

`func (o *CimdPolicy) HasLockoutMaxMaliciousActions() bool`

HasLockoutMaxMaliciousActions returns a boolean if a field has been set.

### GetClientCertIssuerType

`func (o *CimdPolicy) GetClientCertIssuerType() string`

GetClientCertIssuerType returns the ClientCertIssuerType field if non-nil, zero value otherwise.

### GetClientCertIssuerTypeOk

`func (o *CimdPolicy) GetClientCertIssuerTypeOk() (*string, bool)`

GetClientCertIssuerTypeOk returns a tuple with the ClientCertIssuerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertIssuerType

`func (o *CimdPolicy) SetClientCertIssuerType(v string)`

SetClientCertIssuerType sets ClientCertIssuerType field to given value.

### HasClientCertIssuerType

`func (o *CimdPolicy) HasClientCertIssuerType() bool`

HasClientCertIssuerType returns a boolean if a field has been set.

### GetClientCertIssuerRef

`func (o *CimdPolicy) GetClientCertIssuerRef() ResourceLink`

GetClientCertIssuerRef returns the ClientCertIssuerRef field if non-nil, zero value otherwise.

### GetClientCertIssuerRefOk

`func (o *CimdPolicy) GetClientCertIssuerRefOk() (*ResourceLink, bool)`

GetClientCertIssuerRefOk returns a tuple with the ClientCertIssuerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertIssuerRef

`func (o *CimdPolicy) SetClientCertIssuerRef(v ResourceLink)`

SetClientCertIssuerRef sets ClientCertIssuerRef field to given value.

### HasClientCertIssuerRef

`func (o *CimdPolicy) HasClientCertIssuerRef() bool`

HasClientCertIssuerRef returns a boolean if a field has been set.

### GetDpopProofSettings

`func (o *CimdPolicy) GetDpopProofSettings() DpopProofSettings`

GetDpopProofSettings returns the DpopProofSettings field if non-nil, zero value otherwise.

### GetDpopProofSettingsOk

`func (o *CimdPolicy) GetDpopProofSettingsOk() (*DpopProofSettings, bool)`

GetDpopProofSettingsOk returns a tuple with the DpopProofSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopProofSettings

`func (o *CimdPolicy) SetDpopProofSettings(v DpopProofSettings)`

SetDpopProofSettings sets DpopProofSettings field to given value.

### HasDpopProofSettings

`func (o *CimdPolicy) HasDpopProofSettings() bool`

HasDpopProofSettings returns a boolean if a field has been set.

### GetOidcPolicy

`func (o *CimdPolicy) GetOidcPolicy() ClientRegistrationOIDCPolicy`

GetOidcPolicy returns the OidcPolicy field if non-nil, zero value otherwise.

### GetOidcPolicyOk

`func (o *CimdPolicy) GetOidcPolicyOk() (*ClientRegistrationOIDCPolicy, bool)`

GetOidcPolicyOk returns a tuple with the OidcPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcPolicy

`func (o *CimdPolicy) SetOidcPolicy(v ClientRegistrationOIDCPolicy)`

SetOidcPolicy sets OidcPolicy field to given value.

### HasOidcPolicy

`func (o *CimdPolicy) HasOidcPolicy() bool`

HasOidcPolicy returns a boolean if a field has been set.

### GetCibaRequireSignedRequests

`func (o *CimdPolicy) GetCibaRequireSignedRequests() bool`

GetCibaRequireSignedRequests returns the CibaRequireSignedRequests field if non-nil, zero value otherwise.

### GetCibaRequireSignedRequestsOk

`func (o *CimdPolicy) GetCibaRequireSignedRequestsOk() (*bool, bool)`

GetCibaRequireSignedRequestsOk returns a tuple with the CibaRequireSignedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCibaRequireSignedRequests

`func (o *CimdPolicy) SetCibaRequireSignedRequests(v bool)`

SetCibaRequireSignedRequests sets CibaRequireSignedRequests field to given value.

### HasCibaRequireSignedRequests

`func (o *CimdPolicy) HasCibaRequireSignedRequests() bool`

HasCibaRequireSignedRequests returns a boolean if a field has been set.

### GetCibaPollingInterval

`func (o *CimdPolicy) GetCibaPollingInterval() int64`

GetCibaPollingInterval returns the CibaPollingInterval field if non-nil, zero value otherwise.

### GetCibaPollingIntervalOk

`func (o *CimdPolicy) GetCibaPollingIntervalOk() (*int64, bool)`

GetCibaPollingIntervalOk returns a tuple with the CibaPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCibaPollingInterval

`func (o *CimdPolicy) SetCibaPollingInterval(v int64)`

SetCibaPollingInterval sets CibaPollingInterval field to given value.

### HasCibaPollingInterval

`func (o *CimdPolicy) HasCibaPollingInterval() bool`

HasCibaPollingInterval returns a boolean if a field has been set.

### GetCibaRequestPolicyRef

`func (o *CimdPolicy) GetCibaRequestPolicyRef() ResourceLink`

GetCibaRequestPolicyRef returns the CibaRequestPolicyRef field if non-nil, zero value otherwise.

### GetCibaRequestPolicyRefOk

`func (o *CimdPolicy) GetCibaRequestPolicyRefOk() (*ResourceLink, bool)`

GetCibaRequestPolicyRefOk returns a tuple with the CibaRequestPolicyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCibaRequestPolicyRef

`func (o *CimdPolicy) SetCibaRequestPolicyRef(v ResourceLink)`

SetCibaRequestPolicyRef sets CibaRequestPolicyRef field to given value.

### HasCibaRequestPolicyRef

`func (o *CimdPolicy) HasCibaRequestPolicyRef() bool`

HasCibaRequestPolicyRef returns a boolean if a field has been set.

### GetTokenExchangeProcessorPolicyRef

`func (o *CimdPolicy) GetTokenExchangeProcessorPolicyRef() ResourceLink`

GetTokenExchangeProcessorPolicyRef returns the TokenExchangeProcessorPolicyRef field if non-nil, zero value otherwise.

### GetTokenExchangeProcessorPolicyRefOk

`func (o *CimdPolicy) GetTokenExchangeProcessorPolicyRefOk() (*ResourceLink, bool)`

GetTokenExchangeProcessorPolicyRefOk returns a tuple with the TokenExchangeProcessorPolicyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExchangeProcessorPolicyRef

`func (o *CimdPolicy) SetTokenExchangeProcessorPolicyRef(v ResourceLink)`

SetTokenExchangeProcessorPolicyRef sets TokenExchangeProcessorPolicyRef field to given value.

### HasTokenExchangeProcessorPolicyRef

`func (o *CimdPolicy) HasTokenExchangeProcessorPolicyRef() bool`

HasTokenExchangeProcessorPolicyRef returns a boolean if a field has been set.

### GetMetadataUrls

`func (o *CimdPolicy) GetMetadataUrls() []string`

GetMetadataUrls returns the MetadataUrls field if non-nil, zero value otherwise.

### GetMetadataUrlsOk

`func (o *CimdPolicy) GetMetadataUrlsOk() (*[]string, bool)`

GetMetadataUrlsOk returns a tuple with the MetadataUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUrls

`func (o *CimdPolicy) SetMetadataUrls(v []string)`

SetMetadataUrls sets MetadataUrls field to given value.


### GetCreatedAt

`func (o *CimdPolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CimdPolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CimdPolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CimdPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CimdPolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CimdPolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CimdPolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CimdPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTags

`func (o *CimdPolicy) GetTags() []ResourceLink`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CimdPolicy) GetTagsOk() (*[]ResourceLink, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CimdPolicy) SetTags(v []ResourceLink)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CimdPolicy) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnforceReplayPrevention

`func (o *CimdPolicy) GetEnforceReplayPrevention() bool`

GetEnforceReplayPrevention returns the EnforceReplayPrevention field if non-nil, zero value otherwise.

### GetEnforceReplayPreventionOk

`func (o *CimdPolicy) GetEnforceReplayPreventionOk() (*bool, bool)`

GetEnforceReplayPreventionOk returns a tuple with the EnforceReplayPrevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceReplayPrevention

`func (o *CimdPolicy) SetEnforceReplayPrevention(v bool)`

SetEnforceReplayPrevention sets EnforceReplayPrevention field to given value.

### HasEnforceReplayPrevention

`func (o *CimdPolicy) HasEnforceReplayPrevention() bool`

HasEnforceReplayPrevention returns a boolean if a field has been set.

### GetRequireSignedRequests

`func (o *CimdPolicy) GetRequireSignedRequests() bool`

GetRequireSignedRequests returns the RequireSignedRequests field if non-nil, zero value otherwise.

### GetRequireSignedRequestsOk

`func (o *CimdPolicy) GetRequireSignedRequestsOk() (*bool, bool)`

GetRequireSignedRequestsOk returns a tuple with the RequireSignedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedRequests

`func (o *CimdPolicy) SetRequireSignedRequests(v bool)`

SetRequireSignedRequests sets RequireSignedRequests field to given value.

### HasRequireSignedRequests

`func (o *CimdPolicy) HasRequireSignedRequests() bool`

HasRequireSignedRequests returns a boolean if a field has been set.

### GetRequireJwtSecuredAuthorizationResponseMode

`func (o *CimdPolicy) GetRequireJwtSecuredAuthorizationResponseMode() bool`

GetRequireJwtSecuredAuthorizationResponseMode returns the RequireJwtSecuredAuthorizationResponseMode field if non-nil, zero value otherwise.

### GetRequireJwtSecuredAuthorizationResponseModeOk

`func (o *CimdPolicy) GetRequireJwtSecuredAuthorizationResponseModeOk() (*bool, bool)`

GetRequireJwtSecuredAuthorizationResponseModeOk returns a tuple with the RequireJwtSecuredAuthorizationResponseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireJwtSecuredAuthorizationResponseMode

`func (o *CimdPolicy) SetRequireJwtSecuredAuthorizationResponseMode(v bool)`

SetRequireJwtSecuredAuthorizationResponseMode sets RequireJwtSecuredAuthorizationResponseMode field to given value.

### HasRequireJwtSecuredAuthorizationResponseMode

`func (o *CimdPolicy) HasRequireJwtSecuredAuthorizationResponseMode() bool`

HasRequireJwtSecuredAuthorizationResponseMode returns a boolean if a field has been set.

### GetRequireProofKeyForCodeExchange

`func (o *CimdPolicy) GetRequireProofKeyForCodeExchange() bool`

GetRequireProofKeyForCodeExchange returns the RequireProofKeyForCodeExchange field if non-nil, zero value otherwise.

### GetRequireProofKeyForCodeExchangeOk

`func (o *CimdPolicy) GetRequireProofKeyForCodeExchangeOk() (*bool, bool)`

GetRequireProofKeyForCodeExchangeOk returns a tuple with the RequireProofKeyForCodeExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireProofKeyForCodeExchange

`func (o *CimdPolicy) SetRequireProofKeyForCodeExchange(v bool)`

SetRequireProofKeyForCodeExchange sets RequireProofKeyForCodeExchange field to given value.

### HasRequireProofKeyForCodeExchange

`func (o *CimdPolicy) HasRequireProofKeyForCodeExchange() bool`

HasRequireProofKeyForCodeExchange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


