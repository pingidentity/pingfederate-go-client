// Copyright © 2026 Ping Identity Corporation

// Package config provides configuration management and OAuth2 human-login flows
// (authorization_code, device_code) and the client_credentials flow for the PingFederate
// Go Client SDK. It exposes a central Configuration type with a fluent builder that produces
// an oauth2.TokenSource, which can be supplied to the generated client via
// context.WithValue(ctx, configurationapi.ContextOAuth2, ts) or by Configuration.Client.
package config

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	svcOAuth2 "github.com/pingidentity/pingfederate-go-client/v1300/oauth2"
	"github.com/pingidentity/pingfederate-go-client/v1300/oauth2/endpoints"
	"golang.org/x/oauth2"
)

// keychainTokenSource wraps an oauth2.TokenSource and saves tokens to the keychain after refresh.
type keychainTokenSource struct {
	base            oauth2.TokenSource
	keychainStorage *svcOAuth2.KeychainStorage
	tokenKey        string
}

// Token implements the oauth2.TokenSource interface. It obtains a token from the base source
// and persists it to the keychain so a refreshed token survives across process invocations.
func (k *keychainTokenSource) Token() (*oauth2.Token, error) {
	token, err := k.base.Token()
	if err != nil {
		return nil, err
	}

	if saveErr := k.keychainStorage.SaveToken(token); saveErr != nil {
		slog.Warn("Failed to save refreshed token to keychain", "error", saveErr, "tokenKey", k.tokenKey)
		// Don't return an error - the token is still valid even if caching failed.
	} else {
		slog.Debug("Refreshed token saved to keychain", "tokenKey", k.tokenKey, "expires", token.Expiry)
	}

	return token, nil
}

// Storage configures how tokens obtained by the SDK are persisted between invocations.
type Storage struct {
	// KeychainName is the keychain service name under which tokens are stored.
	KeychainName string
	// Type selects the storage backend. When empty, secure local (keychain) storage is used.
	Type StorageType
	// OptionalSuffix allows consumers to append a suffix to the generated token key for
	// disambiguation across contexts (e.g. provider/grant/profile). If empty, no suffix is appended.
	OptionalSuffix string
}

// Configuration represents the OAuth2 authentication configuration for the PingFederate Go
// Client SDK. It holds the runtime endpoint and the credentials for the configured grant type,
// and produces an oauth2.TokenSource via TokenSource. Populate it using the builder methods (With...).
type Configuration struct {
	// Auth contains authentication-related configuration including credentials and grant type.
	Auth struct {
		AccessToken       *string
		AuthorizationCode *AuthorizationCode
		ClientCredentials *ClientCredentials
		DeviceCode        *DeviceCode
		GrantType         *svcOAuth2.GrantType
		Storage           *Storage
	}
	// Endpoint contains the endpoint configuration used to construct the OAuth2 endpoints.
	Endpoint struct {
		// RuntimeBaseURL is the PingFederate runtime engine base URL (scheme, host, optional port),
		// for example "https://pingfederate.example.com:9031". The standard PingFederate OAuth2 paths
		// are derived from it unless an ExplicitEndpoint is set.
		RuntimeBaseURL *string
		// ExplicitEndpoint, when set, overrides RuntimeBaseURL and is used verbatim. It is intended
		// for deployments whose OAuth2 endpoints do not follow the standard runtime paths.
		ExplicitEndpoint *oauth2.Endpoint
	}
}

// NewConfiguration creates a new Configuration instance with default (empty) values.
// Populate it with a runtime base URL, grant type, and credentials using the builder methods
// (With...) before creating a token source or client.
func NewConfiguration() *Configuration {
	return &Configuration{}
}

// GetConfiguration returns the configuration itself, for parity with fluent-builder consumers.
func (c *Configuration) GetConfiguration() *Configuration {
	return c
}

// WithRuntimeBaseURL sets the PingFederate runtime engine base URL used to derive the OAuth2
// endpoints (for example "https://pingfederate.example.com:9031").
func (c *Configuration) WithRuntimeBaseURL(runtimeBaseURL string) *Configuration {
	c.Endpoint.RuntimeBaseURL = &runtimeBaseURL
	return c
}

