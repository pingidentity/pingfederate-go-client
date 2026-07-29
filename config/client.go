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
// c. If httpClient is nil, a default client is used for both the token exchange and the admin API
// calls. It returns the client along with a context carrying the token source
// (configurationapi.ContextOAuth2) so callers can pass the returned context directly to generated
// API methods, without separately constructing and wiring a configurationapi.Configuration.
func (c *Configuration) NewAPIClient(ctx context.Context, adminAPIURL string, httpClient *http.Client) (*configurationapi.APIClient, context.Context, error) {
	ts, err := c.TokenSource(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create token source: %w", err)
	}

	apiCfg := configurationapi.NewConfiguration()
	apiCfg.Servers = configurationapi.ServerConfigurations{{URL: adminAPIURL}}
	apiCfg.HTTPClient = httpClient
	apiCfg.AddDefaultHeader(xsrfHeaderName, "PingFederate")

	client := configurationapi.NewAPIClient(apiCfg)
	authCtx := context.WithValue(ctx, configurationapi.ContextOAuth2, ts)

	return client, authCtx, nil
}
