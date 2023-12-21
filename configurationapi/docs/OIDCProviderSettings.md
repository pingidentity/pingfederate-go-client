# OIDCProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | **string** | Space separated scope values that the OpenID Provider supports. | 
**AuthorizationEndpoint** | **string** | URL of the OpenID Provider&#39;s OAuth 2.0 Authorization Endpoint. | 
**PushedAuthorizationRequestEndpoint** | Pointer to **string** | URL of the OpenID Provider&#39;s OAuth 2.0 Pushed Authorization Request Endpoint. | [optional] 
**LoginType** | **string** | The OpenID Connect login type. These values maps to: &lt;br&gt;  CODE: Authentication using Code Flow &lt;br&gt; POST: Authentication using Form Post &lt;br&gt; POST_AT: Authentication using Form Post with Access Token | 
**AuthenticationScheme** | Pointer to **string** | The OpenID Connect Authentication Scheme. This is required for Authentication using Code Flow.  | [optional] 
**AuthenticationSigningAlgorithm** | Pointer to **string** | The authentication signing algorithm for token endpoint PRIVATE_KEY_JWT or CLIENT_SECRET_JWT authentication. Asymmetric algorithms are allowed for PRIVATE_KEY_JWT and symmetric algorithms are allowed for CLIENT_SECRET_JWT. For RSASSA-PSS signing algorithm, PingFederate must be integrated with a hardware security module (HSM) or Java 11. | [optional] 
**RequestSigningAlgorithm** | Pointer to **string** | The request signing algorithm. Required only if you wish to use signed requests. Only asymmetric algorithms are allowed. For RSASSA-PSS signing algorithm, PingFederate must be integrated with a hardware security module (HSM) or Java 11. | [optional] 
**EnablePKCE** | Pointer to **bool** | Enable Proof Key for Code Exchange (PKCE). When enabled, the client sends an SHA-256 code challenge and corresponding code verifier to the OpenID Provider during the authorization code flow. | [optional] 
**TokenEndpoint** | Pointer to **string** | URL of the OpenID Provider&#39;s OAuth 2.0 Token Endpoint. | [optional] 
**UserInfoEndpoint** | Pointer to **string** | URL of the OpenID Provider&#39;s UserInfo Endpoint. | [optional] 
**LogoutEndpoint** | Pointer to **string** | URL of the OpenID Provider&#39;s RP-Initiated Logout Endpoint. | [optional] 
**JwksURL** | **string** | URL of the OpenID Provider&#39;s JSON Web Key Set [JWK] document. | 
**TrackUserSessionsForLogout** | Pointer to **bool** | Determines whether PingFederate tracks a logout entry when a user signs in, so that the user session can later be terminated via a logout request from the OP. This setting must also be enabled in order for PingFederate to send an RP-initiated logout request to the OP during SLO. | [optional] 
**RequestParameters** | Pointer to [**[]OIDCRequestParameter**](OIDCRequestParameter.md) | A list of request parameters. Request parameters with same name but different attribute values are treated as a multi-valued request parameter. | [optional] 
**RedirectUri** | Pointer to **string** | The redirect URI. This is a read-only parameter. | [optional] 
**BackChannelLogoutUri** | Pointer to **string** | The Back-Channel Logout URI. This read-only parameter is available when user sessions are tracked for logout. | [optional] 
**FrontChannelLogoutUri** | Pointer to **string** | The Front-Channel Logout URI. This is a read-only parameter. | [optional] 
**PostLogoutRedirectUri** | Pointer to **string** | The Post-Logout Redirect URI, where the OpenID Provider may redirect the user when RP-Initiated Logout has completed. This is a read-only parameter. | [optional] 

## Methods

### NewOIDCProviderSettings

`func NewOIDCProviderSettings(scopes string, authorizationEndpoint string, loginType string, jwksURL string, ) *OIDCProviderSettings`

NewOIDCProviderSettings instantiates a new OIDCProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCProviderSettingsWithDefaults

