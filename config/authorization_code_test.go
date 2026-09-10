// Copyright © 2026 Ping Identity Corporation

package config_test

import (
	"bytes"
	"context"
	"net"
	"strings"
	"testing"

	"github.com/pingidentity/pingfederate-go-client/v1300/config"
	"github.com/pingidentity/pingfederate-go-client/v1300/oauth2/endpoints"
	"golang.org/x/oauth2"
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
	// Occupy the default redirect port so the callback server cannot start. This exercises
	// input validation, callback-server startup, and the startup-failure path without
	// requiring a live authorization server or launching a browser.
	listener, err := net.Listen("tcp", ":"+config.GetDefaultAuthorizationCodeRedirectURIPort())
	if err != nil {
		t.Fatalf("failed to occupy redirect port: %v", err)
	}
	defer func() { _ = listener.Close() }()

	clientID := "test-client-id"
	scopes := []string{"openid"}
	authCode := &config.AuthorizationCode{
		AuthorizationCodeClientID: &clientID,
		AuthorizationCodeScopes:   &scopes,
	}

	_, err = authCode.AuthorizationCodeTokenSource(context.Background(), oauth2.Endpoint{})
	if err == nil {
		t.Fatalf("expected error but got none")
	}
	if !strings.Contains(err.Error(), "failed to start callback server") {
		t.Errorf("expected callback-server startup error, got %q", err.Error())
	}
}

func TestAuthorizationCodeTokenSource_DefaultHandlerHonorsOutput(t *testing.T) {
	clientID := "test-client-id"
	scopes := []string{"openid"}
	var buf bytes.Buffer

	// No custom handler: the default handler must be selected and write to Output. The AuthURL
	// deliberately pairs a non-http(s) scheme with a loopback host so that browser.Open rejects
	// it during validation without opening a real browser, and the canceled context
	// deterministically aborts the callback wait once the handler has run.
	authCode := &config.AuthorizationCode{
		AuthorizationCodeClientID: &clientID,
		AuthorizationCodeScopes:   &scopes,
		Output:                    &buf,
	}

	// The default redirect port is used so the local callback server can start normally.
	authCode.AuthorizationCodeRedirectURI = config.AuthorizationCodeRedirectURI{}

	testEndpoint := oauth2.Endpoint{
		AuthURL:  "ftp://127.0.0.1/authorize",
		TokenURL: "https://127.0.0.1/as/token.oauth2",
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := authCode.AuthorizationCodeTokenSource(ctx, testEndpoint)
	if err == nil {
		t.Fatalf("expected error but got none")
	}
	if !strings.Contains(err.Error(), "authorization cancelled") {
		t.Errorf("expected cancellation error, got %q", err.Error())
	}

	out := buf.String()
	if !strings.Contains(out, "Opening browser for authorization") {
		t.Errorf("expected default handler progress output, got %q", out)
	}
	if !strings.Contains(out, "ftp://127.0.0.1/authorize") {
		t.Errorf("expected auth URL in default handler output, got %q", out)
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
