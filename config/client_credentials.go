// Copyright © 2026 Ping Identity Corporation

package config

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// ClientCredentials holds the configuration for the client_credentials grant type.
// Unlike the interactive grant types, this flow authenticates the client directly with
// the authorization server using its client ID and secret and requires no user interaction.
type ClientCredentials struct {
	// ClientCredentialsClientID is the OAuth2 client ID used for the client credentials flow.
	ClientCredentialsClientID *string
	// ClientCredentialsClientSecret is the OAuth2 client secret used for the client credentials flow.
	ClientCredentialsClientSecret *string
	// ClientCredentialsScopes are the OAuth2 scopes requested during the flow.
	ClientCredentialsScopes *[]string
}

// ClientCredentialsTokenSource returns an oauth2.TokenSource using the client_credentials
// grant type. The returned token source automatically obtains and refreshes tokens as needed.
// It requires a client ID and client secret to be configured. The endpoint's TokenURL is used
// to perform the token exchange.
func (c *ClientCredentials) ClientCredentialsTokenSource(ctx context.Context, endpoint oauth2.Endpoint) (oauth2.TokenSource, error) {
	if c.ClientCredentialsClientID == nil || *c.ClientCredentialsClientID == "" {
		return nil, fmt.Errorf("client ID is required for client credentials grant type")
	}

	if c.ClientCredentialsClientSecret == nil || *c.ClientCredentialsClientSecret == "" {
		return nil, fmt.Errorf("client secret is required for client credentials grant type")
	}

	var scopes []string
	if c.ClientCredentialsScopes != nil {
		scopes = *c.ClientCredentialsScopes
	}

	config := &clientcredentials.Config{
		ClientID:     *c.ClientCredentialsClientID,
		ClientSecret: *c.ClientCredentialsClientSecret,
		TokenURL:     endpoint.TokenURL,
		Scopes:       scopes,
		AuthStyle:    oauth2.AuthStyleInHeader,
	}

	return config.TokenSource(ctx), nil
}
