// Copyright © 2026 Ping Identity Corporation

package config_test

import (
	"context"
	"strings"
	"testing"

	"github.com/pingidentity/pingfederate-go-client/v1310/config"
	"github.com/pingidentity/pingfederate-go-client/v1310/oauth2/endpoints"
)

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
