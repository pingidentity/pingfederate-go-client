// Copyright © 2026 Ping Identity Corporation

// Package main demonstrates authorization_code (with PKCE) authentication using the
// PingFederate Go Client SDK. This example shows how to obtain an oauth2.TokenSource
// via a browser-based login and use it to call the PingFederate administrative API.
package main

import (
	"context"
	"crypto/tls"
	"log/slog"
	"net/http"
	"os"

	"github.com/pingidentity/pingfederate-go-client/v1300/config"
	"github.com/pingidentity/pingfederate-go-client/v1300/configurationapi"
	"github.com/pingidentity/pingfederate-go-client/v1300/oauth2"
	xoauth2 "golang.org/x/oauth2"
)

// main demonstrates authorization_code flow authentication.
// This example requires the following environment variables to be set:
//   - PINGFEDERATE_ADMIN_API_URL: administrative API base URL
//     (e.g. https://pingfederate-admin.example.com:9999/pf-admin-api/v1)
//   - PINGFEDERATE_RUNTIME_URL: runtime engine base URL used to derive the OAuth2 endpoints
//     (e.g. https://pingfederate.example.com:9031)
//   - PINGFEDERATE_CLIENT_ID: OAuth2 client ID configured for the authorization_code grant
func main() {
	opts := &slog.HandlerOptions{Level: slog.LevelInfo, AddSource: true}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, opts)))

	adminAPIURL := os.Getenv("PINGFEDERATE_ADMIN_API_URL")
	runtimeURL := os.Getenv("PINGFEDERATE_RUNTIME_URL")
	clientID := os.Getenv("PINGFEDERATE_CLIENT_ID")

	if adminAPIURL == "" || runtimeURL == "" || clientID == "" {
		slog.Error("Missing required environment variables",
			"PINGFEDERATE_ADMIN_API_URL", adminAPIURL != "",
			"PINGFEDERATE_RUNTIME_URL", runtimeURL != "",
			"PINGFEDERATE_CLIENT_ID", clientID != "")
		os.Exit(1)
	}

	// Build the SDK configuration for the authorization_code grant. The runtime base URL is used
	// to derive the OAuth2 endpoints, and tokens are cached in the OS keychain so a subsequent
	// run reuses (and silently refreshes) the token rather than prompting for login again.
	cfg := config.NewConfiguration().
		WithRuntimeBaseURL(runtimeURL).
		WithGrantType(oauth2.GrantTypeAuthorizationCode).
		WithAuthorizationCodeClientID(clientID).
		WithAuthorizationCodeScopes([]string{"openid"}).
		WithStorageName("pingfederate")

	// The loopback redirect URI must match one registered on the OAuth client. The
	// SDK defaults to port 7464; override it via PINGFEDERATE_REDIRECT_PORT (and
	// optionally PINGFEDERATE_REDIRECT_PATH) when the client registers a different one.
	redirectPort := os.Getenv("PINGFEDERATE_REDIRECT_PORT")
	redirectPath := os.Getenv("PINGFEDERATE_REDIRECT_PATH")
	if redirectPort != "" || redirectPath != "" {
		cfg = cfg.WithAuthorizationCodeRedirectURI(config.AuthorizationCodeRedirectURI{
			Port: redirectPort,
			Path: redirectPath,
		})
	}

	// NOTE: TLS verification is disabled here for example brevity only, so the
	// example can reach a PingFederate deployment using a self-signed certificate.
	// Configure proper trust for real deployments. The insecure client must be
	// injected into the context via oauth2.HTTPClient so the token exchange (not
	// just the admin API call below) uses it.
	insecureClient := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}, // #nosec G402 -- example only
	}

	// TokenSource performs the browser-based authorization_code flow unless a valid cached
	// token is available.
	ctx := context.WithValue(context.Background(), xoauth2.HTTPClient, insecureClient)
	tokenSource, err := cfg.TokenSource(ctx)
	if err != nil {
		slog.Error("Authorization code flow failed", "error", err)
		os.Exit(1)
	}

	// Build the generated client and attach the token source via context.
	apiCfg := configurationapi.NewConfiguration()
	apiCfg.Servers = configurationapi.ServerConfigurations{{URL: adminAPIURL}}
	// Reuse the same insecure client for the admin API call.
	apiCfg.HTTPClient = insecureClient
	// The PingFederate admin API requires an X-XSRF-Header on every request for
	// CSRF protection; without it the API responds 400 "xsrf_header_required".
	apiCfg.AddDefaultHeader("X-XSRF-Header", "PingFederate")
	client := configurationapi.NewAPIClient(apiCfg)

	authCtx := context.WithValue(ctx, configurationapi.ContextOAuth2, tokenSource)

	version, resp, err := client.VersionAPI.GetVersion(authCtx).Execute()
	if err != nil {
		slog.Error("Failed to read PingFederate version", "error", err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		slog.Error("Unexpected status reading PingFederate version", "status", resp.Status)
		os.Exit(1)
	}

	slog.Info("Successfully authenticated to PingFederate", "version", version.GetVersion())
}
