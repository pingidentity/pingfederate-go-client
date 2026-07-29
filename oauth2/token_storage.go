// Copyright © 2026 Ping Identity Corporation

package oauth2

import (
	"golang.org/x/oauth2"
)

// TokenStorage defines the interface for storing and retrieving OAuth2 tokens.
type TokenStorage interface {
	// SaveToken stores an OAuth2 token securely.
	SaveToken(token *oauth2.Token) error

	// LoadToken retrieves a stored OAuth2 token.
	LoadToken() (*oauth2.Token, error)

	// ClearToken removes a stored OAuth2 token.
	ClearToken() error

	// HasToken checks if a token exists without loading it.
	HasToken() (bool, error)
}
