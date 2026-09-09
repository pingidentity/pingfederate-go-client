// Copyright © 2026 Ping Identity Corporation

package config_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/pingidentity/pingfederate-go-client/v1310/config"
	svcOAuth2 "github.com/pingidentity/pingfederate-go-client/v1310/oauth2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
	"golang.org/x/oauth2"
)

const (
	testRuntimeURL = "https://pf.example.com:9031"
	testClientID   = "matrix-client"
	testStorage    = "pingfederate-matrix-test"
)

// newTokenServer returns an httptest server that answers OAuth2 token requests with a fixed
// access token, so the client_credentials flow can be exercised end to end without a live PF.
func newTokenServer(t *testing.T, accessToken string) *httptest.Server {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if r.PostForm.Get("grant_type") != "client_credentials" {
			http.Error(w, "unexpected grant_type", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": accessToken,
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	t.Cleanup(srv.Close)
	return srv
}

// seedKeychainToken stores a valid (non-expired, refreshable) token under the exact key the
// Configuration derives for the given grant type, so the interactive-grant cache path resolves
// without a browser or device login.
func seedKeychainToken(t *testing.T, deploymentKey, clientID, grantType, accessToken string) {
	t.Helper()

	key := svcOAuth2.GenerateKeychainAccountNameWithSuffix(deploymentKey, clientID, grantType, "")
	storage, err := svcOAuth2.NewKeychainStorage(testStorage, key)
	require.NoError(t, err)

	require.NoError(t, storage.SaveToken(&oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: "refresh-" + accessToken,
		TokenType:    "Bearer",
		Expiry:       time.Now().Add(time.Hour),
	}))
}

// --- Storage type enum ---------------------------------------------------------------------

func TestStorageTypeValidity(t *testing.T) {
	assert.True(t, config.StorageTypeSecureLocal.IsValid())
	assert.True(t, config.StorageTypeNone.IsValid())
	assert.False(t, config.StorageType("bogus").IsValid())

	assert.Equal(t, "secure_local", config.StorageTypeSecureLocal.String())
	assert.Equal(t, "none", config.StorageTypeNone.String())
}

// --- client_credentials × {SecureLocal, None} ----------------------------------------------

func TestMatrix_ClientCredentials_AllStorageTypes(t *testing.T) {
	keyring.MockInit()

	for _, storageType := range []config.StorageType{config.StorageTypeSecureLocal, config.StorageTypeNone, ""} {
		t.Run("storage="+string(storageType), func(t *testing.T) {
			srv := newTokenServer(t, "cc-access-token")

			cfg := config.NewConfiguration().
				WithExplicitEndpoint(oauth2.Endpoint{TokenURL: srv.URL}).
				WithGrantType(svcOAuth2.GrantTypeClientCredentials).
				WithClientCredentialsClientID(testClientID).
				WithClientCredentialsClientSecret("secret").
				WithClientCredentialsScopes([]string{"openid"}).
				WithStorageType(storageType).
				WithStorageName(testStorage)

			token, err := cfg.GetAccessToken(context.Background())
			require.NoError(t, err)
			assert.Equal(t, "cc-access-token", token)
		})
	}
}

// --- device_code × {SecureLocal (cache hit), None (cache bypassed)} -------------------------

func TestMatrix_DeviceCode_SecureLocal_CacheHit(t *testing.T) {
	keyring.MockInit()
	seedKeychainToken(t, testRuntimeURL, testClientID, "device_code", "cached-device-token")

	cfg := config.NewConfiguration().
		WithRuntimeBaseURL(testRuntimeURL).
		WithGrantType(svcOAuth2.GrantTypeDeviceCode).
		WithDeviceCodeClientID(testClientID).
		WithDeviceCodeScopes([]string{"openid"}).
		WithStorageType(config.StorageTypeSecureLocal).
		WithStorageName(testStorage)

	// A valid cached token must be returned without initiating a device login.
	token, err := cfg.GetAccessToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "cached-device-token", token)
}

