// Copyright © 2026 Ping Identity Corporation

package endpoints_test

import (
	"strings"
	"testing"

	"github.com/pingidentity/pingfederate-go-client/v1300/oauth2/endpoints"
)

func TestPingFederateEndpoint(t *testing.T) {
	tests := []struct {
		name              string
		runtimeBaseURL    string
		expectError       bool
		errorContains     string
		expectedAuthURL   string
		expectedTokenURL  string
		expectedDeviceURL string
	}{
		{
			name:              "Host with port",
			runtimeBaseURL:    "https://pingfederate.example.com:9031",
			expectedAuthURL:   "https://pingfederate.example.com:9031/as/authorization.oauth2",
			expectedTokenURL:  "https://pingfederate.example.com:9031/as/token.oauth2",
			expectedDeviceURL: "https://pingfederate.example.com:9031/as/device_authz.oauth2",
		},
		{
			name:              "Host without port",
			runtimeBaseURL:    "https://auth.example.com",
			expectedAuthURL:   "https://auth.example.com/as/authorization.oauth2",
			expectedTokenURL:  "https://auth.example.com/as/token.oauth2",
			expectedDeviceURL: "https://auth.example.com/as/device_authz.oauth2",
		},
		{
			name:              "Host with trailing slash",
			runtimeBaseURL:    "https://auth.example.com/",
			expectedAuthURL:   "https://auth.example.com/as/authorization.oauth2",
			expectedTokenURL:  "https://auth.example.com/as/token.oauth2",
			expectedDeviceURL: "https://auth.example.com/as/device_authz.oauth2",
		},
		{
			name:              "Host with base path preserved",
			runtimeBaseURL:    "https://gateway.example.com/pf",
			expectedAuthURL:   "https://gateway.example.com/pf/as/authorization.oauth2",
			expectedTokenURL:  "https://gateway.example.com/pf/as/token.oauth2",
			expectedDeviceURL: "https://gateway.example.com/pf/as/device_authz.oauth2",
		},
		{
			name:           "Empty base URL",
			runtimeBaseURL: "",
			expectError:    true,
			errorContains:  "runtime base URL is required",
		},
		{
			name:           "Missing scheme",
			runtimeBaseURL: "pingfederate.example.com:9031",
			expectError:    true,
			errorContains:  "must be an absolute URL",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			endpoint, err := endpoints.PingFederateEndpoint(tc.runtimeBaseURL)

			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
				if tc.errorContains != "" && !strings.Contains(err.Error(), tc.errorContains) {
					t.Errorf("expected error to contain %q, got %q", tc.errorContains, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if endpoint.AuthURL != tc.expectedAuthURL {
				t.Errorf("AuthURL = %q, want %q", endpoint.AuthURL, tc.expectedAuthURL)
			}
			if endpoint.TokenURL != tc.expectedTokenURL {
				t.Errorf("TokenURL = %q, want %q", endpoint.TokenURL, tc.expectedTokenURL)
			}
			if endpoint.DeviceAuthURL != tc.expectedDeviceURL {
				t.Errorf("DeviceAuthURL = %q, want %q", endpoint.DeviceAuthURL, tc.expectedDeviceURL)
			}
		})
	}
}

func TestExplicitEndpoint(t *testing.T) {
	endpoint := endpoints.ExplicitEndpoint(
		"https://auth.example.com/custom/authorize",
		"https://auth.example.com/custom/token",
		"https://auth.example.com/custom/device",
	)

	if endpoint.AuthURL != "https://auth.example.com/custom/authorize" {
		t.Errorf("AuthURL = %q", endpoint.AuthURL)
	}
	if endpoint.TokenURL != "https://auth.example.com/custom/token" {
		t.Errorf("TokenURL = %q", endpoint.TokenURL)
	}
	if endpoint.DeviceAuthURL != "https://auth.example.com/custom/device" {
		t.Errorf("DeviceAuthURL = %q", endpoint.DeviceAuthURL)
	}
}

func TestConstantPaths(t *testing.T) {
	if endpoints.AuthURLPath != "/as/authorization.oauth2" {
		t.Errorf("AuthURLPath = %q", endpoints.AuthURLPath)
	}
	if endpoints.TokenURLPath != "/as/token.oauth2" {
		t.Errorf("TokenURLPath = %q", endpoints.TokenURLPath)
	}
	if endpoints.DeviceAuthURLPath != "/as/device_authz.oauth2" {
		t.Errorf("DeviceAuthURLPath = %q", endpoints.DeviceAuthURLPath)
	}
}