`func NewOIDCProviderSettingsWithDefaults() *OIDCProviderSettings`

NewOIDCProviderSettingsWithDefaults instantiates a new OIDCProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *OIDCProviderSettings) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OIDCProviderSettings) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OIDCProviderSettings) SetScopes(v string)`

SetScopes sets Scopes field to given value.


### GetAuthorizationEndpoint

`func (o *OIDCProviderSettings) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *OIDCProviderSettings) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *OIDCProviderSettings) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.


### GetPushedAuthorizationRequestEndpoint

`func (o *OIDCProviderSettings) GetPushedAuthorizationRequestEndpoint() string`

GetPushedAuthorizationRequestEndpoint returns the PushedAuthorizationRequestEndpoint field if non-nil, zero value otherwise.

### GetPushedAuthorizationRequestEndpointOk

`func (o *OIDCProviderSettings) GetPushedAuthorizationRequestEndpointOk() (*string, bool)`

GetPushedAuthorizationRequestEndpointOk returns a tuple with the PushedAuthorizationRequestEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAuthorizationRequestEndpoint

`func (o *OIDCProviderSettings) SetPushedAuthorizationRequestEndpoint(v string)`

SetPushedAuthorizationRequestEndpoint sets PushedAuthorizationRequestEndpoint field to given value.

### HasPushedAuthorizationRequestEndpoint

`func (o *OIDCProviderSettings) HasPushedAuthorizationRequestEndpoint() bool`

HasPushedAuthorizationRequestEndpoint returns a boolean if a field has been set.

### GetLoginType

`func (o *OIDCProviderSettings) GetLoginType() string`

GetLoginType returns the LoginType field if non-nil, zero value otherwise.

### GetLoginTypeOk

`func (o *OIDCProviderSettings) GetLoginTypeOk() (*string, bool)`

GetLoginTypeOk returns a tuple with the LoginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginType

`func (o *OIDCProviderSettings) SetLoginType(v string)`

SetLoginType sets LoginType field to given value.


### GetAuthenticationScheme

`func (o *OIDCProviderSettings) GetAuthenticationScheme() string`

GetAuthenticationScheme returns the AuthenticationScheme field if non-nil, zero value otherwise.

### GetAuthenticationSchemeOk

`func (o *OIDCProviderSettings) GetAuthenticationSchemeOk() (*string, bool)`

GetAuthenticationSchemeOk returns a tuple with the AuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationScheme

`func (o *OIDCProviderSettings) SetAuthenticationScheme(v string)`

SetAuthenticationScheme sets AuthenticationScheme field to given value.

### HasAuthenticationScheme

`func (o *OIDCProviderSettings) HasAuthenticationScheme() bool`

HasAuthenticationScheme returns a boolean if a field has been set.

### GetAuthenticationSigningAlgorithm

`func (o *OIDCProviderSettings) GetAuthenticationSigningAlgorithm() string`

GetAuthenticationSigningAlgorithm returns the AuthenticationSigningAlgorithm field if non-nil, zero value otherwise.

### GetAuthenticationSigningAlgorithmOk

`func (o *OIDCProviderSettings) GetAuthenticationSigningAlgorithmOk() (*string, bool)`

GetAuthenticationSigningAlgorithmOk returns a tuple with the AuthenticationSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSigningAlgorithm

`func (o *OIDCProviderSettings) SetAuthenticationSigningAlgorithm(v string)`

SetAuthenticationSigningAlgorithm sets AuthenticationSigningAlgorithm field to given value.

### HasAuthenticationSigningAlgorithm

`func (o *OIDCProviderSettings) HasAuthenticationSigningAlgorithm() bool`

HasAuthenticationSigningAlgorithm returns a boolean if a field has been set.

### GetRequestSigningAlgorithm

`func (o *OIDCProviderSettings) GetRequestSigningAlgorithm() string`