// WithExplicitEndpoint sets the OAuth2 endpoint explicitly, overriding the runtime base URL.
// Use this for deployments whose OAuth2 endpoints do not follow the standard PingFederate paths.
func (c *Configuration) WithExplicitEndpoint(endpoint oauth2.Endpoint) *Configuration {
	c.Endpoint.ExplicitEndpoint = &endpoint
	return c
}

// WithGrantType sets the OAuth2 grant type used for token acquisition.
func (c *Configuration) WithGrantType(grantType svcOAuth2.GrantType) *Configuration {
	c.Auth.GrantType = &grantType
	return c
}

// WithAccessToken sets a static access token for API authentication. When set (and the grant type
// is not client_credentials), this bypasses OAuth2 flows and uses the provided token directly.
func (c *Configuration) WithAccessToken(accessToken string) *Configuration {
	c.Auth.AccessToken = &accessToken
	return c
}

// WithClientCredentialsClientID sets the OAuth2 client ID for the client_credentials grant type.
func (c *Configuration) WithClientCredentialsClientID(clientID string) *Configuration {
	if c.Auth.ClientCredentials == nil {
		c.Auth.ClientCredentials = &ClientCredentials{}
	}
	c.Auth.ClientCredentials.ClientCredentialsClientID = &clientID
	return c
}

// WithClientCredentialsClientSecret sets the OAuth2 client secret for the client_credentials grant type.
func (c *Configuration) WithClientCredentialsClientSecret(clientSecret string) *Configuration {
	if c.Auth.ClientCredentials == nil {
		c.Auth.ClientCredentials = &ClientCredentials{}
	}
	c.Auth.ClientCredentials.ClientCredentialsClientSecret = &clientSecret
	return c
}

// WithClientCredentialsScopes sets the OAuth2 scopes requested by the client_credentials grant type.
func (c *Configuration) WithClientCredentialsScopes(scopes []string) *Configuration {
	if c.Auth.ClientCredentials == nil {
		c.Auth.ClientCredentials = &ClientCredentials{}
	}
	c.Auth.ClientCredentials.ClientCredentialsScopes = &scopes
	return c
}

// WithAuthorizationCodeClientID sets the OAuth2 client ID for the authorization_code grant type.
func (c *Configuration) WithAuthorizationCodeClientID(clientID string) *Configuration {
	if c.Auth.AuthorizationCode == nil {
		c.Auth.AuthorizationCode = &AuthorizationCode{}
	}
	c.Auth.AuthorizationCode.AuthorizationCodeClientID = &clientID
	return c
}

// WithAuthorizationCodeScopes sets the OAuth2 scopes requested by the authorization_code grant type.
func (c *Configuration) WithAuthorizationCodeScopes(scopes []string) *Configuration {
	if c.Auth.AuthorizationCode == nil {
		c.Auth.AuthorizationCode = &AuthorizationCode{}
	}
	c.Auth.AuthorizationCode.AuthorizationCodeScopes = &scopes
	return c
}

// WithAuthorizationCodeRedirectURI sets the local redirect (callback) server port and path
// used by the authorization_code grant type.
func (c *Configuration) WithAuthorizationCodeRedirectURI(redirectURI AuthorizationCodeRedirectURI) *Configuration {
	if c.Auth.AuthorizationCode == nil {
		c.Auth.AuthorizationCode = &AuthorizationCode{}
	}
	c.Auth.AuthorizationCode.AuthorizationCodeRedirectURI = redirectURI
	return c
}

// WithAuthorizationCodeOutput sets the writer used by the default authorization_code browser
// handler to print its progress messages. It is honored only when no custom OnOpenBrowser
// handler is set; once OnOpenBrowser is set, that handler is solely responsible for its own
// output and this writer is ignored. The SDK stays quiet by default: pass nil (or leave it
// unset) to keep progress output disabled, os.Stdout to reproduce the interactive v1300.1.0
// behavior, or any other io.Writer to capture or redirect it.
func (c *Configuration) WithAuthorizationCodeOutput(w io.Writer) *Configuration {
	if c.Auth.AuthorizationCode == nil {
		c.Auth.AuthorizationCode = &AuthorizationCode{}
	}
	c.Auth.AuthorizationCode.Output = w
	return c
}

