// Copyright © 2026 Ping Identity Corporation

// Package main demonstrates client_credentials authentication using the PingFederate
// Go Client SDK. This flow is suited to server-to-server (machine-to-machine) integrations
// and automation that run without user interaction: the client authenticates directly with
// its client ID and secret and the SDK refreshes the token automatically as needed.
package main

import (
	"context"
	"crypto/tls"
	"log/slog"
	"net/http"
	"os"

	"github.com/pingidentity/pingfederate-go-client/v1300/config"
	"github.com/pingidentity/pingfederate-go-client/v1300/oauth2"
	xoauth2 "golang.org/x/oauth2"
)

// main demonstrates client_credentials flow authentication.
// This example requires the following environment variables to be set:
//   - PINGFEDERATE_ADMIN_API_URL: administrative API base URL
//     (e.g. https://pingfederate-admin.example.com:9999/pf-admin-api/v1)
//   - PINGFEDERATE_RUNTIME_URL: runtime engine base URL used to derive the OAuth2 endpoints
//     (e.g. https://pingfederate.example.com:9031)
//   - PINGFEDERATE_CLIENT_ID: OAuth2 client ID configured for the client_credentials grant
//   - PINGFEDERATE_CLIENT_SECRET: OAuth2 client secret for the client
func main() {
	opts := &slog.HandlerOptions{Level: slog.LevelInfo, AddSource: true}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, opts)))

	adminAPIURL := os.Getenv("PINGFEDERATE_ADMIN_API_URL")
	runtimeURL := os.Getenv("PINGFEDERATE_RUNTIME_URL")
	clientID := os.Getenv("PINGFEDERATE_CLIENT_ID")
	clientSecret := os.Getenv("PINGFEDERATE_CLIENT_SECRET")

	if adminAPIURL == "" || runtimeURL == "" || clientID == "" || clientSecret == "" {
		slog.Error("Missing required environment variables",
			"PINGFEDERATE_ADMIN_API_URL", adminAPIURL != "",
			"PINGFEDERATE_RUNTIME_URL", runtimeURL != "",
			"PINGFEDERATE_CLIENT_ID", clientID != "",
			"PINGFEDERATE_CLIENT_SECRET", clientSecret != "")
		os.Exit(1)
	}

	// Build the SDK configuration for the client_credentials grant. The returned token source
	// obtains and refreshes tokens automatically; no user interaction or token caching is needed.
	cfg := config.NewConfiguration().
		WithRuntimeBaseURL(runtimeURL).
		WithGrantType(oauth2.GrantTypeClientCredentials).
		WithClientCredentialsClientID(clientID).
		WithClientCredentialsClientSecret(clientSecret).
		WithClientCredentialsScopes([]string{"openid"})

	// NOTE: TLS verification is disabled here for example brevity only, so the
	// example can reach a PingFederate deployment using a self-signed certificate.
	// Configure proper trust for real deployments. The insecure client must be
	// injected into the context via oauth2.HTTPClient so the token exchange (not
	// just the admin API call below) uses it.
	insecureClient := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}, // #nosec G402 -- example only
	}
	ctx := context.WithValue(context.Background(), xoauth2.HTTPClient, insecureClient)

	// NewAPIClient resolves the token source and builds the generated admin API client in one
	// call, so there is no separate configurationapi.Configuration to build and wire by hand.
	client, authCtx, err := cfg.NewAPIClient(ctx, adminAPIURL, insecureClient)
	if err != nil {
		slog.Error("Client credentials flow failed", "error", err)
		os.Exit(1)
	}

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
