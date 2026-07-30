// Copyright © 2026 Ping Identity Corporation

package oauth2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewKeychainStorageValidation(t *testing.T) {
	_, err := NewKeychainStorage("", "")
	require.Error(t, err, "expected error when creating keychain storage with empty strings")

	_, err = NewKeychainStorage("pingfederate", "")
	require.Error(t, err, "expected error when username is empty")

	storage, err := NewKeychainStorage("pingfederate", "test-user")
	require.NoError(t, err)
	assert.Equal(t, "pingfederate", storage.serviceName)
	assert.Equal(t, "test-user", storage.username)
}

func TestGenerateKeychainAccountName(t *testing.T) {
	// No inputs yields a stable default.
	assert.Equal(t, "default-token", GenerateKeychainAccountName("", "", ""))

	// Same inputs are deterministic.
	a := GenerateKeychainAccountName("https://pf.example.com:9031", "client-1", "device_code")
	b := GenerateKeychainAccountName("https://pf.example.com:9031", "client-1", "device_code")
	assert.Equal(t, a, b, "account name generation must be deterministic")

	// Different inputs yield different names.
	c := GenerateKeychainAccountName("https://pf.example.com:9031", "client-2", "device_code")
	assert.NotEqual(t, a, c, "different client IDs must yield different account names")

	d := GenerateKeychainAccountName("https://other.example.com:9031", "client-1", "device_code")
	assert.NotEqual(t, a, d, "different runtime base URLs must yield different account names")
}

func TestGenerateKeychainAccountNameWithSuffix(t *testing.T) {
	base := GenerateKeychainAccountName("https://pf.example.com:9031", "client-1", "device_code")

	// Empty suffix returns the base name unchanged.
	assert.Equal(t, base, GenerateKeychainAccountNameWithSuffix("https://pf.example.com:9031", "client-1", "device_code", ""))

	// A non-empty suffix is appended.
	withSuffix := GenerateKeychainAccountNameWithSuffix("https://pf.example.com:9031", "client-1", "device_code", "profile1")
	assert.Equal(t, base+"_profile1", withSuffix)
}
