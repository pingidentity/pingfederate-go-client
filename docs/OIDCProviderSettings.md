# OIDCProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | **string** | Space separated scope values that the OpenID Provider supports. | 
**AuthorizationEndpoint** | **string** | URL of the OpenID Provider&#39;s OAuth 2.0 Authorization Endpoint. | 
**LoginType** | **string** | The OpenID Connect login type. These values maps to: &lt;br&gt;  CODE: Authentication using Code Flow &lt;br&gt; POST: Authentication using Form Post &lt;br&gt; POST_AT: Authentication using Form Post with Access Token | 
**AuthenticationScheme** | Pointer to **string** | The OpenID Connect Authentication Scheme. This is required for Authentication using Code Flow.  | [optional] 
**AuthenticationSigningAlgorithm** | Pointer to **string** | The authentication signing algorithm for token endpoint PRIVATE_KEY_JWT authentication. Only asymmetric algorithms are allowed. For RSASSA-PSS signing algorithm, PingFederate must be integrated with a hardware security module (HSM) or Java 11. | [optional] 
**RequestSigningAlgorithm** | Pointer to **string** | The request signing algorithm. Required only if you wish to use signed requests. Only asymmetric algorithms are allowed. For RSASSA-PSS signing algorithm, PingFederate must be integrated with a hardware security module (HSM) or Java 11. | [optional] 
**EnablePKCE** | Pointer to **bool** | Enable Proof Key for Code Exchange (PKCE). When enabled, the client sends an SHA-256 code challenge and corresponding code verifier to the OpenID Provider during the authorization code flow. | [optional] 
**TokenEndpoint** | Pointer to **string** | URL of the OpenID Provider&#39;s OAuth 2.0 Token Endpoint. | [optional] 
**UserInfoEndpoint** | Pointer to **string** | URL of the OpenID Provider&#39;s UserInfo Endpoint. | [optional] 
**JwksURL** | **string** | URL of the OpenID Provider&#39;s JSON Web Key Set [JWK] document. | 
**RequestParameters** | Pointer to [**[]OIDCRequestParameter**](OIDCRequestParameter.md) | A list of request parameters. Request parameters with same name but different attribute values are treated as a multi-valued request parameter. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


