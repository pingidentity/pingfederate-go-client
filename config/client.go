// Copyright © 2026 Ping Identity Corporation

package config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pingidentity/pingfederate-go-client/v1300/configurationapi"
)

// xsrfHeaderName is the header the PingFederate administrative API requires on every request
// for CSRF protection; without it the API responds 400 "xsrf_header_required".
const xsrfHeaderName = "X-XSRF-Header"

// NewAPIClient builds a ready-to-use configurationapi.APIClient for the PingFederate
// administrative API, targeting adminAPIURL, and resolves the OAuth2 token source configured on
// c. The returned client's HTTP client already injects the resolved token on every request (via
// Configuration.Client), so callers can use the returned client directly with any context —
// there is no separate auth context to construct or pass around, and no
// configurationapi.Configuration to build and wire by hand. If httpClient is nil, a default
// client is used as the base transport for both the token exchange and the admin API calls.
func (c *Configuration) NewAPIClient(ctx context.Context, adminAPIURL string, httpClient *http.Client) (*configurationapi.APIClient, error) {
	authenticatedClient, err := c.Client(ctx, httpClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create authenticated HTTP client: %w", err)
	}

	apiCfg := configurationapi.NewConfiguration()
	apiCfg.Servers = configurationapi.ServerConfigurations{{URL: adminAPIURL}}
	apiCfg.HTTPClient = authenticatedClient
	apiCfg.AddDefaultHeader(xsrfHeaderName, "PingFederate")

	return configurationapi.NewAPIClient(apiCfg), nil
}