// WithDeviceCodeClientID sets the OAuth2 client ID for the device_code grant type.
func (c *Configuration) WithDeviceCodeClientID(clientID string) *Configuration {
	if c.Auth.DeviceCode == nil {
		c.Auth.DeviceCode = &DeviceCode{}
	}
	c.Auth.DeviceCode.DeviceCodeClientID = &clientID
	return c
}

// WithDeviceCodeScopes sets the OAuth2 scopes requested by the device_code grant type.
func (c *Configuration) WithDeviceCodeScopes(scopes []string) *Configuration {
	if c.Auth.DeviceCode == nil {
		c.Auth.DeviceCode = &DeviceCode{}
	}
	c.Auth.DeviceCode.DeviceCodeScopes = &scopes
	return c
}

// WithDeviceCodeOutput sets the writer used by the default device_code prompt handler to print
// its progress messages. It is honored only when no custom OnDisplayPrompt handler is set; once
// OnDisplayPrompt is set, that handler is solely responsible for its own output and this writer
// is ignored. The SDK stays quiet by default: pass nil (or leave it unset) to keep progress
// output disabled, os.Stdout to reproduce the interactive v1300.1.0 behavior, or any other
// io.Writer to capture or redirect it.
func (c *Configuration) WithDeviceCodeOutput(w io.Writer) *Configuration {
	if c.Auth.DeviceCode == nil {
		c.Auth.DeviceCode = &DeviceCode{}
	}
	c.Auth.DeviceCode.Output = w
	return c
}

// WithStorageType sets the token storage backend.
func (c *Configuration) WithStorageType(storageType StorageType) *Configuration {
	if c.Auth.Storage == nil {
		c.Auth.Storage = &Storage{}
	}
	c.Auth.Storage.Type = storageType
	return c
}

// WithStorageName sets the keychain service name used for token storage.
func (c *Configuration) WithStorageName(name string) *Configuration {
	if c.Auth.Storage == nil {
		c.Auth.Storage = &Storage{}
	}
	c.Auth.Storage.KeychainName = name
	return c
}

// WithStorageOptionalSuffix sets an optional suffix appended to the generated token key,
// allowing consumers to disambiguate keys across contexts (e.g. provider/grant/profile).
func (c *Configuration) WithStorageOptionalSuffix(suffix string) *Configuration {
	if c.Auth.Storage == nil {
		c.Auth.Storage = &Storage{}
	}
	c.Auth.Storage.OptionalSuffix = suffix
	return c
}

// HasBearerToken reports whether a non-empty static access token is configured.
func (c *Configuration) HasBearerToken() bool {
	return c.Auth.AccessToken != nil && *c.Auth.AccessToken != ""
}

// AddBearerTokenToContext adds the configured static access token to the context under key.
// If no static token is configured, the parent context is returned unchanged.
func (c *Configuration) AddBearerTokenToContext(parent context.Context, key any) context.Context {
	if c.HasBearerToken() {
		return context.WithValue(parent, key, *c.Auth.AccessToken)
	}
	return parent
}

// BearerToken creates an OAuth2 token from the configured static access token. Callers should
// verify a token is available via HasBearerToken first.
func (c *Configuration) BearerToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: *c.Auth.AccessToken,
		TokenType:   "Bearer",
	}
}

// AuthEndpoints constructs the OAuth2 endpoint used for authentication. An explicitly-configured
// endpoint takes precedence; otherwise the endpoint is derived from the runtime base URL. It
// returns an error if neither is configured or the runtime base URL is invalid.
func (c *Configuration) AuthEndpoints() (oauth2.Endpoint, error) {
	if ep := c.Endpoint.ExplicitEndpoint; ep != nil {
		return *ep, nil
	}

	if c.Endpoint.RuntimeBaseURL != nil && *c.Endpoint.RuntimeBaseURL != "" {
		return endpoints.PingFederateEndpoint(*c.Endpoint.RuntimeBaseURL)
	}

	return oauth2.Endpoint{}, fmt.Errorf("no valid endpoint configuration found. Must provide a runtime base URL or an explicit endpoint")
}

