// Copyright © 2026 Ping Identity Corporation

package oauth2_test

import (
	"testing"

	"github.com/pingidentity/pingfederate-go-client/v1310/oauth2"
)

func TestIsValidGrantType(t *testing.T) {
	valid := []string{"authorization_code", "client_credentials", "device_code"}
	for _, gt := range valid {
		if !oauth2.IsValidGrantType(gt) {
			t.Errorf("IsValidGrantType(%q) = false, want true", gt)
		}
	}

	invalid := []string{"", "password", "implicit", "refresh_token"}
	for _, gt := range invalid {
		if oauth2.IsValidGrantType(gt) {
			t.Errorf("IsValidGrantType(%q) = true, want false", gt)
		}
	}
}

func TestGrantTypeConstants(t *testing.T) {
	if oauth2.GrantTypeAuthorizationCode != "authorization_code" {
		t.Errorf("GrantTypeAuthorizationCode = %q", oauth2.GrantTypeAuthorizationCode)
	}
	if oauth2.GrantTypeClientCredentials != "client_credentials" {
		t.Errorf("GrantTypeClientCredentials = %q", oauth2.GrantTypeClientCredentials)
	}
	if oauth2.GrantTypeDeviceCode != "device_code" {
		t.Errorf("GrantTypeDeviceCode = %q", oauth2.GrantTypeDeviceCode)
	}
}

func TestAllowedTokenAuthMethods(t *testing.T) {
	for _, gt := range []oauth2.GrantType{
		oauth2.GrantTypeAuthorizationCode,
		oauth2.GrantTypeClientCredentials,
		oauth2.GrantTypeDeviceCode,
	} {
		methods, ok := oauth2.AllowedTokenAuthMethods[gt]
		if !ok {
			t.Errorf("AllowedTokenAuthMethods missing entry for %q", gt)
			continue
		}
		if len(methods) == 0 {
			t.Errorf("AllowedTokenAuthMethods[%q] is empty", gt)
		}
	}
}
