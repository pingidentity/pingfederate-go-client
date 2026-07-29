// Copyright © 2026 Ping Identity Corporation

// Package endpoints provides OAuth2 endpoint construction utilities for PingFederate.
// PingFederate exposes its OAuth2 authorization server on the runtime engine host
// (distinct from the administrative API host), at well-known default paths. This
// package derives the authorization, token, and device-authorization endpoints from
// a single runtime base URL, and also allows each endpoint to be supplied explicitly
// for non-standard deployments.
package endpoints

import (
	"fmt"
	"net/url"

	"golang.org/x/oauth2"
)

const (
	// AuthURLPath is the default PingFederate OAuth2 authorization endpoint path.
	AuthURLPath = "/as/authorization.oauth2"
	// TokenURLPath is the default PingFederate OAuth2 token endpoint path.
	TokenURLPath = "/as/token.oauth2"
	// DeviceAuthURLPath is the default PingFederate OAuth2 device authorization endpoint path.
	DeviceAuthURLPath = "/as/device_authz.oauth2"
)

// PingFederateEndpoint returns an oauth2.Endpoint whose authorization, token, and
// device-authorization URLs are derived from the given PingFederate runtime base URL.
//
// The runtimeBaseURL should be the scheme, host, and optional port of the PingFederate
// runtime engine (for example, "https://pingfederate.example.com:9031"). The standard
// PingFederate OAuth2 paths (/as/authorization.oauth2, /as/token.oauth2, and
// /as/device_authz.oauth2) are appended to it.
//
// An error is returned if runtimeBaseURL is empty or cannot be parsed as an absolute
// URL with a host.
func PingFederateEndpoint(runtimeBaseURL string) (oauth2.Endpoint, error) {
	if runtimeBaseURL == "" {
		return oauth2.Endpoint{}, fmt.Errorf("runtime base URL is required")
	}

	u, err := url.Parse(runtimeBaseURL)
	if err != nil {
		return oauth2.Endpoint{}, fmt.Errorf("invalid runtime base URL %q: %w", runtimeBaseURL, err)
	}

	if u.Scheme == "" || u.Host == "" {
		return oauth2.Endpoint{}, fmt.Errorf("runtime base URL %q must be an absolute URL with a scheme and host", runtimeBaseURL)
	}

	// Preserve any base path configured on the runtime URL by joining the standard
	// OAuth2 paths onto it, rather than replacing the path outright.
	return oauth2.Endpoint{
		AuthURL:       u.JoinPath(AuthURLPath).String(),
		TokenURL:      u.JoinPath(TokenURLPath).String(),
		DeviceAuthURL: u.JoinPath(DeviceAuthURLPath).String(),
	}, nil
}

// ExplicitEndpoint returns an oauth2.Endpoint built from explicitly-supplied URLs.
// It is intended for PingFederate deployments whose OAuth2 endpoints do not follow
// the standard runtime paths. Any empty value is left unset on the returned endpoint.
func ExplicitEndpoint(authURL, tokenURL, deviceAuthURL string) oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:       authURL,
		TokenURL:      tokenURL,
		DeviceAuthURL: deviceAuthURL,
	}
}