// Client creates an HTTP client that automatically manages OAuth2 authentication tokens for
// PingFederate API calls. httpClient, if non-nil, is used both as the transport for any HTTP
// calls TokenSource itself makes (e.g. the browser/device-code token exchange, or minting a
// client_credentials token) and as the base transport for the returned client. If httpClient is
// nil, a default client is used. It returns an error if a token source cannot be created.
func (c *Configuration) Client(ctx context.Context, httpClient *http.Client) (*http.Client, error) {
	if httpClient != nil {
		ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	}

	ts, err := c.TokenSource(ctx)
	if err != nil {
		return nil, err
	}

	return oauth2.NewClient(ctx, ts), nil
}

// GetAccessToken resolves an access token using the configured grant type. It returns an error
// if a token source cannot be created or the resolved token is empty.
func (c *Configuration) GetAccessToken(ctx context.Context) (string, error) {
	ts, err := c.TokenSource(ctx)
	if err != nil {
		return "", err
	}

	token, err := ts.Token()
	if err != nil {
		return "", err
	}

	if token.AccessToken == "" {
		return "", fmt.Errorf("failed to retrieve access token")
	}

	return token.AccessToken, nil
}

// storageEnabled reports whether keychain persistence is active for the current configuration.
// Storage defaults to secure-local (keychain) when unset.
func (c *Configuration) storageEnabled() bool {
	return c.Auth.Storage == nil || c.Auth.Storage.Type == "" || c.Auth.Storage.Type == StorageTypeSecureLocal
}

// generateTokenKey derives the keychain account name for the configured grant type from the
// runtime base URL, client ID, and grant type (plus an optional suffix).
func (c *Configuration) generateTokenKey(grantType svcOAuth2.GrantType) (string, error) {
	var clientID string

	switch grantType {
	case svcOAuth2.GrantTypeAuthorizationCode:
		if c.Auth.AuthorizationCode != nil && c.Auth.AuthorizationCode.AuthorizationCodeClientID != nil {
			clientID = *c.Auth.AuthorizationCode.AuthorizationCodeClientID
		}
	case svcOAuth2.GrantTypeClientCredentials:
		if c.Auth.ClientCredentials != nil && c.Auth.ClientCredentials.ClientCredentialsClientID != nil {
			clientID = *c.Auth.ClientCredentials.ClientCredentialsClientID
		}
	case svcOAuth2.GrantTypeDeviceCode:
		if c.Auth.DeviceCode != nil && c.Auth.DeviceCode.DeviceCodeClientID != nil {
			clientID = *c.Auth.DeviceCode.DeviceCodeClientID
		}
	default:
		return "", fmt.Errorf("unsupported grant type: %s", grantType)
	}

	// The runtime base URL disambiguates tokens across PingFederate deployments. When an explicit
	// endpoint is configured instead (no runtime base URL), fall back to its token URL so keychain
	// storage still works for explicit-endpoint deployments.
	var deploymentKey string
	switch {
	case c.Endpoint.RuntimeBaseURL != nil && *c.Endpoint.RuntimeBaseURL != "":
		deploymentKey = *c.Endpoint.RuntimeBaseURL
	case c.Endpoint.ExplicitEndpoint != nil && c.Endpoint.ExplicitEndpoint.TokenURL != "":
		deploymentKey = c.Endpoint.ExplicitEndpoint.TokenURL
	}

	if deploymentKey == "" || clientID == "" {
		return "", fmt.Errorf("a runtime base URL (or explicit endpoint token URL) and client ID are required for token key generation")
	}

	var suffix string
	if c.Auth.Storage != nil {
		suffix = c.Auth.Storage.OptionalSuffix
	}

	tokenKey := svcOAuth2.GenerateKeychainAccountNameWithSuffix(deploymentKey, clientID, string(grantType), suffix)
	slog.Debug("Generated token key", "clientID", clientID, "grantType", grantType, "tokenKey", tokenKey)

	return tokenKey, nil
}

