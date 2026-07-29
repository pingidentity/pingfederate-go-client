// Copyright © 2026 Ping Identity Corporation

// Package oauth2 provides OAuth2 authentication utilities for the PingFederate Go Client SDK.
// It includes grant type definitions, token authentication methods, token storage, and endpoint
// configuration for the OAuth2 flows supported by PingFederate services.
package oauth2

// GrantType represents the OAuth2 grant type used for token acquisition.
// Grant types define the method by which applications obtain access tokens from the authorization server.
type GrantType string

const (
	// GrantTypeAuthorizationCode represents the authorization code grant type (with PKCE).
	GrantTypeAuthorizationCode GrantType = "authorization_code"

	// GrantTypeClientCredentials represents the client credentials grant type.
	// This grant type is used for server-to-server authentication where the client
	// authenticates directly with the authorization server using its client credentials.
	GrantTypeClientCredentials GrantType = "client_credentials"

	// GrantTypeDeviceCode represents the device authorization grant type (with PKCE).
	GrantTypeDeviceCode GrantType = "device_code"
)

// AllowedTokenAuthMethods maps each grant type to its supported token authentication methods.
// This mapping ensures that only compatible authentication methods are used with each grant type.
// The map helps validate authentication configurations and provides available options for each flow.
var AllowedTokenAuthMethods = map[GrantType][]TokenAuthType{
	GrantTypeAuthorizationCode: {
		TokenAuthTypeNone,
	},
	GrantTypeClientCredentials: {
		TokenAuthTypeClientSecretBasic,
		TokenAuthTypeClientSecretPost,
	},
	GrantTypeDeviceCode: {
		TokenAuthTypeNone,
	},
}

// IsValidGrantType reports whether gt is a grant type supported by this package.
func IsValidGrantType(gt string) bool {
	switch GrantType(gt) {
	case GrantTypeAuthorizationCode, GrantTypeClientCredentials, GrantTypeDeviceCode:
		return true
	}
	return false
}
