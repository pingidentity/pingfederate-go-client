// Copyright © 2026 Ping Identity Corporation

package config_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/pingidentity/pingfederate-go-client/v1300/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A non-http(s) scheme is used for every URL below so that browser.Open rejects it during URL
// validation, before it would otherwise shell out to open a real browser window.
const (
	testAuthURL         = "ftp://example.com/authorize"
	testVerificationURI = "ftp://example.com/device"
	testDeviceUserCode  = "ABCD-1234"
)

func TestDefaultAuthorizationCodeBrowserHandlerTo_WritesToProvidedWriter(t *testing.T) {
	var buf bytes.Buffer

	err := config.DefaultAuthorizationCodeBrowserHandlerTo(&buf)(testAuthURL)
	require.NoError(t, err)

	out := buf.String()
	assert.Contains(t, out, "Opening browser for authorization")
	assert.Contains(t, out, testAuthURL)
	assert.Contains(t, out, "Waiting for authorization callback")
}

func TestDefaultAuthorizationCodeBrowserHandlerTo_DiscardSilencesOutput(t *testing.T) {
	err := config.DefaultAuthorizationCodeBrowserHandlerTo(io.Discard)(testAuthURL)
	require.NoError(t, err)
}

func TestDefaultAuthorizationCodeBrowserHandler_NilOutputWritesToStdout(t *testing.T) {
	out := captureStdout(t, func() {
		err := config.DefaultAuthorizationCodeBrowserHandler(testAuthURL)
		require.NoError(t, err)
	})

	assert.Contains(t, out, "Opening browser for authorization")
	assert.Contains(t, out, testAuthURL)
	assert.Contains(t, out, "Waiting for authorization callback")
}

func TestDefaultDeviceCodePromptHandlerTo_WritesToProvidedWriter(t *testing.T) {
	var buf bytes.Buffer

	err := config.DefaultDeviceCodePromptHandlerTo(&buf)(testVerificationURI, testDeviceUserCode)
	require.NoError(t, err)

	// Assert only on content common to both the browser-available and no-browser branches, since
	// browser.CanOpen() is environment-dependent.
	out := buf.String()
	assert.Contains(t, out, "Device Authorization Required")
	assert.Contains(t, out, testVerificationURI)
	assert.Contains(t, out, testDeviceUserCode)
	assert.Contains(t, out, "Waiting for authorization")
}

func TestDefaultDeviceCodePromptHandlerTo_DiscardSilencesOutput(t *testing.T) {
	err := config.DefaultDeviceCodePromptHandlerTo(io.Discard)(testVerificationURI, testDeviceUserCode)
	require.NoError(t, err)
}

func TestDefaultDeviceCodePromptHandler_NilOutputWritesToStdout(t *testing.T) {
	out := captureStdout(t, func() {
		err := config.DefaultDeviceCodePromptHandler(testVerificationURI, testDeviceUserCode)
		require.NoError(t, err)
	})

	assert.Contains(t, out, "Device Authorization Required")
	assert.Contains(t, out, testVerificationURI)
	assert.Contains(t, out, testDeviceUserCode)
	assert.Contains(t, out, "Waiting for authorization")
}

// captureStdout redirects os.Stdout for the duration of fn and returns everything written to it.
// Not safe to run in parallel with other tests that read or write os.Stdout, so callers must not
// mark their test t.Parallel().
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	r, w, err := os.Pipe()
	require.NoError(t, err)

	original := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = original }()

	outChan := make(chan string, 1)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		outChan <- buf.String()
	}()

	fn()

	require.NoError(t, w.Close())
	return <-outChan
}