// createOAuth2ConfigForRefresh builds a minimal oauth2.Config capable of refreshing a token for
// the configured interactive grant type.
func (c *Configuration) createOAuth2ConfigForRefresh(endpoint oauth2.Endpoint) (*oauth2.Config, error) {
	if c.Auth.GrantType == nil {
		return nil, fmt.Errorf("grant type is required")
	}

	switch *c.Auth.GrantType {
	case svcOAuth2.GrantTypeDeviceCode:
		if c.Auth.DeviceCode != nil && c.Auth.DeviceCode.DeviceCodeClientID != nil {
			var scopes []string
			if c.Auth.DeviceCode.DeviceCodeScopes != nil {
				scopes = *c.Auth.DeviceCode.DeviceCodeScopes
			}
			return &oauth2.Config{
				ClientID: *c.Auth.DeviceCode.DeviceCodeClientID,
				Endpoint: endpoint,
				Scopes:   scopes,
			}, nil
		}
	case svcOAuth2.GrantTypeAuthorizationCode:
		if c.Auth.AuthorizationCode != nil && c.Auth.AuthorizationCode.AuthorizationCodeClientID != nil {
			var scopes []string
			if c.Auth.AuthorizationCode.AuthorizationCodeScopes != nil {
				scopes = *c.Auth.AuthorizationCode.AuthorizationCodeScopes
			}
			return &oauth2.Config{
				ClientID: *c.Auth.AuthorizationCode.AuthorizationCodeClientID,
				Endpoint: endpoint,
				Scopes:   scopes,
			}, nil
		}
	}

	return nil, fmt.Errorf("no valid configuration found for grant type: %s", *c.Auth.GrantType)
}

// loadCachedTokenSource attempts to build a refreshing token source from a keychain-cached token.
// It returns (nil, nil) when there is no usable cached token, so the caller falls through to an
// interactive login.
func (c *Configuration) loadCachedTokenSource(ctx context.Context, grantType svcOAuth2.GrantType) (oauth2.TokenSource, error) {
	if !c.storageEnabled() {
		return nil, nil
	}

	if c.Auth.Storage == nil || c.Auth.Storage.KeychainName == "" {
		return nil, fmt.Errorf("storage name is required when using keychain storage. Use WithStorageName() to set it")
	}

	tokenKey, err := c.generateTokenKey(grantType)
	if err != nil {
		return nil, fmt.Errorf("could not generate token key for caching: %w", err)
	}

	keychainStorage, err := svcOAuth2.NewKeychainStorage(c.Auth.Storage.KeychainName, tokenKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create keychain storage: %w", err)
	}

	existingToken, err := keychainStorage.LoadToken()
	if err != nil || existingToken == nil || existingToken.RefreshToken == "" {
		if err != nil {
			slog.Debug("Unable to load cached token from keychain", "error", err)
		}
		return nil, nil
	}

	endpoint, err := c.AuthEndpoints()
	if err != nil {
		slog.Warn("Failed to get endpoints for token refresh", "error", err)
		return nil, nil
	}

	oauthConfig, err := c.createOAuth2ConfigForRefresh(endpoint)
	if err != nil {
		slog.Warn("Failed to create OAuth2 config for refresh", "error", err)
		return nil, nil
	}

	baseTS := oauthConfig.TokenSource(ctx, existingToken)
	ts := oauth2.ReuseTokenSource(nil, &keychainTokenSource{
		base:            baseTS,
		keychainStorage: keychainStorage,
		tokenKey:        tokenKey,
	})

	// Verify the token source works (uses the cached token if valid, or refreshes if expired).
	if _, err := ts.Token(); err != nil {
		slog.Debug("Cached token source failed, will re-authenticate", "error", err)
		return nil, nil
	}

	slog.Debug("Using cached token source", "tokenKey", tokenKey)
	return ts, nil
}

