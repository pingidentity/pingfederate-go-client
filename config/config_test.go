// Copyright © 2026 Ping Identity Corporation

package config_test

import "errors"

// errTestHandler is a sentinel error returned by test handlers (e.g. OnOpenBrowser) to force an
// early failure so a flow does not block waiting on a browser callback or live authorization server.
var errTestHandler = errors.New("test handler error")
