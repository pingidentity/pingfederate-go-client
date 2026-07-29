// Copyright © 2026 Ping Identity Corporation

package config_test

import (
	"context"
	"testing"

	"github.com/pingidentity/pingfederate-go-client/v1300/config"
)

func TestNewAPIClient(t *testing.T) {
	cfg := config.NewConfiguration().WithAccessToken("static-token")

	client, authCtx, err := cfg.NewAPIClient(context.Background(), "https://pf.example.com:9999/pf-admin-api/v1", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatalf("expected a non-nil API client")
	}
	if authCtx == nil {
		t.Fatalf("expected a non-nil context")
	}
}

func TestNewAPIClientTokenSourceError(t *testing.T) {
	cfg := config.NewConfiguration()

	if _, _, err := cfg.NewAPIClient(context.Background(), "https://pf.example.com:9999/pf-admin-api/v1", nil); err == nil {
		t.Fatalf("expected an error when the grant type is not configured")
	}
}
