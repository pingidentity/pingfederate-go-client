// Copyright © 2026 Ping Identity Corporation

package config_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pingidentity/pingfederate-go-client/v1300/config"
	"github.com/pingidentity/pingfederate-go-client/v1300/oauth2/endpoints"
	"golang.org/x/oauth2"
)

// newDeviceAuthServer returns an httptest server that answers both legs of the RFC 8628 flow —
// the device-authorization request and the token request — with fixed responses, so
// DeviceAuthTokenSource can run to completion without a live PingFederate server.
// verificationURI deliberately uses a non-http(s) scheme so that a default handler exercising it
// does not attempt to open a real browser. interval is the polling interval (in seconds) the
// token leg advertises, letting callers trade poll delay against test runtime.
func newDeviceAuthServer(t *testing.T, deviceCode, userCode, verificationURI string, interval int64) *httptest.Server {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := r.ParseForm(); err != nil || r.PostForm.Get("grant_type") == "" {
			// Device authorization request (no grant_type).
			_ = json.NewEncoder(w).Encode(map[string]any{
				"device_code":      deviceCode,
				"user_code":        userCode,
				"verification_uri": verificationURI,
				"expires_in":       600,
				"interval":         interval,
			})
			return
		}

		// Token request (grant_type=urn:ietf:params:oauth:grant-type:device_code).
		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": "device-access-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	t.Cleanup(srv.Close)
	return srv
}

func TestDeviceAuthTokenSource(t *testing.T) {
	tests := []struct {
		name          string
		setup         func() *config.DeviceCode
		errorContains string
	}{
		{
			name: "MissingClientID",
			setup: func() *config.DeviceCode {
				return &config.DeviceCode{}
			},
			errorContains: "client ID is required",
		},
		{
			name: "EmptyClientID",
			setup: func() *config.DeviceCode {
				clientID := ""
				return &config.DeviceCode{DeviceCodeClientID: &clientID}
			},
			errorContains: "client ID is required",
		},
		{
			name: "ValidClientID_WithScopes",
			setup: func() *config.DeviceCode {
				clientID := "test-client-id"
				scopes := []string{"openid", "profile"}
				return &config.DeviceCode{DeviceCodeClientID: &clientID, DeviceCodeScopes: &scopes}
			},
			// Valid inputs, but the request fails against the unreachable endpoint.
			errorContains: "device auth request failed",
		},
		{
			name: "ValidClientID_WithoutScopes",
			setup: func() *config.DeviceCode {
				clientID := "test-client-id"
				return &config.DeviceCode{DeviceCodeClientID: &clientID}
			},
			errorContains: "device auth request failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deviceCode := tt.setup()
			testEndpoint, err := endpoints.PingFederateEndpoint("https://auth.example.com:9031")
			if err != nil {
				t.Fatalf("failed to build test endpoint: %v", err)
			}

			_, err = deviceCode.DeviceAuthTokenSource(context.Background(), testEndpoint)
			if err == nil {
				t.Fatalf("expected error but got none")
			}
			if !strings.Contains(err.Error(), tt.errorContains) {
				t.Errorf("expected error to contain %q, got %q", tt.errorContains, err.Error())
			}
		})
	}
}

func TestDeviceAuthTokenSource_CustomHandlerPrecedenceOverOutput(t *testing.T) {
	srv := newDeviceAuthServer(t, "test-device-code", "test-user-code", "ftp://example.com/device", 5)

	clientID := "test-client-id"
	var buf bytes.Buffer
	var called bool

	deviceCode := &config.DeviceCode{
		DeviceCodeClientID: &clientID,
		Output:             &buf,
		// A custom handler must take priority over Output, and alone control the flow's output.
		OnDisplayPrompt: func(string, string) error {
			called = true
			return errTestHandler
		},
	}

	// DeviceAccessToken is never reached because the handler fails first, so only DeviceAuthURL
	// needs to be stubbed.
	testEndpoint := oauth2.Endpoint{DeviceAuthURL: srv.URL}

	_, err := deviceCode.DeviceAuthTokenSource(context.Background(), testEndpoint)
	if err == nil {
		t.Fatalf("expected error but got none")
	}
	if !strings.Contains(err.Error(), "prompt handler failed") {
		t.Errorf("expected prompt handler error, got %q", err.Error())
	}
	if !called {
		t.Errorf("expected custom OnDisplayPrompt handler to be called")
	}
	if buf.Len() != 0 {
		t.Errorf("expected Output to be unused when a custom handler is set, got %q", buf.String())
	}
}

func TestDeviceAuthTokenSource_DefaultHandlerHonorsOutput(t *testing.T) {
	srv := newDeviceAuthServer(t, "test-device-code", "test-user-code", "ftp://example.com/device", 1)

	clientID := "test-client-id"
	var buf bytes.Buffer

	// No custom handler: the default handler must be selected and write to Output. The flow runs
	// to completion against the stub server, with the non-http(s) verification URI keeping
	// browser.Open from opening a real browser.
	deviceCode := &config.DeviceCode{
		DeviceCodeClientID: &clientID,
		Output:             &buf,
	}
	testEndpoint := oauth2.Endpoint{DeviceAuthURL: srv.URL, TokenURL: srv.URL}

	ts, err := deviceCode.DeviceAuthTokenSource(context.Background(), testEndpoint)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	token, err := ts.Token()
	if err != nil {
		t.Fatalf("unexpected error getting token: %v", err)
	}
	if token.AccessToken != "device-access-token" {
		t.Errorf("expected device-access-token, got %q", token.AccessToken)
	}

	// The default handler is invoked with the stubbed verification URI and user code.
	out := buf.String()
	if !strings.Contains(out, "Device Authorization Required") {
		t.Errorf("expected default handler progress output, got %q", out)
	}
	if !strings.Contains(out, "ftp://example.com/device") {
		t.Errorf("expected verification URI in default handler output, got %q", out)
	}
	if !strings.Contains(out, "test-user-code") {
		t.Errorf("expected user code in default handler output, got %q", out)
	}
}

func TestDeviceAuthTokenSource_CanceledContext(t *testing.T) {
	clientID := "test-client-id"
	scopes := []string{"openid"}
	deviceCode := &config.DeviceCode{DeviceCodeClientID: &clientID, DeviceCodeScopes: &scopes}

	testEndpoint, err := endpoints.PingFederateEndpoint("https://auth.example.com:9031")
	if err != nil {
		t.Fatalf("failed to build test endpoint: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := deviceCode.DeviceAuthTokenSource(ctx, testEndpoint); err == nil {
		t.Error("expected error with canceled context")
	}
}