func TestMatrix_DeviceCode_None_BypassesCache(t *testing.T) {
	keyring.MockInit()
	// Seed a valid token; storage=None must ignore it and attempt a real device flow, which
	// fails against the unreachable runtime endpoint.
	seedKeychainToken(t, testRuntimeURL, testClientID, "device_code", "should-be-ignored")

	cfg := config.NewConfiguration().
		WithRuntimeBaseURL(testRuntimeURL).
		WithGrantType(svcOAuth2.GrantTypeDeviceCode).
		WithDeviceCodeClientID(testClientID).
		WithStorageType(config.StorageTypeNone).
		WithStorageName(testStorage)

	_, err := cfg.TokenSource(context.Background())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "device auth request failed",
		"storage=None must bypass the cache and attempt the device flow")
}

// --- authorization_code × {SecureLocal (cache hit), None (cache bypassed)} ------------------

func TestMatrix_AuthorizationCode_SecureLocal_CacheHit(t *testing.T) {
	keyring.MockInit()
	seedKeychainToken(t, testRuntimeURL, testClientID, "authorization_code", "cached-authcode-token")

	cfg := config.NewConfiguration().
		WithRuntimeBaseURL(testRuntimeURL).
		WithGrantType(svcOAuth2.GrantTypeAuthorizationCode).
		WithAuthorizationCodeClientID(testClientID).
		WithAuthorizationCodeScopes([]string{"openid"}).
		WithStorageType(config.StorageTypeSecureLocal).
		WithStorageName(testStorage)

	// A valid cached token must be returned without launching a browser flow.
	token, err := cfg.GetAccessToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "cached-authcode-token", token)
}

func TestMatrix_AuthorizationCode_SecureLocal_ExplicitEndpoint_CacheHit(t *testing.T) {
	keyring.MockInit()

	explicit := oauth2.Endpoint{
		AuthURL:  "https://custom.example.com/authorize",
		TokenURL: "https://custom.example.com/token",
	}
	// With an explicit endpoint (no runtime base URL), the token key falls back to the token URL.
	seedKeychainToken(t, explicit.TokenURL, testClientID, "authorization_code", "explicit-cached-token")

	cfg := config.NewConfiguration().
		WithExplicitEndpoint(explicit).
		WithGrantType(svcOAuth2.GrantTypeAuthorizationCode).
		WithAuthorizationCodeClientID(testClientID).
		WithStorageType(config.StorageTypeSecureLocal).
		WithStorageName(testStorage)

	token, err := cfg.GetAccessToken(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "explicit-cached-token", token)
}

func TestMatrix_AuthorizationCode_SecureLocal_MissingStorageName(t *testing.T) {
	keyring.MockInit()

	cfg := config.NewConfiguration().
		WithRuntimeBaseURL(testRuntimeURL).
		WithGrantType(svcOAuth2.GrantTypeAuthorizationCode).
		WithAuthorizationCodeClientID(testClientID).
		WithStorageType(config.StorageTypeSecureLocal)
	// Intentionally no WithStorageName.

	_, err := cfg.TokenSource(context.Background())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "storage name is required")
}

func TestMatrix_AuthorizationCode_None_BypassesCache(t *testing.T) {
	keyring.MockInit()
	seedKeychainToken(t, testRuntimeURL, testClientID, "authorization_code", "should-be-ignored")

	cfg := config.NewConfiguration().
		WithRuntimeBaseURL(testRuntimeURL).
		WithGrantType(svcOAuth2.GrantTypeAuthorizationCode).
		WithAuthorizationCodeClientID(testClientID).
		WithStorageType(config.StorageTypeNone).
		WithStorageName(testStorage).
		// Fail the browser step immediately so the test does not block; reaching this handler
		// proves the cache was bypassed and the interactive flow was attempted.
		WithAuthorizationCodeRedirectURI(config.AuthorizationCodeRedirectURI{})

	cfg.Auth.AuthorizationCode.OnOpenBrowser = func(string) error { return errTestHandler }

	_, err := cfg.TokenSource(context.Background())
	require.Error(t, err)
	assert.True(t,
		strings.Contains(err.Error(), "prompt handler failed"),
		"storage=None must bypass the cache and attempt the authorization_code flow, got: %v", err)
}
