// Copyright © 2026 Ping Identity Corporation

package oauth2

// TokenAuthType represents the method used to authenticate the client when requesting tokens.
// These types define how client credentials are transmitted to the authorization server
// during token requests in OAuth2 flows.
type TokenAuthType string

const (
	// TokenAuthTypeNone indicates no client authentication is required.
	// This is typically used for public clients that cannot securely store credentials.
	TokenAuthTypeNone TokenAuthType = "NONE"

	// TokenAuthTypeClientSecretBasic indicates client authentication using HTTP Basic authentication.
	// The client ID and secret are Base64-encoded and sent in the Authorization header.
	TokenAuthTypeClientSecretBasic TokenAuthType = "CLIENT_SECRET_BASIC"

	// TokenAuthTypeClientSecretPost indicates client authentication using POST parameters.
	// The client ID and secret are sent as form parameters in the request body.
	TokenAuthTypeClientSecretPost TokenAuthType = "CLIENT_SECRET_POST"
)