GetRequestSigningAlgorithm returns the RequestSigningAlgorithm field if non-nil, zero value otherwise.

### GetRequestSigningAlgorithmOk

`func (o *OIDCProviderSettings) GetRequestSigningAlgorithmOk() (*string, bool)`

GetRequestSigningAlgorithmOk returns a tuple with the RequestSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSigningAlgorithm

`func (o *OIDCProviderSettings) SetRequestSigningAlgorithm(v string)`

SetRequestSigningAlgorithm sets RequestSigningAlgorithm field to given value.

### HasRequestSigningAlgorithm

`func (o *OIDCProviderSettings) HasRequestSigningAlgorithm() bool`

HasRequestSigningAlgorithm returns a boolean if a field has been set.

### GetEnablePKCE

`func (o *OIDCProviderSettings) GetEnablePKCE() bool`

GetEnablePKCE returns the EnablePKCE field if non-nil, zero value otherwise.

### GetEnablePKCEOk

`func (o *OIDCProviderSettings) GetEnablePKCEOk() (*bool, bool)`

GetEnablePKCEOk returns a tuple with the EnablePKCE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePKCE

`func (o *OIDCProviderSettings) SetEnablePKCE(v bool)`

SetEnablePKCE sets EnablePKCE field to given value.

### HasEnablePKCE

`func (o *OIDCProviderSettings) HasEnablePKCE() bool`

HasEnablePKCE returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *OIDCProviderSettings) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *OIDCProviderSettings) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *OIDCProviderSettings) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *OIDCProviderSettings) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetUserInfoEndpoint

`func (o *OIDCProviderSettings) GetUserInfoEndpoint() string`

GetUserInfoEndpoint returns the UserInfoEndpoint field if non-nil, zero value otherwise.

### GetUserInfoEndpointOk

`func (o *OIDCProviderSettings) GetUserInfoEndpointOk() (*string, bool)`

GetUserInfoEndpointOk returns a tuple with the UserInfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoEndpoint

`func (o *OIDCProviderSettings) SetUserInfoEndpoint(v string)`

SetUserInfoEndpoint sets UserInfoEndpoint field to given value.

### HasUserInfoEndpoint

`func (o *OIDCProviderSettings) HasUserInfoEndpoint() bool`

HasUserInfoEndpoint returns a boolean if a field has been set.

### GetLogoutEndpoint

`func (o *OIDCProviderSettings) GetLogoutEndpoint() string`

GetLogoutEndpoint returns the LogoutEndpoint field if non-nil, zero value otherwise.

### GetLogoutEndpointOk

`func (o *OIDCProviderSettings) GetLogoutEndpointOk() (*string, bool)`

GetLogoutEndpointOk returns a tuple with the LogoutEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutEndpoint

`func (o *OIDCProviderSettings) SetLogoutEndpoint(v string)`

SetLogoutEndpoint sets LogoutEndpoint field to given value.

### HasLogoutEndpoint

`func (o *OIDCProviderSettings) HasLogoutEndpoint() bool`

HasLogoutEndpoint returns a boolean if a field has been set.

### GetJwksURL

`func (o *OIDCProviderSettings) GetJwksURL() string`

GetJwksURL returns the JwksURL field if non-nil, zero value otherwise.

### GetJwksURLOk

`func (o *OIDCProviderSettings) GetJwksURLOk() (*string, bool)`

GetJwksURLOk returns a tuple with the JwksURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksURL

`func (o *OIDCProviderSettings) SetJwksURL(v string)`

SetJwksURL sets JwksURL field to given value.


### GetTrackUserSessionsForLogout

`func (o *OIDCProviderSettings) GetTrackUserSessionsForLogout() bool`

GetTrackUserSessionsForLogout returns the TrackUserSessionsForLogout field if non-nil, zero value otherwise.

### GetTrackUserSessionsForLogoutOk

`func (o *OIDCProviderSettings) GetTrackUserSessionsForLogoutOk() (*bool, bool)`

