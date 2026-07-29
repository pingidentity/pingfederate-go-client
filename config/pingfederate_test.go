// Copyright © 2026 Ping Identity Corporation

package config_test

import (
	"context"
	"strings"
	"testing"

	"github.com/pingidentity/pingfederate-go-client/v1300/config"
	svcOAuth2 "github.com/pingidentity/pingfederate-go-client/v1300/oauth2"
	"golang.org/x/oauth2"
)

func TestConfigurationBuilder(t *testing.T) {
	cfg := config.NewConfiguration().
		WithRuntimeBaseURL("https://pf.example.com:9031").
		WithGrantType(svcOAuth2.GrantTypeClientCredentials).
		WithClientCredentialsClientID("client-1").
		WithClientCredentialsClientSecret("secret").
		WithClientCredentialsScopes([]string{"openid"}).
		WithStorageName("pingfederate")

	if cfg.Auth.GrantType == nil || *cfg.Auth.GrantType != svcOAuth2.GrantTypeClientCredentials {
		t.Fatalf("grant type not set correctly")
	}
	if cfg.Auth.ClientCredentials == nil || cfg.Auth.ClientCredentials.ClientCredentialsClientID == nil {
		t.Fatalf("client credentials not set correctly")
	}
	if cfg.Endpoint.RuntimeBaseURL == nil || *cfg.Endpoint.RuntimeBaseURL != "https://pf.example.com:9031" {
		t.Fatalf("runtime base URL not set correctly")
	}
}

func TestConfigurationBearerToken(t *testing.T) {
	cfg := config.NewConfiguration().WithAccessToken("static-token")

	if !cfg.HasBearerToken() {
		t.Fatalf("expected HasBearerToken to be true")
	}

	ts, err := cfg.TokenSource(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	token, err := ts.Token()
	if err != nil {
		t.Fatalf("unexpected error getting token: %v", err)
	}
	if token.AccessToken != "static-token" {
		t.Errorf("expected access token %q, got %q", "static-token", token.AccessToken)
	}
}

func TestConfigurationClientNilHTTPClient(t *testing.T) {
	cfg := config.NewConfiguration().WithAccessToken("static-token")

	client, err := cfg.Client(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatalf("expected a non-nil HTTP client")
	}
}

func TestConfigurationAuthEndpointsExplicit(t *testing.T) {
	explicit := oauth2.Endpoint{
		AuthURL:       "https://custom.example.com/authorize",
		TokenURL:      "https://custom.example.com/token",
		DeviceAuthURL: "https://custom.example.com/device",
	}
	cfg := config.NewConfiguration().WithExplicitEndpoint(explicit)

	ep, err := cfg.AuthEndpoints()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ep.TokenURL != explicit.TokenURL {
		t.Errorf("expected token URL %q, got %q", explicit.TokenURL, ep.TokenURL)
	}
}

func TestConfigurationAuthEndpointsMissing(t *testing.T) {
	cfg := config.NewConfiguration()
	if _, err := cfg.AuthEndpoints(); err == nil {
		t.Fatalf("expected error when no endpoint configured")
	}
}

func TestConfigurationTokenSourceMissingGrantType(t *testing.T) {
	cfg := config.NewConfiguration().WithRuntimeBaseURL("https://pf.example.com:9031")
	_, err := cfg.TokenSource(context.Background())
	if err == nil {
		t.Fatalf("expected error when grant type missing")
	}
	if !strings.Contains(err.Error(), "grant type is required") {
		t.Errorf("expected grant type error, got %q", err.Error())
	}
}

func TestClientCredentialsTokenSourceValidation(t *testing.T) {
	testEndpoint := oauth2.Endpoint{TokenURL: "https://pf.example.com:9031/as/token.oauth2"}

	tests := []struct {
		name  string
		setup func() *config.ClientCredentials
	}{
		{
			name:  "MissingClientID",
			setup: func() *config.ClientCredentials { return &config.ClientCredentials{} },
		},
		{
			name: "MissingClientSecret",
			setup: func() *config.ClientCredentials {
				id := "client-1"
				return &config.ClientCredentials{ClientCredentialsClientID: &id}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := tt.setup()
			if _, err := cc.ClientCredentialsTokenSource(context.Background(), testEndpoint); err == nil {
				t.Fatalf("expected error but got none")
			}
		})
	}
}
