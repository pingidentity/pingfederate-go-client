// Copyright © 2026 Ping Identity Corporation

package config_test

import (
	"context"
	"strings"
	"testing"

	"github.com/pingidentity/pingfederate-go-client/v1300/config"
	"github.com/pingidentity/pingfederate-go-client/v1300/oauth2/endpoints"
)

func TestAuthorizationCodeTokenSource_ValidatesClientID(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *config.AuthorizationCode
	}{
		{
			name: "MissingClientID",
			setup: func() *config.AuthorizationCode {
				return &config.AuthorizationCode{}
			},
		},
		{
			name: "EmptyClientID",
			setup: func() *config.AuthorizationCode {
				clientID := ""
				return &config.AuthorizationCode{AuthorizationCodeClientID: &clientID}
			},
		},
	}

	testEndpoint, err := endpoints.PingFederateEndpoint("https://auth.example.com:9031")
	if err != nil {
		t.Fatalf("failed to build test endpoint: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authCode := tt.setup()
			_, err := authCode.AuthorizationCodeTokenSource(context.Background(), testEndpoint)
			if err == nil {
				t.Fatalf("expected error but got none")
			}
			if !strings.Contains(err.Error(), "client ID is required") {
				t.Errorf("expected client ID error, got %q", err.Error())
			}
		})
	}
}

func TestAuthorizationCodeTokenSource_HandlerError(t *testing.T) {
	clientID := "test-client-id"
	scopes := []string{"openid"}
	authCode := &config.AuthorizationCode{
		AuthorizationCodeClientID: &clientID,
		AuthorizationCodeScopes:   &scopes,
		// A handler that fails immediately, so the flow does not block waiting for a
		// browser callback. This exercises input validation and callback-server startup
		// without requiring a live authorization server.
		OnOpenBrowser: func(string) error {
			return errTestHandler
		},
	}

	testEndpoint, err := endpoints.PingFederateEndpoint("https://auth.example.com:9031")
	if err != nil {
		t.Fatalf("failed to build test endpoint: %v", err)
	}

	_, err = authCode.AuthorizationCodeTokenSource(context.Background(), testEndpoint)
	if err == nil {
		t.Fatalf("expected error but got none")
	}
	if !strings.Contains(err.Error(), "prompt handler failed") {
		t.Errorf("expected prompt handler error, got %q", err.Error())
	}
}

func TestGetDefaultRedirectHelpers(t *testing.T) {
	if got := config.GetDefaultAuthorizationCodeRedirectURIPort(); got != "7464" {
		t.Errorf("default port = %q, want 7464", got)
	}
	if got := config.GetDefaultAuthorizationCodeRedirectURIPath(); got != "/callback" {
		t.Errorf("default path = %q, want /callback", got)
	}
	if got := config.GetDefaultAuthorizationCodeRedirectURI(); got != "http://127.0.0.1:7464/callback" {
		t.Errorf("default redirect URI = %q", got)
	}
}