GetTrackUserSessionsForLogoutOk returns a tuple with the TrackUserSessionsForLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackUserSessionsForLogout

`func (o *OIDCProviderSettings) SetTrackUserSessionsForLogout(v bool)`

SetTrackUserSessionsForLogout sets TrackUserSessionsForLogout field to given value.

### HasTrackUserSessionsForLogout

`func (o *OIDCProviderSettings) HasTrackUserSessionsForLogout() bool`

HasTrackUserSessionsForLogout returns a boolean if a field has been set.

### GetRequestParameters

`func (o *OIDCProviderSettings) GetRequestParameters() []OIDCRequestParameter`

GetRequestParameters returns the RequestParameters field if non-nil, zero value otherwise.

### GetRequestParametersOk

`func (o *OIDCProviderSettings) GetRequestParametersOk() (*[]OIDCRequestParameter, bool)`

GetRequestParametersOk returns a tuple with the RequestParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParameters

`func (o *OIDCProviderSettings) SetRequestParameters(v []OIDCRequestParameter)`

SetRequestParameters sets RequestParameters field to given value.

### HasRequestParameters

`func (o *OIDCProviderSettings) HasRequestParameters() bool`

HasRequestParameters returns a boolean if a field has been set.

### GetRedirectUri

`func (o *OIDCProviderSettings) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OIDCProviderSettings) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OIDCProviderSettings) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *OIDCProviderSettings) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetBackChannelLogoutUri

`func (o *OIDCProviderSettings) GetBackChannelLogoutUri() string`

GetBackChannelLogoutUri returns the BackChannelLogoutUri field if non-nil, zero value otherwise.

### GetBackChannelLogoutUriOk

`func (o *OIDCProviderSettings) GetBackChannelLogoutUriOk() (*string, bool)`

GetBackChannelLogoutUriOk returns a tuple with the BackChannelLogoutUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackChannelLogoutUri

`func (o *OIDCProviderSettings) SetBackChannelLogoutUri(v string)`

SetBackChannelLogoutUri sets BackChannelLogoutUri field to given value.

### HasBackChannelLogoutUri

`func (o *OIDCProviderSettings) HasBackChannelLogoutUri() bool`

HasBackChannelLogoutUri returns a boolean if a field has been set.

### GetFrontChannelLogoutUri

`func (o *OIDCProviderSettings) GetFrontChannelLogoutUri() string`

GetFrontChannelLogoutUri returns the FrontChannelLogoutUri field if non-nil, zero value otherwise.

### GetFrontChannelLogoutUriOk

`func (o *OIDCProviderSettings) GetFrontChannelLogoutUriOk() (*string, bool)`

GetFrontChannelLogoutUriOk returns a tuple with the FrontChannelLogoutUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontChannelLogoutUri

`func (o *OIDCProviderSettings) SetFrontChannelLogoutUri(v string)`

SetFrontChannelLogoutUri sets FrontChannelLogoutUri field to given value.

### HasFrontChannelLogoutUri

`func (o *OIDCProviderSettings) HasFrontChannelLogoutUri() bool`

HasFrontChannelLogoutUri returns a boolean if a field has been set.

### GetPostLogoutRedirectUri

`func (o *OIDCProviderSettings) GetPostLogoutRedirectUri() string`

GetPostLogoutRedirectUri returns the PostLogoutRedirectUri field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUriOk

`func (o *OIDCProviderSettings) GetPostLogoutRedirectUriOk() (*string, bool)`

GetPostLogoutRedirectUriOk returns a tuple with the PostLogoutRedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUri

`func (o *OIDCProviderSettings) SetPostLogoutRedirectUri(v string)`

SetPostLogoutRedirectUri sets PostLogoutRedirectUri field to given value.

### HasPostLogoutRedirectUri

`func (o *OIDCProviderSettings) HasPostLogoutRedirectUri() bool`

HasPostLogoutRedirectUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


