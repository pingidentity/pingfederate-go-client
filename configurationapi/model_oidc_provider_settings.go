/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the OIDCProviderSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OIDCProviderSettings{}

// OIDCProviderSettings The OpenID Provider settings.
type OIDCProviderSettings struct {
	// Space separated scope values that the OpenID Provider supports.
	Scopes string `json:"scopes" tfsdk:"scopes"`
	// URL of the OpenID Provider's OAuth 2.0 Authorization Endpoint.
	AuthorizationEndpoint string `json:"authorizationEndpoint" tfsdk:"authorization_endpoint"`
	// URL of the OpenID Provider's OAuth 2.0 Pushed Authorization Request Endpoint.
	PushedAuthorizationRequestEndpoint *string `json:"pushedAuthorizationRequestEndpoint,omitempty" tfsdk:"pushed_authorization_request_endpoint"`
	// The OpenID Connect login type. These values maps to: <br>  CODE: Authentication using Code Flow <br> POST: Authentication using Form Post <br> POST_AT: Authentication using Form Post with Access Token
	LoginType string `json:"loginType" tfsdk:"login_type"`
	// The OpenID Connect Authentication Scheme. This is required for Authentication using Code Flow.
	AuthenticationScheme *string `json:"authenticationScheme,omitempty" tfsdk:"authentication_scheme"`
	// The authentication signing algorithm for token endpoint PRIVATE_KEY_JWT or CLIENT_SECRET_JWT authentication. Asymmetric algorithms are allowed for PRIVATE_KEY_JWT and symmetric algorithms are allowed for CLIENT_SECRET_JWT. For RSASSA-PSS signing algorithm, PingFederate must be integrated with a hardware security module (HSM) or Java 11.
	AuthenticationSigningAlgorithm *string `json:"authenticationSigningAlgorithm,omitempty" tfsdk:"authentication_signing_algorithm"`
	// The request signing algorithm. Required only if you wish to use signed requests. Only asymmetric algorithms are allowed. For RSASSA-PSS signing algorithm, PingFederate must be integrated with a hardware security module (HSM) or Java 11.
	RequestSigningAlgorithm *string `json:"requestSigningAlgorithm,omitempty" tfsdk:"request_signing_algorithm"`
	// Enable Proof Key for Code Exchange (PKCE). When enabled, the client sends an SHA-256 code challenge and corresponding code verifier to the OpenID Provider during the authorization code flow.
	EnablePKCE *bool `json:"enablePKCE,omitempty" tfsdk:"enable_pkce"`
	// URL of the OpenID Provider's OAuth 2.0 Token Endpoint.
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tfsdk:"token_endpoint"`
	// URL of the OpenID Provider's UserInfo Endpoint.
	UserInfoEndpoint *string `json:"userInfoEndpoint,omitempty" tfsdk:"user_info_endpoint"`
	// URL of the OpenID Provider's RP-Initiated Logout Endpoint.
	LogoutEndpoint *string `json:"logoutEndpoint,omitempty" tfsdk:"logout_endpoint"`
	// URL of the OpenID Provider's JSON Web Key Set [JWK] document.
	JwksURL string `json:"jwksURL" tfsdk:"jwks_url"`
	// Determines whether PingFederate tracks a logout entry when a user signs in, so that the user session can later be terminated via a logout request from the OP. This setting must also be enabled in order for PingFederate to send an RP-initiated logout request to the OP during SLO.
	TrackUserSessionsForLogout *bool `json:"trackUserSessionsForLogout,omitempty" tfsdk:"track_user_sessions_for_logout"`
	// A list of request parameters. Request parameters with same name but different attribute values are treated as a multi-valued request parameter.
	RequestParameters []OIDCRequestParameter `json:"requestParameters,omitempty" tfsdk:"request_parameters"`
	// The redirect URI. This is a read-only parameter.
	RedirectUri *string `json:"redirectUri,omitempty" tfsdk:"redirect_uri"`
	// The Back-Channel Logout URI. This read-only parameter is available when user sessions are tracked for logout.
	BackChannelLogoutUri *string `json:"backChannelLogoutUri,omitempty" tfsdk:"back_channel_logout_uri"`
	// The Front-Channel Logout URI. This is a read-only parameter.
	FrontChannelLogoutUri *string `json:"frontChannelLogoutUri,omitempty" tfsdk:"front_channel_logout_uri"`
	// The Post-Logout Redirect URI, where the OpenID Provider may redirect the user when RP-Initiated Logout has completed. This is a read-only parameter.
	PostLogoutRedirectUri *string `json:"postLogoutRedirectUri,omitempty" tfsdk:"post_logout_redirect_uri"`
}

type _OIDCProviderSettings OIDCProviderSettings

// NewOIDCProviderSettings instantiates a new OIDCProviderSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOIDCProviderSettings(scopes string, authorizationEndpoint string, loginType string, jwksURL string) *OIDCProviderSettings {
	this := OIDCProviderSettings{}
	this.Scopes = scopes
	this.AuthorizationEndpoint = authorizationEndpoint
	this.LoginType = loginType
	this.JwksURL = jwksURL
	return &this
}

// NewOIDCProviderSettingsWithDefaults instantiates a new OIDCProviderSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCProviderSettingsWithDefaults() *OIDCProviderSettings {
	this := OIDCProviderSettings{}
	return &this
}

// GetScopes returns the Scopes field value
func (o *OIDCProviderSettings) GetScopes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetScopesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scopes, true
}

// SetScopes sets field value
func (o *OIDCProviderSettings) SetScopes(v string) {
	o.Scopes = v
}

// GetAuthorizationEndpoint returns the AuthorizationEndpoint field value
func (o *OIDCProviderSettings) GetAuthorizationEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationEndpoint
}

// GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field value
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetAuthorizationEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationEndpoint, true
}

// SetAuthorizationEndpoint sets field value
func (o *OIDCProviderSettings) SetAuthorizationEndpoint(v string) {
	o.AuthorizationEndpoint = v
}

// GetPushedAuthorizationRequestEndpoint returns the PushedAuthorizationRequestEndpoint field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetPushedAuthorizationRequestEndpoint() string {
	if o == nil || IsNil(o.PushedAuthorizationRequestEndpoint) {
		var ret string
		return ret
	}
	return *o.PushedAuthorizationRequestEndpoint
}

// GetPushedAuthorizationRequestEndpointOk returns a tuple with the PushedAuthorizationRequestEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetPushedAuthorizationRequestEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.PushedAuthorizationRequestEndpoint) {
		return nil, false
	}
	return o.PushedAuthorizationRequestEndpoint, true
}

// HasPushedAuthorizationRequestEndpoint returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasPushedAuthorizationRequestEndpoint() bool {
	if o != nil && !IsNil(o.PushedAuthorizationRequestEndpoint) {
		return true
	}

	return false
}

// SetPushedAuthorizationRequestEndpoint gets a reference to the given string and assigns it to the PushedAuthorizationRequestEndpoint field.
func (o *OIDCProviderSettings) SetPushedAuthorizationRequestEndpoint(v string) {
	o.PushedAuthorizationRequestEndpoint = &v
}

// GetLoginType returns the LoginType field value
func (o *OIDCProviderSettings) GetLoginType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoginType
}

// GetLoginTypeOk returns a tuple with the LoginType field value
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetLoginTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginType, true
}

// SetLoginType sets field value
func (o *OIDCProviderSettings) SetLoginType(v string) {
	o.LoginType = v
}

// GetAuthenticationScheme returns the AuthenticationScheme field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetAuthenticationScheme() string {
	if o == nil || IsNil(o.AuthenticationScheme) {
		var ret string
		return ret
	}
	return *o.AuthenticationScheme
}

// GetAuthenticationSchemeOk returns a tuple with the AuthenticationScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetAuthenticationSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationScheme) {
		return nil, false
	}
	return o.AuthenticationScheme, true
}

// HasAuthenticationScheme returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasAuthenticationScheme() bool {
	if o != nil && !IsNil(o.AuthenticationScheme) {
		return true
	}

	return false
}

// SetAuthenticationScheme gets a reference to the given string and assigns it to the AuthenticationScheme field.
func (o *OIDCProviderSettings) SetAuthenticationScheme(v string) {
	o.AuthenticationScheme = &v
}

// GetAuthenticationSigningAlgorithm returns the AuthenticationSigningAlgorithm field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetAuthenticationSigningAlgorithm() string {
	if o == nil || IsNil(o.AuthenticationSigningAlgorithm) {
		var ret string
		return ret
	}
	return *o.AuthenticationSigningAlgorithm
}

// GetAuthenticationSigningAlgorithmOk returns a tuple with the AuthenticationSigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetAuthenticationSigningAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationSigningAlgorithm) {
		return nil, false
	}
	return o.AuthenticationSigningAlgorithm, true
}

// HasAuthenticationSigningAlgorithm returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasAuthenticationSigningAlgorithm() bool {
	if o != nil && !IsNil(o.AuthenticationSigningAlgorithm) {
		return true
	}

	return false
}

// SetAuthenticationSigningAlgorithm gets a reference to the given string and assigns it to the AuthenticationSigningAlgorithm field.
func (o *OIDCProviderSettings) SetAuthenticationSigningAlgorithm(v string) {
	o.AuthenticationSigningAlgorithm = &v
}

// GetRequestSigningAlgorithm returns the RequestSigningAlgorithm field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetRequestSigningAlgorithm() string {
	if o == nil || IsNil(o.RequestSigningAlgorithm) {
		var ret string
		return ret
	}
	return *o.RequestSigningAlgorithm
}

// GetRequestSigningAlgorithmOk returns a tuple with the RequestSigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetRequestSigningAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.RequestSigningAlgorithm) {
		return nil, false
	}
	return o.RequestSigningAlgorithm, true
}

// HasRequestSigningAlgorithm returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasRequestSigningAlgorithm() bool {
	if o != nil && !IsNil(o.RequestSigningAlgorithm) {
		return true
	}

	return false
}

// SetRequestSigningAlgorithm gets a reference to the given string and assigns it to the RequestSigningAlgorithm field.
func (o *OIDCProviderSettings) SetRequestSigningAlgorithm(v string) {
	o.RequestSigningAlgorithm = &v
}

// GetEnablePKCE returns the EnablePKCE field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetEnablePKCE() bool {
	if o == nil || IsNil(o.EnablePKCE) {
		var ret bool
		return ret
	}
	return *o.EnablePKCE
}

// GetEnablePKCEOk returns a tuple with the EnablePKCE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetEnablePKCEOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePKCE) {
		return nil, false
	}
	return o.EnablePKCE, true
}

// HasEnablePKCE returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasEnablePKCE() bool {
	if o != nil && !IsNil(o.EnablePKCE) {
		return true
	}

	return false
}

// SetEnablePKCE gets a reference to the given bool and assigns it to the EnablePKCE field.
func (o *OIDCProviderSettings) SetEnablePKCE(v bool) {
	o.EnablePKCE = &v
}

// GetTokenEndpoint returns the TokenEndpoint field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetTokenEndpoint() string {
	if o == nil || IsNil(o.TokenEndpoint) {
		var ret string
		return ret
	}
	return *o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetTokenEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.TokenEndpoint) {
		return nil, false
	}
	return o.TokenEndpoint, true
}

// HasTokenEndpoint returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasTokenEndpoint() bool {
	if o != nil && !IsNil(o.TokenEndpoint) {
		return true
	}

	return false
}

// SetTokenEndpoint gets a reference to the given string and assigns it to the TokenEndpoint field.
func (o *OIDCProviderSettings) SetTokenEndpoint(v string) {
	o.TokenEndpoint = &v
}

// GetUserInfoEndpoint returns the UserInfoEndpoint field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetUserInfoEndpoint() string {
	if o == nil || IsNil(o.UserInfoEndpoint) {
		var ret string
		return ret
	}
	return *o.UserInfoEndpoint
}

// GetUserInfoEndpointOk returns a tuple with the UserInfoEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetUserInfoEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.UserInfoEndpoint) {
		return nil, false
	}
	return o.UserInfoEndpoint, true
}

// HasUserInfoEndpoint returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasUserInfoEndpoint() bool {
	if o != nil && !IsNil(o.UserInfoEndpoint) {
		return true
	}

	return false
}

// SetUserInfoEndpoint gets a reference to the given string and assigns it to the UserInfoEndpoint field.
func (o *OIDCProviderSettings) SetUserInfoEndpoint(v string) {
	o.UserInfoEndpoint = &v
}

// GetLogoutEndpoint returns the LogoutEndpoint field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetLogoutEndpoint() string {
	if o == nil || IsNil(o.LogoutEndpoint) {
		var ret string
		return ret
	}
	return *o.LogoutEndpoint
}

// GetLogoutEndpointOk returns a tuple with the LogoutEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetLogoutEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.LogoutEndpoint) {
		return nil, false
	}
	return o.LogoutEndpoint, true
}

// HasLogoutEndpoint returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasLogoutEndpoint() bool {
	if o != nil && !IsNil(o.LogoutEndpoint) {
		return true
	}

	return false
}

// SetLogoutEndpoint gets a reference to the given string and assigns it to the LogoutEndpoint field.
func (o *OIDCProviderSettings) SetLogoutEndpoint(v string) {
	o.LogoutEndpoint = &v
}

// GetJwksURL returns the JwksURL field value
func (o *OIDCProviderSettings) GetJwksURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JwksURL
}

// GetJwksURLOk returns a tuple with the JwksURL field value
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetJwksURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JwksURL, true
}

// SetJwksURL sets field value
func (o *OIDCProviderSettings) SetJwksURL(v string) {
	o.JwksURL = v
}

// GetTrackUserSessionsForLogout returns the TrackUserSessionsForLogout field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetTrackUserSessionsForLogout() bool {
	if o == nil || IsNil(o.TrackUserSessionsForLogout) {
		var ret bool
		return ret
	}
	return *o.TrackUserSessionsForLogout
}

// GetTrackUserSessionsForLogoutOk returns a tuple with the TrackUserSessionsForLogout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetTrackUserSessionsForLogoutOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackUserSessionsForLogout) {
		return nil, false
	}
	return o.TrackUserSessionsForLogout, true
}

// HasTrackUserSessionsForLogout returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasTrackUserSessionsForLogout() bool {
	if o != nil && !IsNil(o.TrackUserSessionsForLogout) {
		return true
	}

	return false
}

// SetTrackUserSessionsForLogout gets a reference to the given bool and assigns it to the TrackUserSessionsForLogout field.
func (o *OIDCProviderSettings) SetTrackUserSessionsForLogout(v bool) {
	o.TrackUserSessionsForLogout = &v
}

// GetRequestParameters returns the RequestParameters field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetRequestParameters() []OIDCRequestParameter {
	if o == nil || IsNil(o.RequestParameters) {
		var ret []OIDCRequestParameter
		return ret
	}
	return o.RequestParameters
}

// GetRequestParametersOk returns a tuple with the RequestParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetRequestParametersOk() ([]OIDCRequestParameter, bool) {
	if o == nil || IsNil(o.RequestParameters) {
		return nil, false
	}
	return o.RequestParameters, true
}

// HasRequestParameters returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasRequestParameters() bool {
	if o != nil && !IsNil(o.RequestParameters) {
		return true
	}

	return false
}

// SetRequestParameters gets a reference to the given []OIDCRequestParameter and assigns it to the RequestParameters field.
func (o *OIDCProviderSettings) SetRequestParameters(v []OIDCRequestParameter) {
	o.RequestParameters = v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetRedirectUri() string {
	if o == nil || IsNil(o.RedirectUri) {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetRedirectUriOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUri) {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasRedirectUri() bool {
	if o != nil && !IsNil(o.RedirectUri) {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *OIDCProviderSettings) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetBackChannelLogoutUri returns the BackChannelLogoutUri field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetBackChannelLogoutUri() string {
	if o == nil || IsNil(o.BackChannelLogoutUri) {
		var ret string
		return ret
	}
	return *o.BackChannelLogoutUri
}

// GetBackChannelLogoutUriOk returns a tuple with the BackChannelLogoutUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetBackChannelLogoutUriOk() (*string, bool) {
	if o == nil || IsNil(o.BackChannelLogoutUri) {
		return nil, false
	}
	return o.BackChannelLogoutUri, true
}

// HasBackChannelLogoutUri returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasBackChannelLogoutUri() bool {
	if o != nil && !IsNil(o.BackChannelLogoutUri) {
		return true
	}

	return false
}

// SetBackChannelLogoutUri gets a reference to the given string and assigns it to the BackChannelLogoutUri field.
func (o *OIDCProviderSettings) SetBackChannelLogoutUri(v string) {
	o.BackChannelLogoutUri = &v
}

// GetFrontChannelLogoutUri returns the FrontChannelLogoutUri field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetFrontChannelLogoutUri() string {
	if o == nil || IsNil(o.FrontChannelLogoutUri) {
		var ret string
		return ret
	}
	return *o.FrontChannelLogoutUri
}

// GetFrontChannelLogoutUriOk returns a tuple with the FrontChannelLogoutUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetFrontChannelLogoutUriOk() (*string, bool) {
	if o == nil || IsNil(o.FrontChannelLogoutUri) {
		return nil, false
	}
	return o.FrontChannelLogoutUri, true
}

// HasFrontChannelLogoutUri returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasFrontChannelLogoutUri() bool {
	if o != nil && !IsNil(o.FrontChannelLogoutUri) {
		return true
	}

	return false
}

// SetFrontChannelLogoutUri gets a reference to the given string and assigns it to the FrontChannelLogoutUri field.
func (o *OIDCProviderSettings) SetFrontChannelLogoutUri(v string) {
	o.FrontChannelLogoutUri = &v
}

// GetPostLogoutRedirectUri returns the PostLogoutRedirectUri field value if set, zero value otherwise.
func (o *OIDCProviderSettings) GetPostLogoutRedirectUri() string {
	if o == nil || IsNil(o.PostLogoutRedirectUri) {
		var ret string
		return ret
	}
	return *o.PostLogoutRedirectUri
}

// GetPostLogoutRedirectUriOk returns a tuple with the PostLogoutRedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OIDCProviderSettings) GetPostLogoutRedirectUriOk() (*string, bool) {
	if o == nil || IsNil(o.PostLogoutRedirectUri) {
		return nil, false
	}
	return o.PostLogoutRedirectUri, true
}

// HasPostLogoutRedirectUri returns a boolean if a field has been set.
func (o *OIDCProviderSettings) HasPostLogoutRedirectUri() bool {
	if o != nil && !IsNil(o.PostLogoutRedirectUri) {
		return true
	}

	return false
}

// SetPostLogoutRedirectUri gets a reference to the given string and assigns it to the PostLogoutRedirectUri field.
func (o *OIDCProviderSettings) SetPostLogoutRedirectUri(v string) {
	o.PostLogoutRedirectUri = &v
}

func (o OIDCProviderSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OIDCProviderSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scopes"] = o.Scopes
	toSerialize["authorizationEndpoint"] = o.AuthorizationEndpoint
	if !IsNil(o.PushedAuthorizationRequestEndpoint) {
		toSerialize["pushedAuthorizationRequestEndpoint"] = o.PushedAuthorizationRequestEndpoint
	}
	toSerialize["loginType"] = o.LoginType
	if !IsNil(o.AuthenticationScheme) {
		toSerialize["authenticationScheme"] = o.AuthenticationScheme
	}
	if !IsNil(o.AuthenticationSigningAlgorithm) {
		toSerialize["authenticationSigningAlgorithm"] = o.AuthenticationSigningAlgorithm
	}
	if !IsNil(o.RequestSigningAlgorithm) {
		toSerialize["requestSigningAlgorithm"] = o.RequestSigningAlgorithm
	}
	if !IsNil(o.EnablePKCE) {
		toSerialize["enablePKCE"] = o.EnablePKCE
	}
	if !IsNil(o.TokenEndpoint) {
		toSerialize["tokenEndpoint"] = o.TokenEndpoint
	}
	if !IsNil(o.UserInfoEndpoint) {
		toSerialize["userInfoEndpoint"] = o.UserInfoEndpoint
	}
	if !IsNil(o.LogoutEndpoint) {
		toSerialize["logoutEndpoint"] = o.LogoutEndpoint
	}
	toSerialize["jwksURL"] = o.JwksURL
	if !IsNil(o.TrackUserSessionsForLogout) {
		toSerialize["trackUserSessionsForLogout"] = o.TrackUserSessionsForLogout
	}
	if !IsNil(o.RequestParameters) {
		toSerialize["requestParameters"] = o.RequestParameters
	}
	if !IsNil(o.RedirectUri) {
		toSerialize["redirectUri"] = o.RedirectUri
	}
	if !IsNil(o.BackChannelLogoutUri) {
		toSerialize["backChannelLogoutUri"] = o.BackChannelLogoutUri
	}
	if !IsNil(o.FrontChannelLogoutUri) {
		toSerialize["frontChannelLogoutUri"] = o.FrontChannelLogoutUri
	}
	if !IsNil(o.PostLogoutRedirectUri) {
		toSerialize["postLogoutRedirectUri"] = o.PostLogoutRedirectUri
	}
	return toSerialize, nil
}

func (o *OIDCProviderSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"scopes",
		"authorizationEndpoint",
		"loginType",
		"jwksURL",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOIDCProviderSettings := _OIDCProviderSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varOIDCProviderSettings)

	if err != nil {
		return err
	}

	*o = OIDCProviderSettings(varOIDCProviderSettings)

	return err
}

type NullableOIDCProviderSettings struct {
	value *OIDCProviderSettings
	isSet bool
}

func (v NullableOIDCProviderSettings) Get() *OIDCProviderSettings {
	return v.value
}

func (v *NullableOIDCProviderSettings) Set(val *OIDCProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOIDCProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOIDCProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOIDCProviderSettings(val *OIDCProviderSettings) *NullableOIDCProviderSettings {
	return &NullableOIDCProviderSettings{value: val, isSet: true}
}

func (v NullableOIDCProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOIDCProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