// persistAndWrapTokenSource caches the freshly-obtained token and, when the token is refreshable,
// wraps it in a keychain-persisting refreshing source. It returns a token source suitable for use.
func (c *Configuration) persistAndWrapTokenSource(ctx context.Context, grantType svcOAuth2.GrantType, tokenSource oauth2.TokenSource) (oauth2.TokenSource, error) {
	token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	endpoint, endpointErr := c.AuthEndpoints()

	if c.storageEnabled() {
		if c.Auth.Storage == nil || c.Auth.Storage.KeychainName == "" {
			return nil, fmt.Errorf("storage name is required when using keychain storage. Use WithStorageName() to set it")
		}

		tokenKey, err := c.generateTokenKey(grantType)
		if err != nil {
			return nil, fmt.Errorf("failed to generate token key: %w", err)
		}

		keychainStorage, err := svcOAuth2.NewKeychainStorage(c.Auth.Storage.KeychainName, tokenKey)
		if err != nil {
			return nil, fmt.Errorf("failed to create keychain storage: %w", err)
		}

		if err := keychainStorage.SaveToken(token); err != nil {
			slog.Warn("Failed to save token to keychain", "error", err, "tokenKey", tokenKey)
			// Don't return an error - the token is still valid even if caching failed.
		} else {
			slog.Debug("Token saved to keychain", "tokenKey", tokenKey, "expires", token.Expiry)
		}

		if token.RefreshToken != "" && endpointErr == nil {
			if oauthConfig, cfgErr := c.createOAuth2ConfigForRefresh(endpoint); cfgErr == nil {
				baseTS := oauthConfig.TokenSource(ctx, token)
				return oauth2.ReuseTokenSource(nil, &keychainTokenSource{
					base:            baseTS,
					keychainStorage: keychainStorage,
					tokenKey:        tokenKey,
				}), nil
			}
		}

		return oauth2.StaticTokenSource(token), nil
	}

	// Storage disabled: still set up automatic refresh without keychain persistence when possible.
	if token.RefreshToken != "" && endpointErr == nil {
		if oauthConfig, cfgErr := c.createOAuth2ConfigForRefresh(endpoint); cfgErr == nil {
			return oauth2.ReuseTokenSource(nil, oauthConfig.TokenSource(ctx, token)), nil
		}
	}

	return oauth2.StaticTokenSource(token), nil
}

// TokenSource creates an OAuth2 token source based on the configured authentication method.
// If a static access token is configured (and the grant type is not client_credentials), a static
// token source is returned. Otherwise a token source appropriate for the configured grant type is
// created, reusing a keychain-cached token when one is available.
func (c *Configuration) TokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	// A static access token is used directly, except for client_credentials which always uses the
	// OAuth2 flow so tokens can be refreshed automatically.
	if c.HasBearerToken() {
		if c.Auth.GrantType != nil && *c.Auth.GrantType == svcOAuth2.GrantTypeClientCredentials {
			slog.Debug("Skipping static token source for client credentials flow to enable automatic token refresh")
		} else {
			slog.Debug("Using static token source as an access token has been provided")
			return oauth2.StaticTokenSource(c.BearerToken()), nil
		}
	}

	if c.Auth.GrantType == nil {
		return nil, fmt.Errorf("grant type is required")
	}
	grantType := *c.Auth.GrantType

	// The interactive grants (authorization_code, device_code) support token refresh, so try a
	// keychain-cached token before triggering a new user login. client_credentials mints silently.
	if grantType == svcOAuth2.GrantTypeAuthorizationCode || grantType == svcOAuth2.GrantTypeDeviceCode {
		cached, err := c.loadCachedTokenSource(ctx, grantType)
		if err != nil {
			return nil, err
		}
		if cached != nil {
			return cached, nil
		}
	}

	endpoint, err := c.AuthEndpoints()
	if err != nil {
		return nil, err
	}

	switch grantType {
	case svcOAuth2.GrantTypeAuthorizationCode:
		if c.Auth.AuthorizationCode == nil {
			return nil, fmt.Errorf("authorization code configuration is required for authorization_code grant type")
		}
		ts, err := c.Auth.AuthorizationCode.AuthorizationCodeTokenSource(ctx, endpoint)
		if err != nil {
			return nil, err
		}
		return c.persistAndWrapTokenSource(ctx, grantType, ts)
	case svcOAuth2.GrantTypeDeviceCode:
		if c.Auth.DeviceCode == nil {
			return nil, fmt.Errorf("device code configuration is required for device_code grant type")
		}
		ts, err := c.Auth.DeviceCode.DeviceAuthTokenSource(ctx, endpoint)
		if err != nil {
			return nil, err
		}
		return c.persistAndWrapTokenSource(ctx, grantType, ts)
	case svcOAuth2.GrantTypeClientCredentials:
		if c.Auth.ClientCredentials == nil {
			return nil, fmt.Errorf("client credentials configuration is required for client_credentials grant type")
		}
		// The client_credentials source refreshes itself; return it directly so it can re-mint on expiry.
		return c.Auth.ClientCredentials.ClientCredentialsTokenSource(ctx, endpoint)
	default:
		return nil, fmt.Errorf("unsupported grant type: %s", grantType)
	}
}
