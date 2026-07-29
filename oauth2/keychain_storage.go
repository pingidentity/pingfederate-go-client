// Copyright © 2026 Ping Identity Corporation

package oauth2

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/zalando/go-keyring"
	"golang.org/x/oauth2"
)

// KeychainStorage implements TokenStorage using the system keychain.
type KeychainStorage struct {
	serviceName string
	username    string
}

// NewKeychainStorage creates a new KeychainStorage instance.
// Both serviceName and username must be non-empty strings.
func NewKeychainStorage(serviceName, username string) (*KeychainStorage, error) {
	if serviceName == "" {
		return nil, fmt.Errorf("serviceName cannot be empty")
	}
	if username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}
	return &KeychainStorage{
		serviceName: serviceName,
		username:    username,
	}, nil
}

// SaveToken stores the OAuth2 token in the system keychain.
func (k *KeychainStorage) SaveToken(token *oauth2.Token) error {
	if token == nil {
		return errors.New("token cannot be nil")
	}

	tokenJSON, err := json.Marshal(token)
	if err != nil {
		return fmt.Errorf("failed to marshal token: %w", err)
	}

	err = keyring.Set(k.serviceName, k.username, string(tokenJSON))
	if err != nil {
		return fmt.Errorf("failed to save token to keychain: %w", err)
	}

	return nil
}

// LoadToken retrieves the OAuth2 token from the system keychain.
func (k *KeychainStorage) LoadToken() (*oauth2.Token, error) {
	tokenJSON, err := keyring.Get(k.serviceName, k.username)
	if err != nil {
		return nil, fmt.Errorf("failed to load token from keychain: %w", err)
	}

	var token oauth2.Token
	err = json.Unmarshal([]byte(tokenJSON), &token)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal token: %w", err)
	}

	return &token, nil
}

// ClearToken removes the OAuth2 token from the system keychain.
func (k *KeychainStorage) ClearToken() error {
	err := keyring.Delete(k.serviceName, k.username)
	if err != nil && !errors.Is(err, keyring.ErrNotFound) {
		return fmt.Errorf("failed to clear token from keychain: %w", err)
	}
	return nil
}

// ClearAllTokens removes all tokens for the service from the system keychain.
func (k *KeychainStorage) ClearAllTokens() error {
	err := keyring.DeleteAll(k.serviceName)
	if err != nil && !errors.Is(err, keyring.ErrNotFound) {
		return fmt.Errorf("failed to clear all tokens from keychain: %w", err)
	}
	return nil
}

// HasToken checks if a token exists in the system keychain.
func (k *KeychainStorage) HasToken() (bool, error) {
	_, err := keyring.Get(k.serviceName, k.username)
	if err != nil {
		if errors.Is(err, keyring.ErrNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("failed to check token in keychain: %w", err)
	}
	return true, nil
}

// GenerateKeychainAccountName creates a unique account name based on the PingFederate
// runtime base URL, client ID, and grant type. PingFederate has no environment identifier,
// so the runtime base URL disambiguates tokens across different PingFederate deployments.
func GenerateKeychainAccountName(runtimeBaseURL, clientID, grantType string) string {
	if runtimeBaseURL == "" && clientID == "" && grantType == "" {
		// When no inputs are provided, return a stable default.
		return "default-token"
	}

	// Create a hash of runtime base URL + client ID + grant type for uniqueness.
	var b []byte
	b = fmt.Appendf(b, "%s:%s:%s", runtimeBaseURL, clientID, grantType)
	hash := sha256.Sum256(b)
	base := fmt.Sprintf("token-%x", hash[:8]) // Use first 8 bytes of hash for a shorter key.

	return base
}

// GenerateKeychainAccountNameWithSuffix creates a unique account name based on the runtime
// base URL, client ID, grant type, and an optional suffix.
func GenerateKeychainAccountNameWithSuffix(runtimeBaseURL, clientID, grantType, suffix string) string {
	base := GenerateKeychainAccountName(runtimeBaseURL, clientID, grantType)
	if suffix != "" {
		return fmt.Sprintf("%s_%s", base, suffix)
	}
	return base
}
