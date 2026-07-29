// Copyright © 2026 Ping Identity Corporation

// Package main demonstrates device_code (with PKCE) authentication using the PingFederate
// Go Client SDK. This example is suited to CLI tools and devices with limited input
// capabilities: the user completes login on a separate device via a verification URL.
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

// main demonstrates device_code flow authentication.
// This example requires the following environment variables to be set:
//   - PINGFEDERATE_ADMIN_API_URL: administrative API base URL
//     (e.g. https://pingfederate-admin.example.com:9999/pf-admin-api/v1)
//   - PINGFEDERATE_RUNTIME_URL: runtime engine base URL used to derive the OAuth2 endpoints
//     (e.g. https://pingfederate.example.com:9031)
//   - PINGFEDERATE_CLIENT_ID: OAuth2 client ID configured for the device_code grant
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

	// Build the SDK configuration for the device_code grant. The runtime base URL is used to
	// derive the OAuth2 endpoints, and tokens are cached in the OS keychain so a subsequent run
	// reuses (and silently refreshes) the token rather than prompting for login again.
	cfg := config.NewConfiguration().
		WithRuntimeBaseURL(runtimeURL).
		WithGrantType(oauth2.GrantTypeDeviceCode).
		WithDeviceCodeClientID(clientID).
		WithDeviceCodeScopes([]string{"openid"}).
		WithStorageName("pingfederate")

	// NOTE: TLS verification is disabled here for example brevity only, so the
	// example can reach a PingFederate deployment using a self-signed certificate.
	// Configure proper trust for real deployments. The insecure client must be
	// injected into the context via oauth2.HTTPClient so the token exchange (not
	// just the admin API call below) uses it.
	insecureClient := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}, // #nosec G402 -- example only
	}

	// NewAPIClient performs the device_code flow (displaying the verification URL and user code,
	// then polling until login completes, unless a valid cached token is available), then builds
	// the generated admin API client in one call. The returned client's HTTP client already
	// injects the resolved token on every request, so it can be used directly with any context.
	ctx := context.WithValue(context.Background(), xoauth2.HTTPClient, insecureClient)
	client, err := cfg.NewAPIClient(ctx, adminAPIURL, insecureClient)
	if err != nil {
		slog.Error("Device code flow failed", "error", err)
		os.Exit(1)
	}

	version, resp, err := client.VersionAPI.GetVersion(ctx).Execute()
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
